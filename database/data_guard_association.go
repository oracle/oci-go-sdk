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

// DataGuardAssociation. The properties that define a Data Guard association.
// To use any of the API operations, you must be authorized in an IAM policy. If you're not authorized, talk to an
// administrator. If you're an administrator who needs to write policies to give users access, see
// [Getting Started with Policies]({{DOC_SERVER_URL}}/Content/Identity/Concepts/policygetstarted.htm).
// For information about endpoints and signing API requests, see
// [About the API]({{DOC_SERVER_URL}}/Content/API/Concepts/usingapi.htm). For information about available SDKs and tools, see
// [SDKS and Other Tools]({{DOC_SERVER_URL}}/Content/API/Concepts/sdks.htm).
type DataGuardAssociation struct {

	// The [OCID]({{DOC_SERVER_URL}}/Content/General/Concepts/identifiers.htm) of the reporting database.
	DatabaseID *string `mandatory:"true" json:"databaseId,omitempty"`

	// The [OCID]({{DOC_SERVER_URL}}/Content/General/Concepts/identifiers.htm) of the Data Guard association.
	ID *string `mandatory:"true" json:"id,omitempty"`

	// The current state of the Data Guard association.
	LifecycleState DataGuardAssociationLifecycleStateEnum `mandatory:"true" json:"lifecycleState,omitempty"`

	// The [OCID]({{DOC_SERVER_URL}}/Content/General/Concepts/identifiers.htm) of the DB System containing the associated
	// peer database.
	PeerDbSystemID *string `mandatory:"true" json:"peerDbSystemId,omitempty"`

	// The role of the peer database in this Data Guard association.
	PeerRole DataGuardAssociationPeerRoleEnum `mandatory:"true" json:"peerRole,omitempty"`

	// The protection mode of this Data Guard association. For more information, see
	// [Oracle Data Guard Protection Modes](http://docs.oracle.com/database/122/SBYDB/oracle-data-guard-protection-modes.htm#SBYDB02000)
	// in the Oracle Data Guard documentation.
	ProtectionMode DataGuardAssociationProtectionModeEnum `mandatory:"true" json:"protectionMode,omitempty"`

	// The role of the reporting database in this Data Guard association.
	Role DataGuardAssociationRoleEnum `mandatory:"true" json:"role,omitempty"`

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
	TransportType DataGuardAssociationTransportTypeEnum `mandatory:"false" json:"transportType,omitempty"`
}

func (model DataGuardAssociation) String() string {
	return common.PointerString(model)
}

type DataGuardAssociationLifecycleStateEnum string

const (
	DATA_GUARD_ASSOCIATION_LIFECYCLE_STATE_PROVISIONING DataGuardAssociationLifecycleStateEnum = "PROVISIONING"
	DATA_GUARD_ASSOCIATION_LIFECYCLE_STATE_AVAILABLE    DataGuardAssociationLifecycleStateEnum = "AVAILABLE"
	DATA_GUARD_ASSOCIATION_LIFECYCLE_STATE_UPDATING     DataGuardAssociationLifecycleStateEnum = "UPDATING"
	DATA_GUARD_ASSOCIATION_LIFECYCLE_STATE_TERMINATING  DataGuardAssociationLifecycleStateEnum = "TERMINATING"
	DATA_GUARD_ASSOCIATION_LIFECYCLE_STATE_TERMINATED   DataGuardAssociationLifecycleStateEnum = "TERMINATED"
	DATA_GUARD_ASSOCIATION_LIFECYCLE_STATE_FAILED       DataGuardAssociationLifecycleStateEnum = "FAILED"
	DATA_GUARD_ASSOCIATION_LIFECYCLE_STATE_UNKNOWN      DataGuardAssociationLifecycleStateEnum = "UNKNOWN"
)

