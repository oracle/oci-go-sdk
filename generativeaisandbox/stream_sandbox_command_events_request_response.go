// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package generativeaisandbox

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"io"
	"net/http"
	"strings"
)

// StreamSandboxCommandEventsRequest wrapper for the StreamSandboxCommandEvents operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/generativeaisandbox/StreamSandboxCommandEvents.go.html to see an example of how to use StreamSandboxCommandEventsRequest.
type StreamSandboxCommandEventsRequest struct {

	// Sandbox identifier. Newly generated Sandbox IDs use `sbx-` followed by 16 lowercase base32 characters (`a-z2-7`); existing legacy IDs remain accepted.
	SandboxId *string `mandatory:"true" contributesTo:"path" name:"sandboxId"`

	// Sandbox command identifier.
	CommandId *string `mandatory:"true" contributesTo:"path" name:"commandId"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	// The only valid characters for request IDs are letters, numbers,
	// underscore, and dash.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// The first command event sequence number to include.
	FromSeq *int64 `mandatory:"false" contributesTo:"query" name:"fromSeq"`

	// The last command event sequence number to include. If omitted, all available events from `fromSeq` are included.
	UntilSeq *int64 `mandatory:"false" contributesTo:"query" name:"untilSeq"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request StreamSandboxCommandEventsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request StreamSandboxCommandEventsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request StreamSandboxCommandEventsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request StreamSandboxCommandEventsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request StreamSandboxCommandEventsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// StreamSandboxCommandEventsResponse wrapper for the StreamSandboxCommandEvents operation
type StreamSandboxCommandEventsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The io.ReadCloser instance
	Content io.ReadCloser `presentIn:"body" encoding:"binary"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response StreamSandboxCommandEventsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response StreamSandboxCommandEventsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}
