// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Transfer Service API
//
// Data Transfer Service API Specification
//

package dts

import (
	"context"
	"fmt"
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

//TransferApplianceClient a client for TransferAppliance
type TransferApplianceClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewTransferApplianceClientWithConfigurationProvider Creates a new default TransferAppliance client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewTransferApplianceClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client TransferApplianceClient, err error) {
	baseClient, err := common.NewClientWithConfig(configProvider)
	if err != nil {
		return
	}

	return newTransferApplianceClientFromBaseClient(baseClient, configProvider)
}

// NewTransferApplianceClientWithOboToken Creates a new default TransferAppliance client with the given configuration provider.
// The obotoken will be added to default headers and signed; the configuration provider will be used for the signer
//  as well as reading the region
func NewTransferApplianceClientWithOboToken(configProvider common.ConfigurationProvider, oboToken string) (client TransferApplianceClient, err error) {
	baseClient, err := common.NewClientWithOboToken(configProvider, oboToken)
	if err != nil {
		return
	}

	return newTransferApplianceClientFromBaseClient(baseClient, configProvider)
}

func newTransferApplianceClientFromBaseClient(baseClient common.BaseClient, configProvider common.ConfigurationProvider) (client TransferApplianceClient, err error) {
	client = TransferApplianceClient{BaseClient: baseClient}
	client.BasePath = "20171001"
	err = client.setConfigurationProvider(configProvider)
	return
}

// SetRegion overrides the region of this client.
func (client *TransferApplianceClient) SetRegion(region string) {
	client.Host = common.StringToRegion(region).EndpointForTemplate("dts", "https://datatransfer.{region}.oci.{secondLevelDomain}")
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *TransferApplianceClient) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
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
func (client *TransferApplianceClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// CreateTransferAppliance Create a new Transfer Appliance
func (client TransferApplianceClient) CreateTransferAppliance(ctx context.Context, request CreateTransferApplianceRequest) (response CreateTransferApplianceResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.createTransferAppliance, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateTransferApplianceResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateTransferApplianceResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateTransferApplianceResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateTransferApplianceResponse")
	}
	return
}

// createTransferAppliance implements the OCIOperation interface (enables retrying operations)
func (client TransferApplianceClient) createTransferAppliance(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/transferJobs/{id}/transferAppliances")
	if err != nil {
		return nil, err
	}

	var response CreateTransferApplianceResponse
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

// CreateTransferApplianceAdminCredentials Creates an X.509 certificate from a public key
func (client TransferApplianceClient) CreateTransferApplianceAdminCredentials(ctx context.Context, request CreateTransferApplianceAdminCredentialsRequest) (response CreateTransferApplianceAdminCredentialsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.createTransferApplianceAdminCredentials, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateTransferApplianceAdminCredentialsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateTransferApplianceAdminCredentialsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateTransferApplianceAdminCredentialsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateTransferApplianceAdminCredentialsResponse")
	}
	return
}

// createTransferApplianceAdminCredentials implements the OCIOperation interface (enables retrying operations)
func (client TransferApplianceClient) createTransferApplianceAdminCredentials(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/transferJobs/{id}/transferAppliances/{transferApplianceLabel}/admin_credentials")
	if err != nil {
		return nil, err
	}

	var response CreateTransferApplianceAdminCredentialsResponse
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

// DeleteTransferAppliance deletes a transfer Appliance
func (client TransferApplianceClient) DeleteTransferAppliance(ctx context.Context, request DeleteTransferApplianceRequest) (response DeleteTransferApplianceResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteTransferAppliance, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteTransferApplianceResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteTransferApplianceResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteTransferApplianceResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteTransferApplianceResponse")
	}
	return
}

// deleteTransferAppliance implements the OCIOperation interface (enables retrying operations)
func (client TransferApplianceClient) deleteTransferAppliance(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/transferJobs/{id}/transferAppliances/{transferApplianceLabel}")
	if err != nil {
		return nil, err
	}

	var response DeleteTransferApplianceResponse
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

// GetTransferAppliance Describes a transfer appliance in detail
func (client TransferApplianceClient) GetTransferAppliance(ctx context.Context, request GetTransferApplianceRequest) (response GetTransferApplianceResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getTransferAppliance, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetTransferApplianceResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetTransferApplianceResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetTransferApplianceResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetTransferApplianceResponse")
	}
	return
}

