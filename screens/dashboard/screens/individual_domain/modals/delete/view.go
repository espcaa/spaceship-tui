package deletemodal

import "fmt"

func (m *DeleteDNSRecordModel) View() string {
	yes, no := "  Yes  ", "[ No ]"
	if m.confirmed {
		yes, no = "[ Yes ]", "  No  "
	}
	return fmt.Sprintf("Delete %s record '%s'?\n\n%s   %s", m.RecordType, m.RecordName, yes, no)
}
