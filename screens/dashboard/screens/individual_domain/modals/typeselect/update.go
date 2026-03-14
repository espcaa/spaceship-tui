package typeselectmodal

import (
	tea "charm.land/bubbletea/v2"

	"github.com/espcaa/spaceship-tui/shared"
)

func (m *TypeSelectModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyPressMsg:
		switch msg.String() {
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}
		case "down", "j":
			if m.cursor < len(shared.RecordTypes)-1 {
				m.cursor++
			}
		case "enter":
			return m, func() tea.Msg {
				return shared.TypeSelectedMsg{RecordType: shared.RecordTypes[m.cursor]}
			}
		case "esc":
			return m, func() tea.Msg { return shared.CloseModalMsg{} }
		}
	}
	return m, nil
}
