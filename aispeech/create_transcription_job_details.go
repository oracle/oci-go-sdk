// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
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

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment where you want to create the job.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	InputLocation InputLocation `mandatory:"true" json:"inputLocation"`

	OutputLocation *OutputLocation `mandatory:"true" json:"outputLocation"`

	// A user-friendly display name for the job.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// A short description of the job.
	Description *string `mandatory:"false" json:"description"`

	// Transcription Format. By default, the JSON format is used.
	AdditionalTranscriptionFormats []CreateTranscriptionJobDetailsAdditionalTranscriptionFormatsEnum `mandatory:"false" json:"additionalTranscriptionFormats,omitempty"`

	ModelDetails *TranscriptionModelDetails `mandatory:"false" json:"modelDetails"`

	Normalization *TranscriptionNormalization `mandatory:"false" json:"normalization"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`.
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace-1": {"bar-key-1": "value-1", "bar-key-2": "value-2"}, "foo-namespace-2": {"bar-key-1": "value-1", "bar-key-2": "value-2"}}`.
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

	for _, val := range m.AdditionalTranscriptionFormats {
		if _, ok := GetMappingCreateTranscriptionJobDetailsAdditionalTranscriptionFormatsEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for AdditionalTranscriptionFormats: %s. Supported values are: %s.", val, strings.Join(GetCreateTranscriptionJobDetailsAdditionalTranscriptionFormatsEnumStringValues(), ",")))
		}
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *CreateTranscriptionJobDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		DisplayName                    *string                                                           `json:"displayName"`
		Description                    *string                                                           `json:"description"`
		AdditionalTranscriptionFormats []CreateTranscriptionJobDetailsAdditionalTranscriptionFormatsEnum `json:"additionalTranscriptionFormats"`
		ModelDetails                   *TranscriptionModelDetails                                        `json:"modelDetails"`
		Normalization                  *TranscriptionNormalization                                       `json:"normalization"`
		FreeformTags                   map[string]string                                                 `json:"freeformTags"`
		DefinedTags                    map[string]map[string]interface{}                                 `json:"definedTags"`
		CompartmentId                  *string                                                           `json:"compartmentId"`
		InputLocation                  inputlocation                                                     `json:"inputLocation"`
		OutputLocation                 *OutputLocation                                                   `json:"outputLocation"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.DisplayName = model.DisplayName

	m.Description = model.Description

	m.AdditionalTranscriptionFormats = make([]CreateTranscriptionJobDetailsAdditionalTranscriptionFormatsEnum, len(model.AdditionalTranscriptionFormats))
	copy(m.AdditionalTranscriptionFormats, model.AdditionalTranscriptionFormats)
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

// CreateTranscriptionJobDetailsAdditionalTranscriptionFormatsEnum Enum with underlying type: string
type CreateTranscriptionJobDetailsAdditionalTranscriptionFormatsEnum string

// Set of constants representing the allowable values for CreateTranscriptionJobDetailsAdditionalTranscriptionFormatsEnum
const (
	CreateTranscriptionJobDetailsAdditionalTranscriptionFormatsSrt CreateTranscriptionJobDetailsAdditionalTranscriptionFormatsEnum = "SRT"
)

var mappingCreateTranscriptionJobDetailsAdditionalTranscriptionFormatsEnum = map[string]CreateTranscriptionJobDetailsAdditionalTranscriptionFormatsEnum{
	"SRT": CreateTranscriptionJobDetailsAdditionalTranscriptionFormatsSrt,
}

var mappingCreateTranscriptionJobDetailsAdditionalTranscriptionFormatsEnumLowerCase = map[string]CreateTranscriptionJobDetailsAdditionalTranscriptionFormatsEnum{
	"srt": CreateTranscriptionJobDetailsAdditionalTranscriptionFormatsSrt,
}

// GetCreateTranscriptionJobDetailsAdditionalTranscriptionFormatsEnumValues Enumerates the set of values for CreateTranscriptionJobDetailsAdditionalTranscriptionFormatsEnum
func GetCreateTranscriptionJobDetailsAdditionalTranscriptionFormatsEnumValues() []CreateTranscriptionJobDetailsAdditionalTranscriptionFormatsEnum {
	values := make([]CreateTranscriptionJobDetailsAdditionalTranscriptionFormatsEnum, 0)
	for _, v := range mappingCreateTranscriptionJobDetailsAdditionalTranscriptionFormatsEnum {
		values = append(values, v)
	}
	return values
}

// GetCreateTranscriptionJobDetailsAdditionalTranscriptionFormatsEnumStringValues Enumerates the set of values in String for CreateTranscriptionJobDetailsAdditionalTranscriptionFormatsEnum
func GetCreateTranscriptionJobDetailsAdditionalTranscriptionFormatsEnumStringValues() []string {
	return []string{
		"SRT",
	}
}

// GetMappingCreateTranscriptionJobDetailsAdditionalTranscriptionFormatsEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCreateTranscriptionJobDetailsAdditionalTranscriptionFormatsEnum(val string) (CreateTranscriptionJobDetailsAdditionalTranscriptionFormatsEnum, bool) {
	enum, ok := mappingCreateTranscriptionJobDetailsAdditionalTranscriptionFormatsEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
