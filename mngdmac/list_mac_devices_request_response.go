// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package mngdmac

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListMacDevicesRequest wrapper for the ListMacDevices operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/mngdmac/ListMacDevices.go.html to see an example of how to use ListMacDevicesRequest.
type ListMacDevicesRequest struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the MacOrder.
	MacOrderId *string `mandatory:"true" contributesTo:"path" name:"macOrderId"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment in which to list resources.
	CompartmentId *string `mandatory:"false" contributesTo:"query" name:"compartmentId"`

	// The UUID of the MacDevice.
	Id *string `mandatory:"false" contributesTo:"query" name:"id"`

	// The serial number of the MacDevice.
	SerialNumber *string `mandatory:"false" contributesTo:"query" name:"serialNumber"`

	// A filter to return only resources that match the given lifecycle state. The
	// state value is case-insensitive.
	LifecycleState MacDeviceLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// For list pagination. The maximum number of results per page, or items to return in a
	// paginated "List" call. For important details about how pagination works, see
	// List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The value of the opc-next-page response header from the previous
	// "List" call. For important details about how pagination works, see
	// List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either ascending (`ASC`) or descending (`DESC`).
	SortOrder ListMacDevicesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. You can provide only one sort order. Default order for `timeCreated`
	// is descending. Default order for `SerialNumber` is ascending.
	SortBy ListMacDevicesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	// The only valid characters for request IDs are letters, numbers,
	// underscore, and dash.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListMacDevicesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListMacDevicesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListMacDevicesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListMacDevicesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListMacDevicesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingMacDeviceLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetMacDeviceLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListMacDevicesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListMacDevicesSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListMacDevicesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListMacDevicesSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListMacDevicesResponse wrapper for the ListMacDevices operation
type ListMacDevicesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of MacDeviceCollection instances
	MacDeviceCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response, additional pages of results remain. For
	// important details about how pagination works, see List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListMacDevicesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListMacDevicesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListMacDevicesSortOrderEnum Enum with underlying type: string
type ListMacDevicesSortOrderEnum string

// Set of constants representing the allowable values for ListMacDevicesSortOrderEnum
const (
	ListMacDevicesSortOrderAsc  ListMacDevicesSortOrderEnum = "ASC"
	ListMacDevicesSortOrderDesc ListMacDevicesSortOrderEnum = "DESC"
)

var mappingListMacDevicesSortOrderEnum = map[string]ListMacDevicesSortOrderEnum{
	"ASC":  ListMacDevicesSortOrderAsc,
	"DESC": ListMacDevicesSortOrderDesc,
}

var mappingListMacDevicesSortOrderEnumLowerCase = map[string]ListMacDevicesSortOrderEnum{
	"asc":  ListMacDevicesSortOrderAsc,
	"desc": ListMacDevicesSortOrderDesc,
}

// GetListMacDevicesSortOrderEnumValues Enumerates the set of values for ListMacDevicesSortOrderEnum
func GetListMacDevicesSortOrderEnumValues() []ListMacDevicesSortOrderEnum {
	values := make([]ListMacDevicesSortOrderEnum, 0)
	for _, v := range mappingListMacDevicesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListMacDevicesSortOrderEnumStringValues Enumerates the set of values in String for ListMacDevicesSortOrderEnum
func GetListMacDevicesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListMacDevicesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListMacDevicesSortOrderEnum(val string) (ListMacDevicesSortOrderEnum, bool) {
	enum, ok := mappingListMacDevicesSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListMacDevicesSortByEnum Enum with underlying type: string
type ListMacDevicesSortByEnum string

// Set of constants representing the allowable values for ListMacDevicesSortByEnum
const (
	ListMacDevicesSortByTimecreated  ListMacDevicesSortByEnum = "timeCreated"
	ListMacDevicesSortBySerialnumber ListMacDevicesSortByEnum = "SerialNumber"
)

var mappingListMacDevicesSortByEnum = map[string]ListMacDevicesSortByEnum{
	"timeCreated":  ListMacDevicesSortByTimecreated,
	"SerialNumber": ListMacDevicesSortBySerialnumber,
}

var mappingListMacDevicesSortByEnumLowerCase = map[string]ListMacDevicesSortByEnum{
	"timecreated":  ListMacDevicesSortByTimecreated,
	"serialnumber": ListMacDevicesSortBySerialnumber,
}

// GetListMacDevicesSortByEnumValues Enumerates the set of values for ListMacDevicesSortByEnum
func GetListMacDevicesSortByEnumValues() []ListMacDevicesSortByEnum {
	values := make([]ListMacDevicesSortByEnum, 0)
	for _, v := range mappingListMacDevicesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListMacDevicesSortByEnumStringValues Enumerates the set of values in String for ListMacDevicesSortByEnum
func GetListMacDevicesSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"SerialNumber",
	}
}

// GetMappingListMacDevicesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListMacDevicesSortByEnum(val string) (ListMacDevicesSortByEnum, bool) {
	enum, ok := mappingListMacDevicesSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
