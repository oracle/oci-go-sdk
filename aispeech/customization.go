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

// Customization Description of a Customization.
type Customization struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the job.
	Id *string `mandatory:"true" json:"id"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment where you want to create the job.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	ModelDetails *CustomizationModelDetails `mandatory:"true" json:"modelDetails"`

	// A user-friendly display name for the job.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Customization Details Alias
	Alias *string `mandatory:"false" json:"alias"`

	// A short description of the job.
	Description *string `mandatory:"false" json:"description"`

	// The current state of the Job.
	LifecycleState CustomizationLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	// A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// Customization Created time. An RFC3339 formatted datetime string.
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// Customization Updated time. An RFC3339 formatted datetime string.
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	TrainingDataset CustomizationDatasetDetails `mandatory:"false" json:"trainingDataset"`

	// array of referenced entities
	Entities []CustomizationReferencedEntities `mandatory:"false" json:"entities"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`.
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace-1": {"bar-key-1": "value-1", "bar-key-2": "value-2"}, "foo-namespace-2": {"bar-key-1": "value-1", "bar-key-2": "value-2"}}`.
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// Usage of system tag keys. These predefined keys are scoped to namespaces.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`.
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
}

func (m Customization) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m Customization) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingCustomizationLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetCustomizationLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *Customization) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		DisplayName      *string                           `json:"displayName"`
		Alias            *string                           `json:"alias"`
		Description      *string                           `json:"description"`
		LifecycleState   CustomizationLifecycleStateEnum   `json:"lifecycleState"`
		LifecycleDetails *string                           `json:"lifecycleDetails"`
		TimeCreated      *common.SDKTime                   `json:"timeCreated"`
		TimeUpdated      *common.SDKTime                   `json:"timeUpdated"`
		TrainingDataset  customizationdatasetdetails       `json:"trainingDataset"`
		Entities         []CustomizationReferencedEntities `json:"entities"`
		FreeformTags     map[string]string                 `json:"freeformTags"`
		DefinedTags      map[string]map[string]interface{} `json:"definedTags"`
		SystemTags       map[string]map[string]interface{} `json:"systemTags"`
		Id               *string                           `json:"id"`
		CompartmentId    *string                           `json:"compartmentId"`
		ModelDetails     *CustomizationModelDetails        `json:"modelDetails"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.DisplayName = model.DisplayName

	m.Alias = model.Alias

	m.Description = model.Description

	m.LifecycleState = model.LifecycleState

	m.LifecycleDetails = model.LifecycleDetails

	m.TimeCreated = model.TimeCreated

	m.TimeUpdated = model.TimeUpdated

	nn, e = model.TrainingDataset.UnmarshalPolymorphicJSON(model.TrainingDataset.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.TrainingDataset = nn.(CustomizationDatasetDetails)
	} else {
		m.TrainingDataset = nil
	}

	m.Entities = make([]CustomizationReferencedEntities, len(model.Entities))
	copy(m.Entities, model.Entities)
	m.FreeformTags = model.FreeformTags

	m.DefinedTags = model.DefinedTags

	m.SystemTags = model.SystemTags

	m.Id = model.Id

	m.CompartmentId = model.CompartmentId

	m.ModelDetails = model.ModelDetails

	return
}

// CustomizationLifecycleStateEnum Enum with underlying type: string
type CustomizationLifecycleStateEnum string

// Set of constants representing the allowable values for CustomizationLifecycleStateEnum
const (
	CustomizationLifecycleStateDeleting CustomizationLifecycleStateEnum = "DELETING"
	CustomizationLifecycleStateDeleted  CustomizationLifecycleStateEnum = "DELETED"
	CustomizationLifecycleStateFailed   CustomizationLifecycleStateEnum = "FAILED"
	CustomizationLifecycleStateUpdating CustomizationLifecycleStateEnum = "UPDATING"
	CustomizationLifecycleStateActive   CustomizationLifecycleStateEnum = "ACTIVE"
	CustomizationLifecycleStateCreating CustomizationLifecycleStateEnum = "CREATING"
)

var mappingCustomizationLifecycleStateEnum = map[string]CustomizationLifecycleStateEnum{
	"DELETING": CustomizationLifecycleStateDeleting,
	"DELETED":  CustomizationLifecycleStateDeleted,
	"FAILED":   CustomizationLifecycleStateFailed,
	"UPDATING": CustomizationLifecycleStateUpdating,
	"ACTIVE":   CustomizationLifecycleStateActive,
	"CREATING": CustomizationLifecycleStateCreating,
}

var mappingCustomizationLifecycleStateEnumLowerCase = map[string]CustomizationLifecycleStateEnum{
	"deleting": CustomizationLifecycleStateDeleting,
	"deleted":  CustomizationLifecycleStateDeleted,
	"failed":   CustomizationLifecycleStateFailed,
	"updating": CustomizationLifecycleStateUpdating,
	"active":   CustomizationLifecycleStateActive,
	"creating": CustomizationLifecycleStateCreating,
}

// GetCustomizationLifecycleStateEnumValues Enumerates the set of values for CustomizationLifecycleStateEnum
func GetCustomizationLifecycleStateEnumValues() []CustomizationLifecycleStateEnum {
	values := make([]CustomizationLifecycleStateEnum, 0)
	for _, v := range mappingCustomizationLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetCustomizationLifecycleStateEnumStringValues Enumerates the set of values in String for CustomizationLifecycleStateEnum
func GetCustomizationLifecycleStateEnumStringValues() []string {
	return []string{
		"DELETING",
		"DELETED",
		"FAILED",
		"UPDATING",
		"ACTIVE",
		"CREATING",
	}
}

// GetMappingCustomizationLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCustomizationLifecycleStateEnum(val string) (CustomizationLifecycleStateEnum, bool) {
	enum, ok := mappingCustomizationLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
