package app

import (
	tea "github.com/charmbracelet/bubbletea"
)

type AppModel struct {
	// Nothing here yet
}

func NewAppModel() AppModel {
	return AppModel{}
}

func (m AppModel) Init() tea.Cmd {
	return nil
}

func (m AppModel) Update(msg tea.Msg) (AppModel, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit
		case "enter":
			// Handle submission!!
		}
	}

	return m, nil
}

func AppView() string {
	return `
		Welcome to Spaceship tui!

		This is the main application view.

		Press 'q' to quit.
	`
}
