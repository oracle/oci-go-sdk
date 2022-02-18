// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package applicationmigration

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v59/common"
	"net/http"
	"strings"
)

// ListMigrationsRequest wrapper for the ListMigrations operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/applicationmigration/ListMigrations.go.html to see an example of how to use ListMigrationsRequest.
type ListMigrationsRequest struct {

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of a compartment. Retrieves details of objects in the specified compartment.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a
	// particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) on which to query for a migration.
	Id *string `mandatory:"false" contributesTo:"query" name:"id"`

	// The number of items returned in a paginated `List` call. For information about pagination, see
	// List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The value of the `opc-next-page` response header from the preceding `List` call.
	// For information about pagination, see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order, either `ASC` (ascending) or `DESC` (descending).
	SortOrder ListMigrationsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// Specifies the field on which to sort.
	// By default, `TIMECREATED` is ordered descending.
	// By default, `DISPLAYNAME` is ordered ascending. Note that you can sort only on one field.
	SortBy ListMigrationsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Display name on which to query.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// This field is not supported. Do not use.
	LifecycleState ListMigrationsLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListMigrationsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListMigrationsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListMigrationsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListMigrationsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListMigrationsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListMigrationsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListMigrationsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListMigrationsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListMigrationsSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListMigrationsLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetListMigrationsLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListMigrationsResponse wrapper for the ListMigrations operation
type ListMigrationsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []MigrationSummary instances
	Items []MigrationSummary `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response, additional pages of results remain.
	// For details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	// Include this value as the `page` parameter for the subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListMigrationsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListMigrationsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListMigrationsSortOrderEnum Enum with underlying type: string
type ListMigrationsSortOrderEnum string

// Set of constants representing the allowable values for ListMigrationsSortOrderEnum
const (
	ListMigrationsSortOrderAsc  ListMigrationsSortOrderEnum = "ASC"
	ListMigrationsSortOrderDesc ListMigrationsSortOrderEnum = "DESC"
)

var mappingListMigrationsSortOrderEnum = map[string]ListMigrationsSortOrderEnum{
	"ASC":  ListMigrationsSortOrderAsc,
	"DESC": ListMigrationsSortOrderDesc,
}

var mappingListMigrationsSortOrderEnumLowerCase = map[string]ListMigrationsSortOrderEnum{
	"asc":  ListMigrationsSortOrderAsc,
	"desc": ListMigrationsSortOrderDesc,
}

// GetListMigrationsSortOrderEnumValues Enumerates the set of values for ListMigrationsSortOrderEnum
func GetListMigrationsSortOrderEnumValues() []ListMigrationsSortOrderEnum {
	values := make([]ListMigrationsSortOrderEnum, 0)
	for _, v := range mappingListMigrationsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListMigrationsSortOrderEnumStringValues Enumerates the set of values in String for ListMigrationsSortOrderEnum
func GetListMigrationsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListMigrationsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListMigrationsSortOrderEnum(val string) (ListMigrationsSortOrderEnum, bool) {
	enum, ok := mappingListMigrationsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListMigrationsSortByEnum Enum with underlying type: string
type ListMigrationsSortByEnum string

// Set of constants representing the allowable values for ListMigrationsSortByEnum
const (
	ListMigrationsSortByTimecreated ListMigrationsSortByEnum = "TIMECREATED"
	ListMigrationsSortByDisplayname ListMigrationsSortByEnum = "DISPLAYNAME"
)

var mappingListMigrationsSortByEnum = map[string]ListMigrationsSortByEnum{
	"TIMECREATED": ListMigrationsSortByTimecreated,
	"DISPLAYNAME": ListMigrationsSortByDisplayname,
}

var mappingListMigrationsSortByEnumLowerCase = map[string]ListMigrationsSortByEnum{
	"timecreated": ListMigrationsSortByTimecreated,
	"displayname": ListMigrationsSortByDisplayname,
}

// GetListMigrationsSortByEnumValues Enumerates the set of values for ListMigrationsSortByEnum
func GetListMigrationsSortByEnumValues() []ListMigrationsSortByEnum {
	values := make([]ListMigrationsSortByEnum, 0)
	for _, v := range mappingListMigrationsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListMigrationsSortByEnumStringValues Enumerates the set of values in String for ListMigrationsSortByEnum
func GetListMigrationsSortByEnumStringValues() []string {
	return []string{
		"TIMECREATED",
		"DISPLAYNAME",
	}
}

// GetMappingListMigrationsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListMigrationsSortByEnum(val string) (ListMigrationsSortByEnum, bool) {
	enum, ok := mappingListMigrationsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListMigrationsLifecycleStateEnum Enum with underlying type: string
type ListMigrationsLifecycleStateEnum string

// Set of constants representing the allowable values for ListMigrationsLifecycleStateEnum
const (
	ListMigrationsLifecycleStateCreating  ListMigrationsLifecycleStateEnum = "CREATING"
	ListMigrationsLifecycleStateActive    ListMigrationsLifecycleStateEnum = "ACTIVE"
	ListMigrationsLifecycleStateInactive  ListMigrationsLifecycleStateEnum = "INACTIVE"
	ListMigrationsLifecycleStateUpdating  ListMigrationsLifecycleStateEnum = "UPDATING"
	ListMigrationsLifecycleStateSucceeded ListMigrationsLifecycleStateEnum = "SUCCEEDED"
	ListMigrationsLifecycleStateDeleting  ListMigrationsLifecycleStateEnum = "DELETING"
	ListMigrationsLifecycleStateDeleted   ListMigrationsLifecycleStateEnum = "DELETED"
)

var mappingListMigrationsLifecycleStateEnum = map[string]ListMigrationsLifecycleStateEnum{
	"CREATING":  ListMigrationsLifecycleStateCreating,
	"ACTIVE":    ListMigrationsLifecycleStateActive,
	"INACTIVE":  ListMigrationsLifecycleStateInactive,
	"UPDATING":  ListMigrationsLifecycleStateUpdating,
	"SUCCEEDED": ListMigrationsLifecycleStateSucceeded,
	"DELETING":  ListMigrationsLifecycleStateDeleting,
	"DELETED":   ListMigrationsLifecycleStateDeleted,
}

var mappingListMigrationsLifecycleStateEnumLowerCase = map[string]ListMigrationsLifecycleStateEnum{
	"creating":  ListMigrationsLifecycleStateCreating,
	"active":    ListMigrationsLifecycleStateActive,
	"inactive":  ListMigrationsLifecycleStateInactive,
	"updating":  ListMigrationsLifecycleStateUpdating,
	"succeeded": ListMigrationsLifecycleStateSucceeded,
	"deleting":  ListMigrationsLifecycleStateDeleting,
	"deleted":   ListMigrationsLifecycleStateDeleted,
}

// GetListMigrationsLifecycleStateEnumValues Enumerates the set of values for ListMigrationsLifecycleStateEnum
func GetListMigrationsLifecycleStateEnumValues() []ListMigrationsLifecycleStateEnum {
	values := make([]ListMigrationsLifecycleStateEnum, 0)
	for _, v := range mappingListMigrationsLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetListMigrationsLifecycleStateEnumStringValues Enumerates the set of values in String for ListMigrationsLifecycleStateEnum
func GetListMigrationsLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"ACTIVE",
		"INACTIVE",
		"UPDATING",
		"SUCCEEDED",
		"DELETING",
		"DELETED",
	}
}

// GetMappingListMigrationsLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListMigrationsLifecycleStateEnum(val string) (ListMigrationsLifecycleStateEnum, bool) {
	enum, ok := mappingListMigrationsLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
