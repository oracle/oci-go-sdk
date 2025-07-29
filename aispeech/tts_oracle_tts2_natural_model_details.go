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

// TtsOracleTts2NaturalModelDetails Use this schema for specifying properties of TTS_2_NATURAL model from Oracle model family.
type TtsOracleTts2NaturalModelDetails struct {

	// Speaker in whose voice the user wants the output speech to be in.
	// The possible values for `voiceId` can be obtained by calling ListVoices api.
	VoiceId *string `mandatory:"false" json:"voiceId"`

	// Locale value as per given in [https://datatracker.ietf.org/doc/html/rfc5646]. Default en-US
	// - en-US: English - United States
	// - en-GB: English - Great Britain
	// - es-ES: Spanish - Spain
	// - pt-BR: Portuguese - Brazil
	// - hi-IN: Hindi - India
	// - fr-FR: French - France
	// - it-IT: Italian - Italy
	// - ja-JP: Japanese - Japan
	// - zh-CN: Mandarin - China
	LanguageCode *string `mandatory:"false" json:"languageCode"`
}

func (m TtsOracleTts2NaturalModelDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m TtsOracleTts2NaturalModelDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m TtsOracleTts2NaturalModelDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeTtsOracleTts2NaturalModelDetails TtsOracleTts2NaturalModelDetails
	s := struct {
		DiscriminatorParam string `json:"modelName"`
		MarshalTypeTtsOracleTts2NaturalModelDetails
	}{
		"TTS_2_NATURAL",
		(MarshalTypeTtsOracleTts2NaturalModelDetails)(m),
	}

	return json.Marshal(&s)
}
