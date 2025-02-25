// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package marketplacepublisher

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListArtifactsRequest wrapper for the ListArtifacts operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/marketplacepublisher/ListArtifacts.go.html to see an example of how to use ListArtifactsRequest.
type ListArtifactsRequest struct {

	// The ID of the compartment in which to list resources.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// A filter to return only artifacts with their lifecycleState matches the given lifecycleState.
	LifecycleState ArtifactLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// A filter to return only resources that match the entire display name given.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// A filter to return only artifacts with specific status.
	Status ArtifactStatusEnum `mandatory:"false" contributesTo:"query" name:"status" omitEmpty:"true"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// A token representing the position at which to start retrieving results. This must come from the `opc-next-page` header field of a previous response.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'ASC' or 'DESC'.
	SortOrder ListArtifactsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for timeCreated is descending. Default order for displayName is ascending.
	SortBy ListArtifactsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListArtifactsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListArtifactsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListArtifactsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListArtifactsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListArtifactsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingArtifactLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetArtifactLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingArtifactStatusEnum(string(request.Status)); !ok && request.Status != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", request.Status, strings.Join(GetArtifactStatusEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListArtifactsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListArtifactsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListArtifactsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListArtifactsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListArtifactsResponse wrapper for the ListArtifacts operation
type ListArtifactsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of ArtifactCollection instances
	ArtifactCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListArtifactsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListArtifactsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListArtifactsSortOrderEnum Enum with underlying type: string
type ListArtifactsSortOrderEnum string

// Set of constants representing the allowable values for ListArtifactsSortOrderEnum
const (
	ListArtifactsSortOrderAsc  ListArtifactsSortOrderEnum = "ASC"
	ListArtifactsSortOrderDesc ListArtifactsSortOrderEnum = "DESC"
)

var mappingListArtifactsSortOrderEnum = map[string]ListArtifactsSortOrderEnum{
	"ASC":  ListArtifactsSortOrderAsc,
	"DESC": ListArtifactsSortOrderDesc,
}

var mappingListArtifactsSortOrderEnumLowerCase = map[string]ListArtifactsSortOrderEnum{
	"asc":  ListArtifactsSortOrderAsc,
	"desc": ListArtifactsSortOrderDesc,
}

// GetListArtifactsSortOrderEnumValues Enumerates the set of values for ListArtifactsSortOrderEnum
func GetListArtifactsSortOrderEnumValues() []ListArtifactsSortOrderEnum {
	values := make([]ListArtifactsSortOrderEnum, 0)
	for _, v := range mappingListArtifactsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListArtifactsSortOrderEnumStringValues Enumerates the set of values in String for ListArtifactsSortOrderEnum
func GetListArtifactsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListArtifactsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListArtifactsSortOrderEnum(val string) (ListArtifactsSortOrderEnum, bool) {
	enum, ok := mappingListArtifactsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListArtifactsSortByEnum Enum with underlying type: string
type ListArtifactsSortByEnum string

// Set of constants representing the allowable values for ListArtifactsSortByEnum
const (
	ListArtifactsSortByTimecreated ListArtifactsSortByEnum = "timeCreated"
	ListArtifactsSortByDisplayname ListArtifactsSortByEnum = "displayName"
)

var mappingListArtifactsSortByEnum = map[string]ListArtifactsSortByEnum{
	"timeCreated": ListArtifactsSortByTimecreated,
	"displayName": ListArtifactsSortByDisplayname,
}

var mappingListArtifactsSortByEnumLowerCase = map[string]ListArtifactsSortByEnum{
	"timecreated": ListArtifactsSortByTimecreated,
	"displayname": ListArtifactsSortByDisplayname,
}

// GetListArtifactsSortByEnumValues Enumerates the set of values for ListArtifactsSortByEnum
func GetListArtifactsSortByEnumValues() []ListArtifactsSortByEnum {
	values := make([]ListArtifactsSortByEnum, 0)
	for _, v := range mappingListArtifactsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListArtifactsSortByEnumStringValues Enumerates the set of values in String for ListArtifactsSortByEnum
func GetListArtifactsSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListArtifactsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListArtifactsSortByEnum(val string) (ListArtifactsSortByEnum, bool) {
	enum, ok := mappingListArtifactsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
