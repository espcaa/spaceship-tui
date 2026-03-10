package dashboard

import "charm.land/lipgloss/v2"

func (m *DashboardModel) View() string {
	if m.CurrentScreen == nil {
		return lipgloss.NewStyle().Margin(2).Render("Verifying credentials...")
	}
	return m.CurrentScreen.View()
}
