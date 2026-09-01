// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package iot

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListIotFlowRuntimesRequest wrapper for the ListIotFlowRuntimes operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/iot/ListIotFlowRuntimes.go.html to see an example of how to use ListIotFlowRuntimesRequest.
type ListIotFlowRuntimesRequest struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment in which to list resources.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// Filter resources by OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm). Must be a valid OCID of the resource type.
	Id *string `mandatory:"false" contributesTo:"query" name:"id"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the IoT domain in which to list flow runtime.
	IotDomainId *string `mandatory:"false" contributesTo:"query" name:"iotDomainId"`

	// Filter resources whose display name matches the specified value.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// Filter resources whose lifecycleState matches the specified value.
	LifecycleState IotFlowRuntimeLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// For list pagination. The maximum number of results per page, or items to return in a
	// paginated "List" call. For important details about how pagination works, see
	// List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination: The value of the opc-next-page response header from the previous "List" call.
	// For important details on how pagination works, see
	// List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// Specifies sort order to use, either ASC (ascending) or DESC (descending).
	SortOrder ListIotFlowRuntimesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided.
	// Default order for timeCreated is descending. Default order for displayName is ascending.
	SortBy ListIotFlowRuntimesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Unique Oracle-assigned identifier for the request.
	// If you need to contact Oracle about a particular request, provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListIotFlowRuntimesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListIotFlowRuntimesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListIotFlowRuntimesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListIotFlowRuntimesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListIotFlowRuntimesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingIotFlowRuntimeLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetIotFlowRuntimeLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListIotFlowRuntimesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListIotFlowRuntimesSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListIotFlowRuntimesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListIotFlowRuntimesSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListIotFlowRuntimesResponse wrapper for the ListIotFlowRuntimes operation
type ListIotFlowRuntimesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of IotFlowRuntimeCollection instances
	IotFlowRuntimeCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response, additional pages of results remain. For
	// important details about how pagination works, see List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// For list pagination. When this header appears in the response, additional pages of results remain. For
	// important details about how pagination works, see List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcPrevPage *string `presentIn:"header" name:"opc-prev-page"`
}

func (response ListIotFlowRuntimesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListIotFlowRuntimesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListIotFlowRuntimesSortOrderEnum Enum with underlying type: string
type ListIotFlowRuntimesSortOrderEnum string

// Set of constants representing the allowable values for ListIotFlowRuntimesSortOrderEnum
const (
	ListIotFlowRuntimesSortOrderAsc  ListIotFlowRuntimesSortOrderEnum = "ASC"
	ListIotFlowRuntimesSortOrderDesc ListIotFlowRuntimesSortOrderEnum = "DESC"
)

var mappingListIotFlowRuntimesSortOrderEnum = map[string]ListIotFlowRuntimesSortOrderEnum{
	"ASC":  ListIotFlowRuntimesSortOrderAsc,
	"DESC": ListIotFlowRuntimesSortOrderDesc,
}

var mappingListIotFlowRuntimesSortOrderEnumLowerCase = map[string]ListIotFlowRuntimesSortOrderEnum{
	"asc":  ListIotFlowRuntimesSortOrderAsc,
	"desc": ListIotFlowRuntimesSortOrderDesc,
}

// GetListIotFlowRuntimesSortOrderEnumValues Enumerates the set of values for ListIotFlowRuntimesSortOrderEnum
func GetListIotFlowRuntimesSortOrderEnumValues() []ListIotFlowRuntimesSortOrderEnum {
	values := make([]ListIotFlowRuntimesSortOrderEnum, 0)
	for _, v := range mappingListIotFlowRuntimesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListIotFlowRuntimesSortOrderEnumStringValues Enumerates the set of values in String for ListIotFlowRuntimesSortOrderEnum
func GetListIotFlowRuntimesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListIotFlowRuntimesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListIotFlowRuntimesSortOrderEnum(val string) (ListIotFlowRuntimesSortOrderEnum, bool) {
	enum, ok := mappingListIotFlowRuntimesSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListIotFlowRuntimesSortByEnum Enum with underlying type: string
type ListIotFlowRuntimesSortByEnum string

// Set of constants representing the allowable values for ListIotFlowRuntimesSortByEnum
const (
	ListIotFlowRuntimesSortByTimecreated ListIotFlowRuntimesSortByEnum = "timeCreated"
	ListIotFlowRuntimesSortByDisplayname ListIotFlowRuntimesSortByEnum = "displayName"
)

var mappingListIotFlowRuntimesSortByEnum = map[string]ListIotFlowRuntimesSortByEnum{
	"timeCreated": ListIotFlowRuntimesSortByTimecreated,
	"displayName": ListIotFlowRuntimesSortByDisplayname,
}

var mappingListIotFlowRuntimesSortByEnumLowerCase = map[string]ListIotFlowRuntimesSortByEnum{
	"timecreated": ListIotFlowRuntimesSortByTimecreated,
	"displayname": ListIotFlowRuntimesSortByDisplayname,
}

// GetListIotFlowRuntimesSortByEnumValues Enumerates the set of values for ListIotFlowRuntimesSortByEnum
func GetListIotFlowRuntimesSortByEnumValues() []ListIotFlowRuntimesSortByEnum {
	values := make([]ListIotFlowRuntimesSortByEnum, 0)
	for _, v := range mappingListIotFlowRuntimesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListIotFlowRuntimesSortByEnumStringValues Enumerates the set of values in String for ListIotFlowRuntimesSortByEnum
func GetListIotFlowRuntimesSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListIotFlowRuntimesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListIotFlowRuntimesSortByEnum(val string) (ListIotFlowRuntimesSortByEnum, bool) {
	enum, ok := mappingListIotFlowRuntimesSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
