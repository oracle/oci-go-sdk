// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Container Registry API
//
// OCI Container Registry extension endpoints provide compatibility and extensions for the Docker/Open Container Initiative Distribution Registry HTTP API v2. Use together with the Artifacts and Container Images API for control‑plane repository and image management - https://docs.oracle.com/en-us/iaas/api/#/en/registry/20160918/
//

package containerregistry

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// AccessToken Access token response for OCI Container Registry extension authentication.
type AccessToken struct {

	// Access token string.
	Token *string `mandatory:"true" json:"token"`

	// Access token string (alias for token).
	AccessToken *string `mandatory:"true" json:"access_token"`

	// Scope that restricts access to resources.
	Scope *string `mandatory:"true" json:"scope"`

	// Seconds until the token expires.
	ExpiresIn *int `mandatory:"true" json:"expires_in"`
}

func (m AccessToken) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m AccessToken) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