// getTransferAppliance implements the OCIOperation interface (enables retrying operations)
func (client TransferApplianceClient) getTransferAppliance(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/transferJobs/{id}/transferAppliances/{transferApplianceLabel}")
	if err != nil {
		return nil, err
	}

	var response GetTransferApplianceResponse
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

// GetTransferApplianceCertificateAuthorityCertificate Gets the x.509 certificate for the Transfer Appliance's dedicated Certificate Authority (CA)
func (client TransferApplianceClient) GetTransferApplianceCertificateAuthorityCertificate(ctx context.Context, request GetTransferApplianceCertificateAuthorityCertificateRequest) (response GetTransferApplianceCertificateAuthorityCertificateResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getTransferApplianceCertificateAuthorityCertificate, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetTransferApplianceCertificateAuthorityCertificateResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetTransferApplianceCertificateAuthorityCertificateResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetTransferApplianceCertificateAuthorityCertificateResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetTransferApplianceCertificateAuthorityCertificateResponse")
	}
	return
}

// getTransferApplianceCertificateAuthorityCertificate implements the OCIOperation interface (enables retrying operations)
func (client TransferApplianceClient) getTransferApplianceCertificateAuthorityCertificate(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/transferJobs/{id}/transferAppliances/{transferApplianceLabel}/certificate_authority_certificate")
	if err != nil {
		return nil, err
	}

	var response GetTransferApplianceCertificateAuthorityCertificateResponse
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

// GetTransferApplianceEncryptionPassphrase Describes a transfer appliance encryptionPassphrase in detail
func (client TransferApplianceClient) GetTransferApplianceEncryptionPassphrase(ctx context.Context, request GetTransferApplianceEncryptionPassphraseRequest) (response GetTransferApplianceEncryptionPassphraseResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getTransferApplianceEncryptionPassphrase, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetTransferApplianceEncryptionPassphraseResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetTransferApplianceEncryptionPassphraseResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetTransferApplianceEncryptionPassphraseResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetTransferApplianceEncryptionPassphraseResponse")
	}
	return
}

// getTransferApplianceEncryptionPassphrase implements the OCIOperation interface (enables retrying operations)
func (client TransferApplianceClient) getTransferApplianceEncryptionPassphrase(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/transferJobs/{id}/transferAppliances/{transferApplianceLabel}/encryptionPassphrase")
	if err != nil {
		return nil, err
	}

	var response GetTransferApplianceEncryptionPassphraseResponse
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

// ListTransferAppliances Lists Transfer Appliances associated with a transferJob
func (client TransferApplianceClient) ListTransferAppliances(ctx context.Context, request ListTransferAppliancesRequest) (response ListTransferAppliancesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listTransferAppliances, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListTransferAppliancesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListTransferAppliancesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListTransferAppliancesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListTransferAppliancesResponse")
	}
	return
}

// listTransferAppliances implements the OCIOperation interface (enables retrying operations)
func (client TransferApplianceClient) listTransferAppliances(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/transferJobs/{id}/transferAppliances")
	if err != nil {
		return nil, err
	}

	var response ListTransferAppliancesResponse
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

// UpdateTransferAppliance Updates a Transfer Appliance
func (client TransferApplianceClient) UpdateTransferAppliance(ctx context.Context, request UpdateTransferApplianceRequest) (response UpdateTransferApplianceResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateTransferAppliance, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateTransferApplianceResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateTransferApplianceResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateTransferApplianceResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateTransferApplianceResponse")
	}
	return
}

// updateTransferAppliance implements the OCIOperation interface (enables retrying operations)
func (client TransferApplianceClient) updateTransferAppliance(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPut, "/transferJobs/{id}/transferAppliances/{transferApplianceLabel}")
	if err != nil {
		return nil, err
	}

	var response UpdateTransferApplianceResponse
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
