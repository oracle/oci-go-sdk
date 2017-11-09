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
)

var (
	testRegionForBlockstorage = common.REGION_PHX
	validOCID =  "ocidv1:tenancy:oc1:phx:1460406592660:aaaaaaaab4faofrfkxecohhjuivjq262pu"
	validAD = "kIdk:PHX-AD-2"
)

//Volumes CRUD
func TestBlockstorageClient_CreateVolume(t *testing.T) {

	c := core.NewBlockstorageClientForRegion(testRegionForBlockstorage)
	request := core.CreateVolumeRequest{}

	request.AvailabilityDomain=common.String(validAD)
	request.CompartmentID=common.String(validOCID)
	request.DisplayName = common.String("GoSDK2_CreateVolumeDisplayName")

	resCreate, err := c.CreateVolume(context.Background(), request)
	assert.NoError(t, err)
	assert.NotEmpty(t, resCreate.ID)

	//Read
	readTest := func (t *testing.T) {
		c := core.NewBlockstorageClientForRegion(testRegionForBlockstorage)
		request := core.GetVolumeRequest{}
		request.VolumeID = resCreate.ID
		resRead, err := c.GetVolume(context.Background(), request)
		assert.NoError(t, err)
		assert.Equal(t, resCreate.ID, resRead.ID)
		return
	}

	readTest(t)
	return
}

func TestBlockstorageClient_CreateVolumeBackup(t *testing.T) {
	t.Skip("Not implemented")
	c := core.NewBlockstorageClientForRegion(testRegionForBlockstorage)
	request := core.CreateVolumeBackupRequest{}
	r, err := c.CreateVolumeBackup(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

func TestBlockstorageClient_DeleteVolume(t *testing.T) {
	t.Skip("Not implemented")
	c := core.NewBlockstorageClientForRegion(testRegionForBlockstorage)
	request := core.DeleteVolumeRequest{}
	err := c.DeleteVolume(context.Background(), request)
	assert.NoError(t, err)
	return
}

func TestBlockstorageClient_DeleteVolumeBackup(t *testing.T) {
	t.Skip("Not implemented")
	c := core.NewBlockstorageClientForRegion(testRegionForBlockstorage)
	request := core.DeleteVolumeBackupRequest{}
	err := c.DeleteVolumeBackup(context.Background(), request)
	assert.NoError(t, err)
	return
}


func TestBlockstorageClient_GetVolumeBackup(t *testing.T) {
	t.Skip("Not implemented")
	c := core.NewBlockstorageClientForRegion(testRegionForBlockstorage)
	request := core.GetVolumeBackupRequest{}
	r, err := c.GetVolumeBackup(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

func TestBlockstorageClient_ListVolumeBackups(t *testing.T) {
	t.Skip("Not implemented")
	c := core.NewBlockstorageClientForRegion(testRegionForBlockstorage)
	request := core.ListVolumeBackupsRequest{}
	r, err := c.ListVolumeBackups(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

func TestBlockstorageClient_ListVolumes(t *testing.T) {
	c := core.NewBlockstorageClientForRegion(testRegionForBlockstorage)
	request := core.ListVolumesRequest{}
	request.CompartmentID=common.String("ocidv1:tenancy:oc1:phx:1460406592660:aaaaaaaab4faofrfkxecohhjuivjq262pu")
	r, err := c.ListVolumes(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

func TestBlockstorageClient_UpdateVolume(t *testing.T) {
	t.Skip("Not implemented")
	c := core.NewBlockstorageClientForRegion(testRegionForBlockstorage)
	request := core.UpdateVolumeRequest{}
	r, err := c.UpdateVolume(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

func TestBlockstorageClient_UpdateVolumeBackup(t *testing.T) {
	t.Skip("Not implemented")
	c := core.NewBlockstorageClientForRegion(testRegionForBlockstorage)
	request := core.UpdateVolumeBackupRequest{}
	r, err := c.UpdateVolumeBackup(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}