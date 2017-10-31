// Copyright (c) 2016, 2017, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

package identity

import "net/http"

// ListRegionSubscriptionsRequest wrapper for the ListRegionSubscriptions operation
type ListRegionSubscriptionsRequest struct {

	// The OCID of the tenancy.
	TenancyID *string `mandatory:"true" contributesTo:"path" name:"tenancyId"`
}

// ListRegionSubscriptionsResponse wrapper for the ListRegionSubscriptions operation
type ListRegionSubscriptionsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The []RegionSubscription instance
	Items []RegionSubscription `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a
	// particular request, please provide the request ID.
	OpcRequestID *string `presentIn:"header" name:"opc-request-id"`
}
