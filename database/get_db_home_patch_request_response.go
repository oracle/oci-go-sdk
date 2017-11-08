// Copyright (c) 2016, 2017, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

package database

import (
	"bitbucket.aka.lgl.grungy.us/golang-sdk2/common"
	"net/http"
)

// GetDbHomePatchRequest wrapper for the GetDbHomePatch operation
type GetDbHomePatchRequest struct {

	// The database home [OCID]({{DOC_SERVER_URL}}/Content/General/Concepts/identifiers.htm).
	DbHomeID *string `mandatory:"true" contributesTo:"path" name:"dbHomeId"`

	// The OCID of the patch.
	PatchID *string `mandatory:"true" contributesTo:"path" name:"patchId"`
}

func (request GetDbHomePatchRequest) String() string {
	return common.PointerString(request)
}

// GetDbHomePatchResponse wrapper for the GetDbHomePatch operation
type GetDbHomePatchResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The Patch instance
	Patch `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	// a particular request, please provide the request ID.
	OpcRequestID *string `presentIn:"header" name:"opc-request-id"`
}

func (response GetDbHomePatchResponse) String() string {
	return common.PointerString(response)
}
