package deletemodal

import (
	"fmt"

	tea "charm.land/bubbletea/v2"
)

func (m *DeleteDNSRecordModel) View() tea.View {
	yes, no := "  Yes  ", "[ No ]"
	if m.confirmed {
		yes, no = "[ Yes ]", "  No  "
	}
	return tea.NewView(fmt.Sprintf("Delete %s record '%s'?\n\n%s   %s", m.RecordType, m.RecordName, yes, no))
}
