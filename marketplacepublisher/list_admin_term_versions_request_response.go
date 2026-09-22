// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package marketplacepublisher

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListAdminTermVersionsRequest wrapper for the ListAdminTermVersions operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/marketplacepublisher/ListAdminTermVersions.go.html to see an example of how to use ListAdminTermVersionsRequest.
type ListAdminTermVersionsRequest struct {

	// term OCID
	TermId *string `mandatory:"true" contributesTo:"query" name:"termId"`

	// The unique identifier for the compartment.
	CompartmentId *string `mandatory:"false" contributesTo:"query" name:"compartmentId"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// A token representing the position at which to start retrieving results. This must come from the `opc-next-page` header field of a previous response.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'ASC' or 'DESC'.
	SortOrder ListAdminTermVersionsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for timeCreated is descending. Default order for displayName is ascending.
	SortBy ListAdminTermVersionsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// A filter to return only resources that match the entire display name given.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListAdminTermVersionsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListAdminTermVersionsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListAdminTermVersionsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListAdminTermVersionsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListAdminTermVersionsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListAdminTermVersionsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListAdminTermVersionsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListAdminTermVersionsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListAdminTermVersionsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListAdminTermVersionsResponse wrapper for the ListAdminTermVersions operation
type ListAdminTermVersionsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of AdminTermVersionCollection instances
	AdminTermVersionCollection `presentIn:"body"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response ListAdminTermVersionsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListAdminTermVersionsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListAdminTermVersionsSortOrderEnum Enum with underlying type: string
type ListAdminTermVersionsSortOrderEnum string

// Set of constants representing the allowable values for ListAdminTermVersionsSortOrderEnum
const (
	ListAdminTermVersionsSortOrderAsc  ListAdminTermVersionsSortOrderEnum = "ASC"
	ListAdminTermVersionsSortOrderDesc ListAdminTermVersionsSortOrderEnum = "DESC"
)

var mappingListAdminTermVersionsSortOrderEnum = map[string]ListAdminTermVersionsSortOrderEnum{
	"ASC":  ListAdminTermVersionsSortOrderAsc,
	"DESC": ListAdminTermVersionsSortOrderDesc,
}

var mappingListAdminTermVersionsSortOrderEnumLowerCase = map[string]ListAdminTermVersionsSortOrderEnum{
	"asc":  ListAdminTermVersionsSortOrderAsc,
	"desc": ListAdminTermVersionsSortOrderDesc,
}

// GetListAdminTermVersionsSortOrderEnumValues Enumerates the set of values for ListAdminTermVersionsSortOrderEnum
func GetListAdminTermVersionsSortOrderEnumValues() []ListAdminTermVersionsSortOrderEnum {
	values := make([]ListAdminTermVersionsSortOrderEnum, 0)
	for _, v := range mappingListAdminTermVersionsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListAdminTermVersionsSortOrderEnumStringValues Enumerates the set of values in String for ListAdminTermVersionsSortOrderEnum
func GetListAdminTermVersionsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListAdminTermVersionsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListAdminTermVersionsSortOrderEnum(val string) (ListAdminTermVersionsSortOrderEnum, bool) {
	enum, ok := mappingListAdminTermVersionsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListAdminTermVersionsSortByEnum Enum with underlying type: string
type ListAdminTermVersionsSortByEnum string

// Set of constants representing the allowable values for ListAdminTermVersionsSortByEnum
const (
	ListAdminTermVersionsSortByTimecreated ListAdminTermVersionsSortByEnum = "timeCreated"
	ListAdminTermVersionsSortByDisplayname ListAdminTermVersionsSortByEnum = "displayName"
)

var mappingListAdminTermVersionsSortByEnum = map[string]ListAdminTermVersionsSortByEnum{
	"timeCreated": ListAdminTermVersionsSortByTimecreated,
	"displayName": ListAdminTermVersionsSortByDisplayname,
}

var mappingListAdminTermVersionsSortByEnumLowerCase = map[string]ListAdminTermVersionsSortByEnum{
	"timecreated": ListAdminTermVersionsSortByTimecreated,
	"displayname": ListAdminTermVersionsSortByDisplayname,
}

// GetListAdminTermVersionsSortByEnumValues Enumerates the set of values for ListAdminTermVersionsSortByEnum
func GetListAdminTermVersionsSortByEnumValues() []ListAdminTermVersionsSortByEnum {
	values := make([]ListAdminTermVersionsSortByEnum, 0)
	for _, v := range mappingListAdminTermVersionsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListAdminTermVersionsSortByEnumStringValues Enumerates the set of values in String for ListAdminTermVersionsSortByEnum
func GetListAdminTermVersionsSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListAdminTermVersionsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListAdminTermVersionsSortByEnum(val string) (ListAdminTermVersionsSortByEnum, bool) {
	enum, ok := mappingListAdminTermVersionsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
