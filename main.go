package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	app "github.com/espcaa/spaceship-tui/pages"
)

type model struct {
	loggedIn bool
}

func initialModel() model {
	return model{
		loggedIn: false,
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return m, tea.Quit
		}
	}
	return m, nil
}

func (m model) View() string {
	if m.loggedIn {
		return app.AppView()
	}
	return app.LoggedOutView()
}

func main() {
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
