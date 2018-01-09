// Copyright (c) 2016, 2017, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

package core

import (
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

// ListConsoleHistoriesRequest wrapper for the ListConsoleHistories operation
type ListConsoleHistoriesRequest struct {

	// The OCID of the compartment.
	CompartmentID *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// The name of the Availability Domain.
	// Example: `Uocm:PHX-AD-1`
	AvailabilityDomain *string `mandatory:"false" contributesTo:"query" name:"availabilityDomain"`

	// The maximum number of items to return in a paginated "List" call.
	// Example: `500`
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The value of the `opc-next-page` response header from the previous "List" call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The OCID of the instance.
	InstanceID *string `mandatory:"false" contributesTo:"query" name:"instanceId"`

	// The field to sort by. You can provide one sort order (`sortOrder`). Default order for
	// TIMECREATED is descending. Default order for DISPLAYNAME is ascending. The DISPLAYNAME
	// sort order is case sensitive.
	// **Note:** In general, some "List" operations (for example, `ListInstances`) let you
	// optionally filter by Availability Domain if the scope of the resource type is within a
	// single Availability Domain. If you call one of these "List" operations without specifying
	// an Availability Domain, the resources are grouped by Availability Domain, then sorted.
	SortBy ListConsoleHistoriesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to use, either ascending (`ASC`) or descending (`DESC`). The DISPLAYNAME sort order
	// is case sensitive.
	SortOrder ListConsoleHistoriesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// A filter to only return resources that match the given lifecycle state.  The state value is case-insensitive.
	LifecycleState *string `mandatory:"false" contributesTo:"query" name:"lifecycleState"`
}

func (request ListConsoleHistoriesRequest) String() string {
	return common.PointerString(request)
}

// ListConsoleHistoriesResponse wrapper for the ListConsoleHistories operation
type ListConsoleHistoriesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The []ConsoleHistory instance
	Items []ConsoleHistory `presentIn:"body"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	// a particular request, please provide the request ID.
	OpcRequestID *string `presentIn:"header" name:"opc-request-id"`
}

func (response ListConsoleHistoriesResponse) String() string {
	return common.PointerString(response)
}

type ListConsoleHistoriesSortByEnum string

const (
	LIST_CONSOLE_HISTORIES_SORT_BY_TIMECREATED ListConsoleHistoriesSortByEnum = "TIMECREATED"
	LIST_CONSOLE_HISTORIES_SORT_BY_DISPLAYNAME ListConsoleHistoriesSortByEnum = "DISPLAYNAME"
	LIST_CONSOLE_HISTORIES_SORT_BY_UNKNOWN     ListConsoleHistoriesSortByEnum = "UNKNOWN"
)

var mapping_listconsolehistoriessortby = map[string]ListConsoleHistoriesSortByEnum{
	"TIMECREATED": LIST_CONSOLE_HISTORIES_SORT_BY_TIMECREATED,
	"DISPLAYNAME": LIST_CONSOLE_HISTORIES_SORT_BY_DISPLAYNAME,
	"UNKNOWN":     LIST_CONSOLE_HISTORIES_SORT_BY_UNKNOWN,
}

func GetListConsoleHistoriesSortByEnumValues() []ListConsoleHistoriesSortByEnum {
	values := make([]ListConsoleHistoriesSortByEnum, 0)
	for _, v := range mapping_listconsolehistoriessortby {
		if v != LIST_CONSOLE_HISTORIES_SORT_BY_UNKNOWN {
			values = append(values, v)
		}
	}
	return values
}

type ListConsoleHistoriesSortOrderEnum string

const (
	LIST_CONSOLE_HISTORIES_SORT_ORDER_ASC     ListConsoleHistoriesSortOrderEnum = "ASC"
	LIST_CONSOLE_HISTORIES_SORT_ORDER_DESC    ListConsoleHistoriesSortOrderEnum = "DESC"
	LIST_CONSOLE_HISTORIES_SORT_ORDER_UNKNOWN ListConsoleHistoriesSortOrderEnum = "UNKNOWN"
)

var mapping_listconsolehistoriessortorder = map[string]ListConsoleHistoriesSortOrderEnum{
	"ASC":     LIST_CONSOLE_HISTORIES_SORT_ORDER_ASC,
	"DESC":    LIST_CONSOLE_HISTORIES_SORT_ORDER_DESC,
	"UNKNOWN": LIST_CONSOLE_HISTORIES_SORT_ORDER_UNKNOWN,
}

func GetListConsoleHistoriesSortOrderEnumValues() []ListConsoleHistoriesSortOrderEnum {
	values := make([]ListConsoleHistoriesSortOrderEnum, 0)
	for _, v := range mapping_listconsolehistoriessortorder {
		if v != LIST_CONSOLE_HISTORIES_SORT_ORDER_UNKNOWN {
			values = append(values, v)
		}
	}
	return values
}
