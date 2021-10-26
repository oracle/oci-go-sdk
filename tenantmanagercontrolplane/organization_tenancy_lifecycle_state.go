// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Organizations API
//
// The Organizations API allows you to consolidate multiple OCI tenancies into an organization, and centrally manage your tenancies and its resources.
//

package tenantmanagercontrolplane

// OrganizationTenancyLifecycleStateEnum Enum with underlying type: string
type OrganizationTenancyLifecycleStateEnum string

// Set of constants representing the allowable values for OrganizationTenancyLifecycleStateEnum
const (
	OrganizationTenancyLifecycleStateCreating OrganizationTenancyLifecycleStateEnum = "CREATING"
	OrganizationTenancyLifecycleStateActive   OrganizationTenancyLifecycleStateEnum = "ACTIVE"
	OrganizationTenancyLifecycleStateInactive OrganizationTenancyLifecycleStateEnum = "INACTIVE"
	OrganizationTenancyLifecycleStateDeleted  OrganizationTenancyLifecycleStateEnum = "DELETED"
	OrganizationTenancyLifecycleStateFailed   OrganizationTenancyLifecycleStateEnum = "FAILED"
	OrganizationTenancyLifecycleStateDeleting OrganizationTenancyLifecycleStateEnum = "DELETING"
)

var mappingOrganizationTenancyLifecycleState = map[string]OrganizationTenancyLifecycleStateEnum{
	"CREATING": OrganizationTenancyLifecycleStateCreating,
	"ACTIVE":   OrganizationTenancyLifecycleStateActive,
	"INACTIVE": OrganizationTenancyLifecycleStateInactive,
	"DELETED":  OrganizationTenancyLifecycleStateDeleted,
	"FAILED":   OrganizationTenancyLifecycleStateFailed,
	"DELETING": OrganizationTenancyLifecycleStateDeleting,
}

// GetOrganizationTenancyLifecycleStateEnumValues Enumerates the set of values for OrganizationTenancyLifecycleStateEnum
func GetOrganizationTenancyLifecycleStateEnumValues() []OrganizationTenancyLifecycleStateEnum {
	values := make([]OrganizationTenancyLifecycleStateEnum, 0)
	for _, v := range mappingOrganizationTenancyLifecycleState {
		values = append(values, v)
	}
	return values
}
