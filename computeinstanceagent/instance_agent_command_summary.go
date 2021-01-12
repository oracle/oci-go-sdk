// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// InstanceAgentService API
//
// Instance Agent Service API
//

package computeinstanceagent

import (
	"github.com/oracle/oci-go-sdk/v32/common"
)

// InstanceAgentCommandSummary command summary.
type InstanceAgentCommandSummary struct {

	// The command OCID
	InstanceAgentCommandId *string `mandatory:"true" json:"instanceAgentCommandId"`

	// The OCID of the compartment the command is created in.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The command creation date
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The command last updated at date.
	TimeUpdated *common.SDKTime `mandatory:"true" json:"timeUpdated"`

	// The user friendly display name of the command.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Set to true, if the command has been canceled.
	IsCanceled *bool `mandatory:"false" json:"isCanceled"`
}

func (m InstanceAgentCommandSummary) String() string {
	return common.PointerString(m)
}
