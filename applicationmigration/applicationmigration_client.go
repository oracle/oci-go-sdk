// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Application Migration API
//
// Application Migration simplifies the migration of applications from Oracle Cloud Infrastructure Classic to Oracle Cloud Infrastructure.
// You can use Application Migration API to migrate applications, such as Oracle Java Cloud Service, SOA Cloud Service, and Integration Classic
// instances, to Oracle Cloud Infrastructure. For more information, see
// Overview of Application Migration (https://docs.cloud.oracle.com/iaas/application-migration/appmigrationoverview.htm).
//

package applicationmigration

import (
	"context"
	"fmt"
	"github.com/oracle/oci-go-sdk/v56/common"
	"github.com/oracle/oci-go-sdk/v56/common/auth"
	"net/http"
)

//ApplicationMigrationClient a client for ApplicationMigration
type ApplicationMigrationClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewApplicationMigrationClientWithConfigurationProvider Creates a new default ApplicationMigration client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewApplicationMigrationClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client ApplicationMigrationClient, err error) {
	provider, err := auth.GetGenericConfigurationProvider(configProvider)
	if err != nil {
		return client, err
	}
	baseClient, e := common.NewClientWithConfig(provider)
	if e != nil {
		return client, e
	}
	return newApplicationMigrationClientFromBaseClient(baseClient, provider)
}

// NewApplicationMigrationClientWithOboToken Creates a new default ApplicationMigration client with the given configuration provider.
// The obotoken will be added to default headers and signed; the configuration provider will be used for the signer
//  as well as reading the region
func NewApplicationMigrationClientWithOboToken(configProvider common.ConfigurationProvider, oboToken string) (client ApplicationMigrationClient, err error) {
	baseClient, err := common.NewClientWithOboToken(configProvider, oboToken)
	if err != nil {
		return client, err
	}

	return newApplicationMigrationClientFromBaseClient(baseClient, configProvider)
}

func newApplicationMigrationClientFromBaseClient(baseClient common.BaseClient, configProvider common.ConfigurationProvider) (client ApplicationMigrationClient, err error) {
	// ApplicationMigration service default circuit breaker is enabled
	baseClient.Configuration.CircuitBreaker = common.NewCircuitBreaker(common.DefaultCircuitBreakerSetting())
	common.ConfigCircuitBreakerFromEnvVar(&baseClient)
	common.ConfigCircuitBreakerFromGlobalVar(&baseClient)

	client = ApplicationMigrationClient{BaseClient: baseClient}
	client.BasePath = "20191031"
	err = client.setConfigurationProvider(configProvider)
	return
}

// SetRegion overrides the region of this client.
func (client *ApplicationMigrationClient) SetRegion(region string) {
	client.Host = common.StringToRegion(region).EndpointForTemplate("applicationmigration", "https://applicationmigration.{region}.oci.{secondLevelDomain}")
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *ApplicationMigrationClient) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
	if ok, err := common.IsConfigurationProviderValid(configProvider); !ok {
		return err
	}

	// Error has been checked already
	region, _ := configProvider.Region()
	client.SetRegion(region)
	client.config = &configProvider
	return nil
}

// ConfigurationProvider the ConfigurationProvider used in this client, or null if none set
func (client *ApplicationMigrationClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// CancelWorkRequest Cancels the specified work request. When you cancel a work request, it causes the in-progress task to be canceled.
// For example, if the create migration work request is in the accepted or in progress state for a long time, you can cancel the work request.
// When you cancel a work request, the state of the work request changes to cancelling, and then to the cancelled state.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/applicationmigration/CancelWorkRequest.go.html to see an example of how to use CancelWorkRequest API.
func (client ApplicationMigrationClient) CancelWorkRequest(ctx context.Context, request CancelWorkRequestRequest) (response CancelWorkRequestResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.cancelWorkRequest, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CancelWorkRequestResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CancelWorkRequestResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CancelWorkRequestResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CancelWorkRequestResponse")
	}
	return
}

