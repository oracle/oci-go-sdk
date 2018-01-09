// Copyright (c) 2016, 2017, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

package loadbalancer

import (
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

// ListLoadBalancersRequest wrapper for the ListLoadBalancers operation
type ListLoadBalancersRequest struct {

	// The [OCID]({{DOC_SERVER_URL}}/Content/General/Concepts/identifiers.htm) of the compartment containing the load balancers to list.
	CompartmentID *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// The unique Oracle-assigned identifier for the request. If you need to contact Oracle about a
	// particular request, please provide the request ID.
	OpcRequestID *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// The maximum number of items to return in a paginated "List" call.
	// Example: `500`
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The value of the `opc-next-page` response header from the previous "List" call.
	// Example: `3`
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The level of detail to return for each result. Can be `full` or `simple`.
	// Example: `full`
	Detail *string `mandatory:"false" contributesTo:"query" name:"detail"`

	// The field to sort by.  Only one sort order may be provided.  Time created is default ordered as descending.  Display name is default ordered as ascending.
	SortBy ListLoadBalancersSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy"`

	// The sort order to use, either 'asc' or 'desc'
	SortOrder ListLoadBalancersSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder"`

	// A filter to only return resources that match the given display name exactly.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// A filter to only return resources that match the given lifecycle state.
	LifecycleState *string `mandatory:"false" contributesTo:"query" name:"lifecycleState"`
}

func (request ListLoadBalancersRequest) String() string {
	return common.PointerString(request)
}

// ListLoadBalancersResponse wrapper for the ListLoadBalancers operation
type ListLoadBalancersResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The []LoadBalancer instance
	Items []LoadBalancer `presentIn:"body"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	// a particular request, please provide the request ID.
	OpcRequestID *string `presentIn:"header" name:"opc-request-id"`
}

func (response ListLoadBalancersResponse) String() string {
	return common.PointerString(response)
}

type ListLoadBalancersSortByEnum string

const (
	LIST_LOAD_BALANCERS_SORT_BY_TIMECREATED ListLoadBalancersSortByEnum = "TIMECREATED"
	LIST_LOAD_BALANCERS_SORT_BY_DISPLAYNAME ListLoadBalancersSortByEnum = "DISPLAYNAME"
	LIST_LOAD_BALANCERS_SORT_BY_UNKNOWN     ListLoadBalancersSortByEnum = "UNKNOWN"
)

var mapping_listloadbalancerssortby = map[string]ListLoadBalancersSortByEnum{
	"TIMECREATED": LIST_LOAD_BALANCERS_SORT_BY_TIMECREATED,
	"DISPLAYNAME": LIST_LOAD_BALANCERS_SORT_BY_DISPLAYNAME,
	"UNKNOWN":     LIST_LOAD_BALANCERS_SORT_BY_UNKNOWN,
}

func GetListLoadBalancersSortByEnumValues() []ListLoadBalancersSortByEnum {
	values := make([]ListLoadBalancersSortByEnum, 0)
	for _, v := range mapping_listloadbalancerssortby {
		if v != LIST_LOAD_BALANCERS_SORT_BY_UNKNOWN {
			values = append(values, v)
		}
	}
	return values
}

type ListLoadBalancersSortOrderEnum string

const (
	LIST_LOAD_BALANCERS_SORT_ORDER_ASC     ListLoadBalancersSortOrderEnum = "ASC"
	LIST_LOAD_BALANCERS_SORT_ORDER_DESC    ListLoadBalancersSortOrderEnum = "DESC"
	LIST_LOAD_BALANCERS_SORT_ORDER_UNKNOWN ListLoadBalancersSortOrderEnum = "UNKNOWN"
)

var mapping_listloadbalancerssortorder = map[string]ListLoadBalancersSortOrderEnum{
	"ASC":     LIST_LOAD_BALANCERS_SORT_ORDER_ASC,
	"DESC":    LIST_LOAD_BALANCERS_SORT_ORDER_DESC,
	"UNKNOWN": LIST_LOAD_BALANCERS_SORT_ORDER_UNKNOWN,
}

func GetListLoadBalancersSortOrderEnumValues() []ListLoadBalancersSortOrderEnum {
	values := make([]ListLoadBalancersSortOrderEnum, 0)
	for _, v := range mapping_listloadbalancerssortorder {
		if v != LIST_LOAD_BALANCERS_SORT_ORDER_UNKNOWN {
			values = append(values, v)
		}
	}
	return values
}
