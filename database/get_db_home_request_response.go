// Copyright (c) 2016, 2017, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

package database

import (
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

// GetDbHomeRequest wrapper for the GetDbHome operation
type GetDbHomeRequest struct {

	// The database home [OCID]({{DOC_SERVER_URL}}/Content/General/Concepts/identifiers.htm).
	DbHomeId *string `mandatory:"true" contributesTo:"path" name:"dbHomeId"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request GetDbHomeRequest) String() string {
	return common.PointerString(request)
}

// GetHttpRequest implements the OciRequest interface
func (request GetDbHomeRequest) GetHttpRequest(method, path string) (http.Request, error) {
	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// GetRetryPolicy implements the OciRetryableRequest interface
// => assembles retry policy based on specified options and default behavior
func (request GetDbHomeRequest) GetRetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// GetDbHomeResponse wrapper for the GetDbHome operation
type GetDbHomeResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The DbHome instance
	DbHome `presentIn:"body"`

	// For optimistic concurrency control. See `if-match`.
	Etag *string `presentIn:"header" name:"etag"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	// a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response GetDbHomeResponse) String() string {
	return common.PointerString(response)
}

// GetRawResponse implements the OciResponse interface
func (response GetDbHomeResponse) GetRawResponse() *http.Response {
	return response.RawResponse
}

// GetStatefulEntity implements the OciStatefulResponse interface
func (response GetDbHomeResponse) GetStatefulEntity() common.OciPollable {
	return response.DbHome
}
