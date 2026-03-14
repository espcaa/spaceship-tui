package createmodal

import (
	"fmt"
	"strings"

	tea "charm.land/bubbletea/v2"
	"github.com/espcaa/spaceship-tui/shared"
)

func (m *CreateDNSRecordModel) View() tea.View {
	var b strings.Builder

	fmt.Fprintf(&b, "Create new %s record\n\n", m.RecordType)

	for _, input := range m.inputs {
		b.WriteString(input.View() + "\n")
	}

	b.WriteString("\n")

	create := shared.ButtonStyle.Render("Create")
	cancel := shared.ButtonStyle.Render("Cancel")

	if m.focusIndex == len(m.inputs) {
		create = shared.ActiveButtonStyle.Render("Create")
	} else if m.focusIndex == len(m.inputs)+1 {
		cancel = shared.ActiveButtonStyle.Render("Cancel")
	}

	fmt.Fprintf(&b, "%s \n\n%s", create, cancel)

	if m.Error != "" {
		b.WriteString("\n" + m.Error)
	}

	return tea.NewView(b.String())
}
