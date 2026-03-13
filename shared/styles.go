package shared

import "charm.land/lipgloss/v2"

var (
	ButtonStyle = lipgloss.NewStyle().
			Padding(0, 2).
			Background(lipgloss.Color("#3C3C3C")).
			Foreground(lipgloss.Color("#FFFFFF"))
	ActiveButtonStyle = lipgloss.NewStyle().
				Padding(0, 2).
				Background(lipgloss.Color("#475bff")).
				Foreground(lipgloss.Color("#FFFFFF")).
				Bold(true)
)
