// Copyright (c) 2016, 2017, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

package database

import (
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

// GetDbHomePatchHistoryEntryRequest wrapper for the GetDbHomePatchHistoryEntry operation
type GetDbHomePatchHistoryEntryRequest struct {

	// The database home [OCID]({{DOC_SERVER_URL}}/Content/General/Concepts/identifiers.htm).
	DbHomeId *string `mandatory:"true" contributesTo:"path" name:"dbHomeId"`

	// The OCID of the patch history entry.
	PatchHistoryEntryId *string `mandatory:"true" contributesTo:"path" name:"patchHistoryEntryId"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request GetDbHomePatchHistoryEntryRequest) String() string {
	return common.PointerString(request)
}

// GetHttpRequest implements the OciRequest interface
func (request GetDbHomePatchHistoryEntryRequest) GetHttpRequest(method, path string) (http.Request, error) {
	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// GetRetryPolicy implements the OciRetryableRequest interface
// => assembles retry policy based on specified options and default behavior
func (request GetDbHomePatchHistoryEntryRequest) GetRetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// GetDbHomePatchHistoryEntryResponse wrapper for the GetDbHomePatchHistoryEntry operation
type GetDbHomePatchHistoryEntryResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The PatchHistoryEntry instance
	PatchHistoryEntry `presentIn:"body"`

	// For optimistic concurrency control. See `if-match`.
	Etag *string `presentIn:"header" name:"etag"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	// a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response GetDbHomePatchHistoryEntryResponse) String() string {
	return common.PointerString(response)
}

// GetRawResponse implements the OciResponse interface
func (response GetDbHomePatchHistoryEntryResponse) GetRawResponse() *http.Response {
	return response.RawResponse
}

// GetStatefulEntity implements the OciStatefulResponse interface
func (response GetDbHomePatchHistoryEntryResponse) GetStatefulEntity() common.OciPollable {
	return response.PatchHistoryEntry
}
