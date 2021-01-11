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

// InstanceAgentCommandTarget Target instances that will receive a command.
type InstanceAgentCommandTarget struct {

	// The target instance OCID
	InstanceId *string `mandatory:"false" json:"instanceId"`
}

func (m InstanceAgentCommandTarget) String() string {
	return common.PointerString(m)
}
