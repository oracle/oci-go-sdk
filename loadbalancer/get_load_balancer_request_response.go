// Copyright (c) 2016, 2017, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

package loadbalancer

import (
	"bitbucket.aka.lgl.grungy.us/golang-sdk2/common"
	"net/http"
)

// GetLoadBalancerRequest wrapper for the GetLoadBalancer operation
type GetLoadBalancerRequest struct {

	// The [OCID]({{DOC_SERVER_URL}}/Content/General/Concepts/identifiers.htm) of the load balancer to retrieve.
	LoadBalancerID *string `mandatory:"true" contributesTo:"path" name:"loadBalancerId"`

	// The unique Oracle-assigned identifier for the request. If you need to contact Oracle about a
	// particular request, please provide the request ID.
	OpcRequestID *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`
}

func (request GetLoadBalancerRequest) String() string {
	return common.PointerString(request)
}

// GetLoadBalancerResponse wrapper for the GetLoadBalancer operation
type GetLoadBalancerResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The LoadBalancer instance
	LoadBalancer `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	// a particular request, please provide the request ID.
	OpcRequestID *string `presentIn:"header" name:"opc-request-id"`
}

func (response GetLoadBalancerResponse) String() string {
	return common.PointerString(response)
}
