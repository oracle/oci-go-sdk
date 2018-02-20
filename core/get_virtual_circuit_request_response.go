// Copyright (c) 2016, 2017, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

package core

import (
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

// GetVirtualCircuitRequest wrapper for the GetVirtualCircuit operation
type GetVirtualCircuitRequest struct {

	// The OCID of the virtual circuit.
	VirtualCircuitId *string `mandatory:"true" contributesTo:"path" name:"virtualCircuitId"`
}

func (request GetVirtualCircuitRequest) String() string {
	return common.PointerString(request)
}

// GetHttpRequest implements the OciRequest interface
func (request GetVirtualCircuitRequest) GetHttpRequest(method, path string) (http.Request, error) {
	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// GetRetryPolicy implements the OciRetryableRequest interface
// => assembles retry policy based on specified options and default behavior
func (request GetVirtualCircuitRequest) GetRetryPolicy(options ...common.RetryPolicyOption) common.RetryPolicy {
	if len(options) == 0 {
		return common.NoRetryPolicy()
	}
	return common.BuildRetryPolicy(options...)
}

// GetVirtualCircuitResponse wrapper for the GetVirtualCircuit operation
type GetVirtualCircuitResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The VirtualCircuit instance
	VirtualCircuit `presentIn:"body"`

	// For optimistic concurrency control. See `if-match`.
	Etag *string `presentIn:"header" name:"etag"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	// a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response GetVirtualCircuitResponse) String() string {
	return common.PointerString(response)
}

// GetRawResponse implements the OciResponse interface
func (response GetVirtualCircuitResponse) GetRawResponse() *http.Response {
	return response.RawResponse
}

// GetStatefulEntity implements the OciStatefulResponse interface
func (response GetVirtualCircuitResponse) GetStatefulEntity() common.OciPollable {
	return response.VirtualCircuit
}
