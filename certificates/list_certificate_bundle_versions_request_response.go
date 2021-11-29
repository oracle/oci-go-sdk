// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package certificates

import (
	"github.com/oracle/oci-go-sdk/v53/common"
	"net/http"
)

// ListCertificateBundleVersionsRequest wrapper for the ListCertificateBundleVersions operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/certificates/ListCertificateBundleVersions.go.html to see an example of how to use ListCertificateBundleVersionsRequest.
type ListCertificateBundleVersionsRequest struct {

	// The OCID of the certificate.
	CertificateId *string `mandatory:"true" contributesTo:"path" name:"certificateId"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular request,
	// please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// The field to sort by. You can specify only one sort order. The default
	// order for `VERSION_NUMBER` is ascending.
	SortBy ListCertificateBundleVersionsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to use, either ascending (`ASC`) or descending (`DESC`).
	SortOrder ListCertificateBundleVersionsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListCertificateBundleVersionsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListCertificateBundleVersionsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListCertificateBundleVersionsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListCertificateBundleVersionsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ListCertificateBundleVersionsResponse wrapper for the ListCertificateBundleVersions operation
type ListCertificateBundleVersionsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The CertificateBundleVersionCollection instance
	CertificateBundleVersionCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response ListCertificateBundleVersionsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListCertificateBundleVersionsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListCertificateBundleVersionsSortByEnum Enum with underlying type: string
type ListCertificateBundleVersionsSortByEnum string

// Set of constants representing the allowable values for ListCertificateBundleVersionsSortByEnum
const (
	ListCertificateBundleVersionsSortByVersionNumber ListCertificateBundleVersionsSortByEnum = "VERSION_NUMBER"
)

var mappingListCertificateBundleVersionsSortBy = map[string]ListCertificateBundleVersionsSortByEnum{
	"VERSION_NUMBER": ListCertificateBundleVersionsSortByVersionNumber,
}

// GetListCertificateBundleVersionsSortByEnumValues Enumerates the set of values for ListCertificateBundleVersionsSortByEnum
func GetListCertificateBundleVersionsSortByEnumValues() []ListCertificateBundleVersionsSortByEnum {
	values := make([]ListCertificateBundleVersionsSortByEnum, 0)
	for _, v := range mappingListCertificateBundleVersionsSortBy {
		values = append(values, v)
	}
	return values
}

// ListCertificateBundleVersionsSortOrderEnum Enum with underlying type: string
type ListCertificateBundleVersionsSortOrderEnum string

// Set of constants representing the allowable values for ListCertificateBundleVersionsSortOrderEnum
const (
	ListCertificateBundleVersionsSortOrderAsc  ListCertificateBundleVersionsSortOrderEnum = "ASC"
	ListCertificateBundleVersionsSortOrderDesc ListCertificateBundleVersionsSortOrderEnum = "DESC"
)

var mappingListCertificateBundleVersionsSortOrder = map[string]ListCertificateBundleVersionsSortOrderEnum{
	"ASC":  ListCertificateBundleVersionsSortOrderAsc,
	"DESC": ListCertificateBundleVersionsSortOrderDesc,
}

// GetListCertificateBundleVersionsSortOrderEnumValues Enumerates the set of values for ListCertificateBundleVersionsSortOrderEnum
func GetListCertificateBundleVersionsSortOrderEnumValues() []ListCertificateBundleVersionsSortOrderEnum {
	values := make([]ListCertificateBundleVersionsSortOrderEnum, 0)
	for _, v := range mappingListCertificateBundleVersionsSortOrder {
		values = append(values, v)
	}
	return values
}
