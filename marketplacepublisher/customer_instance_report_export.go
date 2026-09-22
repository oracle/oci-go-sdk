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

// CustomerInstanceReportExport The model for a Customer Instance Report export.
type CustomerInstanceReportExport struct {

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
	SortOrder CustomerInstanceReportExportSortOrderEnum `mandatory:"false" json:"sortOrder,omitempty"`

	// The field to sort by.
	SortBy CustomerInstanceReportExportSortByEnum `mandatory:"false" json:"sortBy,omitempty"`

	// A filter to return only report records that match the listing name.
	Name *string `mandatory:"false" json:"name"`

	// A filter to return only report records that match the listing OCID.
	ListingId *string `mandatory:"false" json:"listingId"`

	// A filter to return only report records that match the instance status.
	Status CustomerInstanceReportExportStatusEnum `mandatory:"false" json:"status,omitempty"`

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

func (m CustomerInstanceReportExport) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CustomerInstanceReportExport) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingCustomerInstanceReportExportLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetCustomerInstanceReportExportLifecycleStateEnumStringValues(), ",")))
	}

	if _, ok := GetMappingCustomerInstanceReportExportSortOrderEnum(string(m.SortOrder)); !ok && m.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", m.SortOrder, strings.Join(GetCustomerInstanceReportExportSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingCustomerInstanceReportExportSortByEnum(string(m.SortBy)); !ok && m.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", m.SortBy, strings.Join(GetCustomerInstanceReportExportSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingCustomerInstanceReportExportStatusEnum(string(m.Status)); !ok && m.Status != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", m.Status, strings.Join(GetCustomerInstanceReportExportStatusEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// CustomerInstanceReportExportLifecycleStateEnum Enum with underlying type: string
type CustomerInstanceReportExportLifecycleStateEnum string

// Set of constants representing the allowable values for CustomerInstanceReportExportLifecycleStateEnum
const (
	CustomerInstanceReportExportLifecycleStateInProgress     CustomerInstanceReportExportLifecycleStateEnum = "IN_PROGRESS"
	CustomerInstanceReportExportLifecycleStateAccepted       CustomerInstanceReportExportLifecycleStateEnum = "ACCEPTED"
	CustomerInstanceReportExportLifecycleStateFailed         CustomerInstanceReportExportLifecycleStateEnum = "FAILED"
	CustomerInstanceReportExportLifecycleStateNeedsAttention CustomerInstanceReportExportLifecycleStateEnum = "NEEDS_ATTENTION"
	CustomerInstanceReportExportLifecycleStateSucceeded      CustomerInstanceReportExportLifecycleStateEnum = "SUCCEEDED"
	CustomerInstanceReportExportLifecycleStateCanceled       CustomerInstanceReportExportLifecycleStateEnum = "CANCELED"
	CustomerInstanceReportExportLifecycleStateWaiting        CustomerInstanceReportExportLifecycleStateEnum = "WAITING"
	CustomerInstanceReportExportLifecycleStateCanceling      CustomerInstanceReportExportLifecycleStateEnum = "CANCELING"
)

var mappingCustomerInstanceReportExportLifecycleStateEnum = map[string]CustomerInstanceReportExportLifecycleStateEnum{
	"IN_PROGRESS":     CustomerInstanceReportExportLifecycleStateInProgress,
	"ACCEPTED":        CustomerInstanceReportExportLifecycleStateAccepted,
	"FAILED":          CustomerInstanceReportExportLifecycleStateFailed,
	"NEEDS_ATTENTION": CustomerInstanceReportExportLifecycleStateNeedsAttention,
	"SUCCEEDED":       CustomerInstanceReportExportLifecycleStateSucceeded,
	"CANCELED":        CustomerInstanceReportExportLifecycleStateCanceled,
	"WAITING":         CustomerInstanceReportExportLifecycleStateWaiting,
	"CANCELING":       CustomerInstanceReportExportLifecycleStateCanceling,
}

var mappingCustomerInstanceReportExportLifecycleStateEnumLowerCase = map[string]CustomerInstanceReportExportLifecycleStateEnum{
	"in_progress":     CustomerInstanceReportExportLifecycleStateInProgress,
	"accepted":        CustomerInstanceReportExportLifecycleStateAccepted,
	"failed":          CustomerInstanceReportExportLifecycleStateFailed,
	"needs_attention": CustomerInstanceReportExportLifecycleStateNeedsAttention,
	"succeeded":       CustomerInstanceReportExportLifecycleStateSucceeded,
	"canceled":        CustomerInstanceReportExportLifecycleStateCanceled,
	"waiting":         CustomerInstanceReportExportLifecycleStateWaiting,
	"canceling":       CustomerInstanceReportExportLifecycleStateCanceling,
}

// GetCustomerInstanceReportExportLifecycleStateEnumValues Enumerates the set of values for CustomerInstanceReportExportLifecycleStateEnum
func GetCustomerInstanceReportExportLifecycleStateEnumValues() []CustomerInstanceReportExportLifecycleStateEnum {
	values := make([]CustomerInstanceReportExportLifecycleStateEnum, 0)
	for _, v := range mappingCustomerInstanceReportExportLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetCustomerInstanceReportExportLifecycleStateEnumStringValues Enumerates the set of values in String for CustomerInstanceReportExportLifecycleStateEnum
func GetCustomerInstanceReportExportLifecycleStateEnumStringValues() []string {
	return []string{
		"IN_PROGRESS",
		"ACCEPTED",
		"FAILED",
		"NEEDS_ATTENTION",
		"SUCCEEDED",
		"CANCELED",
		"WAITING",
		"CANCELING",
	}
}

// GetMappingCustomerInstanceReportExportLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCustomerInstanceReportExportLifecycleStateEnum(val string) (CustomerInstanceReportExportLifecycleStateEnum, bool) {
	enum, ok := mappingCustomerInstanceReportExportLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// CustomerInstanceReportExportSortOrderEnum Enum with underlying type: string
type CustomerInstanceReportExportSortOrderEnum string

// Set of constants representing the allowable values for CustomerInstanceReportExportSortOrderEnum
const (
	CustomerInstanceReportExportSortOrderAsc  CustomerInstanceReportExportSortOrderEnum = "ASC"
	CustomerInstanceReportExportSortOrderDesc CustomerInstanceReportExportSortOrderEnum = "DESC"
)

var mappingCustomerInstanceReportExportSortOrderEnum = map[string]CustomerInstanceReportExportSortOrderEnum{
	"ASC":  CustomerInstanceReportExportSortOrderAsc,
	"DESC": CustomerInstanceReportExportSortOrderDesc,
}

var mappingCustomerInstanceReportExportSortOrderEnumLowerCase = map[string]CustomerInstanceReportExportSortOrderEnum{
	"asc":  CustomerInstanceReportExportSortOrderAsc,
	"desc": CustomerInstanceReportExportSortOrderDesc,
}

// GetCustomerInstanceReportExportSortOrderEnumValues Enumerates the set of values for CustomerInstanceReportExportSortOrderEnum
func GetCustomerInstanceReportExportSortOrderEnumValues() []CustomerInstanceReportExportSortOrderEnum {
	values := make([]CustomerInstanceReportExportSortOrderEnum, 0)
	for _, v := range mappingCustomerInstanceReportExportSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetCustomerInstanceReportExportSortOrderEnumStringValues Enumerates the set of values in String for CustomerInstanceReportExportSortOrderEnum
func GetCustomerInstanceReportExportSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingCustomerInstanceReportExportSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCustomerInstanceReportExportSortOrderEnum(val string) (CustomerInstanceReportExportSortOrderEnum, bool) {
	enum, ok := mappingCustomerInstanceReportExportSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// CustomerInstanceReportExportSortByEnum Enum with underlying type: string
type CustomerInstanceReportExportSortByEnum string

// Set of constants representing the allowable values for CustomerInstanceReportExportSortByEnum
const (
	CustomerInstanceReportExportSortByInstanceId CustomerInstanceReportExportSortByEnum = "INSTANCE_ID"
)

var mappingCustomerInstanceReportExportSortByEnum = map[string]CustomerInstanceReportExportSortByEnum{
	"INSTANCE_ID": CustomerInstanceReportExportSortByInstanceId,
}

var mappingCustomerInstanceReportExportSortByEnumLowerCase = map[string]CustomerInstanceReportExportSortByEnum{
	"instance_id": CustomerInstanceReportExportSortByInstanceId,
}

// GetCustomerInstanceReportExportSortByEnumValues Enumerates the set of values for CustomerInstanceReportExportSortByEnum
func GetCustomerInstanceReportExportSortByEnumValues() []CustomerInstanceReportExportSortByEnum {
	values := make([]CustomerInstanceReportExportSortByEnum, 0)
	for _, v := range mappingCustomerInstanceReportExportSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetCustomerInstanceReportExportSortByEnumStringValues Enumerates the set of values in String for CustomerInstanceReportExportSortByEnum
func GetCustomerInstanceReportExportSortByEnumStringValues() []string {
	return []string{
		"INSTANCE_ID",
	}
}

// GetMappingCustomerInstanceReportExportSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCustomerInstanceReportExportSortByEnum(val string) (CustomerInstanceReportExportSortByEnum, bool) {
	enum, ok := mappingCustomerInstanceReportExportSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// CustomerInstanceReportExportStatusEnum Enum with underlying type: string
type CustomerInstanceReportExportStatusEnum string

// Set of constants representing the allowable values for CustomerInstanceReportExportStatusEnum
const (
	CustomerInstanceReportExportStatusProvisioning CustomerInstanceReportExportStatusEnum = "PROVISIONING"
	CustomerInstanceReportExportStatusRunning      CustomerInstanceReportExportStatusEnum = "RUNNING"
	CustomerInstanceReportExportStatusStopped      CustomerInstanceReportExportStatusEnum = "STOPPED"
	CustomerInstanceReportExportStatusTerminating  CustomerInstanceReportExportStatusEnum = "TERMINATING"
	CustomerInstanceReportExportStatusTerminated   CustomerInstanceReportExportStatusEnum = "TERMINATED"
	CustomerInstanceReportExportStatusDisabled     CustomerInstanceReportExportStatusEnum = "DISABLED"
	CustomerInstanceReportExportStatusStarting     CustomerInstanceReportExportStatusEnum = "STARTING"
	CustomerInstanceReportExportStatusStopping     CustomerInstanceReportExportStatusEnum = "STOPPING"
	CustomerInstanceReportExportStatusSnapshotting CustomerInstanceReportExportStatusEnum = "SNAPSHOTTING"
)

var mappingCustomerInstanceReportExportStatusEnum = map[string]CustomerInstanceReportExportStatusEnum{
	"PROVISIONING": CustomerInstanceReportExportStatusProvisioning,
	"RUNNING":      CustomerInstanceReportExportStatusRunning,
	"STOPPED":      CustomerInstanceReportExportStatusStopped,
	"TERMINATING":  CustomerInstanceReportExportStatusTerminating,
	"TERMINATED":   CustomerInstanceReportExportStatusTerminated,
	"DISABLED":     CustomerInstanceReportExportStatusDisabled,
	"STARTING":     CustomerInstanceReportExportStatusStarting,
	"STOPPING":     CustomerInstanceReportExportStatusStopping,
	"SNAPSHOTTING": CustomerInstanceReportExportStatusSnapshotting,
}

var mappingCustomerInstanceReportExportStatusEnumLowerCase = map[string]CustomerInstanceReportExportStatusEnum{
	"provisioning": CustomerInstanceReportExportStatusProvisioning,
	"running":      CustomerInstanceReportExportStatusRunning,
	"stopped":      CustomerInstanceReportExportStatusStopped,
	"terminating":  CustomerInstanceReportExportStatusTerminating,
	"terminated":   CustomerInstanceReportExportStatusTerminated,
	"disabled":     CustomerInstanceReportExportStatusDisabled,
	"starting":     CustomerInstanceReportExportStatusStarting,
	"stopping":     CustomerInstanceReportExportStatusStopping,
	"snapshotting": CustomerInstanceReportExportStatusSnapshotting,
}

// GetCustomerInstanceReportExportStatusEnumValues Enumerates the set of values for CustomerInstanceReportExportStatusEnum
func GetCustomerInstanceReportExportStatusEnumValues() []CustomerInstanceReportExportStatusEnum {
	values := make([]CustomerInstanceReportExportStatusEnum, 0)
	for _, v := range mappingCustomerInstanceReportExportStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetCustomerInstanceReportExportStatusEnumStringValues Enumerates the set of values in String for CustomerInstanceReportExportStatusEnum
func GetCustomerInstanceReportExportStatusEnumStringValues() []string {
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

// GetMappingCustomerInstanceReportExportStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCustomerInstanceReportExportStatusEnum(val string) (CustomerInstanceReportExportStatusEnum, bool) {
	enum, ok := mappingCustomerInstanceReportExportStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
