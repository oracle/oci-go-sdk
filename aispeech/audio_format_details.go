// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Speech API
//
// The OCI Speech Service harnesses the power of spoken language by allowing developers to easily convert file-based data containing human speech into highly accurate text transcriptions.
//

package aispeech

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v59/common"
	"strings"
)

// AudioFormatDetails Audio format details.
type AudioFormatDetails struct {

	// Input file format. Example - WAV.
	Format *string `mandatory:"false" json:"format"`

	// Input file number of channels.
	NumberOfChannels *int `mandatory:"false" json:"numberOfChannels"`

	// Input file encoding. Example - PCM.
	Encoding *string `mandatory:"false" json:"encoding"`

	// Input file sampleRate. Example - 16000
	SampleRateInHz *int `mandatory:"false" json:"sampleRateInHz"`
}

func (m AudioFormatDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m AudioFormatDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
