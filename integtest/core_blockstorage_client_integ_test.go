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
)

var (
	testRegionForBlockstorage = common.REGION_PHX
)

func TestBlockstorageClient_CreateVolume(t *testing.T) {
	t.Skip("Not implemented")
	c := core.NewBlockstorageClientForRegion(testRegionForBlockstorage)
	request := core.CreateVolumeRequest{}
	r, err := c.CreateVolume(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
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

func TestBlockstorageClient_DeleteBootVolume(t *testing.T) {
	t.Skip("Not implemented")
	c := core.NewBlockstorageClientForRegion(testRegionForBlockstorage)
	request := core.DeleteBootVolumeRequest{}
	err := c.DeleteBootVolume(context.Background(), request)
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

func TestBlockstorageClient_GetBootVolume(t *testing.T) {
	t.Skip("Not implemented")
	c := core.NewBlockstorageClientForRegion(testRegionForBlockstorage)
	request := core.GetBootVolumeRequest{}
	r, err := c.GetBootVolume(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

func TestBlockstorageClient_GetVolume(t *testing.T) {
	t.Skip("Not implemented")
	c := core.NewBlockstorageClientForRegion(testRegionForBlockstorage)
	request := core.GetVolumeRequest{}
	r, err := c.GetVolume(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
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

func TestBlockstorageClient_ListBootVolumes(t *testing.T) {
	t.Skip("Not implemented")
	c := core.NewBlockstorageClientForRegion(testRegionForBlockstorage)
	request := core.ListBootVolumesRequest{}
	r, err := c.ListBootVolumes(context.Background(), request)
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
	t.Skip("Not implemented")
	c := core.NewBlockstorageClientForRegion(testRegionForBlockstorage)
	request := core.ListVolumesRequest{}
	r, err := c.ListVolumes(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

func TestBlockstorageClient_UpdateBootVolume(t *testing.T) {
	t.Skip("Not implemented")
	c := core.NewBlockstorageClientForRegion(testRegionForBlockstorage)
	request := core.UpdateBootVolumeRequest{}
	r, err := c.UpdateBootVolume(context.Background(), request)
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
