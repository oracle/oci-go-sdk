// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package dashboardservice

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListDashboardGroupsRequest wrapper for the ListDashboardGroups operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/dashboardservice/ListDashboardGroups.go.html to see an example of how to use ListDashboardGroupsRequest.
type ListDashboardGroupsRequest struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment in which to list resources.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// A filter that returns dashboard groups that match the lifecycle state specified.
	LifecycleState DashboardGroupLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// A case-sensitive filter that returns resources that match the entire display name specified.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the dashboard group.
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

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
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

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListDashboardGroupsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingDashboardGroupLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetDashboardGroupLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListDashboardGroupsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListDashboardGroupsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListDashboardGroupsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListDashboardGroupsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
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

var mappingListDashboardGroupsSortOrderEnum = map[string]ListDashboardGroupsSortOrderEnum{
	"ASC":  ListDashboardGroupsSortOrderAsc,
	"DESC": ListDashboardGroupsSortOrderDesc,
}

var mappingListDashboardGroupsSortOrderEnumLowerCase = map[string]ListDashboardGroupsSortOrderEnum{
	"asc":  ListDashboardGroupsSortOrderAsc,
	"desc": ListDashboardGroupsSortOrderDesc,
}

// GetListDashboardGroupsSortOrderEnumValues Enumerates the set of values for ListDashboardGroupsSortOrderEnum
func GetListDashboardGroupsSortOrderEnumValues() []ListDashboardGroupsSortOrderEnum {
	values := make([]ListDashboardGroupsSortOrderEnum, 0)
	for _, v := range mappingListDashboardGroupsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListDashboardGroupsSortOrderEnumStringValues Enumerates the set of values in String for ListDashboardGroupsSortOrderEnum
func GetListDashboardGroupsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListDashboardGroupsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListDashboardGroupsSortOrderEnum(val string) (ListDashboardGroupsSortOrderEnum, bool) {
	enum, ok := mappingListDashboardGroupsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListDashboardGroupsSortByEnum Enum with underlying type: string
type ListDashboardGroupsSortByEnum string

// Set of constants representing the allowable values for ListDashboardGroupsSortByEnum
const (
	ListDashboardGroupsSortByTimecreated ListDashboardGroupsSortByEnum = "timeCreated"
	ListDashboardGroupsSortByDisplayname ListDashboardGroupsSortByEnum = "displayName"
)

var mappingListDashboardGroupsSortByEnum = map[string]ListDashboardGroupsSortByEnum{
	"timeCreated": ListDashboardGroupsSortByTimecreated,
	"displayName": ListDashboardGroupsSortByDisplayname,
}

var mappingListDashboardGroupsSortByEnumLowerCase = map[string]ListDashboardGroupsSortByEnum{
	"timecreated": ListDashboardGroupsSortByTimecreated,
	"displayname": ListDashboardGroupsSortByDisplayname,
}

// GetListDashboardGroupsSortByEnumValues Enumerates the set of values for ListDashboardGroupsSortByEnum
func GetListDashboardGroupsSortByEnumValues() []ListDashboardGroupsSortByEnum {
	values := make([]ListDashboardGroupsSortByEnum, 0)
	for _, v := range mappingListDashboardGroupsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListDashboardGroupsSortByEnumStringValues Enumerates the set of values in String for ListDashboardGroupsSortByEnum
func GetListDashboardGroupsSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListDashboardGroupsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListDashboardGroupsSortByEnum(val string) (ListDashboardGroupsSortByEnum, bool) {
	enum, ok := mappingListDashboardGroupsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
