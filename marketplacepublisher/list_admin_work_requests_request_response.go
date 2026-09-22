// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package marketplacepublisher

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListAdminWorkRequestsRequest wrapper for the ListAdminWorkRequests operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/marketplacepublisher/ListAdminWorkRequests.go.html to see an example of how to use ListAdminWorkRequestsRequest.
type ListAdminWorkRequestsRequest struct {

	// The ID of the compartment in which to list resources.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// The ID of the asynchronous work request.
	WorkRequestId *string `mandatory:"false" contributesTo:"query" name:"workRequestId"`

	// A filter to return only resources their lifecycleState matches the given OperationStatus.
	Status ListAdminWorkRequestsStatusEnum `mandatory:"false" contributesTo:"query" name:"status" omitEmpty:"true"`

	// The ID of the resource affected by the work request.
	ResourceId *string `mandatory:"false" contributesTo:"query" name:"resourceId"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// A token representing the position at which to start retrieving results. This must come from the `opc-next-page` header field of a previous response.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The sort order to use, either 'ASC' or 'DESC'.
	SortOrder ListAdminWorkRequestsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for timeAccepted is descending.
	SortBy ListAdminWorkRequestsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListAdminWorkRequestsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListAdminWorkRequestsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListAdminWorkRequestsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListAdminWorkRequestsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListAdminWorkRequestsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListAdminWorkRequestsStatusEnum(string(request.Status)); !ok && request.Status != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", request.Status, strings.Join(GetListAdminWorkRequestsStatusEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListAdminWorkRequestsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListAdminWorkRequestsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListAdminWorkRequestsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListAdminWorkRequestsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListAdminWorkRequestsResponse wrapper for the ListAdminWorkRequests operation
type ListAdminWorkRequestsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of AdminWorkRequestSummaryCollection instances
	AdminWorkRequestSummaryCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListAdminWorkRequestsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListAdminWorkRequestsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListAdminWorkRequestsStatusEnum Enum with underlying type: string
type ListAdminWorkRequestsStatusEnum string

// Set of constants representing the allowable values for ListAdminWorkRequestsStatusEnum
const (
	ListAdminWorkRequestsStatusAccepted       ListAdminWorkRequestsStatusEnum = "ACCEPTED"
	ListAdminWorkRequestsStatusInProgress     ListAdminWorkRequestsStatusEnum = "IN_PROGRESS"
	ListAdminWorkRequestsStatusWaiting        ListAdminWorkRequestsStatusEnum = "WAITING"
	ListAdminWorkRequestsStatusNeedsAttention ListAdminWorkRequestsStatusEnum = "NEEDS_ATTENTION"
	ListAdminWorkRequestsStatusFailed         ListAdminWorkRequestsStatusEnum = "FAILED"
	ListAdminWorkRequestsStatusSucceeded      ListAdminWorkRequestsStatusEnum = "SUCCEEDED"
	ListAdminWorkRequestsStatusCanceling      ListAdminWorkRequestsStatusEnum = "CANCELING"
	ListAdminWorkRequestsStatusCanceled       ListAdminWorkRequestsStatusEnum = "CANCELED"
)

var mappingListAdminWorkRequestsStatusEnum = map[string]ListAdminWorkRequestsStatusEnum{
	"ACCEPTED":        ListAdminWorkRequestsStatusAccepted,
	"IN_PROGRESS":     ListAdminWorkRequestsStatusInProgress,
	"WAITING":         ListAdminWorkRequestsStatusWaiting,
	"NEEDS_ATTENTION": ListAdminWorkRequestsStatusNeedsAttention,
	"FAILED":          ListAdminWorkRequestsStatusFailed,
	"SUCCEEDED":       ListAdminWorkRequestsStatusSucceeded,
	"CANCELING":       ListAdminWorkRequestsStatusCanceling,
	"CANCELED":        ListAdminWorkRequestsStatusCanceled,
}

var mappingListAdminWorkRequestsStatusEnumLowerCase = map[string]ListAdminWorkRequestsStatusEnum{
	"accepted":        ListAdminWorkRequestsStatusAccepted,
	"in_progress":     ListAdminWorkRequestsStatusInProgress,
	"waiting":         ListAdminWorkRequestsStatusWaiting,
	"needs_attention": ListAdminWorkRequestsStatusNeedsAttention,
	"failed":          ListAdminWorkRequestsStatusFailed,
	"succeeded":       ListAdminWorkRequestsStatusSucceeded,
	"canceling":       ListAdminWorkRequestsStatusCanceling,
	"canceled":        ListAdminWorkRequestsStatusCanceled,
}

// GetListAdminWorkRequestsStatusEnumValues Enumerates the set of values for ListAdminWorkRequestsStatusEnum
func GetListAdminWorkRequestsStatusEnumValues() []ListAdminWorkRequestsStatusEnum {
	values := make([]ListAdminWorkRequestsStatusEnum, 0)
	for _, v := range mappingListAdminWorkRequestsStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetListAdminWorkRequestsStatusEnumStringValues Enumerates the set of values in String for ListAdminWorkRequestsStatusEnum
func GetListAdminWorkRequestsStatusEnumStringValues() []string {
	return []string{
		"ACCEPTED",
		"IN_PROGRESS",
		"WAITING",
		"NEEDS_ATTENTION",
		"FAILED",
		"SUCCEEDED",
		"CANCELING",
		"CANCELED",
	}
}

// GetMappingListAdminWorkRequestsStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListAdminWorkRequestsStatusEnum(val string) (ListAdminWorkRequestsStatusEnum, bool) {
	enum, ok := mappingListAdminWorkRequestsStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListAdminWorkRequestsSortOrderEnum Enum with underlying type: string
type ListAdminWorkRequestsSortOrderEnum string

// Set of constants representing the allowable values for ListAdminWorkRequestsSortOrderEnum
const (
	ListAdminWorkRequestsSortOrderAsc  ListAdminWorkRequestsSortOrderEnum = "ASC"
	ListAdminWorkRequestsSortOrderDesc ListAdminWorkRequestsSortOrderEnum = "DESC"
)

var mappingListAdminWorkRequestsSortOrderEnum = map[string]ListAdminWorkRequestsSortOrderEnum{
	"ASC":  ListAdminWorkRequestsSortOrderAsc,
	"DESC": ListAdminWorkRequestsSortOrderDesc,
}

var mappingListAdminWorkRequestsSortOrderEnumLowerCase = map[string]ListAdminWorkRequestsSortOrderEnum{
	"asc":  ListAdminWorkRequestsSortOrderAsc,
	"desc": ListAdminWorkRequestsSortOrderDesc,
}

// GetListAdminWorkRequestsSortOrderEnumValues Enumerates the set of values for ListAdminWorkRequestsSortOrderEnum
func GetListAdminWorkRequestsSortOrderEnumValues() []ListAdminWorkRequestsSortOrderEnum {
	values := make([]ListAdminWorkRequestsSortOrderEnum, 0)
	for _, v := range mappingListAdminWorkRequestsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListAdminWorkRequestsSortOrderEnumStringValues Enumerates the set of values in String for ListAdminWorkRequestsSortOrderEnum
func GetListAdminWorkRequestsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListAdminWorkRequestsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListAdminWorkRequestsSortOrderEnum(val string) (ListAdminWorkRequestsSortOrderEnum, bool) {
	enum, ok := mappingListAdminWorkRequestsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListAdminWorkRequestsSortByEnum Enum with underlying type: string
type ListAdminWorkRequestsSortByEnum string

// Set of constants representing the allowable values for ListAdminWorkRequestsSortByEnum
const (
	ListAdminWorkRequestsSortByTimeaccepted ListAdminWorkRequestsSortByEnum = "timeAccepted"
)

var mappingListAdminWorkRequestsSortByEnum = map[string]ListAdminWorkRequestsSortByEnum{
	"timeAccepted": ListAdminWorkRequestsSortByTimeaccepted,
}

var mappingListAdminWorkRequestsSortByEnumLowerCase = map[string]ListAdminWorkRequestsSortByEnum{
	"timeaccepted": ListAdminWorkRequestsSortByTimeaccepted,
}

// GetListAdminWorkRequestsSortByEnumValues Enumerates the set of values for ListAdminWorkRequestsSortByEnum
func GetListAdminWorkRequestsSortByEnumValues() []ListAdminWorkRequestsSortByEnum {
	values := make([]ListAdminWorkRequestsSortByEnum, 0)
	for _, v := range mappingListAdminWorkRequestsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListAdminWorkRequestsSortByEnumStringValues Enumerates the set of values in String for ListAdminWorkRequestsSortByEnum
func GetListAdminWorkRequestsSortByEnumStringValues() []string {
	return []string{
		"timeAccepted",
	}
}

// GetMappingListAdminWorkRequestsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListAdminWorkRequestsSortByEnum(val string) (ListAdminWorkRequestsSortByEnum, bool) {
	enum, ok := mappingListAdminWorkRequestsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
