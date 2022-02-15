// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package dts

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"net/http"
	"strings"
)

// ListTransferJobsRequest wrapper for the ListTransferJobs operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/dts/ListTransferJobs.go.html to see an example of how to use ListTransferJobsRequest.
type ListTransferJobsRequest struct {

	// compartment id
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// filtering by lifecycleState
	LifecycleState ListTransferJobsLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// filtering by displayName
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// For list pagination. The maximum number of results per page, or items to return in a paginated
	// "List" call. For important details about how pagination works, see
	// List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	// Example: `50`
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The value of the `opc-next-page` response header from the previous "List"
	// call. For important details about how pagination works, see
	// List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	// a particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListTransferJobsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListTransferJobsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListTransferJobsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListTransferJobsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListTransferJobsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListTransferJobsLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetListTransferJobsLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListTransferJobsResponse wrapper for the ListTransferJobs operation
type ListTransferJobsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []TransferJobSummary instances
	Items []TransferJobSummary `presentIn:"body"`

	// For list pagination. When this header appears in the response, additional pages
	// of results remain. For important details about how pagination works, see
	// List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response ListTransferJobsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListTransferJobsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListTransferJobsLifecycleStateEnum Enum with underlying type: string
type ListTransferJobsLifecycleStateEnum string

// Set of constants representing the allowable values for ListTransferJobsLifecycleStateEnum
const (
	ListTransferJobsLifecycleStateInitiated ListTransferJobsLifecycleStateEnum = "INITIATED"
	ListTransferJobsLifecycleStatePreparing ListTransferJobsLifecycleStateEnum = "PREPARING"
	ListTransferJobsLifecycleStateActive    ListTransferJobsLifecycleStateEnum = "ACTIVE"
	ListTransferJobsLifecycleStateDeleted   ListTransferJobsLifecycleStateEnum = "DELETED"
	ListTransferJobsLifecycleStateClosed    ListTransferJobsLifecycleStateEnum = "CLOSED"
)

var mappingListTransferJobsLifecycleStateEnum = map[string]ListTransferJobsLifecycleStateEnum{
	"INITIATED": ListTransferJobsLifecycleStateInitiated,
	"PREPARING": ListTransferJobsLifecycleStatePreparing,
	"ACTIVE":    ListTransferJobsLifecycleStateActive,
	"DELETED":   ListTransferJobsLifecycleStateDeleted,
	"CLOSED":    ListTransferJobsLifecycleStateClosed,
}

// GetListTransferJobsLifecycleStateEnumValues Enumerates the set of values for ListTransferJobsLifecycleStateEnum
func GetListTransferJobsLifecycleStateEnumValues() []ListTransferJobsLifecycleStateEnum {
	values := make([]ListTransferJobsLifecycleStateEnum, 0)
	for _, v := range mappingListTransferJobsLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetListTransferJobsLifecycleStateEnumStringValues Enumerates the set of values in String for ListTransferJobsLifecycleStateEnum
func GetListTransferJobsLifecycleStateEnumStringValues() []string {
	return []string{
		"INITIATED",
		"PREPARING",
		"ACTIVE",
		"DELETED",
		"CLOSED",
	}
}

// GetMappingListTransferJobsLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListTransferJobsLifecycleStateEnum(val string) (ListTransferJobsLifecycleStateEnum, bool) {
	mappingListTransferJobsLifecycleStateEnumIgnoreCase := make(map[string]ListTransferJobsLifecycleStateEnum)
	for k, v := range mappingListTransferJobsLifecycleStateEnum {
		mappingListTransferJobsLifecycleStateEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingListTransferJobsLifecycleStateEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
