// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package datalabelingservicedataplane

import (
	"github.com/oracle/oci-go-sdk/v49/common"
	"net/http"
)

// SummarizeRecordAnalyticsRequest wrapper for the SummarizeRecordAnalytics operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datalabelingservicedataplane/SummarizeRecordAnalytics.go.html to see an example of how to use SummarizeRecordAnalyticsRequest.
type SummarizeRecordAnalyticsRequest struct {

	// The ID of the compartment in which to list resources.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// Filter results by the OCID of the dataset.
	DatasetId *string `mandatory:"true" contributesTo:"query" name:"datasetId"`

	// A filter to return only resources their lifecycleState matches the given lifecycleState.
	LifecycleState RecordLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'asc' or 'desc'.
	SortOrder SummarizeRecordAnalyticsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to group by. If no value is specified isLabeled is default.
	RecordGroupBy SummarizeRecordAnalyticsRecordGroupByEnum `mandatory:"false" contributesTo:"query" name:"recordGroupBy" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order is descending. If no value is specified count is default.
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

// SummarizeRecordAnalyticsResponse wrapper for the SummarizeRecordAnalytics operation
type SummarizeRecordAnalyticsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of RecordAnalyticsAggregationCollection instances
	RecordAnalyticsAggregationCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
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

var mappingSummarizeRecordAnalyticsSortOrder = map[string]SummarizeRecordAnalyticsSortOrderEnum{
	"ASC":  SummarizeRecordAnalyticsSortOrderAsc,
	"DESC": SummarizeRecordAnalyticsSortOrderDesc,
}

// GetSummarizeRecordAnalyticsSortOrderEnumValues Enumerates the set of values for SummarizeRecordAnalyticsSortOrderEnum
func GetSummarizeRecordAnalyticsSortOrderEnumValues() []SummarizeRecordAnalyticsSortOrderEnum {
	values := make([]SummarizeRecordAnalyticsSortOrderEnum, 0)
	for _, v := range mappingSummarizeRecordAnalyticsSortOrder {
		values = append(values, v)
	}
	return values
}

// SummarizeRecordAnalyticsRecordGroupByEnum Enum with underlying type: string
type SummarizeRecordAnalyticsRecordGroupByEnum string

// Set of constants representing the allowable values for SummarizeRecordAnalyticsRecordGroupByEnum
const (
	SummarizeRecordAnalyticsRecordGroupByIslabeled               SummarizeRecordAnalyticsRecordGroupByEnum = "isLabeled"
	SummarizeRecordAnalyticsRecordGroupByAnnotationlabelcontains SummarizeRecordAnalyticsRecordGroupByEnum = "annotationLabelContains"
)

var mappingSummarizeRecordAnalyticsRecordGroupBy = map[string]SummarizeRecordAnalyticsRecordGroupByEnum{
	"isLabeled":               SummarizeRecordAnalyticsRecordGroupByIslabeled,
	"annotationLabelContains": SummarizeRecordAnalyticsRecordGroupByAnnotationlabelcontains,
}

// GetSummarizeRecordAnalyticsRecordGroupByEnumValues Enumerates the set of values for SummarizeRecordAnalyticsRecordGroupByEnum
func GetSummarizeRecordAnalyticsRecordGroupByEnumValues() []SummarizeRecordAnalyticsRecordGroupByEnum {
	values := make([]SummarizeRecordAnalyticsRecordGroupByEnum, 0)
	for _, v := range mappingSummarizeRecordAnalyticsRecordGroupBy {
		values = append(values, v)
	}
	return values
}

// SummarizeRecordAnalyticsSortByEnum Enum with underlying type: string
type SummarizeRecordAnalyticsSortByEnum string

// Set of constants representing the allowable values for SummarizeRecordAnalyticsSortByEnum
const (
	SummarizeRecordAnalyticsSortByCount     SummarizeRecordAnalyticsSortByEnum = "count"
	SummarizeRecordAnalyticsSortByIslabeled SummarizeRecordAnalyticsSortByEnum = "isLabeled"
)

var mappingSummarizeRecordAnalyticsSortBy = map[string]SummarizeRecordAnalyticsSortByEnum{
	"count":     SummarizeRecordAnalyticsSortByCount,
	"isLabeled": SummarizeRecordAnalyticsSortByIslabeled,
}

// GetSummarizeRecordAnalyticsSortByEnumValues Enumerates the set of values for SummarizeRecordAnalyticsSortByEnum
func GetSummarizeRecordAnalyticsSortByEnumValues() []SummarizeRecordAnalyticsSortByEnum {
	values := make([]SummarizeRecordAnalyticsSortByEnum, 0)
	for _, v := range mappingSummarizeRecordAnalyticsSortBy {
		values = append(values, v)
	}
	return values
}