var mapping_dataguardassociation_lifecycleState = map[string]DataGuardAssociationLifecycleStateEnum{
	"PROVISIONING": DATA_GUARD_ASSOCIATION_LIFECYCLE_STATE_PROVISIONING,
	"AVAILABLE":    DATA_GUARD_ASSOCIATION_LIFECYCLE_STATE_AVAILABLE,
	"UPDATING":     DATA_GUARD_ASSOCIATION_LIFECYCLE_STATE_UPDATING,
	"TERMINATING":  DATA_GUARD_ASSOCIATION_LIFECYCLE_STATE_TERMINATING,
	"TERMINATED":   DATA_GUARD_ASSOCIATION_LIFECYCLE_STATE_TERMINATED,
	"FAILED":       DATA_GUARD_ASSOCIATION_LIFECYCLE_STATE_FAILED,
	"UNKNOWN":      DATA_GUARD_ASSOCIATION_LIFECYCLE_STATE_UNKNOWN,
}

func GetDataGuardAssociationLifecycleStateEnumValues() []DataGuardAssociationLifecycleStateEnum {
	values := make([]DataGuardAssociationLifecycleStateEnum, 0)
	for _, v := range mapping_dataguardassociation_lifecycleState {
		if v != DATA_GUARD_ASSOCIATION_LIFECYCLE_STATE_UNKNOWN {
			values = append(values, v)
		}
	}
	return values
}

type DataGuardAssociationPeerRoleEnum string

const (
	DATA_GUARD_ASSOCIATION_PEER_ROLE_PRIMARY          DataGuardAssociationPeerRoleEnum = "PRIMARY"
	DATA_GUARD_ASSOCIATION_PEER_ROLE_STANDBY          DataGuardAssociationPeerRoleEnum = "STANDBY"
	DATA_GUARD_ASSOCIATION_PEER_ROLE_DISABLED_STANDBY DataGuardAssociationPeerRoleEnum = "DISABLED_STANDBY"
	DATA_GUARD_ASSOCIATION_PEER_ROLE_UNKNOWN          DataGuardAssociationPeerRoleEnum = "UNKNOWN"
)

var mapping_dataguardassociation_peerRole = map[string]DataGuardAssociationPeerRoleEnum{
	"PRIMARY":          DATA_GUARD_ASSOCIATION_PEER_ROLE_PRIMARY,
	"STANDBY":          DATA_GUARD_ASSOCIATION_PEER_ROLE_STANDBY,
	"DISABLED_STANDBY": DATA_GUARD_ASSOCIATION_PEER_ROLE_DISABLED_STANDBY,
	"UNKNOWN":          DATA_GUARD_ASSOCIATION_PEER_ROLE_UNKNOWN,
}

func GetDataGuardAssociationPeerRoleEnumValues() []DataGuardAssociationPeerRoleEnum {
	values := make([]DataGuardAssociationPeerRoleEnum, 0)
	for _, v := range mapping_dataguardassociation_peerRole {
		if v != DATA_GUARD_ASSOCIATION_PEER_ROLE_UNKNOWN {
			values = append(values, v)
		}
	}
	return values
}

type DataGuardAssociationProtectionModeEnum string

const (
	DATA_GUARD_ASSOCIATION_PROTECTION_MODE_AVAILABILITY DataGuardAssociationProtectionModeEnum = "MAXIMUM_AVAILABILITY"
	DATA_GUARD_ASSOCIATION_PROTECTION_MODE_PERFORMANCE  DataGuardAssociationProtectionModeEnum = "MAXIMUM_PERFORMANCE"
	DATA_GUARD_ASSOCIATION_PROTECTION_MODE_PROTECTION   DataGuardAssociationProtectionModeEnum = "MAXIMUM_PROTECTION"
	DATA_GUARD_ASSOCIATION_PROTECTION_MODE_UNKNOWN      DataGuardAssociationProtectionModeEnum = "UNKNOWN"
)

var mapping_dataguardassociation_protectionMode = map[string]DataGuardAssociationProtectionModeEnum{
	"MAXIMUM_AVAILABILITY": DATA_GUARD_ASSOCIATION_PROTECTION_MODE_AVAILABILITY,
	"MAXIMUM_PERFORMANCE":  DATA_GUARD_ASSOCIATION_PROTECTION_MODE_PERFORMANCE,
	"MAXIMUM_PROTECTION":   DATA_GUARD_ASSOCIATION_PROTECTION_MODE_PROTECTION,
	"UNKNOWN":              DATA_GUARD_ASSOCIATION_PROTECTION_MODE_UNKNOWN,
}

