// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Application Migration API
//
// Application Migration simplifies the migration of applications from Oracle Cloud Infrastructure Classic to Oracle Cloud Infrastructure.
// You can use Application Migration API to migrate applications, such as Oracle Java Cloud Service, SOA Cloud Service, and Integration Classic
// instances, to Oracle Cloud Infrastructure. For more information, see
// Overview of Application Migration (https://docs.cloud.oracle.com/iaas/application-migration/appmigrationoverview.htm).
//

package applicationmigration

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/v46/common"
)

// Source The properties that define a source. Source refers to the source environment from which you migrate an application to Oracle Cloud
// Infrastructure. For more information, see Manage Sources (https://docs.cloud.oracle.com/iaas/application-migration/manage_sources.htm).
// To use any of the API operations, you must be authorized in an IAM policy. If you're not authorized, talk to an administrator. If you're
// an administrator who needs to write policies to give users access, see Getting Started with
// Policies (https://docs.cloud.oracle.com/Content/Identity/Concepts/policygetstarted.htm).
type Source struct {

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the source.
	Id *string `mandatory:"false" json:"id"`

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment that contains the source.
	CompartmentId *string `mandatory:"false" json:"compartmentId"`

	// Name of the source. This helps you to identify the appropriate source environment when you have multiple sources defined.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Description of the source. This helps you to identify the appropriate source environment when you have multiple sources defined.
	Description *string `mandatory:"false" json:"description"`

	// The date and time at which the source was created, in the format defined by RFC3339.
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// The current state of the source.
	LifecycleState SourceLifecycleStatesEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	// Details about the current lifecycle state of the source.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	SourceDetails SourceDetails `mandatory:"false" json:"sourceDetails"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm). Example: `{"Operations": {"CostCenter": "42"}}`
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
