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
	"testing"

	"github.com/oracle/oci-go-sdk/common"
	"github.com/oracle/oci-go-sdk/core"
	"github.com/stretchr/testify/assert"
)

func TestComputeClient_AttachVnic(t *testing.T) {
	t.Skip("Not implemented")
	c, clerr := core.NewComputeClientWithConfigurationProvider(common.DefaultConfigProvider())
	failIfError(t, clerr)
	request := core.AttachVnicRequest{}
	r, err := c.AttachVnic(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

func TestComputeClient_AttachVolume(t *testing.T) {
	t.Skip("Not implemented")
	c, clerr := core.NewComputeClientWithConfigurationProvider(common.DefaultConfigProvider())
	failIfError(t, clerr)
	request := core.AttachVolumeRequest{}
	r, err := c.AttachVolume(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

func TestComputeClient_CaptureConsoleHistory(t *testing.T) {
	t.Skip("Not implemented")
	c, clerr := core.NewComputeClientWithConfigurationProvider(common.DefaultConfigProvider())
	failIfError(t, clerr)
	request := core.CaptureConsoleHistoryRequest{}
	r, err := c.CaptureConsoleHistory(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

func TestComputeClient_CreateImage(t *testing.T) {
	t.Skip("Not implemented")
	c, clerr := core.NewComputeClientWithConfigurationProvider(common.DefaultConfigProvider())
	failIfError(t, clerr)
	request := core.CreateImageRequest{}
	r, err := c.CreateImage(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

func TestComputeClient_CreateInstanceConsoleConnection(t *testing.T) {
	t.Skip("Not implemented")
	c, clerr := core.NewComputeClientWithConfigurationProvider(common.DefaultConfigProvider())
	failIfError(t, clerr)
	request := core.CreateInstanceConsoleConnectionRequest{}
	r, err := c.CreateInstanceConsoleConnection(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

func TestComputeClient_DeleteConsoleHistory(t *testing.T) {
	t.Skip("Not implemented")
	c, clerr := core.NewComputeClientWithConfigurationProvider(common.DefaultConfigProvider())
	failIfError(t, clerr)
	request := core.DeleteConsoleHistoryRequest{}
	_, err := c.DeleteConsoleHistory(context.Background(), request)
	assert.NoError(t, err)
	return
}

func TestComputeClient_DeleteImage(t *testing.T) {
	t.Skip("Not implemented")
	c, clerr := core.NewComputeClientWithConfigurationProvider(common.DefaultConfigProvider())
	failIfError(t, clerr)
	request := core.DeleteImageRequest{}
	_, err := c.DeleteImage(context.Background(), request)
	assert.NoError(t, err)
	return
}

func TestComputeClient_DeleteInstanceConsoleConnection(t *testing.T) {
	t.Skip("Not implemented")
	c, clerr := core.NewComputeClientWithConfigurationProvider(common.DefaultConfigProvider())
	failIfError(t, clerr)
	request := core.DeleteInstanceConsoleConnectionRequest{}
	_, err := c.DeleteInstanceConsoleConnection(context.Background(), request)
	assert.NoError(t, err)
	return
}

func TestComputeClient_DetachVnic(t *testing.T) {
	t.Skip("Not implemented")
	c, clerr := core.NewComputeClientWithConfigurationProvider(common.DefaultConfigProvider())
	failIfError(t, clerr)
	request := core.DetachVnicRequest{}
	_, err := c.DetachVnic(context.Background(), request)
	assert.NoError(t, err)
	return
}

func TestComputeClient_DetachVolume(t *testing.T) {
	t.Skip("Not implemented")
	c, clerr := core.NewComputeClientWithConfigurationProvider(common.DefaultConfigProvider())
	failIfError(t, clerr)
	request := core.DetachVolumeRequest{}
	_, err := c.DetachVolume(context.Background(), request)
	assert.NoError(t, err)
	return
}

func TestComputeClient_ExportImage(t *testing.T) {
	t.Skip("Not implemented")
	c, clerr := core.NewComputeClientWithConfigurationProvider(common.DefaultConfigProvider())
	failIfError(t, clerr)
	request := core.ExportImageRequest{}
	r, err := c.ExportImage(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

func TestComputeClient_GetConsoleHistory(t *testing.T) {
	history := captureOrGetConsoleHistory(t)
	assert.NotEmpty(t, history, fmt.Sprint(history))

	r := getConsoleHistory(t, history.Id)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	return
}

func TestComputeClient_GetConsoleHistoryContent(t *testing.T) {
	t.Skip("Not implemented")
	c, clerr := core.NewComputeClientWithConfigurationProvider(common.DefaultConfigProvider())
	failIfError(t, clerr)
	request := core.GetConsoleHistoryContentRequest{}
	r, err := c.GetConsoleHistoryContent(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

func TestComputeClient_GetImage(t *testing.T) {
	t.Skip("Not implemented")
	c, clerr := core.NewComputeClientWithConfigurationProvider(common.DefaultConfigProvider())
	failIfError(t, clerr)
	request := core.GetImageRequest{}
	r, err := c.GetImage(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

func TestComputeClient_GetInstance(t *testing.T) {
	t.Skip("Not implemented")
	c, clerr := core.NewComputeClientWithConfigurationProvider(common.DefaultConfigProvider())
	failIfError(t, clerr)
	request := core.GetInstanceRequest{}
	r, err := c.GetInstance(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

func TestComputeClient_GetInstanceConsoleConnection(t *testing.T) {
	consoleConnection := createOrGetInstanceConsoleConnection(t)
	assert.NotEmpty(t, consoleConnection)
	assert.NotEmpty(t, consoleConnection.Id)

	c, clerr := core.NewComputeClientWithConfigurationProvider(common.DefaultConfigProvider())
	failIfError(t, clerr)

	request := core.GetInstanceConsoleConnectionRequest{}
	request.InstanceConsoleConnectionId = consoleConnection.Id

	r, err := c.GetInstanceConsoleConnection(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

func TestComputeClient_GetVnicAttachment(t *testing.T) {
	t.Skip("Not implemented")
	c, clerr := core.NewComputeClientWithConfigurationProvider(common.DefaultConfigProvider())
	failIfError(t, clerr)
	request := core.GetVnicAttachmentRequest{}
	r, err := c.GetVnicAttachment(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

func TestComputeClient_GetVolumeAttachment(t *testing.T) {
	t.Skip("Not implemented")
	c, clerr := core.NewComputeClientWithConfigurationProvider(common.DefaultConfigProvider())
	failIfError(t, clerr)
	request := core.GetVolumeAttachmentRequest{}
	r, err := c.GetVolumeAttachment(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

func TestComputeClient_GetWindowsInstanceInitialCredentials(t *testing.T) {
	t.Skip("Not implemented")
	c, clerr := core.NewComputeClientWithConfigurationProvider(common.DefaultConfigProvider())
	failIfError(t, clerr)
	request := core.GetWindowsInstanceInitialCredentialsRequest{}
	r, err := c.GetWindowsInstanceInitialCredentials(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

func TestComputeClient_InstanceAction(t *testing.T) {
	t.Skip("Not implemented")
	c, clerr := core.NewComputeClientWithConfigurationProvider(common.DefaultConfigProvider())
	failIfError(t, clerr)
	request := core.InstanceActionRequest{}
	r, err := c.InstanceAction(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

func TestComputeClient_LaunchInstance(t *testing.T) {
	c, clerr := core.NewComputeClientWithConfigurationProvider(common.DefaultConfigProvider())
	failIfError(t, clerr)

	request := core.LaunchInstanceRequest{}
	request.CompartmentId = common.String(getCompartmentID())
	request.DisplayName = common.String("GOSDK2_Test_Instance")
	request.AvailabilityDomain = common.String(validAD())
	request.SubnetId = createOrGetSubnet(t).Id

	// search image by display name to make integration test running more relaible
	// i.e. ServiceLimitExeceed error etc...
	images := listImagesByDisplayName(t, common.String("Oracle-Linux-7.4-2018.01.20-0"))
	assert.NotEmpty(t, images)
	request.ImageId = images[0].Id

	shapes := listShapesForImage(t, request.ImageId)
	assert.NotEmpty(t, shapes)

	// if more shapes return, use different one to avoid getting service limited error
	if len(shapes) > 0 {
		request.Shape = shapes[1].Shape
	} else {
		request.Shape = shapes[0].Shape
	}

	r, err := c.LaunchInstance(context.Background(), request)
	failIfError(t, err)
	assert.NotEmpty(t, r, fmt.Sprint(r))

	// if we've successfully created a instance during testing, make sure that we delete it
	defer func(client core.ComputeClient, r core.Instance) {
		// delete instance
		request := core.TerminateInstanceRequest{
			InstanceId: r.Id,
		}

		resp, err := client.TerminateInstance(context.Background(), request)
		failIfError(t, err)
		assert.NotEmpty(t, resp.OpcRequestId)
	}(c, r.Instance)

	return
}

func TestComputeClient_ListConsoleHistories(t *testing.T) {
	t.Skip("Not implemented")
	c, clerr := core.NewComputeClientWithConfigurationProvider(common.DefaultConfigProvider())
	failIfError(t, clerr)
	request := core.ListConsoleHistoriesRequest{}
	r, err := c.ListConsoleHistories(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

func TestComputeClient_ListImages(t *testing.T) {
	images := listImages(t)
	assert.NotEmpty(t, images, fmt.Sprint(images))
	return
}

func TestComputeClient_ListInstanceConsoleConnections(t *testing.T) {
	t.Skip("Not implemented")
	c, clerr := core.NewComputeClientWithConfigurationProvider(common.DefaultConfigProvider())
	failIfError(t, clerr)
	request := core.ListInstanceConsoleConnectionsRequest{}
	r, err := c.ListInstanceConsoleConnections(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

func TestComputeClient_ListInstances(t *testing.T) {
	t.Skip("Not implemented")
	c, clerr := core.NewComputeClientWithConfigurationProvider(common.DefaultConfigProvider())
	failIfError(t, clerr)
	request := core.ListInstancesRequest{}
	r, err := c.ListInstances(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

func TestComputeClient_ListShapes(t *testing.T) {
	shapes := listShapes(t)
	assert.NotEmpty(t, shapes, fmt.Sprint(shapes))
	return
}

func TestComputeClient_ListVnicAttachments(t *testing.T) {
	t.Skip("Not implemented")
	c, clerr := core.NewComputeClientWithConfigurationProvider(common.DefaultConfigProvider())
	failIfError(t, clerr)
	request := core.ListVnicAttachmentsRequest{}
	r, err := c.ListVnicAttachments(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

func TestComputeClient_ListVolumeAttachments(t *testing.T) {
	c, clerr := core.NewComputeClientWithConfigurationProvider(common.DefaultConfigProvider())
	failIfError(t, clerr)

	// make sure volume is created
	volume := createOrGetVolume(t)
	instance := createOrGetInstance(t)

	listAttachedVolumes := func() []core.VolumeAttachment {
		request := core.ListVolumeAttachmentsRequest{
			CompartmentId: common.String(getCompartmentID()),
			InstanceId:    instance.Id,
		}
		r, err := c.ListVolumeAttachments(context.Background(), request)
		assert.NotEmpty(t, r, fmt.Sprint(r))
		assert.NoError(t, err)
		return r.Items
	}

	// get list of attached volumes for current instance
	attachedVolumes := listAttachedVolumes()

	// if no volume attached, attach one
	if attachedVolumes == nil ||
		len(attachedVolumes) == 0 {
		// attach volume to instance
		attachRequest := core.AttachVolumeRequest{}
		attachRequest.AttachVolumeDetails = core.AttachIScsiVolumeDetails{
			InstanceId: instance.Id,
			VolumeId:   volume.Id,
		}

		_, err := c.AttachVolume(context.Background(), attachRequest)
		failIfError(t, err)
		attachedVolumes = listAttachedVolumes()
	}

	assert.NotEmpty(t, attachedVolumes, fmt.Sprint(attachedVolumes))
	return
}

func TestComputeClient_ListBootVolumeAttachments(t *testing.T) {
	bootVolumeAttachments := listBootVolumeAttachments(t)
	assert.NotEmpty(t, bootVolumeAttachments)
	return
}

func TestBlockstorageClient_GetBootVolumeAttachment(t *testing.T) {
	bootVolumeAttachments := listBootVolumeAttachments(t)

	c, clerr := core.NewComputeClientWithConfigurationProvider(common.DefaultConfigProvider())
	failIfError(t, clerr)
	request := core.GetBootVolumeAttachmentRequest{
		BootVolumeAttachmentId: bootVolumeAttachments[0].Id,
	}

	r, err := c.GetBootVolumeAttachment(context.Background(), request)
	failIfError(t, err)
	assert.NotEmpty(t, r.Id)
	return
}

func TestComputeClient_TerminateInstance(t *testing.T) {
	t.Skip("Not implemented")
	c, clerr := core.NewComputeClientWithConfigurationProvider(common.DefaultConfigProvider())
	failIfError(t, clerr)
	request := core.TerminateInstanceRequest{}
	_, err := c.TerminateInstance(context.Background(), request)
	assert.NoError(t, err)
	return
}

func TestComputeClient_UpdateConsoleHistory(t *testing.T) {
	t.Skip("Not implemented")
	c, clerr := core.NewComputeClientWithConfigurationProvider(common.DefaultConfigProvider())
	failIfError(t, clerr)
	request := core.UpdateConsoleHistoryRequest{}
	r, err := c.UpdateConsoleHistory(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

func TestComputeClient_UpdateImage(t *testing.T) {
	t.Skip("Not implemented")
	c, clerr := core.NewComputeClientWithConfigurationProvider(common.DefaultConfigProvider())
	failIfError(t, clerr)
	request := core.UpdateImageRequest{}
	r, err := c.UpdateImage(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

func TestComputeClient_UpdateInstance(t *testing.T) {
	t.Skip("Not implemented")
	c, clerr := core.NewComputeClientWithConfigurationProvider(common.DefaultConfigProvider())
	failIfError(t, clerr)
	request := core.UpdateInstanceRequest{}
	r, err := c.UpdateInstance(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}
