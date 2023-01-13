// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
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

// EnforcedGovernanceRule Represents the governance rule shown to the child which is a subset of governance rule resource in parent tenancy.
type EnforcedGovernanceRule struct {

	// The Oracle ID (OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm)) of the enforced governance rule.
	Id *string `mandatory:"true" json:"id"`

	// The Oracle ID (OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm)) of the child's root compartment to which the governance rule is attached.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// Display name of the governance rule.
	GovernanceRuleDisplayName *string `mandatory:"true" json:"governanceRuleDisplayName"`

	// Type of the governance rule, can be one of QUOTA, TAG, ALLOWED_REGIONS.
	// Example: `QUOTA`
	Type GovernanceRuleTypeEnum `mandatory:"true" json:"type"`

	Template Template `mandatory:"true" json:"template"`

	// The current state of the governance rule.
	LifecycleState GovernanceRuleLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// Date and time the governance rule was created. An RFC3339 formatted datetime string.
	// Example: `2016-08-25T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// Date and time the governance rule was updated. An RFC3339 formatted datetime string.
	// Example: `2016-08-25T21:10:29.600Z`
	TimeUpdated *common.SDKTime `mandatory:"true" json:"timeUpdated"`
}

func (m EnforcedGovernanceRule) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m EnforcedGovernanceRule) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingGovernanceRuleTypeEnum(string(m.Type)); !ok && m.Type != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Type: %s. Supported values are: %s.", m.Type, strings.Join(GetGovernanceRuleTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingGovernanceRuleLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetGovernanceRuleLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *EnforcedGovernanceRule) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Id                        *string                          `json:"id"`
		CompartmentId             *string                          `json:"compartmentId"`
		GovernanceRuleDisplayName *string                          `json:"governanceRuleDisplayName"`
		Type                      GovernanceRuleTypeEnum           `json:"type"`
		Template                  template                         `json:"template"`
		LifecycleState            GovernanceRuleLifecycleStateEnum `json:"lifecycleState"`
		TimeCreated               *common.SDKTime                  `json:"timeCreated"`
		TimeUpdated               *common.SDKTime                  `json:"timeUpdated"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.Id = model.Id

	m.CompartmentId = model.CompartmentId

	m.GovernanceRuleDisplayName = model.GovernanceRuleDisplayName

	m.Type = model.Type

	nn, e = model.Template.UnmarshalPolymorphicJSON(model.Template.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.Template = nn.(Template)
	} else {
		m.Template = nil
	}

	m.LifecycleState = model.LifecycleState

	m.TimeCreated = model.TimeCreated

	m.TimeUpdated = model.TimeUpdated

	return
}
