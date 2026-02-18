package main

import (
	"database/sql"
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/espcaa/spaceship-tui/screens/dashboard"
	"github.com/espcaa/spaceship-tui/screens/login"
	_ "modernc.org/sqlite"
)

type AppState int

const (
	StateLoggedOut AppState = iota
	StateLoggedIn
	StateLoading
)

type initialModel struct {
	apiKey        string
	apiSecret     string
	db            *sql.DB
	login         login.LoginModel
	dashboard     *dashboard.DashboardModel
	state         AppState
	width, height int
}

type credentialsLoadedMsg struct {
	apiKey    string
	apiSecret string
}

type credentialsLoadErrorMsg struct {
	Error string
}

func (m initialModel) Init() tea.Cmd {

	return tea.Batch(
		m.login.Init(),
		func() tea.Msg {
			creds, err := LoadCredentials()
			if err != nil {
				return credentialsLoadErrorMsg{Error: err.Error()}
			}
			return credentialsLoadedMsg{apiKey: creds.APIKey, apiSecret: creds.APISecret}
		},
	)
}

func (m initialModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height

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
		return m, tea.Batch(m.dashboard.Init(), func() tea.Msg {
			return tea.WindowSizeMsg{Width: m.width, Height: m.height}
		})

	case credentialsLoadErrorMsg:
		m.state = StateLoggedOut
		return m, nil

	case login.LoginSuccessMsg:
		m.apiKey = msg.ApiKey
		m.apiSecret = msg.ApiSecret
		m.dashboard = dashboard.NewDashboardModel(m.apiKey, m.apiSecret, m.db)
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
		return m, tea.Batch(m.dashboard.Init(), func() tea.Msg {
			return tea.WindowSizeMsg{Width: m.width, Height: m.height}
		})
	}

	if m.state == StateLoggedOut {
		updatedLogin, cmd := m.login.Update(msg)
		m.login = updatedLogin.(login.LoginModel)
		return m, cmd
	}

	if m.state == StateLoggedIn && m.dashboard != nil {
		updatedDashboard, cmd := m.dashboard.Update(msg)
		m.dashboard = updatedDashboard.(*dashboard.DashboardModel)
		return m, cmd
	}

	return m, nil
}

func (m initialModel) View() string {
	if m.state == StateLoggedIn {
		return m.dashboard.View()
	}

	if m.state == StateLoggedOut {
		return m.login.View()
	}
	return ""

}

func main() {

	// init the db
	homedir, err := os.UserHomeDir()
	if err != nil {
		panic("we didn't manage to determine a home folder, this is really strange and shouldn't happen....")
	}
	db, err := sql.Open("sqlite", homedir+"/.config/spaceship-tui/cache.db?mode=rwc")
	if err != nil {
		panic("failed to open database: " + err.Error())
	}
	defer db.Close()
	p := tea.NewProgram(initialModel{
		login: login.NewLoginModel(),
		state: StateLoading,
		db:    db,
	}, tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