// cancelWorkRequest implements the OCIOperation interface (enables retrying operations)
func (client ApplicationMigrationClient) cancelWorkRequest(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/workRequests/{workRequestId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CancelWorkRequestResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ChangeMigrationCompartment Moves the specified migration into a different compartment within the same tenancy. For information about moving resources between compartments,
// see Moving Resources to a Different Compartment (https://docs.cloud.oracle.com/iaas/Content/Identity/Tasks/managingcompartments.htm#moveRes).
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/applicationmigration/ChangeMigrationCompartment.go.html to see an example of how to use ChangeMigrationCompartment API.
func (client ApplicationMigrationClient) ChangeMigrationCompartment(ctx context.Context, request ChangeMigrationCompartmentRequest) (response ChangeMigrationCompartmentResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.changeMigrationCompartment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ChangeMigrationCompartmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ChangeMigrationCompartmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ChangeMigrationCompartmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ChangeMigrationCompartmentResponse")
	}
	return
}

// changeMigrationCompartment implements the OCIOperation interface (enables retrying operations)
func (client ApplicationMigrationClient) changeMigrationCompartment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/migrations/{migrationId}/actions/changeCompartment", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ChangeMigrationCompartmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ChangeSourceCompartment Moves the specified source into a different compartment within the same tenancy. For information about moving resources
// between compartments, see Moving Resources to a Different Compartment (https://docs.cloud.oracle.com/iaas/Content/Identity/Tasks/managingcompartments.htm#moveRes).
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/applicationmigration/ChangeSourceCompartment.go.html to see an example of how to use ChangeSourceCompartment API.
func (client ApplicationMigrationClient) ChangeSourceCompartment(ctx context.Context, request ChangeSourceCompartmentRequest) (response ChangeSourceCompartmentResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.changeSourceCompartment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ChangeSourceCompartmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ChangeSourceCompartmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ChangeSourceCompartmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ChangeSourceCompartmentResponse")
	}
	return
}

// changeSourceCompartment implements the OCIOperation interface (enables retrying operations)
func (client ApplicationMigrationClient) changeSourceCompartment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/sources/{sourceId}/actions/changeCompartment", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ChangeSourceCompartmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateMigration Creates a migration. A migration represents the end-to-end workflow of moving an application from a source environment to Oracle Cloud
// Infrastructure. Each migration moves a single application to Oracle Cloud Infrastructure. For more information,
// see Manage Migrations (https://docs.cloud.oracle.com/iaas/application-migration/manage_migrations.htm).
// When you create a migration, provide the required information to let Application Migration access the source environment.
// Application Migration uses this information to access the application in the source environment and discover application artifacts.
// All Oracle Cloud Infrastructure resources, including migrations, get an Oracle-assigned, unique ID called an Oracle Cloud Identifier (OCID).
// When you create a resource, you can find its OCID in the response. You can also retrieve a resource's OCID by using a List API operation on
// that resource type, or by viewing the resource in the Console. For more information, see Resource Identifiers.
// After you send your request, a migration is created in the compartment that contains the source. The new migration's lifecycle state
// will temporarily be <code>CREATING</code> and the state of the migration will be <code>DISCOVERING_APPLICATION</code>. During this phase,
// Application Migration sets the template for the <code>serviceConfig</code> and <code>applicationConfig</code> fields of the migration.
// When this operation is complete, the state of the migration changes to <code>MISSING_CONFIG_VALUES</code>.
// Next, you'll need to update the migration to provide configuration values. Before updating the
// migration, ensure that its state has changed to <code>MISSING_CONFIG_VALUES</code>.
// To track the progress of this operation, you can monitor the status of the Create Migration and Discover Application work requests
// by using the <code>GetWorkRequest</code> REST API operation on the work request or by viewing the status of the work request in
// the console.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/applicationmigration/CreateMigration.go.html to see an example of how to use CreateMigration API.
func (client ApplicationMigrationClient) CreateMigration(ctx context.Context, request CreateMigrationRequest) (response CreateMigrationResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.createMigration, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateMigrationResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateMigrationResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateMigrationResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateMigrationResponse")
	}
	return
}

