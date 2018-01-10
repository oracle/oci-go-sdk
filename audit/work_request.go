// Copyright (c) 2016, 2017, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Audit API
//
// API for the Audit Service. You can use this API for queries, but not bulk-export operations.
//

package audit

import (
	"github.com/oracle/oci-go-sdk/common"
)

// WorkRequest The response for work request.
type WorkRequest struct {

	// The id of the work request.
	ID *string `mandatory:"true" json:"id,omitempty"`

	// An enum-like description of the type of work the work request is doing.
	OperationType WorkRequestOperationTypeEnum `mandatory:"true" json:"operationType,omitempty"`

	// The current status of the work request.
	Status WorkRequestStatusEnum `mandatory:"true" json:"status,omitempty"`

	// The OCID of the compartment that contains the work request.
	CompartmentID *string `mandatory:"false" json:"compartmentId,omitempty"`

	// The resources this work request affects.
	Resources []WorkRequestResource `mandatory:"false" json:"resources,omitempty"`

	// The errors for work request.
	Errors []WorkRequestError `mandatory:"false" json:"errors,omitempty"`

	// The logs for work request.
	Logs []WorkRequestLogEntry `mandatory:"false" json:"logs,omitempty"`

	// Date and time the work was accepted, in the format defined by RFC3339.
	// Example: `2016-08-25T21:10:29.600Z`
	TimeAccepted *common.SDKTime `mandatory:"false" json:"timeAccepted,omitempty"`

	// Date and time the work started, in the format defined by RFC3339.
	// Example: `2016-08-25T21:10:29.600Z`
	TimeStarted *common.SDKTime `mandatory:"false" json:"timeStarted,omitempty"`

	// Date and time the work completed, in the format defined by RFC3339.
	// Example: `2016-08-25T21:10:29.600Z`
	TimeFinished *common.SDKTime `mandatory:"false" json:"timeFinished,omitempty"`

	// How much progress the operation has made.
	PercentComplete *float32 `mandatory:"false" json:"percentComplete,omitempty"`
}

func (model WorkRequest) String() string {
	return common.PointerString(model)
}

// WorkRequestOperationTypeEnum Enum with underlying type: string
type WorkRequestOperationTypeEnum string

// Set of constants representing the allowable values for WorkRequestOperationType
const (
	WorkRequestOperationTypeConfiguration WorkRequestOperationTypeEnum = "UPDATE_CONFIGURATION"
	WorkRequestOperationTypeUnknown       WorkRequestOperationTypeEnum = "UNKNOWN"
)

var mappingWorkRequestOperationType = map[string]WorkRequestOperationTypeEnum{
	"UPDATE_CONFIGURATION": WorkRequestOperationTypeConfiguration,
	"UNKNOWN":              WorkRequestOperationTypeUnknown,
}

// GetWorkRequestOperationTypeEnumValues Enumerates the set of values for WorkRequestOperationType
func GetWorkRequestOperationTypeEnumValues() []WorkRequestOperationTypeEnum {
	values := make([]WorkRequestOperationTypeEnum, 0)
	for _, v := range mappingWorkRequestOperationType {
		if v != WorkRequestOperationTypeUnknown {
			values = append(values, v)
		}
	}
	return values
}

// WorkRequestStatusEnum Enum with underlying type: string
type WorkRequestStatusEnum string

// Set of constants representing the allowable values for WorkRequestStatus
const (
	WorkRequestStatusAccepted   WorkRequestStatusEnum = "ACCEPTED"
	WorkRequestStatusInProgress WorkRequestStatusEnum = "IN_PROGRESS"
	WorkRequestStatusFailed     WorkRequestStatusEnum = "FAILED"
	WorkRequestStatusSucceeded  WorkRequestStatusEnum = "SUCCEEDED"
	WorkRequestStatusCanceling  WorkRequestStatusEnum = "CANCELING"
	WorkRequestStatusCanceled   WorkRequestStatusEnum = "CANCELED"
	WorkRequestStatusUnknown    WorkRequestStatusEnum = "UNKNOWN"
)

var mappingWorkRequestStatus = map[string]WorkRequestStatusEnum{
	"ACCEPTED":    WorkRequestStatusAccepted,
	"IN_PROGRESS": WorkRequestStatusInProgress,
	"FAILED":      WorkRequestStatusFailed,
	"SUCCEEDED":   WorkRequestStatusSucceeded,
	"CANCELING":   WorkRequestStatusCanceling,
	"CANCELED":    WorkRequestStatusCanceled,
	"UNKNOWN":     WorkRequestStatusUnknown,
}

// GetWorkRequestStatusEnumValues Enumerates the set of values for WorkRequestStatus
func GetWorkRequestStatusEnumValues() []WorkRequestStatusEnum {
	values := make([]WorkRequestStatusEnum, 0)
	for _, v := range mappingWorkRequestStatus {
		if v != WorkRequestStatusUnknown {
			values = append(values, v)
		}
	}
	return values
}
