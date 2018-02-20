// Copyright (c) 2016, 2017, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

package identity

import (
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

// ListIdpGroupMappingsRequest wrapper for the ListIdpGroupMappings operation
type ListIdpGroupMappingsRequest struct {

	// The OCID of the identity provider.
	IdentityProviderId *string `mandatory:"true" contributesTo:"path" name:"identityProviderId"`

	// The value of the `opc-next-page` response header from the previous "List" call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The maximum number of items to return in a paginated "List" call.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`
}

func (request ListIdpGroupMappingsRequest) String() string {
	return common.PointerString(request)
}

// GetHttpRequest implements the OciRequest interface
func (request ListIdpGroupMappingsRequest) GetHttpRequest(method, path string) (http.Request, error) {
	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// GetRetryPolicy implements the OciRetryableRequest interface
// => assembles retry policy based on specified options and default behavior
func (request ListIdpGroupMappingsRequest) GetRetryPolicy(options ...common.RetryPolicyOption) common.RetryPolicy {
	if len(options) == 0 {
		return common.NoRetryPolicy()
	}
	return common.BuildRetryPolicy(options...)
}

// ListIdpGroupMappingsResponse wrapper for the ListIdpGroupMappings operation
type ListIdpGroupMappingsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The []IdpGroupMapping instance
	Items []IdpGroupMapping `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a
	// particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListIdpGroupMappingsResponse) String() string {
	return common.PointerString(response)
}

// GetRawResponse implements the OciResponse interface
func (response ListIdpGroupMappingsResponse) GetRawResponse() *http.Response {
	return response.RawResponse
}
