// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package lockbox

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListLockboxesRequest wrapper for the ListLockboxes operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/lockbox/ListLockboxes.go.html to see an example of how to use ListLockboxesRequest.
type ListLockboxesRequest struct {

	// The ID of the compartment in which to list resources.
	CompartmentId *string `mandatory:"false" contributesTo:"query" name:"compartmentId"`

	// A filter to return only resources their lifecycleState matches the given lifecycleState.
	LifecycleState LockboxLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// A filter to return only resources that match the entire display name given.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// unique Lockbox identifier
	Id *string `mandatory:"false" contributesTo:"query" name:"id"`

	// The ID of the resource associated with the lockbox.
	ResourceId *string `mandatory:"false" contributesTo:"query" name:"resourceId"`

	// The name of the lockbox partner.
	LockboxPartner ListLockboxesLockboxPartnerEnum `mandatory:"false" contributesTo:"query" name:"lockboxPartner" omitEmpty:"true"`

	// The ID of the partner.
	PartnerId *string `mandatory:"false" contributesTo:"query" name:"partnerId"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// A token representing the position at which to start retrieving results. This must come from the `opc-next-page` header field of a previous response.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'ASC' or 'DESC'.
	SortOrder ListLockboxesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for timeCreated is descending. Default order for displayName is ascending.
	SortBy ListLockboxesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListLockboxesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListLockboxesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListLockboxesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListLockboxesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListLockboxesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingLockboxLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetLockboxLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListLockboxesLockboxPartnerEnum(string(request.LockboxPartner)); !ok && request.LockboxPartner != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LockboxPartner: %s. Supported values are: %s.", request.LockboxPartner, strings.Join(GetListLockboxesLockboxPartnerEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListLockboxesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListLockboxesSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListLockboxesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListLockboxesSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListLockboxesResponse wrapper for the ListLockboxes operation
type ListLockboxesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of LockboxCollection instances
	LockboxCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListLockboxesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListLockboxesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListLockboxesLockboxPartnerEnum Enum with underlying type: string
type ListLockboxesLockboxPartnerEnum string

// Set of constants representing the allowable values for ListLockboxesLockboxPartnerEnum
const (
	ListLockboxesLockboxPartnerFaaas  ListLockboxesLockboxPartnerEnum = "FAAAS"
	ListLockboxesLockboxPartnerCanary ListLockboxesLockboxPartnerEnum = "CANARY"
)

var mappingListLockboxesLockboxPartnerEnum = map[string]ListLockboxesLockboxPartnerEnum{
	"FAAAS":  ListLockboxesLockboxPartnerFaaas,
	"CANARY": ListLockboxesLockboxPartnerCanary,
}

var mappingListLockboxesLockboxPartnerEnumLowerCase = map[string]ListLockboxesLockboxPartnerEnum{
	"faaas":  ListLockboxesLockboxPartnerFaaas,
	"canary": ListLockboxesLockboxPartnerCanary,
}

// GetListLockboxesLockboxPartnerEnumValues Enumerates the set of values for ListLockboxesLockboxPartnerEnum
func GetListLockboxesLockboxPartnerEnumValues() []ListLockboxesLockboxPartnerEnum {
	values := make([]ListLockboxesLockboxPartnerEnum, 0)
	for _, v := range mappingListLockboxesLockboxPartnerEnum {
		values = append(values, v)
	}
	return values
}

// GetListLockboxesLockboxPartnerEnumStringValues Enumerates the set of values in String for ListLockboxesLockboxPartnerEnum
func GetListLockboxesLockboxPartnerEnumStringValues() []string {
	return []string{
		"FAAAS",
		"CANARY",
	}
}

// GetMappingListLockboxesLockboxPartnerEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListLockboxesLockboxPartnerEnum(val string) (ListLockboxesLockboxPartnerEnum, bool) {
	enum, ok := mappingListLockboxesLockboxPartnerEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListLockboxesSortOrderEnum Enum with underlying type: string
type ListLockboxesSortOrderEnum string

// Set of constants representing the allowable values for ListLockboxesSortOrderEnum
const (
	ListLockboxesSortOrderAsc  ListLockboxesSortOrderEnum = "ASC"
	ListLockboxesSortOrderDesc ListLockboxesSortOrderEnum = "DESC"
)

var mappingListLockboxesSortOrderEnum = map[string]ListLockboxesSortOrderEnum{
	"ASC":  ListLockboxesSortOrderAsc,
	"DESC": ListLockboxesSortOrderDesc,
}

var mappingListLockboxesSortOrderEnumLowerCase = map[string]ListLockboxesSortOrderEnum{
	"asc":  ListLockboxesSortOrderAsc,
	"desc": ListLockboxesSortOrderDesc,
}

// GetListLockboxesSortOrderEnumValues Enumerates the set of values for ListLockboxesSortOrderEnum
func GetListLockboxesSortOrderEnumValues() []ListLockboxesSortOrderEnum {
	values := make([]ListLockboxesSortOrderEnum, 0)
	for _, v := range mappingListLockboxesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListLockboxesSortOrderEnumStringValues Enumerates the set of values in String for ListLockboxesSortOrderEnum
func GetListLockboxesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListLockboxesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListLockboxesSortOrderEnum(val string) (ListLockboxesSortOrderEnum, bool) {
	enum, ok := mappingListLockboxesSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListLockboxesSortByEnum Enum with underlying type: string
type ListLockboxesSortByEnum string

// Set of constants representing the allowable values for ListLockboxesSortByEnum
const (
	ListLockboxesSortByTimecreated ListLockboxesSortByEnum = "timeCreated"
	ListLockboxesSortByDisplayname ListLockboxesSortByEnum = "displayName"
	ListLockboxesSortById          ListLockboxesSortByEnum = "id"
)

var mappingListLockboxesSortByEnum = map[string]ListLockboxesSortByEnum{
	"timeCreated": ListLockboxesSortByTimecreated,
	"displayName": ListLockboxesSortByDisplayname,
	"id":          ListLockboxesSortById,
}

var mappingListLockboxesSortByEnumLowerCase = map[string]ListLockboxesSortByEnum{
	"timecreated": ListLockboxesSortByTimecreated,
	"displayname": ListLockboxesSortByDisplayname,
	"id":          ListLockboxesSortById,
}

// GetListLockboxesSortByEnumValues Enumerates the set of values for ListLockboxesSortByEnum
func GetListLockboxesSortByEnumValues() []ListLockboxesSortByEnum {
	values := make([]ListLockboxesSortByEnum, 0)
	for _, v := range mappingListLockboxesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListLockboxesSortByEnumStringValues Enumerates the set of values in String for ListLockboxesSortByEnum
func GetListLockboxesSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
		"id",
	}
}

// GetMappingListLockboxesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListLockboxesSortByEnum(val string) (ListLockboxesSortByEnum, bool) {
	enum, ok := mappingListLockboxesSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
