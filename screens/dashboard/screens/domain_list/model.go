package domainlist

import (
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/lipgloss"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/espcaa/spaceship-go"
)

type DomainListErrorMsg struct {
	Error string
}

type DomainListSuccessMsg struct {
	Domains []spaceship.DomainInfo
}

var docStyle = lipgloss.NewStyle().Margin(1, 2)

type DomainListModel struct {
	Domains []spaceship.DomainInfo
	List    list.Model
	Client  *spaceship.Client
}

func NewDomainListModel(client *spaceship.Client, cachedDomains []spaceship.DomainInfo) *DomainListModel {
	l := list.New([]list.Item{}, list.NewDefaultDelegate(), 0, 0)
	l.Title = "Domains"
	return &DomainListModel{
		Domains: cachedDomains,
		List:    l,
		Client:  client,
	}
}

func (m *DomainListModel) Init() tea.Cmd {
	if len(m.Domains) > 0 {
		domains := m.Domains
		return func() tea.Msg {
			return DomainListSuccessMsg{Domains: domains}
		}
	}
	return func() tea.Msg {
		// get all of the domains and return them as a message
		var skip int = 0
		var take int = 100
		var allDomains []spaceship.DomainInfo

		for {
			domains, err := m.Client.ListDomains(take, skip, "name")
			if err != nil {
				return DomainListErrorMsg{Error: err.Error()}
			}

			allDomains = append(allDomains, domains.Items...)

			if len(domains.Items) < take {
				break
			}

			skip += take
		}

		return DomainListSuccessMsg{Domains: allDomains}
	}
}

type item struct {
	title, desc string
}

func (i item) Title() string       { return i.title }
func (i item) Description() string { return i.desc }
func (i item) FilterValue() string { return i.title }
