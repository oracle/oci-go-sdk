// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Apm Traces API
//
// API for APM Trace service. Use this API to query the Traces and associated Spans.
//

package apmtraces

import (
	"github.com/oracle/oci-go-sdk/v53/common"
)

// TraceSpanSummary Trace Span summary object is the summary of the span information for the spans in the span collection
// in the trace window that is being queried.
type TraceSpanSummary struct {

	// Unique identifier (traceId) for the trace that represents the span set.  Note that this field is
	// defined as traceKey in the API to comply with OCI API fields naming conventions.  The traceKey maps to
	// the traceId in the APM repository.
	Key *string `mandatory:"true" json:"key"`

	// Start time of the earliest span in this span collection.
	TimeEarliestSpanStarted *common.SDKTime `mandatory:"true" json:"timeEarliestSpanStarted"`

	// End time of the span that most recently ended in this span collection.
	TimeLatestSpanEnded *common.SDKTime `mandatory:"true" json:"timeLatestSpanEnded"`

	// The number of spans that have been processed by the system for this trace.  Note that there
	// could be additional spans that have not been processed or reported yet if the trace is still
	// in progress.
	SpanCount *int `mandatory:"true" json:"spanCount"`

	// The number of spans with error that have been processed by the system for this trace.
	// Note that the number of spans with errors may be less than the total number of actual spans
	// in this trace.
	ErrorSpanCount *int `mandatory:"true" json:"errorSpanCount"`

	// Time between the start of the earliest span and the end of the most recent span in milliseconds.
	TraceDurationInMs *int `mandatory:"true" json:"traceDurationInMs"`

	// Boolean flag that indicates whether the trace errored out.
	IsFault *bool `mandatory:"true" json:"isFault"`

	// The status of the trace.
	// The trace statuses are defined as follows:
	// complete – a root span has been recorded, but there is no information on the errors.
	// success - a complete root span is recorded there is a successful error type and error code - HTTP 200.
	// incomplete - the root span has not yet been received.
	// error - the root span returned with an error. There may or may not be an associated error code or error type.
	TraceStatus *string `mandatory:"true" json:"traceStatus"`

	// Error type of the trace.
	TraceErrorType *string `mandatory:"true" json:"traceErrorType"`

	// Error code of the trace.
	TraceErrorCode *string `mandatory:"true" json:"traceErrorCode"`

	// Root span name associated with the trace. This is usually the flow start operation name.
	// Null if the root span is not yet completed.
	RootSpanOperationName *string `mandatory:"false" json:"rootSpanOperationName"`

	// Service associated with this trace.
	RootSpanServiceName *string `mandatory:"false" json:"rootSpanServiceName"`

	// Start time of the root span for this span collection.
	TimeRootSpanStarted *common.SDKTime `mandatory:"false" json:"timeRootSpanStarted"`

	// End time of the root span for this span collection.
	TimeRootSpanEnded *common.SDKTime `mandatory:"false" json:"timeRootSpanEnded"`

	// Time taken for the root span operation to complete in milliseconds.
	RootSpanDurationInMs *int `mandatory:"false" json:"rootSpanDurationInMs"`

	// A summary of the spans by service
	ServiceSummaries []TraceServiceSummary `mandatory:"false" json:"serviceSummaries"`
}

func (m TraceSpanSummary) String() string {
	return common.PointerString(m)
}
