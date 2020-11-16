// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
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
	"github.com/oracle/oci-go-sdk/v29/common"
)

// UpdateSourceDetails You can update the authorization details to access the source environment from which you want to migrate applications to
// Oracle Cloud Infrastructure. You can also update the description and tags of a source.
type UpdateSourceDetails struct {

	// Name of the source. This helps you to identify the appropriate source environment when you have multiple sources defined.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Description of the source. This helps you to identify the appropriate source environment when you have multiple sources defined.
	Description *string `mandatory:"false" json:"description"`

	SourceDetails SourceDetails `mandatory:"false" json:"sourceDetails"`

	AuthorizationDetails AuthorizationDetails `mandatory:"false" json:"authorizationDetails"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm). Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m UpdateSourceDetails) String() string {
	return common.PointerString(m)
}

// UnmarshalJSON unmarshals from json
func (m *UpdateSourceDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		DisplayName          *string                           `json:"displayName"`
		Description          *string                           `json:"description"`
		SourceDetails        sourcedetails                     `json:"sourceDetails"`
		AuthorizationDetails authorizationdetails              `json:"authorizationDetails"`
		FreeformTags         map[string]string                 `json:"freeformTags"`
		DefinedTags          map[string]map[string]interface{} `json:"definedTags"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.DisplayName = model.DisplayName

	m.Description = model.Description

	nn, e = model.SourceDetails.UnmarshalPolymorphicJSON(model.SourceDetails.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.SourceDetails = nn.(SourceDetails)
	} else {
		m.SourceDetails = nil
	}

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

	return
}
