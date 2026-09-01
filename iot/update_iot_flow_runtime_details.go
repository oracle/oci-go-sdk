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

// UpdateIotFlowRuntimeDetails Details for updating an IoT flow runtime. This operation supports partial updates. When networkConfig is omitted, the existing network configuration is preserved. When networkConfig is provided, the complete network configuration is replaced. To remove the network configuration, pass networkConfig as null.
type UpdateIotFlowRuntimeDetails struct {

	// A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// A short description of the resource.
	Description *string `mandatory:"false" json:"description"`

	// The scale of the IoT flow runtime. Larger values allocate more CPU and memory for higher throughput and operational headroom. MEDIUM is the default value.
	Scale UpdateIotFlowRuntimeDetailsScaleEnum `mandatory:"false" json:"scale,omitempty"`

	NetworkConfig *NetworkConfigDetails `mandatory:"false" json:"networkConfig"`

	LogConfig *LogConfigDetails `mandatory:"false" json:"logConfig"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m UpdateIotFlowRuntimeDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpdateIotFlowRuntimeDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingUpdateIotFlowRuntimeDetailsScaleEnum(string(m.Scale)); !ok && m.Scale != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Scale: %s. Supported values are: %s.", m.Scale, strings.Join(GetUpdateIotFlowRuntimeDetailsScaleEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UpdateIotFlowRuntimeDetailsScaleEnum Enum with underlying type: string
type UpdateIotFlowRuntimeDetailsScaleEnum string

// Set of constants representing the allowable values for UpdateIotFlowRuntimeDetailsScaleEnum
const (
	UpdateIotFlowRuntimeDetailsScaleLowest  UpdateIotFlowRuntimeDetailsScaleEnum = "LOWEST"
	UpdateIotFlowRuntimeDetailsScaleLow     UpdateIotFlowRuntimeDetailsScaleEnum = "LOW"
	UpdateIotFlowRuntimeDetailsScaleMedium  UpdateIotFlowRuntimeDetailsScaleEnum = "MEDIUM"
	UpdateIotFlowRuntimeDetailsScaleHigh    UpdateIotFlowRuntimeDetailsScaleEnum = "HIGH"
	UpdateIotFlowRuntimeDetailsScaleHighest UpdateIotFlowRuntimeDetailsScaleEnum = "HIGHEST"
)

var mappingUpdateIotFlowRuntimeDetailsScaleEnum = map[string]UpdateIotFlowRuntimeDetailsScaleEnum{
	"LOWEST":  UpdateIotFlowRuntimeDetailsScaleLowest,
	"LOW":     UpdateIotFlowRuntimeDetailsScaleLow,
	"MEDIUM":  UpdateIotFlowRuntimeDetailsScaleMedium,
	"HIGH":    UpdateIotFlowRuntimeDetailsScaleHigh,
	"HIGHEST": UpdateIotFlowRuntimeDetailsScaleHighest,
}

var mappingUpdateIotFlowRuntimeDetailsScaleEnumLowerCase = map[string]UpdateIotFlowRuntimeDetailsScaleEnum{
	"lowest":  UpdateIotFlowRuntimeDetailsScaleLowest,
	"low":     UpdateIotFlowRuntimeDetailsScaleLow,
	"medium":  UpdateIotFlowRuntimeDetailsScaleMedium,
	"high":    UpdateIotFlowRuntimeDetailsScaleHigh,
	"highest": UpdateIotFlowRuntimeDetailsScaleHighest,
}

// GetUpdateIotFlowRuntimeDetailsScaleEnumValues Enumerates the set of values for UpdateIotFlowRuntimeDetailsScaleEnum
func GetUpdateIotFlowRuntimeDetailsScaleEnumValues() []UpdateIotFlowRuntimeDetailsScaleEnum {
	values := make([]UpdateIotFlowRuntimeDetailsScaleEnum, 0)
	for _, v := range mappingUpdateIotFlowRuntimeDetailsScaleEnum {
		values = append(values, v)
	}
	return values
}

// GetUpdateIotFlowRuntimeDetailsScaleEnumStringValues Enumerates the set of values in String for UpdateIotFlowRuntimeDetailsScaleEnum
func GetUpdateIotFlowRuntimeDetailsScaleEnumStringValues() []string {
	return []string{
		"LOWEST",
		"LOW",
		"MEDIUM",
		"HIGH",
		"HIGHEST",
	}
}

// GetMappingUpdateIotFlowRuntimeDetailsScaleEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingUpdateIotFlowRuntimeDetailsScaleEnum(val string) (UpdateIotFlowRuntimeDetailsScaleEnum, bool) {
	enum, ok := mappingUpdateIotFlowRuntimeDetailsScaleEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
