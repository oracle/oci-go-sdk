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

// SandboxSource Base source configuration used to initialize a sandbox.
type SandboxSource interface {
}

type sandboxsource struct {
	JsonData   []byte
	SourceType string `json:"sourceType"`
}

// UnmarshalJSON unmarshals json
func (m *sandboxsource) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalersandboxsource sandboxsource
	s := struct {
		Model Unmarshalersandboxsource
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.SourceType = s.Model.SourceType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *sandboxsource) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.SourceType {
	case "SNAPSHOT":
		mm := SandboxSnapshotSource{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Received unsupported enum value for SandboxSource: %s.", m.SourceType)
		return *m, nil
	}
}

func (m sandboxsource) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m sandboxsource) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// SandboxSourceSourceTypeEnum Enum with underlying type: string
type SandboxSourceSourceTypeEnum string

// Set of constants representing the allowable values for SandboxSourceSourceTypeEnum
const (
	SandboxSourceSourceTypeSnapshot SandboxSourceSourceTypeEnum = "SNAPSHOT"
)

var mappingSandboxSourceSourceTypeEnum = map[string]SandboxSourceSourceTypeEnum{
	"SNAPSHOT": SandboxSourceSourceTypeSnapshot,
}

var mappingSandboxSourceSourceTypeEnumLowerCase = map[string]SandboxSourceSourceTypeEnum{
	"snapshot": SandboxSourceSourceTypeSnapshot,
}

// GetSandboxSourceSourceTypeEnumValues Enumerates the set of values for SandboxSourceSourceTypeEnum
func GetSandboxSourceSourceTypeEnumValues() []SandboxSourceSourceTypeEnum {
	values := make([]SandboxSourceSourceTypeEnum, 0)
	for _, v := range mappingSandboxSourceSourceTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetSandboxSourceSourceTypeEnumStringValues Enumerates the set of values in String for SandboxSourceSourceTypeEnum
func GetSandboxSourceSourceTypeEnumStringValues() []string {
	return []string{
		"SNAPSHOT",
	}
}

// GetMappingSandboxSourceSourceTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSandboxSourceSourceTypeEnum(val string) (SandboxSourceSourceTypeEnum, bool) {
	enum, ok := mappingSandboxSourceSourceTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
