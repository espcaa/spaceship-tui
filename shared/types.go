package shared

import "github.com/espcaa/spaceship-go"

type DashboardErrorMsg struct {
	Error string
}

type DomainSelectedMsg struct {
	Domain spaceship.DomainInfo
}

type Screen int

const (
	ScreenDomainList Screen = iota
	ScreenIndividualDomain
)

type SwitchScreenMsg struct {
	Screen Screen
}
