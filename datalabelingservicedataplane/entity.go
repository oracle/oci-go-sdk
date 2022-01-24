// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Labeling API
//
// Use Data Labeling API to create Annotations on Images, Texts & Documents, and generate snapshots.
//

package datalabelingservicedataplane

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/v56/common"
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
	case "TEXTSELECTION":
		mm := TextSelectionEntity{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		return *m, nil
	}
}

func (m entity) String() string {
	return common.PointerString(m)
}

// EntityEntityTypeEnum Enum with underlying type: string
type EntityEntityTypeEnum string

// Set of constants representing the allowable values for EntityEntityTypeEnum
const (
	EntityEntityTypeGeneric              EntityEntityTypeEnum = "GENERIC"
	EntityEntityTypeImageobjectselection EntityEntityTypeEnum = "IMAGEOBJECTSELECTION"
	EntityEntityTypeTextselection        EntityEntityTypeEnum = "TEXTSELECTION"
)

var mappingEntityEntityType = map[string]EntityEntityTypeEnum{
	"GENERIC":              EntityEntityTypeGeneric,
	"IMAGEOBJECTSELECTION": EntityEntityTypeImageobjectselection,
	"TEXTSELECTION":        EntityEntityTypeTextselection,
}

// GetEntityEntityTypeEnumValues Enumerates the set of values for EntityEntityTypeEnum
func GetEntityEntityTypeEnumValues() []EntityEntityTypeEnum {
	values := make([]EntityEntityTypeEnum, 0)
	for _, v := range mappingEntityEntityType {
		values = append(values, v)
	}
	return values
}
