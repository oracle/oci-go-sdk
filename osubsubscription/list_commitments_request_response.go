// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package osubsubscription

import (
	"github.com/oracle/oci-go-sdk/v56/common"
	"net/http"
)

// ListCommitmentsRequest wrapper for the ListCommitments operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/osubsubscription/ListCommitments.go.html to see an example of how to use ListCommitmentsRequest.
type ListCommitmentsRequest struct {

	// This param is used to get the commitments for a particular subscribed service
	SubscribedServiceId *string `mandatory:"true" contributesTo:"query" name:"subscribedServiceId"`

	// The OCID of the compartment.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// The maximum number of items to return in a paginated "List" call. Default: (`50`)
	// Example: `500`
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The value of the `opc-next-page` response header from the previous "List" call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either ascending (`ASC`) or descending (`DESC`).
	SortOrder ListCommitmentsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. You can provide one sort order (`sortOrder`).
	SortBy ListCommitmentsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// This header is meant to be used only for internal purposes and will be ignored on any public request. The purpose of this header is
	// to help on Gateway to API calls identification.
	XOneGatewaySubscriptionId *string `mandatory:"false" contributesTo:"header" name:"x-one-gateway-subscription-id"`

	// The OCI home region name in case home region is not us-ashburn-1 (IAD), e.g. ap-mumbai-1, us-phoenix-1 etc.
	XOneOriginRegion *string `mandatory:"false" contributesTo:"header" name:"x-one-origin-region"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListCommitmentsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListCommitmentsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListCommitmentsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListCommitmentsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ListCommitmentsResponse wrapper for the ListCommitments operation
type ListCommitmentsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []CommitmentSummary instances
	Items []CommitmentSummary `presentIn:"body"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a
	// particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response ListCommitmentsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListCommitmentsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListCommitmentsSortOrderEnum Enum with underlying type: string
type ListCommitmentsSortOrderEnum string

// Set of constants representing the allowable values for ListCommitmentsSortOrderEnum
const (
	ListCommitmentsSortOrderAsc  ListCommitmentsSortOrderEnum = "ASC"
	ListCommitmentsSortOrderDesc ListCommitmentsSortOrderEnum = "DESC"
)

var mappingListCommitmentsSortOrder = map[string]ListCommitmentsSortOrderEnum{
	"ASC":  ListCommitmentsSortOrderAsc,
	"DESC": ListCommitmentsSortOrderDesc,
}

// GetListCommitmentsSortOrderEnumValues Enumerates the set of values for ListCommitmentsSortOrderEnum
func GetListCommitmentsSortOrderEnumValues() []ListCommitmentsSortOrderEnum {
	values := make([]ListCommitmentsSortOrderEnum, 0)
	for _, v := range mappingListCommitmentsSortOrder {
		values = append(values, v)
	}
	return values
}

// ListCommitmentsSortByEnum Enum with underlying type: string
type ListCommitmentsSortByEnum string

// Set of constants representing the allowable values for ListCommitmentsSortByEnum
const (
	ListCommitmentsSortByTimecreated ListCommitmentsSortByEnum = "TIMECREATED"
	ListCommitmentsSortByTimestart   ListCommitmentsSortByEnum = "TIMESTART"
)

var mappingListCommitmentsSortBy = map[string]ListCommitmentsSortByEnum{
	"TIMECREATED": ListCommitmentsSortByTimecreated,
	"TIMESTART":   ListCommitmentsSortByTimestart,
}

// GetListCommitmentsSortByEnumValues Enumerates the set of values for ListCommitmentsSortByEnum
func GetListCommitmentsSortByEnumValues() []ListCommitmentsSortByEnum {
	values := make([]ListCommitmentsSortByEnum, 0)
	for _, v := range mappingListCommitmentsSortBy {
		values = append(values, v)
	}
	return values
}
