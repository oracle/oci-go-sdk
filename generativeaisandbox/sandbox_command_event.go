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

// SandboxCommandEvent Base command event emitted while a sandbox command runs.
type SandboxCommandEvent interface {

	// Monotonically increasing command event sequence number.
	GetSeq() *int64

	// The date and time the command event was emitted, in RFC 3339 format.
	GetTimestamp() *common.SDKTime

	// Sandbox command identifier.
	GetCommandId() *string

	// Event payload.
	GetData() *string
}

type sandboxcommandevent struct {
	JsonData  []byte
	Data      *string         `mandatory:"false" json:"data"`
	Seq       *int64          `mandatory:"true" json:"seq"`
	Timestamp *common.SDKTime `mandatory:"true" json:"timestamp"`
	CommandId *string         `mandatory:"true" json:"commandId"`
	Event     string          `json:"event"`
}

// UnmarshalJSON unmarshals json
func (m *sandboxcommandevent) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalersandboxcommandevent sandboxcommandevent
	s := struct {
		Model Unmarshalersandboxcommandevent
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Seq = s.Model.Seq
	m.Timestamp = s.Model.Timestamp
	m.CommandId = s.Model.CommandId
	m.Data = s.Model.Data
	m.Event = s.Model.Event

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *sandboxcommandevent) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Event {
	case "STDERR":
		mm := SandboxCommandStderrEvent{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "FINISHED":
		mm := SandboxCommandFinishedEvent{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "STARTED":
		mm := SandboxCommandStartedEvent{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "STDOUT":
		mm := SandboxCommandStdoutEvent{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Received unsupported enum value for SandboxCommandEvent: %s.", m.Event)
		return *m, nil
	}
}

// GetData returns Data
func (m sandboxcommandevent) GetData() *string {
	return m.Data
}

// GetSeq returns Seq
func (m sandboxcommandevent) GetSeq() *int64 {
	return m.Seq
}

// GetTimestamp returns Timestamp
func (m sandboxcommandevent) GetTimestamp() *common.SDKTime {
	return m.Timestamp
}

// GetCommandId returns CommandId
func (m sandboxcommandevent) GetCommandId() *string {
	return m.CommandId
}

func (m sandboxcommandevent) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m sandboxcommandevent) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// SandboxCommandEventEventEnum Enum with underlying type: string
type SandboxCommandEventEventEnum string

// Set of constants representing the allowable values for SandboxCommandEventEventEnum
const (
	SandboxCommandEventEventStarted  SandboxCommandEventEventEnum = "STARTED"
	SandboxCommandEventEventStdout   SandboxCommandEventEventEnum = "STDOUT"
	SandboxCommandEventEventStderr   SandboxCommandEventEventEnum = "STDERR"
	SandboxCommandEventEventFinished SandboxCommandEventEventEnum = "FINISHED"
)

var mappingSandboxCommandEventEventEnum = map[string]SandboxCommandEventEventEnum{
	"STARTED":  SandboxCommandEventEventStarted,
	"STDOUT":   SandboxCommandEventEventStdout,
	"STDERR":   SandboxCommandEventEventStderr,
	"FINISHED": SandboxCommandEventEventFinished,
}

var mappingSandboxCommandEventEventEnumLowerCase = map[string]SandboxCommandEventEventEnum{
	"started":  SandboxCommandEventEventStarted,
	"stdout":   SandboxCommandEventEventStdout,
	"stderr":   SandboxCommandEventEventStderr,
	"finished": SandboxCommandEventEventFinished,
}

// GetSandboxCommandEventEventEnumValues Enumerates the set of values for SandboxCommandEventEventEnum
func GetSandboxCommandEventEventEnumValues() []SandboxCommandEventEventEnum {
	values := make([]SandboxCommandEventEventEnum, 0)
	for _, v := range mappingSandboxCommandEventEventEnum {
		values = append(values, v)
	}
	return values
}

// GetSandboxCommandEventEventEnumStringValues Enumerates the set of values in String for SandboxCommandEventEventEnum
func GetSandboxCommandEventEventEnumStringValues() []string {
	return []string{
		"STARTED",
		"STDOUT",
		"STDERR",
		"FINISHED",
	}
}

// GetMappingSandboxCommandEventEventEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSandboxCommandEventEventEnum(val string) (SandboxCommandEventEventEnum, bool) {
	enum, ok := mappingSandboxCommandEventEventEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
