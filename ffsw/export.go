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


    
 // Export A file system and the path can be used to mount it. Each export
 // resource belongs to exacty one export set.
 // The export's path attribute of the export is not a path in the
 // refenced file system, but the value used by clients for the path
 // component of the remotetarget arugment when mounting the file
 // system.
 // The path must start with a / followed by a sequence of 0 or more
 // /-separated path elements. For any two non-DELETED Export
 // resources associated with the same ExportSet the path element
 // sequence for the first Export resource can't contain the
 // complete path element sequence of the second Export resource.
 // For example:
 // Good:
 //   /foo and /bar
 //   /foo1 and /foo2
 //   /foo and /foo1
 // Bad:
 //   /foo and /foo/bar
 //   / and /foo
 // Also note, paths may not end in /, no path element can be '.'
 // or '..', and no path element can be longer than 255 bytes.
 // No two non-DELETED export resources in the same export set can
 // reference the same file system.
type Export struct {
    
 // The OCID of this export's export set.
    ExportSetId *string `mandatory:"true" json:"exportSetId"`
    
 // The OCID of this export's file system.
    FileSystemId *string `mandatory:"true" json:"fileSystemId"`
    
 // The OCID of the export.
    Id *string `mandatory:"true" json:"id"`
    
 // The current state of the export.
    LifecycleState ExportLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`
    
 // Path used to access the associated file system.
 // Avoid entering confidential information.
 // Example: `/accounting`
    Path *string `mandatory:"true" json:"path"`
    
 // The date and time the export was created, in the format defined by RFC3339.
 // Example: `2016-08-25T21:10:29.600Z`
    TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`
}

func (m Export) String() string {
    return common.PointerString(m)
}


// ExportLifecycleStateEnum Enum with underlying type: string
type ExportLifecycleStateEnum string

// Set of constants representing the allowable values for ExportLifecycleState
const (
    ExportLifecycleStateCreating ExportLifecycleStateEnum = "CREATING"
    ExportLifecycleStateActive ExportLifecycleStateEnum = "ACTIVE"
    ExportLifecycleStateDeleting ExportLifecycleStateEnum = "DELETING"
    ExportLifecycleStateDeleted ExportLifecycleStateEnum = "DELETED"
)

var mappingExportLifecycleState = map[string]ExportLifecycleStateEnum { 
    "CREATING": ExportLifecycleStateCreating,
    "ACTIVE": ExportLifecycleStateActive,
    "DELETING": ExportLifecycleStateDeleting,
    "DELETED": ExportLifecycleStateDeleted,
}

// GetExportLifecycleStateEnumValues Enumerates the set of values for ExportLifecycleState
func GetExportLifecycleStateEnumValues() []ExportLifecycleStateEnum {
   values := make([]ExportLifecycleStateEnum, 0)
   for _, v := range mappingExportLifecycleState {
       values = append(values, v)
   }
   return values
}



