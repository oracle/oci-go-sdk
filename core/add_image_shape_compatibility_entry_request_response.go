// Copyright (c) 2016, 2017, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

package core

import (
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

// AddImageShapeCompatibilityEntryRequest wrapper for the AddImageShapeCompatibilityEntry operation
type AddImageShapeCompatibilityEntryRequest struct {

	// The OCID of the image.
	ImageId *string `mandatory:"true" contributesTo:"path" name:"imageId"`

	// Shape name.
	ShapeName *string `mandatory:"true" contributesTo:"path" name:"shapeName"`
}

func (request AddImageShapeCompatibilityEntryRequest) String() string {
	return common.PointerString(request)
}

// AddImageShapeCompatibilityEntryResponse wrapper for the AddImageShapeCompatibilityEntry operation
type AddImageShapeCompatibilityEntryResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The ImageShapeCompatibilityEntry instance
	ImageShapeCompatibilityEntry `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	// a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response AddImageShapeCompatibilityEntryResponse) String() string {
	return common.PointerString(response)
}
