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

// GenerateCustomerInstanceReportExportDetails Details to generate a Customer Instance Report export.
type GenerateCustomerInstanceReportExportDetails struct {

	// The OCID of the compartment for the export.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The sort order for the generated export.
	SortOrder GenerateCustomerInstanceReportExportDetailsSortOrderEnum `mandatory:"false" json:"sortOrder,omitempty"`

	// The field to sort by.
	SortBy GenerateCustomerInstanceReportExportDetailsSortByEnum `mandatory:"false" json:"sortBy,omitempty"`

	// A filter to return only report records that match the listing name.
	Name *string `mandatory:"false" json:"name"`

	// A filter to return only report records that match the listing OCID.
	ListingId *string `mandatory:"false" json:"listingId"`

	// A filter to return only report records that match the instance status.
	Status GenerateCustomerInstanceReportExportDetailsStatusEnum `mandatory:"false" json:"status,omitempty"`

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
}

func (m GenerateCustomerInstanceReportExportDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m GenerateCustomerInstanceReportExportDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingGenerateCustomerInstanceReportExportDetailsSortOrderEnum(string(m.SortOrder)); !ok && m.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", m.SortOrder, strings.Join(GetGenerateCustomerInstanceReportExportDetailsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingGenerateCustomerInstanceReportExportDetailsSortByEnum(string(m.SortBy)); !ok && m.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", m.SortBy, strings.Join(GetGenerateCustomerInstanceReportExportDetailsSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingGenerateCustomerInstanceReportExportDetailsStatusEnum(string(m.Status)); !ok && m.Status != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", m.Status, strings.Join(GetGenerateCustomerInstanceReportExportDetailsStatusEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// GenerateCustomerInstanceReportExportDetailsSortOrderEnum Enum with underlying type: string
type GenerateCustomerInstanceReportExportDetailsSortOrderEnum string

// Set of constants representing the allowable values for GenerateCustomerInstanceReportExportDetailsSortOrderEnum
const (
	GenerateCustomerInstanceReportExportDetailsSortOrderAsc  GenerateCustomerInstanceReportExportDetailsSortOrderEnum = "ASC"
	GenerateCustomerInstanceReportExportDetailsSortOrderDesc GenerateCustomerInstanceReportExportDetailsSortOrderEnum = "DESC"
)

var mappingGenerateCustomerInstanceReportExportDetailsSortOrderEnum = map[string]GenerateCustomerInstanceReportExportDetailsSortOrderEnum{
	"ASC":  GenerateCustomerInstanceReportExportDetailsSortOrderAsc,
	"DESC": GenerateCustomerInstanceReportExportDetailsSortOrderDesc,
}

var mappingGenerateCustomerInstanceReportExportDetailsSortOrderEnumLowerCase = map[string]GenerateCustomerInstanceReportExportDetailsSortOrderEnum{
	"asc":  GenerateCustomerInstanceReportExportDetailsSortOrderAsc,
	"desc": GenerateCustomerInstanceReportExportDetailsSortOrderDesc,
}

// GetGenerateCustomerInstanceReportExportDetailsSortOrderEnumValues Enumerates the set of values for GenerateCustomerInstanceReportExportDetailsSortOrderEnum
func GetGenerateCustomerInstanceReportExportDetailsSortOrderEnumValues() []GenerateCustomerInstanceReportExportDetailsSortOrderEnum {
	values := make([]GenerateCustomerInstanceReportExportDetailsSortOrderEnum, 0)
	for _, v := range mappingGenerateCustomerInstanceReportExportDetailsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetGenerateCustomerInstanceReportExportDetailsSortOrderEnumStringValues Enumerates the set of values in String for GenerateCustomerInstanceReportExportDetailsSortOrderEnum
func GetGenerateCustomerInstanceReportExportDetailsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingGenerateCustomerInstanceReportExportDetailsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingGenerateCustomerInstanceReportExportDetailsSortOrderEnum(val string) (GenerateCustomerInstanceReportExportDetailsSortOrderEnum, bool) {
	enum, ok := mappingGenerateCustomerInstanceReportExportDetailsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// GenerateCustomerInstanceReportExportDetailsSortByEnum Enum with underlying type: string
type GenerateCustomerInstanceReportExportDetailsSortByEnum string

// Set of constants representing the allowable values for GenerateCustomerInstanceReportExportDetailsSortByEnum
const (
	GenerateCustomerInstanceReportExportDetailsSortByInstanceId GenerateCustomerInstanceReportExportDetailsSortByEnum = "INSTANCE_ID"
)

var mappingGenerateCustomerInstanceReportExportDetailsSortByEnum = map[string]GenerateCustomerInstanceReportExportDetailsSortByEnum{
	"INSTANCE_ID": GenerateCustomerInstanceReportExportDetailsSortByInstanceId,
}

var mappingGenerateCustomerInstanceReportExportDetailsSortByEnumLowerCase = map[string]GenerateCustomerInstanceReportExportDetailsSortByEnum{
	"instance_id": GenerateCustomerInstanceReportExportDetailsSortByInstanceId,
}

// GetGenerateCustomerInstanceReportExportDetailsSortByEnumValues Enumerates the set of values for GenerateCustomerInstanceReportExportDetailsSortByEnum
func GetGenerateCustomerInstanceReportExportDetailsSortByEnumValues() []GenerateCustomerInstanceReportExportDetailsSortByEnum {
	values := make([]GenerateCustomerInstanceReportExportDetailsSortByEnum, 0)
	for _, v := range mappingGenerateCustomerInstanceReportExportDetailsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetGenerateCustomerInstanceReportExportDetailsSortByEnumStringValues Enumerates the set of values in String for GenerateCustomerInstanceReportExportDetailsSortByEnum
func GetGenerateCustomerInstanceReportExportDetailsSortByEnumStringValues() []string {
	return []string{
		"INSTANCE_ID",
	}
}

// GetMappingGenerateCustomerInstanceReportExportDetailsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingGenerateCustomerInstanceReportExportDetailsSortByEnum(val string) (GenerateCustomerInstanceReportExportDetailsSortByEnum, bool) {
	enum, ok := mappingGenerateCustomerInstanceReportExportDetailsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// GenerateCustomerInstanceReportExportDetailsStatusEnum Enum with underlying type: string
type GenerateCustomerInstanceReportExportDetailsStatusEnum string

// Set of constants representing the allowable values for GenerateCustomerInstanceReportExportDetailsStatusEnum
const (
	GenerateCustomerInstanceReportExportDetailsStatusProvisioning GenerateCustomerInstanceReportExportDetailsStatusEnum = "PROVISIONING"
	GenerateCustomerInstanceReportExportDetailsStatusRunning      GenerateCustomerInstanceReportExportDetailsStatusEnum = "RUNNING"
	GenerateCustomerInstanceReportExportDetailsStatusStopped      GenerateCustomerInstanceReportExportDetailsStatusEnum = "STOPPED"
	GenerateCustomerInstanceReportExportDetailsStatusTerminating  GenerateCustomerInstanceReportExportDetailsStatusEnum = "TERMINATING"
	GenerateCustomerInstanceReportExportDetailsStatusTerminated   GenerateCustomerInstanceReportExportDetailsStatusEnum = "TERMINATED"
	GenerateCustomerInstanceReportExportDetailsStatusDisabled     GenerateCustomerInstanceReportExportDetailsStatusEnum = "DISABLED"
	GenerateCustomerInstanceReportExportDetailsStatusStarting     GenerateCustomerInstanceReportExportDetailsStatusEnum = "STARTING"
	GenerateCustomerInstanceReportExportDetailsStatusStopping     GenerateCustomerInstanceReportExportDetailsStatusEnum = "STOPPING"
	GenerateCustomerInstanceReportExportDetailsStatusSnapshotting GenerateCustomerInstanceReportExportDetailsStatusEnum = "SNAPSHOTTING"
)

var mappingGenerateCustomerInstanceReportExportDetailsStatusEnum = map[string]GenerateCustomerInstanceReportExportDetailsStatusEnum{
	"PROVISIONING": GenerateCustomerInstanceReportExportDetailsStatusProvisioning,
	"RUNNING":      GenerateCustomerInstanceReportExportDetailsStatusRunning,
	"STOPPED":      GenerateCustomerInstanceReportExportDetailsStatusStopped,
	"TERMINATING":  GenerateCustomerInstanceReportExportDetailsStatusTerminating,
	"TERMINATED":   GenerateCustomerInstanceReportExportDetailsStatusTerminated,
	"DISABLED":     GenerateCustomerInstanceReportExportDetailsStatusDisabled,
	"STARTING":     GenerateCustomerInstanceReportExportDetailsStatusStarting,
	"STOPPING":     GenerateCustomerInstanceReportExportDetailsStatusStopping,
	"SNAPSHOTTING": GenerateCustomerInstanceReportExportDetailsStatusSnapshotting,
}

var mappingGenerateCustomerInstanceReportExportDetailsStatusEnumLowerCase = map[string]GenerateCustomerInstanceReportExportDetailsStatusEnum{
	"provisioning": GenerateCustomerInstanceReportExportDetailsStatusProvisioning,
	"running":      GenerateCustomerInstanceReportExportDetailsStatusRunning,
	"stopped":      GenerateCustomerInstanceReportExportDetailsStatusStopped,
	"terminating":  GenerateCustomerInstanceReportExportDetailsStatusTerminating,
	"terminated":   GenerateCustomerInstanceReportExportDetailsStatusTerminated,
	"disabled":     GenerateCustomerInstanceReportExportDetailsStatusDisabled,
	"starting":     GenerateCustomerInstanceReportExportDetailsStatusStarting,
	"stopping":     GenerateCustomerInstanceReportExportDetailsStatusStopping,
	"snapshotting": GenerateCustomerInstanceReportExportDetailsStatusSnapshotting,
}

// GetGenerateCustomerInstanceReportExportDetailsStatusEnumValues Enumerates the set of values for GenerateCustomerInstanceReportExportDetailsStatusEnum
func GetGenerateCustomerInstanceReportExportDetailsStatusEnumValues() []GenerateCustomerInstanceReportExportDetailsStatusEnum {
	values := make([]GenerateCustomerInstanceReportExportDetailsStatusEnum, 0)
	for _, v := range mappingGenerateCustomerInstanceReportExportDetailsStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetGenerateCustomerInstanceReportExportDetailsStatusEnumStringValues Enumerates the set of values in String for GenerateCustomerInstanceReportExportDetailsStatusEnum
func GetGenerateCustomerInstanceReportExportDetailsStatusEnumStringValues() []string {
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

// GetMappingGenerateCustomerInstanceReportExportDetailsStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingGenerateCustomerInstanceReportExportDetailsStatusEnum(val string) (GenerateCustomerInstanceReportExportDetailsStatusEnum, bool) {
	enum, ok := mappingGenerateCustomerInstanceReportExportDetailsStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
