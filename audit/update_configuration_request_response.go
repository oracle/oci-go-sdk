// Copyright (c) 2016, 2017, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

package audit

import (
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

// UpdateConfigurationRequest wrapper for the UpdateConfiguration operation
type UpdateConfigurationRequest struct {

	// ID of the root compartment (tenancy)
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// The configuration properties
	UpdateConfigurationDetails `contributesTo:"body"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request UpdateConfigurationRequest) String() string {
	return common.PointerString(request)
}

// GetHttpRequest implements the OciRequest interface
func (request UpdateConfigurationRequest) GetHttpRequest(method, path string) (http.Request, error) {
	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// GetRetryPolicy implements the OciRetryableRequest interface
// => assembles retry policy based on specified options and default behavior
func (request UpdateConfigurationRequest) GetRetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// UpdateConfigurationResponse wrapper for the UpdateConfiguration operation
type UpdateConfigurationResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a
	// particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// The [OCID]({{DOC_SERVER_URL}}/Content/General/Concepts/identifiers.htm) of the work request.
	OpcWorkRequestId *string `presentIn:"header" name:"opc-work-request-id"`
}

func (response UpdateConfigurationResponse) String() string {
	return common.PointerString(response)
}

// GetRawResponse implements the OciResponse interface
func (response UpdateConfigurationResponse) GetRawResponse() *http.Response {
	return response.RawResponse
}
