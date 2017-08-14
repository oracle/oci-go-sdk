// Copyright (c) 2016, 2017, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Identity and Access Management Service API
//
// APIs for managing users, groups, compartments, and policies.
//

package identity

import (
	"time"
)

//Saml2IdentityProvider A special type of [IdentityProvider](#/en/identity/20160918/IdentityProvider/) that\nsupports the SAML 2.0 protocol. For more information, see\n[Identity Providers and Federation](/Content/Identity/Concepts/federation.htm).\n

type Saml2IdentityProvider struct {

	// The OCID of the `IdentityProvider`.
	Id string `json:"id,omitempty"`

	// The OCID of the tenancy containing the `IdentityProvider`.
	CompartmentId string `json:"compartmentId,omitempty"`

	// The name you assign to the `IdentityProvider` during creation. The name\nmust be unique across all `IdentityProvider` objects in the tenancy and\ncannot be changed. This is the name federated users see when choosing\nwhich identity provider to use when signing in to the Oracle Bare Metal Cloud\nServices Console.\n
	Name string `json:"name,omitempty"`

	// The description you assign to the `IdentityProvider` during creation. Does\nnot have to be unique, and it's changeable.\n
	Description string `json:"description,omitempty"`

	// The identity provider service or product (e.g., Oracle Identity Cloud Service).\nAllowed value: `IDCS`.\n\nExample: `IDCS`\n
	ProductType string `json:"productType,omitempty"`

	// Date and time the `IdentityProvider` was created, in the format defined by RFC3339.\n\nExample: `2016-08-25T21:10:29.600Z`\n
	TimeCreated time.Time `json:"timeCreated,omitempty"`

	// The current state. After creating an `IdentityProvider`, make sure its\n`lifecycleState` changes from CREATING to ACTIVE before using it.\n
	LifecycleState string `json:"lifecycleState,omitempty"`

	// The detailed status of INACTIVE lifecycleState.
	InactiveStatus int64 `json:"inactiveStatus,omitempty"`

	// The protocol used for federation. Allowed value: `SAML2`.\n\nExample: `SAML2`\n
	Protocol string `json:"protocol,omitempty"`

	// The URL for retrieving the identity provider's metadata, which\ncontains information required for federating.\n
	MetadataUrl string `json:"metadataUrl,omitempty"`

	// The identity provider's signing certificate used by the IAM Service\nto validate the SAML2 token.\n
	SigningCertificate string `json:"signingCertificate,omitempty"`

	// The URL to redirect federated users to for authentication with the\nidentity provider.\n
	RedirectUrl string `json:"redirectUrl,omitempty"`
}
