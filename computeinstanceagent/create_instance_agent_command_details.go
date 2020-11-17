// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// InstanceAgentService API
//
// Instance Agent Service API
//

package computeinstanceagent

import (
	"github.com/oracle/oci-go-sdk/v29/common"
)

// CreateInstanceAgentCommandDetails Create Command Details
type CreateInstanceAgentCommandDetails struct {

	// The OCID of the compartment you want to create the command.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// Command execution time limit. Zero means no timeout.
	ExecutionTimeOutInSeconds *int `mandatory:"true" json:"executionTimeOutInSeconds"`

	Target *InstanceAgentCommandTarget `mandatory:"true" json:"target"`

	Content *InstanceAgentCommandContent `mandatory:"true" json:"content"`

	// A user-friendly name for the command. It does not have to be unique.
	// Avoid entering confidential information.
	// Example: `Database Backup Command`
	DisplayName *string `mandatory:"false" json:"displayName"`
}

func (m CreateInstanceAgentCommandDetails) String() string {
	return common.PointerString(m)
}
