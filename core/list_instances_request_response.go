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
	CompartmentID *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

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
	LifecycleState *string `mandatory:"false" contributesTo:"query" name:"lifecycleState"`
}

func (request ListInstancesRequest) String() string {
	return common.PointerString(request)
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
	OpcRequestID *string `presentIn:"header" name:"opc-request-id"`
}

func (response ListInstancesResponse) String() string {
	return common.PointerString(response)
}

type ListInstancesSortByEnum string

const (
	LIST_INSTANCES_SORT_BY_TIMECREATED ListInstancesSortByEnum = "TIMECREATED"
	LIST_INSTANCES_SORT_BY_DISPLAYNAME ListInstancesSortByEnum = "DISPLAYNAME"
	LIST_INSTANCES_SORT_BY_UNKNOWN     ListInstancesSortByEnum = "UNKNOWN"
)

var mapping_listinstancessortby = map[string]ListInstancesSortByEnum{
	"TIMECREATED": LIST_INSTANCES_SORT_BY_TIMECREATED,
	"DISPLAYNAME": LIST_INSTANCES_SORT_BY_DISPLAYNAME,
	"UNKNOWN":     LIST_INSTANCES_SORT_BY_UNKNOWN,
}

func GetListInstancesSortByEnumValues() []ListInstancesSortByEnum {
	values := make([]ListInstancesSortByEnum, 0)
	for _, v := range mapping_listinstancessortby {
		if v != LIST_INSTANCES_SORT_BY_UNKNOWN {
			values = append(values, v)
		}
	}
	return values
}

type ListInstancesSortOrderEnum string

const (
	LIST_INSTANCES_SORT_ORDER_ASC     ListInstancesSortOrderEnum = "ASC"
	LIST_INSTANCES_SORT_ORDER_DESC    ListInstancesSortOrderEnum = "DESC"
	LIST_INSTANCES_SORT_ORDER_UNKNOWN ListInstancesSortOrderEnum = "UNKNOWN"
)

var mapping_listinstancessortorder = map[string]ListInstancesSortOrderEnum{
	"ASC":     LIST_INSTANCES_SORT_ORDER_ASC,
	"DESC":    LIST_INSTANCES_SORT_ORDER_DESC,
	"UNKNOWN": LIST_INSTANCES_SORT_ORDER_UNKNOWN,
}

func GetListInstancesSortOrderEnumValues() []ListInstancesSortOrderEnum {
	values := make([]ListInstancesSortOrderEnum, 0)
	for _, v := range mapping_listinstancessortorder {
		if v != LIST_INSTANCES_SORT_ORDER_UNKNOWN {
			values = append(values, v)
		}
	}
	return values
}
