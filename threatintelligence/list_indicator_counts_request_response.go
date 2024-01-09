// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package threatintelligence

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListIndicatorCountsRequest wrapper for the ListIndicatorCounts operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/threatintelligence/ListIndicatorCounts.go.html to see an example of how to use ListIndicatorCountsRequest.
type ListIndicatorCountsRequest struct {

	// The OCID of the tenancy (root compartment) that is used to filter results.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// The sort order to use, either 'ASC' or 'DESC'.
	SortOrder ListIndicatorCountsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListIndicatorCountsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListIndicatorCountsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListIndicatorCountsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListIndicatorCountsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListIndicatorCountsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListIndicatorCountsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListIndicatorCountsSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListIndicatorCountsResponse wrapper for the ListIndicatorCounts operation
type ListIndicatorCountsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The IndicatorCountCollection instance
	IndicatorCountCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response ListIndicatorCountsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListIndicatorCountsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListIndicatorCountsSortOrderEnum Enum with underlying type: string
type ListIndicatorCountsSortOrderEnum string

// Set of constants representing the allowable values for ListIndicatorCountsSortOrderEnum
const (
	ListIndicatorCountsSortOrderAsc  ListIndicatorCountsSortOrderEnum = "ASC"
	ListIndicatorCountsSortOrderDesc ListIndicatorCountsSortOrderEnum = "DESC"
)

var mappingListIndicatorCountsSortOrderEnum = map[string]ListIndicatorCountsSortOrderEnum{
	"ASC":  ListIndicatorCountsSortOrderAsc,
	"DESC": ListIndicatorCountsSortOrderDesc,
}

var mappingListIndicatorCountsSortOrderEnumLowerCase = map[string]ListIndicatorCountsSortOrderEnum{
	"asc":  ListIndicatorCountsSortOrderAsc,
	"desc": ListIndicatorCountsSortOrderDesc,
}

// GetListIndicatorCountsSortOrderEnumValues Enumerates the set of values for ListIndicatorCountsSortOrderEnum
func GetListIndicatorCountsSortOrderEnumValues() []ListIndicatorCountsSortOrderEnum {
	values := make([]ListIndicatorCountsSortOrderEnum, 0)
	for _, v := range mappingListIndicatorCountsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListIndicatorCountsSortOrderEnumStringValues Enumerates the set of values in String for ListIndicatorCountsSortOrderEnum
func GetListIndicatorCountsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListIndicatorCountsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListIndicatorCountsSortOrderEnum(val string) (ListIndicatorCountsSortOrderEnum, bool) {
	enum, ok := mappingListIndicatorCountsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
