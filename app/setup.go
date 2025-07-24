package app

import (
	"net/http"

	"github.com/charmbracelet/bubbles/progress"
	"github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/google/go-querystring/query"
)

type TestQuery struct {
	Take int `url:"take"` // = 1 just for auth test
	Skip int `url:"skip"` // = 0
}

func TestSpaceshipAPI(key, secret string) (bool, string) {
	// Check if either/both or them are not 64 characters long
	if len(key) != 64 || len(secret) != 64 {
		return false, "Key and secret must be 64 characters long!"
	}
	// Check it by calling spaceship api

	params := TestQuery{
		Take: 1,
		Skip: 0,
	}
	v, _ := query.Values(params)
	url := "https://api.spaceship.com/domains?" + v.Encode()
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}
	req.Header.Set("X-API-Key", key)
	req.Header.Set("X-API-Secret", secret)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return false, "the connection failed... are you online?"
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return false, "uh captain we got an error, please check your key and secret!"
	}
	// If we reach here, the key and secret are valid
	return true, "doing some strange work behind the scenes..."
}

type SetupModel struct {
	TextInputKey    textinput.Model
	TextInputSecret textinput.Model
	Progress        progress.Model
	Loading         bool
	ErrorText       string
	Width           int
	Height          int
	Viewport        viewport.Model
}

func NewSetupModel() SetupModel {
	tisecret := textinput.New()
	tisecret.Placeholder = "spaceship secret"
	ti := textinput.New()
	ti.Focus()
	ti.Placeholder = "spaceship api key"
	ti.CharLimit = 64
	ti.Width = 40
	tisecret.Width = 40
	tisecret.CharLimit = 64

	return SetupModel{
		TextInputSecret: tisecret,
		TextInputKey:    ti,
		Progress:        progress.New(progress.WithScaledGradient("#FF5F87", "#FFAFD7")),
		Viewport:        viewport.New(0, 0),
	}
}

func (m SetupModel) Init() tea.Cmd {
	return textinput.Blink
}

func (m SetupModel) Update(msg tea.Msg) (SetupModel, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "enter":
			if m.TextInputSecret.Value() == "" {
				m.ErrorText = "Token cannot be empty!"
			} else {
				m.Loading = true
			}
		case "tab", "shift+tab":
			if m.TextInputKey.Focused() {
				m.TextInputKey.Blur()
				m.TextInputSecret.Focus()
			} else {
				m.TextInputSecret.Blur()
				m.TextInputKey.Focus()
			}
		}

	case tea.WindowSizeMsg:
		m.Width = msg.Width
		m.Height = msg.Height
		m.Viewport.Width = msg.Width
		m.Viewport.Height = msg.Height
	}

	m.TextInputSecret, cmd = m.TextInputSecret.Update(msg)
	return m, cmd
}

func (m SetupModel) View() string {
	m.Viewport.SetContent(m.viewportContent())
	return m.Viewport.View()
}

func (m SetupModel) viewportContent() string {

	var viewport_width = 60
	if m.Width < 60 {
		viewport_width = m.Width - 8
	}

	var viewportStyle = lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		Padding(1, 2).
		Margin(2).
		Width(viewport_width)

	var titleStyle = lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("2"))

	var subtitleStyle = lipgloss.NewStyle().
		Italic(true).
		Foreground(lipgloss.Color("6"))

	var textStyle = lipgloss.NewStyle().
		Foreground(lipgloss.Color("2"))

	var lastText = "Press 'enter' to submit."
	if m.ErrorText != "" {
		lastText = "Error: " + m.ErrorText
	}

	var Textcolor = lipgloss.Color("15")
	if m.ErrorText != "" {
		Textcolor = lipgloss.Color("1")
	}

	return viewportStyle.Render(lipgloss.JoinVertical(lipgloss.Top,
		titleStyle.Render("Welcome to Spaceship TUI!"),
		subtitleStyle.Render("To access your account & domains we need to have a spaceship api key AND secret. You can generate one here : https://www.spaceship.com/application/api-manager/ !"),
		"\n",
		textStyle.Render("First, your api key:"),
		lipgloss.NewStyle().MarginBottom(1).Render(m.TextInputKey.View()),
		textStyle.Render("Then, your api secret:"),
		lipgloss.NewStyle().MarginBottom(1).Render(m.TextInputSecret.View()),
		m.Progress.View(),
		lipgloss.NewStyle().Foreground(lipgloss.Color("red")).Render(m.ErrorText),
		lipgloss.NewStyle().Foreground(Textcolor).Render(lastText),
	))
}
