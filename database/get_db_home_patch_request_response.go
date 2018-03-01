// Copyright (c) 2016, 2017, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

package database

import (
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

// GetDbHomePatchRequest wrapper for the GetDbHomePatch operation
type GetDbHomePatchRequest struct {

	// The database home [OCID]({{DOC_SERVER_URL}}/Content/General/Concepts/identifiers.htm).
	DbHomeId *string `mandatory:"true" contributesTo:"path" name:"dbHomeId"`

	// The OCID of the patch.
	PatchId *string `mandatory:"true" contributesTo:"path" name:"patchId"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request GetDbHomePatchRequest) String() string {
	return common.PointerString(request)
}

// GetHttpRequest implements the OciRequest interface
func (request GetDbHomePatchRequest) GetHttpRequest(method, path string) (http.Request, error) {
	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// GetRetryPolicy implements the OciRetryableRequest interface
// => assembles retry policy based on specified options and default behavior
func (request GetDbHomePatchRequest) GetRetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// GetDbHomePatchResponse wrapper for the GetDbHomePatch operation
type GetDbHomePatchResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The Patch instance
	Patch `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	// a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response GetDbHomePatchResponse) String() string {
	return common.PointerString(response)
}

// GetRawResponse implements the OciResponse interface
func (response GetDbHomePatchResponse) GetRawResponse() *http.Response {
	return response.RawResponse
}

// GetStatefulEntity implements the OciStatefulResponse interface
func (response GetDbHomePatchResponse) GetStatefulEntity() common.OciPollable {
	return response.Patch
}
