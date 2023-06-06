// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package datalabelingservicedataplane

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"io"
	"net/http"
	"strings"
)

// GetRecordContentRequest wrapper for the GetRecordContent operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datalabelingservicedataplane/GetRecordContent.go.html to see an example of how to use GetRecordContentRequest.
type GetRecordContentRequest struct {

	// The OCID of the record annotated.
	RecordId *string `mandatory:"true" contributesTo:"path" name:"recordId"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// For optimistic concurrency control. In the GET call
	// for a resource, set the `if-none-match` parameter to the value of the
	// etag from a previous GET or POST response for that resource.
	// The resource will be fetched only if the etag you provide does
	// not match the resource's current etag value.
	IfNoneMatch *string `mandatory:"false" contributesTo:"header" name:"if-none-match"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request GetRecordContentRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request GetRecordContentRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request GetRecordContentRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request GetRecordContentRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request GetRecordContentRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// GetRecordContentResponse wrapper for the GetRecordContent operation
type GetRecordContentResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The io.ReadCloser instance
	Content io.ReadCloser `presentIn:"body" encoding:"binary"`

	// For optimistic concurrency control. See `if-match`.
	Etag *string `presentIn:"header" name:"etag"`

	// A unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// The content size of the body in bytes.
	ContentLength *int64 `presentIn:"header" name:"content-length"`

	// The content disposition of the body, as described in RFC 2616 (https://tools.ietf.org/rfc/rfc2616), section 19.5.1.
	ContentDisposition *string `presentIn:"header" name:"content-disposition"`

	// The content type as described in RFC 2616 (https://tools.ietf.org/rfc/rfc2616), section 14.17.
	ContentType *string `presentIn:"header" name:"content-type"`

	// Cache-control allows us to tell browsers to cache assets, how long to store them, and whether to validate them.
	CacheControl *string `presentIn:"header" name:"cache-control"`

	// Flag to indicate whether or not the object was modified.  If this is true,
	// the getter for the object itself will return null.  Callers should check this
	// if they specified one of the request params that might result in a conditional
	// response (like 'if-match'/'if-none-match').
	IsNotModified bool
}

func (response GetRecordContentResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response GetRecordContentResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}
