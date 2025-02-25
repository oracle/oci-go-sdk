// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Email Delivery Submission API
//
// Use the Email Delivery API to send high-volume and application-generated emails.
// For more information, see Overview of the Email Delivery Service (https://docs.oracle.com/iaas/Content/Email/Concepts/overview.htm).
//

package emaildataplane

import (
	"context"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"github.com/oracle/oci-go-sdk/v65/common/auth"
	"net/http"
)

// EmailDPClient a client for EmailDP
type EmailDPClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewEmailDPClientWithConfigurationProvider Creates a new default EmailDP client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewEmailDPClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client EmailDPClient, err error) {
	if enabled := common.CheckForEnabledServices("emaildataplane"); !enabled {
		return client, fmt.Errorf("the Developer Tool configuration disabled this service, this behavior is controlled by OciSdkEnabledServicesMap variables. Please check if your local developer-tool-configuration.json file configured the service you're targeting or contact the cloud provider on the availability of this service")
	}
	provider, err := auth.GetGenericConfigurationProvider(configProvider)
	if err != nil {
		return client, err
	}
	baseClient, e := common.NewClientWithConfig(provider)
	if e != nil {
		return client, e
	}
	return newEmailDPClientFromBaseClient(baseClient, provider)
}

// NewEmailDPClientWithOboToken Creates a new default EmailDP client with the given configuration provider.
// The obotoken will be added to default headers and signed; the configuration provider will be used for the signer
//
//	as well as reading the region
func NewEmailDPClientWithOboToken(configProvider common.ConfigurationProvider, oboToken string) (client EmailDPClient, err error) {
	baseClient, err := common.NewClientWithOboToken(configProvider, oboToken)
	if err != nil {
		return client, err
	}

	return newEmailDPClientFromBaseClient(baseClient, configProvider)
}

func newEmailDPClientFromBaseClient(baseClient common.BaseClient, configProvider common.ConfigurationProvider) (client EmailDPClient, err error) {
	// EmailDP service default circuit breaker is enabled
	baseClient.Configuration.CircuitBreaker = common.NewCircuitBreaker(common.DefaultCircuitBreakerSettingWithServiceName("EmailDP"))
	common.ConfigCircuitBreakerFromEnvVar(&baseClient)
	common.ConfigCircuitBreakerFromGlobalVar(&baseClient)

	client = EmailDPClient{BaseClient: baseClient}
	client.BasePath = "20220926"
	err = client.setConfigurationProvider(configProvider)
	return
}

// SetRegion overrides the region of this client.
func (client *EmailDPClient) SetRegion(region string) {
	client.Host = common.StringToRegion(region).EndpointForTemplate("emaildataplane", "https://cell0.submit.email.{region}.oci.{secondLevelDomain}")
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *EmailDPClient) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
	if ok, err := common.IsConfigurationProviderValid(configProvider); !ok {
		return err
	}

	// Error has been checked already
	region, _ := configProvider.Region()
	client.SetRegion(region)
	if client.Host == "" {
		return fmt.Errorf("invalid region or Host. Endpoint cannot be constructed without endpointServiceName or serviceEndpointTemplate for a dotted region")
	}
	client.config = &configProvider
	return nil
}

// ConfigurationProvider the ConfigurationProvider used in this client, or null if none set
func (client *EmailDPClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// SubmitEmail Submits a formatted email.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/emaildataplane/SubmitEmail.go.html to see an example of how to use SubmitEmail API.
// A default retry strategy applies to this operation SubmitEmail()
func (client EmailDPClient) SubmitEmail(ctx context.Context, request SubmitEmailRequest) (response SubmitEmailResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.submitEmail, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = SubmitEmailResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = SubmitEmailResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(SubmitEmailResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into SubmitEmailResponse")
	}
	return
}

// submitEmail implements the OCIOperation interface (enables retrying operations)
func (client EmailDPClient) submitEmail(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/actions/submitEmail", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response SubmitEmailResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/emaildeliverysubmission/20220926/EmailSubmittedResponse/SubmitEmail"
		err = common.PostProcessServiceError(err, "EmailDP", "SubmitEmail", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}
