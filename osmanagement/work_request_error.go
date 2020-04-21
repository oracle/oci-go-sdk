// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// OSMS
//
// OS Management as a Service API definition
//

package osmanagement

import (
	"github.com/oracle/oci-go-sdk/common"
)

// WorkRequestError Human readable error message describing why the work request failed
type WorkRequestError struct {

	// A machine-usable code for the error that occured.
	Code *string `mandatory:"true" json:"code"`

	// A human readable description of the issue encountered.
	Message *string `mandatory:"true" json:"message"`

	// The date and time the error happened, as described in
	// RFC 3339 (https://tools.ietf.org/rfc/rfc3339), section 14.29.
	Timestamp *common.SDKTime `mandatory:"false" json:"timestamp"`
}

func (m WorkRequestError) String() string {
	return common.PointerString(m)
}
