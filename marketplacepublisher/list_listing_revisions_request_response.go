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

// ListListingRevisionsRequest wrapper for the ListListingRevisions operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/marketplacepublisher/ListListingRevisions.go.html to see an example of how to use ListListingRevisionsRequest.
type ListListingRevisionsRequest struct {

	// listing OCID
	ListingId *string `mandatory:"true" contributesTo:"query" name:"listingId"`

	// A filter to return only resources that match the entire display name given.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// A filter to return only listing revisions their lifecycleState matches the given lifecycleState.
	LifecycleState ListingRevisionLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// The unique identifier for the compartment.
	CompartmentId *string `mandatory:"false" contributesTo:"query" name:"compartmentId"`

	// A filter to return only listing revisions their status matches the given listing revision status.
	ListingRevisionStatus ListingRevisionStatusEnum `mandatory:"false" contributesTo:"query" name:"listingRevisionStatus" omitEmpty:"true"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// A token representing the position at which to start retrieving results. This must come from the `opc-next-page` header field of a previous response.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'ASC' or 'DESC'.
	SortOrder ListListingRevisionsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for timeCreated is descending. Default order for displayName is ascending.
	SortBy ListListingRevisionsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListListingRevisionsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListListingRevisionsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListListingRevisionsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListListingRevisionsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListListingRevisionsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListingRevisionLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetListingRevisionLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListingRevisionStatusEnum(string(request.ListingRevisionStatus)); !ok && request.ListingRevisionStatus != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ListingRevisionStatus: %s. Supported values are: %s.", request.ListingRevisionStatus, strings.Join(GetListingRevisionStatusEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListListingRevisionsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListListingRevisionsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListListingRevisionsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListListingRevisionsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListListingRevisionsResponse wrapper for the ListListingRevisions operation
type ListListingRevisionsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of ListingRevisionCollection instances
	ListingRevisionCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListListingRevisionsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListListingRevisionsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListListingRevisionsSortOrderEnum Enum with underlying type: string
type ListListingRevisionsSortOrderEnum string

// Set of constants representing the allowable values for ListListingRevisionsSortOrderEnum
const (
	ListListingRevisionsSortOrderAsc  ListListingRevisionsSortOrderEnum = "ASC"
	ListListingRevisionsSortOrderDesc ListListingRevisionsSortOrderEnum = "DESC"
)

var mappingListListingRevisionsSortOrderEnum = map[string]ListListingRevisionsSortOrderEnum{
	"ASC":  ListListingRevisionsSortOrderAsc,
	"DESC": ListListingRevisionsSortOrderDesc,
}

var mappingListListingRevisionsSortOrderEnumLowerCase = map[string]ListListingRevisionsSortOrderEnum{
	"asc":  ListListingRevisionsSortOrderAsc,
	"desc": ListListingRevisionsSortOrderDesc,
}

// GetListListingRevisionsSortOrderEnumValues Enumerates the set of values for ListListingRevisionsSortOrderEnum
func GetListListingRevisionsSortOrderEnumValues() []ListListingRevisionsSortOrderEnum {
	values := make([]ListListingRevisionsSortOrderEnum, 0)
	for _, v := range mappingListListingRevisionsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListListingRevisionsSortOrderEnumStringValues Enumerates the set of values in String for ListListingRevisionsSortOrderEnum
func GetListListingRevisionsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListListingRevisionsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListListingRevisionsSortOrderEnum(val string) (ListListingRevisionsSortOrderEnum, bool) {
	enum, ok := mappingListListingRevisionsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListListingRevisionsSortByEnum Enum with underlying type: string
type ListListingRevisionsSortByEnum string

// Set of constants representing the allowable values for ListListingRevisionsSortByEnum
const (
	ListListingRevisionsSortByTimecreated ListListingRevisionsSortByEnum = "timeCreated"
	ListListingRevisionsSortByDisplayname ListListingRevisionsSortByEnum = "displayName"
)

var mappingListListingRevisionsSortByEnum = map[string]ListListingRevisionsSortByEnum{
	"timeCreated": ListListingRevisionsSortByTimecreated,
	"displayName": ListListingRevisionsSortByDisplayname,
}

var mappingListListingRevisionsSortByEnumLowerCase = map[string]ListListingRevisionsSortByEnum{
	"timecreated": ListListingRevisionsSortByTimecreated,
	"displayname": ListListingRevisionsSortByDisplayname,
}

// GetListListingRevisionsSortByEnumValues Enumerates the set of values for ListListingRevisionsSortByEnum
func GetListListingRevisionsSortByEnumValues() []ListListingRevisionsSortByEnum {
	values := make([]ListListingRevisionsSortByEnum, 0)
	for _, v := range mappingListListingRevisionsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListListingRevisionsSortByEnumStringValues Enumerates the set of values in String for ListListingRevisionsSortByEnum
func GetListListingRevisionsSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListListingRevisionsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListListingRevisionsSortByEnum(val string) (ListListingRevisionsSortByEnum, bool) {
	enum, ok := mappingListListingRevisionsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
