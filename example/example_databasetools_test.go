// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.

// Examples for creating and validating Database Tools Connections with and without
// Database Tools Private Endpoint Reverse Connections. Examples include:
//
// Oracle:
//   - ExampleCreateADBsConnectionWithPublicIp
//   - ExampleCreateADBsConnectionWithPrivateEndpoint
//
// MySQL:
//   - ExampleCreateMySqlConnectionWithPublicIp
//   - ExampleCreateMySqlDbSystemConnectionWithPrivateEndpoint
//
// Starting the examples:
//   - First, set environment variables accordingly (see comments below)
//   - Run with go test
//     go test ./example -run ^ExampleCreateADBsConnectionWithPublicIp$ -v
//
// Prerequisites and environment variables are noted in the comment preceding
// each Example*. All examples assume an OCI config with a [DEFAULT] profile
package example

import (
	"archive/zip"
	"bytes"
	"context"
	"encoding/base64"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"time"

	"github.com/oracle/oci-go-sdk/v65/common"
	"github.com/oracle/oci-go-sdk/v65/database"
	"github.com/oracle/oci-go-sdk/v65/databasetools"
	"github.com/oracle/oci-go-sdk/v65/example/helpers"
	"github.com/oracle/oci-go-sdk/v65/keymanagement"
	"github.com/oracle/oci-go-sdk/v65/mysql"
	"github.com/oracle/oci-go-sdk/v65/vault"
)

// The config struct used here is an arbitrary collection of the information
// that the examples below require to function. It's helpful to understand the
// scope of the information used during the process of creating connections with
// the Database Tools service, but keep in mind that only certain bits are used
// depending on which SDK is being used. Furthermore, these are academic examples.
// Some liberties are taken with respect to error handling as the main focus is
// on demonstrating the SDK.
type config struct {
	// populated by environment variables
	DatabaseId       *string // because the SDK expects pointers
	VaultId          *string
	Username         *string
	Password         *string
	ConnectionString *string
	Ctx              context.Context

	// populated by reading [home]/.oci/config
	Provider common.ConfigurationProvider

	// SDK clients populated at runtime
	DBToolsClient   databasetools.DatabaseToolsClient
	VaultsClient    vault.VaultsClient
	KmsVaultsClient keymanagement.KmsVaultClient
	DatabaseClient  database.DatabaseClient
	DbSystemClient  mysql.DbSystemClient

	// Resources used in various steps populated at runtime
	VaultCompartmentId  *string
	VaultKeyId          *string
	WalletBase64        *string
	TargetCompartmentId *string
	EndpointServiceId   *string
	SubnetId            *string
}

// Example Use Case: Existing ADB-S with public IP (no ACL)
//
// This example creates a Database Tools Connection to an Autonomous Database
// (ADB) on Shared Exadata Infrastructure, accessible by public IP. Note, since
// this connection will be against a public IP address, a Database Tools Private
// Endpoint Reverse Connection is not required.
//
// Prerequisites:
//   - An existing ADB-S
//   - An existing Vault for storage of secrets
//   - A previously configured .oci/config file with a [DEFAULT] section
//   - The following environment variables set:
//   - OCI_DBS_OCID   : The ocid for an ADB-s database
//   - OCI_VAULT_OCID : The ocid for a vault (to store secrets)
//   - OCI_DB_USER    : The Oracle database user to connect with
//   - OCI_DB_PASS    : The Oracle database password to connect with
//
// High-level Steps:
//
//	1- Locate the Autonomous Database (ADB-S) by the provided OCID
//	2- Locate the Vault by the provided OCID
//	3- Download the wallet for the ADB-S
//	4- Store the secrets in the Vault (as base64 encoded strings)
//	5- Create a Database Tools connection
//	6- Validate the connection
//
//	... cleanup when done (delete the temporary secrets and connection)
//
//	                     Client
//	                       |
//	                       |
//	+----------------------+----------+
//	|                      V          |
//	|              +----------------+ |
//	|              | Database Tools | |
//	|              |    Service     | |
//	|              +----------------+ |
//	|                      |          |
//	| Database             |          |
//	| Tools                |          |
//	| VCN                  |          |
//	+----------------------+----------+
//	                       |
//	                       |
//	+--------------+       |
//	| Customer     |       |
//	| VCN          |       |
//	+--------------+       |
//	                       |
//	                       |
//	+----------------------+----------+
//	|                      |          |
//	|                      V          |
//	|                  ---------      |
//	|                 /  ABD-S  \     |
//	|                 \Public IP/     |
//	|                  ---------      |
//	|                                 |
//	| ADB                             |
//	| Shared                          |
//	| VCN                             |
//	+---------------------------------+
func ExampleCreateADBsConnectionWithPublicIp() {
	// Parses environment variables, .oci/config, and sets up the SDK clients
	cfg := newConfig()

	// Ignoring errors for simplicity
	walletSecretId, _ := createSecretInVault(cfg.WalletBase64, cfg)
	passwdSecretId, _ := createSecretInVault(cfg.Password, cfg)
	dbConnectionId, _ := createDatabaseToolsConnectionADBsPublicIp(walletSecretId, passwdSecretId, cfg)

	if ok := validateDatabaseToolsConnectionOracle(dbConnectionId, cfg); ok {
		log.Println("connection is valid")
	}

	// ... cleanup resources when finished, comment out the following delete
	// calls to keep the resources created above.
	if err := deleteConnection(dbConnectionId, cfg); err != nil {
		log.Printf("error deleting connection: %v\n", err)
	}
	if err := deleteSecret(passwdSecretId, cfg); err != nil {
		log.Printf("error deleting secret: %v\n", err)
	}
	if err := deleteSecret(walletSecretId, cfg); err != nil {
		log.Printf("error deleting secret: %v\n", err)
	}

	fmt.Println("ExampleCreateADBsConnectionWithPublicIp complete")
	// Output:
	// ExampleCreateADBsConnectionWithPublicIp complete
}

