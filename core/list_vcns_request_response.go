// Copyright (c) 2016, 2017, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

package core

import (
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

// ListVcnsRequest wrapper for the ListVcns operation
type ListVcnsRequest struct {

	// The OCID of the compartment.
	CompartmentID *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// The maximum number of items to return in a paginated "List" call.
	// Example: `500`
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The value of the `opc-next-page` response header from the previous "List" call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// A filter to return only resources that match the given display name exactly.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// The field to sort by. You can provide one sort order (`sortOrder`). Default order for
	// TIMECREATED is descending. Default order for DISPLAYNAME is ascending. The DISPLAYNAME
	// sort order is case sensitive.
	// **Note:** In general, some "List" operations (for example, `ListInstances`) let you
	// optionally filter by Availability Domain if the scope of the resource type is within a
	// single Availability Domain. If you call one of these "List" operations without specifying
	// an Availability Domain, the resources are grouped by Availability Domain, then sorted.
	SortBy ListVcnsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy"`

	// The sort order to use, either ascending (`ASC`) or descending (`DESC`). The DISPLAYNAME sort order
	// is case sensitive.
	SortOrder ListVcnsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder"`

	// A filter to only return resources that match the given lifecycle state.  The state value is case-insensitive.
	LifecycleState *string `mandatory:"false" contributesTo:"query" name:"lifecycleState"`
}

func (request ListVcnsRequest) String() string {
	return common.PointerString(request)
}

// ListVcnsResponse wrapper for the ListVcns operation
type ListVcnsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The []Vcn instance
	Items []Vcn `presentIn:"body"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	// a particular request, please provide the request ID.
	OpcRequestID *string `presentIn:"header" name:"opc-request-id"`
}

func (response ListVcnsResponse) String() string {
	return common.PointerString(response)
}

type ListVcnsSortByEnum string

const (
	LIST_VCNS_SORT_BY_TIMECREATED ListVcnsSortByEnum = "TIMECREATED"
	LIST_VCNS_SORT_BY_DISPLAYNAME ListVcnsSortByEnum = "DISPLAYNAME"
	LIST_VCNS_SORT_BY_UNKNOWN     ListVcnsSortByEnum = "UNKNOWN"
)

var mapping_listvcnssortby = map[string]ListVcnsSortByEnum{
	"TIMECREATED": LIST_VCNS_SORT_BY_TIMECREATED,
	"DISPLAYNAME": LIST_VCNS_SORT_BY_DISPLAYNAME,
	"UNKNOWN":     LIST_VCNS_SORT_BY_UNKNOWN,
}

func GetListVcnsSortByEnumValues() []ListVcnsSortByEnum {
	values := make([]ListVcnsSortByEnum, 0)
	for _, v := range mapping_listvcnssortby {
		if v != LIST_VCNS_SORT_BY_UNKNOWN {
			values = append(values, v)
		}
	}
	return values
}

type ListVcnsSortOrderEnum string

const (
	LIST_VCNS_SORT_ORDER_ASC     ListVcnsSortOrderEnum = "ASC"
	LIST_VCNS_SORT_ORDER_DESC    ListVcnsSortOrderEnum = "DESC"
	LIST_VCNS_SORT_ORDER_UNKNOWN ListVcnsSortOrderEnum = "UNKNOWN"
)

var mapping_listvcnssortorder = map[string]ListVcnsSortOrderEnum{
	"ASC":     LIST_VCNS_SORT_ORDER_ASC,
	"DESC":    LIST_VCNS_SORT_ORDER_DESC,
	"UNKNOWN": LIST_VCNS_SORT_ORDER_UNKNOWN,
}

func GetListVcnsSortOrderEnumValues() []ListVcnsSortOrderEnum {
	values := make([]ListVcnsSortOrderEnum, 0)
	for _, v := range mapping_listvcnssortorder {
		if v != LIST_VCNS_SORT_ORDER_UNKNOWN {
			values = append(values, v)
		}
	}
	return values
}
