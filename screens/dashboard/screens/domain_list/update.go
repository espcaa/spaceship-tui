package domainlist

import (
	list "charm.land/bubbles/v2/list"
	tea "charm.land/bubbletea/v2"

	"github.com/espcaa/spaceship-tui/shared"
)

func (m *DomainListModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyPressMsg:
		if m.List.FilterState() == list.Filtering {
			break
		}

		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit
		case "enter":
			if len(m.List.Items()) == 0 {
				return m, nil
			}
			if selectedItem, ok := m.List.SelectedItem().(item); ok {
				for _, domain := range m.Domains {
					if domain.Name == selectedItem.title {
						return m, func() tea.Msg {
							return shared.DomainSelectedMsg{Domain: domain}
						}
					}
				}
			}
		}
	case tea.WindowSizeMsg:
		h, v := docStyle.GetFrameSize()
		m.List.SetSize(msg.Width-h, msg.Height-v)
	case DomainListSuccessMsg:
		m.Domains = msg.Domains

		items := make([]list.Item, len(m.Domains))
		for i, domain := range m.Domains {
			items[i] = item{
				title: domain.Name,
			}
		}
		cmds = append(cmds, m.List.SetItems(items))
	}

	var cmd tea.Cmd
	m.List, cmd = m.List.Update(msg)
	cmds = append(cmds, cmd)

	return m, tea.Batch(cmds...)
}
