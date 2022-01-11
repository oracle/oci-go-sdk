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

// Subscription Subscription information for compartmentId. Only root compartments are allowed.
type Subscription struct {

	// OCID of the subscription details for the particular root compartment or tenancy.
	Id *string `mandatory:"true" json:"id"`

	// Classic subscription ID.
	ClassicSubscriptionId *string `mandatory:"true" json:"classicSubscriptionId"`

	// OCID of the compartment. Always a tenancy OCID.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The type of subscription, such as 'CLOUDCM', 'SAAS', 'ERP', or 'CRM'.
	ServiceName *string `mandatory:"true" json:"serviceName"`

	// Denotes if the subscription is from classic systems or not.
	IsClassicSubscription *bool `mandatory:"false" json:"isClassicSubscription"`

	// The pay model of the subscription, such as 'Pay as you go' or 'Monthly'.
	PaymentModel *string `mandatory:"false" json:"paymentModel"`

	// Region for the subscription.
	RegionAssignment *string `mandatory:"false" json:"regionAssignment"`

	// Lifecycle state of the subscription.
	LifecycleState SubscriptionLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	// List of SKUs linked to this subscription.
	Skus []SubscriptionSku `mandatory:"false" json:"skus"`

	// Subscription start time.
	StartDate *common.SDKTime `mandatory:"false" json:"startDate"`

	// Subscription end time.
	EndDate *common.SDKTime `mandatory:"false" json:"endDate"`

	// Date-time when subscription is updated.
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// Date-time when subscription is created.
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`
}

func (m Subscription) String() string {
	return common.PointerString(m)
}
