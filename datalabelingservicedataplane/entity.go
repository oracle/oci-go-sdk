// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Labeling API
//
// Use Data Labeling API to create Annotations on Images, Texts & Documents, and generate snapshots.
//

package datalabelingservicedataplane

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// Entity An entity allows the labeler to identify an object in the record to label.  This can be, for example, a snippet of text, an entire image, or a bounding box within an image.  All entity types have an array of labels that are indexed. If more than one label is provided, but the annotationType on the corresponding dataset is for a single class, the API rejects the create annotation request.
type Entity interface {
}

type entity struct {
	JsonData   []byte
	EntityType string `json:"entityType"`
}

// UnmarshalJSON unmarshals json
func (m *entity) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerentity entity
	s := struct {
		Model Unmarshalerentity
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.EntityType = s.Model.EntityType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *entity) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.EntityType {
	case "IMAGEOBJECTSELECTION":
		mm := ImageObjectSelectionEntity{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "GENERIC":
		mm := GenericEntity{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "KEYVALUESELECTION":
		mm := KeyValueSelectionEntity{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "TEXTSELECTION":
		mm := TextSelectionEntity{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for Entity: %s.", m.EntityType)
		return *m, nil
	}
}

func (m entity) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m entity) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// EntityEntityTypeEnum Enum with underlying type: string
type EntityEntityTypeEnum string

// Set of constants representing the allowable values for EntityEntityTypeEnum
const (
	EntityEntityTypeGeneric              EntityEntityTypeEnum = "GENERIC"
	EntityEntityTypeImageobjectselection EntityEntityTypeEnum = "IMAGEOBJECTSELECTION"
	EntityEntityTypeTextselection        EntityEntityTypeEnum = "TEXTSELECTION"
	EntityEntityTypeKeyvalueselection    EntityEntityTypeEnum = "KEYVALUESELECTION"
)

var mappingEntityEntityTypeEnum = map[string]EntityEntityTypeEnum{
	"GENERIC":              EntityEntityTypeGeneric,
	"IMAGEOBJECTSELECTION": EntityEntityTypeImageobjectselection,
	"TEXTSELECTION":        EntityEntityTypeTextselection,
	"KEYVALUESELECTION":    EntityEntityTypeKeyvalueselection,
}

var mappingEntityEntityTypeEnumLowerCase = map[string]EntityEntityTypeEnum{
	"generic":              EntityEntityTypeGeneric,
	"imageobjectselection": EntityEntityTypeImageobjectselection,
	"textselection":        EntityEntityTypeTextselection,
	"keyvalueselection":    EntityEntityTypeKeyvalueselection,
}

// GetEntityEntityTypeEnumValues Enumerates the set of values for EntityEntityTypeEnum
func GetEntityEntityTypeEnumValues() []EntityEntityTypeEnum {
	values := make([]EntityEntityTypeEnum, 0)
	for _, v := range mappingEntityEntityTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetEntityEntityTypeEnumStringValues Enumerates the set of values in String for EntityEntityTypeEnum
func GetEntityEntityTypeEnumStringValues() []string {
	return []string{
		"GENERIC",
		"IMAGEOBJECTSELECTION",
		"TEXTSELECTION",
		"KEYVALUESELECTION",
	}
}

// GetMappingEntityEntityTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingEntityEntityTypeEnum(val string) (EntityEntityTypeEnum, bool) {
	enum, ok := mappingEntityEntityTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
