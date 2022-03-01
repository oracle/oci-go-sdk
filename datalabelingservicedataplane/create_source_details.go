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
	"fmt"
	"github.com/oracle/oci-go-sdk/v60/common"
	"strings"
)

// CreateSourceDetails The source information is a polymorphic entity. It captures the details of how to access the data for record creation. The discriminator type must match the dataset's source type. The convention will be enforced by the API. It should only provide the difference in data necessary to access the content, i.e. the object storage path, or the database record id.
type CreateSourceDetails interface {
}

type createsourcedetails struct {
	JsonData   []byte
	SourceType string `json:"sourceType"`
}

// UnmarshalJSON unmarshals json
func (m *createsourcedetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalercreatesourcedetails createsourcedetails
	s := struct {
		Model Unmarshalercreatesourcedetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.SourceType = s.Model.SourceType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *createsourcedetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.SourceType {
	case "OBJECT_STORAGE":
		mm := CreateObjectStorageSourceDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		return *m, nil
	}
}

func (m createsourcedetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m createsourcedetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// CreateSourceDetailsSourceTypeEnum Enum with underlying type: string
type CreateSourceDetailsSourceTypeEnum string

// Set of constants representing the allowable values for CreateSourceDetailsSourceTypeEnum
const (
	CreateSourceDetailsSourceTypeObjectStorage CreateSourceDetailsSourceTypeEnum = "OBJECT_STORAGE"
)

var mappingCreateSourceDetailsSourceTypeEnum = map[string]CreateSourceDetailsSourceTypeEnum{
	"OBJECT_STORAGE": CreateSourceDetailsSourceTypeObjectStorage,
}

var mappingCreateSourceDetailsSourceTypeEnumLowerCase = map[string]CreateSourceDetailsSourceTypeEnum{
	"object_storage": CreateSourceDetailsSourceTypeObjectStorage,
}

// GetCreateSourceDetailsSourceTypeEnumValues Enumerates the set of values for CreateSourceDetailsSourceTypeEnum
func GetCreateSourceDetailsSourceTypeEnumValues() []CreateSourceDetailsSourceTypeEnum {
	values := make([]CreateSourceDetailsSourceTypeEnum, 0)
	for _, v := range mappingCreateSourceDetailsSourceTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetCreateSourceDetailsSourceTypeEnumStringValues Enumerates the set of values in String for CreateSourceDetailsSourceTypeEnum
func GetCreateSourceDetailsSourceTypeEnumStringValues() []string {
	return []string{
		"OBJECT_STORAGE",
	}
}

// GetMappingCreateSourceDetailsSourceTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCreateSourceDetailsSourceTypeEnum(val string) (CreateSourceDetailsSourceTypeEnum, bool) {
	enum, ok := mappingCreateSourceDetailsSourceTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
