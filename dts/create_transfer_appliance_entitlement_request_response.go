// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

package dts

import (
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

// CreateTransferApplianceEntitlementRequest wrapper for the CreateTransferApplianceEntitlement operation
type CreateTransferApplianceEntitlementRequest struct {

	// Creates a Transfer Appliance Entitlement
	CreateTransferApplianceEntitlementDetails `contributesTo:"body"`

	// opc retry token
	OpcRetryToken *string `mandatory:"false" contributesTo:"header" name:"opc-retry-token"`

	// opc request id
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request CreateTransferApplianceEntitlementRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request CreateTransferApplianceEntitlementRequest) HTTPRequest(method, path string) (http.Request, error) {
	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request CreateTransferApplianceEntitlementRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// CreateTransferApplianceEntitlementResponse wrapper for the CreateTransferApplianceEntitlement operation
type CreateTransferApplianceEntitlementResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The TransferApplianceEntitlement instance
	TransferApplianceEntitlement `presentIn:"body"`

	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	Etag *string `presentIn:"header" name:"etag"`
}

func (response CreateTransferApplianceEntitlementResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response CreateTransferApplianceEntitlementResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}
