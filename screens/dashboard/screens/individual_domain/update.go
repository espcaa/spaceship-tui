package individualdomain

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/espcaa/spaceship-tui/shared"
)

func (m *IndividualDomainModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case DomainDetailsSuccessMsg:
		m.RecordsResponse = msg.Response
		m.State = LoadedState
	case DomainDetailsErrorMsg:
		m.Error = msg.Error
		m.State = LoadedState
	case tea.KeyMsg:
		switch msg.String() {
		case "esc":
			return m, func() tea.Msg {
				return shared.SwitchScreenMsg{Screen: shared.ScreenDomainList}
			}
		}
	}
	return m, nil
}
