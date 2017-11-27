// Copyright (c) 2016, 2017, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Database Service API
//
// The API for the Database Service.
//

package database

import (
	"github.com/oci-go-sdk/common"
)

// PatchDetails. The details about what actions to perform and using what patch to the specified target.
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

type PatchDetailsActionEnum string

const (
	PATCH_DETAILS_ACTION_APPLY    PatchDetailsActionEnum = "APPLY"
	PATCH_DETAILS_ACTION_PRECHECK PatchDetailsActionEnum = "PRECHECK"
	PATCH_DETAILS_ACTION_UNKNOWN  PatchDetailsActionEnum = "UNKNOWN"
)

var mapping_patchdetails_action = map[string]PatchDetailsActionEnum{
	"APPLY":    PATCH_DETAILS_ACTION_APPLY,
	"PRECHECK": PATCH_DETAILS_ACTION_PRECHECK,
	"UNKNOWN":  PATCH_DETAILS_ACTION_UNKNOWN,
}

func GetPatchDetailsActionEnumValues() []PatchDetailsActionEnum {
	values := make([]PatchDetailsActionEnum, 0)
	for _, v := range mapping_patchdetails_action {
		if v != PATCH_DETAILS_ACTION_UNKNOWN {
			values = append(values, v)
		}
	}
	return values
}
