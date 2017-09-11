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

// ApiKey. A PEM-format RSA credential for securing requests to the Oracle Bare Metal Cloud Services REST API. Also known
// as an *API signing key*. Specifically, this is the public key from the key pair. The private key remains with
// the user calling the API. For information about generating a key pair
// in the required PEM format, see [Required Keys and OCIDs](/Content/API/Concepts/apisigningkey.htm).
// **Important:** This is **not** the SSH key for accessing compute instances.
// Each user can have a maximum of three API signing keys.
// For more information about user credentials, see [User Credentials](/Content/Identity/Concepts/usercredentials.htm).
type ApiKey struct {

	// An Oracle-assigned identifier for the key, in this format:
	// TENANCY_OCID/USER_OCID/KEY_FINGERPRINT.
	KeyID string `mandatory:"false" json:"keyId,omitempty"`

	// The key's value.
	KeyValue string `mandatory:"false" json:"keyValue,omitempty"`

	// The key's fingerprint (e.g., 12:34:56:78:90:ab:cd:ef:12:34:56:78:90:ab:cd:ef).
	Fingerprint string `mandatory:"false" json:"fingerprint,omitempty"`

	// The OCID of the user the key belongs to.
	UserID string `mandatory:"false" json:"userId,omitempty"`

	// Date and time the `ApiKey` object was created, in the format defined by RFC3339.
	// Example: `2016-08-25T21:10:29.600Z`
	TimeCreated time.Time `mandatory:"false" json:"timeCreated,omitempty"`

	// The API key's current state. After creating an `ApiKey` object, make sure its `lifecycleState` changes from
	// CREATING to ACTIVE before using it.
	LifecycleState string `mandatory:"false" json:"lifecycleState,omitempty"`

	// The detailed status of INACTIVE lifecycleState.
	InactiveStatus int64 `mandatory:"false" json:"inactiveStatus,omitempty"`
}
