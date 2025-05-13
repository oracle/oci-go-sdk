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

// RealtimeMessageResult The websocket result message received from service.
type RealtimeMessageResult struct {

	// List of transcription objects.
	Transcriptions []RealtimeMessageResultTranscription `mandatory:"true" json:"transcriptions"`

	// Session ID for the connected session.
	SessionId *string `mandatory:"false" json:"sessionId"`
}

// GetSessionId returns SessionId
func (m RealtimeMessageResult) GetSessionId() *string {
	return m.SessionId
}

func (m RealtimeMessageResult) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m RealtimeMessageResult) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m RealtimeMessageResult) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeRealtimeMessageResult RealtimeMessageResult
	s := struct {
		DiscriminatorParam string `json:"event"`
		MarshalTypeRealtimeMessageResult
	}{
		"RESULT",
		(MarshalTypeRealtimeMessageResult)(m),
	}

	return json.Marshal(&s)
}
