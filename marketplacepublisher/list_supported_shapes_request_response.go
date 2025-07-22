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

// ListSupportedShapesRequest wrapper for the ListSupportedShapes operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/marketplacepublisher/ListSupportedShapes.go.html to see an example of how to use ListSupportedShapesRequest.
type ListSupportedShapesRequest struct {

	// The ID of the compartment in which to list resources.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// A token representing the position at which to start retrieving results. This must come from the `opc-next-page` header field of a previous response.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'ASC' or 'DESC'.
	SortOrder ListSupportedShapesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for timeCreated is descending. Default order for displayName is ascending.
	SortBy ListSupportedShapesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// A filter to return only resources that match the entire name given.
	Name *string `mandatory:"false" contributesTo:"query" name:"name"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListSupportedShapesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListSupportedShapesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListSupportedShapesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListSupportedShapesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListSupportedShapesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListSupportedShapesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListSupportedShapesSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListSupportedShapesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListSupportedShapesSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListSupportedShapesResponse wrapper for the ListSupportedShapes operation
type ListSupportedShapesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of SupportedShapeCollection instances
	SupportedShapeCollection `presentIn:"body"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response ListSupportedShapesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListSupportedShapesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListSupportedShapesSortOrderEnum Enum with underlying type: string
type ListSupportedShapesSortOrderEnum string

// Set of constants representing the allowable values for ListSupportedShapesSortOrderEnum
const (
	ListSupportedShapesSortOrderAsc  ListSupportedShapesSortOrderEnum = "ASC"
	ListSupportedShapesSortOrderDesc ListSupportedShapesSortOrderEnum = "DESC"
)

var mappingListSupportedShapesSortOrderEnum = map[string]ListSupportedShapesSortOrderEnum{
	"ASC":  ListSupportedShapesSortOrderAsc,
	"DESC": ListSupportedShapesSortOrderDesc,
}

var mappingListSupportedShapesSortOrderEnumLowerCase = map[string]ListSupportedShapesSortOrderEnum{
	"asc":  ListSupportedShapesSortOrderAsc,
	"desc": ListSupportedShapesSortOrderDesc,
}

// GetListSupportedShapesSortOrderEnumValues Enumerates the set of values for ListSupportedShapesSortOrderEnum
func GetListSupportedShapesSortOrderEnumValues() []ListSupportedShapesSortOrderEnum {
	values := make([]ListSupportedShapesSortOrderEnum, 0)
	for _, v := range mappingListSupportedShapesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListSupportedShapesSortOrderEnumStringValues Enumerates the set of values in String for ListSupportedShapesSortOrderEnum
func GetListSupportedShapesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListSupportedShapesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListSupportedShapesSortOrderEnum(val string) (ListSupportedShapesSortOrderEnum, bool) {
	enum, ok := mappingListSupportedShapesSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListSupportedShapesSortByEnum Enum with underlying type: string
type ListSupportedShapesSortByEnum string

// Set of constants representing the allowable values for ListSupportedShapesSortByEnum
const (
	ListSupportedShapesSortByTimecreated ListSupportedShapesSortByEnum = "timeCreated"
	ListSupportedShapesSortByDisplayname ListSupportedShapesSortByEnum = "displayName"
)

var mappingListSupportedShapesSortByEnum = map[string]ListSupportedShapesSortByEnum{
	"timeCreated": ListSupportedShapesSortByTimecreated,
	"displayName": ListSupportedShapesSortByDisplayname,
}

var mappingListSupportedShapesSortByEnumLowerCase = map[string]ListSupportedShapesSortByEnum{
	"timecreated": ListSupportedShapesSortByTimecreated,
	"displayname": ListSupportedShapesSortByDisplayname,
}

// GetListSupportedShapesSortByEnumValues Enumerates the set of values for ListSupportedShapesSortByEnum
func GetListSupportedShapesSortByEnumValues() []ListSupportedShapesSortByEnum {
	values := make([]ListSupportedShapesSortByEnum, 0)
	for _, v := range mappingListSupportedShapesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListSupportedShapesSortByEnumStringValues Enumerates the set of values in String for ListSupportedShapesSortByEnum
func GetListSupportedShapesSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListSupportedShapesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListSupportedShapesSortByEnum(val string) (ListSupportedShapesSortByEnum, bool) {
	enum, ok := mappingListSupportedShapesSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
