// Copyright (c) 2016, 2017, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

package objectstorage

import (
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

// GetNamespaceRequest wrapper for the GetNamespace operation
type GetNamespaceRequest struct {

	// The client request ID for tracing.
	OpcClientRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-client-request-id"`
}

func (request GetNamespaceRequest) String() string {
	return common.PointerString(request)
}

// GetHttpRequest implements the OciRequest interface
func (request GetNamespaceRequest) GetHttpRequest(method, path string) (http.Request, error) {
	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// GetRetryPolicy implements the OciRetryableRequest interface
// => assembles retry policy based on specified options and default behavior
func (request GetNamespaceRequest) GetRetryPolicy(options ...common.RetryPolicyOption) common.RetryPolicy {
	if len(options) == 0 {
		return common.NoRetryPolicy()
	}
	return common.BuildRetryPolicy(options...)
}

// GetNamespaceResponse wrapper for the GetNamespace operation
type GetNamespaceResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The string instance
	Value *string `presentIn:"body"`
}

func (response GetNamespaceResponse) String() string {
	return common.PointerString(response)
}

// GetRawResponse implements the OciResponse interface
func (response GetNamespaceResponse) GetRawResponse() *http.Response {
	return response.RawResponse
}
