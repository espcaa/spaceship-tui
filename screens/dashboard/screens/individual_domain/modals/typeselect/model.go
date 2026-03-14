package typeselectmodal

import tea "charm.land/bubbletea/v2"

type TypeSelectModel struct {
	DomainName string
	cursor     int
}

func NewTypeSelectModel(domainName string) *TypeSelectModel {
	return &TypeSelectModel{
		DomainName: domainName,
		cursor:     0,
	}
}

func (m *TypeSelectModel) Init() tea.Cmd {
	return nil
}
