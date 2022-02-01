// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package dashboardservice

import (
	"github.com/oracle/oci-go-sdk/v56/common"
	"net/http"
)

// ListDashboardGroupsRequest wrapper for the ListDashboardGroups operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/dashboardservice/ListDashboardGroups.go.html to see an example of how to use ListDashboardGroupsRequest.
type ListDashboardGroupsRequest struct {

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment in which to list resources.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// A filter that returns dashboard groups that match the lifecycle state specified.
	LifecycleState DashboardGroupLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// A case-sensitive filter that returns resources that match the entire display name specified.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the dashboard group.
	Id *string `mandatory:"false" contributesTo:"query" name:"id"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This value is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either ascending (`ASC`) or descending (`DESC`).
	SortOrder ListDashboardGroupsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. You can provide one sort order (`sortOrder`).
	// Default order for TIMECREATED is descending.
	// Default order for DISPLAYNAME is ascending.
	// The DISPLAYNAME sort order is case sensitive.
	SortBy ListDashboardGroupsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

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

func (request ListDashboardGroupsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListDashboardGroupsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListDashboardGroupsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListDashboardGroupsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ListDashboardGroupsResponse wrapper for the ListDashboardGroups operation
type ListDashboardGroupsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of DashboardGroupCollection instances
	DashboardGroupCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListDashboardGroupsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListDashboardGroupsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListDashboardGroupsSortOrderEnum Enum with underlying type: string
type ListDashboardGroupsSortOrderEnum string

// Set of constants representing the allowable values for ListDashboardGroupsSortOrderEnum
const (
	ListDashboardGroupsSortOrderAsc  ListDashboardGroupsSortOrderEnum = "ASC"
	ListDashboardGroupsSortOrderDesc ListDashboardGroupsSortOrderEnum = "DESC"
)

var mappingListDashboardGroupsSortOrder = map[string]ListDashboardGroupsSortOrderEnum{
	"ASC":  ListDashboardGroupsSortOrderAsc,
	"DESC": ListDashboardGroupsSortOrderDesc,
}

// GetListDashboardGroupsSortOrderEnumValues Enumerates the set of values for ListDashboardGroupsSortOrderEnum
func GetListDashboardGroupsSortOrderEnumValues() []ListDashboardGroupsSortOrderEnum {
	values := make([]ListDashboardGroupsSortOrderEnum, 0)
	for _, v := range mappingListDashboardGroupsSortOrder {
		values = append(values, v)
	}
	return values
}

// ListDashboardGroupsSortByEnum Enum with underlying type: string
type ListDashboardGroupsSortByEnum string

// Set of constants representing the allowable values for ListDashboardGroupsSortByEnum
const (
	ListDashboardGroupsSortByTimecreated ListDashboardGroupsSortByEnum = "timeCreated"
	ListDashboardGroupsSortByDisplayname ListDashboardGroupsSortByEnum = "displayName"
)

var mappingListDashboardGroupsSortBy = map[string]ListDashboardGroupsSortByEnum{
	"timeCreated": ListDashboardGroupsSortByTimecreated,
	"displayName": ListDashboardGroupsSortByDisplayname,
}

// GetListDashboardGroupsSortByEnumValues Enumerates the set of values for ListDashboardGroupsSortByEnum
func GetListDashboardGroupsSortByEnumValues() []ListDashboardGroupsSortByEnum {
	values := make([]ListDashboardGroupsSortByEnum, 0)
	for _, v := range mappingListDashboardGroupsSortBy {
		values = append(values, v)
	}
	return values
}
