// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// GovernanceRulesControlPlane API
//
// A description of the GovernanceRulesControlPlane API
//

package governancerulescontrolplane

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// UpdateGovernanceRuleDetails Request object for UpdateGovernanceRule operation.
type UpdateGovernanceRuleDetails struct {

	// Display name of the governance rule.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Description of the governance rule.
	Description *string `mandatory:"false" json:"description"`

	Template Template `mandatory:"false" json:"template"`

	// The Oracle ID (OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm)) of the resource, which was used as a template to create this governance rule.
	RelatedResourceId *string `mandatory:"false" json:"relatedResourceId"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m UpdateGovernanceRuleDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpdateGovernanceRuleDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *UpdateGovernanceRuleDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		DisplayName       *string                           `json:"displayName"`
		Description       *string                           `json:"description"`
		Template          template                          `json:"template"`
		RelatedResourceId *string                           `json:"relatedResourceId"`
		FreeformTags      map[string]string                 `json:"freeformTags"`
		DefinedTags       map[string]map[string]interface{} `json:"definedTags"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.DisplayName = model.DisplayName

	m.Description = model.Description

	nn, e = model.Template.UnmarshalPolymorphicJSON(model.Template.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.Template = nn.(Template)
	} else {
		m.Template = nil
	}

	m.RelatedResourceId = model.RelatedResourceId

	m.FreeformTags = model.FreeformTags

	m.DefinedTags = model.DefinedTags

	return
}
