// Copyright (c) 2016, 2017, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

package core

import (
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

// ListInternetGatewaysRequest wrapper for the ListInternetGateways operation
type ListInternetGatewaysRequest struct {

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
	SortBy ListInternetGatewaysSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy"`

	// The sort order to use, either ascending (`ASC`) or descending (`DESC`). The DISPLAYNAME sort order
	// is case sensitive.
	SortOrder ListInternetGatewaysSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder"`

	// A filter to only return resources that match the given lifecycle state.  The state value is case-insensitive.
	LifecycleState *string `mandatory:"false" contributesTo:"query" name:"lifecycleState"`
}

func (request ListInternetGatewaysRequest) String() string {
	return common.PointerString(request)
}

// ListInternetGatewaysResponse wrapper for the ListInternetGateways operation
type ListInternetGatewaysResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The []InternetGateway instance
	Items []InternetGateway `presentIn:"body"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	// a particular request, please provide the request ID.
	OpcRequestID *string `presentIn:"header" name:"opc-request-id"`
}

func (response ListInternetGatewaysResponse) String() string {
	return common.PointerString(response)
}

type ListInternetGatewaysSortByEnum string

const (
	LIST_INTERNET_GATEWAYS_SORT_BY_TIMECREATED ListInternetGatewaysSortByEnum = "TIMECREATED"
	LIST_INTERNET_GATEWAYS_SORT_BY_DISPLAYNAME ListInternetGatewaysSortByEnum = "DISPLAYNAME"
	LIST_INTERNET_GATEWAYS_SORT_BY_UNKNOWN     ListInternetGatewaysSortByEnum = "UNKNOWN"
)

var mapping_listinternetgatewayssortby = map[string]ListInternetGatewaysSortByEnum{
	"TIMECREATED": LIST_INTERNET_GATEWAYS_SORT_BY_TIMECREATED,
	"DISPLAYNAME": LIST_INTERNET_GATEWAYS_SORT_BY_DISPLAYNAME,
	"UNKNOWN":     LIST_INTERNET_GATEWAYS_SORT_BY_UNKNOWN,
}

func GetListInternetGatewaysSortByEnumValues() []ListInternetGatewaysSortByEnum {
	values := make([]ListInternetGatewaysSortByEnum, 0)
	for _, v := range mapping_listinternetgatewayssortby {
		if v != LIST_INTERNET_GATEWAYS_SORT_BY_UNKNOWN {
			values = append(values, v)
		}
	}
	return values
}

type ListInternetGatewaysSortOrderEnum string

const (
	LIST_INTERNET_GATEWAYS_SORT_ORDER_ASC     ListInternetGatewaysSortOrderEnum = "ASC"
	LIST_INTERNET_GATEWAYS_SORT_ORDER_DESC    ListInternetGatewaysSortOrderEnum = "DESC"
	LIST_INTERNET_GATEWAYS_SORT_ORDER_UNKNOWN ListInternetGatewaysSortOrderEnum = "UNKNOWN"
)

var mapping_listinternetgatewayssortorder = map[string]ListInternetGatewaysSortOrderEnum{
	"ASC":     LIST_INTERNET_GATEWAYS_SORT_ORDER_ASC,
	"DESC":    LIST_INTERNET_GATEWAYS_SORT_ORDER_DESC,
	"UNKNOWN": LIST_INTERNET_GATEWAYS_SORT_ORDER_UNKNOWN,
}

func GetListInternetGatewaysSortOrderEnumValues() []ListInternetGatewaysSortOrderEnum {
	values := make([]ListInternetGatewaysSortOrderEnum, 0)
	for _, v := range mapping_listinternetgatewayssortorder {
		if v != LIST_INTERNET_GATEWAYS_SORT_ORDER_UNKNOWN {
			values = append(values, v)
		}
	}
	return values
}
