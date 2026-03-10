package individualdomain

import "charm.land/lipgloss/v2"

var modalStyle = lipgloss.NewStyle().
	Border(lipgloss.RoundedBorder()).
	BorderForeground(lipgloss.Color("#7571F9")).
	Padding(1, 2).
	Width(50).
	Height(15)

func (m *IndividualDomainModel) View() string {
	switch m.State {
	case LoadingState:
		return docStyle.Render("Loading domain details...")
	case LoadedState:
		base := docStyle.Render(m.List.View())
		if m.Error != "" {
			errBox := lipgloss.NewStyle().
				Border(lipgloss.RoundedBorder()).
				BorderForeground(lipgloss.Color("#FF0000")).
				Padding(1, 2).
				Render(m.Error)
			x := (m.width - lipgloss.Width(errBox)) / 2
			y := (m.height - lipgloss.Height(errBox)) / 2
			comp := lipgloss.NewCompositor(
				lipgloss.NewLayer(base),
				lipgloss.NewLayer(errBox).X(x).Y(y).Z(1),
			)
			return comp.Render()
		}
		if m.Modal != nil {
			modal := modalStyle.Render(m.Modal.View())
			x := (m.width - lipgloss.Width(modal)) / 2
			y := (m.height - lipgloss.Height(modal)) / 2

			comp := lipgloss.NewCompositor(
				lipgloss.NewLayer(base),
				lipgloss.NewLayer(modal).X(x).Y(y).Z(1),
			)
			return comp.Render()
		}
		return base
	default:
		return docStyle.Render("Unknown state")
	}
}
