// Copyright (c) 2016, 2017, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Database Service API
//
// The API for the Database Service.
//

package database

import (
	"github.com/oracle/oci-go-sdk/common"
)

// DbHome A directory where Oracle database software is installed. Each DB System can have multiple database homes,
// and each database home can have multiple databases within it. All the databases within a single database home
// must be the same database version, but different database homes can run different versions. For more information,
// see https://docs.us-phoenix-1.oraclecloud.com/Content/Database/Concepts/overview.htm.
// To use any of the API operations, you must be authorized in an IAM policy. If you're not authorized, talk to an
// administrator. If you're an administrator who needs to write policies to give users access,
// see https://docs.us-phoenix-1.oraclecloud.com/Content/Identity/Concepts/policygetstarted.htm.
type DbHome struct {

	// The OCID of the compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId,omitempty"`

	// The Oracle database version.
	DbVersion *string `mandatory:"true" json:"dbVersion,omitempty"`

	// The user-provided name for the database home. It does not need to be unique.
	DisplayName *string `mandatory:"true" json:"displayName,omitempty"`

	// The OCID of the database home.
	Id *string `mandatory:"true" json:"id,omitempty"`

	// The current state of the database home.
	LifecycleState DbHomeLifecycleStateEnum `mandatory:"true" json:"lifecycleState,omitempty"`

	// The OCID of the DB System.
	DbSystemId *string `mandatory:"false" json:"dbSystemId,omitempty"`

	// The OCID of the last patch history. This is updated as soon as a patch operation is started.
	LastPatchHistoryEntryId *string `mandatory:"false" json:"lastPatchHistoryEntryId,omitempty"`

	// The date and time the database home was created.
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated,omitempty"`
}

func (m DbHome) String() string {
	return common.PointerString(m)
}

// DbHomeLifecycleStateEnum Enum with underlying type: string
type DbHomeLifecycleStateEnum string

// Set of constants representing the allowable values for DbHomeLifecycleState
const (
	DbHomeLifecycleStateProvisioning DbHomeLifecycleStateEnum = "PROVISIONING"
	DbHomeLifecycleStateAvailable    DbHomeLifecycleStateEnum = "AVAILABLE"
	DbHomeLifecycleStateUpdating     DbHomeLifecycleStateEnum = "UPDATING"
	DbHomeLifecycleStateTerminating  DbHomeLifecycleStateEnum = "TERMINATING"
	DbHomeLifecycleStateTerminated   DbHomeLifecycleStateEnum = "TERMINATED"
	DbHomeLifecycleStateFailed       DbHomeLifecycleStateEnum = "FAILED"
	DbHomeLifecycleStateUnknown      DbHomeLifecycleStateEnum = "UNKNOWN"
)

var mappingDbHomeLifecycleState = map[string]DbHomeLifecycleStateEnum{
	"PROVISIONING": DbHomeLifecycleStateProvisioning,
	"AVAILABLE":    DbHomeLifecycleStateAvailable,
	"UPDATING":     DbHomeLifecycleStateUpdating,
	"TERMINATING":  DbHomeLifecycleStateTerminating,
	"TERMINATED":   DbHomeLifecycleStateTerminated,
	"FAILED":       DbHomeLifecycleStateFailed,
	"UNKNOWN":      DbHomeLifecycleStateUnknown,
}

// GetDbHomeLifecycleStateEnumValues Enumerates the set of values for DbHomeLifecycleState
func GetDbHomeLifecycleStateEnumValues() []DbHomeLifecycleStateEnum {
	values := make([]DbHomeLifecycleStateEnum, 0)
	for _, v := range mappingDbHomeLifecycleState {
		if v != DbHomeLifecycleStateUnknown {
			values = append(values, v)
		}
	}
	return values
}