// Example Use Case: Existing ADB-S with Private Endpoint
//
// This example creates a Database Tools Connection to a Autonomous Database (ADB-S)
// accessible by private endpoint (PE). Note, since this connection will be against
// a PE, a Database Tools Private Endpoint Reverse Connection is required. This
// example serves as an academic exercise of the SDK.
//
// Prerequisites:
//   - An existing ADB-S with PE and network security group (i.e. ingress on 1522)
//   - Available capacity (limits apply) to create a new Private Endpoint
//   - An existing Vault for storage of secrets
//   - A previously configured .oci/config file with a [DEFAULT] section
//   - The following environment variables set:
//   - OCI_DBS_OCID   : The ocid for an ADB-s database
//   - OCI_VAULT_OCID : The ocid for a vault (to store secrets)
//   - OCI_DB_USER    : The Oracle database user to connect with
//   - OCI_DB_PASS    : The Oracle database password to connect with
//
// High-level Steps:
//
//	1- Locate the Autonomous Database (ADB-S) by the provided OCID
//	2- Locate the Vault by the provided OCID
//	3- Download the wallet for the ADB-S
//	4- Store the secrets in the Vault (as base64 encoded strings)
//	5- Create a Database Tools Private Endpoint for a Reverse Connection to the Private Endpoint of the ADB-S
//	6- Create a Database Tools connection
//	7- Validate the connection
//
//	... cleanup when done (delete the temporary secrets, connection, and PE)
//
//	                     Client
//	                       |
//	                       |
//	+----------------------+----------+
//	|                      V          |
//	|              +----------------+ |
//	|              | Database Tools | |
//	|              |    Service     | |
//	|              +----------------+ |
//	|                      |          |
//	| Database             |          |
//	| Tools                |          |
//	| VCN                  |          |
//	+----------------------+----------+
//	                       |
//	                       |
//	+----------------------+----------+
//	|                      |          |
//	|                      V          |
//	|                +-----------+    |
//	|                | Database  |    |
//	|                |  Tools    |    |
//	|                | Private   |    |
//	|                | Endpoint  |    |
//	|                |  Reverse  |    |
//	|                | Connection|    |
//	|                +-----------+    |
//	|                      |          |
//	|                      V          |
//	|                +-----------+    |
//	|                |   ADB-S   |    |
//	|                |  Private  |    |
//	|                |  Endpoint |    |
//	|                +-----------+    |
//	|                      |          |
//	| Customer             |          |
//	| VCN                  |          |
//	+----------------------+----------+
//	                       |
//	                       |
//	+----------------------+----------+
//	|                      |          |
//	|                      V          |
//	|                  ---------      |
//	|                 /  ABD-S  \     |
//	|                 | Private |     |
//	|                 \ Endpoint/     |
//	|                  ---------      |
//	|                                 |
//	| ADB                             |
//	| Shared                          |
//	| VCN                             |
//	+---------------------------------+
func ExampleCreateADBsConnectionWithPrivateEndpoint() {
	// Parses environment variables, .oci/config, and sets up the SDK clients
	cfg := newConfig()

	// Ignoring errors for simplicity
	privateEndpointId, _ := createDbToolsPrivateEndpoint(cfg)
	walletSecretId, _ := createSecretInVault(cfg.WalletBase64, cfg)
	passwdSecretId, _ := createSecretInVault(cfg.Password, cfg)
	dbConnectionId, _ := createDatabaseToolsConnectionADBs(walletSecretId, passwdSecretId, privateEndpointId, cfg)

	if ok := validateDatabaseToolsConnectionOracle(dbConnectionId, cfg); ok {
		log.Println("connection is valid")
	}

	// ... cleanup resources when finished, comment out the following delete
	// calls to keep the resources created above.
	if err := deleteConnection(dbConnectionId, cfg); err != nil {
		log.Printf("error deleting connection: %v\n", err)
	}
	if err := deleteSecret(passwdSecretId, cfg); err != nil {
		log.Printf("error deleting secret: %v\n", err)
	}
	if err := deleteSecret(walletSecretId, cfg); err != nil {
		log.Printf("error deleting secret: %v\n", err)
	}
	if err := deletePrivateEndpoint(privateEndpointId, cfg); err != nil {
		log.Printf("error deleting private endpoint: %v\n", err)
	}

	fmt.Println("ExampleCreateADBsConnectionWithPrivateEndpoint complete")
	// Output:
	// ExampleCreateADBsConnectionWithPrivateEndpoint complete
}

