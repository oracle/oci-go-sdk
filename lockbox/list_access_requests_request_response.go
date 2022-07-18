// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package lockbox

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListAccessRequestsRequest wrapper for the ListAccessRequests operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/lockbox/ListAccessRequests.go.html to see an example of how to use ListAccessRequestsRequest.
type ListAccessRequestsRequest struct {

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// The unique identifier (OCID) of the associated lockbox.
	LockboxId *string `mandatory:"false" contributesTo:"query" name:"lockboxId"`

	// A generic Id query param used to filter lockbox, access request and approval template by Id.
	Id *string `mandatory:"false" contributesTo:"query" name:"id"`

	// A filter to return only resources that match the entire display name given.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// A filter to return only resources their lifecycleState matches the given lifecycleState.
	LifecycleState AccessRequestLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// The name of the lockbox partner.
	LockboxPartner ListAccessRequestsLockboxPartnerEnum `mandatory:"false" contributesTo:"query" name:"lockboxPartner" omitEmpty:"true"`

	// The unique identifier (OCID) of the requestor in which to list resources.
	RequestorId *string `mandatory:"false" contributesTo:"query" name:"requestorId"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// A token representing the position at which to start retrieving results. This must come from the `opc-next-page` header field of a previous response.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'ASC' or 'DESC'.
	SortOrder ListAccessRequestsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for timeCreated is descending. Default order for displayName is ascending.
	SortBy ListAccessRequestsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListAccessRequestsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListAccessRequestsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListAccessRequestsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListAccessRequestsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListAccessRequestsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingAccessRequestLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetAccessRequestLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListAccessRequestsLockboxPartnerEnum(string(request.LockboxPartner)); !ok && request.LockboxPartner != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LockboxPartner: %s. Supported values are: %s.", request.LockboxPartner, strings.Join(GetListAccessRequestsLockboxPartnerEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListAccessRequestsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListAccessRequestsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListAccessRequestsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListAccessRequestsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListAccessRequestsResponse wrapper for the ListAccessRequests operation
type ListAccessRequestsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of AccessRequestCollection instances
	AccessRequestCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListAccessRequestsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListAccessRequestsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListAccessRequestsLockboxPartnerEnum Enum with underlying type: string
type ListAccessRequestsLockboxPartnerEnum string

// Set of constants representing the allowable values for ListAccessRequestsLockboxPartnerEnum
const (
	ListAccessRequestsLockboxPartnerFaaas  ListAccessRequestsLockboxPartnerEnum = "FAAAS"
	ListAccessRequestsLockboxPartnerCanary ListAccessRequestsLockboxPartnerEnum = "CANARY"
)

var mappingListAccessRequestsLockboxPartnerEnum = map[string]ListAccessRequestsLockboxPartnerEnum{
	"FAAAS":  ListAccessRequestsLockboxPartnerFaaas,
	"CANARY": ListAccessRequestsLockboxPartnerCanary,
}

var mappingListAccessRequestsLockboxPartnerEnumLowerCase = map[string]ListAccessRequestsLockboxPartnerEnum{
	"faaas":  ListAccessRequestsLockboxPartnerFaaas,
	"canary": ListAccessRequestsLockboxPartnerCanary,
}

// GetListAccessRequestsLockboxPartnerEnumValues Enumerates the set of values for ListAccessRequestsLockboxPartnerEnum
func GetListAccessRequestsLockboxPartnerEnumValues() []ListAccessRequestsLockboxPartnerEnum {
	values := make([]ListAccessRequestsLockboxPartnerEnum, 0)
	for _, v := range mappingListAccessRequestsLockboxPartnerEnum {
		values = append(values, v)
	}
	return values
}

// GetListAccessRequestsLockboxPartnerEnumStringValues Enumerates the set of values in String for ListAccessRequestsLockboxPartnerEnum
func GetListAccessRequestsLockboxPartnerEnumStringValues() []string {
	return []string{
		"FAAAS",
		"CANARY",
	}
}

// GetMappingListAccessRequestsLockboxPartnerEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListAccessRequestsLockboxPartnerEnum(val string) (ListAccessRequestsLockboxPartnerEnum, bool) {
	enum, ok := mappingListAccessRequestsLockboxPartnerEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListAccessRequestsSortOrderEnum Enum with underlying type: string
type ListAccessRequestsSortOrderEnum string

// Set of constants representing the allowable values for ListAccessRequestsSortOrderEnum
const (
	ListAccessRequestsSortOrderAsc  ListAccessRequestsSortOrderEnum = "ASC"
	ListAccessRequestsSortOrderDesc ListAccessRequestsSortOrderEnum = "DESC"
)

var mappingListAccessRequestsSortOrderEnum = map[string]ListAccessRequestsSortOrderEnum{
	"ASC":  ListAccessRequestsSortOrderAsc,
	"DESC": ListAccessRequestsSortOrderDesc,
}

var mappingListAccessRequestsSortOrderEnumLowerCase = map[string]ListAccessRequestsSortOrderEnum{
	"asc":  ListAccessRequestsSortOrderAsc,
	"desc": ListAccessRequestsSortOrderDesc,
}

// GetListAccessRequestsSortOrderEnumValues Enumerates the set of values for ListAccessRequestsSortOrderEnum
func GetListAccessRequestsSortOrderEnumValues() []ListAccessRequestsSortOrderEnum {
	values := make([]ListAccessRequestsSortOrderEnum, 0)
	for _, v := range mappingListAccessRequestsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListAccessRequestsSortOrderEnumStringValues Enumerates the set of values in String for ListAccessRequestsSortOrderEnum
func GetListAccessRequestsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListAccessRequestsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListAccessRequestsSortOrderEnum(val string) (ListAccessRequestsSortOrderEnum, bool) {
	enum, ok := mappingListAccessRequestsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListAccessRequestsSortByEnum Enum with underlying type: string
type ListAccessRequestsSortByEnum string

// Set of constants representing the allowable values for ListAccessRequestsSortByEnum
const (
	ListAccessRequestsSortByTimecreated ListAccessRequestsSortByEnum = "timeCreated"
	ListAccessRequestsSortByDisplayname ListAccessRequestsSortByEnum = "displayName"
	ListAccessRequestsSortById          ListAccessRequestsSortByEnum = "id"
)

var mappingListAccessRequestsSortByEnum = map[string]ListAccessRequestsSortByEnum{
	"timeCreated": ListAccessRequestsSortByTimecreated,
	"displayName": ListAccessRequestsSortByDisplayname,
	"id":          ListAccessRequestsSortById,
}

var mappingListAccessRequestsSortByEnumLowerCase = map[string]ListAccessRequestsSortByEnum{
	"timecreated": ListAccessRequestsSortByTimecreated,
	"displayname": ListAccessRequestsSortByDisplayname,
	"id":          ListAccessRequestsSortById,
}

// GetListAccessRequestsSortByEnumValues Enumerates the set of values for ListAccessRequestsSortByEnum
func GetListAccessRequestsSortByEnumValues() []ListAccessRequestsSortByEnum {
	values := make([]ListAccessRequestsSortByEnum, 0)
	for _, v := range mappingListAccessRequestsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListAccessRequestsSortByEnumStringValues Enumerates the set of values in String for ListAccessRequestsSortByEnum
func GetListAccessRequestsSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
		"id",
	}
}

// GetMappingListAccessRequestsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListAccessRequestsSortByEnum(val string) (ListAccessRequestsSortByEnum, bool) {
	enum, ok := mappingListAccessRequestsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
