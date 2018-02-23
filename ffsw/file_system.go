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


    
 // FileSystem A file system. A file system can be made accessible by adding it
 // to an export set and associating the export set with a mount
 // target. The same file system can be in multiple export sets and
 // associated with multiple mount targets.
 // To use any of the API operations, you must be authorized in an
 // IAM policy. If you are not authorized, talk to an
 // administrator. If you are an administrator who needs to write
 // policies to give users access, see Getting Started with
 // Policies ({{DOC_SERVER_URL}}/Content/Identity/Concepts/policygetstarted.htm).
type FileSystem struct {
    
 // The number of bytes consumed by the file system, including
 // snapshot. This number reflects the metered size of the file
 // system, and is updated asynchronously with respect to
 // updates to the file system. For details on file system
 // metering see File System Metering ({{DOC_SERVER_URL}}/Content/File/Concepts/metering.htm).
    MeteredBytes *int `mandatory:"true" json:"meteredBytes"`
    
 // The OCID of the compartment that contains the file system.
    CompartmentId *string `mandatory:"true" json:"compartmentId"`
    
 // A user-friendly name. Does not have to be unique, and it is changeable.
 // Avoid entering confidential information.
 // Example: `My file system`
    DisplayName *string `mandatory:"true" json:"displayName"`
    
 // The OCID of the file system.
    Id *string `mandatory:"true" json:"id"`
    
 // The current state of the file system.
    LifecycleState FileSystemLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`
    
 // The date and time the file system was created, in the format defined by RFC3339.
 // Example: `2016-08-25T21:10:29.600Z`
    TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`
    
 // The Availability Domain the file system is in. May be unset.
 // Example: `Uocm:PHX-AD-1`
    AvailabilityDomain *string `mandatory:"false" json:"availabilityDomain"`
}

func (m FileSystem) String() string {
    return common.PointerString(m)
}


// FileSystemLifecycleStateEnum Enum with underlying type: string
type FileSystemLifecycleStateEnum string

// Set of constants representing the allowable values for FileSystemLifecycleState
const (
    FileSystemLifecycleStateCreating FileSystemLifecycleStateEnum = "CREATING"
    FileSystemLifecycleStateActive FileSystemLifecycleStateEnum = "ACTIVE"
    FileSystemLifecycleStateDeleting FileSystemLifecycleStateEnum = "DELETING"
    FileSystemLifecycleStateDeleted FileSystemLifecycleStateEnum = "DELETED"
)

var mappingFileSystemLifecycleState = map[string]FileSystemLifecycleStateEnum { 
    "CREATING": FileSystemLifecycleStateCreating,
    "ACTIVE": FileSystemLifecycleStateActive,
    "DELETING": FileSystemLifecycleStateDeleting,
    "DELETED": FileSystemLifecycleStateDeleted,
}

// GetFileSystemLifecycleStateEnumValues Enumerates the set of values for FileSystemLifecycleState
func GetFileSystemLifecycleStateEnumValues() []FileSystemLifecycleStateEnum {
   values := make([]FileSystemLifecycleStateEnum, 0)
   for _, v := range mappingFileSystemLifecycleState {
       values = append(values, v)
   }
   return values
}



