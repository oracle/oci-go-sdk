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

// TtsOracleTts1StandardModelDetails Use this schema for specifying properties of TTS_1_STANDARD model from Oracle model family.
type TtsOracleTts1StandardModelDetails struct {

	// Speaker in whose voice the user wants the output speech to be in.
	// The possible values for `voiceId` can be obtained by calling ListVoices api.
	VoiceId *string `mandatory:"false" json:"voiceId"`
}

func (m TtsOracleTts1StandardModelDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m TtsOracleTts1StandardModelDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m TtsOracleTts1StandardModelDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeTtsOracleTts1StandardModelDetails TtsOracleTts1StandardModelDetails
	s := struct {
		DiscriminatorParam string `json:"modelName"`
		MarshalTypeTtsOracleTts1StandardModelDetails
	}{
		"TTS_1_STANDARD",
		(MarshalTypeTtsOracleTts1StandardModelDetails)(m),
	}

	return json.Marshal(&s)
}
