// Copyright (c) 2016, 2017, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

package core

import (
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

// GetCrossConnectGroupRequest wrapper for the GetCrossConnectGroup operation
type GetCrossConnectGroupRequest struct {

	// The OCID of the cross-connect group.
	CrossConnectGroupId *string `mandatory:"true" contributesTo:"path" name:"crossConnectGroupId"`
}

func (request GetCrossConnectGroupRequest) String() string {
	return common.PointerString(request)
}

// GetHttpRequest implements the OciRequest interface
func (request GetCrossConnectGroupRequest) GetHttpRequest(method, path string) (http.Request, error) {
	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// GetRetryPolicy implements the OciRetryableRequest interface
// => assembles retry policy based on specified options and default behavior
func (request GetCrossConnectGroupRequest) GetRetryPolicy(options ...common.RetryPolicyOption) common.RetryPolicy {
	if len(options) == 0 {
		return common.NoRetryPolicy()
	}
	return common.BuildRetryPolicy(options...)
}

// GetCrossConnectGroupResponse wrapper for the GetCrossConnectGroup operation
type GetCrossConnectGroupResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The CrossConnectGroup instance
	CrossConnectGroup `presentIn:"body"`

	// For optimistic concurrency control. See `if-match`.
	Etag *string `presentIn:"header" name:"etag"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	// a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response GetCrossConnectGroupResponse) String() string {
	return common.PointerString(response)
}

// GetRawResponse implements the OciResponse interface
func (response GetCrossConnectGroupResponse) GetRawResponse() *http.Response {
	return response.RawResponse
}

// GetStatefulEntity implements the OciStatefulResponse interface
func (response GetCrossConnectGroupResponse) GetStatefulEntity() common.OciPollable {
	return response.CrossConnectGroup
}
