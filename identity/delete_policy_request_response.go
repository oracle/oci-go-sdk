// Copyright (c) 2016, 2017, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

package identity

import (
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

// DeletePolicyRequest wrapper for the DeletePolicy operation
type DeletePolicyRequest struct {

	// The OCID of the policy.
	PolicyId *string `mandatory:"true" contributesTo:"path" name:"policyId"`

	// For optimistic concurrency control. In the PUT or DELETE call for a resource, set the `if-match`
	// parameter to the value of the etag from a previous GET or POST response for that resource.  The resource
	// will be updated or deleted only if the etag you provide matches the resource's current etag value.
	IfMatch *string `mandatory:"false" contributesTo:"header" name:"if-match"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request DeletePolicyRequest) String() string {
	return common.PointerString(request)
}

// GetHttpRequest implements the OciRequest interface
func (request DeletePolicyRequest) GetHttpRequest(method, path string) (http.Request, error) {
	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// GetRetryPolicy implements the OciRetryableRequest interface
// => assembles retry policy based on specified options and default behavior
func (request DeletePolicyRequest) GetRetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// DeletePolicyResponse wrapper for the DeletePolicy operation
type DeletePolicyResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a
	// particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response DeletePolicyResponse) String() string {
	return common.PointerString(response)
}

// GetRawResponse implements the OciResponse interface
func (response DeletePolicyResponse) GetRawResponse() *http.Response {
	return response.RawResponse
}
