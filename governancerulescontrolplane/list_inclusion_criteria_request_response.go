// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package governancerulescontrolplane

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListInclusionCriteriaRequest wrapper for the ListInclusionCriteria operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/governancerulescontrolplane/ListInclusionCriteria.go.html to see an example of how to use ListInclusionCriteriaRequest.
type ListInclusionCriteriaRequest struct {

	// Unique governance rule identifier.
	GovernanceRuleId *string `mandatory:"true" contributesTo:"query" name:"governanceRuleId"`

	// Unique inclusion criterion identifier.
	InclusionCriterionId *string `mandatory:"false" contributesTo:"query" name:"inclusionCriterionId"`

	// A filter to return only resources when their lifecycle state matches the given lifecycle state.
	LifecycleState InclusionCriterionLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// A token representing the position at which to start retrieving results. This must come from the `opc-next-page` header field of a previous response.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'ASC' or 'DESC'.
	SortOrder ListInclusionCriteriaSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for timeCreated is descending. Default order for displayName is ascending.
	SortBy ListInclusionCriteriaSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListInclusionCriteriaRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListInclusionCriteriaRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListInclusionCriteriaRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListInclusionCriteriaRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListInclusionCriteriaRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingInclusionCriterionLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetInclusionCriterionLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListInclusionCriteriaSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListInclusionCriteriaSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListInclusionCriteriaSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListInclusionCriteriaSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListInclusionCriteriaResponse wrapper for the ListInclusionCriteria operation
type ListInclusionCriteriaResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of InclusionCriterionCollection instances
	InclusionCriterionCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListInclusionCriteriaResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListInclusionCriteriaResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListInclusionCriteriaSortOrderEnum Enum with underlying type: string
type ListInclusionCriteriaSortOrderEnum string

// Set of constants representing the allowable values for ListInclusionCriteriaSortOrderEnum
const (
	ListInclusionCriteriaSortOrderAsc  ListInclusionCriteriaSortOrderEnum = "ASC"
	ListInclusionCriteriaSortOrderDesc ListInclusionCriteriaSortOrderEnum = "DESC"
)

var mappingListInclusionCriteriaSortOrderEnum = map[string]ListInclusionCriteriaSortOrderEnum{
	"ASC":  ListInclusionCriteriaSortOrderAsc,
	"DESC": ListInclusionCriteriaSortOrderDesc,
}

var mappingListInclusionCriteriaSortOrderEnumLowerCase = map[string]ListInclusionCriteriaSortOrderEnum{
	"asc":  ListInclusionCriteriaSortOrderAsc,
	"desc": ListInclusionCriteriaSortOrderDesc,
}

// GetListInclusionCriteriaSortOrderEnumValues Enumerates the set of values for ListInclusionCriteriaSortOrderEnum
func GetListInclusionCriteriaSortOrderEnumValues() []ListInclusionCriteriaSortOrderEnum {
	values := make([]ListInclusionCriteriaSortOrderEnum, 0)
	for _, v := range mappingListInclusionCriteriaSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListInclusionCriteriaSortOrderEnumStringValues Enumerates the set of values in String for ListInclusionCriteriaSortOrderEnum
func GetListInclusionCriteriaSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListInclusionCriteriaSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListInclusionCriteriaSortOrderEnum(val string) (ListInclusionCriteriaSortOrderEnum, bool) {
	enum, ok := mappingListInclusionCriteriaSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListInclusionCriteriaSortByEnum Enum with underlying type: string
type ListInclusionCriteriaSortByEnum string

// Set of constants representing the allowable values for ListInclusionCriteriaSortByEnum
const (
	ListInclusionCriteriaSortByTimecreated ListInclusionCriteriaSortByEnum = "timeCreated"
	ListInclusionCriteriaSortByDisplayname ListInclusionCriteriaSortByEnum = "displayName"
)

var mappingListInclusionCriteriaSortByEnum = map[string]ListInclusionCriteriaSortByEnum{
	"timeCreated": ListInclusionCriteriaSortByTimecreated,
	"displayName": ListInclusionCriteriaSortByDisplayname,
}

var mappingListInclusionCriteriaSortByEnumLowerCase = map[string]ListInclusionCriteriaSortByEnum{
	"timecreated": ListInclusionCriteriaSortByTimecreated,
	"displayname": ListInclusionCriteriaSortByDisplayname,
}

// GetListInclusionCriteriaSortByEnumValues Enumerates the set of values for ListInclusionCriteriaSortByEnum
func GetListInclusionCriteriaSortByEnumValues() []ListInclusionCriteriaSortByEnum {
	values := make([]ListInclusionCriteriaSortByEnum, 0)
	for _, v := range mappingListInclusionCriteriaSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListInclusionCriteriaSortByEnumStringValues Enumerates the set of values in String for ListInclusionCriteriaSortByEnum
func GetListInclusionCriteriaSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListInclusionCriteriaSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListInclusionCriteriaSortByEnum(val string) (ListInclusionCriteriaSortByEnum, bool) {
	enum, ok := mappingListInclusionCriteriaSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
