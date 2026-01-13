// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Speech API
//
// The OCI Speech Service harnesses the power of spoken language by allowing developers to easily convert file-based data containing human speech into highly accurate text transcriptions.
//

package aispeech

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// RealtimeMessageOut Websocket messages sent by client to the service.
type RealtimeMessageOut interface {
}

type realtimemessageout struct {
	JsonData []byte
	Event    string `json:"event"`
}

// UnmarshalJSON unmarshals json
func (m *realtimemessageout) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerrealtimemessageout realtimemessageout
	s := struct {
		Model Unmarshalerrealtimemessageout
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Event = s.Model.Event

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *realtimemessageout) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Event {
	case "SEND_FINAL_RESULT":
		mm := RealtimeMessageSendFinalResult{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Received unsupported enum value for RealtimeMessageOut: %s.", m.Event)
		return *m, nil
	}
}

func (m realtimemessageout) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m realtimemessageout) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// RealtimeMessageOutEventEnum Enum with underlying type: string
type RealtimeMessageOutEventEnum string

// Set of constants representing the allowable values for RealtimeMessageOutEventEnum
const (
	RealtimeMessageOutEventSendFinalResult RealtimeMessageOutEventEnum = "SEND_FINAL_RESULT"
)

var mappingRealtimeMessageOutEventEnum = map[string]RealtimeMessageOutEventEnum{
	"SEND_FINAL_RESULT": RealtimeMessageOutEventSendFinalResult,
}

var mappingRealtimeMessageOutEventEnumLowerCase = map[string]RealtimeMessageOutEventEnum{
	"send_final_result": RealtimeMessageOutEventSendFinalResult,
}

// GetRealtimeMessageOutEventEnumValues Enumerates the set of values for RealtimeMessageOutEventEnum
func GetRealtimeMessageOutEventEnumValues() []RealtimeMessageOutEventEnum {
	values := make([]RealtimeMessageOutEventEnum, 0)
	for _, v := range mappingRealtimeMessageOutEventEnum {
		values = append(values, v)
	}
	return values
}

// GetRealtimeMessageOutEventEnumStringValues Enumerates the set of values in String for RealtimeMessageOutEventEnum
func GetRealtimeMessageOutEventEnumStringValues() []string {
	return []string{
		"SEND_FINAL_RESULT",
	}
}

// GetMappingRealtimeMessageOutEventEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingRealtimeMessageOutEventEnum(val string) (RealtimeMessageOutEventEnum, bool) {
	enum, ok := mappingRealtimeMessageOutEventEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
