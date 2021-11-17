// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// DlsDataPlane API
//
// A description of the DlsDataPlane API.
//

package datalabelingservicedataplane

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/v52/common"
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

// CreateSourceDetailsSourceTypeEnum Enum with underlying type: string
type CreateSourceDetailsSourceTypeEnum string

// Set of constants representing the allowable values for CreateSourceDetailsSourceTypeEnum
const (
	CreateSourceDetailsSourceTypeObjectStorage CreateSourceDetailsSourceTypeEnum = "OBJECT_STORAGE"
)

var mappingCreateSourceDetailsSourceType = map[string]CreateSourceDetailsSourceTypeEnum{
	"OBJECT_STORAGE": CreateSourceDetailsSourceTypeObjectStorage,
}

// GetCreateSourceDetailsSourceTypeEnumValues Enumerates the set of values for CreateSourceDetailsSourceTypeEnum
func GetCreateSourceDetailsSourceTypeEnumValues() []CreateSourceDetailsSourceTypeEnum {
	values := make([]CreateSourceDetailsSourceTypeEnum, 0)
	for _, v := range mappingCreateSourceDetailsSourceType {
		values = append(values, v)
	}
	return values
}
