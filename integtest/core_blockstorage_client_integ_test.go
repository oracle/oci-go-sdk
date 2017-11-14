// Copyright (c) 2016, 2017, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Core Services API
//
// APIs for Networking Service, Compute Service, and Block Volume Service.
//

package integtest

import (
	"bitbucket.aka.lgl.grungy.us/golang-sdk2/common"
	"bitbucket.aka.lgl.grungy.us/golang-sdk2/core"
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

var (
	testRegionForBlockstorage = common.REGION_PHX
	validOCID                 = "ocidv1:tenancy:oc1:phx:1460406592660:aaaaaaaab4faofrfkxecohhjuivjq262pu"
	validAD                   = "kIdk:PHX-AD-2"
	validVolume               = "ocid1.volume.oc1.phx.abyhqljrm66bxtplnokqq762przj3s67ewvqlk7eidgpvwdjm35f2hxd6c5a"
	tick                      = time.Tick(500 * time.Millisecond)
	timeout                   = time.After(30 * time.Second)
)

//Volumes CRUDL
func TestBlockstorageClient_CreateVolume(t *testing.T) {

	c := core.NewBlockstorageClientForRegion(testRegionForBlockstorage)
	request := core.CreateVolumeRequest{}

	request.AvailabilityDomain = common.String(validAD)
	request.CompartmentID = common.String(validOCID)
	request.DisplayName = common.String("GoSDK2_TestBlockStorageCreateVolumeDisplayName")

	resCreate, err := c.CreateVolume(context.Background(), request)
	assert.NoError(t, err)
	assert.NotEmpty(t, resCreate.ID)

	//Read
	readTest := func() (interface{}, error) {
		c := core.NewBlockstorageClientForRegion(testRegionForBlockstorage)
		request := core.GetVolumeRequest{}
		request.VolumeID = resCreate.ID
		resRead, err := c.GetVolume(context.Background(), request)
		return resRead, err
	}
	assert.NoError(t,
		retryUntilTrueOrError(readTest,
			checkLifecycleState(string(core.VOLUME_LIFECYCLE_STATE_AVAILABLE)), tick, timeout))

	//update
	updateTest := func(t *testing.T) {
		c := core.NewBlockstorageClientForRegion(testRegionForBlockstorage)
		request := core.UpdateVolumeRequest{}
		request.VolumeID = resCreate.ID
		request.DisplayName = common.String("GoSDK2_TestBlockStorageCreateVolumeDisplayNameUpdated")
		r, err := c.UpdateVolume(context.Background(), request)
		assert.NotEmpty(t, r, fmt.Sprint(r))
		assert.NoError(t, err)
		return
	}
	updateTest(t)

	//list
	listTest := func(t *testing.T) {
		c := core.NewBlockstorageClientForRegion(testRegionForBlockstorage)
		request := core.ListVolumesRequest{}
		request.CompartmentID = common.String(validOCID)
		r, err := c.ListVolumes(context.Background(), request)
		assert.NoError(t, err)
		assert.NotEmpty(t, r.Items)
		return
	}
	listTest(t)

	//delete
	deleteTest := func(t *testing.T) {
		c := core.NewBlockstorageClientForRegion(testRegionForBlockstorage)
		request := core.DeleteVolumeRequest{}
		request.VolumeID = resCreate.ID
		err := c.DeleteVolume(context.Background(), request)
		assert.NoError(t, err)
		return
	}
	deleteTest(t)

	return
}

func TestBlockstorageClient_CreateVolumeBackup(t *testing.T) {
	//Create
	c := core.NewBlockstorageClientForRegion(testRegionForBlockstorage)
	request := core.CreateVolumeBackupRequest{}
	request.VolumeID = common.String(validVolume)
	request.DisplayName = common.String("GoSDK2_TestBlockStorageCreateVolumeBackup")
	r, err := c.CreateVolumeBackup(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	failIfError(t, err)

	//Read
	readTest := func() (interface{}, error) {
		c := core.NewBlockstorageClientForRegion(testRegionForBlockstorage)
		request := core.GetVolumeBackupRequest{}
		request.VolumeBackupID = r.ID
		rRead, err := c.GetVolumeBackup(context.Background(), request)
		return rRead, err
	}

	failIfError(t,
		retryUntilTrueOrError(readTest,
			checkLifecycleState(string(core.VOLUME_BACKUP_LIFECYCLE_STATE_AVAILABLE)), tick, time.After(120*time.Second)))

	//List
	listTest := func(t *testing.T) {
		c := core.NewBlockstorageClientForRegion(testRegionForBlockstorage)
		request := core.ListVolumeBackupsRequest{}
		r, err := c.ListVolumeBackups(context.Background(), request)
		failIfError(t, err)
		assert.True(t, len(r.Items) > 0)
		return
	}
	listTest(t)

	//Update
	updateTest := func(t *testing.T) {
		c := core.NewBlockstorageClientForRegion(testRegionForBlockstorage)
		request := core.UpdateVolumeBackupRequest{}
		request.VolumeBackupID = r.ID
		request.DisplayName = common.String("GoSDK2_TestBlockStorageVolumeBackupUpdate")
		r, err := c.UpdateVolumeBackup(context.Background(), request)
		assert.NotEmpty(t, r, fmt.Sprint(r))
		failIfError(t, err)
		return
	}
	updateTest(t)

	//Delete
	deleteTest := func(t *testing.T) {
		c := core.NewBlockstorageClientForRegion(testRegionForBlockstorage)
		request := core.DeleteVolumeBackupRequest{}
		request.VolumeBackupID = r.ID
		err := c.DeleteVolumeBackup(context.Background(), request)
		failIfError(t, err)
		return
	}

	deleteTest(t)
	return
}
