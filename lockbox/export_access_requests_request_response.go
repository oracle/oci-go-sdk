// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package lockbox

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"io"
	"net/http"
	"strings"
)

// ExportAccessRequestsRequest wrapper for the ExportAccessRequests operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/lockbox/ExportAccessRequests.go.html to see an example of how to use ExportAccessRequestsRequest.
type ExportAccessRequestsRequest struct {

	// Exports the list of access requests for given date range in text format
	ExportAccessRequestsDetails `contributesTo:"body"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// A generic Id query param used to filter lockbox, access request and approval template by Id.
	Id *string `mandatory:"false" contributesTo:"query" name:"id"`

	// A filter to return only resources that match the entire display name given.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// A filter to return only resources their lifecycleState matches the given lifecycleState.
	LifecycleState AccessRequestLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// The name of the lockbox partner.
	LockboxPartner ExportAccessRequestsLockboxPartnerEnum `mandatory:"false" contributesTo:"query" name:"lockboxPartner" omitEmpty:"true"`

	// The ID of the partner.
	PartnerId *string `mandatory:"false" contributesTo:"query" name:"partnerId"`

	// The unique identifier (OCID) of the requestor in which to list resources.
	RequestorId *string `mandatory:"false" contributesTo:"query" name:"requestorId"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// A token representing the position at which to start retrieving results. This must come from the `opc-next-page` header field of a previous response.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'ASC' or 'DESC'.
	SortOrder ExportAccessRequestsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for timeCreated is descending. Default order for displayName is ascending.
	SortBy ExportAccessRequestsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ExportAccessRequestsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ExportAccessRequestsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ExportAccessRequestsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ExportAccessRequestsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ExportAccessRequestsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingAccessRequestLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetAccessRequestLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingExportAccessRequestsLockboxPartnerEnum(string(request.LockboxPartner)); !ok && request.LockboxPartner != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LockboxPartner: %s. Supported values are: %s.", request.LockboxPartner, strings.Join(GetExportAccessRequestsLockboxPartnerEnumStringValues(), ",")))
	}
	if _, ok := GetMappingExportAccessRequestsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetExportAccessRequestsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingExportAccessRequestsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetExportAccessRequestsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ExportAccessRequestsResponse wrapper for the ExportAccessRequests operation
type ExportAccessRequestsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of io.ReadCloser instances
	Content io.ReadCloser `presentIn:"body" encoding:"binary"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For optimistic concurrency control. See `if-match`.
	Etag *string `presentIn:"header" name:"etag"`
}

func (response ExportAccessRequestsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ExportAccessRequestsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ExportAccessRequestsLockboxPartnerEnum Enum with underlying type: string
type ExportAccessRequestsLockboxPartnerEnum string

// Set of constants representing the allowable values for ExportAccessRequestsLockboxPartnerEnum
const (
	ExportAccessRequestsLockboxPartnerFaaas  ExportAccessRequestsLockboxPartnerEnum = "FAAAS"
	ExportAccessRequestsLockboxPartnerCanary ExportAccessRequestsLockboxPartnerEnum = "CANARY"
)

var mappingExportAccessRequestsLockboxPartnerEnum = map[string]ExportAccessRequestsLockboxPartnerEnum{
	"FAAAS":  ExportAccessRequestsLockboxPartnerFaaas,
	"CANARY": ExportAccessRequestsLockboxPartnerCanary,
}

var mappingExportAccessRequestsLockboxPartnerEnumLowerCase = map[string]ExportAccessRequestsLockboxPartnerEnum{
	"faaas":  ExportAccessRequestsLockboxPartnerFaaas,
	"canary": ExportAccessRequestsLockboxPartnerCanary,
}

// GetExportAccessRequestsLockboxPartnerEnumValues Enumerates the set of values for ExportAccessRequestsLockboxPartnerEnum
func GetExportAccessRequestsLockboxPartnerEnumValues() []ExportAccessRequestsLockboxPartnerEnum {
	values := make([]ExportAccessRequestsLockboxPartnerEnum, 0)
	for _, v := range mappingExportAccessRequestsLockboxPartnerEnum {
		values = append(values, v)
	}
	return values
}

// GetExportAccessRequestsLockboxPartnerEnumStringValues Enumerates the set of values in String for ExportAccessRequestsLockboxPartnerEnum
func GetExportAccessRequestsLockboxPartnerEnumStringValues() []string {
	return []string{
		"FAAAS",
		"CANARY",
	}
}

// GetMappingExportAccessRequestsLockboxPartnerEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingExportAccessRequestsLockboxPartnerEnum(val string) (ExportAccessRequestsLockboxPartnerEnum, bool) {
	enum, ok := mappingExportAccessRequestsLockboxPartnerEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ExportAccessRequestsSortOrderEnum Enum with underlying type: string
type ExportAccessRequestsSortOrderEnum string

// Set of constants representing the allowable values for ExportAccessRequestsSortOrderEnum
const (
	ExportAccessRequestsSortOrderAsc  ExportAccessRequestsSortOrderEnum = "ASC"
	ExportAccessRequestsSortOrderDesc ExportAccessRequestsSortOrderEnum = "DESC"
)

var mappingExportAccessRequestsSortOrderEnum = map[string]ExportAccessRequestsSortOrderEnum{
	"ASC":  ExportAccessRequestsSortOrderAsc,
	"DESC": ExportAccessRequestsSortOrderDesc,
}

var mappingExportAccessRequestsSortOrderEnumLowerCase = map[string]ExportAccessRequestsSortOrderEnum{
	"asc":  ExportAccessRequestsSortOrderAsc,
	"desc": ExportAccessRequestsSortOrderDesc,
}

// GetExportAccessRequestsSortOrderEnumValues Enumerates the set of values for ExportAccessRequestsSortOrderEnum
func GetExportAccessRequestsSortOrderEnumValues() []ExportAccessRequestsSortOrderEnum {
	values := make([]ExportAccessRequestsSortOrderEnum, 0)
	for _, v := range mappingExportAccessRequestsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetExportAccessRequestsSortOrderEnumStringValues Enumerates the set of values in String for ExportAccessRequestsSortOrderEnum
func GetExportAccessRequestsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingExportAccessRequestsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingExportAccessRequestsSortOrderEnum(val string) (ExportAccessRequestsSortOrderEnum, bool) {
	enum, ok := mappingExportAccessRequestsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ExportAccessRequestsSortByEnum Enum with underlying type: string
type ExportAccessRequestsSortByEnum string

// Set of constants representing the allowable values for ExportAccessRequestsSortByEnum
const (
	ExportAccessRequestsSortByTimecreated ExportAccessRequestsSortByEnum = "timeCreated"
	ExportAccessRequestsSortByDisplayname ExportAccessRequestsSortByEnum = "displayName"
	ExportAccessRequestsSortById          ExportAccessRequestsSortByEnum = "id"
)

var mappingExportAccessRequestsSortByEnum = map[string]ExportAccessRequestsSortByEnum{
	"timeCreated": ExportAccessRequestsSortByTimecreated,
	"displayName": ExportAccessRequestsSortByDisplayname,
	"id":          ExportAccessRequestsSortById,
}

var mappingExportAccessRequestsSortByEnumLowerCase = map[string]ExportAccessRequestsSortByEnum{
	"timecreated": ExportAccessRequestsSortByTimecreated,
	"displayname": ExportAccessRequestsSortByDisplayname,
	"id":          ExportAccessRequestsSortById,
}

// GetExportAccessRequestsSortByEnumValues Enumerates the set of values for ExportAccessRequestsSortByEnum
func GetExportAccessRequestsSortByEnumValues() []ExportAccessRequestsSortByEnum {
	values := make([]ExportAccessRequestsSortByEnum, 0)
	for _, v := range mappingExportAccessRequestsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetExportAccessRequestsSortByEnumStringValues Enumerates the set of values in String for ExportAccessRequestsSortByEnum
func GetExportAccessRequestsSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
		"id",
	}
}

// GetMappingExportAccessRequestsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingExportAccessRequestsSortByEnum(val string) (ExportAccessRequestsSortByEnum, bool) {
	enum, ok := mappingExportAccessRequestsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
