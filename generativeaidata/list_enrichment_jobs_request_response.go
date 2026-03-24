// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package generativeaidata

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListEnrichmentJobsRequest wrapper for the ListEnrichmentJobs operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/generativeaidata/ListEnrichmentJobs.go.html to see an example of how to use ListEnrichmentJobsRequest.
type ListEnrichmentJobsRequest struct {

	// The OCID of the semantic store
	SemanticStoreId *string `mandatory:"true" contributesTo:"path" name:"semanticStoreId"`

	// The ID of the compartment in which to list resources.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// A filter to return only resources that match the entire display name given.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// A filter to return only resources their lifecycleState matches the given lifecycleState.
	LifecycleState ListEnrichmentJobsLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// A token representing the position at which to start retrieving results. This must come from the `opc-next-page` header field of a previous response.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The field to sort by. Only one sort order may be provided. Default order for timeCreated is descending. Default order for displayName is ascending.
	SortBy ListEnrichmentJobsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to use, either 'ASC' or 'DESC'.
	SortOrder ListEnrichmentJobsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListEnrichmentJobsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListEnrichmentJobsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListEnrichmentJobsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListEnrichmentJobsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListEnrichmentJobsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListEnrichmentJobsLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetListEnrichmentJobsLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListEnrichmentJobsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListEnrichmentJobsSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListEnrichmentJobsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListEnrichmentJobsSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListEnrichmentJobsResponse wrapper for the ListEnrichmentJobs operation
type ListEnrichmentJobsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of EnrichmentJobCollection instances
	EnrichmentJobCollection `presentIn:"body"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// For list pagination. When this header appears in the response, previous pages of results remain. For
	// important details about how pagination works, see List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcPrevPage *string `presentIn:"header" name:"opc-prev-page"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response ListEnrichmentJobsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListEnrichmentJobsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListEnrichmentJobsLifecycleStateEnum Enum with underlying type: string
type ListEnrichmentJobsLifecycleStateEnum string

// Set of constants representing the allowable values for ListEnrichmentJobsLifecycleStateEnum
const (
	ListEnrichmentJobsLifecycleStateAccepted   ListEnrichmentJobsLifecycleStateEnum = "ACCEPTED"
	ListEnrichmentJobsLifecycleStateInProgress ListEnrichmentJobsLifecycleStateEnum = "IN_PROGRESS"
	ListEnrichmentJobsLifecycleStateFailed     ListEnrichmentJobsLifecycleStateEnum = "FAILED"
	ListEnrichmentJobsLifecycleStateSucceeded  ListEnrichmentJobsLifecycleStateEnum = "SUCCEEDED"
	ListEnrichmentJobsLifecycleStateCanceling  ListEnrichmentJobsLifecycleStateEnum = "CANCELING"
	ListEnrichmentJobsLifecycleStateCanceled   ListEnrichmentJobsLifecycleStateEnum = "CANCELED"
)

var mappingListEnrichmentJobsLifecycleStateEnum = map[string]ListEnrichmentJobsLifecycleStateEnum{
	"ACCEPTED":    ListEnrichmentJobsLifecycleStateAccepted,
	"IN_PROGRESS": ListEnrichmentJobsLifecycleStateInProgress,
	"FAILED":      ListEnrichmentJobsLifecycleStateFailed,
	"SUCCEEDED":   ListEnrichmentJobsLifecycleStateSucceeded,
	"CANCELING":   ListEnrichmentJobsLifecycleStateCanceling,
	"CANCELED":    ListEnrichmentJobsLifecycleStateCanceled,
}

var mappingListEnrichmentJobsLifecycleStateEnumLowerCase = map[string]ListEnrichmentJobsLifecycleStateEnum{
	"accepted":    ListEnrichmentJobsLifecycleStateAccepted,
	"in_progress": ListEnrichmentJobsLifecycleStateInProgress,
	"failed":      ListEnrichmentJobsLifecycleStateFailed,
	"succeeded":   ListEnrichmentJobsLifecycleStateSucceeded,
	"canceling":   ListEnrichmentJobsLifecycleStateCanceling,
	"canceled":    ListEnrichmentJobsLifecycleStateCanceled,
}

// GetListEnrichmentJobsLifecycleStateEnumValues Enumerates the set of values for ListEnrichmentJobsLifecycleStateEnum
func GetListEnrichmentJobsLifecycleStateEnumValues() []ListEnrichmentJobsLifecycleStateEnum {
	values := make([]ListEnrichmentJobsLifecycleStateEnum, 0)
	for _, v := range mappingListEnrichmentJobsLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetListEnrichmentJobsLifecycleStateEnumStringValues Enumerates the set of values in String for ListEnrichmentJobsLifecycleStateEnum
func GetListEnrichmentJobsLifecycleStateEnumStringValues() []string {
	return []string{
		"ACCEPTED",
		"IN_PROGRESS",
		"FAILED",
		"SUCCEEDED",
		"CANCELING",
		"CANCELED",
	}
}

// GetMappingListEnrichmentJobsLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListEnrichmentJobsLifecycleStateEnum(val string) (ListEnrichmentJobsLifecycleStateEnum, bool) {
	enum, ok := mappingListEnrichmentJobsLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListEnrichmentJobsSortByEnum Enum with underlying type: string
type ListEnrichmentJobsSortByEnum string

// Set of constants representing the allowable values for ListEnrichmentJobsSortByEnum
const (
	ListEnrichmentJobsSortByTimecreated ListEnrichmentJobsSortByEnum = "timeCreated"
	ListEnrichmentJobsSortByDisplayname ListEnrichmentJobsSortByEnum = "displayName"
)

var mappingListEnrichmentJobsSortByEnum = map[string]ListEnrichmentJobsSortByEnum{
	"timeCreated": ListEnrichmentJobsSortByTimecreated,
	"displayName": ListEnrichmentJobsSortByDisplayname,
}

var mappingListEnrichmentJobsSortByEnumLowerCase = map[string]ListEnrichmentJobsSortByEnum{
	"timecreated": ListEnrichmentJobsSortByTimecreated,
	"displayname": ListEnrichmentJobsSortByDisplayname,
}

// GetListEnrichmentJobsSortByEnumValues Enumerates the set of values for ListEnrichmentJobsSortByEnum
func GetListEnrichmentJobsSortByEnumValues() []ListEnrichmentJobsSortByEnum {
	values := make([]ListEnrichmentJobsSortByEnum, 0)
	for _, v := range mappingListEnrichmentJobsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListEnrichmentJobsSortByEnumStringValues Enumerates the set of values in String for ListEnrichmentJobsSortByEnum
func GetListEnrichmentJobsSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListEnrichmentJobsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListEnrichmentJobsSortByEnum(val string) (ListEnrichmentJobsSortByEnum, bool) {
	enum, ok := mappingListEnrichmentJobsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListEnrichmentJobsSortOrderEnum Enum with underlying type: string
type ListEnrichmentJobsSortOrderEnum string

// Set of constants representing the allowable values for ListEnrichmentJobsSortOrderEnum
const (
	ListEnrichmentJobsSortOrderAsc  ListEnrichmentJobsSortOrderEnum = "ASC"
	ListEnrichmentJobsSortOrderDesc ListEnrichmentJobsSortOrderEnum = "DESC"
)

var mappingListEnrichmentJobsSortOrderEnum = map[string]ListEnrichmentJobsSortOrderEnum{
	"ASC":  ListEnrichmentJobsSortOrderAsc,
	"DESC": ListEnrichmentJobsSortOrderDesc,
}

var mappingListEnrichmentJobsSortOrderEnumLowerCase = map[string]ListEnrichmentJobsSortOrderEnum{
	"asc":  ListEnrichmentJobsSortOrderAsc,
	"desc": ListEnrichmentJobsSortOrderDesc,
}

// GetListEnrichmentJobsSortOrderEnumValues Enumerates the set of values for ListEnrichmentJobsSortOrderEnum
func GetListEnrichmentJobsSortOrderEnumValues() []ListEnrichmentJobsSortOrderEnum {
	values := make([]ListEnrichmentJobsSortOrderEnum, 0)
	for _, v := range mappingListEnrichmentJobsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListEnrichmentJobsSortOrderEnumStringValues Enumerates the set of values in String for ListEnrichmentJobsSortOrderEnum
func GetListEnrichmentJobsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListEnrichmentJobsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListEnrichmentJobsSortOrderEnum(val string) (ListEnrichmentJobsSortOrderEnum, bool) {
	enum, ok := mappingListEnrichmentJobsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
