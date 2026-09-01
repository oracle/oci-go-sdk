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

// IotFlowRuntimeSummary Summary information about an IoT flow runtime.
type IotFlowRuntimeSummary struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the resource.
	Id *string `mandatory:"true" json:"id"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the IoT domain.
	IotDomainId *string `mandatory:"true" json:"iotDomainId"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment corresponding to the resource.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The scale of the IoT flow runtime. Larger values allocate more CPU and memory for higher throughput and operational headroom. MEDIUM is the default value.
	Scale IotFlowRuntimeSummaryScaleEnum `mandatory:"true" json:"scale"`

	// A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The current state of the IoT flow runtime.
	LifecycleState IotFlowRuntimeLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The date and time when the resource was created, in the format defined by RFC 3339 (https://tools.ietf.org/html/rfc3339).
	// Example: `2016-08-25T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// A short description of the resource.
	Description *string `mandatory:"false" json:"description"`

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

func (m IotFlowRuntimeSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m IotFlowRuntimeSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingIotFlowRuntimeSummaryScaleEnum(string(m.Scale)); !ok && m.Scale != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Scale: %s. Supported values are: %s.", m.Scale, strings.Join(GetIotFlowRuntimeSummaryScaleEnumStringValues(), ",")))
	}
	if _, ok := GetMappingIotFlowRuntimeLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetIotFlowRuntimeLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// IotFlowRuntimeSummaryScaleEnum Enum with underlying type: string
type IotFlowRuntimeSummaryScaleEnum string

// Set of constants representing the allowable values for IotFlowRuntimeSummaryScaleEnum
const (
	IotFlowRuntimeSummaryScaleLowest  IotFlowRuntimeSummaryScaleEnum = "LOWEST"
	IotFlowRuntimeSummaryScaleLow     IotFlowRuntimeSummaryScaleEnum = "LOW"
	IotFlowRuntimeSummaryScaleMedium  IotFlowRuntimeSummaryScaleEnum = "MEDIUM"
	IotFlowRuntimeSummaryScaleHigh    IotFlowRuntimeSummaryScaleEnum = "HIGH"
	IotFlowRuntimeSummaryScaleHighest IotFlowRuntimeSummaryScaleEnum = "HIGHEST"
)

var mappingIotFlowRuntimeSummaryScaleEnum = map[string]IotFlowRuntimeSummaryScaleEnum{
	"LOWEST":  IotFlowRuntimeSummaryScaleLowest,
	"LOW":     IotFlowRuntimeSummaryScaleLow,
	"MEDIUM":  IotFlowRuntimeSummaryScaleMedium,
	"HIGH":    IotFlowRuntimeSummaryScaleHigh,
	"HIGHEST": IotFlowRuntimeSummaryScaleHighest,
}

var mappingIotFlowRuntimeSummaryScaleEnumLowerCase = map[string]IotFlowRuntimeSummaryScaleEnum{
	"lowest":  IotFlowRuntimeSummaryScaleLowest,
	"low":     IotFlowRuntimeSummaryScaleLow,
	"medium":  IotFlowRuntimeSummaryScaleMedium,
	"high":    IotFlowRuntimeSummaryScaleHigh,
	"highest": IotFlowRuntimeSummaryScaleHighest,
}

// GetIotFlowRuntimeSummaryScaleEnumValues Enumerates the set of values for IotFlowRuntimeSummaryScaleEnum
func GetIotFlowRuntimeSummaryScaleEnumValues() []IotFlowRuntimeSummaryScaleEnum {
	values := make([]IotFlowRuntimeSummaryScaleEnum, 0)
	for _, v := range mappingIotFlowRuntimeSummaryScaleEnum {
		values = append(values, v)
	}
	return values
}

// GetIotFlowRuntimeSummaryScaleEnumStringValues Enumerates the set of values in String for IotFlowRuntimeSummaryScaleEnum
func GetIotFlowRuntimeSummaryScaleEnumStringValues() []string {
	return []string{
		"LOWEST",
		"LOW",
		"MEDIUM",
		"HIGH",
		"HIGHEST",
	}
}

// GetMappingIotFlowRuntimeSummaryScaleEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingIotFlowRuntimeSummaryScaleEnum(val string) (IotFlowRuntimeSummaryScaleEnum, bool) {
	enum, ok := mappingIotFlowRuntimeSummaryScaleEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
