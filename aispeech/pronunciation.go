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

// Pronunciation Pronunciation Object Reference
type Pronunciation struct {

	// Written phonetic representation of entity value
	SoundsLike *string `mandatory:"false" json:"soundsLike"`

	Audio LocationDetails `mandatory:"false" json:"audio"`
}

func (m Pronunciation) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m Pronunciation) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *Pronunciation) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		SoundsLike *string         `json:"soundsLike"`
		Audio      locationdetails `json:"audio"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.SoundsLike = model.SoundsLike

	nn, e = model.Audio.UnmarshalPolymorphicJSON(model.Audio.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.Audio = nn.(LocationDetails)
	} else {
		m.Audio = nil
	}

	return
}
