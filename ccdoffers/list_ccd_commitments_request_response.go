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

// ListCcdCommitmentsRequest wrapper for the ListCcdCommitments operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/ccdoffers/ListCcdCommitments.go.html to see an example of how to use ListCcdCommitmentsRequest.
type ListCcdCommitmentsRequest struct {

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

	// The lifecycle of a CCD commitment that's being requested, which always defaults to Active. If you want to filter by
	// scheduled, current, awaiting provisioning, canceled, or ended, use the `lifecycleDetails` query parameter.
	LifecycleState CcdCommitmentLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// The lifecycle details of a CCD commitment that's being requested. The accepted values are Scheduled, Current,
	// Awaiting Provisioning, Canceled, or Ended.
	// If a null or empty value is passed, it will return for all resources.
	LifecycleDetails CcdCommitmentLifecycleDetailsEnum `mandatory:"false" contributesTo:"query" name:"lifecycleDetails" omitEmpty:"true"`

	// Filter CCD offers or commitments by the commercial workflow that created the associated offer.
	SourceType CcdOfferSourceTypeEnum `mandatory:"false" contributesTo:"query" name:"sourceType" omitEmpty:"true"`

	// The sort order to use, whether "asc" or "desc". By default, it is ascending.
	SortOrder ListCcdCommitmentsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// Filter by the resource display name.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// The sort by field to use, whether "DISPLAYNAME" or "TIMECREATED". By default, it is TIMECREATED.
	SortBy ListCcdCommitmentsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListCcdCommitmentsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListCcdCommitmentsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListCcdCommitmentsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListCcdCommitmentsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListCcdCommitmentsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingCcdCommitmentLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetCcdCommitmentLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingCcdCommitmentLifecycleDetailsEnum(string(request.LifecycleDetails)); !ok && request.LifecycleDetails != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleDetails: %s. Supported values are: %s.", request.LifecycleDetails, strings.Join(GetCcdCommitmentLifecycleDetailsEnumStringValues(), ",")))
	}
	if _, ok := GetMappingCcdOfferSourceTypeEnum(string(request.SourceType)); !ok && request.SourceType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SourceType: %s. Supported values are: %s.", request.SourceType, strings.Join(GetCcdOfferSourceTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListCcdCommitmentsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListCcdCommitmentsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListCcdCommitmentsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListCcdCommitmentsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListCcdCommitmentsResponse wrapper for the ListCcdCommitments operation
type ListCcdCommitmentsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of CcdCommitmentCollection instances
	CcdCommitmentCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListCcdCommitmentsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListCcdCommitmentsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListCcdCommitmentsSortOrderEnum Enum with underlying type: string
type ListCcdCommitmentsSortOrderEnum string

// Set of constants representing the allowable values for ListCcdCommitmentsSortOrderEnum
const (
	ListCcdCommitmentsSortOrderAsc  ListCcdCommitmentsSortOrderEnum = "ASC"
	ListCcdCommitmentsSortOrderDesc ListCcdCommitmentsSortOrderEnum = "DESC"
)

var mappingListCcdCommitmentsSortOrderEnum = map[string]ListCcdCommitmentsSortOrderEnum{
	"ASC":  ListCcdCommitmentsSortOrderAsc,
	"DESC": ListCcdCommitmentsSortOrderDesc,
}

var mappingListCcdCommitmentsSortOrderEnumLowerCase = map[string]ListCcdCommitmentsSortOrderEnum{
	"asc":  ListCcdCommitmentsSortOrderAsc,
	"desc": ListCcdCommitmentsSortOrderDesc,
}

// GetListCcdCommitmentsSortOrderEnumValues Enumerates the set of values for ListCcdCommitmentsSortOrderEnum
func GetListCcdCommitmentsSortOrderEnumValues() []ListCcdCommitmentsSortOrderEnum {
	values := make([]ListCcdCommitmentsSortOrderEnum, 0)
	for _, v := range mappingListCcdCommitmentsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListCcdCommitmentsSortOrderEnumStringValues Enumerates the set of values in String for ListCcdCommitmentsSortOrderEnum
func GetListCcdCommitmentsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListCcdCommitmentsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListCcdCommitmentsSortOrderEnum(val string) (ListCcdCommitmentsSortOrderEnum, bool) {
	enum, ok := mappingListCcdCommitmentsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListCcdCommitmentsSortByEnum Enum with underlying type: string
type ListCcdCommitmentsSortByEnum string

// Set of constants representing the allowable values for ListCcdCommitmentsSortByEnum
const (
	ListCcdCommitmentsSortByDisplayname ListCcdCommitmentsSortByEnum = "DISPLAYNAME"
	ListCcdCommitmentsSortByTimecreated ListCcdCommitmentsSortByEnum = "TIMECREATED"
)

var mappingListCcdCommitmentsSortByEnum = map[string]ListCcdCommitmentsSortByEnum{
	"DISPLAYNAME": ListCcdCommitmentsSortByDisplayname,
	"TIMECREATED": ListCcdCommitmentsSortByTimecreated,
}

var mappingListCcdCommitmentsSortByEnumLowerCase = map[string]ListCcdCommitmentsSortByEnum{
	"displayname": ListCcdCommitmentsSortByDisplayname,
	"timecreated": ListCcdCommitmentsSortByTimecreated,
}

// GetListCcdCommitmentsSortByEnumValues Enumerates the set of values for ListCcdCommitmentsSortByEnum
func GetListCcdCommitmentsSortByEnumValues() []ListCcdCommitmentsSortByEnum {
	values := make([]ListCcdCommitmentsSortByEnum, 0)
	for _, v := range mappingListCcdCommitmentsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListCcdCommitmentsSortByEnumStringValues Enumerates the set of values in String for ListCcdCommitmentsSortByEnum
func GetListCcdCommitmentsSortByEnumStringValues() []string {
	return []string{
		"DISPLAYNAME",
		"TIMECREATED",
	}
}

// GetMappingListCcdCommitmentsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListCcdCommitmentsSortByEnum(val string) (ListCcdCommitmentsSortByEnum, bool) {
	enum, ok := mappingListCcdCommitmentsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
