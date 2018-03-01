// Copyright (c) 2016, 2017, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Core Services API
//
// APIs for Networking Service, Compute Service, and Block Volume Service.
//

package core

import (
	"context"
	"fmt"
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

//BlockstorageClient a client for Blockstorage
type BlockstorageClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewBlockstorageClientWithConfigurationProvider Creates a new default Blockstorage client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewBlockstorageClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client BlockstorageClient, err error) {
	baseClient, err := common.NewClientWithConfig(configProvider)
	if err != nil {
		return
	}

	client = BlockstorageClient{BaseClient: baseClient}
	client.BasePath = "20160918"
	err = client.setConfigurationProvider(configProvider)
	return
}

// SetRegion overrides the region of this client.
func (client *BlockstorageClient) SetRegion(region string) {
	client.Host = fmt.Sprintf(common.DefaultHostURLTemplate, "iaas", region)
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *BlockstorageClient) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
	if ok, err := common.IsConfigurationProviderValid(configProvider); !ok {
		return err
	}

	// Error has been checked already
	region, _ := configProvider.Region()
	client.config = &configProvider
	client.SetRegion(region)
	return nil
}

// ConfigurationProvider the ConfigurationProvider used in this client, or null if none set
func (client *BlockstorageClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// CreateVolume Creates a new volume in the specified compartment. Volumes can be created in sizes ranging from
// 50 GB (51200 MB) to 16 TB (16777216 MB), in 1 GB (1024 MB) increments. By default, volumes are 1 TB (1048576 MB).
// For general information about block volumes, see
// [Overview of Block Volume Service]({{DOC_SERVER_URL}}/Content/Block/Concepts/overview.htm).
// A volume and instance can be in separate compartments but must be in the same Availability Domain.
// For information about access control and compartments, see
// [Overview of the IAM Service]({{DOC_SERVER_URL}}/Content/Identity/Concepts/overview.htm). For information about
// Availability Domains, see [Regions and Availability Domains]({{DOC_SERVER_URL}}/Content/General/Concepts/regions.htm).
// To get a list of Availability Domains, use the `ListAvailabilityDomains` operation
// in the Identity and Access Management Service API.
// You may optionally specify a *display name* for the volume, which is simply a friendly name or
// description. It does not have to be unique, and you can change it. Avoid entering confidential information.
func (client BlockstorageClient) CreateVolume(ctx context.Context, request CreateVolumeRequest) (response CreateVolumeResponse, err error) {
	ociResponse, e := client.Retry(ctx, request, client.createVolume, request.GetRetryPolicy())
	if e != nil {
		err = e
		return
	}
	if convertedResponse, convertedResponseOk := ociResponse.(CreateVolumeResponse); convertedResponseOk {
		response = convertedResponse
	}
	return
}

// lower camel case request => not exported
func (client BlockstorageClient) createVolume(ctx context.Context, request common.OciRequest) (common.OciResponse, error) {
	httpRequest, err := request.GetHttpRequest(http.MethodPost, "/volumes")
	if err != nil {
		return nil, err
	}

	var response CreateVolumeResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, nil
}

// CreateVolumeBackup Creates a new backup of the specified volume. For general information about volume backups,
// see [Overview of Block Volume Service Backups]({{DOC_SERVER_URL}}/Content/Block/Concepts/blockvolumebackups.htm)
// When the request is received, the backup object is in a REQUEST_RECEIVED state.
// When the data is imaged, it goes into a CREATING state.
// After the backup is fully uploaded to the cloud, it goes into an AVAILABLE state.
func (client BlockstorageClient) CreateVolumeBackup(ctx context.Context, request CreateVolumeBackupRequest) (response CreateVolumeBackupResponse, err error) {
	ociResponse, e := client.Retry(ctx, request, client.createVolumeBackup, request.GetRetryPolicy())
	if e != nil {
		err = e
		return
	}
	if convertedResponse, convertedResponseOk := ociResponse.(CreateVolumeBackupResponse); convertedResponseOk {
		response = convertedResponse
	}
	return
}

// lower camel case request => not exported
func (client BlockstorageClient) createVolumeBackup(ctx context.Context, request common.OciRequest) (common.OciResponse, error) {
	httpRequest, err := request.GetHttpRequest(http.MethodPost, "/volumeBackups")
	if err != nil {
		return nil, err
	}

	var response CreateVolumeBackupResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, nil
}

// DeleteBootVolume Deletes the specified boot volume. The volume cannot have an active connection to an instance.
// To disconnect the boot volume from a connected instance, see
// [Disconnecting From a Boot Volume]({{DOC_SERVER_URL}}/Content/Block/Tasks/deletingbootvolume.htm).
// **Warning:** All data on the boot volume will be permanently lost when the boot volume is deleted.
func (client BlockstorageClient) DeleteBootVolume(ctx context.Context, request DeleteBootVolumeRequest) (response DeleteBootVolumeResponse, err error) {
	ociResponse, e := client.Retry(ctx, request, client.deleteBootVolume, request.GetRetryPolicy())
	if e != nil {
		err = e
		return
	}
	if convertedResponse, convertedResponseOk := ociResponse.(DeleteBootVolumeResponse); convertedResponseOk {
		response = convertedResponse
	}
	return
}

// lower camel case request => not exported
func (client BlockstorageClient) deleteBootVolume(ctx context.Context, request common.OciRequest) (common.OciResponse, error) {
	httpRequest, err := request.GetHttpRequest(http.MethodDelete, "/bootVolumes/{bootVolumeId}")
	if err != nil {
		return nil, err
	}

	var response DeleteBootVolumeResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, nil
}

// DeleteVolume Deletes the specified volume. The volume cannot have an active connection to an instance.
// To disconnect the volume from a connected instance, see
// [Disconnecting From a Volume]({{DOC_SERVER_URL}}/Content/Block/Tasks/disconnectingfromavolume.htm).
// **Warning:** All data on the volume will be permanently lost when the volume is deleted.
func (client BlockstorageClient) DeleteVolume(ctx context.Context, request DeleteVolumeRequest) (response DeleteVolumeResponse, err error) {
	ociResponse, e := client.Retry(ctx, request, client.deleteVolume, request.GetRetryPolicy())
	if e != nil {
		err = e
		return
	}
	if convertedResponse, convertedResponseOk := ociResponse.(DeleteVolumeResponse); convertedResponseOk {
		response = convertedResponse
	}
	return
}

// lower camel case request => not exported
func (client BlockstorageClient) deleteVolume(ctx context.Context, request common.OciRequest) (common.OciResponse, error) {
	httpRequest, err := request.GetHttpRequest(http.MethodDelete, "/volumes/{volumeId}")
	if err != nil {
		return nil, err
	}

	var response DeleteVolumeResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, nil
}

// DeleteVolumeBackup Deletes a volume backup.
func (client BlockstorageClient) DeleteVolumeBackup(ctx context.Context, request DeleteVolumeBackupRequest) (response DeleteVolumeBackupResponse, err error) {
	ociResponse, e := client.Retry(ctx, request, client.deleteVolumeBackup, request.GetRetryPolicy())
	if e != nil {
		err = e
		return
	}
	if convertedResponse, convertedResponseOk := ociResponse.(DeleteVolumeBackupResponse); convertedResponseOk {
		response = convertedResponse
	}
	return
}

// lower camel case request => not exported
func (client BlockstorageClient) deleteVolumeBackup(ctx context.Context, request common.OciRequest) (common.OciResponse, error) {
	httpRequest, err := request.GetHttpRequest(http.MethodDelete, "/volumeBackups/{volumeBackupId}")
	if err != nil {
		return nil, err
	}

	var response DeleteVolumeBackupResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, nil
}

// GetBootVolume Gets information for the specified boot volume.
func (client BlockstorageClient) GetBootVolume(ctx context.Context, request GetBootVolumeRequest) (response GetBootVolumeResponse, err error) {
	ociResponse, e := client.Retry(ctx, request, client.getBootVolume, request.GetRetryPolicy())
	if e != nil {
		err = e
		return
	}
	if convertedResponse, convertedResponseOk := ociResponse.(GetBootVolumeResponse); convertedResponseOk {
		response = convertedResponse
	}
	return
}

// lower camel case request => not exported
func (client BlockstorageClient) getBootVolume(ctx context.Context, request common.OciRequest) (common.OciResponse, error) {
	httpRequest, err := request.GetHttpRequest(http.MethodGet, "/bootVolumes/{bootVolumeId}")
	if err != nil {
		return nil, err
	}

	var response GetBootVolumeResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, nil
}

// GetVolume Gets information for the specified volume.
func (client BlockstorageClient) GetVolume(ctx context.Context, request GetVolumeRequest) (response GetVolumeResponse, err error) {
	ociResponse, e := client.Retry(ctx, request, client.getVolume, request.GetRetryPolicy())
	if e != nil {
		err = e
		return
	}
	if convertedResponse, convertedResponseOk := ociResponse.(GetVolumeResponse); convertedResponseOk {
		response = convertedResponse
	}
	return
}

// lower camel case request => not exported
func (client BlockstorageClient) getVolume(ctx context.Context, request common.OciRequest) (common.OciResponse, error) {
	httpRequest, err := request.GetHttpRequest(http.MethodGet, "/volumes/{volumeId}")
	if err != nil {
		return nil, err
	}

	var response GetVolumeResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, nil
}

// GetVolumeBackup Gets information for the specified volume backup.
func (client BlockstorageClient) GetVolumeBackup(ctx context.Context, request GetVolumeBackupRequest) (response GetVolumeBackupResponse, err error) {
	ociResponse, e := client.Retry(ctx, request, client.getVolumeBackup, request.GetRetryPolicy())
	if e != nil {
		err = e
		return
	}
	if convertedResponse, convertedResponseOk := ociResponse.(GetVolumeBackupResponse); convertedResponseOk {
		response = convertedResponse
	}
	return
}

// lower camel case request => not exported
func (client BlockstorageClient) getVolumeBackup(ctx context.Context, request common.OciRequest) (common.OciResponse, error) {
	httpRequest, err := request.GetHttpRequest(http.MethodGet, "/volumeBackups/{volumeBackupId}")
	if err != nil {
		return nil, err
	}

	var response GetVolumeBackupResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, nil
}

// ListBootVolumes Lists the boot volumes in the specified compartment and Availability Domain.
func (client BlockstorageClient) ListBootVolumes(ctx context.Context, request ListBootVolumesRequest) (response ListBootVolumesResponse, err error) {
	ociResponse, e := client.Retry(ctx, request, client.listBootVolumes, request.GetRetryPolicy())
	if e != nil {
		err = e
		return
	}
	if convertedResponse, convertedResponseOk := ociResponse.(ListBootVolumesResponse); convertedResponseOk {
		response = convertedResponse
	}
	return
}

// lower camel case request => not exported
func (client BlockstorageClient) listBootVolumes(ctx context.Context, request common.OciRequest) (common.OciResponse, error) {
	httpRequest, err := request.GetHttpRequest(http.MethodGet, "/bootVolumes")
	if err != nil {
		return nil, err
	}

	var response ListBootVolumesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, nil
}

// ListVolumeBackups Lists the volume backups in the specified compartment. You can filter the results by volume.
func (client BlockstorageClient) ListVolumeBackups(ctx context.Context, request ListVolumeBackupsRequest) (response ListVolumeBackupsResponse, err error) {
	ociResponse, e := client.Retry(ctx, request, client.listVolumeBackups, request.GetRetryPolicy())
	if e != nil {
		err = e
		return
	}
	if convertedResponse, convertedResponseOk := ociResponse.(ListVolumeBackupsResponse); convertedResponseOk {
		response = convertedResponse
	}
	return
}

// lower camel case request => not exported
func (client BlockstorageClient) listVolumeBackups(ctx context.Context, request common.OciRequest) (common.OciResponse, error) {
	httpRequest, err := request.GetHttpRequest(http.MethodGet, "/volumeBackups")
	if err != nil {
		return nil, err
	}

	var response ListVolumeBackupsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, nil
}

// ListVolumes Lists the volumes in the specified compartment and Availability Domain.
func (client BlockstorageClient) ListVolumes(ctx context.Context, request ListVolumesRequest) (response ListVolumesResponse, err error) {
	ociResponse, e := client.Retry(ctx, request, client.listVolumes, request.GetRetryPolicy())
	if e != nil {
		err = e
		return
	}
	if convertedResponse, convertedResponseOk := ociResponse.(ListVolumesResponse); convertedResponseOk {
		response = convertedResponse
	}
	return
}

// lower camel case request => not exported
func (client BlockstorageClient) listVolumes(ctx context.Context, request common.OciRequest) (common.OciResponse, error) {
	httpRequest, err := request.GetHttpRequest(http.MethodGet, "/volumes")
	if err != nil {
		return nil, err
	}

	var response ListVolumesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, nil
}

// UpdateBootVolume Updates the specified boot volume's display name.
func (client BlockstorageClient) UpdateBootVolume(ctx context.Context, request UpdateBootVolumeRequest) (response UpdateBootVolumeResponse, err error) {
	ociResponse, e := client.Retry(ctx, request, client.updateBootVolume, request.GetRetryPolicy())
	if e != nil {
		err = e
		return
	}
	if convertedResponse, convertedResponseOk := ociResponse.(UpdateBootVolumeResponse); convertedResponseOk {
		response = convertedResponse
	}
	return
}

// lower camel case request => not exported
func (client BlockstorageClient) updateBootVolume(ctx context.Context, request common.OciRequest) (common.OciResponse, error) {
	httpRequest, err := request.GetHttpRequest(http.MethodPut, "/bootVolumes/{bootVolumeId}")
	if err != nil {
		return nil, err
	}

	var response UpdateBootVolumeResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, nil
}

// UpdateVolume Updates the specified volume's display name.
// Avoid entering confidential information.
func (client BlockstorageClient) UpdateVolume(ctx context.Context, request UpdateVolumeRequest) (response UpdateVolumeResponse, err error) {
	ociResponse, e := client.Retry(ctx, request, client.updateVolume, request.GetRetryPolicy())
	if e != nil {
		err = e
		return
	}
	if convertedResponse, convertedResponseOk := ociResponse.(UpdateVolumeResponse); convertedResponseOk {
		response = convertedResponse
	}
	return
}

// lower camel case request => not exported
func (client BlockstorageClient) updateVolume(ctx context.Context, request common.OciRequest) (common.OciResponse, error) {
	httpRequest, err := request.GetHttpRequest(http.MethodPut, "/volumes/{volumeId}")
	if err != nil {
		return nil, err
	}

	var response UpdateVolumeResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, nil
}

// UpdateVolumeBackup Updates the display name for the specified volume backup.
// Avoid entering confidential information.
func (client BlockstorageClient) UpdateVolumeBackup(ctx context.Context, request UpdateVolumeBackupRequest) (response UpdateVolumeBackupResponse, err error) {
	ociResponse, e := client.Retry(ctx, request, client.updateVolumeBackup, request.GetRetryPolicy())
	if e != nil {
		err = e
		return
	}
	if convertedResponse, convertedResponseOk := ociResponse.(UpdateVolumeBackupResponse); convertedResponseOk {
		response = convertedResponse
	}
	return
}

// lower camel case request => not exported
func (client BlockstorageClient) updateVolumeBackup(ctx context.Context, request common.OciRequest) (common.OciResponse, error) {
	httpRequest, err := request.GetHttpRequest(http.MethodPut, "/volumeBackups/{volumeBackupId}")
	if err != nil {
		return nil, err
	}

	var response UpdateVolumeBackupResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, nil
}
