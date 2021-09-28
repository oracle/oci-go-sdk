// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package applicationmigration

import (
	"github.com/oracle/oci-go-sdk/v48/common"
	"net/http"
)

// ListSourceApplicationsRequest wrapper for the ListSourceApplications operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/applicationmigration/ListSourceApplications.go.html to see an example of how to use ListSourceApplicationsRequest.
type ListSourceApplicationsRequest struct {

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the source.
	SourceId *string `mandatory:"true" contributesTo:"path" name:"sourceId"`

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of a compartment. Retrieves details of objects in the specified compartment.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a
	// particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// The number of items returned in a paginated `List` call. For information about pagination, see
	// List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The value of the `opc-next-page` response header from the preceding `List` call.
	// For information about pagination, see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order, either `ASC` (ascending) or `DESC` (descending).
	SortOrder ListSourceApplicationsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// Specifies the field on which to sort.
	// By default, `TIMECREATED` is ordered descending.
	// By default, `DISPLAYNAME` is ordered ascending. Note that you can sort only on one field.
	SortBy ListSourceApplicationsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Resource name on which to query.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListSourceApplicationsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListSourceApplicationsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListSourceApplicationsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListSourceApplicationsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ListSourceApplicationsResponse wrapper for the ListSourceApplications operation
type ListSourceApplicationsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []SourceApplicationSummary instances
	Items []SourceApplicationSummary `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response, additional pages of results remain.
	// For details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	// Include this value as the `page` parameter for the subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListSourceApplicationsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListSourceApplicationsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListSourceApplicationsSortOrderEnum Enum with underlying type: string
type ListSourceApplicationsSortOrderEnum string

// Set of constants representing the allowable values for ListSourceApplicationsSortOrderEnum
const (
	ListSourceApplicationsSortOrderAsc  ListSourceApplicationsSortOrderEnum = "ASC"
	ListSourceApplicationsSortOrderDesc ListSourceApplicationsSortOrderEnum = "DESC"
)

var mappingListSourceApplicationsSortOrder = map[string]ListSourceApplicationsSortOrderEnum{
	"ASC":  ListSourceApplicationsSortOrderAsc,
	"DESC": ListSourceApplicationsSortOrderDesc,
}

// GetListSourceApplicationsSortOrderEnumValues Enumerates the set of values for ListSourceApplicationsSortOrderEnum
func GetListSourceApplicationsSortOrderEnumValues() []ListSourceApplicationsSortOrderEnum {
	values := make([]ListSourceApplicationsSortOrderEnum, 0)
	for _, v := range mappingListSourceApplicationsSortOrder {
		values = append(values, v)
	}
	return values
}

// ListSourceApplicationsSortByEnum Enum with underlying type: string
type ListSourceApplicationsSortByEnum string

// Set of constants representing the allowable values for ListSourceApplicationsSortByEnum
const (
	ListSourceApplicationsSortByTimecreated ListSourceApplicationsSortByEnum = "TIMECREATED"
	ListSourceApplicationsSortByDisplayname ListSourceApplicationsSortByEnum = "DISPLAYNAME"
)

var mappingListSourceApplicationsSortBy = map[string]ListSourceApplicationsSortByEnum{
	"TIMECREATED": ListSourceApplicationsSortByTimecreated,
	"DISPLAYNAME": ListSourceApplicationsSortByDisplayname,
}

// GetListSourceApplicationsSortByEnumValues Enumerates the set of values for ListSourceApplicationsSortByEnum
func GetListSourceApplicationsSortByEnumValues() []ListSourceApplicationsSortByEnum {
	values := make([]ListSourceApplicationsSortByEnum, 0)
	for _, v := range mappingListSourceApplicationsSortBy {
		values = append(values, v)
	}
	return values
}
