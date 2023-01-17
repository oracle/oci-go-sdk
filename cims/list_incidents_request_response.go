// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package cims

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListIncidentsRequest wrapper for the ListIncidents operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/cims/ListIncidents.go.html to see an example of how to use ListIncidentsRequest.
type ListIncidentsRequest struct {

	// The Customer Support Identifier associated with the support account.
	Csi *string `mandatory:"true" contributesTo:"header" name:"csi"`

	// The OCID of the tenancy.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// User OCID for Oracle Identity Cloud Service (IDCS) users who also have a federated Oracle Cloud Infrastructure account.
	Ocid *string `mandatory:"true" contributesTo:"header" name:"ocid"`

	// For list pagination. The maximum number of results per page, or items to return in a paginated "List" call. For important details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The key to use to sort the returned items.
	SortBy ListIncidentsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The order to sort the results in.
	SortOrder ListIncidentsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The current state of the ticket.
	LifecycleState ListIncidentsLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// For list pagination. The value of the `opc-next-page` response header from the previous "List" call. For important details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// The region of the tenancy.
	Homeregion *string `mandatory:"false" contributesTo:"header" name:"homeregion"`

	// The kind of support request.
	ProblemType *string `mandatory:"false" contributesTo:"query" name:"problemType"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListIncidentsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListIncidentsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListIncidentsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListIncidentsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListIncidentsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListIncidentsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListIncidentsSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListIncidentsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListIncidentsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListIncidentsLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetListIncidentsLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListIncidentsResponse wrapper for the ListIncidents operation
type ListIncidentsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []IncidentSummary instances
	Items []IncidentSummary `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response, additional pages of results remain. For important details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
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

var mappingListIncidentsSortByEnum = map[string]ListIncidentsSortByEnum{
	"dateUpdated": ListIncidentsSortByDateupdated,
	"severity":    ListIncidentsSortBySeverity,
}

var mappingListIncidentsSortByEnumLowerCase = map[string]ListIncidentsSortByEnum{
	"dateupdated": ListIncidentsSortByDateupdated,
	"severity":    ListIncidentsSortBySeverity,
}

// GetListIncidentsSortByEnumValues Enumerates the set of values for ListIncidentsSortByEnum
func GetListIncidentsSortByEnumValues() []ListIncidentsSortByEnum {
	values := make([]ListIncidentsSortByEnum, 0)
	for _, v := range mappingListIncidentsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListIncidentsSortByEnumStringValues Enumerates the set of values in String for ListIncidentsSortByEnum
func GetListIncidentsSortByEnumStringValues() []string {
	return []string{
		"dateUpdated",
		"severity",
	}
}

// GetMappingListIncidentsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListIncidentsSortByEnum(val string) (ListIncidentsSortByEnum, bool) {
	enum, ok := mappingListIncidentsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListIncidentsSortOrderEnum Enum with underlying type: string
type ListIncidentsSortOrderEnum string

// Set of constants representing the allowable values for ListIncidentsSortOrderEnum
const (
	ListIncidentsSortOrderAsc  ListIncidentsSortOrderEnum = "ASC"
	ListIncidentsSortOrderDesc ListIncidentsSortOrderEnum = "DESC"
)

var mappingListIncidentsSortOrderEnum = map[string]ListIncidentsSortOrderEnum{
	"ASC":  ListIncidentsSortOrderAsc,
	"DESC": ListIncidentsSortOrderDesc,
}

var mappingListIncidentsSortOrderEnumLowerCase = map[string]ListIncidentsSortOrderEnum{
	"asc":  ListIncidentsSortOrderAsc,
	"desc": ListIncidentsSortOrderDesc,
}

// GetListIncidentsSortOrderEnumValues Enumerates the set of values for ListIncidentsSortOrderEnum
func GetListIncidentsSortOrderEnumValues() []ListIncidentsSortOrderEnum {
	values := make([]ListIncidentsSortOrderEnum, 0)
	for _, v := range mappingListIncidentsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListIncidentsSortOrderEnumStringValues Enumerates the set of values in String for ListIncidentsSortOrderEnum
func GetListIncidentsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListIncidentsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListIncidentsSortOrderEnum(val string) (ListIncidentsSortOrderEnum, bool) {
	enum, ok := mappingListIncidentsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListIncidentsLifecycleStateEnum Enum with underlying type: string
type ListIncidentsLifecycleStateEnum string

// Set of constants representing the allowable values for ListIncidentsLifecycleStateEnum
const (
	ListIncidentsLifecycleStateActive ListIncidentsLifecycleStateEnum = "ACTIVE"
	ListIncidentsLifecycleStateClosed ListIncidentsLifecycleStateEnum = "CLOSED"
)

var mappingListIncidentsLifecycleStateEnum = map[string]ListIncidentsLifecycleStateEnum{
	"ACTIVE": ListIncidentsLifecycleStateActive,
	"CLOSED": ListIncidentsLifecycleStateClosed,
}

var mappingListIncidentsLifecycleStateEnumLowerCase = map[string]ListIncidentsLifecycleStateEnum{
	"active": ListIncidentsLifecycleStateActive,
	"closed": ListIncidentsLifecycleStateClosed,
}

// GetListIncidentsLifecycleStateEnumValues Enumerates the set of values for ListIncidentsLifecycleStateEnum
func GetListIncidentsLifecycleStateEnumValues() []ListIncidentsLifecycleStateEnum {
	values := make([]ListIncidentsLifecycleStateEnum, 0)
	for _, v := range mappingListIncidentsLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetListIncidentsLifecycleStateEnumStringValues Enumerates the set of values in String for ListIncidentsLifecycleStateEnum
func GetListIncidentsLifecycleStateEnumStringValues() []string {
	return []string{
		"ACTIVE",
		"CLOSED",
	}
}

// GetMappingListIncidentsLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListIncidentsLifecycleStateEnum(val string) (ListIncidentsLifecycleStateEnum, bool) {
	enum, ok := mappingListIncidentsLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
