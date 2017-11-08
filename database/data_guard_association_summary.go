// Copyright (c) 2016, 2017, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Database Service API
//
// The API for the Database Service.
//

package database

import (
	"bitbucket.aka.lgl.grungy.us/golang-sdk2/common"
)

// DataGuardAssociationSummary. The properties that define a Data Guard association.
// To use any of the API operations, you must be authorized in an IAM policy. If you're not authorized, talk to an
// administrator. If you're an administrator who needs to write policies to give users access, see
// [Getting Started with Policies]({{DOC_SERVER_URL}}/Content/Identity/Concepts/policygetstarted.htm).
// For information about endpoints and signing API requests, see
// [About the API]({{DOC_SERVER_URL}}/Content/API/Concepts/usingapi.htm). For information about available SDKs and tools, see
// [SDKS and Other Tools]({{DOC_SERVER_URL}}/Content/API/Concepts/sdks.htm).
type DataGuardAssociationSummary struct {

	// The [OCID]({{DOC_SERVER_URL}}/Content/General/Concepts/identifiers.htm) of the reporting database.
	DatabaseID *string `mandatory:"true" json:"databaseId,omitempty"`

	// The [OCID]({{DOC_SERVER_URL}}/Content/General/Concepts/identifiers.htm) of the Data Guard association.
	ID *string `mandatory:"true" json:"id,omitempty"`

	// The current state of the Data Guard association.
	LifecycleState DataGuardAssociationSummaryLifecycleStateEnum `mandatory:"true" json:"lifecycleState,omitempty"`

	// The [OCID]({{DOC_SERVER_URL}}/Content/General/Concepts/identifiers.htm) of the DB System containing the associated
	// peer database.
	PeerDbSystemID *string `mandatory:"true" json:"peerDbSystemId,omitempty"`

	// The role of the peer database in this Data Guard association.
	PeerRole DataGuardAssociationSummaryPeerRoleEnum `mandatory:"true" json:"peerRole,omitempty"`

	// The protection mode of this Data Guard association. For more information, see
	// [Oracle Data Guard Protection Modes](http://docs.oracle.com/database/122/SBYDB/oracle-data-guard-protection-modes.htm#SBYDB02000)
	// in the Oracle Data Guard documentation.
	ProtectionMode DataGuardAssociationSummaryProtectionModeEnum `mandatory:"true" json:"protectionMode,omitempty"`

	// The role of the reporting database in this Data Guard association.
	Role DataGuardAssociationSummaryRoleEnum `mandatory:"true" json:"role,omitempty"`

	// The lag time between updates to the primary database and application of the redo data on the standby database,
	// as computed by the reporting database.
	// Example: `9 seconds`
	ApplyLag *string `mandatory:"false" json:"applyLag,omitempty"`

	// The rate at which redo logs are synced between the associated databases.
	// Example: `180 Mb per second`
	ApplyRate *string `mandatory:"false" json:"applyRate,omitempty"`

	// Additional information about the current lifecycleState, if available.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails,omitempty"`

	// The [OCID]({{DOC_SERVER_URL}}/Content/General/Concepts/identifiers.htm) of the peer database's Data Guard association.
	PeerDataGuardAssociationID *string `mandatory:"false" json:"peerDataGuardAssociationId,omitempty"`

	// The [OCID]({{DOC_SERVER_URL}}/Content/General/Concepts/identifiers.htm) of the associated peer database.
	PeerDatabaseID *string `mandatory:"false" json:"peerDatabaseId,omitempty"`

	// The [OCID]({{DOC_SERVER_URL}}/Content/General/Concepts/identifiers.htm) of the database home containing the associated peer database.
	PeerDbHomeID *string `mandatory:"false" json:"peerDbHomeId,omitempty"`

	// The date and time the Data Guard Association was created.
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated,omitempty"`

	// The redo transport type used by this Data Guard association.  For more information, see
	// [Redo Transport Services](http://docs.oracle.com/database/122/SBYDB/oracle-data-guard-redo-transport-services.htm#SBYDB00400)
	// in the Oracle Data Guard documentation.
	TransportType DataGuardAssociationSummaryTransportTypeEnum `mandatory:"false" json:"transportType,omitempty"`
}

func (model DataGuardAssociationSummary) String() string {
	return common.PointerString(model)
}

type DataGuardAssociationSummaryLifecycleStateEnum string

const (
	DATA_GUARD_ASSOCIATION_SUMMARY_LIFECYCLE_STATE_PROVISIONING DataGuardAssociationSummaryLifecycleStateEnum = "PROVISIONING"
	DATA_GUARD_ASSOCIATION_SUMMARY_LIFECYCLE_STATE_AVAILABLE    DataGuardAssociationSummaryLifecycleStateEnum = "AVAILABLE"
	DATA_GUARD_ASSOCIATION_SUMMARY_LIFECYCLE_STATE_UPDATING     DataGuardAssociationSummaryLifecycleStateEnum = "UPDATING"
	DATA_GUARD_ASSOCIATION_SUMMARY_LIFECYCLE_STATE_TERMINATING  DataGuardAssociationSummaryLifecycleStateEnum = "TERMINATING"
	DATA_GUARD_ASSOCIATION_SUMMARY_LIFECYCLE_STATE_TERMINATED   DataGuardAssociationSummaryLifecycleStateEnum = "TERMINATED"
	DATA_GUARD_ASSOCIATION_SUMMARY_LIFECYCLE_STATE_FAILED       DataGuardAssociationSummaryLifecycleStateEnum = "FAILED"
	DATA_GUARD_ASSOCIATION_SUMMARY_LIFECYCLE_STATE_UNKNOWN      DataGuardAssociationSummaryLifecycleStateEnum = "UNKNOWN"
)

