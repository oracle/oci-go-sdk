// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// DlsDataPlane API
//
// A description of the DlsDataPlane API.
//

package datalabelingservicedataplane

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/v52/common"
)

// CreateAnnotationDetails This is the payload sent in the CreateAnnotation operation.  It contains all the information required for a user to create an Annotation for a record.
type CreateAnnotationDetails struct {

	// The OCID of the record annotated
	RecordId *string `mandatory:"true" json:"recordId"`

	// The OCID of the compartment for the annotation
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The entity types will be validated against the dataset to ensure consistency.
	Entities []Entity `mandatory:"true" json:"entities"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m CreateAnnotationDetails) String() string {
	return common.PointerString(m)
}

// UnmarshalJSON unmarshals from json
func (m *CreateAnnotationDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		FreeformTags  map[string]string                 `json:"freeformTags"`
		DefinedTags   map[string]map[string]interface{} `json:"definedTags"`
		RecordId      *string                           `json:"recordId"`
		CompartmentId *string                           `json:"compartmentId"`
		Entities      []entity                          `json:"entities"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.FreeformTags = model.FreeformTags

	m.DefinedTags = model.DefinedTags

	m.RecordId = model.RecordId

	m.CompartmentId = model.CompartmentId

	m.Entities = make([]Entity, len(model.Entities))
	for i, n := range model.Entities {
		nn, e = n.UnmarshalPolymorphicJSON(n.JsonData)
		if e != nil {
			return e
		}
		if nn != nil {
			m.Entities[i] = nn.(Entity)
		} else {
			m.Entities[i] = nil
		}
	}

	return
}
