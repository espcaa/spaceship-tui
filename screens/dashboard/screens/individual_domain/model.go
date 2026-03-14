package individualdomain

import (
	"database/sql"
	"fmt"

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
			key.NewBinding(key.WithKeys("enter"), key.WithHelp("enter", "modify")),
			key.NewBinding(key.WithKeys("a"), key.WithHelp("a", "add")),
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

func recordToItem(record spaceship.DNSRecord) item {
	switch r := record.(type) {
	case spaceship.ARecord:
		return item{title: "A : " + r.Name, desc: r.Adress, record: r}
	case spaceship.AAAARecord:
		return item{title: "AAAA : " + r.Name, desc: r.Adress, record: r}
	case spaceship.CNAMERecord:
		return item{title: "CNAME : " + r.Name, desc: r.CNAME, record: r}
	case spaceship.TXTRecord:
		return item{title: "TXT : " + r.Name, desc: r.Value, record: r}
	case spaceship.MXRecord:
		return item{title: "MX : " + r.Name, desc: fmt.Sprintf("%s (priority %d)", r.Exchange, r.Preference), record: r}
	case spaceship.NSRecord:
		return item{title: "NS : " + r.Name, desc: r.Nameserver, record: r}
	case spaceship.SRVRecord:
		return item{title: "SRV : " + r.Name, desc: fmt.Sprintf("%s:%d (priority %d, weight %d)", r.Target, r.Port, r.Priority, r.Weight), record: r}
	case spaceship.CAARecord:
		return item{title: "CAA : " + r.Name, desc: fmt.Sprintf("%d %s \"%s\"", r.Flag, r.Tag, r.Value), record: r}
	case spaceship.AliasRecord:
		return item{title: "ALIAS : " + r.Name, desc: r.AliasTarget, record: r}
	case spaceship.PTRRecord:
		return item{title: "PTR : " + r.Name, desc: r.Pointer, record: r}
	case spaceship.TLSARecord:
		return item{title: "TLSA : " + r.Name, desc: fmt.Sprintf("usage %d, selector %d, matching %d", r.Usage, r.Selector, r.Matching), record: r}
	case spaceship.HTTPSRecord:
		return item{title: "HTTPS : " + r.Name, desc: fmt.Sprintf("%s (priority %d)", r.TargetName, r.SvcPriority), record: r}
	case spaceship.SVCBRecord:
		return item{title: "SVCB : " + r.Name, desc: fmt.Sprintf("%s (priority %d)", r.TargetName, r.SvcPriority), record: r}
	default:
		return item{title: "Unknown", desc: "", record: record}
	}
}
