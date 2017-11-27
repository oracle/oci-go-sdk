// Copyright (c) 2016, 2017, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Load Balancing Service API
//
// API for the Load Balancing Service
//

package loadbalancer

import (
	"github.com/oci-go-sdk/common"
)

// WorkRequestError. An object returned in the event of a work request error.
type WorkRequestError struct {
	ErrorCode WorkRequestErrorErrorCodeEnum `mandatory:"true" json:"errorCode,omitempty"`

	// A human-readable error string.
	Message *string `mandatory:"true" json:"message,omitempty"`
}

func (model WorkRequestError) String() string {
	return common.PointerString(model)
}

type WorkRequestErrorErrorCodeEnum string

const (
	WORK_REQUEST_ERROR_ERROR_CODE_BAD_INPUT      WorkRequestErrorErrorCodeEnum = "BAD_INPUT"
	WORK_REQUEST_ERROR_ERROR_CODE_INTERNAL_ERROR WorkRequestErrorErrorCodeEnum = "INTERNAL_ERROR"
	WORK_REQUEST_ERROR_ERROR_CODE_UNKNOWN        WorkRequestErrorErrorCodeEnum = "UNKNOWN"
)

var mapping_workrequesterror_errorCode = map[string]WorkRequestErrorErrorCodeEnum{
	"BAD_INPUT":      WORK_REQUEST_ERROR_ERROR_CODE_BAD_INPUT,
	"INTERNAL_ERROR": WORK_REQUEST_ERROR_ERROR_CODE_INTERNAL_ERROR,
	"UNKNOWN":        WORK_REQUEST_ERROR_ERROR_CODE_UNKNOWN,
}

func GetWorkRequestErrorErrorCodeEnumValues() []WorkRequestErrorErrorCodeEnum {
	values := make([]WorkRequestErrorErrorCodeEnum, 0)
	for _, v := range mapping_workrequesterror_errorCode {
		if v != WORK_REQUEST_ERROR_ERROR_CODE_UNKNOWN {
			values = append(values, v)
		}
	}
	return values
}
