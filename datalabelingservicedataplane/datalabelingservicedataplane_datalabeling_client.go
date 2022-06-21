// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Labeling API
//
// Use Data Labeling API to create Annotations on Images, Texts & Documents, and generate snapshots.
//

package datalabelingservicedataplane

import (
	"context"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"github.com/oracle/oci-go-sdk/v65/common/auth"
	"net/http"
)

//DataLabelingClient a client for DataLabeling
type DataLabelingClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewDataLabelingClientWithConfigurationProvider Creates a new default DataLabeling client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewDataLabelingClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client DataLabelingClient, err error) {
	provider, err := auth.GetGenericConfigurationProvider(configProvider)
	if err != nil {
		return client, err
	}
	baseClient, e := common.NewClientWithConfig(provider)
	if e != nil {
		return client, e
	}
	return newDataLabelingClientFromBaseClient(baseClient, provider)
}

// NewDataLabelingClientWithOboToken Creates a new default DataLabeling client with the given configuration provider.
// The obotoken will be added to default headers and signed; the configuration provider will be used for the signer
//  as well as reading the region
func NewDataLabelingClientWithOboToken(configProvider common.ConfigurationProvider, oboToken string) (client DataLabelingClient, err error) {
	baseClient, err := common.NewClientWithOboToken(configProvider, oboToken)
	if err != nil {
		return client, err
	}

	return newDataLabelingClientFromBaseClient(baseClient, configProvider)
}

func newDataLabelingClientFromBaseClient(baseClient common.BaseClient, configProvider common.ConfigurationProvider) (client DataLabelingClient, err error) {
	// DataLabeling service default circuit breaker is enabled
	baseClient.Configuration.CircuitBreaker = common.NewCircuitBreaker(common.DefaultCircuitBreakerSettingWithServiceName("DataLabeling"))
	common.ConfigCircuitBreakerFromEnvVar(&baseClient)
	common.ConfigCircuitBreakerFromGlobalVar(&baseClient)

	client = DataLabelingClient{BaseClient: baseClient}
	client.BasePath = "20211001"
	err = client.setConfigurationProvider(configProvider)
	return
}

// SetRegion overrides the region of this client.
func (client *DataLabelingClient) SetRegion(region string) {
	client.Host = common.StringToRegion(region).EndpointForTemplate("datalabelingservicedataplane", "https://datalabeling-dp.{region}.oci.{secondLevelDomain}")
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *DataLabelingClient) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
	if ok, err := common.IsConfigurationProviderValid(configProvider); !ok {
		return err
	}

	// Error has been checked already
	region, _ := configProvider.Region()
	client.SetRegion(region)
	if client.Host == "" {
		return fmt.Errorf("Invalid region or Host. Endpoint cannot be constructed without endpointServiceName or serviceEndpointTemplate for a dotted region")
	}
	client.config = &configProvider
	return nil
}

// ConfigurationProvider the ConfigurationProvider used in this client, or null if none set
func (client *DataLabelingClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// CreateAnnotation Creates an annotation.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datalabelingservicedataplane/CreateAnnotation.go.html to see an example of how to use CreateAnnotation API.
// A default retry strategy applies to this operation CreateAnnotation()
func (client DataLabelingClient) CreateAnnotation(ctx context.Context, request CreateAnnotationRequest) (response CreateAnnotationResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.createAnnotation, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateAnnotationResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateAnnotationResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateAnnotationResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateAnnotationResponse")
	}
	return
}

