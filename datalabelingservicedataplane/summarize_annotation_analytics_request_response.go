// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package datalabelingservicedataplane

import (
	"github.com/oracle/oci-go-sdk/v50/common"
	"net/http"
)

// SummarizeAnnotationAnalyticsRequest wrapper for the SummarizeAnnotationAnalytics operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datalabelingservicedataplane/SummarizeAnnotationAnalytics.go.html to see an example of how to use SummarizeAnnotationAnalyticsRequest.
type SummarizeAnnotationAnalyticsRequest struct {

	// The ID of the compartment in which to list resources.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// Filter results by the OCID of the dataset.
	DatasetId *string `mandatory:"true" contributesTo:"query" name:"datasetId"`

	// A filter to return only resources their lifecycleState matches the given lifecycleState.
	LifecycleState AnnotationLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// This field is used to summarize annotations with specified label.
	Label *string `mandatory:"false" contributesTo:"query" name:"label"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'asc' or 'desc'.
	SortOrder SummarizeAnnotationAnalyticsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order is descending. If no value is specified updatedBy is default.
	SortBy SummarizeAnnotationAnalyticsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The field to group by. If no value is specified updatedBy is default.
	AnnotationGroupBy SummarizeAnnotationAnalyticsAnnotationGroupByEnum `mandatory:"false" contributesTo:"query" name:"annotationGroupBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request SummarizeAnnotationAnalyticsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request SummarizeAnnotationAnalyticsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request SummarizeAnnotationAnalyticsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request SummarizeAnnotationAnalyticsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// SummarizeAnnotationAnalyticsResponse wrapper for the SummarizeAnnotationAnalytics operation
type SummarizeAnnotationAnalyticsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of AnnotationAnalyticsAggregationCollection instances
	AnnotationAnalyticsAggregationCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response SummarizeAnnotationAnalyticsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response SummarizeAnnotationAnalyticsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// SummarizeAnnotationAnalyticsSortOrderEnum Enum with underlying type: string
type SummarizeAnnotationAnalyticsSortOrderEnum string

// Set of constants representing the allowable values for SummarizeAnnotationAnalyticsSortOrderEnum
const (
	SummarizeAnnotationAnalyticsSortOrderAsc  SummarizeAnnotationAnalyticsSortOrderEnum = "ASC"
	SummarizeAnnotationAnalyticsSortOrderDesc SummarizeAnnotationAnalyticsSortOrderEnum = "DESC"
)

var mappingSummarizeAnnotationAnalyticsSortOrder = map[string]SummarizeAnnotationAnalyticsSortOrderEnum{
	"ASC":  SummarizeAnnotationAnalyticsSortOrderAsc,
	"DESC": SummarizeAnnotationAnalyticsSortOrderDesc,
}

// GetSummarizeAnnotationAnalyticsSortOrderEnumValues Enumerates the set of values for SummarizeAnnotationAnalyticsSortOrderEnum
func GetSummarizeAnnotationAnalyticsSortOrderEnumValues() []SummarizeAnnotationAnalyticsSortOrderEnum {
	values := make([]SummarizeAnnotationAnalyticsSortOrderEnum, 0)
	for _, v := range mappingSummarizeAnnotationAnalyticsSortOrder {
		values = append(values, v)
	}
	return values
}

// SummarizeAnnotationAnalyticsSortByEnum Enum with underlying type: string
type SummarizeAnnotationAnalyticsSortByEnum string

// Set of constants representing the allowable values for SummarizeAnnotationAnalyticsSortByEnum
const (
	SummarizeAnnotationAnalyticsSortByCount     SummarizeAnnotationAnalyticsSortByEnum = "count"
	SummarizeAnnotationAnalyticsSortByLabel     SummarizeAnnotationAnalyticsSortByEnum = "label"
	SummarizeAnnotationAnalyticsSortByUpdatedby SummarizeAnnotationAnalyticsSortByEnum = "updatedBy"
)

var mappingSummarizeAnnotationAnalyticsSortBy = map[string]SummarizeAnnotationAnalyticsSortByEnum{
	"count":     SummarizeAnnotationAnalyticsSortByCount,
	"label":     SummarizeAnnotationAnalyticsSortByLabel,
	"updatedBy": SummarizeAnnotationAnalyticsSortByUpdatedby,
}

// GetSummarizeAnnotationAnalyticsSortByEnumValues Enumerates the set of values for SummarizeAnnotationAnalyticsSortByEnum
func GetSummarizeAnnotationAnalyticsSortByEnumValues() []SummarizeAnnotationAnalyticsSortByEnum {
	values := make([]SummarizeAnnotationAnalyticsSortByEnum, 0)
	for _, v := range mappingSummarizeAnnotationAnalyticsSortBy {
		values = append(values, v)
	}
	return values
}

// SummarizeAnnotationAnalyticsAnnotationGroupByEnum Enum with underlying type: string
type SummarizeAnnotationAnalyticsAnnotationGroupByEnum string

// Set of constants representing the allowable values for SummarizeAnnotationAnalyticsAnnotationGroupByEnum
const (
	SummarizeAnnotationAnalyticsAnnotationGroupByUpdatedby SummarizeAnnotationAnalyticsAnnotationGroupByEnum = "updatedBy"
	SummarizeAnnotationAnalyticsAnnotationGroupByLabel     SummarizeAnnotationAnalyticsAnnotationGroupByEnum = "label"
)

var mappingSummarizeAnnotationAnalyticsAnnotationGroupBy = map[string]SummarizeAnnotationAnalyticsAnnotationGroupByEnum{
	"updatedBy": SummarizeAnnotationAnalyticsAnnotationGroupByUpdatedby,
	"label":     SummarizeAnnotationAnalyticsAnnotationGroupByLabel,
}

// GetSummarizeAnnotationAnalyticsAnnotationGroupByEnumValues Enumerates the set of values for SummarizeAnnotationAnalyticsAnnotationGroupByEnum
func GetSummarizeAnnotationAnalyticsAnnotationGroupByEnumValues() []SummarizeAnnotationAnalyticsAnnotationGroupByEnum {
	values := make([]SummarizeAnnotationAnalyticsAnnotationGroupByEnum, 0)
	for _, v := range mappingSummarizeAnnotationAnalyticsAnnotationGroupBy {
		values = append(values, v)
	}
	return values
}
