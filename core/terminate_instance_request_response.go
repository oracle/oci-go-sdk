// Copyright (c) 2016, 2017, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

package core

import (
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

// TerminateInstanceRequest wrapper for the TerminateInstance operation
type TerminateInstanceRequest struct {

	// The OCID of the instance.
	InstanceId *string `mandatory:"true" contributesTo:"path" name:"instanceId"`

	// For optimistic concurrency control. In the PUT or DELETE call for a resource, set the `if-match`
	// parameter to the value of the etag from a previous GET or POST response for that resource.  The resource
	// will be updated or deleted only if the etag you provide matches the resource's current etag value.
	IfMatch *string `mandatory:"false" contributesTo:"header" name:"if-match"`

	// Specifies whether to delete or preserve the boot volume when terminating an instance.
	// The default value is false.
	PreserveBootVolume *bool `mandatory:"false" contributesTo:"query" name:"preserveBootVolume"`
}

func (request TerminateInstanceRequest) String() string {
	return common.PointerString(request)
}

// GetHttpRequest implements the OciRequest interface
func (request TerminateInstanceRequest) GetHttpRequest(method, path string) (http.Request, error) {
	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// GetRetryPolicy implements the OciRetryableRequest interface
// => assembles retry policy based on specified options and default behavior
func (request TerminateInstanceRequest) GetRetryPolicy(options ...common.RetryPolicyOption) common.RetryPolicy {
	if len(options) == 0 {
		return common.NoRetryPolicy()
	}
	return common.BuildRetryPolicy(options...)
}

// TerminateInstanceResponse wrapper for the TerminateInstance operation
type TerminateInstanceResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	// a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response TerminateInstanceResponse) String() string {
	return common.PointerString(response)
}

// GetRawResponse implements the OciResponse interface
func (response TerminateInstanceResponse) GetRawResponse() *http.Response {
	return response.RawResponse
}
