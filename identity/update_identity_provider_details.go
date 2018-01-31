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

// UpdateIdentityProviderDetails The representation of UpdateIdentityProviderDetails
type UpdateIdentityProviderDetails interface {

	// The description you assign to the `IdentityProvider`. Does not have to
	// be unique, and it's changeable.
	GetDescription() *string
}

type updateidentityproviderdetails struct {
	JsonData    []byte
	Description *string `mandatory:"false" json:"description"`
	Protocol    string  `json:"protocol"`
}

// UnmarshalJSON unmarshals json
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

// UnmarshalPolymorphicJSON unmarshals polymorphic json
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

//GetDescription returns Description
func (m updateidentityproviderdetails) GetDescription() *string {
	return m.Description
}

func (m updateidentityproviderdetails) String() string {
	return common.PointerString(m)
}

//listupdateidentityproviderdetails allows to unmarshal list of polymorphic UpdateIdentityProviderDetails
type listupdateidentityproviderdetails []UpdateIdentityProviderDetails

//UnmarshalPolymorphicJSON unmarshals polymorphic json list of items
func (m *listupdateidentityproviderdetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {
	type listMarshalHelper []updateidentityproviderdetails
	n := make(listMarshalHelper, 0)

	json.Unmarshal(data, &n)
	res := make([]UpdateIdentityProviderDetails, len(n))
	for i, v := range n {
		nn, err := v.UnmarshalPolymorphicJSON(v.JsonData)
		if err != nil {
			return nil, err
		}
		res[i] = nn.(UpdateIdentityProviderDetails)
	}
	return res, nil
}
