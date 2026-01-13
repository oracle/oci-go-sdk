// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
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

// Tag Details of the tag that is being created.
type Tag struct {

	// The name you assign to the tag during creation. This is the tag key definition.
	// The name must be unique within the tag namespace and cannot be changed.
	Name *string `mandatory:"true" json:"name"`

	// The description assigned to the tag during creation.
	Description *string `mandatory:"false" json:"description"`

	// Indicates whether the tag is enabled for cost tracking.
	IsCostTracking *bool `mandatory:"false" json:"isCostTracking"`

	Validator BaseTagDefinitionValidator `mandatory:"false" json:"validator"`
}

func (m Tag) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m Tag) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *Tag) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Description    *string                    `json:"description"`
		IsCostTracking *bool                      `json:"isCostTracking"`
		Validator      basetagdefinitionvalidator `json:"validator"`
		Name           *string                    `json:"name"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.Description = model.Description

	m.IsCostTracking = model.IsCostTracking

	nn, e = model.Validator.UnmarshalPolymorphicJSON(model.Validator.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.Validator = nn.(BaseTagDefinitionValidator)
	} else {
		m.Validator = nil
	}

	m.Name = model.Name

	return
}
