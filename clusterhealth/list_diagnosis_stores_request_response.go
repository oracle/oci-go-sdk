// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package clusterhealth

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListDiagnosisStoresRequest wrapper for the ListDiagnosisStores operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/clusterhealth/ListDiagnosisStores.go.html to see an example of how to use ListDiagnosisStoresRequest.
type ListDiagnosisStoresRequest struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment in which to list resources.
	CompartmentId *string `mandatory:"false" contributesTo:"query" name:"compartmentId"`

	// A filter to return only resources that match the given lifecycle state. The
	// state value is case-insensitive.
	LifecycleState DiagnosisStoreLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// A filter to return only resources that match the given display name exactly.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Diagnosis.
	DiagnosisStoreId *string `mandatory:"false" contributesTo:"query" name:"diagnosisStoreId"`

	// For list pagination. The maximum number of results per page, or items to return in a
	// paginated "List" call. For important details about how pagination works, see
	// List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The value of the opc-next-page response header from the previous
	// "List" call. For important details about how pagination works, see
	// List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either ascending (`ASC`) or descending (`DESC`).
	SortOrder ListDiagnosisStoresSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. You can provide only one sort order. Default order for `timeCreated`
	// is descending. Default order for `displayName` is ascending.
	SortBy ListDiagnosisStoresSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	// The only valid characters for request IDs are letters, numbers,
	// underscore, and dash.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListDiagnosisStoresRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListDiagnosisStoresRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListDiagnosisStoresRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListDiagnosisStoresRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListDiagnosisStoresRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingDiagnosisStoreLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetDiagnosisStoreLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListDiagnosisStoresSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListDiagnosisStoresSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListDiagnosisStoresSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListDiagnosisStoresSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListDiagnosisStoresResponse wrapper for the ListDiagnosisStores operation
type ListDiagnosisStoresResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of DiagnosisStoreCollection instances
	DiagnosisStoreCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response, additional pages of results remain. For
	// important details about how pagination works, see List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListDiagnosisStoresResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListDiagnosisStoresResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListDiagnosisStoresSortOrderEnum Enum with underlying type: string
type ListDiagnosisStoresSortOrderEnum string

// Set of constants representing the allowable values for ListDiagnosisStoresSortOrderEnum
const (
	ListDiagnosisStoresSortOrderAsc  ListDiagnosisStoresSortOrderEnum = "ASC"
	ListDiagnosisStoresSortOrderDesc ListDiagnosisStoresSortOrderEnum = "DESC"
)

var mappingListDiagnosisStoresSortOrderEnum = map[string]ListDiagnosisStoresSortOrderEnum{
	"ASC":  ListDiagnosisStoresSortOrderAsc,
	"DESC": ListDiagnosisStoresSortOrderDesc,
}

var mappingListDiagnosisStoresSortOrderEnumLowerCase = map[string]ListDiagnosisStoresSortOrderEnum{
	"asc":  ListDiagnosisStoresSortOrderAsc,
	"desc": ListDiagnosisStoresSortOrderDesc,
}

// GetListDiagnosisStoresSortOrderEnumValues Enumerates the set of values for ListDiagnosisStoresSortOrderEnum
func GetListDiagnosisStoresSortOrderEnumValues() []ListDiagnosisStoresSortOrderEnum {
	values := make([]ListDiagnosisStoresSortOrderEnum, 0)
	for _, v := range mappingListDiagnosisStoresSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListDiagnosisStoresSortOrderEnumStringValues Enumerates the set of values in String for ListDiagnosisStoresSortOrderEnum
func GetListDiagnosisStoresSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListDiagnosisStoresSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListDiagnosisStoresSortOrderEnum(val string) (ListDiagnosisStoresSortOrderEnum, bool) {
	enum, ok := mappingListDiagnosisStoresSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListDiagnosisStoresSortByEnum Enum with underlying type: string
type ListDiagnosisStoresSortByEnum string

// Set of constants representing the allowable values for ListDiagnosisStoresSortByEnum
const (
	ListDiagnosisStoresSortByTimecreated ListDiagnosisStoresSortByEnum = "timeCreated"
	ListDiagnosisStoresSortByDisplayname ListDiagnosisStoresSortByEnum = "displayName"
)

var mappingListDiagnosisStoresSortByEnum = map[string]ListDiagnosisStoresSortByEnum{
	"timeCreated": ListDiagnosisStoresSortByTimecreated,
	"displayName": ListDiagnosisStoresSortByDisplayname,
}

var mappingListDiagnosisStoresSortByEnumLowerCase = map[string]ListDiagnosisStoresSortByEnum{
	"timecreated": ListDiagnosisStoresSortByTimecreated,
	"displayname": ListDiagnosisStoresSortByDisplayname,
}

// GetListDiagnosisStoresSortByEnumValues Enumerates the set of values for ListDiagnosisStoresSortByEnum
func GetListDiagnosisStoresSortByEnumValues() []ListDiagnosisStoresSortByEnum {
	values := make([]ListDiagnosisStoresSortByEnum, 0)
	for _, v := range mappingListDiagnosisStoresSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListDiagnosisStoresSortByEnumStringValues Enumerates the set of values in String for ListDiagnosisStoresSortByEnum
func GetListDiagnosisStoresSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListDiagnosisStoresSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListDiagnosisStoresSortByEnum(val string) (ListDiagnosisStoresSortByEnum, bool) {
	enum, ok := mappingListDiagnosisStoresSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
