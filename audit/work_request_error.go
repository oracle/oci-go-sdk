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

// WorkRequestError The error entity.
type WorkRequestError struct {

	// A machine-usable code for the error that occured.
	Code *string `mandatory:"true" json:"code,omitempty"`

	// A human-readable error string.
	Message *string `mandatory:"true" json:"message,omitempty"`

	// Date and time the error happened, in the format defined by RFC3339.
	// Example: `2016-08-25T21:10:29.600Z`
	Timestamp *common.SDKTime `mandatory:"false" json:"timestamp,omitempty"`
}

func (model WorkRequestError) String() string {
	return common.PointerString(model)
}
