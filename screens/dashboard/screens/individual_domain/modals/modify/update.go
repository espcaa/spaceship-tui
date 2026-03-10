package modifymodal

import (
	"fmt"
	"strconv"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/espcaa/spaceship-tui/shared"
)

func (m *ModifyDNSRecordModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	totalPositions := len(m.inputs) + 2

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "esc":
			return m, func() tea.Msg { return shared.CloseModalMsg{} }
		case "tab", "down":
			m.focusIndex = (m.focusIndex + 1) % totalPositions
		case "shift+tab", "up":
			m.focusIndex = (m.focusIndex - 1 + totalPositions) % totalPositions
		case "enter":
			if m.focusIndex == len(m.inputs) {
				// Save button
				ttl, err := strconv.Atoi(m.inputs[1].Value())
				if err != nil {
					m.Error = "Invalid TTL value"
					return m, nil
				}

				for i, def := range m.fieldDefs {
					val := m.inputs[2+i].Value()
					switch def.Type {
					case shared.FieldUint8:
						if _, err := strconv.ParseUint(val, 10, 8); err != nil {
							m.Error = fmt.Sprintf("Invalid %s value", def.Label)
							return m, nil
						}
					case shared.FieldUint16:
						if _, err := strconv.ParseUint(val, 10, 16); err != nil {
							m.Error = fmt.Sprintf("Invalid %s value", def.Label)
							return m, nil
						}
					}
				}

				fields := make(map[string]string)
				for i, def := range m.fieldDefs {
					fields[def.Key] = m.inputs[2+i].Value()
				}

				name := m.inputs[0].Value()
				modified, err := shared.BuildRecord(m.RecordType, name, ttl, m.Group, fields)
				if err != nil {
					m.Error = err.Error()
					return m, nil
				}

				return m, func() tea.Msg {
					return CloseModifyDNSRecordMsg{Confirmed: true, Original: m.Original, Modified: modified}
				}
			} else if m.focusIndex == len(m.inputs)+1 {
				// Cancel button
				return m, func() tea.Msg { return shared.CloseModalMsg{} }
			} else {
				// On an input field, move to next
				m.focusIndex++
			}
		}

		for i := range m.inputs {
			m.inputs[i].Blur()
		}
		if m.focusIndex < len(m.inputs) {
			m.inputs[m.focusIndex].Focus()
		}
	}

	if m.focusIndex < len(m.inputs) {
		m.inputs[m.focusIndex], cmd = m.inputs[m.focusIndex].Update(msg)
	}

	return m, cmd
}
