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

// ListCustomerInstanceReportExportsRequest wrapper for the ListCustomerInstanceReportExports operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/marketplacepublisher/ListCustomerInstanceReportExports.go.html to see an example of how to use ListCustomerInstanceReportExportsRequest.
type ListCustomerInstanceReportExportsRequest struct {

	// The ID of the compartment in which to list resources.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// A token representing the position at which to start retrieving results. This must come from the `opc-next-page` header field of a previous response.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'ASC' or 'DESC'.
	SortOrder ListCustomerInstanceReportExportsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for timeCreated is descending. Default order for fileName is ascending.
	SortBy ListCustomerInstanceReportExportsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// A filter to return only resources that match the entire name given.
	Name *string `mandatory:"false" contributesTo:"query" name:"name"`

	// Listing OCID to query resource against.
	ListingId *string `mandatory:"false" contributesTo:"query" name:"listingId"`

	// A filter to return only report records that match the instance status.
	Status ListCustomerInstanceReportExportsStatusEnum `mandatory:"false" contributesTo:"query" name:"status" omitEmpty:"true"`

	// A filter to return only report records that match the instance shape.
	Shape *string `mandatory:"false" contributesTo:"query" name:"shape"`

	// A filter to return only report records that match the instance region.
	Region *string `mandatory:"false" contributesTo:"query" name:"region"`

	// A filter to return only report records that match the instance realm.
	Realm *string `mandatory:"false" contributesTo:"query" name:"realm"`

	// A filter to return only report records that match the tenant administrator domain.
	TenantAdminDomain *string `mandatory:"false" contributesTo:"query" name:"tenantAdminDomain"`

	// A filter to return only report records that match the package version.
	PackageVersion *string `mandatory:"false" contributesTo:"query" name:"packageVersion"`

	// A filter to return only report records that match the instance OCID.
	InstanceOcid *string `mandatory:"false" contributesTo:"query" name:"instanceOcid"`

	// The inclusive earliest instance creation time, in RFC 3339 format.
	TimeInstanceCreationFromDate *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timeInstanceCreationFromDate"`

	// The inclusive latest instance creation time, in RFC 3339 format.
	TimeInstanceCreationToDate *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timeInstanceCreationToDate"`

	// The inclusive earliest instance termination time, in RFC 3339 format.
	TimeInstanceTerminationFromDate *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timeInstanceTerminationFromDate"`

	// The inclusive latest instance termination time, in RFC 3339 format.
	TimeInstanceTerminationToDate *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timeInstanceTerminationToDate"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListCustomerInstanceReportExportsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListCustomerInstanceReportExportsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListCustomerInstanceReportExportsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListCustomerInstanceReportExportsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListCustomerInstanceReportExportsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListCustomerInstanceReportExportsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListCustomerInstanceReportExportsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListCustomerInstanceReportExportsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListCustomerInstanceReportExportsSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListCustomerInstanceReportExportsStatusEnum(string(request.Status)); !ok && request.Status != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", request.Status, strings.Join(GetListCustomerInstanceReportExportsStatusEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListCustomerInstanceReportExportsResponse wrapper for the ListCustomerInstanceReportExports operation
type ListCustomerInstanceReportExportsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of CustomerInstanceReportExportCollection instances
	CustomerInstanceReportExportCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListCustomerInstanceReportExportsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListCustomerInstanceReportExportsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListCustomerInstanceReportExportsSortOrderEnum Enum with underlying type: string
type ListCustomerInstanceReportExportsSortOrderEnum string

// Set of constants representing the allowable values for ListCustomerInstanceReportExportsSortOrderEnum
const (
	ListCustomerInstanceReportExportsSortOrderAsc  ListCustomerInstanceReportExportsSortOrderEnum = "ASC"
	ListCustomerInstanceReportExportsSortOrderDesc ListCustomerInstanceReportExportsSortOrderEnum = "DESC"
)

var mappingListCustomerInstanceReportExportsSortOrderEnum = map[string]ListCustomerInstanceReportExportsSortOrderEnum{
	"ASC":  ListCustomerInstanceReportExportsSortOrderAsc,
	"DESC": ListCustomerInstanceReportExportsSortOrderDesc,
}

var mappingListCustomerInstanceReportExportsSortOrderEnumLowerCase = map[string]ListCustomerInstanceReportExportsSortOrderEnum{
	"asc":  ListCustomerInstanceReportExportsSortOrderAsc,
	"desc": ListCustomerInstanceReportExportsSortOrderDesc,
}

// GetListCustomerInstanceReportExportsSortOrderEnumValues Enumerates the set of values for ListCustomerInstanceReportExportsSortOrderEnum
func GetListCustomerInstanceReportExportsSortOrderEnumValues() []ListCustomerInstanceReportExportsSortOrderEnum {
	values := make([]ListCustomerInstanceReportExportsSortOrderEnum, 0)
	for _, v := range mappingListCustomerInstanceReportExportsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListCustomerInstanceReportExportsSortOrderEnumStringValues Enumerates the set of values in String for ListCustomerInstanceReportExportsSortOrderEnum
func GetListCustomerInstanceReportExportsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListCustomerInstanceReportExportsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListCustomerInstanceReportExportsSortOrderEnum(val string) (ListCustomerInstanceReportExportsSortOrderEnum, bool) {
	enum, ok := mappingListCustomerInstanceReportExportsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListCustomerInstanceReportExportsSortByEnum Enum with underlying type: string
type ListCustomerInstanceReportExportsSortByEnum string

// Set of constants representing the allowable values for ListCustomerInstanceReportExportsSortByEnum
const (
	ListCustomerInstanceReportExportsSortByTimecreated ListCustomerInstanceReportExportsSortByEnum = "timeCreated"
	ListCustomerInstanceReportExportsSortByFilename    ListCustomerInstanceReportExportsSortByEnum = "fileName"
)

var mappingListCustomerInstanceReportExportsSortByEnum = map[string]ListCustomerInstanceReportExportsSortByEnum{
	"timeCreated": ListCustomerInstanceReportExportsSortByTimecreated,
	"fileName":    ListCustomerInstanceReportExportsSortByFilename,
}

var mappingListCustomerInstanceReportExportsSortByEnumLowerCase = map[string]ListCustomerInstanceReportExportsSortByEnum{
	"timecreated": ListCustomerInstanceReportExportsSortByTimecreated,
	"filename":    ListCustomerInstanceReportExportsSortByFilename,
}

// GetListCustomerInstanceReportExportsSortByEnumValues Enumerates the set of values for ListCustomerInstanceReportExportsSortByEnum
func GetListCustomerInstanceReportExportsSortByEnumValues() []ListCustomerInstanceReportExportsSortByEnum {
	values := make([]ListCustomerInstanceReportExportsSortByEnum, 0)
	for _, v := range mappingListCustomerInstanceReportExportsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListCustomerInstanceReportExportsSortByEnumStringValues Enumerates the set of values in String for ListCustomerInstanceReportExportsSortByEnum
func GetListCustomerInstanceReportExportsSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"fileName",
	}
}

// GetMappingListCustomerInstanceReportExportsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListCustomerInstanceReportExportsSortByEnum(val string) (ListCustomerInstanceReportExportsSortByEnum, bool) {
	enum, ok := mappingListCustomerInstanceReportExportsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListCustomerInstanceReportExportsStatusEnum Enum with underlying type: string
type ListCustomerInstanceReportExportsStatusEnum string

// Set of constants representing the allowable values for ListCustomerInstanceReportExportsStatusEnum
const (
	ListCustomerInstanceReportExportsStatusProvisioning ListCustomerInstanceReportExportsStatusEnum = "PROVISIONING"
	ListCustomerInstanceReportExportsStatusRunning      ListCustomerInstanceReportExportsStatusEnum = "RUNNING"
	ListCustomerInstanceReportExportsStatusStopped      ListCustomerInstanceReportExportsStatusEnum = "STOPPED"
	ListCustomerInstanceReportExportsStatusTerminating  ListCustomerInstanceReportExportsStatusEnum = "TERMINATING"
	ListCustomerInstanceReportExportsStatusTerminated   ListCustomerInstanceReportExportsStatusEnum = "TERMINATED"
	ListCustomerInstanceReportExportsStatusDisabled     ListCustomerInstanceReportExportsStatusEnum = "DISABLED"
	ListCustomerInstanceReportExportsStatusStarting     ListCustomerInstanceReportExportsStatusEnum = "STARTING"
	ListCustomerInstanceReportExportsStatusStopping     ListCustomerInstanceReportExportsStatusEnum = "STOPPING"
	ListCustomerInstanceReportExportsStatusSnapshotting ListCustomerInstanceReportExportsStatusEnum = "SNAPSHOTTING"
)

var mappingListCustomerInstanceReportExportsStatusEnum = map[string]ListCustomerInstanceReportExportsStatusEnum{
	"PROVISIONING": ListCustomerInstanceReportExportsStatusProvisioning,
	"RUNNING":      ListCustomerInstanceReportExportsStatusRunning,
	"STOPPED":      ListCustomerInstanceReportExportsStatusStopped,
	"TERMINATING":  ListCustomerInstanceReportExportsStatusTerminating,
	"TERMINATED":   ListCustomerInstanceReportExportsStatusTerminated,
	"DISABLED":     ListCustomerInstanceReportExportsStatusDisabled,
	"STARTING":     ListCustomerInstanceReportExportsStatusStarting,
	"STOPPING":     ListCustomerInstanceReportExportsStatusStopping,
	"SNAPSHOTTING": ListCustomerInstanceReportExportsStatusSnapshotting,
}

var mappingListCustomerInstanceReportExportsStatusEnumLowerCase = map[string]ListCustomerInstanceReportExportsStatusEnum{
	"provisioning": ListCustomerInstanceReportExportsStatusProvisioning,
	"running":      ListCustomerInstanceReportExportsStatusRunning,
	"stopped":      ListCustomerInstanceReportExportsStatusStopped,
	"terminating":  ListCustomerInstanceReportExportsStatusTerminating,
	"terminated":   ListCustomerInstanceReportExportsStatusTerminated,
	"disabled":     ListCustomerInstanceReportExportsStatusDisabled,
	"starting":     ListCustomerInstanceReportExportsStatusStarting,
	"stopping":     ListCustomerInstanceReportExportsStatusStopping,
	"snapshotting": ListCustomerInstanceReportExportsStatusSnapshotting,
}

// GetListCustomerInstanceReportExportsStatusEnumValues Enumerates the set of values for ListCustomerInstanceReportExportsStatusEnum
func GetListCustomerInstanceReportExportsStatusEnumValues() []ListCustomerInstanceReportExportsStatusEnum {
	values := make([]ListCustomerInstanceReportExportsStatusEnum, 0)
	for _, v := range mappingListCustomerInstanceReportExportsStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetListCustomerInstanceReportExportsStatusEnumStringValues Enumerates the set of values in String for ListCustomerInstanceReportExportsStatusEnum
func GetListCustomerInstanceReportExportsStatusEnumStringValues() []string {
	return []string{
		"PROVISIONING",
		"RUNNING",
		"STOPPED",
		"TERMINATING",
		"TERMINATED",
		"DISABLED",
		"STARTING",
		"STOPPING",
		"SNAPSHOTTING",
	}
}

// GetMappingListCustomerInstanceReportExportsStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListCustomerInstanceReportExportsStatusEnum(val string) (ListCustomerInstanceReportExportsStatusEnum, bool) {
	enum, ok := mappingListCustomerInstanceReportExportsStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
