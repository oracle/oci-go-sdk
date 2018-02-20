// Copyright (c) 2016, 2017, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

package loadbalancer

import (
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

// GetLoadBalancerRequest wrapper for the GetLoadBalancer operation
type GetLoadBalancerRequest struct {

	// The [OCID]({{DOC_SERVER_URL}}/Content/General/Concepts/identifiers.htm) of the load balancer to retrieve.
	LoadBalancerId *string `mandatory:"true" contributesTo:"path" name:"loadBalancerId"`

	// The unique Oracle-assigned identifier for the request. If you need to contact Oracle about a
	// particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`
}

func (request GetLoadBalancerRequest) String() string {
	return common.PointerString(request)
}

// GetHttpRequest implements the OciRequest interface
func (request GetLoadBalancerRequest) GetHttpRequest(method, path string) (http.Request, error) {
	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// GetRetryPolicy implements the OciRetryableRequest interface
// => assembles retry policy based on specified options and default behavior
func (request GetLoadBalancerRequest) GetRetryPolicy(options ...common.RetryPolicyOption) common.RetryPolicy {
	if len(options) == 0 {
		return common.NoRetryPolicy()
	}
	return common.BuildRetryPolicy(options...)
}

// GetLoadBalancerResponse wrapper for the GetLoadBalancer operation
type GetLoadBalancerResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The LoadBalancer instance
	LoadBalancer `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	// a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response GetLoadBalancerResponse) String() string {
	return common.PointerString(response)
}

// GetRawResponse implements the OciResponse interface
func (response GetLoadBalancerResponse) GetRawResponse() *http.Response {
	return response.RawResponse
}

// GetStatefulEntity implements the OciStatefulResponse interface
func (response GetLoadBalancerResponse) GetStatefulEntity() common.OciPollable {
	return response.LoadBalancer
}
