// Copyright (c) 2016, 2017, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

package database

import (
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

// SwitchoverDataGuardAssociationRequest wrapper for the SwitchoverDataGuardAssociation operation
type SwitchoverDataGuardAssociationRequest struct {

	// The database [OCID]({{DOC_SERVER_URL}}/Content/General/Concepts/identifiers.htm).
	DatabaseId *string `mandatory:"true" contributesTo:"path" name:"databaseId"`

	// The Data Guard association's [OCID]({{DOC_SERVER_URL}}/Content/General/Concepts/identifiers.htm).
	DataGuardAssociationId *string `mandatory:"true" contributesTo:"path" name:"dataGuardAssociationId"`

	// Request to swtichover a primary to a standby.
	SwitchoverDataGuardAssociationDetails `contributesTo:"body"`

	// For optimistic concurrency control. In the PUT or DELETE call for a resource, set the `if-match`
	// parameter to the value of the etag from a previous GET or POST response for that resource.  The resource
	// will be updated or deleted only if the etag you provide matches the resource's current etag value.
	IfMatch *string `mandatory:"false" contributesTo:"header" name:"if-match"`
}

func (request SwitchoverDataGuardAssociationRequest) String() string {
	return common.PointerString(request)
}

// GetHttpRequest implements the OciRequest interface
func (request SwitchoverDataGuardAssociationRequest) GetHttpRequest(method, path string) (http.Request, error) {
	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// GetRetryPolicy implements the OciRetryableRequest interface
// => assembles retry policy based on specified options and default behavior
func (request SwitchoverDataGuardAssociationRequest) GetRetryPolicy(options ...common.RetryPolicyOption) common.RetryPolicy {
	if len(options) == 0 {
		return common.NoRetryPolicy()
	}
	return common.BuildRetryPolicy(options...)
}

// SwitchoverDataGuardAssociationResponse wrapper for the SwitchoverDataGuardAssociation operation
type SwitchoverDataGuardAssociationResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The DataGuardAssociation instance
	DataGuardAssociation `presentIn:"body"`

	// For optimistic concurrency control. See `if-match`.
	Etag *string `presentIn:"header" name:"etag"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	// a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response SwitchoverDataGuardAssociationResponse) String() string {
	return common.PointerString(response)
}

// GetRawResponse implements the OciResponse interface
func (response SwitchoverDataGuardAssociationResponse) GetRawResponse() *http.Response {
	return response.RawResponse
}

// GetStatefulEntity implements the OciStatefulResponse interface
func (response SwitchoverDataGuardAssociationResponse) GetStatefulEntity() common.OciPollable {
	return response.DataGuardAssociation
}
