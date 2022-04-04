// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package dts

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListApplianceExportJobsRequest wrapper for the ListApplianceExportJobs operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/dts/ListApplianceExportJobs.go.html to see an example of how to use ListApplianceExportJobsRequest.
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
func (request ListApplianceExportJobsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListApplianceExportJobsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListApplianceExportJobsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListApplianceExportJobsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListApplianceExportJobsLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetListApplianceExportJobsLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
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

var mappingListApplianceExportJobsLifecycleStateEnum = map[string]ListApplianceExportJobsLifecycleStateEnum{
	"CREATING":   ListApplianceExportJobsLifecycleStateCreating,
	"ACTIVE":     ListApplianceExportJobsLifecycleStateActive,
	"INPROGRESS": ListApplianceExportJobsLifecycleStateInprogress,
	"SUCCEEDED":  ListApplianceExportJobsLifecycleStateSucceeded,
	"FAILED":     ListApplianceExportJobsLifecycleStateFailed,
	"CANCELLED":  ListApplianceExportJobsLifecycleStateCancelled,
	"DELETED":    ListApplianceExportJobsLifecycleStateDeleted,
}

var mappingListApplianceExportJobsLifecycleStateEnumLowerCase = map[string]ListApplianceExportJobsLifecycleStateEnum{
	"creating":   ListApplianceExportJobsLifecycleStateCreating,
	"active":     ListApplianceExportJobsLifecycleStateActive,
	"inprogress": ListApplianceExportJobsLifecycleStateInprogress,
	"succeeded":  ListApplianceExportJobsLifecycleStateSucceeded,
	"failed":     ListApplianceExportJobsLifecycleStateFailed,
	"cancelled":  ListApplianceExportJobsLifecycleStateCancelled,
	"deleted":    ListApplianceExportJobsLifecycleStateDeleted,
}

// GetListApplianceExportJobsLifecycleStateEnumValues Enumerates the set of values for ListApplianceExportJobsLifecycleStateEnum
func GetListApplianceExportJobsLifecycleStateEnumValues() []ListApplianceExportJobsLifecycleStateEnum {
	values := make([]ListApplianceExportJobsLifecycleStateEnum, 0)
	for _, v := range mappingListApplianceExportJobsLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetListApplianceExportJobsLifecycleStateEnumStringValues Enumerates the set of values in String for ListApplianceExportJobsLifecycleStateEnum
func GetListApplianceExportJobsLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"ACTIVE",
		"INPROGRESS",
		"SUCCEEDED",
		"FAILED",
		"CANCELLED",
		"DELETED",
	}
}

// GetMappingListApplianceExportJobsLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListApplianceExportJobsLifecycleStateEnum(val string) (ListApplianceExportJobsLifecycleStateEnum, bool) {
	enum, ok := mappingListApplianceExportJobsLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
