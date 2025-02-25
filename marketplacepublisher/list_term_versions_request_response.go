// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package marketplacepublisher

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListTermVersionsRequest wrapper for the ListTermVersions operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/marketplacepublisher/ListTermVersions.go.html to see an example of how to use ListTermVersionsRequest.
type ListTermVersionsRequest struct {

	// term OCID
	TermId *string `mandatory:"true" contributesTo:"query" name:"termId"`

	// The unique identifier for the compartment.
	CompartmentId *string `mandatory:"false" contributesTo:"query" name:"compartmentId"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// A token representing the position at which to start retrieving results. This must come from the `opc-next-page` header field of a previous response.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'ASC' or 'DESC'.
	SortOrder ListTermVersionsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for timeCreated is descending. Default order for displayName is ascending.
	SortBy ListTermVersionsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// A filter to return only resources that match the entire display name given.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListTermVersionsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListTermVersionsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListTermVersionsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListTermVersionsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListTermVersionsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListTermVersionsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListTermVersionsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListTermVersionsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListTermVersionsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListTermVersionsResponse wrapper for the ListTermVersions operation
type ListTermVersionsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of TermVersionCollection instances
	TermVersionCollection `presentIn:"body"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response ListTermVersionsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListTermVersionsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListTermVersionsSortOrderEnum Enum with underlying type: string
type ListTermVersionsSortOrderEnum string

// Set of constants representing the allowable values for ListTermVersionsSortOrderEnum
const (
	ListTermVersionsSortOrderAsc  ListTermVersionsSortOrderEnum = "ASC"
	ListTermVersionsSortOrderDesc ListTermVersionsSortOrderEnum = "DESC"
)

var mappingListTermVersionsSortOrderEnum = map[string]ListTermVersionsSortOrderEnum{
	"ASC":  ListTermVersionsSortOrderAsc,
	"DESC": ListTermVersionsSortOrderDesc,
}

var mappingListTermVersionsSortOrderEnumLowerCase = map[string]ListTermVersionsSortOrderEnum{
	"asc":  ListTermVersionsSortOrderAsc,
	"desc": ListTermVersionsSortOrderDesc,
}

// GetListTermVersionsSortOrderEnumValues Enumerates the set of values for ListTermVersionsSortOrderEnum
func GetListTermVersionsSortOrderEnumValues() []ListTermVersionsSortOrderEnum {
	values := make([]ListTermVersionsSortOrderEnum, 0)
	for _, v := range mappingListTermVersionsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListTermVersionsSortOrderEnumStringValues Enumerates the set of values in String for ListTermVersionsSortOrderEnum
func GetListTermVersionsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListTermVersionsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListTermVersionsSortOrderEnum(val string) (ListTermVersionsSortOrderEnum, bool) {
	enum, ok := mappingListTermVersionsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListTermVersionsSortByEnum Enum with underlying type: string
type ListTermVersionsSortByEnum string

// Set of constants representing the allowable values for ListTermVersionsSortByEnum
const (
	ListTermVersionsSortByTimecreated ListTermVersionsSortByEnum = "timeCreated"
	ListTermVersionsSortByDisplayname ListTermVersionsSortByEnum = "displayName"
)

var mappingListTermVersionsSortByEnum = map[string]ListTermVersionsSortByEnum{
	"timeCreated": ListTermVersionsSortByTimecreated,
	"displayName": ListTermVersionsSortByDisplayname,
}

var mappingListTermVersionsSortByEnumLowerCase = map[string]ListTermVersionsSortByEnum{
	"timecreated": ListTermVersionsSortByTimecreated,
	"displayname": ListTermVersionsSortByDisplayname,
}

// GetListTermVersionsSortByEnumValues Enumerates the set of values for ListTermVersionsSortByEnum
func GetListTermVersionsSortByEnumValues() []ListTermVersionsSortByEnum {
	values := make([]ListTermVersionsSortByEnum, 0)
	for _, v := range mappingListTermVersionsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListTermVersionsSortByEnumStringValues Enumerates the set of values in String for ListTermVersionsSortByEnum
func GetListTermVersionsSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListTermVersionsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListTermVersionsSortByEnum(val string) (ListTermVersionsSortByEnum, bool) {
	enum, ok := mappingListTermVersionsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
