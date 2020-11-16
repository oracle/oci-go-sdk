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
	"github.com/oracle/oci-go-sdk/v29/common"
)

// InstanceAgentCommandExecutionOutputViaObjectStorageUriDetails command execution output via uri.
type InstanceAgentCommandExecutionOutputViaObjectStorageUriDetails struct {

	// command exit code.
	ExitCode *int `mandatory:"true" json:"exitCode"`

	// The Object Storage URL or PAR for the command output.
	OutputUri *string `mandatory:"true" json:"outputUri"`

	// optional status message that agent's can populate for additional troubleshooting.
	Message *string `mandatory:"false" json:"message"`
}

//GetExitCode returns ExitCode
func (m InstanceAgentCommandExecutionOutputViaObjectStorageUriDetails) GetExitCode() *int {
	return m.ExitCode
}

//GetMessage returns Message
func (m InstanceAgentCommandExecutionOutputViaObjectStorageUriDetails) GetMessage() *string {
	return m.Message
}

func (m InstanceAgentCommandExecutionOutputViaObjectStorageUriDetails) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m InstanceAgentCommandExecutionOutputViaObjectStorageUriDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeInstanceAgentCommandExecutionOutputViaObjectStorageUriDetails InstanceAgentCommandExecutionOutputViaObjectStorageUriDetails
	s := struct {
		DiscriminatorParam string `json:"outputType"`
		MarshalTypeInstanceAgentCommandExecutionOutputViaObjectStorageUriDetails
	}{
		"OBJECT_STORAGE_URI",
		(MarshalTypeInstanceAgentCommandExecutionOutputViaObjectStorageUriDetails)(m),
	}

	return json.Marshal(&s)
}
