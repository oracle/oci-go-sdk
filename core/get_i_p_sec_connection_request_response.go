// Copyright (c) 2016, 2017, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

package core

import (
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

// GetIPSecConnectionRequest wrapper for the GetIPSecConnection operation
type GetIPSecConnectionRequest struct {

	// The OCID of the IPSec connection.
	IpscId *string `mandatory:"true" contributesTo:"path" name:"ipscId"`
}

func (request GetIPSecConnectionRequest) String() string {
	return common.PointerString(request)
}

// GetHttpRequest implements the OciRequest interface
func (request GetIPSecConnectionRequest) GetHttpRequest(method, path string) (http.Request, error) {
	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// GetRetryPolicy implements the OciRetryableRequest interface
// => assembles retry policy based on specified options and default behavior
func (request GetIPSecConnectionRequest) GetRetryPolicy(options ...common.RetryPolicyOption) common.RetryPolicy {
	if len(options) == 0 {
		return common.NoRetryPolicy()
	}
	return common.BuildRetryPolicy(options...)
}

// GetIPSecConnectionResponse wrapper for the GetIPSecConnection operation
type GetIPSecConnectionResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The IpSecConnection instance
	IpSecConnection `presentIn:"body"`

	// For optimistic concurrency control. See `if-match`.
	Etag *string `presentIn:"header" name:"etag"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	// a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response GetIPSecConnectionResponse) String() string {
	return common.PointerString(response)
}

// GetRawResponse implements the OciResponse interface
func (response GetIPSecConnectionResponse) GetRawResponse() *http.Response {
	return response.RawResponse
}
