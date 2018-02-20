// Copyright (c) 2016, 2017, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

package core

import (
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

// GetIPSecConnectionDeviceConfigRequest wrapper for the GetIPSecConnectionDeviceConfig operation
type GetIPSecConnectionDeviceConfigRequest struct {

	// The OCID of the IPSec connection.
	IpscId *string `mandatory:"true" contributesTo:"path" name:"ipscId"`
}

func (request GetIPSecConnectionDeviceConfigRequest) String() string {
	return common.PointerString(request)
}

// GetHttpRequest implements the OciRequest interface
func (request GetIPSecConnectionDeviceConfigRequest) GetHttpRequest(method, path string) (http.Request, error) {
	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// GetRetryPolicy implements the OciRetryableRequest interface
// => assembles retry policy based on specified options and default behavior
func (request GetIPSecConnectionDeviceConfigRequest) GetRetryPolicy(options ...common.RetryPolicyOption) common.RetryPolicy {
	if len(options) == 0 {
		return common.NoRetryPolicy()
	}
	return common.BuildRetryPolicy(options...)
}

// GetIPSecConnectionDeviceConfigResponse wrapper for the GetIPSecConnectionDeviceConfig operation
type GetIPSecConnectionDeviceConfigResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The IpSecConnectionDeviceConfig instance
	IpSecConnectionDeviceConfig `presentIn:"body"`

	// For optimistic concurrency control. See `if-match`.
	Etag *string `presentIn:"header" name:"etag"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	// a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response GetIPSecConnectionDeviceConfigResponse) String() string {
	return common.PointerString(response)
}

// GetRawResponse implements the OciResponse interface
func (response GetIPSecConnectionDeviceConfigResponse) GetRawResponse() *http.Response {
	return response.RawResponse
}
