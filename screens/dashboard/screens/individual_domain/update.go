package individualdomain

import (
	"fmt"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/espcaa/spaceship-go"
	deletemodal "github.com/espcaa/spaceship-tui/screens/dashboard/screens/individual_domain/modals/delete"
	modifymodal "github.com/espcaa/spaceship-tui/screens/dashboard/screens/individual_domain/modals/modify"
	"github.com/espcaa/spaceship-tui/shared"
)

func (m *IndividualDomainModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	if closeMsg, ok := msg.(deletemodal.CloseDeleteDNSRecordMsg); ok {
		m.Modal = nil
		if closeMsg.Confirmed {
			idx := m.List.Index()
			items := m.List.Items()
			m.List.SetItems(append(items[:idx], items[idx+1:]...))
			record := closeMsg.Record
			return m, func() tea.Msg {
				err := m.Client.DeleteDNSRecords(m.Domain.Name, []spaceship.DNSRecord{record})
				if err != nil {
					return DomainDetailsErrorMsg{Error: err.Error()}
				}
				return nil
			}
		}
		return m, nil
	}

	if m.Modal != nil {
		var cmd tea.Cmd
		m.Modal, cmd = m.Modal.Update(msg)
		return m, cmd
	}

	switch msg := msg.(type) {

	case tea.WindowSizeMsg:
		h, v := docStyle.GetFrameSize()
		m.List.SetSize(msg.Width-h, msg.Height-v)
		m.width = msg.Width
		m.height = msg.Height
	case DomainDetailsSuccessMsg:
		m.RecordsResponse = msg.Response

		items := make([]list.Item, 0, len(m.RecordsResponse.Items))
		for _, record := range m.RecordsResponse.Items {
			switch r := record.(type) {
			case spaceship.ARecord:
				items = append(items, item{
					title:  "A : " + r.Name,
					desc:   r.Adress,
					record: r,
				})
			case spaceship.AAAARecord:
				items = append(items, item{
					title:  "AAAA : " + r.Name,
					desc:   r.Adress,
					record: r,
				})
			case spaceship.CNAMERecord:
				items = append(items, item{
					title:  "CNAME : " + r.Name,
					desc:   r.CNAME,
					record: r,
				})
			case spaceship.TXTRecord:
				items = append(items, item{
					title:  "TXT : " + r.Name,
					desc:   r.Value,
					record: r,
				})
			case spaceship.MXRecord:
				items = append(items, item{
					title:  "MX : " + r.Name,
					desc:   fmt.Sprintf("%s (priority %d)", r.Exchange, r.Preference),
					record: r,
				})
			case spaceship.NSRecord:
				items = append(items, item{
					title:  "NS : " + r.Name,
					desc:   r.Nameserver,
					record: r,
				})
			case spaceship.SRVRecord:
				items = append(items, item{
					title:  "SRV : " + r.Name,
					desc:   fmt.Sprintf("%s:%d (priority %d, weight %d)", r.Target, r.Port, r.Priority, r.Weight),
					record: r,
				})
			case spaceship.CAARecord:
				items = append(items, item{
					title:  "CAA : " + r.Name,
					desc:   fmt.Sprintf("%d %s \"%s\"", r.Flag, r.Tag, r.Value),
					record: r,
				})
			case spaceship.AliasRecord:
				items = append(items, item{
					title:  "ALIAS : " + r.Name,
					desc:   r.AliasTarget,
					record: r,
				})
			case spaceship.PTRRecord:
				items = append(items, item{
					title:  "PTR : " + r.Name,
					desc:   r.Pointer,
					record: r,
				})
			case spaceship.TLSARecord:
				items = append(items, item{
					title:  "TLSA : " + r.Name,
					desc:   fmt.Sprintf("usage %d, selector %d, matching %d", r.Usage, r.Selector, r.Matching),
					record: r,
				})
			case spaceship.HTTPSRecord:
				items = append(items, item{
					title:  "HTTPS : " + r.Name,
					desc:   fmt.Sprintf("%s (priority %d)", r.TargetName, r.SvcPriority),
					record: r,
				})
			case spaceship.SVCBRecord:
				items = append(items, item{
					title:  "SVCB : " + r.Name,
					desc:   fmt.Sprintf("%s (priority %d)", r.TargetName, r.SvcPriority),
					record: r,
				})
			}
		}

		m.List.SetItems(items)
		m.State = LoadedState
	case DomainDetailsErrorMsg:
		m.Error = msg.Error
		m.State = LoadedState
	case tea.KeyMsg:
		if m.Error != "" {
			m.Error = ""
			return m, nil
		}
		switch msg.String() {
		case "esc":
			return m, func() tea.Msg {
				return shared.SwitchScreenMsg{Screen: shared.ScreenDomainList}
			}

		case "d":
			selected, ok := m.List.SelectedItem().(item)
			if ok {
				m.Modal = deletemodal.NewDeleteDNSRecordModel(
					m.Domain.Name, selected.record,
					selected.title, selected.record.GetType(),
				)
			}
		case "m":
			selected, ok := m.List.SelectedItem().(item)
			if ok {
				m.Modal = modifymodal.NewModifyDNSRecordModel(
					m.Domain.Name, selected.record,
					selected.title, selected.record.GetGroup(),
				)
			}
		}
	}

	var cmd tea.Cmd
	m.List, cmd = m.List.Update(msg)
	return m, cmd
}