// createMigration implements the OCIOperation interface (enables retrying operations)
func (client ApplicationMigrationClient) createMigration(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/migrations", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateMigrationResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateSource Creates a source in the specified compartment. In Application Migration, a source refers to the environment from which the application
// is being migrated. For more information, see Manage Sources (https://docs.cloud.oracle.com/iaas/application-migration/manage_sources.htm).
// All Oracle Cloud Infrastructure resources, including sources, get an Oracle-assigned, unique ID called an Oracle Cloud Identifier (OCID).
// When you create a resource, you can find its OCID in the response. You can also retrieve a resource's OCID by using a List API operation
// on that resource type, or by viewing the resource in the Console.
// After you send your request, a source is created in the specified compartment. The new source's lifecycle state will temporarily be
// <code>CREATING</code>. Application Migration connects to the source environment with the authentication credentials that you have provided.
// If the connection is established, the status of the source changes to <code>ACTIVE</code> and Application Migration fetches the list of
// applications that are available for migration in the source environment.
// To track the progress of the operation, you can monitor the status of the Create Source work request by using the
// <code>GetWorkRequest</code> REST API operation on the work request or by viewing the status of the work request in the console.
// Ensure that the state of the source has changed to <code>ACTIVE</code>, before you retrieve the list of applications from
// the source environment using the <code>ListSourceApplications</code> REST API call.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/applicationmigration/CreateSource.go.html to see an example of how to use CreateSource API.
func (client ApplicationMigrationClient) CreateSource(ctx context.Context, request CreateSourceRequest) (response CreateSourceResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.createSource, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateSourceResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateSourceResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateSourceResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateSourceResponse")
	}
	return
}

// createSource implements the OCIOperation interface (enables retrying operations)
func (client ApplicationMigrationClient) createSource(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/sources", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateSourceResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteMigration Deletes the specified migration.
// If you have migrated the application or for any other reason if you no longer require a migration, then you can delete the
// relevant migration. You can delete a migration, irrespective of its state. If any work request is being processed for the migration
// that you want to delete, then the associated work requests are cancelled and then the migration is deleted.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/applicationmigration/DeleteMigration.go.html to see an example of how to use DeleteMigration API.
func (client ApplicationMigrationClient) DeleteMigration(ctx context.Context, request DeleteMigrationRequest) (response DeleteMigrationResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteMigration, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteMigrationResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteMigrationResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteMigrationResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteMigrationResponse")
	}
	return
}

