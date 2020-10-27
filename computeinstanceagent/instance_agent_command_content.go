// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// InstanceAgentService API
//
// Instance Agent Service API
//

package computeinstanceagent

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/v27/common"
)

// InstanceAgentCommandContent Command content.
type InstanceAgentCommandContent struct {
	Source InstanceAgentCommandSourceDetails `mandatory:"true" json:"source"`

	Output InstanceAgentCommandOutputDetails `mandatory:"false" json:"output"`
}

func (m InstanceAgentCommandContent) String() string {
	return common.PointerString(m)
}

// UnmarshalJSON unmarshals from json
func (m *InstanceAgentCommandContent) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Output instanceagentcommandoutputdetails `json:"output"`
		Source instanceagentcommandsourcedetails `json:"source"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	nn, e = model.Output.UnmarshalPolymorphicJSON(model.Output.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.Output = nn.(InstanceAgentCommandOutputDetails)
	} else {
		m.Output = nil
	}

	nn, e = model.Source.UnmarshalPolymorphicJSON(model.Source.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.Source = nn.(InstanceAgentCommandSourceDetails)
	} else {
		m.Source = nil
	}

	return
}
