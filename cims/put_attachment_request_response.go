// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package cims

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"io"
	"net/http"
	"strings"
)

// PutAttachmentRequest wrapper for the PutAttachment operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/cims/PutAttachment.go.html to see an example of how to use PutAttachmentRequest.
type PutAttachmentRequest struct {

	// File to be uploaded as attachment to the Service Request.
	PutAttachmentDetails io.ReadCloser `mandatory:"true" contributesTo:"body" encoding:"binary"`

	// Unique identifier for the support request.
	IncidentKey *string `mandatory:"true" contributesTo:"path" name:"incidentKey"`

	// The name of the file to attach to the support request. Avoid entering confidential information.
	AttachmentName *string `mandatory:"true" contributesTo:"query" name:"attachmentName"`

	// The OCID of the tenancy.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// Set to `true` when the attachment contains personal information (PI)
	// or protected health information (PHI).
	IsRestrictedFlag *bool `mandatory:"true" contributesTo:"query" name:"isRestrictedFlag"`

	// The Customer Support Identifier (CSI) number associated with the support account.
	// The CSI is optional for all support request types.
	Csi *string `mandatory:"false" contributesTo:"header" name:"csi"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// User OCID for Oracle Identity Cloud Service (IDCS) users who also have a federated Oracle Cloud Infrastructure account.
	// User OCID is mandatory for OCI Users and optional for Multicloud users.
	Ocid *string `mandatory:"false" contributesTo:"header" name:"ocid"`

	// For optimistic concurrency control. In the PUT or DELETE call for a resource, set the `if-match` parameter to the value of the etag from a previous GET or POST response for that resource. The resource will be updated or deleted only if the etag you provide matches the resource's current etag value.
	IfMatch *string `mandatory:"false" contributesTo:"header" name:"if-match"`

	// The region of the tenancy.
	Homeregion *string `mandatory:"false" contributesTo:"header" name:"homeregion"`

	// The kind of support request.
	Problemtype *string `mandatory:"false" contributesTo:"header" name:"problemtype"`

	// Token type that determine which cloud provider the request come from.
	Bearertokentype *string `mandatory:"false" contributesTo:"header" name:"bearertokentype"`

	// Token that provided by multi cloud provider, which help to validate the email.
	Bearertoken *string `mandatory:"false" contributesTo:"header" name:"bearertoken"`

	// IdToken that provided by multi cloud provider, which help to validate the email.
	Idtoken *string `mandatory:"false" contributesTo:"header" name:"idtoken"`

	// The OCID of identity domain.
	// DomainID is mandatory if the user is part of Non Default Identity domain.
	Domainid *string `mandatory:"false" contributesTo:"header" name:"domainid"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request PutAttachmentRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request PutAttachmentRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {
	httpRequest, err := common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
	if err == nil && binaryRequestBody.Seekable() {
		common.UpdateRequestBinaryBody(&httpRequest, binaryRequestBody)
	}
	return httpRequest, err
}

// BinaryRequestBody implements the OCIRequest interface
func (request PutAttachmentRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {
	rsc := common.NewOCIReadSeekCloser(request.PutAttachmentDetails)
	if rsc.Seekable() {
		return rsc, true
	}
	return nil, true

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request PutAttachmentRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request PutAttachmentRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// PutAttachmentResponse wrapper for the PutAttachment operation
type PutAttachmentResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The Incident instance
	Incident `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response PutAttachmentResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response PutAttachmentResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}
