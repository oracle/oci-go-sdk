// Copyright (c) 2016, 2017, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

package core

import (
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

// GetConsoleHistoryRequest wrapper for the GetConsoleHistory operation
type GetConsoleHistoryRequest struct {

	// The OCID of the console history.
	InstanceConsoleHistoryId *string `mandatory:"true" contributesTo:"path" name:"instanceConsoleHistoryId"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request GetConsoleHistoryRequest) String() string {
	return common.PointerString(request)
}

// GetHttpRequest implements the OciRequest interface
func (request GetConsoleHistoryRequest) GetHttpRequest(method, path string) (http.Request, error) {
	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// GetRetryPolicy implements the OciRetryableRequest interface
// => assembles retry policy based on specified options and default behavior
func (request GetConsoleHistoryRequest) GetRetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// GetConsoleHistoryResponse wrapper for the GetConsoleHistory operation
type GetConsoleHistoryResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The ConsoleHistory instance
	ConsoleHistory `presentIn:"body"`

	// For optimistic concurrency control. See `if-match`.
	Etag *string `presentIn:"header" name:"etag"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	// a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response GetConsoleHistoryResponse) String() string {
	return common.PointerString(response)
}

// GetRawResponse implements the OciResponse interface
func (response GetConsoleHistoryResponse) GetRawResponse() *http.Response {
	return response.RawResponse
}

// GetStatefulEntity implements the OciStatefulResponse interface
func (response GetConsoleHistoryResponse) GetStatefulEntity() common.OciPollable {
	return response.ConsoleHistory
}
