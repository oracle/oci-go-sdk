// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Language API
//
// OCI Language Service solutions can help enterprise customers integrate AI into their products immediately using our proven,
//     pre-trained and custom models or containers, without a need to set up an house team of AI and ML experts.
//     This allows enterprises to focus on business drivers and development work rather than AI and ML operations, which shortens the time to market.
//

package ailanguage

import (
	"context"
	"fmt"
	"github.com/oracle/oci-go-sdk/v64/common"
	"github.com/oracle/oci-go-sdk/v64/common/auth"
	"net/http"
)

//AIServiceLanguageClient a client for AIServiceLanguage
type AIServiceLanguageClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewAIServiceLanguageClientWithConfigurationProvider Creates a new default AIServiceLanguage client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewAIServiceLanguageClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client AIServiceLanguageClient, err error) {
	provider, err := auth.GetGenericConfigurationProvider(configProvider)
	if err != nil {
		return client, err
	}
	baseClient, e := common.NewClientWithConfig(provider)
	if e != nil {
		return client, e
	}
	return newAIServiceLanguageClientFromBaseClient(baseClient, provider)
}

// NewAIServiceLanguageClientWithOboToken Creates a new default AIServiceLanguage client with the given configuration provider.
// The obotoken will be added to default headers and signed; the configuration provider will be used for the signer
//  as well as reading the region
func NewAIServiceLanguageClientWithOboToken(configProvider common.ConfigurationProvider, oboToken string) (client AIServiceLanguageClient, err error) {
	baseClient, err := common.NewClientWithOboToken(configProvider, oboToken)
	if err != nil {
		return client, err
	}

	return newAIServiceLanguageClientFromBaseClient(baseClient, configProvider)
}

func newAIServiceLanguageClientFromBaseClient(baseClient common.BaseClient, configProvider common.ConfigurationProvider) (client AIServiceLanguageClient, err error) {
	// AIServiceLanguage service default circuit breaker is enabled
	baseClient.Configuration.CircuitBreaker = common.NewCircuitBreaker(common.DefaultCircuitBreakerSettingWithServiceName())
	common.ConfigCircuitBreakerFromEnvVar(&baseClient)
	common.ConfigCircuitBreakerFromGlobalVar(&baseClient)

	client = AIServiceLanguageClient{BaseClient: baseClient}
	client.BasePath = "20210101"
	err = client.setConfigurationProvider(configProvider)
	return
}

