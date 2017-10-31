// Copyright (c) 2016, 2017, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Identity and Access Management Service API
//
// APIs for managing users, groups, compartments, and policies.
//

package identity

type UpdateIdentityProviderDetails struct {

	// The protocol used for federation.
	// Example: `SAML2`
	Protocol *string `mandatory:"true" json:"protocol,omitempty"`

	// The description you assign to the `IdentityProvider`. Does not have to
	// be unique, and it's changeable.
	Description *string `mandatory:"false" json:"description,omitempty"`
}
