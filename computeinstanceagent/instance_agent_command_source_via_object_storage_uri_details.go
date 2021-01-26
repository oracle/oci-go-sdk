// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// InstanceAgentService API
//
// Instance Agent Service API
//

package computeinstanceagent

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/v34/common"
)

// InstanceAgentCommandSourceViaObjectStorageUriDetails Command content via uri.
type InstanceAgentCommandSourceViaObjectStorageUriDetails struct {

	// The Object Storage URL or PAR for the command.
	SourceUri *string `mandatory:"true" json:"sourceUri"`
}

func (m InstanceAgentCommandSourceViaObjectStorageUriDetails) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m InstanceAgentCommandSourceViaObjectStorageUriDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeInstanceAgentCommandSourceViaObjectStorageUriDetails InstanceAgentCommandSourceViaObjectStorageUriDetails
	s := struct {
		DiscriminatorParam string `json:"sourceType"`
		MarshalTypeInstanceAgentCommandSourceViaObjectStorageUriDetails
	}{
		"OBJECT_STORAGE_URI",
		(MarshalTypeInstanceAgentCommandSourceViaObjectStorageUriDetails)(m),
	}

	return json.Marshal(&s)
}
