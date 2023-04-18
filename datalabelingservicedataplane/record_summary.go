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

// RecordSummary A record summary is the representation returned in list views.  It is usually a subset of the full record entity and should not contain any potentially sensitive information.
type RecordSummary struct {

	// The OCID of the record.
	Id *string `mandatory:"true" json:"id"`

	// The name is automatically assigned by the service. It is unique and immutable
	Name *string `mandatory:"true" json:"name"`

	// The date and time the resource was created, in the timestamp format defined by RFC3339.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The date and time the resource was updated, in the timestamp format defined by RFC3339.
	TimeUpdated *common.SDKTime `mandatory:"true" json:"timeUpdated"`

	// The OCID of the dataset to associate the record with.
	DatasetId *string `mandatory:"true" json:"datasetId"`

	// The OCID of the compartment for the task.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// Whether or not the record has been labeled and has associated annotations.
	IsLabeled *bool `mandatory:"true" json:"isLabeled"`

	// Describes the lifecycle state.
	LifecycleState RecordLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	RecordMetadata RecordMetadata `mandatory:"false" json:"recordMetadata"`

	// A simple key-value pair that is applied without any predefined name, type, or scope. It exists for cross-compatibility only.
	// For example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// The defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m RecordSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m RecordSummary) ValidateEnumValue() (bool, error) {
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
func (m *RecordSummary) UnmarshalJSON(data []byte) (e error) {
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

	m.IsLabeled = model.IsLabeled

	m.LifecycleState = model.LifecycleState

	return
}
