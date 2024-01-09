// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package marketplacepublisher

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListListingRevisionNotesRequest wrapper for the ListListingRevisionNotes operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/marketplacepublisher/ListListingRevisionNotes.go.html to see an example of how to use ListListingRevisionNotesRequest.
type ListListingRevisionNotesRequest struct {

	// OCID of the listing revision.
	ListingRevisionId *string `mandatory:"true" contributesTo:"query" name:"listingRevisionId"`

	// The unique identifier for the compartment.
	CompartmentId *string `mandatory:"false" contributesTo:"query" name:"compartmentId"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// A token representing the position at which to start retrieving results. This must come from the `opc-next-page` header field of a previous response.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'ASC' or 'DESC'.
	SortOrder ListListingRevisionNotesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for timeCreated is descending. Default order for displayName is ascending.
	SortBy ListListingRevisionNotesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListListingRevisionNotesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListListingRevisionNotesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListListingRevisionNotesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListListingRevisionNotesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListListingRevisionNotesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListListingRevisionNotesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListListingRevisionNotesSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListListingRevisionNotesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListListingRevisionNotesSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListListingRevisionNotesResponse wrapper for the ListListingRevisionNotes operation
type ListListingRevisionNotesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of ListingRevisionNoteCollection instances
	ListingRevisionNoteCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListListingRevisionNotesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListListingRevisionNotesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListListingRevisionNotesSortOrderEnum Enum with underlying type: string
type ListListingRevisionNotesSortOrderEnum string

// Set of constants representing the allowable values for ListListingRevisionNotesSortOrderEnum
const (
	ListListingRevisionNotesSortOrderAsc  ListListingRevisionNotesSortOrderEnum = "ASC"
	ListListingRevisionNotesSortOrderDesc ListListingRevisionNotesSortOrderEnum = "DESC"
)

var mappingListListingRevisionNotesSortOrderEnum = map[string]ListListingRevisionNotesSortOrderEnum{
	"ASC":  ListListingRevisionNotesSortOrderAsc,
	"DESC": ListListingRevisionNotesSortOrderDesc,
}

var mappingListListingRevisionNotesSortOrderEnumLowerCase = map[string]ListListingRevisionNotesSortOrderEnum{
	"asc":  ListListingRevisionNotesSortOrderAsc,
	"desc": ListListingRevisionNotesSortOrderDesc,
}

// GetListListingRevisionNotesSortOrderEnumValues Enumerates the set of values for ListListingRevisionNotesSortOrderEnum
func GetListListingRevisionNotesSortOrderEnumValues() []ListListingRevisionNotesSortOrderEnum {
	values := make([]ListListingRevisionNotesSortOrderEnum, 0)
	for _, v := range mappingListListingRevisionNotesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListListingRevisionNotesSortOrderEnumStringValues Enumerates the set of values in String for ListListingRevisionNotesSortOrderEnum
func GetListListingRevisionNotesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListListingRevisionNotesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListListingRevisionNotesSortOrderEnum(val string) (ListListingRevisionNotesSortOrderEnum, bool) {
	enum, ok := mappingListListingRevisionNotesSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListListingRevisionNotesSortByEnum Enum with underlying type: string
type ListListingRevisionNotesSortByEnum string

// Set of constants representing the allowable values for ListListingRevisionNotesSortByEnum
const (
	ListListingRevisionNotesSortByTimecreated ListListingRevisionNotesSortByEnum = "timeCreated"
	ListListingRevisionNotesSortByDisplayname ListListingRevisionNotesSortByEnum = "displayName"
)

var mappingListListingRevisionNotesSortByEnum = map[string]ListListingRevisionNotesSortByEnum{
	"timeCreated": ListListingRevisionNotesSortByTimecreated,
	"displayName": ListListingRevisionNotesSortByDisplayname,
}

var mappingListListingRevisionNotesSortByEnumLowerCase = map[string]ListListingRevisionNotesSortByEnum{
	"timecreated": ListListingRevisionNotesSortByTimecreated,
	"displayname": ListListingRevisionNotesSortByDisplayname,
}

// GetListListingRevisionNotesSortByEnumValues Enumerates the set of values for ListListingRevisionNotesSortByEnum
func GetListListingRevisionNotesSortByEnumValues() []ListListingRevisionNotesSortByEnum {
	values := make([]ListListingRevisionNotesSortByEnum, 0)
	for _, v := range mappingListListingRevisionNotesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListListingRevisionNotesSortByEnumStringValues Enumerates the set of values in String for ListListingRevisionNotesSortByEnum
func GetListListingRevisionNotesSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListListingRevisionNotesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListListingRevisionNotesSortByEnum(val string) (ListListingRevisionNotesSortByEnum, bool) {
	enum, ok := mappingListListingRevisionNotesSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
