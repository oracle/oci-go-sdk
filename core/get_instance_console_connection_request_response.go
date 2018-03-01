// Copyright (c) 2016, 2017, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

package core

import (
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

// GetInstanceConsoleConnectionRequest wrapper for the GetInstanceConsoleConnection operation
type GetInstanceConsoleConnectionRequest struct {

	// The OCID of the intance console connection
	InstanceConsoleConnectionId *string `mandatory:"true" contributesTo:"path" name:"instanceConsoleConnectionId"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request GetInstanceConsoleConnectionRequest) String() string {
	return common.PointerString(request)
}

// GetHttpRequest implements the OciRequest interface
func (request GetInstanceConsoleConnectionRequest) GetHttpRequest(method, path string) (http.Request, error) {
	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// GetRetryPolicy implements the OciRetryableRequest interface
// => assembles retry policy based on specified options and default behavior
func (request GetInstanceConsoleConnectionRequest) GetRetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// GetInstanceConsoleConnectionResponse wrapper for the GetInstanceConsoleConnection operation
type GetInstanceConsoleConnectionResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The InstanceConsoleConnection instance
	InstanceConsoleConnection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	// a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response GetInstanceConsoleConnectionResponse) String() string {
	return common.PointerString(response)
}

// GetRawResponse implements the OciResponse interface
func (response GetInstanceConsoleConnectionResponse) GetRawResponse() *http.Response {
	return response.RawResponse
}

// GetStatefulEntity implements the OciStatefulResponse interface
func (response GetInstanceConsoleConnectionResponse) GetStatefulEntity() common.OciPollable {
	return response.InstanceConsoleConnection
}
