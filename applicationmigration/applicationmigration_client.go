// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Application Migration Service API
//
// API for the Application Migration service. Use this API to migrate applications from Oracle Cloud Infrastructure - Classic to Oracle Cloud Infrastructure.
//

package applicationmigration

import (
	"context"
	"fmt"
	"github.com/oracle/oci-go-sdk/common"
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
	baseClient, err := common.NewClientWithConfig(configProvider)
	if err != nil {
		return
	}

	return newApplicationMigrationClientFromBaseClient(baseClient, configProvider)
}

// NewApplicationMigrationClientWithOboToken Creates a new default ApplicationMigration client with the given configuration provider.
// The obotoken will be added to default headers and signed; the configuration provider will be used for the signer
//  as well as reading the region
func NewApplicationMigrationClientWithOboToken(configProvider common.ConfigurationProvider, oboToken string) (client ApplicationMigrationClient, err error) {
	baseClient, err := common.NewClientWithOboToken(configProvider, oboToken)
	if err != nil {
		return
	}

	return newApplicationMigrationClientFromBaseClient(baseClient, configProvider)
}

func newApplicationMigrationClientFromBaseClient(baseClient common.BaseClient, configProvider common.ConfigurationProvider) (client ApplicationMigrationClient, err error) {
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

// CancelWorkRequest Cancels the specified work request
func (client ApplicationMigrationClient) CancelWorkRequest(ctx context.Context, request CancelWorkRequestRequest) (response CancelWorkRequestResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
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
func (client ApplicationMigrationClient) cancelWorkRequest(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/workRequests/{workRequestId}")
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

// ChangeMigrationCompartment Moves a Migration into a different compartment.
func (client ApplicationMigrationClient) ChangeMigrationCompartment(ctx context.Context, request ChangeMigrationCompartmentRequest) (response ChangeMigrationCompartmentResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
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
func (client ApplicationMigrationClient) changeMigrationCompartment(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/migrations/{migrationId}/actions/changeCompartment")
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

// ChangeSourceCompartment Moves a Source into a different compartment.
func (client ApplicationMigrationClient) ChangeSourceCompartment(ctx context.Context, request ChangeSourceCompartmentRequest) (response ChangeSourceCompartmentResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
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
func (client ApplicationMigrationClient) changeSourceCompartment(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/sources/{sourceId}/actions/changeCompartment")
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

// CreateMigration Creates an application migration in the specified compartment.
// Specify the compartment using the compartment ID.
func (client ApplicationMigrationClient) CreateMigration(ctx context.Context, request CreateMigrationRequest) (response CreateMigrationResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
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
func (client ApplicationMigrationClient) createMigration(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/migrations")
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

// CreateSource Creates a migration source in the specified compartment.
// Specify the compartment using the compartment ID.
func (client ApplicationMigrationClient) CreateSource(ctx context.Context, request CreateSourceRequest) (response CreateSourceResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
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
func (client ApplicationMigrationClient) createSource(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/sources")
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

// DeleteMigration Deletes the specified Application object.
func (client ApplicationMigrationClient) DeleteMigration(ctx context.Context, request DeleteMigrationRequest) (response DeleteMigrationResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
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
func (client ApplicationMigrationClient) deleteMigration(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/migrations/{migrationId}")
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

// DeleteSource Deletes the specified Source object.
func (client ApplicationMigrationClient) DeleteSource(ctx context.Context, request DeleteSourceRequest) (response DeleteSourceResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
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
func (client ApplicationMigrationClient) deleteSource(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/sources/{sourceId}")
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

// GetMigration Gets an application migration using the ID.
func (client ApplicationMigrationClient) GetMigration(ctx context.Context, request GetMigrationRequest) (response GetMigrationResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
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
func (client ApplicationMigrationClient) getMigration(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/migrations/{migrationId}")
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

// GetSource Gets a migration source using the source ID.
func (client ApplicationMigrationClient) GetSource(ctx context.Context, request GetSourceRequest) (response GetSourceResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
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
func (client ApplicationMigrationClient) getSource(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/sources/{sourceId}")
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

// GetWorkRequest Gets the details of a work request.
func (client ApplicationMigrationClient) GetWorkRequest(ctx context.Context, request GetWorkRequestRequest) (response GetWorkRequestResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
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
func (client ApplicationMigrationClient) getWorkRequest(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/workRequests/{workRequestId}")
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

// ListMigrations Returns a list of migrations in a given compartment.
func (client ApplicationMigrationClient) ListMigrations(ctx context.Context, request ListMigrationsRequest) (response ListMigrationsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
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
func (client ApplicationMigrationClient) listMigrations(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/migrations")
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

// ListSourceApplications Returns a list of applications running in the source environment. This list is generated dynamically by interrogating the source and changes as applications are started or stopped in that environment.
func (client ApplicationMigrationClient) ListSourceApplications(ctx context.Context, request ListSourceApplicationsRequest) (response ListSourceApplicationsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
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
func (client ApplicationMigrationClient) listSourceApplications(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/sources/{sourceId}/applications")
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

// ListSources Returns a list of migration sources in a specified compartment.
func (client ApplicationMigrationClient) ListSources(ctx context.Context, request ListSourcesRequest) (response ListSourcesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
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
func (client ApplicationMigrationClient) listSources(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/sources")
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

// ListWorkRequestErrors Gets the errors for a work request.
func (client ApplicationMigrationClient) ListWorkRequestErrors(ctx context.Context, request ListWorkRequestErrorsRequest) (response ListWorkRequestErrorsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
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
func (client ApplicationMigrationClient) listWorkRequestErrors(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/workRequests/{workRequestId}/errors")
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

// ListWorkRequestLogs Gets the logs for a work request.
func (client ApplicationMigrationClient) ListWorkRequestLogs(ctx context.Context, request ListWorkRequestLogsRequest) (response ListWorkRequestLogsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
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
func (client ApplicationMigrationClient) listWorkRequestLogs(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/workRequests/{workRequestId}/logs")
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

// ListWorkRequests Lists the work requests in a compartment or for a specified resource.
func (client ApplicationMigrationClient) ListWorkRequests(ctx context.Context, request ListWorkRequestsRequest) (response ListWorkRequestsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
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
func (client ApplicationMigrationClient) listWorkRequests(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/workRequests")
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

// MigrateApplication Validates target configuration and migrates a PaaS application running in a Source environment into the customers Oracle Cloud Infrastructure tenancy. This an optional action and only required if automatic start of migration was not selected when creating the migration.
func (client ApplicationMigrationClient) MigrateApplication(ctx context.Context, request MigrateApplicationRequest) (response MigrateApplicationResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
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
func (client ApplicationMigrationClient) migrateApplication(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/migrations/{migrationId}/actions/migrate")
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

// UpdateMigration Update the configuration for an application migration.
func (client ApplicationMigrationClient) UpdateMigration(ctx context.Context, request UpdateMigrationRequest) (response UpdateMigrationResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
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
func (client ApplicationMigrationClient) updateMigration(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPut, "/migrations/{migrationId}")
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

// UpdateSource Update source details.
func (client ApplicationMigrationClient) UpdateSource(ctx context.Context, request UpdateSourceRequest) (response UpdateSourceResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
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
func (client ApplicationMigrationClient) updateSource(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPut, "/sources/{sourceId}")
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
