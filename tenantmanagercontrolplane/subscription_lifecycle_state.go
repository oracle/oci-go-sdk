// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Organizations API
//
// The Organizations API allows you to consolidate multiple OCI tenancies into an organization, and centrally manage your tenancies and its resources.
//

package tenantmanagercontrolplane

// SubscriptionLifecycleStateEnum Enum with underlying type: string
type SubscriptionLifecycleStateEnum string

// Set of constants representing the allowable values for SubscriptionLifecycleStateEnum
const (
	SubscriptionLifecycleStateCreating SubscriptionLifecycleStateEnum = "CREATING"
	SubscriptionLifecycleStateActive   SubscriptionLifecycleStateEnum = "ACTIVE"
	SubscriptionLifecycleStateInactive SubscriptionLifecycleStateEnum = "INACTIVE"
	SubscriptionLifecycleStateUpdating SubscriptionLifecycleStateEnum = "UPDATING"
	SubscriptionLifecycleStateDeleting SubscriptionLifecycleStateEnum = "DELETING"
	SubscriptionLifecycleStateDeleted  SubscriptionLifecycleStateEnum = "DELETED"
	SubscriptionLifecycleStateFailed   SubscriptionLifecycleStateEnum = "FAILED"
)

var mappingSubscriptionLifecycleState = map[string]SubscriptionLifecycleStateEnum{
	"CREATING": SubscriptionLifecycleStateCreating,
	"ACTIVE":   SubscriptionLifecycleStateActive,
	"INACTIVE": SubscriptionLifecycleStateInactive,
	"UPDATING": SubscriptionLifecycleStateUpdating,
	"DELETING": SubscriptionLifecycleStateDeleting,
	"DELETED":  SubscriptionLifecycleStateDeleted,
	"FAILED":   SubscriptionLifecycleStateFailed,
}

// GetSubscriptionLifecycleStateEnumValues Enumerates the set of values for SubscriptionLifecycleStateEnum
func GetSubscriptionLifecycleStateEnumValues() []SubscriptionLifecycleStateEnum {
	values := make([]SubscriptionLifecycleStateEnum, 0)
	for _, v := range mappingSubscriptionLifecycleState {
		values = append(values, v)
	}
	return values
}
