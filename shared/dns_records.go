package shared

import (
	"fmt"
	"strconv"

	"github.com/espcaa/spaceship-go"
)

func GetFieldValue(record spaceship.DNSRecord, key string) string {
	switch r := record.(type) {
	case spaceship.ARecord:
		switch key {
		case "name":
			return r.Name
		case "ttl":
			return strconv.Itoa(r.TTL)
		case "address":
			return r.Adress
		}
	case spaceship.AAAARecord:
		switch key {
		case "name":
			return r.Name
		case "ttl":
			return strconv.Itoa(r.TTL)
		case "address":
			return r.Adress
		}
	case spaceship.AliasRecord:
		switch key {
		case "name":
			return r.Name
		case "ttl":
			return strconv.Itoa(r.TTL)
		case "aliasTarget":
			return r.AliasTarget
		}
	case spaceship.CNAMERecord:
		switch key {
		case "name":
			return r.Name
		case "ttl":
			return strconv.Itoa(r.TTL)
		case "cname":
			return r.CNAME
		}
	case spaceship.TXTRecord:
		switch key {
		case "name":
			return r.Name
		case "ttl":
			return strconv.Itoa(r.TTL)
		case "value":
			return r.Value
		}
	case spaceship.NSRecord:
		switch key {
		case "name":
			return r.Name
		case "ttl":
			return strconv.Itoa(r.TTL)
		case "nameserver":
			return r.Nameserver
		}
	case spaceship.PTRRecord:
		switch key {
		case "name":
			return r.Name
		case "ttl":
			return strconv.Itoa(r.TTL)
		case "pointer":
			return r.Pointer
		}
	case spaceship.MXRecord:
		switch key {
		case "name":
			return r.Name
		case "ttl":
			return strconv.Itoa(r.TTL)
		case "exchange":
			return r.Exchange
		case "preference":
			return fmt.Sprintf("%d", r.Preference)
		}
	case spaceship.CAARecord:
		switch key {
		case "name":
			return r.Name
		case "ttl":
			return strconv.Itoa(r.TTL)
		case "flag":
			return fmt.Sprintf("%d", r.Flag)
		case "tag":
			return r.Tag
		case "value":
			return r.Value
		}
	case spaceship.SRVRecord:
		switch key {
		case "name":
			return r.Name
		case "ttl":
			return strconv.Itoa(r.TTL)
		case "service":
			return r.Service
		case "protocol":
			return r.Protocol
		case "priority":
			return fmt.Sprintf("%d", r.Priority)
		case "weight":
			return fmt.Sprintf("%d", r.Weight)
		case "port":
			return fmt.Sprintf("%d", r.Port)
		case "target":
			return r.Target
		}
	case spaceship.TLSARecord:
		switch key {
		case "name":
			return r.Name
		case "ttl":
			return strconv.Itoa(r.TTL)
		case "port":
			return r.Port
		case "protocol":
			return r.Protocol
		case "usage":
			return fmt.Sprintf("%d", r.Usage)
		case "selector":
			return fmt.Sprintf("%d", r.Selector)
		case "matching":
			return fmt.Sprintf("%d", r.Matching)
		case "associationData":
			return r.AssociationData
		}
	case spaceship.HTTPSRecord:
		switch key {
		case "name":
			return r.Name
		case "ttl":
			return strconv.Itoa(r.TTL)
		case "port":
			return r.Port
		case "scheme":
			return r.Scheme
		case "svcPriority":
			return fmt.Sprintf("%d", r.SvcPriority)
		case "targetName":
			return r.TargetName
		case "svcParams":
			return r.SvcParams
		}
	case spaceship.SVCBRecord:
		switch key {
		case "name":
			return r.Name
		case "ttl":
			return strconv.Itoa(r.TTL)
		case "port":
			return r.Port
		case "scheme":
			return r.Scheme
		case "svcPriority":
			return fmt.Sprintf("%d", r.SvcPriority)
		case "targetName":
			return r.TargetName
		case "svcParams":
			return r.SvcParams
		}
	}
	return ""
}

