// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Organizations API
//
// The Organizations API allows you to consolidate multiple OCI tenancies into an organization, and centrally manage your tenancies and its resources.
//

package tenantmanagercontrolplane

import (
	"github.com/oracle/oci-go-sdk/v51/common"
)

// AssignedSubscription Assigned subscription information.
type AssignedSubscription struct {

	// OCID of the subscription.
	Id *string `mandatory:"true" json:"id"`

	// OCID of the compartment. Always a tenancy OCID.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// Subscription id.
	ClassicSubscriptionId *string `mandatory:"true" json:"classicSubscriptionId"`

	// The type of subscription, such as 'CLOUDCM', 'SAAS', 'ERP', or 'CRM'.
	ServiceName *string `mandatory:"true" json:"serviceName"`

	// Denotes if the subscription is legacy or not.
	IsClassicSubscription *bool `mandatory:"false" json:"isClassicSubscription"`

	// Region for the subscription.
	RegionAssignment *string `mandatory:"false" json:"regionAssignment"`

	// Lifecycle state of the subscription.
	LifecycleState SubscriptionLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	// List of sku linked to the subscription.
	Skus []SubscriptionSku `mandatory:"false" json:"skus"`

	// List of subscription order OCIDs that contributed to this subscription.
	OrderIds []string `mandatory:"false" json:"orderIds"`

	// Subscription start time.
	StartDate *common.SDKTime `mandatory:"false" json:"startDate"`

	// Subscription end time.
	EndDate *common.SDKTime `mandatory:"false" json:"endDate"`

	// Date-time when subscription is updated.
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// Date-time when subscription is created.
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`
}

func (m AssignedSubscription) String() string {
	return common.PointerString(m)
}
