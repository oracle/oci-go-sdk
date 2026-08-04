// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package marketplaceprivateoffer

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListQuoteAttachmentsRequest wrapper for the ListQuoteAttachments operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/marketplaceprivateoffer/ListQuoteAttachments.go.html to see an example of how to use ListQuoteAttachmentsRequest.
type ListQuoteAttachmentsRequest struct {

	// Unique offer quote identifier.
	OfferQuoteId *string `mandatory:"true" contributesTo:"path" name:"offerQuoteId"`

	// OCID of the reseller's compartment.
	ResellerCompartmentId *string `mandatory:"false" contributesTo:"query" name:"resellerCompartmentId"`

	// OCID of the ISV's compartment.
	IsvCompartmentId *string `mandatory:"false" contributesTo:"query" name:"isvCompartmentId"`

	// A filter to return only resources their lifecycleState matches the given lifecycleState.
	LifecycleState QuoteAttachmentLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// A filter to return only resources that match the entire display name given.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// Unique offer quote attachment identifier.
	Id *string `mandatory:"false" contributesTo:"query" name:"id"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// A token representing the position at which to start retrieving results. This must come from the `opc-next-page` header field of a previous response.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'ASC' or 'DESC'.
	SortOrder ListQuoteAttachmentsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for timeCreated is descending. Default order for displayName is ascending.
	SortBy ListQuoteAttachmentsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListQuoteAttachmentsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListQuoteAttachmentsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListQuoteAttachmentsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListQuoteAttachmentsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListQuoteAttachmentsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingQuoteAttachmentLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetQuoteAttachmentLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListQuoteAttachmentsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListQuoteAttachmentsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListQuoteAttachmentsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListQuoteAttachmentsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListQuoteAttachmentsResponse wrapper for the ListQuoteAttachments operation
type ListQuoteAttachmentsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of QuoteAttachmentCollection instances
	QuoteAttachmentCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListQuoteAttachmentsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListQuoteAttachmentsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListQuoteAttachmentsSortOrderEnum Enum with underlying type: string
type ListQuoteAttachmentsSortOrderEnum string

// Set of constants representing the allowable values for ListQuoteAttachmentsSortOrderEnum
const (
	ListQuoteAttachmentsSortOrderAsc  ListQuoteAttachmentsSortOrderEnum = "ASC"
	ListQuoteAttachmentsSortOrderDesc ListQuoteAttachmentsSortOrderEnum = "DESC"
)

var mappingListQuoteAttachmentsSortOrderEnum = map[string]ListQuoteAttachmentsSortOrderEnum{
	"ASC":  ListQuoteAttachmentsSortOrderAsc,
	"DESC": ListQuoteAttachmentsSortOrderDesc,
}

var mappingListQuoteAttachmentsSortOrderEnumLowerCase = map[string]ListQuoteAttachmentsSortOrderEnum{
	"asc":  ListQuoteAttachmentsSortOrderAsc,
	"desc": ListQuoteAttachmentsSortOrderDesc,
}

// GetListQuoteAttachmentsSortOrderEnumValues Enumerates the set of values for ListQuoteAttachmentsSortOrderEnum
func GetListQuoteAttachmentsSortOrderEnumValues() []ListQuoteAttachmentsSortOrderEnum {
	values := make([]ListQuoteAttachmentsSortOrderEnum, 0)
	for _, v := range mappingListQuoteAttachmentsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListQuoteAttachmentsSortOrderEnumStringValues Enumerates the set of values in String for ListQuoteAttachmentsSortOrderEnum
func GetListQuoteAttachmentsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListQuoteAttachmentsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListQuoteAttachmentsSortOrderEnum(val string) (ListQuoteAttachmentsSortOrderEnum, bool) {
	enum, ok := mappingListQuoteAttachmentsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListQuoteAttachmentsSortByEnum Enum with underlying type: string
type ListQuoteAttachmentsSortByEnum string

// Set of constants representing the allowable values for ListQuoteAttachmentsSortByEnum
const (
	ListQuoteAttachmentsSortByTimecreated ListQuoteAttachmentsSortByEnum = "timeCreated"
	ListQuoteAttachmentsSortByDisplayname ListQuoteAttachmentsSortByEnum = "displayName"
)

var mappingListQuoteAttachmentsSortByEnum = map[string]ListQuoteAttachmentsSortByEnum{
	"timeCreated": ListQuoteAttachmentsSortByTimecreated,
	"displayName": ListQuoteAttachmentsSortByDisplayname,
}

var mappingListQuoteAttachmentsSortByEnumLowerCase = map[string]ListQuoteAttachmentsSortByEnum{
	"timecreated": ListQuoteAttachmentsSortByTimecreated,
	"displayname": ListQuoteAttachmentsSortByDisplayname,
}

// GetListQuoteAttachmentsSortByEnumValues Enumerates the set of values for ListQuoteAttachmentsSortByEnum
func GetListQuoteAttachmentsSortByEnumValues() []ListQuoteAttachmentsSortByEnum {
	values := make([]ListQuoteAttachmentsSortByEnum, 0)
	for _, v := range mappingListQuoteAttachmentsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListQuoteAttachmentsSortByEnumStringValues Enumerates the set of values in String for ListQuoteAttachmentsSortByEnum
func GetListQuoteAttachmentsSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListQuoteAttachmentsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListQuoteAttachmentsSortByEnum(val string) (ListQuoteAttachmentsSortByEnum, bool) {
	enum, ok := mappingListQuoteAttachmentsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
