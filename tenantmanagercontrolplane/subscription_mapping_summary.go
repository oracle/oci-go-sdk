// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Organizations API
//
// The Organizations API allows you to consolidate multiple OCI tenancies into an organization, and centrally manage your tenancies and its resources.
//

package tenantmanagercontrolplane

import (
	"github.com/oracle/oci-go-sdk/v52/common"
)

// SubscriptionMappingSummary Subscription mapping information.
type SubscriptionMappingSummary struct {

	// OCID of the mapping between subscription and compartment identified by the tenancy.
	Id *string `mandatory:"true" json:"id"`

	// OCID of the subscription.
	SubscriptionId *string `mandatory:"true" json:"subscriptionId"`

	// OCID of the compartment. Always a tenancy OCID.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// Denotes if the subscription is explicity assigned to the root compartment or tenancy.
	IsExplicitlyAssigned *bool `mandatory:"true" json:"isExplicitlyAssigned"`

	// Lifecycle state of the subscription mapping.
	LifecycleState SubscriptionMappingLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// Date-time when subscription mapping was created.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// Date-time when subscription mapping was updated.
	TimeUpdated *common.SDKTime `mandatory:"true" json:"timeUpdated"`

	// Date-time when subscription mapping was terminated.
	TimeTerminated *common.SDKTime `mandatory:"false" json:"timeTerminated"`
}

func (m SubscriptionMappingSummary) String() string {
	return common.PointerString(m)
}
