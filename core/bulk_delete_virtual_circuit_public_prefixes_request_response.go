// Copyright (c) 2016, 2017, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

package core

import (
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

// BulkDeleteVirtualCircuitPublicPrefixesRequest wrapper for the BulkDeleteVirtualCircuitPublicPrefixes operation
type BulkDeleteVirtualCircuitPublicPrefixesRequest struct {

	// The OCID of the virtual circuit.
	VirtualCircuitId *string `mandatory:"true" contributesTo:"path" name:"virtualCircuitId"`

	// Request with publix prefixes to be deleted from the virtual circuit
	BulkDeleteVirtualCircuitPublicPrefixesDetails `contributesTo:"body"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request BulkDeleteVirtualCircuitPublicPrefixesRequest) String() string {
	return common.PointerString(request)
}

// GetHttpRequest implements the OciRequest interface
func (request BulkDeleteVirtualCircuitPublicPrefixesRequest) GetHttpRequest(method, path string) (http.Request, error) {
	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// GetRetryPolicy implements the OciRetryableRequest interface
// => assembles retry policy based on specified options and default behavior
func (request BulkDeleteVirtualCircuitPublicPrefixesRequest) GetRetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// BulkDeleteVirtualCircuitPublicPrefixesResponse wrapper for the BulkDeleteVirtualCircuitPublicPrefixes operation
type BulkDeleteVirtualCircuitPublicPrefixesResponse struct {

	// The underlying http response
	RawResponse *http.Response
}

func (response BulkDeleteVirtualCircuitPublicPrefixesResponse) String() string {
	return common.PointerString(response)
}

// GetRawResponse implements the OciResponse interface
func (response BulkDeleteVirtualCircuitPublicPrefixesResponse) GetRawResponse() *http.Response {
	return response.RawResponse
}
