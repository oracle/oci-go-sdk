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

// RunSandboxCommandDetails Request body used to run a command in a sandbox.
type RunSandboxCommandDetails struct {

	// The command to execute in the sandbox runtime.
	Command *string `mandatory:"true" json:"command"`

	// Absolute directory path to use as the working directory for command execution in the sandbox.
	Cwd *string `mandatory:"false" json:"cwd"`

	// Command arguments.
	Args []string `mandatory:"false" json:"args"`

	// Environment variables to inject into the command execution.
	EnvironmentVariables []EnvironmentVariable `mandatory:"false" json:"environmentVariables"`

	// Maximum execution time for the command, in ISO 8601 duration format.
	Timeout *string `mandatory:"false" json:"timeout"`
}

func (m RunSandboxCommandDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m RunSandboxCommandDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
