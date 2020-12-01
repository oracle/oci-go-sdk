// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// InstanceAgentService API
//
// Instance Agent Service API
//

package computeinstanceagent

import (
	"github.com/oracle/oci-go-sdk/v30/common"
)

// InstanceAgentCommandContentInfo The command content.
type InstanceAgentCommandContentInfo struct {

	// The command ocid
	InstanceAgentCommandId *string `mandatory:"true" json:"instanceAgentCommandId"`

	// The OCID of the compartment the command is created in.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	Content *InstanceAgentCommandContent `mandatory:"true" json:"content"`

	// created at time of command.
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// updated time of command.
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// Whether the command has been canceled.
	IsCanceled *bool `mandatory:"false" json:"isCanceled"`

	// The last command time.
	ExecutionTimeOutInSeconds *int `mandatory:"false" json:"executionTimeOutInSeconds"`
}

func (m InstanceAgentCommandContentInfo) String() string {
	return common.PointerString(m)
}
