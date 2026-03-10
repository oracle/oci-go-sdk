// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package limitsincrease

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// UpdateLimitsIncreaseRequestRequest wrapper for the UpdateLimitsIncreaseRequest operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/limitsincrease/UpdateLimitsIncreaseRequest.go.html to see an example of how to use UpdateLimitsIncreaseRequestRequest.
type UpdateLimitsIncreaseRequestRequest struct {

	// Details for updating Limits Increase request
	UpdateLimitsIncreaseRequestDetails `contributesTo:"body"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the limit increase request.
	LimitsIncreaseRequestId *string `mandatory:"true" contributesTo:"path" name:"limitsIncreaseRequestId"`

	// Override request id for request tracking purposes.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// This value (etag) should be passed when you want to ensure that no-one else has changed the value while you're making
	// an update. To get the current etag, make a GET call and read the current etag header.
	// If GET returns 404, and you still want to ensure that noone else has executed a SET operation, pass the following
	// header instead: `if-none-match: *`
	IfMatch *string `mandatory:"false" contributesTo:"header" name:"if-match"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request UpdateLimitsIncreaseRequestRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request UpdateLimitsIncreaseRequestRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request UpdateLimitsIncreaseRequestRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request UpdateLimitsIncreaseRequestRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request UpdateLimitsIncreaseRequestRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UpdateLimitsIncreaseRequestResponse wrapper for the UpdateLimitsIncreaseRequest operation
type UpdateLimitsIncreaseRequestResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The LimitsIncreaseRequest instance
	LimitsIncreaseRequest `presentIn:"body"`

	// For optimistic concurrency control. See `if-match`.
	Etag *string `presentIn:"header" name:"etag"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response UpdateLimitsIncreaseRequestResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response UpdateLimitsIncreaseRequestResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}
