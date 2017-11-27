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

// PatchSummary. A Patch for a DB System or DB Home.
// To use any of the API operations, you must be authorized in an IAM policy. If you're not authorized,
// talk to an administrator. If you're an administrator who needs to write policies to give users access,
// see [Getting Started with Policies]({{DOC_SERVER_URL}}/Content/Identity/Concepts/policygetstarted.htm).
type PatchSummary struct {

	// The text describing this patch package.
	Description *string `mandatory:"true" json:"description,omitempty"`

	// The OCID of the patch.
	ID *string `mandatory:"true" json:"id,omitempty"`

	// The date and time that the patch was released.
	TimeReleased *common.SDKTime `mandatory:"true" json:"timeReleased,omitempty"`

	// The version of this patch package.
	Version *string `mandatory:"true" json:"version,omitempty"`

	// Actions that can possibly be performed using this patch.
	AvailableActions []PatchSummaryAvailableActionsEnum `mandatory:"false" json:"availableActions,omitempty"`

	// Action that is currently being performed or was completed last.
	LastAction PatchSummaryLastActionEnum `mandatory:"false" json:"lastAction,omitempty"`

	// A descriptive text associated with the lifecycleState.
	// Typically can contain additional displayable text.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails,omitempty"`

	// The current state of the patch as a result of lastAction.
	LifecycleState PatchSummaryLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`
}

func (model PatchSummary) String() string {
	return common.PointerString(model)
}

type PatchSummaryAvailableActionsEnum string

const (
	PATCH_SUMMARY_AVAILABLE_ACTIONS_APPLY    PatchSummaryAvailableActionsEnum = "APPLY"
	PATCH_SUMMARY_AVAILABLE_ACTIONS_PRECHECK PatchSummaryAvailableActionsEnum = "PRECHECK"
	PATCH_SUMMARY_AVAILABLE_ACTIONS_UNKNOWN  PatchSummaryAvailableActionsEnum = "UNKNOWN"
)

var mapping_patchsummary_availableActions = map[string]PatchSummaryAvailableActionsEnum{
	"APPLY":    PATCH_SUMMARY_AVAILABLE_ACTIONS_APPLY,
	"PRECHECK": PATCH_SUMMARY_AVAILABLE_ACTIONS_PRECHECK,
	"UNKNOWN":  PATCH_SUMMARY_AVAILABLE_ACTIONS_UNKNOWN,
}

func GetPatchSummaryAvailableActionsEnumValues() []PatchSummaryAvailableActionsEnum {
	values := make([]PatchSummaryAvailableActionsEnum, 0)
	for _, v := range mapping_patchsummary_availableActions {
		if v != PATCH_SUMMARY_AVAILABLE_ACTIONS_UNKNOWN {
			values = append(values, v)
		}
	}
	return values
}

type PatchSummaryLastActionEnum string

const (
	PATCH_SUMMARY_LAST_ACTION_APPLY    PatchSummaryLastActionEnum = "APPLY"
	PATCH_SUMMARY_LAST_ACTION_PRECHECK PatchSummaryLastActionEnum = "PRECHECK"
	PATCH_SUMMARY_LAST_ACTION_UNKNOWN  PatchSummaryLastActionEnum = "UNKNOWN"
)

var mapping_patchsummary_lastAction = map[string]PatchSummaryLastActionEnum{
	"APPLY":    PATCH_SUMMARY_LAST_ACTION_APPLY,
	"PRECHECK": PATCH_SUMMARY_LAST_ACTION_PRECHECK,
	"UNKNOWN":  PATCH_SUMMARY_LAST_ACTION_UNKNOWN,
}

func GetPatchSummaryLastActionEnumValues() []PatchSummaryLastActionEnum {
	values := make([]PatchSummaryLastActionEnum, 0)
	for _, v := range mapping_patchsummary_lastAction {
		if v != PATCH_SUMMARY_LAST_ACTION_UNKNOWN {
			values = append(values, v)
		}
	}
	return values
}

type PatchSummaryLifecycleStateEnum string

const (
	PATCH_SUMMARY_LIFECYCLE_STATE_AVAILABLE   PatchSummaryLifecycleStateEnum = "AVAILABLE"
	PATCH_SUMMARY_LIFECYCLE_STATE_SUCCESS     PatchSummaryLifecycleStateEnum = "SUCCESS"
	PATCH_SUMMARY_LIFECYCLE_STATE_IN_PROGRESS PatchSummaryLifecycleStateEnum = "IN_PROGRESS"
	PATCH_SUMMARY_LIFECYCLE_STATE_FAILED      PatchSummaryLifecycleStateEnum = "FAILED"
	PATCH_SUMMARY_LIFECYCLE_STATE_UNKNOWN     PatchSummaryLifecycleStateEnum = "UNKNOWN"
)

var mapping_patchsummary_lifecycleState = map[string]PatchSummaryLifecycleStateEnum{
	"AVAILABLE":   PATCH_SUMMARY_LIFECYCLE_STATE_AVAILABLE,
	"SUCCESS":     PATCH_SUMMARY_LIFECYCLE_STATE_SUCCESS,
	"IN_PROGRESS": PATCH_SUMMARY_LIFECYCLE_STATE_IN_PROGRESS,
	"FAILED":      PATCH_SUMMARY_LIFECYCLE_STATE_FAILED,
	"UNKNOWN":     PATCH_SUMMARY_LIFECYCLE_STATE_UNKNOWN,
}

func GetPatchSummaryLifecycleStateEnumValues() []PatchSummaryLifecycleStateEnum {
	values := make([]PatchSummaryLifecycleStateEnum, 0)
	for _, v := range mapping_patchsummary_lifecycleState {
		if v != PATCH_SUMMARY_LIFECYCLE_STATE_UNKNOWN {
			values = append(values, v)
		}
	}
	return values
}
