package individualdomain

func (m *IndividualDomainModel) View() string {
	switch m.State {
	case LoadingState:
		return docStyle.Render("Loading domain details...")
	case LoadedState:
		return docStyle.Render(m.List.View())
	default:
		return docStyle.Render("Unknown state")
	}
}
