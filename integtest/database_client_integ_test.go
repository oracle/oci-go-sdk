// Copyright (c) 2016, 2017, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Database Service API
//
// The API for the Database Service.
//

package integtest

import (
	"context"
	"fmt"
	"testing"

	"github.com/oracle/oci-go-sdk/common"
	"github.com/oracle/oci-go-sdk/database"
	"github.com/stretchr/testify/assert"
)

func TestDatabaseClient_CreateBackup(t *testing.T) {
	t.Skip("Not implemented")
	c, clerr := database.NewDatabaseClientWithConfigurationProvider(common.DefaultConfigProvider())
	failIfError(t, clerr)
	request := database.CreateBackupRequest{}
	r, err := c.CreateBackup(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

func TestDatabaseClient_CreateDataGuardAssociation(t *testing.T) {
	t.Skip("Not implemented")
	c, clerr := database.NewDatabaseClientWithConfigurationProvider(common.DefaultConfigProvider())
	failIfError(t, clerr)
	request := database.CreateDataGuardAssociationRequest{}
	r, err := c.CreateDataGuardAssociation(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

func TestDatabaseClient_CreateDbHome(t *testing.T) {
	t.Skip("Not implemented")
	c, clerr := database.NewDatabaseClientWithConfigurationProvider(common.DefaultConfigProvider())
	failIfError(t, clerr)
	request := database.CreateDbHomeRequest{}
	r, err := c.CreateDbHome(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

func TestDatabaseClient_DbNodeAction(t *testing.T) {
	t.Skip("Not implemented")
	c, clerr := database.NewDatabaseClientWithConfigurationProvider(common.DefaultConfigProvider())
	failIfError(t, clerr)
	request := database.DbNodeActionRequest{}
	r, err := c.DbNodeAction(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

func TestDatabaseClient_DeleteBackup(t *testing.T) {
	t.Skip("Not implemented")
	c, clerr := database.NewDatabaseClientWithConfigurationProvider(common.DefaultConfigProvider())
	failIfError(t, clerr)
	request := database.DeleteBackupRequest{}
	_, err := c.DeleteBackup(context.Background(), request)
	assert.NoError(t, err)
	return
}

func TestDatabaseClient_DeleteDbHome(t *testing.T) {
	t.Skip("Not implemented")
	c, clerr := database.NewDatabaseClientWithConfigurationProvider(common.DefaultConfigProvider())
	failIfError(t, clerr)
	request := database.DeleteDbHomeRequest{}
	_, err := c.DeleteDbHome(context.Background(), request)
	assert.NoError(t, err)
	return
}

func TestDatabaseClient_FailoverDataGuardAssociation(t *testing.T) {
	t.Skip("Not implemented")
	c, clerr := database.NewDatabaseClientWithConfigurationProvider(common.DefaultConfigProvider())
	failIfError(t, clerr)
	request := database.FailoverDataGuardAssociationRequest{}
	r, err := c.FailoverDataGuardAssociation(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

func TestDatabaseClient_GetBackup(t *testing.T) {
	t.Skip("Not implemented")
	c, clerr := database.NewDatabaseClientWithConfigurationProvider(common.DefaultConfigProvider())
	failIfError(t, clerr)
	request := database.GetBackupRequest{}
	r, err := c.GetBackup(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

func TestDatabaseClient_GetDataGuardAssociation(t *testing.T) {
	t.Skip("Not implemented")
	c, clerr := database.NewDatabaseClientWithConfigurationProvider(common.DefaultConfigProvider())
	failIfError(t, clerr)
	request := database.GetDataGuardAssociationRequest{}
	r, err := c.GetDataGuardAssociation(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

func TestDatabaseClient_GetDatabase(t *testing.T) {
	t.Skip("Not implemented")
	c, clerr := database.NewDatabaseClientWithConfigurationProvider(common.DefaultConfigProvider())
	failIfError(t, clerr)
	request := database.GetDatabaseRequest{}
	r, err := c.GetDatabase(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

func TestDatabaseClient_GetDbHome(t *testing.T) {
	t.Skip("Not implemented")
	c, clerr := database.NewDatabaseClientWithConfigurationProvider(common.DefaultConfigProvider())
	failIfError(t, clerr)
	request := database.GetDbHomeRequest{}
	r, err := c.GetDbHome(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

func TestDatabaseClient_GetDbHomePatch(t *testing.T) {
	t.Skip("Not implemented")
	c, clerr := database.NewDatabaseClientWithConfigurationProvider(common.DefaultConfigProvider())
	failIfError(t, clerr)
	request := database.GetDbHomePatchRequest{}
	r, err := c.GetDbHomePatch(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

func TestDatabaseClient_GetDbHomePatchHistoryEntry(t *testing.T) {
	t.Skip("Not implemented")
	c, clerr := database.NewDatabaseClientWithConfigurationProvider(common.DefaultConfigProvider())
	failIfError(t, clerr)
	request := database.GetDbHomePatchHistoryEntryRequest{}
	r, err := c.GetDbHomePatchHistoryEntry(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

func TestDatabaseClient_GetDbNode(t *testing.T) {
	t.Skip("Not implemented")
	c, clerr := database.NewDatabaseClientWithConfigurationProvider(common.DefaultConfigProvider())
	failIfError(t, clerr)
	request := database.GetDbNodeRequest{}
	r, err := c.GetDbNode(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

func TestDatabaseClient_GetDbSystem(t *testing.T) {
	t.Skip("Not implemented")
	c, clerr := database.NewDatabaseClientWithConfigurationProvider(common.DefaultConfigProvider())
	failIfError(t, clerr)
	request := database.GetDbSystemRequest{}
	r, err := c.GetDbSystem(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

func TestDatabaseClient_GetDbSystemPatch(t *testing.T) {
	t.Skip("Not implemented")
	c, clerr := database.NewDatabaseClientWithConfigurationProvider(common.DefaultConfigProvider())
	failIfError(t, clerr)
	request := database.GetDbSystemPatchRequest{}
	r, err := c.GetDbSystemPatch(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

func TestDatabaseClient_GetDbSystemPatchHistoryEntry(t *testing.T) {
	t.Skip("Not implemented")
	c, clerr := database.NewDatabaseClientWithConfigurationProvider(common.DefaultConfigProvider())
	failIfError(t, clerr)
	request := database.GetDbSystemPatchHistoryEntryRequest{}
	r, err := c.GetDbSystemPatchHistoryEntry(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

func TestDatabaseClient_LaunchDbSystem(t *testing.T) {

	if !getRunExpensiveTests() {
		t.Skip("expensive test excluded")
		return
	}

	c, clerr := database.NewDatabaseClientWithConfigurationProvider(common.DefaultConfigProvider())
	failIfError(t, clerr)

	request := database.LaunchDbSystemRequest{}
	request.AvailabilityDomain = common.String(validAD())
	request.CompartmentId = common.String(getCompartmentID())
	request.CpuCoreCount = common.Int(2)
	request.DatabaseEdition = "STANDARD_EDITION"
	request.DisplayName = common.String(dbSystemDisplayName)
	request.Shape = common.String("BM.DenseIO1.36") // this shape will not get service limit error for now

	buffer, err := readTestPubKey()
	failIfError(t, err)
	request.SshPublicKeys = []string{string(buffer)}

	subnet := createOrGetSubnet(t)
	request.SubnetId = subnet.Id
	request.Hostname = common.String("test")
	request.LicenseModel = "LICENSE_INCLUDED"
	request.DiskRedundancy = "NORMAL"

	request.DbHome = &database.CreateDbHomeDetails{
		DbVersion: common.String("11.2.0.4"),
		Database: &database.CreateDatabaseDetails{
			DbName:        common.String("Name"),
			AdminPassword: common.String("OraclE12--"),
			DbWorkload:    "OLTP",
		},
	}

	r, err := c.LaunchDbSystem(context.Background(), request)
	failIfError(t, err)

	// if we've successfully created a group during testing, make sure that we delete it
	defer func() {
		//Delete
		rDel := database.TerminateDbSystemRequest{DbSystemId: r.DbSystem.Id}
		resp, err := c.TerminateDbSystem(context.Background(), rDel)
		failIfError(t, err)
		assert.NotEmpty(t, resp.OpcRequestId)
	}()
	assert.NotEmpty(t, r)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

func TestDatabaseClient_ListBackups(t *testing.T) {
	t.Skip("Not implemented")
	c, clerr := database.NewDatabaseClientWithConfigurationProvider(common.DefaultConfigProvider())
	failIfError(t, clerr)
	request := database.ListBackupsRequest{}
	r, err := c.ListBackups(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

func TestDatabaseClient_ListDataGuardAssociations(t *testing.T) {
	t.Skip("Not implemented")
	c, clerr := database.NewDatabaseClientWithConfigurationProvider(common.DefaultConfigProvider())
	failIfError(t, clerr)
	request := database.ListDataGuardAssociationsRequest{}
	r, err := c.ListDataGuardAssociations(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

func TestDatabaseClient_ListDatabases(t *testing.T) {
	t.Skip("Not implemented")
	c, clerr := database.NewDatabaseClientWithConfigurationProvider(common.DefaultConfigProvider())
	failIfError(t, clerr)
	request := database.ListDatabasesRequest{}
	r, err := c.ListDatabases(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

func TestDatabaseClient_ListDbHomePatchHistoryEntries(t *testing.T) {
	t.Skip("Not implemented")
	c, clerr := database.NewDatabaseClientWithConfigurationProvider(common.DefaultConfigProvider())
	failIfError(t, clerr)
	request := database.ListDbHomePatchHistoryEntriesRequest{}
	r, err := c.ListDbHomePatchHistoryEntries(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

func TestDatabaseClient_ListDbHomePatches(t *testing.T) {
	t.Skip("Not implemented")
	c, clerr := database.NewDatabaseClientWithConfigurationProvider(common.DefaultConfigProvider())
	failIfError(t, clerr)
	request := database.ListDbHomePatchesRequest{}
	r, err := c.ListDbHomePatches(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

func TestDatabaseClient_ListDbHomes(t *testing.T) {
	t.Skip("Not implemented")
	c, clerr := database.NewDatabaseClientWithConfigurationProvider(common.DefaultConfigProvider())
	failIfError(t, clerr)
	request := database.ListDbHomesRequest{}
	r, err := c.ListDbHomes(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

func TestDatabaseClient_ListDbNodes(t *testing.T) {
	t.Skip("Not implemented")
	c, clerr := database.NewDatabaseClientWithConfigurationProvider(common.DefaultConfigProvider())
	failIfError(t, clerr)
	request := database.ListDbNodesRequest{}
	r, err := c.ListDbNodes(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

func TestDatabaseClient_ListDbSystemPatchHistoryEntries(t *testing.T) {
	t.Skip("Not implemented")
	c, clerr := database.NewDatabaseClientWithConfigurationProvider(common.DefaultConfigProvider())
	failIfError(t, clerr)
	request := database.ListDbSystemPatchHistoryEntriesRequest{}
	r, err := c.ListDbSystemPatchHistoryEntries(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

func TestDatabaseClient_ListDbSystemPatches(t *testing.T) {
	t.Skip("Not implemented")
	c, clerr := database.NewDatabaseClientWithConfigurationProvider(common.DefaultConfigProvider())
	failIfError(t, clerr)
	request := database.ListDbSystemPatchesRequest{}
	r, err := c.ListDbSystemPatches(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

func TestDatabaseClient_ListDbSystemShapes(t *testing.T) {
	t.Skip("Not implemented")
	c, clerr := database.NewDatabaseClientWithConfigurationProvider(common.DefaultConfigProvider())
	failIfError(t, clerr)
	request := database.ListDbSystemShapesRequest{}
	r, err := c.ListDbSystemShapes(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

func TestDatabaseClient_ListDbSystems(t *testing.T) {
	t.Skip("Not implemented")
	c, clerr := database.NewDatabaseClientWithConfigurationProvider(common.DefaultConfigProvider())
	failIfError(t, clerr)
	request := database.ListDbSystemsRequest{}
	r, err := c.ListDbSystems(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

func TestDatabaseClient_ListDbVersions(t *testing.T) {
	t.Skip("Not implemented")
	c, clerr := database.NewDatabaseClientWithConfigurationProvider(common.DefaultConfigProvider())
	failIfError(t, clerr)
	request := database.ListDbVersionsRequest{}
	r, err := c.ListDbVersions(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

func TestDatabaseClient_ReinstateDataGuardAssociation(t *testing.T) {
	t.Skip("Not implemented")
	c, clerr := database.NewDatabaseClientWithConfigurationProvider(common.DefaultConfigProvider())
	failIfError(t, clerr)
	request := database.ReinstateDataGuardAssociationRequest{}
	r, err := c.ReinstateDataGuardAssociation(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

func TestDatabaseClient_RestoreDatabase(t *testing.T) {
	t.Skip("Not implemented")
	c, clerr := database.NewDatabaseClientWithConfigurationProvider(common.DefaultConfigProvider())
	failIfError(t, clerr)
	request := database.RestoreDatabaseRequest{}
	r, err := c.RestoreDatabase(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

func TestDatabaseClient_SwitchoverDataGuardAssociation(t *testing.T) {
	t.Skip("Not implemented")
	c, clerr := database.NewDatabaseClientWithConfigurationProvider(common.DefaultConfigProvider())
	failIfError(t, clerr)
	request := database.SwitchoverDataGuardAssociationRequest{}
	r, err := c.SwitchoverDataGuardAssociation(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

func TestDatabaseClient_TerminateDbSystem(t *testing.T) {
	t.Skip("Not implemented")
	c, clerr := database.NewDatabaseClientWithConfigurationProvider(common.DefaultConfigProvider())
	failIfError(t, clerr)
	request := database.TerminateDbSystemRequest{}
	_, err := c.TerminateDbSystem(context.Background(), request)
	assert.NoError(t, err)
	return
}

func TestDatabaseClient_UpdateDatabase(t *testing.T) {
	t.Skip("Not implemented")
	c, clerr := database.NewDatabaseClientWithConfigurationProvider(common.DefaultConfigProvider())
	failIfError(t, clerr)
	request := database.UpdateDatabaseRequest{}
	r, err := c.UpdateDatabase(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

func TestDatabaseClient_UpdateDbHome(t *testing.T) {
	t.Skip("Not implemented")
	c, clerr := database.NewDatabaseClientWithConfigurationProvider(common.DefaultConfigProvider())
	failIfError(t, clerr)
	request := database.UpdateDbHomeRequest{}
	r, err := c.UpdateDbHome(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

func TestDatabaseClient_UpdateDbSystem(t *testing.T) {
	t.Skip("Not implemented")
	c, clerr := database.NewDatabaseClientWithConfigurationProvider(common.DefaultConfigProvider())
	failIfError(t, clerr)
	request := database.UpdateDbSystemRequest{}
	r, err := c.UpdateDbSystem(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}
