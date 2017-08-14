// Copyright (c) 2016, 2017, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Identity and Access Management Service API
//
// APIs for managing users, groups, compartments, and policies.
//

package identity


type CreateSaml2IdentityProviderDetails struct {

    // The OCID of your tenancy.
    CompartmentId string `json:"compartmentId,omitempty"`

    // The name you assign to the `IdentityProvider` during creation.\nThe name must be unique across all `IdentityProvider` objects in the\ntenancy and cannot be changed.\n
    Name string `json:"name,omitempty"`

    // The description you assign to the `IdentityProvider` during creation.\nDoes not have to be unique, and it's changeable.\n
    Description string `json:"description,omitempty"`

    // The identity provider service or product (e.g., Oracle Identity Cloud Service).\n\nExample: `IDCS`\n
    ProductType string `json:"productType,omitempty"`

    // The protocol used for federation.\n\nExample: `SAML2`\n
    Protocol string `json:"protocol,omitempty"`

    // The URL for retrieving the identity provider's metadata,\nwhich contains information required for federating.\n
    MetadataUrl string `json:"metadataUrl,omitempty"`

    // The XML that contains the information required for federating.\n
    Metadata string `json:"metadata,omitempty"`
}
