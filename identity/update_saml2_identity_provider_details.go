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

type UpdateSaml2IdentityProviderDetails struct {

	// The description you assign to the `IdentityProvider`. Does not have to
	// be unique, and it's changeable.
	Description *string `mandatory:"false" json:"description,omitempty"`

	// The URL for retrieving the identity provider's metadata,
	// which contains information required for federating.
	MetadataURL *string `mandatory:"false" json:"metadataUrl,omitempty"`

	// The XML that contains the information required for federating.
	Metadata *string `mandatory:"false" json:"metadata,omitempty"`
}

//GetDescription returns Description
func (m UpdateSaml2IdentityProviderDetails) GetDescription() *string {
	return m.Description
}

func (m UpdateSaml2IdentityProviderDetails) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m UpdateSaml2IdentityProviderDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeUpdateSaml2IdentityProviderDetails UpdateSaml2IdentityProviderDetails
	s := struct {
		DiscriminatorParam string `json:"protocol"`
		MarshalTypeUpdateSaml2IdentityProviderDetails
	}{
		"SAML2",
		(MarshalTypeUpdateSaml2IdentityProviderDetails)(m),
	}

	return json.Marshal(&s)
}
