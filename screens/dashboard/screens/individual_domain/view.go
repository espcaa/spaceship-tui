package individualdomain

import (
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
)

var modalStyle = lipgloss.NewStyle().
	Border(lipgloss.RoundedBorder()).
	BorderForeground(lipgloss.Color("#7571F9")).
	Padding(1, 3)

func (m *IndividualDomainModel) View() tea.View {
	switch m.State {
	case LoadingState:
		return tea.NewView(docStyle.Render("Loading domain details..."))
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
			return tea.NewView(comp.Render())
		}
		if m.Modal != nil {
			modalContent := modalStyle.Render(m.Modal.View().Content)
			x := (m.width - lipgloss.Width(modalContent)) / 2
			y := (m.height - lipgloss.Height(modalContent)) / 2

			comp := lipgloss.NewCompositor(
				lipgloss.NewLayer(base),
				lipgloss.NewLayer(modalContent).X(x).Y(y).Z(1),
			)
			return tea.NewView(comp.Render())
		}
		return tea.NewView(base)
	default:
		return tea.NewView(docStyle.Render("Unknown state"))
	}
}
