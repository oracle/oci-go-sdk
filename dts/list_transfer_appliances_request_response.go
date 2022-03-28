// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package dts

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v64/common"
	"net/http"
	"strings"
)

// ListTransferAppliancesRequest wrapper for the ListTransferAppliances operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/dts/ListTransferAppliances.go.html to see an example of how to use ListTransferAppliancesRequest.
type ListTransferAppliancesRequest struct {

	// ID of the Transfer Job
	Id *string `mandatory:"true" contributesTo:"path" name:"id"`

	// filtering by lifecycleState
	LifecycleState ListTransferAppliancesLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// Unique Oracle-assigned identifier for the request.
	// If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListTransferAppliancesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListTransferAppliancesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListTransferAppliancesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListTransferAppliancesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListTransferAppliancesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListTransferAppliancesLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetListTransferAppliancesLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListTransferAppliancesResponse wrapper for the ListTransferAppliances operation
type ListTransferAppliancesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The MultipleTransferAppliances instance
	MultipleTransferAppliances `presentIn:"body"`

	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response ListTransferAppliancesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListTransferAppliancesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListTransferAppliancesLifecycleStateEnum Enum with underlying type: string
type ListTransferAppliancesLifecycleStateEnum string

// Set of constants representing the allowable values for ListTransferAppliancesLifecycleStateEnum
const (
	ListTransferAppliancesLifecycleStateRequested               ListTransferAppliancesLifecycleStateEnum = "REQUESTED"
	ListTransferAppliancesLifecycleStateOraclePreparing         ListTransferAppliancesLifecycleStateEnum = "ORACLE_PREPARING"
	ListTransferAppliancesLifecycleStateShipping                ListTransferAppliancesLifecycleStateEnum = "SHIPPING"
	ListTransferAppliancesLifecycleStateDelivered               ListTransferAppliancesLifecycleStateEnum = "DELIVERED"
	ListTransferAppliancesLifecycleStatePreparing               ListTransferAppliancesLifecycleStateEnum = "PREPARING"
	ListTransferAppliancesLifecycleStateFinalized               ListTransferAppliancesLifecycleStateEnum = "FINALIZED"
	ListTransferAppliancesLifecycleStateReturnDelayed           ListTransferAppliancesLifecycleStateEnum = "RETURN_DELAYED"
	ListTransferAppliancesLifecycleStateReturnShipped           ListTransferAppliancesLifecycleStateEnum = "RETURN_SHIPPED"
	ListTransferAppliancesLifecycleStateReturnShippedCancelled  ListTransferAppliancesLifecycleStateEnum = "RETURN_SHIPPED_CANCELLED"
	ListTransferAppliancesLifecycleStateOracleReceived          ListTransferAppliancesLifecycleStateEnum = "ORACLE_RECEIVED"
	ListTransferAppliancesLifecycleStateOracleReceivedCancelled ListTransferAppliancesLifecycleStateEnum = "ORACLE_RECEIVED_CANCELLED"
	ListTransferAppliancesLifecycleStateProcessing              ListTransferAppliancesLifecycleStateEnum = "PROCESSING"
	ListTransferAppliancesLifecycleStateComplete                ListTransferAppliancesLifecycleStateEnum = "COMPLETE"
	ListTransferAppliancesLifecycleStateCustomerNeverReceived   ListTransferAppliancesLifecycleStateEnum = "CUSTOMER_NEVER_RECEIVED"
	ListTransferAppliancesLifecycleStateOracleNeverReceived     ListTransferAppliancesLifecycleStateEnum = "ORACLE_NEVER_RECEIVED"
	ListTransferAppliancesLifecycleStateCustomerLost            ListTransferAppliancesLifecycleStateEnum = "CUSTOMER_LOST"
	ListTransferAppliancesLifecycleStateCancelled               ListTransferAppliancesLifecycleStateEnum = "CANCELLED"
	ListTransferAppliancesLifecycleStateDeleted                 ListTransferAppliancesLifecycleStateEnum = "DELETED"
	ListTransferAppliancesLifecycleStateRejected                ListTransferAppliancesLifecycleStateEnum = "REJECTED"
	ListTransferAppliancesLifecycleStateError                   ListTransferAppliancesLifecycleStateEnum = "ERROR"
)

var mappingListTransferAppliancesLifecycleStateEnum = map[string]ListTransferAppliancesLifecycleStateEnum{
	"REQUESTED":                 ListTransferAppliancesLifecycleStateRequested,
	"ORACLE_PREPARING":          ListTransferAppliancesLifecycleStateOraclePreparing,
	"SHIPPING":                  ListTransferAppliancesLifecycleStateShipping,
	"DELIVERED":                 ListTransferAppliancesLifecycleStateDelivered,
	"PREPARING":                 ListTransferAppliancesLifecycleStatePreparing,
	"FINALIZED":                 ListTransferAppliancesLifecycleStateFinalized,
	"RETURN_DELAYED":            ListTransferAppliancesLifecycleStateReturnDelayed,
	"RETURN_SHIPPED":            ListTransferAppliancesLifecycleStateReturnShipped,
	"RETURN_SHIPPED_CANCELLED":  ListTransferAppliancesLifecycleStateReturnShippedCancelled,
	"ORACLE_RECEIVED":           ListTransferAppliancesLifecycleStateOracleReceived,
	"ORACLE_RECEIVED_CANCELLED": ListTransferAppliancesLifecycleStateOracleReceivedCancelled,
	"PROCESSING":                ListTransferAppliancesLifecycleStateProcessing,
	"COMPLETE":                  ListTransferAppliancesLifecycleStateComplete,
	"CUSTOMER_NEVER_RECEIVED":   ListTransferAppliancesLifecycleStateCustomerNeverReceived,
	"ORACLE_NEVER_RECEIVED":     ListTransferAppliancesLifecycleStateOracleNeverReceived,
	"CUSTOMER_LOST":             ListTransferAppliancesLifecycleStateCustomerLost,
	"CANCELLED":                 ListTransferAppliancesLifecycleStateCancelled,
	"DELETED":                   ListTransferAppliancesLifecycleStateDeleted,
	"REJECTED":                  ListTransferAppliancesLifecycleStateRejected,
	"ERROR":                     ListTransferAppliancesLifecycleStateError,
}

var mappingListTransferAppliancesLifecycleStateEnumLowerCase = map[string]ListTransferAppliancesLifecycleStateEnum{
	"requested":                 ListTransferAppliancesLifecycleStateRequested,
	"oracle_preparing":          ListTransferAppliancesLifecycleStateOraclePreparing,
	"shipping":                  ListTransferAppliancesLifecycleStateShipping,
	"delivered":                 ListTransferAppliancesLifecycleStateDelivered,
	"preparing":                 ListTransferAppliancesLifecycleStatePreparing,
	"finalized":                 ListTransferAppliancesLifecycleStateFinalized,
	"return_delayed":            ListTransferAppliancesLifecycleStateReturnDelayed,
	"return_shipped":            ListTransferAppliancesLifecycleStateReturnShipped,
	"return_shipped_cancelled":  ListTransferAppliancesLifecycleStateReturnShippedCancelled,
	"oracle_received":           ListTransferAppliancesLifecycleStateOracleReceived,
	"oracle_received_cancelled": ListTransferAppliancesLifecycleStateOracleReceivedCancelled,
	"processing":                ListTransferAppliancesLifecycleStateProcessing,
	"complete":                  ListTransferAppliancesLifecycleStateComplete,
	"customer_never_received":   ListTransferAppliancesLifecycleStateCustomerNeverReceived,
	"oracle_never_received":     ListTransferAppliancesLifecycleStateOracleNeverReceived,
	"customer_lost":             ListTransferAppliancesLifecycleStateCustomerLost,
	"cancelled":                 ListTransferAppliancesLifecycleStateCancelled,
	"deleted":                   ListTransferAppliancesLifecycleStateDeleted,
	"rejected":                  ListTransferAppliancesLifecycleStateRejected,
	"error":                     ListTransferAppliancesLifecycleStateError,
}

// GetListTransferAppliancesLifecycleStateEnumValues Enumerates the set of values for ListTransferAppliancesLifecycleStateEnum
func GetListTransferAppliancesLifecycleStateEnumValues() []ListTransferAppliancesLifecycleStateEnum {
	values := make([]ListTransferAppliancesLifecycleStateEnum, 0)
	for _, v := range mappingListTransferAppliancesLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetListTransferAppliancesLifecycleStateEnumStringValues Enumerates the set of values in String for ListTransferAppliancesLifecycleStateEnum
func GetListTransferAppliancesLifecycleStateEnumStringValues() []string {
	return []string{
		"REQUESTED",
		"ORACLE_PREPARING",
		"SHIPPING",
		"DELIVERED",
		"PREPARING",
		"FINALIZED",
		"RETURN_DELAYED",
		"RETURN_SHIPPED",
		"RETURN_SHIPPED_CANCELLED",
		"ORACLE_RECEIVED",
		"ORACLE_RECEIVED_CANCELLED",
		"PROCESSING",
		"COMPLETE",
		"CUSTOMER_NEVER_RECEIVED",
		"ORACLE_NEVER_RECEIVED",
		"CUSTOMER_LOST",
		"CANCELLED",
		"DELETED",
		"REJECTED",
		"ERROR",
	}
}

// GetMappingListTransferAppliancesLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListTransferAppliancesLifecycleStateEnum(val string) (ListTransferAppliancesLifecycleStateEnum, bool) {
	enum, ok := mappingListTransferAppliancesLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
