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

// SynthesizeSpeechDetails Input JSON to get audio inference from TTS Service.
type SynthesizeSpeechDetails struct {

	// The text input to get the inference audio from TTS Service.
	Text *string `mandatory:"true" json:"text"`

	// If set to true, response will be sent in the chunked transfer-encoding and audio chunks
	// are sent back as and when they are ready. If set to false, response will be sent only once
	// the entire audio is generated.
	IsStreamEnabled *bool `mandatory:"false" json:"isStreamEnabled"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment where the user has access to call `SpeechSynthesize` api. But default user access will be checked at tenancy level.
	CompartmentId *string `mandatory:"false" json:"compartmentId"`

	Configuration TtsConfiguration `mandatory:"false" json:"configuration"`

	AudioConfig TtsAudioConfig `mandatory:"false" json:"audioConfig"`
}

func (m SynthesizeSpeechDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m SynthesizeSpeechDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *SynthesizeSpeechDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		IsStreamEnabled *bool            `json:"isStreamEnabled"`
		CompartmentId   *string          `json:"compartmentId"`
		Configuration   ttsconfiguration `json:"configuration"`
		AudioConfig     ttsaudioconfig   `json:"audioConfig"`
		Text            *string          `json:"text"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.IsStreamEnabled = model.IsStreamEnabled

	m.CompartmentId = model.CompartmentId

	nn, e = model.Configuration.UnmarshalPolymorphicJSON(model.Configuration.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.Configuration = nn.(TtsConfiguration)
	} else {
		m.Configuration = nil
	}

	nn, e = model.AudioConfig.UnmarshalPolymorphicJSON(model.AudioConfig.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.AudioConfig = nn.(TtsAudioConfig)
	} else {
		m.AudioConfig = nil
	}

	m.Text = model.Text

	return
}
