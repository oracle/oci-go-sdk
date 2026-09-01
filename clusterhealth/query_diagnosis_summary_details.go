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

// QueryDiagnosisSummaryDetails Query filters for retrieving diagnosis summaries.
type QueryDiagnosisSummaryDetails struct {

	// A list of instance OCIDs to match.
	InstanceIds []string `mandatory:"true" json:"instanceIds"`

	// The test type, PASSIVE or ACTIVE
	TestType QueryDiagnosisSummaryDetailsTestTypeEnum `mandatory:"false" json:"testType,omitempty"`

	// Start of the time range for tests run, in RFC 3339 format.
	TimeTestRunIntervalStart *common.SDKTime `mandatory:"false" json:"timeTestRunIntervalStart"`

	// End of the time range for tests run, in RFC 3339 format.
	TimeTestRunIntervalEnd *common.SDKTime `mandatory:"false" json:"timeTestRunIntervalEnd"`
}

func (m QueryDiagnosisSummaryDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m QueryDiagnosisSummaryDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingQueryDiagnosisSummaryDetailsTestTypeEnum(string(m.TestType)); !ok && m.TestType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for TestType: %s. Supported values are: %s.", m.TestType, strings.Join(GetQueryDiagnosisSummaryDetailsTestTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// QueryDiagnosisSummaryDetailsTestTypeEnum Enum with underlying type: string
type QueryDiagnosisSummaryDetailsTestTypeEnum string

// Set of constants representing the allowable values for QueryDiagnosisSummaryDetailsTestTypeEnum
const (
	QueryDiagnosisSummaryDetailsTestTypePassive QueryDiagnosisSummaryDetailsTestTypeEnum = "PASSIVE"
	QueryDiagnosisSummaryDetailsTestTypeActive  QueryDiagnosisSummaryDetailsTestTypeEnum = "ACTIVE"
)

var mappingQueryDiagnosisSummaryDetailsTestTypeEnum = map[string]QueryDiagnosisSummaryDetailsTestTypeEnum{
	"PASSIVE": QueryDiagnosisSummaryDetailsTestTypePassive,
	"ACTIVE":  QueryDiagnosisSummaryDetailsTestTypeActive,
}

var mappingQueryDiagnosisSummaryDetailsTestTypeEnumLowerCase = map[string]QueryDiagnosisSummaryDetailsTestTypeEnum{
	"passive": QueryDiagnosisSummaryDetailsTestTypePassive,
	"active":  QueryDiagnosisSummaryDetailsTestTypeActive,
}

// GetQueryDiagnosisSummaryDetailsTestTypeEnumValues Enumerates the set of values for QueryDiagnosisSummaryDetailsTestTypeEnum
func GetQueryDiagnosisSummaryDetailsTestTypeEnumValues() []QueryDiagnosisSummaryDetailsTestTypeEnum {
	values := make([]QueryDiagnosisSummaryDetailsTestTypeEnum, 0)
	for _, v := range mappingQueryDiagnosisSummaryDetailsTestTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetQueryDiagnosisSummaryDetailsTestTypeEnumStringValues Enumerates the set of values in String for QueryDiagnosisSummaryDetailsTestTypeEnum
func GetQueryDiagnosisSummaryDetailsTestTypeEnumStringValues() []string {
	return []string{
		"PASSIVE",
		"ACTIVE",
	}
}

// GetMappingQueryDiagnosisSummaryDetailsTestTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingQueryDiagnosisSummaryDetailsTestTypeEnum(val string) (QueryDiagnosisSummaryDetailsTestTypeEnum, bool) {
	enum, ok := mappingQueryDiagnosisSummaryDetailsTestTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
