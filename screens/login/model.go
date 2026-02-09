package login

import (
	"strings"
	"time"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/espcaa/spaceship-tui/utils"
)

type LoginModel struct {
	apiKey        string
	apiSecret     string
	focusIndex    int
	textInputs    []textinput.Model
	wavePos       int
	letters       []Letter
	logoSize      Vector2
	errorText     string
	currentSpaces float64
}

type Letter struct {
	char rune
	pos  Vector2
}

var inputs = []string{"API Key", "API Secret"}

func NewLoginModel() LoginModel {

	var textInputs []textinput.Model

	for i := range inputs {
		ti := textinput.New()
		ti.Placeholder = inputs[i]
		ti.PromptStyle = ti.PromptStyle.Bold(true)
		ti.TextStyle = ti.TextStyle.Bold(true)
		if i == 1 {
			ti.EchoMode = textinput.EchoPassword
			ti.EchoCharacter = '•'
		}
		ti.CharLimit = 64
		ti.Width = 40

		if i == 0 {
			ti.Focus()
		}

		textInputs = append(textInputs, ti)
	}

	m := LoginModel{
		apiKey:     "",
		apiSecret:  "",
		textInputs: textInputs,
	}

	lines := strings.Split(utils.Logo, "\n")
	for y, line := range lines {
		for x, ch := range line {
			m.letters = append(m.letters, Letter{
				char: ch,
				pos:  Vector2{X: x, Y: y},
			})

		}
		m.letters = append(m.letters, Letter{
			char: '\n',
			pos:  Vector2{X: 0, Y: y + 1},
		})
	}
	m.currentSpaces = 0
	m.logoSize = Vector2{X: len(lines[0]), Y: len(lines)}

	return m
}

type waveTickMsg time.Time

type LoginErrorMsg struct {
	Error string
}

func waveTick() tea.Cmd {
	return tea.Tick(time.Millisecond*20, func(t time.Time) tea.Msg {
		return waveTickMsg(t)
	})
}

func (m LoginModel) Init() tea.Cmd {
	return tea.Batch(textinput.Blink, waveTick())
}
