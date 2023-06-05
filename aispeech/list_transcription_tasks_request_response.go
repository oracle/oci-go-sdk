// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package aispeech

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListTranscriptionTasksRequest wrapper for the ListTranscriptionTasks operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/aispeech/ListTranscriptionTasks.go.html to see an example of how to use ListTranscriptionTasksRequest.
type ListTranscriptionTasksRequest struct {

	// Unique Transcription Job identifier.
	TranscriptionJobId *string `mandatory:"true" contributesTo:"path" name:"transcriptionJobId"`

	// A filter to return only resources whose lifecycleState matches the given lifecycleState.
	LifecycleState TranscriptionTaskLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// A filter to return only resources that match the entire display name given.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// Unique identifier(OCID).
	Id *string `mandatory:"false" contributesTo:"query" name:"id"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// A token representing the position at which to start retrieving results. This must come from the `opc-next-page` header field of a previous response.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'ASC' or 'DESC'.
	SortOrder ListTranscriptionTasksSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for timeCreated is descending. Default order for displayName is ascending.
	SortBy ListTranscriptionTasksSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListTranscriptionTasksRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListTranscriptionTasksRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListTranscriptionTasksRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListTranscriptionTasksRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListTranscriptionTasksRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingTranscriptionTaskLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetTranscriptionTaskLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListTranscriptionTasksSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListTranscriptionTasksSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListTranscriptionTasksSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListTranscriptionTasksSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListTranscriptionTasksResponse wrapper for the ListTranscriptionTasks operation
type ListTranscriptionTasksResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of TranscriptionTaskCollection instances
	TranscriptionTaskCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// Retrieves the previous page of results. When this header appears in the response, previous pages of results exist. See List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcPrevPage *string `presentIn:"header" name:"opc-prev-page"`
}

func (response ListTranscriptionTasksResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListTranscriptionTasksResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListTranscriptionTasksSortOrderEnum Enum with underlying type: string
type ListTranscriptionTasksSortOrderEnum string

// Set of constants representing the allowable values for ListTranscriptionTasksSortOrderEnum
const (
	ListTranscriptionTasksSortOrderAsc  ListTranscriptionTasksSortOrderEnum = "ASC"
	ListTranscriptionTasksSortOrderDesc ListTranscriptionTasksSortOrderEnum = "DESC"
)

var mappingListTranscriptionTasksSortOrderEnum = map[string]ListTranscriptionTasksSortOrderEnum{
	"ASC":  ListTranscriptionTasksSortOrderAsc,
	"DESC": ListTranscriptionTasksSortOrderDesc,
}

var mappingListTranscriptionTasksSortOrderEnumLowerCase = map[string]ListTranscriptionTasksSortOrderEnum{
	"asc":  ListTranscriptionTasksSortOrderAsc,
	"desc": ListTranscriptionTasksSortOrderDesc,
}

// GetListTranscriptionTasksSortOrderEnumValues Enumerates the set of values for ListTranscriptionTasksSortOrderEnum
func GetListTranscriptionTasksSortOrderEnumValues() []ListTranscriptionTasksSortOrderEnum {
	values := make([]ListTranscriptionTasksSortOrderEnum, 0)
	for _, v := range mappingListTranscriptionTasksSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListTranscriptionTasksSortOrderEnumStringValues Enumerates the set of values in String for ListTranscriptionTasksSortOrderEnum
func GetListTranscriptionTasksSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListTranscriptionTasksSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListTranscriptionTasksSortOrderEnum(val string) (ListTranscriptionTasksSortOrderEnum, bool) {
	enum, ok := mappingListTranscriptionTasksSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListTranscriptionTasksSortByEnum Enum with underlying type: string
type ListTranscriptionTasksSortByEnum string

// Set of constants representing the allowable values for ListTranscriptionTasksSortByEnum
const (
	ListTranscriptionTasksSortByTimecreated ListTranscriptionTasksSortByEnum = "timeCreated"
	ListTranscriptionTasksSortByDisplayname ListTranscriptionTasksSortByEnum = "displayName"
)

var mappingListTranscriptionTasksSortByEnum = map[string]ListTranscriptionTasksSortByEnum{
	"timeCreated": ListTranscriptionTasksSortByTimecreated,
	"displayName": ListTranscriptionTasksSortByDisplayname,
}

var mappingListTranscriptionTasksSortByEnumLowerCase = map[string]ListTranscriptionTasksSortByEnum{
	"timecreated": ListTranscriptionTasksSortByTimecreated,
	"displayname": ListTranscriptionTasksSortByDisplayname,
}

// GetListTranscriptionTasksSortByEnumValues Enumerates the set of values for ListTranscriptionTasksSortByEnum
func GetListTranscriptionTasksSortByEnumValues() []ListTranscriptionTasksSortByEnum {
	values := make([]ListTranscriptionTasksSortByEnum, 0)
	for _, v := range mappingListTranscriptionTasksSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListTranscriptionTasksSortByEnumStringValues Enumerates the set of values in String for ListTranscriptionTasksSortByEnum
func GetListTranscriptionTasksSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListTranscriptionTasksSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListTranscriptionTasksSortByEnum(val string) (ListTranscriptionTasksSortByEnum, bool) {
	enum, ok := mappingListTranscriptionTasksSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
