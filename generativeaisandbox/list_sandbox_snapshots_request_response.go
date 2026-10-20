// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package generativeaisandbox

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListSandboxSnapshotsRequest wrapper for the ListSandboxSnapshots operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/generativeaisandbox/ListSandboxSnapshots.go.html to see an example of how to use ListSandboxSnapshotsRequest.
type ListSandboxSnapshotsRequest struct {

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	// The only valid characters for request IDs are letters, numbers,
	// underscore, and dash.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// The project identifier used to filter listed sandbox snapshot resources.
	ProjectId *string `mandatory:"false" contributesTo:"query" name:"projectId"`

	// Sandbox snapshot display name used to filter listed sandbox snapshot resources.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// Sandbox identifier used to filter listed sandbox snapshot resources. Newly generated Sandbox IDs use `sbx-` followed by 16 lowercase base32 characters (`a-z2-7`); existing legacy IDs remain accepted.
	SandboxId *string `mandatory:"false" contributesTo:"query" name:"sandboxId"`

	// Sandbox snapshot identifier used to filter listed sandbox snapshot resources.
	SnapshotId *string `mandatory:"false" contributesTo:"query" name:"snapshotId"`

	// A filter to return only sandbox snapshot resources whose state matches the given state.
	LifecycleState SandboxSnapshotLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for `timeCreated` is ascending.
	SortBy ListSandboxSnapshotsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to use, either ascending (`ASC`) or descending (`DESC`).
	SortOrder ListSandboxSnapshotsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// A token representing the position at which to start retrieving results.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListSandboxSnapshotsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListSandboxSnapshotsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListSandboxSnapshotsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListSandboxSnapshotsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListSandboxSnapshotsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingSandboxSnapshotLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetSandboxSnapshotLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListSandboxSnapshotsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListSandboxSnapshotsSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListSandboxSnapshotsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListSandboxSnapshotsSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListSandboxSnapshotsResponse wrapper for the ListSandboxSnapshots operation
type ListSandboxSnapshotsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of SandboxSnapshotCollection instances
	SandboxSnapshotCollection `presentIn:"body"`

	// For list pagination. When this header appears in the response, additional pages of results remain. For
	// important details about how pagination works, see List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response ListSandboxSnapshotsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListSandboxSnapshotsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListSandboxSnapshotsSortByEnum Enum with underlying type: string
type ListSandboxSnapshotsSortByEnum string

// Set of constants representing the allowable values for ListSandboxSnapshotsSortByEnum
const (
	ListSandboxSnapshotsSortByTimecreated ListSandboxSnapshotsSortByEnum = "timeCreated"
	ListSandboxSnapshotsSortByDisplayname ListSandboxSnapshotsSortByEnum = "displayName"
)

var mappingListSandboxSnapshotsSortByEnum = map[string]ListSandboxSnapshotsSortByEnum{
	"timeCreated": ListSandboxSnapshotsSortByTimecreated,
	"displayName": ListSandboxSnapshotsSortByDisplayname,
}

var mappingListSandboxSnapshotsSortByEnumLowerCase = map[string]ListSandboxSnapshotsSortByEnum{
	"timecreated": ListSandboxSnapshotsSortByTimecreated,
	"displayname": ListSandboxSnapshotsSortByDisplayname,
}

// GetListSandboxSnapshotsSortByEnumValues Enumerates the set of values for ListSandboxSnapshotsSortByEnum
func GetListSandboxSnapshotsSortByEnumValues() []ListSandboxSnapshotsSortByEnum {
	values := make([]ListSandboxSnapshotsSortByEnum, 0)
	for _, v := range mappingListSandboxSnapshotsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListSandboxSnapshotsSortByEnumStringValues Enumerates the set of values in String for ListSandboxSnapshotsSortByEnum
func GetListSandboxSnapshotsSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListSandboxSnapshotsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListSandboxSnapshotsSortByEnum(val string) (ListSandboxSnapshotsSortByEnum, bool) {
	enum, ok := mappingListSandboxSnapshotsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListSandboxSnapshotsSortOrderEnum Enum with underlying type: string
type ListSandboxSnapshotsSortOrderEnum string

// Set of constants representing the allowable values for ListSandboxSnapshotsSortOrderEnum
const (
	ListSandboxSnapshotsSortOrderAsc  ListSandboxSnapshotsSortOrderEnum = "ASC"
	ListSandboxSnapshotsSortOrderDesc ListSandboxSnapshotsSortOrderEnum = "DESC"
)

var mappingListSandboxSnapshotsSortOrderEnum = map[string]ListSandboxSnapshotsSortOrderEnum{
	"ASC":  ListSandboxSnapshotsSortOrderAsc,
	"DESC": ListSandboxSnapshotsSortOrderDesc,
}

var mappingListSandboxSnapshotsSortOrderEnumLowerCase = map[string]ListSandboxSnapshotsSortOrderEnum{
	"asc":  ListSandboxSnapshotsSortOrderAsc,
	"desc": ListSandboxSnapshotsSortOrderDesc,
}

// GetListSandboxSnapshotsSortOrderEnumValues Enumerates the set of values for ListSandboxSnapshotsSortOrderEnum
func GetListSandboxSnapshotsSortOrderEnumValues() []ListSandboxSnapshotsSortOrderEnum {
	values := make([]ListSandboxSnapshotsSortOrderEnum, 0)
	for _, v := range mappingListSandboxSnapshotsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListSandboxSnapshotsSortOrderEnumStringValues Enumerates the set of values in String for ListSandboxSnapshotsSortOrderEnum
func GetListSandboxSnapshotsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListSandboxSnapshotsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListSandboxSnapshotsSortOrderEnum(val string) (ListSandboxSnapshotsSortOrderEnum, bool) {
	enum, ok := mappingListSandboxSnapshotsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
