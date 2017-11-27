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

// CreateDataGuardAssociationDetails. The configuration details for creating a Data Guard association between databases.
// **NOTE:**
// "ExistingDbSystem" is the only supported `creationType` value. Therefore, all
// CreateDataGuardAssociation
// requests must include the `peerDbSystemId` parameter found in the
// CreateDataGuardAssociationToExistingDbSystemDetails
// object.
type CreateDataGuardAssociationDetails struct {

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
	ProtectionMode CreateDataGuardAssociationDetailsProtectionModeEnum `mandatory:"true" json:"protectionMode,omitempty"`

	// The redo transport type to use for this Data Guard association.  Valid values depend on the specified `protectionMode`:
	// * MAXIMUM_AVAILABILITY - SYNC or FASTSYNC
	// * MAXIMUM_PERFORMANCE - ASYNC
	// * MAXIMUM_PROTECTION - SYNC
	// For more information, see
	// [Redo Transport Services](http://docs.oracle.com/database/122/SBYDB/oracle-data-guard-redo-transport-services.htm#SBYDB00400)
	// in the Oracle Data Guard documentation.
	// **IMPORTANT** - The only transport type currently supported by the Database Service is ASYNC.
	TransportType CreateDataGuardAssociationDetailsTransportTypeEnum `mandatory:"true" json:"transportType,omitempty"`
}

func (model CreateDataGuardAssociationDetails) String() string {
	return common.PointerString(model)
}

type CreateDataGuardAssociationDetailsProtectionModeEnum string

const (
	CREATE_DATA_GUARD_ASSOCIATION_DETAILS_PROTECTION_MODE_AVAILABILITY CreateDataGuardAssociationDetailsProtectionModeEnum = "MAXIMUM_AVAILABILITY"
	CREATE_DATA_GUARD_ASSOCIATION_DETAILS_PROTECTION_MODE_PERFORMANCE  CreateDataGuardAssociationDetailsProtectionModeEnum = "MAXIMUM_PERFORMANCE"
	CREATE_DATA_GUARD_ASSOCIATION_DETAILS_PROTECTION_MODE_PROTECTION   CreateDataGuardAssociationDetailsProtectionModeEnum = "MAXIMUM_PROTECTION"
	CREATE_DATA_GUARD_ASSOCIATION_DETAILS_PROTECTION_MODE_UNKNOWN      CreateDataGuardAssociationDetailsProtectionModeEnum = "UNKNOWN"
)

var mapping_createdataguardassociationdetails_protectionMode = map[string]CreateDataGuardAssociationDetailsProtectionModeEnum{
	"MAXIMUM_AVAILABILITY": CREATE_DATA_GUARD_ASSOCIATION_DETAILS_PROTECTION_MODE_AVAILABILITY,
	"MAXIMUM_PERFORMANCE":  CREATE_DATA_GUARD_ASSOCIATION_DETAILS_PROTECTION_MODE_PERFORMANCE,
	"MAXIMUM_PROTECTION":   CREATE_DATA_GUARD_ASSOCIATION_DETAILS_PROTECTION_MODE_PROTECTION,
	"UNKNOWN":              CREATE_DATA_GUARD_ASSOCIATION_DETAILS_PROTECTION_MODE_UNKNOWN,
}

func GetCreateDataGuardAssociationDetailsProtectionModeEnumValues() []CreateDataGuardAssociationDetailsProtectionModeEnum {
	values := make([]CreateDataGuardAssociationDetailsProtectionModeEnum, 0)
	for _, v := range mapping_createdataguardassociationdetails_protectionMode {
		if v != CREATE_DATA_GUARD_ASSOCIATION_DETAILS_PROTECTION_MODE_UNKNOWN {
			values = append(values, v)
		}
	}
	return values
}

type CreateDataGuardAssociationDetailsTransportTypeEnum string

const (
	CREATE_DATA_GUARD_ASSOCIATION_DETAILS_TRANSPORT_TYPE_SYNC     CreateDataGuardAssociationDetailsTransportTypeEnum = "SYNC"
	CREATE_DATA_GUARD_ASSOCIATION_DETAILS_TRANSPORT_TYPE_ASYNC    CreateDataGuardAssociationDetailsTransportTypeEnum = "ASYNC"
	CREATE_DATA_GUARD_ASSOCIATION_DETAILS_TRANSPORT_TYPE_FASTSYNC CreateDataGuardAssociationDetailsTransportTypeEnum = "FASTSYNC"
	CREATE_DATA_GUARD_ASSOCIATION_DETAILS_TRANSPORT_TYPE_UNKNOWN  CreateDataGuardAssociationDetailsTransportTypeEnum = "UNKNOWN"
)

var mapping_createdataguardassociationdetails_transportType = map[string]CreateDataGuardAssociationDetailsTransportTypeEnum{
	"SYNC":     CREATE_DATA_GUARD_ASSOCIATION_DETAILS_TRANSPORT_TYPE_SYNC,
	"ASYNC":    CREATE_DATA_GUARD_ASSOCIATION_DETAILS_TRANSPORT_TYPE_ASYNC,
	"FASTSYNC": CREATE_DATA_GUARD_ASSOCIATION_DETAILS_TRANSPORT_TYPE_FASTSYNC,
	"UNKNOWN":  CREATE_DATA_GUARD_ASSOCIATION_DETAILS_TRANSPORT_TYPE_UNKNOWN,
}

func GetCreateDataGuardAssociationDetailsTransportTypeEnumValues() []CreateDataGuardAssociationDetailsTransportTypeEnum {
	values := make([]CreateDataGuardAssociationDetailsTransportTypeEnum, 0)
	for _, v := range mapping_createdataguardassociationdetails_transportType {
		if v != CREATE_DATA_GUARD_ASSOCIATION_DETAILS_TRANSPORT_TYPE_UNKNOWN {
			values = append(values, v)
		}
	}
	return values
}
