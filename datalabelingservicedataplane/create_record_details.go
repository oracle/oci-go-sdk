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
	"github.com/oracle/oci-go-sdk/v50/common"
)

// CreateRecordDetails A record represents an entry in a dataset that needs labeling.
type CreateRecordDetails struct {

	// This will be automatically assigned by the service. It will be unique and immutable.
	Name *string `mandatory:"true" json:"name"`

	// The OCID of the dataset to associate the record with.
	DatasetId *string `mandatory:"true" json:"datasetId"`

	// The OCID of the compartment for the record. This is tied to the dataset. It is not changeable on the record itself.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	SourceDetails CreateSourceDetails `mandatory:"true" json:"sourceDetails"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m CreateRecordDetails) String() string {
	return common.PointerString(m)
}

// UnmarshalJSON unmarshals from json
func (m *CreateRecordDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		FreeformTags  map[string]string                 `json:"freeformTags"`
		DefinedTags   map[string]map[string]interface{} `json:"definedTags"`
		Name          *string                           `json:"name"`
		DatasetId     *string                           `json:"datasetId"`
		CompartmentId *string                           `json:"compartmentId"`
		SourceDetails createsourcedetails               `json:"sourceDetails"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.FreeformTags = model.FreeformTags

	m.DefinedTags = model.DefinedTags

	m.Name = model.Name

	m.DatasetId = model.DatasetId

	m.CompartmentId = model.CompartmentId

	nn, e = model.SourceDetails.UnmarshalPolymorphicJSON(model.SourceDetails.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.SourceDetails = nn.(CreateSourceDetails)
	} else {
		m.SourceDetails = nil
	}

	return
}
