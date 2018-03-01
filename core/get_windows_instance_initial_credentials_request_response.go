// Copyright (c) 2016, 2017, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

package core

import (
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

// GetWindowsInstanceInitialCredentialsRequest wrapper for the GetWindowsInstanceInitialCredentials operation
type GetWindowsInstanceInitialCredentialsRequest struct {

	// The OCID of the instance.
	InstanceId *string `mandatory:"true" contributesTo:"path" name:"instanceId"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request GetWindowsInstanceInitialCredentialsRequest) String() string {
	return common.PointerString(request)
}

// GetHttpRequest implements the OciRequest interface
func (request GetWindowsInstanceInitialCredentialsRequest) GetHttpRequest(method, path string) (http.Request, error) {
	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// GetRetryPolicy implements the OciRetryableRequest interface
// => assembles retry policy based on specified options and default behavior
func (request GetWindowsInstanceInitialCredentialsRequest) GetRetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// GetWindowsInstanceInitialCredentialsResponse wrapper for the GetWindowsInstanceInitialCredentials operation
type GetWindowsInstanceInitialCredentialsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The InstanceCredentials instance
	InstanceCredentials `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	// a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response GetWindowsInstanceInitialCredentialsResponse) String() string {
	return common.PointerString(response)
}

// GetRawResponse implements the OciResponse interface
func (response GetWindowsInstanceInitialCredentialsResponse) GetRawResponse() *http.Response {
	return response.RawResponse
}
