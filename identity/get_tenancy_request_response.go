// Copyright (c) 2016, 2017, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

package identity

import (
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

// GetTenancyRequest wrapper for the GetTenancy operation
type GetTenancyRequest struct {

	// The OCID of the tenancy.
	TenancyId *string `mandatory:"true" contributesTo:"path" name:"tenancyId"`
}

func (request GetTenancyRequest) String() string {
	return common.PointerString(request)
}

// GetHttpRequest implements the OciRequest interface
func (request GetTenancyRequest) GetHttpRequest(method, path string) (http.Request, error) {
	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// GetRetryPolicy implements the OciRetryableRequest interface
// => assembles retry policy based on specified options and default behavior
func (request GetTenancyRequest) GetRetryPolicy(options ...common.RetryPolicyOption) common.RetryPolicy {
	if len(options) == 0 {
		return common.NoRetryPolicy()
	}
	return common.BuildRetryPolicy(options...)
}

// GetTenancyResponse wrapper for the GetTenancy operation
type GetTenancyResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The Tenancy instance
	Tenancy `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a
	// particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response GetTenancyResponse) String() string {
	return common.PointerString(response)
}

// GetRawResponse implements the OciResponse interface
func (response GetTenancyResponse) GetRawResponse() *http.Response {
	return response.RawResponse
}
