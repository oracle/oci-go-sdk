// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Organizations API
//
// The Organizations API allows you to consolidate multiple OCI tenancies into an organization, and centrally manage your tenancies and its resources.
//

package tenantmanagercontrolplane

import (
	"github.com/oracle/oci-go-sdk/v55/common"
)

// OrganizationSummary An organization entity.
type OrganizationSummary struct {

	// OCID of the organization.
	Id *string `mandatory:"true" json:"id"`

	// OCID of the compartment containing the organization. Always a tenancy OCID.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// OCID of the default Annual Universal Credits subscription. Any tenancy joining the organization will automatically get assigned this subscription, if a subscription is not explictly assigned.
	DefaultUcmSubscriptionId *string `mandatory:"true" json:"defaultUcmSubscriptionId"`

	// Lifecycle state of the organization.
	LifecycleState OrganizationLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// Date-time when this organization was created.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// A display name for the organization.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// The name of the tenancy that is the organization parent.
	ParentName *string `mandatory:"false" json:"parentName"`

	// Date-time when this organization was last updated.
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`
}

func (m OrganizationSummary) String() string {
	return common.PointerString(m)
}
