// Copyright (c) 2016, 2017, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

package core

import (
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

// GetCrossConnectStatusRequest wrapper for the GetCrossConnectStatus operation
type GetCrossConnectStatusRequest struct {

	// The OCID of the cross-connect.
	CrossConnectId *string `mandatory:"true" contributesTo:"path" name:"crossConnectId"`
}

func (request GetCrossConnectStatusRequest) String() string {
	return common.PointerString(request)
}

// GetHttpRequest implements the OciRequest interface
func (request GetCrossConnectStatusRequest) GetHttpRequest(method, path string) (http.Request, error) {
	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// GetRetryPolicy implements the OciRetryableRequest interface
// => assembles retry policy based on specified options and default behavior
func (request GetCrossConnectStatusRequest) GetRetryPolicy(options ...common.RetryPolicyOption) common.RetryPolicy {
	if len(options) == 0 {
		return common.NoRetryPolicy()
	}
	return common.BuildRetryPolicy(options...)
}

// GetCrossConnectStatusResponse wrapper for the GetCrossConnectStatus operation
type GetCrossConnectStatusResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The CrossConnectStatus instance
	CrossConnectStatus `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	// a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response GetCrossConnectStatusResponse) String() string {
	return common.PointerString(response)
}

// GetRawResponse implements the OciResponse interface
func (response GetCrossConnectStatusResponse) GetRawResponse() *http.Response {
	return response.RawResponse
}
