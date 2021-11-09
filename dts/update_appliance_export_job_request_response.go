// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package dts

import (
	"github.com/oracle/oci-go-sdk/v51/common"
	"net/http"
)

// UpdateApplianceExportJobRequest wrapper for the UpdateApplianceExportJob operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/dts/UpdateApplianceExportJob.go.html to see an example of how to use UpdateApplianceExportJobRequest.
type UpdateApplianceExportJobRequest struct {

	// ID of the Appliance Export Job
	ApplianceExportJobId *string `mandatory:"true" contributesTo:"path" name:"applianceExportJobId"`

	// fields to update
	UpdateApplianceExportJobDetails `contributesTo:"body"`

	// The entity tag to match. Optional, if set, the update will be successful only if the
	// object's tag matches the tag specified in the request.
	IfMatch *string `mandatory:"false" contributesTo:"header" name:"if-match"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	// a particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request UpdateApplianceExportJobRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request UpdateApplianceExportJobRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request UpdateApplianceExportJobRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request UpdateApplianceExportJobRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// UpdateApplianceExportJobResponse wrapper for the UpdateApplianceExportJob operation
type UpdateApplianceExportJobResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The ApplianceExportJob instance
	ApplianceExportJob `presentIn:"body"`

	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	Etag *string `presentIn:"header" name:"etag"`
}

func (response UpdateApplianceExportJobResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response UpdateApplianceExportJobResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}
