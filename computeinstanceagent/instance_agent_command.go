// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// InstanceAgentService API
//
// Instance Agent Service API
//

package computeinstanceagent

import (
	"github.com/oracle/oci-go-sdk/v35/common"
)

// InstanceAgentCommand The command payload.
type InstanceAgentCommand struct {

	// The command OCID
	Id *string `mandatory:"true" json:"id"`

	// The OCID of the compartment the command is created in.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	Target *InstanceAgentCommandTarget `mandatory:"true" json:"target"`

	Content *InstanceAgentCommandContent `mandatory:"true" json:"content"`

	// The user friendly display name of the command.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// the time command was created at.
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// the time command was updated at.
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// Whether the command has been requested to be canceled.
	IsCanceled *bool `mandatory:"false" json:"isCanceled"`

	// Command execution time limit that the instance agent will honor when executing the command inside the instance. This timer starts when the instance agent starts the commond. Zero means no timeout.
	ExecutionTimeOutInSeconds *int `mandatory:"false" json:"executionTimeOutInSeconds"`
}

func (m InstanceAgentCommand) String() string {
	return common.PointerString(m)
}
