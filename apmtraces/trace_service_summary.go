// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Apm Traces API
//
// API for APM Trace service. Use this API to query the Traces and associated Spans.
//

package apmtraces

import (
	"github.com/oracle/oci-go-sdk/v50/common"
)

// TraceServiceSummary A summary of the spans in a trace by service.
type TraceServiceSummary struct {

	// Name associated with the service.
	SpanServiceName *string `mandatory:"true" json:"spanServiceName"`

	// Number of spans for serviceName in this trace.
	TotalSpans *int64 `mandatory:"true" json:"totalSpans"`

	// Number of spans with errorsfor serviceName in this trace.
	ErrorSpans *int64 `mandatory:"true" json:"errorSpans"`
}

func (m TraceServiceSummary) String() string {
	return common.PointerString(m)
}
