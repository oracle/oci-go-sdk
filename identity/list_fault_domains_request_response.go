// Copyright (c) 2016, 2017, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

package identity

import (
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

// ListFaultDomainsRequest wrapper for the ListFaultDomains operation
type ListFaultDomainsRequest struct {

	// The OCID of the compartment (remember that the tenancy is simply the root compartment).
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// The name of the availibilityDomain.
	AvailabilityDomain *string `mandatory:"true" contributesTo:"query" name:"availabilityDomain"`
}

func (request ListFaultDomainsRequest) String() string {
	return common.PointerString(request)
}

// GetHttpRequest implements the OciRequest interface
func (request ListFaultDomainsRequest) GetHttpRequest(method, path string) (http.Request, error) {
	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// GetRetryPolicy implements the OciRetryableRequest interface
// => assembles retry policy based on specified options and default behavior
func (request ListFaultDomainsRequest) GetRetryPolicy(options ...common.RetryPolicyOption) common.RetryPolicy {
	if len(options) == 0 {
		return common.NoRetryPolicy()
	}
	return common.BuildRetryPolicy(options...)
}

// ListFaultDomainsResponse wrapper for the ListFaultDomains operation
type ListFaultDomainsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The []FaultDomain instance
	Items []FaultDomain `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a
	// particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response ListFaultDomainsResponse) String() string {
	return common.PointerString(response)
}

// GetRawResponse implements the OciResponse interface
func (response ListFaultDomainsResponse) GetRawResponse() *http.Response {
	return response.RawResponse
}
