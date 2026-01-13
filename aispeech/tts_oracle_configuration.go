// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
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

// TtsOracleConfiguration Use this configuration for selecting a model from Oracle model family.
type TtsOracleConfiguration struct {
	ModelDetails TtsOracleModelDetails `mandatory:"false" json:"modelDetails"`

	SpeechSettings *TtsOracleSpeechSettings `mandatory:"false" json:"speechSettings"`
}

func (m TtsOracleConfiguration) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m TtsOracleConfiguration) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m TtsOracleConfiguration) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeTtsOracleConfiguration TtsOracleConfiguration
	s := struct {
		DiscriminatorParam string `json:"modelFamily"`
		MarshalTypeTtsOracleConfiguration
	}{
		"ORACLE",
		(MarshalTypeTtsOracleConfiguration)(m),
	}

	return json.Marshal(&s)
}

// UnmarshalJSON unmarshals from json
func (m *TtsOracleConfiguration) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		ModelDetails   ttsoraclemodeldetails    `json:"modelDetails"`
		SpeechSettings *TtsOracleSpeechSettings `json:"speechSettings"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	nn, e = model.ModelDetails.UnmarshalPolymorphicJSON(model.ModelDetails.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.ModelDetails = nn.(TtsOracleModelDetails)
	} else {
		m.ModelDetails = nil
	}

	m.SpeechSettings = model.SpeechSettings

	return
}
