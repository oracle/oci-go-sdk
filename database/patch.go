// Copyright (c) 2016, 2017, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Database Service API
//
// The API for the Database Service.
//

package database

import (
	"bitbucket.aka.lgl.grungy.us/golang-sdk2/common"
)

// Patch. A Patch for a DB System or DB Home.
// To use any of the API operations, you must be authorized in an IAM policy. If you're not authorized,
// talk to an administrator. If you're an administrator who needs to write policies to give users access,
// see [Getting Started with Policies]({{DOC_SERVER_URL}}/Content/Identity/Concepts/policygetstarted.htm).
type Patch struct {

	// The text describing this patch package.
	Description *string `mandatory:"true" json:"description,omitempty"`

	// The OCID of the patch.
	ID *string `mandatory:"true" json:"id,omitempty"`

	// The date and time that the patch was released.
	TimeReleased *common.SDKTime `mandatory:"true" json:"timeReleased,omitempty"`

	// The version of this patch package.
	Version *string `mandatory:"true" json:"version,omitempty"`

	// Actions that can possibly be performed using this patch.
	AvailableActions []PatchAvailableActionsEnum `mandatory:"false" json:"availableActions,omitempty"`

	// Action that is currently being performed or was completed last.
	LastAction PatchLastActionEnum `mandatory:"false" json:"lastAction,omitempty"`

	// A descriptive text associated with the lifecycleState.
	// Typically can contain additional displayable text.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails,omitempty"`

	// The current state of the patch as a result of lastAction.
	LifecycleState PatchLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`
}

func (model Patch) String() string {
	return common.PointerString(model)
}

type PatchAvailableActionsEnum string

const (
	PATCH_AVAILABLE_ACTIONS_APPLY    PatchAvailableActionsEnum = "APPLY"
	PATCH_AVAILABLE_ACTIONS_PRECHECK PatchAvailableActionsEnum = "PRECHECK"
	PATCH_AVAILABLE_ACTIONS_UNKNOWN  PatchAvailableActionsEnum = "UNKNOWN"
)

var mapping_patch_availableActions = map[string]PatchAvailableActionsEnum{
	"APPLY":    PATCH_AVAILABLE_ACTIONS_APPLY,
	"PRECHECK": PATCH_AVAILABLE_ACTIONS_PRECHECK,
	"UNKNOWN":  PATCH_AVAILABLE_ACTIONS_UNKNOWN,
}

func GetPatchAvailableActionsEnumValues() []PatchAvailableActionsEnum {
	values := make([]PatchAvailableActionsEnum, 0)
	for _, v := range mapping_patch_availableActions {
		if v != PATCH_AVAILABLE_ACTIONS_UNKNOWN {
			values = append(values, v)
		}
	}
	return values
}

type PatchLastActionEnum string

const (
	PATCH_LAST_ACTION_APPLY    PatchLastActionEnum = "APPLY"
	PATCH_LAST_ACTION_PRECHECK PatchLastActionEnum = "PRECHECK"
	PATCH_LAST_ACTION_UNKNOWN  PatchLastActionEnum = "UNKNOWN"
)

var mapping_patch_lastAction = map[string]PatchLastActionEnum{
	"APPLY":    PATCH_LAST_ACTION_APPLY,
	"PRECHECK": PATCH_LAST_ACTION_PRECHECK,
	"UNKNOWN":  PATCH_LAST_ACTION_UNKNOWN,
}

func GetPatchLastActionEnumValues() []PatchLastActionEnum {
	values := make([]PatchLastActionEnum, 0)
	for _, v := range mapping_patch_lastAction {
		if v != PATCH_LAST_ACTION_UNKNOWN {
			values = append(values, v)
		}
	}
	return values
}

type PatchLifecycleStateEnum string

const (
	PATCH_LIFECYCLE_STATE_AVAILABLE   PatchLifecycleStateEnum = "AVAILABLE"
	PATCH_LIFECYCLE_STATE_SUCCESS     PatchLifecycleStateEnum = "SUCCESS"
	PATCH_LIFECYCLE_STATE_IN_PROGRESS PatchLifecycleStateEnum = "IN_PROGRESS"
	PATCH_LIFECYCLE_STATE_FAILED      PatchLifecycleStateEnum = "FAILED"
	PATCH_LIFECYCLE_STATE_UNKNOWN     PatchLifecycleStateEnum = "UNKNOWN"
)

var mapping_patch_lifecycleState = map[string]PatchLifecycleStateEnum{
	"AVAILABLE":   PATCH_LIFECYCLE_STATE_AVAILABLE,
	"SUCCESS":     PATCH_LIFECYCLE_STATE_SUCCESS,
	"IN_PROGRESS": PATCH_LIFECYCLE_STATE_IN_PROGRESS,
	"FAILED":      PATCH_LIFECYCLE_STATE_FAILED,
	"UNKNOWN":     PATCH_LIFECYCLE_STATE_UNKNOWN,
}

func GetPatchLifecycleStateEnumValues() []PatchLifecycleStateEnum {
	values := make([]PatchLifecycleStateEnum, 0)
	for _, v := range mapping_patch_lifecycleState {
		if v != PATCH_LIFECYCLE_STATE_UNKNOWN {
			values = append(values, v)
		}
	}
	return values
}
