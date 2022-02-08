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
	"github.com/oracle/oci-go-sdk/v57/common"
	"strings"
)

// ServiceCategory Information about the incident classifier.
type ServiceCategory struct {

	// The unique ID that identifies a classifier.
	Key *string `mandatory:"false" json:"key"`

	// The name of the classifier.
	Name *string `mandatory:"false" json:"name"`

	// The label for the classifier.
	Label *string `mandatory:"false" json:"label"`

	// The text describing the classifier.
	Description *string `mandatory:"false" json:"description"`

	// The list of issues.
	IssueTypeList []IssueType `mandatory:"false" json:"issueTypeList"`

	// The scope of the incident.
	Scope ScopeEnum `mandatory:"false" json:"scope,omitempty"`

	// The unit to use to measure the service category or resource.
	Unit UnitEnum `mandatory:"false" json:"unit,omitempty"`

	// The unique ID for the limit.
	LimitId *string `mandatory:"false" json:"limitId"`
}

func (m ServiceCategory) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ServiceCategory) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := mappingScopeEnum[string(m.Scope)]; !ok && m.Scope != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Scope: %s. Supported values are: %s.", m.Scope, strings.Join(GetScopeEnumStringValues(), ",")))
	}
	if _, ok := mappingUnitEnum[string(m.Unit)]; !ok && m.Unit != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Unit: %s. Supported values are: %s.", m.Unit, strings.Join(GetUnitEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
