// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package tenantmanagercontrolplane

import (
	"github.com/oracle/oci-go-sdk/v56/common"
	"net/http"
)

// ListSubscriptionsRequest wrapper for the ListSubscriptions operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/tenantmanagercontrolplane/ListSubscriptions.go.html to see an example of how to use ListSubscriptionsRequest.
type ListSubscriptionsRequest struct {

	// The ID of the compartment in which to list resources.
	CompartmentId *string `mandatory:"false" contributesTo:"query" name:"compartmentId"`

	// The ID of the subscription to which the tenancy is associated.
	SubscriptionId *string `mandatory:"false" contributesTo:"query" name:"subscriptionId"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The sort order to use, whether 'asc' or 'desc'.
	SortOrder ListSubscriptionsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order can be provided.
	// * The default order for timeCreated is descending.
	// * The default order for displayName is ascending.
	// * If no value is specified, timeCreated is the default.
	SortBy ListSubscriptionsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListSubscriptionsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListSubscriptionsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListSubscriptionsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListSubscriptionsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ListSubscriptionsResponse wrapper for the ListSubscriptions operation
type ListSubscriptionsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of SubscriptionCollection instances
	SubscriptionCollection `presentIn:"body"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response ListSubscriptionsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListSubscriptionsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListSubscriptionsSortOrderEnum Enum with underlying type: string
type ListSubscriptionsSortOrderEnum string

// Set of constants representing the allowable values for ListSubscriptionsSortOrderEnum
const (
	ListSubscriptionsSortOrderAsc  ListSubscriptionsSortOrderEnum = "ASC"
	ListSubscriptionsSortOrderDesc ListSubscriptionsSortOrderEnum = "DESC"
)

var mappingListSubscriptionsSortOrder = map[string]ListSubscriptionsSortOrderEnum{
	"ASC":  ListSubscriptionsSortOrderAsc,
	"DESC": ListSubscriptionsSortOrderDesc,
}

// GetListSubscriptionsSortOrderEnumValues Enumerates the set of values for ListSubscriptionsSortOrderEnum
func GetListSubscriptionsSortOrderEnumValues() []ListSubscriptionsSortOrderEnum {
	values := make([]ListSubscriptionsSortOrderEnum, 0)
	for _, v := range mappingListSubscriptionsSortOrder {
		values = append(values, v)
	}
	return values
}

// ListSubscriptionsSortByEnum Enum with underlying type: string
type ListSubscriptionsSortByEnum string

// Set of constants representing the allowable values for ListSubscriptionsSortByEnum
const (
	ListSubscriptionsSortByTimecreated ListSubscriptionsSortByEnum = "timeCreated"
	ListSubscriptionsSortByDisplayname ListSubscriptionsSortByEnum = "displayName"
)

var mappingListSubscriptionsSortBy = map[string]ListSubscriptionsSortByEnum{
	"timeCreated": ListSubscriptionsSortByTimecreated,
	"displayName": ListSubscriptionsSortByDisplayname,
}

// GetListSubscriptionsSortByEnumValues Enumerates the set of values for ListSubscriptionsSortByEnum
func GetListSubscriptionsSortByEnumValues() []ListSubscriptionsSortByEnum {
	values := make([]ListSubscriptionsSortByEnum, 0)
	for _, v := range mappingListSubscriptionsSortBy {
		values = append(values, v)
	}
	return values
}
