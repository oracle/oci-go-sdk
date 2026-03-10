// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package limitsincrease

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListLimitsIncreaseItemRequestsRequest wrapper for the ListLimitsIncreaseItemRequests operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/limitsincrease/ListLimitsIncreaseItemRequests.go.html to see an example of how to use ListLimitsIncreaseItemRequestsRequest.
type ListLimitsIncreaseItemRequestsRequest struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the parent compartment.
	// Note: The tenancy is the root compartment.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// Name of service that owns the limit.
	Service *string `mandatory:"false" contributesTo:"query" name:"service"`

	// Lifecycle state of the limit increase request.
	LifecycleState *string `mandatory:"false" contributesTo:"query" name:"lifecycleState"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the limit increase request.
	LimitsIncreaseRequestId *string `mandatory:"false" contributesTo:"query" name:"limitsIncreaseRequestId"`

	// The sort order to use, either 'asc' or 'desc'. By default, it is ascending.
	SortOrder ListLimitsIncreaseItemRequestsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. You can specify only one sort order. The default order for `timeCreated` is descending.
	SortBy ListLimitsIncreaseItemRequestsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Maximum number of items returned in a page.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The next page token.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// Override request id for request tracking purposes.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListLimitsIncreaseItemRequestsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListLimitsIncreaseItemRequestsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListLimitsIncreaseItemRequestsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListLimitsIncreaseItemRequestsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListLimitsIncreaseItemRequestsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListLimitsIncreaseItemRequestsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListLimitsIncreaseItemRequestsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListLimitsIncreaseItemRequestsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListLimitsIncreaseItemRequestsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListLimitsIncreaseItemRequestsResponse wrapper for the ListLimitsIncreaseItemRequests operation
type ListLimitsIncreaseItemRequestsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of LimitsIncreaseItemRequestCollection instances
	LimitsIncreaseItemRequestCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response, additional pages
	// of results remain. For important details about how pagination works, see
	// List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListLimitsIncreaseItemRequestsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListLimitsIncreaseItemRequestsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListLimitsIncreaseItemRequestsSortOrderEnum Enum with underlying type: string
type ListLimitsIncreaseItemRequestsSortOrderEnum string

// Set of constants representing the allowable values for ListLimitsIncreaseItemRequestsSortOrderEnum
const (
	ListLimitsIncreaseItemRequestsSortOrderAsc  ListLimitsIncreaseItemRequestsSortOrderEnum = "ASC"
	ListLimitsIncreaseItemRequestsSortOrderDesc ListLimitsIncreaseItemRequestsSortOrderEnum = "DESC"
)

var mappingListLimitsIncreaseItemRequestsSortOrderEnum = map[string]ListLimitsIncreaseItemRequestsSortOrderEnum{
	"ASC":  ListLimitsIncreaseItemRequestsSortOrderAsc,
	"DESC": ListLimitsIncreaseItemRequestsSortOrderDesc,
}

var mappingListLimitsIncreaseItemRequestsSortOrderEnumLowerCase = map[string]ListLimitsIncreaseItemRequestsSortOrderEnum{
	"asc":  ListLimitsIncreaseItemRequestsSortOrderAsc,
	"desc": ListLimitsIncreaseItemRequestsSortOrderDesc,
}

// GetListLimitsIncreaseItemRequestsSortOrderEnumValues Enumerates the set of values for ListLimitsIncreaseItemRequestsSortOrderEnum
func GetListLimitsIncreaseItemRequestsSortOrderEnumValues() []ListLimitsIncreaseItemRequestsSortOrderEnum {
	values := make([]ListLimitsIncreaseItemRequestsSortOrderEnum, 0)
	for _, v := range mappingListLimitsIncreaseItemRequestsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListLimitsIncreaseItemRequestsSortOrderEnumStringValues Enumerates the set of values in String for ListLimitsIncreaseItemRequestsSortOrderEnum
func GetListLimitsIncreaseItemRequestsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListLimitsIncreaseItemRequestsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListLimitsIncreaseItemRequestsSortOrderEnum(val string) (ListLimitsIncreaseItemRequestsSortOrderEnum, bool) {
	enum, ok := mappingListLimitsIncreaseItemRequestsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListLimitsIncreaseItemRequestsSortByEnum Enum with underlying type: string
type ListLimitsIncreaseItemRequestsSortByEnum string

// Set of constants representing the allowable values for ListLimitsIncreaseItemRequestsSortByEnum
const (
	ListLimitsIncreaseItemRequestsSortByTimecreated ListLimitsIncreaseItemRequestsSortByEnum = "timeCreated"
)

var mappingListLimitsIncreaseItemRequestsSortByEnum = map[string]ListLimitsIncreaseItemRequestsSortByEnum{
	"timeCreated": ListLimitsIncreaseItemRequestsSortByTimecreated,
}

var mappingListLimitsIncreaseItemRequestsSortByEnumLowerCase = map[string]ListLimitsIncreaseItemRequestsSortByEnum{
	"timecreated": ListLimitsIncreaseItemRequestsSortByTimecreated,
}

// GetListLimitsIncreaseItemRequestsSortByEnumValues Enumerates the set of values for ListLimitsIncreaseItemRequestsSortByEnum
func GetListLimitsIncreaseItemRequestsSortByEnumValues() []ListLimitsIncreaseItemRequestsSortByEnum {
	values := make([]ListLimitsIncreaseItemRequestsSortByEnum, 0)
	for _, v := range mappingListLimitsIncreaseItemRequestsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListLimitsIncreaseItemRequestsSortByEnumStringValues Enumerates the set of values in String for ListLimitsIncreaseItemRequestsSortByEnum
func GetListLimitsIncreaseItemRequestsSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
	}
}

// GetMappingListLimitsIncreaseItemRequestsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListLimitsIncreaseItemRequestsSortByEnum(val string) (ListLimitsIncreaseItemRequestsSortByEnum, bool) {
	enum, ok := mappingListLimitsIncreaseItemRequestsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
