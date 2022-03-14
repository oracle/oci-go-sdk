// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Support Management API
//
// Use the Support Management API to manage support requests. For more information, see Getting Help and Contacting Support (https://docs.cloud.oracle.com/iaas/Content/GSG/Tasks/contactingsupport.htm). **Note**: Before you can create service requests with this API, you need to have an Oracle Single Sign On (SSO) account, and you need to register your Customer Support Identifier (CSI) with My Oracle Support.
//

package cims

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v62/common"
	"strings"
)

// Classifier Details about the incident classifier object.
type Classifier struct {

	// Unique identifier of the classifier.
	Id *string `mandatory:"false" json:"id"`

	// The display name of the classifier.
	Name *string `mandatory:"false" json:"name"`

	// The label associated with the classifier.
	Label *string `mandatory:"false" json:"label"`

	// The description of the classifier.
	Description *string `mandatory:"false" json:"description"`

	// The list of issues.
	IssueTypeList []IssueType `mandatory:"false" json:"issueTypeList"`

	// The scope of the service category or resource.
	Scope ClassifierScopeEnum `mandatory:"false" json:"scope,omitempty"`

	// The unit to use to measure the service category or resource.
	Unit ClassifierUnitEnum `mandatory:"false" json:"unit,omitempty"`
}

func (m Classifier) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m Classifier) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingClassifierScopeEnum(string(m.Scope)); !ok && m.Scope != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Scope: %s. Supported values are: %s.", m.Scope, strings.Join(GetClassifierScopeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingClassifierUnitEnum(string(m.Unit)); !ok && m.Unit != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Unit: %s. Supported values are: %s.", m.Unit, strings.Join(GetClassifierUnitEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ClassifierScopeEnum Enum with underlying type: string
type ClassifierScopeEnum string

// Set of constants representing the allowable values for ClassifierScopeEnum
const (
	ClassifierScopeAd      ClassifierScopeEnum = "AD"
	ClassifierScopeRegion  ClassifierScopeEnum = "REGION"
	ClassifierScopeTenancy ClassifierScopeEnum = "TENANCY"
	ClassifierScopeNone    ClassifierScopeEnum = "NONE"
)

var mappingClassifierScopeEnum = map[string]ClassifierScopeEnum{
	"AD":      ClassifierScopeAd,
	"REGION":  ClassifierScopeRegion,
	"TENANCY": ClassifierScopeTenancy,
	"NONE":    ClassifierScopeNone,
}

var mappingClassifierScopeEnumLowerCase = map[string]ClassifierScopeEnum{
	"ad":      ClassifierScopeAd,
	"region":  ClassifierScopeRegion,
	"tenancy": ClassifierScopeTenancy,
	"none":    ClassifierScopeNone,
}

// GetClassifierScopeEnumValues Enumerates the set of values for ClassifierScopeEnum
func GetClassifierScopeEnumValues() []ClassifierScopeEnum {
	values := make([]ClassifierScopeEnum, 0)
	for _, v := range mappingClassifierScopeEnum {
		values = append(values, v)
	}
	return values
}

// GetClassifierScopeEnumStringValues Enumerates the set of values in String for ClassifierScopeEnum
func GetClassifierScopeEnumStringValues() []string {
	return []string{
		"AD",
		"REGION",
		"TENANCY",
		"NONE",
	}
}

// GetMappingClassifierScopeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingClassifierScopeEnum(val string) (ClassifierScopeEnum, bool) {
	enum, ok := mappingClassifierScopeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ClassifierUnitEnum Enum with underlying type: string
type ClassifierUnitEnum string

// Set of constants representing the allowable values for ClassifierUnitEnum
const (
	ClassifierUnitCount ClassifierUnitEnum = "COUNT"
	ClassifierUnitGb    ClassifierUnitEnum = "GB"
	ClassifierUnitNone  ClassifierUnitEnum = "NONE"
)

var mappingClassifierUnitEnum = map[string]ClassifierUnitEnum{
	"COUNT": ClassifierUnitCount,
	"GB":    ClassifierUnitGb,
	"NONE":  ClassifierUnitNone,
}

var mappingClassifierUnitEnumLowerCase = map[string]ClassifierUnitEnum{
	"count": ClassifierUnitCount,
	"gb":    ClassifierUnitGb,
	"none":  ClassifierUnitNone,
}

// GetClassifierUnitEnumValues Enumerates the set of values for ClassifierUnitEnum
func GetClassifierUnitEnumValues() []ClassifierUnitEnum {
	values := make([]ClassifierUnitEnum, 0)
	for _, v := range mappingClassifierUnitEnum {
		values = append(values, v)
	}
	return values
}

// GetClassifierUnitEnumStringValues Enumerates the set of values in String for ClassifierUnitEnum
func GetClassifierUnitEnumStringValues() []string {
	return []string{
		"COUNT",
		"GB",
		"NONE",
	}
}

// GetMappingClassifierUnitEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingClassifierUnitEnum(val string) (ClassifierUnitEnum, bool) {
	enum, ok := mappingClassifierUnitEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
