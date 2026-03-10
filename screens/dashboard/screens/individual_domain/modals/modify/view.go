package modifymodal

import (
	"fmt"
	"strings"

	tea "charm.land/bubbletea/v2"
)

func (m *ModifyDNSRecordModel) View() tea.View {
	var b strings.Builder

	b.WriteString(fmt.Sprintf("Modify %s Record\n", m.RecordType))

	for _, input := range m.inputs {
		b.WriteString(input.View() + "\n")
	}

	b.WriteString("\n")

	save, cancel := "  Save  ", "  Cancel  "
	if m.focusIndex == len(m.inputs) {
		save = "[ Save ]"
	} else if m.focusIndex == len(m.inputs)+1 {
		cancel = "[ Cancel ]"
	}
	b.WriteString(fmt.Sprintf("%s   %s", save, cancel))

	if m.Error != "" {
		b.WriteString("\n" + m.Error)
	}

	return tea.NewView(b.String())
}