// createAnnotation implements the OCIOperation interface (enables retrying operations)
func (client DataLabelingClient) createAnnotation(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/annotations", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateAnnotationResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/datalabeling-dp/20211001/Annotation/CreateAnnotation"
		err = common.PostProcessServiceError(err, "DataLabeling", "CreateAnnotation", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateRecord Creates a record.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datalabelingservicedataplane/CreateRecord.go.html to see an example of how to use CreateRecord API.
// A default retry strategy applies to this operation CreateRecord()
func (client DataLabelingClient) CreateRecord(ctx context.Context, request CreateRecordRequest) (response CreateRecordResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.createRecord, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateRecordResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateRecordResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateRecordResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateRecordResponse")
	}
	return
}

// createRecord implements the OCIOperation interface (enables retrying operations)
func (client DataLabelingClient) createRecord(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/records", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateRecordResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/datalabeling-dp/20211001/Record/CreateRecord"
		err = common.PostProcessServiceError(err, "DataLabeling", "CreateRecord", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteAnnotation It deletes an annotation resource by identifier.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datalabelingservicedataplane/DeleteAnnotation.go.html to see an example of how to use DeleteAnnotation API.
// A default retry strategy applies to this operation DeleteAnnotation()
func (client DataLabelingClient) DeleteAnnotation(ctx context.Context, request DeleteAnnotationRequest) (response DeleteAnnotationResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteAnnotation, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteAnnotationResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteAnnotationResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteAnnotationResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteAnnotationResponse")
	}
	return
}

// deleteAnnotation implements the OCIOperation interface (enables retrying operations)
func (client DataLabelingClient) deleteAnnotation(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/annotations/{annotationId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteAnnotationResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/datalabeling-dp/20211001/Annotation/DeleteAnnotation"
		err = common.PostProcessServiceError(err, "DataLabeling", "DeleteAnnotation", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteRecord Deletes a record resource by identifier.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datalabelingservicedataplane/DeleteRecord.go.html to see an example of how to use DeleteRecord API.
// A default retry strategy applies to this operation DeleteRecord()
func (client DataLabelingClient) DeleteRecord(ctx context.Context, request DeleteRecordRequest) (response DeleteRecordResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteRecord, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteRecordResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteRecordResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteRecordResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteRecordResponse")
	}
	return
}

// deleteRecord implements the OCIOperation interface (enables retrying operations)
func (client DataLabelingClient) deleteRecord(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/records/{recordId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteRecordResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/datalabeling-dp/20211001/Record/DeleteRecord"
		err = common.PostProcessServiceError(err, "DataLabeling", "DeleteRecord", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetAnnotation Gets an annotation.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datalabelingservicedataplane/GetAnnotation.go.html to see an example of how to use GetAnnotation API.
// A default retry strategy applies to this operation GetAnnotation()
func (client DataLabelingClient) GetAnnotation(ctx context.Context, request GetAnnotationRequest) (response GetAnnotationResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getAnnotation, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetAnnotationResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetAnnotationResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetAnnotationResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetAnnotationResponse")
	}
	return
}

// getAnnotation implements the OCIOperation interface (enables retrying operations)
func (client DataLabelingClient) getAnnotation(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/annotations/{annotationId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetAnnotationResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/datalabeling-dp/20211001/Annotation/GetAnnotation"
		err = common.PostProcessServiceError(err, "DataLabeling", "GetAnnotation", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetDataset Gets a dataset by identifier.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datalabelingservicedataplane/GetDataset.go.html to see an example of how to use GetDataset API.
// A default retry strategy applies to this operation GetDataset()
func (client DataLabelingClient) GetDataset(ctx context.Context, request GetDatasetRequest) (response GetDatasetResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getDataset, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetDatasetResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetDatasetResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetDatasetResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetDatasetResponse")
	}
	return
}

// getDataset implements the OCIOperation interface (enables retrying operations)
func (client DataLabelingClient) getDataset(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/datasets/{datasetId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetDatasetResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/datalabeling-dp/20211001/Dataset/GetDataset"
		err = common.PostProcessServiceError(err, "DataLabeling", "GetDataset", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetRecord Gets a record.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datalabelingservicedataplane/GetRecord.go.html to see an example of how to use GetRecord API.
// A default retry strategy applies to this operation GetRecord()
func (client DataLabelingClient) GetRecord(ctx context.Context, request GetRecordRequest) (response GetRecordResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getRecord, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetRecordResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetRecordResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetRecordResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetRecordResponse")
	}
	return
}

// getRecord implements the OCIOperation interface (enables retrying operations)
func (client DataLabelingClient) getRecord(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/records/{recordId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetRecordResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/datalabeling-dp/20211001/Record/GetRecord"
		err = common.PostProcessServiceError(err, "DataLabeling", "GetRecord", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetRecordContent Retrieves the content of the record from the dataset source.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datalabelingservicedataplane/GetRecordContent.go.html to see an example of how to use GetRecordContent API.
// A default retry strategy applies to this operation GetRecordContent()
func (client DataLabelingClient) GetRecordContent(ctx context.Context, request GetRecordContentRequest) (response GetRecordContentResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getRecordContent, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetRecordContentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetRecordContentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetRecordContentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetRecordContentResponse")
	}
	return
}

// getRecordContent implements the OCIOperation interface (enables retrying operations)
func (client DataLabelingClient) getRecordContent(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/records/{recordId}/content", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetRecordContentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/datalabeling-dp/20211001/Record/GetRecordContent"
		err = common.PostProcessServiceError(err, "DataLabeling", "GetRecordContent", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetRecordPreviewContent Retrieves the preview of the record content from the dataset source.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datalabelingservicedataplane/GetRecordPreviewContent.go.html to see an example of how to use GetRecordPreviewContent API.
// A default retry strategy applies to this operation GetRecordPreviewContent()
func (client DataLabelingClient) GetRecordPreviewContent(ctx context.Context, request GetRecordPreviewContentRequest) (response GetRecordPreviewContentResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getRecordPreviewContent, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetRecordPreviewContentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetRecordPreviewContentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetRecordPreviewContentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetRecordPreviewContentResponse")
	}
	return
}

// getRecordPreviewContent implements the OCIOperation interface (enables retrying operations)
func (client DataLabelingClient) getRecordPreviewContent(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/records/{recordId}/preview/content", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetRecordPreviewContentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/datalabeling-dp/20211001/Record/GetRecordPreviewContent"
		err = common.PostProcessServiceError(err, "DataLabeling", "GetRecordPreviewContent", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListAnnotations Returns a list of annotations.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datalabelingservicedataplane/ListAnnotations.go.html to see an example of how to use ListAnnotations API.
// A default retry strategy applies to this operation ListAnnotations()
func (client DataLabelingClient) ListAnnotations(ctx context.Context, request ListAnnotationsRequest) (response ListAnnotationsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listAnnotations, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListAnnotationsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListAnnotationsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListAnnotationsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListAnnotationsResponse")
	}
	return
}

// listAnnotations implements the OCIOperation interface (enables retrying operations)
func (client DataLabelingClient) listAnnotations(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/annotations", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListAnnotationsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/datalabeling-dp/20211001/AnnotationCollection/ListAnnotations"
		err = common.PostProcessServiceError(err, "DataLabeling", "ListAnnotations", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListRecords The list of records in the specified compartment.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datalabelingservicedataplane/ListRecords.go.html to see an example of how to use ListRecords API.
// A default retry strategy applies to this operation ListRecords()
func (client DataLabelingClient) ListRecords(ctx context.Context, request ListRecordsRequest) (response ListRecordsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listRecords, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListRecordsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListRecordsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListRecordsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListRecordsResponse")
	}
	return
}

// listRecords implements the OCIOperation interface (enables retrying operations)
func (client DataLabelingClient) listRecords(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/records", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListRecordsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/datalabeling-dp/20211001/RecordCollection/ListRecords"
		err = common.PostProcessServiceError(err, "DataLabeling", "ListRecords", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// SummarizeAnnotationAnalytics Summarize the annotations created for a given dataset.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datalabelingservicedataplane/SummarizeAnnotationAnalytics.go.html to see an example of how to use SummarizeAnnotationAnalytics API.
// A default retry strategy applies to this operation SummarizeAnnotationAnalytics()
func (client DataLabelingClient) SummarizeAnnotationAnalytics(ctx context.Context, request SummarizeAnnotationAnalyticsRequest) (response SummarizeAnnotationAnalyticsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.summarizeAnnotationAnalytics, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = SummarizeAnnotationAnalyticsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = SummarizeAnnotationAnalyticsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(SummarizeAnnotationAnalyticsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into SummarizeAnnotationAnalyticsResponse")
	}
	return
}

// summarizeAnnotationAnalytics implements the OCIOperation interface (enables retrying operations)
func (client DataLabelingClient) summarizeAnnotationAnalytics(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/annotationAnalytics", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response SummarizeAnnotationAnalyticsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/datalabeling-dp/20211001/AnnotationAnalyticsAggregationCollection/SummarizeAnnotationAnalytics"
		err = common.PostProcessServiceError(err, "DataLabeling", "SummarizeAnnotationAnalytics", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// SummarizeRecordAnalytics Summarize the records created for a given dataset.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datalabelingservicedataplane/SummarizeRecordAnalytics.go.html to see an example of how to use SummarizeRecordAnalytics API.
// A default retry strategy applies to this operation SummarizeRecordAnalytics()
func (client DataLabelingClient) SummarizeRecordAnalytics(ctx context.Context, request SummarizeRecordAnalyticsRequest) (response SummarizeRecordAnalyticsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.summarizeRecordAnalytics, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = SummarizeRecordAnalyticsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = SummarizeRecordAnalyticsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(SummarizeRecordAnalyticsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into SummarizeRecordAnalyticsResponse")
	}
	return
}

// summarizeRecordAnalytics implements the OCIOperation interface (enables retrying operations)
func (client DataLabelingClient) summarizeRecordAnalytics(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/recordAnalytics", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response SummarizeRecordAnalyticsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/datalabeling-dp/20211001/RecordAnalyticsAggregationCollection/SummarizeRecordAnalytics"
		err = common.PostProcessServiceError(err, "DataLabeling", "SummarizeRecordAnalytics", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateAnnotation Updates an annotation.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datalabelingservicedataplane/UpdateAnnotation.go.html to see an example of how to use UpdateAnnotation API.
// A default retry strategy applies to this operation UpdateAnnotation()
func (client DataLabelingClient) UpdateAnnotation(ctx context.Context, request UpdateAnnotationRequest) (response UpdateAnnotationResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateAnnotation, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateAnnotationResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateAnnotationResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateAnnotationResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateAnnotationResponse")
	}
	return
}

// updateAnnotation implements the OCIOperation interface (enables retrying operations)
func (client DataLabelingClient) updateAnnotation(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/annotations/{annotationId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateAnnotationResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/datalabeling-dp/20211001/Annotation/UpdateAnnotation"
		err = common.PostProcessServiceError(err, "DataLabeling", "UpdateAnnotation", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateRecord Updates a record.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datalabelingservicedataplane/UpdateRecord.go.html to see an example of how to use UpdateRecord API.
// A default retry strategy applies to this operation UpdateRecord()
func (client DataLabelingClient) UpdateRecord(ctx context.Context, request UpdateRecordRequest) (response UpdateRecordResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateRecord, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateRecordResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateRecordResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateRecordResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateRecordResponse")
	}
	return
}

// updateRecord implements the OCIOperation interface (enables retrying operations)
func (client DataLabelingClient) updateRecord(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/records/{recordId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateRecordResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/datalabeling-dp/20211001/Record/UpdateRecord"
		err = common.PostProcessServiceError(err, "DataLabeling", "UpdateRecord", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}
