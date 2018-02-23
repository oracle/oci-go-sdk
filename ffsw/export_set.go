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


    
 // ExportSet A set of file systems to be exported via one (or more) mount
 // targets. Composed of zero or more export resource.
type ExportSet struct {
    
 // The OCID of the compartment that contains the export set.
    CompartmentId *string `mandatory:"true" json:"compartmentId"`
    
 // A user-friendly name. Does not have to be unique, and it is changeable.
 // Avoid entering confidential information.
 // Example: `My export set`
    DisplayName *string `mandatory:"true" json:"displayName"`
    
 // The OCID of the export set.
    Id *string `mandatory:"true" json:"id"`
    
 // The current state of the export set.
    LifecycleState ExportSetLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`
    
 // The date and time the export set was created, in the format defined by RFC3339.
 // Example: `2016-08-25T21:10:29.600Z`
    TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`
    
 // The OCID of the VCN the export set is in.
    VcnId *string `mandatory:"true" json:"vcnId"`
    
 // The Availability Domain the export set is in. May be unset.
 // Example: `Uocm:PHX-AD-1`
    AvailabilityDomain *string `mandatory:"false" json:"availabilityDomain"`
    
 // Controls the maximum `tbytes`, `fbytes`, and `abytes`,
 // values reported by `NFS FSSTAT` calls though any assocated
 // mount targets. This is an advanced feature, most
 // applications should leave this set to its default value. The
 // `tbytes` value reported by `FSSTAT` will be
 // `maxFsStatBytes`. `fbytes` and `abytes` value will be
 // `maxFsStatBytes` minus the metered size of the file
 // system. If the metered size is larger than `maxFsStatBytes`
 // then `fbytes` and `abytes` will both be 0.
    MaxFsStatBytes *int `mandatory:"false" json:"maxFsStatBytes"`
    
 // Controls the maximum `ffiles`, `ffiles`, and `afiles`,
 // values reported by `NFS FSSTAT` calls though any assocated
 // mount targets. This is an advanced feature, most
 // applications should leave this set to its default value. The
 // `tfiles` value reported by `FSSTAT` will be
 // `maxFsStatFiles`. `ffiles` and `afiles` value will be
 // `maxFsStatFiles` minus the metered size of the file
 // system. If the metered size is larger than `maxFsStatFiles`
 // then `ffiles` and `afiles` will both be 0.
    MaxFsStatFiles *int `mandatory:"false" json:"maxFsStatFiles"`
}

func (m ExportSet) String() string {
    return common.PointerString(m)
}


// ExportSetLifecycleStateEnum Enum with underlying type: string
type ExportSetLifecycleStateEnum string

// Set of constants representing the allowable values for ExportSetLifecycleState
const (
    ExportSetLifecycleStateCreating ExportSetLifecycleStateEnum = "CREATING"
    ExportSetLifecycleStateActive ExportSetLifecycleStateEnum = "ACTIVE"
    ExportSetLifecycleStateDeleting ExportSetLifecycleStateEnum = "DELETING"
    ExportSetLifecycleStateDeleted ExportSetLifecycleStateEnum = "DELETED"
)

var mappingExportSetLifecycleState = map[string]ExportSetLifecycleStateEnum { 
    "CREATING": ExportSetLifecycleStateCreating,
    "ACTIVE": ExportSetLifecycleStateActive,
    "DELETING": ExportSetLifecycleStateDeleting,
    "DELETED": ExportSetLifecycleStateDeleted,
}

// GetExportSetLifecycleStateEnumValues Enumerates the set of values for ExportSetLifecycleState
func GetExportSetLifecycleStateEnumValues() []ExportSetLifecycleStateEnum {
   values := make([]ExportSetLifecycleStateEnum, 0)
   for _, v := range mappingExportSetLifecycleState {
       values = append(values, v)
   }
   return values
}



