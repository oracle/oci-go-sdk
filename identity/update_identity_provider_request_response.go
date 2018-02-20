// Copyright (c) 2016, 2017, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

package identity

import (
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

// UpdateIdentityProviderRequest wrapper for the UpdateIdentityProvider operation
type UpdateIdentityProviderRequest struct {

	// The OCID of the identity provider.
	IdentityProviderId *string `mandatory:"true" contributesTo:"path" name:"identityProviderId"`

	// Request object for updating a identity provider.
	UpdateIdentityProviderDetails `contributesTo:"body"`

	// For optimistic concurrency control. In the PUT or DELETE call for a resource, set the `if-match`
	// parameter to the value of the etag from a previous GET or POST response for that resource.  The resource
	// will be updated or deleted only if the etag you provide matches the resource's current etag value.
	IfMatch *string `mandatory:"false" contributesTo:"header" name:"if-match"`
}

func (request UpdateIdentityProviderRequest) String() string {
	return common.PointerString(request)
}

// GetHttpRequest implements the OciRequest interface
func (request UpdateIdentityProviderRequest) GetHttpRequest(method, path string) (http.Request, error) {
	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// GetRetryPolicy implements the OciRetryableRequest interface
// => assembles retry policy based on specified options and default behavior
func (request UpdateIdentityProviderRequest) GetRetryPolicy(options ...common.RetryPolicyOption) common.RetryPolicy {
	if len(options) == 0 {
		return common.NoRetryPolicy()
	}
	return common.BuildRetryPolicy(options...)
}

// UpdateIdentityProviderResponse wrapper for the UpdateIdentityProvider operation
type UpdateIdentityProviderResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The IdentityProvider instance
	IdentityProvider `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a
	// particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For optimistic concurrency control. See `if-match`.
	Etag *string `presentIn:"header" name:"etag"`
}

func (response UpdateIdentityProviderResponse) String() string {
	return common.PointerString(response)
}

// GetRawResponse implements the OciResponse interface
func (response UpdateIdentityProviderResponse) GetRawResponse() *http.Response {
	return response.RawResponse
}

// GetStatefulEntity implements the OciStatefulResponse interface
func (response UpdateIdentityProviderResponse) GetStatefulEntity() common.OciPollable {
	return response.IdentityProvider
}
