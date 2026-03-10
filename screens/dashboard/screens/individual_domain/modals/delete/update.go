package deletemodal

import (
	tea "charm.land/bubbletea/v2"

	"github.com/espcaa/spaceship-tui/shared"
)

func (m *DeleteDNSRecordModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyPressMsg:
		switch msg.String() {
		case "y", "Y":
			m.confirmed = true
		case "n", "N":
			m.confirmed = false
		case "left", "right", "tab":
			m.confirmed = !m.confirmed
		case "enter":
			return m, func() tea.Msg { return CloseDeleteDNSRecordMsg{Confirmed: m.confirmed} }
		case "esc":
			return m, func() tea.Msg { return shared.CloseModalMsg{} }
		}
	}
	return m, nil
}
