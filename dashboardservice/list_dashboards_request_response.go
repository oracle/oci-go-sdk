// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package dashboardservice

import (
	"github.com/oracle/oci-go-sdk/v56/common"
	"net/http"
)

// ListDashboardsRequest wrapper for the ListDashboards operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/dashboardservice/ListDashboards.go.html to see an example of how to use ListDashboardsRequest.
type ListDashboardsRequest struct {

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the dashboard group that the dashboard belongs to.
	DashboardGroupId *string `mandatory:"true" contributesTo:"query" name:"dashboardGroupId"`

	// A filter that returns dashboard resources that match the lifecycle state specified.
	LifecycleState DashboardLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// A case-sensitive filter that returns resources that match the entire display name specified.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the dashboard.
	Id *string `mandatory:"false" contributesTo:"query" name:"id"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This value is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either ascending (`ASC`) or descending (`DESC`).
	SortOrder ListDashboardsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. You can provide one sort order (`sortOrder`).
	// Default order for TIMECREATED is descending.
	// Default order for DISPLAYNAME is ascending.
	// The DISPLAYNAME sort order is case sensitive.
	SortBy ListDashboardsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// To identify if the call is cross-regional. In CRUD calls for a resource, to
	// identify that the call originates from different region, set the
	// `CrossRegionIdentifierHeader` parameter to a region name (ex - `US-ASHBURN-1`)
	// The call will be served from a Replicated bucket.
	// For same-region calls, the value is unassigned.
	OpcCrossRegion *string `mandatory:"false" contributesTo:"header" name:"opc-cross-region"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListDashboardsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListDashboardsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListDashboardsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListDashboardsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ListDashboardsResponse wrapper for the ListDashboards operation
type ListDashboardsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of DashboardCollection instances
	DashboardCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListDashboardsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListDashboardsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListDashboardsSortOrderEnum Enum with underlying type: string
type ListDashboardsSortOrderEnum string

// Set of constants representing the allowable values for ListDashboardsSortOrderEnum
const (
	ListDashboardsSortOrderAsc  ListDashboardsSortOrderEnum = "ASC"
	ListDashboardsSortOrderDesc ListDashboardsSortOrderEnum = "DESC"
)

var mappingListDashboardsSortOrder = map[string]ListDashboardsSortOrderEnum{
	"ASC":  ListDashboardsSortOrderAsc,
	"DESC": ListDashboardsSortOrderDesc,
}

// GetListDashboardsSortOrderEnumValues Enumerates the set of values for ListDashboardsSortOrderEnum
func GetListDashboardsSortOrderEnumValues() []ListDashboardsSortOrderEnum {
	values := make([]ListDashboardsSortOrderEnum, 0)
	for _, v := range mappingListDashboardsSortOrder {
		values = append(values, v)
	}
	return values
}

// ListDashboardsSortByEnum Enum with underlying type: string
type ListDashboardsSortByEnum string

// Set of constants representing the allowable values for ListDashboardsSortByEnum
const (
	ListDashboardsSortByTimecreated ListDashboardsSortByEnum = "timeCreated"
	ListDashboardsSortByDisplayname ListDashboardsSortByEnum = "displayName"
)

var mappingListDashboardsSortBy = map[string]ListDashboardsSortByEnum{
	"timeCreated": ListDashboardsSortByTimecreated,
	"displayName": ListDashboardsSortByDisplayname,
}

// GetListDashboardsSortByEnumValues Enumerates the set of values for ListDashboardsSortByEnum
func GetListDashboardsSortByEnumValues() []ListDashboardsSortByEnum {
	values := make([]ListDashboardsSortByEnum, 0)
	for _, v := range mappingListDashboardsSortBy {
		values = append(values, v)
	}
	return values
}
