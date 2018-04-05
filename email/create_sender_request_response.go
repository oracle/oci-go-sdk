// Copyright (c) 2016, 2018, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

package email

import (
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

// CreateSenderRequest wrapper for the CreateSender operation
type CreateSenderRequest struct {

	// Create a sender.
	CreateSenderDetails `contributesTo:"body"`
}

func (request CreateSenderRequest) String() string {
	return common.PointerString(request)
}

// CreateSenderResponse wrapper for the CreateSender operation
type CreateSenderResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The Sender instance
	Sender `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need
	// to contact Oracle about a particular request, please provide the
	// request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response CreateSenderResponse) String() string {
	return common.PointerString(response)
}
