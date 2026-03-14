package typeselectmodal

import (
	"strings"

	tea "charm.land/bubbletea/v2"
	"github.com/espcaa/spaceship-tui/shared"
)

func (m *TypeSelectModel) View() tea.View {
	var b strings.Builder

	b.WriteString("Select record type:\n\n")

	for i, rt := range shared.RecordTypes {
		if i == m.cursor {
			b.WriteString("  " + shared.ActiveButtonStyle.Render(rt) + "\n")
		} else {
			b.WriteString("  " + shared.ButtonStyle.Render(rt) + "\n")
		}
	}

	return tea.NewView(b.String())
}
