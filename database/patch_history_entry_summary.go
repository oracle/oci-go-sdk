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

// PatchHistoryEntrySummary. The record of a patch action on a specified target.
type PatchHistoryEntrySummary struct {

	// The OCID of the patch history entry.
	ID *string `mandatory:"true" json:"id,omitempty"`

	// The current state of the action.
	LifecycleState PatchHistoryEntrySummaryLifecycleStateEnum `mandatory:"true" json:"lifecycleState,omitempty"`

	// The OCID of the patch.
	PatchID *string `mandatory:"true" json:"patchId,omitempty"`

	// The date and time when the patch action started.
	TimeStarted *common.SDKTime `mandatory:"true" json:"timeStarted,omitempty"`

	// The action being performed or was completed.
	Action PatchHistoryEntrySummaryActionEnum `mandatory:"false" json:"action,omitempty"`

	// A descriptive text associated with the lifecycleState.
	// Typically contains additional displayable text.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails,omitempty"`

	// The date and time when the patch action completed.
	TimeEnded *common.SDKTime `mandatory:"false" json:"timeEnded,omitempty"`
}

func (model PatchHistoryEntrySummary) String() string {
	return common.PointerString(model)
}

type PatchHistoryEntrySummaryActionEnum string

const (
	PATCH_HISTORY_ENTRY_SUMMARY_ACTION_APPLY    PatchHistoryEntrySummaryActionEnum = "APPLY"
	PATCH_HISTORY_ENTRY_SUMMARY_ACTION_PRECHECK PatchHistoryEntrySummaryActionEnum = "PRECHECK"
	PATCH_HISTORY_ENTRY_SUMMARY_ACTION_UNKNOWN  PatchHistoryEntrySummaryActionEnum = "UNKNOWN"
)

var mapping_patchhistoryentrysummary_action = map[string]PatchHistoryEntrySummaryActionEnum{
	"APPLY":    PATCH_HISTORY_ENTRY_SUMMARY_ACTION_APPLY,
	"PRECHECK": PATCH_HISTORY_ENTRY_SUMMARY_ACTION_PRECHECK,
	"UNKNOWN":  PATCH_HISTORY_ENTRY_SUMMARY_ACTION_UNKNOWN,
}

func GetPatchHistoryEntrySummaryActionEnumValues() []PatchHistoryEntrySummaryActionEnum {
	values := make([]PatchHistoryEntrySummaryActionEnum, 0)
	for _, v := range mapping_patchhistoryentrysummary_action {
		if v != PATCH_HISTORY_ENTRY_SUMMARY_ACTION_UNKNOWN {
			values = append(values, v)
		}
	}
	return values
}

type PatchHistoryEntrySummaryLifecycleStateEnum string

const (
	PATCH_HISTORY_ENTRY_SUMMARY_LIFECYCLE_STATE_IN_PROGRESS PatchHistoryEntrySummaryLifecycleStateEnum = "IN_PROGRESS"
	PATCH_HISTORY_ENTRY_SUMMARY_LIFECYCLE_STATE_SUCCEEDED   PatchHistoryEntrySummaryLifecycleStateEnum = "SUCCEEDED"
	PATCH_HISTORY_ENTRY_SUMMARY_LIFECYCLE_STATE_FAILED      PatchHistoryEntrySummaryLifecycleStateEnum = "FAILED"
	PATCH_HISTORY_ENTRY_SUMMARY_LIFECYCLE_STATE_UNKNOWN     PatchHistoryEntrySummaryLifecycleStateEnum = "UNKNOWN"
)

var mapping_patchhistoryentrysummary_lifecycleState = map[string]PatchHistoryEntrySummaryLifecycleStateEnum{
	"IN_PROGRESS": PATCH_HISTORY_ENTRY_SUMMARY_LIFECYCLE_STATE_IN_PROGRESS,
	"SUCCEEDED":   PATCH_HISTORY_ENTRY_SUMMARY_LIFECYCLE_STATE_SUCCEEDED,
	"FAILED":      PATCH_HISTORY_ENTRY_SUMMARY_LIFECYCLE_STATE_FAILED,
	"UNKNOWN":     PATCH_HISTORY_ENTRY_SUMMARY_LIFECYCLE_STATE_UNKNOWN,
}

func GetPatchHistoryEntrySummaryLifecycleStateEnumValues() []PatchHistoryEntrySummaryLifecycleStateEnum {
	values := make([]PatchHistoryEntrySummaryLifecycleStateEnum, 0)
	for _, v := range mapping_patchhistoryentrysummary_lifecycleState {
		if v != PATCH_HISTORY_ENTRY_SUMMARY_LIFECYCLE_STATE_UNKNOWN {
			values = append(values, v)
		}
	}
	return values
}
