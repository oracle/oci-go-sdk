// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package cims

import (
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

// GetIncidentRequest wrapper for the GetIncident operation
type GetIncidentRequest struct {

	// Unique identifier for the support ticket.
	IncidentKey *string `mandatory:"true" contributesTo:"path" name:"incidentKey"`

	// The Customer Support Identifier associated with the support account.
	Csi *string `mandatory:"true" contributesTo:"header" name:"csi"`

	// User OCID for Oracle Identity Cloud Service (IDCS) users who also have a federated Oracle Cloud Infrastructure account.
	Ocid *string `mandatory:"true" contributesTo:"header" name:"ocid"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// The region of the tenancy.
	Homeregion *string `mandatory:"false" contributesTo:"header" name:"homeregion"`

	// The kind of support request.
	ProblemType *string `mandatory:"false" contributesTo:"header" name:"problem-type"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request GetIncidentRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request GetIncidentRequest) HTTPRequest(method, path string) (http.Request, error) {
	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request GetIncidentRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// GetIncidentResponse wrapper for the GetIncident operation
type GetIncidentResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The Incident instance
	Incident `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response GetIncidentResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response GetIncidentResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}
