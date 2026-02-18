package individualdomain

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/espcaa/spaceship-go"
)

var docStyle = lipgloss.NewStyle().Margin(1, 2)

type State int

const (
	LoadingState State = iota
	LoadedState
)

type IndividualDomainModel struct {
	Domain          spaceship.DomainInfo
	Client          *spaceship.Client
	RecordsResponse spaceship.ListDNSRecordsResponse
	State           State
	Error           string
}

type DomainDetailsSuccessMsg struct {
	Response spaceship.ListDNSRecordsResponse
}

type DomainDetailsErrorMsg struct {
	Error string
}

func NewIndividualDomainModel(domain spaceship.DomainInfo, client *spaceship.Client) *IndividualDomainModel {
	return &IndividualDomainModel{
		Domain: domain,
		Client: client,
	}
}

func (m *IndividualDomainModel) Init() tea.Cmd {
	return func() tea.Msg {
		var skip int = 0
		var take int = 100
		var allRecords spaceship.ListDNSRecordsResponse

		for {
			records, err := m.Client.GetDomainDNSRecords(m.Domain.Name, take, skip, "name")
			if err != nil {
				return DomainDetailsErrorMsg{Error: err.Error()}
			}

			allRecords = spaceship.ListDNSRecordsResponse{
				Items: append(allRecords.Items, records.Items...),
				Total: allRecords.Total + records.Total,
			}

			if len(records.Items) < take {
				break
			}

			skip += take
		}

		return DomainDetailsSuccessMsg{Response: allRecords}
	}
}
