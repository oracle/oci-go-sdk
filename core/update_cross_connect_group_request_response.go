// Copyright (c) 2016, 2017, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

package core

import (
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

// UpdateCrossConnectGroupRequest wrapper for the UpdateCrossConnectGroup operation
type UpdateCrossConnectGroupRequest struct {

	// The OCID of the cross-connect group.
	CrossConnectGroupId *string `mandatory:"true" contributesTo:"path" name:"crossConnectGroupId"`

	// Update CrossConnectGroup fields
	UpdateCrossConnectGroupDetails `contributesTo:"body"`

	// For optimistic concurrency control. In the PUT or DELETE call for a resource, set the `if-match`
	// parameter to the value of the etag from a previous GET or POST response for that resource.  The resource
	// will be updated or deleted only if the etag you provide matches the resource's current etag value.
	IfMatch *string `mandatory:"false" contributesTo:"header" name:"if-match"`
}

func (request UpdateCrossConnectGroupRequest) String() string {
	return common.PointerString(request)
}

// GetHttpRequest implements the OciRequest interface
func (request UpdateCrossConnectGroupRequest) GetHttpRequest(method, path string) (http.Request, error) {
	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// GetRetryPolicy implements the OciRetryableRequest interface
// => assembles retry policy based on specified options and default behavior
func (request UpdateCrossConnectGroupRequest) GetRetryPolicy(options ...common.RetryPolicyOption) common.RetryPolicy {
	if len(options) == 0 {
		return common.NoRetryPolicy()
	}
	return common.BuildRetryPolicy(options...)
}

// UpdateCrossConnectGroupResponse wrapper for the UpdateCrossConnectGroup operation
type UpdateCrossConnectGroupResponse struct {

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

func (response UpdateCrossConnectGroupResponse) String() string {
	return common.PointerString(response)
}

// GetRawResponse implements the OciResponse interface
func (response UpdateCrossConnectGroupResponse) GetRawResponse() *http.Response {
	return response.RawResponse
}

// GetStatefulEntity implements the OciStatefulResponse interface
func (response UpdateCrossConnectGroupResponse) GetStatefulEntity() common.OciPollable {
	return response.CrossConnectGroup
}
