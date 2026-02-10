package dashboard

import (
	tea "github.com/charmbracelet/bubbletea"
	domainlist "github.com/espcaa/spaceship-tui/screens/dashboard/screens/domain_list"
)

func (m *DashboardModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case SwitchScreenMsg:
		m.ActiveScreen = msg.Screen
		switch msg.Screen {
		case ScreenDomainList:
			m.CurrentScreen = domainlist.NewDomainListModel(m.Client)
		}
		return m, m.CurrentScreen.Init()

	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit
		}
	}

	var cmd tea.Cmd
	m.CurrentScreen, cmd = m.CurrentScreen.Update(msg)

	return m, cmd
}
