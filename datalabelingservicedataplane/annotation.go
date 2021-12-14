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
	"github.com/oracle/oci-go-sdk/v54/common"
)

// Annotation An Annotation represents a user/machine generated annotation for a given record.  The details of the annotation are captured in the RecordAnnotationDetails.
type Annotation struct {

	// The OCID of the annotation
	Id *string `mandatory:"true" json:"id"`

	// The date and time the annotation was created, in the timestamp format defined by RFC3339.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The date and time the resource was updated, in the timestamp format defined by RFC3339.
	TimeUpdated *common.SDKTime `mandatory:"true" json:"timeUpdated"`

	// The OCID of the principal who created the annotation
	CreatedBy *string `mandatory:"true" json:"createdBy"`

	// The OCID of the principal who updated the annotation
	UpdatedBy *string `mandatory:"true" json:"updatedBy"`

	// The OCID of the record annotated
	RecordId *string `mandatory:"true" json:"recordId"`

	// The entity types will be validated against the dataset to ensure consistency.
	Entities []Entity `mandatory:"true" json:"entities"`

	// The OCID of the compartment for the annotation. This is tied to the dataset. It is not changeable on the record itself.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// Lifecycle State of an Annotation.
	// ACTIVE - Annotation is active to be used for labeling.
	// INACTIVE - Annotation has been marked as inactive and should not be used for labeling.
	// DELETED - Annotation been deleted and no longer available for labeling.
	LifecycleState AnnotationLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m Annotation) String() string {
	return common.PointerString(m)
}

// UnmarshalJSON unmarshals from json
func (m *Annotation) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		FreeformTags   map[string]string                 `json:"freeformTags"`
		DefinedTags    map[string]map[string]interface{} `json:"definedTags"`
		Id             *string                           `json:"id"`
		TimeCreated    *common.SDKTime                   `json:"timeCreated"`
		TimeUpdated    *common.SDKTime                   `json:"timeUpdated"`
		CreatedBy      *string                           `json:"createdBy"`
		UpdatedBy      *string                           `json:"updatedBy"`
		RecordId       *string                           `json:"recordId"`
		Entities       []entity                          `json:"entities"`
		CompartmentId  *string                           `json:"compartmentId"`
		LifecycleState AnnotationLifecycleStateEnum      `json:"lifecycleState"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.FreeformTags = model.FreeformTags

	m.DefinedTags = model.DefinedTags

	m.Id = model.Id

	m.TimeCreated = model.TimeCreated

	m.TimeUpdated = model.TimeUpdated

	m.CreatedBy = model.CreatedBy

	m.UpdatedBy = model.UpdatedBy

	m.RecordId = model.RecordId

	m.Entities = make([]Entity, len(model.Entities))
	for i, n := range model.Entities {
		nn, e = n.UnmarshalPolymorphicJSON(n.JsonData)
		if e != nil {
			return e
		}
		if nn != nil {
			m.Entities[i] = nn.(Entity)
		} else {
			m.Entities[i] = nil
		}
	}

	m.CompartmentId = model.CompartmentId

	m.LifecycleState = model.LifecycleState

	return
}

// AnnotationLifecycleStateEnum Enum with underlying type: string
type AnnotationLifecycleStateEnum string

// Set of constants representing the allowable values for AnnotationLifecycleStateEnum
const (
	AnnotationLifecycleStateActive   AnnotationLifecycleStateEnum = "ACTIVE"
	AnnotationLifecycleStateInactive AnnotationLifecycleStateEnum = "INACTIVE"
	AnnotationLifecycleStateDeleted  AnnotationLifecycleStateEnum = "DELETED"
)

var mappingAnnotationLifecycleState = map[string]AnnotationLifecycleStateEnum{
	"ACTIVE":   AnnotationLifecycleStateActive,
	"INACTIVE": AnnotationLifecycleStateInactive,
	"DELETED":  AnnotationLifecycleStateDeleted,
}

// GetAnnotationLifecycleStateEnumValues Enumerates the set of values for AnnotationLifecycleStateEnum
func GetAnnotationLifecycleStateEnumValues() []AnnotationLifecycleStateEnum {
	values := make([]AnnotationLifecycleStateEnum, 0)
	for _, v := range mappingAnnotationLifecycleState {
		values = append(values, v)
	}
	return values
}
