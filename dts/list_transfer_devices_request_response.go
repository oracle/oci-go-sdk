// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package dts

import (
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

// ListTransferDevicesRequest wrapper for the ListTransferDevices operation
type ListTransferDevicesRequest struct {

	// ID of the Transfer Job
	Id *string `mandatory:"true" contributesTo:"path" name:"id"`

	// filtering by lifecycleState
	LifecycleState ListTransferDevicesLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// filtering by displayName
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// Unique Oracle-assigned identifier for the request.
	// If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListTransferDevicesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListTransferDevicesRequest) HTTPRequest(method, path string) (http.Request, error) {
	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListTransferDevicesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ListTransferDevicesResponse wrapper for the ListTransferDevices operation
type ListTransferDevicesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The MultipleTransferDevices instance
	MultipleTransferDevices `presentIn:"body"`

	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response ListTransferDevicesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListTransferDevicesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListTransferDevicesLifecycleStateEnum Enum with underlying type: string
type ListTransferDevicesLifecycleStateEnum string

// Set of constants representing the allowable values for ListTransferDevicesLifecycleStateEnum
const (
	ListTransferDevicesLifecycleStatePreparing  ListTransferDevicesLifecycleStateEnum = "PREPARING"
	ListTransferDevicesLifecycleStateReady      ListTransferDevicesLifecycleStateEnum = "READY"
	ListTransferDevicesLifecycleStatePackaged   ListTransferDevicesLifecycleStateEnum = "PACKAGED"
	ListTransferDevicesLifecycleStateActive     ListTransferDevicesLifecycleStateEnum = "ACTIVE"
	ListTransferDevicesLifecycleStateProcessing ListTransferDevicesLifecycleStateEnum = "PROCESSING"
	ListTransferDevicesLifecycleStateComplete   ListTransferDevicesLifecycleStateEnum = "COMPLETE"
	ListTransferDevicesLifecycleStateMissing    ListTransferDevicesLifecycleStateEnum = "MISSING"
	ListTransferDevicesLifecycleStateError      ListTransferDevicesLifecycleStateEnum = "ERROR"
	ListTransferDevicesLifecycleStateDeleted    ListTransferDevicesLifecycleStateEnum = "DELETED"
	ListTransferDevicesLifecycleStateCancelled  ListTransferDevicesLifecycleStateEnum = "CANCELLED"
)

var mappingListTransferDevicesLifecycleState = map[string]ListTransferDevicesLifecycleStateEnum{
	"PREPARING":  ListTransferDevicesLifecycleStatePreparing,
	"READY":      ListTransferDevicesLifecycleStateReady,
	"PACKAGED":   ListTransferDevicesLifecycleStatePackaged,
	"ACTIVE":     ListTransferDevicesLifecycleStateActive,
	"PROCESSING": ListTransferDevicesLifecycleStateProcessing,
	"COMPLETE":   ListTransferDevicesLifecycleStateComplete,
	"MISSING":    ListTransferDevicesLifecycleStateMissing,
	"ERROR":      ListTransferDevicesLifecycleStateError,
	"DELETED":    ListTransferDevicesLifecycleStateDeleted,
	"CANCELLED":  ListTransferDevicesLifecycleStateCancelled,
}

// GetListTransferDevicesLifecycleStateEnumValues Enumerates the set of values for ListTransferDevicesLifecycleStateEnum
func GetListTransferDevicesLifecycleStateEnumValues() []ListTransferDevicesLifecycleStateEnum {
	values := make([]ListTransferDevicesLifecycleStateEnum, 0)
	for _, v := range mappingListTransferDevicesLifecycleState {
		values = append(values, v)
	}
	return values
}
