// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package certificatesmanagement

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListCertificateVersionsRequest wrapper for the ListCertificateVersions operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/certificatesmanagement/ListCertificateVersions.go.html to see an example of how to use ListCertificateVersionsRequest.
type ListCertificateVersionsRequest struct {

	// The OCID of the certificate.
	CertificateId *string `mandatory:"true" contributesTo:"path" name:"certificateId"`

	// Unique Oracle-assigned identifier for the request. If provided, the returned request ID
	// will include this value. Otherwise, a random request ID will be
	// generated by the service.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// A filter that returns only resources that match the specified version number. The default value is 0, which means that this filter is not applied.
	VersionNumber *int64 `mandatory:"false" contributesTo:"query" name:"versionNumber"`

	// The maximum number of items to return in a paginated "List" call.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The value of the `opc-next-page` response header
	// from the previous "List" call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The field to sort by. You can specify only one sort order. The default order for 'VERSION_NUMBER' is ascending.
	SortBy ListCertificateVersionsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to use, either ascending (`ASC`) or descending (`DESC`).
	SortOrder ListCertificateVersionsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListCertificateVersionsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListCertificateVersionsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListCertificateVersionsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListCertificateVersionsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListCertificateVersionsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListCertificateVersionsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListCertificateVersionsSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListCertificateVersionsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListCertificateVersionsSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListCertificateVersionsResponse wrapper for the ListCertificateVersions operation
type ListCertificateVersionsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of CertificateVersionCollection instances
	CertificateVersionCollection `presentIn:"body"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then there are additional items still to get. Include this value as the `page` parameter for the
	// subsequent GET request. For information about pagination, see
	// List Pagination (https://docs.cloud.oracle.com/Content/API/Concepts/usingapi.htm#List_Pagination).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	// a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response ListCertificateVersionsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListCertificateVersionsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListCertificateVersionsSortByEnum Enum with underlying type: string
type ListCertificateVersionsSortByEnum string

// Set of constants representing the allowable values for ListCertificateVersionsSortByEnum
const (
	ListCertificateVersionsSortByVersionNumber ListCertificateVersionsSortByEnum = "VERSION_NUMBER"
)

var mappingListCertificateVersionsSortByEnum = map[string]ListCertificateVersionsSortByEnum{
	"VERSION_NUMBER": ListCertificateVersionsSortByVersionNumber,
}

var mappingListCertificateVersionsSortByEnumLowerCase = map[string]ListCertificateVersionsSortByEnum{
	"version_number": ListCertificateVersionsSortByVersionNumber,
}

// GetListCertificateVersionsSortByEnumValues Enumerates the set of values for ListCertificateVersionsSortByEnum
func GetListCertificateVersionsSortByEnumValues() []ListCertificateVersionsSortByEnum {
	values := make([]ListCertificateVersionsSortByEnum, 0)
	for _, v := range mappingListCertificateVersionsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListCertificateVersionsSortByEnumStringValues Enumerates the set of values in String for ListCertificateVersionsSortByEnum
func GetListCertificateVersionsSortByEnumStringValues() []string {
	return []string{
		"VERSION_NUMBER",
	}
}

// GetMappingListCertificateVersionsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListCertificateVersionsSortByEnum(val string) (ListCertificateVersionsSortByEnum, bool) {
	enum, ok := mappingListCertificateVersionsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListCertificateVersionsSortOrderEnum Enum with underlying type: string
type ListCertificateVersionsSortOrderEnum string

// Set of constants representing the allowable values for ListCertificateVersionsSortOrderEnum
const (
	ListCertificateVersionsSortOrderAsc  ListCertificateVersionsSortOrderEnum = "ASC"
	ListCertificateVersionsSortOrderDesc ListCertificateVersionsSortOrderEnum = "DESC"
)

var mappingListCertificateVersionsSortOrderEnum = map[string]ListCertificateVersionsSortOrderEnum{
	"ASC":  ListCertificateVersionsSortOrderAsc,
	"DESC": ListCertificateVersionsSortOrderDesc,
}

var mappingListCertificateVersionsSortOrderEnumLowerCase = map[string]ListCertificateVersionsSortOrderEnum{
	"asc":  ListCertificateVersionsSortOrderAsc,
	"desc": ListCertificateVersionsSortOrderDesc,
}

// GetListCertificateVersionsSortOrderEnumValues Enumerates the set of values for ListCertificateVersionsSortOrderEnum
func GetListCertificateVersionsSortOrderEnumValues() []ListCertificateVersionsSortOrderEnum {
	values := make([]ListCertificateVersionsSortOrderEnum, 0)
	for _, v := range mappingListCertificateVersionsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListCertificateVersionsSortOrderEnumStringValues Enumerates the set of values in String for ListCertificateVersionsSortOrderEnum
func GetListCertificateVersionsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListCertificateVersionsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListCertificateVersionsSortOrderEnum(val string) (ListCertificateVersionsSortOrderEnum, bool) {
	enum, ok := mappingListCertificateVersionsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
