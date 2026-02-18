package dashboard

import (
	tea "github.com/charmbracelet/bubbletea"
	domainlist "github.com/espcaa/spaceship-tui/screens/dashboard/screens/domain_list"
	individualdomain "github.com/espcaa/spaceship-tui/screens/dashboard/screens/individual_domain"
	"github.com/espcaa/spaceship-tui/shared"
)

func (m *DashboardModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height

	case credentialsVerifiedMsg:
		m.Client = msg.Client
		m.CurrentScreen = domainlist.NewDomainListModel(m.Client, m.Domains)
		return m, tea.Batch(m.CurrentScreen.Init(), func() tea.Msg {
			return tea.WindowSizeMsg{Width: m.width, Height: m.height}
		})

	case domainlist.DomainListSuccessMsg:
		m.Domains = msg.Domains

	case shared.DomainSelectedMsg:
		m.ActiveScreen = shared.ScreenIndividualDomain
		m.CurrentScreen = individualdomain.NewIndividualDomainModel(msg.Domain, m.Client)
		return m, m.CurrentScreen.Init()

	case shared.SwitchScreenMsg:
		m.ActiveScreen = msg.Screen
		switch msg.Screen {
		case shared.ScreenDomainList:
			m.CurrentScreen = domainlist.NewDomainListModel(m.Client, m.Domains)
		}
		return m, tea.Batch(m.CurrentScreen.Init(), func() tea.Msg {
			return tea.WindowSizeMsg{Width: m.width, Height: m.height}
		})

	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit
		}
	}

	if m.CurrentScreen != nil {
		var cmd tea.Cmd
		m.CurrentScreen, cmd = m.CurrentScreen.Update(msg)
		return m, cmd
	}

	return m, nil
}
