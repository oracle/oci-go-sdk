// Copyright (c) 2016, 2017, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

package database

import (
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

// GetDbNodeRequest wrapper for the GetDbNode operation
type GetDbNodeRequest struct {

	// The database node [OCID]({{DOC_SERVER_URL}}/Content/General/Concepts/identifiers.htm).
	DbNodeId *string `mandatory:"true" contributesTo:"path" name:"dbNodeId"`
}

func (request GetDbNodeRequest) String() string {
	return common.PointerString(request)
}

// GetHttpRequest implements the OciRequest interface
func (request GetDbNodeRequest) GetHttpRequest(method, path string) (http.Request, error) {
	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// GetRetryPolicy implements the OciRetryableRequest interface
// => assembles retry policy based on specified options and default behavior
func (request GetDbNodeRequest) GetRetryPolicy(options ...common.RetryPolicyOption) common.RetryPolicy {
	if len(options) == 0 {
		return common.NoRetryPolicy()
	}
	return common.BuildRetryPolicy(options...)
}

// GetDbNodeResponse wrapper for the GetDbNode operation
type GetDbNodeResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The DbNode instance
	DbNode `presentIn:"body"`

	// For optimistic concurrency control. See `if-match`.
	Etag *string `presentIn:"header" name:"etag"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	// a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response GetDbNodeResponse) String() string {
	return common.PointerString(response)
}

// GetRawResponse implements the OciResponse interface
func (response GetDbNodeResponse) GetRawResponse() *http.Response {
	return response.RawResponse
}

// GetStatefulEntity implements the OciStatefulResponse interface
func (response GetDbNodeResponse) GetStatefulEntity() common.OciPollable {
	return response.DbNode
}
