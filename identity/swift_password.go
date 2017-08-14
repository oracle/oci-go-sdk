// Copyright (c) 2016, 2017, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Identity and Access Management Service API
//
// APIs for managing users, groups, compartments, and policies.
//

package identity

import (
        "time"
)


//SwiftPassword Swift is the OpenStack object storage service. A `SwiftPassword` is an Oracle-provided password for using a\nSwift client with the Oracle Bare Metal Cloud Services Object Storage Service. This password is associated with\nthe user's Console login. Swift passwords never expire. A user can have up to two Swift passwords at a time.\n\n**Note:** The password is always an Oracle-generated string; you can't change it to a string of your choice.\n\nFor more information, see [Managing User Credentials](/Content/Identity/Tasks/managingcredentials.htm).\n

type SwiftPassword struct {

    // The Swift password. The value is available only in the response for `CreateSwiftPassword`, and not\nfor `ListSwiftPasswords` or `UpdateSwiftPassword`.\n
    Password string `json:"password,omitempty"`

    // The OCID of the Swift password.
    Id string `json:"id,omitempty"`

    // The OCID of the user the password belongs to.
    UserId string `json:"userId,omitempty"`

    // The description you assign to the Swift password. Does not have to be unique, and it's changeable.
    Description string `json:"description,omitempty"`

    // Date and time the `SwiftPassword` object was created, in the format defined by RFC3339.\n\nExample: `2016-08-25T21:10:29.600Z`\n
    TimeCreated time.Time `json:"timeCreated,omitempty"`

    // Date and time when this password will expire, in the format defined by RFC3339.\nNull if it never expires.\n\nExample: `2016-08-25T21:10:29.600Z`\n
    ExpiresOn time.Time `json:"expiresOn,omitempty"`

    // The password's current state. After creating a password, make sure its `lifecycleState` changes from\nCREATING to ACTIVE before using it.\n
    LifecycleState string `json:"lifecycleState,omitempty"`

    // The detailed status of INACTIVE lifecycleState.
    InactiveStatus int64 `json:"inactiveStatus,omitempty"`
}
