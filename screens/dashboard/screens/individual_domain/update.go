package individualdomain

import (
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/espcaa/spaceship-go"
	"github.com/espcaa/spaceship-tui/shared"
)

func (m *IndividualDomainModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.WindowSizeMsg:
		h, v := docStyle.GetFrameSize()
		m.List.SetSize(msg.Width-h, msg.Height-v)
	case DomainDetailsSuccessMsg:
		m.RecordsResponse = msg.Response

		items := make([]list.Item, 0, len(m.RecordsResponse.Items))
		for _, record := range m.RecordsResponse.Items {
			switch r := record.(type) {
			case spaceship.ARecord:
				items = append(items, item{
					title: "A : " + r.Name,
					desc:  r.Adress,
				})
			case spaceship.CNAMERecord:
				items = append(items, item{
					title: "CNAME : " + r.Name,
					desc:  r.CNAME,
				})
			case spaceship.TXTRecord:
				items = append(items, item{
					title: "TXT : " + r.Name,
					desc:  r.Value,
				})
			}

		}

		m.List.SetItems(items)
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

	var cmd tea.Cmd
	m.List, cmd = m.List.Update(msg)
	return m, cmd
}
