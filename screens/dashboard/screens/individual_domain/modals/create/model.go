package createmodal

import (
	"charm.land/bubbles/v2/textinput"
	tea "charm.land/bubbletea/v2"

	"github.com/espcaa/spaceship-go"
	"github.com/espcaa/spaceship-tui/shared"
)

type CreateDNSRecordModel struct {
	Domain     string
	RecordType string
	Group      string
	inputs     []textinput.Model
	fieldDefs  []shared.FieldDef
	focusIndex int
	Error      string
}

type CloseCreateDNSRecordMsg struct {
	Confirmed bool
	Record    spaceship.DNSRecord
}

func NewCreateDNSRecordModel(domain string, recordType string, group string) *CreateDNSRecordModel {
	fieldDefs := shared.RecordFieldDefs[recordType]

	inputs := make([]textinput.Model, 2+len(fieldDefs))

	inputs[0] = textinput.New()
	inputs[0].Prompt = "Name: "
	inputs[0].CharLimit = 256
	inputs[0].SetWidth(40)
	inputs[0].Focus()

	inputs[1] = textinput.New()
	inputs[1].Prompt = "TTL: "
	inputs[1].CharLimit = 256
	inputs[1].SetWidth(40)

	// set default TTL
	inputs[1].SetValue("3600")

	for i, def := range fieldDefs {
		ti := textinput.New()
		ti.Prompt = def.Label + ": "
		ti.CharLimit = 256
		ti.SetWidth(40)
		inputs[2+i] = ti
	}

	return &CreateDNSRecordModel{
		Domain:     domain,
		RecordType: recordType,
		Group:      group,
		inputs:     inputs,
		fieldDefs:  fieldDefs,
		focusIndex: 0,
		Error:      "",
	}
}

func (m *CreateDNSRecordModel) Init() tea.Cmd {
	return textinput.Blink
}
