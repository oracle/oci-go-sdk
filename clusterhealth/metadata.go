// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// ITV Control Plane - Diagnosis Store API
//
// Use the ITV Control Plane Diagnosis Store API to manage diagnosis.
//

package clusterhealth

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// Metadata Metadata associated with the diagnosis recommendation report.
type Metadata struct {

	// OCID of the host that ran the test.
	HostId *string `mandatory:"false" json:"hostId"`

	// Identifier of the health check run.
	HealthCheckKey *string `mandatory:"false" json:"healthCheckKey"`

	// OCID of the instance related to the test.
	InstanceId *string `mandatory:"false" json:"instanceId"`

	// Shape of the instance under test.
	Shape *string `mandatory:"false" json:"shape"`

	// Test type provided by the diagnosis payload.
	TestType MetadataTestTypeEnum `mandatory:"false" json:"testType,omitempty"`
}

func (m Metadata) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m Metadata) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingMetadataTestTypeEnum(string(m.TestType)); !ok && m.TestType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for TestType: %s. Supported values are: %s.", m.TestType, strings.Join(GetMetadataTestTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MetadataTestTypeEnum Enum with underlying type: string
type MetadataTestTypeEnum string

// Set of constants representing the allowable values for MetadataTestTypeEnum
const (
	MetadataTestTypePassive MetadataTestTypeEnum = "PASSIVE"
	MetadataTestTypeActive  MetadataTestTypeEnum = "ACTIVE"
)

var mappingMetadataTestTypeEnum = map[string]MetadataTestTypeEnum{
	"PASSIVE": MetadataTestTypePassive,
	"ACTIVE":  MetadataTestTypeActive,
}

var mappingMetadataTestTypeEnumLowerCase = map[string]MetadataTestTypeEnum{
	"passive": MetadataTestTypePassive,
	"active":  MetadataTestTypeActive,
}

// GetMetadataTestTypeEnumValues Enumerates the set of values for MetadataTestTypeEnum
func GetMetadataTestTypeEnumValues() []MetadataTestTypeEnum {
	values := make([]MetadataTestTypeEnum, 0)
	for _, v := range mappingMetadataTestTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetMetadataTestTypeEnumStringValues Enumerates the set of values in String for MetadataTestTypeEnum
func GetMetadataTestTypeEnumStringValues() []string {
	return []string{
		"PASSIVE",
		"ACTIVE",
	}
}

// GetMappingMetadataTestTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingMetadataTestTypeEnum(val string) (MetadataTestTypeEnum, bool) {
	enum, ok := mappingMetadataTestTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
