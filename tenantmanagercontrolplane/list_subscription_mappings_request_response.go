// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package tenantmanagercontrolplane

import (
	"github.com/oracle/oci-go-sdk/v55/common"
	"net/http"
)

// ListSubscriptionMappingsRequest wrapper for the ListSubscriptionMappings operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/tenantmanagercontrolplane/ListSubscriptionMappings.go.html to see an example of how to use ListSubscriptionMappingsRequest.
type ListSubscriptionMappingsRequest struct {

	// The ID of the subscription to which the tenancy is associated.
	SubscriptionId *string `mandatory:"false" contributesTo:"query" name:"subscriptionId"`

	// SubscriptionMappingId is a unique ID for subscription and tenancy mapping.
	SubscriptionMappingId *string `mandatory:"false" contributesTo:"query" name:"subscriptionMappingId"`

	// The ID of the compartment in which to list resources.
	CompartmentId *string `mandatory:"false" contributesTo:"query" name:"compartmentId"`

	// The lifecycle state of the resource.
	LifecycleState SubscriptionMappingLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The sort order to use, whether 'asc' or 'desc'.
	SortOrder ListSubscriptionMappingsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order can be provided.
	// * The default order for timeCreated is descending.
	// * The default order for displayName is ascending.
	// * If no value is specified, timeCreated is the default.
	SortBy ListSubscriptionMappingsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListSubscriptionMappingsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListSubscriptionMappingsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListSubscriptionMappingsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListSubscriptionMappingsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ListSubscriptionMappingsResponse wrapper for the ListSubscriptionMappings operation
type ListSubscriptionMappingsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of SubscriptionMappingCollection instances
	SubscriptionMappingCollection `presentIn:"body"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response ListSubscriptionMappingsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListSubscriptionMappingsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListSubscriptionMappingsSortOrderEnum Enum with underlying type: string
type ListSubscriptionMappingsSortOrderEnum string

// Set of constants representing the allowable values for ListSubscriptionMappingsSortOrderEnum
const (
	ListSubscriptionMappingsSortOrderAsc  ListSubscriptionMappingsSortOrderEnum = "ASC"
	ListSubscriptionMappingsSortOrderDesc ListSubscriptionMappingsSortOrderEnum = "DESC"
)

var mappingListSubscriptionMappingsSortOrder = map[string]ListSubscriptionMappingsSortOrderEnum{
	"ASC":  ListSubscriptionMappingsSortOrderAsc,
	"DESC": ListSubscriptionMappingsSortOrderDesc,
}

// GetListSubscriptionMappingsSortOrderEnumValues Enumerates the set of values for ListSubscriptionMappingsSortOrderEnum
func GetListSubscriptionMappingsSortOrderEnumValues() []ListSubscriptionMappingsSortOrderEnum {
	values := make([]ListSubscriptionMappingsSortOrderEnum, 0)
	for _, v := range mappingListSubscriptionMappingsSortOrder {
		values = append(values, v)
	}
	return values
}

// ListSubscriptionMappingsSortByEnum Enum with underlying type: string
type ListSubscriptionMappingsSortByEnum string

// Set of constants representing the allowable values for ListSubscriptionMappingsSortByEnum
const (
	ListSubscriptionMappingsSortByTimecreated ListSubscriptionMappingsSortByEnum = "timeCreated"
	ListSubscriptionMappingsSortByDisplayname ListSubscriptionMappingsSortByEnum = "displayName"
)

var mappingListSubscriptionMappingsSortBy = map[string]ListSubscriptionMappingsSortByEnum{
	"timeCreated": ListSubscriptionMappingsSortByTimecreated,
	"displayName": ListSubscriptionMappingsSortByDisplayname,
}

// GetListSubscriptionMappingsSortByEnumValues Enumerates the set of values for ListSubscriptionMappingsSortByEnum
func GetListSubscriptionMappingsSortByEnumValues() []ListSubscriptionMappingsSortByEnum {
	values := make([]ListSubscriptionMappingsSortByEnum, 0)
	for _, v := range mappingListSubscriptionMappingsSortBy {
		values = append(values, v)
	}
	return values
}
