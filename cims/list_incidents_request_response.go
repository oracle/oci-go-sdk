// Copyright (c) 2016, 2018, , Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package cims

import (
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

// ListIncidentsRequest wrapper for the ListIncidents operation
type ListIncidentsRequest struct {

	// Customer Support Identifier of the support account
	Csi *string `mandatory:"true" contributesTo:"header" name:"csi"`

	// Tenancy Ocid
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// User OCID for IDCS users that have a shadow in OCI
	Ocid *string `mandatory:"true" contributesTo:"header" name:"ocid"`

	// Limit query for number of returned results
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The key to sort the returned items by
	SortBy ListIncidentsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The order in which to sort the results
	SortOrder ListIncidentsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The order in which to sort the results
	LifecycleState ListIncidentsLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// Pagination for Incident list
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// Unique Header for request id
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListIncidentsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListIncidentsRequest) HTTPRequest(method, path string) (http.Request, error) {
	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListIncidentsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ListIncidentsResponse wrapper for the ListIncidents operation
type ListIncidentsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []IncidentSummary instances
	Items []IncidentSummary `presentIn:"body"`

	// OPC Request Id
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// OPC next page
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// e-Tag
	Etag *string `presentIn:"header" name:"etag"`
}

func (response ListIncidentsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListIncidentsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListIncidentsSortByEnum Enum with underlying type: string
type ListIncidentsSortByEnum string

// Set of constants representing the allowable values for ListIncidentsSortByEnum
const (
	ListIncidentsSortByDateupdated ListIncidentsSortByEnum = "dateUpdated"
	ListIncidentsSortBySeverity    ListIncidentsSortByEnum = "severity"
)

var mappingListIncidentsSortBy = map[string]ListIncidentsSortByEnum{
	"dateUpdated": ListIncidentsSortByDateupdated,
	"severity":    ListIncidentsSortBySeverity,
}

// GetListIncidentsSortByEnumValues Enumerates the set of values for ListIncidentsSortByEnum
func GetListIncidentsSortByEnumValues() []ListIncidentsSortByEnum {
	values := make([]ListIncidentsSortByEnum, 0)
	for _, v := range mappingListIncidentsSortBy {
		values = append(values, v)
	}
	return values
}

// ListIncidentsSortOrderEnum Enum with underlying type: string
type ListIncidentsSortOrderEnum string

// Set of constants representing the allowable values for ListIncidentsSortOrderEnum
const (
	ListIncidentsSortOrderAsc  ListIncidentsSortOrderEnum = "ASC"
	ListIncidentsSortOrderDesc ListIncidentsSortOrderEnum = "DESC"
)

var mappingListIncidentsSortOrder = map[string]ListIncidentsSortOrderEnum{
	"ASC":  ListIncidentsSortOrderAsc,
	"DESC": ListIncidentsSortOrderDesc,
}

// GetListIncidentsSortOrderEnumValues Enumerates the set of values for ListIncidentsSortOrderEnum
func GetListIncidentsSortOrderEnumValues() []ListIncidentsSortOrderEnum {
	values := make([]ListIncidentsSortOrderEnum, 0)
	for _, v := range mappingListIncidentsSortOrder {
		values = append(values, v)
	}
	return values
}

// ListIncidentsLifecycleStateEnum Enum with underlying type: string
type ListIncidentsLifecycleStateEnum string

// Set of constants representing the allowable values for ListIncidentsLifecycleStateEnum
const (
	ListIncidentsLifecycleStateActive ListIncidentsLifecycleStateEnum = "ACTIVE"
	ListIncidentsLifecycleStateClosed ListIncidentsLifecycleStateEnum = "CLOSED"
)

var mappingListIncidentsLifecycleState = map[string]ListIncidentsLifecycleStateEnum{
	"ACTIVE": ListIncidentsLifecycleStateActive,
	"CLOSED": ListIncidentsLifecycleStateClosed,
}

// GetListIncidentsLifecycleStateEnumValues Enumerates the set of values for ListIncidentsLifecycleStateEnum
func GetListIncidentsLifecycleStateEnumValues() []ListIncidentsLifecycleStateEnum {
	values := make([]ListIncidentsLifecycleStateEnum, 0)
	for _, v := range mappingListIncidentsLifecycleState {
		values = append(values, v)
	}
	return values
}
