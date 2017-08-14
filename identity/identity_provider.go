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


//IdentityProvider The resulting base object when you add an identity provider to your tenancy. A\n[Saml2IdentityProvider](#/en/identity/20160918/Saml2IdentityProvider/)\nis a specific type of `IdentityProvider` that supports the SAML 2.0 protocol. Each\n`IdentityProvider` object has its own OCID. For more information, see\n[Identity Providers and Federation](/Content/Identity/Concepts/federation.htm).\n\nTo use any of the API operations, you must be authorized in an IAM policy. If you're not authorized,\ntalk to an administrator. If you're an administrator who needs to write policies to give users access,\nsee [Getting Started with Policies](/Content/Identity/Concepts/policygetstarted.htm).\n

type IdentityProvider struct {

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
}
