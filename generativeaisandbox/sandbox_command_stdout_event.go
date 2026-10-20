// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// GenAI Sandbox API
//
// API for managing sandbox resources and exchanging sandbox access tokens.
//

package generativeaisandbox

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// SandboxCommandStdoutEvent Event containing a standard output chunk.
type SandboxCommandStdoutEvent struct {

	// Monotonically increasing command event sequence number.
	Seq *int64 `mandatory:"true" json:"seq"`

	// The date and time the command event was emitted, in RFC 3339 format.
	Timestamp *common.SDKTime `mandatory:"true" json:"timestamp"`

	// Sandbox command identifier.
	CommandId *string `mandatory:"true" json:"commandId"`

	// Event payload.
	Data *string `mandatory:"false" json:"data"`
}

// GetSeq returns Seq
func (m SandboxCommandStdoutEvent) GetSeq() *int64 {
	return m.Seq
}

// GetTimestamp returns Timestamp
func (m SandboxCommandStdoutEvent) GetTimestamp() *common.SDKTime {
	return m.Timestamp
}

// GetCommandId returns CommandId
func (m SandboxCommandStdoutEvent) GetCommandId() *string {
	return m.CommandId
}

// GetData returns Data
func (m SandboxCommandStdoutEvent) GetData() *string {
	return m.Data
}

func (m SandboxCommandStdoutEvent) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m SandboxCommandStdoutEvent) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m SandboxCommandStdoutEvent) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeSandboxCommandStdoutEvent SandboxCommandStdoutEvent
	s := struct {
		DiscriminatorParam string `json:"event"`
		MarshalTypeSandboxCommandStdoutEvent
	}{
		"STDOUT",
		(MarshalTypeSandboxCommandStdoutEvent)(m),
	}

	return json.Marshal(&s)
}
