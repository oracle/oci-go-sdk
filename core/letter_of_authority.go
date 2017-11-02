// Copyright (c) 2016, 2017, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Core Services API
//
// APIs for Networking Service, Compute Service, and Block Volume Service.
//

package core

import (
	"bitbucket.aka.lgl.grungy.us/golang-sdk2/common"
)

// LetterOfAuthority. The Letter of Authority for the cross-connect. You must submit this letter when
// requesting cabling for the cross-connect at the FastConnect location.
type LetterOfAuthority struct {

	// The type of cross-connect fiber, termination, and optical specification.
	CircuitType LetterOfAuthorityCircuitTypeEnum `mandatory:"false" json:"circuitType,omitempty"`

	// The OCID of the cross-connect.
	CrossConnectID *string `mandatory:"false" json:"crossConnectId,omitempty"`

	// The address of the FastConnect location.
	FacilityLocation *string `mandatory:"false" json:"facilityLocation,omitempty"`

	// The meet-me room port for this cross-connect.
	PortName *string `mandatory:"false" json:"portName,omitempty"`

	// The date and time when the Letter of Authority expires, in the format defined by RFC3339.
	TimeExpires *common.SDKTime `mandatory:"false" json:"timeExpires,omitempty"`

	// The date and time the Letter of Authority was created, in the format defined by RFC3339.
	// Example: `2016-08-25T21:10:29.600Z`
	TimeIssued *common.SDKTime `mandatory:"false" json:"timeIssued,omitempty"`
}

func (model LetterOfAuthority) String() string {
	return common.PointerString(model)
}

type LetterOfAuthorityCircuitTypeEnum string
type LetterOfAuthorityCircuitType struct{}

const (
	LETTER_OF_AUTHORITY_CIRCUIT_TYPE_LC      LetterOfAuthorityCircuitTypeEnum = "Single_mode_LC"
	LETTER_OF_AUTHORITY_CIRCUIT_TYPE_SC      LetterOfAuthorityCircuitTypeEnum = "Single_mode_SC"
	LETTER_OF_AUTHORITY_CIRCUIT_TYPE_UNKNOWN LetterOfAuthorityCircuitTypeEnum = "UNKNOWN"
)

var mapping_letterofauthority_circuitType = map[string]LetterOfAuthorityCircuitTypeEnum{
	"Single_mode_LC": LETTER_OF_AUTHORITY_CIRCUIT_TYPE_LC,
	"Single_mode_SC": LETTER_OF_AUTHORITY_CIRCUIT_TYPE_SC,
	"UNKNOWN":        LETTER_OF_AUTHORITY_CIRCUIT_TYPE_UNKNOWN,
}

func (receiver LetterOfAuthorityCircuitType) Values() []LetterOfAuthorityCircuitTypeEnum {
	values := make([]LetterOfAuthorityCircuitTypeEnum, 0)
	for _, v := range mapping_letterofauthority_circuitType {
		if v != LETTER_OF_AUTHORITY_CIRCUIT_TYPE_UNKNOWN {
			values = append(values, v)
		}
	}
	return values
}

func (receiver LetterOfAuthorityCircuitType) IsValid(toBeChecked string) bool {
	for _, v := range receiver.Values() {
		if LetterOfAuthorityCircuitTypeEnum(toBeChecked) == v {
			return true
		}
	}
	return false
}

func (receiver LetterOfAuthorityCircuitType) From(toBeConverted string) LetterOfAuthorityCircuitTypeEnum {
	if val, ok := mapping_letterofauthority_circuitType[toBeConverted]; ok {
		return val
	}
	return LETTER_OF_AUTHORITY_CIRCUIT_TYPE_UNKNOWN
}
