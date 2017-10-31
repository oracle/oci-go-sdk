// Copyright (c) 2016, 2017, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

package identity

import "net/http"

// ListIdentityProvidersRequest wrapper for the ListIdentityProviders operation
type ListIdentityProvidersRequest struct {

	// The protocol used for federation.
	Protocol *string `mandatory:"true" contributesTo:"query" name:"protocol"`

	// The OCID of the compartment (remember that the tenancy is simply the root compartment).
	CompartmentID *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// The value of the `opc-next-page` response header from the previous "List" call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The maximum number of items to return in a paginated "List" call.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`
}

// ListIdentityProvidersResponse wrapper for the ListIdentityProviders operation
type ListIdentityProvidersResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The []IdentityProvider instance
	Items []IdentityProvider `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a
	// particular request, please provide the request ID.
	OpcRequestID *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}
