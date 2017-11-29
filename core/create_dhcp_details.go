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

func (model CreateDhcpDetails) String() string {
	return common.PointerString(model)
}

func (model *CreateDhcpDetails) UnmarshalJSON(data []byte) (e error) {
	m := struct {
		DisplayName   *string      `mandatory:"true" json:"displayName,omitempty"`
		CompartmentID *string      `mandatory:"true" json:"compartmentId,omitempty"`
		Options       []dhcpoption `mandatory:"true" json:"options,omitempty"`
		VcnID         *string      `mandatory:"true" json:"vcnId,omitempty"`
	}{}

	e = json.Unmarshal(data, &m)
	if e != nil {
		return
	}
	model.DisplayName = m.DisplayName
	model.CompartmentID = m.CompartmentID
	model.Options = make([]DhcpOption, len(m.Options))
	for i, n := range m.Options {
		nn, err := n.UnmarshalPolymorphicJSON(n.JsonData)
		if err != nil {
			return err
		}
		model.Options[i] = nn
	}
	model.VcnID = m.VcnID
	return
}
