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

// RealtimeMessageAuthentication The payload for authentication.
type RealtimeMessageAuthentication interface {

	// Compartment ID to be used for authentication/authorization.
	GetCompartmentId() *string
}

type realtimemessageauthentication struct {
	JsonData           []byte
	CompartmentId      *string `mandatory:"true" json:"compartmentId"`
	AuthenticationType string  `json:"authenticationType"`
}

// UnmarshalJSON unmarshals json
func (m *realtimemessageauthentication) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerrealtimemessageauthentication realtimemessageauthentication
	s := struct {
		Model Unmarshalerrealtimemessageauthentication
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.CompartmentId = s.Model.CompartmentId
	m.AuthenticationType = s.Model.AuthenticationType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *realtimemessageauthentication) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.AuthenticationType {
	case "TOKEN":
		mm := RealtimeMessageAuthenticationToken{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "CREDENTIALS":
		mm := RealtimeMessageAuthenticationCredentials{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Received unsupported enum value for RealtimeMessageAuthentication: %s.", m.AuthenticationType)
		return *m, nil
	}
}

// GetCompartmentId returns CompartmentId
func (m realtimemessageauthentication) GetCompartmentId() *string {
	return m.CompartmentId
}

func (m realtimemessageauthentication) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m realtimemessageauthentication) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// RealtimeMessageAuthenticationAuthenticationTypeEnum Enum with underlying type: string
type RealtimeMessageAuthenticationAuthenticationTypeEnum string

// Set of constants representing the allowable values for RealtimeMessageAuthenticationAuthenticationTypeEnum
const (
	RealtimeMessageAuthenticationAuthenticationTypeCredentials RealtimeMessageAuthenticationAuthenticationTypeEnum = "CREDENTIALS"
	RealtimeMessageAuthenticationAuthenticationTypeToken       RealtimeMessageAuthenticationAuthenticationTypeEnum = "TOKEN"
)

var mappingRealtimeMessageAuthenticationAuthenticationTypeEnum = map[string]RealtimeMessageAuthenticationAuthenticationTypeEnum{
	"CREDENTIALS": RealtimeMessageAuthenticationAuthenticationTypeCredentials,
	"TOKEN":       RealtimeMessageAuthenticationAuthenticationTypeToken,
}

var mappingRealtimeMessageAuthenticationAuthenticationTypeEnumLowerCase = map[string]RealtimeMessageAuthenticationAuthenticationTypeEnum{
	"credentials": RealtimeMessageAuthenticationAuthenticationTypeCredentials,
	"token":       RealtimeMessageAuthenticationAuthenticationTypeToken,
}

// GetRealtimeMessageAuthenticationAuthenticationTypeEnumValues Enumerates the set of values for RealtimeMessageAuthenticationAuthenticationTypeEnum
func GetRealtimeMessageAuthenticationAuthenticationTypeEnumValues() []RealtimeMessageAuthenticationAuthenticationTypeEnum {
	values := make([]RealtimeMessageAuthenticationAuthenticationTypeEnum, 0)
	for _, v := range mappingRealtimeMessageAuthenticationAuthenticationTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetRealtimeMessageAuthenticationAuthenticationTypeEnumStringValues Enumerates the set of values in String for RealtimeMessageAuthenticationAuthenticationTypeEnum
func GetRealtimeMessageAuthenticationAuthenticationTypeEnumStringValues() []string {
	return []string{
		"CREDENTIALS",
		"TOKEN",
	}
}

// GetMappingRealtimeMessageAuthenticationAuthenticationTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingRealtimeMessageAuthenticationAuthenticationTypeEnum(val string) (RealtimeMessageAuthenticationAuthenticationTypeEnum, bool) {
	enum, ok := mappingRealtimeMessageAuthenticationAuthenticationTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