// Example Use Case: Existing MySQL database with public IP (customer-managed)
//
// This example creates a Database Tools Connection to a MySQL database
// accessible by public IP. Note, since this connection will be against a
// public IP address, a Database Tools Private Endpoint Reverse Connection
// is not required. Exposing a database directly to the Internet is not a
// recommended practice for security reasons. This example serves as an
// academic exercise of the SDK and proof of concept only.
//
// Prerequisites:
//   - An existing MySQL database on a compute node, for example
//   - Firewall or security list entries allowing TCP traffic to MySQL
//   - An existing Vault for storage of secrets
//   - A previously configured .oci/config file with a [DEFAULT] section
//   - The following environment variables set:
//   - OCI_VAULT_OCID  : The ocid for a vault (to store secrets)
//   - OCI_DB_USER     : The MySQL database user to connect with
//   - OCI_DB_PASS     : The MySQL database password to connect with
//   - OCI_CONN_STRING : The MySQL connection string, asin  mysql://host:port
//
// High-level Steps:
//
//	1- Locate the Vault by the provided OCID
//	2- Store the secret in the Vault (as base64 encoded string)
//	3- Create a Database Tools Connection
//	4- Validate the connection
//
//	... cleanup when done (delete the temporary secret and connection)
//
//	                     Client
//	                       |
//	                       |
//	+----------------------+----------+
//	|                      V          |
//	|              +----------------+ |
//	|              | Database Tools | |
//	|              |    Service     | |
//	|              +----------------+ |
//	|                      |          |
//	| Database             |          |
//	| Tools                |          |
//	| VCN                  |          |
//	+----------------------+----------+
//	                       |
//	                       |
//	+----------------------+----------+
//	| Compute              |          |
//	| Node                 |          |
//	|                      |          |
//	|                      |          |
//	|                      V          |
//	|                  ---------      |
//	|                 /  MySQL  \     |
//	|                 \Public IP/     |
//	|                  ---------      |
//	|                                 |
//	+---------------------------------+
func ExampleCreateMySqlConnectionWithPublicIp() {
	// Parses environment variables, .oci/config, and sets up the SDK clients
	cfg := newConfig()

	// Ignoring errors for simplicity
	passwdSecretId, _ := createSecretInVault(cfg.Password, cfg)
	dbConnectionId, _ := createDatabaseToolsConnectionMySqlPublicIp(passwdSecretId, cfg)

	if ok := validateDatabaseToolsConnectionMySQL(dbConnectionId, cfg); ok {
		log.Println("connection is valid")
	}

	// ... cleanup resources when finished, comment out the following delete
	// calls to keep the resources created above.
	if err := deleteConnection(dbConnectionId, cfg); err != nil {
		log.Printf("error deleting connection: %v\n", err)
	}
	if err := deleteSecret(passwdSecretId, cfg); err != nil {
		log.Printf("error deleting secret: %v\n", err)
	}

	fmt.Println("ExampleCreateMySqlConnectionWithPublicIp complete")
	// Output:
	// ExampleCreateMySqlConnectionWithPublicIp complete
}

// Example Use Case: MySQL DB System with Database Tools Private Endpoint
//
// This example creates a Database Tools Connection to a MySQL DB System
// accessible by private IP. Note, since this connection will be against a
// private IP address, a Database Tools Private Endpoint Reverse Connection
// is required. This example serves as an academic exercise of the SDK.
//
// Prerequisites:
//   - An existing MySQL DB System in a VCN and associated subnet
//   - Available capacity (limits apply) to create a new Private Endpoint
//   - An existing Vault for storage of secrets
//   - A previously configured .oci/config file with a [DEFAULT] section
//   - The following environment variables set:
//   - OCI_DBS_OCID   : The ocid for a MySQL DB System
//   - OCI_VAULT_OCID : The ocid for a vault (to store secrets)
//   - OCI_DB_USER    : The MySQL database user to connect with
//   - OCI_DB_PASS    : The MySQL database password to connect with
//
// High-level Steps:
//
//	1- Locate the MySQL DB System by provided OCID
//	2- Locate the Vault by provided OCID
//	3- Store the secret in the Vault (as base64 encoded string)
//	4- Create a Database Tools Private Endpoint Reverse Connection
//	5- Create a Database Tools connection
//	6- Validate the connection
//
//	... cleanup when done (delete the temporary secret, connection, and PE)
//
//	                     Client
//	                       |
//	                       |
//	+----------------------+----------+
//	|                      V          |
//	|              +----------------+ |
//	|              | Database Tools | |
//	|              |    Service     | |
//	|              +----------------+ |
//	|                      |          |
//	| Database             |          |
//	| Tools                |          |
//	| VCN                  |          |
//	+----------------------+----------+
//	                       |
//	                       |
//	+----------------------+----------+
//	|                      |          |
//	|                      V          |
//	|                +-----------+    |
//	|                | Database  |    |
//	|                |  Tools    |    |
//	|                | Private   |    |
//	|                | Endpoint  |    |
//	|                +-----------+    |
//	|                      |          |
//	|                      |          |
//	|                      V          |
//	|                  ---------      |
//	|                 /  MDS    \     |
//	|                | Private  |     |
//	|                \   IP    /      |
//	|                 ---------       |
//	|                                 |
//	| Customer                        |
//	| VCN (jump host not required)    |
//	+---------------------------------+
func ExampleCreateMySqlDbSystemConnectionWithPrivateEndpoint() {
	// Parses environment variables, .oci/config, and sets up the SDK clients
	cfg := newConfig()

	// Ignoring errors for simplicity
	privateEndpointId, _ := createDbToolsPrivateEndpoint(cfg)
	passwdSecretId, _ := createSecretInVault(cfg.Password, cfg)
	dbConnectionId, _ := createDatabaseToolsConnectionMySql(passwdSecretId, privateEndpointId, cfg)

	if ok := validateDatabaseToolsConnectionMySQL(dbConnectionId, cfg); ok {
		log.Println("connection is valid")
	}

	// ... cleanup resources when finished, comment out the following delete
	// calls to keep the resources created above.
	if err := deleteConnection(dbConnectionId, cfg); err != nil {
		log.Printf("error deleting connection: %v\n", err)
	}
	if err := deleteSecret(passwdSecretId, cfg); err != nil {
		log.Printf("error deleting secret: %v\n", err)
	}
	if err := deletePrivateEndpoint(privateEndpointId, cfg); err != nil {
		log.Printf("error deleting private endpoint: %v\n", err)
	}

	fmt.Println("ExampleCreateMySqlDbSystemConnectionWithPrivateEndpoint complete")
	// Output:
	// ExampleCreateMySqlDbSystemConnectionWithPrivateEndpoint complete
}

