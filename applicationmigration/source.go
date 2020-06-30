// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Application Migration Service API
//
// API for the Application Migration service. Use this API to migrate applications from Oracle Cloud Infrastructure - Classic to Oracle Cloud Infrastructure.
//

package applicationmigration

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/common"
)

// Source The Source object. Sources represent external locations from which
// applications may be imported into an OCI tenancy.
type Source struct {

	// Unique identifier (OCID) for the source
	Id *string `mandatory:"false" json:"id"`

	// Unique idenfifier (OCID) for the compartment where the Source is located.
	CompartmentId *string `mandatory:"false" json:"compartmentId"`

	// Human-readable name of the source.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Description of the source.
	Description *string `mandatory:"false" json:"description"`

	// The date and time at which the source was created
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// The current state of the Source
	LifecycleState SourceLifecycleStatesEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	// Details about the current lifecycle state
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	SourceDetails SourceDetails `mandatory:"false" json:"sourceDetails"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m Source) String() string {
	return common.PointerString(m)
}

// UnmarshalJSON unmarshals from json
func (m *Source) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Id               *string                           `json:"id"`
		CompartmentId    *string                           `json:"compartmentId"`
		DisplayName      *string                           `json:"displayName"`
		Description      *string                           `json:"description"`
		TimeCreated      *common.SDKTime                   `json:"timeCreated"`
		LifecycleState   SourceLifecycleStatesEnum         `json:"lifecycleState"`
		LifecycleDetails *string                           `json:"lifecycleDetails"`
		SourceDetails    sourcedetails                     `json:"sourceDetails"`
		FreeformTags     map[string]string                 `json:"freeformTags"`
		DefinedTags      map[string]map[string]interface{} `json:"definedTags"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.Id = model.Id

	m.CompartmentId = model.CompartmentId

	m.DisplayName = model.DisplayName

	m.Description = model.Description

	m.TimeCreated = model.TimeCreated

	m.LifecycleState = model.LifecycleState

	m.LifecycleDetails = model.LifecycleDetails

	nn, e = model.SourceDetails.UnmarshalPolymorphicJSON(model.SourceDetails.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.SourceDetails = nn.(SourceDetails)
	} else {
		m.SourceDetails = nil
	}

	m.FreeformTags = model.FreeformTags

	m.DefinedTags = model.DefinedTags

	return
}
