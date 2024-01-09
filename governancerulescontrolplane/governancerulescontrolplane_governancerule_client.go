// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// GovernanceRulesControlPlane API
//
// A description of the GovernanceRulesControlPlane API
//

package governancerulescontrolplane

import (
	"context"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"github.com/oracle/oci-go-sdk/v65/common/auth"
	"net/http"
)

// GovernanceRuleClient a client for GovernanceRule
type GovernanceRuleClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewGovernanceRuleClientWithConfigurationProvider Creates a new default GovernanceRule client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewGovernanceRuleClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client GovernanceRuleClient, err error) {
	if enabled := common.CheckForEnabledServices("governancerulescontrolplane"); !enabled {
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
	return newGovernanceRuleClientFromBaseClient(baseClient, provider)
}

// NewGovernanceRuleClientWithOboToken Creates a new default GovernanceRule client with the given configuration provider.
// The obotoken will be added to default headers and signed; the configuration provider will be used for the signer
//
//	as well as reading the region
func NewGovernanceRuleClientWithOboToken(configProvider common.ConfigurationProvider, oboToken string) (client GovernanceRuleClient, err error) {
	baseClient, err := common.NewClientWithOboToken(configProvider, oboToken)
	if err != nil {
		return client, err
	}

	return newGovernanceRuleClientFromBaseClient(baseClient, configProvider)
}

func newGovernanceRuleClientFromBaseClient(baseClient common.BaseClient, configProvider common.ConfigurationProvider) (client GovernanceRuleClient, err error) {
	// GovernanceRule service default circuit breaker is enabled
	baseClient.Configuration.CircuitBreaker = common.NewCircuitBreaker(common.DefaultCircuitBreakerSettingWithServiceName("GovernanceRule"))
	common.ConfigCircuitBreakerFromEnvVar(&baseClient)
	common.ConfigCircuitBreakerFromGlobalVar(&baseClient)

	client = GovernanceRuleClient{BaseClient: baseClient}
	client.BasePath = "20220504"
	err = client.setConfigurationProvider(configProvider)
	return
}

// SetRegion overrides the region of this client.
func (client *GovernanceRuleClient) SetRegion(region string) {
	client.Host = common.StringToRegion(region).EndpointForTemplate("governancerulescontrolplane", "https://governance-rules.organizations.{region}.oci.{secondLevelDomain}")
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *GovernanceRuleClient) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
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
func (client *GovernanceRuleClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// CreateGovernanceRule Create governance rule in the root compartment only. Either relatedResourceId or template must be supplied.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/governancerulescontrolplane/CreateGovernanceRule.go.html to see an example of how to use CreateGovernanceRule API.
func (client GovernanceRuleClient) CreateGovernanceRule(ctx context.Context, request CreateGovernanceRuleRequest) (response CreateGovernanceRuleResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createGovernanceRule, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateGovernanceRuleResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateGovernanceRuleResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateGovernanceRuleResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateGovernanceRuleResponse")
	}
	return
}

// createGovernanceRule implements the OCIOperation interface (enables retrying operations)
func (client GovernanceRuleClient) createGovernanceRule(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/governanceRules", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateGovernanceRuleResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "GovernanceRule", "CreateGovernanceRule", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateInclusionCriterion Create inclusion criterion of type tenancy or tag for the governance rule.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/governancerulescontrolplane/CreateInclusionCriterion.go.html to see an example of how to use CreateInclusionCriterion API.
func (client GovernanceRuleClient) CreateInclusionCriterion(ctx context.Context, request CreateInclusionCriterionRequest) (response CreateInclusionCriterionResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createInclusionCriterion, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateInclusionCriterionResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateInclusionCriterionResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateInclusionCriterionResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateInclusionCriterionResponse")
	}
	return
}

// createInclusionCriterion implements the OCIOperation interface (enables retrying operations)
func (client GovernanceRuleClient) createInclusionCriterion(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/inclusionCriteria", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateInclusionCriterionResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "GovernanceRule", "CreateInclusionCriterion", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteGovernanceRule Delete the specified governance rule.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/governancerulescontrolplane/DeleteGovernanceRule.go.html to see an example of how to use DeleteGovernanceRule API.
func (client GovernanceRuleClient) DeleteGovernanceRule(ctx context.Context, request DeleteGovernanceRuleRequest) (response DeleteGovernanceRuleResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteGovernanceRule, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteGovernanceRuleResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteGovernanceRuleResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteGovernanceRuleResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteGovernanceRuleResponse")
	}
	return
}

// deleteGovernanceRule implements the OCIOperation interface (enables retrying operations)
func (client GovernanceRuleClient) deleteGovernanceRule(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/governanceRules/{governanceRuleId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteGovernanceRuleResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "GovernanceRule", "DeleteGovernanceRule", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteInclusionCriterion Delete the specified inclusion criterion.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/governancerulescontrolplane/DeleteInclusionCriterion.go.html to see an example of how to use DeleteInclusionCriterion API.
func (client GovernanceRuleClient) DeleteInclusionCriterion(ctx context.Context, request DeleteInclusionCriterionRequest) (response DeleteInclusionCriterionResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteInclusionCriterion, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteInclusionCriterionResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteInclusionCriterionResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteInclusionCriterionResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteInclusionCriterionResponse")
	}
	return
}

// deleteInclusionCriterion implements the OCIOperation interface (enables retrying operations)
func (client GovernanceRuleClient) deleteInclusionCriterion(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/inclusionCriteria/{inclusionCriterionId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteInclusionCriterionResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "GovernanceRule", "DeleteInclusionCriterion", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetEnforcedGovernanceRule Get the specified enforced governance rule's information.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/governancerulescontrolplane/GetEnforcedGovernanceRule.go.html to see an example of how to use GetEnforcedGovernanceRule API.
func (client GovernanceRuleClient) GetEnforcedGovernanceRule(ctx context.Context, request GetEnforcedGovernanceRuleRequest) (response GetEnforcedGovernanceRuleResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getEnforcedGovernanceRule, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetEnforcedGovernanceRuleResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetEnforcedGovernanceRuleResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetEnforcedGovernanceRuleResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetEnforcedGovernanceRuleResponse")
	}
	return
}

// getEnforcedGovernanceRule implements the OCIOperation interface (enables retrying operations)
func (client GovernanceRuleClient) getEnforcedGovernanceRule(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/enforcedGovernanceRules/{enforcedGovernanceRuleId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetEnforcedGovernanceRuleResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "GovernanceRule", "GetEnforcedGovernanceRule", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetGovernanceRule Get the specified governance rule's information.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/governancerulescontrolplane/GetGovernanceRule.go.html to see an example of how to use GetGovernanceRule API.
func (client GovernanceRuleClient) GetGovernanceRule(ctx context.Context, request GetGovernanceRuleRequest) (response GetGovernanceRuleResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getGovernanceRule, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetGovernanceRuleResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetGovernanceRuleResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetGovernanceRuleResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetGovernanceRuleResponse")
	}
	return
}

// getGovernanceRule implements the OCIOperation interface (enables retrying operations)
func (client GovernanceRuleClient) getGovernanceRule(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/governanceRules/{governanceRuleId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetGovernanceRuleResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "GovernanceRule", "GetGovernanceRule", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetInclusionCriterion Get the specified inclusion criterion's information.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/governancerulescontrolplane/GetInclusionCriterion.go.html to see an example of how to use GetInclusionCriterion API.
func (client GovernanceRuleClient) GetInclusionCriterion(ctx context.Context, request GetInclusionCriterionRequest) (response GetInclusionCriterionResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getInclusionCriterion, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetInclusionCriterionResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetInclusionCriterionResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetInclusionCriterionResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetInclusionCriterionResponse")
	}
	return
}

// getInclusionCriterion implements the OCIOperation interface (enables retrying operations)
func (client GovernanceRuleClient) getInclusionCriterion(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/inclusionCriteria/{inclusionCriterionId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetInclusionCriterionResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "GovernanceRule", "GetInclusionCriterion", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetTenancyAttachment Get the specified tenancy attachment's information.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/governancerulescontrolplane/GetTenancyAttachment.go.html to see an example of how to use GetTenancyAttachment API.
func (client GovernanceRuleClient) GetTenancyAttachment(ctx context.Context, request GetTenancyAttachmentRequest) (response GetTenancyAttachmentResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getTenancyAttachment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetTenancyAttachmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetTenancyAttachmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetTenancyAttachmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetTenancyAttachmentResponse")
	}
	return
}

// getTenancyAttachment implements the OCIOperation interface (enables retrying operations)
func (client GovernanceRuleClient) getTenancyAttachment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/tenancyAttachments/{tenancyAttachmentId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetTenancyAttachmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "GovernanceRule", "GetTenancyAttachment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListEnforcedGovernanceRules List enforced governance rules. Either compartment id or enforced governance rule id must be supplied.
// An optional governance rule type or a display name can also be supplied.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/governancerulescontrolplane/ListEnforcedGovernanceRules.go.html to see an example of how to use ListEnforcedGovernanceRules API.
func (client GovernanceRuleClient) ListEnforcedGovernanceRules(ctx context.Context, request ListEnforcedGovernanceRulesRequest) (response ListEnforcedGovernanceRulesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listEnforcedGovernanceRules, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListEnforcedGovernanceRulesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListEnforcedGovernanceRulesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListEnforcedGovernanceRulesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListEnforcedGovernanceRulesResponse")
	}
	return
}

// listEnforcedGovernanceRules implements the OCIOperation interface (enables retrying operations)
func (client GovernanceRuleClient) listEnforcedGovernanceRules(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/enforcedGovernanceRules", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListEnforcedGovernanceRulesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "GovernanceRule", "ListEnforcedGovernanceRules", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListGovernanceRules List governance rules. Either compartment id or governance rule id must be supplied.
// An optional lifecycle state, display name or a governance rule type can also be supplied.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/governancerulescontrolplane/ListGovernanceRules.go.html to see an example of how to use ListGovernanceRules API.
func (client GovernanceRuleClient) ListGovernanceRules(ctx context.Context, request ListGovernanceRulesRequest) (response ListGovernanceRulesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listGovernanceRules, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListGovernanceRulesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListGovernanceRulesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListGovernanceRulesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListGovernanceRulesResponse")
	}
	return
}

// listGovernanceRules implements the OCIOperation interface (enables retrying operations)
func (client GovernanceRuleClient) listGovernanceRules(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/governanceRules", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListGovernanceRulesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "GovernanceRule", "ListGovernanceRules", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListInclusionCriteria List inclusion criteria associated with a governance rule. Governance rule id must be supplied.
// An optional inclusion criterion id or a lifecycle state can also be supplied.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/governancerulescontrolplane/ListInclusionCriteria.go.html to see an example of how to use ListInclusionCriteria API.
func (client GovernanceRuleClient) ListInclusionCriteria(ctx context.Context, request ListInclusionCriteriaRequest) (response ListInclusionCriteriaResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listInclusionCriteria, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListInclusionCriteriaResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListInclusionCriteriaResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListInclusionCriteriaResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListInclusionCriteriaResponse")
	}
	return
}

// listInclusionCriteria implements the OCIOperation interface (enables retrying operations)
func (client GovernanceRuleClient) listInclusionCriteria(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/inclusionCriteria", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListInclusionCriteriaResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "GovernanceRule", "ListInclusionCriteria", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListTenancyAttachments List tenancy attachments. Either compartment id, governance rule id or tenancy attachment id must be supplied.
// An optional lifecycle state or a child tenancy id can also be supplied.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/governancerulescontrolplane/ListTenancyAttachments.go.html to see an example of how to use ListTenancyAttachments API.
func (client GovernanceRuleClient) ListTenancyAttachments(ctx context.Context, request ListTenancyAttachmentsRequest) (response ListTenancyAttachmentsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listTenancyAttachments, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListTenancyAttachmentsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListTenancyAttachmentsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListTenancyAttachmentsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListTenancyAttachmentsResponse")
	}
	return
}

// listTenancyAttachments implements the OCIOperation interface (enables retrying operations)
func (client GovernanceRuleClient) listTenancyAttachments(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/tenancyAttachments", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListTenancyAttachmentsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "GovernanceRule", "ListTenancyAttachments", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// RetryGovernanceRule Retry the creation of the specified governance rule.
// Used by the tenancy admins when all the workflow retries have exhausted.
// When provided, If-Match is checked against ETag values of the resource.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/governancerulescontrolplane/RetryGovernanceRule.go.html to see an example of how to use RetryGovernanceRule API.
func (client GovernanceRuleClient) RetryGovernanceRule(ctx context.Context, request RetryGovernanceRuleRequest) (response RetryGovernanceRuleResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.retryGovernanceRule, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = RetryGovernanceRuleResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = RetryGovernanceRuleResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(RetryGovernanceRuleResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into RetryGovernanceRuleResponse")
	}
	return
}

// retryGovernanceRule implements the OCIOperation interface (enables retrying operations)
func (client GovernanceRuleClient) retryGovernanceRule(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/governanceRules/{governanceRuleId}/actions/retry", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response RetryGovernanceRuleResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "GovernanceRule", "RetryGovernanceRule", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// RetryTenancyAttachment Retry governance rule application for the specified tenancy attachment id.
// Used by the tenancy admins when all the workflow retries have exhausted.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/governancerulescontrolplane/RetryTenancyAttachment.go.html to see an example of how to use RetryTenancyAttachment API.
func (client GovernanceRuleClient) RetryTenancyAttachment(ctx context.Context, request RetryTenancyAttachmentRequest) (response RetryTenancyAttachmentResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.retryTenancyAttachment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = RetryTenancyAttachmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = RetryTenancyAttachmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(RetryTenancyAttachmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into RetryTenancyAttachmentResponse")
	}
	return
}

// retryTenancyAttachment implements the OCIOperation interface (enables retrying operations)
func (client GovernanceRuleClient) retryTenancyAttachment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/tenancyAttachments/{tenancyAttachmentId}/actions/retry", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response RetryTenancyAttachmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "GovernanceRule", "RetryTenancyAttachment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateGovernanceRule Update the specified governance rule.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/governancerulescontrolplane/UpdateGovernanceRule.go.html to see an example of how to use UpdateGovernanceRule API.
func (client GovernanceRuleClient) UpdateGovernanceRule(ctx context.Context, request UpdateGovernanceRuleRequest) (response UpdateGovernanceRuleResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateGovernanceRule, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateGovernanceRuleResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateGovernanceRuleResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateGovernanceRuleResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateGovernanceRuleResponse")
	}
	return
}

// updateGovernanceRule implements the OCIOperation interface (enables retrying operations)
func (client GovernanceRuleClient) updateGovernanceRule(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/governanceRules/{governanceRuleId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateGovernanceRuleResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "GovernanceRule", "UpdateGovernanceRule", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}
