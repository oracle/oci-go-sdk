// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Support Management API
//
// Use the Support Management API to manage support requests. For more information, see Getting Help and Contacting Support (https://docs.cloud.oracle.com/iaas/Content/GSG/Tasks/contactingsupport.htm). **Note**: Before you can create service requests with this API, you need to have an Oracle Single Sign On (SSO) account, and you need to register your Customer Support Identifier (CSI) with My Oracle Support.
//

package cims

import (
	"context"
	"fmt"
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

//UserClient a client for User
type UserClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewUserClientWithConfigurationProvider Creates a new default User client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewUserClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client UserClient, err error) {
	baseClient, err := common.NewClientWithConfig(configProvider)
	if err != nil {
		return
	}

	return newUserClientFromBaseClient(baseClient, configProvider)
}

// NewUserClientWithOboToken Creates a new default User client with the given configuration provider.
// The obotoken will be added to default headers and signed; the configuration provider will be used for the signer
//  as well as reading the region
func NewUserClientWithOboToken(configProvider common.ConfigurationProvider, oboToken string) (client UserClient, err error) {
	baseClient, err := common.NewClientWithOboToken(configProvider, oboToken)
	if err != nil {
		return
	}

	return newUserClientFromBaseClient(baseClient, configProvider)
}

func newUserClientFromBaseClient(baseClient common.BaseClient, configProvider common.ConfigurationProvider) (client UserClient, err error) {
	client = UserClient{BaseClient: baseClient}
	client.BasePath = "20181231"
	err = client.setConfigurationProvider(configProvider)
	return
}

// SetRegion overrides the region of this client.
func (client *UserClient) SetRegion(region string) {
	client.Host = common.StringToRegion(region).EndpointForTemplate("cims", "https://incidentmanagement.{region}.{secondLevelDomain}")
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *UserClient) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
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
func (client *UserClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// CreateUser Create user to request Customer Support Identifier(CSI) to Customer User Administrator(CUA).
func (client UserClient) CreateUser(ctx context.Context, request CreateUserRequest) (response CreateUserResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.createUser, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateUserResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateUserResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateUserResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateUserResponse")
	}
	return
}

// createUser implements the OCIOperation interface (enables retrying operations)
func (client UserClient) createUser(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/v2/users")
	if err != nil {
		return nil, err
	}

	var response CreateUserResponse
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
