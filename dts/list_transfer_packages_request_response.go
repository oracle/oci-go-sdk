// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package dts

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v57/common"
	"net/http"
	"strings"
)

// ListTransferPackagesRequest wrapper for the ListTransferPackages operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/dts/ListTransferPackages.go.html to see an example of how to use ListTransferPackagesRequest.
type ListTransferPackagesRequest struct {

	// ID of the Transfer Job
	Id *string `mandatory:"true" contributesTo:"path" name:"id"`

	// filtering by lifecycleState
	LifecycleState ListTransferPackagesLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// filtering by displayName
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// Unique Oracle-assigned identifier for the request.
	// If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListTransferPackagesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListTransferPackagesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListTransferPackagesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListTransferPackagesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListTransferPackagesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := mappingListTransferPackagesLifecycleStateEnum[string(request.LifecycleState)]; !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetListTransferPackagesLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListTransferPackagesResponse wrapper for the ListTransferPackages operation
type ListTransferPackagesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The MultipleTransferPackages instance
	MultipleTransferPackages `presentIn:"body"`

	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response ListTransferPackagesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListTransferPackagesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListTransferPackagesLifecycleStateEnum Enum with underlying type: string
type ListTransferPackagesLifecycleStateEnum string

// Set of constants representing the allowable values for ListTransferPackagesLifecycleStateEnum
const (
	ListTransferPackagesLifecycleStatePreparing         ListTransferPackagesLifecycleStateEnum = "PREPARING"
	ListTransferPackagesLifecycleStateShipping          ListTransferPackagesLifecycleStateEnum = "SHIPPING"
	ListTransferPackagesLifecycleStateReceived          ListTransferPackagesLifecycleStateEnum = "RECEIVED"
	ListTransferPackagesLifecycleStateProcessing        ListTransferPackagesLifecycleStateEnum = "PROCESSING"
	ListTransferPackagesLifecycleStateProcessed         ListTransferPackagesLifecycleStateEnum = "PROCESSED"
	ListTransferPackagesLifecycleStateReturned          ListTransferPackagesLifecycleStateEnum = "RETURNED"
	ListTransferPackagesLifecycleStateDeleted           ListTransferPackagesLifecycleStateEnum = "DELETED"
	ListTransferPackagesLifecycleStateCancelled         ListTransferPackagesLifecycleStateEnum = "CANCELLED"
	ListTransferPackagesLifecycleStateCancelledReturned ListTransferPackagesLifecycleStateEnum = "CANCELLED_RETURNED"
)

var mappingListTransferPackagesLifecycleStateEnum = map[string]ListTransferPackagesLifecycleStateEnum{
	"PREPARING":          ListTransferPackagesLifecycleStatePreparing,
	"SHIPPING":           ListTransferPackagesLifecycleStateShipping,
	"RECEIVED":           ListTransferPackagesLifecycleStateReceived,
	"PROCESSING":         ListTransferPackagesLifecycleStateProcessing,
	"PROCESSED":          ListTransferPackagesLifecycleStateProcessed,
	"RETURNED":           ListTransferPackagesLifecycleStateReturned,
	"DELETED":            ListTransferPackagesLifecycleStateDeleted,
	"CANCELLED":          ListTransferPackagesLifecycleStateCancelled,
	"CANCELLED_RETURNED": ListTransferPackagesLifecycleStateCancelledReturned,
}

// GetListTransferPackagesLifecycleStateEnumValues Enumerates the set of values for ListTransferPackagesLifecycleStateEnum
func GetListTransferPackagesLifecycleStateEnumValues() []ListTransferPackagesLifecycleStateEnum {
	values := make([]ListTransferPackagesLifecycleStateEnum, 0)
	for _, v := range mappingListTransferPackagesLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetListTransferPackagesLifecycleStateEnumStringValues Enumerates the set of values in String for ListTransferPackagesLifecycleStateEnum
func GetListTransferPackagesLifecycleStateEnumStringValues() []string {
	return []string{
		"PREPARING",
		"SHIPPING",
		"RECEIVED",
		"PROCESSING",
		"PROCESSED",
		"RETURNED",
		"DELETED",
		"CANCELLED",
		"CANCELLED_RETURNED",
	}
}
