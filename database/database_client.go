// Copyright (c) 2016, 2017, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Database Service API
//
// The API for the Database Service.
//

package database

import (
	"context"
	"fmt"
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

//DatabaseClient a client for Database
type DatabaseClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewDatabaseClientWithConfigurationProvider Creates a new default Database client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewDatabaseClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client DatabaseClient, err error) {
	baseClient, err := common.NewClientWithConfig(configProvider)
	if err != nil {
		return
	}

	client = DatabaseClient{BaseClient: baseClient}
	client.BasePath = "20160918"
	err = client.setConfigurationProvider(configProvider)
	return
}

// SetRegion overrides the region of this client.
func (client *DatabaseClient) SetRegion(region string) {
	client.Host = fmt.Sprintf(common.DefaultHostURLTemplate, "database", region)
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *DatabaseClient) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
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
func (client *DatabaseClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// CreateBackup Creates a new backup in the specified database based on the request parameters you provide. If you previously used RMAN or dbcli to configure backups and then you switch to using the Console or the API for backups, a new backup configuration is created and associated with your database. This means that you can no longer rely on your previously configured unmanaged backups to work.
func (client DatabaseClient) CreateBackup(ctx context.Context, request CreateBackupRequest) (response CreateBackupResponse, err error) {
	ociResponse, e := client.Retry(ctx, request, client.createBackup, request.GetRetryPolicy())
	if e != nil {
		err = e
		return
	}
	if convertedResponse, convertedResponseOk := ociResponse.(CreateBackupResponse); convertedResponseOk {
		response = convertedResponse
	}
	return
}

// lower camel case request => not exported
func (client DatabaseClient) createBackup(ctx context.Context, request common.OciRequest) (common.OciResponse, error) {
	httpRequest, err := request.GetHttpRequest(http.MethodPost, "/backups")
	if err != nil {
		return nil, err
	}

	var response CreateBackupResponse
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

// CreateDataGuardAssociation Creates a new Data Guard association.  A Data Guard association represents the replication relationship between the
// specified database and a peer database. For more information, see [Using Oracle Data Guard]({{DOC_SERVER_URL}}/Content/Database/Tasks/usingdataguard.htm).
// All Oracle Cloud Infrastructure resources, including Data Guard associations, get an Oracle-assigned, unique ID
// called an Oracle Cloud Identifier (OCID). When you create a resource, you can find its OCID in the response.
// You can also retrieve a resource's OCID by using a List API operation on that resource type, or by viewing the
// resource in the Console. Fore more information, see
// [Resource Identifiers](http://localhost:8000/Content/General/Concepts/identifiers.htm).
func (client DatabaseClient) CreateDataGuardAssociation(ctx context.Context, request CreateDataGuardAssociationRequest) (response CreateDataGuardAssociationResponse, err error) {
	ociResponse, e := client.Retry(ctx, request, client.createDataGuardAssociation, request.GetRetryPolicy())
	if e != nil {
		err = e
		return
	}
	if convertedResponse, convertedResponseOk := ociResponse.(CreateDataGuardAssociationResponse); convertedResponseOk {
		response = convertedResponse
	}
	return
}

// lower camel case request => not exported
func (client DatabaseClient) createDataGuardAssociation(ctx context.Context, request common.OciRequest) (common.OciResponse, error) {
	httpRequest, err := request.GetHttpRequest(http.MethodPost, "/databases/{databaseId}/dataGuardAssociations")
	if err != nil {
		return nil, err
	}

	var response CreateDataGuardAssociationResponse
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

// CreateDbHome Creates a new DB Home in the specified DB System based on the request parameters you provide.
func (client DatabaseClient) CreateDbHome(ctx context.Context, request CreateDbHomeRequest) (response CreateDbHomeResponse, err error) {
	ociResponse, e := client.Retry(ctx, request, client.createDbHome, request.GetRetryPolicy())
	if e != nil {
		err = e
		return
	}
	if convertedResponse, convertedResponseOk := ociResponse.(CreateDbHomeResponse); convertedResponseOk {
		response = convertedResponse
	}
	return
}

// lower camel case request => not exported
func (client DatabaseClient) createDbHome(ctx context.Context, request common.OciRequest) (common.OciResponse, error) {
	httpRequest, err := request.GetHttpRequest(http.MethodPost, "/dbHomes")
	if err != nil {
		return nil, err
	}

	var response CreateDbHomeResponse
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

// DbNodeAction Performs an action, such as one of the power actions (start, stop, softreset, or reset), on the specified DB Node.
// **start** - power on
// **stop** - power off
// **softreset** - ACPI shutdown and power on
// **reset** - power off and power on
// Note that the **stop** state has no effect on the resources you consume.
// Billing continues for DB Nodes that you stop, and related resources continue
// to apply against any relevant quotas. You must terminate the DB System
// (TerminateDbSystem)
// to remove its resources from billing and quotas.
func (client DatabaseClient) DbNodeAction(ctx context.Context, request DbNodeActionRequest) (response DbNodeActionResponse, err error) {
	ociResponse, e := client.Retry(ctx, request, client.dbNodeAction, request.GetRetryPolicy())
	if e != nil {
		err = e
		return
	}
	if convertedResponse, convertedResponseOk := ociResponse.(DbNodeActionResponse); convertedResponseOk {
		response = convertedResponse
	}
	return
}

// lower camel case request => not exported
func (client DatabaseClient) dbNodeAction(ctx context.Context, request common.OciRequest) (common.OciResponse, error) {
	httpRequest, err := request.GetHttpRequest(http.MethodPost, "/dbNodes/{dbNodeId}")
	if err != nil {
		return nil, err
	}

	var response DbNodeActionResponse
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

// DeleteBackup Deletes a full backup. You cannot delete automatic backups using this API.
func (client DatabaseClient) DeleteBackup(ctx context.Context, request DeleteBackupRequest) (response DeleteBackupResponse, err error) {
	ociResponse, e := client.Retry(ctx, request, client.deleteBackup, request.GetRetryPolicy())
	if e != nil {
		err = e
		return
	}
	if convertedResponse, convertedResponseOk := ociResponse.(DeleteBackupResponse); convertedResponseOk {
		response = convertedResponse
	}
	return
}

// lower camel case request => not exported
func (client DatabaseClient) deleteBackup(ctx context.Context, request common.OciRequest) (common.OciResponse, error) {
	httpRequest, err := request.GetHttpRequest(http.MethodDelete, "/backups/{backupId}")
	if err != nil {
		return nil, err
	}

	var response DeleteBackupResponse
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

// DeleteDbHome Deletes a DB Home. The DB Home and its database data are local to the DB System and will be lost when it is deleted. Oracle recommends that you back up any data in the DB System prior to deleting it.
func (client DatabaseClient) DeleteDbHome(ctx context.Context, request DeleteDbHomeRequest) (response DeleteDbHomeResponse, err error) {
	ociResponse, e := client.Retry(ctx, request, client.deleteDbHome, request.GetRetryPolicy())
	if e != nil {
		err = e
		return
	}
	if convertedResponse, convertedResponseOk := ociResponse.(DeleteDbHomeResponse); convertedResponseOk {
		response = convertedResponse
	}
	return
}

// lower camel case request => not exported
func (client DatabaseClient) deleteDbHome(ctx context.Context, request common.OciRequest) (common.OciResponse, error) {
	httpRequest, err := request.GetHttpRequest(http.MethodDelete, "/dbHomes/{dbHomeId}")
	if err != nil {
		return nil, err
	}

	var response DeleteDbHomeResponse
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

// FailoverDataGuardAssociation Performs a failover to transition the standby database identified by the `databaseId` parameter into the
// specified Data Guard association's primary role after the existing primary database fails or becomes unreachable.
// A failover might result in data loss depending on the protection mode in effect at the time of the primary
// database failure.
func (client DatabaseClient) FailoverDataGuardAssociation(ctx context.Context, request FailoverDataGuardAssociationRequest) (response FailoverDataGuardAssociationResponse, err error) {
	ociResponse, e := client.Retry(ctx, request, client.failoverDataGuardAssociation, request.GetRetryPolicy())
	if e != nil {
		err = e
		return
	}
	if convertedResponse, convertedResponseOk := ociResponse.(FailoverDataGuardAssociationResponse); convertedResponseOk {
		response = convertedResponse
	}
	return
}

// lower camel case request => not exported
func (client DatabaseClient) failoverDataGuardAssociation(ctx context.Context, request common.OciRequest) (common.OciResponse, error) {
	httpRequest, err := request.GetHttpRequest(http.MethodPost, "/databases/{databaseId}/dataGuardAssociations/{dataGuardAssociationId}/actions/failover")
	if err != nil {
		return nil, err
	}

	var response FailoverDataGuardAssociationResponse
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

// GetBackup Gets information about the specified backup.
func (client DatabaseClient) GetBackup(ctx context.Context, request GetBackupRequest) (response GetBackupResponse, err error) {
	ociResponse, e := client.Retry(ctx, request, client.getBackup, request.GetRetryPolicy())
	if e != nil {
		err = e
		return
	}
	if convertedResponse, convertedResponseOk := ociResponse.(GetBackupResponse); convertedResponseOk {
		response = convertedResponse
	}
	return
}

// lower camel case request => not exported
func (client DatabaseClient) getBackup(ctx context.Context, request common.OciRequest) (common.OciResponse, error) {
	httpRequest, err := request.GetHttpRequest(http.MethodGet, "/backups/{backupId}")
	if err != nil {
		return nil, err
	}

	var response GetBackupResponse
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

// GetDataGuardAssociation Gets the specified Data Guard association's configuration information.
func (client DatabaseClient) GetDataGuardAssociation(ctx context.Context, request GetDataGuardAssociationRequest) (response GetDataGuardAssociationResponse, err error) {
	ociResponse, e := client.Retry(ctx, request, client.getDataGuardAssociation, request.GetRetryPolicy())
	if e != nil {
		err = e
		return
	}
	if convertedResponse, convertedResponseOk := ociResponse.(GetDataGuardAssociationResponse); convertedResponseOk {
		response = convertedResponse
	}
	return
}

// lower camel case request => not exported
func (client DatabaseClient) getDataGuardAssociation(ctx context.Context, request common.OciRequest) (common.OciResponse, error) {
	httpRequest, err := request.GetHttpRequest(http.MethodGet, "/databases/{databaseId}/dataGuardAssociations/{dataGuardAssociationId}")
	if err != nil {
		return nil, err
	}

	var response GetDataGuardAssociationResponse
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

// GetDatabase Gets information about a specific database.
func (client DatabaseClient) GetDatabase(ctx context.Context, request GetDatabaseRequest) (response GetDatabaseResponse, err error) {
	ociResponse, e := client.Retry(ctx, request, client.getDatabase, request.GetRetryPolicy())
	if e != nil {
		err = e
		return
	}
	if convertedResponse, convertedResponseOk := ociResponse.(GetDatabaseResponse); convertedResponseOk {
		response = convertedResponse
	}
	return
}

// lower camel case request => not exported
func (client DatabaseClient) getDatabase(ctx context.Context, request common.OciRequest) (common.OciResponse, error) {
	httpRequest, err := request.GetHttpRequest(http.MethodGet, "/databases/{databaseId}")
	if err != nil {
		return nil, err
	}

	var response GetDatabaseResponse
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

// GetDbHome Gets information about the specified database home.
func (client DatabaseClient) GetDbHome(ctx context.Context, request GetDbHomeRequest) (response GetDbHomeResponse, err error) {
	ociResponse, e := client.Retry(ctx, request, client.getDbHome, request.GetRetryPolicy())
	if e != nil {
		err = e
		return
	}
	if convertedResponse, convertedResponseOk := ociResponse.(GetDbHomeResponse); convertedResponseOk {
		response = convertedResponse
	}
	return
}

// lower camel case request => not exported
func (client DatabaseClient) getDbHome(ctx context.Context, request common.OciRequest) (common.OciResponse, error) {
	httpRequest, err := request.GetHttpRequest(http.MethodGet, "/dbHomes/{dbHomeId}")
	if err != nil {
		return nil, err
	}

	var response GetDbHomeResponse
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

// GetDbHomePatch Gets information about a specified patch package.
func (client DatabaseClient) GetDbHomePatch(ctx context.Context, request GetDbHomePatchRequest) (response GetDbHomePatchResponse, err error) {
	ociResponse, e := client.Retry(ctx, request, client.getDbHomePatch, request.GetRetryPolicy())
	if e != nil {
		err = e
		return
	}
	if convertedResponse, convertedResponseOk := ociResponse.(GetDbHomePatchResponse); convertedResponseOk {
		response = convertedResponse
	}
	return
}

// lower camel case request => not exported
func (client DatabaseClient) getDbHomePatch(ctx context.Context, request common.OciRequest) (common.OciResponse, error) {
	httpRequest, err := request.GetHttpRequest(http.MethodGet, "/dbHomes/{dbHomeId}/patches/{patchId}")
	if err != nil {
		return nil, err
	}

	var response GetDbHomePatchResponse
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

// GetDbHomePatchHistoryEntry Gets the patch history details for the specified patchHistoryEntryId
func (client DatabaseClient) GetDbHomePatchHistoryEntry(ctx context.Context, request GetDbHomePatchHistoryEntryRequest) (response GetDbHomePatchHistoryEntryResponse, err error) {
	ociResponse, e := client.Retry(ctx, request, client.getDbHomePatchHistoryEntry, request.GetRetryPolicy())
	if e != nil {
		err = e
		return
	}
	if convertedResponse, convertedResponseOk := ociResponse.(GetDbHomePatchHistoryEntryResponse); convertedResponseOk {
		response = convertedResponse
	}
	return
}

// lower camel case request => not exported
func (client DatabaseClient) getDbHomePatchHistoryEntry(ctx context.Context, request common.OciRequest) (common.OciResponse, error) {
	httpRequest, err := request.GetHttpRequest(http.MethodGet, "/dbHomes/{dbHomeId}/patchHistoryEntries/{patchHistoryEntryId}")
	if err != nil {
		return nil, err
	}

	var response GetDbHomePatchHistoryEntryResponse
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

// GetDbNode Gets information about the specified database node.
func (client DatabaseClient) GetDbNode(ctx context.Context, request GetDbNodeRequest) (response GetDbNodeResponse, err error) {
	ociResponse, e := client.Retry(ctx, request, client.getDbNode, request.GetRetryPolicy())
	if e != nil {
		err = e
		return
	}
	if convertedResponse, convertedResponseOk := ociResponse.(GetDbNodeResponse); convertedResponseOk {
		response = convertedResponse
	}
	return
}

// lower camel case request => not exported
func (client DatabaseClient) getDbNode(ctx context.Context, request common.OciRequest) (common.OciResponse, error) {
	httpRequest, err := request.GetHttpRequest(http.MethodGet, "/dbNodes/{dbNodeId}")
	if err != nil {
		return nil, err
	}

	var response GetDbNodeResponse
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

// GetDbSystem Gets information about the specified DB System.
func (client DatabaseClient) GetDbSystem(ctx context.Context, request GetDbSystemRequest) (response GetDbSystemResponse, err error) {
	ociResponse, e := client.Retry(ctx, request, client.getDbSystem, request.GetRetryPolicy())
	if e != nil {
		err = e
		return
	}
	if convertedResponse, convertedResponseOk := ociResponse.(GetDbSystemResponse); convertedResponseOk {
		response = convertedResponse
	}
	return
}

// lower camel case request => not exported
func (client DatabaseClient) getDbSystem(ctx context.Context, request common.OciRequest) (common.OciResponse, error) {
	httpRequest, err := request.GetHttpRequest(http.MethodGet, "/dbSystems/{dbSystemId}")
	if err != nil {
		return nil, err
	}

	var response GetDbSystemResponse
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

// GetDbSystemPatch Gets information about a specified patch package.
func (client DatabaseClient) GetDbSystemPatch(ctx context.Context, request GetDbSystemPatchRequest) (response GetDbSystemPatchResponse, err error) {
	ociResponse, e := client.Retry(ctx, request, client.getDbSystemPatch, request.GetRetryPolicy())
	if e != nil {
		err = e
		return
	}
	if convertedResponse, convertedResponseOk := ociResponse.(GetDbSystemPatchResponse); convertedResponseOk {
		response = convertedResponse
	}
	return
}

// lower camel case request => not exported
func (client DatabaseClient) getDbSystemPatch(ctx context.Context, request common.OciRequest) (common.OciResponse, error) {
	httpRequest, err := request.GetHttpRequest(http.MethodGet, "/dbSystems/{dbSystemId}/patches/{patchId}")
	if err != nil {
		return nil, err
	}

	var response GetDbSystemPatchResponse
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

// GetDbSystemPatchHistoryEntry Gets the patch history details for the specified patchHistoryEntryId.
func (client DatabaseClient) GetDbSystemPatchHistoryEntry(ctx context.Context, request GetDbSystemPatchHistoryEntryRequest) (response GetDbSystemPatchHistoryEntryResponse, err error) {
	ociResponse, e := client.Retry(ctx, request, client.getDbSystemPatchHistoryEntry, request.GetRetryPolicy())
	if e != nil {
		err = e
		return
	}
	if convertedResponse, convertedResponseOk := ociResponse.(GetDbSystemPatchHistoryEntryResponse); convertedResponseOk {
		response = convertedResponse
	}
	return
}

// lower camel case request => not exported
func (client DatabaseClient) getDbSystemPatchHistoryEntry(ctx context.Context, request common.OciRequest) (common.OciResponse, error) {
	httpRequest, err := request.GetHttpRequest(http.MethodGet, "/dbSystems/{dbSystemId}/patchHistoryEntries/{patchHistoryEntryId}")
	if err != nil {
		return nil, err
	}

	var response GetDbSystemPatchHistoryEntryResponse
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

// LaunchDbSystem Launches a new DB System in the specified compartment and Availability Domain. You'll specify a single Oracle
// Database Edition that applies to all the databases on that DB System. The selected edition cannot be changed.
// An initial database is created on the DB System based on the request parameters you provide and some default
// options. For more information,
// see [Default Options for the Initial Database]({{DOC_SERVER_URL}}/Content/Database/Tasks/launchingDB.htm#Default_Options_for_the_Initial_Database).
// The DB System will include a command line interface (CLI) that you can use to create additional databases and
// manage existing databases. For more information, see the
// [Oracle Database CLI Reference]({{DOC_SERVER_URL}}/Content/Database/References/odacli.htm#Oracle_Database_CLI_Reference).
func (client DatabaseClient) LaunchDbSystem(ctx context.Context, request LaunchDbSystemRequest) (response LaunchDbSystemResponse, err error) {
	ociResponse, e := client.Retry(ctx, request, client.launchDbSystem, request.GetRetryPolicy())
	if e != nil {
		err = e
		return
	}
	if convertedResponse, convertedResponseOk := ociResponse.(LaunchDbSystemResponse); convertedResponseOk {
		response = convertedResponse
	}
	return
}

// lower camel case request => not exported
func (client DatabaseClient) launchDbSystem(ctx context.Context, request common.OciRequest) (common.OciResponse, error) {
	httpRequest, err := request.GetHttpRequest(http.MethodPost, "/dbSystems")
	if err != nil {
		return nil, err
	}

	var response LaunchDbSystemResponse
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

// ListBackups Gets a list of backups based on the databaseId or compartmentId specified. Either one of the query parameters must be provided.
func (client DatabaseClient) ListBackups(ctx context.Context, request ListBackupsRequest) (response ListBackupsResponse, err error) {
	ociResponse, e := client.Retry(ctx, request, client.listBackups, request.GetRetryPolicy())
	if e != nil {
		err = e
		return
	}
	if convertedResponse, convertedResponseOk := ociResponse.(ListBackupsResponse); convertedResponseOk {
		response = convertedResponse
	}
	return
}

// lower camel case request => not exported
func (client DatabaseClient) listBackups(ctx context.Context, request common.OciRequest) (common.OciResponse, error) {
	httpRequest, err := request.GetHttpRequest(http.MethodGet, "/backups")
	if err != nil {
		return nil, err
	}

	var response ListBackupsResponse
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

// ListDataGuardAssociations Lists all Data Guard associations for the specified database.
func (client DatabaseClient) ListDataGuardAssociations(ctx context.Context, request ListDataGuardAssociationsRequest) (response ListDataGuardAssociationsResponse, err error) {
	ociResponse, e := client.Retry(ctx, request, client.listDataGuardAssociations, request.GetRetryPolicy())
	if e != nil {
		err = e
		return
	}
	if convertedResponse, convertedResponseOk := ociResponse.(ListDataGuardAssociationsResponse); convertedResponseOk {
		response = convertedResponse
	}
	return
}

// lower camel case request => not exported
func (client DatabaseClient) listDataGuardAssociations(ctx context.Context, request common.OciRequest) (common.OciResponse, error) {
	httpRequest, err := request.GetHttpRequest(http.MethodGet, "/databases/{databaseId}/dataGuardAssociations")
	if err != nil {
		return nil, err
	}

	var response ListDataGuardAssociationsResponse
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

// ListDatabases Gets a list of the databases in the specified database home.
func (client DatabaseClient) ListDatabases(ctx context.Context, request ListDatabasesRequest) (response ListDatabasesResponse, err error) {
	ociResponse, e := client.Retry(ctx, request, client.listDatabases, request.GetRetryPolicy())
	if e != nil {
		err = e
		return
	}
	if convertedResponse, convertedResponseOk := ociResponse.(ListDatabasesResponse); convertedResponseOk {
		response = convertedResponse
	}
	return
}

// lower camel case request => not exported
func (client DatabaseClient) listDatabases(ctx context.Context, request common.OciRequest) (common.OciResponse, error) {
	httpRequest, err := request.GetHttpRequest(http.MethodGet, "/databases")
	if err != nil {
		return nil, err
	}

	var response ListDatabasesResponse
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

// ListDbHomePatchHistoryEntries Gets history of the actions taken for patches for the specified database home.
func (client DatabaseClient) ListDbHomePatchHistoryEntries(ctx context.Context, request ListDbHomePatchHistoryEntriesRequest) (response ListDbHomePatchHistoryEntriesResponse, err error) {
	ociResponse, e := client.Retry(ctx, request, client.listDbHomePatchHistoryEntries, request.GetRetryPolicy())
	if e != nil {
		err = e
		return
	}
	if convertedResponse, convertedResponseOk := ociResponse.(ListDbHomePatchHistoryEntriesResponse); convertedResponseOk {
		response = convertedResponse
	}
	return
}

// lower camel case request => not exported
func (client DatabaseClient) listDbHomePatchHistoryEntries(ctx context.Context, request common.OciRequest) (common.OciResponse, error) {
	httpRequest, err := request.GetHttpRequest(http.MethodGet, "/dbHomes/{dbHomeId}/patchHistoryEntries")
	if err != nil {
		return nil, err
	}

	var response ListDbHomePatchHistoryEntriesResponse
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

// ListDbHomePatches Lists patches applicable to the requested database home.
func (client DatabaseClient) ListDbHomePatches(ctx context.Context, request ListDbHomePatchesRequest) (response ListDbHomePatchesResponse, err error) {
	ociResponse, e := client.Retry(ctx, request, client.listDbHomePatches, request.GetRetryPolicy())
	if e != nil {
		err = e
		return
	}
	if convertedResponse, convertedResponseOk := ociResponse.(ListDbHomePatchesResponse); convertedResponseOk {
		response = convertedResponse
	}
	return
}

// lower camel case request => not exported
func (client DatabaseClient) listDbHomePatches(ctx context.Context, request common.OciRequest) (common.OciResponse, error) {
	httpRequest, err := request.GetHttpRequest(http.MethodGet, "/dbHomes/{dbHomeId}/patches")
	if err != nil {
		return nil, err
	}

	var response ListDbHomePatchesResponse
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

// ListDbHomes Gets a list of database homes in the specified DB System and compartment. A database home is a directory where Oracle database software is installed.
func (client DatabaseClient) ListDbHomes(ctx context.Context, request ListDbHomesRequest) (response ListDbHomesResponse, err error) {
	ociResponse, e := client.Retry(ctx, request, client.listDbHomes, request.GetRetryPolicy())
	if e != nil {
		err = e
		return
	}
	if convertedResponse, convertedResponseOk := ociResponse.(ListDbHomesResponse); convertedResponseOk {
		response = convertedResponse
	}
	return
}

// lower camel case request => not exported
func (client DatabaseClient) listDbHomes(ctx context.Context, request common.OciRequest) (common.OciResponse, error) {
	httpRequest, err := request.GetHttpRequest(http.MethodGet, "/dbHomes")
	if err != nil {
		return nil, err
	}

	var response ListDbHomesResponse
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

// ListDbNodes Gets a list of database nodes in the specified DB System and compartment. A database node is a server running database software.
func (client DatabaseClient) ListDbNodes(ctx context.Context, request ListDbNodesRequest) (response ListDbNodesResponse, err error) {
	ociResponse, e := client.Retry(ctx, request, client.listDbNodes, request.GetRetryPolicy())
	if e != nil {
		err = e
		return
	}
	if convertedResponse, convertedResponseOk := ociResponse.(ListDbNodesResponse); convertedResponseOk {
		response = convertedResponse
	}
	return
}

// lower camel case request => not exported
func (client DatabaseClient) listDbNodes(ctx context.Context, request common.OciRequest) (common.OciResponse, error) {
	httpRequest, err := request.GetHttpRequest(http.MethodGet, "/dbNodes")
	if err != nil {
		return nil, err
	}

	var response ListDbNodesResponse
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

// ListDbSystemPatchHistoryEntries Gets the history of the patch actions performed on the specified DB System.
func (client DatabaseClient) ListDbSystemPatchHistoryEntries(ctx context.Context, request ListDbSystemPatchHistoryEntriesRequest) (response ListDbSystemPatchHistoryEntriesResponse, err error) {
	ociResponse, e := client.Retry(ctx, request, client.listDbSystemPatchHistoryEntries, request.GetRetryPolicy())
	if e != nil {
		err = e
		return
	}
	if convertedResponse, convertedResponseOk := ociResponse.(ListDbSystemPatchHistoryEntriesResponse); convertedResponseOk {
		response = convertedResponse
	}
	return
}

// lower camel case request => not exported
func (client DatabaseClient) listDbSystemPatchHistoryEntries(ctx context.Context, request common.OciRequest) (common.OciResponse, error) {
	httpRequest, err := request.GetHttpRequest(http.MethodGet, "/dbSystems/{dbSystemId}/patchHistoryEntries")
	if err != nil {
		return nil, err
	}

	var response ListDbSystemPatchHistoryEntriesResponse
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

// ListDbSystemPatches Lists the patches applicable to the requested DB System.
func (client DatabaseClient) ListDbSystemPatches(ctx context.Context, request ListDbSystemPatchesRequest) (response ListDbSystemPatchesResponse, err error) {
	ociResponse, e := client.Retry(ctx, request, client.listDbSystemPatches, request.GetRetryPolicy())
	if e != nil {
		err = e
		return
	}
	if convertedResponse, convertedResponseOk := ociResponse.(ListDbSystemPatchesResponse); convertedResponseOk {
		response = convertedResponse
	}
	return
}

// lower camel case request => not exported
func (client DatabaseClient) listDbSystemPatches(ctx context.Context, request common.OciRequest) (common.OciResponse, error) {
	httpRequest, err := request.GetHttpRequest(http.MethodGet, "/dbSystems/{dbSystemId}/patches")
	if err != nil {
		return nil, err
	}

	var response ListDbSystemPatchesResponse
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

// ListDbSystemShapes Gets a list of the shapes that can be used to launch a new DB System. The shape determines resources to allocate to the DB system - CPU cores and memory for VM shapes; CPU cores, memory and storage for non-VM (or bare metal) shapes.
func (client DatabaseClient) ListDbSystemShapes(ctx context.Context, request ListDbSystemShapesRequest) (response ListDbSystemShapesResponse, err error) {
	ociResponse, e := client.Retry(ctx, request, client.listDbSystemShapes, request.GetRetryPolicy())
	if e != nil {
		err = e
		return
	}
	if convertedResponse, convertedResponseOk := ociResponse.(ListDbSystemShapesResponse); convertedResponseOk {
		response = convertedResponse
	}
	return
}

// lower camel case request => not exported
func (client DatabaseClient) listDbSystemShapes(ctx context.Context, request common.OciRequest) (common.OciResponse, error) {
	httpRequest, err := request.GetHttpRequest(http.MethodGet, "/dbSystemShapes")
	if err != nil {
		return nil, err
	}

	var response ListDbSystemShapesResponse
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

// ListDbSystems Gets a list of the DB Systems in the specified compartment.
//
func (client DatabaseClient) ListDbSystems(ctx context.Context, request ListDbSystemsRequest) (response ListDbSystemsResponse, err error) {
	ociResponse, e := client.Retry(ctx, request, client.listDbSystems, request.GetRetryPolicy())
	if e != nil {
		err = e
		return
	}
	if convertedResponse, convertedResponseOk := ociResponse.(ListDbSystemsResponse); convertedResponseOk {
		response = convertedResponse
	}
	return
}

// lower camel case request => not exported
func (client DatabaseClient) listDbSystems(ctx context.Context, request common.OciRequest) (common.OciResponse, error) {
	httpRequest, err := request.GetHttpRequest(http.MethodGet, "/dbSystems")
	if err != nil {
		return nil, err
	}

	var response ListDbSystemsResponse
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

// ListDbVersions Gets a list of supported Oracle database versions.
func (client DatabaseClient) ListDbVersions(ctx context.Context, request ListDbVersionsRequest) (response ListDbVersionsResponse, err error) {
	ociResponse, e := client.Retry(ctx, request, client.listDbVersions, request.GetRetryPolicy())
	if e != nil {
		err = e
		return
	}
	if convertedResponse, convertedResponseOk := ociResponse.(ListDbVersionsResponse); convertedResponseOk {
		response = convertedResponse
	}
	return
}

// lower camel case request => not exported
func (client DatabaseClient) listDbVersions(ctx context.Context, request common.OciRequest) (common.OciResponse, error) {
	httpRequest, err := request.GetHttpRequest(http.MethodGet, "/dbVersions")
	if err != nil {
		return nil, err
	}

	var response ListDbVersionsResponse
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

// ReinstateDataGuardAssociation Reinstates the database identified by the `databaseId` parameter into the standby role in a Data Guard association.
func (client DatabaseClient) ReinstateDataGuardAssociation(ctx context.Context, request ReinstateDataGuardAssociationRequest) (response ReinstateDataGuardAssociationResponse, err error) {
	ociResponse, e := client.Retry(ctx, request, client.reinstateDataGuardAssociation, request.GetRetryPolicy())
	if e != nil {
		err = e
		return
	}
	if convertedResponse, convertedResponseOk := ociResponse.(ReinstateDataGuardAssociationResponse); convertedResponseOk {
		response = convertedResponse
	}
	return
}

// lower camel case request => not exported
func (client DatabaseClient) reinstateDataGuardAssociation(ctx context.Context, request common.OciRequest) (common.OciResponse, error) {
	httpRequest, err := request.GetHttpRequest(http.MethodPost, "/databases/{databaseId}/dataGuardAssociations/{dataGuardAssociationId}/actions/reinstate")
	if err != nil {
		return nil, err
	}

	var response ReinstateDataGuardAssociationResponse
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

// RestoreDatabase Restore a Database based on the request parameters you provide.
func (client DatabaseClient) RestoreDatabase(ctx context.Context, request RestoreDatabaseRequest) (response RestoreDatabaseResponse, err error) {
	ociResponse, e := client.Retry(ctx, request, client.restoreDatabase, request.GetRetryPolicy())
	if e != nil {
		err = e
		return
	}
	if convertedResponse, convertedResponseOk := ociResponse.(RestoreDatabaseResponse); convertedResponseOk {
		response = convertedResponse
	}
	return
}

// lower camel case request => not exported
func (client DatabaseClient) restoreDatabase(ctx context.Context, request common.OciRequest) (common.OciResponse, error) {
	httpRequest, err := request.GetHttpRequest(http.MethodPost, "/databases/{databaseId}/actions/restore")
	if err != nil {
		return nil, err
	}

	var response RestoreDatabaseResponse
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

// SwitchoverDataGuardAssociation Performs a switchover to transition the primary database of a Data Guard association into a standby role. The
// standby database associated with the `dataGuardAssociationId` assumes the primary database role.
// A switchover guarantees no data loss.
func (client DatabaseClient) SwitchoverDataGuardAssociation(ctx context.Context, request SwitchoverDataGuardAssociationRequest) (response SwitchoverDataGuardAssociationResponse, err error) {
	ociResponse, e := client.Retry(ctx, request, client.switchoverDataGuardAssociation, request.GetRetryPolicy())
	if e != nil {
		err = e
		return
	}
	if convertedResponse, convertedResponseOk := ociResponse.(SwitchoverDataGuardAssociationResponse); convertedResponseOk {
		response = convertedResponse
	}
	return
}

// lower camel case request => not exported
func (client DatabaseClient) switchoverDataGuardAssociation(ctx context.Context, request common.OciRequest) (common.OciResponse, error) {
	httpRequest, err := request.GetHttpRequest(http.MethodPost, "/databases/{databaseId}/dataGuardAssociations/{dataGuardAssociationId}/actions/switchover")
	if err != nil {
		return nil, err
	}

	var response SwitchoverDataGuardAssociationResponse
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

// TerminateDbSystem Terminates a DB System and permanently deletes it and any databases running on it, and any storage volumes attached to it. The database data is local to the DB System and will be lost when the system is terminated. Oracle recommends that you back up any data in the DB System prior to terminating it.
func (client DatabaseClient) TerminateDbSystem(ctx context.Context, request TerminateDbSystemRequest) (response TerminateDbSystemResponse, err error) {
	ociResponse, e := client.Retry(ctx, request, client.terminateDbSystem, request.GetRetryPolicy())
	if e != nil {
		err = e
		return
	}
	if convertedResponse, convertedResponseOk := ociResponse.(TerminateDbSystemResponse); convertedResponseOk {
		response = convertedResponse
	}
	return
}

// lower camel case request => not exported
func (client DatabaseClient) terminateDbSystem(ctx context.Context, request common.OciRequest) (common.OciResponse, error) {
	httpRequest, err := request.GetHttpRequest(http.MethodDelete, "/dbSystems/{dbSystemId}")
	if err != nil {
		return nil, err
	}

	var response TerminateDbSystemResponse
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

// UpdateDatabase Update a Database based on the request parameters you provide.
func (client DatabaseClient) UpdateDatabase(ctx context.Context, request UpdateDatabaseRequest) (response UpdateDatabaseResponse, err error) {
	ociResponse, e := client.Retry(ctx, request, client.updateDatabase, request.GetRetryPolicy())
	if e != nil {
		err = e
		return
	}
	if convertedResponse, convertedResponseOk := ociResponse.(UpdateDatabaseResponse); convertedResponseOk {
		response = convertedResponse
	}
	return
}

// lower camel case request => not exported
func (client DatabaseClient) updateDatabase(ctx context.Context, request common.OciRequest) (common.OciResponse, error) {
	httpRequest, err := request.GetHttpRequest(http.MethodPut, "/databases/{databaseId}")
	if err != nil {
		return nil, err
	}

	var response UpdateDatabaseResponse
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

// UpdateDbHome Patches the specified dbHome.
func (client DatabaseClient) UpdateDbHome(ctx context.Context, request UpdateDbHomeRequest) (response UpdateDbHomeResponse, err error) {
	ociResponse, e := client.Retry(ctx, request, client.updateDbHome, request.GetRetryPolicy())
	if e != nil {
		err = e
		return
	}
	if convertedResponse, convertedResponseOk := ociResponse.(UpdateDbHomeResponse); convertedResponseOk {
		response = convertedResponse
	}
	return
}

// lower camel case request => not exported
func (client DatabaseClient) updateDbHome(ctx context.Context, request common.OciRequest) (common.OciResponse, error) {
	httpRequest, err := request.GetHttpRequest(http.MethodPut, "/dbHomes/{dbHomeId}")
	if err != nil {
		return nil, err
	}

	var response UpdateDbHomeResponse
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

// UpdateDbSystem Updates the properties of a DB System, such as the CPU core count.
func (client DatabaseClient) UpdateDbSystem(ctx context.Context, request UpdateDbSystemRequest) (response UpdateDbSystemResponse, err error) {
	ociResponse, e := client.Retry(ctx, request, client.updateDbSystem, request.GetRetryPolicy())
	if e != nil {
		err = e
		return
	}
	if convertedResponse, convertedResponseOk := ociResponse.(UpdateDbSystemResponse); convertedResponseOk {
		response = convertedResponse
	}
	return
}

// lower camel case request => not exported
func (client DatabaseClient) updateDbSystem(ctx context.Context, request common.OciRequest) (common.OciResponse, error) {
	httpRequest, err := request.GetHttpRequest(http.MethodPut, "/dbSystems/{dbSystemId}")
	if err != nil {
		return nil, err
	}

	var response UpdateDbSystemResponse
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
