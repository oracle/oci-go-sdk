// Copyright (c) 2016, 2018, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Key Management Service API
//
// API for managing and performing operations with keys and vaults.
//

package keymanagement

import (
	"github.com/oracle/oci-go-sdk/common"
)

// DecryptDataDetails The representation of DecryptDataDetails
type DecryptDataDetails struct {

	// The encrypted data to decrypt.
	Ciphertext *string `mandatory:"true" json:"ciphertext"`

	// The OCID of the key used to encrypt the ciphertext.
	KeyId *string `mandatory:"true" json:"keyId"`

	// Information that can be used to provide an encryption context for the
	// encrypted data. The length of the string representation of the associatedData
	// must be fewer than 4096 characters.
	AssociatedData map[string]string `mandatory:"false" json:"associatedData"`
}

func (m DecryptDataDetails) String() string {
	return common.PointerString(m)
}
