// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Speech API
//
// The OCI Speech Service harnesses the power of spoken language by allowing developers to easily convert file-based data containing human speech into highly accurate text transcriptions.
//

package aispeech

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// UpdateCustomizationDetails The information to be updated.
type UpdateCustomizationDetails struct {

	// Customization Details Alias
	Alias *string `mandatory:"false" json:"alias"`

	// A user-friendly display name for the customization.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// A short description of the customization.
	Description *string `mandatory:"false" json:"description"`

	ModelDetails *CustomizationModelDetails `mandatory:"false" json:"modelDetails"`

	TrainingDataset CustomizationDatasetDetails `mandatory:"false" json:"trainingDataset"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`.
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace-1": {"bar-key-1": "value-1", "bar-key-2": "value-2"}, "foo-namespace-2": {"bar-key-1": "value-1", "bar-key-2": "value-2"}}`.
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m UpdateCustomizationDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpdateCustomizationDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *UpdateCustomizationDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Alias           *string                           `json:"alias"`
		DisplayName     *string                           `json:"displayName"`
		Description     *string                           `json:"description"`
		ModelDetails    *CustomizationModelDetails        `json:"modelDetails"`
		TrainingDataset customizationdatasetdetails       `json:"trainingDataset"`
		FreeformTags    map[string]string                 `json:"freeformTags"`
		DefinedTags     map[string]map[string]interface{} `json:"definedTags"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.Alias = model.Alias

	m.DisplayName = model.DisplayName

	m.Description = model.Description

	m.ModelDetails = model.ModelDetails

	nn, e = model.TrainingDataset.UnmarshalPolymorphicJSON(model.TrainingDataset.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.TrainingDataset = nn.(CustomizationDatasetDetails)
	} else {
		m.TrainingDataset = nil
	}

	m.FreeformTags = model.FreeformTags

	m.DefinedTags = model.DefinedTags

	return
}
