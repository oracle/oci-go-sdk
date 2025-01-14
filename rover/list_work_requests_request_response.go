// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package rover

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListWorkRequestsRequest wrapper for the ListWorkRequests operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/rover/ListWorkRequests.go.html to see an example of how to use ListWorkRequestsRequest.
type ListWorkRequestsRequest struct {

	// The OCID of the compartment in which to list resources.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// <b>Filter</b> results by the type of the operation associated with the work request.
	OperationType ListWorkRequestsOperationTypeEnum `mandatory:"false" contributesTo:"query" name:"operationType" omitEmpty:"true"`

	// <b>Filter</b> results by work request status.
	Status ListWorkRequestsStatusEnum `mandatory:"false" contributesTo:"query" name:"status" omitEmpty:"true"`

	// Filter work requests by the resource ocid.
	ResourceId *string `mandatory:"false" contributesTo:"query" name:"resourceId"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The sort order to use, either 'asc' or 'desc'.
	SortOrder ListWorkRequestsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// Specifies the field to sort by. Accepts only one field. By default, when you sort by time fields, the results are shown in descending order. All other fields default to ascending order.
	SortBy ListWorkRequestsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListWorkRequestsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListWorkRequestsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListWorkRequestsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListWorkRequestsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListWorkRequestsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListWorkRequestsOperationTypeEnum(string(request.OperationType)); !ok && request.OperationType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for OperationType: %s. Supported values are: %s.", request.OperationType, strings.Join(GetListWorkRequestsOperationTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListWorkRequestsStatusEnum(string(request.Status)); !ok && request.Status != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", request.Status, strings.Join(GetListWorkRequestsStatusEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListWorkRequestsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListWorkRequestsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListWorkRequestsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListWorkRequestsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListWorkRequestsResponse wrapper for the ListWorkRequests operation
type ListWorkRequestsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of WorkRequestCollection instances
	WorkRequestCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the previous batch of items.
	OpcPrevPage *string `presentIn:"header" name:"opc-prev-page"`
}

func (response ListWorkRequestsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListWorkRequestsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListWorkRequestsOperationTypeEnum Enum with underlying type: string
type ListWorkRequestsOperationTypeEnum string

// Set of constants representing the allowable values for ListWorkRequestsOperationTypeEnum
const (
	ListWorkRequestsOperationTypeAddNodes ListWorkRequestsOperationTypeEnum = "ADD_NODES"
)

var mappingListWorkRequestsOperationTypeEnum = map[string]ListWorkRequestsOperationTypeEnum{
	"ADD_NODES": ListWorkRequestsOperationTypeAddNodes,
}

var mappingListWorkRequestsOperationTypeEnumLowerCase = map[string]ListWorkRequestsOperationTypeEnum{
	"add_nodes": ListWorkRequestsOperationTypeAddNodes,
}

// GetListWorkRequestsOperationTypeEnumValues Enumerates the set of values for ListWorkRequestsOperationTypeEnum
func GetListWorkRequestsOperationTypeEnumValues() []ListWorkRequestsOperationTypeEnum {
	values := make([]ListWorkRequestsOperationTypeEnum, 0)
	for _, v := range mappingListWorkRequestsOperationTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetListWorkRequestsOperationTypeEnumStringValues Enumerates the set of values in String for ListWorkRequestsOperationTypeEnum
func GetListWorkRequestsOperationTypeEnumStringValues() []string {
	return []string{
		"ADD_NODES",
	}
}

// GetMappingListWorkRequestsOperationTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListWorkRequestsOperationTypeEnum(val string) (ListWorkRequestsOperationTypeEnum, bool) {
	enum, ok := mappingListWorkRequestsOperationTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListWorkRequestsStatusEnum Enum with underlying type: string
type ListWorkRequestsStatusEnum string

// Set of constants representing the allowable values for ListWorkRequestsStatusEnum
const (
	ListWorkRequestsStatusAccepted       ListWorkRequestsStatusEnum = "ACCEPTED"
	ListWorkRequestsStatusInProgress     ListWorkRequestsStatusEnum = "IN_PROGRESS"
	ListWorkRequestsStatusWaiting        ListWorkRequestsStatusEnum = "WAITING"
	ListWorkRequestsStatusNeedsAttention ListWorkRequestsStatusEnum = "NEEDS_ATTENTION"
	ListWorkRequestsStatusFailed         ListWorkRequestsStatusEnum = "FAILED"
	ListWorkRequestsStatusSucceeded      ListWorkRequestsStatusEnum = "SUCCEEDED"
	ListWorkRequestsStatusCanceling      ListWorkRequestsStatusEnum = "CANCELING"
	ListWorkRequestsStatusCanceled       ListWorkRequestsStatusEnum = "CANCELED"
)

var mappingListWorkRequestsStatusEnum = map[string]ListWorkRequestsStatusEnum{
	"ACCEPTED":        ListWorkRequestsStatusAccepted,
	"IN_PROGRESS":     ListWorkRequestsStatusInProgress,
	"WAITING":         ListWorkRequestsStatusWaiting,
	"NEEDS_ATTENTION": ListWorkRequestsStatusNeedsAttention,
	"FAILED":          ListWorkRequestsStatusFailed,
	"SUCCEEDED":       ListWorkRequestsStatusSucceeded,
	"CANCELING":       ListWorkRequestsStatusCanceling,
	"CANCELED":        ListWorkRequestsStatusCanceled,
}

var mappingListWorkRequestsStatusEnumLowerCase = map[string]ListWorkRequestsStatusEnum{
	"accepted":        ListWorkRequestsStatusAccepted,
	"in_progress":     ListWorkRequestsStatusInProgress,
	"waiting":         ListWorkRequestsStatusWaiting,
	"needs_attention": ListWorkRequestsStatusNeedsAttention,
	"failed":          ListWorkRequestsStatusFailed,
	"succeeded":       ListWorkRequestsStatusSucceeded,
	"canceling":       ListWorkRequestsStatusCanceling,
	"canceled":        ListWorkRequestsStatusCanceled,
}

// GetListWorkRequestsStatusEnumValues Enumerates the set of values for ListWorkRequestsStatusEnum
func GetListWorkRequestsStatusEnumValues() []ListWorkRequestsStatusEnum {
	values := make([]ListWorkRequestsStatusEnum, 0)
	for _, v := range mappingListWorkRequestsStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetListWorkRequestsStatusEnumStringValues Enumerates the set of values in String for ListWorkRequestsStatusEnum
func GetListWorkRequestsStatusEnumStringValues() []string {
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

// GetMappingListWorkRequestsStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListWorkRequestsStatusEnum(val string) (ListWorkRequestsStatusEnum, bool) {
	enum, ok := mappingListWorkRequestsStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListWorkRequestsSortOrderEnum Enum with underlying type: string
type ListWorkRequestsSortOrderEnum string

// Set of constants representing the allowable values for ListWorkRequestsSortOrderEnum
const (
	ListWorkRequestsSortOrderAsc  ListWorkRequestsSortOrderEnum = "ASC"
	ListWorkRequestsSortOrderDesc ListWorkRequestsSortOrderEnum = "DESC"
)

var mappingListWorkRequestsSortOrderEnum = map[string]ListWorkRequestsSortOrderEnum{
	"ASC":  ListWorkRequestsSortOrderAsc,
	"DESC": ListWorkRequestsSortOrderDesc,
}

var mappingListWorkRequestsSortOrderEnumLowerCase = map[string]ListWorkRequestsSortOrderEnum{
	"asc":  ListWorkRequestsSortOrderAsc,
	"desc": ListWorkRequestsSortOrderDesc,
}

// GetListWorkRequestsSortOrderEnumValues Enumerates the set of values for ListWorkRequestsSortOrderEnum
func GetListWorkRequestsSortOrderEnumValues() []ListWorkRequestsSortOrderEnum {
	values := make([]ListWorkRequestsSortOrderEnum, 0)
	for _, v := range mappingListWorkRequestsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListWorkRequestsSortOrderEnumStringValues Enumerates the set of values in String for ListWorkRequestsSortOrderEnum
func GetListWorkRequestsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListWorkRequestsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListWorkRequestsSortOrderEnum(val string) (ListWorkRequestsSortOrderEnum, bool) {
	enum, ok := mappingListWorkRequestsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListWorkRequestsSortByEnum Enum with underlying type: string
type ListWorkRequestsSortByEnum string

// Set of constants representing the allowable values for ListWorkRequestsSortByEnum
const (
	ListWorkRequestsSortByOperationtype ListWorkRequestsSortByEnum = "operationType"
	ListWorkRequestsSortByStatus        ListWorkRequestsSortByEnum = "status"
	ListWorkRequestsSortByTimeaccepted  ListWorkRequestsSortByEnum = "timeAccepted"
)

var mappingListWorkRequestsSortByEnum = map[string]ListWorkRequestsSortByEnum{
	"operationType": ListWorkRequestsSortByOperationtype,
	"status":        ListWorkRequestsSortByStatus,
	"timeAccepted":  ListWorkRequestsSortByTimeaccepted,
}

var mappingListWorkRequestsSortByEnumLowerCase = map[string]ListWorkRequestsSortByEnum{
	"operationtype": ListWorkRequestsSortByOperationtype,
	"status":        ListWorkRequestsSortByStatus,
	"timeaccepted":  ListWorkRequestsSortByTimeaccepted,
}

// GetListWorkRequestsSortByEnumValues Enumerates the set of values for ListWorkRequestsSortByEnum
func GetListWorkRequestsSortByEnumValues() []ListWorkRequestsSortByEnum {
	values := make([]ListWorkRequestsSortByEnum, 0)
	for _, v := range mappingListWorkRequestsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListWorkRequestsSortByEnumStringValues Enumerates the set of values in String for ListWorkRequestsSortByEnum
func GetListWorkRequestsSortByEnumStringValues() []string {
	return []string{
		"operationType",
		"status",
		"timeAccepted",
	}
}

// GetMappingListWorkRequestsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListWorkRequestsSortByEnum(val string) (ListWorkRequestsSortByEnum, bool) {
	enum, ok := mappingListWorkRequestsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
