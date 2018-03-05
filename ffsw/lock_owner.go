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


    
 // LockOwner Identifies a client that owns one or more locks in the given file system.
type LockOwner struct {
    
 // The OCID of this file system where the locks are held.
    FileSystemId *string `mandatory:"true" json:"fileSystemId"`
    
 // The IP address of the client hold the locks.
    ClientIpAddress *string `mandatory:"true" json:"clientIpAddress"`
    
 // The OCID of the mount target the locks are held through.
    MountTargetId *string `mandatory:"true" json:"mountTargetId"`
    
 // The OCID of the mount target's VCN.
    VcnId *string `mandatory:"true" json:"vcnId"`
}

func (m LockOwner) String() string {
    return common.PointerString(m)
}





