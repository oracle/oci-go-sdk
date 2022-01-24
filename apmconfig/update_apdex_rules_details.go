// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Apm Configuration API
//
// An API for the APM Configuration service. Use this API to query and set APM configuration.
//

package apmconfig

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/v56/common"
)

// UpdateApdexRulesDetails The set of Apdex rules to be used in Apdex computation.
type UpdateApdexRulesDetails struct {
	Rules []Apdex `mandatory:"true" json:"rules"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// The name by which this rule set can be displayed to the user.
	DisplayName *string `mandatory:"false" json:"displayName"`
}

//GetFreeformTags returns FreeformTags
func (m UpdateApdexRulesDetails) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

//GetDefinedTags returns DefinedTags
func (m UpdateApdexRulesDetails) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

func (m UpdateApdexRulesDetails) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m UpdateApdexRulesDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeUpdateApdexRulesDetails UpdateApdexRulesDetails
	s := struct {
		DiscriminatorParam string `json:"configType"`
		MarshalTypeUpdateApdexRulesDetails
	}{
		"APDEX",
		(MarshalTypeUpdateApdexRulesDetails)(m),
	}

	return json.Marshal(&s)
}
