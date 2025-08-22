// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
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

// ResourceBundle Resource Bundle associated with an Offer
type ResourceBundle struct {

	// The type of resources in the bundle
	Type ResourceBundleTypeEnum `mandatory:"false" json:"type,omitempty"`

	// The quantity of a resources associated with the bundle
	Quantity *int64 `mandatory:"false" json:"quantity"`

	// The unit of measurement for the resource bundle
	UnitOfMeasurement ResourceBundleUnitOfMeasurementEnum `mandatory:"false" json:"unitOfMeasurement,omitempty"`

	// the ids of the resources in the Offer
	ResourceIds []string `mandatory:"false" json:"resourceIds"`
}

func (m ResourceBundle) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ResourceBundle) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingResourceBundleTypeEnum(string(m.Type)); !ok && m.Type != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Type: %s. Supported values are: %s.", m.Type, strings.Join(GetResourceBundleTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingResourceBundleUnitOfMeasurementEnum(string(m.UnitOfMeasurement)); !ok && m.UnitOfMeasurement != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for UnitOfMeasurement: %s. Supported values are: %s.", m.UnitOfMeasurement, strings.Join(GetResourceBundleUnitOfMeasurementEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ResourceBundleTypeEnum Enum with underlying type: string
type ResourceBundleTypeEnum string

// Set of constants representing the allowable values for ResourceBundleTypeEnum
const (
	ResourceBundleTypeListing ResourceBundleTypeEnum = "LISTING"
)

var mappingResourceBundleTypeEnum = map[string]ResourceBundleTypeEnum{
	"LISTING": ResourceBundleTypeListing,
}

var mappingResourceBundleTypeEnumLowerCase = map[string]ResourceBundleTypeEnum{
	"listing": ResourceBundleTypeListing,
}

// GetResourceBundleTypeEnumValues Enumerates the set of values for ResourceBundleTypeEnum
func GetResourceBundleTypeEnumValues() []ResourceBundleTypeEnum {
	values := make([]ResourceBundleTypeEnum, 0)
	for _, v := range mappingResourceBundleTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetResourceBundleTypeEnumStringValues Enumerates the set of values in String for ResourceBundleTypeEnum
func GetResourceBundleTypeEnumStringValues() []string {
	return []string{
		"LISTING",
	}
}

// GetMappingResourceBundleTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingResourceBundleTypeEnum(val string) (ResourceBundleTypeEnum, bool) {
	enum, ok := mappingResourceBundleTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ResourceBundleUnitOfMeasurementEnum Enum with underlying type: string
type ResourceBundleUnitOfMeasurementEnum string

// Set of constants representing the allowable values for ResourceBundleUnitOfMeasurementEnum
const (
	ResourceBundleUnitOfMeasurementOcpuPerHour     ResourceBundleUnitOfMeasurementEnum = "OCPU_PER_HOUR"
	ResourceBundleUnitOfMeasurementInstancePerHour ResourceBundleUnitOfMeasurementEnum = "INSTANCE_PER_HOUR"
	ResourceBundleUnitOfMeasurementCredits         ResourceBundleUnitOfMeasurementEnum = "CREDITS"
	ResourceBundleUnitOfMeasurementInstances       ResourceBundleUnitOfMeasurementEnum = "INSTANCES"
	ResourceBundleUnitOfMeasurementNodes           ResourceBundleUnitOfMeasurementEnum = "NODES"
)

var mappingResourceBundleUnitOfMeasurementEnum = map[string]ResourceBundleUnitOfMeasurementEnum{
	"OCPU_PER_HOUR":     ResourceBundleUnitOfMeasurementOcpuPerHour,
	"INSTANCE_PER_HOUR": ResourceBundleUnitOfMeasurementInstancePerHour,
	"CREDITS":           ResourceBundleUnitOfMeasurementCredits,
	"INSTANCES":         ResourceBundleUnitOfMeasurementInstances,
	"NODES":             ResourceBundleUnitOfMeasurementNodes,
}

var mappingResourceBundleUnitOfMeasurementEnumLowerCase = map[string]ResourceBundleUnitOfMeasurementEnum{
	"ocpu_per_hour":     ResourceBundleUnitOfMeasurementOcpuPerHour,
	"instance_per_hour": ResourceBundleUnitOfMeasurementInstancePerHour,
	"credits":           ResourceBundleUnitOfMeasurementCredits,
	"instances":         ResourceBundleUnitOfMeasurementInstances,
	"nodes":             ResourceBundleUnitOfMeasurementNodes,
}

// GetResourceBundleUnitOfMeasurementEnumValues Enumerates the set of values for ResourceBundleUnitOfMeasurementEnum
func GetResourceBundleUnitOfMeasurementEnumValues() []ResourceBundleUnitOfMeasurementEnum {
	values := make([]ResourceBundleUnitOfMeasurementEnum, 0)
	for _, v := range mappingResourceBundleUnitOfMeasurementEnum {
		values = append(values, v)
	}
	return values
}

// GetResourceBundleUnitOfMeasurementEnumStringValues Enumerates the set of values in String for ResourceBundleUnitOfMeasurementEnum
func GetResourceBundleUnitOfMeasurementEnumStringValues() []string {
	return []string{
		"OCPU_PER_HOUR",
		"INSTANCE_PER_HOUR",
		"CREDITS",
		"INSTANCES",
		"NODES",
	}
}

// GetMappingResourceBundleUnitOfMeasurementEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingResourceBundleUnitOfMeasurementEnum(val string) (ResourceBundleUnitOfMeasurementEnum, bool) {
	enum, ok := mappingResourceBundleUnitOfMeasurementEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
