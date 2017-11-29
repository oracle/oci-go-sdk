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

type UpdateDhcpDetails struct {

	// A user-friendly name. Does not have to be unique, and it's changeable.
	// Avoid entering confidential information.
	DisplayName *string `mandatory:"false" json:"displayName,omitempty"`

	Options []DhcpOption `mandatory:"false" json:"options,omitempty"`
}

func (model UpdateDhcpDetails) String() string {
	return common.PointerString(model)
}

func (model *UpdateDhcpDetails) UnmarshalJSON(data []byte) (e error) {
	m := struct {
		DisplayName *string      `mandatory:"true" json:"displayName,omitempty"`
		Options     []dhcpoption `mandatory:"true" json:"options,omitempty"`
	}{}

	e = json.Unmarshal(data, &m)
	if e != nil {
		return
	}
	model.DisplayName = m.DisplayName
	model.Options = make([]DhcpOption, len(m.Options))
	for i, n := range m.Options {
		nn, err := n.UnmarshalPolymorphicJSON(n.JsonData)
		if err != nil {
			return err
		}
		model.Options[i] = nn
	}
	return
}
