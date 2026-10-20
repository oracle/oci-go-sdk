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

// SandboxCommandOutput Captured command output details.
type SandboxCommandOutput struct {

	// Sandbox command identifier.
	CommandId *string `mandatory:"true" json:"commandId"`

	// Last command event sequence number included in the output.
	LastSeq *int64 `mandatory:"true" json:"lastSeq"`

	// Process exit code when available.
	ExitCode *int `mandatory:"false" json:"exitCode"`

	// Standard output captured from the requested event sequence.
	Stdout *string `mandatory:"false" json:"stdout"`

	// Standard error captured from the requested event sequence.
	Stderr *string `mandatory:"false" json:"stderr"`
}

func (m SandboxCommandOutput) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m SandboxCommandOutput) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
