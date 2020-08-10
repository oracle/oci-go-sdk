// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package cims

import (
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

// ListIncidentResourceTypesRequest wrapper for the ListIncidentResourceTypes operation
type ListIncidentResourceTypesRequest struct {

	// The kind of support request.
	ProblemType *string `mandatory:"true" contributesTo:"query" name:"problemType"`

	// The OCID of the tenancy.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// The Customer Support Identifier associated with the support account.
	Csi *string `mandatory:"true" contributesTo:"header" name:"csi"`

	// User OCID for Oracle Identity Cloud Service (IDCS) users who also have a federated Oracle Cloud Infrastructure account.
	Ocid *string `mandatory:"true" contributesTo:"header" name:"ocid"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// For list pagination. The maximum number of results per page, or items to return in a paginated "List" call. For important details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The value of the `opc-next-page` response header from the previous "List" call. For important details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The key to use to sort the returned items.
	SortBy ListIncidentResourceTypesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The order to sort the results in.
	SortOrder ListIncidentResourceTypesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The user-friendly name of the incident type.
	Name *string `mandatory:"false" contributesTo:"query" name:"name"`

	// The region of the tenancy.
	Homeregion *string `mandatory:"false" contributesTo:"header" name:"homeregion"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListIncidentResourceTypesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListIncidentResourceTypesRequest) HTTPRequest(method, path string) (http.Request, error) {
	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListIncidentResourceTypesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ListIncidentResourceTypesResponse wrapper for the ListIncidentResourceTypes operation
type ListIncidentResourceTypesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []IncidentResourceType instances
	Items []IncidentResourceType `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response, additional pages of results remain. For important details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListIncidentResourceTypesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListIncidentResourceTypesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListIncidentResourceTypesSortByEnum Enum with underlying type: string
type ListIncidentResourceTypesSortByEnum string

// Set of constants representing the allowable values for ListIncidentResourceTypesSortByEnum
const (
	ListIncidentResourceTypesSortByDateupdated ListIncidentResourceTypesSortByEnum = "dateUpdated"
	ListIncidentResourceTypesSortBySeverity    ListIncidentResourceTypesSortByEnum = "severity"
)

var mappingListIncidentResourceTypesSortBy = map[string]ListIncidentResourceTypesSortByEnum{
	"dateUpdated": ListIncidentResourceTypesSortByDateupdated,
	"severity":    ListIncidentResourceTypesSortBySeverity,
}

// GetListIncidentResourceTypesSortByEnumValues Enumerates the set of values for ListIncidentResourceTypesSortByEnum
func GetListIncidentResourceTypesSortByEnumValues() []ListIncidentResourceTypesSortByEnum {
	values := make([]ListIncidentResourceTypesSortByEnum, 0)
	for _, v := range mappingListIncidentResourceTypesSortBy {
		values = append(values, v)
	}
	return values
}

// ListIncidentResourceTypesSortOrderEnum Enum with underlying type: string
type ListIncidentResourceTypesSortOrderEnum string

// Set of constants representing the allowable values for ListIncidentResourceTypesSortOrderEnum
const (
	ListIncidentResourceTypesSortOrderAsc  ListIncidentResourceTypesSortOrderEnum = "ASC"
	ListIncidentResourceTypesSortOrderDesc ListIncidentResourceTypesSortOrderEnum = "DESC"
)

var mappingListIncidentResourceTypesSortOrder = map[string]ListIncidentResourceTypesSortOrderEnum{
	"ASC":  ListIncidentResourceTypesSortOrderAsc,
	"DESC": ListIncidentResourceTypesSortOrderDesc,
}

// GetListIncidentResourceTypesSortOrderEnumValues Enumerates the set of values for ListIncidentResourceTypesSortOrderEnum
func GetListIncidentResourceTypesSortOrderEnumValues() []ListIncidentResourceTypesSortOrderEnum {
	values := make([]ListIncidentResourceTypesSortOrderEnum, 0)
	for _, v := range mappingListIncidentResourceTypesSortOrder {
		values = append(values, v)
	}
	return values
}
