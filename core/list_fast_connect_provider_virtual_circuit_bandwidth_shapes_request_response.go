// Copyright (c) 2016, 2017, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

package core

import (
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

// ListFastConnectProviderVirtualCircuitBandwidthShapesRequest wrapper for the ListFastConnectProviderVirtualCircuitBandwidthShapes operation
type ListFastConnectProviderVirtualCircuitBandwidthShapesRequest struct {

	// The OCID of the provider service.
	ProviderServiceId *string `mandatory:"true" contributesTo:"path" name:"providerServiceId"`

	// The maximum number of items to return in a paginated "List" call.
	// Example: `500`
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The value of the `opc-next-page` response header from the previous "List" call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListFastConnectProviderVirtualCircuitBandwidthShapesRequest) String() string {
	return common.PointerString(request)
}

// GetHttpRequest implements the OciRequest interface
func (request ListFastConnectProviderVirtualCircuitBandwidthShapesRequest) GetHttpRequest(method, path string) (http.Request, error) {
	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// GetRetryPolicy implements the OciRetryableRequest interface
// => assembles retry policy based on specified options and default behavior
func (request ListFastConnectProviderVirtualCircuitBandwidthShapesRequest) GetRetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ListFastConnectProviderVirtualCircuitBandwidthShapesResponse wrapper for the ListFastConnectProviderVirtualCircuitBandwidthShapes operation
type ListFastConnectProviderVirtualCircuitBandwidthShapesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The []VirtualCircuitBandwidthShape instance
	Items []VirtualCircuitBandwidthShape `presentIn:"body"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	// a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response ListFastConnectProviderVirtualCircuitBandwidthShapesResponse) String() string {
	return common.PointerString(response)
}

// GetRawResponse implements the OciResponse interface
func (response ListFastConnectProviderVirtualCircuitBandwidthShapesResponse) GetRawResponse() *http.Response {
	return response.RawResponse
}
