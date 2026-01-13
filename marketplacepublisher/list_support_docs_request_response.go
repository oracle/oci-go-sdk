// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package marketplacepublisher

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListSupportDocsRequest wrapper for the ListSupportDocs operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/marketplacepublisher/ListSupportDocs.go.html to see an example of how to use ListSupportDocsRequest.
type ListSupportDocsRequest struct {

	// The ID of the compartment in which to list resources.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// A token representing the position at which to start retrieving results. This must come from the `opc-next-page` header field of a previous response.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'ASC' or 'DESC'.
	SortOrder ListSupportDocsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for timeCreated is descending. Default order for displayName is ascending.
	SortBy ListSupportDocsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// SupportDoc Group
	SupportDocGroup *string `mandatory:"false" contributesTo:"query" name:"supportDocGroup"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// A filter to return only resources that match the entire name given.
	Name *string `mandatory:"false" contributesTo:"query" name:"name"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListSupportDocsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListSupportDocsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListSupportDocsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListSupportDocsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListSupportDocsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListSupportDocsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListSupportDocsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListSupportDocsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListSupportDocsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListSupportDocsResponse wrapper for the ListSupportDocs operation
type ListSupportDocsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of SupportDocCollection instances
	SupportDocCollection `presentIn:"body"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response ListSupportDocsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListSupportDocsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListSupportDocsSortOrderEnum Enum with underlying type: string
type ListSupportDocsSortOrderEnum string

// Set of constants representing the allowable values for ListSupportDocsSortOrderEnum
const (
	ListSupportDocsSortOrderAsc  ListSupportDocsSortOrderEnum = "ASC"
	ListSupportDocsSortOrderDesc ListSupportDocsSortOrderEnum = "DESC"
)

var mappingListSupportDocsSortOrderEnum = map[string]ListSupportDocsSortOrderEnum{
	"ASC":  ListSupportDocsSortOrderAsc,
	"DESC": ListSupportDocsSortOrderDesc,
}

var mappingListSupportDocsSortOrderEnumLowerCase = map[string]ListSupportDocsSortOrderEnum{
	"asc":  ListSupportDocsSortOrderAsc,
	"desc": ListSupportDocsSortOrderDesc,
}

// GetListSupportDocsSortOrderEnumValues Enumerates the set of values for ListSupportDocsSortOrderEnum
func GetListSupportDocsSortOrderEnumValues() []ListSupportDocsSortOrderEnum {
	values := make([]ListSupportDocsSortOrderEnum, 0)
	for _, v := range mappingListSupportDocsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListSupportDocsSortOrderEnumStringValues Enumerates the set of values in String for ListSupportDocsSortOrderEnum
func GetListSupportDocsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListSupportDocsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListSupportDocsSortOrderEnum(val string) (ListSupportDocsSortOrderEnum, bool) {
	enum, ok := mappingListSupportDocsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListSupportDocsSortByEnum Enum with underlying type: string
type ListSupportDocsSortByEnum string

// Set of constants representing the allowable values for ListSupportDocsSortByEnum
const (
	ListSupportDocsSortByTimecreated ListSupportDocsSortByEnum = "timeCreated"
	ListSupportDocsSortByDisplayname ListSupportDocsSortByEnum = "displayName"
)

var mappingListSupportDocsSortByEnum = map[string]ListSupportDocsSortByEnum{
	"timeCreated": ListSupportDocsSortByTimecreated,
	"displayName": ListSupportDocsSortByDisplayname,
}

var mappingListSupportDocsSortByEnumLowerCase = map[string]ListSupportDocsSortByEnum{
	"timecreated": ListSupportDocsSortByTimecreated,
	"displayname": ListSupportDocsSortByDisplayname,
}

// GetListSupportDocsSortByEnumValues Enumerates the set of values for ListSupportDocsSortByEnum
func GetListSupportDocsSortByEnumValues() []ListSupportDocsSortByEnum {
	values := make([]ListSupportDocsSortByEnum, 0)
	for _, v := range mappingListSupportDocsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListSupportDocsSortByEnumStringValues Enumerates the set of values in String for ListSupportDocsSortByEnum
func GetListSupportDocsSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListSupportDocsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListSupportDocsSortByEnum(val string) (ListSupportDocsSortByEnum, bool) {
	enum, ok := mappingListSupportDocsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
