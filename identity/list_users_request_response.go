// Copyright (c) 2016, 2017, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

package identity

// ListUsersRequest wrapper for the ListUsers operation
type ListUsersRequest struct {

	// The OCID of the compartment (remember that the tenancy is simply the root compartment).
	CompartmentID string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// The value of the `opc-next-page` response header from the previous "List" call.
	Page string `mandatory:"false" contributesTo:"query" name:"page"`

	// The maximum number of items to return in a paginated "List" call.
	Limit int32 `mandatory:"false" contributesTo:"query" name:"limit"`
}

// ListUsersResponse wrapper for the ListUsers operation
type ListUsersResponse struct {

	// The []User instance
	Items []User `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a
	// particular request, please provide the request ID.
	OpcRequestID string `presentIn:"header" name:"opcrequestid"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage string `presentIn:"header" name:"opcnextpage"`
}
