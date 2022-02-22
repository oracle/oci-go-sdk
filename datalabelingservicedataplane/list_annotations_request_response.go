// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package datalabelingservicedataplane

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v59/common"
	"net/http"
	"strings"
)

// ListAnnotationsRequest wrapper for the ListAnnotations operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datalabelingservicedataplane/ListAnnotations.go.html to see an example of how to use ListAnnotationsRequest.
type ListAnnotationsRequest struct {

	// The ID of the compartment in which to list resources.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// Filter the results by the OCID of the dataset.
	DatasetId *string `mandatory:"true" contributesTo:"query" name:"datasetId"`

	// A filter to return only resources whose lifecycleState matches the given lifecycleState.
	LifecycleState AnnotationLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// The unique OCID identifier.
	Id *string `mandatory:"false" contributesTo:"query" name:"id"`

	// The OCID of the principal which updated the annotation.
	UpdatedBy *string `mandatory:"false" contributesTo:"query" name:"updatedBy"`

	// The OCID of the record annotated.
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

	// The field to sort by. Only one sort order may be provided. The default order for timeCreated is descending. If no value is specified timeCreated is used by default.
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

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
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

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListAnnotationsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingAnnotationLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetAnnotationLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListAnnotationsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListAnnotationsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListAnnotationsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListAnnotationsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListAnnotationsResponse wrapper for the ListAnnotations operation
type ListAnnotationsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of AnnotationCollection instances
	AnnotationCollection `presentIn:"body"`

	// A unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For the pagination of a list of items. When paging through a list, if this header appears in the response,
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

var mappingListAnnotationsSortOrderEnum = map[string]ListAnnotationsSortOrderEnum{
	"ASC":  ListAnnotationsSortOrderAsc,
	"DESC": ListAnnotationsSortOrderDesc,
}

var mappingListAnnotationsSortOrderEnumLowerCase = map[string]ListAnnotationsSortOrderEnum{
	"asc":  ListAnnotationsSortOrderAsc,
	"desc": ListAnnotationsSortOrderDesc,
}

// GetListAnnotationsSortOrderEnumValues Enumerates the set of values for ListAnnotationsSortOrderEnum
func GetListAnnotationsSortOrderEnumValues() []ListAnnotationsSortOrderEnum {
	values := make([]ListAnnotationsSortOrderEnum, 0)
	for _, v := range mappingListAnnotationsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListAnnotationsSortOrderEnumStringValues Enumerates the set of values in String for ListAnnotationsSortOrderEnum
func GetListAnnotationsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListAnnotationsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListAnnotationsSortOrderEnum(val string) (ListAnnotationsSortOrderEnum, bool) {
	enum, ok := mappingListAnnotationsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListAnnotationsSortByEnum Enum with underlying type: string
type ListAnnotationsSortByEnum string

// Set of constants representing the allowable values for ListAnnotationsSortByEnum
const (
	ListAnnotationsSortByTimecreated ListAnnotationsSortByEnum = "timeCreated"
	ListAnnotationsSortByLabel       ListAnnotationsSortByEnum = "label"
)

var mappingListAnnotationsSortByEnum = map[string]ListAnnotationsSortByEnum{
	"timeCreated": ListAnnotationsSortByTimecreated,
	"label":       ListAnnotationsSortByLabel,
}

var mappingListAnnotationsSortByEnumLowerCase = map[string]ListAnnotationsSortByEnum{
	"timecreated": ListAnnotationsSortByTimecreated,
	"label":       ListAnnotationsSortByLabel,
}

// GetListAnnotationsSortByEnumValues Enumerates the set of values for ListAnnotationsSortByEnum
func GetListAnnotationsSortByEnumValues() []ListAnnotationsSortByEnum {
	values := make([]ListAnnotationsSortByEnum, 0)
	for _, v := range mappingListAnnotationsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListAnnotationsSortByEnumStringValues Enumerates the set of values in String for ListAnnotationsSortByEnum
func GetListAnnotationsSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"label",
	}
}

// GetMappingListAnnotationsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListAnnotationsSortByEnum(val string) (ListAnnotationsSortByEnum, bool) {
	enum, ok := mappingListAnnotationsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
