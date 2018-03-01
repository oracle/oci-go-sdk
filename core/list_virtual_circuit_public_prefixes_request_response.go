// Copyright (c) 2016, 2017, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

package core

import (
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

// ListVirtualCircuitPublicPrefixesRequest wrapper for the ListVirtualCircuitPublicPrefixes operation
type ListVirtualCircuitPublicPrefixesRequest struct {

	// The OCID of the virtual circuit.
	VirtualCircuitId *string `mandatory:"true" contributesTo:"path" name:"virtualCircuitId"`

	// A filter to only return resources that match the given verification state.
	// The state value is case-insensitive.
	VerificationState VirtualCircuitPublicPrefixVerificationStateEnum `mandatory:"false" contributesTo:"query" name:"verificationState" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListVirtualCircuitPublicPrefixesRequest) String() string {
	return common.PointerString(request)
}

// GetHttpRequest implements the OciRequest interface
func (request ListVirtualCircuitPublicPrefixesRequest) GetHttpRequest(method, path string) (http.Request, error) {
	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// GetRetryPolicy implements the OciRetryableRequest interface
// => assembles retry policy based on specified options and default behavior
func (request ListVirtualCircuitPublicPrefixesRequest) GetRetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ListVirtualCircuitPublicPrefixesResponse wrapper for the ListVirtualCircuitPublicPrefixes operation
type ListVirtualCircuitPublicPrefixesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The []VirtualCircuitPublicPrefix instance
	Items []VirtualCircuitPublicPrefix `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	// a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response ListVirtualCircuitPublicPrefixesResponse) String() string {
	return common.PointerString(response)
}

// GetRawResponse implements the OciResponse interface
func (response ListVirtualCircuitPublicPrefixesResponse) GetRawResponse() *http.Response {
	return response.RawResponse
}
