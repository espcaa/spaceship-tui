package login

import (
	"fmt"
	"math"

	"github.com/charmbracelet/lipgloss"
)

var boxStyle = lipgloss.NewStyle().Padding(1, 2).
	Border(lipgloss.NormalBorder()).
	BorderForeground(lipgloss.Color("15")).
	Padding(1, 2).
	Margin(1, 2).
	Align(lipgloss.Center)

var errorStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("9")).Bold(true)

var subtitleStyle = lipgloss.NewStyle().Italic(true).Foreground(lipgloss.Color("8"))

var baseColor = lipgloss.Color("#475bff")
var waveColor = lipgloss.Color("15")

func (m LoginModel) View() string {
	var buttons string
	for _, ti := range m.textInputs {
		buttons += ti.View() + "\n"
	}

	var logoBlock = m.renderAnimatedLogo(int(m.currentSpaces))

	form := lipgloss.JoinVertical(lipgloss.Top,
		lipgloss.NewStyle().Bold(true).Render("Spaceship"),
		subtitleStyle.Render("Enter your Spaceship API secrets and key below:\n\n"),
		buttons,
		subtitleStyle.Render("\nPress Tab to switch between fields. Press Enter to submit."),
		errorStyle.Render(m.errorText),
	)

	content := lipgloss.JoinHorizontal(lipgloss.Top, logoBlock, form)

	return boxStyle.Render(content)
}

func (m LoginModel) renderAnimatedLogo(numberOfSpacesPerLine ...int) string {
	var logostring string
	spaceSinceLastLine := 0
	spacesAllowed := 0
	if len(numberOfSpacesPerLine) > 0 {
		spacesAllowed = numberOfSpacesPerLine[0]
	}

	for i := range m.letters {
		l := m.letters[i]

		if l.char == '\n' {
			logostring += "\n"
			spaceSinceLastLine = 0
			continue
		}

		glow := 0.15*math.Sin(float64(m.wavePos)*0.08) + 1.0
		baseColor := hexGlow("#475bff", glow)
		var styled string
		styled = lipgloss.NewStyle().Foreground(baseColor).Render(string(l.char))

		if l.char != ' ' {
			logostring += styled
		} else if spaceSinceLastLine < spacesAllowed {
			logostring += styled
			spaceSinceLastLine++
		}
	}

	return logostring
}

func hexGlow(base string, glow float64) lipgloss.Color {
	var r, g, b uint8
	fmt.Sscanf(base, "#%02x%02x%02x", &r, &g, &b)

	// glow: 0.0 → 1.0
	// scale brightness subtly
	factor := lerp(0.85, 1.15, glow)

	nr := clamp(float64(r)*factor, 0, 255)
	ng := clamp(float64(g)*factor, 0, 255)
	nb := clamp(float64(b)*factor, 0, 255)

	return lipgloss.Color(fmt.Sprintf("#%02x%02x%02x", int(nr), int(ng), int(nb)))
}

func lerp(a, b, t float64) float64 {
	return a + (b-a)*t
}

func clamp(v, min, max float64) float64 {
	if v < min {
		return min
	}
	if v > max {
		return max
	}
	return v
}
