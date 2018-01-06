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

// WorkRequestLogEntry. The log entity.
type WorkRequestLogEntry struct {

	// A human-readable error string.
	Message *string `mandatory:"true" json:"message,omitempty"`

	// Date and time the log was written, in the format defined by RFC3339.
	// Example: `2016-08-25T21:10:29.600Z`
	Timestamp *common.SDKTime `mandatory:"false" json:"timestamp,omitempty"`
}

func (model WorkRequestLogEntry) String() string {
	return common.PointerString(model)
}
