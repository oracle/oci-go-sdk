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

// ListMarketsRequest wrapper for the ListMarkets operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/marketplacepublisher/ListMarkets.go.html to see an example of how to use ListMarketsRequest.
type ListMarketsRequest struct {

	// The ID of the compartment in which to list resources.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// A token representing the position at which to start retrieving results. This must come from the `opc-next-page` header field of a previous response.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'ASC' or 'DESC'.
	SortOrder ListMarketsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for timeCreated is descending. Default order for displayName is ascending.
	SortBy ListMarketsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// A filter to return only resources that match the entire name given.
	Name *string `mandatory:"false" contributesTo:"query" name:"name"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListMarketsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListMarketsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListMarketsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListMarketsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListMarketsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListMarketsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListMarketsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListMarketsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListMarketsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListMarketsResponse wrapper for the ListMarkets operation
type ListMarketsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of MarketCollection instances
	MarketCollection `presentIn:"body"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response ListMarketsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListMarketsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListMarketsSortOrderEnum Enum with underlying type: string
type ListMarketsSortOrderEnum string

// Set of constants representing the allowable values for ListMarketsSortOrderEnum
const (
	ListMarketsSortOrderAsc  ListMarketsSortOrderEnum = "ASC"
	ListMarketsSortOrderDesc ListMarketsSortOrderEnum = "DESC"
)

var mappingListMarketsSortOrderEnum = map[string]ListMarketsSortOrderEnum{
	"ASC":  ListMarketsSortOrderAsc,
	"DESC": ListMarketsSortOrderDesc,
}

var mappingListMarketsSortOrderEnumLowerCase = map[string]ListMarketsSortOrderEnum{
	"asc":  ListMarketsSortOrderAsc,
	"desc": ListMarketsSortOrderDesc,
}

// GetListMarketsSortOrderEnumValues Enumerates the set of values for ListMarketsSortOrderEnum
func GetListMarketsSortOrderEnumValues() []ListMarketsSortOrderEnum {
	values := make([]ListMarketsSortOrderEnum, 0)
	for _, v := range mappingListMarketsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListMarketsSortOrderEnumStringValues Enumerates the set of values in String for ListMarketsSortOrderEnum
func GetListMarketsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListMarketsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListMarketsSortOrderEnum(val string) (ListMarketsSortOrderEnum, bool) {
	enum, ok := mappingListMarketsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListMarketsSortByEnum Enum with underlying type: string
type ListMarketsSortByEnum string

// Set of constants representing the allowable values for ListMarketsSortByEnum
const (
	ListMarketsSortByTimecreated ListMarketsSortByEnum = "timeCreated"
	ListMarketsSortByDisplayname ListMarketsSortByEnum = "displayName"
)

var mappingListMarketsSortByEnum = map[string]ListMarketsSortByEnum{
	"timeCreated": ListMarketsSortByTimecreated,
	"displayName": ListMarketsSortByDisplayname,
}

var mappingListMarketsSortByEnumLowerCase = map[string]ListMarketsSortByEnum{
	"timecreated": ListMarketsSortByTimecreated,
	"displayname": ListMarketsSortByDisplayname,
}

// GetListMarketsSortByEnumValues Enumerates the set of values for ListMarketsSortByEnum
func GetListMarketsSortByEnumValues() []ListMarketsSortByEnum {
	values := make([]ListMarketsSortByEnum, 0)
	for _, v := range mappingListMarketsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListMarketsSortByEnumStringValues Enumerates the set of values in String for ListMarketsSortByEnum
func GetListMarketsSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListMarketsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListMarketsSortByEnum(val string) (ListMarketsSortByEnum, bool) {
	enum, ok := mappingListMarketsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
