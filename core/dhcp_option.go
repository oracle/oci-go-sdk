// Copyright (c) 2016, 2017, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Core Services API
//
// APIs for Networking Service, Compute Service, and Block Volume Service.
//

package core

import (
	"github.com/oci-go-sdk/common"
	"encoding/json"
)

// DhcpOption. A single DHCP option according to [RFC 1533](https://tools.ietf.org/html/rfc1533).
// The two options available to use are DhcpDnsOption
// and DhcpSearchDomainOption. For more
// information, see [DNS in Your Virtual Cloud Network]({{DOC_SERVER_URL}}/Content/Network/Concepts/dns.htm)
// and [DHCP Options]({{DOC_SERVER_URL}}/Content/Network/Tasks/managingDHCP.htm).
type DhcpOption interface {
}

type dhcpoption struct {
	Type string `json:"type"`
}

func (m *dhcpoption) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {
	err := json.Unmarshal(data, m)
	if err != nil {
		return nil, err
	}

	switch m.Type {
	case "DomainNameServer":
		mm := DhcpDnsOption{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "SearchDomain":
		mm := DhcpSearchDomainOption{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		return m, nil
	}
}

func (model dhcpoption) String() string {
	return common.PointerString(model)
}
