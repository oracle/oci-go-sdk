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

// CreateDataGuardAssociationToExistingDbSystemDetails. The configuration details for creating a Data Guard association to an existing database.
type CreateDataGuardAssociationToExistingDbSystemDetails struct {

	// Specifies where to create the associated database.
	// "ExistingDbSystem" is the only supported `creationType` value.
	CreationType *string `mandatory:"true" json:"creationType,omitempty"`

	// A strong password for the `SYS`, `SYSTEM`, and `PDB Admin` users to apply during standby creation.
	// The password must contain no fewer than nine characters and include:
	// * At least two uppercase characters.
	// * At least two lowercase characters.
	// * At least two numeric characters.
	// * At least two special characters. Valid special characters include "_", "#", and "-" only.
	// **The password MUST be the same as the primary admin password.**
	DatabaseAdminPassword *string `mandatory:"true" json:"databaseAdminPassword,omitempty"`

	// The protection mode to set up between the primary and standby databases. For more information, see
	// [Oracle Data Guard Protection Modes](http://docs.oracle.com/database/122/SBYDB/oracle-data-guard-protection-modes.htm#SBYDB02000)
	// in the Oracle Data Guard documentation.
	// **IMPORTANT** - The only protection mode currently supported by the Database Service is MAXIMUM_PERFORMANCE.
	ProtectionMode CreateDataGuardAssociationToExistingDbSystemDetailsProtectionModeEnum `mandatory:"true" json:"protectionMode,omitempty"`

	// The redo transport type to use for this Data Guard association.  Valid values depend on the specified `protectionMode`:
	// * MAXIMUM_AVAILABILITY - SYNC or FASTSYNC
	// * MAXIMUM_PERFORMANCE - ASYNC
	// * MAXIMUM_PROTECTION - SYNC
	// For more information, see
	// [Redo Transport Services](http://docs.oracle.com/database/122/SBYDB/oracle-data-guard-redo-transport-services.htm#SBYDB00400)
	// in the Oracle Data Guard documentation.
	// **IMPORTANT** - The only transport type currently supported by the Database Service is ASYNC.
	TransportType CreateDataGuardAssociationToExistingDbSystemDetailsTransportTypeEnum `mandatory:"true" json:"transportType,omitempty"`

	// The [OCID]({{DOC_SERVER_URL}}/Content/General/Concepts/identifiers.htm) of the DB System to create the standby database on.
	PeerDbSystemID *string `mandatory:"false" json:"peerDbSystemId,omitempty"`
}

func (model CreateDataGuardAssociationToExistingDbSystemDetails) String() string {
	return common.PointerString(model)
}

type CreateDataGuardAssociationToExistingDbSystemDetailsProtectionModeEnum string

const (
	CREATE_DATA_GUARD_ASSOCIATION_TO_EXISTING_DB_SYSTEM_DETAILS_PROTECTION_MODE_AVAILABILITY CreateDataGuardAssociationToExistingDbSystemDetailsProtectionModeEnum = "MAXIMUM_AVAILABILITY"
	CREATE_DATA_GUARD_ASSOCIATION_TO_EXISTING_DB_SYSTEM_DETAILS_PROTECTION_MODE_PERFORMANCE  CreateDataGuardAssociationToExistingDbSystemDetailsProtectionModeEnum = "MAXIMUM_PERFORMANCE"
	CREATE_DATA_GUARD_ASSOCIATION_TO_EXISTING_DB_SYSTEM_DETAILS_PROTECTION_MODE_PROTECTION   CreateDataGuardAssociationToExistingDbSystemDetailsProtectionModeEnum = "MAXIMUM_PROTECTION"
	CREATE_DATA_GUARD_ASSOCIATION_TO_EXISTING_DB_SYSTEM_DETAILS_PROTECTION_MODE_UNKNOWN      CreateDataGuardAssociationToExistingDbSystemDetailsProtectionModeEnum = "UNKNOWN"
)

