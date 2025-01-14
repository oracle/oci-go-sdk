// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package rover

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListRoverNodeRoverBundleRequestsRequest wrapper for the ListRoverNodeRoverBundleRequests operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/rover/ListRoverNodeRoverBundleRequests.go.html to see an example of how to use ListRoverNodeRoverBundleRequestsRequest.
type ListRoverNodeRoverBundleRequestsRequest struct {

	// Unique RoverNode identifier
	RoverNodeId *string `mandatory:"true" contributesTo:"path" name:"roverNodeId"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'asc' or 'desc'.
	SortOrder ListRoverNodeRoverBundleRequestsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for timeTaskCreated is descending. If no value is specified timeTaskCreated is default.
	SortBy ListRoverNodeRoverBundleRequestsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListRoverNodeRoverBundleRequestsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListRoverNodeRoverBundleRequestsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListRoverNodeRoverBundleRequestsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListRoverNodeRoverBundleRequestsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListRoverNodeRoverBundleRequestsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListRoverNodeRoverBundleRequestsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListRoverNodeRoverBundleRequestsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListRoverNodeRoverBundleRequestsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListRoverNodeRoverBundleRequestsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListRoverNodeRoverBundleRequestsResponse wrapper for the ListRoverNodeRoverBundleRequests operation
type ListRoverNodeRoverBundleRequestsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of RoverBundleRequestCollection instances
	RoverBundleRequestCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListRoverNodeRoverBundleRequestsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListRoverNodeRoverBundleRequestsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListRoverNodeRoverBundleRequestsSortOrderEnum Enum with underlying type: string
type ListRoverNodeRoverBundleRequestsSortOrderEnum string

// Set of constants representing the allowable values for ListRoverNodeRoverBundleRequestsSortOrderEnum
const (
	ListRoverNodeRoverBundleRequestsSortOrderAsc  ListRoverNodeRoverBundleRequestsSortOrderEnum = "ASC"
	ListRoverNodeRoverBundleRequestsSortOrderDesc ListRoverNodeRoverBundleRequestsSortOrderEnum = "DESC"
)

var mappingListRoverNodeRoverBundleRequestsSortOrderEnum = map[string]ListRoverNodeRoverBundleRequestsSortOrderEnum{
	"ASC":  ListRoverNodeRoverBundleRequestsSortOrderAsc,
	"DESC": ListRoverNodeRoverBundleRequestsSortOrderDesc,
}

var mappingListRoverNodeRoverBundleRequestsSortOrderEnumLowerCase = map[string]ListRoverNodeRoverBundleRequestsSortOrderEnum{
	"asc":  ListRoverNodeRoverBundleRequestsSortOrderAsc,
	"desc": ListRoverNodeRoverBundleRequestsSortOrderDesc,
}

// GetListRoverNodeRoverBundleRequestsSortOrderEnumValues Enumerates the set of values for ListRoverNodeRoverBundleRequestsSortOrderEnum
func GetListRoverNodeRoverBundleRequestsSortOrderEnumValues() []ListRoverNodeRoverBundleRequestsSortOrderEnum {
	values := make([]ListRoverNodeRoverBundleRequestsSortOrderEnum, 0)
	for _, v := range mappingListRoverNodeRoverBundleRequestsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListRoverNodeRoverBundleRequestsSortOrderEnumStringValues Enumerates the set of values in String for ListRoverNodeRoverBundleRequestsSortOrderEnum
func GetListRoverNodeRoverBundleRequestsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListRoverNodeRoverBundleRequestsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListRoverNodeRoverBundleRequestsSortOrderEnum(val string) (ListRoverNodeRoverBundleRequestsSortOrderEnum, bool) {
	enum, ok := mappingListRoverNodeRoverBundleRequestsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListRoverNodeRoverBundleRequestsSortByEnum Enum with underlying type: string
type ListRoverNodeRoverBundleRequestsSortByEnum string

// Set of constants representing the allowable values for ListRoverNodeRoverBundleRequestsSortByEnum
const (
	ListRoverNodeRoverBundleRequestsSortByTimetaskcreated ListRoverNodeRoverBundleRequestsSortByEnum = "timeTaskCreated"
)

var mappingListRoverNodeRoverBundleRequestsSortByEnum = map[string]ListRoverNodeRoverBundleRequestsSortByEnum{
	"timeTaskCreated": ListRoverNodeRoverBundleRequestsSortByTimetaskcreated,
}

var mappingListRoverNodeRoverBundleRequestsSortByEnumLowerCase = map[string]ListRoverNodeRoverBundleRequestsSortByEnum{
	"timetaskcreated": ListRoverNodeRoverBundleRequestsSortByTimetaskcreated,
}

// GetListRoverNodeRoverBundleRequestsSortByEnumValues Enumerates the set of values for ListRoverNodeRoverBundleRequestsSortByEnum
func GetListRoverNodeRoverBundleRequestsSortByEnumValues() []ListRoverNodeRoverBundleRequestsSortByEnum {
	values := make([]ListRoverNodeRoverBundleRequestsSortByEnum, 0)
	for _, v := range mappingListRoverNodeRoverBundleRequestsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListRoverNodeRoverBundleRequestsSortByEnumStringValues Enumerates the set of values in String for ListRoverNodeRoverBundleRequestsSortByEnum
func GetListRoverNodeRoverBundleRequestsSortByEnumStringValues() []string {
	return []string{
		"timeTaskCreated",
	}
}

// GetMappingListRoverNodeRoverBundleRequestsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListRoverNodeRoverBundleRequestsSortByEnum(val string) (ListRoverNodeRoverBundleRequestsSortByEnum, bool) {
	enum, ok := mappingListRoverNodeRoverBundleRequestsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
