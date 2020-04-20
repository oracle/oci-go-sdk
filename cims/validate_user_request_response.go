// Copyright (c) 2016, 2018, , Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package cims

import (
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

// ValidateUserRequest wrapper for the ValidateUser operation
type ValidateUserRequest struct {

	// Customer support identifier of the support account
	Csi *string `mandatory:"true" contributesTo:"header" name:"csi"`

	// User OCID for IDCS users that have a shadow in OCI
	Ocid *string `mandatory:"true" contributesTo:"header" name:"ocid"`

	// Retry-token header
	OpcRetryToken *string `mandatory:"false" contributesTo:"header" name:"opc-retry-token"`

	// Unique request id
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Problem Type of Taxonomy - tech/limit
	ProblemType *string `mandatory:"false" contributesTo:"query" name:"problemType"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ValidateUserRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ValidateUserRequest) HTTPRequest(method, path string) (http.Request, error) {
	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ValidateUserRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateUserResponse wrapper for the ValidateUser operation
type ValidateUserResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The ValidationResponse instance
	ValidationResponse `presentIn:"body"`

	// OPC Request Id
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// e-Tag
	Etag *string `presentIn:"header" name:"etag"`
}

func (response ValidateUserResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ValidateUserResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}
