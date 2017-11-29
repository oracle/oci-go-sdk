// Copyright (c) 2016, 2017, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Identity and Access Management Service API
//
// APIs for managing users, groups, compartments, and policies.
//

package identity

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/common"
)

type UpdateIdentityProviderDetails interface {

	// The description you assign to the `IdentityProvider`. Does not have to
	// be unique, and it's changeable.
	GetDescription() *string
}

type updateidentityproviderdetails struct {
	JsonData    []byte
	Description *string `mandatory:"false" json:"description,omitempty"`
	Protocol    string  `json:"protocol"`
}

func (m *updateidentityproviderdetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerupdateidentityproviderdetails updateidentityproviderdetails
	s := struct {
		Model Unmarshalerupdateidentityproviderdetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Description = s.Model.Description
	m.Protocol = s.Model.Protocol

	return err
}

func (m *updateidentityproviderdetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {
	var err error
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
