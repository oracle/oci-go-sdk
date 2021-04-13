// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package rover

import (
	"github.com/oracle/oci-go-sdk/v39/common"
	"net/http"
)

// ListRoverEntitlementsRequest wrapper for the ListRoverEntitlements operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/rover/ListRoverEntitlements.go.html to see an example of how to use ListRoverEntitlementsRequest.
type ListRoverEntitlementsRequest struct {

	// The OCID of the compartment in which to list resources.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// filtering by Rover Device Entitlement id
	Id *string `mandatory:"false" contributesTo:"query" name:"id"`

	// A filter to return only resources that match the entire display name given.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// A filter to return only resources their lifecycleState matches the given lifecycleState.
	LifecycleState ListRoverEntitlementsLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'asc' or 'desc'.
	SortOrder ListRoverEntitlementsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for timeCreated is descending. Default order for displayName is ascending. If no value is specified timeCreated is default.
	SortBy ListRoverEntitlementsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListRoverEntitlementsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListRoverEntitlementsRequest) HTTPRequest(method, path string) (http.Request, error) {
	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListRoverEntitlementsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ListRoverEntitlementsResponse wrapper for the ListRoverEntitlements operation
type ListRoverEntitlementsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of RoverEntitlementCollection instances
	RoverEntitlementCollection `presentIn:"body"`

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

func (response ListRoverEntitlementsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListRoverEntitlementsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListRoverEntitlementsLifecycleStateEnum Enum with underlying type: string
type ListRoverEntitlementsLifecycleStateEnum string

// Set of constants representing the allowable values for ListRoverEntitlementsLifecycleStateEnum
const (
	ListRoverEntitlementsLifecycleStateCreating ListRoverEntitlementsLifecycleStateEnum = "CREATING"
	ListRoverEntitlementsLifecycleStateUpdating ListRoverEntitlementsLifecycleStateEnum = "UPDATING"
	ListRoverEntitlementsLifecycleStateActive   ListRoverEntitlementsLifecycleStateEnum = "ACTIVE"
	ListRoverEntitlementsLifecycleStateDeleting ListRoverEntitlementsLifecycleStateEnum = "DELETING"
	ListRoverEntitlementsLifecycleStateDeleted  ListRoverEntitlementsLifecycleStateEnum = "DELETED"
	ListRoverEntitlementsLifecycleStateFailed   ListRoverEntitlementsLifecycleStateEnum = "FAILED"
)

var mappingListRoverEntitlementsLifecycleState = map[string]ListRoverEntitlementsLifecycleStateEnum{
	"CREATING": ListRoverEntitlementsLifecycleStateCreating,
	"UPDATING": ListRoverEntitlementsLifecycleStateUpdating,
	"ACTIVE":   ListRoverEntitlementsLifecycleStateActive,
	"DELETING": ListRoverEntitlementsLifecycleStateDeleting,
	"DELETED":  ListRoverEntitlementsLifecycleStateDeleted,
	"FAILED":   ListRoverEntitlementsLifecycleStateFailed,
}

// GetListRoverEntitlementsLifecycleStateEnumValues Enumerates the set of values for ListRoverEntitlementsLifecycleStateEnum
func GetListRoverEntitlementsLifecycleStateEnumValues() []ListRoverEntitlementsLifecycleStateEnum {
	values := make([]ListRoverEntitlementsLifecycleStateEnum, 0)
	for _, v := range mappingListRoverEntitlementsLifecycleState {
		values = append(values, v)
	}
	return values
}

// ListRoverEntitlementsSortOrderEnum Enum with underlying type: string
type ListRoverEntitlementsSortOrderEnum string

// Set of constants representing the allowable values for ListRoverEntitlementsSortOrderEnum
const (
	ListRoverEntitlementsSortOrderAsc  ListRoverEntitlementsSortOrderEnum = "ASC"
	ListRoverEntitlementsSortOrderDesc ListRoverEntitlementsSortOrderEnum = "DESC"
)

var mappingListRoverEntitlementsSortOrder = map[string]ListRoverEntitlementsSortOrderEnum{
	"ASC":  ListRoverEntitlementsSortOrderAsc,
	"DESC": ListRoverEntitlementsSortOrderDesc,
}

// GetListRoverEntitlementsSortOrderEnumValues Enumerates the set of values for ListRoverEntitlementsSortOrderEnum
func GetListRoverEntitlementsSortOrderEnumValues() []ListRoverEntitlementsSortOrderEnum {
	values := make([]ListRoverEntitlementsSortOrderEnum, 0)
	for _, v := range mappingListRoverEntitlementsSortOrder {
		values = append(values, v)
	}
	return values
}

// ListRoverEntitlementsSortByEnum Enum with underlying type: string
type ListRoverEntitlementsSortByEnum string

// Set of constants representing the allowable values for ListRoverEntitlementsSortByEnum
const (
	ListRoverEntitlementsSortByTimecreated ListRoverEntitlementsSortByEnum = "timeCreated"
	ListRoverEntitlementsSortByDisplayname ListRoverEntitlementsSortByEnum = "displayName"
)

var mappingListRoverEntitlementsSortBy = map[string]ListRoverEntitlementsSortByEnum{
	"timeCreated": ListRoverEntitlementsSortByTimecreated,
	"displayName": ListRoverEntitlementsSortByDisplayname,
}

// GetListRoverEntitlementsSortByEnumValues Enumerates the set of values for ListRoverEntitlementsSortByEnum
func GetListRoverEntitlementsSortByEnumValues() []ListRoverEntitlementsSortByEnum {
	values := make([]ListRoverEntitlementsSortByEnum, 0)
	for _, v := range mappingListRoverEntitlementsSortBy {
		values = append(values, v)
	}
	return values
}
