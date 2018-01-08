// Copyright (c) 2016, 2017, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

package core

import (
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

// ListSubnetsRequest wrapper for the ListSubnets operation
type ListSubnetsRequest struct {

	// The OCID of the compartment.
	CompartmentID *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// The OCID of the VCN.
	VcnID *string `mandatory:"true" contributesTo:"query" name:"vcnId"`

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
	SortBy ListSubnetsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy"`

	// The sort order to use, either ascending (`ASC`) or descending (`DESC`). The DISPLAYNAME sort order
	// is case sensitive.
	SortOrder ListSubnetsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder"`

	// A filter to only return resources that match the given lifecycle state.  The state value is case-insensitive.
	LifecycleState *string `mandatory:"false" contributesTo:"query" name:"lifecycleState"`
}

func (request ListSubnetsRequest) String() string {
	return common.PointerString(request)
}

// ListSubnetsResponse wrapper for the ListSubnets operation
type ListSubnetsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The []Subnet instance
	Items []Subnet `presentIn:"body"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	// a particular request, please provide the request ID.
	OpcRequestID *string `presentIn:"header" name:"opc-request-id"`
}

func (response ListSubnetsResponse) String() string {
	return common.PointerString(response)
}

type ListSubnetsSortByEnum string

const (
	LIST_SUBNETS_SORT_BY_TIMECREATED ListSubnetsSortByEnum = "TIMECREATED"
	LIST_SUBNETS_SORT_BY_DISPLAYNAME ListSubnetsSortByEnum = "DISPLAYNAME"
	LIST_SUBNETS_SORT_BY_UNKNOWN     ListSubnetsSortByEnum = "UNKNOWN"
)

var mapping_listsubnetssortby = map[string]ListSubnetsSortByEnum{
	"TIMECREATED": LIST_SUBNETS_SORT_BY_TIMECREATED,
	"DISPLAYNAME": LIST_SUBNETS_SORT_BY_DISPLAYNAME,
	"UNKNOWN":     LIST_SUBNETS_SORT_BY_UNKNOWN,
}

func GetListSubnetsSortByEnumValues() []ListSubnetsSortByEnum {
	values := make([]ListSubnetsSortByEnum, 0)
	for _, v := range mapping_listsubnetssortby {
		if v != LIST_SUBNETS_SORT_BY_UNKNOWN {
			values = append(values, v)
		}
	}
	return values
}

type ListSubnetsSortOrderEnum string

const (
	LIST_SUBNETS_SORT_ORDER_ASC     ListSubnetsSortOrderEnum = "ASC"
	LIST_SUBNETS_SORT_ORDER_DESC    ListSubnetsSortOrderEnum = "DESC"
	LIST_SUBNETS_SORT_ORDER_UNKNOWN ListSubnetsSortOrderEnum = "UNKNOWN"
)

var mapping_listsubnetssortorder = map[string]ListSubnetsSortOrderEnum{
	"ASC":     LIST_SUBNETS_SORT_ORDER_ASC,
	"DESC":    LIST_SUBNETS_SORT_ORDER_DESC,
	"UNKNOWN": LIST_SUBNETS_SORT_ORDER_UNKNOWN,
}

func GetListSubnetsSortOrderEnumValues() []ListSubnetsSortOrderEnum {
	values := make([]ListSubnetsSortOrderEnum, 0)
	for _, v := range mapping_listsubnetssortorder {
		if v != LIST_SUBNETS_SORT_ORDER_UNKNOWN {
			values = append(values, v)
		}
	}
	return values
}
