// Copyright (c) 2016, 2017, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Identity and Access Management Service API
//
// APIs for managing users, groups, compartments, and policies.
//

package identity

type UpdateSaml2IdentityProviderDetails struct {

	// The protocol used for federation.
	// Example: `SAML2`
	Protocol string `json:"protocol,omitempty"`

	// The description you assign to the `IdentityProvider`. Does not have to
	// be unique, and it's changeable.
	Description string `json:"description,omitempty"`

	// The URL for retrieving the identity provider's metadata,
	// which contains information required for federating.
	MetadataUrl string `json:"metadataUrl,omitempty"`

	// The XML that contains the information required for federating.
	Metadata string `json:"metadata,omitempty"`
}