var mapping_createdataguardassociationtoexistingdbsystemdetails_protectionMode = map[string]CreateDataGuardAssociationToExistingDbSystemDetailsProtectionModeEnum{
	"MAXIMUM_AVAILABILITY": CREATE_DATA_GUARD_ASSOCIATION_TO_EXISTING_DB_SYSTEM_DETAILS_PROTECTION_MODE_AVAILABILITY,
	"MAXIMUM_PERFORMANCE":  CREATE_DATA_GUARD_ASSOCIATION_TO_EXISTING_DB_SYSTEM_DETAILS_PROTECTION_MODE_PERFORMANCE,
	"MAXIMUM_PROTECTION":   CREATE_DATA_GUARD_ASSOCIATION_TO_EXISTING_DB_SYSTEM_DETAILS_PROTECTION_MODE_PROTECTION,
	"UNKNOWN":              CREATE_DATA_GUARD_ASSOCIATION_TO_EXISTING_DB_SYSTEM_DETAILS_PROTECTION_MODE_UNKNOWN,
}

func GetCreateDataGuardAssociationToExistingDbSystemDetailsProtectionModeEnumValues() []CreateDataGuardAssociationToExistingDbSystemDetailsProtectionModeEnum {
	values := make([]CreateDataGuardAssociationToExistingDbSystemDetailsProtectionModeEnum, 0)
	for _, v := range mapping_createdataguardassociationtoexistingdbsystemdetails_protectionMode {
		if v != CREATE_DATA_GUARD_ASSOCIATION_TO_EXISTING_DB_SYSTEM_DETAILS_PROTECTION_MODE_UNKNOWN {
			values = append(values, v)
		}
	}
	return values
}

type CreateDataGuardAssociationToExistingDbSystemDetailsTransportTypeEnum string

const (
	CREATE_DATA_GUARD_ASSOCIATION_TO_EXISTING_DB_SYSTEM_DETAILS_TRANSPORT_TYPE_SYNC     CreateDataGuardAssociationToExistingDbSystemDetailsTransportTypeEnum = "SYNC"
	CREATE_DATA_GUARD_ASSOCIATION_TO_EXISTING_DB_SYSTEM_DETAILS_TRANSPORT_TYPE_ASYNC    CreateDataGuardAssociationToExistingDbSystemDetailsTransportTypeEnum = "ASYNC"
	CREATE_DATA_GUARD_ASSOCIATION_TO_EXISTING_DB_SYSTEM_DETAILS_TRANSPORT_TYPE_FASTSYNC CreateDataGuardAssociationToExistingDbSystemDetailsTransportTypeEnum = "FASTSYNC"
	CREATE_DATA_GUARD_ASSOCIATION_TO_EXISTING_DB_SYSTEM_DETAILS_TRANSPORT_TYPE_UNKNOWN  CreateDataGuardAssociationToExistingDbSystemDetailsTransportTypeEnum = "UNKNOWN"
)

var mapping_createdataguardassociationtoexistingdbsystemdetails_transportType = map[string]CreateDataGuardAssociationToExistingDbSystemDetailsTransportTypeEnum{
	"SYNC":     CREATE_DATA_GUARD_ASSOCIATION_TO_EXISTING_DB_SYSTEM_DETAILS_TRANSPORT_TYPE_SYNC,
	"ASYNC":    CREATE_DATA_GUARD_ASSOCIATION_TO_EXISTING_DB_SYSTEM_DETAILS_TRANSPORT_TYPE_ASYNC,
	"FASTSYNC": CREATE_DATA_GUARD_ASSOCIATION_TO_EXISTING_DB_SYSTEM_DETAILS_TRANSPORT_TYPE_FASTSYNC,
	"UNKNOWN":  CREATE_DATA_GUARD_ASSOCIATION_TO_EXISTING_DB_SYSTEM_DETAILS_TRANSPORT_TYPE_UNKNOWN,
}

func GetCreateDataGuardAssociationToExistingDbSystemDetailsTransportTypeEnumValues() []CreateDataGuardAssociationToExistingDbSystemDetailsTransportTypeEnum {
	values := make([]CreateDataGuardAssociationToExistingDbSystemDetailsTransportTypeEnum, 0)
	for _, v := range mapping_createdataguardassociationtoexistingdbsystemdetails_transportType {
		if v != CREATE_DATA_GUARD_ASSOCIATION_TO_EXISTING_DB_SYSTEM_DETAILS_TRANSPORT_TYPE_UNKNOWN {
			values = append(values, v)
		}
	}
	return values
}
