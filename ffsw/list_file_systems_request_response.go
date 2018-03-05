// Copyright (c) 2016, 2018, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

package ffsw

import (
    "github.com/oracle/oci-go-sdk/common"
    "net/http"
)

// ListFileSystemsRequest wrapper for the ListFileSystems operation
type ListFileSystemsRequest struct {
        
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
        LifecycleState ListFileSystemsLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`
        
 // Filter results by OCID. Must be an OCID of the correct type for
 // the resouce type. 
        Id *string `mandatory:"false" contributesTo:"query" name:"id"`
        
 // The field to sort by.  Only one sort order may be provided.
 // Time created is default ordered as descending.  Display name,
 // path, and name is default ordered as ascending. 
        SortBy ListFileSystemsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`
        
 // The sort order to use, either 'asc' or 'desc'. 
        SortOrder ListFileSystemsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`
}

func (request ListFileSystemsRequest) String() string {
    return common.PointerString(request)
}

// ListFileSystemsResponse wrapper for the ListFileSystems operation
type ListFileSystemsResponse struct {

    // The underlying http response
    RawResponse *http.Response
    
 // The []FileSystemSummary instance
    Items []FileSystemSummary `presentIn:"body"`

    
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

func (response ListFileSystemsResponse) String() string {
    return common.PointerString(response)
}

// ListFileSystemsLifecycleStateEnum Enum with underlying type: string
type ListFileSystemsLifecycleStateEnum string

// Set of constants representing the allowable values for ListFileSystemsLifecycleState
const (
    ListFileSystemsLifecycleStateCreating ListFileSystemsLifecycleStateEnum = "CREATING"
    ListFileSystemsLifecycleStateActive ListFileSystemsLifecycleStateEnum = "ACTIVE"
    ListFileSystemsLifecycleStateDeleting ListFileSystemsLifecycleStateEnum = "DELETING"
    ListFileSystemsLifecycleStateDeleted ListFileSystemsLifecycleStateEnum = "DELETED"
    ListFileSystemsLifecycleStateFailed ListFileSystemsLifecycleStateEnum = "FAILED"
)

var mappingListFileSystemsLifecycleState = map[string]ListFileSystemsLifecycleStateEnum { 
    "CREATING": ListFileSystemsLifecycleStateCreating,
    "ACTIVE": ListFileSystemsLifecycleStateActive,
    "DELETING": ListFileSystemsLifecycleStateDeleting,
    "DELETED": ListFileSystemsLifecycleStateDeleted,
    "FAILED": ListFileSystemsLifecycleStateFailed,
}

// GetListFileSystemsLifecycleStateEnumValues Enumerates the set of values for ListFileSystemsLifecycleState
func GetListFileSystemsLifecycleStateEnumValues() []ListFileSystemsLifecycleStateEnum {
   values := make([]ListFileSystemsLifecycleStateEnum, 0)
   for _, v := range mappingListFileSystemsLifecycleState {
       values = append(values, v)
   }
   return values
}

// ListFileSystemsSortByEnum Enum with underlying type: string
type ListFileSystemsSortByEnum string

// Set of constants representing the allowable values for ListFileSystemsSortBy
const (
    ListFileSystemsSortByTimecreated ListFileSystemsSortByEnum = "TIMECREATED"
    ListFileSystemsSortByDisplayname ListFileSystemsSortByEnum = "DISPLAYNAME"
    ListFileSystemsSortByName ListFileSystemsSortByEnum = "NAME"
    ListFileSystemsSortByPath ListFileSystemsSortByEnum = "PATH"
)

var mappingListFileSystemsSortBy = map[string]ListFileSystemsSortByEnum { 
    "TIMECREATED": ListFileSystemsSortByTimecreated,
    "DISPLAYNAME": ListFileSystemsSortByDisplayname,
    "NAME": ListFileSystemsSortByName,
    "PATH": ListFileSystemsSortByPath,
}

// GetListFileSystemsSortByEnumValues Enumerates the set of values for ListFileSystemsSortBy
func GetListFileSystemsSortByEnumValues() []ListFileSystemsSortByEnum {
   values := make([]ListFileSystemsSortByEnum, 0)
   for _, v := range mappingListFileSystemsSortBy {
       values = append(values, v)
   }
   return values
}

// ListFileSystemsSortOrderEnum Enum with underlying type: string
type ListFileSystemsSortOrderEnum string

// Set of constants representing the allowable values for ListFileSystemsSortOrder
const (
    ListFileSystemsSortOrderAsc ListFileSystemsSortOrderEnum = "ASC"
    ListFileSystemsSortOrderDesc ListFileSystemsSortOrderEnum = "DESC"
)

var mappingListFileSystemsSortOrder = map[string]ListFileSystemsSortOrderEnum { 
    "ASC": ListFileSystemsSortOrderAsc,
    "DESC": ListFileSystemsSortOrderDesc,
}

// GetListFileSystemsSortOrderEnumValues Enumerates the set of values for ListFileSystemsSortOrder
func GetListFileSystemsSortOrderEnumValues() []ListFileSystemsSortOrderEnum {
   values := make([]ListFileSystemsSortOrderEnum, 0)
   for _, v := range mappingListFileSystemsSortOrder {
       values = append(values, v)
   }
   return values
}

