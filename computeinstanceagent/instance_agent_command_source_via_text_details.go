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

// InstanceAgentCommandSourceViaTextDetails Command content via text.
type InstanceAgentCommandSourceViaTextDetails struct {

	// The plain text command.
	Text *string `mandatory:"true" json:"text"`

	// Sha256 checksum value of the text content
	TextSha256 *string `mandatory:"false" json:"textSha256"`
}

func (m InstanceAgentCommandSourceViaTextDetails) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m InstanceAgentCommandSourceViaTextDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeInstanceAgentCommandSourceViaTextDetails InstanceAgentCommandSourceViaTextDetails
	s := struct {
		DiscriminatorParam string `json:"sourceType"`
		MarshalTypeInstanceAgentCommandSourceViaTextDetails
	}{
		"TEXT",
		(MarshalTypeInstanceAgentCommandSourceViaTextDetails)(m),
	}

	return json.Marshal(&s)
}
