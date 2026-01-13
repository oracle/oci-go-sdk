// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package threatintelligence

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListThreatTypesRequest wrapper for the ListThreatTypes operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/threatintelligence/ListThreatTypes.go.html to see an example of how to use ListThreatTypesRequest.
type ListThreatTypesRequest struct {

	// The OCID of the tenancy (root compartment) that is used to filter results.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// A token representing the position at which to start retrieving results. This must come from the `opc-next-page` header field of a previous response.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'ASC' or 'DESC'.
	SortOrder ListThreatTypesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListThreatTypesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListThreatTypesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListThreatTypesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListThreatTypesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListThreatTypesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListThreatTypesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListThreatTypesSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListThreatTypesResponse wrapper for the ListThreatTypes operation
type ListThreatTypesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of ThreatTypesCollection instances
	ThreatTypesCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListThreatTypesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListThreatTypesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListThreatTypesSortOrderEnum Enum with underlying type: string
type ListThreatTypesSortOrderEnum string

// Set of constants representing the allowable values for ListThreatTypesSortOrderEnum
const (
	ListThreatTypesSortOrderAsc  ListThreatTypesSortOrderEnum = "ASC"
	ListThreatTypesSortOrderDesc ListThreatTypesSortOrderEnum = "DESC"
)

var mappingListThreatTypesSortOrderEnum = map[string]ListThreatTypesSortOrderEnum{
	"ASC":  ListThreatTypesSortOrderAsc,
	"DESC": ListThreatTypesSortOrderDesc,
}

var mappingListThreatTypesSortOrderEnumLowerCase = map[string]ListThreatTypesSortOrderEnum{
	"asc":  ListThreatTypesSortOrderAsc,
	"desc": ListThreatTypesSortOrderDesc,
}

// GetListThreatTypesSortOrderEnumValues Enumerates the set of values for ListThreatTypesSortOrderEnum
func GetListThreatTypesSortOrderEnumValues() []ListThreatTypesSortOrderEnum {
	values := make([]ListThreatTypesSortOrderEnum, 0)
	for _, v := range mappingListThreatTypesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListThreatTypesSortOrderEnumStringValues Enumerates the set of values in String for ListThreatTypesSortOrderEnum
func GetListThreatTypesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListThreatTypesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListThreatTypesSortOrderEnum(val string) (ListThreatTypesSortOrderEnum, bool) {
	enum, ok := mappingListThreatTypesSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
