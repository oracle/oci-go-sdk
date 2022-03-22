// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Transfer Service API
//
// Data Transfer Service API Specification
//

package dts

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v63/common"
	"strings"
)

// ApplianceExportJobSummary The representation of ApplianceExportJobSummary
type ApplianceExportJobSummary struct {
	Id *string `mandatory:"false" json:"id"`

	BucketName *string `mandatory:"false" json:"bucketName"`

	DisplayName *string `mandatory:"false" json:"displayName"`

	CreationTime *common.SDKTime `mandatory:"false" json:"creationTime"`

	LifecycleState ApplianceExportJobSummaryLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	// A property that can contain details on the lifecycle.
	LifecycleStateDetails *string `mandatory:"false" json:"lifecycleStateDetails"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no
	// predefined name, type, or namespace. For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a
	// namespace. For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m ApplianceExportJobSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ApplianceExportJobSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingApplianceExportJobSummaryLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetApplianceExportJobSummaryLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ApplianceExportJobSummaryLifecycleStateEnum Enum with underlying type: string
type ApplianceExportJobSummaryLifecycleStateEnum string

// Set of constants representing the allowable values for ApplianceExportJobSummaryLifecycleStateEnum
const (
	ApplianceExportJobSummaryLifecycleStateCreating   ApplianceExportJobSummaryLifecycleStateEnum = "CREATING"
	ApplianceExportJobSummaryLifecycleStateActive     ApplianceExportJobSummaryLifecycleStateEnum = "ACTIVE"
	ApplianceExportJobSummaryLifecycleStateInprogress ApplianceExportJobSummaryLifecycleStateEnum = "INPROGRESS"
	ApplianceExportJobSummaryLifecycleStateSucceeded  ApplianceExportJobSummaryLifecycleStateEnum = "SUCCEEDED"
	ApplianceExportJobSummaryLifecycleStateFailed     ApplianceExportJobSummaryLifecycleStateEnum = "FAILED"
	ApplianceExportJobSummaryLifecycleStateCancelled  ApplianceExportJobSummaryLifecycleStateEnum = "CANCELLED"
	ApplianceExportJobSummaryLifecycleStateDeleted    ApplianceExportJobSummaryLifecycleStateEnum = "DELETED"
)

var mappingApplianceExportJobSummaryLifecycleStateEnum = map[string]ApplianceExportJobSummaryLifecycleStateEnum{
	"CREATING":   ApplianceExportJobSummaryLifecycleStateCreating,
	"ACTIVE":     ApplianceExportJobSummaryLifecycleStateActive,
	"INPROGRESS": ApplianceExportJobSummaryLifecycleStateInprogress,
	"SUCCEEDED":  ApplianceExportJobSummaryLifecycleStateSucceeded,
	"FAILED":     ApplianceExportJobSummaryLifecycleStateFailed,
	"CANCELLED":  ApplianceExportJobSummaryLifecycleStateCancelled,
	"DELETED":    ApplianceExportJobSummaryLifecycleStateDeleted,
}

var mappingApplianceExportJobSummaryLifecycleStateEnumLowerCase = map[string]ApplianceExportJobSummaryLifecycleStateEnum{
	"creating":   ApplianceExportJobSummaryLifecycleStateCreating,
	"active":     ApplianceExportJobSummaryLifecycleStateActive,
	"inprogress": ApplianceExportJobSummaryLifecycleStateInprogress,
	"succeeded":  ApplianceExportJobSummaryLifecycleStateSucceeded,
	"failed":     ApplianceExportJobSummaryLifecycleStateFailed,
	"cancelled":  ApplianceExportJobSummaryLifecycleStateCancelled,
	"deleted":    ApplianceExportJobSummaryLifecycleStateDeleted,
}

// GetApplianceExportJobSummaryLifecycleStateEnumValues Enumerates the set of values for ApplianceExportJobSummaryLifecycleStateEnum
func GetApplianceExportJobSummaryLifecycleStateEnumValues() []ApplianceExportJobSummaryLifecycleStateEnum {
	values := make([]ApplianceExportJobSummaryLifecycleStateEnum, 0)
	for _, v := range mappingApplianceExportJobSummaryLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetApplianceExportJobSummaryLifecycleStateEnumStringValues Enumerates the set of values in String for ApplianceExportJobSummaryLifecycleStateEnum
func GetApplianceExportJobSummaryLifecycleStateEnumStringValues() []string {
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

// GetMappingApplianceExportJobSummaryLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingApplianceExportJobSummaryLifecycleStateEnum(val string) (ApplianceExportJobSummaryLifecycleStateEnum, bool) {
	enum, ok := mappingApplianceExportJobSummaryLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
