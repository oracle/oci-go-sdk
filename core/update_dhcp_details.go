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

func (m UpdateDhcpDetails) String() string {
	return common.PointerString(m)
}

// UnmarshalJSON unmarshals from json
func (m *UpdateDhcpDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		DisplayName *string      `json:"displayName,omitempty"`
		Options     []dhcpoption `json:"options,omitempty"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	m.DisplayName = model.DisplayName
	m.Options = make([]DhcpOption, len(model.Options))
	for i, n := range model.Options {
		nn, err := n.UnmarshalPolymorphicJSON(n.JsonData)
		if err != nil {
			return err
		}
		m.Options[i] = nn
	}
	return
}
