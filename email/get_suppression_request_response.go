// Copyright (c) 2016, 2018, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

package email

import (
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

// GetSuppressionRequest wrapper for the GetSuppression operation
type GetSuppressionRequest struct {

	// The unique OCID of the suppression.
	SuppressionId *string `mandatory:"true" contributesTo:"path" name:"suppressionId"`
}

func (request GetSuppressionRequest) String() string {
	return common.PointerString(request)
}

// GetSuppressionResponse wrapper for the GetSuppression operation
type GetSuppressionResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The Suppression instance
	Suppression `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response GetSuppressionResponse) String() string {
	return common.PointerString(response)
}
