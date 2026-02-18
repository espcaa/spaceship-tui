package dashboard

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/espcaa/spaceship-go"
	"github.com/espcaa/spaceship-tui/shared"
)

type credentialsVerifiedMsg struct {
	Client *spaceship.Client
}

type DashboardModel struct {
	ApiKey        string
	ApiSecret     string
	ActiveScreen  shared.Screen
	Client        *spaceship.Client
	CurrentScreen tea.Model
	Domains       []spaceship.DomainInfo
	width, height int
}

func NewDashboardModel(apiKey, apiSecret string) *DashboardModel {
	return &DashboardModel{
		ApiKey:       apiKey,
		ApiSecret:    apiSecret,
		ActiveScreen: shared.ScreenDomainList,
	}
}

func (m *DashboardModel) Init() tea.Cmd {
	return func() tea.Msg {
		client := spaceship.NewClient(m.ApiKey, m.ApiSecret)
		if err := client.VerifyCredentials(); err != nil {
			return shared.DashboardErrorMsg{Error: "Invalid API credentials"}
		}
		return credentialsVerifiedMsg{Client: client}
	}
}
