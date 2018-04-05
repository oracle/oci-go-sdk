// Copyright (c) 2016, 2018, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

package email

import (
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

// DeleteSenderRequest wrapper for the DeleteSender operation
type DeleteSenderRequest struct {

	// The unique OCID of the sender.
	SenderId *string `mandatory:"true" contributesTo:"path" name:"senderId"`
}

func (request DeleteSenderRequest) String() string {
	return common.PointerString(request)
}

// DeleteSenderResponse wrapper for the DeleteSender operation
type DeleteSenderResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// Unique Oracle-assigned identifier for the request. If you need to
	// contact Oracle about a particular request, please provide the request
	// ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response DeleteSenderResponse) String() string {
	return common.PointerString(response)
}
