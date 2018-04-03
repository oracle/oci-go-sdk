// Copyright (c) 2016, 2018, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

package email

import (
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

// GetSenderRequest wrapper for the GetSender operation
type GetSenderRequest struct {

	// The unique OCID of the sender.
	SenderId *string `mandatory:"true" contributesTo:"path" name:"senderId"`
}

func (request GetSenderRequest) String() string {
	return common.PointerString(request)
}

// GetSenderResponse wrapper for the GetSender operation
type GetSenderResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The Sender instance
	Sender `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need
	// to contact Oracle about a particular request, please provide the
	// request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response GetSenderResponse) String() string {
	return common.PointerString(response)
}
