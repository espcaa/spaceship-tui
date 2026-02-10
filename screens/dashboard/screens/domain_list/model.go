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

func NewDomainListModel(client *spaceship.Client) *DomainListModel {
	l := list.New([]list.Item{}, list.NewDefaultDelegate(), 0, 0)
	l.Title = "Domains"
	return &DomainListModel{
		Domains: []spaceship.DomainInfo{},
		List:    l,
		Client:  client,
	}
}

func (m *DomainListModel) Init() tea.Cmd {
	return func() tea.Msg {
		// get all of the domains and return them as a message
		var skip int = 0
		var take int = 100
		var allDomains []spaceship.DomainInfo

		for {
			domains, err := m.Client.ListDomains(skip, take, "name")
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
