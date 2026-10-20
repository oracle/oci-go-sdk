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

// SandboxCommandSummary Sandbox command resource summary.
type SandboxCommandSummary struct {

	// Sandbox command identifier.
	Id *string `mandatory:"true" json:"id"`

	// Sandbox identifier that this command runs against.
	SandboxId *string `mandatory:"true" json:"sandboxId"`

	// The command executed in the sandbox runtime.
	Command *string `mandatory:"true" json:"command"`

	// Absolute directory path used as the working directory for command execution in the sandbox.
	Cwd *string `mandatory:"false" json:"cwd"`

	// Command arguments.
	Args []string `mandatory:"false" json:"args"`

	// Process exit code when available.
	ExitCode *int `mandatory:"false" json:"exitCode"`

	// The date and time the command was created, in RFC 3339 format.
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// The date and time the command was last updated, in RFC 3339 format.
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// The date and time command execution started, in RFC 3339 format.
	TimeStarted *common.SDKTime `mandatory:"false" json:"timeStarted"`

	// The date and time command execution finished, in RFC 3339 format.
	TimeFinished *common.SDKTime `mandatory:"false" json:"timeFinished"`
}

func (m SandboxCommandSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m SandboxCommandSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
