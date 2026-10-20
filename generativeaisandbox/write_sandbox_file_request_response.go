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

// WriteSandboxFileRequest wrapper for the WriteSandboxFile operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/generativeaisandbox/WriteSandboxFile.go.html to see an example of how to use WriteSandboxFileRequest.
type WriteSandboxFileRequest struct {

	// Sandbox identifier. Newly generated Sandbox IDs use `sbx-` followed by 16 lowercase base32 characters (`a-z2-7`); existing legacy IDs remain accepted.
	SandboxId *string `mandatory:"true" contributesTo:"path" name:"sandboxId"`

	// Absolute path under /workspace where the file will be written.
	FilePath *string `mandatory:"true" contributesTo:"query" name:"filePath"`

	// Raw file bytes to write.
	WriteSandboxFileContentDetails io.ReadCloser `mandatory:"true" contributesTo:"body" encoding:"binary"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	// The only valid characters for request IDs are letters, numbers,
	// underscore, and dash.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// A token that uniquely identifies a request so it can be retried in case of a timeout or
	// server error without risk of running that same action again. Retry tokens expire after 24
	// hours, but can be invalidated before then due to conflicting operations. For example, if a resource
	// has been deleted and removed from the system, then a retry of the original creation request
	// might be rejected.
	OpcRetryToken *string `mandatory:"false" contributesTo:"header" name:"opc-retry-token"`

	// Optional base64-encoded SHA-256 checksum of the sandbox file content. When supported by the sandbox runtime,
	// the runtime validates the uploaded file content against this value and rejects the request if the values do not match.
	OpcContentSha256 *string `mandatory:"false" contributesTo:"header" name:"opc-content-sha256"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request WriteSandboxFileRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request WriteSandboxFileRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {
	httpRequest, err := common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
	if err == nil && binaryRequestBody.Seekable() {
		common.UpdateRequestBinaryBody(&httpRequest, binaryRequestBody)
	}
	return httpRequest, err
}

// BinaryRequestBody implements the OCIRequest interface
func (request WriteSandboxFileRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {
	rsc := common.NewOCIReadSeekCloser(request.WriteSandboxFileContentDetails)
	if rsc.Seekable() {
		return rsc, true
	}
	return nil, true

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request WriteSandboxFileRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request WriteSandboxFileRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// WriteSandboxFileResponse wrapper for the WriteSandboxFile operation
type WriteSandboxFileResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The SandboxFile instance
	SandboxFile `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// Base64-encoded SHA-256 checksum of the sandbox file content.
	OpcContentSha256 *string `presentIn:"header" name:"opc-content-sha256"`
}

func (response WriteSandboxFileResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response WriteSandboxFileResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}
