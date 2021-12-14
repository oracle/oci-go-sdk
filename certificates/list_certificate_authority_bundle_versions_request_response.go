// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package certificates

import (
	"github.com/oracle/oci-go-sdk/v54/common"
	"net/http"
)

// ListCertificateAuthorityBundleVersionsRequest wrapper for the ListCertificateAuthorityBundleVersions operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/certificates/ListCertificateAuthorityBundleVersions.go.html to see an example of how to use ListCertificateAuthorityBundleVersionsRequest.
type ListCertificateAuthorityBundleVersionsRequest struct {

	// The OCID of the certificate authority (CA).
	CertificateAuthorityId *string `mandatory:"true" contributesTo:"path" name:"certificateAuthorityId"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular request,
	// please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// The field to sort by. You can specify only one sort order. The default
	// order for `VERSION_NUMBER` is ascending.
	SortBy ListCertificateAuthorityBundleVersionsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to use, either ascending (`ASC`) or descending (`DESC`).
	SortOrder ListCertificateAuthorityBundleVersionsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListCertificateAuthorityBundleVersionsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListCertificateAuthorityBundleVersionsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListCertificateAuthorityBundleVersionsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListCertificateAuthorityBundleVersionsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ListCertificateAuthorityBundleVersionsResponse wrapper for the ListCertificateAuthorityBundleVersions operation
type ListCertificateAuthorityBundleVersionsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The CertificateAuthorityBundleVersionCollection instance
	CertificateAuthorityBundleVersionCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response ListCertificateAuthorityBundleVersionsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListCertificateAuthorityBundleVersionsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListCertificateAuthorityBundleVersionsSortByEnum Enum with underlying type: string
type ListCertificateAuthorityBundleVersionsSortByEnum string

// Set of constants representing the allowable values for ListCertificateAuthorityBundleVersionsSortByEnum
const (
	ListCertificateAuthorityBundleVersionsSortByVersionNumber ListCertificateAuthorityBundleVersionsSortByEnum = "VERSION_NUMBER"
)

var mappingListCertificateAuthorityBundleVersionsSortBy = map[string]ListCertificateAuthorityBundleVersionsSortByEnum{
	"VERSION_NUMBER": ListCertificateAuthorityBundleVersionsSortByVersionNumber,
}

// GetListCertificateAuthorityBundleVersionsSortByEnumValues Enumerates the set of values for ListCertificateAuthorityBundleVersionsSortByEnum
func GetListCertificateAuthorityBundleVersionsSortByEnumValues() []ListCertificateAuthorityBundleVersionsSortByEnum {
	values := make([]ListCertificateAuthorityBundleVersionsSortByEnum, 0)
	for _, v := range mappingListCertificateAuthorityBundleVersionsSortBy {
		values = append(values, v)
	}
	return values
}

// ListCertificateAuthorityBundleVersionsSortOrderEnum Enum with underlying type: string
type ListCertificateAuthorityBundleVersionsSortOrderEnum string

// Set of constants representing the allowable values for ListCertificateAuthorityBundleVersionsSortOrderEnum
const (
	ListCertificateAuthorityBundleVersionsSortOrderAsc  ListCertificateAuthorityBundleVersionsSortOrderEnum = "ASC"
	ListCertificateAuthorityBundleVersionsSortOrderDesc ListCertificateAuthorityBundleVersionsSortOrderEnum = "DESC"
)

var mappingListCertificateAuthorityBundleVersionsSortOrder = map[string]ListCertificateAuthorityBundleVersionsSortOrderEnum{
	"ASC":  ListCertificateAuthorityBundleVersionsSortOrderAsc,
	"DESC": ListCertificateAuthorityBundleVersionsSortOrderDesc,
}

// GetListCertificateAuthorityBundleVersionsSortOrderEnumValues Enumerates the set of values for ListCertificateAuthorityBundleVersionsSortOrderEnum
func GetListCertificateAuthorityBundleVersionsSortOrderEnumValues() []ListCertificateAuthorityBundleVersionsSortOrderEnum {
	values := make([]ListCertificateAuthorityBundleVersionsSortOrderEnum, 0)
	for _, v := range mappingListCertificateAuthorityBundleVersionsSortOrder {
		values = append(values, v)
	}
	return values
}
