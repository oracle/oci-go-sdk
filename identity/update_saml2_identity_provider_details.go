// Copyright (c) 2016, 2017, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Identity and Access Management Service API
//
// APIs for managing users, groups, compartments, and policies.
//

package identity

import (
	"bitbucket.aka.lgl.grungy.us/golang-sdk2/common"
)

type UpdateSaml2IdentityProviderDetails struct {

	// The protocol used for federation.
	// Example: `SAML2`
	Protocol UpdateSaml2IdentityProviderDetailsProtocolEnum `mandatory:"true" json:"protocol,omitempty"`

	// The description you assign to the `IdentityProvider`. Does not have to
	// be unique, and it's changeable.
	Description *string `mandatory:"false" json:"description,omitempty"`

	// The URL for retrieving the identity provider's metadata,
	// which contains information required for federating.
	MetadataURL *string `mandatory:"false" json:"metadataUrl,omitempty"`

	// The XML that contains the information required for federating.
	Metadata *string `mandatory:"false" json:"metadata,omitempty"`
}

func (model UpdateSaml2IdentityProviderDetails) String() string {
	return common.PointerString(model)
}

type UpdateSaml2IdentityProviderDetailsProtocolEnum string
type UpdateSaml2IdentityProviderDetailsProtocol struct{}

const (
	UPDATE_SAML2_IDENTITY_PROVIDER_DETAILS_PROTOCOL_SAML2   UpdateSaml2IdentityProviderDetailsProtocolEnum = "SAML2"
	UPDATE_SAML2_IDENTITY_PROVIDER_DETAILS_PROTOCOL_UNKNOWN UpdateSaml2IdentityProviderDetailsProtocolEnum = "UNKNOWN"
)

var mapping_updatesaml2identityproviderdetails_protocol = map[string]UpdateSaml2IdentityProviderDetailsProtocolEnum{
	"SAML2":   UPDATE_SAML2_IDENTITY_PROVIDER_DETAILS_PROTOCOL_SAML2,
	"UNKNOWN": UPDATE_SAML2_IDENTITY_PROVIDER_DETAILS_PROTOCOL_UNKNOWN,
}

func (receiver UpdateSaml2IdentityProviderDetailsProtocol) Values() []UpdateSaml2IdentityProviderDetailsProtocolEnum {
	values := make([]UpdateSaml2IdentityProviderDetailsProtocolEnum, 0)
	for _, v := range mapping_updatesaml2identityproviderdetails_protocol {
		if v != UPDATE_SAML2_IDENTITY_PROVIDER_DETAILS_PROTOCOL_UNKNOWN {
			values = append(values, v)
		}
	}
	return values
}

func (receiver UpdateSaml2IdentityProviderDetailsProtocol) IsValid(toBeChecked string) bool {
	for _, v := range receiver.Values() {
		if UpdateSaml2IdentityProviderDetailsProtocolEnum(toBeChecked) == v {
			return true
		}
	}
	return false
}

func (receiver UpdateSaml2IdentityProviderDetailsProtocol) From(toBeConverted string) UpdateSaml2IdentityProviderDetailsProtocolEnum {
	if val, ok := mapping_updatesaml2identityproviderdetails_protocol[toBeConverted]; ok {
		return val
	}
	return UPDATE_SAML2_IDENTITY_PROVIDER_DETAILS_PROTOCOL_UNKNOWN
}
