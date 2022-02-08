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
	"github.com/oracle/oci-go-sdk/v57/common"
	"strings"
)

// SourceDetails The source information is a polymorphic entity. It captures the details of data used for record creation. The discriminator type must match the dataset's source type. The convention is enforced by the API.
type SourceDetails interface {
}

type sourcedetails struct {
	JsonData   []byte
	SourceType string `json:"sourceType"`
}

// UnmarshalJSON unmarshals json
func (m *sourcedetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalersourcedetails sourcedetails
	s := struct {
		Model Unmarshalersourcedetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.SourceType = s.Model.SourceType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *sourcedetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.SourceType {
	case "OBJECT_STORAGE":
		mm := ObjectStorageSourceDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		return *m, nil
	}
}

func (m sourcedetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m sourcedetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// SourceDetailsSourceTypeEnum Enum with underlying type: string
type SourceDetailsSourceTypeEnum string

// Set of constants representing the allowable values for SourceDetailsSourceTypeEnum
const (
	SourceDetailsSourceTypeObjectStorage SourceDetailsSourceTypeEnum = "OBJECT_STORAGE"
)

var mappingSourceDetailsSourceTypeEnum = map[string]SourceDetailsSourceTypeEnum{
	"OBJECT_STORAGE": SourceDetailsSourceTypeObjectStorage,
}

// GetSourceDetailsSourceTypeEnumValues Enumerates the set of values for SourceDetailsSourceTypeEnum
func GetSourceDetailsSourceTypeEnumValues() []SourceDetailsSourceTypeEnum {
	values := make([]SourceDetailsSourceTypeEnum, 0)
	for _, v := range mappingSourceDetailsSourceTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetSourceDetailsSourceTypeEnumStringValues Enumerates the set of values in String for SourceDetailsSourceTypeEnum
func GetSourceDetailsSourceTypeEnumStringValues() []string {
	return []string{
		"OBJECT_STORAGE",
	}
}