func deletePrivateEndpoint(id *string, cfg config) error {
	if id == nil {
		log.Println("=== Skipping delete since no private endpoint was found")
		return nil
	}
	log.Printf("=== Deleting private endpoint: %+v\n", toString(id))
	client := cfg.DBToolsClient
	_, err := client.DeleteDatabaseToolsPrivateEndpoint(cfg.Ctx,
		databasetools.DeleteDatabaseToolsPrivateEndpointRequest{
			DatabaseToolsPrivateEndpointId: id,
		})
	if err != nil {
		return err
	}

	err = waitForPrivateEndpoint(id, databasetools.LifecycleStateDeleted, cfg)
	if err == nil {
		log.Println("=== Private endpoint deleted")
	}
	return err
}

func deleteConnection(id *string, cfg config) error {
	if id == nil {
		log.Println("=== Skipping delete since no connection was found")
		return nil
	}
	log.Printf("=== Deleting database tools connection: %+v\n", toString(id))
	client := cfg.DBToolsClient
	_, err := client.DeleteDatabaseToolsConnection(cfg.Ctx,
		databasetools.DeleteDatabaseToolsConnectionRequest{
			DatabaseToolsConnectionId: id,
		})
	if err != nil {
		return err
	}

	err = waitForConnection(id, databasetools.LifecycleStateDeleted, cfg)
	if err == nil {
		log.Println("=== Connection deleted")
	}
	return err
}

func deleteSecret(id *string, cfg config) error {
	if id == nil {
		log.Println("=== Skipping delete since no secret was found")
		return nil
	}
	log.Printf("=== Scheduling secret deletion: %+v\n", toString(id))
	client := cfg.VaultsClient
	_, err := client.ScheduleSecretDeletion(cfg.Ctx, vault.ScheduleSecretDeletionRequest{
		SecretId: id,
		ScheduleSecretDeletionDetails: vault.ScheduleSecretDeletionDetails{
			TimeOfDeletion: &common.SDKTime{time.Now().UTC().Add(time.Hour * 48)},
		},
	})
	if err != nil {
		return err
	}

	err = waitForSecret(id, vault.SecretLifecycleStatePendingDeletion, cfg)
	if err == nil {
		log.Println("=== Secret deletion is scheduled")
	}
	return err
}

func validateDatabaseToolsConnectionOracle(id *string, cfg config) bool {
	connectionDetailsType := databasetools.ValidateDatabaseToolsConnectionOracleDatabaseDetails{}
	return validateDatabaseToolsConnection(id, connectionDetailsType, cfg)
}

func validateDatabaseToolsConnectionMySQL(id *string, cfg config) bool {
	connectionDetailsType := databasetools.ValidateDatabaseToolsConnectionMySqlDetails{}
	return validateDatabaseToolsConnection(id, connectionDetailsType, cfg)
}

// validateDatabaseToolsConnection is an example of how to validate a database tools
// connection. Note, the validateion details should be supplied for either Oracle
// or MySql connection type structs. (see above)
func validateDatabaseToolsConnection(id *string, connectionDetailsType databasetools.ValidateDatabaseToolsConnectionDetails, cfg config) bool {
	if id == nil || connectionDetailsType == nil {
		log.Println("=== Skipping validate (connection is nil)")
		return false
	}
	log.Printf("=== Validating connection: %v\n", toString(id))
	response, err := cfg.DBToolsClient.ValidateDatabaseToolsConnection(cfg.Ctx,
		databasetools.ValidateDatabaseToolsConnectionRequest{
			DatabaseToolsConnectionId:              id,
			ValidateDatabaseToolsConnectionDetails: connectionDetailsType,
		})
	if err != nil {
		log.Printf("error validating connection: %v\n", err)
		return false
	}

	if toString(response.GetCode()) != "OK" {
		log.Printf("Validation response: %+v\n", response)
		return false
	}
	return true
}

func createDatabaseToolsConnectionMySqlPublicIp(passwordSecretId *string, cfg config) (*string, error) {
	return createDatabaseToolsConnectionMySql(passwordSecretId, nil, cfg)
}

