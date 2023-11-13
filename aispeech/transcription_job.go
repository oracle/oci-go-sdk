// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
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

// TranscriptionJob Description of Transcription Job.
type TranscriptionJob struct {

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the job.
	Id *string `mandatory:"true" json:"id"`

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment where you want to create the job.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	ModelDetails *TranscriptionModelDetails `mandatory:"true" json:"modelDetails"`

	InputLocation InputLocation `mandatory:"true" json:"inputLocation"`

	OutputLocation *OutputLocation `mandatory:"true" json:"outputLocation"`

	// A user-friendly display name for the job.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// A short description of the job.
	Description *string `mandatory:"false" json:"description"`

	Normalization *TranscriptionNormalization `mandatory:"false" json:"normalization"`

	// Job accepted time.
	TimeAccepted *common.SDKTime `mandatory:"false" json:"timeAccepted"`

	// Job started time.
	TimeStarted *common.SDKTime `mandatory:"false" json:"timeStarted"`

	// Job finished time.
	TimeFinished *common.SDKTime `mandatory:"false" json:"timeFinished"`

	// Total tasks in a job.
	TotalTasks *int `mandatory:"false" json:"totalTasks"`

	// Total outstanding tasks in a job.
	OutstandingTasks *int `mandatory:"false" json:"outstandingTasks"`

	// Total successful tasks in a job.
	SuccessfulTasks *int `mandatory:"false" json:"successfulTasks"`

	// Time to live duration in days for Job. Job will be available till max 90 days.
	TtlInDays *int `mandatory:"false" json:"ttlInDays"`

	// How much progress the operation has made, vs the total amount of work that must be performed.
	PercentComplete *int `mandatory:"false" json:"percentComplete"`

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the user who created the job.
	CreatedBy *string `mandatory:"false" json:"createdBy"`

	// Transcription format. JSON format will always be provided in addition to any formats in this list.
	AdditionalTranscriptionFormats []TranscriptionJobAdditionalTranscriptionFormatsEnum `mandatory:"false" json:"additionalTranscriptionFormats,omitempty"`

	// The current state of the Job.
	LifecycleState TranscriptionJobLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	// A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`.
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace-1": {"bar-key-1": "value-1", "bar-key-2": "value-2"}, "foo-namespace-2": {"bar-key-1": "value-1", "bar-key-2": "value-2"}}`.
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// Usage of system tag keys. These predefined keys are scoped to namespaces.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`.
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
}