var mapping_dataguardassociationsummary_lifecycleState = map[string]DataGuardAssociationSummaryLifecycleStateEnum{
	"PROVISIONING": DATA_GUARD_ASSOCIATION_SUMMARY_LIFECYCLE_STATE_PROVISIONING,
	"AVAILABLE":    DATA_GUARD_ASSOCIATION_SUMMARY_LIFECYCLE_STATE_AVAILABLE,
	"UPDATING":     DATA_GUARD_ASSOCIATION_SUMMARY_LIFECYCLE_STATE_UPDATING,
	"TERMINATING":  DATA_GUARD_ASSOCIATION_SUMMARY_LIFECYCLE_STATE_TERMINATING,
	"TERMINATED":   DATA_GUARD_ASSOCIATION_SUMMARY_LIFECYCLE_STATE_TERMINATED,
	"FAILED":       DATA_GUARD_ASSOCIATION_SUMMARY_LIFECYCLE_STATE_FAILED,
	"UNKNOWN":      DATA_GUARD_ASSOCIATION_SUMMARY_LIFECYCLE_STATE_UNKNOWN,
}

func GetDataGuardAssociationSummaryLifecycleStateEnumValues() []DataGuardAssociationSummaryLifecycleStateEnum {
	values := make([]DataGuardAssociationSummaryLifecycleStateEnum, 0)
	for _, v := range mapping_dataguardassociationsummary_lifecycleState {
		if v != DATA_GUARD_ASSOCIATION_SUMMARY_LIFECYCLE_STATE_UNKNOWN {
			values = append(values, v)
		}
	}
	return values
}

type DataGuardAssociationSummaryPeerRoleEnum string

const (
	DATA_GUARD_ASSOCIATION_SUMMARY_PEER_ROLE_PRIMARY          DataGuardAssociationSummaryPeerRoleEnum = "PRIMARY"
	DATA_GUARD_ASSOCIATION_SUMMARY_PEER_ROLE_STANDBY          DataGuardAssociationSummaryPeerRoleEnum = "STANDBY"
	DATA_GUARD_ASSOCIATION_SUMMARY_PEER_ROLE_DISABLED_STANDBY DataGuardAssociationSummaryPeerRoleEnum = "DISABLED_STANDBY"
	DATA_GUARD_ASSOCIATION_SUMMARY_PEER_ROLE_UNKNOWN          DataGuardAssociationSummaryPeerRoleEnum = "UNKNOWN"
)

var mapping_dataguardassociationsummary_peerRole = map[string]DataGuardAssociationSummaryPeerRoleEnum{
	"PRIMARY":          DATA_GUARD_ASSOCIATION_SUMMARY_PEER_ROLE_PRIMARY,
	"STANDBY":          DATA_GUARD_ASSOCIATION_SUMMARY_PEER_ROLE_STANDBY,
	"DISABLED_STANDBY": DATA_GUARD_ASSOCIATION_SUMMARY_PEER_ROLE_DISABLED_STANDBY,
	"UNKNOWN":          DATA_GUARD_ASSOCIATION_SUMMARY_PEER_ROLE_UNKNOWN,
}

func GetDataGuardAssociationSummaryPeerRoleEnumValues() []DataGuardAssociationSummaryPeerRoleEnum {
	values := make([]DataGuardAssociationSummaryPeerRoleEnum, 0)
	for _, v := range mapping_dataguardassociationsummary_peerRole {
		if v != DATA_GUARD_ASSOCIATION_SUMMARY_PEER_ROLE_UNKNOWN {
			values = append(values, v)
		}
	}
	return values
}

type DataGuardAssociationSummaryProtectionModeEnum string

const (
	DATA_GUARD_ASSOCIATION_SUMMARY_PROTECTION_MODE_AVAILABILITY DataGuardAssociationSummaryProtectionModeEnum = "MAXIMUM_AVAILABILITY"
	DATA_GUARD_ASSOCIATION_SUMMARY_PROTECTION_MODE_PERFORMANCE  DataGuardAssociationSummaryProtectionModeEnum = "MAXIMUM_PERFORMANCE"
	DATA_GUARD_ASSOCIATION_SUMMARY_PROTECTION_MODE_PROTECTION   DataGuardAssociationSummaryProtectionModeEnum = "MAXIMUM_PROTECTION"
	DATA_GUARD_ASSOCIATION_SUMMARY_PROTECTION_MODE_UNKNOWN      DataGuardAssociationSummaryProtectionModeEnum = "UNKNOWN"
)