// createDatabaseToolsConnectionMySql creates a database tools connection to a
// MySQL database. privateEndpointId can be nil, in which case the the
// connection should be for a database that can be reached by public IP.
func createDatabaseToolsConnectionMySql(passwordSecretId *string, privateEndpointId *string, cfg config) (*string, error) {
	displayName := common.String("dbtools-temp-connection-" + helpers.GetRandomString(10))

	if passwordSecretId == nil || cfg.TargetCompartmentId == nil || toString(cfg.ConnectionString) == "" || toString(cfg.Username) == "" {
		return nil, errors.New("connection cannot be created due to missing details")
	}

	connectionDetails := databasetools.CreateDatabaseToolsConnectionMySqlDetails{
		UserName:         cfg.Username,
		UserPassword:     databasetools.DatabaseToolsUserPasswordSecretIdDetails{SecretId: passwordSecretId},
		CompartmentId:    cfg.TargetCompartmentId,
		ConnectionString: cfg.ConnectionString,
		DisplayName:      displayName,
	}

	if privateEndpointId != nil {
		connectionDetails.PrivateEndpointId = privateEndpointId
	}

	return createDatabaseToolsConnection(connectionDetails, cfg)
}

// createDatabaseToolsConnectionADBsPublicIp creates a database tools connection
// to an Autonomous Database without a private endpoint. For this to work the ADBs
// needs to be accessible by public IP address with mTLS connections.
func createDatabaseToolsConnectionADBsPublicIp(walletSecretId *string, passwordSecretId *string, cfg config) (*string, error) {
	return createDatabaseToolsConnectionADBs(walletSecretId, passwordSecretId, nil, cfg)
}

// createDatabaseToolsConnectionADBs creates a database tools connection to an
// Autonomous Database. privateEndpointId can be nil, in which case the the
// connection should be for a database that can be reached by public IP.
func createDatabaseToolsConnectionADBs(walletSecretId *string, passwordSecretId *string, privateEndpointId *string, cfg config) (*string, error) {
	displayName := common.String("dbtools-temp-connection-" + helpers.GetRandomString(10))

	if walletSecretId == nil || passwordSecretId == nil || cfg.TargetCompartmentId == nil || toString(cfg.ConnectionString) == "" || toString(cfg.Username) == "" {
		return nil, errors.New("connection cannot be created due to missing details")
	}

	keyStore := databasetools.DatabaseToolsKeyStoreDetails{
		KeyStoreType: databasetools.KeyStoreTypeSso,
		KeyStoreContent: databasetools.DatabaseToolsKeyStoreContentSecretIdDetails{
			SecretId: walletSecretId,
		},
	}

	connectionDetails := databasetools.CreateDatabaseToolsConnectionOracleDatabaseDetails{
		KeyStores:        []databasetools.DatabaseToolsKeyStoreDetails{keyStore},
		UserName:         cfg.Username,
		UserPassword:     databasetools.DatabaseToolsUserPasswordSecretIdDetails{SecretId: passwordSecretId},
		CompartmentId:    cfg.TargetCompartmentId,
		ConnectionString: cfg.ConnectionString,
		DisplayName:      displayName,
	}

	if privateEndpointId != nil {
		connectionDetails.PrivateEndpointId = privateEndpointId
	}

	return createDatabaseToolsConnection(connectionDetails, cfg)
}

// createDatabaseToolsConnection is an example of how to setup a database tools
// connection. Note, the connectionDetails should be supplied fpr either Oracle
// or MySql connection detail structs. (see above)
func createDatabaseToolsConnection(connectionDetails databasetools.CreateDatabaseToolsConnectionDetails, cfg config) (*string, error) {
	client := cfg.DBToolsClient

	log.Printf("=== Creating connection %s in compartment %s\n",
		toString(connectionDetails.GetDisplayName()),
		toString(connectionDetails.GetCompartmentId()))

	log.Printf("Connection details provided: %+v\n", connectionDetails)

	response, err := client.CreateDatabaseToolsConnection(cfg.Ctx,
		databasetools.CreateDatabaseToolsConnectionRequest{
			CreateDatabaseToolsConnectionDetails: connectionDetails,
		})
	if err != nil {
		return nil, err
	}

	err = waitForConnection(response.GetId(), databasetools.LifecycleStateActive, cfg)
	if err == nil {
		log.Println("=== Connection is created")
	}
	return response.GetId(), err
}

// createDbToolsPrivateEndpoint is an example of how to setup a private endpoint
// using information gathered from an ADBs database (or a MySQL DB System). Given
// the appropriate tenancy and VCN details, a private endpoint can also be setup
// for Oracle or MySQL databases running on a compute node / subnet in a tenancy.
func createDbToolsPrivateEndpoint(cfg config) (*string, error) {
	client := cfg.DBToolsClient
	displayName := common.String("dbtools-temp-pe-" + helpers.GetRandomString(10))

	if cfg.TargetCompartmentId == nil || cfg.EndpointServiceId == nil || cfg.SubnetId == nil {
		return nil, errors.New("private endpoint cannot be created due to missing details")
	}

	log.Printf("=== Creating private endpoint %s with service %s\n", toString(displayName), toString(cfg.EndpointServiceId))
	response, err := client.CreateDatabaseToolsPrivateEndpoint(cfg.Ctx,
		databasetools.CreateDatabaseToolsPrivateEndpointRequest{
			CreateDatabaseToolsPrivateEndpointDetails: databasetools.CreateDatabaseToolsPrivateEndpointDetails{
				DisplayName:       displayName,
				CompartmentId:     cfg.TargetCompartmentId,
				EndpointServiceId: cfg.EndpointServiceId,
				SubnetId:          cfg.SubnetId,
			},
		})
	if err != nil {
		return nil, err
	}

	err = waitForPrivateEndpoint(response.Id, databasetools.LifecycleStateActive, cfg)
	if err == nil {
		log.Println("=== Private endpoint is created")
	}
	return response.Id, err
}

