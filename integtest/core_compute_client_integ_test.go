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
	testRegionForCompute = common.REGION_PHX
)

func TestComputeClient_AddImageShapeCompatibilityEntry(t *testing.T) {
	t.Skip("Not implemented")
	c := core.NewComputeClientForRegion(testRegionForCompute)
	request := core.AddImageShapeCompatibilityEntryRequest{}
	r, err := c.AddImageShapeCompatibilityEntry(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

func TestComputeClient_AttachBootVolume(t *testing.T) {
	t.Skip("Not implemented")
	c := core.NewComputeClientForRegion(testRegionForCompute)
	request := core.AttachBootVolumeRequest{}
	r, err := c.AttachBootVolume(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

func TestComputeClient_AttachVnic(t *testing.T) {
	t.Skip("Not implemented")
	c := core.NewComputeClientForRegion(testRegionForCompute)
	request := core.AttachVnicRequest{}
	r, err := c.AttachVnic(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

func TestComputeClient_AttachVolume(t *testing.T) {
	t.Skip("Not implemented")
	c := core.NewComputeClientForRegion(testRegionForCompute)
	request := core.AttachVolumeRequest{}
	r, err := c.AttachVolume(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

func TestComputeClient_CaptureConsoleHistory(t *testing.T) {
	t.Skip("Not implemented")
	c := core.NewComputeClientForRegion(testRegionForCompute)
	request := core.CaptureConsoleHistoryRequest{}
	r, err := c.CaptureConsoleHistory(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

func TestComputeClient_CreateImage(t *testing.T) {
	t.Skip("Not implemented")
	c := core.NewComputeClientForRegion(testRegionForCompute)
	request := core.CreateImageRequest{}
	r, err := c.CreateImage(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

func TestComputeClient_CreateInstanceConsoleConnection(t *testing.T) {
	t.Skip("Not implemented")
	c := core.NewComputeClientForRegion(testRegionForCompute)
	request := core.CreateInstanceConsoleConnectionRequest{}
	r, err := c.CreateInstanceConsoleConnection(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

func TestComputeClient_DeleteConsoleHistory(t *testing.T) {
	t.Skip("Not implemented")
	c := core.NewComputeClientForRegion(testRegionForCompute)
	request := core.DeleteConsoleHistoryRequest{}
	err := c.DeleteConsoleHistory(context.Background(), request)
	assert.NoError(t, err)
	return
}

func TestComputeClient_DeleteImage(t *testing.T) {
	t.Skip("Not implemented")
	c := core.NewComputeClientForRegion(testRegionForCompute)
	request := core.DeleteImageRequest{}
	err := c.DeleteImage(context.Background(), request)
	assert.NoError(t, err)
	return
}

func TestComputeClient_DeleteInstanceConsoleConnection(t *testing.T) {
	t.Skip("Not implemented")
	c := core.NewComputeClientForRegion(testRegionForCompute)
	request := core.DeleteInstanceConsoleConnectionRequest{}
	err := c.DeleteInstanceConsoleConnection(context.Background(), request)
	assert.NoError(t, err)
	return
}

func TestComputeClient_DetachBootVolume(t *testing.T) {
	t.Skip("Not implemented")
	c := core.NewComputeClientForRegion(testRegionForCompute)
	request := core.DetachBootVolumeRequest{}
	err := c.DetachBootVolume(context.Background(), request)
	assert.NoError(t, err)
	return
}

func TestComputeClient_DetachVnic(t *testing.T) {
	t.Skip("Not implemented")
	c := core.NewComputeClientForRegion(testRegionForCompute)
	request := core.DetachVnicRequest{}
	err := c.DetachVnic(context.Background(), request)
	assert.NoError(t, err)
	return
}

func TestComputeClient_DetachVolume(t *testing.T) {
	t.Skip("Not implemented")
	c := core.NewComputeClientForRegion(testRegionForCompute)
	request := core.DetachVolumeRequest{}
	err := c.DetachVolume(context.Background(), request)
	assert.NoError(t, err)
	return
}

func TestComputeClient_ExportImage(t *testing.T) {
	t.Skip("Not implemented")
	c := core.NewComputeClientForRegion(testRegionForCompute)
	request := core.ExportImageRequest{}
	r, err := c.ExportImage(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

func TestComputeClient_GetBootVolumeAttachment(t *testing.T) {
	t.Skip("Not implemented")
	c := core.NewComputeClientForRegion(testRegionForCompute)
	request := core.GetBootVolumeAttachmentRequest{}
	r, err := c.GetBootVolumeAttachment(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

func TestComputeClient_GetConsoleHistory(t *testing.T) {
	t.Skip("Not implemented")
	c := core.NewComputeClientForRegion(testRegionForCompute)
	request := core.GetConsoleHistoryRequest{}
	r, err := c.GetConsoleHistory(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

func TestComputeClient_GetConsoleHistoryContent(t *testing.T) {
	t.Skip("Not implemented")
	c := core.NewComputeClientForRegion(testRegionForCompute)
	request := core.GetConsoleHistoryContentRequest{}
	r, err := c.GetConsoleHistoryContent(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

func TestComputeClient_GetImage(t *testing.T) {
	t.Skip("Not implemented")
	c := core.NewComputeClientForRegion(testRegionForCompute)
	request := core.GetImageRequest{}
	r, err := c.GetImage(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

func TestComputeClient_GetInstance(t *testing.T) {
	t.Skip("Not implemented")
	c := core.NewComputeClientForRegion(testRegionForCompute)
	request := core.GetInstanceRequest{}
	r, err := c.GetInstance(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

func TestComputeClient_GetInstanceConsoleConnection(t *testing.T) {
	t.Skip("Not implemented")
	c := core.NewComputeClientForRegion(testRegionForCompute)
	request := core.GetInstanceConsoleConnectionRequest{}
	r, err := c.GetInstanceConsoleConnection(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

func TestComputeClient_GetVnicAttachment(t *testing.T) {
	t.Skip("Not implemented")
	c := core.NewComputeClientForRegion(testRegionForCompute)
	request := core.GetVnicAttachmentRequest{}
	r, err := c.GetVnicAttachment(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

func TestComputeClient_GetVolumeAttachment(t *testing.T) {
	t.Skip("Not implemented")
	c := core.NewComputeClientForRegion(testRegionForCompute)
	request := core.GetVolumeAttachmentRequest{}
	r, err := c.GetVolumeAttachment(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

func TestComputeClient_GetWindowsInstanceInitialCredentials(t *testing.T) {
	t.Skip("Not implemented")
	c := core.NewComputeClientForRegion(testRegionForCompute)
	request := core.GetWindowsInstanceInitialCredentialsRequest{}
	r, err := c.GetWindowsInstanceInitialCredentials(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

func TestComputeClient_InstanceAction(t *testing.T) {
	t.Skip("Not implemented")
	c := core.NewComputeClientForRegion(testRegionForCompute)
	request := core.InstanceActionRequest{}
	r, err := c.InstanceAction(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

func TestComputeClient_LaunchInstance(t *testing.T) {
	t.Skip("Not implemented")
	c := core.NewComputeClientForRegion(testRegionForCompute)
	request := core.LaunchInstanceRequest{}
	r, err := c.LaunchInstance(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

func TestComputeClient_ListBootVolumeAttachments(t *testing.T) {
	t.Skip("Not implemented")
	c := core.NewComputeClientForRegion(testRegionForCompute)
	request := core.ListBootVolumeAttachmentsRequest{}
	r, err := c.ListBootVolumeAttachments(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

func TestComputeClient_ListConsoleHistories(t *testing.T) {
	t.Skip("Not implemented")
	c := core.NewComputeClientForRegion(testRegionForCompute)
	request := core.ListConsoleHistoriesRequest{}
	r, err := c.ListConsoleHistories(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

func TestComputeClient_ListImages(t *testing.T) {
	t.Skip("Not implemented")
	c := core.NewComputeClientForRegion(testRegionForCompute)
	request := core.ListImagesRequest{}
	r, err := c.ListImages(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

func TestComputeClient_ListInstanceConsoleConnections(t *testing.T) {
	t.Skip("Not implemented")
	c := core.NewComputeClientForRegion(testRegionForCompute)
	request := core.ListInstanceConsoleConnectionsRequest{}
	r, err := c.ListInstanceConsoleConnections(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

func TestComputeClient_ListInstances(t *testing.T) {
	t.Skip("Not implemented")
	c := core.NewComputeClientForRegion(testRegionForCompute)
	request := core.ListInstancesRequest{}
	r, err := c.ListInstances(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

func TestComputeClient_ListShapes(t *testing.T) {
	t.Skip("Not implemented")
	c := core.NewComputeClientForRegion(testRegionForCompute)
	request := core.ListShapesRequest{}
	r, err := c.ListShapes(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

func TestComputeClient_ListVnicAttachments(t *testing.T) {
	t.Skip("Not implemented")
	c := core.NewComputeClientForRegion(testRegionForCompute)
	request := core.ListVnicAttachmentsRequest{}
	r, err := c.ListVnicAttachments(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

func TestComputeClient_ListVolumeAttachments(t *testing.T) {
	t.Skip("Not implemented")
	c := core.NewComputeClientForRegion(testRegionForCompute)
	request := core.ListVolumeAttachmentsRequest{}
	r, err := c.ListVolumeAttachments(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

func TestComputeClient_RemoveImageShapeCompatibilityEntry(t *testing.T) {
	t.Skip("Not implemented")
	c := core.NewComputeClientForRegion(testRegionForCompute)
	request := core.RemoveImageShapeCompatibilityEntryRequest{}
	err := c.RemoveImageShapeCompatibilityEntry(context.Background(), request)
	assert.NoError(t, err)
	return
}

func TestComputeClient_TerminateInstance(t *testing.T) {
	t.Skip("Not implemented")
	c := core.NewComputeClientForRegion(testRegionForCompute)
	request := core.TerminateInstanceRequest{}
	err := c.TerminateInstance(context.Background(), request)
	assert.NoError(t, err)
	return
}

func TestComputeClient_UpdateConsoleHistory(t *testing.T) {
	t.Skip("Not implemented")
	c := core.NewComputeClientForRegion(testRegionForCompute)
	request := core.UpdateConsoleHistoryRequest{}
	r, err := c.UpdateConsoleHistory(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

func TestComputeClient_UpdateImage(t *testing.T) {
	t.Skip("Not implemented")
	c := core.NewComputeClientForRegion(testRegionForCompute)
	request := core.UpdateImageRequest{}
	r, err := c.UpdateImage(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

func TestComputeClient_UpdateInstance(t *testing.T) {
	t.Skip("Not implemented")
	c := core.NewComputeClientForRegion(testRegionForCompute)
	request := core.UpdateInstanceRequest{}
	r, err := c.UpdateInstance(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}
