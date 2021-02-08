// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package computeinstanceagent

import (
	"github.com/oracle/oci-go-sdk/v35/common"
	"net/http"
)

// CreateInstanceAgentCommandRequest wrapper for the CreateInstanceAgentCommand operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/computeinstanceagent/CreateInstanceAgentCommand.go.html to see an example of how to use CreateInstanceAgentCommandRequest.
type CreateInstanceAgentCommandRequest struct {

	// Create Instance agent command details
	CreateInstanceAgentCommandDetails `contributesTo:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular request,
	// please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// A token that uniquely identifies a request so it can be retried in case of a timeout or
	// server error without risk of executing that same action again. Retry tokens expire after 24
	// hours, but can be invalidated before then due to conflicting operations (for example, if a resource
	// has been deleted and purged from the system, then a retry of the original creation request
	// may be rejected).
	OpcRetryToken *string `mandatory:"false" contributesTo:"header" name:"opc-retry-token"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request CreateInstanceAgentCommandRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request CreateInstanceAgentCommandRequest) HTTPRequest(method, path string) (http.Request, error) {
	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request CreateInstanceAgentCommandRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// CreateInstanceAgentCommandResponse wrapper for the CreateInstanceAgentCommand operation
type CreateInstanceAgentCommandResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The InstanceAgentCommand instance
	InstanceAgentCommand `presentIn:"body"`

	// For optimistic concurrency control. See `if-match`.
	Etag *string `presentIn:"header" name:"etag"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response CreateInstanceAgentCommandResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response CreateInstanceAgentCommandResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}
