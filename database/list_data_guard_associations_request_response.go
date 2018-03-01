// Copyright (c) 2016, 2017, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

package database

import (
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

// ListDataGuardAssociationsRequest wrapper for the ListDataGuardAssociations operation
type ListDataGuardAssociationsRequest struct {

	// The database [OCID]({{DOC_SERVER_URL}}/Content/General/Concepts/identifiers.htm).
	DatabaseId *string `mandatory:"true" contributesTo:"path" name:"databaseId"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The pagination token to continue listing from.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListDataGuardAssociationsRequest) String() string {
	return common.PointerString(request)
}

// GetHttpRequest implements the OciRequest interface
func (request ListDataGuardAssociationsRequest) GetHttpRequest(method, path string) (http.Request, error) {
	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// GetRetryPolicy implements the OciRetryableRequest interface
// => assembles retry policy based on specified options and default behavior
func (request ListDataGuardAssociationsRequest) GetRetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ListDataGuardAssociationsResponse wrapper for the ListDataGuardAssociations operation
type ListDataGuardAssociationsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The []DataGuardAssociationSummary instance
	Items []DataGuardAssociationSummary `presentIn:"body"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then there are additional items still to get. Include this value as the `page` parameter for the
	// subsequent GET request. For information about pagination, see
	// [List Pagination]({{DOC_SERVER_URL}}/Content/API/Concepts/usingapi.htm#List_Pagination).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	// a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response ListDataGuardAssociationsResponse) String() string {
	return common.PointerString(response)
}

// GetRawResponse implements the OciResponse interface
func (response ListDataGuardAssociationsResponse) GetRawResponse() *http.Response {
	return response.RawResponse
}
