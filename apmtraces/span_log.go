// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Apm Traces API
//
// API for APM Trace service. Use this API to query the Traces and associated Spans.
//

package apmtraces

import (
	"github.com/oracle/oci-go-sdk/v56/common"
)

// SpanLog Definition of a log which is a key value pair of log data.
type SpanLog struct {

	// Key that specifies the log name.
	LogKey *string `mandatory:"true" json:"logKey"`

	// Value associated with the log key.
	LogValue *string `mandatory:"true" json:"logValue"`
}

func (m SpanLog) String() string {
	return common.PointerString(m)
}
