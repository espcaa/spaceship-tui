package individualdomain

import (
	"strings"

	"github.com/espcaa/spaceship-go"
)

func (m *IndividualDomainModel) View() string {
	switch m.State {
	case LoadingState:
		return docStyle.Render("Loading domain details...")
	case LoadedState:
		return m.renderDomainDetails()
	default:
		return docStyle.Render("Unknown state")
	}
}

func (m *IndividualDomainModel) renderDomainDetails() string {
	if m.Error != "" {
		return docStyle.Render("Error: " + m.Error)
	}

	if m.RecordsResponse.Total == 0 {
		return docStyle.Render("No DNS records found for this domain.")
	}

	var result strings.Builder
	for _, record := range m.RecordsResponse.Items {
		switch r := record.(type) {
		case spaceship.ARecord:
			result.WriteString("A Record: " + r.Name + " -> " + r.Adress + "\n")
		case spaceship.TXTRecord:
			result.WriteString("TXT Record: " + r.Name + " -> " + r.Value + "\n")
		case spaceship.CNAMERecord:
			result.WriteString("CNAME Record: " + r.Name + " -> " + r.CNAME + "\n")
		case spaceship.MXRecord:
			result.WriteString("MX Record: " + r.Name + " -> " + r.Exchange + " (Priority: " + string(r.Preference) + ")\n")
		case spaceship.NSRecord:
			result.WriteString("NS Record: " + r.Name + " -> " + r.Nameserver + "\n")
		}

	}

	return docStyle.Render(result.String())
}
