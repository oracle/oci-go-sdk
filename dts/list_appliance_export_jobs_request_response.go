// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package dts

import (
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

// ListApplianceExportJobsRequest wrapper for the ListApplianceExportJobs operation
type ListApplianceExportJobsRequest struct {

	// compartment id
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// filtering by lifecycleState
	LifecycleState ListApplianceExportJobsLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

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

func (request ListApplianceExportJobsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListApplianceExportJobsRequest) HTTPRequest(method, path string) (http.Request, error) {
	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListApplianceExportJobsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ListApplianceExportJobsResponse wrapper for the ListApplianceExportJobs operation
type ListApplianceExportJobsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []ApplianceExportJobSummary instances
	Items []ApplianceExportJobSummary `presentIn:"body"`

	// For list pagination. When this header appears in the response, additional pages
	// of results remain. For important details about how pagination works, see
	// List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response ListApplianceExportJobsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListApplianceExportJobsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListApplianceExportJobsLifecycleStateEnum Enum with underlying type: string
type ListApplianceExportJobsLifecycleStateEnum string

// Set of constants representing the allowable values for ListApplianceExportJobsLifecycleStateEnum
const (
	ListApplianceExportJobsLifecycleStateCreating   ListApplianceExportJobsLifecycleStateEnum = "CREATING"
	ListApplianceExportJobsLifecycleStateActive     ListApplianceExportJobsLifecycleStateEnum = "ACTIVE"
	ListApplianceExportJobsLifecycleStateInprogress ListApplianceExportJobsLifecycleStateEnum = "INPROGRESS"
	ListApplianceExportJobsLifecycleStateSucceeded  ListApplianceExportJobsLifecycleStateEnum = "SUCCEEDED"
	ListApplianceExportJobsLifecycleStateFailed     ListApplianceExportJobsLifecycleStateEnum = "FAILED"
	ListApplianceExportJobsLifecycleStateCancelled  ListApplianceExportJobsLifecycleStateEnum = "CANCELLED"
	ListApplianceExportJobsLifecycleStateDeleted    ListApplianceExportJobsLifecycleStateEnum = "DELETED"
)

var mappingListApplianceExportJobsLifecycleState = map[string]ListApplianceExportJobsLifecycleStateEnum{
	"CREATING":   ListApplianceExportJobsLifecycleStateCreating,
	"ACTIVE":     ListApplianceExportJobsLifecycleStateActive,
	"INPROGRESS": ListApplianceExportJobsLifecycleStateInprogress,
	"SUCCEEDED":  ListApplianceExportJobsLifecycleStateSucceeded,
	"FAILED":     ListApplianceExportJobsLifecycleStateFailed,
	"CANCELLED":  ListApplianceExportJobsLifecycleStateCancelled,
	"DELETED":    ListApplianceExportJobsLifecycleStateDeleted,
}

// GetListApplianceExportJobsLifecycleStateEnumValues Enumerates the set of values for ListApplianceExportJobsLifecycleStateEnum
func GetListApplianceExportJobsLifecycleStateEnumValues() []ListApplianceExportJobsLifecycleStateEnum {
	values := make([]ListApplianceExportJobsLifecycleStateEnum, 0)
	for _, v := range mappingListApplianceExportJobsLifecycleState {
		values = append(values, v)
	}
	return values
}
