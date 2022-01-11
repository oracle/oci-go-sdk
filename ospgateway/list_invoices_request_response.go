// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package ospgateway

import (
	"github.com/oracle/oci-go-sdk/v55/common"
	"net/http"
)

// ListInvoicesRequest wrapper for the ListInvoices operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/ospgateway/ListInvoices.go.html to see an example of how to use ListInvoicesRequest.
type ListInvoicesRequest struct {

	// The home region's public name of the logged in user.
	OspHomeRegion *string `mandatory:"true" contributesTo:"query" name:"ospHomeRegion"`

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// The unique Oracle-assigned identifier for the request. If you need to contact Oracle about a
	// particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// The invoice query param (not unique).
	InvoiceId *string `mandatory:"false" contributesTo:"query" name:"invoiceId"`

	// A filter to only return resources that match the given type exactly.
	Type []ListInvoicesTypeEnum `contributesTo:"query" name:"type" omitEmpty:"true" collectionFormat:"multi"`

	// A filter to only return resources that match the given value.
	// Looking for partial matches in the following fileds:
	// Invoice No., Reference No. (plan number), Payment Ref, Total Amount(plan number), Balance Due(plan number)
	// and Party/Customer Name
	SearchText *string `mandatory:"false" contributesTo:"query" name:"searchText"`

	// description: Start time (UTC) of the target invoice date range for which to fetch invoice data (inclusive).
	TimeInvoiceStart *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timeInvoiceStart"`

	// description: End time (UTC) of the target invoice date range for which to fetch invoice data (exclusive).
	TimeInvoiceEnd *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timeInvoiceEnd"`

	// description: Start time (UTC) of the target payment date range for which to fetch invoice data (inclusive).
	TimePaymentStart *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timePaymentStart"`

	// description: End time (UTC) of the target payment date range for which to fetch invoice data (exclusive).
	TimePaymentEnd *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timePaymentEnd"`

	// A filter to only return resources that match one of the status elements.
	Status []ListInvoicesStatusEnum `contributesTo:"query" name:"status" omitEmpty:"true" collectionFormat:"multi"`

	// For list pagination. The value of the opc-next-page response header from the previous "List" call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// For list pagination. The maximum number of results per page, or items to return in a paginated "List" call.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The field to sort by. Only one field can be selected for sorting.
	SortBy ListInvoicesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to use (ascending or descending).
	SortOrder ListInvoicesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListInvoicesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListInvoicesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListInvoicesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListInvoicesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ListInvoicesResponse wrapper for the ListInvoices operation
type ListInvoicesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of InvoiceCollection instances
	InvoiceCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a
	// particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. If this header appears in the response, then this
	// is a partial list of invoices. Include this value as the `page` parameter in a subsequent
	// GET request to get the next batch of invoices.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// The total number of items that match the query.
	OpcTotalItems *int `presentIn:"header" name:"opc-total-items"`
}

