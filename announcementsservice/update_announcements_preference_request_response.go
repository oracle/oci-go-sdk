// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package announcementsservice

import (
	"github.com/oracle/oci-go-sdk/v48/common"
	"net/http"
)

// UpdateAnnouncementsPreferenceRequest wrapper for the UpdateAnnouncementsPreference operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/announcementsservice/UpdateAnnouncementsPreference.go.html to see an example of how to use UpdateAnnouncementsPreferenceRequest.
type UpdateAnnouncementsPreferenceRequest struct {

	// The ID of the preference.
	PreferenceId *string `mandatory:"true" contributesTo:"path" name:"preferenceId"`

	// The object that contains details about tenancy preferences for receiving announcements by email.
	AnnouncementsPreferenceDetails UpdateAnnouncementsPreferencesDetails `contributesTo:"body"`

	// The unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	// a particular request, please provide the complete request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// The locking version, used for optimistic concurrency control.
	IfMatch *string `mandatory:"false" contributesTo:"header" name:"if-match"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request UpdateAnnouncementsPreferenceRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request UpdateAnnouncementsPreferenceRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request UpdateAnnouncementsPreferenceRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request UpdateAnnouncementsPreferenceRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// UpdateAnnouncementsPreferenceResponse wrapper for the UpdateAnnouncementsPreference operation
type UpdateAnnouncementsPreferenceResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The AnnouncementsPreferencesSummary instance
	AnnouncementsPreferencesSummary `presentIn:"body"`

	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	Etag *string `presentIn:"header" name:"etag"`
}

func (response UpdateAnnouncementsPreferenceResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response UpdateAnnouncementsPreferenceResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}