// deleteMigration implements the OCIOperation interface (enables retrying operations)
func (client ApplicationMigrationClient) deleteMigration(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/migrations/{migrationId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteMigrationResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteSource Deletes the specified source.
// Before deleting a source, you must delete all the migrations associated with the source.
// If you have migrated all the required applications in a source or for any other reason you no longer require a source, then you can
// delete the relevant source.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/applicationmigration/DeleteSource.go.html to see an example of how to use DeleteSource API.
func (client ApplicationMigrationClient) DeleteSource(ctx context.Context, request DeleteSourceRequest) (response DeleteSourceResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteSource, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteSourceResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteSourceResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteSourceResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteSourceResponse")
	}
	return
}

// deleteSource implements the OCIOperation interface (enables retrying operations)
func (client ApplicationMigrationClient) deleteSource(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/sources/{sourceId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteSourceResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetMigration Retrieves details of the specified migration.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/applicationmigration/GetMigration.go.html to see an example of how to use GetMigration API.
func (client ApplicationMigrationClient) GetMigration(ctx context.Context, request GetMigrationRequest) (response GetMigrationResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getMigration, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetMigrationResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetMigrationResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetMigrationResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetMigrationResponse")
	}
	return
}

// getMigration implements the OCIOperation interface (enables retrying operations)
func (client ApplicationMigrationClient) getMigration(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/migrations/{migrationId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetMigrationResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetSource Retrieves details of the specified source. Specify the OCID of the source for which you want to retrieve details.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/applicationmigration/GetSource.go.html to see an example of how to use GetSource API.
func (client ApplicationMigrationClient) GetSource(ctx context.Context, request GetSourceRequest) (response GetSourceResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getSource, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetSourceResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetSourceResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetSourceResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetSourceResponse")
	}
	return
}

// getSource implements the OCIOperation interface (enables retrying operations)
func (client ApplicationMigrationClient) getSource(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/sources/{sourceId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetSourceResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetWorkRequest Gets the details of the specified work request.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/applicationmigration/GetWorkRequest.go.html to see an example of how to use GetWorkRequest API.
func (client ApplicationMigrationClient) GetWorkRequest(ctx context.Context, request GetWorkRequestRequest) (response GetWorkRequestResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getWorkRequest, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetWorkRequestResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetWorkRequestResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetWorkRequestResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetWorkRequestResponse")
	}
	return
}

// getWorkRequest implements the OCIOperation interface (enables retrying operations)
func (client ApplicationMigrationClient) getWorkRequest(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/workRequests/{workRequestId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetWorkRequestResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListMigrations Retrieves details of all the migrations that are available in the specified compartment.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/applicationmigration/ListMigrations.go.html to see an example of how to use ListMigrations API.
func (client ApplicationMigrationClient) ListMigrations(ctx context.Context, request ListMigrationsRequest) (response ListMigrationsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listMigrations, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListMigrationsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListMigrationsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListMigrationsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListMigrationsResponse")
	}
	return
}

// listMigrations implements the OCIOperation interface (enables retrying operations)
func (client ApplicationMigrationClient) listMigrations(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/migrations", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListMigrationsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListSourceApplications Retrieves details of all the applications associated with the specified source.
// This list is generated dynamically by interrogating the source and the list changes as applications are started or
// stopped in the source environment.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/applicationmigration/ListSourceApplications.go.html to see an example of how to use ListSourceApplications API.
func (client ApplicationMigrationClient) ListSourceApplications(ctx context.Context, request ListSourceApplicationsRequest) (response ListSourceApplicationsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listSourceApplications, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListSourceApplicationsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListSourceApplicationsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListSourceApplicationsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListSourceApplicationsResponse")
	}
	return
}

// listSourceApplications implements the OCIOperation interface (enables retrying operations)
func (client ApplicationMigrationClient) listSourceApplications(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/sources/{sourceId}/applications", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListSourceApplicationsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListSources Retrieves details of all the sources that are available in the specified compartment and match the specified query criteria.
// If you don't specify any query criteria, then details of all the sources are displayed.
// To filter the retrieved results, you can pass one or more of the following query parameters, by appending them to the URI
// as shown in the following example.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/applicationmigration/ListSources.go.html to see an example of how to use ListSources API.
func (client ApplicationMigrationClient) ListSources(ctx context.Context, request ListSourcesRequest) (response ListSourcesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listSources, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListSourcesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListSourcesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListSourcesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListSourcesResponse")
	}
	return
}

// listSources implements the OCIOperation interface (enables retrying operations)
func (client ApplicationMigrationClient) listSources(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/sources", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListSourcesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListWorkRequestErrors Retrieves details of the errors encountered while executing an operation that is tracked by the specified work request.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/applicationmigration/ListWorkRequestErrors.go.html to see an example of how to use ListWorkRequestErrors API.
func (client ApplicationMigrationClient) ListWorkRequestErrors(ctx context.Context, request ListWorkRequestErrorsRequest) (response ListWorkRequestErrorsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listWorkRequestErrors, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListWorkRequestErrorsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListWorkRequestErrorsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListWorkRequestErrorsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListWorkRequestErrorsResponse")
	}
	return
}

// listWorkRequestErrors implements the OCIOperation interface (enables retrying operations)
func (client ApplicationMigrationClient) listWorkRequestErrors(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/workRequests/{workRequestId}/errors", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListWorkRequestErrorsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListWorkRequestLogs Retrieves logs for the specified work request.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/applicationmigration/ListWorkRequestLogs.go.html to see an example of how to use ListWorkRequestLogs API.
func (client ApplicationMigrationClient) ListWorkRequestLogs(ctx context.Context, request ListWorkRequestLogsRequest) (response ListWorkRequestLogsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listWorkRequestLogs, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListWorkRequestLogsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListWorkRequestLogsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListWorkRequestLogsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListWorkRequestLogsResponse")
	}
	return
}

// listWorkRequestLogs implements the OCIOperation interface (enables retrying operations)
func (client ApplicationMigrationClient) listWorkRequestLogs(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/workRequests/{workRequestId}/logs", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListWorkRequestLogsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListWorkRequests Retrieves details of all the work requests and match the specified query criteria. To filter the retrieved results, you can pass one or more of the following query parameters, by appending them to the URI as shown in the following example.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/applicationmigration/ListWorkRequests.go.html to see an example of how to use ListWorkRequests API.
func (client ApplicationMigrationClient) ListWorkRequests(ctx context.Context, request ListWorkRequestsRequest) (response ListWorkRequestsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listWorkRequests, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListWorkRequestsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListWorkRequestsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListWorkRequestsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListWorkRequestsResponse")
	}
	return
}

// listWorkRequests implements the OCIOperation interface (enables retrying operations)
func (client ApplicationMigrationClient) listWorkRequests(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/workRequests", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListWorkRequestsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// MigrateApplication Starts migrating the specified application to Oracle Cloud Infrastructure.
// Before sending this request, ensure that you have provided configuration details to update the migration and the state of the migration
// is <code>READY</code>.
// After you send this request, the migration's state will temporarily be <code>MIGRATING</code>.
// To track the progress of the operation, you can monitor the status of the Migrate Application work request by using the
// <code>GetWorkRequest</code> REST API operation on the work request or by viewing the status of the work request in the console.
// When this work request is processed successfully, Application Migration creates the required resources in the target environment
// and the state of the migration changes to <code>MIGRATION_SUCCEEDED</code>.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/applicationmigration/MigrateApplication.go.html to see an example of how to use MigrateApplication API.
func (client ApplicationMigrationClient) MigrateApplication(ctx context.Context, request MigrateApplicationRequest) (response MigrateApplicationResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.migrateApplication, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = MigrateApplicationResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = MigrateApplicationResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(MigrateApplicationResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into MigrateApplicationResponse")
	}
	return
}

// migrateApplication implements the OCIOperation interface (enables retrying operations)
func (client ApplicationMigrationClient) migrateApplication(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/migrations/{migrationId}/actions/migrate", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response MigrateApplicationResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateMigration Updates the configuration details for the specified migration.
// When you create a migration, Application Migration sets the template for the <code>serviceConfig</code> and <code>applicationConfig</code>
// attributes of the migration.
// When you update the migration, you must provide values for these fields to specify configuration information for the application in the
// target environment.
//
// Before updating the migration, complete the following tasks:
// <ol>
// <li>Identify the migration that you want to update and ensure that the migration is in the <code>MISSING_CONFIG_VALUES</code> state.</li>
// <li>Get details of the migration using the <code>GetMigration</code> command. This returns the  template for the <code>serviceConfig</code>
// and <code>applicationConfig</code> attributes of the migration.</li>
// <li>You must fill out the required details for the <code>serviceConfig</code> and <code>applicationConfig</code> attributes.
// The <code>isRequired</code> attribute of a configuration property indicates whether it is mandatory to provide a value.</li>
// <li>You can provide values for the optional configuration properties or you can delete the optional properties for which you do not
// provide values. Note that you cannot add any property that is not present in the template.</li>
// </ol>
// To update the migration, pass the configuration values in the request body. The information that you must provide depends on the type
// of application that you are migrating. For reference information about configuration fields, see
// Provide Configuration Information (https://docs.cloud.oracle.com/iaas/application-migration/manage_migrations.htm#provide_configuration_details).
// To track the progress of the operation, you can monitor the status of the Update Migration work request by using the
// <code>GetWorkRequest</code> REST API operation on the work request or by viewing the status of the work request in the console.
// When the migration has been updated, the state of the migration changes to <code>READY</code>. After updating the migration,
// you can start the migration whenever you are ready.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/applicationmigration/UpdateMigration.go.html to see an example of how to use UpdateMigration API.
func (client ApplicationMigrationClient) UpdateMigration(ctx context.Context, request UpdateMigrationRequest) (response UpdateMigrationResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.updateMigration, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateMigrationResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateMigrationResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateMigrationResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateMigrationResponse")
	}
	return
}

// updateMigration implements the OCIOperation interface (enables retrying operations)
func (client ApplicationMigrationClient) updateMigration(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/migrations/{migrationId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateMigrationResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateSource You can update the authorization details to access the source environment from which you want to migrate applications to Oracle Cloud
// Infrastructure. You can also update the description and tags of a source.
// **Warning:** Oracle recommends that you avoid using any confidential information when you supply string values using the API.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/applicationmigration/UpdateSource.go.html to see an example of how to use UpdateSource API.
func (client ApplicationMigrationClient) UpdateSource(ctx context.Context, request UpdateSourceRequest) (response UpdateSourceResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateSource, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateSourceResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateSourceResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateSourceResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateSourceResponse")
	}
	return
}

// updateSource implements the OCIOperation interface (enables retrying operations)
func (client ApplicationMigrationClient) updateSource(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/sources/{sourceId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateSourceResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}