// createSecretInVault creates a secret in a vault within the target tenancy.
// Database Tools Connections do not store secrets directly. Only a reference
// (ocid) of the secret is stored with the connection. Database Tools calls out
// to the Key Management and Vaults services to use secrets.
func createSecretInVault(secretValue *string, cfg config) (*string, error) {
	client := cfg.VaultsClient
	displayName := common.String("dbtools-temp-secret-" + helpers.GetRandomString(10))

	if secretValue == nil || cfg.VaultId == nil || cfg.VaultKeyId == nil || cfg.TargetCompartmentId == nil {
		return nil, errors.New("secret cannot be created due to missing details")
	}

	log.Printf("=== Creating secret %s in vault %s\n", toString(displayName), toString(cfg.VaultId))
	response, err := client.CreateSecret(cfg.Ctx, vault.CreateSecretRequest{
		CreateSecretDetails: vault.CreateSecretDetails{
			CompartmentId: cfg.TargetCompartmentId,
			SecretContent: vault.Base64SecretContentDetails{
				Content: secretValue,
			},
			SecretName: displayName,
			VaultId:    cfg.VaultId,
			KeyId:      cfg.VaultKeyId, // Documented as not mandatory, but is required
		},
	})
	if err != nil {
		return nil, err
	}

	err = waitForSecret(response.Secret.Id, vault.SecretLifecycleStateActive, cfg)
	if err == nil {
		log.Println("=== Secret is created")
	}
	return response.Secret.Id, err
}

// The process of creating or deleting a secret is asynchronous so if a secret
// is to be used after the creation starts, for example, wait for the lifecycle
// state to be ACTIVE.
func waitForSecret(secretId *string, targetState vault.SecretLifecycleStateEnum, cfg config) error {
	client := cfg.VaultsClient
	for {
		secret, err := client.GetSecret(cfg.Ctx, vault.GetSecretRequest{
			SecretId: secretId,
		})
		if err != nil {
			return err
		}
		log.Printf("waitForSecret: lifecycle state = %v\n", secret.LifecycleState)

		if secret.LifecycleState == targetState {
			return nil
		}
		time.Sleep(5 * time.Second)
	}
}

// The process of creating or deleting a connection is asynchronous so if a
// connection is to be used after the creation starts, for example, wait for the
// lifecycle state to be ACTIVE.
func waitForConnection(connectionId *string, targetState databasetools.LifecycleStateEnum, cfg config) error {
	client := cfg.DBToolsClient
	for {
		connection, err := client.GetDatabaseToolsConnection(cfg.Ctx,
			databasetools.GetDatabaseToolsConnectionRequest{
				DatabaseToolsConnectionId: connectionId,
			})
		if err != nil {
			return err
		}
		log.Printf("waitForConnection: lifecycle state = %v\n", connection.GetLifecycleState())

		if connection.GetLifecycleState() == targetState {
			return nil
		}
		time.Sleep(10 * time.Second)
	}
}

// The process of creating or deleting a private endpoint is asynchronous so if
// a private endpoint is to be used after the creation starts, for example, wait
// for the lifecycle state to be ACTIVE.
func waitForPrivateEndpoint(endpointId *string, targetState databasetools.LifecycleStateEnum, cfg config) error {
	client := cfg.DBToolsClient
	for {
		pe, err := client.GetDatabaseToolsPrivateEndpoint(cfg.Ctx,
			databasetools.GetDatabaseToolsPrivateEndpointRequest{
				DatabaseToolsPrivateEndpointId: endpointId,
			})
		if err != nil {
			return err
		}
		log.Printf("waitForPrivateEndpoint: lifecycle state = %v\n", pe.LifecycleState)

		if pe.LifecycleState == targetState {
			return nil
		}
		time.Sleep(10 * time.Second)
	}
}

// newConfig returns a collection of the environment variables and SDK clients
// used by the examples. These are stored in a config struct. Helper functions
// are attached to config to make the examples above a little less verbose but
// this is purely to help focus the examples and integration with services.
func newConfig() config {
	cfg := config{
		Ctx:              context.Background(),
		Provider:         common.DefaultConfigProvider(),
		DatabaseId:       common.String(os.Getenv("OCI_DBS_OCID")),
		VaultId:          common.String(os.Getenv("OCI_VAULT_OCID")),
		Username:         common.String(os.Getenv("OCI_DB_USER")),
		ConnectionString: common.String(os.Getenv("OCI_CONN_STRING")),
		Password: common.String(
			base64.StdEncoding.EncodeToString(
				[]byte(os.Getenv("OCI_DB_PASS")))),
	}
	cfg.DBToolsClient = cfg.getDatabaseToolsClient()
	cfg.VaultsClient = cfg.getVaultsClient()
	cfg.KmsVaultsClient = cfg.getKmsVaultsClient()
	cfg.DatabaseClient = cfg.getDatabaseClient()
	cfg.DbSystemClient = cfg.getMySqlDbSystemClient()

	// Things that need to be initialized AFTER the SDK clients
	cfg.VaultCompartmentId = cfg.getVaultCompartmentId()
	cfg.VaultKeyId = cfg.getVaultKeyId()
	cfg.WalletBase64 = cfg.getADBsWalletAsBase64String()
	cfg.TargetCompartmentId = cfg.getTargetCompartmentId()
	cfg.EndpointServiceId = cfg.getEndpointServiceId() // depends on TargetCompartmentId
	cfg.SubnetId = cfg.getDatabaseSubnetId()

	cfg.ConnectionString = cfg.getConnectionString() // respects OCI_CONN_STRING if defined

	return cfg
}

