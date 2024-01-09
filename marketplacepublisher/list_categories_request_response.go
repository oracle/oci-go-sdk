// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package marketplacepublisher

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListCategoriesRequest wrapper for the ListCategories operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/marketplacepublisher/ListCategories.go.html to see an example of how to use ListCategoriesRequest.
type ListCategoriesRequest struct {

	// The ID of the compartment in which to list resources.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// A token representing the position at which to start retrieving results. This must come from the `opc-next-page` header field of a previous response.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// product code to filter
	ProductCode *string `mandatory:"false" contributesTo:"query" name:"productCode"`

	// The sort order to use, either 'ASC' or 'DESC'.
	SortOrder ListCategoriesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for timeCreated is descending. Default order for displayName is ascending.
	SortBy ListCategoriesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// A filter to return only resources that match the entire name given.
	Name *string `mandatory:"false" contributesTo:"query" name:"name"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListCategoriesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListCategoriesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListCategoriesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListCategoriesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListCategoriesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListCategoriesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListCategoriesSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListCategoriesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListCategoriesSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListCategoriesResponse wrapper for the ListCategories operation
type ListCategoriesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of CategoryCollection instances
	CategoryCollection `presentIn:"body"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response ListCategoriesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListCategoriesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListCategoriesSortOrderEnum Enum with underlying type: string
type ListCategoriesSortOrderEnum string

// Set of constants representing the allowable values for ListCategoriesSortOrderEnum
const (
	ListCategoriesSortOrderAsc  ListCategoriesSortOrderEnum = "ASC"
	ListCategoriesSortOrderDesc ListCategoriesSortOrderEnum = "DESC"
)

var mappingListCategoriesSortOrderEnum = map[string]ListCategoriesSortOrderEnum{
	"ASC":  ListCategoriesSortOrderAsc,
	"DESC": ListCategoriesSortOrderDesc,
}

var mappingListCategoriesSortOrderEnumLowerCase = map[string]ListCategoriesSortOrderEnum{
	"asc":  ListCategoriesSortOrderAsc,
	"desc": ListCategoriesSortOrderDesc,
}

// GetListCategoriesSortOrderEnumValues Enumerates the set of values for ListCategoriesSortOrderEnum
func GetListCategoriesSortOrderEnumValues() []ListCategoriesSortOrderEnum {
	values := make([]ListCategoriesSortOrderEnum, 0)
	for _, v := range mappingListCategoriesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListCategoriesSortOrderEnumStringValues Enumerates the set of values in String for ListCategoriesSortOrderEnum
func GetListCategoriesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListCategoriesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListCategoriesSortOrderEnum(val string) (ListCategoriesSortOrderEnum, bool) {
	enum, ok := mappingListCategoriesSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListCategoriesSortByEnum Enum with underlying type: string
type ListCategoriesSortByEnum string

// Set of constants representing the allowable values for ListCategoriesSortByEnum
const (
	ListCategoriesSortByTimecreated ListCategoriesSortByEnum = "timeCreated"
	ListCategoriesSortByDisplayname ListCategoriesSortByEnum = "displayName"
)

var mappingListCategoriesSortByEnum = map[string]ListCategoriesSortByEnum{
	"timeCreated": ListCategoriesSortByTimecreated,
	"displayName": ListCategoriesSortByDisplayname,
}

var mappingListCategoriesSortByEnumLowerCase = map[string]ListCategoriesSortByEnum{
	"timecreated": ListCategoriesSortByTimecreated,
	"displayname": ListCategoriesSortByDisplayname,
}

// GetListCategoriesSortByEnumValues Enumerates the set of values for ListCategoriesSortByEnum
func GetListCategoriesSortByEnumValues() []ListCategoriesSortByEnum {
	values := make([]ListCategoriesSortByEnum, 0)
	for _, v := range mappingListCategoriesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListCategoriesSortByEnumStringValues Enumerates the set of values in String for ListCategoriesSortByEnum
func GetListCategoriesSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListCategoriesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListCategoriesSortByEnum(val string) (ListCategoriesSortByEnum, bool) {
	enum, ok := mappingListCategoriesSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
