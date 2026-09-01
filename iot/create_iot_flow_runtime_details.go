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

// CreateIotFlowRuntimeDetails Details for creating an IoT flow runtime in an IoT domain.
type CreateIotFlowRuntimeDetails struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the IoT domain.
	IotDomainId *string `mandatory:"true" json:"iotDomainId"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment corresponding to the resource.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The scale of the IoT flow runtime. Larger values allocate more CPU and memory for higher throughput and operational headroom. MEDIUM is the default value.
	Scale CreateIotFlowRuntimeDetailsScaleEnum `mandatory:"false" json:"scale,omitempty"`

	NetworkConfig *NetworkConfigDetails `mandatory:"false" json:"networkConfig"`

	LogConfig *LogConfigDetails `mandatory:"false" json:"logConfig"`

	// A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.
	DisplayName *string `mandatory:"false" json:"displayName"`

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
}

func (m CreateIotFlowRuntimeDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateIotFlowRuntimeDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingCreateIotFlowRuntimeDetailsScaleEnum(string(m.Scale)); !ok && m.Scale != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Scale: %s. Supported values are: %s.", m.Scale, strings.Join(GetCreateIotFlowRuntimeDetailsScaleEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// CreateIotFlowRuntimeDetailsScaleEnum Enum with underlying type: string
type CreateIotFlowRuntimeDetailsScaleEnum string

// Set of constants representing the allowable values for CreateIotFlowRuntimeDetailsScaleEnum
const (
	CreateIotFlowRuntimeDetailsScaleLowest  CreateIotFlowRuntimeDetailsScaleEnum = "LOWEST"
	CreateIotFlowRuntimeDetailsScaleLow     CreateIotFlowRuntimeDetailsScaleEnum = "LOW"
	CreateIotFlowRuntimeDetailsScaleMedium  CreateIotFlowRuntimeDetailsScaleEnum = "MEDIUM"
	CreateIotFlowRuntimeDetailsScaleHigh    CreateIotFlowRuntimeDetailsScaleEnum = "HIGH"
	CreateIotFlowRuntimeDetailsScaleHighest CreateIotFlowRuntimeDetailsScaleEnum = "HIGHEST"
)

var mappingCreateIotFlowRuntimeDetailsScaleEnum = map[string]CreateIotFlowRuntimeDetailsScaleEnum{
	"LOWEST":  CreateIotFlowRuntimeDetailsScaleLowest,
	"LOW":     CreateIotFlowRuntimeDetailsScaleLow,
	"MEDIUM":  CreateIotFlowRuntimeDetailsScaleMedium,
	"HIGH":    CreateIotFlowRuntimeDetailsScaleHigh,
	"HIGHEST": CreateIotFlowRuntimeDetailsScaleHighest,
}

var mappingCreateIotFlowRuntimeDetailsScaleEnumLowerCase = map[string]CreateIotFlowRuntimeDetailsScaleEnum{
	"lowest":  CreateIotFlowRuntimeDetailsScaleLowest,
	"low":     CreateIotFlowRuntimeDetailsScaleLow,
	"medium":  CreateIotFlowRuntimeDetailsScaleMedium,
	"high":    CreateIotFlowRuntimeDetailsScaleHigh,
	"highest": CreateIotFlowRuntimeDetailsScaleHighest,
}

// GetCreateIotFlowRuntimeDetailsScaleEnumValues Enumerates the set of values for CreateIotFlowRuntimeDetailsScaleEnum
func GetCreateIotFlowRuntimeDetailsScaleEnumValues() []CreateIotFlowRuntimeDetailsScaleEnum {
	values := make([]CreateIotFlowRuntimeDetailsScaleEnum, 0)
	for _, v := range mappingCreateIotFlowRuntimeDetailsScaleEnum {
		values = append(values, v)
	}
	return values
}

// GetCreateIotFlowRuntimeDetailsScaleEnumStringValues Enumerates the set of values in String for CreateIotFlowRuntimeDetailsScaleEnum
func GetCreateIotFlowRuntimeDetailsScaleEnumStringValues() []string {
	return []string{
		"LOWEST",
		"LOW",
		"MEDIUM",
		"HIGH",
		"HIGHEST",
	}
}

// GetMappingCreateIotFlowRuntimeDetailsScaleEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCreateIotFlowRuntimeDetailsScaleEnum(val string) (CreateIotFlowRuntimeDetailsScaleEnum, bool) {
	enum, ok := mappingCreateIotFlowRuntimeDetailsScaleEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
