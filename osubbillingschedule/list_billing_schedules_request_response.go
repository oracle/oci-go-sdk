// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package osubbillingschedule

import (
	"github.com/oracle/oci-go-sdk/v56/common"
	"net/http"
)

// ListBillingSchedulesRequest wrapper for the ListBillingSchedules operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/osubbillingschedule/ListBillingSchedules.go.html to see an example of how to use ListBillingSchedulesRequest.
type ListBillingSchedulesRequest struct {

	// The OCID of the compartment.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// This param is used to get only the billing schedules for a particular Subscription Id
	SubscriptionId *string `mandatory:"true" contributesTo:"query" name:"subscriptionId"`

	// This param is used to get only the billing schedules for a particular Subscribed Service
	SubscribedServiceId *string `mandatory:"false" contributesTo:"query" name:"subscribedServiceId"`

	// The maximum number of items to return in a paginated "List" call. Default: (`50`)
	// Example: '500'
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The value of the 'opc-next-page' response header from the previous "List" call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either ascending ('ASC') or descending ('DESC').
	SortOrder ListBillingSchedulesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. You can provide one sort order ('sortOrder').
	SortBy ListBillingSchedulesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// The OCI home region name in case home region is not us-ashburn-1 (IAD), e.g. ap-mumbai-1, us-phoenix-1 etc.
	XOneOriginRegion *string `mandatory:"false" contributesTo:"header" name:"x-one-origin-region"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListBillingSchedulesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListBillingSchedulesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListBillingSchedulesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListBillingSchedulesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ListBillingSchedulesResponse wrapper for the ListBillingSchedules operation
type ListBillingSchedulesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []BillingScheduleSummary instances
	Items []BillingScheduleSummary `presentIn:"body"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the 'page' parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a
	// particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response ListBillingSchedulesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListBillingSchedulesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListBillingSchedulesSortOrderEnum Enum with underlying type: string
type ListBillingSchedulesSortOrderEnum string

// Set of constants representing the allowable values for ListBillingSchedulesSortOrderEnum
const (
	ListBillingSchedulesSortOrderAsc  ListBillingSchedulesSortOrderEnum = "ASC"
	ListBillingSchedulesSortOrderDesc ListBillingSchedulesSortOrderEnum = "DESC"
)

var mappingListBillingSchedulesSortOrder = map[string]ListBillingSchedulesSortOrderEnum{
	"ASC":  ListBillingSchedulesSortOrderAsc,
	"DESC": ListBillingSchedulesSortOrderDesc,
}

// GetListBillingSchedulesSortOrderEnumValues Enumerates the set of values for ListBillingSchedulesSortOrderEnum
func GetListBillingSchedulesSortOrderEnumValues() []ListBillingSchedulesSortOrderEnum {
	values := make([]ListBillingSchedulesSortOrderEnum, 0)
	for _, v := range mappingListBillingSchedulesSortOrder {
		values = append(values, v)
	}
	return values
}

// ListBillingSchedulesSortByEnum Enum with underlying type: string
type ListBillingSchedulesSortByEnum string

// Set of constants representing the allowable values for ListBillingSchedulesSortByEnum
const (
	ListBillingSchedulesSortByOrdernumber   ListBillingSchedulesSortByEnum = "ORDERNUMBER"
	ListBillingSchedulesSortByTimeinvoicing ListBillingSchedulesSortByEnum = "TIMEINVOICING"
)

var mappingListBillingSchedulesSortBy = map[string]ListBillingSchedulesSortByEnum{
	"ORDERNUMBER":   ListBillingSchedulesSortByOrdernumber,
	"TIMEINVOICING": ListBillingSchedulesSortByTimeinvoicing,
}

// GetListBillingSchedulesSortByEnumValues Enumerates the set of values for ListBillingSchedulesSortByEnum
func GetListBillingSchedulesSortByEnumValues() []ListBillingSchedulesSortByEnum {
	values := make([]ListBillingSchedulesSortByEnum, 0)
	for _, v := range mappingListBillingSchedulesSortBy {
		values = append(values, v)
	}
	return values
}
