package domainlist

import tea "charm.land/bubbletea/v2"

func (m *DomainListModel) View() tea.View {
	return tea.NewView(docStyle.Render(m.List.View()))
}
