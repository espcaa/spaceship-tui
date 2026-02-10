package dashboard

func (m *DashboardModel) View() string {
	if m.CurrentScreen == nil {
		return "Loading..."
	}
	return m.CurrentScreen.View()
}
