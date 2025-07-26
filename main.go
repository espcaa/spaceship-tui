package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	app "github.com/espcaa/spaceship-tui/app"
)

type model struct {
	loggedIn bool
	setup    app.SetupModel
	app      app.AppModel
}

func initialModel() model {
	return model{
		loggedIn: false,
		setup:    app.NewSetupModel(),
		app:      app.NewAppModel(),
	}
}

func (m model) Init() tea.Cmd {
	m.setup.Init()
	m.app.Init()
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit
		}
	}

	var cmd tea.Cmd
	if m.loggedIn {

		m.app, cmd = m.app.Update(msg)
	} else {
		m.setup, cmd = m.setup.Update(msg)
	}
	return m, cmd
}

func (m model) View() string {
	if m.loggedIn {
		return app.AppView()
	}
	return m.setup.View()
}

func main() {
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("OOoops i think we have a problem: %v", err)
		os.Exit(1)
	}
}
