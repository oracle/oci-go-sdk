// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Secrets
//
// API for retrieving secrets from vaults.
//

package secrets

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/common"
)

// Base64SecretBundleContentDetails The contents of the secret.
type Base64SecretBundleContentDetails struct {

	// The base64-encoded content of the secret.
	Content *string `mandatory:"false" json:"content"`
}

func (m Base64SecretBundleContentDetails) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m Base64SecretBundleContentDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeBase64SecretBundleContentDetails Base64SecretBundleContentDetails
	s := struct {
		DiscriminatorParam string `json:"contentType"`
		MarshalTypeBase64SecretBundleContentDetails
	}{
		"BASE64",
		(MarshalTypeBase64SecretBundleContentDetails)(m),
	}

	return json.Marshal(&s)
}
