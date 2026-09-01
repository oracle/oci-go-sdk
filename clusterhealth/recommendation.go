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

// Recommendation recommendation
type Recommendation struct {

	// Severity of the recommendation.
	Type RecommendationTypeEnum `mandatory:"true" json:"type"`

	// name of the test
	TestName *string `mandatory:"true" json:"testName"`

	// description of the issue
	Issue *string `mandatory:"true" json:"issue"`

	// suggested actions
	Suggestion *string `mandatory:"true" json:"suggestion"`

	// Optional action to take (empty string when no action is required).
	Action *string `mandatory:"true" json:"action"`

	// Optional fault code identifier.
	FaultCode *string `mandatory:"false" json:"faultCode"`
}

func (m Recommendation) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m Recommendation) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingRecommendationTypeEnum(string(m.Type)); !ok && m.Type != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Type: %s. Supported values are: %s.", m.Type, strings.Join(GetRecommendationTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// RecommendationTypeEnum Enum with underlying type: string
type RecommendationTypeEnum string

// Set of constants representing the allowable values for RecommendationTypeEnum
const (
	RecommendationTypeCritical RecommendationTypeEnum = "CRITICAL"
	RecommendationTypeWarning  RecommendationTypeEnum = "WARNING"
	RecommendationTypeInfo     RecommendationTypeEnum = "INFO"
)

var mappingRecommendationTypeEnum = map[string]RecommendationTypeEnum{
	"CRITICAL": RecommendationTypeCritical,
	"WARNING":  RecommendationTypeWarning,
	"INFO":     RecommendationTypeInfo,
}

var mappingRecommendationTypeEnumLowerCase = map[string]RecommendationTypeEnum{
	"critical": RecommendationTypeCritical,
	"warning":  RecommendationTypeWarning,
	"info":     RecommendationTypeInfo,
}

// GetRecommendationTypeEnumValues Enumerates the set of values for RecommendationTypeEnum
func GetRecommendationTypeEnumValues() []RecommendationTypeEnum {
	values := make([]RecommendationTypeEnum, 0)
	for _, v := range mappingRecommendationTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetRecommendationTypeEnumStringValues Enumerates the set of values in String for RecommendationTypeEnum
func GetRecommendationTypeEnumStringValues() []string {
	return []string{
		"CRITICAL",
		"WARNING",
		"INFO",
	}
}

// GetMappingRecommendationTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingRecommendationTypeEnum(val string) (RecommendationTypeEnum, bool) {
	enum, ok := mappingRecommendationTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
