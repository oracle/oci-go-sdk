// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package mngdmac

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListMacOrdersRequest wrapper for the ListMacOrders operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/mngdmac/ListMacOrders.go.html to see an example of how to use ListMacOrdersRequest.
type ListMacOrdersRequest struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment in which to list resources.
	CompartmentId *string `mandatory:"false" contributesTo:"query" name:"compartmentId"`

	// A filter to return only resources that match the given lifecycle state. The
	// state value is case-insensitive.
	LifecycleState MacOrderLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// A filter to return only resources that match the given display name exactly.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the MacOrder.
	Id *string `mandatory:"false" contributesTo:"query" name:"id"`

	// For list pagination. The maximum number of results per page, or items to return in a
	// paginated "List" call. For important details about how pagination works, see
	// List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The value of the opc-next-page response header from the previous
	// "List" call. For important details about how pagination works, see
	// List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either ascending (`ASC`) or descending (`DESC`).
	SortOrder ListMacOrdersSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. You can provide only one sort order. Default order for `timeCreated`
	// is descending. Default order for `displayName` is ascending.
	SortBy ListMacOrdersSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	// The only valid characters for request IDs are letters, numbers,
	// underscore, and dash.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListMacOrdersRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListMacOrdersRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListMacOrdersRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListMacOrdersRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListMacOrdersRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingMacOrderLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetMacOrderLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListMacOrdersSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListMacOrdersSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListMacOrdersSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListMacOrdersSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListMacOrdersResponse wrapper for the ListMacOrders operation
type ListMacOrdersResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of MacOrderCollection instances
	MacOrderCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response, additional pages of results remain. For
	// important details about how pagination works, see List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListMacOrdersResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListMacOrdersResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListMacOrdersSortOrderEnum Enum with underlying type: string
type ListMacOrdersSortOrderEnum string

// Set of constants representing the allowable values for ListMacOrdersSortOrderEnum
const (
	ListMacOrdersSortOrderAsc  ListMacOrdersSortOrderEnum = "ASC"
	ListMacOrdersSortOrderDesc ListMacOrdersSortOrderEnum = "DESC"
)

var mappingListMacOrdersSortOrderEnum = map[string]ListMacOrdersSortOrderEnum{
	"ASC":  ListMacOrdersSortOrderAsc,
	"DESC": ListMacOrdersSortOrderDesc,
}

var mappingListMacOrdersSortOrderEnumLowerCase = map[string]ListMacOrdersSortOrderEnum{
	"asc":  ListMacOrdersSortOrderAsc,
	"desc": ListMacOrdersSortOrderDesc,
}

// GetListMacOrdersSortOrderEnumValues Enumerates the set of values for ListMacOrdersSortOrderEnum
func GetListMacOrdersSortOrderEnumValues() []ListMacOrdersSortOrderEnum {
	values := make([]ListMacOrdersSortOrderEnum, 0)
	for _, v := range mappingListMacOrdersSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListMacOrdersSortOrderEnumStringValues Enumerates the set of values in String for ListMacOrdersSortOrderEnum
func GetListMacOrdersSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListMacOrdersSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListMacOrdersSortOrderEnum(val string) (ListMacOrdersSortOrderEnum, bool) {
	enum, ok := mappingListMacOrdersSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListMacOrdersSortByEnum Enum with underlying type: string
type ListMacOrdersSortByEnum string

// Set of constants representing the allowable values for ListMacOrdersSortByEnum
const (
	ListMacOrdersSortByTimecreated ListMacOrdersSortByEnum = "timeCreated"
	ListMacOrdersSortByDisplayname ListMacOrdersSortByEnum = "displayName"
)

var mappingListMacOrdersSortByEnum = map[string]ListMacOrdersSortByEnum{
	"timeCreated": ListMacOrdersSortByTimecreated,
	"displayName": ListMacOrdersSortByDisplayname,
}

var mappingListMacOrdersSortByEnumLowerCase = map[string]ListMacOrdersSortByEnum{
	"timecreated": ListMacOrdersSortByTimecreated,
	"displayname": ListMacOrdersSortByDisplayname,
}

// GetListMacOrdersSortByEnumValues Enumerates the set of values for ListMacOrdersSortByEnum
func GetListMacOrdersSortByEnumValues() []ListMacOrdersSortByEnum {
	values := make([]ListMacOrdersSortByEnum, 0)
	for _, v := range mappingListMacOrdersSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListMacOrdersSortByEnumStringValues Enumerates the set of values in String for ListMacOrdersSortByEnum
func GetListMacOrdersSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListMacOrdersSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListMacOrdersSortByEnum(val string) (ListMacOrdersSortByEnum, bool) {
	enum, ok := mappingListMacOrdersSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
