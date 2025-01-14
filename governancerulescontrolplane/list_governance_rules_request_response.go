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

// ListGovernanceRulesRequest wrapper for the ListGovernanceRules operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/governancerulescontrolplane/ListGovernanceRules.go.html to see an example of how to use ListGovernanceRulesRequest.
type ListGovernanceRulesRequest struct {

	// The ID of the compartment in which to list resources.
	CompartmentId *string `mandatory:"false" contributesTo:"query" name:"compartmentId"`

	// Unique governance rule identifier.
	GovernanceRuleId *string `mandatory:"false" contributesTo:"query" name:"governanceRuleId"`

	// A filter to return only resources whose lifecycle state matches the given lifecycle state.
	LifecycleState ListGovernanceRulesLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// A filter to return only resources that match the entire name given.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// A filter to return only resources that match the type given.
	GovernanceRuleType ListGovernanceRulesGovernanceRuleTypeEnum `mandatory:"false" contributesTo:"query" name:"governanceRuleType" omitEmpty:"true"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// A token representing the position at which to start retrieving results. This must come from the `opc-next-page` header field of a previous response.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'ASC' or 'DESC'.
	SortOrder ListGovernanceRulesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for timeCreated is descending. Default order for displayName is ascending.
	SortBy ListGovernanceRulesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListGovernanceRulesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListGovernanceRulesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListGovernanceRulesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListGovernanceRulesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListGovernanceRulesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListGovernanceRulesLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetListGovernanceRulesLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListGovernanceRulesGovernanceRuleTypeEnum(string(request.GovernanceRuleType)); !ok && request.GovernanceRuleType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for GovernanceRuleType: %s. Supported values are: %s.", request.GovernanceRuleType, strings.Join(GetListGovernanceRulesGovernanceRuleTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListGovernanceRulesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListGovernanceRulesSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListGovernanceRulesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListGovernanceRulesSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListGovernanceRulesResponse wrapper for the ListGovernanceRules operation
type ListGovernanceRulesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of GovernanceRuleCollection instances
	GovernanceRuleCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListGovernanceRulesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListGovernanceRulesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListGovernanceRulesLifecycleStateEnum Enum with underlying type: string
type ListGovernanceRulesLifecycleStateEnum string

// Set of constants representing the allowable values for ListGovernanceRulesLifecycleStateEnum
const (
	ListGovernanceRulesLifecycleStateActive  ListGovernanceRulesLifecycleStateEnum = "ACTIVE"
	ListGovernanceRulesLifecycleStateDeleted ListGovernanceRulesLifecycleStateEnum = "DELETED"
)

var mappingListGovernanceRulesLifecycleStateEnum = map[string]ListGovernanceRulesLifecycleStateEnum{
	"ACTIVE":  ListGovernanceRulesLifecycleStateActive,
	"DELETED": ListGovernanceRulesLifecycleStateDeleted,
}

var mappingListGovernanceRulesLifecycleStateEnumLowerCase = map[string]ListGovernanceRulesLifecycleStateEnum{
	"active":  ListGovernanceRulesLifecycleStateActive,
	"deleted": ListGovernanceRulesLifecycleStateDeleted,
}

// GetListGovernanceRulesLifecycleStateEnumValues Enumerates the set of values for ListGovernanceRulesLifecycleStateEnum
func GetListGovernanceRulesLifecycleStateEnumValues() []ListGovernanceRulesLifecycleStateEnum {
	values := make([]ListGovernanceRulesLifecycleStateEnum, 0)
	for _, v := range mappingListGovernanceRulesLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetListGovernanceRulesLifecycleStateEnumStringValues Enumerates the set of values in String for ListGovernanceRulesLifecycleStateEnum
func GetListGovernanceRulesLifecycleStateEnumStringValues() []string {
	return []string{
		"ACTIVE",
		"DELETED",
	}
}

// GetMappingListGovernanceRulesLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListGovernanceRulesLifecycleStateEnum(val string) (ListGovernanceRulesLifecycleStateEnum, bool) {
	enum, ok := mappingListGovernanceRulesLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListGovernanceRulesGovernanceRuleTypeEnum Enum with underlying type: string
type ListGovernanceRulesGovernanceRuleTypeEnum string

// Set of constants representing the allowable values for ListGovernanceRulesGovernanceRuleTypeEnum
const (
	ListGovernanceRulesGovernanceRuleTypeQuota          ListGovernanceRulesGovernanceRuleTypeEnum = "QUOTA"
	ListGovernanceRulesGovernanceRuleTypeTag            ListGovernanceRulesGovernanceRuleTypeEnum = "TAG"
	ListGovernanceRulesGovernanceRuleTypeAllowedRegions ListGovernanceRulesGovernanceRuleTypeEnum = "ALLOWED_REGIONS"
)

var mappingListGovernanceRulesGovernanceRuleTypeEnum = map[string]ListGovernanceRulesGovernanceRuleTypeEnum{
	"QUOTA":           ListGovernanceRulesGovernanceRuleTypeQuota,
	"TAG":             ListGovernanceRulesGovernanceRuleTypeTag,
	"ALLOWED_REGIONS": ListGovernanceRulesGovernanceRuleTypeAllowedRegions,
}

var mappingListGovernanceRulesGovernanceRuleTypeEnumLowerCase = map[string]ListGovernanceRulesGovernanceRuleTypeEnum{
	"quota":           ListGovernanceRulesGovernanceRuleTypeQuota,
	"tag":             ListGovernanceRulesGovernanceRuleTypeTag,
	"allowed_regions": ListGovernanceRulesGovernanceRuleTypeAllowedRegions,
}

// GetListGovernanceRulesGovernanceRuleTypeEnumValues Enumerates the set of values for ListGovernanceRulesGovernanceRuleTypeEnum
func GetListGovernanceRulesGovernanceRuleTypeEnumValues() []ListGovernanceRulesGovernanceRuleTypeEnum {
	values := make([]ListGovernanceRulesGovernanceRuleTypeEnum, 0)
	for _, v := range mappingListGovernanceRulesGovernanceRuleTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetListGovernanceRulesGovernanceRuleTypeEnumStringValues Enumerates the set of values in String for ListGovernanceRulesGovernanceRuleTypeEnum
func GetListGovernanceRulesGovernanceRuleTypeEnumStringValues() []string {
	return []string{
		"QUOTA",
		"TAG",
		"ALLOWED_REGIONS",
	}
}

// GetMappingListGovernanceRulesGovernanceRuleTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListGovernanceRulesGovernanceRuleTypeEnum(val string) (ListGovernanceRulesGovernanceRuleTypeEnum, bool) {
	enum, ok := mappingListGovernanceRulesGovernanceRuleTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListGovernanceRulesSortOrderEnum Enum with underlying type: string
type ListGovernanceRulesSortOrderEnum string

// Set of constants representing the allowable values for ListGovernanceRulesSortOrderEnum
const (
	ListGovernanceRulesSortOrderAsc  ListGovernanceRulesSortOrderEnum = "ASC"
	ListGovernanceRulesSortOrderDesc ListGovernanceRulesSortOrderEnum = "DESC"
)

var mappingListGovernanceRulesSortOrderEnum = map[string]ListGovernanceRulesSortOrderEnum{
	"ASC":  ListGovernanceRulesSortOrderAsc,
	"DESC": ListGovernanceRulesSortOrderDesc,
}

var mappingListGovernanceRulesSortOrderEnumLowerCase = map[string]ListGovernanceRulesSortOrderEnum{
	"asc":  ListGovernanceRulesSortOrderAsc,
	"desc": ListGovernanceRulesSortOrderDesc,
}

// GetListGovernanceRulesSortOrderEnumValues Enumerates the set of values for ListGovernanceRulesSortOrderEnum
func GetListGovernanceRulesSortOrderEnumValues() []ListGovernanceRulesSortOrderEnum {
	values := make([]ListGovernanceRulesSortOrderEnum, 0)
	for _, v := range mappingListGovernanceRulesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListGovernanceRulesSortOrderEnumStringValues Enumerates the set of values in String for ListGovernanceRulesSortOrderEnum
func GetListGovernanceRulesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListGovernanceRulesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListGovernanceRulesSortOrderEnum(val string) (ListGovernanceRulesSortOrderEnum, bool) {
	enum, ok := mappingListGovernanceRulesSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListGovernanceRulesSortByEnum Enum with underlying type: string
type ListGovernanceRulesSortByEnum string

// Set of constants representing the allowable values for ListGovernanceRulesSortByEnum
const (
	ListGovernanceRulesSortByTimecreated ListGovernanceRulesSortByEnum = "timeCreated"
	ListGovernanceRulesSortByDisplayname ListGovernanceRulesSortByEnum = "displayName"
)

var mappingListGovernanceRulesSortByEnum = map[string]ListGovernanceRulesSortByEnum{
	"timeCreated": ListGovernanceRulesSortByTimecreated,
	"displayName": ListGovernanceRulesSortByDisplayname,
}

var mappingListGovernanceRulesSortByEnumLowerCase = map[string]ListGovernanceRulesSortByEnum{
	"timecreated": ListGovernanceRulesSortByTimecreated,
	"displayname": ListGovernanceRulesSortByDisplayname,
}

// GetListGovernanceRulesSortByEnumValues Enumerates the set of values for ListGovernanceRulesSortByEnum
func GetListGovernanceRulesSortByEnumValues() []ListGovernanceRulesSortByEnum {
	values := make([]ListGovernanceRulesSortByEnum, 0)
	for _, v := range mappingListGovernanceRulesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListGovernanceRulesSortByEnumStringValues Enumerates the set of values in String for ListGovernanceRulesSortByEnum
func GetListGovernanceRulesSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListGovernanceRulesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListGovernanceRulesSortByEnum(val string) (ListGovernanceRulesSortByEnum, bool) {
	enum, ok := mappingListGovernanceRulesSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
