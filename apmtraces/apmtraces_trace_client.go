// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Apm Traces API
//
// API for APM Trace service. Use this API to query the Traces and associated Spans.
//

package apmtraces

import (
	"context"
	"fmt"
	"github.com/oracle/oci-go-sdk/v40/common"
	"github.com/oracle/oci-go-sdk/v40/common/auth"
	"net/http"
)

//TraceClient a client for Trace
type TraceClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewTraceClientWithConfigurationProvider Creates a new default Trace client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewTraceClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client TraceClient, err error) {
	provider, err := auth.GetGenericConfigurationProvider(configProvider)
	if err != nil {
		return client, err
	}
	baseClient, e := common.NewClientWithConfig(provider)
	if e != nil {
		return client, e
	}
	return newTraceClientFromBaseClient(baseClient, provider)
}

// NewTraceClientWithOboToken Creates a new default Trace client with the given configuration provider.
// The obotoken will be added to default headers and signed; the configuration provider will be used for the signer
//  as well as reading the region
func NewTraceClientWithOboToken(configProvider common.ConfigurationProvider, oboToken string) (client TraceClient, err error) {
	baseClient, err := common.NewClientWithOboToken(configProvider, oboToken)
	if err != nil {
		return client, err
	}

	return newTraceClientFromBaseClient(baseClient, configProvider)
}

func newTraceClientFromBaseClient(baseClient common.BaseClient, configProvider common.ConfigurationProvider) (client TraceClient, err error) {
	client = TraceClient{BaseClient: baseClient}
	client.BasePath = "20200630"
	err = client.setConfigurationProvider(configProvider)
	return
}

// SetRegion overrides the region of this client.
func (client *TraceClient) SetRegion(region string) {
	client.Host = common.StringToRegion(region).EndpointForTemplate("apmtraces", "https://apm-trace.{region}.oci.{secondLevelDomain}")
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *TraceClient) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
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
func (client *TraceClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// GetSpan Get the span details identified by spanId
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/apmtraces/GetSpan.go.html to see an example of how to use GetSpan API.
func (client TraceClient) GetSpan(ctx context.Context, request GetSpanRequest) (response GetSpanResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getSpan, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetSpanResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetSpanResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetSpanResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetSpanResponse")
	}
	return
}

// getSpan implements the OCIOperation interface (enables retrying operations)
func (client TraceClient) getSpan(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/spans/{traceKey}/{spanKey}", binaryReqBody)
	if err != nil {
		return nil, err
	}

	var response GetSpanResponse
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

// GetTrace Get the trace details identified by traceId
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/apmtraces/GetTrace.go.html to see an example of how to use GetTrace API.
func (client TraceClient) GetTrace(ctx context.Context, request GetTraceRequest) (response GetTraceResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getTrace, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetTraceResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetTraceResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetTraceResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetTraceResponse")
	}
	return
}

// getTrace implements the OCIOperation interface (enables retrying operations)
func (client TraceClient) getTrace(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/traces/{traceKey}", binaryReqBody)
	if err != nil {
		return nil, err
	}

	var response GetTraceResponse
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
