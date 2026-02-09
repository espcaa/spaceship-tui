package main

import (
	"encoding/json"
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
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
	state     AppState
}

func loadCredentials() (string, string, error) {

	// load from json file
	homedir, err := os.UserHomeDir()
	path := homedir + "/.config/spaceship-tui/secrets.json"
	file, err := os.ReadFile(path)
	if err != nil {
		return "", "", err
	}

	var creds struct {
		APIKey    string `json:"api_key"`
		APISecret string `json:"api_secret"`
	}

	err = json.Unmarshal(file, &creds)
	if err != nil {
		return "", "", err
	}

	return creds.APIKey, creds.APISecret, nil
}

type credentialsLoadedMsg struct {
	apiKey    string
	apiSecret string
}

func (m initialModel) Init() tea.Cmd {

	return tea.Batch(
		m.login.Init(),
		func() tea.Msg {
			apiKey, apiSecret, err := loadCredentials()
			if err != nil {
				return nil
			}
			return credentialsLoadedMsg{apiKey: apiKey, apiSecret: apiSecret}
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
		m.state = StateLoggedIn
		return m, nil

	case login.LoginSuccessMsg:
		m.apiKey = msg.ApiKey
		m.apiSecret = msg.ApiSecret
		m.state = StateLoggedIn
		SaveCredentials(msg.ApiKey, msg.ApiSecret)
		return m, nil
	}

	if m.state == StateLoggedOut {
		updatedLogin, cmd := m.login.Update(msg)
		m.login = updatedLogin.(login.LoginModel)
		return m, cmd
	}

	return m, nil
}

func (m initialModel) View() string {
	if m.state == StateLoggedIn {
		return "Logged in! Press q to quit.\n"
	} else {
		return m.login.View()
	}
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

func SaveCredentials(apiKey, apiSecret string) error {
	creds := struct {
		APIKey    string `json:"api_key"`
		APISecret string `json:"api_secret"`
	}{
		APIKey:    apiKey,
		APISecret: apiSecret,
	}

	data, err := json.MarshalIndent(creds, "", "  ")
	if err != nil {
		return err
	}

	homedir, err := os.UserHomeDir()
	if err != nil {
		return err
	}
	path := homedir + "/.config/spaceship-tui/secrets.json"

	err = os.MkdirAll(homedir+"/.config/spaceship-tui", 0700)
	if err != nil {
		return err
	}

	err = os.WriteFile(path, data, 0600)
	if err != nil {
		return err
	}

	return nil
}
