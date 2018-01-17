// Copyright (c) 2016, 2017, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Core Services API
//
// APIs for Networking Service, Compute Service, and Block Volume Service.
//

package integtest

import (
	"context"
	"fmt"
	"github.com/oracle/oci-go-sdk/common"
	"github.com/oracle/oci-go-sdk/core"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

var (
	tick    = time.Tick(500 * time.Millisecond)
	timeout = time.After(30 * time.Second)
)

//Volumes CRUDL
func TestBlockstorageClient_CreateVolume(t *testing.T) {

	c, clerr := core.NewBlockstorageClientWithConfigurationProvider(common.DefaultConfigProvider())
	failIfError(t, clerr)
	request := core.CreateVolumeRequest{}

	request.AvailabilityDomain = common.String(validAD())
	request.CompartmentId = common.String(getTenancyID())
	request.DisplayName = common.String("GoSDK2_TestBlockStorageCreateVolumeDisplayName")

	resCreate, err := c.CreateVolume(context.Background(), request)
	assert.NoError(t, err)
	assert.NotEmpty(t, resCreate.Id)

	//Read
	readTest := func() (interface{}, error) {
		c, clerr := core.NewBlockstorageClientWithConfigurationProvider(common.DefaultConfigProvider())
		failIfError(t, clerr)
		request := core.GetVolumeRequest{}
		request.VolumeId = resCreate.Id
		resRead, err := c.GetVolume(context.Background(), request)
		return resRead, err
	}
	assert.NoError(t,
		retryUntilTrueOrError(readTest,
			checkLifecycleState(string(core.VolumeLifecycleStateAvailable)), tick, timeout))

	//update
	updateTest := func(t *testing.T) {
		c, clerr := core.NewBlockstorageClientWithConfigurationProvider(common.DefaultConfigProvider())
		failIfError(t, clerr)
		request := core.UpdateVolumeRequest{}
		request.VolumeId = resCreate.Id
		request.DisplayName = common.String("GoSDK2_TestBlockStorageCreateVolumeDisplayNameUpdated")
		r, err := c.UpdateVolume(context.Background(), request)
		assert.NotEmpty(t, r, fmt.Sprint(r))
		assert.NoError(t, err)
		return
	}
	updateTest(t)

	//list
	listTest := func(t *testing.T) {
		c, clerr := core.NewBlockstorageClientWithConfigurationProvider(common.DefaultConfigProvider())
		failIfError(t, clerr)
		request := core.ListVolumesRequest{}
		request.CompartmentId = common.String(getTenancyID())
		r, err := c.ListVolumes(context.Background(), request)
		assert.NoError(t, err)
		assert.NotEmpty(t, r.Items)
		return
	}
	listTest(t)

	//delete
	deleteTest := func(t *testing.T) {
		c, clerr := core.NewBlockstorageClientWithConfigurationProvider(common.DefaultConfigProvider())
		failIfError(t, clerr)
		request := core.DeleteVolumeRequest{}
		request.VolumeId = resCreate.Id
		err := c.DeleteVolume(context.Background(), request)
		assert.NoError(t, err)
		return
	}
	deleteTest(t)

	return
}

func TestBlockstorageClient_CreateVolumeBackup(t *testing.T) {

	c, clerr := core.NewBlockstorageClientWithConfigurationProvider(common.DefaultConfigProvider())
	failIfError(t, clerr)

	var volumeId *string
	//First create a volume
	createVol := func(t *testing.T) {
		rCreateVol := core.CreateVolumeRequest{}
		rCreateVol.AvailabilityDomain = common.String(validAD())
		rCreateVol.CompartmentId = common.String(getTenancyID())
		rCreateVol.DisplayName = common.String("GoSDK2_TestBlockStorageCreateVolumeForBackup")
		resCreate, err := c.CreateVolume(context.Background(), rCreateVol)
		volumeId = resCreate.Id
		failIfError(t, err)
	}

	deleteVol := func() {
		c, clerr := core.NewBlockstorageClientWithConfigurationProvider(common.DefaultConfigProvider())
		failIfError(t, clerr)
		request := core.DeleteVolumeRequest{}
		request.VolumeId = volumeId
		c.DeleteVolume(context.Background(), request)
		return
	}

	readVol := func() (interface{}, error) {
		c, clerr := core.NewBlockstorageClientWithConfigurationProvider(common.DefaultConfigProvider())
		failIfError(t, clerr)
		request := core.GetVolumeRequest{}
		request.VolumeId = volumeId
		resRead, err := c.GetVolume(context.Background(), request)
		return resRead, err
	}

	createVol(t)
	defer deleteVol()
	failIfError(t,
		retryUntilTrueOrError(readVol,
			checkLifecycleState(string(core.VolumeLifecycleStateAvailable)), tick, time.After(120*time.Second)))

	//Create
	request := core.CreateVolumeBackupRequest{}
	request.VolumeId = volumeId
	request.DisplayName = common.String("GoSDK2_TestBlockStorageCreateVolumeBackup")
	r, err := c.CreateVolumeBackup(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	failIfError(t, err)

	//Read
	readTest := func() (interface{}, error) {
		c, clerr := core.NewBlockstorageClientWithConfigurationProvider(common.DefaultConfigProvider())
		failIfError(t, clerr)
		request := core.GetVolumeBackupRequest{}
		request.VolumeBackupId = r.Id
		rRead, err := c.GetVolumeBackup(context.Background(), request)
		return rRead, err
	}

	failIfError(t,
		retryUntilTrueOrError(readTest,
			checkLifecycleState(string(core.VolumeBackupLifecycleStateAvailable)), tick, time.After(120*time.Second)))

	//List
	listTest := func(t *testing.T) {
		c, clerr := core.NewBlockstorageClientWithConfigurationProvider(common.DefaultConfigProvider())
		failIfError(t, clerr)
		request := core.ListVolumeBackupsRequest{}
		request.CompartmentId = common.String(getTenancyID())
		request.VolumeId = volumeId
		r, err := c.ListVolumeBackups(context.Background(), request)
		failIfError(t, err)
		assert.True(t, len(r.Items) > 0)
		return
	}
	listTest(t)

	//Update
	updateTest := func(t *testing.T) {
		c, clerr := core.NewBlockstorageClientWithConfigurationProvider(common.DefaultConfigProvider())
		failIfError(t, clerr)
		request := core.UpdateVolumeBackupRequest{}
		request.VolumeBackupId = r.Id
		request.DisplayName = common.String("GoSDK2_TestBlockStorageVolumeBackupUpdate")
		r, err := c.UpdateVolumeBackup(context.Background(), request)
		assert.NotEmpty(t, r, fmt.Sprint(r))
		failIfError(t, err)
		return
	}
	updateTest(t)

	//Delete
	deleteTest := func(t *testing.T) {
		c, clerr := core.NewBlockstorageClientWithConfigurationProvider(common.DefaultConfigProvider())
		failIfError(t, clerr)
		request := core.DeleteVolumeBackupRequest{}
		request.VolumeBackupId = r.Id
		err := c.DeleteVolumeBackup(context.Background(), request)
		failIfError(t, err)
		return
	}

	deleteTest(t)
	return
}
