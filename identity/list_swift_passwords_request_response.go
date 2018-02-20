// Copyright (c) 2016, 2017, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

package identity

import (
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

// ListSwiftPasswordsRequest wrapper for the ListSwiftPasswords operation
type ListSwiftPasswordsRequest struct {

	// The OCID of the user.
	UserId *string `mandatory:"true" contributesTo:"path" name:"userId"`
}

func (request ListSwiftPasswordsRequest) String() string {
	return common.PointerString(request)
}

// GetHttpRequest implements the OciRequest interface
func (request ListSwiftPasswordsRequest) GetHttpRequest(method, path string) (http.Request, error) {
	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// GetRetryPolicy implements the OciRetryableRequest interface
// => assembles retry policy based on specified options and default behavior
func (request ListSwiftPasswordsRequest) GetRetryPolicy(options ...common.RetryPolicyOption) common.RetryPolicy {
	if len(options) == 0 {
		return common.NoRetryPolicy()
	}
	return common.BuildRetryPolicy(options...)
}

// ListSwiftPasswordsResponse wrapper for the ListSwiftPasswords operation
type ListSwiftPasswordsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The []SwiftPassword instance
	Items []SwiftPassword `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a
	// particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListSwiftPasswordsResponse) String() string {
	return common.PointerString(response)
}

// GetRawResponse implements the OciResponse interface
func (response ListSwiftPasswordsResponse) GetRawResponse() *http.Response {
	return response.RawResponse
}
