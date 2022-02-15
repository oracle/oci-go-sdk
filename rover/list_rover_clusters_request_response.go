// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package rover

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"net/http"
	"strings"
)

// ListRoverClustersRequest wrapper for the ListRoverClusters operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/rover/ListRoverClusters.go.html to see an example of how to use ListRoverClustersRequest.
type ListRoverClustersRequest struct {

	// The OCID of the compartment in which to list resources.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// A filter to return only resources that match the entire display name given.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// A filter to return only resources their lifecycleState matches the given lifecycleState.
	LifecycleState ListRoverClustersLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// The sort order to use, either 'asc' or 'desc'.
	SortOrder ListRoverClustersSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for timeCreated is descending. Default order for displayName is ascending. If no value is specified timeCreated is default.
	SortBy ListRoverClustersSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListRoverClustersRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListRoverClustersRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListRoverClustersRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListRoverClustersRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListRoverClustersRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListRoverClustersLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetListRoverClustersLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListRoverClustersSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListRoverClustersSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListRoverClustersSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListRoverClustersSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListRoverClustersResponse wrapper for the ListRoverClusters operation
type ListRoverClustersResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of RoverClusterCollection instances
	RoverClusterCollection `presentIn:"body"`

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

func (response ListRoverClustersResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListRoverClustersResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListRoverClustersLifecycleStateEnum Enum with underlying type: string
type ListRoverClustersLifecycleStateEnum string

// Set of constants representing the allowable values for ListRoverClustersLifecycleStateEnum
const (
	ListRoverClustersLifecycleStateCreating ListRoverClustersLifecycleStateEnum = "CREATING"
	ListRoverClustersLifecycleStateUpdating ListRoverClustersLifecycleStateEnum = "UPDATING"
	ListRoverClustersLifecycleStateActive   ListRoverClustersLifecycleStateEnum = "ACTIVE"
	ListRoverClustersLifecycleStateDeleting ListRoverClustersLifecycleStateEnum = "DELETING"
	ListRoverClustersLifecycleStateDeleted  ListRoverClustersLifecycleStateEnum = "DELETED"
	ListRoverClustersLifecycleStateFailed   ListRoverClustersLifecycleStateEnum = "FAILED"
)

var mappingListRoverClustersLifecycleStateEnum = map[string]ListRoverClustersLifecycleStateEnum{
	"CREATING": ListRoverClustersLifecycleStateCreating,
	"UPDATING": ListRoverClustersLifecycleStateUpdating,
	"ACTIVE":   ListRoverClustersLifecycleStateActive,
	"DELETING": ListRoverClustersLifecycleStateDeleting,
	"DELETED":  ListRoverClustersLifecycleStateDeleted,
	"FAILED":   ListRoverClustersLifecycleStateFailed,
}

// GetListRoverClustersLifecycleStateEnumValues Enumerates the set of values for ListRoverClustersLifecycleStateEnum
func GetListRoverClustersLifecycleStateEnumValues() []ListRoverClustersLifecycleStateEnum {
	values := make([]ListRoverClustersLifecycleStateEnum, 0)
	for _, v := range mappingListRoverClustersLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetListRoverClustersLifecycleStateEnumStringValues Enumerates the set of values in String for ListRoverClustersLifecycleStateEnum
func GetListRoverClustersLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"UPDATING",
		"ACTIVE",
		"DELETING",
		"DELETED",
		"FAILED",
	}
}

// GetMappingListRoverClustersLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListRoverClustersLifecycleStateEnum(val string) (ListRoverClustersLifecycleStateEnum, bool) {
	mappingListRoverClustersLifecycleStateEnumIgnoreCase := make(map[string]ListRoverClustersLifecycleStateEnum)
	for k, v := range mappingListRoverClustersLifecycleStateEnum {
		mappingListRoverClustersLifecycleStateEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingListRoverClustersLifecycleStateEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}

// ListRoverClustersSortOrderEnum Enum with underlying type: string
type ListRoverClustersSortOrderEnum string

// Set of constants representing the allowable values for ListRoverClustersSortOrderEnum
const (
	ListRoverClustersSortOrderAsc  ListRoverClustersSortOrderEnum = "ASC"
	ListRoverClustersSortOrderDesc ListRoverClustersSortOrderEnum = "DESC"
)

var mappingListRoverClustersSortOrderEnum = map[string]ListRoverClustersSortOrderEnum{
	"ASC":  ListRoverClustersSortOrderAsc,
	"DESC": ListRoverClustersSortOrderDesc,
}

// GetListRoverClustersSortOrderEnumValues Enumerates the set of values for ListRoverClustersSortOrderEnum
func GetListRoverClustersSortOrderEnumValues() []ListRoverClustersSortOrderEnum {
	values := make([]ListRoverClustersSortOrderEnum, 0)
	for _, v := range mappingListRoverClustersSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListRoverClustersSortOrderEnumStringValues Enumerates the set of values in String for ListRoverClustersSortOrderEnum
func GetListRoverClustersSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListRoverClustersSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListRoverClustersSortOrderEnum(val string) (ListRoverClustersSortOrderEnum, bool) {
	mappingListRoverClustersSortOrderEnumIgnoreCase := make(map[string]ListRoverClustersSortOrderEnum)
	for k, v := range mappingListRoverClustersSortOrderEnum {
		mappingListRoverClustersSortOrderEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingListRoverClustersSortOrderEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}

// ListRoverClustersSortByEnum Enum with underlying type: string
type ListRoverClustersSortByEnum string

// Set of constants representing the allowable values for ListRoverClustersSortByEnum
const (
	ListRoverClustersSortByTimecreated ListRoverClustersSortByEnum = "timeCreated"
	ListRoverClustersSortByDisplayname ListRoverClustersSortByEnum = "displayName"
)

var mappingListRoverClustersSortByEnum = map[string]ListRoverClustersSortByEnum{
	"timeCreated": ListRoverClustersSortByTimecreated,
	"displayName": ListRoverClustersSortByDisplayname,
}

// GetListRoverClustersSortByEnumValues Enumerates the set of values for ListRoverClustersSortByEnum
func GetListRoverClustersSortByEnumValues() []ListRoverClustersSortByEnum {
	values := make([]ListRoverClustersSortByEnum, 0)
	for _, v := range mappingListRoverClustersSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListRoverClustersSortByEnumStringValues Enumerates the set of values in String for ListRoverClustersSortByEnum
func GetListRoverClustersSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListRoverClustersSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListRoverClustersSortByEnum(val string) (ListRoverClustersSortByEnum, bool) {
	mappingListRoverClustersSortByEnumIgnoreCase := make(map[string]ListRoverClustersSortByEnum)
	for k, v := range mappingListRoverClustersSortByEnum {
		mappingListRoverClustersSortByEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingListRoverClustersSortByEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
