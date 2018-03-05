// Copyright (c) 2016, 2018, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

package ffsw

import (
    "github.com/oracle/oci-go-sdk/common"
    "net/http"
)

// ListExportSetsRequest wrapper for the ListExportSets operation
type ListExportSetsRequest struct {
        
 // The OCID of the compartment. 
        CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`
        
 // The name of the Availability Domain.
 // Example: `Uocm:PHX-AD-1` 
        AvailabilityDomain *string `mandatory:"true" contributesTo:"query" name:"availabilityDomain"`
        
 // The maximum number of items to return in a paginated "List" call.
 // Example: `500` 
        Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`
        
 // The value of the `opc-next-page` response header from the previous "List" call. 
        Page *string `mandatory:"false" contributesTo:"query" name:"page"`
        
 // A user-friendly name. Does not have to be unique, and it is changeable.
 // Example: `My resource` 
        DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`
        
 // Filter results by the specified lifecycle state. Must be a valid
 // state for the resource type. 
        LifecycleState ListExportSetsLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`
        
 // Filter results by OCID. Must be an OCID of the correct type for
 // the resouce type. 
        Id *string `mandatory:"false" contributesTo:"query" name:"id"`
        
 // The field to sort by.  Only one sort order may be provided.
 // Time created is default ordered as descending.  Display name,
 // path, and name is default ordered as ascending. 
        SortBy ListExportSetsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`
        
 // The sort order to use, either 'asc' or 'desc'. 
        SortOrder ListExportSetsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`
}

func (request ListExportSetsRequest) String() string {
    return common.PointerString(request)
}

// ListExportSetsResponse wrapper for the ListExportSets operation
type ListExportSetsResponse struct {

    // The underlying http response
    RawResponse *http.Response
    
 // The []ExportSetSummary instance
    Items []ExportSetSummary `presentIn:"body"`

    
 // For pagination of a list of items. When paging through
 // a list, if this header appears in the response, then a
 // partial list might have been returned. Include this
 // value as the `page` parameter for the subsequent GET
 // request to get the next batch of items.
    OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
    
 // Unique Oracle-assigned identifier for the request. If
 // you need to contact Oracle about a particular request,
 // please provide the request ID.
    OpcRequestId *string `presentIn:"header" name:"opc-request-id"`


}

func (response ListExportSetsResponse) String() string {
    return common.PointerString(response)
}

// ListExportSetsLifecycleStateEnum Enum with underlying type: string
type ListExportSetsLifecycleStateEnum string

// Set of constants representing the allowable values for ListExportSetsLifecycleState
const (
    ListExportSetsLifecycleStateCreating ListExportSetsLifecycleStateEnum = "CREATING"
    ListExportSetsLifecycleStateActive ListExportSetsLifecycleStateEnum = "ACTIVE"
    ListExportSetsLifecycleStateDeleting ListExportSetsLifecycleStateEnum = "DELETING"
    ListExportSetsLifecycleStateDeleted ListExportSetsLifecycleStateEnum = "DELETED"
    ListExportSetsLifecycleStateFailed ListExportSetsLifecycleStateEnum = "FAILED"
)

var mappingListExportSetsLifecycleState = map[string]ListExportSetsLifecycleStateEnum { 
    "CREATING": ListExportSetsLifecycleStateCreating,
    "ACTIVE": ListExportSetsLifecycleStateActive,
    "DELETING": ListExportSetsLifecycleStateDeleting,
    "DELETED": ListExportSetsLifecycleStateDeleted,
    "FAILED": ListExportSetsLifecycleStateFailed,
}

// GetListExportSetsLifecycleStateEnumValues Enumerates the set of values for ListExportSetsLifecycleState
func GetListExportSetsLifecycleStateEnumValues() []ListExportSetsLifecycleStateEnum {
   values := make([]ListExportSetsLifecycleStateEnum, 0)
   for _, v := range mappingListExportSetsLifecycleState {
       values = append(values, v)
   }
   return values
}

// ListExportSetsSortByEnum Enum with underlying type: string
type ListExportSetsSortByEnum string

// Set of constants representing the allowable values for ListExportSetsSortBy
const (
    ListExportSetsSortByTimecreated ListExportSetsSortByEnum = "TIMECREATED"
    ListExportSetsSortByDisplayname ListExportSetsSortByEnum = "DISPLAYNAME"
    ListExportSetsSortByName ListExportSetsSortByEnum = "NAME"
    ListExportSetsSortByPath ListExportSetsSortByEnum = "PATH"
)

var mappingListExportSetsSortBy = map[string]ListExportSetsSortByEnum { 
    "TIMECREATED": ListExportSetsSortByTimecreated,
    "DISPLAYNAME": ListExportSetsSortByDisplayname,
    "NAME": ListExportSetsSortByName,
    "PATH": ListExportSetsSortByPath,
}

// GetListExportSetsSortByEnumValues Enumerates the set of values for ListExportSetsSortBy
func GetListExportSetsSortByEnumValues() []ListExportSetsSortByEnum {
   values := make([]ListExportSetsSortByEnum, 0)
   for _, v := range mappingListExportSetsSortBy {
       values = append(values, v)
   }
   return values
}

// ListExportSetsSortOrderEnum Enum with underlying type: string
type ListExportSetsSortOrderEnum string

// Set of constants representing the allowable values for ListExportSetsSortOrder
const (
    ListExportSetsSortOrderAsc ListExportSetsSortOrderEnum = "ASC"
    ListExportSetsSortOrderDesc ListExportSetsSortOrderEnum = "DESC"
)

var mappingListExportSetsSortOrder = map[string]ListExportSetsSortOrderEnum { 
    "ASC": ListExportSetsSortOrderAsc,
    "DESC": ListExportSetsSortOrderDesc,
}

// GetListExportSetsSortOrderEnumValues Enumerates the set of values for ListExportSetsSortOrder
func GetListExportSetsSortOrderEnumValues() []ListExportSetsSortOrderEnum {
   values := make([]ListExportSetsSortOrderEnum, 0)
   for _, v := range mappingListExportSetsSortOrder {
       values = append(values, v)
   }
   return values
}

