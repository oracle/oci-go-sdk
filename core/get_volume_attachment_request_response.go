// Copyright (c) 2016, 2017, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

package core

import (
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

// GetVolumeAttachmentRequest wrapper for the GetVolumeAttachment operation
type GetVolumeAttachmentRequest struct {

	// The OCID of the volume attachment.
	VolumeAttachmentId *string `mandatory:"true" contributesTo:"path" name:"volumeAttachmentId"`
}

func (request GetVolumeAttachmentRequest) String() string {
	return common.PointerString(request)
}

// GetHttpRequest implements the OciRequest interface
func (request GetVolumeAttachmentRequest) GetHttpRequest(method, path string) (http.Request, error) {
	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// GetRetryPolicy implements the OciRetryableRequest interface
// => assembles retry policy based on specified options and default behavior
func (request GetVolumeAttachmentRequest) GetRetryPolicy(options ...common.RetryPolicyOption) common.RetryPolicy {
	if len(options) == 0 {
		return common.NoRetryPolicy()
	}
	return common.BuildRetryPolicy(options...)
}

// GetVolumeAttachmentResponse wrapper for the GetVolumeAttachment operation
type GetVolumeAttachmentResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The VolumeAttachment instance
	VolumeAttachment `presentIn:"body"`

	// For optimistic concurrency control. See `if-match`.
	Etag *string `presentIn:"header" name:"etag"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	// a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response GetVolumeAttachmentResponse) String() string {
	return common.PointerString(response)
}

// GetRawResponse implements the OciResponse interface
func (response GetVolumeAttachmentResponse) GetRawResponse() *http.Response {
	return response.RawResponse
}

// GetStatefulEntity implements the OciStatefulResponse interface
func (response GetVolumeAttachmentResponse) GetStatefulEntity() common.OciPollable {
	return response.VolumeAttachment
}
