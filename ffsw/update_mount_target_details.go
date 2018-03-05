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


    
 // UpdateMountTargetDetails The representation of UpdateMountTargetDetails
type UpdateMountTargetDetails struct {
    
 // A user-friendly name. Does not have to be unique, and it is changeable.
 // Avoid entering confidential information.
 // Example: `My mount target`
    DisplayName *string `mandatory:"false" json:"displayName"`
}

func (m UpdateMountTargetDetails) String() string {
    return common.PointerString(m)
}





