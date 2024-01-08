// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
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

// TranscriptionTaskSummary Summary of the Transcription Task.
type TranscriptionTaskSummary struct {

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the task.
	Id *string `mandatory:"true" json:"id"`

	// A user-friendly display name for the task.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// How much progress the operation has made, vs the total amount of work that must be performed.
	PercentComplete *int `mandatory:"false" json:"percentComplete"`

	// Size of input file in Bytes.
	FileSizeInBytes *int `mandatory:"false" json:"fileSizeInBytes"`

	// Duration of input file in Seconds.
	FileDurationInSeconds *int `mandatory:"false" json:"fileDurationInSeconds"`

	// Task proccessing duration, which excludes waiting time in the system.
	ProcessingDurationInSeconds *int `mandatory:"false" json:"processingDurationInSeconds"`

	// Task started time
	TimeStarted *common.SDKTime `mandatory:"false" json:"timeStarted"`

	// Job finished time
	TimeFinished *common.SDKTime `mandatory:"false" json:"timeFinished"`

	// The current state of the Speech Job.
	LifecycleState TranscriptionTaskLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	// A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`
}

func (m TranscriptionTaskSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m TranscriptionTaskSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingTranscriptionTaskLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetTranscriptionTaskLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
