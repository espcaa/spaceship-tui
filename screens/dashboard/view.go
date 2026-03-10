package dashboard

import (
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
)

func (m *DashboardModel) View() tea.View {
	if m.CurrentScreen == nil {
		return tea.NewView(lipgloss.NewStyle().Margin(2).Render("Verifying credentials..."))
	}
	return m.CurrentScreen.View()
}
