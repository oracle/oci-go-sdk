// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package marketplacepublisher

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListListingsRequest wrapper for the ListListings operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/marketplacepublisher/ListListings.go.html to see an example of how to use ListListingsRequest.
type ListListingsRequest struct {

	// The ID of the compartment in which to list resources.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// A filter to return only listings their lifecycleState matches the given lifecycleState.
	LifecycleState ListingLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// A filter to return only resources that match the entire name given.
	Name *string `mandatory:"false" contributesTo:"query" name:"name"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// A token representing the position at which to start retrieving results. This must come from the `opc-next-page` header field of a previous response.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'ASC' or 'DESC'.
	SortOrder ListListingsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for timeCreated is descending. Default order for displayName is ascending.
	SortBy ListListingsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListListingsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListListingsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListListingsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListListingsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListListingsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListingLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetListingLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListListingsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListListingsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListListingsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListListingsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListListingsResponse wrapper for the ListListings operation
type ListListingsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of ListingCollection instances
	ListingCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListListingsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListListingsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListListingsSortOrderEnum Enum with underlying type: string
type ListListingsSortOrderEnum string

// Set of constants representing the allowable values for ListListingsSortOrderEnum
const (
	ListListingsSortOrderAsc  ListListingsSortOrderEnum = "ASC"
	ListListingsSortOrderDesc ListListingsSortOrderEnum = "DESC"
)

var mappingListListingsSortOrderEnum = map[string]ListListingsSortOrderEnum{
	"ASC":  ListListingsSortOrderAsc,
	"DESC": ListListingsSortOrderDesc,
}

var mappingListListingsSortOrderEnumLowerCase = map[string]ListListingsSortOrderEnum{
	"asc":  ListListingsSortOrderAsc,
	"desc": ListListingsSortOrderDesc,
}

// GetListListingsSortOrderEnumValues Enumerates the set of values for ListListingsSortOrderEnum
func GetListListingsSortOrderEnumValues() []ListListingsSortOrderEnum {
	values := make([]ListListingsSortOrderEnum, 0)
	for _, v := range mappingListListingsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListListingsSortOrderEnumStringValues Enumerates the set of values in String for ListListingsSortOrderEnum
func GetListListingsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListListingsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListListingsSortOrderEnum(val string) (ListListingsSortOrderEnum, bool) {
	enum, ok := mappingListListingsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListListingsSortByEnum Enum with underlying type: string
type ListListingsSortByEnum string

// Set of constants representing the allowable values for ListListingsSortByEnum
const (
	ListListingsSortByTimecreated ListListingsSortByEnum = "timeCreated"
	ListListingsSortByDisplayname ListListingsSortByEnum = "displayName"
)

var mappingListListingsSortByEnum = map[string]ListListingsSortByEnum{
	"timeCreated": ListListingsSortByTimecreated,
	"displayName": ListListingsSortByDisplayname,
}

var mappingListListingsSortByEnumLowerCase = map[string]ListListingsSortByEnum{
	"timecreated": ListListingsSortByTimecreated,
	"displayname": ListListingsSortByDisplayname,
}

// GetListListingsSortByEnumValues Enumerates the set of values for ListListingsSortByEnum
func GetListListingsSortByEnumValues() []ListListingsSortByEnum {
	values := make([]ListListingsSortByEnum, 0)
	for _, v := range mappingListListingsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListListingsSortByEnumStringValues Enumerates the set of values in String for ListListingsSortByEnum
func GetListListingsSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListListingsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListListingsSortByEnum(val string) (ListListingsSortByEnum, bool) {
	enum, ok := mappingListListingsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
