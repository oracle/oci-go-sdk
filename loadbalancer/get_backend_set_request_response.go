// Copyright (c) 2016, 2017, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

package loadbalancer

import (
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

// GetBackendSetRequest wrapper for the GetBackendSet operation
type GetBackendSetRequest struct {

	// The [OCID]({{DOC_SERVER_URL}}/Content/General/Concepts/identifiers.htm) of the specified load balancer.
	LoadBalancerId *string `mandatory:"true" contributesTo:"path" name:"loadBalancerId"`

	// The name of the backend set to retrieve.
	// Example: `My_backend_set`
	BackendSetName *string `mandatory:"true" contributesTo:"path" name:"backendSetName"`

	// The unique Oracle-assigned identifier for the request. If you need to contact Oracle about a
	// particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`
}

func (request GetBackendSetRequest) String() string {
	return common.PointerString(request)
}

// GetHttpRequest implements the OciRequest interface
func (request GetBackendSetRequest) GetHttpRequest(method, path string) (http.Request, error) {
	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// GetRetryPolicy implements the OciRetryableRequest interface
// => assembles retry policy based on specified options and default behavior
func (request GetBackendSetRequest) GetRetryPolicy(options ...common.RetryPolicyOption) common.RetryPolicy {
	if len(options) == 0 {
		return common.NoRetryPolicy()
	}
	return common.BuildRetryPolicy(options...)
}

// GetBackendSetResponse wrapper for the GetBackendSet operation
type GetBackendSetResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The BackendSet instance
	BackendSet `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	// a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response GetBackendSetResponse) String() string {
	return common.PointerString(response)
}

// GetRawResponse implements the OciResponse interface
func (response GetBackendSetResponse) GetRawResponse() *http.Response {
	return response.RawResponse
}
