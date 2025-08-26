// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Speech API
//
// The OCI Speech Service harnesses the power of spoken language by allowing developers to easily convert file-based data containing human speech into highly accurate text transcriptions.
//

package aispeech

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// TranscriptionTask Description of Transcription Task.
type TranscriptionTask struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the task.
	Id *string `mandatory:"true" json:"id"`

	// A user-friendly display name for the task.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Task started time.
	TimeStarted *common.SDKTime `mandatory:"false" json:"timeStarted"`

	// Task finished time.
	TimeFinished *common.SDKTime `mandatory:"false" json:"timeFinished"`

	// How much progress the operation has made, vs the total amount of work that must be performed.
	PercentComplete *int `mandatory:"false" json:"percentComplete"`

	// Time to live duration in days for tasks. Task will be available till max 90 days.
	TtlInDays *int `mandatory:"false" json:"ttlInDays"`

	ModelDetails *TranscriptionModelDetails `mandatory:"false" json:"modelDetails"`

	AudioFormatDetails *AudioFormatDetails `mandatory:"false" json:"audioFormatDetails"`

	// Size of input file in Bytes.
	FileSizeInBytes *int `mandatory:"false" json:"fileSizeInBytes"`

	// Duration of input file in Seconds.
	FileDurationInSeconds *int `mandatory:"false" json:"fileDurationInSeconds"`

	// Task proccessing duration, which excludes waiting time in the system.
	ProcessingDurationInSeconds *int `mandatory:"false" json:"processingDurationInSeconds"`

	InputLocation *ObjectLocation `mandatory:"false" json:"inputLocation"`

	OutputLocation *ObjectLocation `mandatory:"false" json:"outputLocation"`

	// The current state of the Task.
	LifecycleState TranscriptionTaskLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	// A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`
}

func (m TranscriptionTask) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m TranscriptionTask) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingTranscriptionTaskLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetTranscriptionTaskLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// TranscriptionTaskLifecycleStateEnum Enum with underlying type: string
type TranscriptionTaskLifecycleStateEnum string

// Set of constants representing the allowable values for TranscriptionTaskLifecycleStateEnum
const (
	TranscriptionTaskLifecycleStateAccepted   TranscriptionTaskLifecycleStateEnum = "ACCEPTED"
	TranscriptionTaskLifecycleStateInProgress TranscriptionTaskLifecycleStateEnum = "IN_PROGRESS"
	TranscriptionTaskLifecycleStateSucceeded  TranscriptionTaskLifecycleStateEnum = "SUCCEEDED"
	TranscriptionTaskLifecycleStateFailed     TranscriptionTaskLifecycleStateEnum = "FAILED"
	TranscriptionTaskLifecycleStateCanceled   TranscriptionTaskLifecycleStateEnum = "CANCELED"
)

var mappingTranscriptionTaskLifecycleStateEnum = map[string]TranscriptionTaskLifecycleStateEnum{
	"ACCEPTED":    TranscriptionTaskLifecycleStateAccepted,
	"IN_PROGRESS": TranscriptionTaskLifecycleStateInProgress,
	"SUCCEEDED":   TranscriptionTaskLifecycleStateSucceeded,
	"FAILED":      TranscriptionTaskLifecycleStateFailed,
	"CANCELED":    TranscriptionTaskLifecycleStateCanceled,
}

var mappingTranscriptionTaskLifecycleStateEnumLowerCase = map[string]TranscriptionTaskLifecycleStateEnum{
	"accepted":    TranscriptionTaskLifecycleStateAccepted,
	"in_progress": TranscriptionTaskLifecycleStateInProgress,
	"succeeded":   TranscriptionTaskLifecycleStateSucceeded,
	"failed":      TranscriptionTaskLifecycleStateFailed,
	"canceled":    TranscriptionTaskLifecycleStateCanceled,
}

// GetTranscriptionTaskLifecycleStateEnumValues Enumerates the set of values for TranscriptionTaskLifecycleStateEnum
func GetTranscriptionTaskLifecycleStateEnumValues() []TranscriptionTaskLifecycleStateEnum {
	values := make([]TranscriptionTaskLifecycleStateEnum, 0)
	for _, v := range mappingTranscriptionTaskLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetTranscriptionTaskLifecycleStateEnumStringValues Enumerates the set of values in String for TranscriptionTaskLifecycleStateEnum
func GetTranscriptionTaskLifecycleStateEnumStringValues() []string {
	return []string{
		"ACCEPTED",
		"IN_PROGRESS",
		"SUCCEEDED",
		"FAILED",
		"CANCELED",
	}
}

// GetMappingTranscriptionTaskLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingTranscriptionTaskLifecycleStateEnum(val string) (TranscriptionTaskLifecycleStateEnum, bool) {
	enum, ok := mappingTranscriptionTaskLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
