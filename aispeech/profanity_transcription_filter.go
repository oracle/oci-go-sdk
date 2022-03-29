// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
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
	"github.com/oracle/oci-go-sdk/v64/common"
	"strings"
)

// ProfanityTranscriptionFilter Profanity transcription filter.
type ProfanityTranscriptionFilter struct {

	// The mode of filters.
	// Allowed values are:
	// - `MASK`: Will mask detected profanity in transcription.
	// - `REMOVE`: Will replace profane word with * in transcription.
	// - `TAG`: Will tag profane word as profanity but will show actual word.
	Mode ProfanityTranscriptionFilterModeEnum `mandatory:"true" json:"mode"`
}

func (m ProfanityTranscriptionFilter) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ProfanityTranscriptionFilter) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingProfanityTranscriptionFilterModeEnum(string(m.Mode)); !ok && m.Mode != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Mode: %s. Supported values are: %s.", m.Mode, strings.Join(GetProfanityTranscriptionFilterModeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m ProfanityTranscriptionFilter) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeProfanityTranscriptionFilter ProfanityTranscriptionFilter
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeProfanityTranscriptionFilter
	}{
		"PROFANITY",
		(MarshalTypeProfanityTranscriptionFilter)(m),
	}

	return json.Marshal(&s)
}

// ProfanityTranscriptionFilterModeEnum Enum with underlying type: string
type ProfanityTranscriptionFilterModeEnum string

// Set of constants representing the allowable values for ProfanityTranscriptionFilterModeEnum
const (
	ProfanityTranscriptionFilterModeMask   ProfanityTranscriptionFilterModeEnum = "MASK"
	ProfanityTranscriptionFilterModeRemove ProfanityTranscriptionFilterModeEnum = "REMOVE"
	ProfanityTranscriptionFilterModeTag    ProfanityTranscriptionFilterModeEnum = "TAG"
)

var mappingProfanityTranscriptionFilterModeEnum = map[string]ProfanityTranscriptionFilterModeEnum{
	"MASK":   ProfanityTranscriptionFilterModeMask,
	"REMOVE": ProfanityTranscriptionFilterModeRemove,
	"TAG":    ProfanityTranscriptionFilterModeTag,
}

var mappingProfanityTranscriptionFilterModeEnumLowerCase = map[string]ProfanityTranscriptionFilterModeEnum{
	"mask":   ProfanityTranscriptionFilterModeMask,
	"remove": ProfanityTranscriptionFilterModeRemove,
	"tag":    ProfanityTranscriptionFilterModeTag,
}

// GetProfanityTranscriptionFilterModeEnumValues Enumerates the set of values for ProfanityTranscriptionFilterModeEnum
func GetProfanityTranscriptionFilterModeEnumValues() []ProfanityTranscriptionFilterModeEnum {
	values := make([]ProfanityTranscriptionFilterModeEnum, 0)
	for _, v := range mappingProfanityTranscriptionFilterModeEnum {
		values = append(values, v)
	}
	return values
}

// GetProfanityTranscriptionFilterModeEnumStringValues Enumerates the set of values in String for ProfanityTranscriptionFilterModeEnum
func GetProfanityTranscriptionFilterModeEnumStringValues() []string {
	return []string{
		"MASK",
		"REMOVE",
		"TAG",
	}
}

// GetMappingProfanityTranscriptionFilterModeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingProfanityTranscriptionFilterModeEnum(val string) (ProfanityTranscriptionFilterModeEnum, bool) {
	enum, ok := mappingProfanityTranscriptionFilterModeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
