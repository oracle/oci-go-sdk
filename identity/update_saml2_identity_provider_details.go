// Copyright (c) 2016, 2017, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Identity and Access Management Service API
//
// APIs for managing users, groups, compartments, and policies.
//

package identity


type UpdateSaml2IdentityProviderDetails struct {

    // The protocol used for federation.\n\nExample: `SAML2`\n
    Protocol string `json:"protocol,omitempty"`

    // The description you assign to the `IdentityProvider`. Does not have to\nbe unique, and it's changeable.\n
    Description string `json:"description,omitempty"`

    // The URL for retrieving the identity provider's metadata,\nwhich contains information required for federating.\n
    MetadataUrl string `json:"metadataUrl,omitempty"`

    // The XML that contains the information required for federating.\n
    Metadata string `json:"metadata,omitempty"`
}
