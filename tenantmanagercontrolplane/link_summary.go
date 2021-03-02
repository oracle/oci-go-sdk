// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// TenantManager API
//
// A description of the TenantManager API
//

package tenantmanagercontrolplane

import (
	"github.com/oracle/oci-go-sdk/v36/common"
)

// LinkSummary The summary of a link between a parent tenancy and a child tenancy.
type LinkSummary struct {

	// OCID of the link.
	Id *string `mandatory:"true" json:"id"`

	// OCID of the parent tenancy.
	ParentTenancyId *string `mandatory:"true" json:"parentTenancyId"`

	// OCID of the child tenancy.
	ChildTenancyId *string `mandatory:"true" json:"childTenancyId"`

	// Date-time when this link was created
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// Lifecycle state of the link.
	LifecycleState LifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	// Date-time when this link was last updated.
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// Date-time when this link was terminated.
	TimeTerminated *common.SDKTime `mandatory:"false" json:"timeTerminated"`
}

func (m LinkSummary) String() string {
	return common.PointerString(m)
}
