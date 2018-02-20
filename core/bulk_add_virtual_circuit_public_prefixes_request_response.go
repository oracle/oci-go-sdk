// Copyright (c) 2016, 2017, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

package core

import (
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

// BulkAddVirtualCircuitPublicPrefixesRequest wrapper for the BulkAddVirtualCircuitPublicPrefixes operation
type BulkAddVirtualCircuitPublicPrefixesRequest struct {

	// The OCID of the virtual circuit.
	VirtualCircuitId *string `mandatory:"true" contributesTo:"path" name:"virtualCircuitId"`

	// Request with publix prefixes to be added to the virtual circuit
	BulkAddVirtualCircuitPublicPrefixesDetails `contributesTo:"body"`
}

func (request BulkAddVirtualCircuitPublicPrefixesRequest) String() string {
	return common.PointerString(request)
}

// GetHttpRequest implements the OciRequest interface
func (request BulkAddVirtualCircuitPublicPrefixesRequest) GetHttpRequest(method, path string) (http.Request, error) {
	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// GetRetryPolicy implements the OciRetryableRequest interface
// => assembles retry policy based on specified options and default behavior
func (request BulkAddVirtualCircuitPublicPrefixesRequest) GetRetryPolicy(options ...common.RetryPolicyOption) common.RetryPolicy {
	if len(options) == 0 {
		return common.NoRetryPolicy()
	}
	return common.BuildRetryPolicy(options...)
}

// BulkAddVirtualCircuitPublicPrefixesResponse wrapper for the BulkAddVirtualCircuitPublicPrefixes operation
type BulkAddVirtualCircuitPublicPrefixesResponse struct {

	// The underlying http response
	RawResponse *http.Response
}

func (response BulkAddVirtualCircuitPublicPrefixesResponse) String() string {
	return common.PointerString(response)
}

// GetRawResponse implements the OciResponse interface
func (response BulkAddVirtualCircuitPublicPrefixesResponse) GetRawResponse() *http.Response {
	return response.RawResponse
}
