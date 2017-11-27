// Copyright (c) 2016, 2017, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Identity and Access Management Service API
//
// APIs for managing users, groups, compartments, and policies.
//

package identity

import (
	"github.com/oracle/oci-go-sdk/common"
	"encoding/json"
)

type UpdateIdentityProviderDetails interface {

	// The description you assign to the `IdentityProvider`. Does not have to
	// be unique, and it's changeable.
	GetDescription() *string
}

type updateidentityproviderdetails struct {
	Description *string `mandatory:"false" json:"description,omitempty"`
	Protocol    string  `json:"protocol"`
}

func (m *updateidentityproviderdetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {
	err := json.Unmarshal(data, m)
	if err != nil {
		return nil, err
	}

	switch m.Protocol {
	case "SAML2":
		mm := UpdateSaml2IdentityProviderDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		return m, nil
	}
}

func (m updateidentityproviderdetails) GetDescription() *string {
	return m.Description
}

func (model updateidentityproviderdetails) String() string {
	return common.PointerString(model)
}
