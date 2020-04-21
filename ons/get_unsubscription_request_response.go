// Copyright (c) 2016, 2018, , Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package ons

import (
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

// GetUnsubscriptionRequest wrapper for the GetUnsubscription operation
type GetUnsubscriptionRequest struct {

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the subscription to unsubscribe from.
	Id *string `mandatory:"true" contributesTo:"path" name:"id"`

	// The subscription confirmation token.
	Token *string `mandatory:"true" contributesTo:"query" name:"token"`

	// The protocol used for the subscription.
	// Allowed values:
	//   * `CUSTOM_HTTPS`
	//   * `EMAIL`
	//   * `HTTPS` (deprecated; for PagerDuty endpoints, use `PAGERDUTY`)
	//   * `PAGERDUTY`
	//   * `SLACK`
	//   * `ORACLE_FUNCTIONS`
	// For information about subscription protocols, see
	// To create a subscription (https://docs.cloud.oracle.com/iaas/Content/Notification/Tasks/managingtopicsandsubscriptions.htm#createSub).
	Protocol *string `mandatory:"true" contributesTo:"query" name:"protocol"`

	// The unique Oracle-assigned identifier for the request. If you need to contact Oracle about a
	// particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request GetUnsubscriptionRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request GetUnsubscriptionRequest) HTTPRequest(method, path string) (http.Request, error) {
	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request GetUnsubscriptionRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// GetUnsubscriptionResponse wrapper for the GetUnsubscription operation
type GetUnsubscriptionResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The string instance
	Value *string `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	// a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response GetUnsubscriptionResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response GetUnsubscriptionResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}
