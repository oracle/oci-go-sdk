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

// PatchHistoryEntry. The record of a patch action on a specified target.
type PatchHistoryEntry struct {

	// The OCID of the patch history entry.
	ID *string `mandatory:"true" json:"id,omitempty"`

	// The current state of the action.
	LifecycleState PatchHistoryEntryLifecycleStateEnum `mandatory:"true" json:"lifecycleState,omitempty"`

	// The OCID of the patch.
	PatchID *string `mandatory:"true" json:"patchId,omitempty"`

	// The date and time when the patch action started.
	TimeStarted *common.SDKTime `mandatory:"true" json:"timeStarted,omitempty"`

	// The action being performed or was completed.
	Action PatchHistoryEntryActionEnum `mandatory:"false" json:"action,omitempty"`

	// A descriptive text associated with the lifecycleState.
	// Typically contains additional displayable text.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails,omitempty"`

	// The date and time when the patch action completed.
	TimeEnded *common.SDKTime `mandatory:"false" json:"timeEnded,omitempty"`
}

func (model PatchHistoryEntry) String() string {
	return common.PointerString(model)
}

type PatchHistoryEntryActionEnum string

const (
	PATCH_HISTORY_ENTRY_ACTION_APPLY    PatchHistoryEntryActionEnum = "APPLY"
	PATCH_HISTORY_ENTRY_ACTION_PRECHECK PatchHistoryEntryActionEnum = "PRECHECK"
	PATCH_HISTORY_ENTRY_ACTION_UNKNOWN  PatchHistoryEntryActionEnum = "UNKNOWN"
)

var mapping_patchhistoryentry_action = map[string]PatchHistoryEntryActionEnum{
	"APPLY":    PATCH_HISTORY_ENTRY_ACTION_APPLY,
	"PRECHECK": PATCH_HISTORY_ENTRY_ACTION_PRECHECK,
	"UNKNOWN":  PATCH_HISTORY_ENTRY_ACTION_UNKNOWN,
}

func GetPatchHistoryEntryActionEnumValues() []PatchHistoryEntryActionEnum {
	values := make([]PatchHistoryEntryActionEnum, 0)
	for _, v := range mapping_patchhistoryentry_action {
		if v != PATCH_HISTORY_ENTRY_ACTION_UNKNOWN {
			values = append(values, v)
		}
	}
	return values
}

type PatchHistoryEntryLifecycleStateEnum string

const (
	PATCH_HISTORY_ENTRY_LIFECYCLE_STATE_IN_PROGRESS PatchHistoryEntryLifecycleStateEnum = "IN_PROGRESS"
	PATCH_HISTORY_ENTRY_LIFECYCLE_STATE_SUCCEEDED   PatchHistoryEntryLifecycleStateEnum = "SUCCEEDED"
	PATCH_HISTORY_ENTRY_LIFECYCLE_STATE_FAILED      PatchHistoryEntryLifecycleStateEnum = "FAILED"
	PATCH_HISTORY_ENTRY_LIFECYCLE_STATE_UNKNOWN     PatchHistoryEntryLifecycleStateEnum = "UNKNOWN"
)

var mapping_patchhistoryentry_lifecycleState = map[string]PatchHistoryEntryLifecycleStateEnum{
	"IN_PROGRESS": PATCH_HISTORY_ENTRY_LIFECYCLE_STATE_IN_PROGRESS,
	"SUCCEEDED":   PATCH_HISTORY_ENTRY_LIFECYCLE_STATE_SUCCEEDED,
	"FAILED":      PATCH_HISTORY_ENTRY_LIFECYCLE_STATE_FAILED,
	"UNKNOWN":     PATCH_HISTORY_ENTRY_LIFECYCLE_STATE_UNKNOWN,
}

func GetPatchHistoryEntryLifecycleStateEnumValues() []PatchHistoryEntryLifecycleStateEnum {
	values := make([]PatchHistoryEntryLifecycleStateEnum, 0)
	for _, v := range mapping_patchhistoryentry_lifecycleState {
		if v != PATCH_HISTORY_ENTRY_LIFECYCLE_STATE_UNKNOWN {
			values = append(values, v)
		}
	}
	return values
}
