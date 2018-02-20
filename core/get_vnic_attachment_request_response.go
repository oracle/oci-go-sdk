// Copyright (c) 2016, 2017, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

package core

import (
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

// GetVnicAttachmentRequest wrapper for the GetVnicAttachment operation
type GetVnicAttachmentRequest struct {

	// The OCID of the VNIC attachment.
	VnicAttachmentId *string `mandatory:"true" contributesTo:"path" name:"vnicAttachmentId"`
}

func (request GetVnicAttachmentRequest) String() string {
	return common.PointerString(request)
}

// GetHttpRequest implements the OciRequest interface
func (request GetVnicAttachmentRequest) GetHttpRequest(method, path string) (http.Request, error) {
	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// GetRetryPolicy implements the OciRetryableRequest interface
// => assembles retry policy based on specified options and default behavior
func (request GetVnicAttachmentRequest) GetRetryPolicy(options ...common.RetryPolicyOption) common.RetryPolicy {
	if len(options) == 0 {
		return common.NoRetryPolicy()
	}
	return common.BuildRetryPolicy(options...)
}

// GetVnicAttachmentResponse wrapper for the GetVnicAttachment operation
type GetVnicAttachmentResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The VnicAttachment instance
	VnicAttachment `presentIn:"body"`

	// For optimistic concurrency control. See `if-match`.
	Etag *string `presentIn:"header" name:"etag"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	// a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response GetVnicAttachmentResponse) String() string {
	return common.PointerString(response)
}

// GetRawResponse implements the OciResponse interface
func (response GetVnicAttachmentResponse) GetRawResponse() *http.Response {
	return response.RawResponse
}

// GetStatefulEntity implements the OciStatefulResponse interface
func (response GetVnicAttachmentResponse) GetStatefulEntity() common.OciPollable {
	return response.VnicAttachment
}
