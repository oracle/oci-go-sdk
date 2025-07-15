// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
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

// RealtimeMessage Websocket messages sent between client and service.
type RealtimeMessage interface {

	// Session ID for the connected session.
	GetSessionId() *string
}

type realtimemessage struct {
	JsonData  []byte
	SessionId *string `mandatory:"false" json:"sessionId"`
	Event     string  `json:"event"`
}

// UnmarshalJSON unmarshals json
func (m *realtimemessage) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerrealtimemessage realtimemessage
	s := struct {
		Model Unmarshalerrealtimemessage
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.SessionId = s.Model.SessionId
	m.Event = s.Model.Event

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *realtimemessage) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Event {
	case "ERROR":
		mm := RealtimeMessageError{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "ACKAUDIO":
		mm := RealtimeMessageAckAudio{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "CONNECT":
		mm := RealtimeMessageConnect{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "RESULT":
		mm := RealtimeMessageResult{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Received unsupported enum value for RealtimeMessage: %s.", m.Event)
		return *m, nil
	}
}

// GetSessionId returns SessionId
func (m realtimemessage) GetSessionId() *string {
	return m.SessionId
}

func (m realtimemessage) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m realtimemessage) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// RealtimeMessageEventEnum Enum with underlying type: string
type RealtimeMessageEventEnum string

// Set of constants representing the allowable values for RealtimeMessageEventEnum
const (
	RealtimeMessageEventResult   RealtimeMessageEventEnum = "RESULT"
	RealtimeMessageEventAckaudio RealtimeMessageEventEnum = "ACKAUDIO"
	RealtimeMessageEventError    RealtimeMessageEventEnum = "ERROR"
	RealtimeMessageEventConnect  RealtimeMessageEventEnum = "CONNECT"
)

var mappingRealtimeMessageEventEnum = map[string]RealtimeMessageEventEnum{
	"RESULT":   RealtimeMessageEventResult,
	"ACKAUDIO": RealtimeMessageEventAckaudio,
	"ERROR":    RealtimeMessageEventError,
	"CONNECT":  RealtimeMessageEventConnect,
}

var mappingRealtimeMessageEventEnumLowerCase = map[string]RealtimeMessageEventEnum{
	"result":   RealtimeMessageEventResult,
	"ackaudio": RealtimeMessageEventAckaudio,
	"error":    RealtimeMessageEventError,
	"connect":  RealtimeMessageEventConnect,
}

// GetRealtimeMessageEventEnumValues Enumerates the set of values for RealtimeMessageEventEnum
func GetRealtimeMessageEventEnumValues() []RealtimeMessageEventEnum {
	values := make([]RealtimeMessageEventEnum, 0)
	for _, v := range mappingRealtimeMessageEventEnum {
		values = append(values, v)
	}
	return values
}

// GetRealtimeMessageEventEnumStringValues Enumerates the set of values in String for RealtimeMessageEventEnum
func GetRealtimeMessageEventEnumStringValues() []string {
	return []string{
		"RESULT",
		"ACKAUDIO",
		"ERROR",
		"CONNECT",
	}
}

// GetMappingRealtimeMessageEventEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingRealtimeMessageEventEnum(val string) (RealtimeMessageEventEnum, bool) {
	enum, ok := mappingRealtimeMessageEventEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
