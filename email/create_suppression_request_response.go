// Copyright (c) 2016, 2018, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

package email

import (
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

// CreateSuppressionRequest wrapper for the CreateSuppression operation
type CreateSuppressionRequest struct {

	// Adds a single email address to the suppression list for a compartment's tenancy.
	CreateSuppressionDetails `contributesTo:"body"`
}

func (request CreateSuppressionRequest) String() string {
	return common.PointerString(request)
}

// CreateSuppressionResponse wrapper for the CreateSuppression operation
type CreateSuppressionResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The Suppression instance
	Suppression `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response CreateSuppressionResponse) String() string {
	return common.PointerString(response)
}
