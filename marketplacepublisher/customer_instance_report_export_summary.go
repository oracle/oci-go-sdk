// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// MarketplacePublisherService API
//
// Use the Marketplace Publisher API to manage the publishing of applications in Oracle Cloud Infrastructure Marketplace.
//

package marketplacepublisher

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CustomerInstanceReportExportSummary Summary model for a Customer Instance Report export.
type CustomerInstanceReportExportSummary struct {

	// Customer Instance Report export identifier.
	Id *string `mandatory:"true" json:"id"`

	// The OCID of the compartment for the export.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// Customer Instance Report export lifecycle state.
	LifecycleState CustomerInstanceReportExportLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The date and time that the export was created, expressed in RFC 3339 (https://tools.ietf.org/html/rfc3339)
	// timestamp format.
	// Example: `2016-08-25T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The work request identifier associated with the export generation.
	WorkRequestId *string `mandatory:"false" json:"workRequestId"`

	// The generated CSV file name.
	FileName *string `mandatory:"false" json:"fileName"`

	// The sort order for the generated export.
	SortOrder CustomerInstanceReportExportSummarySortOrderEnum `mandatory:"false" json:"sortOrder,omitempty"`

	// The field to sort by.
	SortBy CustomerInstanceReportExportSummarySortByEnum `mandatory:"false" json:"sortBy,omitempty"`

	// A filter to return only report records that match the listing name.
	Name *string `mandatory:"false" json:"name"`

	// A filter to return only report records that match the listing OCID.
	ListingId *string `mandatory:"false" json:"listingId"`

	// A filter to return only report records that match the instance status.
	Status CustomerInstanceReportExportSummaryStatusEnum `mandatory:"false" json:"status,omitempty"`

	// A filter to return only report records that match the instance shape.
	Shape *string `mandatory:"false" json:"shape"`

	// A filter to return only report records that match the instance region.
	Region *string `mandatory:"false" json:"region"`

	// A filter to return only report records that match the instance realm.
	Realm *string `mandatory:"false" json:"realm"`

	// A filter to return only report records that match the tenant administrator email domain.
	TenantAdminDomain *string `mandatory:"false" json:"tenantAdminDomain"`

	// A filter to return only report records that match the package version.
	PackageVersion *string `mandatory:"false" json:"packageVersion"`

	// A filter to return only report records that match the instance OCID.
	InstanceId *string `mandatory:"false" json:"instanceId"`

	// The inclusive earliest instance creation time, in RFC 3339 format.
	TimeInstanceCreationFromDate *common.SDKTime `mandatory:"false" json:"timeInstanceCreationFromDate"`

	// The inclusive latest instance creation time, in RFC 3339 format.
	TimeInstanceCreationToDate *common.SDKTime `mandatory:"false" json:"timeInstanceCreationToDate"`

	// The inclusive earliest instance termination time, in RFC 3339 format.
	TimeInstanceTerminationFromDate *common.SDKTime `mandatory:"false" json:"timeInstanceTerminationFromDate"`

	// The inclusive latest instance termination time, in RFC 3339 format.
	TimeInstanceTerminationToDate *common.SDKTime `mandatory:"false" json:"timeInstanceTerminationToDate"`

	// The date and time that the export expires, expressed in RFC 3339 (https://tools.ietf.org/html/rfc3339)
	// timestamp format.
	// Example: `2016-08-25T21:10:29.600Z`
	TimeExpires *common.SDKTime `mandatory:"false" json:"timeExpires"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// System tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
}

func (m CustomerInstanceReportExportSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CustomerInstanceReportExportSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingCustomerInstanceReportExportLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetCustomerInstanceReportExportLifecycleStateEnumStringValues(), ",")))
	}

	if _, ok := GetMappingCustomerInstanceReportExportSummarySortOrderEnum(string(m.SortOrder)); !ok && m.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", m.SortOrder, strings.Join(GetCustomerInstanceReportExportSummarySortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingCustomerInstanceReportExportSummarySortByEnum(string(m.SortBy)); !ok && m.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", m.SortBy, strings.Join(GetCustomerInstanceReportExportSummarySortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingCustomerInstanceReportExportSummaryStatusEnum(string(m.Status)); !ok && m.Status != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", m.Status, strings.Join(GetCustomerInstanceReportExportSummaryStatusEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// CustomerInstanceReportExportSummarySortOrderEnum Enum with underlying type: string
type CustomerInstanceReportExportSummarySortOrderEnum string

// Set of constants representing the allowable values for CustomerInstanceReportExportSummarySortOrderEnum
const (
	CustomerInstanceReportExportSummarySortOrderAsc  CustomerInstanceReportExportSummarySortOrderEnum = "ASC"
	CustomerInstanceReportExportSummarySortOrderDesc CustomerInstanceReportExportSummarySortOrderEnum = "DESC"
)

var mappingCustomerInstanceReportExportSummarySortOrderEnum = map[string]CustomerInstanceReportExportSummarySortOrderEnum{
	"ASC":  CustomerInstanceReportExportSummarySortOrderAsc,
	"DESC": CustomerInstanceReportExportSummarySortOrderDesc,
}

var mappingCustomerInstanceReportExportSummarySortOrderEnumLowerCase = map[string]CustomerInstanceReportExportSummarySortOrderEnum{
	"asc":  CustomerInstanceReportExportSummarySortOrderAsc,
	"desc": CustomerInstanceReportExportSummarySortOrderDesc,
}

// GetCustomerInstanceReportExportSummarySortOrderEnumValues Enumerates the set of values for CustomerInstanceReportExportSummarySortOrderEnum
func GetCustomerInstanceReportExportSummarySortOrderEnumValues() []CustomerInstanceReportExportSummarySortOrderEnum {
	values := make([]CustomerInstanceReportExportSummarySortOrderEnum, 0)
	for _, v := range mappingCustomerInstanceReportExportSummarySortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetCustomerInstanceReportExportSummarySortOrderEnumStringValues Enumerates the set of values in String for CustomerInstanceReportExportSummarySortOrderEnum
func GetCustomerInstanceReportExportSummarySortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingCustomerInstanceReportExportSummarySortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCustomerInstanceReportExportSummarySortOrderEnum(val string) (CustomerInstanceReportExportSummarySortOrderEnum, bool) {
	enum, ok := mappingCustomerInstanceReportExportSummarySortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// CustomerInstanceReportExportSummarySortByEnum Enum with underlying type: string
type CustomerInstanceReportExportSummarySortByEnum string

// Set of constants representing the allowable values for CustomerInstanceReportExportSummarySortByEnum
const (
	CustomerInstanceReportExportSummarySortByInstanceId CustomerInstanceReportExportSummarySortByEnum = "INSTANCE_ID"
)

var mappingCustomerInstanceReportExportSummarySortByEnum = map[string]CustomerInstanceReportExportSummarySortByEnum{
	"INSTANCE_ID": CustomerInstanceReportExportSummarySortByInstanceId,
}

var mappingCustomerInstanceReportExportSummarySortByEnumLowerCase = map[string]CustomerInstanceReportExportSummarySortByEnum{
	"instance_id": CustomerInstanceReportExportSummarySortByInstanceId,
}

// GetCustomerInstanceReportExportSummarySortByEnumValues Enumerates the set of values for CustomerInstanceReportExportSummarySortByEnum
func GetCustomerInstanceReportExportSummarySortByEnumValues() []CustomerInstanceReportExportSummarySortByEnum {
	values := make([]CustomerInstanceReportExportSummarySortByEnum, 0)
	for _, v := range mappingCustomerInstanceReportExportSummarySortByEnum {
		values = append(values, v)
	}
	return values
}

// GetCustomerInstanceReportExportSummarySortByEnumStringValues Enumerates the set of values in String for CustomerInstanceReportExportSummarySortByEnum
func GetCustomerInstanceReportExportSummarySortByEnumStringValues() []string {
	return []string{
		"INSTANCE_ID",
	}
}

// GetMappingCustomerInstanceReportExportSummarySortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCustomerInstanceReportExportSummarySortByEnum(val string) (CustomerInstanceReportExportSummarySortByEnum, bool) {
	enum, ok := mappingCustomerInstanceReportExportSummarySortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// CustomerInstanceReportExportSummaryStatusEnum Enum with underlying type: string
type CustomerInstanceReportExportSummaryStatusEnum string

// Set of constants representing the allowable values for CustomerInstanceReportExportSummaryStatusEnum
const (
	CustomerInstanceReportExportSummaryStatusProvisioning CustomerInstanceReportExportSummaryStatusEnum = "PROVISIONING"
	CustomerInstanceReportExportSummaryStatusRunning      CustomerInstanceReportExportSummaryStatusEnum = "RUNNING"
	CustomerInstanceReportExportSummaryStatusStopped      CustomerInstanceReportExportSummaryStatusEnum = "STOPPED"
	CustomerInstanceReportExportSummaryStatusTerminating  CustomerInstanceReportExportSummaryStatusEnum = "TERMINATING"
	CustomerInstanceReportExportSummaryStatusTerminated   CustomerInstanceReportExportSummaryStatusEnum = "TERMINATED"
	CustomerInstanceReportExportSummaryStatusDisabled     CustomerInstanceReportExportSummaryStatusEnum = "DISABLED"
	CustomerInstanceReportExportSummaryStatusStarting     CustomerInstanceReportExportSummaryStatusEnum = "STARTING"
	CustomerInstanceReportExportSummaryStatusStopping     CustomerInstanceReportExportSummaryStatusEnum = "STOPPING"
	CustomerInstanceReportExportSummaryStatusSnapshotting CustomerInstanceReportExportSummaryStatusEnum = "SNAPSHOTTING"
)

var mappingCustomerInstanceReportExportSummaryStatusEnum = map[string]CustomerInstanceReportExportSummaryStatusEnum{
	"PROVISIONING": CustomerInstanceReportExportSummaryStatusProvisioning,
	"RUNNING":      CustomerInstanceReportExportSummaryStatusRunning,
	"STOPPED":      CustomerInstanceReportExportSummaryStatusStopped,
	"TERMINATING":  CustomerInstanceReportExportSummaryStatusTerminating,
	"TERMINATED":   CustomerInstanceReportExportSummaryStatusTerminated,
	"DISABLED":     CustomerInstanceReportExportSummaryStatusDisabled,
	"STARTING":     CustomerInstanceReportExportSummaryStatusStarting,
	"STOPPING":     CustomerInstanceReportExportSummaryStatusStopping,
	"SNAPSHOTTING": CustomerInstanceReportExportSummaryStatusSnapshotting,
}

var mappingCustomerInstanceReportExportSummaryStatusEnumLowerCase = map[string]CustomerInstanceReportExportSummaryStatusEnum{
	"provisioning": CustomerInstanceReportExportSummaryStatusProvisioning,
	"running":      CustomerInstanceReportExportSummaryStatusRunning,
	"stopped":      CustomerInstanceReportExportSummaryStatusStopped,
	"terminating":  CustomerInstanceReportExportSummaryStatusTerminating,
	"terminated":   CustomerInstanceReportExportSummaryStatusTerminated,
	"disabled":     CustomerInstanceReportExportSummaryStatusDisabled,
	"starting":     CustomerInstanceReportExportSummaryStatusStarting,
	"stopping":     CustomerInstanceReportExportSummaryStatusStopping,
	"snapshotting": CustomerInstanceReportExportSummaryStatusSnapshotting,
}

// GetCustomerInstanceReportExportSummaryStatusEnumValues Enumerates the set of values for CustomerInstanceReportExportSummaryStatusEnum
func GetCustomerInstanceReportExportSummaryStatusEnumValues() []CustomerInstanceReportExportSummaryStatusEnum {
	values := make([]CustomerInstanceReportExportSummaryStatusEnum, 0)
	for _, v := range mappingCustomerInstanceReportExportSummaryStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetCustomerInstanceReportExportSummaryStatusEnumStringValues Enumerates the set of values in String for CustomerInstanceReportExportSummaryStatusEnum
func GetCustomerInstanceReportExportSummaryStatusEnumStringValues() []string {
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

// GetMappingCustomerInstanceReportExportSummaryStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCustomerInstanceReportExportSummaryStatusEnum(val string) (CustomerInstanceReportExportSummaryStatusEnum, bool) {
	enum, ok := mappingCustomerInstanceReportExportSummaryStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
