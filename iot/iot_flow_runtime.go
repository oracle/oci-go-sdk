// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Internet of Things API
//
// Use the Internet of Things (IoT) API to manage IoT domain groups, domains, and digital twin resources including models, adapters, instances, and relationships.
// For more information, see Internet of Things (https://docs.oracle.com/iaas/Content/internet-of-things/home.htm).
//

package iot

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// IotFlowRuntime An IoT flow runtime is a managed Node-RED runtime in an IoT domain for building and running Node-RED flows.
type IotFlowRuntime struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the resource.
	Id *string `mandatory:"true" json:"id"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the IoT domain.
	IotDomainId *string `mandatory:"true" json:"iotDomainId"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment corresponding to the resource.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The scale of the IoT flow runtime. Larger values allocate more CPU and memory for higher throughput and operational headroom. MEDIUM is the default value.
	Scale IotFlowRuntimeScaleEnum `mandatory:"true" json:"scale"`

	// The current state of the IoT flow runtime.
	LifecycleState IotFlowRuntimeLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The date and time when the resource was created, in the format defined by RFC 3339 (https://tools.ietf.org/html/rfc3339).
	// Example: `2016-08-25T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// A short description of the resource.
	Description *string `mandatory:"false" json:"description"`

	LogConfig *LogConfigDetails `mandatory:"false" json:"logConfig"`

	// Hostname of the IoT flow runtime.
	FlowRuntimeHost *string `mandatory:"false" json:"flowRuntimeHost"`

	NetworkConfig *NetworkConfigDetails `mandatory:"false" json:"networkConfig"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// System tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`

	// The date and time when the resource was last updated, in the format defined by RFC 3339 (https://tools.ietf.org/html/rfc3339).
	// Example: `2016-08-25T21:10:29.600Z`
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`
}

func (m IotFlowRuntime) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m IotFlowRuntime) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingIotFlowRuntimeScaleEnum(string(m.Scale)); !ok && m.Scale != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Scale: %s. Supported values are: %s.", m.Scale, strings.Join(GetIotFlowRuntimeScaleEnumStringValues(), ",")))
	}
	if _, ok := GetMappingIotFlowRuntimeLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetIotFlowRuntimeLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// IotFlowRuntimeScaleEnum Enum with underlying type: string
type IotFlowRuntimeScaleEnum string

// Set of constants representing the allowable values for IotFlowRuntimeScaleEnum
const (
	IotFlowRuntimeScaleLowest  IotFlowRuntimeScaleEnum = "LOWEST"
	IotFlowRuntimeScaleLow     IotFlowRuntimeScaleEnum = "LOW"
	IotFlowRuntimeScaleMedium  IotFlowRuntimeScaleEnum = "MEDIUM"
	IotFlowRuntimeScaleHigh    IotFlowRuntimeScaleEnum = "HIGH"
	IotFlowRuntimeScaleHighest IotFlowRuntimeScaleEnum = "HIGHEST"
)

var mappingIotFlowRuntimeScaleEnum = map[string]IotFlowRuntimeScaleEnum{
	"LOWEST":  IotFlowRuntimeScaleLowest,
	"LOW":     IotFlowRuntimeScaleLow,
	"MEDIUM":  IotFlowRuntimeScaleMedium,
	"HIGH":    IotFlowRuntimeScaleHigh,
	"HIGHEST": IotFlowRuntimeScaleHighest,
}

var mappingIotFlowRuntimeScaleEnumLowerCase = map[string]IotFlowRuntimeScaleEnum{
	"lowest":  IotFlowRuntimeScaleLowest,
	"low":     IotFlowRuntimeScaleLow,
	"medium":  IotFlowRuntimeScaleMedium,
	"high":    IotFlowRuntimeScaleHigh,
	"highest": IotFlowRuntimeScaleHighest,
}

// GetIotFlowRuntimeScaleEnumValues Enumerates the set of values for IotFlowRuntimeScaleEnum
func GetIotFlowRuntimeScaleEnumValues() []IotFlowRuntimeScaleEnum {
	values := make([]IotFlowRuntimeScaleEnum, 0)
	for _, v := range mappingIotFlowRuntimeScaleEnum {
		values = append(values, v)
	}
	return values
}

// GetIotFlowRuntimeScaleEnumStringValues Enumerates the set of values in String for IotFlowRuntimeScaleEnum
func GetIotFlowRuntimeScaleEnumStringValues() []string {
	return []string{
		"LOWEST",
		"LOW",
		"MEDIUM",
		"HIGH",
		"HIGHEST",
	}
}

// GetMappingIotFlowRuntimeScaleEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingIotFlowRuntimeScaleEnum(val string) (IotFlowRuntimeScaleEnum, bool) {
	enum, ok := mappingIotFlowRuntimeScaleEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// IotFlowRuntimeLifecycleStateEnum Enum with underlying type: string
type IotFlowRuntimeLifecycleStateEnum string

// Set of constants representing the allowable values for IotFlowRuntimeLifecycleStateEnum
const (
	IotFlowRuntimeLifecycleStateCreating IotFlowRuntimeLifecycleStateEnum = "CREATING"
	IotFlowRuntimeLifecycleStateUpdating IotFlowRuntimeLifecycleStateEnum = "UPDATING"
	IotFlowRuntimeLifecycleStateActive   IotFlowRuntimeLifecycleStateEnum = "ACTIVE"
	IotFlowRuntimeLifecycleStateInactive IotFlowRuntimeLifecycleStateEnum = "INACTIVE"
	IotFlowRuntimeLifecycleStateDeleting IotFlowRuntimeLifecycleStateEnum = "DELETING"
	IotFlowRuntimeLifecycleStateDeleted  IotFlowRuntimeLifecycleStateEnum = "DELETED"
	IotFlowRuntimeLifecycleStateFailed   IotFlowRuntimeLifecycleStateEnum = "FAILED"
)

var mappingIotFlowRuntimeLifecycleStateEnum = map[string]IotFlowRuntimeLifecycleStateEnum{
	"CREATING": IotFlowRuntimeLifecycleStateCreating,
	"UPDATING": IotFlowRuntimeLifecycleStateUpdating,
	"ACTIVE":   IotFlowRuntimeLifecycleStateActive,
	"INACTIVE": IotFlowRuntimeLifecycleStateInactive,
	"DELETING": IotFlowRuntimeLifecycleStateDeleting,
	"DELETED":  IotFlowRuntimeLifecycleStateDeleted,
	"FAILED":   IotFlowRuntimeLifecycleStateFailed,
}

var mappingIotFlowRuntimeLifecycleStateEnumLowerCase = map[string]IotFlowRuntimeLifecycleStateEnum{
	"creating": IotFlowRuntimeLifecycleStateCreating,
	"updating": IotFlowRuntimeLifecycleStateUpdating,
	"active":   IotFlowRuntimeLifecycleStateActive,
	"inactive": IotFlowRuntimeLifecycleStateInactive,
	"deleting": IotFlowRuntimeLifecycleStateDeleting,
	"deleted":  IotFlowRuntimeLifecycleStateDeleted,
	"failed":   IotFlowRuntimeLifecycleStateFailed,
}

// GetIotFlowRuntimeLifecycleStateEnumValues Enumerates the set of values for IotFlowRuntimeLifecycleStateEnum
func GetIotFlowRuntimeLifecycleStateEnumValues() []IotFlowRuntimeLifecycleStateEnum {
	values := make([]IotFlowRuntimeLifecycleStateEnum, 0)
	for _, v := range mappingIotFlowRuntimeLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetIotFlowRuntimeLifecycleStateEnumStringValues Enumerates the set of values in String for IotFlowRuntimeLifecycleStateEnum
func GetIotFlowRuntimeLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"UPDATING",
		"ACTIVE",
		"INACTIVE",
		"DELETING",
		"DELETED",
		"FAILED",
	}
}

// GetMappingIotFlowRuntimeLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingIotFlowRuntimeLifecycleStateEnum(val string) (IotFlowRuntimeLifecycleStateEnum, bool) {
	enum, ok := mappingIotFlowRuntimeLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
