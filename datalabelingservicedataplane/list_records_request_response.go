// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package datalabelingservicedataplane

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v57/common"
	"net/http"
	"strings"
)

// ListRecordsRequest wrapper for the ListRecords operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datalabelingservicedataplane/ListRecords.go.html to see an example of how to use ListRecordsRequest.
type ListRecordsRequest struct {

	// The ID of the compartment in which to list resources.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// Filter the results by the OCID of the dataset.
	DatasetId *string `mandatory:"true" contributesTo:"query" name:"datasetId"`

	// A filter to return only resources whose lifecycleState matches the given lifecycleState.
	LifecycleState RecordLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// The name of the record.
	Name *string `mandatory:"false" contributesTo:"query" name:"name"`

	// The unique OCID identifier.
	Id *string `mandatory:"false" contributesTo:"query" name:"id"`

	// Whether the record has been labeled and has associated annotations.
	IsLabeled *bool `mandatory:"false" contributesTo:"query" name:"isLabeled"`

	// Lets the user filter records based on the related annotations.
	AnnotationLabelsContains []string `contributesTo:"query" name:"annotationLabelsContains" collectionFormat:"multi"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'asc' or 'desc'.
	SortOrder ListRecordsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. The default order for timeCreated is descending. The default order for name is ascending. If no value is specified, timeCreated is used by default.
	SortBy ListRecordsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListRecordsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListRecordsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListRecordsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListRecordsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListRecordsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := mappingRecordLifecycleStateEnum[string(request.LifecycleState)]; !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetRecordLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := mappingListRecordsSortOrderEnum[string(request.SortOrder)]; !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListRecordsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := mappingListRecordsSortByEnum[string(request.SortBy)]; !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListRecordsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListRecordsResponse wrapper for the ListRecords operation
type ListRecordsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of RecordCollection instances
	RecordCollection `presentIn:"body"`

	// A unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For the pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListRecordsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListRecordsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListRecordsSortOrderEnum Enum with underlying type: string
type ListRecordsSortOrderEnum string

// Set of constants representing the allowable values for ListRecordsSortOrderEnum
const (
	ListRecordsSortOrderAsc  ListRecordsSortOrderEnum = "ASC"
	ListRecordsSortOrderDesc ListRecordsSortOrderEnum = "DESC"
)

var mappingListRecordsSortOrderEnum = map[string]ListRecordsSortOrderEnum{
	"ASC":  ListRecordsSortOrderAsc,
	"DESC": ListRecordsSortOrderDesc,
}

// GetListRecordsSortOrderEnumValues Enumerates the set of values for ListRecordsSortOrderEnum
func GetListRecordsSortOrderEnumValues() []ListRecordsSortOrderEnum {
	values := make([]ListRecordsSortOrderEnum, 0)
	for _, v := range mappingListRecordsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListRecordsSortOrderEnumStringValues Enumerates the set of values in String for ListRecordsSortOrderEnum
func GetListRecordsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// ListRecordsSortByEnum Enum with underlying type: string
type ListRecordsSortByEnum string

// Set of constants representing the allowable values for ListRecordsSortByEnum
const (
	ListRecordsSortByTimecreated ListRecordsSortByEnum = "timeCreated"
	ListRecordsSortByName        ListRecordsSortByEnum = "name"
)

var mappingListRecordsSortByEnum = map[string]ListRecordsSortByEnum{
	"timeCreated": ListRecordsSortByTimecreated,
	"name":        ListRecordsSortByName,
}

// GetListRecordsSortByEnumValues Enumerates the set of values for ListRecordsSortByEnum
func GetListRecordsSortByEnumValues() []ListRecordsSortByEnum {
	values := make([]ListRecordsSortByEnum, 0)
	for _, v := range mappingListRecordsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListRecordsSortByEnumStringValues Enumerates the set of values in String for ListRecordsSortByEnum
func GetListRecordsSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"name",
	}
}
