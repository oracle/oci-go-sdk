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
	"github.com/oracle/oci-go-sdk/v63/common"
	"strings"
)

// Record A record represents an entry in a dataset that needs labeling.
type Record struct {

	// The OCID of the record.
	Id *string `mandatory:"true" json:"id"`

	// The name is created by the user. It is unique and immutable.
	Name *string `mandatory:"true" json:"name"`

	// The date and time the resource was created, in the timestamp format defined by RFC3339.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The date and time the resource was updated, in the timestamp format defined by RFC3339.
	TimeUpdated *common.SDKTime `mandatory:"true" json:"timeUpdated"`

	// The OCID of the dataset to associate the record with.
	DatasetId *string `mandatory:"true" json:"datasetId"`

	// The OCID of the compartment for the task.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	SourceDetails SourceDetails `mandatory:"true" json:"sourceDetails"`

	// Whether or not the record has been labeled and has associated annotations.
	IsLabeled *bool `mandatory:"true" json:"isLabeled"`

	// The lifecycle state of the record.
	// ACTIVE - The record is active and ready for labeling.
	// INACTIVE - The record has been marked as inactive and should not be used for labeling.
	// DELETED - The record has been deleted and is no longer available for labeling.
	LifecycleState RecordLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	RecordMetadata RecordMetadata `mandatory:"false" json:"recordMetadata"`

	// A simple key-value pair that is applied without any predefined name, type, or scope. It exists for cross-compatibility only.
	// For example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// The defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m Record) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m Record) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingRecordLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetRecordLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *Record) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		RecordMetadata recordmetadata                    `json:"recordMetadata"`
		FreeformTags   map[string]string                 `json:"freeformTags"`
		DefinedTags    map[string]map[string]interface{} `json:"definedTags"`
		Id             *string                           `json:"id"`
		Name           *string                           `json:"name"`
		TimeCreated    *common.SDKTime                   `json:"timeCreated"`
		TimeUpdated    *common.SDKTime                   `json:"timeUpdated"`
		DatasetId      *string                           `json:"datasetId"`
		CompartmentId  *string                           `json:"compartmentId"`
		SourceDetails  sourcedetails                     `json:"sourceDetails"`
		IsLabeled      *bool                             `json:"isLabeled"`
		LifecycleState RecordLifecycleStateEnum          `json:"lifecycleState"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	nn, e = model.RecordMetadata.UnmarshalPolymorphicJSON(model.RecordMetadata.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.RecordMetadata = nn.(RecordMetadata)
	} else {
		m.RecordMetadata = nil
	}

	m.FreeformTags = model.FreeformTags

	m.DefinedTags = model.DefinedTags

	m.Id = model.Id

	m.Name = model.Name

	m.TimeCreated = model.TimeCreated

	m.TimeUpdated = model.TimeUpdated

	m.DatasetId = model.DatasetId

	m.CompartmentId = model.CompartmentId

	nn, e = model.SourceDetails.UnmarshalPolymorphicJSON(model.SourceDetails.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.SourceDetails = nn.(SourceDetails)
	} else {
		m.SourceDetails = nil
	}

	m.IsLabeled = model.IsLabeled

	m.LifecycleState = model.LifecycleState

	return
}

// RecordLifecycleStateEnum Enum with underlying type: string
type RecordLifecycleStateEnum string

// Set of constants representing the allowable values for RecordLifecycleStateEnum
const (
	RecordLifecycleStateActive   RecordLifecycleStateEnum = "ACTIVE"
	RecordLifecycleStateInactive RecordLifecycleStateEnum = "INACTIVE"
	RecordLifecycleStateDeleted  RecordLifecycleStateEnum = "DELETED"
)

var mappingRecordLifecycleStateEnum = map[string]RecordLifecycleStateEnum{
	"ACTIVE":   RecordLifecycleStateActive,
	"INACTIVE": RecordLifecycleStateInactive,
	"DELETED":  RecordLifecycleStateDeleted,
}

var mappingRecordLifecycleStateEnumLowerCase = map[string]RecordLifecycleStateEnum{
	"active":   RecordLifecycleStateActive,
	"inactive": RecordLifecycleStateInactive,
	"deleted":  RecordLifecycleStateDeleted,
}

// GetRecordLifecycleStateEnumValues Enumerates the set of values for RecordLifecycleStateEnum
func GetRecordLifecycleStateEnumValues() []RecordLifecycleStateEnum {
	values := make([]RecordLifecycleStateEnum, 0)
	for _, v := range mappingRecordLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetRecordLifecycleStateEnumStringValues Enumerates the set of values in String for RecordLifecycleStateEnum
func GetRecordLifecycleStateEnumStringValues() []string {
	return []string{
		"ACTIVE",
		"INACTIVE",
		"DELETED",
	}
}

// GetMappingRecordLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingRecordLifecycleStateEnum(val string) (RecordLifecycleStateEnum, bool) {
	enum, ok := mappingRecordLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
