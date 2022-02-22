// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Labeling API
//
// Use Data Labeling API to create Annotations on Images, Texts & Documents, and generate snapshots.
//

package datalabelingservicedataplane

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v59/common"
	"strings"
)

// CreateRecordDetails A record represents an entry in a dataset that needs labeling.
type CreateRecordDetails struct {

	// The name is automatically assigned by the service. It is unique and immutable.
	Name *string `mandatory:"true" json:"name"`

	// The OCID of the dataset to associate the record with.
	DatasetId *string `mandatory:"true" json:"datasetId"`

	// The OCID of the compartment for the record. This is tied to the dataset. It is not changeable on the record itself.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	SourceDetails CreateSourceDetails `mandatory:"true" json:"sourceDetails"`

	RecordMetadata RecordMetadata `mandatory:"false" json:"recordMetadata"`

	// A simple key-value pair that is applied without any predefined name, type, or scope. It exists for cross-compatibility only.
	// For example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// The defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m CreateRecordDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateRecordDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *CreateRecordDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		RecordMetadata recordmetadata                    `json:"recordMetadata"`
		FreeformTags   map[string]string                 `json:"freeformTags"`
		DefinedTags    map[string]map[string]interface{} `json:"definedTags"`
		Name           *string                           `json:"name"`
		DatasetId      *string                           `json:"datasetId"`
		CompartmentId  *string                           `json:"compartmentId"`
		SourceDetails  createsourcedetails               `json:"sourceDetails"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	nn, e = model.RecordMetadata.UnmarshalPolymorphicJSON(model.RecordMetadata.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.RecordMetadata = nn.(RecordMetadata)
	} else {
		m.RecordMetadata = nil
	}

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
