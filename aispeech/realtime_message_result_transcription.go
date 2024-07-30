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

// RealtimeMessageResultTranscription Transcription object.
type RealtimeMessageResultTranscription struct {

	// Transcription text.
	Transcription *string `mandatory:"true" json:"transcription"`

	// Whether the transcription is final or partial.
	IsFinal *bool `mandatory:"true" json:"isFinal"`

	// Start time in milliseconds for the transcription text.
	StartTimeInMs *int `mandatory:"true" json:"startTimeInMs"`

	// End time in milliseconds for the transcription text.
	EndTimeInMs *int `mandatory:"true" json:"endTimeInMs"`

	// Confidence for the transcription text.
	Confidence *float32 `mandatory:"true" json:"confidence"`

	// Trailing silence after the transcription text.
	TrailingSilence *int `mandatory:"true" json:"trailingSilence"`

	// Array of individual transcription tokens.
	Tokens []RealtimeMessageResultTranscriptionToken `mandatory:"true" json:"tokens"`
}

func (m RealtimeMessageResultTranscription) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m RealtimeMessageResultTranscription) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
