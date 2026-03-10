package shared

type FieldType int

const (
	FieldString FieldType = iota
	FieldUint8
	FieldUint16
)

type FieldDef struct {
	Label string
	Key   string
	Type  FieldType
}

var RecordFieldDefs = map[string][]FieldDef{
	"A": {
		{Label: "Address", Key: "address", Type: FieldString},
	},
	"AAAA": {
		{Label: "Address", Key: "address", Type: FieldString},
	},
	"ALIAS": {
		{Label: "Alias Target", Key: "aliasTarget", Type: FieldString},
	},
	"CNAME": {
		{Label: "CNAME", Key: "cname", Type: FieldString},
	},
	"TXT": {
		{Label: "Value", Key: "value", Type: FieldString},
	},
	"NS": {
		{Label: "Nameserver", Key: "nameserver", Type: FieldString},
	},
	"PTR": {
		{Label: "Pointer", Key: "pointer", Type: FieldString},
	},
	"MX": {
		{Label: "Exchange", Key: "exchange", Type: FieldString},
		{Label: "Preference", Key: "preference", Type: FieldUint16},
	},
	"CAA": {
		{Label: "Flag", Key: "flag", Type: FieldUint8},
		{Label: "Tag", Key: "tag", Type: FieldString},
		{Label: "Value", Key: "value", Type: FieldString},
	},
	"SRV": {
		{Label: "Service", Key: "service", Type: FieldString},
		{Label: "Protocol", Key: "protocol", Type: FieldString},
		{Label: "Priority", Key: "priority", Type: FieldUint16},
		{Label: "Weight", Key: "weight", Type: FieldUint16},
		{Label: "Port", Key: "port", Type: FieldUint16},
		{Label: "Target", Key: "target", Type: FieldString},
	},
	"TLSA": {
		{Label: "Port", Key: "port", Type: FieldString},
		{Label: "Protocol", Key: "protocol", Type: FieldString},
		{Label: "Usage", Key: "usage", Type: FieldUint16},
		{Label: "Selector", Key: "selector", Type: FieldUint16},
		{Label: "Matching", Key: "matching", Type: FieldUint16},
		{Label: "Association Data", Key: "associationData", Type: FieldString},
	},
	"HTTPS": {
		{Label: "Port", Key: "port", Type: FieldString},
		{Label: "Scheme", Key: "scheme", Type: FieldString},
		{Label: "SVC Priority", Key: "svcPriority", Type: FieldUint16},
		{Label: "Target Name", Key: "targetName", Type: FieldString},
		{Label: "SVC Params", Key: "svcParams", Type: FieldString},
	},
	"SVCB": {
		{Label: "Port", Key: "port", Type: FieldString},
		{Label: "Scheme", Key: "scheme", Type: FieldString},
		{Label: "SVC Priority", Key: "svcPriority", Type: FieldUint16},
		{Label: "Target Name", Key: "targetName", Type: FieldString},
		{Label: "SVC Params", Key: "svcParams", Type: FieldString},
	},
}

var RecordTypes = []string{
	"A", "AAAA", "ALIAS", "CNAME", "TXT", "NS", "PTR",
	"MX", "CAA", "SRV", "TLSA", "HTTPS", "SVCB",
}
