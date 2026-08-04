// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package marketplaceprivateoffer

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListOfferQuotesRequest wrapper for the ListOfferQuotes operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/marketplaceprivateoffer/ListOfferQuotes.go.html to see an example of how to use ListOfferQuotesRequest.
type ListOfferQuotesRequest struct {

	// OCID of the reseller's compartment.
	ResellerCompartmentId *string `mandatory:"false" contributesTo:"query" name:"resellerCompartmentId"`

	// OCID of the ISV's compartment.
	IsvCompartmentId *string `mandatory:"false" contributesTo:"query" name:"isvCompartmentId"`

	// A filter to return resources whose lifecycle state matches the given lifecycle state.
	LifecycleState OfferQuoteLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// A filter to return only resources that match the entire display name given.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// Unique offer quote identifier.
	Id *string `mandatory:"false" contributesTo:"query" name:"id"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// A token representing the position at which to start retrieving results. This must come from the `opc-next-page` header field of a previous response.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'ASC' or 'DESC'.
	SortOrder ListOfferQuotesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for timeCreated is descending. Default order for displayName is ascending.
	SortBy ListOfferQuotesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListOfferQuotesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListOfferQuotesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListOfferQuotesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListOfferQuotesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListOfferQuotesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingOfferQuoteLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetOfferQuoteLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListOfferQuotesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListOfferQuotesSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListOfferQuotesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListOfferQuotesSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListOfferQuotesResponse wrapper for the ListOfferQuotes operation
type ListOfferQuotesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of OfferQuoteCollection instances
	OfferQuoteCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListOfferQuotesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListOfferQuotesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListOfferQuotesSortOrderEnum Enum with underlying type: string
type ListOfferQuotesSortOrderEnum string

// Set of constants representing the allowable values for ListOfferQuotesSortOrderEnum
const (
	ListOfferQuotesSortOrderAsc  ListOfferQuotesSortOrderEnum = "ASC"
	ListOfferQuotesSortOrderDesc ListOfferQuotesSortOrderEnum = "DESC"
)

var mappingListOfferQuotesSortOrderEnum = map[string]ListOfferQuotesSortOrderEnum{
	"ASC":  ListOfferQuotesSortOrderAsc,
	"DESC": ListOfferQuotesSortOrderDesc,
}

var mappingListOfferQuotesSortOrderEnumLowerCase = map[string]ListOfferQuotesSortOrderEnum{
	"asc":  ListOfferQuotesSortOrderAsc,
	"desc": ListOfferQuotesSortOrderDesc,
}

// GetListOfferQuotesSortOrderEnumValues Enumerates the set of values for ListOfferQuotesSortOrderEnum
func GetListOfferQuotesSortOrderEnumValues() []ListOfferQuotesSortOrderEnum {
	values := make([]ListOfferQuotesSortOrderEnum, 0)
	for _, v := range mappingListOfferQuotesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListOfferQuotesSortOrderEnumStringValues Enumerates the set of values in String for ListOfferQuotesSortOrderEnum
func GetListOfferQuotesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListOfferQuotesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListOfferQuotesSortOrderEnum(val string) (ListOfferQuotesSortOrderEnum, bool) {
	enum, ok := mappingListOfferQuotesSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListOfferQuotesSortByEnum Enum with underlying type: string
type ListOfferQuotesSortByEnum string

// Set of constants representing the allowable values for ListOfferQuotesSortByEnum
const (
	ListOfferQuotesSortByTimecreated ListOfferQuotesSortByEnum = "timeCreated"
	ListOfferQuotesSortByDisplayname ListOfferQuotesSortByEnum = "displayName"
)

var mappingListOfferQuotesSortByEnum = map[string]ListOfferQuotesSortByEnum{
	"timeCreated": ListOfferQuotesSortByTimecreated,
	"displayName": ListOfferQuotesSortByDisplayname,
}

var mappingListOfferQuotesSortByEnumLowerCase = map[string]ListOfferQuotesSortByEnum{
	"timecreated": ListOfferQuotesSortByTimecreated,
	"displayname": ListOfferQuotesSortByDisplayname,
}

// GetListOfferQuotesSortByEnumValues Enumerates the set of values for ListOfferQuotesSortByEnum
func GetListOfferQuotesSortByEnumValues() []ListOfferQuotesSortByEnum {
	values := make([]ListOfferQuotesSortByEnum, 0)
	for _, v := range mappingListOfferQuotesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListOfferQuotesSortByEnumStringValues Enumerates the set of values in String for ListOfferQuotesSortByEnum
func GetListOfferQuotesSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListOfferQuotesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListOfferQuotesSortByEnum(val string) (ListOfferQuotesSortByEnum, bool) {
	enum, ok := mappingListOfferQuotesSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
