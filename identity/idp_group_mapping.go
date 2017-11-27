// Copyright (c) 2016, 2017, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Identity and Access Management Service API
//
// APIs for managing users, groups, compartments, and policies.
//

package identity

import (
	"github.com/oracle/oci-go-sdk/common"
)

// IdpGroupMapping. A mapping between a single group defined by the identity provider (IdP) you're federating with
// and a single IAM Service Group in Oracle Cloud Infrastructure.
// For more information about group mappings and what they're for, see
// [Identity Providers and Federation]({{DOC_SERVER_URL}}/Content/Identity/Concepts/federation.htm).
// A given IdP group can be mapped to zero, one, or multiple IAM Service groups, and vice versa.
// But each `IdPGroupMapping` object is between only a single IdP group and IAM Service group.
// Each `IdPGroupMapping` object has its own OCID.
// **Note:** Any users who are in more than 50 IdP groups cannot be authenticated to use the Oracle
// Cloud Infrastructure Console.
type IdpGroupMapping struct {

	// The OCID of the `IdpGroupMapping`.
	ID *string `mandatory:"true" json:"id,omitempty"`

	// The OCID of the `IdentityProvider` this mapping belongs to.
	IdpID *string `mandatory:"true" json:"idpId,omitempty"`

	// The name of the IdP group that is mapped to the IAM Service group.
	IdpGroupName *string `mandatory:"true" json:"idpGroupName,omitempty"`

	// The OCID of the IAM Service group that is mapped to the IdP group.
	GroupID *string `mandatory:"true" json:"groupId,omitempty"`

	// The OCID of the tenancy containing the `IdentityProvider`.
	CompartmentID *string `mandatory:"true" json:"compartmentId,omitempty"`

	// Date and time the mapping was created, in the format defined by RFC3339.
	// Example: `2016-08-25T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated,omitempty"`

	// The mapping's current state.  After creating a mapping object, make sure its `lifecycleState` changes
	// from CREATING to ACTIVE before using it.
	LifecycleState IdpGroupMappingLifecycleStateEnum `mandatory:"true" json:"lifecycleState,omitempty"`

	// The detailed status of INACTIVE lifecycleState.
	InactiveStatus *int `mandatory:"false" json:"inactiveStatus,omitempty"`
}

func (model IdpGroupMapping) String() string {
	return common.PointerString(model)
}

type IdpGroupMappingLifecycleStateEnum string

const (
	IDP_GROUP_MAPPING_LIFECYCLE_STATE_CREATING IdpGroupMappingLifecycleStateEnum = "CREATING"
	IDP_GROUP_MAPPING_LIFECYCLE_STATE_ACTIVE   IdpGroupMappingLifecycleStateEnum = "ACTIVE"
	IDP_GROUP_MAPPING_LIFECYCLE_STATE_INACTIVE IdpGroupMappingLifecycleStateEnum = "INACTIVE"
	IDP_GROUP_MAPPING_LIFECYCLE_STATE_DELETING IdpGroupMappingLifecycleStateEnum = "DELETING"
	IDP_GROUP_MAPPING_LIFECYCLE_STATE_DELETED  IdpGroupMappingLifecycleStateEnum = "DELETED"
	IDP_GROUP_MAPPING_LIFECYCLE_STATE_UNKNOWN  IdpGroupMappingLifecycleStateEnum = "UNKNOWN"
)

var mapping_idpgroupmapping_lifecycleState = map[string]IdpGroupMappingLifecycleStateEnum{
	"CREATING": IDP_GROUP_MAPPING_LIFECYCLE_STATE_CREATING,
	"ACTIVE":   IDP_GROUP_MAPPING_LIFECYCLE_STATE_ACTIVE,
	"INACTIVE": IDP_GROUP_MAPPING_LIFECYCLE_STATE_INACTIVE,
	"DELETING": IDP_GROUP_MAPPING_LIFECYCLE_STATE_DELETING,
	"DELETED":  IDP_GROUP_MAPPING_LIFECYCLE_STATE_DELETED,
	"UNKNOWN":  IDP_GROUP_MAPPING_LIFECYCLE_STATE_UNKNOWN,
}

func GetIdpGroupMappingLifecycleStateEnumValues() []IdpGroupMappingLifecycleStateEnum {
	values := make([]IdpGroupMappingLifecycleStateEnum, 0)
	for _, v := range mapping_idpgroupmapping_lifecycleState {
		if v != IDP_GROUP_MAPPING_LIFECYCLE_STATE_UNKNOWN {
			values = append(values, v)
		}
	}
	return values
}
