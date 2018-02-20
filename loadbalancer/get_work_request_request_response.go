// Copyright (c) 2016, 2017, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

package loadbalancer

import (
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

// GetWorkRequestRequest wrapper for the GetWorkRequest operation
type GetWorkRequestRequest struct {

	// The [OCID]({{DOC_SERVER_URL}}/Content/General/Concepts/identifiers.htm) of the work request to retrieve.
	WorkRequestId *string `mandatory:"true" contributesTo:"path" name:"workRequestId"`

	// The unique Oracle-assigned identifier for the request. If you need to contact Oracle about a
	// particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`
}

func (request GetWorkRequestRequest) String() string {
	return common.PointerString(request)
}

// GetHttpRequest implements the OciRequest interface
func (request GetWorkRequestRequest) GetHttpRequest(method, path string) (http.Request, error) {
	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// GetRetryPolicy implements the OciRetryableRequest interface
// => assembles retry policy based on specified options and default behavior
func (request GetWorkRequestRequest) GetRetryPolicy(options ...common.RetryPolicyOption) common.RetryPolicy {
	if len(options) == 0 {
		return common.NoRetryPolicy()
	}
	return common.BuildRetryPolicy(options...)
}

// GetWorkRequestResponse wrapper for the GetWorkRequest operation
type GetWorkRequestResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The WorkRequest instance
	WorkRequest `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	// a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response GetWorkRequestResponse) String() string {
	return common.PointerString(response)
}

// GetRawResponse implements the OciResponse interface
func (response GetWorkRequestResponse) GetRawResponse() *http.Response {
	return response.RawResponse
}

// GetStatefulEntity implements the OciStatefulResponse interface
func (response GetWorkRequestResponse) GetStatefulEntity() common.OciPollable {
	return response.WorkRequest
}
