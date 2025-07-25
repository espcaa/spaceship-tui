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

type ErrorMsg struct {
	Text string
}

type SuccessMsg struct {
	Text string
}

func TestSpaceshipAPI(key, secret string) tea.Msg {
	if len(key) != 64 || len(secret) != 64 {
		return ErrorMsg{
			Text: "Invalid key or secret length. Both should be 64 characters long.",
		}
	}

	params := TestQuery{
		Take: 1,
		Skip: 0,
	}
	v, _ := query.Values(params)
	url := "https://api.spaceship.com/domains?" + v.Encode()
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return ErrorMsg{
			Text: "failed to build request",
		}
	}
	req.Header.Set("X-API-Key", key)
	req.Header.Set("X-API-Secret", secret)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return ErrorMsg{
			Text: "the connection failed... are you online?",
		}
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return ErrorMsg{
			Text: "uh captain we got an error, please check your key and secret!",
		}
	}

	return SuccessMsg{
		Text: "doing some strange work behind the scenes...",
	}

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

func testSpaceshipCmd(key, secret string) tea.Cmd {
	return func() tea.Msg {
		return TestSpaceshipAPI(key, secret)
	}
}

func (m SetupModel) Update(msg tea.Msg) (SetupModel, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "enter":
			if !m.Loading {
				key := m.TextInputKey.Value()
				secret := m.TextInputSecret.Value()

				if key == "" || secret == "" {
					m.ErrorText = "i think you forgot one of your api key somewhere..."
					return m, nil
				}

				m.Loading = true
				m.ErrorText = ""
				m.Progress = progress.New(progress.WithScaledGradient("#FF5F87", "#FFAFD7"))
				m.Progress.SetPercent(0.0)

				return m, tea.Batch(
					testSpaceshipCmd(key, secret),
				)

			}
		case "tab", "shift+tab":
			if !m.Loading {
				if m.TextInputKey.Focused() {
					m.TextInputKey.Blur()
					m.TextInputSecret.Focus()
				} else {
					m.TextInputSecret.Blur()
					m.TextInputKey.Focus()
				}
			}
		}

	case tea.WindowSizeMsg:
		m.Width = msg.Width
		m.Height = msg.Height
		m.Viewport.Width = msg.Width
		m.Viewport.Height = msg.Height

	case SuccessMsg:
		m.Loading = false
		m.ErrorText = ""

	case ErrorMsg:
		m.Loading = false
		m.ErrorText = msg.Text
	}

	var cmd1, cmd2 tea.Cmd
	m.TextInputKey, cmd1 = m.TextInputKey.Update(msg)
	m.TextInputSecret, cmd2 = m.TextInputSecret.Update(msg)
	return m, tea.Batch(
		cmd1,
		cmd2,
	)

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
		lastText = m.ErrorText
	}

	var Textcolor = lipgloss.Color("15")
	if m.ErrorText != "" {
		Textcolor = lipgloss.Color("red")
	}

	var content []string

	content = append(content,
		titleStyle.Render("Welcome to Spaceship TUI!"),
		subtitleStyle.Render("To access your account & domains we need to have a spaceship api key AND secret. You can generate one here : https://www.spaceship.com/application/api-manager/ !"),
		"\n",
		textStyle.Render("First, your api key:"),
		lipgloss.NewStyle().MarginBottom(1).Render(m.TextInputKey.View()),
		textStyle.Render("Then, your api secret:"),
		lipgloss.NewStyle().MarginBottom(1).Render(m.TextInputSecret.View()),
	)

	if m.Loading {
		progress := m.Progress.View()
		if progress != "" {
			content = append(content, lipgloss.NewStyle().MarginBottom(1).Render(progress))
		}
	}

	content = append(content, lipgloss.NewStyle().Foreground(Textcolor).Render(lastText))

	return viewportStyle.Render(lipgloss.JoinVertical(lipgloss.Top, content...))
}
