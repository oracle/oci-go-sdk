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

// ListCustomerInstanceReportRecordsRequest wrapper for the ListCustomerInstanceReportRecords operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/marketplacepublisher/ListCustomerInstanceReportRecords.go.html to see an example of how to use ListCustomerInstanceReportRecordsRequest.
type ListCustomerInstanceReportRecordsRequest struct {

	// The ID of the compartment in which to list resources.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// A token representing the position at which to start retrieving results. This must come from the `opc-next-page` header field of a previous response.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'ASC' or 'DESC'.
	SortOrder ListCustomerInstanceReportRecordsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for instance_id is ascending.
	SortBy ListCustomerInstanceReportRecordsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// A filter to return only resources that match the entire name given.
	Name *string `mandatory:"false" contributesTo:"query" name:"name"`

	// Listing OCID to query resource against.
	ListingId *string `mandatory:"false" contributesTo:"query" name:"listingId"`

	// A filter to return only report records that match the instance status.
	Status ListCustomerInstanceReportRecordsStatusEnum `mandatory:"false" contributesTo:"query" name:"status" omitEmpty:"true"`

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

func (request ListCustomerInstanceReportRecordsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListCustomerInstanceReportRecordsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListCustomerInstanceReportRecordsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListCustomerInstanceReportRecordsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListCustomerInstanceReportRecordsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListCustomerInstanceReportRecordsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListCustomerInstanceReportRecordsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListCustomerInstanceReportRecordsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListCustomerInstanceReportRecordsSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListCustomerInstanceReportRecordsStatusEnum(string(request.Status)); !ok && request.Status != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", request.Status, strings.Join(GetListCustomerInstanceReportRecordsStatusEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListCustomerInstanceReportRecordsResponse wrapper for the ListCustomerInstanceReportRecords operation
type ListCustomerInstanceReportRecordsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of CustomerInstanceReportRecordCollection instances
	CustomerInstanceReportRecordCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListCustomerInstanceReportRecordsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListCustomerInstanceReportRecordsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListCustomerInstanceReportRecordsSortOrderEnum Enum with underlying type: string
type ListCustomerInstanceReportRecordsSortOrderEnum string

// Set of constants representing the allowable values for ListCustomerInstanceReportRecordsSortOrderEnum
const (
	ListCustomerInstanceReportRecordsSortOrderAsc  ListCustomerInstanceReportRecordsSortOrderEnum = "ASC"
	ListCustomerInstanceReportRecordsSortOrderDesc ListCustomerInstanceReportRecordsSortOrderEnum = "DESC"
)

var mappingListCustomerInstanceReportRecordsSortOrderEnum = map[string]ListCustomerInstanceReportRecordsSortOrderEnum{
	"ASC":  ListCustomerInstanceReportRecordsSortOrderAsc,
	"DESC": ListCustomerInstanceReportRecordsSortOrderDesc,
}

var mappingListCustomerInstanceReportRecordsSortOrderEnumLowerCase = map[string]ListCustomerInstanceReportRecordsSortOrderEnum{
	"asc":  ListCustomerInstanceReportRecordsSortOrderAsc,
	"desc": ListCustomerInstanceReportRecordsSortOrderDesc,
}

// GetListCustomerInstanceReportRecordsSortOrderEnumValues Enumerates the set of values for ListCustomerInstanceReportRecordsSortOrderEnum
func GetListCustomerInstanceReportRecordsSortOrderEnumValues() []ListCustomerInstanceReportRecordsSortOrderEnum {
	values := make([]ListCustomerInstanceReportRecordsSortOrderEnum, 0)
	for _, v := range mappingListCustomerInstanceReportRecordsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListCustomerInstanceReportRecordsSortOrderEnumStringValues Enumerates the set of values in String for ListCustomerInstanceReportRecordsSortOrderEnum
func GetListCustomerInstanceReportRecordsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListCustomerInstanceReportRecordsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListCustomerInstanceReportRecordsSortOrderEnum(val string) (ListCustomerInstanceReportRecordsSortOrderEnum, bool) {
	enum, ok := mappingListCustomerInstanceReportRecordsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListCustomerInstanceReportRecordsSortByEnum Enum with underlying type: string
type ListCustomerInstanceReportRecordsSortByEnum string

// Set of constants representing the allowable values for ListCustomerInstanceReportRecordsSortByEnum
const (
	ListCustomerInstanceReportRecordsSortByInstanceId ListCustomerInstanceReportRecordsSortByEnum = "INSTANCE_ID"
)

var mappingListCustomerInstanceReportRecordsSortByEnum = map[string]ListCustomerInstanceReportRecordsSortByEnum{
	"INSTANCE_ID": ListCustomerInstanceReportRecordsSortByInstanceId,
}

var mappingListCustomerInstanceReportRecordsSortByEnumLowerCase = map[string]ListCustomerInstanceReportRecordsSortByEnum{
	"instance_id": ListCustomerInstanceReportRecordsSortByInstanceId,
}

// GetListCustomerInstanceReportRecordsSortByEnumValues Enumerates the set of values for ListCustomerInstanceReportRecordsSortByEnum
func GetListCustomerInstanceReportRecordsSortByEnumValues() []ListCustomerInstanceReportRecordsSortByEnum {
	values := make([]ListCustomerInstanceReportRecordsSortByEnum, 0)
	for _, v := range mappingListCustomerInstanceReportRecordsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListCustomerInstanceReportRecordsSortByEnumStringValues Enumerates the set of values in String for ListCustomerInstanceReportRecordsSortByEnum
func GetListCustomerInstanceReportRecordsSortByEnumStringValues() []string {
	return []string{
		"INSTANCE_ID",
	}
}

// GetMappingListCustomerInstanceReportRecordsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListCustomerInstanceReportRecordsSortByEnum(val string) (ListCustomerInstanceReportRecordsSortByEnum, bool) {
	enum, ok := mappingListCustomerInstanceReportRecordsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListCustomerInstanceReportRecordsStatusEnum Enum with underlying type: string
type ListCustomerInstanceReportRecordsStatusEnum string

// Set of constants representing the allowable values for ListCustomerInstanceReportRecordsStatusEnum
const (
	ListCustomerInstanceReportRecordsStatusProvisioning ListCustomerInstanceReportRecordsStatusEnum = "PROVISIONING"
	ListCustomerInstanceReportRecordsStatusRunning      ListCustomerInstanceReportRecordsStatusEnum = "RUNNING"
	ListCustomerInstanceReportRecordsStatusStopped      ListCustomerInstanceReportRecordsStatusEnum = "STOPPED"
	ListCustomerInstanceReportRecordsStatusTerminating  ListCustomerInstanceReportRecordsStatusEnum = "TERMINATING"
	ListCustomerInstanceReportRecordsStatusTerminated   ListCustomerInstanceReportRecordsStatusEnum = "TERMINATED"
	ListCustomerInstanceReportRecordsStatusDisabled     ListCustomerInstanceReportRecordsStatusEnum = "DISABLED"
	ListCustomerInstanceReportRecordsStatusStarting     ListCustomerInstanceReportRecordsStatusEnum = "STARTING"
	ListCustomerInstanceReportRecordsStatusStopping     ListCustomerInstanceReportRecordsStatusEnum = "STOPPING"
	ListCustomerInstanceReportRecordsStatusSnapshotting ListCustomerInstanceReportRecordsStatusEnum = "SNAPSHOTTING"
)

var mappingListCustomerInstanceReportRecordsStatusEnum = map[string]ListCustomerInstanceReportRecordsStatusEnum{
	"PROVISIONING": ListCustomerInstanceReportRecordsStatusProvisioning,
	"RUNNING":      ListCustomerInstanceReportRecordsStatusRunning,
	"STOPPED":      ListCustomerInstanceReportRecordsStatusStopped,
	"TERMINATING":  ListCustomerInstanceReportRecordsStatusTerminating,
	"TERMINATED":   ListCustomerInstanceReportRecordsStatusTerminated,
	"DISABLED":     ListCustomerInstanceReportRecordsStatusDisabled,
	"STARTING":     ListCustomerInstanceReportRecordsStatusStarting,
	"STOPPING":     ListCustomerInstanceReportRecordsStatusStopping,
	"SNAPSHOTTING": ListCustomerInstanceReportRecordsStatusSnapshotting,
}

var mappingListCustomerInstanceReportRecordsStatusEnumLowerCase = map[string]ListCustomerInstanceReportRecordsStatusEnum{
	"provisioning": ListCustomerInstanceReportRecordsStatusProvisioning,
	"running":      ListCustomerInstanceReportRecordsStatusRunning,
	"stopped":      ListCustomerInstanceReportRecordsStatusStopped,
	"terminating":  ListCustomerInstanceReportRecordsStatusTerminating,
	"terminated":   ListCustomerInstanceReportRecordsStatusTerminated,
	"disabled":     ListCustomerInstanceReportRecordsStatusDisabled,
	"starting":     ListCustomerInstanceReportRecordsStatusStarting,
	"stopping":     ListCustomerInstanceReportRecordsStatusStopping,
	"snapshotting": ListCustomerInstanceReportRecordsStatusSnapshotting,
}

// GetListCustomerInstanceReportRecordsStatusEnumValues Enumerates the set of values for ListCustomerInstanceReportRecordsStatusEnum
func GetListCustomerInstanceReportRecordsStatusEnumValues() []ListCustomerInstanceReportRecordsStatusEnum {
	values := make([]ListCustomerInstanceReportRecordsStatusEnum, 0)
	for _, v := range mappingListCustomerInstanceReportRecordsStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetListCustomerInstanceReportRecordsStatusEnumStringValues Enumerates the set of values in String for ListCustomerInstanceReportRecordsStatusEnum
func GetListCustomerInstanceReportRecordsStatusEnumStringValues() []string {
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

// GetMappingListCustomerInstanceReportRecordsStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListCustomerInstanceReportRecordsStatusEnum(val string) (ListCustomerInstanceReportRecordsStatusEnum, bool) {
	enum, ok := mappingListCustomerInstanceReportRecordsStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
