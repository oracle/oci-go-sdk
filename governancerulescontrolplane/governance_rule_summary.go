// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// GovernanceRulesControlPlane API
//
// A description of the GovernanceRulesControlPlane API
//

package governancerulescontrolplane

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// GovernanceRuleSummary A summary of the governance rule.
type GovernanceRuleSummary struct {

	// The Oracle ID (OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm)) of the governance rule.
	Id *string `mandatory:"true" json:"id"`

	// The Oracle ID (OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm)) of the root compartment containing the governance rule.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// Display name of the governance rule.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// Type of the governance rule, can be one of QUOTA, TAG, ALLOWED_REGIONS.
	// Example: `QUOTA`
	Type GovernanceRuleTypeEnum `mandatory:"true" json:"type"`

	// The type of option used to create the governance rule, could be one of TEMPLATE or CLONE.
	// Example: `TEMPLATE`
	CreationOption CreationOptionEnum `mandatory:"true" json:"creationOption"`

	// Date and time the governance rule was created. An RFC3339 formatted datetime string.
	// Example: `2016-08-25T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// Date and time the governance rule was updated. An RFC3339 formatted datetime string.
	// Example: `2016-08-25T21:10:29.600Z`
	TimeUpdated *common.SDKTime `mandatory:"true" json:"timeUpdated"`

	// The current state of the governance rule.
	LifecycleState GovernanceRuleLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"true" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"true" json:"definedTags"`

	// System tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
}

func (m GovernanceRuleSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m GovernanceRuleSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingGovernanceRuleTypeEnum(string(m.Type)); !ok && m.Type != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Type: %s. Supported values are: %s.", m.Type, strings.Join(GetGovernanceRuleTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingCreationOptionEnum(string(m.CreationOption)); !ok && m.CreationOption != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for CreationOption: %s. Supported values are: %s.", m.CreationOption, strings.Join(GetCreationOptionEnumStringValues(), ",")))
	}
	if _, ok := GetMappingGovernanceRuleLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetGovernanceRuleLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
