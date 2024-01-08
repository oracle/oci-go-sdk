// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package accessgovernancecp

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListGovernanceInstancesRequest wrapper for the ListGovernanceInstances operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/accessgovernancecp/ListGovernanceInstances.go.html to see an example of how to use ListGovernanceInstancesRequest.
type ListGovernanceInstancesRequest struct {

	// The OCID of the compartment in which resources are listed.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// The lifecycle state to filter on.
	LifecycleState *string `mandatory:"false" contributesTo:"query" name:"lifecycleState"`

	// A filter to return only resources that match the entire display name given.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// The OCID of the GovernanceInstance
	Id *string `mandatory:"false" contributesTo:"query" name:"id"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// A token representing the position at which to start retrieving results. This must come from the `opc-next-page` header field of a previous response.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'ASC' or 'DESC'.
	SortOrder ListGovernanceInstancesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for timeCreated is descending.
	SortBy ListGovernanceInstancesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListGovernanceInstancesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListGovernanceInstancesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListGovernanceInstancesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListGovernanceInstancesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListGovernanceInstancesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListGovernanceInstancesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListGovernanceInstancesSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListGovernanceInstancesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListGovernanceInstancesSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListGovernanceInstancesResponse wrapper for the ListGovernanceInstances operation
type ListGovernanceInstancesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of GovernanceInstanceCollection instances
	GovernanceInstanceCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListGovernanceInstancesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListGovernanceInstancesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListGovernanceInstancesSortOrderEnum Enum with underlying type: string
type ListGovernanceInstancesSortOrderEnum string

// Set of constants representing the allowable values for ListGovernanceInstancesSortOrderEnum
const (
	ListGovernanceInstancesSortOrderAsc  ListGovernanceInstancesSortOrderEnum = "ASC"
	ListGovernanceInstancesSortOrderDesc ListGovernanceInstancesSortOrderEnum = "DESC"
)

var mappingListGovernanceInstancesSortOrderEnum = map[string]ListGovernanceInstancesSortOrderEnum{
	"ASC":  ListGovernanceInstancesSortOrderAsc,
	"DESC": ListGovernanceInstancesSortOrderDesc,
}

var mappingListGovernanceInstancesSortOrderEnumLowerCase = map[string]ListGovernanceInstancesSortOrderEnum{
	"asc":  ListGovernanceInstancesSortOrderAsc,
	"desc": ListGovernanceInstancesSortOrderDesc,
}

// GetListGovernanceInstancesSortOrderEnumValues Enumerates the set of values for ListGovernanceInstancesSortOrderEnum
func GetListGovernanceInstancesSortOrderEnumValues() []ListGovernanceInstancesSortOrderEnum {
	values := make([]ListGovernanceInstancesSortOrderEnum, 0)
	for _, v := range mappingListGovernanceInstancesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListGovernanceInstancesSortOrderEnumStringValues Enumerates the set of values in String for ListGovernanceInstancesSortOrderEnum
func GetListGovernanceInstancesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListGovernanceInstancesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListGovernanceInstancesSortOrderEnum(val string) (ListGovernanceInstancesSortOrderEnum, bool) {
	enum, ok := mappingListGovernanceInstancesSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListGovernanceInstancesSortByEnum Enum with underlying type: string
type ListGovernanceInstancesSortByEnum string

// Set of constants representing the allowable values for ListGovernanceInstancesSortByEnum
const (
	ListGovernanceInstancesSortByTimecreated    ListGovernanceInstancesSortByEnum = "timeCreated"
	ListGovernanceInstancesSortByDisplayname    ListGovernanceInstancesSortByEnum = "displayName"
	ListGovernanceInstancesSortByTimeupdated    ListGovernanceInstancesSortByEnum = "timeUpdated"
	ListGovernanceInstancesSortByLifecyclestate ListGovernanceInstancesSortByEnum = "lifecycleState"
)

var mappingListGovernanceInstancesSortByEnum = map[string]ListGovernanceInstancesSortByEnum{
	"timeCreated":    ListGovernanceInstancesSortByTimecreated,
	"displayName":    ListGovernanceInstancesSortByDisplayname,
	"timeUpdated":    ListGovernanceInstancesSortByTimeupdated,
	"lifecycleState": ListGovernanceInstancesSortByLifecyclestate,
}

var mappingListGovernanceInstancesSortByEnumLowerCase = map[string]ListGovernanceInstancesSortByEnum{
	"timecreated":    ListGovernanceInstancesSortByTimecreated,
	"displayname":    ListGovernanceInstancesSortByDisplayname,
	"timeupdated":    ListGovernanceInstancesSortByTimeupdated,
	"lifecyclestate": ListGovernanceInstancesSortByLifecyclestate,
}

// GetListGovernanceInstancesSortByEnumValues Enumerates the set of values for ListGovernanceInstancesSortByEnum
func GetListGovernanceInstancesSortByEnumValues() []ListGovernanceInstancesSortByEnum {
	values := make([]ListGovernanceInstancesSortByEnum, 0)
	for _, v := range mappingListGovernanceInstancesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListGovernanceInstancesSortByEnumStringValues Enumerates the set of values in String for ListGovernanceInstancesSortByEnum
func GetListGovernanceInstancesSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
		"timeUpdated",
		"lifecycleState",
	}
}

// GetMappingListGovernanceInstancesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListGovernanceInstancesSortByEnum(val string) (ListGovernanceInstancesSortByEnum, bool) {
	enum, ok := mappingListGovernanceInstancesSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
