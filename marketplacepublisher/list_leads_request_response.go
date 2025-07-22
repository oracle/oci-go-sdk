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

// ListLeadsRequest wrapper for the ListLeads operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/marketplacepublisher/ListLeads.go.html to see an example of how to use ListLeadsRequest.
type ListLeadsRequest struct {

	// The ID of the compartment in which to list leads.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// Listing OCID to query resource against.
	ListingId *string `mandatory:"false" contributesTo:"query" name:"listingId"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// A token representing the position at which to start retrieving results. This must come from the `opc-next-page` header field of a previous response.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'ASC' or 'DESC'.
	SortOrder ListLeadsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for timeCreated is descending.
	SortBy ListLeadsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListLeadsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListLeadsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListLeadsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListLeadsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListLeadsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListLeadsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListLeadsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListLeadsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListLeadsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListLeadsResponse wrapper for the ListLeads operation
type ListLeadsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of LeadCollection instances
	LeadCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListLeadsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListLeadsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListLeadsSortOrderEnum Enum with underlying type: string
type ListLeadsSortOrderEnum string

// Set of constants representing the allowable values for ListLeadsSortOrderEnum
const (
	ListLeadsSortOrderAsc  ListLeadsSortOrderEnum = "ASC"
	ListLeadsSortOrderDesc ListLeadsSortOrderEnum = "DESC"
)

var mappingListLeadsSortOrderEnum = map[string]ListLeadsSortOrderEnum{
	"ASC":  ListLeadsSortOrderAsc,
	"DESC": ListLeadsSortOrderDesc,
}

var mappingListLeadsSortOrderEnumLowerCase = map[string]ListLeadsSortOrderEnum{
	"asc":  ListLeadsSortOrderAsc,
	"desc": ListLeadsSortOrderDesc,
}

// GetListLeadsSortOrderEnumValues Enumerates the set of values for ListLeadsSortOrderEnum
func GetListLeadsSortOrderEnumValues() []ListLeadsSortOrderEnum {
	values := make([]ListLeadsSortOrderEnum, 0)
	for _, v := range mappingListLeadsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListLeadsSortOrderEnumStringValues Enumerates the set of values in String for ListLeadsSortOrderEnum
func GetListLeadsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListLeadsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListLeadsSortOrderEnum(val string) (ListLeadsSortOrderEnum, bool) {
	enum, ok := mappingListLeadsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListLeadsSortByEnum Enum with underlying type: string
type ListLeadsSortByEnum string

// Set of constants representing the allowable values for ListLeadsSortByEnum
const (
	ListLeadsSortByTimecreated ListLeadsSortByEnum = "timeCreated"
)

var mappingListLeadsSortByEnum = map[string]ListLeadsSortByEnum{
	"timeCreated": ListLeadsSortByTimecreated,
}

var mappingListLeadsSortByEnumLowerCase = map[string]ListLeadsSortByEnum{
	"timecreated": ListLeadsSortByTimecreated,
}

// GetListLeadsSortByEnumValues Enumerates the set of values for ListLeadsSortByEnum
func GetListLeadsSortByEnumValues() []ListLeadsSortByEnum {
	values := make([]ListLeadsSortByEnum, 0)
	for _, v := range mappingListLeadsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListLeadsSortByEnumStringValues Enumerates the set of values in String for ListLeadsSortByEnum
func GetListLeadsSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
	}
}

// GetMappingListLeadsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListLeadsSortByEnum(val string) (ListLeadsSortByEnum, bool) {
	enum, ok := mappingListLeadsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
