package app

import (
	"github.com/charmbracelet/bubbles/progress"
	"github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type SetupModel struct {
	TextInput textinput.Model
	Progress  progress.Model
	Loading   bool
	ErrorText string
	Width     int
	Height    int
	Viewport  viewport.Model
}

func NewSetupModel() SetupModel {
	ti := textinput.New()
	ti.Placeholder = "spaceship token"
	ti.Focus()
	ti.CharLimit = 64
	ti.Width = 40

	return SetupModel{
		TextInput: ti,
		Progress:  progress.New(progress.WithScaledGradient("#FF5F87", "#FFAFD7")),
		Viewport:  viewport.New(0, 0),
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
			if m.TextInput.Value() == "" {
				m.ErrorText = "Token cannot be empty!"
			} else {
				m.Loading = true
			}
		}

	case tea.WindowSizeMsg:
		m.Width = msg.Width
		m.Height = msg.Height
		m.Viewport.Width = msg.Width
		m.Viewport.Height = msg.Height
		m.Viewport.SetContent(m.viewportContent())
	}

	m.TextInput, cmd = m.TextInput.Update(msg)
	return m, cmd
}

func (m SetupModel) View() string {
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

	// This is just wakatime test api thingy

	var subtitleStyle = lipgloss.NewStyle().
		Italic(true).
		Foreground(lipgloss.Color("6"))

	return viewportStyle.Render(lipgloss.JoinVertical(lipgloss.Top,
		titleStyle.Render("Welcome to Spaceship TUI!"),
		subtitleStyle.Render("To access your account & domains we need to have a spaceship user token. You can generate one here : https://www.spaceship.com/application/api-manager/ !"),
		"\n",
		lipgloss.NewStyle().MarginBottom(1).Render(m.TextInput.View()),
		m.Progress.View(),
		lipgloss.NewStyle().Foreground(lipgloss.Color("red")).Render(m.ErrorText),
		lipgloss.NewStyle().Render("Press 'enter' to submit."),
	))
}