func (cfg config) getADBsWalletAsBase64String() *string {
	client := cfg.DatabaseClient
	id := cfg.DatabaseId
	if !strings.Contains(toString(id), ".autonomousdatabase.") {
		if toString(cfg.DatabaseId) != "" {
			log.Printf("Database ID provided is not an Autonomous Database (no wallet to download): %s\n", toString(id))
		}
		return nil
	}
	response, err := client.GenerateAutonomousDatabaseWallet(cfg.Ctx,
		database.GenerateAutonomousDatabaseWalletRequest{
			AutonomousDatabaseId: id,
			GenerateAutonomousDatabaseWalletDetails: database.GenerateAutonomousDatabaseWalletDetails{
				Password:     common.String("1StrongPassword###"),
				GenerateType: "SINGLE",
			},
		})
	defer response.Content.Close()
	if err != nil {
		log.Printf("error generating ADBs wallet: %v\n", err)
		return nil
	}

	zipFileBytes, err := ioutil.ReadAll(response.Content)
	if err != nil {
		log.Printf("error reading ADBs wallet zip: %v\n", err)
		return nil
	}

	cwalletBytes, err := cfg.getFileFromZip(zipFileBytes, "cwallet.sso")
	if err != nil {
		log.Printf("error reading cwallet.sso from zip: %v\n", err)
		return nil
	}

	return common.String(
		base64.StdEncoding.EncodeToString(cwalletBytes))
}

func (cfg config) getEndpointServiceId() *string {
	client := cfg.DBToolsClient
	if cfg.TargetCompartmentId == nil {
		log.Println("unable to get endpoint service (compartment id is nil)")
		return nil
	}
	service, err := client.ListDatabaseToolsEndpointServices(cfg.Ctx,
		databasetools.ListDatabaseToolsEndpointServicesRequest{
			CompartmentId: cfg.TargetCompartmentId,
			Name:          common.String("DATABASE_TOOLS"),
		})
	if err != nil {
		log.Printf("error getting endpoint service: %v\n", err)
		return nil
	}
	if len(service.Items) > 0 {
		return service.Items[0].Id
	}
	return nil
}

func (cfg config) getTargetCompartmentId() *string {
	id := toString(cfg.DatabaseId)
	if strings.Contains(id, ".mysqldbsystem.") {
		db, err := cfg.getMySqlDbSystem()
		if err != nil {
			log.Printf("error getting MySQL DB system compartment: %v\n", err)
			return nil
		}
		return db.CompartmentId
	} else if strings.Contains(id, ".autonomousdatabase.") {
		db, err := cfg.getADBsDatabase()
		if err != nil {
			log.Printf("error getting ADBs compartment: %v\n", err)
			return nil
		}
		return db.CompartmentId
	} else if id == "" {
		tenancy, err := cfg.Provider.TenancyOCID()
		if err != nil {
			log.Printf("error getting root compartment from provider: %v\n", err)
			return nil
		}
		return common.String(tenancy)
	}
	return nil
}

func (cfg config) getDatabaseSubnetId() *string {
	id := toString(cfg.DatabaseId)
	if strings.Contains(id, ".mysqldbsystem.") {
		db, err := cfg.getMySqlDbSystem()
		if err != nil {
			log.Printf("error getting MySQL DB system subnet: %v\n", err)
			return nil
		}
		return db.SubnetId
	} else if strings.Contains(id, ".autonomousdatabase.") {
		db, err := cfg.getADBsDatabase()
		if err != nil {
			log.Printf("error getting ADBs subnet: %v\n", err)
			return nil
		}
		return db.SubnetId
	} else {
		return nil
	}
}

func (cfg config) getConnectionString() (connectionString *string) {
	log.Println("=== Getting database connection string")
	if toString(cfg.ConnectionString) != "" {
		connectionString = cfg.ConnectionString
	} else {
		id := toString(cfg.DatabaseId)
		if strings.Contains(id, ".mysqldbsystem.") {
			connectionString = cfg.getMySqlDbSystemConnectionString()
		} else if strings.Contains(id, ".autonomousdatabase.") {
			connectionString = cfg.getADBsConnectionString("low")
		} else {
			log.Println("Connection string is not defined")
		}
	}
	return
}

func (cfg config) getMySqlDbSystemConnectionString() *string {
	db, err := cfg.getMySqlDbSystem()
	if err != nil {
		log.Printf("error getting MySQL DB system connection string: %v\n", err)
		return nil
	}
	for _, endpoint := range db.Endpoints {
		if endpoint.Port != nil && endpoint.IpAddress != nil {
			connectionString := fmt.Sprintf("mysql://%s:%d", *endpoint.IpAddress, *endpoint.Port)
			return common.String(connectionString)
		}
	}
	return nil
}

