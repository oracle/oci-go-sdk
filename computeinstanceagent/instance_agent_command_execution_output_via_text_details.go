// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// InstanceAgentService API
//
// Instance Agent Service API
//

package computeinstanceagent

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/v28/common"
)

// InstanceAgentCommandExecutionOutputViaTextDetails command execution output via text.
type InstanceAgentCommandExecutionOutputViaTextDetails struct {

	// command exit code.
	ExitCode *int `mandatory:"true" json:"exitCode"`

	// The command response output.
	Text *string `mandatory:"true" json:"text"`

	// optional status message that agent's can populate for additional troubleshooting.
	Message *string `mandatory:"false" json:"message"`

	// Sha256 checksum value of the text content
	TextSha256 *string `mandatory:"false" json:"textSha256"`
}

//GetExitCode returns ExitCode
func (m InstanceAgentCommandExecutionOutputViaTextDetails) GetExitCode() *int {
	return m.ExitCode
}

//GetMessage returns Message
func (m InstanceAgentCommandExecutionOutputViaTextDetails) GetMessage() *string {
	return m.Message
}

func (m InstanceAgentCommandExecutionOutputViaTextDetails) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m InstanceAgentCommandExecutionOutputViaTextDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeInstanceAgentCommandExecutionOutputViaTextDetails InstanceAgentCommandExecutionOutputViaTextDetails
	s := struct {
		DiscriminatorParam string `json:"outputType"`
		MarshalTypeInstanceAgentCommandExecutionOutputViaTextDetails
	}{
		"TEXT",
		(MarshalTypeInstanceAgentCommandExecutionOutputViaTextDetails)(m),
	}

	return json.Marshal(&s)
}
