package deletemodal

import (
	"fmt"

	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
)

var (
	buttonStyle = lipgloss.NewStyle().
			Padding(0, 2).
			Background(lipgloss.Color("#3C3C3C")).
			Foreground(lipgloss.Color("#FFFFFF"))
	activeButtonStyle = lipgloss.NewStyle().
				Padding(0, 2).
				Background(lipgloss.Color("#475bff")).
				Foreground(lipgloss.Color("#FFFFFF")).
				Bold(true)
)

func (m *DeleteDNSRecordModel) View() tea.View {
	yes := buttonStyle.Render("Yes")
	no := buttonStyle.Render("No")
	if m.confirmed == true {
		yes = activeButtonStyle.Render("Yes")
	} else {
		no = activeButtonStyle.Render("No")
	}

	return tea.NewView(fmt.Sprintf("Delete %s record '%s'?\n\n%s   %s", m.RecordType, m.RecordName, yes, no))
}
