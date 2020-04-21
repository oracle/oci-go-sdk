// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package dts

import (
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

// ListTransferApplianceEntitlementRequest wrapper for the ListTransferApplianceEntitlement operation
type ListTransferApplianceEntitlementRequest struct {

	// compartment id
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// filtering by Transfer Appliance Entitlement id
	Id *string `mandatory:"false" contributesTo:"query" name:"id"`

	// filtering by displayName
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	// a particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListTransferApplianceEntitlementRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListTransferApplianceEntitlementRequest) HTTPRequest(method, path string) (http.Request, error) {
	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListTransferApplianceEntitlementRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ListTransferApplianceEntitlementResponse wrapper for the ListTransferApplianceEntitlement operation
type ListTransferApplianceEntitlementResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The []TransferApplianceEntitlementSummary instance
	Items []TransferApplianceEntitlementSummary `presentIn:"body"`

	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response ListTransferApplianceEntitlementResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListTransferApplianceEntitlementResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}
