// Copyright (c) 2016, 2018, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

package database

import (
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

// ReinstateDataGuardAssociationRequest wrapper for the ReinstateDataGuardAssociation operation
type ReinstateDataGuardAssociationRequest struct {

	// The database OCID https://docs.us-phoenix-1.oraclecloud.com//Content/General/Concepts/identifiers.htm.
	DatabaseId *string `mandatory:"true" contributesTo:"path" name:"databaseId"`

	// The Data Guard association's OCID https://docs.us-phoenix-1.oraclecloud.com//Content/General/Concepts/identifiers.htm.
	DataGuardAssociationId *string `mandatory:"true" contributesTo:"path" name:"dataGuardAssociationId"`

	// A request to reinstate a database in a standby role.
	ReinstateDataGuardAssociationDetails `contributesTo:"body"`

	// For optimistic concurrency control. In the PUT or DELETE call for a resource, set the `if-match`
	// parameter to the value of the etag from a previous GET or POST response for that resource.  The resource
	// will be updated or deleted only if the etag you provide matches the resource's current etag value.
	IfMatch *string `mandatory:"false" contributesTo:"header" name:"if-match"`
}

func (request ReinstateDataGuardAssociationRequest) String() string {
	return common.PointerString(request)
}

// ReinstateDataGuardAssociationResponse wrapper for the ReinstateDataGuardAssociation operation
type ReinstateDataGuardAssociationResponse struct {

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

func (response ReinstateDataGuardAssociationResponse) String() string {
	return common.PointerString(response)
}
