// Copyright (c) 2016, 2017, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Identity and Access Management Service API
//
// APIs for managing users, groups, compartments, and policies.
//

package identity

import (
	"bitbucket.aka.lgl.grungy.us/golang-sdk2/common"
)

// CustomerSecretKey. A `CustomerSecretKey` is an Oracle-provided key for using the Object Storage Service's
// [Amazon S3 compatible API]({{DOC_SERVER_URL}}/Content/Object/Tasks/s3compatibleapi.htm).
// A user can have up to two secret keys at a time.
// **Note:** The secret key is always an Oracle-generated string; you can't change it to a string of your choice.
// For more information, see [Managing User Credentials]({{DOC_SERVER_URL}}/Content/Identity/Tasks/managingcredentials.htm).
type CustomerSecretKey struct {

	// The secret key.
	Key *string `mandatory:"false" json:"key,omitempty"`

	// The OCID of the secret key.
	ID *string `mandatory:"false" json:"id,omitempty"`

	// The OCID of the user the password belongs to.
	UserID *string `mandatory:"false" json:"userId,omitempty"`

	// The display name you assign to the secret key. Does not have to be unique, and it's changeable.
	DisplayName *string `mandatory:"false" json:"displayName,omitempty"`

	// Date and time the `CustomerSecretKey` object was created, in the format defined by RFC3339.
	// Example: `2016-08-25T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated,omitempty"`

	// Date and time when this password will expire, in the format defined by RFC3339.
	// Null if it never expires.
	// Example: `2016-08-25T21:10:29.600Z`
	TimeExpires *common.SDKTime `mandatory:"false" json:"timeExpires,omitempty"`

	// The secret key's current state. After creating a secret key, make sure its `lifecycleState` changes from
	// CREATING to ACTIVE before using it.
	LifecycleState CustomerSecretKeyLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	// The detailed status of INACTIVE lifecycleState.
	InactiveStatus *int `mandatory:"false" json:"inactiveStatus,omitempty"`
}

func (model CustomerSecretKey) String() string {
	return common.PointerString(model)
}

type CustomerSecretKeyLifecycleStateEnum string

const (
	CUSTOMER_SECRET_KEY_LIFECYCLE_STATE_CREATING CustomerSecretKeyLifecycleStateEnum = "CREATING"
	CUSTOMER_SECRET_KEY_LIFECYCLE_STATE_ACTIVE   CustomerSecretKeyLifecycleStateEnum = "ACTIVE"
	CUSTOMER_SECRET_KEY_LIFECYCLE_STATE_INACTIVE CustomerSecretKeyLifecycleStateEnum = "INACTIVE"
	CUSTOMER_SECRET_KEY_LIFECYCLE_STATE_DELETING CustomerSecretKeyLifecycleStateEnum = "DELETING"
	CUSTOMER_SECRET_KEY_LIFECYCLE_STATE_DELETED  CustomerSecretKeyLifecycleStateEnum = "DELETED"
	CUSTOMER_SECRET_KEY_LIFECYCLE_STATE_UNKNOWN  CustomerSecretKeyLifecycleStateEnum = "UNKNOWN"
)

var mapping_customersecretkey_lifecycleState = map[string]CustomerSecretKeyLifecycleStateEnum{
	"CREATING": CUSTOMER_SECRET_KEY_LIFECYCLE_STATE_CREATING,
	"ACTIVE":   CUSTOMER_SECRET_KEY_LIFECYCLE_STATE_ACTIVE,
	"INACTIVE": CUSTOMER_SECRET_KEY_LIFECYCLE_STATE_INACTIVE,
	"DELETING": CUSTOMER_SECRET_KEY_LIFECYCLE_STATE_DELETING,
	"DELETED":  CUSTOMER_SECRET_KEY_LIFECYCLE_STATE_DELETED,
	"UNKNOWN":  CUSTOMER_SECRET_KEY_LIFECYCLE_STATE_UNKNOWN,
}

func GetCustomerSecretKeyLifecycleStateEnumValues() []CustomerSecretKeyLifecycleStateEnum {
	values := make([]CustomerSecretKeyLifecycleStateEnum, 0)
	for _, v := range mapping_customersecretkey_lifecycleState {
		if v != CUSTOMER_SECRET_KEY_LIFECYCLE_STATE_UNKNOWN {
			values = append(values, v)
		}
	}
	return values
}