// SetRegion overrides the region of this client.
func (client *AIServiceLanguageClient) SetRegion(region string) {
	client.Host = common.StringToRegion(region).EndpointForTemplate("ailanguage", "https://language.aiservice.{region}.oci.{secondLevelDomain}")
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *AIServiceLanguageClient) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
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
func (client *AIServiceLanguageClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// BatchDetectDominantLanguage Make a detect call to language detection pre-deployed model.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/ailanguage/BatchDetectDominantLanguage.go.html to see an example of how to use BatchDetectDominantLanguage API.
func (client AIServiceLanguageClient) BatchDetectDominantLanguage(ctx context.Context, request BatchDetectDominantLanguageRequest) (response BatchDetectDominantLanguageResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.batchDetectDominantLanguage, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = BatchDetectDominantLanguageResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = BatchDetectDominantLanguageResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(BatchDetectDominantLanguageResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into BatchDetectDominantLanguageResponse")
	}
	return
}

// batchDetectDominantLanguage implements the OCIOperation interface (enables retrying operations)
func (client AIServiceLanguageClient) batchDetectDominantLanguage(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/actions/batchDetectDominantLanguage", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response BatchDetectDominantLanguageResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/language/20210101/BatchDetectDominantLanguage/BatchDetectDominantLanguage"
		err = common.PostProcessServiceError(err, "AIServiceLanguage", "BatchDetectDominantLanguage", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// BatchDetectLanguageEntities Make a batch detect call to entity pre-deployed model
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/ailanguage/BatchDetectLanguageEntities.go.html to see an example of how to use BatchDetectLanguageEntities API.
func (client AIServiceLanguageClient) BatchDetectLanguageEntities(ctx context.Context, request BatchDetectLanguageEntitiesRequest) (response BatchDetectLanguageEntitiesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.batchDetectLanguageEntities, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = BatchDetectLanguageEntitiesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = BatchDetectLanguageEntitiesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(BatchDetectLanguageEntitiesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into BatchDetectLanguageEntitiesResponse")
	}
	return
}

// batchDetectLanguageEntities implements the OCIOperation interface (enables retrying operations)
func (client AIServiceLanguageClient) batchDetectLanguageEntities(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/actions/batchDetectLanguageEntities", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response BatchDetectLanguageEntitiesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/language/20210101/BatchDetectLanguageEntities/BatchDetectLanguageEntities"
		err = common.PostProcessServiceError(err, "AIServiceLanguage", "BatchDetectLanguageEntities", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// BatchDetectLanguageKeyPhrases Make a detect call to the keyPhrase pre-deployed model.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/ailanguage/BatchDetectLanguageKeyPhrases.go.html to see an example of how to use BatchDetectLanguageKeyPhrases API.
func (client AIServiceLanguageClient) BatchDetectLanguageKeyPhrases(ctx context.Context, request BatchDetectLanguageKeyPhrasesRequest) (response BatchDetectLanguageKeyPhrasesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.batchDetectLanguageKeyPhrases, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = BatchDetectLanguageKeyPhrasesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = BatchDetectLanguageKeyPhrasesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(BatchDetectLanguageKeyPhrasesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into BatchDetectLanguageKeyPhrasesResponse")
	}
	return
}

// batchDetectLanguageKeyPhrases implements the OCIOperation interface (enables retrying operations)
func (client AIServiceLanguageClient) batchDetectLanguageKeyPhrases(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/actions/batchDetectLanguageKeyPhrases", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response BatchDetectLanguageKeyPhrasesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/language/20210101/BatchDetectLanguageKeyPhrases/BatchDetectLanguageKeyPhrases"
		err = common.PostProcessServiceError(err, "AIServiceLanguage", "BatchDetectLanguageKeyPhrases", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// BatchDetectLanguageSentiments Make a detect call to sentiment pre-deployed model.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/ailanguage/BatchDetectLanguageSentiments.go.html to see an example of how to use BatchDetectLanguageSentiments API.
func (client AIServiceLanguageClient) BatchDetectLanguageSentiments(ctx context.Context, request BatchDetectLanguageSentimentsRequest) (response BatchDetectLanguageSentimentsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.batchDetectLanguageSentiments, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = BatchDetectLanguageSentimentsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = BatchDetectLanguageSentimentsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(BatchDetectLanguageSentimentsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into BatchDetectLanguageSentimentsResponse")
	}
	return
}

// batchDetectLanguageSentiments implements the OCIOperation interface (enables retrying operations)
func (client AIServiceLanguageClient) batchDetectLanguageSentiments(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/actions/batchDetectLanguageSentiments", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response BatchDetectLanguageSentimentsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/language/20210101/BatchDetectLanguageSentiments/BatchDetectLanguageSentiments"
		err = common.PostProcessServiceError(err, "AIServiceLanguage", "BatchDetectLanguageSentiments", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// BatchDetectLanguageTextClassification Make a detect call to text classification from the pre-deployed model.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/ailanguage/BatchDetectLanguageTextClassification.go.html to see an example of how to use BatchDetectLanguageTextClassification API.
func (client AIServiceLanguageClient) BatchDetectLanguageTextClassification(ctx context.Context, request BatchDetectLanguageTextClassificationRequest) (response BatchDetectLanguageTextClassificationResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.batchDetectLanguageTextClassification, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = BatchDetectLanguageTextClassificationResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = BatchDetectLanguageTextClassificationResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(BatchDetectLanguageTextClassificationResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into BatchDetectLanguageTextClassificationResponse")
	}
	return
}

// batchDetectLanguageTextClassification implements the OCIOperation interface (enables retrying operations)
func (client AIServiceLanguageClient) batchDetectLanguageTextClassification(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/actions/batchDetectLanguageTextClassification", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response BatchDetectLanguageTextClassificationResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/language/20210101/BatchDetectLanguageTextClassification/BatchDetectLanguageTextClassification"
		err = common.PostProcessServiceError(err, "AIServiceLanguage", "BatchDetectLanguageTextClassification", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DetectDominantLanguage Make a detect call to language detection pre-deployed model.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/ailanguage/DetectDominantLanguage.go.html to see an example of how to use DetectDominantLanguage API.
func (client AIServiceLanguageClient) DetectDominantLanguage(ctx context.Context, request DetectDominantLanguageRequest) (response DetectDominantLanguageResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.detectDominantLanguage, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DetectDominantLanguageResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DetectDominantLanguageResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DetectDominantLanguageResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DetectDominantLanguageResponse")
	}
	return
}

// detectDominantLanguage implements the OCIOperation interface (enables retrying operations)
func (client AIServiceLanguageClient) detectDominantLanguage(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/actions/detectDominantLanguage", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DetectDominantLanguageResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/language/20210101/DetectDominantLanguage/DetectDominantLanguage"
		err = common.PostProcessServiceError(err, "AIServiceLanguage", "DetectDominantLanguage", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DetectLanguageEntities Make a detect call to enitiy pre-deployed model
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/ailanguage/DetectLanguageEntities.go.html to see an example of how to use DetectLanguageEntities API.
func (client AIServiceLanguageClient) DetectLanguageEntities(ctx context.Context, request DetectLanguageEntitiesRequest) (response DetectLanguageEntitiesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.detectLanguageEntities, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DetectLanguageEntitiesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DetectLanguageEntitiesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DetectLanguageEntitiesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DetectLanguageEntitiesResponse")
	}
	return
}

// detectLanguageEntities implements the OCIOperation interface (enables retrying operations)
func (client AIServiceLanguageClient) detectLanguageEntities(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/actions/detectLanguageEntities", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DetectLanguageEntitiesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/language/20210101/DetectLanguageEntities/DetectLanguageEntities"
		err = common.PostProcessServiceError(err, "AIServiceLanguage", "DetectLanguageEntities", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DetectLanguageKeyPhrases Make a detect call to the keyPhrase pre-deployed model.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/ailanguage/DetectLanguageKeyPhrases.go.html to see an example of how to use DetectLanguageKeyPhrases API.
func (client AIServiceLanguageClient) DetectLanguageKeyPhrases(ctx context.Context, request DetectLanguageKeyPhrasesRequest) (response DetectLanguageKeyPhrasesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.detectLanguageKeyPhrases, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DetectLanguageKeyPhrasesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DetectLanguageKeyPhrasesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DetectLanguageKeyPhrasesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DetectLanguageKeyPhrasesResponse")
	}
	return
}

// detectLanguageKeyPhrases implements the OCIOperation interface (enables retrying operations)
func (client AIServiceLanguageClient) detectLanguageKeyPhrases(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/actions/detectLanguageKeyPhrases", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DetectLanguageKeyPhrasesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/language/20210101/DetectLanguageKeyPhrases/DetectLanguageKeyPhrases"
		err = common.PostProcessServiceError(err, "AIServiceLanguage", "DetectLanguageKeyPhrases", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DetectLanguageSentiments Make a detect call to sentiment pre-deployed model.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/ailanguage/DetectLanguageSentiments.go.html to see an example of how to use DetectLanguageSentiments API.
func (client AIServiceLanguageClient) DetectLanguageSentiments(ctx context.Context, request DetectLanguageSentimentsRequest) (response DetectLanguageSentimentsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.detectLanguageSentiments, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DetectLanguageSentimentsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DetectLanguageSentimentsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DetectLanguageSentimentsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DetectLanguageSentimentsResponse")
	}
	return
}

// detectLanguageSentiments implements the OCIOperation interface (enables retrying operations)
func (client AIServiceLanguageClient) detectLanguageSentiments(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/actions/detectLanguageSentiments", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DetectLanguageSentimentsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/language/20210101/DetectLanguageSentiments/DetectLanguageSentiments"
		err = common.PostProcessServiceError(err, "AIServiceLanguage", "DetectLanguageSentiments", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DetectLanguageTextClassification Make a detect call to text classification from the pre-deployed model.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/ailanguage/DetectLanguageTextClassification.go.html to see an example of how to use DetectLanguageTextClassification API.
func (client AIServiceLanguageClient) DetectLanguageTextClassification(ctx context.Context, request DetectLanguageTextClassificationRequest) (response DetectLanguageTextClassificationResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.detectLanguageTextClassification, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DetectLanguageTextClassificationResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DetectLanguageTextClassificationResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DetectLanguageTextClassificationResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DetectLanguageTextClassificationResponse")
	}
	return
}

// detectLanguageTextClassification implements the OCIOperation interface (enables retrying operations)
func (client AIServiceLanguageClient) detectLanguageTextClassification(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/actions/detectLanguageTextClassification", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DetectLanguageTextClassificationResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/language/20210101/DetectLanguageTextClassification/DetectLanguageTextClassification"
		err = common.PostProcessServiceError(err, "AIServiceLanguage", "DetectLanguageTextClassification", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}
