// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package oda

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// UpdateDigitalAssistantRequest wrapper for the UpdateDigitalAssistant operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/oda/UpdateDigitalAssistant.go.html to see an example of how to use UpdateDigitalAssistantRequest.
type UpdateDigitalAssistantRequest struct {

	// Unique Digital Assistant instance identifier.
	OdaInstanceId *string `mandatory:"true" contributesTo:"path" name:"odaInstanceId"`

	// Unique Digital Assistant identifier.
	DigitalAssistantId *string `mandatory:"true" contributesTo:"path" name:"digitalAssistantId"`

	// Property values to update the Digital Assistant.
	UpdateDigitalAssistantDetails `contributesTo:"body"`

	// For optimistic concurrency control in a PUT or DELETE call for
	// a Digital Assistant instance, set the `if-match` query parameter
	// to the value of the `ETAG` header from a previous GET or POST
	// response for that instance. The service updates or deletes the
	// instance only if the etag that you provide matches the instance's
	// current etag value.
	IfMatch *string `mandatory:"false" contributesTo:"header" name:"if-match"`

	// The client request ID for tracing. This value is included in the opc-request-id response header.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request UpdateDigitalAssistantRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request UpdateDigitalAssistantRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request UpdateDigitalAssistantRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request UpdateDigitalAssistantRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request UpdateDigitalAssistantRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UpdateDigitalAssistantResponse wrapper for the UpdateDigitalAssistant operation
type UpdateDigitalAssistantResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The DigitalAssistant instance
	DigitalAssistant `presentIn:"body"`

	// For use in a PUT or DELETE `if-match` query parameter for optimistic concurrency control.
	Etag *string `presentIn:"header" name:"etag"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response UpdateDigitalAssistantResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response UpdateDigitalAssistantResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}
