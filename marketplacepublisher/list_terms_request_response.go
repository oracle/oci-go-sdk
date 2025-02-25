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

// ListTermsRequest wrapper for the ListTerms operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/marketplacepublisher/ListTerms.go.html to see an example of how to use ListTermsRequest.
type ListTermsRequest struct {

	// The ID of the compartment in which to list resources.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// A token representing the position at which to start retrieving results. This must come from the `opc-next-page` header field of a previous response.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'ASC' or 'DESC'.
	SortOrder ListTermsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for timeCreated is descending. Default order for displayName is ascending.
	SortBy ListTermsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// A filter to return only resources that match the entire name given.
	Name *string `mandatory:"false" contributesTo:"query" name:"name"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListTermsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListTermsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListTermsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListTermsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListTermsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListTermsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListTermsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListTermsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListTermsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListTermsResponse wrapper for the ListTerms operation
type ListTermsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of TermCollection instances
	TermCollection `presentIn:"body"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response ListTermsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListTermsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListTermsSortOrderEnum Enum with underlying type: string
type ListTermsSortOrderEnum string

// Set of constants representing the allowable values for ListTermsSortOrderEnum
const (
	ListTermsSortOrderAsc  ListTermsSortOrderEnum = "ASC"
	ListTermsSortOrderDesc ListTermsSortOrderEnum = "DESC"
)

var mappingListTermsSortOrderEnum = map[string]ListTermsSortOrderEnum{
	"ASC":  ListTermsSortOrderAsc,
	"DESC": ListTermsSortOrderDesc,
}

var mappingListTermsSortOrderEnumLowerCase = map[string]ListTermsSortOrderEnum{
	"asc":  ListTermsSortOrderAsc,
	"desc": ListTermsSortOrderDesc,
}

// GetListTermsSortOrderEnumValues Enumerates the set of values for ListTermsSortOrderEnum
func GetListTermsSortOrderEnumValues() []ListTermsSortOrderEnum {
	values := make([]ListTermsSortOrderEnum, 0)
	for _, v := range mappingListTermsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListTermsSortOrderEnumStringValues Enumerates the set of values in String for ListTermsSortOrderEnum
func GetListTermsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListTermsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListTermsSortOrderEnum(val string) (ListTermsSortOrderEnum, bool) {
	enum, ok := mappingListTermsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListTermsSortByEnum Enum with underlying type: string
type ListTermsSortByEnum string

// Set of constants representing the allowable values for ListTermsSortByEnum
const (
	ListTermsSortByTimecreated ListTermsSortByEnum = "timeCreated"
	ListTermsSortByDisplayname ListTermsSortByEnum = "displayName"
)

var mappingListTermsSortByEnum = map[string]ListTermsSortByEnum{
	"timeCreated": ListTermsSortByTimecreated,
	"displayName": ListTermsSortByDisplayname,
}

var mappingListTermsSortByEnumLowerCase = map[string]ListTermsSortByEnum{
	"timecreated": ListTermsSortByTimecreated,
	"displayname": ListTermsSortByDisplayname,
}

// GetListTermsSortByEnumValues Enumerates the set of values for ListTermsSortByEnum
func GetListTermsSortByEnumValues() []ListTermsSortByEnum {
	values := make([]ListTermsSortByEnum, 0)
	for _, v := range mappingListTermsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListTermsSortByEnumStringValues Enumerates the set of values in String for ListTermsSortByEnum
func GetListTermsSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListTermsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListTermsSortByEnum(val string) (ListTermsSortByEnum, bool) {
	enum, ok := mappingListTermsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
