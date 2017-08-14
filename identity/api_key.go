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

//ApiKey A PEM-format RSA credential for securing requests to the Oracle Bare Metal Cloud Services REST API. Also known\nas an *API signing key*. Specifically, this is the public key from the key pair. The private key remains with\nthe user calling the API. For information about generating a key pair\nin the required PEM format, see [Required Keys and OCIDs](/Content/API/Concepts/apisigningkey.htm).\n\n**Important:** This is **not** the SSH key for accessing compute instances.\n\nEach user can have a maximum of three API signing keys.\n\nFor more information about user credentials, see [User Credentials](/Content/Identity/Concepts/usercredentials.htm).\n

type ApiKey struct {

	// An Oracle-assigned identifier for the key, in this format:\nTENANCY_OCID/USER_OCID/KEY_FINGERPRINT.\n
	KeyId string `json:"keyId,omitempty"`

	// The key's value.
	KeyValue string `json:"keyValue,omitempty"`

	// The key's fingerprint (e.g., 12:34:56:78:90:ab:cd:ef:12:34:56:78:90:ab:cd:ef).
	Fingerprint string `json:"fingerprint,omitempty"`

	// The OCID of the user the key belongs to.
	UserId string `json:"userId,omitempty"`

	// Date and time the `ApiKey` object was created, in the format defined by RFC3339.\n\nExample: `2016-08-25T21:10:29.600Z`\n
	TimeCreated time.Time `json:"timeCreated,omitempty"`

	// The API key's current state. After creating an `ApiKey` object, make sure its `lifecycleState` changes from\nCREATING to ACTIVE before using it.\n
	LifecycleState string `json:"lifecycleState,omitempty"`

	// The detailed status of INACTIVE lifecycleState.
	InactiveStatus int64 `json:"inactiveStatus,omitempty"`
}
