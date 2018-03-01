// Copyright (c) 2016, 2017, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

package core

import (
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

// GetBootVolumeAttachmentRequest wrapper for the GetBootVolumeAttachment operation
type GetBootVolumeAttachmentRequest struct {

	// The OCID of the boot volume attachment.
	BootVolumeAttachmentId *string `mandatory:"true" contributesTo:"path" name:"bootVolumeAttachmentId"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request GetBootVolumeAttachmentRequest) String() string {
	return common.PointerString(request)
}

// GetHttpRequest implements the OciRequest interface
func (request GetBootVolumeAttachmentRequest) GetHttpRequest(method, path string) (http.Request, error) {
	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// GetRetryPolicy implements the OciRetryableRequest interface
// => assembles retry policy based on specified options and default behavior
func (request GetBootVolumeAttachmentRequest) GetRetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// GetBootVolumeAttachmentResponse wrapper for the GetBootVolumeAttachment operation
type GetBootVolumeAttachmentResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The BootVolumeAttachment instance
	BootVolumeAttachment `presentIn:"body"`

	// For optimistic concurrency control. See `if-match`.
	Etag *string `presentIn:"header" name:"etag"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	// a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response GetBootVolumeAttachmentResponse) String() string {
	return common.PointerString(response)
}

// GetRawResponse implements the OciResponse interface
func (response GetBootVolumeAttachmentResponse) GetRawResponse() *http.Response {
	return response.RawResponse
}

// GetStatefulEntity implements the OciStatefulResponse interface
func (response GetBootVolumeAttachmentResponse) GetStatefulEntity() common.OciPollable {
	return response.BootVolumeAttachment
}
