package domainlist

func (m *DomainListModel) View() string {
	return docStyle.Render(m.List.View())
}
