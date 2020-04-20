// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Secrets
//
// API for retrieving secrets from vaults.
//

package secrets

import (
	"context"
	"fmt"
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

//SecretsClient a client for Secrets
type SecretsClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewSecretsClientWithConfigurationProvider Creates a new default Secrets client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewSecretsClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client SecretsClient, err error) {
	baseClient, err := common.NewClientWithConfig(configProvider)
	if err != nil {
		return
	}

	return newSecretsClientFromBaseClient(baseClient, configProvider)
}

// NewSecretsClientWithOboToken Creates a new default Secrets client with the given configuration provider.
// The obotoken will be added to default headers and signed; the configuration provider will be used for the signer
//  as well as reading the region
func NewSecretsClientWithOboToken(configProvider common.ConfigurationProvider, oboToken string) (client SecretsClient, err error) {
	baseClient, err := common.NewClientWithOboToken(configProvider, oboToken)
	if err != nil {
		return
	}

	return newSecretsClientFromBaseClient(baseClient, configProvider)
}

func newSecretsClientFromBaseClient(baseClient common.BaseClient, configProvider common.ConfigurationProvider) (client SecretsClient, err error) {
	client = SecretsClient{BaseClient: baseClient}
	client.BasePath = "20190301"
	err = client.setConfigurationProvider(configProvider)
	return
}

// SetRegion overrides the region of this client.
func (client *SecretsClient) SetRegion(region string) {
	client.Host = common.StringToRegion(region).EndpointForTemplate("secrets", "https://secrets.vaults.{region}.oci.{secondLevelDomain}")
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *SecretsClient) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
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
func (client *SecretsClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// GetSecretBundle Gets a secret bundle that matches either the specified `stage`, `label`, or `versionNumber` parameter.
// If none of these parameters are provided, the bundle for the secret version marked as `CURRENT` will be returned.
func (client SecretsClient) GetSecretBundle(ctx context.Context, request GetSecretBundleRequest) (response GetSecretBundleResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getSecretBundle, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetSecretBundleResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetSecretBundleResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetSecretBundleResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetSecretBundleResponse")
	}
	return
}

// getSecretBundle implements the OCIOperation interface (enables retrying operations)
func (client SecretsClient) getSecretBundle(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/secretbundles/{secretId}")
	if err != nil {
		return nil, err
	}

	var response GetSecretBundleResponse
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

// ListSecretBundleVersions Lists all secret bundle versions for the specified secret.
func (client SecretsClient) ListSecretBundleVersions(ctx context.Context, request ListSecretBundleVersionsRequest) (response ListSecretBundleVersionsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listSecretBundleVersions, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListSecretBundleVersionsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListSecretBundleVersionsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListSecretBundleVersionsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListSecretBundleVersionsResponse")
	}
	return
}

// listSecretBundleVersions implements the OCIOperation interface (enables retrying operations)
func (client SecretsClient) listSecretBundleVersions(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/secretbundles/{secretId}/versions")
	if err != nil {
		return nil, err
	}

	var response ListSecretBundleVersionsResponse
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
