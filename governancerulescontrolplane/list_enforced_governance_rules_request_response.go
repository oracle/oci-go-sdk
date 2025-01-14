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

// ListEnforcedGovernanceRulesRequest wrapper for the ListEnforcedGovernanceRules operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/governancerulescontrolplane/ListEnforcedGovernanceRules.go.html to see an example of how to use ListEnforcedGovernanceRulesRequest.
type ListEnforcedGovernanceRulesRequest struct {

	// The ID of the compartment in which to list resources.
	CompartmentId *string `mandatory:"false" contributesTo:"query" name:"compartmentId"`

	// Unique enforced governance rule identifier.
	EnforcedGovernanceRuleId *string `mandatory:"false" contributesTo:"query" name:"enforcedGovernanceRuleId"`

	// A filter to return only resources that match the type given.
	GovernanceRuleType ListEnforcedGovernanceRulesGovernanceRuleTypeEnum `mandatory:"false" contributesTo:"query" name:"governanceRuleType" omitEmpty:"true"`

	// A filter to return only resources that match the entire name given.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// A token representing the position at which to start retrieving results. This must come from the `opc-next-page` header field of a previous response.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'ASC' or 'DESC'.
	SortOrder ListEnforcedGovernanceRulesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for timeCreated is descending. Default order for displayName is ascending.
	SortBy ListEnforcedGovernanceRulesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListEnforcedGovernanceRulesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListEnforcedGovernanceRulesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListEnforcedGovernanceRulesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListEnforcedGovernanceRulesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListEnforcedGovernanceRulesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListEnforcedGovernanceRulesGovernanceRuleTypeEnum(string(request.GovernanceRuleType)); !ok && request.GovernanceRuleType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for GovernanceRuleType: %s. Supported values are: %s.", request.GovernanceRuleType, strings.Join(GetListEnforcedGovernanceRulesGovernanceRuleTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListEnforcedGovernanceRulesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListEnforcedGovernanceRulesSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListEnforcedGovernanceRulesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListEnforcedGovernanceRulesSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListEnforcedGovernanceRulesResponse wrapper for the ListEnforcedGovernanceRules operation
type ListEnforcedGovernanceRulesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of EnforcedGovernanceRuleCollection instances
	EnforcedGovernanceRuleCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListEnforcedGovernanceRulesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListEnforcedGovernanceRulesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListEnforcedGovernanceRulesGovernanceRuleTypeEnum Enum with underlying type: string
type ListEnforcedGovernanceRulesGovernanceRuleTypeEnum string

// Set of constants representing the allowable values for ListEnforcedGovernanceRulesGovernanceRuleTypeEnum
const (
	ListEnforcedGovernanceRulesGovernanceRuleTypeQuota          ListEnforcedGovernanceRulesGovernanceRuleTypeEnum = "QUOTA"
	ListEnforcedGovernanceRulesGovernanceRuleTypeTag            ListEnforcedGovernanceRulesGovernanceRuleTypeEnum = "TAG"
	ListEnforcedGovernanceRulesGovernanceRuleTypeAllowedRegions ListEnforcedGovernanceRulesGovernanceRuleTypeEnum = "ALLOWED_REGIONS"
)

var mappingListEnforcedGovernanceRulesGovernanceRuleTypeEnum = map[string]ListEnforcedGovernanceRulesGovernanceRuleTypeEnum{
	"QUOTA":           ListEnforcedGovernanceRulesGovernanceRuleTypeQuota,
	"TAG":             ListEnforcedGovernanceRulesGovernanceRuleTypeTag,
	"ALLOWED_REGIONS": ListEnforcedGovernanceRulesGovernanceRuleTypeAllowedRegions,
}

var mappingListEnforcedGovernanceRulesGovernanceRuleTypeEnumLowerCase = map[string]ListEnforcedGovernanceRulesGovernanceRuleTypeEnum{
	"quota":           ListEnforcedGovernanceRulesGovernanceRuleTypeQuota,
	"tag":             ListEnforcedGovernanceRulesGovernanceRuleTypeTag,
	"allowed_regions": ListEnforcedGovernanceRulesGovernanceRuleTypeAllowedRegions,
}

// GetListEnforcedGovernanceRulesGovernanceRuleTypeEnumValues Enumerates the set of values for ListEnforcedGovernanceRulesGovernanceRuleTypeEnum
func GetListEnforcedGovernanceRulesGovernanceRuleTypeEnumValues() []ListEnforcedGovernanceRulesGovernanceRuleTypeEnum {
	values := make([]ListEnforcedGovernanceRulesGovernanceRuleTypeEnum, 0)
	for _, v := range mappingListEnforcedGovernanceRulesGovernanceRuleTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetListEnforcedGovernanceRulesGovernanceRuleTypeEnumStringValues Enumerates the set of values in String for ListEnforcedGovernanceRulesGovernanceRuleTypeEnum
func GetListEnforcedGovernanceRulesGovernanceRuleTypeEnumStringValues() []string {
	return []string{
		"QUOTA",
		"TAG",
		"ALLOWED_REGIONS",
	}
}

// GetMappingListEnforcedGovernanceRulesGovernanceRuleTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListEnforcedGovernanceRulesGovernanceRuleTypeEnum(val string) (ListEnforcedGovernanceRulesGovernanceRuleTypeEnum, bool) {
	enum, ok := mappingListEnforcedGovernanceRulesGovernanceRuleTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListEnforcedGovernanceRulesSortOrderEnum Enum with underlying type: string
type ListEnforcedGovernanceRulesSortOrderEnum string

// Set of constants representing the allowable values for ListEnforcedGovernanceRulesSortOrderEnum
const (
	ListEnforcedGovernanceRulesSortOrderAsc  ListEnforcedGovernanceRulesSortOrderEnum = "ASC"
	ListEnforcedGovernanceRulesSortOrderDesc ListEnforcedGovernanceRulesSortOrderEnum = "DESC"
)

var mappingListEnforcedGovernanceRulesSortOrderEnum = map[string]ListEnforcedGovernanceRulesSortOrderEnum{
	"ASC":  ListEnforcedGovernanceRulesSortOrderAsc,
	"DESC": ListEnforcedGovernanceRulesSortOrderDesc,
}

var mappingListEnforcedGovernanceRulesSortOrderEnumLowerCase = map[string]ListEnforcedGovernanceRulesSortOrderEnum{
	"asc":  ListEnforcedGovernanceRulesSortOrderAsc,
	"desc": ListEnforcedGovernanceRulesSortOrderDesc,
}

// GetListEnforcedGovernanceRulesSortOrderEnumValues Enumerates the set of values for ListEnforcedGovernanceRulesSortOrderEnum
func GetListEnforcedGovernanceRulesSortOrderEnumValues() []ListEnforcedGovernanceRulesSortOrderEnum {
	values := make([]ListEnforcedGovernanceRulesSortOrderEnum, 0)
	for _, v := range mappingListEnforcedGovernanceRulesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListEnforcedGovernanceRulesSortOrderEnumStringValues Enumerates the set of values in String for ListEnforcedGovernanceRulesSortOrderEnum
func GetListEnforcedGovernanceRulesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListEnforcedGovernanceRulesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListEnforcedGovernanceRulesSortOrderEnum(val string) (ListEnforcedGovernanceRulesSortOrderEnum, bool) {
	enum, ok := mappingListEnforcedGovernanceRulesSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListEnforcedGovernanceRulesSortByEnum Enum with underlying type: string
type ListEnforcedGovernanceRulesSortByEnum string

// Set of constants representing the allowable values for ListEnforcedGovernanceRulesSortByEnum
const (
	ListEnforcedGovernanceRulesSortByTimecreated ListEnforcedGovernanceRulesSortByEnum = "timeCreated"
	ListEnforcedGovernanceRulesSortByDisplayname ListEnforcedGovernanceRulesSortByEnum = "displayName"
)

var mappingListEnforcedGovernanceRulesSortByEnum = map[string]ListEnforcedGovernanceRulesSortByEnum{
	"timeCreated": ListEnforcedGovernanceRulesSortByTimecreated,
	"displayName": ListEnforcedGovernanceRulesSortByDisplayname,
}

var mappingListEnforcedGovernanceRulesSortByEnumLowerCase = map[string]ListEnforcedGovernanceRulesSortByEnum{
	"timecreated": ListEnforcedGovernanceRulesSortByTimecreated,
	"displayname": ListEnforcedGovernanceRulesSortByDisplayname,
}

// GetListEnforcedGovernanceRulesSortByEnumValues Enumerates the set of values for ListEnforcedGovernanceRulesSortByEnum
func GetListEnforcedGovernanceRulesSortByEnumValues() []ListEnforcedGovernanceRulesSortByEnum {
	values := make([]ListEnforcedGovernanceRulesSortByEnum, 0)
	for _, v := range mappingListEnforcedGovernanceRulesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListEnforcedGovernanceRulesSortByEnumStringValues Enumerates the set of values in String for ListEnforcedGovernanceRulesSortByEnum
func GetListEnforcedGovernanceRulesSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListEnforcedGovernanceRulesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListEnforcedGovernanceRulesSortByEnum(val string) (ListEnforcedGovernanceRulesSortByEnum, bool) {
	enum, ok := mappingListEnforcedGovernanceRulesSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
