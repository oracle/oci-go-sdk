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

// ListAdminListingRevisionsRequest wrapper for the ListAdminListingRevisions operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/marketplacepublisher/ListAdminListingRevisions.go.html to see an example of how to use ListAdminListingRevisionsRequest.
type ListAdminListingRevisionsRequest struct {

	// listing OCID
	ListingId *string `mandatory:"false" contributesTo:"query" name:"listingId"`

	// The unique identifier for the compartment.
	CompartmentId *string `mandatory:"false" contributesTo:"query" name:"compartmentId"`

	// A filter to return listings based on their type
	ListingType ListAdminListingRevisionsListingTypeEnum `mandatory:"false" contributesTo:"query" name:"listingType" omitEmpty:"true"`

	// A filter to return only resources that match the entire display name given.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// Unique publisher identifier.
	PublisherId *string `mandatory:"false" contributesTo:"query" name:"publisherId"`

	// A filter to return only listing revisions their lifecycleState matches the given lifecycleState.
	LifecycleState ListingRevisionLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// A filter to return only listing revisions their status matches the given listing revision status.
	ListingRevisionStatus ListingRevisionStatusEnum `mandatory:"false" contributesTo:"query" name:"listingRevisionStatus" omitEmpty:"true"`

	// Comma-separated listing revision status values used to filter OpenSearch admin listing revision results.
	StatusCsv *string `mandatory:"false" contributesTo:"query" name:"statusCsv"`

	// Product value used to filter OpenSearch admin listing revision results.
	Product *string `mandatory:"false" contributesTo:"query" name:"product"`

	// Filters data created after the specified date.
	TimeCreatedAfter *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timeCreatedAfter"`

	// Filters data updated after the specified date.
	TimeUpdatedAfter *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timeUpdatedAfter"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// A token representing the position at which to start retrieving results. This must come from the `opc-next-page` header field of a previous response.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'ASC' or 'DESC'.
	SortOrder ListAdminListingRevisionsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for timeCreated is descending. Default order for timeUpdated is descending. Default order for displayName is ascending.
	SortBy ListAdminListingRevisionsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Optional pricing type filter for listing admin listing revisions. This filter is only applied when dataSource is opensearch.
	PricingType *string `mandatory:"false" contributesTo:"query" name:"pricingType"`

	// Optional private offer enabled filter for listing admin listing revisions. This filter is only applied when dataSource is opensearch.
	IsPrivateOfferEnabled *bool `mandatory:"false" contributesTo:"query" name:"isPrivateOfferEnabled"`

	// Optional source for listing admin listing revisions. Use opensearch to read from OpenSearch.
	DataSource *string `mandatory:"false" contributesTo:"query" name:"dataSource"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListAdminListingRevisionsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListAdminListingRevisionsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListAdminListingRevisionsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListAdminListingRevisionsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListAdminListingRevisionsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListAdminListingRevisionsListingTypeEnum(string(request.ListingType)); !ok && request.ListingType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ListingType: %s. Supported values are: %s.", request.ListingType, strings.Join(GetListAdminListingRevisionsListingTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListingRevisionLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetListingRevisionLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListingRevisionStatusEnum(string(request.ListingRevisionStatus)); !ok && request.ListingRevisionStatus != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ListingRevisionStatus: %s. Supported values are: %s.", request.ListingRevisionStatus, strings.Join(GetListingRevisionStatusEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListAdminListingRevisionsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListAdminListingRevisionsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListAdminListingRevisionsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListAdminListingRevisionsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListAdminListingRevisionsResponse wrapper for the ListAdminListingRevisions operation
type ListAdminListingRevisionsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of AdminListingRevisionCollection instances
	AdminListingRevisionCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListAdminListingRevisionsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListAdminListingRevisionsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListAdminListingRevisionsListingTypeEnum Enum with underlying type: string
type ListAdminListingRevisionsListingTypeEnum string

// Set of constants representing the allowable values for ListAdminListingRevisionsListingTypeEnum
const (
	ListAdminListingRevisionsListingTypeOciApplication ListAdminListingRevisionsListingTypeEnum = "OCI_APPLICATION"
	ListAdminListingRevisionsListingTypeLeadGeneration ListAdminListingRevisionsListingTypeEnum = "LEAD_GENERATION"
	ListAdminListingRevisionsListingTypeService        ListAdminListingRevisionsListingTypeEnum = "SERVICE"
)

var mappingListAdminListingRevisionsListingTypeEnum = map[string]ListAdminListingRevisionsListingTypeEnum{
	"OCI_APPLICATION": ListAdminListingRevisionsListingTypeOciApplication,
	"LEAD_GENERATION": ListAdminListingRevisionsListingTypeLeadGeneration,
	"SERVICE":         ListAdminListingRevisionsListingTypeService,
}

var mappingListAdminListingRevisionsListingTypeEnumLowerCase = map[string]ListAdminListingRevisionsListingTypeEnum{
	"oci_application": ListAdminListingRevisionsListingTypeOciApplication,
	"lead_generation": ListAdminListingRevisionsListingTypeLeadGeneration,
	"service":         ListAdminListingRevisionsListingTypeService,
}

// GetListAdminListingRevisionsListingTypeEnumValues Enumerates the set of values for ListAdminListingRevisionsListingTypeEnum
func GetListAdminListingRevisionsListingTypeEnumValues() []ListAdminListingRevisionsListingTypeEnum {
	values := make([]ListAdminListingRevisionsListingTypeEnum, 0)
	for _, v := range mappingListAdminListingRevisionsListingTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetListAdminListingRevisionsListingTypeEnumStringValues Enumerates the set of values in String for ListAdminListingRevisionsListingTypeEnum
func GetListAdminListingRevisionsListingTypeEnumStringValues() []string {
	return []string{
		"OCI_APPLICATION",
		"LEAD_GENERATION",
		"SERVICE",
	}
}

// GetMappingListAdminListingRevisionsListingTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListAdminListingRevisionsListingTypeEnum(val string) (ListAdminListingRevisionsListingTypeEnum, bool) {
	enum, ok := mappingListAdminListingRevisionsListingTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListAdminListingRevisionsSortOrderEnum Enum with underlying type: string
type ListAdminListingRevisionsSortOrderEnum string

// Set of constants representing the allowable values for ListAdminListingRevisionsSortOrderEnum
const (
	ListAdminListingRevisionsSortOrderAsc  ListAdminListingRevisionsSortOrderEnum = "ASC"
	ListAdminListingRevisionsSortOrderDesc ListAdminListingRevisionsSortOrderEnum = "DESC"
)

var mappingListAdminListingRevisionsSortOrderEnum = map[string]ListAdminListingRevisionsSortOrderEnum{
	"ASC":  ListAdminListingRevisionsSortOrderAsc,
	"DESC": ListAdminListingRevisionsSortOrderDesc,
}

var mappingListAdminListingRevisionsSortOrderEnumLowerCase = map[string]ListAdminListingRevisionsSortOrderEnum{
	"asc":  ListAdminListingRevisionsSortOrderAsc,
	"desc": ListAdminListingRevisionsSortOrderDesc,
}

// GetListAdminListingRevisionsSortOrderEnumValues Enumerates the set of values for ListAdminListingRevisionsSortOrderEnum
func GetListAdminListingRevisionsSortOrderEnumValues() []ListAdminListingRevisionsSortOrderEnum {
	values := make([]ListAdminListingRevisionsSortOrderEnum, 0)
	for _, v := range mappingListAdminListingRevisionsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListAdminListingRevisionsSortOrderEnumStringValues Enumerates the set of values in String for ListAdminListingRevisionsSortOrderEnum
func GetListAdminListingRevisionsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListAdminListingRevisionsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListAdminListingRevisionsSortOrderEnum(val string) (ListAdminListingRevisionsSortOrderEnum, bool) {
	enum, ok := mappingListAdminListingRevisionsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListAdminListingRevisionsSortByEnum Enum with underlying type: string
type ListAdminListingRevisionsSortByEnum string

// Set of constants representing the allowable values for ListAdminListingRevisionsSortByEnum
const (
	ListAdminListingRevisionsSortByTimecreated ListAdminListingRevisionsSortByEnum = "timeCreated"
	ListAdminListingRevisionsSortByTimeupdated ListAdminListingRevisionsSortByEnum = "timeUpdated"
	ListAdminListingRevisionsSortByDisplayname ListAdminListingRevisionsSortByEnum = "displayName"
)

var mappingListAdminListingRevisionsSortByEnum = map[string]ListAdminListingRevisionsSortByEnum{
	"timeCreated": ListAdminListingRevisionsSortByTimecreated,
	"timeUpdated": ListAdminListingRevisionsSortByTimeupdated,
	"displayName": ListAdminListingRevisionsSortByDisplayname,
}

var mappingListAdminListingRevisionsSortByEnumLowerCase = map[string]ListAdminListingRevisionsSortByEnum{
	"timecreated": ListAdminListingRevisionsSortByTimecreated,
	"timeupdated": ListAdminListingRevisionsSortByTimeupdated,
	"displayname": ListAdminListingRevisionsSortByDisplayname,
}

// GetListAdminListingRevisionsSortByEnumValues Enumerates the set of values for ListAdminListingRevisionsSortByEnum
func GetListAdminListingRevisionsSortByEnumValues() []ListAdminListingRevisionsSortByEnum {
	values := make([]ListAdminListingRevisionsSortByEnum, 0)
	for _, v := range mappingListAdminListingRevisionsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListAdminListingRevisionsSortByEnumStringValues Enumerates the set of values in String for ListAdminListingRevisionsSortByEnum
func GetListAdminListingRevisionsSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"timeUpdated",
		"displayName",
	}
}

// GetMappingListAdminListingRevisionsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListAdminListingRevisionsSortByEnum(val string) (ListAdminListingRevisionsSortByEnum, bool) {
	enum, ok := mappingListAdminListingRevisionsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
