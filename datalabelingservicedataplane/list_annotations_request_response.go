// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package datalabelingservicedataplane

import (
	"github.com/oracle/oci-go-sdk/v50/common"
	"net/http"
)

// ListAnnotationsRequest wrapper for the ListAnnotations operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datalabelingservicedataplane/ListAnnotations.go.html to see an example of how to use ListAnnotationsRequest.
type ListAnnotationsRequest struct {

	// The ID of the compartment in which to list resources.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// Filter results by the OCID of the dataset.
	DatasetId *string `mandatory:"true" contributesTo:"query" name:"datasetId"`

	// A filter to return only resources their lifecycleState matches the given lifecycleState.
	LifecycleState AnnotationLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// Unique OCID identifier
	Id *string `mandatory:"false" contributesTo:"query" name:"id"`

	// The OCID of the principal who updated the annotation.
	UpdatedBy *string `mandatory:"false" contributesTo:"query" name:"updatedBy"`

	// The OCID of the record annotated
	RecordId *string `mandatory:"false" contributesTo:"query" name:"recordId"`

	// The date and time the resource was created, in the timestamp format defined by RFC3339.
	TimeCreatedGreaterThanOrEqualTo *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timeCreatedGreaterThanOrEqualTo"`

	// The date and time the resource was created, in the timestamp format defined by RFC3339.
	TimeCreatedLessThanOrEqualTo *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timeCreatedLessThanOrEqualTo"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'asc' or 'desc'.
	SortOrder ListAnnotationsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for timeCreated is descending. If no value is specified timeCreated is default.
	SortBy ListAnnotationsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListAnnotationsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListAnnotationsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListAnnotationsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListAnnotationsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ListAnnotationsResponse wrapper for the ListAnnotations operation
type ListAnnotationsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of AnnotationCollection instances
	AnnotationCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListAnnotationsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListAnnotationsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListAnnotationsSortOrderEnum Enum with underlying type: string
type ListAnnotationsSortOrderEnum string

// Set of constants representing the allowable values for ListAnnotationsSortOrderEnum
const (
	ListAnnotationsSortOrderAsc  ListAnnotationsSortOrderEnum = "ASC"
	ListAnnotationsSortOrderDesc ListAnnotationsSortOrderEnum = "DESC"
)

var mappingListAnnotationsSortOrder = map[string]ListAnnotationsSortOrderEnum{
	"ASC":  ListAnnotationsSortOrderAsc,
	"DESC": ListAnnotationsSortOrderDesc,
}

// GetListAnnotationsSortOrderEnumValues Enumerates the set of values for ListAnnotationsSortOrderEnum
func GetListAnnotationsSortOrderEnumValues() []ListAnnotationsSortOrderEnum {
	values := make([]ListAnnotationsSortOrderEnum, 0)
	for _, v := range mappingListAnnotationsSortOrder {
		values = append(values, v)
	}
	return values
}

// ListAnnotationsSortByEnum Enum with underlying type: string
type ListAnnotationsSortByEnum string

// Set of constants representing the allowable values for ListAnnotationsSortByEnum
const (
	ListAnnotationsSortByTimecreated ListAnnotationsSortByEnum = "timeCreated"
	ListAnnotationsSortByLabel       ListAnnotationsSortByEnum = "label"
)

var mappingListAnnotationsSortBy = map[string]ListAnnotationsSortByEnum{
	"timeCreated": ListAnnotationsSortByTimecreated,
	"label":       ListAnnotationsSortByLabel,
}

// GetListAnnotationsSortByEnumValues Enumerates the set of values for ListAnnotationsSortByEnum
func GetListAnnotationsSortByEnumValues() []ListAnnotationsSortByEnum {
	values := make([]ListAnnotationsSortByEnum, 0)
	for _, v := range mappingListAnnotationsSortBy {
		values = append(values, v)
	}
	return values
}
