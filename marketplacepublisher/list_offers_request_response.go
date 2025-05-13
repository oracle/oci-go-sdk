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

// ListOffersRequest wrapper for the ListOffers operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/marketplacepublisher/ListOffers.go.html to see an example of how to use ListOffersRequest.
type ListOffersRequest struct {

	// The ID of the buyer compartment this offer is associated with.
	BuyerCompartmentId *string `mandatory:"false" contributesTo:"query" name:"buyerCompartmentId"`

	// The ID of the seller compartment this offer is associated with.
	SellerCompartmentId *string `mandatory:"false" contributesTo:"query" name:"sellerCompartmentId"`

	// A filter to return only resources their lifecycleState matches the given lifecycleState.
	LifecycleState OfferLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// A filter to return only resources that match the entire display name given.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// unique Offer identifier
	Id *string `mandatory:"false" contributesTo:"query" name:"id"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// A token representing the position at which to start retrieving results. This must come from the `opc-next-page` header field of a previous response.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'ASC' or 'DESC'.
	SortOrder ListOffersSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for timeCreated is descending. Default order for displayName is ascending.
	SortBy ListOffersSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListOffersRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListOffersRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListOffersRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListOffersRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListOffersRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingOfferLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetOfferLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListOffersSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListOffersSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListOffersSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListOffersSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListOffersResponse wrapper for the ListOffers operation
type ListOffersResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of OfferCollection instances
	OfferCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListOffersResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListOffersResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListOffersSortOrderEnum Enum with underlying type: string
type ListOffersSortOrderEnum string

// Set of constants representing the allowable values for ListOffersSortOrderEnum
const (
	ListOffersSortOrderAsc  ListOffersSortOrderEnum = "ASC"
	ListOffersSortOrderDesc ListOffersSortOrderEnum = "DESC"
)

var mappingListOffersSortOrderEnum = map[string]ListOffersSortOrderEnum{
	"ASC":  ListOffersSortOrderAsc,
	"DESC": ListOffersSortOrderDesc,
}

var mappingListOffersSortOrderEnumLowerCase = map[string]ListOffersSortOrderEnum{
	"asc":  ListOffersSortOrderAsc,
	"desc": ListOffersSortOrderDesc,
}

// GetListOffersSortOrderEnumValues Enumerates the set of values for ListOffersSortOrderEnum
func GetListOffersSortOrderEnumValues() []ListOffersSortOrderEnum {
	values := make([]ListOffersSortOrderEnum, 0)
	for _, v := range mappingListOffersSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListOffersSortOrderEnumStringValues Enumerates the set of values in String for ListOffersSortOrderEnum
func GetListOffersSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListOffersSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListOffersSortOrderEnum(val string) (ListOffersSortOrderEnum, bool) {
	enum, ok := mappingListOffersSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListOffersSortByEnum Enum with underlying type: string
type ListOffersSortByEnum string

// Set of constants representing the allowable values for ListOffersSortByEnum
const (
	ListOffersSortByTimecreated ListOffersSortByEnum = "timeCreated"
	ListOffersSortByDisplayname ListOffersSortByEnum = "displayName"
)

var mappingListOffersSortByEnum = map[string]ListOffersSortByEnum{
	"timeCreated": ListOffersSortByTimecreated,
	"displayName": ListOffersSortByDisplayname,
}

var mappingListOffersSortByEnumLowerCase = map[string]ListOffersSortByEnum{
	"timecreated": ListOffersSortByTimecreated,
	"displayname": ListOffersSortByDisplayname,
}

// GetListOffersSortByEnumValues Enumerates the set of values for ListOffersSortByEnum
func GetListOffersSortByEnumValues() []ListOffersSortByEnum {
	values := make([]ListOffersSortByEnum, 0)
	for _, v := range mappingListOffersSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListOffersSortByEnumStringValues Enumerates the set of values in String for ListOffersSortByEnum
func GetListOffersSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListOffersSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListOffersSortByEnum(val string) (ListOffersSortByEnum, bool) {
	enum, ok := mappingListOffersSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
