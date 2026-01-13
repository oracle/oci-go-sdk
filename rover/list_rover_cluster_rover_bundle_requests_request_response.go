// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package rover

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListRoverClusterRoverBundleRequestsRequest wrapper for the ListRoverClusterRoverBundleRequests operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/rover/ListRoverClusterRoverBundleRequests.go.html to see an example of how to use ListRoverClusterRoverBundleRequestsRequest.
type ListRoverClusterRoverBundleRequestsRequest struct {

	// Unique RoverCluster identifier
	RoverClusterId *string `mandatory:"true" contributesTo:"path" name:"roverClusterId"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'asc' or 'desc'.
	SortOrder ListRoverClusterRoverBundleRequestsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for timeTaskCreated is descending. If no value is specified timeTaskCreated is default.
	SortBy ListRoverClusterRoverBundleRequestsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListRoverClusterRoverBundleRequestsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListRoverClusterRoverBundleRequestsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListRoverClusterRoverBundleRequestsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListRoverClusterRoverBundleRequestsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListRoverClusterRoverBundleRequestsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListRoverClusterRoverBundleRequestsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListRoverClusterRoverBundleRequestsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListRoverClusterRoverBundleRequestsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListRoverClusterRoverBundleRequestsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListRoverClusterRoverBundleRequestsResponse wrapper for the ListRoverClusterRoverBundleRequests operation
type ListRoverClusterRoverBundleRequestsResponse struct {

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

func (response ListRoverClusterRoverBundleRequestsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListRoverClusterRoverBundleRequestsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListRoverClusterRoverBundleRequestsSortOrderEnum Enum with underlying type: string
type ListRoverClusterRoverBundleRequestsSortOrderEnum string

// Set of constants representing the allowable values for ListRoverClusterRoverBundleRequestsSortOrderEnum
const (
	ListRoverClusterRoverBundleRequestsSortOrderAsc  ListRoverClusterRoverBundleRequestsSortOrderEnum = "ASC"
	ListRoverClusterRoverBundleRequestsSortOrderDesc ListRoverClusterRoverBundleRequestsSortOrderEnum = "DESC"
)

var mappingListRoverClusterRoverBundleRequestsSortOrderEnum = map[string]ListRoverClusterRoverBundleRequestsSortOrderEnum{
	"ASC":  ListRoverClusterRoverBundleRequestsSortOrderAsc,
	"DESC": ListRoverClusterRoverBundleRequestsSortOrderDesc,
}

var mappingListRoverClusterRoverBundleRequestsSortOrderEnumLowerCase = map[string]ListRoverClusterRoverBundleRequestsSortOrderEnum{
	"asc":  ListRoverClusterRoverBundleRequestsSortOrderAsc,
	"desc": ListRoverClusterRoverBundleRequestsSortOrderDesc,
}

// GetListRoverClusterRoverBundleRequestsSortOrderEnumValues Enumerates the set of values for ListRoverClusterRoverBundleRequestsSortOrderEnum
func GetListRoverClusterRoverBundleRequestsSortOrderEnumValues() []ListRoverClusterRoverBundleRequestsSortOrderEnum {
	values := make([]ListRoverClusterRoverBundleRequestsSortOrderEnum, 0)
	for _, v := range mappingListRoverClusterRoverBundleRequestsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListRoverClusterRoverBundleRequestsSortOrderEnumStringValues Enumerates the set of values in String for ListRoverClusterRoverBundleRequestsSortOrderEnum
func GetListRoverClusterRoverBundleRequestsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListRoverClusterRoverBundleRequestsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListRoverClusterRoverBundleRequestsSortOrderEnum(val string) (ListRoverClusterRoverBundleRequestsSortOrderEnum, bool) {
	enum, ok := mappingListRoverClusterRoverBundleRequestsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListRoverClusterRoverBundleRequestsSortByEnum Enum with underlying type: string
type ListRoverClusterRoverBundleRequestsSortByEnum string

// Set of constants representing the allowable values for ListRoverClusterRoverBundleRequestsSortByEnum
const (
	ListRoverClusterRoverBundleRequestsSortByTimetaskcreated ListRoverClusterRoverBundleRequestsSortByEnum = "timeTaskCreated"
)

var mappingListRoverClusterRoverBundleRequestsSortByEnum = map[string]ListRoverClusterRoverBundleRequestsSortByEnum{
	"timeTaskCreated": ListRoverClusterRoverBundleRequestsSortByTimetaskcreated,
}

var mappingListRoverClusterRoverBundleRequestsSortByEnumLowerCase = map[string]ListRoverClusterRoverBundleRequestsSortByEnum{
	"timetaskcreated": ListRoverClusterRoverBundleRequestsSortByTimetaskcreated,
}

// GetListRoverClusterRoverBundleRequestsSortByEnumValues Enumerates the set of values for ListRoverClusterRoverBundleRequestsSortByEnum
func GetListRoverClusterRoverBundleRequestsSortByEnumValues() []ListRoverClusterRoverBundleRequestsSortByEnum {
	values := make([]ListRoverClusterRoverBundleRequestsSortByEnum, 0)
	for _, v := range mappingListRoverClusterRoverBundleRequestsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListRoverClusterRoverBundleRequestsSortByEnumStringValues Enumerates the set of values in String for ListRoverClusterRoverBundleRequestsSortByEnum
func GetListRoverClusterRoverBundleRequestsSortByEnumStringValues() []string {
	return []string{
		"timeTaskCreated",
	}
}

// GetMappingListRoverClusterRoverBundleRequestsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListRoverClusterRoverBundleRequestsSortByEnum(val string) (ListRoverClusterRoverBundleRequestsSortByEnum, bool) {
	enum, ok := mappingListRoverClusterRoverBundleRequestsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
