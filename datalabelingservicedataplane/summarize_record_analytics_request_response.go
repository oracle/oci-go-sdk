// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package datalabelingservicedataplane

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// SummarizeRecordAnalyticsRequest wrapper for the SummarizeRecordAnalytics operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datalabelingservicedataplane/SummarizeRecordAnalytics.go.html to see an example of how to use SummarizeRecordAnalyticsRequest.
type SummarizeRecordAnalyticsRequest struct {

	// The ID of the compartment in which to list resources.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// Filter the results by the OCID of the dataset.
	DatasetId *string `mandatory:"true" contributesTo:"query" name:"datasetId"`

	// A filter to return only resources whose lifecycleState matches the given lifecycleState.
	LifecycleState RecordLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'asc' or 'desc'.
	SortOrder SummarizeRecordAnalyticsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to group by. If no value is specified isLabeled is used by default.
	RecordGroupBy SummarizeRecordAnalyticsRecordGroupByEnum `mandatory:"false" contributesTo:"query" name:"recordGroupBy" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. The default order is descending. If no value is specified, count is used by default.
	SortBy SummarizeRecordAnalyticsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request SummarizeRecordAnalyticsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request SummarizeRecordAnalyticsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request SummarizeRecordAnalyticsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request SummarizeRecordAnalyticsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request SummarizeRecordAnalyticsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingRecordLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetRecordLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingSummarizeRecordAnalyticsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetSummarizeRecordAnalyticsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingSummarizeRecordAnalyticsRecordGroupByEnum(string(request.RecordGroupBy)); !ok && request.RecordGroupBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for RecordGroupBy: %s. Supported values are: %s.", request.RecordGroupBy, strings.Join(GetSummarizeRecordAnalyticsRecordGroupByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingSummarizeRecordAnalyticsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetSummarizeRecordAnalyticsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// SummarizeRecordAnalyticsResponse wrapper for the SummarizeRecordAnalytics operation
type SummarizeRecordAnalyticsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of RecordAnalyticsAggregationCollection instances
	RecordAnalyticsAggregationCollection `presentIn:"body"`

	// A unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For the pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response SummarizeRecordAnalyticsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response SummarizeRecordAnalyticsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// SummarizeRecordAnalyticsSortOrderEnum Enum with underlying type: string
type SummarizeRecordAnalyticsSortOrderEnum string

// Set of constants representing the allowable values for SummarizeRecordAnalyticsSortOrderEnum
const (
	SummarizeRecordAnalyticsSortOrderAsc  SummarizeRecordAnalyticsSortOrderEnum = "ASC"
	SummarizeRecordAnalyticsSortOrderDesc SummarizeRecordAnalyticsSortOrderEnum = "DESC"
)

var mappingSummarizeRecordAnalyticsSortOrderEnum = map[string]SummarizeRecordAnalyticsSortOrderEnum{
	"ASC":  SummarizeRecordAnalyticsSortOrderAsc,
	"DESC": SummarizeRecordAnalyticsSortOrderDesc,
}

var mappingSummarizeRecordAnalyticsSortOrderEnumLowerCase = map[string]SummarizeRecordAnalyticsSortOrderEnum{
	"asc":  SummarizeRecordAnalyticsSortOrderAsc,
	"desc": SummarizeRecordAnalyticsSortOrderDesc,
}

// GetSummarizeRecordAnalyticsSortOrderEnumValues Enumerates the set of values for SummarizeRecordAnalyticsSortOrderEnum
func GetSummarizeRecordAnalyticsSortOrderEnumValues() []SummarizeRecordAnalyticsSortOrderEnum {
	values := make([]SummarizeRecordAnalyticsSortOrderEnum, 0)
	for _, v := range mappingSummarizeRecordAnalyticsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetSummarizeRecordAnalyticsSortOrderEnumStringValues Enumerates the set of values in String for SummarizeRecordAnalyticsSortOrderEnum
func GetSummarizeRecordAnalyticsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingSummarizeRecordAnalyticsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSummarizeRecordAnalyticsSortOrderEnum(val string) (SummarizeRecordAnalyticsSortOrderEnum, bool) {
	enum, ok := mappingSummarizeRecordAnalyticsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// SummarizeRecordAnalyticsRecordGroupByEnum Enum with underlying type: string
type SummarizeRecordAnalyticsRecordGroupByEnum string

// Set of constants representing the allowable values for SummarizeRecordAnalyticsRecordGroupByEnum
const (
	SummarizeRecordAnalyticsRecordGroupByIslabeled               SummarizeRecordAnalyticsRecordGroupByEnum = "isLabeled"
	SummarizeRecordAnalyticsRecordGroupByAnnotationlabelcontains SummarizeRecordAnalyticsRecordGroupByEnum = "annotationLabelContains"
)

var mappingSummarizeRecordAnalyticsRecordGroupByEnum = map[string]SummarizeRecordAnalyticsRecordGroupByEnum{
	"isLabeled":               SummarizeRecordAnalyticsRecordGroupByIslabeled,
	"annotationLabelContains": SummarizeRecordAnalyticsRecordGroupByAnnotationlabelcontains,
}

var mappingSummarizeRecordAnalyticsRecordGroupByEnumLowerCase = map[string]SummarizeRecordAnalyticsRecordGroupByEnum{
	"islabeled":               SummarizeRecordAnalyticsRecordGroupByIslabeled,
	"annotationlabelcontains": SummarizeRecordAnalyticsRecordGroupByAnnotationlabelcontains,
}

// GetSummarizeRecordAnalyticsRecordGroupByEnumValues Enumerates the set of values for SummarizeRecordAnalyticsRecordGroupByEnum
func GetSummarizeRecordAnalyticsRecordGroupByEnumValues() []SummarizeRecordAnalyticsRecordGroupByEnum {
	values := make([]SummarizeRecordAnalyticsRecordGroupByEnum, 0)
	for _, v := range mappingSummarizeRecordAnalyticsRecordGroupByEnum {
		values = append(values, v)
	}
	return values
}

// GetSummarizeRecordAnalyticsRecordGroupByEnumStringValues Enumerates the set of values in String for SummarizeRecordAnalyticsRecordGroupByEnum
func GetSummarizeRecordAnalyticsRecordGroupByEnumStringValues() []string {
	return []string{
		"isLabeled",
		"annotationLabelContains",
	}
}

// GetMappingSummarizeRecordAnalyticsRecordGroupByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSummarizeRecordAnalyticsRecordGroupByEnum(val string) (SummarizeRecordAnalyticsRecordGroupByEnum, bool) {
	enum, ok := mappingSummarizeRecordAnalyticsRecordGroupByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// SummarizeRecordAnalyticsSortByEnum Enum with underlying type: string
type SummarizeRecordAnalyticsSortByEnum string

// Set of constants representing the allowable values for SummarizeRecordAnalyticsSortByEnum
const (
	SummarizeRecordAnalyticsSortByCount     SummarizeRecordAnalyticsSortByEnum = "count"
	SummarizeRecordAnalyticsSortByIslabeled SummarizeRecordAnalyticsSortByEnum = "isLabeled"
)

var mappingSummarizeRecordAnalyticsSortByEnum = map[string]SummarizeRecordAnalyticsSortByEnum{
	"count":     SummarizeRecordAnalyticsSortByCount,
	"isLabeled": SummarizeRecordAnalyticsSortByIslabeled,
}

var mappingSummarizeRecordAnalyticsSortByEnumLowerCase = map[string]SummarizeRecordAnalyticsSortByEnum{
	"count":     SummarizeRecordAnalyticsSortByCount,
	"islabeled": SummarizeRecordAnalyticsSortByIslabeled,
}

// GetSummarizeRecordAnalyticsSortByEnumValues Enumerates the set of values for SummarizeRecordAnalyticsSortByEnum
func GetSummarizeRecordAnalyticsSortByEnumValues() []SummarizeRecordAnalyticsSortByEnum {
	values := make([]SummarizeRecordAnalyticsSortByEnum, 0)
	for _, v := range mappingSummarizeRecordAnalyticsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetSummarizeRecordAnalyticsSortByEnumStringValues Enumerates the set of values in String for SummarizeRecordAnalyticsSortByEnum
func GetSummarizeRecordAnalyticsSortByEnumStringValues() []string {
	return []string{
		"count",
		"isLabeled",
	}
}

// GetMappingSummarizeRecordAnalyticsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSummarizeRecordAnalyticsSortByEnum(val string) (SummarizeRecordAnalyticsSortByEnum, bool) {
	enum, ok := mappingSummarizeRecordAnalyticsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
