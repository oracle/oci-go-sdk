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

// ServiceCategory Incident Classifier details
type ServiceCategory struct {

	// Unique ID that identifies a classifier
	Key *string `mandatory:"false" json:"key"`

	// Name of classifier. eg: LIMIT Increase
	Name *string `mandatory:"false" json:"name"`

	// Label of classifier
	Label *string `mandatory:"false" json:"label"`

	// Description of classifier
	Description *string `mandatory:"false" json:"description"`

	// List of Issues
	IssueTypeList []IssueType `mandatory:"false" json:"issueTypeList"`

	// List of Scope
	Scope ScopeEnum `mandatory:"false" json:"scope,omitempty"`

	// List of Units
	Unit UnitEnum `mandatory:"false" json:"unit,omitempty"`

	// Limit's unique id
	LimitId *string `mandatory:"false" json:"limitId"`
}

func (m ServiceCategory) String() string {
	return common.PointerString(m)
}
