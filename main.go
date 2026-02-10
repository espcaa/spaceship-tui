package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/espcaa/spaceship-tui/screens/dashboard"
	"github.com/espcaa/spaceship-tui/screens/login"
)

type AppState int

const (
	StateLoggedOut AppState = iota
	StateLoggedIn
)

type initialModel struct {
	apiKey    string
	apiSecret string
	login     login.LoginModel
	dashboard *dashboard.DashboardModel
	state     AppState
}

type credentialsLoadedMsg struct {
	apiKey    string
	apiSecret string
}

func (m initialModel) Init() tea.Cmd {

	return tea.Batch(
		m.login.Init(),
		func() tea.Msg {
			creds, err := LoadCredentials()
			if err != nil {
				return nil
			}
			return credentialsLoadedMsg{apiKey: creds.APIKey, apiSecret: creds.APISecret}
		},
	)
}

func (m initialModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit
		}

	case credentialsLoadedMsg:
		m.apiKey = msg.apiKey
		m.apiSecret = msg.apiSecret
		m.dashboard = dashboard.NewDashboardModel(m.apiKey, m.apiSecret)
		m.state = StateLoggedIn
		return m, m.dashboard.Init()

	case login.LoginSuccessMsg:
		m.apiKey = msg.ApiKey
		m.apiSecret = msg.ApiSecret
		m.dashboard = dashboard.NewDashboardModel(m.apiKey, m.apiSecret)
		m.state = StateLoggedIn
		err := SaveCredentials(Credentials{
			APIKey:    msg.ApiKey,
			APISecret: msg.ApiSecret,
		})
		if err != nil {
			return m, func() tea.Msg {
				return login.LoginErrorMsg{Error: "Failed to save credentials: " + err.Error()}
			}
		}
		return m, m.dashboard.Init()
	}

	if m.state == StateLoggedOut {
		updatedLogin, cmd := m.login.Update(msg)
		m.login = updatedLogin.(login.LoginModel)
		return m, cmd
	}

	updatedDashboard, cmd := m.dashboard.Update(msg)
	m.dashboard = updatedDashboard.(*dashboard.DashboardModel)
	return m, cmd
}

func (m initialModel) View() string {
	if m.state == StateLoggedIn {
		return m.dashboard.View()
	}
	return m.login.View()
}

func main() {
	p := tea.NewProgram(initialModel{
		login: login.NewLoginModel(),
		state: StateLoggedOut,
	}, tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
