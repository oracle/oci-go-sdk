// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package announcementsservice

import (
	"github.com/oracle/oci-go-sdk/v49/common"
	"net/http"
)

// ListAnnouncementsPreferencesRequest wrapper for the ListAnnouncementsPreferences operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/announcementsservice/ListAnnouncementsPreferences.go.html to see an example of how to use ListAnnouncementsPreferencesRequest.
type ListAnnouncementsPreferencesRequest struct {

	// The OCID of the compartment. Because announcements are specific to a tenancy, this is the
	// OCID of the root compartment.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// The maximum number of items to return in a paginated "List" call.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The value of the `opc-next-page` response header from the previous "List" call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	// a particular request, please provide the complete request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListAnnouncementsPreferencesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListAnnouncementsPreferencesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListAnnouncementsPreferencesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListAnnouncementsPreferencesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ListAnnouncementsPreferencesResponse wrapper for the ListAnnouncementsPreferences operation
type ListAnnouncementsPreferencesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []AnnouncementsPreferencesSummary instances
	Items []AnnouncementsPreferencesSummary `presentIn:"body"`

	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response ListAnnouncementsPreferencesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListAnnouncementsPreferencesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}
