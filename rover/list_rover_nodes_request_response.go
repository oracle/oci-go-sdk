// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package rover

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListRoverNodesRequest wrapper for the ListRoverNodes operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/rover/ListRoverNodes.go.html to see an example of how to use ListRoverNodesRequest.
type ListRoverNodesRequest struct {

	// The OCID of the compartment in which to list resources.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// A filter to return only resources that match the entire display name given.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// A filter to return only Nodes of type matched with the given node type.
	NodeType ListRoverNodesNodeTypeEnum `mandatory:"false" contributesTo:"query" name:"nodeType" omitEmpty:"true"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// A filter to return only resources their lifecycleState matches the given lifecycleState.
	LifecycleState ListRoverNodesLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// The sort order to use, either 'asc' or 'desc'.
	SortOrder ListRoverNodesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for timeCreated is descending. Default order for displayName is ascending. If no value is specified timeCreated is default.
	SortBy ListRoverNodesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListRoverNodesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListRoverNodesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListRoverNodesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListRoverNodesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListRoverNodesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListRoverNodesNodeTypeEnum(string(request.NodeType)); !ok && request.NodeType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for NodeType: %s. Supported values are: %s.", request.NodeType, strings.Join(GetListRoverNodesNodeTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListRoverNodesLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetListRoverNodesLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListRoverNodesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListRoverNodesSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListRoverNodesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListRoverNodesSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListRoverNodesResponse wrapper for the ListRoverNodes operation
type ListRoverNodesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of RoverNodeCollection instances
	RoverNodeCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the previous batch of items.
	OpcPrevPage *string `presentIn:"header" name:"opc-prev-page"`
}

func (response ListRoverNodesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListRoverNodesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListRoverNodesNodeTypeEnum Enum with underlying type: string
type ListRoverNodesNodeTypeEnum string

// Set of constants representing the allowable values for ListRoverNodesNodeTypeEnum
const (
	ListRoverNodesNodeTypeStandalone ListRoverNodesNodeTypeEnum = "STANDALONE"
	ListRoverNodesNodeTypeClustered  ListRoverNodesNodeTypeEnum = "CLUSTERED"
	ListRoverNodesNodeTypeStation    ListRoverNodesNodeTypeEnum = "STATION"
)

var mappingListRoverNodesNodeTypeEnum = map[string]ListRoverNodesNodeTypeEnum{
	"STANDALONE": ListRoverNodesNodeTypeStandalone,
	"CLUSTERED":  ListRoverNodesNodeTypeClustered,
	"STATION":    ListRoverNodesNodeTypeStation,
}

var mappingListRoverNodesNodeTypeEnumLowerCase = map[string]ListRoverNodesNodeTypeEnum{
	"standalone": ListRoverNodesNodeTypeStandalone,
	"clustered":  ListRoverNodesNodeTypeClustered,
	"station":    ListRoverNodesNodeTypeStation,
}

// GetListRoverNodesNodeTypeEnumValues Enumerates the set of values for ListRoverNodesNodeTypeEnum
func GetListRoverNodesNodeTypeEnumValues() []ListRoverNodesNodeTypeEnum {
	values := make([]ListRoverNodesNodeTypeEnum, 0)
	for _, v := range mappingListRoverNodesNodeTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetListRoverNodesNodeTypeEnumStringValues Enumerates the set of values in String for ListRoverNodesNodeTypeEnum
func GetListRoverNodesNodeTypeEnumStringValues() []string {
	return []string{
		"STANDALONE",
		"CLUSTERED",
		"STATION",
	}
}

// GetMappingListRoverNodesNodeTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListRoverNodesNodeTypeEnum(val string) (ListRoverNodesNodeTypeEnum, bool) {
	enum, ok := mappingListRoverNodesNodeTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListRoverNodesLifecycleStateEnum Enum with underlying type: string
type ListRoverNodesLifecycleStateEnum string

// Set of constants representing the allowable values for ListRoverNodesLifecycleStateEnum
const (
	ListRoverNodesLifecycleStateCreating ListRoverNodesLifecycleStateEnum = "CREATING"
	ListRoverNodesLifecycleStateUpdating ListRoverNodesLifecycleStateEnum = "UPDATING"
	ListRoverNodesLifecycleStateActive   ListRoverNodesLifecycleStateEnum = "ACTIVE"
	ListRoverNodesLifecycleStateDeleting ListRoverNodesLifecycleStateEnum = "DELETING"
	ListRoverNodesLifecycleStateDeleted  ListRoverNodesLifecycleStateEnum = "DELETED"
	ListRoverNodesLifecycleStateFailed   ListRoverNodesLifecycleStateEnum = "FAILED"
)

var mappingListRoverNodesLifecycleStateEnum = map[string]ListRoverNodesLifecycleStateEnum{
	"CREATING": ListRoverNodesLifecycleStateCreating,
	"UPDATING": ListRoverNodesLifecycleStateUpdating,
	"ACTIVE":   ListRoverNodesLifecycleStateActive,
	"DELETING": ListRoverNodesLifecycleStateDeleting,
	"DELETED":  ListRoverNodesLifecycleStateDeleted,
	"FAILED":   ListRoverNodesLifecycleStateFailed,
}

var mappingListRoverNodesLifecycleStateEnumLowerCase = map[string]ListRoverNodesLifecycleStateEnum{
	"creating": ListRoverNodesLifecycleStateCreating,
	"updating": ListRoverNodesLifecycleStateUpdating,
	"active":   ListRoverNodesLifecycleStateActive,
	"deleting": ListRoverNodesLifecycleStateDeleting,
	"deleted":  ListRoverNodesLifecycleStateDeleted,
	"failed":   ListRoverNodesLifecycleStateFailed,
}

// GetListRoverNodesLifecycleStateEnumValues Enumerates the set of values for ListRoverNodesLifecycleStateEnum
func GetListRoverNodesLifecycleStateEnumValues() []ListRoverNodesLifecycleStateEnum {
	values := make([]ListRoverNodesLifecycleStateEnum, 0)
	for _, v := range mappingListRoverNodesLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetListRoverNodesLifecycleStateEnumStringValues Enumerates the set of values in String for ListRoverNodesLifecycleStateEnum
func GetListRoverNodesLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"UPDATING",
		"ACTIVE",
		"DELETING",
		"DELETED",
		"FAILED",
	}
}

// GetMappingListRoverNodesLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListRoverNodesLifecycleStateEnum(val string) (ListRoverNodesLifecycleStateEnum, bool) {
	enum, ok := mappingListRoverNodesLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListRoverNodesSortOrderEnum Enum with underlying type: string
type ListRoverNodesSortOrderEnum string

// Set of constants representing the allowable values for ListRoverNodesSortOrderEnum
const (
	ListRoverNodesSortOrderAsc  ListRoverNodesSortOrderEnum = "ASC"
	ListRoverNodesSortOrderDesc ListRoverNodesSortOrderEnum = "DESC"
)

var mappingListRoverNodesSortOrderEnum = map[string]ListRoverNodesSortOrderEnum{
	"ASC":  ListRoverNodesSortOrderAsc,
	"DESC": ListRoverNodesSortOrderDesc,
}

var mappingListRoverNodesSortOrderEnumLowerCase = map[string]ListRoverNodesSortOrderEnum{
	"asc":  ListRoverNodesSortOrderAsc,
	"desc": ListRoverNodesSortOrderDesc,
}

// GetListRoverNodesSortOrderEnumValues Enumerates the set of values for ListRoverNodesSortOrderEnum
func GetListRoverNodesSortOrderEnumValues() []ListRoverNodesSortOrderEnum {
	values := make([]ListRoverNodesSortOrderEnum, 0)
	for _, v := range mappingListRoverNodesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListRoverNodesSortOrderEnumStringValues Enumerates the set of values in String for ListRoverNodesSortOrderEnum
func GetListRoverNodesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListRoverNodesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListRoverNodesSortOrderEnum(val string) (ListRoverNodesSortOrderEnum, bool) {
	enum, ok := mappingListRoverNodesSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListRoverNodesSortByEnum Enum with underlying type: string
type ListRoverNodesSortByEnum string

// Set of constants representing the allowable values for ListRoverNodesSortByEnum
const (
	ListRoverNodesSortByTimecreated ListRoverNodesSortByEnum = "timeCreated"
	ListRoverNodesSortByDisplayname ListRoverNodesSortByEnum = "displayName"
)

var mappingListRoverNodesSortByEnum = map[string]ListRoverNodesSortByEnum{
	"timeCreated": ListRoverNodesSortByTimecreated,
	"displayName": ListRoverNodesSortByDisplayname,
}

var mappingListRoverNodesSortByEnumLowerCase = map[string]ListRoverNodesSortByEnum{
	"timecreated": ListRoverNodesSortByTimecreated,
	"displayname": ListRoverNodesSortByDisplayname,
}

// GetListRoverNodesSortByEnumValues Enumerates the set of values for ListRoverNodesSortByEnum
func GetListRoverNodesSortByEnumValues() []ListRoverNodesSortByEnum {
	values := make([]ListRoverNodesSortByEnum, 0)
	for _, v := range mappingListRoverNodesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListRoverNodesSortByEnumStringValues Enumerates the set of values in String for ListRoverNodesSortByEnum
func GetListRoverNodesSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListRoverNodesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListRoverNodesSortByEnum(val string) (ListRoverNodesSortByEnum, bool) {
	enum, ok := mappingListRoverNodesSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
