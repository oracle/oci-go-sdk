// Copyright (c) 2016, 2017, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

package objectstorage

import (
	"bitbucket.aka.lgl.grungy.us/golang-sdk2/common"
	"net/http"
)

// GetNamespaceRequest wrapper for the GetNamespace operation
type GetNamespaceRequest struct {

	// The client request ID for tracing.
	OpcClientRequestID *string `mandatory:"false" contributesTo:"header" name:"opc-client-request-id"`
}

func (request GetNamespaceRequest) String() string {
	return common.PointerString(request)
}

// GetNamespaceResponse wrapper for the GetNamespace operation
type GetNamespaceResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The string instance
	value string `presentIn:"body"`
}

func (response GetNamespaceResponse) String() string {
	return common.PointerString(response)
}
