package individualdomain

import (
	"database/sql"

	"charm.land/bubbles/v2/key"
	"charm.land/bubbles/v2/list"
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
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
	Db              *sql.DB
	List            list.Model
	Modal           tea.Model
	height          int
	width           int
}

type DomainDetailsSuccessMsg struct {
	Response spaceship.ListDNSRecordsResponse
}

type DomainDetailsErrorMsg struct {
	Error string
}

func NewIndividualDomainModel(domain spaceship.DomainInfo, client *spaceship.Client, db *sql.DB) *IndividualDomainModel {
	l := list.New([]list.Item{}, list.NewDefaultDelegate(), 0, 0)
	l.Title = "DNS Records: " + domain.Name
	l.AdditionalShortHelpKeys = func() []key.Binding {
		return []key.Binding{
			key.NewBinding(key.WithKeys("del"), key.WithHelp("del", "delete")),
		}
	}
	return &IndividualDomainModel{
		Domain: domain,
		List:   l,
		Db:     db,
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

type item struct {
	title, desc string
	record      spaceship.DNSRecord
}

func (i item) Title() string       { return i.title }
func (i item) Description() string { return i.desc }
func (i item) FilterValue() string { return i.title }
