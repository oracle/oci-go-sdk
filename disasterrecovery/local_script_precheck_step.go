// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Full Stack Disaster Recovery API
//
// Use the Full Stack Disaster Recovery (FSDR) API to manage disaster recovery for business applications.
// FSDR is an OCI disaster recovery orchestration and management service that provides comprehensive disaster recovery
// capabilities for all layers of an application stack, including infrastructure, middleware, database, and application.
//

package disasterrecovery

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// LocalScriptPrecheckStep Run local script precheck step details.
type LocalScriptPrecheckStep struct {
}

func (m LocalScriptPrecheckStep) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m LocalScriptPrecheckStep) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m LocalScriptPrecheckStep) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeLocalScriptPrecheckStep LocalScriptPrecheckStep
	s := struct {
		DiscriminatorParam string `json:"stepType"`
		MarshalTypeLocalScriptPrecheckStep
	}{
		"RUN_LOCAL_SCRIPT_PRECHECK",
		(MarshalTypeLocalScriptPrecheckStep)(m),
	}

	return json.Marshal(&s)
}