func (m TranscriptionJob) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m TranscriptionJob) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	for _, val := range m.AdditionalTranscriptionFormats {
		if _, ok := GetMappingTranscriptionJobAdditionalTranscriptionFormatsEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for AdditionalTranscriptionFormats: %s. Supported values are: %s.", val, strings.Join(GetTranscriptionJobAdditionalTranscriptionFormatsEnumStringValues(), ",")))
		}
	}

	if _, ok := GetMappingTranscriptionJobLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetTranscriptionJobLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *TranscriptionJob) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		DisplayName                    *string                                              `json:"displayName"`
		Description                    *string                                              `json:"description"`
		Normalization                  *TranscriptionNormalization                          `json:"normalization"`
		TimeAccepted                   *common.SDKTime                                      `json:"timeAccepted"`
		TimeStarted                    *common.SDKTime                                      `json:"timeStarted"`
		TimeFinished                   *common.SDKTime                                      `json:"timeFinished"`
		TotalTasks                     *int                                                 `json:"totalTasks"`
		OutstandingTasks               *int                                                 `json:"outstandingTasks"`
		SuccessfulTasks                *int                                                 `json:"successfulTasks"`
		TtlInDays                      *int                                                 `json:"ttlInDays"`
		PercentComplete                *int                                                 `json:"percentComplete"`
		CreatedBy                      *string                                              `json:"createdBy"`
		AdditionalTranscriptionFormats []TranscriptionJobAdditionalTranscriptionFormatsEnum `json:"additionalTranscriptionFormats"`
		LifecycleState                 TranscriptionJobLifecycleStateEnum                   `json:"lifecycleState"`
		LifecycleDetails               *string                                              `json:"lifecycleDetails"`
		FreeformTags                   map[string]string                                    `json:"freeformTags"`
		DefinedTags                    map[string]map[string]interface{}                    `json:"definedTags"`
		SystemTags                     map[string]map[string]interface{}                    `json:"systemTags"`
		Id                             *string                                              `json:"id"`
		CompartmentId                  *string                                              `json:"compartmentId"`
		ModelDetails                   *TranscriptionModelDetails                           `json:"modelDetails"`
		InputLocation                  inputlocation                                        `json:"inputLocation"`
		OutputLocation                 *OutputLocation                                      `json:"outputLocation"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.DisplayName = model.DisplayName

	m.Description = model.Description

	m.Normalization = model.Normalization

	m.TimeAccepted = model.TimeAccepted

	m.TimeStarted = model.TimeStarted

	m.TimeFinished = model.TimeFinished

	m.TotalTasks = model.TotalTasks

	m.OutstandingTasks = model.OutstandingTasks

	m.SuccessfulTasks = model.SuccessfulTasks

	m.TtlInDays = model.TtlInDays

	m.PercentComplete = model.PercentComplete

	m.CreatedBy = model.CreatedBy

	m.AdditionalTranscriptionFormats = make([]TranscriptionJobAdditionalTranscriptionFormatsEnum, len(model.AdditionalTranscriptionFormats))
	copy(m.AdditionalTranscriptionFormats, model.AdditionalTranscriptionFormats)
	m.LifecycleState = model.LifecycleState

	m.LifecycleDetails = model.LifecycleDetails

	m.FreeformTags = model.FreeformTags

	m.DefinedTags = model.DefinedTags

	m.SystemTags = model.SystemTags

	m.Id = model.Id

	m.CompartmentId = model.CompartmentId

	m.ModelDetails = model.ModelDetails

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

// TranscriptionJobAdditionalTranscriptionFormatsEnum Enum with underlying type: string
type TranscriptionJobAdditionalTranscriptionFormatsEnum string

// Set of constants representing the allowable values for TranscriptionJobAdditionalTranscriptionFormatsEnum
const (
	TranscriptionJobAdditionalTranscriptionFormatsSrt TranscriptionJobAdditionalTranscriptionFormatsEnum = "SRT"
)

var mappingTranscriptionJobAdditionalTranscriptionFormatsEnum = map[string]TranscriptionJobAdditionalTranscriptionFormatsEnum{
	"SRT": TranscriptionJobAdditionalTranscriptionFormatsSrt,
}

var mappingTranscriptionJobAdditionalTranscriptionFormatsEnumLowerCase = map[string]TranscriptionJobAdditionalTranscriptionFormatsEnum{
	"srt": TranscriptionJobAdditionalTranscriptionFormatsSrt,
}

// GetTranscriptionJobAdditionalTranscriptionFormatsEnumValues Enumerates the set of values for TranscriptionJobAdditionalTranscriptionFormatsEnum
func GetTranscriptionJobAdditionalTranscriptionFormatsEnumValues() []TranscriptionJobAdditionalTranscriptionFormatsEnum {
	values := make([]TranscriptionJobAdditionalTranscriptionFormatsEnum, 0)
	for _, v := range mappingTranscriptionJobAdditionalTranscriptionFormatsEnum {
		values = append(values, v)
	}
	return values
}

// GetTranscriptionJobAdditionalTranscriptionFormatsEnumStringValues Enumerates the set of values in String for TranscriptionJobAdditionalTranscriptionFormatsEnum
func GetTranscriptionJobAdditionalTranscriptionFormatsEnumStringValues() []string {
	return []string{
		"SRT",
	}
}

// GetMappingTranscriptionJobAdditionalTranscriptionFormatsEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingTranscriptionJobAdditionalTranscriptionFormatsEnum(val string) (TranscriptionJobAdditionalTranscriptionFormatsEnum, bool) {
	enum, ok := mappingTranscriptionJobAdditionalTranscriptionFormatsEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// TranscriptionJobLifecycleStateEnum Enum with underlying type: string
type TranscriptionJobLifecycleStateEnum string

// Set of constants representing the allowable values for TranscriptionJobLifecycleStateEnum
const (
	TranscriptionJobLifecycleStateAccepted   TranscriptionJobLifecycleStateEnum = "ACCEPTED"
	TranscriptionJobLifecycleStateInProgress TranscriptionJobLifecycleStateEnum = "IN_PROGRESS"
	TranscriptionJobLifecycleStateSucceeded  TranscriptionJobLifecycleStateEnum = "SUCCEEDED"
	TranscriptionJobLifecycleStateFailed     TranscriptionJobLifecycleStateEnum = "FAILED"
	TranscriptionJobLifecycleStateCanceling  TranscriptionJobLifecycleStateEnum = "CANCELING"
	TranscriptionJobLifecycleStateCanceled   TranscriptionJobLifecycleStateEnum = "CANCELED"
)

var mappingTranscriptionJobLifecycleStateEnum = map[string]TranscriptionJobLifecycleStateEnum{
	"ACCEPTED":    TranscriptionJobLifecycleStateAccepted,
	"IN_PROGRESS": TranscriptionJobLifecycleStateInProgress,
	"SUCCEEDED":   TranscriptionJobLifecycleStateSucceeded,
	"FAILED":      TranscriptionJobLifecycleStateFailed,
	"CANCELING":   TranscriptionJobLifecycleStateCanceling,
	"CANCELED":    TranscriptionJobLifecycleStateCanceled,
}

var mappingTranscriptionJobLifecycleStateEnumLowerCase = map[string]TranscriptionJobLifecycleStateEnum{
	"accepted":    TranscriptionJobLifecycleStateAccepted,
	"in_progress": TranscriptionJobLifecycleStateInProgress,
	"succeeded":   TranscriptionJobLifecycleStateSucceeded,
	"failed":      TranscriptionJobLifecycleStateFailed,
	"canceling":   TranscriptionJobLifecycleStateCanceling,
	"canceled":    TranscriptionJobLifecycleStateCanceled,
}

// GetTranscriptionJobLifecycleStateEnumValues Enumerates the set of values for TranscriptionJobLifecycleStateEnum
func GetTranscriptionJobLifecycleStateEnumValues() []TranscriptionJobLifecycleStateEnum {
	values := make([]TranscriptionJobLifecycleStateEnum, 0)
	for _, v := range mappingTranscriptionJobLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetTranscriptionJobLifecycleStateEnumStringValues Enumerates the set of values in String for TranscriptionJobLifecycleStateEnum
func GetTranscriptionJobLifecycleStateEnumStringValues() []string {
	return []string{
		"ACCEPTED",
		"IN_PROGRESS",
		"SUCCEEDED",
		"FAILED",
		"CANCELING",
		"CANCELED",
	}
}

// GetMappingTranscriptionJobLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingTranscriptionJobLifecycleStateEnum(val string) (TranscriptionJobLifecycleStateEnum, bool) {
	enum, ok := mappingTranscriptionJobLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
