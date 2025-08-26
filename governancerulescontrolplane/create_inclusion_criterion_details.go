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

// CreateInclusionCriterionDetails Request object for Createinclusion criterion operation.
type CreateInclusionCriterionDetails struct {

	// The Oracle ID (OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm)) of the governance rule. Every inclusion criterion is associated with a governance rule.
	GovernanceRuleId *string `mandatory:"true" json:"governanceRuleId"`

	// Type of inclusion criterion - TENANCY, ALL or TAG. We support TENANCY and ALL for now.
	Type InclusionCriterionTypeEnum `mandatory:"true" json:"type"`

	Association Association `mandatory:"false" json:"association"`
}

func (m CreateInclusionCriterionDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateInclusionCriterionDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingInclusionCriterionTypeEnum(string(m.Type)); !ok && m.Type != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Type: %s. Supported values are: %s.", m.Type, strings.Join(GetInclusionCriterionTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *CreateInclusionCriterionDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Association      association                `json:"association"`
		GovernanceRuleId *string                    `json:"governanceRuleId"`
		Type             InclusionCriterionTypeEnum `json:"type"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	nn, e = model.Association.UnmarshalPolymorphicJSON(model.Association.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.Association = nn.(Association)
	} else {
		m.Association = nil
	}

	m.GovernanceRuleId = model.GovernanceRuleId

	m.Type = model.Type

	return
}
