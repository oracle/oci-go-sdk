// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
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

// CreateTranscriptionJobDetails The information about new Transcription Job.
type CreateTranscriptionJobDetails struct {

	// The OCID of the compartment that contains the transcriptionJob.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	InputLocation InputLocation `mandatory:"true" json:"inputLocation"`

	OutputLocation *OutputLocation `mandatory:"true" json:"outputLocation"`

	// Transcription job name.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Transcription job description.
	Description *string `mandatory:"false" json:"description"`

	ModelDetails *TranscriptionModelDetails `mandatory:"false" json:"modelDetails"`

	Normalization *TranscriptionNormalization `mandatory:"false" json:"normalization"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m CreateTranscriptionJobDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateTranscriptionJobDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *CreateTranscriptionJobDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		DisplayName    *string                           `json:"displayName"`
		Description    *string                           `json:"description"`
		ModelDetails   *TranscriptionModelDetails        `json:"modelDetails"`
		Normalization  *TranscriptionNormalization       `json:"normalization"`
		FreeformTags   map[string]string                 `json:"freeformTags"`
		DefinedTags    map[string]map[string]interface{} `json:"definedTags"`
		CompartmentId  *string                           `json:"compartmentId"`
		InputLocation  inputlocation                     `json:"inputLocation"`
		OutputLocation *OutputLocation                   `json:"outputLocation"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.DisplayName = model.DisplayName

	m.Description = model.Description

	m.ModelDetails = model.ModelDetails

	m.Normalization = model.Normalization

	m.FreeformTags = model.FreeformTags

	m.DefinedTags = model.DefinedTags

	m.CompartmentId = model.CompartmentId

	nn, e = model.InputLocation.UnmarshalPolymorphicJSON(model.InputLocation.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.InputLocation = nn.(InputLocation)
	} else {
		m.InputLocation = nil
	}

	m.OutputLocation = model.OutputLocation

	return
}
