// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Support Management API
//
// Use the Support Management API to manage support requests. For more information, see Getting Help and Contacting Support (https://docs.cloud.oracle.com/iaas/Content/GSG/Tasks/contactingsupport.htm). **Note**: Before you can create service requests with this API, you need to have an Oracle Single Sign On (SSO) account, and you need to register your Customer Support Identifier (CSI) with My Oracle Support.
//

package cims

import (
	"github.com/oracle/oci-go-sdk/common"
)

// Classifier Incident Classifier details
type Classifier struct {

	// Unique ID that identifies a classifier
	Id *string `mandatory:"false" json:"id"`

	// Name of classifier. eg: LIMIT Increase
	Name *string `mandatory:"false" json:"name"`

	// Label of classifier
	Label *string `mandatory:"false" json:"label"`

	// Description of classifier
	Description *string `mandatory:"false" json:"description"`

	// List of Issues
	IssueTypeList []IssueType `mandatory:"false" json:"issueTypeList"`

	// Scope of Service category/resource
	Scope ClassifierScopeEnum `mandatory:"false" json:"scope,omitempty"`

	// Unit to measure Service category/ resource
	Unit ClassifierUnitEnum `mandatory:"false" json:"unit,omitempty"`
}

func (m Classifier) String() string {
	return common.PointerString(m)
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

var mappingClassifierScope = map[string]ClassifierScopeEnum{
	"AD":      ClassifierScopeAd,
	"REGION":  ClassifierScopeRegion,
	"TENANCY": ClassifierScopeTenancy,
	"NONE":    ClassifierScopeNone,
}

// GetClassifierScopeEnumValues Enumerates the set of values for ClassifierScopeEnum
func GetClassifierScopeEnumValues() []ClassifierScopeEnum {
	values := make([]ClassifierScopeEnum, 0)
	for _, v := range mappingClassifierScope {
		values = append(values, v)
	}
	return values
}

// ClassifierUnitEnum Enum with underlying type: string
type ClassifierUnitEnum string

// Set of constants representing the allowable values for ClassifierUnitEnum
const (
	ClassifierUnitCount ClassifierUnitEnum = "COUNT"
	ClassifierUnitGb    ClassifierUnitEnum = "GB"
	ClassifierUnitNone  ClassifierUnitEnum = "NONE"
)

var mappingClassifierUnit = map[string]ClassifierUnitEnum{
	"COUNT": ClassifierUnitCount,
	"GB":    ClassifierUnitGb,
	"NONE":  ClassifierUnitNone,
}

// GetClassifierUnitEnumValues Enumerates the set of values for ClassifierUnitEnum
func GetClassifierUnitEnumValues() []ClassifierUnitEnum {
	values := make([]ClassifierUnitEnum, 0)
	for _, v := range mappingClassifierUnit {
		values = append(values, v)
	}
	return values
}