func (response ListInvoicesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListInvoicesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListInvoicesTypeEnum Enum with underlying type: string
type ListInvoicesTypeEnum string

// Set of constants representing the allowable values for ListInvoicesTypeEnum
const (
	ListInvoicesTypeHardware     ListInvoicesTypeEnum = "HARDWARE"
	ListInvoicesTypeSubscription ListInvoicesTypeEnum = "SUBSCRIPTION"
	ListInvoicesTypeSupport      ListInvoicesTypeEnum = "SUPPORT"
	ListInvoicesTypeLicense      ListInvoicesTypeEnum = "LICENSE"
	ListInvoicesTypeEducation    ListInvoicesTypeEnum = "EDUCATION"
	ListInvoicesTypeConsulting   ListInvoicesTypeEnum = "CONSULTING"
	ListInvoicesTypeService      ListInvoicesTypeEnum = "SERVICE"
	ListInvoicesTypeUsage        ListInvoicesTypeEnum = "USAGE"
)

var mappingListInvoicesType = map[string]ListInvoicesTypeEnum{
	"HARDWARE":     ListInvoicesTypeHardware,
	"SUBSCRIPTION": ListInvoicesTypeSubscription,
	"SUPPORT":      ListInvoicesTypeSupport,
	"LICENSE":      ListInvoicesTypeLicense,
	"EDUCATION":    ListInvoicesTypeEducation,
	"CONSULTING":   ListInvoicesTypeConsulting,
	"SERVICE":      ListInvoicesTypeService,
	"USAGE":        ListInvoicesTypeUsage,
}

// GetListInvoicesTypeEnumValues Enumerates the set of values for ListInvoicesTypeEnum
func GetListInvoicesTypeEnumValues() []ListInvoicesTypeEnum {
	values := make([]ListInvoicesTypeEnum, 0)
	for _, v := range mappingListInvoicesType {
		values = append(values, v)
	}
	return values
}

// ListInvoicesStatusEnum Enum with underlying type: string
type ListInvoicesStatusEnum string

// Set of constants representing the allowable values for ListInvoicesStatusEnum
const (
	ListInvoicesStatusOpen             ListInvoicesStatusEnum = "OPEN"
	ListInvoicesStatusPastDue          ListInvoicesStatusEnum = "PAST_DUE"
	ListInvoicesStatusPaymentSubmitted ListInvoicesStatusEnum = "PAYMENT_SUBMITTED"
	ListInvoicesStatusClosed           ListInvoicesStatusEnum = "CLOSED"
)

var mappingListInvoicesStatus = map[string]ListInvoicesStatusEnum{
	"OPEN":              ListInvoicesStatusOpen,
	"PAST_DUE":          ListInvoicesStatusPastDue,
	"PAYMENT_SUBMITTED": ListInvoicesStatusPaymentSubmitted,
	"CLOSED":            ListInvoicesStatusClosed,
}

// GetListInvoicesStatusEnumValues Enumerates the set of values for ListInvoicesStatusEnum
func GetListInvoicesStatusEnumValues() []ListInvoicesStatusEnum {
	values := make([]ListInvoicesStatusEnum, 0)
	for _, v := range mappingListInvoicesStatus {
		values = append(values, v)
	}
	return values
}

// ListInvoicesSortByEnum Enum with underlying type: string
type ListInvoicesSortByEnum string

// Set of constants representing the allowable values for ListInvoicesSortByEnum
const (
	ListInvoicesSortByInvoiceNo   ListInvoicesSortByEnum = "INVOICE_NO"
	ListInvoicesSortByRefNo       ListInvoicesSortByEnum = "REF_NO"
	ListInvoicesSortByStatus      ListInvoicesSortByEnum = "STATUS"
	ListInvoicesSortByType        ListInvoicesSortByEnum = "TYPE"
	ListInvoicesSortByInvoiceDate ListInvoicesSortByEnum = "INVOICE_DATE"
	ListInvoicesSortByDueDate     ListInvoicesSortByEnum = "DUE_DATE"
	ListInvoicesSortByPaymRef     ListInvoicesSortByEnum = "PAYM_REF"
	ListInvoicesSortByTotalAmount ListInvoicesSortByEnum = "TOTAL_AMOUNT"
	ListInvoicesSortByBalanceDue  ListInvoicesSortByEnum = "BALANCE_DUE"
)

var mappingListInvoicesSortBy = map[string]ListInvoicesSortByEnum{
	"INVOICE_NO":   ListInvoicesSortByInvoiceNo,
	"REF_NO":       ListInvoicesSortByRefNo,
	"STATUS":       ListInvoicesSortByStatus,
	"TYPE":         ListInvoicesSortByType,
	"INVOICE_DATE": ListInvoicesSortByInvoiceDate,
	"DUE_DATE":     ListInvoicesSortByDueDate,
	"PAYM_REF":     ListInvoicesSortByPaymRef,
	"TOTAL_AMOUNT": ListInvoicesSortByTotalAmount,
	"BALANCE_DUE":  ListInvoicesSortByBalanceDue,
}

// GetListInvoicesSortByEnumValues Enumerates the set of values for ListInvoicesSortByEnum
func GetListInvoicesSortByEnumValues() []ListInvoicesSortByEnum {
	values := make([]ListInvoicesSortByEnum, 0)
	for _, v := range mappingListInvoicesSortBy {
		values = append(values, v)
	}
	return values
}

// ListInvoicesSortOrderEnum Enum with underlying type: string
type ListInvoicesSortOrderEnum string

// Set of constants representing the allowable values for ListInvoicesSortOrderEnum
const (
	ListInvoicesSortOrderAsc  ListInvoicesSortOrderEnum = "ASC"
	ListInvoicesSortOrderDesc ListInvoicesSortOrderEnum = "DESC"
)

var mappingListInvoicesSortOrder = map[string]ListInvoicesSortOrderEnum{
	"ASC":  ListInvoicesSortOrderAsc,
	"DESC": ListInvoicesSortOrderDesc,
}

// GetListInvoicesSortOrderEnumValues Enumerates the set of values for ListInvoicesSortOrderEnum
func GetListInvoicesSortOrderEnumValues() []ListInvoicesSortOrderEnum {
	values := make([]ListInvoicesSortOrderEnum, 0)
	for _, v := range mappingListInvoicesSortOrder {
		values = append(values, v)
	}
	return values
}
