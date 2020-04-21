// Copyright (c) 2016, 2018, , Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package database

import (
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

// ListAutonomousDbVersionsRequest wrapper for the ListAutonomousDbVersions operation
type ListAutonomousDbVersionsRequest struct {

	// The compartment OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm).
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// The maximum number of items to return per page.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The pagination token to continue listing from.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// Unique identifier for the request.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// A filter to return only autonomous database resources that match the specified workload type.
	DbWorkload AutonomousDatabaseSummaryDbWorkloadEnum `mandatory:"false" contributesTo:"query" name:"dbWorkload" omitEmpty:"true"`

	// The sort order to use, either ascending (`ASC`) or descending (`DESC`).
	SortOrder ListAutonomousDbVersionsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListAutonomousDbVersionsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListAutonomousDbVersionsRequest) HTTPRequest(method, path string) (http.Request, error) {
	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListAutonomousDbVersionsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ListAutonomousDbVersionsResponse wrapper for the ListAutonomousDbVersions operation
type ListAutonomousDbVersionsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []AutonomousDbVersionSummary instances
	Items []AutonomousDbVersionSummary `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	// a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then there are additional items still to get. Include this value as the `page` parameter for the
	// subsequent GET request. For information about pagination, see
	// List Pagination (https://docs.cloud.oracle.com/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListAutonomousDbVersionsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListAutonomousDbVersionsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListAutonomousDbVersionsSortOrderEnum Enum with underlying type: string
type ListAutonomousDbVersionsSortOrderEnum string

// Set of constants representing the allowable values for ListAutonomousDbVersionsSortOrderEnum
const (
	ListAutonomousDbVersionsSortOrderAsc  ListAutonomousDbVersionsSortOrderEnum = "ASC"
	ListAutonomousDbVersionsSortOrderDesc ListAutonomousDbVersionsSortOrderEnum = "DESC"
)

var mappingListAutonomousDbVersionsSortOrder = map[string]ListAutonomousDbVersionsSortOrderEnum{
	"ASC":  ListAutonomousDbVersionsSortOrderAsc,
	"DESC": ListAutonomousDbVersionsSortOrderDesc,
}

// GetListAutonomousDbVersionsSortOrderEnumValues Enumerates the set of values for ListAutonomousDbVersionsSortOrderEnum
func GetListAutonomousDbVersionsSortOrderEnumValues() []ListAutonomousDbVersionsSortOrderEnum {
	values := make([]ListAutonomousDbVersionsSortOrderEnum, 0)
	for _, v := range mappingListAutonomousDbVersionsSortOrder {
		values = append(values, v)
	}
	return values
}
