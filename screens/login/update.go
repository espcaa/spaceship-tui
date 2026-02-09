package login

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/espcaa/spaceship-go"
)

type LoginSuccessMsg struct {
	ApiKey    string
	ApiSecret string
}

type Vector2 struct {
	X, Y int
}

func loginCmd(apiKey, apiSecret string) tea.Cmd {
	return func() tea.Msg {
		err := spaceship.NewClient(apiKey, apiSecret).VerifyCredentials()

		if err != nil {
			return LoginErrorMsg{Error: err.Error()}
		}

		return LoginSuccessMsg{
			ApiKey:    apiKey,
			ApiSecret: apiSecret,
		}
	}
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

	case LoginErrorMsg:
		m.errorText = msg.Error
		return m, nil

	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEnter:
			return m, loginCmd(m.textInputs[0].Value(), m.textInputs[1].Value())

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
