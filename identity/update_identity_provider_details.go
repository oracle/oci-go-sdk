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

type UpdateIdentityProviderDetails struct {

	// The protocol used for federation.
	// Example: `SAML2`
	Protocol UpdateIdentityProviderDetailsProtocolEnum `mandatory:"true" json:"protocol,omitempty"`

	// The description you assign to the `IdentityProvider`. Does not have to
	// be unique, and it's changeable.
	Description *string `mandatory:"false" json:"description,omitempty"`
}

func (model UpdateIdentityProviderDetails) String() string {
	return common.PointerString(model)
}

type UpdateIdentityProviderDetailsProtocolEnum string

const (
	UPDATE_IDENTITY_PROVIDER_DETAILS_PROTOCOL_SAML2   UpdateIdentityProviderDetailsProtocolEnum = "SAML2"
	UPDATE_IDENTITY_PROVIDER_DETAILS_PROTOCOL_UNKNOWN UpdateIdentityProviderDetailsProtocolEnum = "UNKNOWN"
)

var mapping_updateidentityproviderdetails_protocol = map[string]UpdateIdentityProviderDetailsProtocolEnum{
	"SAML2":   UPDATE_IDENTITY_PROVIDER_DETAILS_PROTOCOL_SAML2,
	"UNKNOWN": UPDATE_IDENTITY_PROVIDER_DETAILS_PROTOCOL_UNKNOWN,
}

func GetUpdateIdentityProviderDetailsProtocolEnumValues() []UpdateIdentityProviderDetailsProtocolEnum {
	values := make([]UpdateIdentityProviderDetailsProtocolEnum, 0)
	for _, v := range mapping_updateidentityproviderdetails_protocol {
		if v != UPDATE_IDENTITY_PROVIDER_DETAILS_PROTOCOL_UNKNOWN {
			values = append(values, v)
		}
	}
	return values
}
