// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package marketplacepublisher

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListAdminArtifactsRequest wrapper for the ListAdminArtifacts operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/marketplacepublisher/ListAdminArtifacts.go.html to see an example of how to use ListAdminArtifactsRequest.
type ListAdminArtifactsRequest struct {

	// The unique identifier for the compartment.
	CompartmentId *string `mandatory:"false" contributesTo:"query" name:"compartmentId"`

	// Unique publisher identifier.
	PublisherId *string `mandatory:"false" contributesTo:"query" name:"publisherId"`

	// A filter to return only artifacts with specific type.
	ArtifactType ListAdminArtifactsArtifactTypeEnum `mandatory:"false" contributesTo:"query" name:"artifactType" omitEmpty:"true"`

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
	SortOrder ListAdminArtifactsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort admin artifacts by. The default order for timeCreated and timeUpdated is descending.
	// The default order for displayName is ascending.
	SortBy ListAdminArtifactsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListAdminArtifactsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListAdminArtifactsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListAdminArtifactsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListAdminArtifactsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListAdminArtifactsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListAdminArtifactsArtifactTypeEnum(string(request.ArtifactType)); !ok && request.ArtifactType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ArtifactType: %s. Supported values are: %s.", request.ArtifactType, strings.Join(GetListAdminArtifactsArtifactTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingArtifactLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetArtifactLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingArtifactStatusEnum(string(request.Status)); !ok && request.Status != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", request.Status, strings.Join(GetArtifactStatusEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListAdminArtifactsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListAdminArtifactsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListAdminArtifactsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListAdminArtifactsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListAdminArtifactsResponse wrapper for the ListAdminArtifacts operation
type ListAdminArtifactsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of AdminArtifactCollection instances
	AdminArtifactCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListAdminArtifactsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListAdminArtifactsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListAdminArtifactsArtifactTypeEnum Enum with underlying type: string
type ListAdminArtifactsArtifactTypeEnum string

// Set of constants representing the allowable values for ListAdminArtifactsArtifactTypeEnum
const (
	ListAdminArtifactsArtifactTypeContainerImage ListAdminArtifactsArtifactTypeEnum = "CONTAINER_IMAGE"
	ListAdminArtifactsArtifactTypeHelmChart      ListAdminArtifactsArtifactTypeEnum = "HELM_CHART"
	ListAdminArtifactsArtifactTypeMachineImage   ListAdminArtifactsArtifactTypeEnum = "MACHINE_IMAGE"
	ListAdminArtifactsArtifactTypeStack          ListAdminArtifactsArtifactTypeEnum = "STACK"
)

var mappingListAdminArtifactsArtifactTypeEnum = map[string]ListAdminArtifactsArtifactTypeEnum{
	"CONTAINER_IMAGE": ListAdminArtifactsArtifactTypeContainerImage,
	"HELM_CHART":      ListAdminArtifactsArtifactTypeHelmChart,
	"MACHINE_IMAGE":   ListAdminArtifactsArtifactTypeMachineImage,
	"STACK":           ListAdminArtifactsArtifactTypeStack,
}

var mappingListAdminArtifactsArtifactTypeEnumLowerCase = map[string]ListAdminArtifactsArtifactTypeEnum{
	"container_image": ListAdminArtifactsArtifactTypeContainerImage,
	"helm_chart":      ListAdminArtifactsArtifactTypeHelmChart,
	"machine_image":   ListAdminArtifactsArtifactTypeMachineImage,
	"stack":           ListAdminArtifactsArtifactTypeStack,
}

// GetListAdminArtifactsArtifactTypeEnumValues Enumerates the set of values for ListAdminArtifactsArtifactTypeEnum
func GetListAdminArtifactsArtifactTypeEnumValues() []ListAdminArtifactsArtifactTypeEnum {
	values := make([]ListAdminArtifactsArtifactTypeEnum, 0)
	for _, v := range mappingListAdminArtifactsArtifactTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetListAdminArtifactsArtifactTypeEnumStringValues Enumerates the set of values in String for ListAdminArtifactsArtifactTypeEnum
func GetListAdminArtifactsArtifactTypeEnumStringValues() []string {
	return []string{
		"CONTAINER_IMAGE",
		"HELM_CHART",
		"MACHINE_IMAGE",
		"STACK",
	}
}

// GetMappingListAdminArtifactsArtifactTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListAdminArtifactsArtifactTypeEnum(val string) (ListAdminArtifactsArtifactTypeEnum, bool) {
	enum, ok := mappingListAdminArtifactsArtifactTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListAdminArtifactsSortOrderEnum Enum with underlying type: string
type ListAdminArtifactsSortOrderEnum string

// Set of constants representing the allowable values for ListAdminArtifactsSortOrderEnum
const (
	ListAdminArtifactsSortOrderAsc  ListAdminArtifactsSortOrderEnum = "ASC"
	ListAdminArtifactsSortOrderDesc ListAdminArtifactsSortOrderEnum = "DESC"
)

var mappingListAdminArtifactsSortOrderEnum = map[string]ListAdminArtifactsSortOrderEnum{
	"ASC":  ListAdminArtifactsSortOrderAsc,
	"DESC": ListAdminArtifactsSortOrderDesc,
}

var mappingListAdminArtifactsSortOrderEnumLowerCase = map[string]ListAdminArtifactsSortOrderEnum{
	"asc":  ListAdminArtifactsSortOrderAsc,
	"desc": ListAdminArtifactsSortOrderDesc,
}

// GetListAdminArtifactsSortOrderEnumValues Enumerates the set of values for ListAdminArtifactsSortOrderEnum
func GetListAdminArtifactsSortOrderEnumValues() []ListAdminArtifactsSortOrderEnum {
	values := make([]ListAdminArtifactsSortOrderEnum, 0)
	for _, v := range mappingListAdminArtifactsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListAdminArtifactsSortOrderEnumStringValues Enumerates the set of values in String for ListAdminArtifactsSortOrderEnum
func GetListAdminArtifactsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListAdminArtifactsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListAdminArtifactsSortOrderEnum(val string) (ListAdminArtifactsSortOrderEnum, bool) {
	enum, ok := mappingListAdminArtifactsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListAdminArtifactsSortByEnum Enum with underlying type: string
type ListAdminArtifactsSortByEnum string

// Set of constants representing the allowable values for ListAdminArtifactsSortByEnum
const (
	ListAdminArtifactsSortByTimecreated ListAdminArtifactsSortByEnum = "timeCreated"
	ListAdminArtifactsSortByTimeupdated ListAdminArtifactsSortByEnum = "timeUpdated"
	ListAdminArtifactsSortByDisplayname ListAdminArtifactsSortByEnum = "displayName"
)

var mappingListAdminArtifactsSortByEnum = map[string]ListAdminArtifactsSortByEnum{
	"timeCreated": ListAdminArtifactsSortByTimecreated,
	"timeUpdated": ListAdminArtifactsSortByTimeupdated,
	"displayName": ListAdminArtifactsSortByDisplayname,
}

var mappingListAdminArtifactsSortByEnumLowerCase = map[string]ListAdminArtifactsSortByEnum{
	"timecreated": ListAdminArtifactsSortByTimecreated,
	"timeupdated": ListAdminArtifactsSortByTimeupdated,
	"displayname": ListAdminArtifactsSortByDisplayname,
}

// GetListAdminArtifactsSortByEnumValues Enumerates the set of values for ListAdminArtifactsSortByEnum
func GetListAdminArtifactsSortByEnumValues() []ListAdminArtifactsSortByEnum {
	values := make([]ListAdminArtifactsSortByEnum, 0)
	for _, v := range mappingListAdminArtifactsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListAdminArtifactsSortByEnumStringValues Enumerates the set of values in String for ListAdminArtifactsSortByEnum
func GetListAdminArtifactsSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"timeUpdated",
		"displayName",
	}
}

// GetMappingListAdminArtifactsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListAdminArtifactsSortByEnum(val string) (ListAdminArtifactsSortByEnum, bool) {
	enum, ok := mappingListAdminArtifactsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
