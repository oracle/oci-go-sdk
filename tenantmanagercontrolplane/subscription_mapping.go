// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Organizations API
//
// The Organizations API allows you to consolidate multiple OCI tenancies into an organization, and centrally manage your tenancies and its resources.
//

package tenantmanagercontrolplane

import (
	"github.com/oracle/oci-go-sdk/v54/common"
)

// SubscriptionMapping Subscription mapping information.
type SubscriptionMapping struct {

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

func (m SubscriptionMapping) String() string {
	return common.PointerString(m)
}

// SubscriptionMappingLifecycleStateEnum Enum with underlying type: string
type SubscriptionMappingLifecycleStateEnum string

// Set of constants representing the allowable values for SubscriptionMappingLifecycleStateEnum
const (
	SubscriptionMappingLifecycleStateCreating SubscriptionMappingLifecycleStateEnum = "CREATING"
	SubscriptionMappingLifecycleStateActive   SubscriptionMappingLifecycleStateEnum = "ACTIVE"
	SubscriptionMappingLifecycleStateInactive SubscriptionMappingLifecycleStateEnum = "INACTIVE"
	SubscriptionMappingLifecycleStateUpdating SubscriptionMappingLifecycleStateEnum = "UPDATING"
	SubscriptionMappingLifecycleStateDeleting SubscriptionMappingLifecycleStateEnum = "DELETING"
	SubscriptionMappingLifecycleStateDeleted  SubscriptionMappingLifecycleStateEnum = "DELETED"
	SubscriptionMappingLifecycleStateFailed   SubscriptionMappingLifecycleStateEnum = "FAILED"
)

var mappingSubscriptionMappingLifecycleState = map[string]SubscriptionMappingLifecycleStateEnum{
	"CREATING": SubscriptionMappingLifecycleStateCreating,
	"ACTIVE":   SubscriptionMappingLifecycleStateActive,
	"INACTIVE": SubscriptionMappingLifecycleStateInactive,
	"UPDATING": SubscriptionMappingLifecycleStateUpdating,
	"DELETING": SubscriptionMappingLifecycleStateDeleting,
	"DELETED":  SubscriptionMappingLifecycleStateDeleted,
	"FAILED":   SubscriptionMappingLifecycleStateFailed,
}

// GetSubscriptionMappingLifecycleStateEnumValues Enumerates the set of values for SubscriptionMappingLifecycleStateEnum
func GetSubscriptionMappingLifecycleStateEnumValues() []SubscriptionMappingLifecycleStateEnum {
	values := make([]SubscriptionMappingLifecycleStateEnum, 0)
	for _, v := range mappingSubscriptionMappingLifecycleState {
		values = append(values, v)
	}
	return values
}
