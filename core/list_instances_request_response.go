// Copyright (c) 2016, 2017, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

package core

import (
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

// ListInstancesRequest wrapper for the ListInstances operation
type ListInstancesRequest struct {

	// The OCID of the compartment.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// The name of the Availability Domain.
	// Example: `Uocm:PHX-AD-1`
	AvailabilityDomain *string `mandatory:"false" contributesTo:"query" name:"availabilityDomain"`

	// A filter to return only resources that match the given display name exactly.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// The maximum number of items to return in a paginated "List" call.
	// Example: `500`
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The value of the `opc-next-page` response header from the previous "List" call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The field to sort by. You can provide one sort order (`sortOrder`). Default order for
	// TIMECREATED is descending. Default order for DISPLAYNAME is ascending. The DISPLAYNAME
	// sort order is case sensitive.
	// **Note:** In general, some "List" operations (for example, `ListInstances`) let you
	// optionally filter by Availability Domain if the scope of the resource type is within a
	// single Availability Domain. If you call one of these "List" operations without specifying
	// an Availability Domain, the resources are grouped by Availability Domain, then sorted.
	SortBy ListInstancesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to use, either ascending (`ASC`) or descending (`DESC`). The DISPLAYNAME sort order
	// is case sensitive.
	SortOrder ListInstancesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// A filter to only return resources that match the given lifecycle state.  The state value is case-insensitive.
	LifecycleState InstanceLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListInstancesRequest) String() string {
	return common.PointerString(request)
}

// GetHttpRequest implements the OciRequest interface
func (request ListInstancesRequest) GetHttpRequest(method, path string) (http.Request, error) {
	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// GetRetryPolicy implements the OciRetryableRequest interface
// => assembles retry policy based on specified options and default behavior
func (request ListInstancesRequest) GetRetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ListInstancesResponse wrapper for the ListInstances operation
type ListInstancesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The []Instance instance
	Items []Instance `presentIn:"body"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	// a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response ListInstancesResponse) String() string {
	return common.PointerString(response)
}

// GetRawResponse implements the OciResponse interface
func (response ListInstancesResponse) GetRawResponse() *http.Response {
	return response.RawResponse
}

// ListInstancesSortByEnum Enum with underlying type: string
type ListInstancesSortByEnum string

// Set of constants representing the allowable values for ListInstancesSortBy
const (
	ListInstancesSortByTimecreated ListInstancesSortByEnum = "TIMECREATED"
	ListInstancesSortByDisplayname ListInstancesSortByEnum = "DISPLAYNAME"
	ListInstancesSortByUnknown     ListInstancesSortByEnum = "UNKNOWN"
)

var mappingListInstancesSortBy = map[string]ListInstancesSortByEnum{
	"TIMECREATED": ListInstancesSortByTimecreated,
	"DISPLAYNAME": ListInstancesSortByDisplayname,
	"UNKNOWN":     ListInstancesSortByUnknown,
}

// GetListInstancesSortByEnumValues Enumerates the set of values for ListInstancesSortBy
func GetListInstancesSortByEnumValues() []ListInstancesSortByEnum {
	values := make([]ListInstancesSortByEnum, 0)
	for _, v := range mappingListInstancesSortBy {
		if v != ListInstancesSortByUnknown {
			values = append(values, v)
		}
	}
	return values
}

// ListInstancesSortOrderEnum Enum with underlying type: string
type ListInstancesSortOrderEnum string

// Set of constants representing the allowable values for ListInstancesSortOrder
const (
	ListInstancesSortOrderAsc     ListInstancesSortOrderEnum = "ASC"
	ListInstancesSortOrderDesc    ListInstancesSortOrderEnum = "DESC"
	ListInstancesSortOrderUnknown ListInstancesSortOrderEnum = "UNKNOWN"
)

var mappingListInstancesSortOrder = map[string]ListInstancesSortOrderEnum{
	"ASC":     ListInstancesSortOrderAsc,
	"DESC":    ListInstancesSortOrderDesc,
	"UNKNOWN": ListInstancesSortOrderUnknown,
}

// GetListInstancesSortOrderEnumValues Enumerates the set of values for ListInstancesSortOrder
func GetListInstancesSortOrderEnumValues() []ListInstancesSortOrderEnum {
	values := make([]ListInstancesSortOrderEnum, 0)
	for _, v := range mappingListInstancesSortOrder {
		if v != ListInstancesSortOrderUnknown {
			values = append(values, v)
		}
	}
	return values
}
