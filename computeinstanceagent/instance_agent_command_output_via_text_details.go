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
	"github.com/oracle/oci-go-sdk/v27/common"
)

// InstanceAgentCommandOutputViaTextDetails Command output via text.
type InstanceAgentCommandOutputViaTextDetails struct {
}

func (m InstanceAgentCommandOutputViaTextDetails) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m InstanceAgentCommandOutputViaTextDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeInstanceAgentCommandOutputViaTextDetails InstanceAgentCommandOutputViaTextDetails
	s := struct {
		DiscriminatorParam string `json:"outputType"`
		MarshalTypeInstanceAgentCommandOutputViaTextDetails
	}{
		"TEXT",
		(MarshalTypeInstanceAgentCommandOutputViaTextDetails)(m),
	}

	return json.Marshal(&s)
}
