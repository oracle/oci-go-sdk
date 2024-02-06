// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Globally Distributed Database
//
// Use the Globally Distributed Database service APIs to create and manage distributed databases.
//

package globallydistributeddatabase

import (
	"strings"
)

// OperationTypeEnum Enum with underlying type: string
type OperationTypeEnum string

// Set of constants representing the allowable values for OperationTypeEnum
const (
	OperationTypeCreateShardedDatabase             OperationTypeEnum = "CREATE_SHARDED_DATABASE"
	OperationTypeDeleteShardedDatabase             OperationTypeEnum = "DELETE_SHARDED_DATABASE"
	OperationTypeUpdateShardedDatabase             OperationTypeEnum = "UPDATE_SHARDED_DATABASE"
	OperationTypeConfigureShardedDatabaseGsms      OperationTypeEnum = "CONFIGURE_SHARDED_DATABASE_GSMS"
	OperationTypeStartShardedDatabase              OperationTypeEnum = "START_SHARDED_DATABASE"
	OperationTypeStopShardedDatabase               OperationTypeEnum = "STOP_SHARDED_DATABASE"
	OperationTypeValidateNetwork                   OperationTypeEnum = "VALIDATE_NETWORK"
	OperationTypeChangeShardedDbCompartment        OperationTypeEnum = "CHANGE_SHARDED_DB_COMPARTMENT"
	OperationTypeCreatePrivateEndpoint             OperationTypeEnum = "CREATE_PRIVATE_ENDPOINT"
	OperationTypeUpdatePrivateEndpoint             OperationTypeEnum = "UPDATE_PRIVATE_ENDPOINT"
	OperationTypeDeletePrivateEndpoint             OperationTypeEnum = "DELETE_PRIVATE_ENDPOINT"
	OperationTypeChangePrivateEndpointCompartment  OperationTypeEnum = "CHANGE_PRIVATE_ENDPOINT_COMPARTMENT"
	OperationTypeInsertShards                      OperationTypeEnum = "INSERT_SHARDS"
	OperationTypeRemoveShards                      OperationTypeEnum = "REMOVE_SHARDS"
	OperationTypeMergeCatalogShards                OperationTypeEnum = "MERGE_CATALOG_SHARDS"
	OperationTypeUploadSignedCertAndGenerateWallet OperationTypeEnum = "UPLOAD_SIGNED_CERT_AND_GENERATE_WALLET"
	OperationTypeGenerateGsmCertSigningReq         OperationTypeEnum = "GENERATE_GSM_CERT_SIGNING_REQ"
	OperationTypeConfigureSharding                 OperationTypeEnum = "CONFIGURE_SHARDING"
)

var mappingOperationTypeEnum = map[string]OperationTypeEnum{
	"CREATE_SHARDED_DATABASE":                OperationTypeCreateShardedDatabase,
	"DELETE_SHARDED_DATABASE":                OperationTypeDeleteShardedDatabase,
	"UPDATE_SHARDED_DATABASE":                OperationTypeUpdateShardedDatabase,
	"CONFIGURE_SHARDED_DATABASE_GSMS":        OperationTypeConfigureShardedDatabaseGsms,
	"START_SHARDED_DATABASE":                 OperationTypeStartShardedDatabase,
	"STOP_SHARDED_DATABASE":                  OperationTypeStopShardedDatabase,
	"VALIDATE_NETWORK":                       OperationTypeValidateNetwork,
	"CHANGE_SHARDED_DB_COMPARTMENT":          OperationTypeChangeShardedDbCompartment,
	"CREATE_PRIVATE_ENDPOINT":                OperationTypeCreatePrivateEndpoint,
	"UPDATE_PRIVATE_ENDPOINT":                OperationTypeUpdatePrivateEndpoint,
	"DELETE_PRIVATE_ENDPOINT":                OperationTypeDeletePrivateEndpoint,
	"CHANGE_PRIVATE_ENDPOINT_COMPARTMENT":    OperationTypeChangePrivateEndpointCompartment,
	"INSERT_SHARDS":                          OperationTypeInsertShards,
	"REMOVE_SHARDS":                          OperationTypeRemoveShards,
	"MERGE_CATALOG_SHARDS":                   OperationTypeMergeCatalogShards,
	"UPLOAD_SIGNED_CERT_AND_GENERATE_WALLET": OperationTypeUploadSignedCertAndGenerateWallet,
	"GENERATE_GSM_CERT_SIGNING_REQ":          OperationTypeGenerateGsmCertSigningReq,
	"CONFIGURE_SHARDING":                     OperationTypeConfigureSharding,
}

var mappingOperationTypeEnumLowerCase = map[string]OperationTypeEnum{
	"create_sharded_database":                OperationTypeCreateShardedDatabase,
	"delete_sharded_database":                OperationTypeDeleteShardedDatabase,
	"update_sharded_database":                OperationTypeUpdateShardedDatabase,
	"configure_sharded_database_gsms":        OperationTypeConfigureShardedDatabaseGsms,
	"start_sharded_database":                 OperationTypeStartShardedDatabase,
	"stop_sharded_database":                  OperationTypeStopShardedDatabase,
	"validate_network":                       OperationTypeValidateNetwork,
	"change_sharded_db_compartment":          OperationTypeChangeShardedDbCompartment,
	"create_private_endpoint":                OperationTypeCreatePrivateEndpoint,
	"update_private_endpoint":                OperationTypeUpdatePrivateEndpoint,
	"delete_private_endpoint":                OperationTypeDeletePrivateEndpoint,
	"change_private_endpoint_compartment":    OperationTypeChangePrivateEndpointCompartment,
	"insert_shards":                          OperationTypeInsertShards,
	"remove_shards":                          OperationTypeRemoveShards,
	"merge_catalog_shards":                   OperationTypeMergeCatalogShards,
	"upload_signed_cert_and_generate_wallet": OperationTypeUploadSignedCertAndGenerateWallet,
	"generate_gsm_cert_signing_req":          OperationTypeGenerateGsmCertSigningReq,
	"configure_sharding":                     OperationTypeConfigureSharding,
}

// GetOperationTypeEnumValues Enumerates the set of values for OperationTypeEnum
func GetOperationTypeEnumValues() []OperationTypeEnum {
	values := make([]OperationTypeEnum, 0)
	for _, v := range mappingOperationTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetOperationTypeEnumStringValues Enumerates the set of values in String for OperationTypeEnum
func GetOperationTypeEnumStringValues() []string {
	return []string{
		"CREATE_SHARDED_DATABASE",
		"DELETE_SHARDED_DATABASE",
		"UPDATE_SHARDED_DATABASE",
		"CONFIGURE_SHARDED_DATABASE_GSMS",
		"START_SHARDED_DATABASE",
		"STOP_SHARDED_DATABASE",
		"VALIDATE_NETWORK",
		"CHANGE_SHARDED_DB_COMPARTMENT",
		"CREATE_PRIVATE_ENDPOINT",
		"UPDATE_PRIVATE_ENDPOINT",
		"DELETE_PRIVATE_ENDPOINT",
		"CHANGE_PRIVATE_ENDPOINT_COMPARTMENT",
		"INSERT_SHARDS",
		"REMOVE_SHARDS",
		"MERGE_CATALOG_SHARDS",
		"UPLOAD_SIGNED_CERT_AND_GENERATE_WALLET",
		"GENERATE_GSM_CERT_SIGNING_REQ",
		"CONFIGURE_SHARDING",
	}
}

// GetMappingOperationTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingOperationTypeEnum(val string) (OperationTypeEnum, bool) {
	enum, ok := mappingOperationTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
