// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Full Stack Disaster Recovery API
//
// Use the Full Stack Disaster Recovery (DR) API to manage disaster recovery for business applications.
// Full Stack DR is an OCI disaster recovery orchestration and management service that provides comprehensive disaster
// recovery capabilities for all layers of an application stack, including infrastructure, middleware, database,
// and application.
//

package disasterrecovery

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// RunObjectStoreScriptUserDefinedStep Run Object Store Script step details.
type RunObjectStoreScriptUserDefinedStep struct {

	// The OCID of the instance on which this script or command should be executed.
	// **For moving instances:** *runOnInstanceId* must be the OCID of the instance in the region where the
	// instance is currently present.
	// **For non-moving instances:** *runOnInstanceId* must be the OCID of the non-moving instance.
	// Example: `ocid1.instance.oc1..uniqueID`
	RunOnInstanceId *string `mandatory:"true" json:"runOnInstanceId"`

	// The region of the instance where this script or command should be executed.
	// Example: `us-ashburn-1`
	RunOnInstanceRegion *string `mandatory:"true" json:"runOnInstanceRegion"`

	ObjectStorageScriptLocation *ObjectStorageScriptLocation `mandatory:"true" json:"objectStorageScriptLocation"`
}

func (m RunObjectStoreScriptUserDefinedStep) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m RunObjectStoreScriptUserDefinedStep) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m RunObjectStoreScriptUserDefinedStep) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeRunObjectStoreScriptUserDefinedStep RunObjectStoreScriptUserDefinedStep
	s := struct {
		DiscriminatorParam string `json:"stepType"`
		MarshalTypeRunObjectStoreScriptUserDefinedStep
	}{
		"RUN_OBJECTSTORE_SCRIPT",
		(MarshalTypeRunObjectStoreScriptUserDefinedStep)(m),
	}

	return json.Marshal(&s)
}
