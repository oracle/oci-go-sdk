// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package cims

import (
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

// GetStatusRequest wrapper for the GetStatus operation
type GetStatusRequest struct {

	// Source is a downstream system. Eg: JIRA or MOS or any other source in future.
	Source *string `mandatory:"true" contributesTo:"path" name:"source"`

	// User OCID for IDCS users that have a shadow in OCI
	Ocid *string `mandatory:"true" contributesTo:"header" name:"ocid"`

	// Unique Header for request id
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request GetStatusRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request GetStatusRequest) HTTPRequest(method, path string) (http.Request, error) {
	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request GetStatusRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// GetStatusResponse wrapper for the GetStatus operation
type GetStatusResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The Status instance
	Status `presentIn:"body"`

	// OPC Request Id
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// e-Tag
	Etag *string `presentIn:"header" name:"etag"`
}

func (response GetStatusResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response GetStatusResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}
