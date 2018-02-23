// Copyright (c) 2016, 2018, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

package ffsw

import (
    "github.com/oracle/oci-go-sdk/common"
    "net/http"
)

// ListLockOwnersRequest wrapper for the ListLockOwners operation
type ListLockOwnersRequest struct {
        
 // The OCID of the file system. This feature is currently in preview and may change before public release. Do not use it for production workloads.
        FileSystemId *string `mandatory:"true" contributesTo:"path" name:"fileSystemId"`
        
 // The OCID of a mount target. This feature is currently in preview and may change before public release. Do not use it for production workloads.
        MountTarget *string `mandatory:"false" contributesTo:"query" name:"mountTarget"`
        
 // The maximum number of items to return in a paginated "List" call.
 // Example: `500` 
        Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`
        
 // The value of the `opc-next-page` response header from the previous "List" call. 
        Page *string `mandatory:"false" contributesTo:"query" name:"page"`
}

func (request ListLockOwnersRequest) String() string {
    return common.PointerString(request)
}

// ListLockOwnersResponse wrapper for the ListLockOwners operation
type ListLockOwnersResponse struct {

    // The underlying http response
    RawResponse *http.Response
    
 // The []LockOwner instance
    Items []LockOwner `presentIn:"body"`

    
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

func (response ListLockOwnersResponse) String() string {
    return common.PointerString(response)
}