func (cfg config) getADBsConnectionString(suffix string) *string {
	db, err := cfg.getADBsDatabase()
	if err != nil {
		log.Printf("error getting ADBs connection string: %v\n", err)
		return nil
	}
	for _, c := range db.ConnectionStrings.Profiles {
		if len(suffix) == 0 || strings.HasSuffix(toString(c.DisplayName), suffix) {
			return c.Value
		}
	}
	return nil
}

func (cfg config) getMySqlDbSystem() (*mysql.GetDbSystemResponse, error) {
	client := cfg.DbSystemClient
	id := cfg.DatabaseId
	if !strings.Contains(toString(id), ".mysqldbsystem.") {
		return nil, errors.New("database id provided is not a valid mysql db system " + toString(id))
	}
	db, err := client.GetDbSystem(cfg.Ctx, mysql.GetDbSystemRequest{
		DbSystemId: id,
	})
	if err != nil {
		return nil, err
	}
	return &db, nil
}

func (cfg config) getADBsDatabase() (*database.GetAutonomousDatabaseResponse, error) {
	client := cfg.DatabaseClient
	id := cfg.DatabaseId
	if !strings.Contains(toString(id), ".autonomousdatabase.") {
		return nil, errors.New("database id provided is not a valid autonomous database " + toString(id))
	}
	db, err := client.GetAutonomousDatabase(cfg.Ctx, database.GetAutonomousDatabaseRequest{
		AutonomousDatabaseId: cfg.DatabaseId,
	})
	if err != nil {
		return nil, err
	}
	return &db, nil
}

func (cfg config) getVaultCompartmentId() *string {
	client := cfg.KmsVaultsClient
	vault, err := client.GetVault(cfg.Ctx, keymanagement.GetVaultRequest{
		VaultId: cfg.VaultId,
	})
	if err != nil {
		log.Printf("error getting vault compartment: %v\n", err)
		return nil
	}
	return vault.CompartmentId
}

func (cfg config) getVaultKeyId() *string {
	vc := cfg.KmsVaultsClient
	vault, err := vc.GetVault(cfg.Ctx, keymanagement.GetVaultRequest{
		VaultId: cfg.VaultId,
	})
	if err != nil {
		log.Printf("error getting vault: %v\n", err)
		return nil
	}

	mc, err := keymanagement.NewKmsManagementClientWithConfigurationProvider(
		cfg.Provider,
		toString(vault.ManagementEndpoint))
	if err != nil {
		log.Printf("error getting KMS management client: %v\n", err)
		return nil
	}

	keys, err := mc.ListKeys(cfg.Ctx, keymanagement.ListKeysRequest{
		CompartmentId: cfg.VaultCompartmentId,
	})
	if err != nil {
		log.Printf("error getting vault key: %v\n", err)
		return nil
	}

	if len(keys.Items) == 0 {
		log.Println("error getting key from vault (no keys found)")
		return nil
	}
	return keys.Items[0].Id
}

func (cfg config) getMySqlDbSystemClient() mysql.DbSystemClient {
	client, err := mysql.NewDbSystemClientWithConfigurationProvider(cfg.Provider)
	logFatalIfErr(err)
	return client
}

func (cfg config) getDatabaseClient() database.DatabaseClient {
	client, err := database.NewDatabaseClientWithConfigurationProvider(cfg.Provider)
	logFatalIfErr(err)
	return client
}

func (cfg config) getDatabaseToolsClient() databasetools.DatabaseToolsClient {
	client, err := databasetools.NewDatabaseToolsClientWithConfigurationProvider(cfg.Provider)
	logFatalIfErr(err)
	return client
}

func (cfg config) getVaultsClient() vault.VaultsClient {
	client, err := vault.NewVaultsClientWithConfigurationProvider(cfg.Provider)
	logFatalIfErr(err)
	return client
}

func (cfg config) getKmsVaultsClient() keymanagement.KmsVaultClient {
	client, err := keymanagement.NewKmsVaultClientWithConfigurationProvider(cfg.Provider)
	logFatalIfErr(err)
	return client
}

// Note: log.Fatalln calls os.Exit which kills the process
func logFatalIfErr(err error) {
	if err != nil {
		log.Fatalln(err.Error())
	}
}

// Protect against dereferncing a nil string pointer which will cause a panic.
func toString(s *string) string {
	if s != nil {
		return *s
	}
	return ""
}

// Read the contents of a zip file in memory and return the bytes
// of a file specified within if found, or an error otherwise.
func (cfg config) getFileFromZip(zipFileBytes []byte, fileName string) ([]byte, error) {
	zipReader, err := zip.NewReader(bytes.NewReader(zipFileBytes), int64(len(zipFileBytes)))
	if err != nil {
		return nil, err
	}
	for _, f := range zipReader.File {
		if f.Name == fileName {
			fileReader, err := f.Open()
			defer fileReader.Close()
			if err != nil {
				return nil, err
			}

			fileBytes, err := ioutil.ReadAll(fileReader)
			if err != nil {
				return nil, err
			}
			return fileBytes, nil
		}
	}
	return nil, errors.New("file not found in zip: " + fileName)
}
