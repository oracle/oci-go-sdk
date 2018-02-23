// Copyright (c) 2016, 2018, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

package ffsw

import (
    "github.com/oracle/oci-go-sdk/common"
    "net/http"
)

// ListExportsRequest wrapper for the ListExports operation
type ListExportsRequest struct {
        
 // The OCID of the compartment. 
        CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`
        
 // The maximum number of items to return in a paginated "List" call.
 // Example: `500` 
        Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`
        
 // The value of the `opc-next-page` response header from the previous "List" call. 
        Page *string `mandatory:"false" contributesTo:"query" name:"page"`
        
 // The OCID of the export set. This feature is currently in preview and may change before public release. Do not use it for production workloads.
        ExportSetId *string `mandatory:"false" contributesTo:"query" name:"exportSetId"`
        
 // The OCID of the export set. This feature is currently in preview and may change before public release. Do not use it for production workloads.
        ExportSet *string `mandatory:"false" contributesTo:"query" name:"exportSet"`
        
 // The OCID of the file system. This feature is currently in preview and may change before public release. Do not use it for production workloads.
        FileSystemId *string `mandatory:"false" contributesTo:"query" name:"fileSystemId"`
        
 // Filter results by the specified lifecycle state. Must be a valid
 // state for the resource type. 
        LifecycleState ListExportsLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`
        
 // Filter results by OCID. Must be an OCID of the correct type for
 // the resouce type. 
        Id *string `mandatory:"false" contributesTo:"query" name:"id"`
        
 // The field to sort by.  Only one sort order may be provided.
 // Time created is default ordered as descending.  Display name,
 // path, and name is default ordered as ascending. 
        SortBy ListExportsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`
        
 // The sort order to use, either 'asc' or 'desc'. 
        SortOrder ListExportsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`
}

func (request ListExportsRequest) String() string {
    return common.PointerString(request)
}

// ListExportsResponse wrapper for the ListExports operation
type ListExportsResponse struct {

    // The underlying http response
    RawResponse *http.Response
    
 // The []ExportSummary instance
    Items []ExportSummary `presentIn:"body"`

    
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

func (response ListExportsResponse) String() string {
    return common.PointerString(response)
}

// ListExportsLifecycleStateEnum Enum with underlying type: string
type ListExportsLifecycleStateEnum string

// Set of constants representing the allowable values for ListExportsLifecycleState
const (
    ListExportsLifecycleStateCreating ListExportsLifecycleStateEnum = "CREATING"
    ListExportsLifecycleStateActive ListExportsLifecycleStateEnum = "ACTIVE"
    ListExportsLifecycleStateDeleting ListExportsLifecycleStateEnum = "DELETING"
    ListExportsLifecycleStateDeleted ListExportsLifecycleStateEnum = "DELETED"
    ListExportsLifecycleStateFailed ListExportsLifecycleStateEnum = "FAILED"
)

var mappingListExportsLifecycleState = map[string]ListExportsLifecycleStateEnum { 
    "CREATING": ListExportsLifecycleStateCreating,
    "ACTIVE": ListExportsLifecycleStateActive,
    "DELETING": ListExportsLifecycleStateDeleting,
    "DELETED": ListExportsLifecycleStateDeleted,
    "FAILED": ListExportsLifecycleStateFailed,
}

// GetListExportsLifecycleStateEnumValues Enumerates the set of values for ListExportsLifecycleState
func GetListExportsLifecycleStateEnumValues() []ListExportsLifecycleStateEnum {
   values := make([]ListExportsLifecycleStateEnum, 0)
   for _, v := range mappingListExportsLifecycleState {
       values = append(values, v)
   }
   return values
}

// ListExportsSortByEnum Enum with underlying type: string
type ListExportsSortByEnum string

// Set of constants representing the allowable values for ListExportsSortBy
const (
    ListExportsSortByTimecreated ListExportsSortByEnum = "TIMECREATED"
    ListExportsSortByDisplayname ListExportsSortByEnum = "DISPLAYNAME"
    ListExportsSortByName ListExportsSortByEnum = "NAME"
    ListExportsSortByPath ListExportsSortByEnum = "PATH"
)

var mappingListExportsSortBy = map[string]ListExportsSortByEnum { 
    "TIMECREATED": ListExportsSortByTimecreated,
    "DISPLAYNAME": ListExportsSortByDisplayname,
    "NAME": ListExportsSortByName,
    "PATH": ListExportsSortByPath,
}

// GetListExportsSortByEnumValues Enumerates the set of values for ListExportsSortBy
func GetListExportsSortByEnumValues() []ListExportsSortByEnum {
   values := make([]ListExportsSortByEnum, 0)
   for _, v := range mappingListExportsSortBy {
       values = append(values, v)
   }
   return values
}

// ListExportsSortOrderEnum Enum with underlying type: string
type ListExportsSortOrderEnum string

// Set of constants representing the allowable values for ListExportsSortOrder
const (
    ListExportsSortOrderAsc ListExportsSortOrderEnum = "ASC"
    ListExportsSortOrderDesc ListExportsSortOrderEnum = "DESC"
)

var mappingListExportsSortOrder = map[string]ListExportsSortOrderEnum { 
    "ASC": ListExportsSortOrderAsc,
    "DESC": ListExportsSortOrderDesc,
}

// GetListExportsSortOrderEnumValues Enumerates the set of values for ListExportsSortOrder
func GetListExportsSortOrderEnumValues() []ListExportsSortOrderEnum {
   values := make([]ListExportsSortOrderEnum, 0)
   for _, v := range mappingListExportsSortOrder {
       values = append(values, v)
   }
   return values
}

