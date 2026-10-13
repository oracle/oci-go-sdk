// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package ccdoffers

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListCcdOffersRequest wrapper for the ListCcdOffers operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/ccdoffers/ListCcdOffers.go.html to see an example of how to use ListCcdOffersRequest.
type ListCcdOffersRequest struct {

	// The subscription ID associated with the CCD offer.
	SubscriptionId *string `mandatory:"true" contributesTo:"query" name:"subscriptionId"`

	// The OCID of the compartment (remember that the tenancy is simply the root compartment).
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// The value of the `opc-next-page` response header from the previous "List" call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The maximum number of items to return in a paginated "List" call.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The life cycle of a CCD offer that's being requested, which always defaults to `Active`. If you want to filter by
	// scheduled, current, canceled, or ended, use the `lifecycleDetails` query parameter.
	LifecycleState CcdOfferLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// The life cycle details of a CCD offer that's being requested. The accepted values are Scheduled, Current, Canceled, or Ended.
	// If a null or empty value is passed, it will return for all resources.
	LifecycleDetails CcdOfferLifecycleDetailsEnum `mandatory:"false" contributesTo:"query" name:"lifecycleDetails" omitEmpty:"true"`

	// The CCD offer type to use when filtering results.
	TypeOfOffer CcdOfferTypeOfOfferEnum `mandatory:"false" contributesTo:"query" name:"typeOfOffer" omitEmpty:"true"`

	// Filter CCD offers or commitments by the commercial workflow that created the associated offer.
	SourceType CcdOfferSourceTypeEnum `mandatory:"false" contributesTo:"query" name:"sourceType" omitEmpty:"true"`

	// The sort order to use, whether "asc" or "desc". By default, it is ascending.
	SortOrder ListCcdOffersSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The sort by field to use, whether "DISPLAYNAME" or "TIMECREATED". By default, it is TIMECREATED.
	SortBy ListCcdOffersSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Filter by the resource display name.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListCcdOffersRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListCcdOffersRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListCcdOffersRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListCcdOffersRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListCcdOffersRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingCcdOfferLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetCcdOfferLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingCcdOfferLifecycleDetailsEnum(string(request.LifecycleDetails)); !ok && request.LifecycleDetails != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleDetails: %s. Supported values are: %s.", request.LifecycleDetails, strings.Join(GetCcdOfferLifecycleDetailsEnumStringValues(), ",")))
	}
	if _, ok := GetMappingCcdOfferTypeOfOfferEnum(string(request.TypeOfOffer)); !ok && request.TypeOfOffer != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for TypeOfOffer: %s. Supported values are: %s.", request.TypeOfOffer, strings.Join(GetCcdOfferTypeOfOfferEnumStringValues(), ",")))
	}
	if _, ok := GetMappingCcdOfferSourceTypeEnum(string(request.SourceType)); !ok && request.SourceType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SourceType: %s. Supported values are: %s.", request.SourceType, strings.Join(GetCcdOfferSourceTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListCcdOffersSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListCcdOffersSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListCcdOffersSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListCcdOffersSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListCcdOffersResponse wrapper for the ListCcdOffers operation
type ListCcdOffersResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of CcdOfferCollection instances
	CcdOfferCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListCcdOffersResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListCcdOffersResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListCcdOffersSortOrderEnum Enum with underlying type: string
type ListCcdOffersSortOrderEnum string

// Set of constants representing the allowable values for ListCcdOffersSortOrderEnum
const (
	ListCcdOffersSortOrderAsc  ListCcdOffersSortOrderEnum = "ASC"
	ListCcdOffersSortOrderDesc ListCcdOffersSortOrderEnum = "DESC"
)

var mappingListCcdOffersSortOrderEnum = map[string]ListCcdOffersSortOrderEnum{
	"ASC":  ListCcdOffersSortOrderAsc,
	"DESC": ListCcdOffersSortOrderDesc,
}

var mappingListCcdOffersSortOrderEnumLowerCase = map[string]ListCcdOffersSortOrderEnum{
	"asc":  ListCcdOffersSortOrderAsc,
	"desc": ListCcdOffersSortOrderDesc,
}

// GetListCcdOffersSortOrderEnumValues Enumerates the set of values for ListCcdOffersSortOrderEnum
func GetListCcdOffersSortOrderEnumValues() []ListCcdOffersSortOrderEnum {
	values := make([]ListCcdOffersSortOrderEnum, 0)
	for _, v := range mappingListCcdOffersSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListCcdOffersSortOrderEnumStringValues Enumerates the set of values in String for ListCcdOffersSortOrderEnum
func GetListCcdOffersSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListCcdOffersSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListCcdOffersSortOrderEnum(val string) (ListCcdOffersSortOrderEnum, bool) {
	enum, ok := mappingListCcdOffersSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListCcdOffersSortByEnum Enum with underlying type: string
type ListCcdOffersSortByEnum string

// Set of constants representing the allowable values for ListCcdOffersSortByEnum
const (
	ListCcdOffersSortByDisplayname ListCcdOffersSortByEnum = "DISPLAYNAME"
	ListCcdOffersSortByTimecreated ListCcdOffersSortByEnum = "TIMECREATED"
)

var mappingListCcdOffersSortByEnum = map[string]ListCcdOffersSortByEnum{
	"DISPLAYNAME": ListCcdOffersSortByDisplayname,
	"TIMECREATED": ListCcdOffersSortByTimecreated,
}

var mappingListCcdOffersSortByEnumLowerCase = map[string]ListCcdOffersSortByEnum{
	"displayname": ListCcdOffersSortByDisplayname,
	"timecreated": ListCcdOffersSortByTimecreated,
}

// GetListCcdOffersSortByEnumValues Enumerates the set of values for ListCcdOffersSortByEnum
func GetListCcdOffersSortByEnumValues() []ListCcdOffersSortByEnum {
	values := make([]ListCcdOffersSortByEnum, 0)
	for _, v := range mappingListCcdOffersSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListCcdOffersSortByEnumStringValues Enumerates the set of values in String for ListCcdOffersSortByEnum
func GetListCcdOffersSortByEnumStringValues() []string {
	return []string{
		"DISPLAYNAME",
		"TIMECREATED",
	}
}

// GetMappingListCcdOffersSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListCcdOffersSortByEnum(val string) (ListCcdOffersSortByEnum, bool) {
	enum, ok := mappingListCcdOffersSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
