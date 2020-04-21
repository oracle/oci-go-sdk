// Copyright (c) 2016, 2018, , Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package healthchecks

import (
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

// ChangePingMonitorCompartmentRequest wrapper for the ChangePingMonitorCompartment operation
type ChangePingMonitorCompartmentRequest struct {

	// The OCID of a monitor.
	MonitorId *string `mandatory:"true" contributesTo:"path" name:"monitorId"`

	// The details needed to move the monitor.
	ChangePingMonitorCompartmentDetails `contributesTo:"body"`

	// For optimistic concurrency control. In the PUT or DELETE call for a resource,
	// set the `if-match` parameter to the value of the etag from a previous GET
	// or POST response for that resource.  The resource will be updated or deleted
	// only if the etag you provide matches the resource's current etag value.
	IfMatch *string `mandatory:"false" contributesTo:"header" name:"if-match"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a
	// particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// A token that uniquely identifies a request that can be retried in case of a timeout or
	// server error without risk of executing the same action again. Retry tokens expire after 24
	// hours.
	// *Note:* Retry tokens can be invalidated before the 24 hour time limit due to conflicting
	// operations, such as a resource being deleted or purged from the system.
	OpcRetryToken *string `mandatory:"false" contributesTo:"header" name:"opc-retry-token"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ChangePingMonitorCompartmentRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ChangePingMonitorCompartmentRequest) HTTPRequest(method, path string) (http.Request, error) {
	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ChangePingMonitorCompartmentRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ChangePingMonitorCompartmentResponse wrapper for the ChangePingMonitorCompartment operation
type ChangePingMonitorCompartmentResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// Unique Oracle-assigned identifier for the request. If you need to
	// contact Oracle about a particular request, please provide
	// the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response ChangePingMonitorCompartmentResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ChangePingMonitorCompartmentResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}
