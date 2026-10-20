// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// GenAI Sandbox API
//
// API for managing sandbox resources and exchanging sandbox access tokens.
//

package generativeaisandbox

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// PublishedPort Published port key/value pair. Requests to the published URL may supply
// X-App-Token with the raw application token, without a Bearer prefix.
// After the sandbox's access policy succeeds, the proxy builds the upstream
// Authorization header as Bearer followed by the unchanged token and removes
// X-App-Token. This applies to HTTP, SSE,
// and WebSocket opening handshakes, and does not replace required OCI authentication.
// Exactly one nonempty printable ASCII token of at most 8192 bytes is accepted;
// whitespace, commas, control characters, and duplicate values are
// rejected with HTTP 400. Both header names must be treated as sensitive credentials.
type PublishedPort struct {

	// Key string.
	Key *string `mandatory:"true" json:"key"`

	// Value string.
	Value *string `mandatory:"true" json:"value"`
}

func (m PublishedPort) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m PublishedPort) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
