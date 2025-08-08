// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Threat Intelligence API
//
// Use the Threat Intelligence API to search for information about known threat indicators, including suspicious IP addresses, domain names, and other digital fingerprints. Threat Intelligence is a managed database of curated threat intelligence that comes from first party Oracle security insights, open source feeds, and vendor-procured data. For more information, see the Threat Intelligence documentation (https://docs.oracle.com/iaas/Content/threat-intel/home.htm).
//

package threatintelligence

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// EntityReference A reference to a resource or other entity.
type EntityReference interface {
}

type entityreference struct {
	JsonData []byte
	Type     string `json:"type"`
}

// UnmarshalJSON unmarshals json
func (m *entityreference) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerentityreference entityreference
	s := struct {
		Model Unmarshalerentityreference
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Type = s.Model.Type

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *entityreference) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Type {
	case "INDICATOR":
		mm := IndicatorReference{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Received unsupported enum value for EntityReference: %s.", m.Type)
		return *m, nil
	}
}

func (m entityreference) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m entityreference) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// EntityReferenceTypeEnum Enum with underlying type: string
type EntityReferenceTypeEnum string

// Set of constants representing the allowable values for EntityReferenceTypeEnum
const (
	EntityReferenceTypeIndicator EntityReferenceTypeEnum = "INDICATOR"
)

var mappingEntityReferenceTypeEnum = map[string]EntityReferenceTypeEnum{
	"INDICATOR": EntityReferenceTypeIndicator,
}

var mappingEntityReferenceTypeEnumLowerCase = map[string]EntityReferenceTypeEnum{
	"indicator": EntityReferenceTypeIndicator,
}

// GetEntityReferenceTypeEnumValues Enumerates the set of values for EntityReferenceTypeEnum
func GetEntityReferenceTypeEnumValues() []EntityReferenceTypeEnum {
	values := make([]EntityReferenceTypeEnum, 0)
	for _, v := range mappingEntityReferenceTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetEntityReferenceTypeEnumStringValues Enumerates the set of values in String for EntityReferenceTypeEnum
func GetEntityReferenceTypeEnumStringValues() []string {
	return []string{
		"INDICATOR",
	}
}

// GetMappingEntityReferenceTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingEntityReferenceTypeEnum(val string) (EntityReferenceTypeEnum, bool) {
	enum, ok := mappingEntityReferenceTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
