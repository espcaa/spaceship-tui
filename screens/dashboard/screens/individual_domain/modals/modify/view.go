package modifymodal

import (
	"fmt"
	"strings"

	tea "charm.land/bubbletea/v2"
	"github.com/espcaa/spaceship-tui/shared"
)

func (m *ModifyDNSRecordModel) View() tea.View {
	var b strings.Builder

	b.WriteString(fmt.Sprintf("Modify %s Record\n", m.RecordType))

	for _, input := range m.inputs {
		b.WriteString(input.View() + "\n")
	}

	b.WriteString("\n")

	save := shared.ButtonStyle.Render("Save")
	cancel := shared.ButtonStyle.Render("Cancel")

	if m.focusIndex == len(m.inputs) {
		save = shared.ActiveButtonStyle.Render("Save")
	} else if m.focusIndex == len(m.inputs)+1 {
		cancel = shared.ActiveButtonStyle.Render("Cancel")
	}

	fmt.Fprintf(&b, "%s \n\n%s", save, cancel)

	if m.Error != "" {
		b.WriteString("\n" + m.Error)
	}

	return tea.NewView(b.String())
}
