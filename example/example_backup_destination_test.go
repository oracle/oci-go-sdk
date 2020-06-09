// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.

// Example code for Backup Destination.
// It demonstrates CRUD(Create, Read, Update and Delete) operations on BackupDestination, Create DbHome with
// BackupDestination and Update Database with BackupDestination.

package example

import (
	"context"
	"fmt"
	"github.com/oracle/oci-go-sdk/common"
	"github.com/oracle/oci-go-sdk/database"
	"github.com/oracle/oci-go-sdk/example/helpers"
	"os"
	"strings"
)

var (
	displayNameBackupDestination, localMountPath string
)

func ExampleCreateNFSBackupDestination() {
	c, clerr := database.NewDatabaseClientWithConfigurationProvider(common.DefaultConfigProvider())
	helpers.FatalIfError(clerr)
	displayNameBackupDestination := helpers.GetRandomString(32)
	compartmentId := os.Getenv("OCI_COMPARTMENT_ID")

	createBackupDestinationDetails := database.CreateNfsBackupDestinationDetails{
		CompartmentId:       &compartmentId,
		DisplayName:         &displayNameBackupDestination,
		LocalMountPointPath: common.String("path"),
	}

	createbackupdestinationReq := database.CreateBackupDestinationRequest{
		CreateBackupDestinationDetails: createBackupDestinationDetails,
	}

	_, err := c.CreateBackupDestination(context.Background(), createbackupdestinationReq)
	helpers.FatalIfError(err)

	fmt.Println("create backup destination is successful")

	// Output:
	// create backup destination is successful
}

func ExampleGetBackupDestination() {
	c, clerr := database.NewDatabaseClientWithConfigurationProvider(common.DefaultConfigProvider())
	helpers.FatalIfError(clerr)

	getbackupDestinationReq := database.GetBackupDestinationRequest{
		BackupDestinationId: common.String("backup-destination-ocid"),
	}

	_, err := c.GetBackupDestination(context.Background(), getbackupDestinationReq)
	helpers.FatalIfError(err)

	fmt.Println("get backup destination is successful")

	// Output:
	// get backup destination is successful
}

func ExampleUpdateBackupDestination() {
	c, clerr := database.NewDatabaseClientWithConfigurationProvider(common.DefaultConfigProvider())
	helpers.FatalIfError(clerr)

	updateBackupDestinationDetails := database.UpdateBackupDestinationDetails{
		LocalMountPointPath: &localMountPath,
	}

	updatebackupdestinationReq := database.UpdateBackupDestinationRequest{
		UpdateBackupDestinationDetails: updateBackupDestinationDetails,
		BackupDestinationId:            common.String("backup-destination-ocid"),
	}

	_, err := c.UpdateBackupDestination(context.Background(), updatebackupdestinationReq)
	helpers.FatalIfError(err)

	fmt.Println("update backup destination is successful")

	// Output:
	// update backup destination is successful
}

func ExampleUpdateDbBackupBackupDestination() {
	c, clerr := database.NewDatabaseClientWithConfigurationProvider(common.DefaultConfigProvider())
	helpers.FatalIfError(clerr)

	backupDestination := database.BackupDestinationDetails{
		Type: database.BackupDestinationDetailsTypeEnum("NFS"),
		Id:   common.String("backup-destination-ocid"),
	}

	dbBackupConfig := database.DbBackupConfig{
		BackupDestinationDetails: []database.BackupDestinationDetails{backupDestination},
	}

	updatedatabaseDetails := database.UpdateDatabaseDetails{
		DbBackupConfig: &dbBackupConfig,
	}

	updateDatabaseReq := database.UpdateDatabaseRequest{
		UpdateDatabaseDetails: updatedatabaseDetails,
		DatabaseId:            common.String("database-ocid"),
	}

	_, err := c.UpdateDatabase(context.Background(), updateDatabaseReq)
	helpers.FatalIfError(err)

	fmt.Println("update backup destination is successful")

	// Output:
	// update backup destination is successful
}

func ExampleCreateDbHomeBackupDestination() {
	c, clerr := database.NewDatabaseClientWithConfigurationProvider(common.DefaultConfigProvider())
	helpers.FatalIfError(clerr)

	dbName := strings.ToLower(helpers.GetRandomString(8))
	dbUniqueName := dbName + "_" + strings.ToLower(helpers.GetRandomString(20))
	dbVersion := "18.0.0.0"
	adminPassword := "DBaaS12345_#"
	displayName := helpers.GetRandomString(32)

	backupDestination := database.BackupDestinationDetails{
		Type: database.BackupDestinationDetailsTypeEnum("NFS"),
		Id:   common.String("backup-destination-ocid"),
	}

	dbBackupConfig := database.DbBackupConfig{
		BackupDestinationDetails: []database.BackupDestinationDetails{backupDestination},
	}

	// create database details
	createDatabaseDetails := database.CreateDatabaseDetails{
		AdminPassword:  &adminPassword,
		DbName:         &dbName,
		DbUniqueName:   &dbUniqueName,
		DbBackupConfig: &dbBackupConfig,
	}

	// create dbhome details
	createDbHomeDetails := database.CreateDbHomeWithVmClusterIdDetails{
		DisplayName: &displayName,
		Database:    &createDatabaseDetails,
		VmClusterId: common.String("vm-cluster-ocid"),
		DbVersion:   &dbVersion,
	}

	// create dbome request
	request := database.CreateDbHomeRequest{CreateDbHomeWithDbSystemIdDetails: createDbHomeDetails}

	_, createErrors := c.CreateDbHome(context.Background(), request)

	helpers.FatalIfError(createErrors)

	fmt.Printf("Create DB Home with backupDestination completed")

	// Output:
	// Create DB Home with backupDestination completed
}

func ExampleDeleteBackupDestination() {
	c, clerr := database.NewDatabaseClientWithConfigurationProvider(common.DefaultConfigProvider())
	helpers.FatalIfError(clerr)

	deletebackupDestinationReq := database.DeleteBackupDestinationRequest{
		BackupDestinationId: common.String("backup-destination-ocid"),
	}

	_, err := c.DeleteBackupDestination(context.Background(), deletebackupDestinationReq)
	helpers.FatalIfError(err)

	fmt.Println("delete backup destination is successful")

	// Output:
	// delete backup destination is successful
}
