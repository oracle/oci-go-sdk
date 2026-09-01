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

// DiagnosisSummary Diagnosis summary of a host and tenancy.
type DiagnosisSummary struct {

	// Type of the diagnosis test recorded in this summary (e.g., PASSIVE or ACTIVE).
	TestType DiagnosisSummaryTestTypeEnum `mandatory:"true" json:"testType"`

	// id of the host
	HostId *string `mandatory:"false" json:"hostId"`

	// OCID of the instance associated with this summary.
	InstanceId *string `mandatory:"false" json:"instanceId"`

	// Identifier of the health check that produced this diagnosis summary.
	HealthCheckKey *string `mandatory:"false" json:"healthCheckKey"`

	// time the test was run
	TimeTestRan *common.SDKTime `mandatory:"false" json:"timeTestRan"`

	// time the recommendation report was generated
	TimeGenerated *common.SDKTime `mandatory:"false" json:"timeGenerated"`

	// Each summary contains no more than 50 failed tests
	FailedTests []FailedTest `mandatory:"false" json:"failedTests"`

	// if the host is healthy
	Health DiagnosisSummaryHealthEnum `mandatory:"false" json:"health,omitempty"`
}

func (m DiagnosisSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DiagnosisSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingDiagnosisSummaryTestTypeEnum(string(m.TestType)); !ok && m.TestType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for TestType: %s. Supported values are: %s.", m.TestType, strings.Join(GetDiagnosisSummaryTestTypeEnumStringValues(), ",")))
	}

	if _, ok := GetMappingDiagnosisSummaryHealthEnum(string(m.Health)); !ok && m.Health != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Health: %s. Supported values are: %s.", m.Health, strings.Join(GetDiagnosisSummaryHealthEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// DiagnosisSummaryHealthEnum Enum with underlying type: string
type DiagnosisSummaryHealthEnum string

// Set of constants representing the allowable values for DiagnosisSummaryHealthEnum
const (
	DiagnosisSummaryHealthHealthy   DiagnosisSummaryHealthEnum = "HEALTHY"
	DiagnosisSummaryHealthUnhealthy DiagnosisSummaryHealthEnum = "UNHEALTHY"
)

var mappingDiagnosisSummaryHealthEnum = map[string]DiagnosisSummaryHealthEnum{
	"HEALTHY":   DiagnosisSummaryHealthHealthy,
	"UNHEALTHY": DiagnosisSummaryHealthUnhealthy,
}

var mappingDiagnosisSummaryHealthEnumLowerCase = map[string]DiagnosisSummaryHealthEnum{
	"healthy":   DiagnosisSummaryHealthHealthy,
	"unhealthy": DiagnosisSummaryHealthUnhealthy,
}

// GetDiagnosisSummaryHealthEnumValues Enumerates the set of values for DiagnosisSummaryHealthEnum
func GetDiagnosisSummaryHealthEnumValues() []DiagnosisSummaryHealthEnum {
	values := make([]DiagnosisSummaryHealthEnum, 0)
	for _, v := range mappingDiagnosisSummaryHealthEnum {
		values = append(values, v)
	}
	return values
}

// GetDiagnosisSummaryHealthEnumStringValues Enumerates the set of values in String for DiagnosisSummaryHealthEnum
func GetDiagnosisSummaryHealthEnumStringValues() []string {
	return []string{
		"HEALTHY",
		"UNHEALTHY",
	}
}

// GetMappingDiagnosisSummaryHealthEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDiagnosisSummaryHealthEnum(val string) (DiagnosisSummaryHealthEnum, bool) {
	enum, ok := mappingDiagnosisSummaryHealthEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// DiagnosisSummaryTestTypeEnum Enum with underlying type: string
type DiagnosisSummaryTestTypeEnum string

// Set of constants representing the allowable values for DiagnosisSummaryTestTypeEnum
const (
	DiagnosisSummaryTestTypePassive DiagnosisSummaryTestTypeEnum = "PASSIVE"
	DiagnosisSummaryTestTypeActive  DiagnosisSummaryTestTypeEnum = "ACTIVE"
)

var mappingDiagnosisSummaryTestTypeEnum = map[string]DiagnosisSummaryTestTypeEnum{
	"PASSIVE": DiagnosisSummaryTestTypePassive,
	"ACTIVE":  DiagnosisSummaryTestTypeActive,
}

var mappingDiagnosisSummaryTestTypeEnumLowerCase = map[string]DiagnosisSummaryTestTypeEnum{
	"passive": DiagnosisSummaryTestTypePassive,
	"active":  DiagnosisSummaryTestTypeActive,
}

// GetDiagnosisSummaryTestTypeEnumValues Enumerates the set of values for DiagnosisSummaryTestTypeEnum
func GetDiagnosisSummaryTestTypeEnumValues() []DiagnosisSummaryTestTypeEnum {
	values := make([]DiagnosisSummaryTestTypeEnum, 0)
	for _, v := range mappingDiagnosisSummaryTestTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetDiagnosisSummaryTestTypeEnumStringValues Enumerates the set of values in String for DiagnosisSummaryTestTypeEnum
func GetDiagnosisSummaryTestTypeEnumStringValues() []string {
	return []string{
		"PASSIVE",
		"ACTIVE",
	}
}

// GetMappingDiagnosisSummaryTestTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDiagnosisSummaryTestTypeEnum(val string) (DiagnosisSummaryTestTypeEnum, bool) {
	enum, ok := mappingDiagnosisSummaryTestTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
