// Copyright (c) 2016, 2018, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

package loadbalancer

import (
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

// UpdateLoadBalancerRequest wrapper for the UpdateLoadBalancer operation
type UpdateLoadBalancerRequest struct {

	// The details for updating a load balancer's configuration.
	UpdateLoadBalancerDetails `contributesTo:"body"`

	// The OCID https://docs.us-phoenix-1.oraclecloud.com//Content/General/Concepts/identifiers.htm of the load balancer to update.
	LoadBalancerId *string `mandatory:"true" contributesTo:"path" name:"loadBalancerId"`

	// The unique Oracle-assigned identifier for the request. If you need to contact Oracle about a
	// particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// A token that uniquely identifies a request so it can be retried in case of a timeout or
	// server error without risk of executing that same action again. Retry tokens expire after 24
	// hours, but can be invalidated before then due to conflicting operations (e.g., if a resource
	// has been deleted and purged from the system, then a retry of the original creation request
	// may be rejected).
	OpcRetryToken *string `mandatory:"false" contributesTo:"header" name:"opc-retry-token"`
}

func (request UpdateLoadBalancerRequest) String() string {
	return common.PointerString(request)
}

// UpdateLoadBalancerResponse wrapper for the UpdateLoadBalancer operation
type UpdateLoadBalancerResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	// a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// The OCID https://docs.us-phoenix-1.oraclecloud.com//Content/General/Concepts/identifiers.htm of the work request.
	OpcWorkRequestId *string `presentIn:"header" name:"opc-work-request-id"`
}

func (response UpdateLoadBalancerResponse) String() string {
	return common.PointerString(response)
}
