// Copyright (c) 2016, 2017, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

package audit

import (
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

// GetConfigurationRequest wrapper for the GetConfiguration operation
type GetConfigurationRequest struct {

	// ID of the root compartment (tenancy)
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request GetConfigurationRequest) String() string {
	return common.PointerString(request)
}

// GetHttpRequest implements the OciRequest interface
func (request GetConfigurationRequest) GetHttpRequest(method, path string) (http.Request, error) {
	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// GetRetryPolicy implements the OciRetryableRequest interface
// => assembles retry policy based on specified options and default behavior
func (request GetConfigurationRequest) GetRetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// GetConfigurationResponse wrapper for the GetConfiguration operation
type GetConfigurationResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The Configuration instance
	Configuration `presentIn:"body"`
}

func (response GetConfigurationResponse) String() string {
	return common.PointerString(response)
}

// GetRawResponse implements the OciResponse interface
func (response GetConfigurationResponse) GetRawResponse() *http.Response {
	return response.RawResponse
}
