// Copyright (c) 2016, 2017, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

package database

import (
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

// GetDbSystemPatchRequest wrapper for the GetDbSystemPatch operation
type GetDbSystemPatchRequest struct {

	// The DB System [OCID]({{DOC_SERVER_URL}}/Content/General/Concepts/identifiers.htm).
	DbSystemId *string `mandatory:"true" contributesTo:"path" name:"dbSystemId"`

	// The OCID of the patch.
	PatchId *string `mandatory:"true" contributesTo:"path" name:"patchId"`
}

func (request GetDbSystemPatchRequest) String() string {
	return common.PointerString(request)
}

// GetHttpRequest implements the OciRequest interface
func (request GetDbSystemPatchRequest) GetHttpRequest(method, path string) (http.Request, error) {
	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// GetRetryPolicy implements the OciRetryableRequest interface
// => assembles retry policy based on specified options and default behavior
func (request GetDbSystemPatchRequest) GetRetryPolicy(options ...common.RetryPolicyOption) common.RetryPolicy {
	if len(options) == 0 {
		return common.NoRetryPolicy()
	}
	return common.BuildRetryPolicy(options...)
}

// GetDbSystemPatchResponse wrapper for the GetDbSystemPatch operation
type GetDbSystemPatchResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The Patch instance
	Patch `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	// a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response GetDbSystemPatchResponse) String() string {
	return common.PointerString(response)
}

// GetRawResponse implements the OciResponse interface
func (response GetDbSystemPatchResponse) GetRawResponse() *http.Response {
	return response.RawResponse
}

// GetStatefulEntity implements the OciStatefulResponse interface
func (response GetDbSystemPatchResponse) GetStatefulEntity() common.OciPollable {
	return response.Patch
}
