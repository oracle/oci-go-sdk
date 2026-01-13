// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
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

// CustomizationDatasetDetails Customization Training Dataset
type CustomizationDatasetDetails interface {
}

type customizationdatasetdetails struct {
	JsonData    []byte
	DatasetType string `json:"datasetType"`
}

// UnmarshalJSON unmarshals json
func (m *customizationdatasetdetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalercustomizationdatasetdetails customizationdatasetdetails
	s := struct {
		Model Unmarshalercustomizationdatasetdetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.DatasetType = s.Model.DatasetType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *customizationdatasetdetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.DatasetType {
	case "OBJECT_STORAGE":
		mm := ObjectStorageDataset{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "ENTITY_LIST":
		mm := EntityListDataset{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Received unsupported enum value for CustomizationDatasetDetails: %s.", m.DatasetType)
		return *m, nil
	}
}

func (m customizationdatasetdetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m customizationdatasetdetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// CustomizationDatasetDetailsDatasetTypeEnum Enum with underlying type: string
type CustomizationDatasetDetailsDatasetTypeEnum string

// Set of constants representing the allowable values for CustomizationDatasetDetailsDatasetTypeEnum
const (
	CustomizationDatasetDetailsDatasetTypeObjectStorage CustomizationDatasetDetailsDatasetTypeEnum = "OBJECT_STORAGE"
	CustomizationDatasetDetailsDatasetTypeEntityList    CustomizationDatasetDetailsDatasetTypeEnum = "ENTITY_LIST"
)

var mappingCustomizationDatasetDetailsDatasetTypeEnum = map[string]CustomizationDatasetDetailsDatasetTypeEnum{
	"OBJECT_STORAGE": CustomizationDatasetDetailsDatasetTypeObjectStorage,
	"ENTITY_LIST":    CustomizationDatasetDetailsDatasetTypeEntityList,
}

var mappingCustomizationDatasetDetailsDatasetTypeEnumLowerCase = map[string]CustomizationDatasetDetailsDatasetTypeEnum{
	"object_storage": CustomizationDatasetDetailsDatasetTypeObjectStorage,
	"entity_list":    CustomizationDatasetDetailsDatasetTypeEntityList,
}

// GetCustomizationDatasetDetailsDatasetTypeEnumValues Enumerates the set of values for CustomizationDatasetDetailsDatasetTypeEnum
func GetCustomizationDatasetDetailsDatasetTypeEnumValues() []CustomizationDatasetDetailsDatasetTypeEnum {
	values := make([]CustomizationDatasetDetailsDatasetTypeEnum, 0)
	for _, v := range mappingCustomizationDatasetDetailsDatasetTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetCustomizationDatasetDetailsDatasetTypeEnumStringValues Enumerates the set of values in String for CustomizationDatasetDetailsDatasetTypeEnum
func GetCustomizationDatasetDetailsDatasetTypeEnumStringValues() []string {
	return []string{
		"OBJECT_STORAGE",
		"ENTITY_LIST",
	}
}

// GetMappingCustomizationDatasetDetailsDatasetTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCustomizationDatasetDetailsDatasetTypeEnum(val string) (CustomizationDatasetDetailsDatasetTypeEnum, bool) {
	enum, ok := mappingCustomizationDatasetDetailsDatasetTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
