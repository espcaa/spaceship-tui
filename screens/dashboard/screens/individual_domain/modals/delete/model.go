package deletemodal

import (
	tea "charm.land/bubbletea/v2"

	"github.com/espcaa/spaceship-go"
)

type DeleteDNSRecordModel struct {
	Domain     string
	Record     spaceship.DNSRecord
	RecordName string
	RecordType string
	confirmed  bool
	Error      string
}

type CloseDeleteDNSRecordMsg struct {
	Confirmed bool
	Record    spaceship.DNSRecord
}

func NewDeleteDNSRecordModel(domain string, record spaceship.DNSRecord, recordName string, recordType string) *DeleteDNSRecordModel {
	return &DeleteDNSRecordModel{
		Domain:     domain,
		Record:     record,
		RecordName: recordName,
		RecordType: recordType,
		Error:      "",
	}
}

func (m *DeleteDNSRecordModel) Init() tea.Cmd {
	return nil
}
