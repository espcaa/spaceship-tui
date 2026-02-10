package dashboard

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/espcaa/spaceship-go"
	domainlist "github.com/espcaa/spaceship-tui/screens/dashboard/screens/domain_list"
	"github.com/espcaa/spaceship-tui/shared"
)

type Screen int

const (
	ScreenDomainList Screen = iota
)

type SwitchScreenMsg struct {
	Screen Screen
}

type DashboardModel struct {
	ApiKey        string
	ApiSecret     string
	ActiveScreen  Screen
	Client        *spaceship.Client
	CurrentScreen tea.Model
}

func NewDashboardModel(apiKey, apiSecret string) *DashboardModel {
	return &DashboardModel{
		ApiKey:       apiKey,
		ApiSecret:    apiSecret,
		ActiveScreen: ScreenDomainList,
	}
}

func (m *DashboardModel) Init() tea.Cmd {
	// initialize the client
	client := spaceship.NewClient(m.ApiKey, m.ApiSecret)
	if err := client.VerifyCredentials(); err != nil {
		return func() tea.Msg {
			return shared.DashboardErrorMsg{Error: "Invalid API credentials"}
		}
	}
	m.Client = client
	m.CurrentScreen = domainlist.NewDomainListModel(m.Client)

	return m.CurrentScreen.Init()
}
