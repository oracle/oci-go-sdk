// Copyright (c) 2016, 2017, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Core Services API
//
// APIs for Networking Service, Compute Service, and Block Volume Service.
//

package core

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/common"
)

// CreateDhcpDetails The representation of CreateDhcpDetails
type CreateDhcpDetails struct {

	// The OCID of the compartment to contain the set of DHCP options.
	CompartmentID *string `mandatory:"true" json:"compartmentId,omitempty"`

	// A set of DHCP options.
	Options []DhcpOption `mandatory:"true" json:"options,omitempty"`

	// The OCID of the VCN the set of DHCP options belongs to.
	VcnID *string `mandatory:"true" json:"vcnId,omitempty"`

	// A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.
	DisplayName *string `mandatory:"false" json:"displayName,omitempty"`
}

func (m CreateDhcpDetails) String() string {
	return common.PointerString(m)
}

// UnmarshalJSON unmarshals from json
func (m *CreateDhcpDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		DisplayName   *string      `json:"displayName,omitempty"`
		CompartmentID *string      `json:"compartmentId,omitempty"`
		Options       []dhcpoption `json:"options,omitempty"`
		VcnID         *string      `json:"vcnId,omitempty"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	m.DisplayName = model.DisplayName
	m.CompartmentID = model.CompartmentID
	m.Options = make([]DhcpOption, len(model.Options))
	for i, n := range model.Options {
		nn, err := n.UnmarshalPolymorphicJSON(n.JsonData)
		if err != nil {
			return err
		}
		m.Options[i] = nn
	}
	m.VcnID = model.VcnID
	return
}
