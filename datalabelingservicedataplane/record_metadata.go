// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
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
		common.Logf("Received unsupported enum value for RecordMetadata: %s.", m.RecordType)
		return *m, nil
	}
}

func (m recordmetadata) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m recordmetadata) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// RecordMetadataRecordTypeEnum Enum with underlying type: string
type RecordMetadataRecordTypeEnum string

// Set of constants representing the allowable values for RecordMetadataRecordTypeEnum
const (
	RecordMetadataRecordTypeImageMetadata    RecordMetadataRecordTypeEnum = "IMAGE_METADATA"
	RecordMetadataRecordTypeTextMetadata     RecordMetadataRecordTypeEnum = "TEXT_METADATA"
	RecordMetadataRecordTypeDocumentMetadata RecordMetadataRecordTypeEnum = "DOCUMENT_METADATA"
)

var mappingRecordMetadataRecordTypeEnum = map[string]RecordMetadataRecordTypeEnum{
	"IMAGE_METADATA":    RecordMetadataRecordTypeImageMetadata,
	"TEXT_METADATA":     RecordMetadataRecordTypeTextMetadata,
	"DOCUMENT_METADATA": RecordMetadataRecordTypeDocumentMetadata,
}

var mappingRecordMetadataRecordTypeEnumLowerCase = map[string]RecordMetadataRecordTypeEnum{
	"image_metadata":    RecordMetadataRecordTypeImageMetadata,
	"text_metadata":     RecordMetadataRecordTypeTextMetadata,
	"document_metadata": RecordMetadataRecordTypeDocumentMetadata,
}

// GetRecordMetadataRecordTypeEnumValues Enumerates the set of values for RecordMetadataRecordTypeEnum
func GetRecordMetadataRecordTypeEnumValues() []RecordMetadataRecordTypeEnum {
	values := make([]RecordMetadataRecordTypeEnum, 0)
	for _, v := range mappingRecordMetadataRecordTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetRecordMetadataRecordTypeEnumStringValues Enumerates the set of values in String for RecordMetadataRecordTypeEnum
func GetRecordMetadataRecordTypeEnumStringValues() []string {
	return []string{
		"IMAGE_METADATA",
		"TEXT_METADATA",
		"DOCUMENT_METADATA",
	}
}

// GetMappingRecordMetadataRecordTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingRecordMetadataRecordTypeEnum(val string) (RecordMetadataRecordTypeEnum, bool) {
	enum, ok := mappingRecordMetadataRecordTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
