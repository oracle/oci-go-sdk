// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.

// Example code for blockstorage API.

package example

import (
	"context"
	"fmt"
	"time"

	"github.com/oracle/oci-go-sdk/common"
	"github.com/oracle/oci-go-sdk/core"
	"github.com/oracle/oci-go-sdk/example/helpers"
)

// Copies a volume backup to another region.
// Polls the copied volume backup in the destination region until it's lifecycle is Available.
func ExampleCopyVolumeBackup() {
	sourceBackupId := "REPLACE_WITH_VOLUME_BACKUP_OCID"
	destinationRegion := "REPLACE_WITH_DESTINATION_REGION_NAME"
	// displayName can be empty, in which case the copied backup will have the same display name as the original backup
	displayName := ""
	// kmsKey is optional too. If not specified, the copied backup is going to be encrypted with oracle provided
	// encryption keys.
	kmsKeyId := ""

	//Creating the copyVolumeBackupRequest
	request := core.CopyVolumeBackupRequest{
		CopyVolumeBackupDetails: core.CopyVolumeBackupDetails{
			DestinationRegion: common.String(destinationRegion),
		},
		VolumeBackupId: common.String(sourceBackupId),
	}
	if len(displayName) > 0 {
		request.CopyVolumeBackupDetails.DisplayName = common.String(displayName)
	}

	if len(kmsKeyId) > 0 {
		request.CopyVolumeBackupDetails.KmsKeyId = common.String(kmsKeyId)
	}

	// Creating a Blockstorage client in the source region to initiate the copy.
	bs, err := core.NewBlockstorageClientWithConfigurationProvider(common.DefaultConfigProvider())
	helpers.FatalIfError(err)
	fmt.Println("Copying backup.")
	copyResponse, err := bs.CopyVolumeBackup(context.Background(), request)
	volumeBackup := copyResponse.VolumeBackup
	helpers.FatalIfError(err)
	fmt.Println("Copy backup request accepted waiting for the backup to be in Available state.")
	backupState := volumeBackup.LifecycleState

	// Creating a Blockstorage client in the destination region and
	// poll on the copied volume backup's lifecycle state.
	destinationBS, err := core.NewBlockstorageClientWithConfigurationProvider(common.DefaultConfigProvider())
	destinationBS.SetRegion(destinationRegion)
	helpers.FatalIfError(err)
	for backupState != core.VolumeBackupLifecycleStateAvailable {
		time.Sleep(15 * time.Second)
		getVolumeBackupRequest := core.GetVolumeBackupRequest{
			VolumeBackupId: volumeBackup.Id,
		}

		getVolumeBackupResponse, err := destinationBS.GetVolumeBackup(context.Background(), getVolumeBackupRequest)
		helpers.FatalIfError(err)
		backupState = getVolumeBackupResponse.VolumeBackup.LifecycleState
	}
	// VolumeBackup's lifecycle state reached Available.
	fmt.Println("Backup copy succeeded.")

	// Output:
	// Copying backup.
	// Copy backup request accepted waiting for the backup to be in Available state.
	// Backup copy succeeded.
}
