// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Transfer Service API
//
// Data Transfer Service API Specification
//

package dts

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// UpdateApplianceExportJobDetails The representation of UpdateApplianceExportJobDetails
type UpdateApplianceExportJobDetails struct {
	BucketName *string `mandatory:"false" json:"bucketName"`

	// List of objects with names matching this prefix would be part of this export job.
	Prefix *string `mandatory:"false" json:"prefix"`

	// Object names returned by a list query must be greater or equal to this parameter.
	RangeStart *string `mandatory:"false" json:"rangeStart"`

	// Object names returned by a list query must be strictly less than this parameter.
	RangeEnd *string `mandatory:"false" json:"rangeEnd"`

	DisplayName *string `mandatory:"false" json:"displayName"`

	LifecycleState UpdateApplianceExportJobDetailsLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	// A property that can contain details on the lifecycle.
	LifecycleStateDetails *string `mandatory:"false" json:"lifecycleStateDetails"`

	// Manifest File associated with this export job.
	ManifestFile *string `mandatory:"false" json:"manifestFile"`

	// md5 digest of the manifest file.
	ManifestMd5 *string `mandatory:"false" json:"manifestMd5"`

	// Total number of objects that are exported in this job.
	NumberOfObjects *string `mandatory:"false" json:"numberOfObjects"`

	// Total size of objects in Bytes that are exported in this job.
	TotalSizeInBytes *string `mandatory:"false" json:"totalSizeInBytes"`

	// First object in the list of objects that are exported in this job.
	FirstObject *string `mandatory:"false" json:"firstObject"`

	// Last object in the list of objects that are exported in this job.
	LastObject *string `mandatory:"false" json:"lastObject"`

	// First object from which the next potential export job could start.
	NextObject *string `mandatory:"false" json:"nextObject"`

	// Expected return date from customer for the device, time portion should be zero.
	ExpectedReturnDate *common.SDKTime `mandatory:"false" json:"expectedReturnDate"`

	// Start time for the window to pickup the device from customer.
	PickupWindowStartTime *common.SDKTime `mandatory:"false" json:"pickupWindowStartTime"`

	// End time for the window to pickup the device from customer.
	PickupWindowEndTime *common.SDKTime `mandatory:"false" json:"pickupWindowEndTime"`

	CustomerShippingAddress *ShippingAddress `mandatory:"false" json:"customerShippingAddress"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no
	// predefined name, type, or namespace. For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a
	// namespace. For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m UpdateApplianceExportJobDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpdateApplianceExportJobDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingUpdateApplianceExportJobDetailsLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetUpdateApplianceExportJobDetailsLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UpdateApplianceExportJobDetailsLifecycleStateEnum Enum with underlying type: string
type UpdateApplianceExportJobDetailsLifecycleStateEnum string

// Set of constants representing the allowable values for UpdateApplianceExportJobDetailsLifecycleStateEnum
const (
	UpdateApplianceExportJobDetailsLifecycleStateCreating   UpdateApplianceExportJobDetailsLifecycleStateEnum = "CREATING"
	UpdateApplianceExportJobDetailsLifecycleStateActive     UpdateApplianceExportJobDetailsLifecycleStateEnum = "ACTIVE"
	UpdateApplianceExportJobDetailsLifecycleStateInprogress UpdateApplianceExportJobDetailsLifecycleStateEnum = "INPROGRESS"
	UpdateApplianceExportJobDetailsLifecycleStateSucceeded  UpdateApplianceExportJobDetailsLifecycleStateEnum = "SUCCEEDED"
	UpdateApplianceExportJobDetailsLifecycleStateFailed     UpdateApplianceExportJobDetailsLifecycleStateEnum = "FAILED"
	UpdateApplianceExportJobDetailsLifecycleStateCancelled  UpdateApplianceExportJobDetailsLifecycleStateEnum = "CANCELLED"
	UpdateApplianceExportJobDetailsLifecycleStateDeleted    UpdateApplianceExportJobDetailsLifecycleStateEnum = "DELETED"
)

var mappingUpdateApplianceExportJobDetailsLifecycleStateEnum = map[string]UpdateApplianceExportJobDetailsLifecycleStateEnum{
	"CREATING":   UpdateApplianceExportJobDetailsLifecycleStateCreating,
	"ACTIVE":     UpdateApplianceExportJobDetailsLifecycleStateActive,
	"INPROGRESS": UpdateApplianceExportJobDetailsLifecycleStateInprogress,
	"SUCCEEDED":  UpdateApplianceExportJobDetailsLifecycleStateSucceeded,
	"FAILED":     UpdateApplianceExportJobDetailsLifecycleStateFailed,
	"CANCELLED":  UpdateApplianceExportJobDetailsLifecycleStateCancelled,
	"DELETED":    UpdateApplianceExportJobDetailsLifecycleStateDeleted,
}

var mappingUpdateApplianceExportJobDetailsLifecycleStateEnumLowerCase = map[string]UpdateApplianceExportJobDetailsLifecycleStateEnum{
	"creating":   UpdateApplianceExportJobDetailsLifecycleStateCreating,
	"active":     UpdateApplianceExportJobDetailsLifecycleStateActive,
	"inprogress": UpdateApplianceExportJobDetailsLifecycleStateInprogress,
	"succeeded":  UpdateApplianceExportJobDetailsLifecycleStateSucceeded,
	"failed":     UpdateApplianceExportJobDetailsLifecycleStateFailed,
	"cancelled":  UpdateApplianceExportJobDetailsLifecycleStateCancelled,
	"deleted":    UpdateApplianceExportJobDetailsLifecycleStateDeleted,
}

// GetUpdateApplianceExportJobDetailsLifecycleStateEnumValues Enumerates the set of values for UpdateApplianceExportJobDetailsLifecycleStateEnum
func GetUpdateApplianceExportJobDetailsLifecycleStateEnumValues() []UpdateApplianceExportJobDetailsLifecycleStateEnum {
	values := make([]UpdateApplianceExportJobDetailsLifecycleStateEnum, 0)
	for _, v := range mappingUpdateApplianceExportJobDetailsLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetUpdateApplianceExportJobDetailsLifecycleStateEnumStringValues Enumerates the set of values in String for UpdateApplianceExportJobDetailsLifecycleStateEnum
func GetUpdateApplianceExportJobDetailsLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"ACTIVE",
		"INPROGRESS",
		"SUCCEEDED",
		"FAILED",
		"CANCELLED",
		"DELETED",
	}
}

// GetMappingUpdateApplianceExportJobDetailsLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingUpdateApplianceExportJobDetailsLifecycleStateEnum(val string) (UpdateApplianceExportJobDetailsLifecycleStateEnum, bool) {
	enum, ok := mappingUpdateApplianceExportJobDetailsLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
