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


    
 // UpdateExportSetDetails The representation of UpdateExportSetDetails
type UpdateExportSetDetails struct {
    
 // A user-friendly name. Does not have to be unique, and it is changeable.
 // Avoid entering confidential information.
 // Example: `My export set`
    DisplayName *string `mandatory:"false" json:"displayName"`
    
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

func (m UpdateExportSetDetails) String() string {
    return common.PointerString(m)
}





