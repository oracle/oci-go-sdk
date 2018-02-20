// Copyright (c) 2016, 2017, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

package database

import (
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

// GetDbSystemPatchHistoryEntryRequest wrapper for the GetDbSystemPatchHistoryEntry operation
type GetDbSystemPatchHistoryEntryRequest struct {

	// The DB System [OCID]({{DOC_SERVER_URL}}/Content/General/Concepts/identifiers.htm).
	DbSystemId *string `mandatory:"true" contributesTo:"path" name:"dbSystemId"`

	// The OCID of the patch history entry.
	PatchHistoryEntryId *string `mandatory:"true" contributesTo:"path" name:"patchHistoryEntryId"`
}

func (request GetDbSystemPatchHistoryEntryRequest) String() string {
	return common.PointerString(request)
}

// GetHttpRequest implements the OciRequest interface
func (request GetDbSystemPatchHistoryEntryRequest) GetHttpRequest(method, path string) (http.Request, error) {
	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// GetRetryPolicy implements the OciRetryableRequest interface
// => assembles retry policy based on specified options and default behavior
func (request GetDbSystemPatchHistoryEntryRequest) GetRetryPolicy(options ...common.RetryPolicyOption) common.RetryPolicy {
	if len(options) == 0 {
		return common.NoRetryPolicy()
	}
	return common.BuildRetryPolicy(options...)
}

// GetDbSystemPatchHistoryEntryResponse wrapper for the GetDbSystemPatchHistoryEntry operation
type GetDbSystemPatchHistoryEntryResponse struct {

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

func (response GetDbSystemPatchHistoryEntryResponse) String() string {
	return common.PointerString(response)
}

// GetRawResponse implements the OciResponse interface
func (response GetDbSystemPatchHistoryEntryResponse) GetRawResponse() *http.Response {
	return response.RawResponse
}

// GetStatefulEntity implements the OciStatefulResponse interface
func (response GetDbSystemPatchHistoryEntryResponse) GetStatefulEntity() common.OciPollable {
	return response.PatchHistoryEntry
}
