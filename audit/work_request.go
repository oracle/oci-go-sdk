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

// WorkRequest. The response for work request.
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

type WorkRequestOperationTypeEnum string

const (
	WORK_REQUEST_OPERATION_TYPE_CONFIGURATION WorkRequestOperationTypeEnum = "UPDATE_CONFIGURATION"
	WORK_REQUEST_OPERATION_TYPE_UNKNOWN       WorkRequestOperationTypeEnum = "UNKNOWN"
)

var mapping_workrequest_operationType = map[string]WorkRequestOperationTypeEnum{
	"UPDATE_CONFIGURATION": WORK_REQUEST_OPERATION_TYPE_CONFIGURATION,
	"UNKNOWN":              WORK_REQUEST_OPERATION_TYPE_UNKNOWN,
}

func GetWorkRequestOperationTypeEnumValues() []WorkRequestOperationTypeEnum {
	values := make([]WorkRequestOperationTypeEnum, 0)
	for _, v := range mapping_workrequest_operationType {
		if v != WORK_REQUEST_OPERATION_TYPE_UNKNOWN {
			values = append(values, v)
		}
	}
	return values
}

type WorkRequestStatusEnum string

const (
	WORK_REQUEST_STATUS_ACCEPTED    WorkRequestStatusEnum = "ACCEPTED"
	WORK_REQUEST_STATUS_IN_PROGRESS WorkRequestStatusEnum = "IN_PROGRESS"
	WORK_REQUEST_STATUS_FAILED      WorkRequestStatusEnum = "FAILED"
	WORK_REQUEST_STATUS_SUCCEEDED   WorkRequestStatusEnum = "SUCCEEDED"
	WORK_REQUEST_STATUS_CANCELING   WorkRequestStatusEnum = "CANCELING"
	WORK_REQUEST_STATUS_CANCELED    WorkRequestStatusEnum = "CANCELED"
	WORK_REQUEST_STATUS_UNKNOWN     WorkRequestStatusEnum = "UNKNOWN"
)

var mapping_workrequest_status = map[string]WorkRequestStatusEnum{
	"ACCEPTED":    WORK_REQUEST_STATUS_ACCEPTED,
	"IN_PROGRESS": WORK_REQUEST_STATUS_IN_PROGRESS,
	"FAILED":      WORK_REQUEST_STATUS_FAILED,
	"SUCCEEDED":   WORK_REQUEST_STATUS_SUCCEEDED,
	"CANCELING":   WORK_REQUEST_STATUS_CANCELING,
	"CANCELED":    WORK_REQUEST_STATUS_CANCELED,
	"UNKNOWN":     WORK_REQUEST_STATUS_UNKNOWN,
}

func GetWorkRequestStatusEnumValues() []WorkRequestStatusEnum {
	values := make([]WorkRequestStatusEnum, 0)
	for _, v := range mapping_workrequest_status {
		if v != WORK_REQUEST_STATUS_UNKNOWN {
			values = append(values, v)
		}
	}
	return values
}
