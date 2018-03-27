// Copyright (c) 2016, 2018, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

package filestorage

import (
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

// DeleteExportRequest wrapper for the DeleteExport operation
type DeleteExportRequest struct {

	// The OCID of the export.
	ExportId *string `mandatory:"true" contributesTo:"path" name:"exportId"`

	// For optimistic concurrency control. In the PUT or DELETE call
	// for a resource, set the `if-match` parameter to the value of the
	// etag from a previous GET or POST response for that resource.
	// The resource will be updated or deleted only if the etag you
	// provide matches the resource's current etag value.
	IfMatch *string `mandatory:"false" contributesTo:"header" name:"if-match"`
}

func (request DeleteExportRequest) String() string {
	return common.PointerString(request)
}

// DeleteExportResponse wrapper for the DeleteExport operation
type DeleteExportResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// Unique Oracle-assigned identifier for the request. If
	// you need to contact Oracle about a particular request,
	// please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response DeleteExportResponse) String() string {
	return common.PointerString(response)
}
