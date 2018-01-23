// Copyright (c) 2016, 2017, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

package loadbalancer

import (
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

// ListBackendSetsRequest wrapper for the ListBackendSets operation
type ListBackendSetsRequest struct {

	// The https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/identifiers.htm of the load balancer associated with the backend sets to retrieve.
	LoadBalancerId *string `mandatory:"true" contributesTo:"path" name:"loadBalancerId"`

	// The unique Oracle-assigned identifier for the request. If you need to contact Oracle about a
	// particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`
}

func (request ListBackendSetsRequest) String() string {
	return common.PointerString(request)
}

// ListBackendSetsResponse wrapper for the ListBackendSets operation
type ListBackendSetsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The []BackendSet instance
	Items []BackendSet `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	// a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response ListBackendSetsResponse) String() string {
	return common.PointerString(response)
}
