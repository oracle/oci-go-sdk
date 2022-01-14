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
	"github.com/oracle/oci-go-sdk/v55/common"
)

// RecordMetadata Collection of record's metadata.  This can be, for example, the height, width or depth of image for an image record.
type RecordMetadata interface {
}

type recordmetadata struct {
	JsonData   []byte
	RecordType string `json:"recordType"`
}

// UnmarshalJSON unmarshals json
func (m *recordmetadata) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerrecordmetadata recordmetadata
	s := struct {
		Model Unmarshalerrecordmetadata
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.RecordType = s.Model.RecordType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *recordmetadata) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.RecordType {
	case "DOCUMENT_METADATA":
		mm := DocumentMetadata{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "IMAGE_METADATA":
		mm := ImageMetadata{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "TEXT_METADATA":
		mm := TextMetadata{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		return *m, nil
	}
}

func (m recordmetadata) String() string {
	return common.PointerString(m)
}

// RecordMetadataRecordTypeEnum Enum with underlying type: string
type RecordMetadataRecordTypeEnum string

// Set of constants representing the allowable values for RecordMetadataRecordTypeEnum
const (
	RecordMetadataRecordTypeImageMetadata    RecordMetadataRecordTypeEnum = "IMAGE_METADATA"
	RecordMetadataRecordTypeTextMetadata     RecordMetadataRecordTypeEnum = "TEXT_METADATA"
	RecordMetadataRecordTypeDocumentMetadata RecordMetadataRecordTypeEnum = "DOCUMENT_METADATA"
)

var mappingRecordMetadataRecordType = map[string]RecordMetadataRecordTypeEnum{
	"IMAGE_METADATA":    RecordMetadataRecordTypeImageMetadata,
	"TEXT_METADATA":     RecordMetadataRecordTypeTextMetadata,
	"DOCUMENT_METADATA": RecordMetadataRecordTypeDocumentMetadata,
}

// GetRecordMetadataRecordTypeEnumValues Enumerates the set of values for RecordMetadataRecordTypeEnum
func GetRecordMetadataRecordTypeEnumValues() []RecordMetadataRecordTypeEnum {
	values := make([]RecordMetadataRecordTypeEnum, 0)
	for _, v := range mappingRecordMetadataRecordType {
		values = append(values, v)
	}
	return values
}
