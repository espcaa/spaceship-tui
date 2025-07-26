package app

import (
	"log"
	"net/http"
	"time"

	"github.com/charmbracelet/bubbles/progress"
	"github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/google/go-querystring/query"
)

type tickMsg time.Time

func tickCmd() tea.Cmd {
	return tea.Tick(time.Second, func(t time.Time) tea.Msg {
		return tickMsg(t)
	})
}

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
	if len(key) < 10 || len(secret) < 63 {
		return ErrorMsg{
			Text: "uh i think your api keys are invalid or switched, please check them again!",
		}
	}

	params := TestQuery{
		Take: 1,
		Skip: 0,
	}
	v, _ := query.Values(params)
	url := "https://spaceship.dev/api/v1/domains?" + v.Encode()
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
		log.Printf("Error during request: %+v\n", err)
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
		Text: "API test succeeded!",
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
	done            bool
	animateProgress bool
	WaitingError    string
}

func NewSetupModel() SetupModel {
	tisecret := textinput.New()
	tisecret.Placeholder = "spaceship secret"
	tti := textinput.New()
	tti.Focus()
	tti.Placeholder = "spaceship api key"
	tti.CharLimit = 64
	tti.Width = 40
	tisecret.Width = 40
	tisecret.CharLimit = 64

	return SetupModel{
		TextInputSecret: tisecret,
		TextInputKey:    tti,
		Progress:        progress.New(progress.WithScaledGradient("#FF0000", "#00FF00"), progress.WithDefaultGradient()),
		Loading:         false,
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
	var cmds []tea.Cmd

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

				m.Progress = progress.New(progress.WithScaledGradient("#FF0000", "#00FF00"), progress.WithDefaultGradient())
				cmds = append(cmds, testSpaceshipCmd(key, secret))
			}
			return m, tea.Batch(cmds...)

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
			return m, nil

		}

	case tea.WindowSizeMsg:
		m.Width = msg.Width
		m.Height = msg.Height
		m.Viewport.Width = msg.Width
		m.Viewport.Height = msg.Height

	case SuccessMsg:
		m.ErrorText = ""
		m.WaitingError = ""
		m.Progress.SetPercent(0.0)
		m.animateProgress = true
		cmds = append(cmds, tickCmd())

	case ErrorMsg:
		m.WaitingError = msg.Text
		m.Progress.SetPercent(0.0)
		m.animateProgress = true
		cmds = append(cmds, tickCmd())

	case tickMsg:
		if m.animateProgress {
			if m.Progress.Percent() < 1.0 {
				cmd := m.Progress.IncrPercent(1.0)
				cmds = append(cmds, tickCmd(), cmd)
			} else {
				if m.WaitingError != "" {
					m.ErrorText = m.WaitingError
				} else {
					m.ErrorText = "Spaceship API test succeeded! You can now use the app."
					m.done = true
				}
				m.animateProgress = false
				m.Loading = false
			}
		}
	}

	if !m.Loading {
		var cmd1, cmd2 tea.Cmd
		m.TextInputKey, cmd1 = m.TextInputKey.Update(msg)
		m.TextInputSecret, cmd2 = m.TextInputSecret.Update(msg)
		cmds = append(cmds, cmd1, cmd2)
	} else {
		var updated tea.Model
		var progressCmd tea.Cmd
		updated, progressCmd = m.Progress.Update(msg)
		m.Progress = updated.(progress.Model)
		cmds = append(cmds, progressCmd)
	}

	return m, tea.Batch(cmds...)
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
		Textcolor = lipgloss.Color("1")
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
		progressView := m.Progress.View()
		if progressView != "" {
			content = append(content, lipgloss.NewStyle().MarginBottom(1).Render(progressView))
		}
	}

	content = append(content, lipgloss.NewStyle().Foreground(Textcolor).Render(lastText))

	return viewportStyle.Render(lipgloss.JoinVertical(lipgloss.Top, content...))
}
