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

// ListPublishersRequest wrapper for the ListPublishers operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/marketplacepublisher/ListPublishers.go.html to see an example of how to use ListPublishersRequest.
type ListPublishersRequest struct {

	// The ID of the compartment in which to list resources.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// A token representing the position at which to start retrieving results. This must come from the `opc-next-page` header field of a previous response.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'ASC' or 'DESC'.
	SortOrder ListPublishersSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for timeCreated is descending. Default order for displayName is ascending.
	SortBy ListPublishersSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// A filter to return only resources that match the entire name given.
	Name *string `mandatory:"false" contributesTo:"query" name:"name"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListPublishersRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListPublishersRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListPublishersRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListPublishersRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListPublishersRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListPublishersSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListPublishersSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListPublishersSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListPublishersSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListPublishersResponse wrapper for the ListPublishers operation
type ListPublishersResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of PublisherCollection instances
	PublisherCollection `presentIn:"body"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response ListPublishersResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListPublishersResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListPublishersSortOrderEnum Enum with underlying type: string
type ListPublishersSortOrderEnum string

// Set of constants representing the allowable values for ListPublishersSortOrderEnum
const (
	ListPublishersSortOrderAsc  ListPublishersSortOrderEnum = "ASC"
	ListPublishersSortOrderDesc ListPublishersSortOrderEnum = "DESC"
)

var mappingListPublishersSortOrderEnum = map[string]ListPublishersSortOrderEnum{
	"ASC":  ListPublishersSortOrderAsc,
	"DESC": ListPublishersSortOrderDesc,
}

var mappingListPublishersSortOrderEnumLowerCase = map[string]ListPublishersSortOrderEnum{
	"asc":  ListPublishersSortOrderAsc,
	"desc": ListPublishersSortOrderDesc,
}

// GetListPublishersSortOrderEnumValues Enumerates the set of values for ListPublishersSortOrderEnum
func GetListPublishersSortOrderEnumValues() []ListPublishersSortOrderEnum {
	values := make([]ListPublishersSortOrderEnum, 0)
	for _, v := range mappingListPublishersSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListPublishersSortOrderEnumStringValues Enumerates the set of values in String for ListPublishersSortOrderEnum
func GetListPublishersSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListPublishersSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListPublishersSortOrderEnum(val string) (ListPublishersSortOrderEnum, bool) {
	enum, ok := mappingListPublishersSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListPublishersSortByEnum Enum with underlying type: string
type ListPublishersSortByEnum string

// Set of constants representing the allowable values for ListPublishersSortByEnum
const (
	ListPublishersSortByTimecreated ListPublishersSortByEnum = "timeCreated"
	ListPublishersSortByDisplayname ListPublishersSortByEnum = "displayName"
)

var mappingListPublishersSortByEnum = map[string]ListPublishersSortByEnum{
	"timeCreated": ListPublishersSortByTimecreated,
	"displayName": ListPublishersSortByDisplayname,
}

var mappingListPublishersSortByEnumLowerCase = map[string]ListPublishersSortByEnum{
	"timecreated": ListPublishersSortByTimecreated,
	"displayname": ListPublishersSortByDisplayname,
}

// GetListPublishersSortByEnumValues Enumerates the set of values for ListPublishersSortByEnum
func GetListPublishersSortByEnumValues() []ListPublishersSortByEnum {
	values := make([]ListPublishersSortByEnum, 0)
	for _, v := range mappingListPublishersSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListPublishersSortByEnumStringValues Enumerates the set of values in String for ListPublishersSortByEnum
func GetListPublishersSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListPublishersSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListPublishersSortByEnum(val string) (ListPublishersSortByEnum, bool) {
	enum, ok := mappingListPublishersSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
