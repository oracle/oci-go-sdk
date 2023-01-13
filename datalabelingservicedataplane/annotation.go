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

// Annotation An annotation represents a user- or machine-generated annotation for a given record.  The details of the annotation are captured in the RecordAnnotationDetails.
type Annotation struct {

	// The OCID of the annotation.
	Id *string `mandatory:"true" json:"id"`

	// The date and time the annotation was created, in the timestamp format defined by RFC3339.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The date and time the resource was updated, in the timestamp format defined by RFC3339.
	TimeUpdated *common.SDKTime `mandatory:"true" json:"timeUpdated"`

	// The OCID of the principal which created the annotation.
	CreatedBy *string `mandatory:"true" json:"createdBy"`

	// The OCID of the principal which updated the annotation.
	UpdatedBy *string `mandatory:"true" json:"updatedBy"`

	// The OCID of the record annotated.
	RecordId *string `mandatory:"true" json:"recordId"`

	// The entity types are validated against the dataset to ensure consistency.
	Entities []Entity `mandatory:"true" json:"entities"`

	// The OCID of the compartment for the annotation. This is tied to the dataset. It is not changeable on the record itself.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The lifecycle state of an annotation.
	// ACTIVE - The annotation is active to be used for labeling.
	// INACTIVE - The annotation has been marked as inactive and should not be used for labeling.
	// DELETED - Tha annotation been deleted and no longer available for labeling.
	LifecycleState AnnotationLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// An integer value used in achieving concurrency control, this field will be used to generate eTags.
	LifetimeLogicalClock *int `mandatory:"true" json:"lifetimeLogicalClock"`

	// A simple key-value pair that is applied without any predefined name, type, or scope. It exists for cross-compatibility only.
	// For example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// The defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m Annotation) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m Annotation) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingAnnotationLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetAnnotationLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *Annotation) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		FreeformTags         map[string]string                 `json:"freeformTags"`
		DefinedTags          map[string]map[string]interface{} `json:"definedTags"`
		Id                   *string                           `json:"id"`
		TimeCreated          *common.SDKTime                   `json:"timeCreated"`
		TimeUpdated          *common.SDKTime                   `json:"timeUpdated"`
		CreatedBy            *string                           `json:"createdBy"`
		UpdatedBy            *string                           `json:"updatedBy"`
		RecordId             *string                           `json:"recordId"`
		Entities             []entity                          `json:"entities"`
		CompartmentId        *string                           `json:"compartmentId"`
		LifecycleState       AnnotationLifecycleStateEnum      `json:"lifecycleState"`
		LifetimeLogicalClock *int                              `json:"lifetimeLogicalClock"`
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

	m.LifetimeLogicalClock = model.LifetimeLogicalClock

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

var mappingAnnotationLifecycleStateEnum = map[string]AnnotationLifecycleStateEnum{
	"ACTIVE":   AnnotationLifecycleStateActive,
	"INACTIVE": AnnotationLifecycleStateInactive,
	"DELETED":  AnnotationLifecycleStateDeleted,
}

var mappingAnnotationLifecycleStateEnumLowerCase = map[string]AnnotationLifecycleStateEnum{
	"active":   AnnotationLifecycleStateActive,
	"inactive": AnnotationLifecycleStateInactive,
	"deleted":  AnnotationLifecycleStateDeleted,
}

// GetAnnotationLifecycleStateEnumValues Enumerates the set of values for AnnotationLifecycleStateEnum
func GetAnnotationLifecycleStateEnumValues() []AnnotationLifecycleStateEnum {
	values := make([]AnnotationLifecycleStateEnum, 0)
	for _, v := range mappingAnnotationLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetAnnotationLifecycleStateEnumStringValues Enumerates the set of values in String for AnnotationLifecycleStateEnum
func GetAnnotationLifecycleStateEnumStringValues() []string {
	return []string{
		"ACTIVE",
		"INACTIVE",
		"DELETED",
	}
}

// GetMappingAnnotationLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAnnotationLifecycleStateEnum(val string) (AnnotationLifecycleStateEnum, bool) {
	enum, ok := mappingAnnotationLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