func BuildRecord(recordType string, name string, ttl int, group spaceship.DNSRecordGroup, fields map[string]string) (spaceship.DNSRecord, error) {
	switch recordType {
	case "A":
		return spaceship.ARecord{
			Name:   name,
			TTL:    ttl,
			Group:  group,
			Adress: fields["address"],
		}, nil
	case "AAAA":
		return spaceship.AAAARecord{
			Name:   name,
			TTL:    ttl,
			Group:  group,
			Adress: fields["address"],
		}, nil
	case "ALIAS":
		return spaceship.AliasRecord{
			Name:        name,
			TTL:         ttl,
			Group:       group,
			AliasTarget: fields["aliasTarget"],
		}, nil
	case "CNAME":
		return spaceship.CNAMERecord{
			Name:  name,
			TTL:   ttl,
			Group: group,
			CNAME: fields["cname"],
		}, nil
	case "TXT":
		return spaceship.TXTRecord{
			Name:  name,
			TTL:   ttl,
			Group: group,
			Value: fields["value"],
		}, nil
	case "NS":
		return spaceship.NSRecord{
			Name:       name,
			TTL:        ttl,
			Group:      group,
			Nameserver: fields["nameserver"],
		}, nil
	case "PTR":
		return spaceship.PTRRecord{
			Name:    name,
			TTL:     ttl,
			Group:   group,
			Pointer: fields["pointer"],
		}, nil
	case "MX":
		preference, err := parseUint16(fields, "preference")
		if err != nil {
			return nil, err
		}
		return spaceship.MXRecord{
			Name:       name,
			TTL:        ttl,
			Group:      group,
			Exchange:   fields["exchange"],
			Preference: preference,
		}, nil
	case "CAA":
		flag, err := parseUint8(fields, "flag")
		if err != nil {
			return nil, err
		}
		return spaceship.CAARecord{
			Name:  name,
			TTL:   ttl,
			Group: group,
			Flag:  flag,
			Tag:   fields["tag"],
			Value: fields["value"],
		}, nil
	case "SRV":
		priority, err := parseUint16(fields, "priority")
		if err != nil {
			return nil, err
		}
		weight, err := parseUint16(fields, "weight")
		if err != nil {
			return nil, err
		}
		port, err := parseUint16(fields, "port")
		if err != nil {
			return nil, err
		}
		return spaceship.SRVRecord{
			Name:     name,
			TTL:      ttl,
			Group:    group,
			Service:  fields["service"],
			Protocol: fields["protocol"],
			Priority: priority,
			Weight:   weight,
			Port:     port,
			Target:   fields["target"],
		}, nil
	case "TLSA":
		usage, err := parseUint16(fields, "usage")
		if err != nil {
			return nil, err
		}
		selector, err := parseUint16(fields, "selector")
		if err != nil {
			return nil, err
		}
		matching, err := parseUint16(fields, "matching")
		if err != nil {
			return nil, err
		}
		return spaceship.TLSARecord{
			Name:            name,
			TTL:             ttl,
			Group:           group,
			Port:            fields["port"],
			Protocol:        fields["protocol"],
			Usage:           usage,
			Selector:        selector,
			Matching:        matching,
			AssociationData: fields["associationData"],
		}, nil
	case "HTTPS":
		svcPriority, err := parseUint16(fields, "svcPriority")
		if err != nil {
			return nil, err
		}
		return spaceship.HTTPSRecord{
			Name:        name,
			TTL:         ttl,
			Group:       group,
			Port:        fields["port"],
			Scheme:      fields["scheme"],
			SvcPriority: svcPriority,
			TargetName:  fields["targetName"],
			SvcParams:   fields["svcParams"],
		}, nil
	case "SVCB":
		svcPriority, err := parseUint16(fields, "svcPriority")
		if err != nil {
			return nil, err
		}
		return spaceship.SVCBRecord{
			Name:        name,
			TTL:         ttl,
			Group:       group,
			Port:        fields["port"],
			Scheme:      fields["scheme"],
			SvcPriority: svcPriority,
			TargetName:  fields["targetName"],
			SvcParams:   fields["svcParams"],
		}, nil
	}
	return nil, fmt.Errorf("unknown record type: %s", recordType)
}

func parseUint8(fields map[string]string, key string) (uint8, error) {
	val, err := strconv.ParseUint(fields[key], 10, 8)
	if err != nil {
		return 0, fmt.Errorf("invalid %s: %w", key, err)
	}
	return uint8(val), nil
}

func parseUint16(fields map[string]string, key string) (uint16, error) {
	val, err := strconv.ParseUint(fields[key], 10, 16)
	if err != nil {
		return 0, fmt.Errorf("invalid %s: %w", key, err)
	}
	return uint16(val), nil
}
