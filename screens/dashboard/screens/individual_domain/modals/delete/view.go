package deletemodal

import (
	"fmt"

	tea "charm.land/bubbletea/v2"
	"github.com/espcaa/spaceship-tui/shared"
)

func (m *DeleteDNSRecordModel) View() tea.View {
	yes := shared.ButtonStyle.Render("Yes")
	no := shared.ButtonStyle.Render("No")
	if m.confirmed == true {
		yes = shared.ActiveButtonStyle.Render("Yes")
	} else {
		no = shared.ActiveButtonStyle.Render("No")
	}

	return tea.NewView(fmt.Sprintf("Delete %s record '%s'?\n\n%s   %s", m.RecordType, m.RecordName, yes, no))
}