var mapping_dataguardassociationsummary_protectionMode = map[string]DataGuardAssociationSummaryProtectionModeEnum{
	"MAXIMUM_AVAILABILITY": DATA_GUARD_ASSOCIATION_SUMMARY_PROTECTION_MODE_AVAILABILITY,
	"MAXIMUM_PERFORMANCE":  DATA_GUARD_ASSOCIATION_SUMMARY_PROTECTION_MODE_PERFORMANCE,
	"MAXIMUM_PROTECTION":   DATA_GUARD_ASSOCIATION_SUMMARY_PROTECTION_MODE_PROTECTION,
	"UNKNOWN":              DATA_GUARD_ASSOCIATION_SUMMARY_PROTECTION_MODE_UNKNOWN,
}

func GetDataGuardAssociationSummaryProtectionModeEnumValues() []DataGuardAssociationSummaryProtectionModeEnum {
	values := make([]DataGuardAssociationSummaryProtectionModeEnum, 0)
	for _, v := range mapping_dataguardassociationsummary_protectionMode {
		if v != DATA_GUARD_ASSOCIATION_SUMMARY_PROTECTION_MODE_UNKNOWN {
			values = append(values, v)
		}
	}
	return values
}

type DataGuardAssociationSummaryRoleEnum string

const (
	DATA_GUARD_ASSOCIATION_SUMMARY_ROLE_PRIMARY          DataGuardAssociationSummaryRoleEnum = "PRIMARY"
	DATA_GUARD_ASSOCIATION_SUMMARY_ROLE_STANDBY          DataGuardAssociationSummaryRoleEnum = "STANDBY"
	DATA_GUARD_ASSOCIATION_SUMMARY_ROLE_DISABLED_STANDBY DataGuardAssociationSummaryRoleEnum = "DISABLED_STANDBY"
	DATA_GUARD_ASSOCIATION_SUMMARY_ROLE_UNKNOWN          DataGuardAssociationSummaryRoleEnum = "UNKNOWN"
)

var mapping_dataguardassociationsummary_role = map[string]DataGuardAssociationSummaryRoleEnum{
	"PRIMARY":          DATA_GUARD_ASSOCIATION_SUMMARY_ROLE_PRIMARY,
	"STANDBY":          DATA_GUARD_ASSOCIATION_SUMMARY_ROLE_STANDBY,
	"DISABLED_STANDBY": DATA_GUARD_ASSOCIATION_SUMMARY_ROLE_DISABLED_STANDBY,
	"UNKNOWN":          DATA_GUARD_ASSOCIATION_SUMMARY_ROLE_UNKNOWN,
}

func GetDataGuardAssociationSummaryRoleEnumValues() []DataGuardAssociationSummaryRoleEnum {
	values := make([]DataGuardAssociationSummaryRoleEnum, 0)
	for _, v := range mapping_dataguardassociationsummary_role {
		if v != DATA_GUARD_ASSOCIATION_SUMMARY_ROLE_UNKNOWN {
			values = append(values, v)
		}
	}
	return values
}

type DataGuardAssociationSummaryTransportTypeEnum string

const (
	DATA_GUARD_ASSOCIATION_SUMMARY_TRANSPORT_TYPE_SYNC     DataGuardAssociationSummaryTransportTypeEnum = "SYNC"
	DATA_GUARD_ASSOCIATION_SUMMARY_TRANSPORT_TYPE_ASYNC    DataGuardAssociationSummaryTransportTypeEnum = "ASYNC"
	DATA_GUARD_ASSOCIATION_SUMMARY_TRANSPORT_TYPE_FASTSYNC DataGuardAssociationSummaryTransportTypeEnum = "FASTSYNC"
	DATA_GUARD_ASSOCIATION_SUMMARY_TRANSPORT_TYPE_UNKNOWN  DataGuardAssociationSummaryTransportTypeEnum = "UNKNOWN"
)

var mapping_dataguardassociationsummary_transportType = map[string]DataGuardAssociationSummaryTransportTypeEnum{
	"SYNC":     DATA_GUARD_ASSOCIATION_SUMMARY_TRANSPORT_TYPE_SYNC,
	"ASYNC":    DATA_GUARD_ASSOCIATION_SUMMARY_TRANSPORT_TYPE_ASYNC,
	"FASTSYNC": DATA_GUARD_ASSOCIATION_SUMMARY_TRANSPORT_TYPE_FASTSYNC,
	"UNKNOWN":  DATA_GUARD_ASSOCIATION_SUMMARY_TRANSPORT_TYPE_UNKNOWN,
}

func GetDataGuardAssociationSummaryTransportTypeEnumValues() []DataGuardAssociationSummaryTransportTypeEnum {
	values := make([]DataGuardAssociationSummaryTransportTypeEnum, 0)
	for _, v := range mapping_dataguardassociationsummary_transportType {
		if v != DATA_GUARD_ASSOCIATION_SUMMARY_TRANSPORT_TYPE_UNKNOWN {
			values = append(values, v)
		}
	}
	return values
}
