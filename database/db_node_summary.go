// Copyright (c) 2016, 2017, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Database Service API
//
// The API for the Database Service.
//

package database

import (
	"github.com/oci-go-sdk/common"
)

// DbNodeSummary. A server where Oracle database software is running.
// To use any of the API operations, you must be authorized in an IAM policy. If you're not authorized, talk to an administrator. If you're an administrator who needs to write policies to give users access, see [Getting Started with Policies]({{DOC_SERVER_URL}}/Content/Identity/Concepts/policygetstarted.htm).
type DbNodeSummary struct {

	// The OCID of the DB System.
	DbSystemID *string `mandatory:"true" json:"dbSystemId,omitempty"`

	// The OCID of the DB Node.
	ID *string `mandatory:"true" json:"id,omitempty"`

	// The current state of the database node.
	LifecycleState DbNodeSummaryLifecycleStateEnum `mandatory:"true" json:"lifecycleState,omitempty"`

	// The date and time that the DB Node was created.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated,omitempty"`

	// The OCID of the VNIC.
	VnicID *string `mandatory:"true" json:"vnicId,omitempty"`

	// The OCID of the backup VNIC.
	BackupVnicID *string `mandatory:"false" json:"backupVnicId,omitempty"`

	// The host name for the DB Node.
	Hostname *string `mandatory:"false" json:"hostname,omitempty"`

	// Storage size, in GBs, of the software volume that is allocated to the DB system. This is applicable only for VM-based DBs.
	SoftwareStorageSizeInGB *int `mandatory:"false" json:"softwareStorageSizeInGB,omitempty"`
}

func (model DbNodeSummary) String() string {
	return common.PointerString(model)
}

type DbNodeSummaryLifecycleStateEnum string

const (
	DB_NODE_SUMMARY_LIFECYCLE_STATE_PROVISIONING DbNodeSummaryLifecycleStateEnum = "PROVISIONING"
	DB_NODE_SUMMARY_LIFECYCLE_STATE_AVAILABLE    DbNodeSummaryLifecycleStateEnum = "AVAILABLE"
	DB_NODE_SUMMARY_LIFECYCLE_STATE_UPDATING     DbNodeSummaryLifecycleStateEnum = "UPDATING"
	DB_NODE_SUMMARY_LIFECYCLE_STATE_STOPPING     DbNodeSummaryLifecycleStateEnum = "STOPPING"
	DB_NODE_SUMMARY_LIFECYCLE_STATE_STOPPED      DbNodeSummaryLifecycleStateEnum = "STOPPED"
	DB_NODE_SUMMARY_LIFECYCLE_STATE_STARTING     DbNodeSummaryLifecycleStateEnum = "STARTING"
	DB_NODE_SUMMARY_LIFECYCLE_STATE_TERMINATING  DbNodeSummaryLifecycleStateEnum = "TERMINATING"
	DB_NODE_SUMMARY_LIFECYCLE_STATE_TERMINATED   DbNodeSummaryLifecycleStateEnum = "TERMINATED"
	DB_NODE_SUMMARY_LIFECYCLE_STATE_FAILED       DbNodeSummaryLifecycleStateEnum = "FAILED"
	DB_NODE_SUMMARY_LIFECYCLE_STATE_UNKNOWN      DbNodeSummaryLifecycleStateEnum = "UNKNOWN"
)

var mapping_dbnodesummary_lifecycleState = map[string]DbNodeSummaryLifecycleStateEnum{
	"PROVISIONING": DB_NODE_SUMMARY_LIFECYCLE_STATE_PROVISIONING,
	"AVAILABLE":    DB_NODE_SUMMARY_LIFECYCLE_STATE_AVAILABLE,
	"UPDATING":     DB_NODE_SUMMARY_LIFECYCLE_STATE_UPDATING,
	"STOPPING":     DB_NODE_SUMMARY_LIFECYCLE_STATE_STOPPING,
	"STOPPED":      DB_NODE_SUMMARY_LIFECYCLE_STATE_STOPPED,
	"STARTING":     DB_NODE_SUMMARY_LIFECYCLE_STATE_STARTING,
	"TERMINATING":  DB_NODE_SUMMARY_LIFECYCLE_STATE_TERMINATING,
	"TERMINATED":   DB_NODE_SUMMARY_LIFECYCLE_STATE_TERMINATED,
	"FAILED":       DB_NODE_SUMMARY_LIFECYCLE_STATE_FAILED,
	"UNKNOWN":      DB_NODE_SUMMARY_LIFECYCLE_STATE_UNKNOWN,
}

func GetDbNodeSummaryLifecycleStateEnumValues() []DbNodeSummaryLifecycleStateEnum {
	values := make([]DbNodeSummaryLifecycleStateEnum, 0)
	for _, v := range mapping_dbnodesummary_lifecycleState {
		if v != DB_NODE_SUMMARY_LIFECYCLE_STATE_UNKNOWN {
			values = append(values, v)
		}
	}
	return values
}
