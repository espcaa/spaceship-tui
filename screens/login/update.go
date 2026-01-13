package login

import (
	tea "github.com/charmbracelet/bubbletea"
)

type LoginSuccessMsg struct {
	ApiKey    string
	ApiSecret string
}

type Vector2 struct {
	X, Y int
}

func (m LoginModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case waveTickMsg:
		m.wavePos += 1
		if m.wavePos > (m.logoSize.X+m.logoSize.Y)*3 {
			m.wavePos = 0
		}

		m.currentSpaces += 1.0

		return m, waveTick()

	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyTab, tea.KeyShiftTab, tea.KeyUp, tea.KeyDown:
			s := msg.String()

			if s == "up" || s == "shift+tab" {
				m.focusIndex--
			} else {
				m.focusIndex++
			}

			if m.focusIndex > len(m.textInputs)-1 {
				m.focusIndex = 0
			} else if m.focusIndex < 0 {
				m.focusIndex = len(m.textInputs) - 1
			}

			for i := 0; i < len(m.textInputs); i++ {
				if i == m.focusIndex {
					m.textInputs[i].Focus()
					continue
				}
				m.textInputs[i].Blur()
			}
			return m, nil
		}

	}

	for i := range m.textInputs {
		m.textInputs[i], cmd = m.textInputs[i].Update(msg)
	}
	return m, cmd
}
