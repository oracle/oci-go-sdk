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

// ListListingRevisionPackagesRequest wrapper for the ListListingRevisionPackages operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/marketplacepublisher/ListListingRevisionPackages.go.html to see an example of how to use ListListingRevisionPackagesRequest.
type ListListingRevisionPackagesRequest struct {

	// OCID of the listing revision.
	ListingRevisionId *string `mandatory:"true" contributesTo:"query" name:"listingRevisionId"`

	// The unique identifier for the compartment.
	CompartmentId *string `mandatory:"false" contributesTo:"query" name:"compartmentId"`

	// A filter to return only resources that match the entire display name given.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// A filter to return only packages their lifecycleState matches the given lifecycleState.
	LifecycleState ListingRevisionPackageLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// A token representing the position at which to start retrieving results. This must come from the `opc-next-page` header field of a previous response.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'ASC' or 'DESC'.
	SortOrder ListListingRevisionPackagesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for timeCreated is descending. Default order for displayName is ascending.
	SortBy ListListingRevisionPackagesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListListingRevisionPackagesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListListingRevisionPackagesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListListingRevisionPackagesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListListingRevisionPackagesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListListingRevisionPackagesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListingRevisionPackageLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetListingRevisionPackageLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListListingRevisionPackagesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListListingRevisionPackagesSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListListingRevisionPackagesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListListingRevisionPackagesSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListListingRevisionPackagesResponse wrapper for the ListListingRevisionPackages operation
type ListListingRevisionPackagesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of ListingRevisionPackageCollection instances
	ListingRevisionPackageCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListListingRevisionPackagesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListListingRevisionPackagesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListListingRevisionPackagesSortOrderEnum Enum with underlying type: string
type ListListingRevisionPackagesSortOrderEnum string

// Set of constants representing the allowable values for ListListingRevisionPackagesSortOrderEnum
const (
	ListListingRevisionPackagesSortOrderAsc  ListListingRevisionPackagesSortOrderEnum = "ASC"
	ListListingRevisionPackagesSortOrderDesc ListListingRevisionPackagesSortOrderEnum = "DESC"
)

var mappingListListingRevisionPackagesSortOrderEnum = map[string]ListListingRevisionPackagesSortOrderEnum{
	"ASC":  ListListingRevisionPackagesSortOrderAsc,
	"DESC": ListListingRevisionPackagesSortOrderDesc,
}

var mappingListListingRevisionPackagesSortOrderEnumLowerCase = map[string]ListListingRevisionPackagesSortOrderEnum{
	"asc":  ListListingRevisionPackagesSortOrderAsc,
	"desc": ListListingRevisionPackagesSortOrderDesc,
}

// GetListListingRevisionPackagesSortOrderEnumValues Enumerates the set of values for ListListingRevisionPackagesSortOrderEnum
func GetListListingRevisionPackagesSortOrderEnumValues() []ListListingRevisionPackagesSortOrderEnum {
	values := make([]ListListingRevisionPackagesSortOrderEnum, 0)
	for _, v := range mappingListListingRevisionPackagesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListListingRevisionPackagesSortOrderEnumStringValues Enumerates the set of values in String for ListListingRevisionPackagesSortOrderEnum
func GetListListingRevisionPackagesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListListingRevisionPackagesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListListingRevisionPackagesSortOrderEnum(val string) (ListListingRevisionPackagesSortOrderEnum, bool) {
	enum, ok := mappingListListingRevisionPackagesSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListListingRevisionPackagesSortByEnum Enum with underlying type: string
type ListListingRevisionPackagesSortByEnum string

// Set of constants representing the allowable values for ListListingRevisionPackagesSortByEnum
const (
	ListListingRevisionPackagesSortByTimecreated ListListingRevisionPackagesSortByEnum = "timeCreated"
	ListListingRevisionPackagesSortByDisplayname ListListingRevisionPackagesSortByEnum = "displayName"
)

var mappingListListingRevisionPackagesSortByEnum = map[string]ListListingRevisionPackagesSortByEnum{
	"timeCreated": ListListingRevisionPackagesSortByTimecreated,
	"displayName": ListListingRevisionPackagesSortByDisplayname,
}

var mappingListListingRevisionPackagesSortByEnumLowerCase = map[string]ListListingRevisionPackagesSortByEnum{
	"timecreated": ListListingRevisionPackagesSortByTimecreated,
	"displayname": ListListingRevisionPackagesSortByDisplayname,
}

// GetListListingRevisionPackagesSortByEnumValues Enumerates the set of values for ListListingRevisionPackagesSortByEnum
func GetListListingRevisionPackagesSortByEnumValues() []ListListingRevisionPackagesSortByEnum {
	values := make([]ListListingRevisionPackagesSortByEnum, 0)
	for _, v := range mappingListListingRevisionPackagesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListListingRevisionPackagesSortByEnumStringValues Enumerates the set of values in String for ListListingRevisionPackagesSortByEnum
func GetListListingRevisionPackagesSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListListingRevisionPackagesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListListingRevisionPackagesSortByEnum(val string) (ListListingRevisionPackagesSortByEnum, bool) {
	enum, ok := mappingListListingRevisionPackagesSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