func GetDataGuardAssociationProtectionModeEnumValues() []DataGuardAssociationProtectionModeEnum {
	values := make([]DataGuardAssociationProtectionModeEnum, 0)
	for _, v := range mapping_dataguardassociation_protectionMode {
		if v != DATA_GUARD_ASSOCIATION_PROTECTION_MODE_UNKNOWN {
			values = append(values, v)
		}
	}
	return values
}

type DataGuardAssociationRoleEnum string

const (
	DATA_GUARD_ASSOCIATION_ROLE_PRIMARY          DataGuardAssociationRoleEnum = "PRIMARY"
	DATA_GUARD_ASSOCIATION_ROLE_STANDBY          DataGuardAssociationRoleEnum = "STANDBY"
	DATA_GUARD_ASSOCIATION_ROLE_DISABLED_STANDBY DataGuardAssociationRoleEnum = "DISABLED_STANDBY"
	DATA_GUARD_ASSOCIATION_ROLE_UNKNOWN          DataGuardAssociationRoleEnum = "UNKNOWN"
)

var mapping_dataguardassociation_role = map[string]DataGuardAssociationRoleEnum{
	"PRIMARY":          DATA_GUARD_ASSOCIATION_ROLE_PRIMARY,
	"STANDBY":          DATA_GUARD_ASSOCIATION_ROLE_STANDBY,
	"DISABLED_STANDBY": DATA_GUARD_ASSOCIATION_ROLE_DISABLED_STANDBY,
	"UNKNOWN":          DATA_GUARD_ASSOCIATION_ROLE_UNKNOWN,
}

func GetDataGuardAssociationRoleEnumValues() []DataGuardAssociationRoleEnum {
	values := make([]DataGuardAssociationRoleEnum, 0)
	for _, v := range mapping_dataguardassociation_role {
		if v != DATA_GUARD_ASSOCIATION_ROLE_UNKNOWN {
			values = append(values, v)
		}
	}
	return values
}

type DataGuardAssociationTransportTypeEnum string

const (
	DATA_GUARD_ASSOCIATION_TRANSPORT_TYPE_SYNC     DataGuardAssociationTransportTypeEnum = "SYNC"
	DATA_GUARD_ASSOCIATION_TRANSPORT_TYPE_ASYNC    DataGuardAssociationTransportTypeEnum = "ASYNC"
	DATA_GUARD_ASSOCIATION_TRANSPORT_TYPE_FASTSYNC DataGuardAssociationTransportTypeEnum = "FASTSYNC"
	DATA_GUARD_ASSOCIATION_TRANSPORT_TYPE_UNKNOWN  DataGuardAssociationTransportTypeEnum = "UNKNOWN"
)

var mapping_dataguardassociation_transportType = map[string]DataGuardAssociationTransportTypeEnum{
	"SYNC":     DATA_GUARD_ASSOCIATION_TRANSPORT_TYPE_SYNC,
	"ASYNC":    DATA_GUARD_ASSOCIATION_TRANSPORT_TYPE_ASYNC,
	"FASTSYNC": DATA_GUARD_ASSOCIATION_TRANSPORT_TYPE_FASTSYNC,
	"UNKNOWN":  DATA_GUARD_ASSOCIATION_TRANSPORT_TYPE_UNKNOWN,
}

func GetDataGuardAssociationTransportTypeEnumValues() []DataGuardAssociationTransportTypeEnum {
	values := make([]DataGuardAssociationTransportTypeEnum, 0)
	for _, v := range mapping_dataguardassociation_transportType {
		if v != DATA_GUARD_ASSOCIATION_TRANSPORT_TYPE_UNKNOWN {
			values = append(values, v)
		}
	}
	return values
}
