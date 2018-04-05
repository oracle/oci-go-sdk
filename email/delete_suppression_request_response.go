// Copyright (c) 2016, 2018, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

package email

import (
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

// DeleteSuppressionRequest wrapper for the DeleteSuppression operation
type DeleteSuppressionRequest struct {

	// The unique OCID of the suppression.
	SuppressionId *string `mandatory:"true" contributesTo:"path" name:"suppressionId"`
}

func (request DeleteSuppressionRequest) String() string {
	return common.PointerString(request)
}

// DeleteSuppressionResponse wrapper for the DeleteSuppression operation
type DeleteSuppressionResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// Unique Oracle-assigned identifier for the request. If you need to
	// contact Oracle about a particular request, please provide the request
	// ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response DeleteSuppressionResponse) String() string {
	return common.PointerString(response)
}
