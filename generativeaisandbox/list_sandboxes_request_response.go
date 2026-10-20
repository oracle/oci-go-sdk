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

// ListSandboxesRequest wrapper for the ListSandboxes operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/generativeaisandbox/ListSandboxes.go.html to see an example of how to use ListSandboxesRequest.
type ListSandboxesRequest struct {

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	// The only valid characters for request IDs are letters, numbers,
	// underscore, and dash.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// The project identifier used to filter listed sandbox resources.
	ProjectId *string `mandatory:"false" contributesTo:"query" name:"projectId"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment in which to list resources.
	CompartmentId *string `mandatory:"false" contributesTo:"query" name:"compartmentId"`

	// Sandbox identifier used to filter listed sandbox resources. Newly generated Sandbox IDs use `sbx-` followed by 16 lowercase base32 characters (`a-z2-7`); existing legacy IDs remain accepted.
	SandboxId *string `mandatory:"false" contributesTo:"query" name:"sandboxId"`

	// Sandbox display name used to filter listed sandbox resources.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// A filter to return only resources whose state matches the given state.
	LifecycleState SandboxLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for `timeCreated` is ascending.
	SortBy ListSandboxesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to use, either ascending (`ASC`) or descending (`DESC`).
	SortOrder ListSandboxesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// A token representing the position at which to start retrieving results.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListSandboxesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListSandboxesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListSandboxesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListSandboxesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListSandboxesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingSandboxLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetSandboxLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListSandboxesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListSandboxesSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListSandboxesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListSandboxesSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListSandboxesResponse wrapper for the ListSandboxes operation
type ListSandboxesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of SandboxCollection instances
	SandboxCollection `presentIn:"body"`

	// For list pagination. When this header appears in the response, additional pages of results remain. For
	// important details about how pagination works, see List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response ListSandboxesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListSandboxesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListSandboxesSortByEnum Enum with underlying type: string
type ListSandboxesSortByEnum string

// Set of constants representing the allowable values for ListSandboxesSortByEnum
const (
	ListSandboxesSortByTimecreated ListSandboxesSortByEnum = "timeCreated"
	ListSandboxesSortByDisplayname ListSandboxesSortByEnum = "displayName"
)

var mappingListSandboxesSortByEnum = map[string]ListSandboxesSortByEnum{
	"timeCreated": ListSandboxesSortByTimecreated,
	"displayName": ListSandboxesSortByDisplayname,
}

var mappingListSandboxesSortByEnumLowerCase = map[string]ListSandboxesSortByEnum{
	"timecreated": ListSandboxesSortByTimecreated,
	"displayname": ListSandboxesSortByDisplayname,
}

// GetListSandboxesSortByEnumValues Enumerates the set of values for ListSandboxesSortByEnum
func GetListSandboxesSortByEnumValues() []ListSandboxesSortByEnum {
	values := make([]ListSandboxesSortByEnum, 0)
	for _, v := range mappingListSandboxesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListSandboxesSortByEnumStringValues Enumerates the set of values in String for ListSandboxesSortByEnum
func GetListSandboxesSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListSandboxesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListSandboxesSortByEnum(val string) (ListSandboxesSortByEnum, bool) {
	enum, ok := mappingListSandboxesSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListSandboxesSortOrderEnum Enum with underlying type: string
type ListSandboxesSortOrderEnum string

// Set of constants representing the allowable values for ListSandboxesSortOrderEnum
const (
	ListSandboxesSortOrderAsc  ListSandboxesSortOrderEnum = "ASC"
	ListSandboxesSortOrderDesc ListSandboxesSortOrderEnum = "DESC"
)

var mappingListSandboxesSortOrderEnum = map[string]ListSandboxesSortOrderEnum{
	"ASC":  ListSandboxesSortOrderAsc,
	"DESC": ListSandboxesSortOrderDesc,
}

var mappingListSandboxesSortOrderEnumLowerCase = map[string]ListSandboxesSortOrderEnum{
	"asc":  ListSandboxesSortOrderAsc,
	"desc": ListSandboxesSortOrderDesc,
}

// GetListSandboxesSortOrderEnumValues Enumerates the set of values for ListSandboxesSortOrderEnum
func GetListSandboxesSortOrderEnumValues() []ListSandboxesSortOrderEnum {
	values := make([]ListSandboxesSortOrderEnum, 0)
	for _, v := range mappingListSandboxesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListSandboxesSortOrderEnumStringValues Enumerates the set of values in String for ListSandboxesSortOrderEnum
func GetListSandboxesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListSandboxesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListSandboxesSortOrderEnum(val string) (ListSandboxesSortOrderEnum, bool) {
	enum, ok := mappingListSandboxesSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
