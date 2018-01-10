// Copyright (c) 2016, 2017, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Database Service API
//
// The API for the Database Service.
//

package database

import (
	"github.com/oracle/oci-go-sdk/common"
)

// PatchDetails The details about what actions to perform and using what patch to the specified target.
// This is part of an update request that is applied to a version field on the target such
// as DB System, database home, etc.
type PatchDetails struct {

	// The action to perform on the patch.
	Action PatchDetailsActionEnum `mandatory:"false" json:"action,omitempty"`

	// The OCID of the patch.
	PatchID *string `mandatory:"false" json:"patchId,omitempty"`
}

func (model PatchDetails) String() string {
	return common.PointerString(model)
}

// PatchDetailsActionEnum Enum with underlying type: string
type PatchDetailsActionEnum string

// Set of constants representing the allowable values for PatchDetailsAction
const (
	PatchDetailsActionApply    PatchDetailsActionEnum = "APPLY"
	PatchDetailsActionPrecheck PatchDetailsActionEnum = "PRECHECK"
	PatchDetailsActionUnknown  PatchDetailsActionEnum = "UNKNOWN"
)

var mappingPatchDetailsAction = map[string]PatchDetailsActionEnum{
	"APPLY":    PatchDetailsActionApply,
	"PRECHECK": PatchDetailsActionPrecheck,
	"UNKNOWN":  PatchDetailsActionUnknown,
}

// GetPatchDetailsActionEnumValues Enumerates the set of values for PatchDetailsAction
func GetPatchDetailsActionEnumValues() []PatchDetailsActionEnum {
	values := make([]PatchDetailsActionEnum, 0)
	for _, v := range mappingPatchDetailsAction {
		if v != PatchDetailsActionUnknown {
			values = append(values, v)
		}
	}
	return values
}
