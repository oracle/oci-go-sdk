// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Model Deployment Data Plane API
//
// Model deployments are a managed resource in the OCI Data Science service to use to deploy machine learning models as HTTP endpoints in OCI. Deploying machine learning models as web applications (HTTP API endpoints) serving predictions in real time is the most common way that models are productionized. HTTP endpoints are flexible and can serve requests for model predictions.
// For more information, see Model Deployments (https://docs.oracle.com/en-us/iaas/data-science/using/model-dep-about.htm)
//

package modeldeployment

import (
	"context"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"github.com/oracle/oci-go-sdk/v65/common/auth"
	"net/http"

	"regexp"
)

// ModelDeploymentClient a client for ModelDeployment
type ModelDeploymentClient struct {
	common.BaseClient
	config                   *common.ConfigurationProvider
	requiredParamsInEndpoint map[string][]common.TemplateParamForPerRealmEndpoint
}

// NewModelDeploymentClientWithConfigurationProvider Creates a new default ModelDeployment client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewModelDeploymentClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client ModelDeploymentClient, err error) {
	if enabled := common.CheckForEnabledServices("modeldeployment"); !enabled {
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
	return newModelDeploymentClientFromBaseClient(baseClient, provider)
}

// NewModelDeploymentClientWithOboToken Creates a new default ModelDeployment client with the given configuration provider.
// The obotoken will be added to default headers and signed; the configuration provider will be used for the signer
//
//	as well as reading the region
func NewModelDeploymentClientWithOboToken(configProvider common.ConfigurationProvider, oboToken string) (client ModelDeploymentClient, err error) {
	baseClient, err := common.NewClientWithOboToken(configProvider, oboToken)
	if err != nil {
		return client, err
	}

	return newModelDeploymentClientFromBaseClient(baseClient, configProvider)
}

func newModelDeploymentClientFromBaseClient(baseClient common.BaseClient, configProvider common.ConfigurationProvider) (client ModelDeploymentClient, err error) {
	// ModelDeployment service default circuit breaker is enabled
	baseClient.Configuration.CircuitBreaker = common.NewCircuitBreaker(common.DefaultCircuitBreakerSettingWithServiceName("ModelDeployment"))
	common.ConfigCircuitBreakerFromEnvVar(&baseClient)
	common.ConfigCircuitBreakerFromGlobalVar(&baseClient)

	client = ModelDeploymentClient{BaseClient: baseClient}
	client.BasePath = ""
	err = client.setConfigurationProvider(configProvider)
	return
}

// SetRegion overrides the region of this client.
func (client *ModelDeploymentClient) SetRegion(region string) {
	client.Host = common.StringToRegion(region).EndpointForTemplate("modeldeployment", "https://md.datascience.{region}.oci.{secondLevelDomain}")
	client.parseEndpointTemplatePerRealm()
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *ModelDeploymentClient) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
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
func (client *ModelDeploymentClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// getEndpointTemplatePerRealm returns the endpoint template for the given region, if not found, returns the default endpoint template
func (client *ModelDeploymentClient) getEndpointTemplatePerRealm(region string) string {
	if client.IsOciRealmSpecificServiceEndpointTemplateEnabled() {
		realm, _ := common.StringToRegion(region).RealmID()
		templatePerRealmDict := map[string]string{
			"oc1": "https://modeldeployment.{region}.oci.customer-oci.com",
		}
		if template, ok := templatePerRealmDict[realm]; ok {
			return template
		}
	}
	return "https://md.datascience.{region}.oci.{secondLevelDomain}"
}

// parseEndpointTemplatePerRealm parses the endpoint template per realm from the service endpoint template
// This function will build a map of template params to their values, this map is used when building the API endpoint
func (client *ModelDeploymentClient) parseEndpointTemplatePerRealm() {
	client.requiredParamsInEndpoint = make(map[string][]common.TemplateParamForPerRealmEndpoint)
	templateRegex := regexp.MustCompile(`{.*?}`)
	templateSubRegex := regexp.MustCompile(`{(.+)\+Dot}`)
	templates := templateRegex.FindAllString(client.Host, -1)
	for _, template := range templates {
		templateParam := templateSubRegex.FindStringSubmatch(template)
		if len(templateParam) > 1 {
			client.requiredParamsInEndpoint[templateParam[1]] = append(client.requiredParamsInEndpoint[templateParam[1]], common.TemplateParamForPerRealmEndpoint{
				Template:    templateParam[0],
				EndsWithDot: true,
			})
		} else {
			templateParam := template[1 : len(template)-1]
			client.requiredParamsInEndpoint[templateParam] = append(client.requiredParamsInEndpoint[templateParam], common.TemplateParamForPerRealmEndpoint{
				Template:    template,
				EndsWithDot: false,
			})
		}
	}
}

// SetCustomClientConfiguration sets client with retry and other custom configurations
func (client *ModelDeploymentClient) SetCustomClientConfiguration(config common.CustomClientConfiguration) {
	client.Configuration = config
	client.refreshRegion()
}

// refreshRegion will refresh the region of this client, this function will be called after setting the CustomClientConfiguration
func (client *ModelDeploymentClient) refreshRegion() {
	configProvider := *client.config
	region, _ := configProvider.Region()
	client.SetRegion(region)
}

// Predict Invoking a model deployment calls the predict endpoint of the model deployment URI.
// This endpoint takes sample data as input and is processed using the predict() function in score.py model artifact file
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/modeldeployment/Predict.go.html to see an example of how to use Predict API.
// A default retry strategy applies to this operation Predict()
func (client ModelDeploymentClient) Predict(ctx context.Context, request PredictRequest) (response PredictResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.predict, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = PredictResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = PredictResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(PredictResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into PredictResponse")
	}
	return
}

// predict implements the OCIOperation interface (enables retrying operations)
func (client ModelDeploymentClient) predict(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/{modelDeploymentId}/predict", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	host := client.Host

	common.SetMissingTemplateParams(&client.BaseClient)
	defer func() {
		client.Host = host
	}()

	var response PredictResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "ModelDeployment", "Predict", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// PredictWithResponseStream Invoking a model deployment calls the predictWithResponseStream endpoint of the model deployment URI to get the streaming result.
// This endpoint takes sample data as input and is processed using the predict() function in score.py model artifact file
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/modeldeployment/PredictWithResponseStream.go.html to see an example of how to use PredictWithResponseStream API.
// A default retry strategy applies to this operation PredictWithResponseStream()
func (client ModelDeploymentClient) PredictWithResponseStream(ctx context.Context, request PredictWithResponseStreamRequest) (response PredictWithResponseStreamResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.predictWithResponseStream, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = PredictWithResponseStreamResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = PredictWithResponseStreamResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(PredictWithResponseStreamResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into PredictWithResponseStreamResponse")
	}
	return
}

// predictWithResponseStream implements the OCIOperation interface (enables retrying operations)
func (client ModelDeploymentClient) predictWithResponseStream(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/{modelDeploymentId}/predictWithResponseStream", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	host := client.Host

	common.SetMissingTemplateParams(&client.BaseClient)
	defer func() {
		client.Host = host
	}()

	var response PredictWithResponseStreamResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "ModelDeployment", "PredictWithResponseStream", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}
