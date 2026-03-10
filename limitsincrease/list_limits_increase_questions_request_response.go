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

// ListLimitsIncreaseQuestionsRequest wrapper for the ListLimitsIncreaseQuestions operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/limitsincrease/ListLimitsIncreaseQuestions.go.html to see an example of how to use ListLimitsIncreaseQuestionsRequest.
type ListLimitsIncreaseQuestionsRequest struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the parent compartment.
	// Note: The tenancy is the root compartment.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// The name of the service.
	ServiceName *string `mandatory:"false" contributesTo:"query" name:"serviceName"`

	// The name of the limit.
	LimitName *string `mandatory:"false" contributesTo:"query" name:"limitName"`

	// Override request id for request tracking purposes.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// The sort order to use, either 'asc' or 'desc'. By default, it is ascending.
	SortOrder ListLimitsIncreaseQuestionsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Both fields will sort alphabetically
	SortBy ListLimitsIncreaseQuestionsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Maximum number of items returned in a page.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The next page token.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListLimitsIncreaseQuestionsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListLimitsIncreaseQuestionsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListLimitsIncreaseQuestionsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListLimitsIncreaseQuestionsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListLimitsIncreaseQuestionsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListLimitsIncreaseQuestionsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListLimitsIncreaseQuestionsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListLimitsIncreaseQuestionsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListLimitsIncreaseQuestionsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListLimitsIncreaseQuestionsResponse wrapper for the ListLimitsIncreaseQuestions operation
type ListLimitsIncreaseQuestionsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of LimitsIncreaseQuestionCollection instances
	LimitsIncreaseQuestionCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response, additional pages
	// of results remain. For important details about how pagination works, see
	// List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListLimitsIncreaseQuestionsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListLimitsIncreaseQuestionsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListLimitsIncreaseQuestionsSortOrderEnum Enum with underlying type: string
type ListLimitsIncreaseQuestionsSortOrderEnum string

// Set of constants representing the allowable values for ListLimitsIncreaseQuestionsSortOrderEnum
const (
	ListLimitsIncreaseQuestionsSortOrderAsc  ListLimitsIncreaseQuestionsSortOrderEnum = "ASC"
	ListLimitsIncreaseQuestionsSortOrderDesc ListLimitsIncreaseQuestionsSortOrderEnum = "DESC"
)

var mappingListLimitsIncreaseQuestionsSortOrderEnum = map[string]ListLimitsIncreaseQuestionsSortOrderEnum{
	"ASC":  ListLimitsIncreaseQuestionsSortOrderAsc,
	"DESC": ListLimitsIncreaseQuestionsSortOrderDesc,
}

var mappingListLimitsIncreaseQuestionsSortOrderEnumLowerCase = map[string]ListLimitsIncreaseQuestionsSortOrderEnum{
	"asc":  ListLimitsIncreaseQuestionsSortOrderAsc,
	"desc": ListLimitsIncreaseQuestionsSortOrderDesc,
}

// GetListLimitsIncreaseQuestionsSortOrderEnumValues Enumerates the set of values for ListLimitsIncreaseQuestionsSortOrderEnum
func GetListLimitsIncreaseQuestionsSortOrderEnumValues() []ListLimitsIncreaseQuestionsSortOrderEnum {
	values := make([]ListLimitsIncreaseQuestionsSortOrderEnum, 0)
	for _, v := range mappingListLimitsIncreaseQuestionsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListLimitsIncreaseQuestionsSortOrderEnumStringValues Enumerates the set of values in String for ListLimitsIncreaseQuestionsSortOrderEnum
func GetListLimitsIncreaseQuestionsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListLimitsIncreaseQuestionsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListLimitsIncreaseQuestionsSortOrderEnum(val string) (ListLimitsIncreaseQuestionsSortOrderEnum, bool) {
	enum, ok := mappingListLimitsIncreaseQuestionsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListLimitsIncreaseQuestionsSortByEnum Enum with underlying type: string
type ListLimitsIncreaseQuestionsSortByEnum string

// Set of constants representing the allowable values for ListLimitsIncreaseQuestionsSortByEnum
const (
	ListLimitsIncreaseQuestionsSortByServicename ListLimitsIncreaseQuestionsSortByEnum = "serviceName"
	ListLimitsIncreaseQuestionsSortByLimitname   ListLimitsIncreaseQuestionsSortByEnum = "limitName"
)

var mappingListLimitsIncreaseQuestionsSortByEnum = map[string]ListLimitsIncreaseQuestionsSortByEnum{
	"serviceName": ListLimitsIncreaseQuestionsSortByServicename,
	"limitName":   ListLimitsIncreaseQuestionsSortByLimitname,
}

var mappingListLimitsIncreaseQuestionsSortByEnumLowerCase = map[string]ListLimitsIncreaseQuestionsSortByEnum{
	"servicename": ListLimitsIncreaseQuestionsSortByServicename,
	"limitname":   ListLimitsIncreaseQuestionsSortByLimitname,
}

// GetListLimitsIncreaseQuestionsSortByEnumValues Enumerates the set of values for ListLimitsIncreaseQuestionsSortByEnum
func GetListLimitsIncreaseQuestionsSortByEnumValues() []ListLimitsIncreaseQuestionsSortByEnum {
	values := make([]ListLimitsIncreaseQuestionsSortByEnum, 0)
	for _, v := range mappingListLimitsIncreaseQuestionsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListLimitsIncreaseQuestionsSortByEnumStringValues Enumerates the set of values in String for ListLimitsIncreaseQuestionsSortByEnum
func GetListLimitsIncreaseQuestionsSortByEnumStringValues() []string {
	return []string{
		"serviceName",
		"limitName",
	}
}

// GetMappingListLimitsIncreaseQuestionsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListLimitsIncreaseQuestionsSortByEnum(val string) (ListLimitsIncreaseQuestionsSortByEnum, bool) {
	enum, ok := mappingListLimitsIncreaseQuestionsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
