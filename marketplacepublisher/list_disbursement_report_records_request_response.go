// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package marketplacepublisher

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListDisbursementReportRecordsRequest wrapper for the ListDisbursementReportRecords operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/marketplacepublisher/ListDisbursementReportRecords.go.html to see an example of how to use ListDisbursementReportRecordsRequest.
type ListDisbursementReportRecordsRequest struct {

	// The ID of the compartment in which to list resources.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// The date range of the report
	DateRange ListDisbursementReportRecordsDateRangeEnum `mandatory:"true" contributesTo:"query" name:"dateRange" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// A token representing the position at which to start retrieving results. This must come from the `opc-next-page` header field of a previous response.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'ASC' or 'DESC'.
	SortOrder ListDisbursementReportRecordsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for timeCreated is descending. Default order for displayName is ascending.
	SortBy ListDisbursementReportRecordsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// A filter to return only resources that match the entire name given.
	Name *string `mandatory:"false" contributesTo:"query" name:"name"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListDisbursementReportRecordsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListDisbursementReportRecordsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListDisbursementReportRecordsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListDisbursementReportRecordsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListDisbursementReportRecordsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListDisbursementReportRecordsDateRangeEnum(string(request.DateRange)); !ok && request.DateRange != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DateRange: %s. Supported values are: %s.", request.DateRange, strings.Join(GetListDisbursementReportRecordsDateRangeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListDisbursementReportRecordsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListDisbursementReportRecordsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListDisbursementReportRecordsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListDisbursementReportRecordsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListDisbursementReportRecordsResponse wrapper for the ListDisbursementReportRecords operation
type ListDisbursementReportRecordsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of DisbursementReportRecordCollection instances
	DisbursementReportRecordCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListDisbursementReportRecordsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListDisbursementReportRecordsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListDisbursementReportRecordsDateRangeEnum Enum with underlying type: string
type ListDisbursementReportRecordsDateRangeEnum string

// Set of constants representing the allowable values for ListDisbursementReportRecordsDateRangeEnum
const (
	ListDisbursementReportRecordsDateRangeMonth       ListDisbursementReportRecordsDateRangeEnum = "LAST_MONTH"
	ListDisbursementReportRecordsDateRangeThreeMonths ListDisbursementReportRecordsDateRangeEnum = "LAST_THREE_MONTHS"
)

var mappingListDisbursementReportRecordsDateRangeEnum = map[string]ListDisbursementReportRecordsDateRangeEnum{
	"LAST_MONTH":        ListDisbursementReportRecordsDateRangeMonth,
	"LAST_THREE_MONTHS": ListDisbursementReportRecordsDateRangeThreeMonths,
}

var mappingListDisbursementReportRecordsDateRangeEnumLowerCase = map[string]ListDisbursementReportRecordsDateRangeEnum{
	"last_month":        ListDisbursementReportRecordsDateRangeMonth,
	"last_three_months": ListDisbursementReportRecordsDateRangeThreeMonths,
}

// GetListDisbursementReportRecordsDateRangeEnumValues Enumerates the set of values for ListDisbursementReportRecordsDateRangeEnum
func GetListDisbursementReportRecordsDateRangeEnumValues() []ListDisbursementReportRecordsDateRangeEnum {
	values := make([]ListDisbursementReportRecordsDateRangeEnum, 0)
	for _, v := range mappingListDisbursementReportRecordsDateRangeEnum {
		values = append(values, v)
	}
	return values
}

// GetListDisbursementReportRecordsDateRangeEnumStringValues Enumerates the set of values in String for ListDisbursementReportRecordsDateRangeEnum
func GetListDisbursementReportRecordsDateRangeEnumStringValues() []string {
	return []string{
		"LAST_MONTH",
		"LAST_THREE_MONTHS",
	}
}

// GetMappingListDisbursementReportRecordsDateRangeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListDisbursementReportRecordsDateRangeEnum(val string) (ListDisbursementReportRecordsDateRangeEnum, bool) {
	enum, ok := mappingListDisbursementReportRecordsDateRangeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListDisbursementReportRecordsSortOrderEnum Enum with underlying type: string
type ListDisbursementReportRecordsSortOrderEnum string

// Set of constants representing the allowable values for ListDisbursementReportRecordsSortOrderEnum
const (
	ListDisbursementReportRecordsSortOrderAsc  ListDisbursementReportRecordsSortOrderEnum = "ASC"
	ListDisbursementReportRecordsSortOrderDesc ListDisbursementReportRecordsSortOrderEnum = "DESC"
)

var mappingListDisbursementReportRecordsSortOrderEnum = map[string]ListDisbursementReportRecordsSortOrderEnum{
	"ASC":  ListDisbursementReportRecordsSortOrderAsc,
	"DESC": ListDisbursementReportRecordsSortOrderDesc,
}

var mappingListDisbursementReportRecordsSortOrderEnumLowerCase = map[string]ListDisbursementReportRecordsSortOrderEnum{
	"asc":  ListDisbursementReportRecordsSortOrderAsc,
	"desc": ListDisbursementReportRecordsSortOrderDesc,
}

// GetListDisbursementReportRecordsSortOrderEnumValues Enumerates the set of values for ListDisbursementReportRecordsSortOrderEnum
func GetListDisbursementReportRecordsSortOrderEnumValues() []ListDisbursementReportRecordsSortOrderEnum {
	values := make([]ListDisbursementReportRecordsSortOrderEnum, 0)
	for _, v := range mappingListDisbursementReportRecordsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListDisbursementReportRecordsSortOrderEnumStringValues Enumerates the set of values in String for ListDisbursementReportRecordsSortOrderEnum
func GetListDisbursementReportRecordsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListDisbursementReportRecordsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListDisbursementReportRecordsSortOrderEnum(val string) (ListDisbursementReportRecordsSortOrderEnum, bool) {
	enum, ok := mappingListDisbursementReportRecordsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListDisbursementReportRecordsSortByEnum Enum with underlying type: string
type ListDisbursementReportRecordsSortByEnum string

// Set of constants representing the allowable values for ListDisbursementReportRecordsSortByEnum
const (
	ListDisbursementReportRecordsSortByTimecreated ListDisbursementReportRecordsSortByEnum = "timeCreated"
	ListDisbursementReportRecordsSortByDisplayname ListDisbursementReportRecordsSortByEnum = "displayName"
)

var mappingListDisbursementReportRecordsSortByEnum = map[string]ListDisbursementReportRecordsSortByEnum{
	"timeCreated": ListDisbursementReportRecordsSortByTimecreated,
	"displayName": ListDisbursementReportRecordsSortByDisplayname,
}

var mappingListDisbursementReportRecordsSortByEnumLowerCase = map[string]ListDisbursementReportRecordsSortByEnum{
	"timecreated": ListDisbursementReportRecordsSortByTimecreated,
	"displayname": ListDisbursementReportRecordsSortByDisplayname,
}

// GetListDisbursementReportRecordsSortByEnumValues Enumerates the set of values for ListDisbursementReportRecordsSortByEnum
func GetListDisbursementReportRecordsSortByEnumValues() []ListDisbursementReportRecordsSortByEnum {
	values := make([]ListDisbursementReportRecordsSortByEnum, 0)
	for _, v := range mappingListDisbursementReportRecordsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListDisbursementReportRecordsSortByEnumStringValues Enumerates the set of values in String for ListDisbursementReportRecordsSortByEnum
func GetListDisbursementReportRecordsSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListDisbursementReportRecordsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListDisbursementReportRecordsSortByEnum(val string) (ListDisbursementReportRecordsSortByEnum, bool) {
	enum, ok := mappingListDisbursementReportRecordsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
