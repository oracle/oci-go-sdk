// Copyright (c) 2016, 2017, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

package core

import (
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

// ListCrossConnectsRequest wrapper for the ListCrossConnects operation
type ListCrossConnectsRequest struct {

	// The OCID of the compartment.
	CompartmentID *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// The OCID of the cross-connect group.
	CrossConnectGroupID *string `mandatory:"false" contributesTo:"query" name:"crossConnectGroupId"`

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
	SortBy ListCrossConnectsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to use, either ascending (`ASC`) or descending (`DESC`). The DISPLAYNAME sort order
	// is case sensitive.
	SortOrder ListCrossConnectsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// A filter to return only resources that match the specified lifecycle state. The value is case insensitive.
	LifecycleState *string `mandatory:"false" contributesTo:"query" name:"lifecycleState"`
}

func (request ListCrossConnectsRequest) String() string {
	return common.PointerString(request)
}

// ListCrossConnectsResponse wrapper for the ListCrossConnects operation
type ListCrossConnectsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The []CrossConnect instance
	Items []CrossConnect `presentIn:"body"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	// a particular request, please provide the request ID.
	OpcRequestID *string `presentIn:"header" name:"opc-request-id"`
}

func (response ListCrossConnectsResponse) String() string {
	return common.PointerString(response)
}

type ListCrossConnectsSortByEnum string

const (
	LIST_CROSS_CONNECTS_SORT_BY_TIMECREATED ListCrossConnectsSortByEnum = "TIMECREATED"
	LIST_CROSS_CONNECTS_SORT_BY_DISPLAYNAME ListCrossConnectsSortByEnum = "DISPLAYNAME"
	LIST_CROSS_CONNECTS_SORT_BY_UNKNOWN     ListCrossConnectsSortByEnum = "UNKNOWN"
)

var mapping_listcrossconnectssortby = map[string]ListCrossConnectsSortByEnum{
	"TIMECREATED": LIST_CROSS_CONNECTS_SORT_BY_TIMECREATED,
	"DISPLAYNAME": LIST_CROSS_CONNECTS_SORT_BY_DISPLAYNAME,
	"UNKNOWN":     LIST_CROSS_CONNECTS_SORT_BY_UNKNOWN,
}

func GetListCrossConnectsSortByEnumValues() []ListCrossConnectsSortByEnum {
	values := make([]ListCrossConnectsSortByEnum, 0)
	for _, v := range mapping_listcrossconnectssortby {
		if v != LIST_CROSS_CONNECTS_SORT_BY_UNKNOWN {
			values = append(values, v)
		}
	}
	return values
}

type ListCrossConnectsSortOrderEnum string

const (
	LIST_CROSS_CONNECTS_SORT_ORDER_ASC     ListCrossConnectsSortOrderEnum = "ASC"
	LIST_CROSS_CONNECTS_SORT_ORDER_DESC    ListCrossConnectsSortOrderEnum = "DESC"
	LIST_CROSS_CONNECTS_SORT_ORDER_UNKNOWN ListCrossConnectsSortOrderEnum = "UNKNOWN"
)

var mapping_listcrossconnectssortorder = map[string]ListCrossConnectsSortOrderEnum{
	"ASC":     LIST_CROSS_CONNECTS_SORT_ORDER_ASC,
	"DESC":    LIST_CROSS_CONNECTS_SORT_ORDER_DESC,
	"UNKNOWN": LIST_CROSS_CONNECTS_SORT_ORDER_UNKNOWN,
}

func GetListCrossConnectsSortOrderEnumValues() []ListCrossConnectsSortOrderEnum {
	values := make([]ListCrossConnectsSortOrderEnum, 0)
	for _, v := range mapping_listcrossconnectssortorder {
		if v != LIST_CROSS_CONNECTS_SORT_ORDER_UNKNOWN {
			values = append(values, v)
		}
	}
	return values
}
