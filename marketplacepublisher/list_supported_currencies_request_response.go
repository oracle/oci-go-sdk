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

// ListSupportedCurrenciesRequest wrapper for the ListSupportedCurrencies operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/marketplacepublisher/ListSupportedCurrencies.go.html to see an example of how to use ListSupportedCurrenciesRequest.
type ListSupportedCurrenciesRequest struct {

	// The ID of the compartment in which to list resources.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// A token representing the position at which to start retrieving results. This must come from the `opc-next-page` header field of a previous response.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'ASC' or 'DESC'.
	SortOrder ListSupportedCurrenciesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for timeCreated is descending. Default order for displayName is ascending.
	SortBy ListSupportedCurrenciesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListSupportedCurrenciesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListSupportedCurrenciesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListSupportedCurrenciesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListSupportedCurrenciesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListSupportedCurrenciesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListSupportedCurrenciesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListSupportedCurrenciesSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListSupportedCurrenciesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListSupportedCurrenciesSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListSupportedCurrenciesResponse wrapper for the ListSupportedCurrencies operation
type ListSupportedCurrenciesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of SupportedCurrencyCollection instances
	SupportedCurrencyCollection `presentIn:"body"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response ListSupportedCurrenciesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListSupportedCurrenciesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListSupportedCurrenciesSortOrderEnum Enum with underlying type: string
type ListSupportedCurrenciesSortOrderEnum string

// Set of constants representing the allowable values for ListSupportedCurrenciesSortOrderEnum
const (
	ListSupportedCurrenciesSortOrderAsc  ListSupportedCurrenciesSortOrderEnum = "ASC"
	ListSupportedCurrenciesSortOrderDesc ListSupportedCurrenciesSortOrderEnum = "DESC"
)

var mappingListSupportedCurrenciesSortOrderEnum = map[string]ListSupportedCurrenciesSortOrderEnum{
	"ASC":  ListSupportedCurrenciesSortOrderAsc,
	"DESC": ListSupportedCurrenciesSortOrderDesc,
}

var mappingListSupportedCurrenciesSortOrderEnumLowerCase = map[string]ListSupportedCurrenciesSortOrderEnum{
	"asc":  ListSupportedCurrenciesSortOrderAsc,
	"desc": ListSupportedCurrenciesSortOrderDesc,
}

// GetListSupportedCurrenciesSortOrderEnumValues Enumerates the set of values for ListSupportedCurrenciesSortOrderEnum
func GetListSupportedCurrenciesSortOrderEnumValues() []ListSupportedCurrenciesSortOrderEnum {
	values := make([]ListSupportedCurrenciesSortOrderEnum, 0)
	for _, v := range mappingListSupportedCurrenciesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListSupportedCurrenciesSortOrderEnumStringValues Enumerates the set of values in String for ListSupportedCurrenciesSortOrderEnum
func GetListSupportedCurrenciesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListSupportedCurrenciesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListSupportedCurrenciesSortOrderEnum(val string) (ListSupportedCurrenciesSortOrderEnum, bool) {
	enum, ok := mappingListSupportedCurrenciesSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListSupportedCurrenciesSortByEnum Enum with underlying type: string
type ListSupportedCurrenciesSortByEnum string

// Set of constants representing the allowable values for ListSupportedCurrenciesSortByEnum
const (
	ListSupportedCurrenciesSortByTimecreated ListSupportedCurrenciesSortByEnum = "timeCreated"
	ListSupportedCurrenciesSortByDisplayname ListSupportedCurrenciesSortByEnum = "displayName"
)

var mappingListSupportedCurrenciesSortByEnum = map[string]ListSupportedCurrenciesSortByEnum{
	"timeCreated": ListSupportedCurrenciesSortByTimecreated,
	"displayName": ListSupportedCurrenciesSortByDisplayname,
}

var mappingListSupportedCurrenciesSortByEnumLowerCase = map[string]ListSupportedCurrenciesSortByEnum{
	"timecreated": ListSupportedCurrenciesSortByTimecreated,
	"displayname": ListSupportedCurrenciesSortByDisplayname,
}

// GetListSupportedCurrenciesSortByEnumValues Enumerates the set of values for ListSupportedCurrenciesSortByEnum
func GetListSupportedCurrenciesSortByEnumValues() []ListSupportedCurrenciesSortByEnum {
	values := make([]ListSupportedCurrenciesSortByEnum, 0)
	for _, v := range mappingListSupportedCurrenciesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListSupportedCurrenciesSortByEnumStringValues Enumerates the set of values in String for ListSupportedCurrenciesSortByEnum
func GetListSupportedCurrenciesSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListSupportedCurrenciesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListSupportedCurrenciesSortByEnum(val string) (ListSupportedCurrenciesSortByEnum, bool) {
	enum, ok := mappingListSupportedCurrenciesSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
