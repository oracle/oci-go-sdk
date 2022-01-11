// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
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
	"github.com/oracle/oci-go-sdk/v55/common"
)

// CreateSourceDetails The configuration details for creating a source.
// When you create a source, provide the required information to let Application Migration access the source environment.
// You must also assign a name and provide a description for the source. This helps you to identify the appropriate source environment when you
// have multiple sources defined.
// **Warning:** Oracle recommends that you avoid using any confidential information when you supply string values using the API.
type CreateSourceDetails struct {

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment that contains the source.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	SourceDetails SourceDetails `mandatory:"true" json:"sourceDetails"`

	// Name of the source. This helps you to identify the appropriate source environment when you have multiple sources defined.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Description of the source. This helps you to identify the appropriate source environment when you have multiple sources defined.
	Description *string `mandatory:"false" json:"description"`

	AuthorizationDetails AuthorizationDetails `mandatory:"false" json:"authorizationDetails"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm). Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m CreateSourceDetails) String() string {
	return common.PointerString(m)
}

// UnmarshalJSON unmarshals from json
func (m *CreateSourceDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		DisplayName          *string                           `json:"displayName"`
		Description          *string                           `json:"description"`
		AuthorizationDetails authorizationdetails              `json:"authorizationDetails"`
		FreeformTags         map[string]string                 `json:"freeformTags"`
		DefinedTags          map[string]map[string]interface{} `json:"definedTags"`
		CompartmentId        *string                           `json:"compartmentId"`
		SourceDetails        sourcedetails                     `json:"sourceDetails"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.DisplayName = model.DisplayName

	m.Description = model.Description

	nn, e = model.AuthorizationDetails.UnmarshalPolymorphicJSON(model.AuthorizationDetails.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.AuthorizationDetails = nn.(AuthorizationDetails)
	} else {
		m.AuthorizationDetails = nil
	}

	m.FreeformTags = model.FreeformTags

	m.DefinedTags = model.DefinedTags

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

	return
}
