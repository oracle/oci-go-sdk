// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package applicationmigration

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v62/common"
	"net/http"
	"strings"
)

// ListSourcesRequest wrapper for the ListSources operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/applicationmigration/ListSources.go.html to see an example of how to use ListSourcesRequest.
type ListSourcesRequest struct {

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of a compartment. Retrieves details of objects in the specified compartment.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a
	// particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) on which to query for a source.
	Id *string `mandatory:"false" contributesTo:"query" name:"id"`

	// The number of items returned in a paginated `List` call. For information about pagination, see
	// List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The value of the `opc-next-page` response header from the preceding `List` call.
	// For information about pagination, see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order, either `ASC` (ascending) or `DESC` (descending).
	SortOrder ListSourcesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// Specifies the field on which to sort.
	// By default, `TIMECREATED` is ordered descending.
	// By default, `DISPLAYNAME` is ordered ascending. Note that you can sort only on one field.
	SortBy ListSourcesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Display name on which to query.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// Retrieves details of sources in the specified lifecycle state.
	LifecycleState ListSourcesLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListSourcesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListSourcesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListSourcesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListSourcesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListSourcesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListSourcesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListSourcesSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListSourcesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListSourcesSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListSourcesLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetListSourcesLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListSourcesResponse wrapper for the ListSources operation
type ListSourcesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []SourceSummary instances
	Items []SourceSummary `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response, additional pages of results remain.
	// For details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	// Include this value as the `page` parameter for the subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListSourcesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListSourcesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListSourcesSortOrderEnum Enum with underlying type: string
type ListSourcesSortOrderEnum string

// Set of constants representing the allowable values for ListSourcesSortOrderEnum
const (
	ListSourcesSortOrderAsc  ListSourcesSortOrderEnum = "ASC"
	ListSourcesSortOrderDesc ListSourcesSortOrderEnum = "DESC"
)

var mappingListSourcesSortOrderEnum = map[string]ListSourcesSortOrderEnum{
	"ASC":  ListSourcesSortOrderAsc,
	"DESC": ListSourcesSortOrderDesc,
}

var mappingListSourcesSortOrderEnumLowerCase = map[string]ListSourcesSortOrderEnum{
	"asc":  ListSourcesSortOrderAsc,
	"desc": ListSourcesSortOrderDesc,
}

// GetListSourcesSortOrderEnumValues Enumerates the set of values for ListSourcesSortOrderEnum
func GetListSourcesSortOrderEnumValues() []ListSourcesSortOrderEnum {
	values := make([]ListSourcesSortOrderEnum, 0)
	for _, v := range mappingListSourcesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListSourcesSortOrderEnumStringValues Enumerates the set of values in String for ListSourcesSortOrderEnum
func GetListSourcesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListSourcesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListSourcesSortOrderEnum(val string) (ListSourcesSortOrderEnum, bool) {
	enum, ok := mappingListSourcesSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListSourcesSortByEnum Enum with underlying type: string
type ListSourcesSortByEnum string

// Set of constants representing the allowable values for ListSourcesSortByEnum
const (
	ListSourcesSortByTimecreated ListSourcesSortByEnum = "TIMECREATED"
	ListSourcesSortByDisplayname ListSourcesSortByEnum = "DISPLAYNAME"
)

var mappingListSourcesSortByEnum = map[string]ListSourcesSortByEnum{
	"TIMECREATED": ListSourcesSortByTimecreated,
	"DISPLAYNAME": ListSourcesSortByDisplayname,
}

var mappingListSourcesSortByEnumLowerCase = map[string]ListSourcesSortByEnum{
	"timecreated": ListSourcesSortByTimecreated,
	"displayname": ListSourcesSortByDisplayname,
}

// GetListSourcesSortByEnumValues Enumerates the set of values for ListSourcesSortByEnum
func GetListSourcesSortByEnumValues() []ListSourcesSortByEnum {
	values := make([]ListSourcesSortByEnum, 0)
	for _, v := range mappingListSourcesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListSourcesSortByEnumStringValues Enumerates the set of values in String for ListSourcesSortByEnum
func GetListSourcesSortByEnumStringValues() []string {
	return []string{
		"TIMECREATED",
		"DISPLAYNAME",
	}
}

// GetMappingListSourcesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListSourcesSortByEnum(val string) (ListSourcesSortByEnum, bool) {
	enum, ok := mappingListSourcesSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListSourcesLifecycleStateEnum Enum with underlying type: string
type ListSourcesLifecycleStateEnum string

// Set of constants representing the allowable values for ListSourcesLifecycleStateEnum
const (
	ListSourcesLifecycleStateCreating ListSourcesLifecycleStateEnum = "CREATING"
	ListSourcesLifecycleStateDeleting ListSourcesLifecycleStateEnum = "DELETING"
	ListSourcesLifecycleStateUpdating ListSourcesLifecycleStateEnum = "UPDATING"
	ListSourcesLifecycleStateActive   ListSourcesLifecycleStateEnum = "ACTIVE"
	ListSourcesLifecycleStateInactive ListSourcesLifecycleStateEnum = "INACTIVE"
	ListSourcesLifecycleStateDeleted  ListSourcesLifecycleStateEnum = "DELETED"
)

var mappingListSourcesLifecycleStateEnum = map[string]ListSourcesLifecycleStateEnum{
	"CREATING": ListSourcesLifecycleStateCreating,
	"DELETING": ListSourcesLifecycleStateDeleting,
	"UPDATING": ListSourcesLifecycleStateUpdating,
	"ACTIVE":   ListSourcesLifecycleStateActive,
	"INACTIVE": ListSourcesLifecycleStateInactive,
	"DELETED":  ListSourcesLifecycleStateDeleted,
}

var mappingListSourcesLifecycleStateEnumLowerCase = map[string]ListSourcesLifecycleStateEnum{
	"creating": ListSourcesLifecycleStateCreating,
	"deleting": ListSourcesLifecycleStateDeleting,
	"updating": ListSourcesLifecycleStateUpdating,
	"active":   ListSourcesLifecycleStateActive,
	"inactive": ListSourcesLifecycleStateInactive,
	"deleted":  ListSourcesLifecycleStateDeleted,
}

// GetListSourcesLifecycleStateEnumValues Enumerates the set of values for ListSourcesLifecycleStateEnum
func GetListSourcesLifecycleStateEnumValues() []ListSourcesLifecycleStateEnum {
	values := make([]ListSourcesLifecycleStateEnum, 0)
	for _, v := range mappingListSourcesLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetListSourcesLifecycleStateEnumStringValues Enumerates the set of values in String for ListSourcesLifecycleStateEnum
func GetListSourcesLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"DELETING",
		"UPDATING",
		"ACTIVE",
		"INACTIVE",
		"DELETED",
	}
}

// GetMappingListSourcesLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListSourcesLifecycleStateEnum(val string) (ListSourcesLifecycleStateEnum, bool) {
	enum, ok := mappingListSourcesLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
