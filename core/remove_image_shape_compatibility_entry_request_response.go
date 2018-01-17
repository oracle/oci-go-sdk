// Copyright (c) 2016, 2017, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

package core

import (
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

// RemoveImageShapeCompatibilityEntryRequest wrapper for the RemoveImageShapeCompatibilityEntry operation
type RemoveImageShapeCompatibilityEntryRequest struct {

	// The OCID of the image.
	ImageId *string `mandatory:"true" contributesTo:"path" name:"imageId"`

	// Shape name.
	ShapeName *string `mandatory:"true" contributesTo:"path" name:"shapeName"`
}

func (request RemoveImageShapeCompatibilityEntryRequest) String() string {
	return common.PointerString(request)
}

// RemoveImageShapeCompatibilityEntryResponse wrapper for the RemoveImageShapeCompatibilityEntry operation
type RemoveImageShapeCompatibilityEntryResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	// a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response RemoveImageShapeCompatibilityEntryResponse) String() string {
	return common.PointerString(response)
}
