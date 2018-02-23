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


    
 // UpdateFileSystemDetails The representation of UpdateFileSystemDetails
type UpdateFileSystemDetails struct {
    
 // A user-friendly name. Does not have to be unique, and it is changeable.
 // Avoid entering confidential information.
 // Example: `My file system`
    DisplayName *string `mandatory:"false" json:"displayName"`
}

func (m UpdateFileSystemDetails) String() string {
    return common.PointerString(m)
}





