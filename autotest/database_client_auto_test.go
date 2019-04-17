package autotest

import (
	"github.com/oracle/oci-go-sdk/common"
	"github.com/oracle/oci-go-sdk/database"

	"context"
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func createDatabaseClientWithProvider(p common.ConfigurationProvider, testConfig TestingConfig) (interface{}, error) {

	client, err := database.NewDatabaseClientWithConfigurationProvider(p)
	if testConfig.Endpoint != "" {
		client.Host = testConfig.Endpoint
	} else {
		client.SetRegion(testConfig.Region)
	}
	return client, err
}

// IssueRoutingInfo tag="default" email="sic_dbaas_cp_us_grp@oracle.com" jiraProject="DBAAS" opsJiraProject="DBAASOPS"
func TestDatabaseClientCompleteExternalBackupJob(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("database", "CompleteExternalBackupJob")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("CompleteExternalBackupJob is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("database", "Database", "CompleteExternalBackupJob", createDatabaseClientWithProvider)
	assert.NoError(t, err)
	c := cc.(database.DatabaseClient)

	body, err := testClient.getRequests("database", "CompleteExternalBackupJob")
	assert.NoError(t, err)

	type CompleteExternalBackupJobRequestInfo struct {
		ContainerId string
		Request     database.CompleteExternalBackupJobRequest
	}

	var requests []CompleteExternalBackupJobRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.CompleteExternalBackupJob(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="sic_dbaas_cp_us_grp@oracle.com" jiraProject="DBAAS" opsJiraProject="DBAASOPS"
func TestDatabaseClientCreateAutonomousDataWarehouse(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("database", "CreateAutonomousDataWarehouse")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("CreateAutonomousDataWarehouse is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("database", "Database", "CreateAutonomousDataWarehouse", createDatabaseClientWithProvider)
	assert.NoError(t, err)
	c := cc.(database.DatabaseClient)

	body, err := testClient.getRequests("database", "CreateAutonomousDataWarehouse")
	assert.NoError(t, err)

	type CreateAutonomousDataWarehouseRequestInfo struct {
		ContainerId string
		Request     database.CreateAutonomousDataWarehouseRequest
	}

	var requests []CreateAutonomousDataWarehouseRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.CreateAutonomousDataWarehouse(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="sic_dbaas_cp_us_grp@oracle.com" jiraProject="DBAAS" opsJiraProject="DBAASOPS"
func TestDatabaseClientCreateAutonomousDataWarehouseBackup(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("database", "CreateAutonomousDataWarehouseBackup")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("CreateAutonomousDataWarehouseBackup is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("database", "Database", "CreateAutonomousDataWarehouseBackup", createDatabaseClientWithProvider)
	assert.NoError(t, err)
	c := cc.(database.DatabaseClient)

	body, err := testClient.getRequests("database", "CreateAutonomousDataWarehouseBackup")
	assert.NoError(t, err)

	type CreateAutonomousDataWarehouseBackupRequestInfo struct {
		ContainerId string
		Request     database.CreateAutonomousDataWarehouseBackupRequest
	}

	var requests []CreateAutonomousDataWarehouseBackupRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.CreateAutonomousDataWarehouseBackup(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="dbaas-atp-d" email="sic_dbaas_cp_us_grp@oracle.com" jiraProject="DBAAS" opsJiraProject="DBAASOPS"
func TestDatabaseClientCreateAutonomousDatabase(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("database", "CreateAutonomousDatabase")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("CreateAutonomousDatabase is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("database", "Database", "CreateAutonomousDatabase", createDatabaseClientWithProvider)
	assert.NoError(t, err)
	c := cc.(database.DatabaseClient)

	body, err := testClient.getRequests("database", "CreateAutonomousDatabase")
	assert.NoError(t, err)

	type CreateAutonomousDatabaseRequestInfo struct {
		ContainerId string
		Request     database.CreateAutonomousDatabaseRequest
	}

	var requests []CreateAutonomousDatabaseRequestInfo
	var pr []map[string]interface{}
	err = json.Unmarshal([]byte(body), &pr)
	assert.NoError(t, err)
	requests = make([]CreateAutonomousDatabaseRequestInfo, len(pr))
	polymorphicRequestInfo := map[string]PolymorphicRequestUnmarshallingInfo{}
	polymorphicRequestInfo["CreateAutonomousDatabaseBase"] =
		PolymorphicRequestUnmarshallingInfo{
			DiscriminatorName: "source",
			DiscriminatorValuesAndTypes: map[string]interface{}{
				"DATABASE": &database.CreateAutonomousDatabaseCloneDetails{},
				"NONE":     &database.CreateAutonomousDatabaseDetails{},
			},
		}

	for i, ppr := range pr {
		conditionalStructCopy(ppr, &requests[i], polymorphicRequestInfo, testClient.Log)
	}

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.CreateAutonomousDatabase(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="dbaas-atp-d" email="sic_dbaas_cp_us_grp@oracle.com" jiraProject="DBAAS" opsJiraProject="DBAASOPS"
func TestDatabaseClientCreateAutonomousDatabaseBackup(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("database", "CreateAutonomousDatabaseBackup")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("CreateAutonomousDatabaseBackup is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("database", "Database", "CreateAutonomousDatabaseBackup", createDatabaseClientWithProvider)
	assert.NoError(t, err)
	c := cc.(database.DatabaseClient)

	body, err := testClient.getRequests("database", "CreateAutonomousDatabaseBackup")
	assert.NoError(t, err)

	type CreateAutonomousDatabaseBackupRequestInfo struct {
		ContainerId string
		Request     database.CreateAutonomousDatabaseBackupRequest
	}

	var requests []CreateAutonomousDatabaseBackupRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.CreateAutonomousDatabaseBackup(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="sic_dbaas_cp_us_grp@oracle.com" jiraProject="DBAAS" opsJiraProject="DBAASOPS"
func TestDatabaseClientCreateBackup(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("database", "CreateBackup")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("CreateBackup is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("database", "Database", "CreateBackup", createDatabaseClientWithProvider)
	assert.NoError(t, err)
	c := cc.(database.DatabaseClient)

	body, err := testClient.getRequests("database", "CreateBackup")
	assert.NoError(t, err)

	type CreateBackupRequestInfo struct {
		ContainerId string
		Request     database.CreateBackupRequest
	}

	var requests []CreateBackupRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.CreateBackup(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="sic_dbaas_cp_us_grp@oracle.com" jiraProject="DBAAS" opsJiraProject="DBAASOPS"
func TestDatabaseClientCreateDataGuardAssociation(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("database", "CreateDataGuardAssociation")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("CreateDataGuardAssociation is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("database", "Database", "CreateDataGuardAssociation", createDatabaseClientWithProvider)
	assert.NoError(t, err)
	c := cc.(database.DatabaseClient)

	body, err := testClient.getRequests("database", "CreateDataGuardAssociation")
	assert.NoError(t, err)

	type CreateDataGuardAssociationRequestInfo struct {
		ContainerId string
		Request     database.CreateDataGuardAssociationRequest
	}

	var requests []CreateDataGuardAssociationRequestInfo
	var pr []map[string]interface{}
	err = json.Unmarshal([]byte(body), &pr)
	assert.NoError(t, err)
	requests = make([]CreateDataGuardAssociationRequestInfo, len(pr))
	polymorphicRequestInfo := map[string]PolymorphicRequestUnmarshallingInfo{}
	polymorphicRequestInfo["CreateDataGuardAssociationDetails"] =
		PolymorphicRequestUnmarshallingInfo{
			DiscriminatorName: "creationType",
			DiscriminatorValuesAndTypes: map[string]interface{}{
				"NewDbSystem":      &database.CreateDataGuardAssociationWithNewDbSystemDetails{},
				"ExistingDbSystem": &database.CreateDataGuardAssociationToExistingDbSystemDetails{},
			},
		}

	for i, ppr := range pr {
		conditionalStructCopy(ppr, &requests[i], polymorphicRequestInfo, testClient.Log)
	}

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.CreateDataGuardAssociation(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="sic_dbaas_cp_us_grp@oracle.com" jiraProject="DBAAS" opsJiraProject="DBAASOPS"
func TestDatabaseClientCreateDbHome(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("database", "CreateDbHome")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("CreateDbHome is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("database", "Database", "CreateDbHome", createDatabaseClientWithProvider)
	assert.NoError(t, err)
	c := cc.(database.DatabaseClient)

	body, err := testClient.getRequests("database", "CreateDbHome")
	assert.NoError(t, err)

	type CreateDbHomeRequestInfo struct {
		ContainerId string
		Request     database.CreateDbHomeRequest
	}

	var requests []CreateDbHomeRequestInfo
	var pr []map[string]interface{}
	err = json.Unmarshal([]byte(body), &pr)
	assert.NoError(t, err)
	requests = make([]CreateDbHomeRequestInfo, len(pr))
	polymorphicRequestInfo := map[string]PolymorphicRequestUnmarshallingInfo{}
	polymorphicRequestInfo["CreateDbHomeWithDbSystemIdBase"] =
		PolymorphicRequestUnmarshallingInfo{
			DiscriminatorName: "source",
			DiscriminatorValuesAndTypes: map[string]interface{}{
				"DB_BACKUP": &database.CreateDbHomeWithDbSystemIdFromBackupDetails{},
				"NONE":      &database.CreateDbHomeWithDbSystemIdDetails{},
			},
		}

	for i, ppr := range pr {
		conditionalStructCopy(ppr, &requests[i], polymorphicRequestInfo, testClient.Log)
	}

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.CreateDbHome(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="sic_dbaas_cp_us_grp@oracle.com" jiraProject="DBAAS" opsJiraProject="DBAASOPS"
func TestDatabaseClientCreateExternalBackupJob(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("database", "CreateExternalBackupJob")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("CreateExternalBackupJob is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("database", "Database", "CreateExternalBackupJob", createDatabaseClientWithProvider)
	assert.NoError(t, err)
	c := cc.(database.DatabaseClient)

	body, err := testClient.getRequests("database", "CreateExternalBackupJob")
	assert.NoError(t, err)

	type CreateExternalBackupJobRequestInfo struct {
		ContainerId string
		Request     database.CreateExternalBackupJobRequest
	}

	var requests []CreateExternalBackupJobRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.CreateExternalBackupJob(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="sic_dbaas_cp_us_grp@oracle.com" jiraProject="DBAAS" opsJiraProject="DBAASOPS"
func TestDatabaseClientDbNodeAction(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("database", "DbNodeAction")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("DbNodeAction is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("database", "Database", "DbNodeAction", createDatabaseClientWithProvider)
	assert.NoError(t, err)
	c := cc.(database.DatabaseClient)

	body, err := testClient.getRequests("database", "DbNodeAction")
	assert.NoError(t, err)

	type DbNodeActionRequestInfo struct {
		ContainerId string
		Request     database.DbNodeActionRequest
	}

	var requests []DbNodeActionRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.DbNodeAction(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="sic_dbaas_cp_us_grp@oracle.com" jiraProject="DBAAS" opsJiraProject="DBAASOPS"
func TestDatabaseClientDeleteAutonomousDataWarehouse(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("database", "DeleteAutonomousDataWarehouse")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("DeleteAutonomousDataWarehouse is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("database", "Database", "DeleteAutonomousDataWarehouse", createDatabaseClientWithProvider)
	assert.NoError(t, err)
	c := cc.(database.DatabaseClient)

	body, err := testClient.getRequests("database", "DeleteAutonomousDataWarehouse")
	assert.NoError(t, err)

	type DeleteAutonomousDataWarehouseRequestInfo struct {
		ContainerId string
		Request     database.DeleteAutonomousDataWarehouseRequest
	}

	var requests []DeleteAutonomousDataWarehouseRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.DeleteAutonomousDataWarehouse(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="dbaas-atp-d" email="sic_dbaas_cp_us_grp@oracle.com" jiraProject="DBAAS" opsJiraProject="DBAASOPS"
func TestDatabaseClientDeleteAutonomousDatabase(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("database", "DeleteAutonomousDatabase")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("DeleteAutonomousDatabase is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("database", "Database", "DeleteAutonomousDatabase", createDatabaseClientWithProvider)
	assert.NoError(t, err)
	c := cc.(database.DatabaseClient)

	body, err := testClient.getRequests("database", "DeleteAutonomousDatabase")
	assert.NoError(t, err)

	type DeleteAutonomousDatabaseRequestInfo struct {
		ContainerId string
		Request     database.DeleteAutonomousDatabaseRequest
	}

	var requests []DeleteAutonomousDatabaseRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.DeleteAutonomousDatabase(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="sic_dbaas_cp_us_grp@oracle.com" jiraProject="DBAAS" opsJiraProject="DBAASOPS"
func TestDatabaseClientDeleteBackup(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("database", "DeleteBackup")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("DeleteBackup is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("database", "Database", "DeleteBackup", createDatabaseClientWithProvider)
	assert.NoError(t, err)
	c := cc.(database.DatabaseClient)

	body, err := testClient.getRequests("database", "DeleteBackup")
	assert.NoError(t, err)

	type DeleteBackupRequestInfo struct {
		ContainerId string
		Request     database.DeleteBackupRequest
	}

	var requests []DeleteBackupRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.DeleteBackup(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="sic_dbaas_cp_us_grp@oracle.com" jiraProject="DBAAS" opsJiraProject="DBAASOPS"
func TestDatabaseClientDeleteDbHome(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("database", "DeleteDbHome")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("DeleteDbHome is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("database", "Database", "DeleteDbHome", createDatabaseClientWithProvider)
	assert.NoError(t, err)
	c := cc.(database.DatabaseClient)

	body, err := testClient.getRequests("database", "DeleteDbHome")
	assert.NoError(t, err)

	type DeleteDbHomeRequestInfo struct {
		ContainerId string
		Request     database.DeleteDbHomeRequest
	}

	var requests []DeleteDbHomeRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.DeleteDbHome(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="sic_dbaas_cp_us_grp@oracle.com" jiraProject="DBAAS" opsJiraProject="DBAASOPS"
func TestDatabaseClientFailoverDataGuardAssociation(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("database", "FailoverDataGuardAssociation")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("FailoverDataGuardAssociation is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("database", "Database", "FailoverDataGuardAssociation", createDatabaseClientWithProvider)
	assert.NoError(t, err)
	c := cc.(database.DatabaseClient)

	body, err := testClient.getRequests("database", "FailoverDataGuardAssociation")
	assert.NoError(t, err)

	type FailoverDataGuardAssociationRequestInfo struct {
		ContainerId string
		Request     database.FailoverDataGuardAssociationRequest
	}

	var requests []FailoverDataGuardAssociationRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.FailoverDataGuardAssociation(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="sic_dbaas_cp_us_grp@oracle.com" jiraProject="DBAAS" opsJiraProject="DBAASOPS"
func TestDatabaseClientGenerateAutonomousDataWarehouseWallet(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("database", "GenerateAutonomousDataWarehouseWallet")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GenerateAutonomousDataWarehouseWallet is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("database", "Database", "GenerateAutonomousDataWarehouseWallet", createDatabaseClientWithProvider)
	assert.NoError(t, err)
	c := cc.(database.DatabaseClient)

	body, err := testClient.getRequests("database", "GenerateAutonomousDataWarehouseWallet")
	assert.NoError(t, err)

	type GenerateAutonomousDataWarehouseWalletRequestInfo struct {
		ContainerId string
		Request     database.GenerateAutonomousDataWarehouseWalletRequest
	}

	var requests []GenerateAutonomousDataWarehouseWalletRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.GenerateAutonomousDataWarehouseWallet(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="dbaas-atp-d" email="sic_dbaas_cp_us_grp@oracle.com" jiraProject="DBAAS" opsJiraProject="DBAASOPS"
func TestDatabaseClientGenerateAutonomousDatabaseWallet(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("database", "GenerateAutonomousDatabaseWallet")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GenerateAutonomousDatabaseWallet is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("database", "Database", "GenerateAutonomousDatabaseWallet", createDatabaseClientWithProvider)
	assert.NoError(t, err)
	c := cc.(database.DatabaseClient)

	body, err := testClient.getRequests("database", "GenerateAutonomousDatabaseWallet")
	assert.NoError(t, err)

	type GenerateAutonomousDatabaseWalletRequestInfo struct {
		ContainerId string
		Request     database.GenerateAutonomousDatabaseWalletRequest
	}

	var requests []GenerateAutonomousDatabaseWalletRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.GenerateAutonomousDatabaseWallet(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="sic_dbaas_cp_us_grp@oracle.com" jiraProject="DBAAS" opsJiraProject="DBAASOPS"
func TestDatabaseClientGetAutonomousDataWarehouse(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("database", "GetAutonomousDataWarehouse")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetAutonomousDataWarehouse is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("database", "Database", "GetAutonomousDataWarehouse", createDatabaseClientWithProvider)
	assert.NoError(t, err)
	c := cc.(database.DatabaseClient)

	body, err := testClient.getRequests("database", "GetAutonomousDataWarehouse")
	assert.NoError(t, err)

	type GetAutonomousDataWarehouseRequestInfo struct {
		ContainerId string
		Request     database.GetAutonomousDataWarehouseRequest
	}

	var requests []GetAutonomousDataWarehouseRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.GetAutonomousDataWarehouse(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="sic_dbaas_cp_us_grp@oracle.com" jiraProject="DBAAS" opsJiraProject="DBAASOPS"
func TestDatabaseClientGetAutonomousDataWarehouseBackup(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("database", "GetAutonomousDataWarehouseBackup")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetAutonomousDataWarehouseBackup is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("database", "Database", "GetAutonomousDataWarehouseBackup", createDatabaseClientWithProvider)
	assert.NoError(t, err)
	c := cc.(database.DatabaseClient)

	body, err := testClient.getRequests("database", "GetAutonomousDataWarehouseBackup")
	assert.NoError(t, err)

	type GetAutonomousDataWarehouseBackupRequestInfo struct {
		ContainerId string
		Request     database.GetAutonomousDataWarehouseBackupRequest
	}

	var requests []GetAutonomousDataWarehouseBackupRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.GetAutonomousDataWarehouseBackup(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="dbaas-atp-d" email="sic_dbaas_cp_us_grp@oracle.com" jiraProject="DBAAS" opsJiraProject="DBAASOPS"
func TestDatabaseClientGetAutonomousDatabase(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("database", "GetAutonomousDatabase")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetAutonomousDatabase is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("database", "Database", "GetAutonomousDatabase", createDatabaseClientWithProvider)
	assert.NoError(t, err)
	c := cc.(database.DatabaseClient)

	body, err := testClient.getRequests("database", "GetAutonomousDatabase")
	assert.NoError(t, err)

	type GetAutonomousDatabaseRequestInfo struct {
		ContainerId string
		Request     database.GetAutonomousDatabaseRequest
	}

	var requests []GetAutonomousDatabaseRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.GetAutonomousDatabase(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="dbaas-atp-d" email="sic_dbaas_cp_us_grp@oracle.com" jiraProject="DBAAS" opsJiraProject="DBAASOPS"
func TestDatabaseClientGetAutonomousDatabaseBackup(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("database", "GetAutonomousDatabaseBackup")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetAutonomousDatabaseBackup is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("database", "Database", "GetAutonomousDatabaseBackup", createDatabaseClientWithProvider)
	assert.NoError(t, err)
	c := cc.(database.DatabaseClient)

	body, err := testClient.getRequests("database", "GetAutonomousDatabaseBackup")
	assert.NoError(t, err)

	type GetAutonomousDatabaseBackupRequestInfo struct {
		ContainerId string
		Request     database.GetAutonomousDatabaseBackupRequest
	}

	var requests []GetAutonomousDatabaseBackupRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.GetAutonomousDatabaseBackup(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="sic_dbaas_cp_us_grp@oracle.com" jiraProject="DBAAS" opsJiraProject="DBAASOPS"
func TestDatabaseClientGetBackup(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("database", "GetBackup")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetBackup is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("database", "Database", "GetBackup", createDatabaseClientWithProvider)
	assert.NoError(t, err)
	c := cc.(database.DatabaseClient)

	body, err := testClient.getRequests("database", "GetBackup")
	assert.NoError(t, err)

	type GetBackupRequestInfo struct {
		ContainerId string
		Request     database.GetBackupRequest
	}

	var requests []GetBackupRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.GetBackup(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="sic_dbaas_cp_us_grp@oracle.com" jiraProject="DBAAS" opsJiraProject="DBAASOPS"
func TestDatabaseClientGetDataGuardAssociation(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("database", "GetDataGuardAssociation")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetDataGuardAssociation is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("database", "Database", "GetDataGuardAssociation", createDatabaseClientWithProvider)
	assert.NoError(t, err)
	c := cc.(database.DatabaseClient)

	body, err := testClient.getRequests("database", "GetDataGuardAssociation")
	assert.NoError(t, err)

	type GetDataGuardAssociationRequestInfo struct {
		ContainerId string
		Request     database.GetDataGuardAssociationRequest
	}

	var requests []GetDataGuardAssociationRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.GetDataGuardAssociation(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="sic_dbaas_cp_us_grp@oracle.com" jiraProject="DBAAS" opsJiraProject="DBAASOPS"
func TestDatabaseClientGetDatabase(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("database", "GetDatabase")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetDatabase is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("database", "Database", "GetDatabase", createDatabaseClientWithProvider)
	assert.NoError(t, err)
	c := cc.(database.DatabaseClient)

	body, err := testClient.getRequests("database", "GetDatabase")
	assert.NoError(t, err)

	type GetDatabaseRequestInfo struct {
		ContainerId string
		Request     database.GetDatabaseRequest
	}

	var requests []GetDatabaseRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.GetDatabase(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="sic_dbaas_cp_us_grp@oracle.com" jiraProject="DBAAS" opsJiraProject="DBAASOPS"
func TestDatabaseClientGetDbHome(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("database", "GetDbHome")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetDbHome is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("database", "Database", "GetDbHome", createDatabaseClientWithProvider)
	assert.NoError(t, err)
	c := cc.(database.DatabaseClient)

	body, err := testClient.getRequests("database", "GetDbHome")
	assert.NoError(t, err)

	type GetDbHomeRequestInfo struct {
		ContainerId string
		Request     database.GetDbHomeRequest
	}

	var requests []GetDbHomeRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.GetDbHome(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="sic_dbaas_cp_us_grp@oracle.com" jiraProject="DBAAS" opsJiraProject="DBAASOPS"
func TestDatabaseClientGetDbHomePatch(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("database", "GetDbHomePatch")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetDbHomePatch is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("database", "Database", "GetDbHomePatch", createDatabaseClientWithProvider)
	assert.NoError(t, err)
	c := cc.(database.DatabaseClient)

	body, err := testClient.getRequests("database", "GetDbHomePatch")
	assert.NoError(t, err)

	type GetDbHomePatchRequestInfo struct {
		ContainerId string
		Request     database.GetDbHomePatchRequest
	}

	var requests []GetDbHomePatchRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.GetDbHomePatch(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="sic_dbaas_cp_us_grp@oracle.com" jiraProject="DBAAS" opsJiraProject="DBAASOPS"
func TestDatabaseClientGetDbHomePatchHistoryEntry(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("database", "GetDbHomePatchHistoryEntry")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetDbHomePatchHistoryEntry is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("database", "Database", "GetDbHomePatchHistoryEntry", createDatabaseClientWithProvider)
	assert.NoError(t, err)
	c := cc.(database.DatabaseClient)

	body, err := testClient.getRequests("database", "GetDbHomePatchHistoryEntry")
	assert.NoError(t, err)

	type GetDbHomePatchHistoryEntryRequestInfo struct {
		ContainerId string
		Request     database.GetDbHomePatchHistoryEntryRequest
	}

	var requests []GetDbHomePatchHistoryEntryRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.GetDbHomePatchHistoryEntry(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="sic_dbaas_cp_us_grp@oracle.com" jiraProject="DBAAS" opsJiraProject="DBAASOPS"
func TestDatabaseClientGetDbNode(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("database", "GetDbNode")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetDbNode is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("database", "Database", "GetDbNode", createDatabaseClientWithProvider)
	assert.NoError(t, err)
	c := cc.(database.DatabaseClient)

	body, err := testClient.getRequests("database", "GetDbNode")
	assert.NoError(t, err)

	type GetDbNodeRequestInfo struct {
		ContainerId string
		Request     database.GetDbNodeRequest
	}

	var requests []GetDbNodeRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.GetDbNode(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="sic_dbaas_cp_us_grp@oracle.com" jiraProject="DBAAS" opsJiraProject="DBAASOPS"
func TestDatabaseClientGetDbSystem(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("database", "GetDbSystem")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetDbSystem is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("database", "Database", "GetDbSystem", createDatabaseClientWithProvider)
	assert.NoError(t, err)
	c := cc.(database.DatabaseClient)

	body, err := testClient.getRequests("database", "GetDbSystem")
	assert.NoError(t, err)

	type GetDbSystemRequestInfo struct {
		ContainerId string
		Request     database.GetDbSystemRequest
	}

	var requests []GetDbSystemRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.GetDbSystem(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="sic_dbaas_cp_us_grp@oracle.com" jiraProject="DBAAS" opsJiraProject="DBAASOPS"
func TestDatabaseClientGetDbSystemPatch(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("database", "GetDbSystemPatch")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetDbSystemPatch is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("database", "Database", "GetDbSystemPatch", createDatabaseClientWithProvider)
	assert.NoError(t, err)
	c := cc.(database.DatabaseClient)

	body, err := testClient.getRequests("database", "GetDbSystemPatch")
	assert.NoError(t, err)

	type GetDbSystemPatchRequestInfo struct {
		ContainerId string
		Request     database.GetDbSystemPatchRequest
	}

	var requests []GetDbSystemPatchRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.GetDbSystemPatch(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="sic_dbaas_cp_us_grp@oracle.com" jiraProject="DBAAS" opsJiraProject="DBAASOPS"
func TestDatabaseClientGetDbSystemPatchHistoryEntry(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("database", "GetDbSystemPatchHistoryEntry")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetDbSystemPatchHistoryEntry is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("database", "Database", "GetDbSystemPatchHistoryEntry", createDatabaseClientWithProvider)
	assert.NoError(t, err)
	c := cc.(database.DatabaseClient)

	body, err := testClient.getRequests("database", "GetDbSystemPatchHistoryEntry")
	assert.NoError(t, err)

	type GetDbSystemPatchHistoryEntryRequestInfo struct {
		ContainerId string
		Request     database.GetDbSystemPatchHistoryEntryRequest
	}

	var requests []GetDbSystemPatchHistoryEntryRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.GetDbSystemPatchHistoryEntry(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="sic_dbaas_cp_us_grp@oracle.com" jiraProject="DBAAS" opsJiraProject="DBAASOPS"
func TestDatabaseClientGetExadataIormConfig(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("database", "GetExadataIormConfig")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetExadataIormConfig is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("database", "Database", "GetExadataIormConfig", createDatabaseClientWithProvider)
	assert.NoError(t, err)
	c := cc.(database.DatabaseClient)

	body, err := testClient.getRequests("database", "GetExadataIormConfig")
	assert.NoError(t, err)

	type GetExadataIormConfigRequestInfo struct {
		ContainerId string
		Request     database.GetExadataIormConfigRequest
	}

	var requests []GetExadataIormConfigRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.GetExadataIormConfig(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="sic_dbaas_cp_us_grp@oracle.com" jiraProject="DBAAS" opsJiraProject="DBAASOPS"
func TestDatabaseClientGetExternalBackupJob(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("database", "GetExternalBackupJob")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetExternalBackupJob is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("database", "Database", "GetExternalBackupJob", createDatabaseClientWithProvider)
	assert.NoError(t, err)
	c := cc.(database.DatabaseClient)

	body, err := testClient.getRequests("database", "GetExternalBackupJob")
	assert.NoError(t, err)

	type GetExternalBackupJobRequestInfo struct {
		ContainerId string
		Request     database.GetExternalBackupJobRequest
	}

	var requests []GetExternalBackupJobRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.GetExternalBackupJob(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="sic_dbaas_cp_us_grp@oracle.com" jiraProject="DBAAS" opsJiraProject="DBAASOPS"
func TestDatabaseClientLaunchDbSystem(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("database", "LaunchDbSystem")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("LaunchDbSystem is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("database", "Database", "LaunchDbSystem", createDatabaseClientWithProvider)
	assert.NoError(t, err)
	c := cc.(database.DatabaseClient)

	body, err := testClient.getRequests("database", "LaunchDbSystem")
	assert.NoError(t, err)

	type LaunchDbSystemRequestInfo struct {
		ContainerId string
		Request     database.LaunchDbSystemRequest
	}

	var requests []LaunchDbSystemRequestInfo
	var pr []map[string]interface{}
	err = json.Unmarshal([]byte(body), &pr)
	assert.NoError(t, err)
	requests = make([]LaunchDbSystemRequestInfo, len(pr))
	polymorphicRequestInfo := map[string]PolymorphicRequestUnmarshallingInfo{}
	polymorphicRequestInfo["LaunchDbSystemBase"] =
		PolymorphicRequestUnmarshallingInfo{
			DiscriminatorName: "source",
			DiscriminatorValuesAndTypes: map[string]interface{}{
				"NONE":      &database.LaunchDbSystemDetails{},
				"DB_BACKUP": &database.LaunchDbSystemFromBackupDetails{},
			},
		}

	for i, ppr := range pr {
		conditionalStructCopy(ppr, &requests[i], polymorphicRequestInfo, testClient.Log)
	}

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.LaunchDbSystem(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="sic_dbaas_cp_us_grp@oracle.com" jiraProject="DBAAS" opsJiraProject="DBAASOPS"
func TestDatabaseClientListAutonomousDataWarehouseBackups(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("database", "ListAutonomousDataWarehouseBackups")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListAutonomousDataWarehouseBackups is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("database", "Database", "ListAutonomousDataWarehouseBackups", createDatabaseClientWithProvider)
	assert.NoError(t, err)
	c := cc.(database.DatabaseClient)

	body, err := testClient.getRequests("database", "ListAutonomousDataWarehouseBackups")
	assert.NoError(t, err)

	type ListAutonomousDataWarehouseBackupsRequestInfo struct {
		ContainerId string
		Request     database.ListAutonomousDataWarehouseBackupsRequest
	}

	var requests []ListAutonomousDataWarehouseBackupsRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, request := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			request.Request.RequestMetadata.RetryPolicy = retryPolicy
			listFn := func(req common.OCIRequest) (common.OCIResponse, error) {
				r := req.(*database.ListAutonomousDataWarehouseBackupsRequest)
				return c.ListAutonomousDataWarehouseBackups(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]database.ListAutonomousDataWarehouseBackupsResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(database.ListAutonomousDataWarehouseBackupsResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="sic_dbaas_cp_us_grp@oracle.com" jiraProject="DBAAS" opsJiraProject="DBAASOPS"
func TestDatabaseClientListAutonomousDataWarehouses(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("database", "ListAutonomousDataWarehouses")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListAutonomousDataWarehouses is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("database", "Database", "ListAutonomousDataWarehouses", createDatabaseClientWithProvider)
	assert.NoError(t, err)
	c := cc.(database.DatabaseClient)

	body, err := testClient.getRequests("database", "ListAutonomousDataWarehouses")
	assert.NoError(t, err)

	type ListAutonomousDataWarehousesRequestInfo struct {
		ContainerId string
		Request     database.ListAutonomousDataWarehousesRequest
	}

	var requests []ListAutonomousDataWarehousesRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, request := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			request.Request.RequestMetadata.RetryPolicy = retryPolicy
			listFn := func(req common.OCIRequest) (common.OCIResponse, error) {
				r := req.(*database.ListAutonomousDataWarehousesRequest)
				return c.ListAutonomousDataWarehouses(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]database.ListAutonomousDataWarehousesResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(database.ListAutonomousDataWarehousesResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="dbaas-atp-d" email="sic_dbaas_cp_us_grp@oracle.com" jiraProject="DBAAS" opsJiraProject="DBAASOPS"
func TestDatabaseClientListAutonomousDatabaseBackups(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("database", "ListAutonomousDatabaseBackups")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListAutonomousDatabaseBackups is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("database", "Database", "ListAutonomousDatabaseBackups", createDatabaseClientWithProvider)
	assert.NoError(t, err)
	c := cc.(database.DatabaseClient)

	body, err := testClient.getRequests("database", "ListAutonomousDatabaseBackups")
	assert.NoError(t, err)

	type ListAutonomousDatabaseBackupsRequestInfo struct {
		ContainerId string
		Request     database.ListAutonomousDatabaseBackupsRequest
	}

	var requests []ListAutonomousDatabaseBackupsRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, request := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			request.Request.RequestMetadata.RetryPolicy = retryPolicy
			listFn := func(req common.OCIRequest) (common.OCIResponse, error) {
				r := req.(*database.ListAutonomousDatabaseBackupsRequest)
				return c.ListAutonomousDatabaseBackups(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]database.ListAutonomousDatabaseBackupsResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(database.ListAutonomousDatabaseBackupsResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="dbaas-atp-d" email="sic_dbaas_cp_us_grp@oracle.com" jiraProject="DBAAS" opsJiraProject="DBAASOPS"
func TestDatabaseClientListAutonomousDatabases(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("database", "ListAutonomousDatabases")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListAutonomousDatabases is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("database", "Database", "ListAutonomousDatabases", createDatabaseClientWithProvider)
	assert.NoError(t, err)
	c := cc.(database.DatabaseClient)

	body, err := testClient.getRequests("database", "ListAutonomousDatabases")
	assert.NoError(t, err)

	type ListAutonomousDatabasesRequestInfo struct {
		ContainerId string
		Request     database.ListAutonomousDatabasesRequest
	}

	var requests []ListAutonomousDatabasesRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, request := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			request.Request.RequestMetadata.RetryPolicy = retryPolicy
			listFn := func(req common.OCIRequest) (common.OCIResponse, error) {
				r := req.(*database.ListAutonomousDatabasesRequest)
				return c.ListAutonomousDatabases(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]database.ListAutonomousDatabasesResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(database.ListAutonomousDatabasesResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="sic_dbaas_cp_us_grp@oracle.com" jiraProject="DBAAS" opsJiraProject="DBAASOPS"
func TestDatabaseClientListBackups(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("database", "ListBackups")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListBackups is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("database", "Database", "ListBackups", createDatabaseClientWithProvider)
	assert.NoError(t, err)
	c := cc.(database.DatabaseClient)

	body, err := testClient.getRequests("database", "ListBackups")
	assert.NoError(t, err)

	type ListBackupsRequestInfo struct {
		ContainerId string
		Request     database.ListBackupsRequest
	}

	var requests []ListBackupsRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, request := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			request.Request.RequestMetadata.RetryPolicy = retryPolicy
			listFn := func(req common.OCIRequest) (common.OCIResponse, error) {
				r := req.(*database.ListBackupsRequest)
				return c.ListBackups(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]database.ListBackupsResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(database.ListBackupsResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="sic_dbaas_cp_us_grp@oracle.com" jiraProject="DBAAS" opsJiraProject="DBAASOPS"
func TestDatabaseClientListDataGuardAssociations(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("database", "ListDataGuardAssociations")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListDataGuardAssociations is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("database", "Database", "ListDataGuardAssociations", createDatabaseClientWithProvider)
	assert.NoError(t, err)
	c := cc.(database.DatabaseClient)

	body, err := testClient.getRequests("database", "ListDataGuardAssociations")
	assert.NoError(t, err)

	type ListDataGuardAssociationsRequestInfo struct {
		ContainerId string
		Request     database.ListDataGuardAssociationsRequest
	}

	var requests []ListDataGuardAssociationsRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, request := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			request.Request.RequestMetadata.RetryPolicy = retryPolicy
			listFn := func(req common.OCIRequest) (common.OCIResponse, error) {
				r := req.(*database.ListDataGuardAssociationsRequest)
				return c.ListDataGuardAssociations(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]database.ListDataGuardAssociationsResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(database.ListDataGuardAssociationsResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="sic_dbaas_cp_us_grp@oracle.com" jiraProject="DBAAS" opsJiraProject="DBAASOPS"
func TestDatabaseClientListDatabases(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("database", "ListDatabases")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListDatabases is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("database", "Database", "ListDatabases", createDatabaseClientWithProvider)
	assert.NoError(t, err)
	c := cc.(database.DatabaseClient)

	body, err := testClient.getRequests("database", "ListDatabases")
	assert.NoError(t, err)

	type ListDatabasesRequestInfo struct {
		ContainerId string
		Request     database.ListDatabasesRequest
	}

	var requests []ListDatabasesRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, request := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			request.Request.RequestMetadata.RetryPolicy = retryPolicy
			listFn := func(req common.OCIRequest) (common.OCIResponse, error) {
				r := req.(*database.ListDatabasesRequest)
				return c.ListDatabases(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]database.ListDatabasesResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(database.ListDatabasesResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="sic_dbaas_cp_us_grp@oracle.com" jiraProject="DBAAS" opsJiraProject="DBAASOPS"
func TestDatabaseClientListDbHomePatchHistoryEntries(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("database", "ListDbHomePatchHistoryEntries")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListDbHomePatchHistoryEntries is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("database", "Database", "ListDbHomePatchHistoryEntries", createDatabaseClientWithProvider)
	assert.NoError(t, err)
	c := cc.(database.DatabaseClient)

	body, err := testClient.getRequests("database", "ListDbHomePatchHistoryEntries")
	assert.NoError(t, err)

	type ListDbHomePatchHistoryEntriesRequestInfo struct {
		ContainerId string
		Request     database.ListDbHomePatchHistoryEntriesRequest
	}

	var requests []ListDbHomePatchHistoryEntriesRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, request := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			request.Request.RequestMetadata.RetryPolicy = retryPolicy
			listFn := func(req common.OCIRequest) (common.OCIResponse, error) {
				r := req.(*database.ListDbHomePatchHistoryEntriesRequest)
				return c.ListDbHomePatchHistoryEntries(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]database.ListDbHomePatchHistoryEntriesResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(database.ListDbHomePatchHistoryEntriesResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="sic_dbaas_cp_us_grp@oracle.com" jiraProject="DBAAS" opsJiraProject="DBAASOPS"
func TestDatabaseClientListDbHomePatches(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("database", "ListDbHomePatches")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListDbHomePatches is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("database", "Database", "ListDbHomePatches", createDatabaseClientWithProvider)
	assert.NoError(t, err)
	c := cc.(database.DatabaseClient)

	body, err := testClient.getRequests("database", "ListDbHomePatches")
	assert.NoError(t, err)

	type ListDbHomePatchesRequestInfo struct {
		ContainerId string
		Request     database.ListDbHomePatchesRequest
	}

	var requests []ListDbHomePatchesRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, request := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			request.Request.RequestMetadata.RetryPolicy = retryPolicy
			listFn := func(req common.OCIRequest) (common.OCIResponse, error) {
				r := req.(*database.ListDbHomePatchesRequest)
				return c.ListDbHomePatches(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]database.ListDbHomePatchesResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(database.ListDbHomePatchesResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="sic_dbaas_cp_us_grp@oracle.com" jiraProject="DBAAS" opsJiraProject="DBAASOPS"
func TestDatabaseClientListDbHomes(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("database", "ListDbHomes")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListDbHomes is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("database", "Database", "ListDbHomes", createDatabaseClientWithProvider)
	assert.NoError(t, err)
	c := cc.(database.DatabaseClient)

	body, err := testClient.getRequests("database", "ListDbHomes")
	assert.NoError(t, err)

	type ListDbHomesRequestInfo struct {
		ContainerId string
		Request     database.ListDbHomesRequest
	}

	var requests []ListDbHomesRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, request := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			request.Request.RequestMetadata.RetryPolicy = retryPolicy
			listFn := func(req common.OCIRequest) (common.OCIResponse, error) {
				r := req.(*database.ListDbHomesRequest)
				return c.ListDbHomes(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]database.ListDbHomesResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(database.ListDbHomesResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="sic_dbaas_cp_us_grp@oracle.com" jiraProject="DBAAS" opsJiraProject="DBAASOPS"
func TestDatabaseClientListDbNodes(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("database", "ListDbNodes")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListDbNodes is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("database", "Database", "ListDbNodes", createDatabaseClientWithProvider)
	assert.NoError(t, err)
	c := cc.(database.DatabaseClient)

	body, err := testClient.getRequests("database", "ListDbNodes")
	assert.NoError(t, err)

	type ListDbNodesRequestInfo struct {
		ContainerId string
		Request     database.ListDbNodesRequest
	}

	var requests []ListDbNodesRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, request := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			request.Request.RequestMetadata.RetryPolicy = retryPolicy
			listFn := func(req common.OCIRequest) (common.OCIResponse, error) {
				r := req.(*database.ListDbNodesRequest)
				return c.ListDbNodes(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]database.ListDbNodesResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(database.ListDbNodesResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="sic_dbaas_cp_us_grp@oracle.com" jiraProject="DBAAS" opsJiraProject="DBAASOPS"
func TestDatabaseClientListDbSystemPatchHistoryEntries(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("database", "ListDbSystemPatchHistoryEntries")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListDbSystemPatchHistoryEntries is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("database", "Database", "ListDbSystemPatchHistoryEntries", createDatabaseClientWithProvider)
	assert.NoError(t, err)
	c := cc.(database.DatabaseClient)

	body, err := testClient.getRequests("database", "ListDbSystemPatchHistoryEntries")
	assert.NoError(t, err)

	type ListDbSystemPatchHistoryEntriesRequestInfo struct {
		ContainerId string
		Request     database.ListDbSystemPatchHistoryEntriesRequest
	}

	var requests []ListDbSystemPatchHistoryEntriesRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, request := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			request.Request.RequestMetadata.RetryPolicy = retryPolicy
			listFn := func(req common.OCIRequest) (common.OCIResponse, error) {
				r := req.(*database.ListDbSystemPatchHistoryEntriesRequest)
				return c.ListDbSystemPatchHistoryEntries(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]database.ListDbSystemPatchHistoryEntriesResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(database.ListDbSystemPatchHistoryEntriesResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="sic_dbaas_cp_us_grp@oracle.com" jiraProject="DBAAS" opsJiraProject="DBAASOPS"
func TestDatabaseClientListDbSystemPatches(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("database", "ListDbSystemPatches")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListDbSystemPatches is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("database", "Database", "ListDbSystemPatches", createDatabaseClientWithProvider)
	assert.NoError(t, err)
	c := cc.(database.DatabaseClient)

	body, err := testClient.getRequests("database", "ListDbSystemPatches")
	assert.NoError(t, err)

	type ListDbSystemPatchesRequestInfo struct {
		ContainerId string
		Request     database.ListDbSystemPatchesRequest
	}

	var requests []ListDbSystemPatchesRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, request := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			request.Request.RequestMetadata.RetryPolicy = retryPolicy
			listFn := func(req common.OCIRequest) (common.OCIResponse, error) {
				r := req.(*database.ListDbSystemPatchesRequest)
				return c.ListDbSystemPatches(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]database.ListDbSystemPatchesResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(database.ListDbSystemPatchesResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="sic_dbaas_cp_us_grp@oracle.com" jiraProject="DBAAS" opsJiraProject="DBAASOPS"
func TestDatabaseClientListDbSystemShapes(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("database", "ListDbSystemShapes")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListDbSystemShapes is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("database", "Database", "ListDbSystemShapes", createDatabaseClientWithProvider)
	assert.NoError(t, err)
	c := cc.(database.DatabaseClient)

	body, err := testClient.getRequests("database", "ListDbSystemShapes")
	assert.NoError(t, err)

	type ListDbSystemShapesRequestInfo struct {
		ContainerId string
		Request     database.ListDbSystemShapesRequest
	}

	var requests []ListDbSystemShapesRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, request := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			request.Request.RequestMetadata.RetryPolicy = retryPolicy
			listFn := func(req common.OCIRequest) (common.OCIResponse, error) {
				r := req.(*database.ListDbSystemShapesRequest)
				return c.ListDbSystemShapes(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]database.ListDbSystemShapesResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(database.ListDbSystemShapesResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="sic_dbaas_cp_us_grp@oracle.com" jiraProject="DBAAS" opsJiraProject="DBAASOPS"
func TestDatabaseClientListDbSystems(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("database", "ListDbSystems")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListDbSystems is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("database", "Database", "ListDbSystems", createDatabaseClientWithProvider)
	assert.NoError(t, err)
	c := cc.(database.DatabaseClient)

	body, err := testClient.getRequests("database", "ListDbSystems")
	assert.NoError(t, err)

	type ListDbSystemsRequestInfo struct {
		ContainerId string
		Request     database.ListDbSystemsRequest
	}

	var requests []ListDbSystemsRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, request := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			request.Request.RequestMetadata.RetryPolicy = retryPolicy
			listFn := func(req common.OCIRequest) (common.OCIResponse, error) {
				r := req.(*database.ListDbSystemsRequest)
				return c.ListDbSystems(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]database.ListDbSystemsResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(database.ListDbSystemsResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="sic_dbaas_cp_us_grp@oracle.com" jiraProject="DBAAS" opsJiraProject="DBAASOPS"
func TestDatabaseClientListDbVersions(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("database", "ListDbVersions")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListDbVersions is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("database", "Database", "ListDbVersions", createDatabaseClientWithProvider)
	assert.NoError(t, err)
	c := cc.(database.DatabaseClient)

	body, err := testClient.getRequests("database", "ListDbVersions")
	assert.NoError(t, err)

	type ListDbVersionsRequestInfo struct {
		ContainerId string
		Request     database.ListDbVersionsRequest
	}

	var requests []ListDbVersionsRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, request := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			request.Request.RequestMetadata.RetryPolicy = retryPolicy
			listFn := func(req common.OCIRequest) (common.OCIResponse, error) {
				r := req.(*database.ListDbVersionsRequest)
				return c.ListDbVersions(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]database.ListDbVersionsResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(database.ListDbVersionsResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="sic_dbaas_cp_us_grp@oracle.com" jiraProject="DBAAS" opsJiraProject="DBAASOPS"
func TestDatabaseClientReinstateDataGuardAssociation(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("database", "ReinstateDataGuardAssociation")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ReinstateDataGuardAssociation is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("database", "Database", "ReinstateDataGuardAssociation", createDatabaseClientWithProvider)
	assert.NoError(t, err)
	c := cc.(database.DatabaseClient)

	body, err := testClient.getRequests("database", "ReinstateDataGuardAssociation")
	assert.NoError(t, err)

	type ReinstateDataGuardAssociationRequestInfo struct {
		ContainerId string
		Request     database.ReinstateDataGuardAssociationRequest
	}

	var requests []ReinstateDataGuardAssociationRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.ReinstateDataGuardAssociation(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="sic_dbaas_cp_us_grp@oracle.com" jiraProject="DBAAS" opsJiraProject="DBAASOPS"
func TestDatabaseClientRestoreAutonomousDataWarehouse(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("database", "RestoreAutonomousDataWarehouse")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("RestoreAutonomousDataWarehouse is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("database", "Database", "RestoreAutonomousDataWarehouse", createDatabaseClientWithProvider)
	assert.NoError(t, err)
	c := cc.(database.DatabaseClient)

	body, err := testClient.getRequests("database", "RestoreAutonomousDataWarehouse")
	assert.NoError(t, err)

	type RestoreAutonomousDataWarehouseRequestInfo struct {
		ContainerId string
		Request     database.RestoreAutonomousDataWarehouseRequest
	}

	var requests []RestoreAutonomousDataWarehouseRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.RestoreAutonomousDataWarehouse(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="dbaas-atp-d" email="sic_dbaas_cp_us_grp@oracle.com" jiraProject="DBAAS" opsJiraProject="DBAASOPS"
func TestDatabaseClientRestoreAutonomousDatabase(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("database", "RestoreAutonomousDatabase")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("RestoreAutonomousDatabase is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("database", "Database", "RestoreAutonomousDatabase", createDatabaseClientWithProvider)
	assert.NoError(t, err)
	c := cc.(database.DatabaseClient)

	body, err := testClient.getRequests("database", "RestoreAutonomousDatabase")
	assert.NoError(t, err)

	type RestoreAutonomousDatabaseRequestInfo struct {
		ContainerId string
		Request     database.RestoreAutonomousDatabaseRequest
	}

	var requests []RestoreAutonomousDatabaseRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.RestoreAutonomousDatabase(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="sic_dbaas_cp_us_grp@oracle.com" jiraProject="DBAAS" opsJiraProject="DBAASOPS"
func TestDatabaseClientRestoreDatabase(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("database", "RestoreDatabase")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("RestoreDatabase is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("database", "Database", "RestoreDatabase", createDatabaseClientWithProvider)
	assert.NoError(t, err)
	c := cc.(database.DatabaseClient)

	body, err := testClient.getRequests("database", "RestoreDatabase")
	assert.NoError(t, err)

	type RestoreDatabaseRequestInfo struct {
		ContainerId string
		Request     database.RestoreDatabaseRequest
	}

	var requests []RestoreDatabaseRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.RestoreDatabase(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="sic_dbaas_cp_us_grp@oracle.com" jiraProject="DBAAS" opsJiraProject="DBAASOPS"
func TestDatabaseClientStartAutonomousDataWarehouse(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("database", "StartAutonomousDataWarehouse")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("StartAutonomousDataWarehouse is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("database", "Database", "StartAutonomousDataWarehouse", createDatabaseClientWithProvider)
	assert.NoError(t, err)
	c := cc.(database.DatabaseClient)

	body, err := testClient.getRequests("database", "StartAutonomousDataWarehouse")
	assert.NoError(t, err)

	type StartAutonomousDataWarehouseRequestInfo struct {
		ContainerId string
		Request     database.StartAutonomousDataWarehouseRequest
	}

	var requests []StartAutonomousDataWarehouseRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.StartAutonomousDataWarehouse(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="dbaas-atp-d" email="sic_dbaas_cp_us_grp@oracle.com" jiraProject="DBAAS" opsJiraProject="DBAASOPS"
func TestDatabaseClientStartAutonomousDatabase(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("database", "StartAutonomousDatabase")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("StartAutonomousDatabase is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("database", "Database", "StartAutonomousDatabase", createDatabaseClientWithProvider)
	assert.NoError(t, err)
	c := cc.(database.DatabaseClient)

	body, err := testClient.getRequests("database", "StartAutonomousDatabase")
	assert.NoError(t, err)

	type StartAutonomousDatabaseRequestInfo struct {
		ContainerId string
		Request     database.StartAutonomousDatabaseRequest
	}

	var requests []StartAutonomousDatabaseRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.StartAutonomousDatabase(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="sic_dbaas_cp_us_grp@oracle.com" jiraProject="DBAAS" opsJiraProject="DBAASOPS"
func TestDatabaseClientStopAutonomousDataWarehouse(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("database", "StopAutonomousDataWarehouse")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("StopAutonomousDataWarehouse is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("database", "Database", "StopAutonomousDataWarehouse", createDatabaseClientWithProvider)
	assert.NoError(t, err)
	c := cc.(database.DatabaseClient)

	body, err := testClient.getRequests("database", "StopAutonomousDataWarehouse")
	assert.NoError(t, err)

	type StopAutonomousDataWarehouseRequestInfo struct {
		ContainerId string
		Request     database.StopAutonomousDataWarehouseRequest
	}

	var requests []StopAutonomousDataWarehouseRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.StopAutonomousDataWarehouse(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="dbaas-atp-d" email="sic_dbaas_cp_us_grp@oracle.com" jiraProject="DBAAS" opsJiraProject="DBAASOPS"
func TestDatabaseClientStopAutonomousDatabase(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("database", "StopAutonomousDatabase")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("StopAutonomousDatabase is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("database", "Database", "StopAutonomousDatabase", createDatabaseClientWithProvider)
	assert.NoError(t, err)
	c := cc.(database.DatabaseClient)

	body, err := testClient.getRequests("database", "StopAutonomousDatabase")
	assert.NoError(t, err)

	type StopAutonomousDatabaseRequestInfo struct {
		ContainerId string
		Request     database.StopAutonomousDatabaseRequest
	}

	var requests []StopAutonomousDatabaseRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.StopAutonomousDatabase(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="sic_dbaas_cp_us_grp@oracle.com" jiraProject="DBAAS" opsJiraProject="DBAASOPS"
func TestDatabaseClientSwitchoverDataGuardAssociation(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("database", "SwitchoverDataGuardAssociation")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("SwitchoverDataGuardAssociation is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("database", "Database", "SwitchoverDataGuardAssociation", createDatabaseClientWithProvider)
	assert.NoError(t, err)
	c := cc.(database.DatabaseClient)

	body, err := testClient.getRequests("database", "SwitchoverDataGuardAssociation")
	assert.NoError(t, err)

	type SwitchoverDataGuardAssociationRequestInfo struct {
		ContainerId string
		Request     database.SwitchoverDataGuardAssociationRequest
	}

	var requests []SwitchoverDataGuardAssociationRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.SwitchoverDataGuardAssociation(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="sic_dbaas_cp_us_grp@oracle.com" jiraProject="DBAAS" opsJiraProject="DBAASOPS"
func TestDatabaseClientTerminateDbSystem(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("database", "TerminateDbSystem")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("TerminateDbSystem is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("database", "Database", "TerminateDbSystem", createDatabaseClientWithProvider)
	assert.NoError(t, err)
	c := cc.(database.DatabaseClient)

	body, err := testClient.getRequests("database", "TerminateDbSystem")
	assert.NoError(t, err)

	type TerminateDbSystemRequestInfo struct {
		ContainerId string
		Request     database.TerminateDbSystemRequest
	}

	var requests []TerminateDbSystemRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.TerminateDbSystem(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="sic_dbaas_cp_us_grp@oracle.com" jiraProject="DBAAS" opsJiraProject="DBAASOPS"
func TestDatabaseClientUpdateAutonomousDataWarehouse(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("database", "UpdateAutonomousDataWarehouse")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("UpdateAutonomousDataWarehouse is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("database", "Database", "UpdateAutonomousDataWarehouse", createDatabaseClientWithProvider)
	assert.NoError(t, err)
	c := cc.(database.DatabaseClient)

	body, err := testClient.getRequests("database", "UpdateAutonomousDataWarehouse")
	assert.NoError(t, err)

	type UpdateAutonomousDataWarehouseRequestInfo struct {
		ContainerId string
		Request     database.UpdateAutonomousDataWarehouseRequest
	}

	var requests []UpdateAutonomousDataWarehouseRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.UpdateAutonomousDataWarehouse(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="dbaas-atp-d" email="sic_dbaas_cp_us_grp@oracle.com" jiraProject="DBAAS" opsJiraProject="DBAASOPS"
func TestDatabaseClientUpdateAutonomousDatabase(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("database", "UpdateAutonomousDatabase")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("UpdateAutonomousDatabase is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("database", "Database", "UpdateAutonomousDatabase", createDatabaseClientWithProvider)
	assert.NoError(t, err)
	c := cc.(database.DatabaseClient)

	body, err := testClient.getRequests("database", "UpdateAutonomousDatabase")
	assert.NoError(t, err)

	type UpdateAutonomousDatabaseRequestInfo struct {
		ContainerId string
		Request     database.UpdateAutonomousDatabaseRequest
	}

	var requests []UpdateAutonomousDatabaseRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.UpdateAutonomousDatabase(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="sic_dbaas_cp_us_grp@oracle.com" jiraProject="DBAAS" opsJiraProject="DBAASOPS"
func TestDatabaseClientUpdateDatabase(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("database", "UpdateDatabase")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("UpdateDatabase is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("database", "Database", "UpdateDatabase", createDatabaseClientWithProvider)
	assert.NoError(t, err)
	c := cc.(database.DatabaseClient)

	body, err := testClient.getRequests("database", "UpdateDatabase")
	assert.NoError(t, err)

	type UpdateDatabaseRequestInfo struct {
		ContainerId string
		Request     database.UpdateDatabaseRequest
	}

	var requests []UpdateDatabaseRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.UpdateDatabase(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="sic_dbaas_cp_us_grp@oracle.com" jiraProject="DBAAS" opsJiraProject="DBAASOPS"
func TestDatabaseClientUpdateDbHome(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("database", "UpdateDbHome")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("UpdateDbHome is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("database", "Database", "UpdateDbHome", createDatabaseClientWithProvider)
	assert.NoError(t, err)
	c := cc.(database.DatabaseClient)

	body, err := testClient.getRequests("database", "UpdateDbHome")
	assert.NoError(t, err)

	type UpdateDbHomeRequestInfo struct {
		ContainerId string
		Request     database.UpdateDbHomeRequest
	}

	var requests []UpdateDbHomeRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.UpdateDbHome(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="sic_dbaas_cp_us_grp@oracle.com" jiraProject="DBAAS" opsJiraProject="DBAASOPS"
func TestDatabaseClientUpdateDbSystem(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("database", "UpdateDbSystem")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("UpdateDbSystem is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("database", "Database", "UpdateDbSystem", createDatabaseClientWithProvider)
	assert.NoError(t, err)
	c := cc.(database.DatabaseClient)

	body, err := testClient.getRequests("database", "UpdateDbSystem")
	assert.NoError(t, err)

	type UpdateDbSystemRequestInfo struct {
		ContainerId string
		Request     database.UpdateDbSystemRequest
	}

	var requests []UpdateDbSystemRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.UpdateDbSystem(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="sic_dbaas_cp_us_grp@oracle.com" jiraProject="DBAAS" opsJiraProject="DBAASOPS"
func TestDatabaseClientUpdateExadataIormConfig(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("database", "UpdateExadataIormConfig")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("UpdateExadataIormConfig is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("database", "Database", "UpdateExadataIormConfig", createDatabaseClientWithProvider)
	assert.NoError(t, err)
	c := cc.(database.DatabaseClient)

	body, err := testClient.getRequests("database", "UpdateExadataIormConfig")
	assert.NoError(t, err)

	type UpdateExadataIormConfigRequestInfo struct {
		ContainerId string
		Request     database.UpdateExadataIormConfigRequest
	}

	var requests []UpdateExadataIormConfigRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.UpdateExadataIormConfig(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}
