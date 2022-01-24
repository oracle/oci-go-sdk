// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package osuborganizationsubscription

import (
	"github.com/oracle/oci-go-sdk/v56/common"
	"net/http"
)

// ListOrganizationSubscriptionsRequest wrapper for the ListOrganizationSubscriptions operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/osuborganizationsubscription/ListOrganizationSubscriptions.go.html to see an example of how to use ListOrganizationSubscriptionsRequest.
type ListOrganizationSubscriptionsRequest struct {

	// The OCID of the compartment.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// Comma separated list of subscription ids
	SubscriptionIds *string `mandatory:"true" contributesTo:"query" name:"subscriptionIds"`

	// The maximum number of items to return in a paginated "List" call. Default: (`50`)
	// Example: `500`
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The value of the `opc-next-page` response header from the previous "List" call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either ascending (`ASC`) or descending (`DESC`).
	SortOrder ListOrganizationSubscriptionsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. You can provide one sort order (`sortOrder`).
	SortBy ListOrganizationSubscriptionsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// The OCI home region name in case home region is not us-ashburn-1 (IAD), e.g. ap-mumbai-1, us-phoenix-1 etc.
	XOneOriginRegion *string `mandatory:"false" contributesTo:"header" name:"x-one-origin-region"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListOrganizationSubscriptionsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListOrganizationSubscriptionsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListOrganizationSubscriptionsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListOrganizationSubscriptionsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ListOrganizationSubscriptionsResponse wrapper for the ListOrganizationSubscriptions operation
type ListOrganizationSubscriptionsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []SubscriptionSummary instances
	Items []SubscriptionSummary `presentIn:"body"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a
	// particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response ListOrganizationSubscriptionsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListOrganizationSubscriptionsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListOrganizationSubscriptionsSortOrderEnum Enum with underlying type: string
type ListOrganizationSubscriptionsSortOrderEnum string

// Set of constants representing the allowable values for ListOrganizationSubscriptionsSortOrderEnum
const (
	ListOrganizationSubscriptionsSortOrderAsc  ListOrganizationSubscriptionsSortOrderEnum = "ASC"
	ListOrganizationSubscriptionsSortOrderDesc ListOrganizationSubscriptionsSortOrderEnum = "DESC"
)

var mappingListOrganizationSubscriptionsSortOrder = map[string]ListOrganizationSubscriptionsSortOrderEnum{
	"ASC":  ListOrganizationSubscriptionsSortOrderAsc,
	"DESC": ListOrganizationSubscriptionsSortOrderDesc,
}

// GetListOrganizationSubscriptionsSortOrderEnumValues Enumerates the set of values for ListOrganizationSubscriptionsSortOrderEnum
func GetListOrganizationSubscriptionsSortOrderEnumValues() []ListOrganizationSubscriptionsSortOrderEnum {
	values := make([]ListOrganizationSubscriptionsSortOrderEnum, 0)
	for _, v := range mappingListOrganizationSubscriptionsSortOrder {
		values = append(values, v)
	}
	return values
}

// ListOrganizationSubscriptionsSortByEnum Enum with underlying type: string
type ListOrganizationSubscriptionsSortByEnum string

// Set of constants representing the allowable values for ListOrganizationSubscriptionsSortByEnum
const (
	ListOrganizationSubscriptionsSortBySubscriptionid ListOrganizationSubscriptionsSortByEnum = "SUBSCRIPTIONID"
	ListOrganizationSubscriptionsSortByTimestart      ListOrganizationSubscriptionsSortByEnum = "TIMESTART"
)

var mappingListOrganizationSubscriptionsSortBy = map[string]ListOrganizationSubscriptionsSortByEnum{
	"SUBSCRIPTIONID": ListOrganizationSubscriptionsSortBySubscriptionid,
	"TIMESTART":      ListOrganizationSubscriptionsSortByTimestart,
}

// GetListOrganizationSubscriptionsSortByEnumValues Enumerates the set of values for ListOrganizationSubscriptionsSortByEnum
func GetListOrganizationSubscriptionsSortByEnumValues() []ListOrganizationSubscriptionsSortByEnum {
	values := make([]ListOrganizationSubscriptionsSortByEnum, 0)
	for _, v := range mappingListOrganizationSubscriptionsSortBy {
		values = append(values, v)
	}
	return values
}
