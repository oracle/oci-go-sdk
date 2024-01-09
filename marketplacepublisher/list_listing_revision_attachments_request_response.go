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

// ListListingRevisionAttachmentsRequest wrapper for the ListListingRevisionAttachments operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/marketplacepublisher/ListListingRevisionAttachments.go.html to see an example of how to use ListListingRevisionAttachmentsRequest.
type ListListingRevisionAttachmentsRequest struct {

	// OCID of the listing revision.
	ListingRevisionId *string `mandatory:"true" contributesTo:"query" name:"listingRevisionId"`

	// A filter to return only resources that match the entire display name given.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// A filter to return only attachments their lifecycleState matches the given lifecycleState.
	LifecycleState ListingRevisionAttachmentLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// The unique identifier for the compartment.
	CompartmentId *string `mandatory:"false" contributesTo:"query" name:"compartmentId"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// A token representing the position at which to start retrieving results. This must come from the `opc-next-page` header field of a previous response.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'ASC' or 'DESC'.
	SortOrder ListListingRevisionAttachmentsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for timeCreated is descending. Default order for displayName is ascending.
	SortBy ListListingRevisionAttachmentsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListListingRevisionAttachmentsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListListingRevisionAttachmentsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListListingRevisionAttachmentsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListListingRevisionAttachmentsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListListingRevisionAttachmentsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListingRevisionAttachmentLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetListingRevisionAttachmentLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListListingRevisionAttachmentsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListListingRevisionAttachmentsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListListingRevisionAttachmentsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListListingRevisionAttachmentsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListListingRevisionAttachmentsResponse wrapper for the ListListingRevisionAttachments operation
type ListListingRevisionAttachmentsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of ListingRevisionAttachmentCollection instances
	ListingRevisionAttachmentCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListListingRevisionAttachmentsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListListingRevisionAttachmentsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListListingRevisionAttachmentsSortOrderEnum Enum with underlying type: string
type ListListingRevisionAttachmentsSortOrderEnum string

// Set of constants representing the allowable values for ListListingRevisionAttachmentsSortOrderEnum
const (
	ListListingRevisionAttachmentsSortOrderAsc  ListListingRevisionAttachmentsSortOrderEnum = "ASC"
	ListListingRevisionAttachmentsSortOrderDesc ListListingRevisionAttachmentsSortOrderEnum = "DESC"
)

var mappingListListingRevisionAttachmentsSortOrderEnum = map[string]ListListingRevisionAttachmentsSortOrderEnum{
	"ASC":  ListListingRevisionAttachmentsSortOrderAsc,
	"DESC": ListListingRevisionAttachmentsSortOrderDesc,
}

var mappingListListingRevisionAttachmentsSortOrderEnumLowerCase = map[string]ListListingRevisionAttachmentsSortOrderEnum{
	"asc":  ListListingRevisionAttachmentsSortOrderAsc,
	"desc": ListListingRevisionAttachmentsSortOrderDesc,
}

// GetListListingRevisionAttachmentsSortOrderEnumValues Enumerates the set of values for ListListingRevisionAttachmentsSortOrderEnum
func GetListListingRevisionAttachmentsSortOrderEnumValues() []ListListingRevisionAttachmentsSortOrderEnum {
	values := make([]ListListingRevisionAttachmentsSortOrderEnum, 0)
	for _, v := range mappingListListingRevisionAttachmentsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListListingRevisionAttachmentsSortOrderEnumStringValues Enumerates the set of values in String for ListListingRevisionAttachmentsSortOrderEnum
func GetListListingRevisionAttachmentsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListListingRevisionAttachmentsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListListingRevisionAttachmentsSortOrderEnum(val string) (ListListingRevisionAttachmentsSortOrderEnum, bool) {
	enum, ok := mappingListListingRevisionAttachmentsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListListingRevisionAttachmentsSortByEnum Enum with underlying type: string
type ListListingRevisionAttachmentsSortByEnum string

// Set of constants representing the allowable values for ListListingRevisionAttachmentsSortByEnum
const (
	ListListingRevisionAttachmentsSortByTimecreated ListListingRevisionAttachmentsSortByEnum = "timeCreated"
	ListListingRevisionAttachmentsSortByDisplayname ListListingRevisionAttachmentsSortByEnum = "displayName"
)

var mappingListListingRevisionAttachmentsSortByEnum = map[string]ListListingRevisionAttachmentsSortByEnum{
	"timeCreated": ListListingRevisionAttachmentsSortByTimecreated,
	"displayName": ListListingRevisionAttachmentsSortByDisplayname,
}

var mappingListListingRevisionAttachmentsSortByEnumLowerCase = map[string]ListListingRevisionAttachmentsSortByEnum{
	"timecreated": ListListingRevisionAttachmentsSortByTimecreated,
	"displayname": ListListingRevisionAttachmentsSortByDisplayname,
}

// GetListListingRevisionAttachmentsSortByEnumValues Enumerates the set of values for ListListingRevisionAttachmentsSortByEnum
func GetListListingRevisionAttachmentsSortByEnumValues() []ListListingRevisionAttachmentsSortByEnum {
	values := make([]ListListingRevisionAttachmentsSortByEnum, 0)
	for _, v := range mappingListListingRevisionAttachmentsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListListingRevisionAttachmentsSortByEnumStringValues Enumerates the set of values in String for ListListingRevisionAttachmentsSortByEnum
func GetListListingRevisionAttachmentsSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListListingRevisionAttachmentsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListListingRevisionAttachmentsSortByEnum(val string) (ListListingRevisionAttachmentsSortByEnum, bool) {
	enum, ok := mappingListListingRevisionAttachmentsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
