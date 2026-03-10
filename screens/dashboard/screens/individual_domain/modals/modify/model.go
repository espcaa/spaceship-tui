package modifymodal

import (
	"charm.land/bubbles/v2/textinput"
	tea "charm.land/bubbletea/v2"

	"github.com/espcaa/spaceship-go"
	"github.com/espcaa/spaceship-tui/shared"
)

type ModifyDNSRecordModel struct {
	Domain     string
	Original   spaceship.DNSRecord
	RecordType string
	Group      spaceship.DNSRecordGroup
	inputs     []textinput.Model
	fieldDefs  []shared.FieldDef
	focusIndex int
	Error      string
}

type CloseModifyDNSRecordMsg struct {
	Confirmed bool
	Original  spaceship.DNSRecord
	Modified  spaceship.DNSRecord
}

func NewModifyDNSRecordModel(domain string, record spaceship.DNSRecord, recordType string, group spaceship.DNSRecordGroup) *ModifyDNSRecordModel {
	fieldDefs := shared.RecordFieldDefs[recordType]

	inputs := make([]textinput.Model, 2+len(fieldDefs))

	inputs[0] = textinput.New()
	inputs[0].Prompt = "Name: "
	inputs[0].CharLimit = 256
	inputs[0].SetWidth(40)
	inputs[0].SetValue(shared.GetFieldValue(record, "name"))
	inputs[0].Focus()

	inputs[1] = textinput.New()
	inputs[1].Prompt = "TTL: "
	inputs[1].CharLimit = 256
	inputs[1].SetWidth(40)
	inputs[1].SetValue(shared.GetFieldValue(record, "ttl"))

	for i, def := range fieldDefs {
		ti := textinput.New()
		ti.Prompt = def.Label + ": "
		ti.CharLimit = 256
		ti.SetWidth(40)
		ti.SetValue(shared.GetFieldValue(record, def.Key))
		inputs[2+i] = ti
	}

	return &ModifyDNSRecordModel{
		Domain:     domain,
		Original:   record,
		RecordType: recordType,
		Group:      group,
		inputs:     inputs,
		fieldDefs:  fieldDefs,
		focusIndex: 0,
		Error:      "",
	}
}

func (m *ModifyDNSRecordModel) Init() tea.Cmd {
	return textinput.Blink
}
