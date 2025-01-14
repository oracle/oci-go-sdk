// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package aispeech

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListCustomizationsRequest wrapper for the ListCustomizations operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/aispeech/ListCustomizations.go.html to see an example of how to use ListCustomizationsRequest.
type ListCustomizationsRequest struct {

	// The ID of the compartment in which to list resources.
	CompartmentId *string `mandatory:"false" contributesTo:"query" name:"compartmentId"`

	// A filter to return only resources whose lifecycleState matches the given lifecycleState.
	LifecycleState CustomizationLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// A filter to return only resources that match the entire display name given.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// Unique identifier(OCID).
	Id *string `mandatory:"false" contributesTo:"query" name:"id"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// A token representing the position at which to start retrieving results. This must come from the `opc-next-page` header field of a previous response.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'ASC' or 'DESC'.
	SortOrder ListCustomizationsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for timeCreated is descending. Default order for displayName is ascending.
	SortBy ListCustomizationsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListCustomizationsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListCustomizationsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListCustomizationsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListCustomizationsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListCustomizationsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingCustomizationLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetCustomizationLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListCustomizationsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListCustomizationsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListCustomizationsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListCustomizationsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListCustomizationsResponse wrapper for the ListCustomizations operation
type ListCustomizationsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of CustomizationCollection instances
	CustomizationCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// Retrieves the previous page of results. When this header appears in the response, previous pages of results exist. See List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcPrevPage *string `presentIn:"header" name:"opc-prev-page"`
}

func (response ListCustomizationsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListCustomizationsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListCustomizationsSortOrderEnum Enum with underlying type: string
type ListCustomizationsSortOrderEnum string

// Set of constants representing the allowable values for ListCustomizationsSortOrderEnum
const (
	ListCustomizationsSortOrderAsc  ListCustomizationsSortOrderEnum = "ASC"
	ListCustomizationsSortOrderDesc ListCustomizationsSortOrderEnum = "DESC"
)

var mappingListCustomizationsSortOrderEnum = map[string]ListCustomizationsSortOrderEnum{
	"ASC":  ListCustomizationsSortOrderAsc,
	"DESC": ListCustomizationsSortOrderDesc,
}

var mappingListCustomizationsSortOrderEnumLowerCase = map[string]ListCustomizationsSortOrderEnum{
	"asc":  ListCustomizationsSortOrderAsc,
	"desc": ListCustomizationsSortOrderDesc,
}

// GetListCustomizationsSortOrderEnumValues Enumerates the set of values for ListCustomizationsSortOrderEnum
func GetListCustomizationsSortOrderEnumValues() []ListCustomizationsSortOrderEnum {
	values := make([]ListCustomizationsSortOrderEnum, 0)
	for _, v := range mappingListCustomizationsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListCustomizationsSortOrderEnumStringValues Enumerates the set of values in String for ListCustomizationsSortOrderEnum
func GetListCustomizationsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListCustomizationsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListCustomizationsSortOrderEnum(val string) (ListCustomizationsSortOrderEnum, bool) {
	enum, ok := mappingListCustomizationsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListCustomizationsSortByEnum Enum with underlying type: string
type ListCustomizationsSortByEnum string

// Set of constants representing the allowable values for ListCustomizationsSortByEnum
const (
	ListCustomizationsSortByTimecreated ListCustomizationsSortByEnum = "timeCreated"
	ListCustomizationsSortByDisplayname ListCustomizationsSortByEnum = "displayName"
)

var mappingListCustomizationsSortByEnum = map[string]ListCustomizationsSortByEnum{
	"timeCreated": ListCustomizationsSortByTimecreated,
	"displayName": ListCustomizationsSortByDisplayname,
}

var mappingListCustomizationsSortByEnumLowerCase = map[string]ListCustomizationsSortByEnum{
	"timecreated": ListCustomizationsSortByTimecreated,
	"displayname": ListCustomizationsSortByDisplayname,
}

// GetListCustomizationsSortByEnumValues Enumerates the set of values for ListCustomizationsSortByEnum
func GetListCustomizationsSortByEnumValues() []ListCustomizationsSortByEnum {
	values := make([]ListCustomizationsSortByEnum, 0)
	for _, v := range mappingListCustomizationsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListCustomizationsSortByEnumStringValues Enumerates the set of values in String for ListCustomizationsSortByEnum
func GetListCustomizationsSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListCustomizationsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListCustomizationsSortByEnum(val string) (ListCustomizationsSortByEnum, bool) {
	enum, ok := mappingListCustomizationsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
