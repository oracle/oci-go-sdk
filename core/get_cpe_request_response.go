// Copyright (c) 2016, 2017, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

package core

import (
	"bitbucket.aka.lgl.grungy.us/golang-sdk2/common"
	"net/http"
)

// GetCpeRequest wrapper for the GetCpe operation
type GetCpeRequest struct {

	// The OCID of the CPE.
	CpeID *string `mandatory:"true" contributesTo:"path" name:"cpeId"`
}

func (request GetCpeRequest) String() string {
	return common.PointerString(request)
}

// GetCpeResponse wrapper for the GetCpe operation
type GetCpeResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The Cpe instance
	Cpe `presentIn:"body"`

	// For optimistic concurrency control. See `if-match`.
	Etag *string `presentIn:"header" name:"etag"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	// a particular request, please provide the request ID.
	OpcRequestID *string `presentIn:"header" name:"opc-request-id"`
}

func (response GetCpeResponse) String() string {
	return common.PointerString(response)
}
