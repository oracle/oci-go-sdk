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

// ListSandboxCommandRequest wrapper for the ListSandboxCommand operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/generativeaisandbox/ListSandboxCommand.go.html to see an example of how to use ListSandboxCommandRequest.
type ListSandboxCommandRequest struct {

	// Sandbox identifier. Newly generated Sandbox IDs use `sbx-` followed by 16 lowercase base32 characters (`a-z2-7`); existing legacy IDs remain accepted.
	SandboxId *string `mandatory:"true" contributesTo:"path" name:"sandboxId"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	// The only valid characters for request IDs are letters, numbers,
	// underscore, and dash.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Sandbox command identifier used to filter listed command resources.
	CommandId *string `mandatory:"false" contributesTo:"query" name:"commandId"`

	// The field to sort by. Only one sort order may be provided. Default order for `timeCreated` is ascending.
	SortBy ListSandboxCommandSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to use, either ascending (`ASC`) or descending (`DESC`).
	SortOrder ListSandboxCommandSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// A token representing the position at which to start retrieving results.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListSandboxCommandRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListSandboxCommandRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListSandboxCommandRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListSandboxCommandRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListSandboxCommandRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListSandboxCommandSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListSandboxCommandSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListSandboxCommandSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListSandboxCommandSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListSandboxCommandResponse wrapper for the ListSandboxCommand operation
type ListSandboxCommandResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of SandboxCommandCollection instances
	SandboxCommandCollection `presentIn:"body"`

	// For list pagination. When this header appears in the response, additional pages of results remain. For
	// important details about how pagination works, see List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response ListSandboxCommandResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListSandboxCommandResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListSandboxCommandSortByEnum Enum with underlying type: string
type ListSandboxCommandSortByEnum string

// Set of constants representing the allowable values for ListSandboxCommandSortByEnum
const (
	ListSandboxCommandSortByTimecreated  ListSandboxCommandSortByEnum = "timeCreated"
	ListSandboxCommandSortByTimestarted  ListSandboxCommandSortByEnum = "timeStarted"
	ListSandboxCommandSortByTimefinished ListSandboxCommandSortByEnum = "timeFinished"
)

var mappingListSandboxCommandSortByEnum = map[string]ListSandboxCommandSortByEnum{
	"timeCreated":  ListSandboxCommandSortByTimecreated,
	"timeStarted":  ListSandboxCommandSortByTimestarted,
	"timeFinished": ListSandboxCommandSortByTimefinished,
}

var mappingListSandboxCommandSortByEnumLowerCase = map[string]ListSandboxCommandSortByEnum{
	"timecreated":  ListSandboxCommandSortByTimecreated,
	"timestarted":  ListSandboxCommandSortByTimestarted,
	"timefinished": ListSandboxCommandSortByTimefinished,
}

// GetListSandboxCommandSortByEnumValues Enumerates the set of values for ListSandboxCommandSortByEnum
func GetListSandboxCommandSortByEnumValues() []ListSandboxCommandSortByEnum {
	values := make([]ListSandboxCommandSortByEnum, 0)
	for _, v := range mappingListSandboxCommandSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListSandboxCommandSortByEnumStringValues Enumerates the set of values in String for ListSandboxCommandSortByEnum
func GetListSandboxCommandSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"timeStarted",
		"timeFinished",
	}
}

// GetMappingListSandboxCommandSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListSandboxCommandSortByEnum(val string) (ListSandboxCommandSortByEnum, bool) {
	enum, ok := mappingListSandboxCommandSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListSandboxCommandSortOrderEnum Enum with underlying type: string
type ListSandboxCommandSortOrderEnum string

// Set of constants representing the allowable values for ListSandboxCommandSortOrderEnum
const (
	ListSandboxCommandSortOrderAsc  ListSandboxCommandSortOrderEnum = "ASC"
	ListSandboxCommandSortOrderDesc ListSandboxCommandSortOrderEnum = "DESC"
)

var mappingListSandboxCommandSortOrderEnum = map[string]ListSandboxCommandSortOrderEnum{
	"ASC":  ListSandboxCommandSortOrderAsc,
	"DESC": ListSandboxCommandSortOrderDesc,
}

var mappingListSandboxCommandSortOrderEnumLowerCase = map[string]ListSandboxCommandSortOrderEnum{
	"asc":  ListSandboxCommandSortOrderAsc,
	"desc": ListSandboxCommandSortOrderDesc,
}

// GetListSandboxCommandSortOrderEnumValues Enumerates the set of values for ListSandboxCommandSortOrderEnum
func GetListSandboxCommandSortOrderEnumValues() []ListSandboxCommandSortOrderEnum {
	values := make([]ListSandboxCommandSortOrderEnum, 0)
	for _, v := range mappingListSandboxCommandSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListSandboxCommandSortOrderEnumStringValues Enumerates the set of values in String for ListSandboxCommandSortOrderEnum
func GetListSandboxCommandSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListSandboxCommandSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListSandboxCommandSortOrderEnum(val string) (ListSandboxCommandSortOrderEnum, bool) {
	enum, ok := mappingListSandboxCommandSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
