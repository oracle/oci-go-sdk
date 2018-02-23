// Copyright (c) 2016, 2018, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// File Storage Service API
//
// APIs for OCI file storage service.
//

package ffsw

import (
    "github.com/oracle/oci-go-sdk/common"
)


    
 // MountTarget Provides access to a collection of file systems via VNIC(s) on a
 // specified subnet. The set of file system is controlled via the
 // referenced export set.
type MountTarget struct {
    
 // The OCID of the compartment that contains the mount target.
    CompartmentId *string `mandatory:"true" json:"compartmentId"`
    
 // A user-friendly name. Does not have to be unique, and it is changeable.
 // Avoid entering confidential information.
 // Example: `My mount target`
    DisplayName *string `mandatory:"true" json:"displayName"`
    
 // The OCID of the mount target.
    Id *string `mandatory:"true" json:"id"`
    
 // Additional information about the current lifecycleState.
    LifecycleDetails *string `mandatory:"true" json:"lifecycleDetails"`
    
 // The current state of the mount target.
    LifecycleState MountTargetLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`
    
 // The OCIDs of the private IPs associated with this mount target.
    PrivateIpIds []string `mandatory:"true" json:"privateIpIds"`
    
 // The OCID of the subnet the mount target is in.
    SubnetId *string `mandatory:"true" json:"subnetId"`
    
 // The date and time the mount target was created, in the format defined by RFC3339.
 // Example: `2016-08-25T21:10:29.600Z`
    TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`
    
 // The Availability Domain the mount target is in. May be unset.
 // Example: `Uocm:PHX-AD-1`
    AvailabilityDomain *string `mandatory:"false" json:"availabilityDomain"`
    
 // The OCID of the associated export set. Controls what file
 // systems will be exported via NFS on this mount target.
    ExportSetId *string `mandatory:"false" json:"exportSetId"`
}

func (m MountTarget) String() string {
    return common.PointerString(m)
}


// MountTargetLifecycleStateEnum Enum with underlying type: string
type MountTargetLifecycleStateEnum string

// Set of constants representing the allowable values for MountTargetLifecycleState
const (
    MountTargetLifecycleStateCreating MountTargetLifecycleStateEnum = "CREATING"
    MountTargetLifecycleStateActive MountTargetLifecycleStateEnum = "ACTIVE"
    MountTargetLifecycleStateDeleting MountTargetLifecycleStateEnum = "DELETING"
    MountTargetLifecycleStateDeleted MountTargetLifecycleStateEnum = "DELETED"
    MountTargetLifecycleStateFailed MountTargetLifecycleStateEnum = "FAILED"
)

var mappingMountTargetLifecycleState = map[string]MountTargetLifecycleStateEnum { 
    "CREATING": MountTargetLifecycleStateCreating,
    "ACTIVE": MountTargetLifecycleStateActive,
    "DELETING": MountTargetLifecycleStateDeleting,
    "DELETED": MountTargetLifecycleStateDeleted,
    "FAILED": MountTargetLifecycleStateFailed,
}

// GetMountTargetLifecycleStateEnumValues Enumerates the set of values for MountTargetLifecycleState
func GetMountTargetLifecycleStateEnumValues() []MountTargetLifecycleStateEnum {
   values := make([]MountTargetLifecycleStateEnum, 0)
   for _, v := range mappingMountTargetLifecycleState {
       values = append(values, v)
   }
   return values
}



