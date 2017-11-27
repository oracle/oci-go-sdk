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

// DbHomeSummary. A directory where Oracle database software is installed. Each DB System can have multiple database homes,
// and each database home can have multiple databases within it. All the databases within a single database home
// must be the same database version, but different database homes can run different versions. For more information,
// see [Managing Oracle Databases]({{DOC_SERVER_URL}}/Content/Database/Concepts/overview.htm).
// To use any of the API operations, you must be authorized in an IAM policy. If you're not authorized, talk to an
// administrator. If you're an administrator who needs to write policies to give users access,
// see [Getting Started with Policies]({{DOC_SERVER_URL}}/Content/Identity/Concepts/policygetstarted.htm).
type DbHomeSummary struct {

	// The OCID of the compartment.
	CompartmentID *string `mandatory:"true" json:"compartmentId,omitempty"`

	// The Oracle database version.
	DbVersion *string `mandatory:"true" json:"dbVersion,omitempty"`

	// The user-provided name for the database home. It does not need to be unique.
	DisplayName *string `mandatory:"true" json:"displayName,omitempty"`

	// The OCID of the database home.
	ID *string `mandatory:"true" json:"id,omitempty"`

	// The current state of the database home.
	LifecycleState DbHomeSummaryLifecycleStateEnum `mandatory:"true" json:"lifecycleState,omitempty"`

	// The OCID of the DB System.
	DbSystemID *string `mandatory:"false" json:"dbSystemId,omitempty"`

	// The OCID of the last patch history. This is updated as soon as a patch operation is started.
	LastPatchHistoryEntryID *string `mandatory:"false" json:"lastPatchHistoryEntryId,omitempty"`

	// The date and time the database home was created.
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated,omitempty"`
}

func (model DbHomeSummary) String() string {
	return common.PointerString(model)
}

type DbHomeSummaryLifecycleStateEnum string

const (
	DB_HOME_SUMMARY_LIFECYCLE_STATE_PROVISIONING DbHomeSummaryLifecycleStateEnum = "PROVISIONING"
	DB_HOME_SUMMARY_LIFECYCLE_STATE_AVAILABLE    DbHomeSummaryLifecycleStateEnum = "AVAILABLE"
	DB_HOME_SUMMARY_LIFECYCLE_STATE_UPDATING     DbHomeSummaryLifecycleStateEnum = "UPDATING"
	DB_HOME_SUMMARY_LIFECYCLE_STATE_TERMINATING  DbHomeSummaryLifecycleStateEnum = "TERMINATING"
	DB_HOME_SUMMARY_LIFECYCLE_STATE_TERMINATED   DbHomeSummaryLifecycleStateEnum = "TERMINATED"
	DB_HOME_SUMMARY_LIFECYCLE_STATE_FAILED       DbHomeSummaryLifecycleStateEnum = "FAILED"
	DB_HOME_SUMMARY_LIFECYCLE_STATE_UNKNOWN      DbHomeSummaryLifecycleStateEnum = "UNKNOWN"
)

var mapping_dbhomesummary_lifecycleState = map[string]DbHomeSummaryLifecycleStateEnum{
	"PROVISIONING": DB_HOME_SUMMARY_LIFECYCLE_STATE_PROVISIONING,
	"AVAILABLE":    DB_HOME_SUMMARY_LIFECYCLE_STATE_AVAILABLE,
	"UPDATING":     DB_HOME_SUMMARY_LIFECYCLE_STATE_UPDATING,
	"TERMINATING":  DB_HOME_SUMMARY_LIFECYCLE_STATE_TERMINATING,
	"TERMINATED":   DB_HOME_SUMMARY_LIFECYCLE_STATE_TERMINATED,
	"FAILED":       DB_HOME_SUMMARY_LIFECYCLE_STATE_FAILED,
	"UNKNOWN":      DB_HOME_SUMMARY_LIFECYCLE_STATE_UNKNOWN,
}

func GetDbHomeSummaryLifecycleStateEnumValues() []DbHomeSummaryLifecycleStateEnum {
	values := make([]DbHomeSummaryLifecycleStateEnum, 0)
	for _, v := range mapping_dbhomesummary_lifecycleState {
		if v != DB_HOME_SUMMARY_LIFECYCLE_STATE_UNKNOWN {
			values = append(values, v)
		}
	}
	return values
}
