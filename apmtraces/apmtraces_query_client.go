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
	"github.com/oracle/oci-go-sdk/v37/common"
	"github.com/oracle/oci-go-sdk/v37/common/auth"
	"net/http"
)

//QueryClient a client for Query
type QueryClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewQueryClientWithConfigurationProvider Creates a new default Query client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewQueryClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client QueryClient, err error) {
	provider, err := auth.GetGenericConfigurationProvider(configProvider)
	if err != nil {
		return client, err
	}
	baseClient, e := common.NewClientWithConfig(provider)
	if e != nil {
		return client, e
	}
	return newQueryClientFromBaseClient(baseClient, provider)
}

// NewQueryClientWithOboToken Creates a new default Query client with the given configuration provider.
// The obotoken will be added to default headers and signed; the configuration provider will be used for the signer
//  as well as reading the region
func NewQueryClientWithOboToken(configProvider common.ConfigurationProvider, oboToken string) (client QueryClient, err error) {
	baseClient, err := common.NewClientWithOboToken(configProvider, oboToken)
	if err != nil {
		return client, err
	}

	return newQueryClientFromBaseClient(baseClient, configProvider)
}

func newQueryClientFromBaseClient(baseClient common.BaseClient, configProvider common.ConfigurationProvider) (client QueryClient, err error) {
	client = QueryClient{BaseClient: baseClient}
	client.BasePath = "20200630"
	err = client.setConfigurationProvider(configProvider)
	return
}

// SetRegion overrides the region of this client.
func (client *QueryClient) SetRegion(region string) {
	client.Host = common.StringToRegion(region).EndpointForTemplate("apmtraces", "https://apm-trace.{region}.oci.{secondLevelDomain}")
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *QueryClient) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
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
func (client *QueryClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// ListQuickPicks Returns a list of predefined quick pick queries intended to assist the user
// to choose a query to run.  There is no sorting applied on the results.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/apmtraces/ListQuickPicks.go.html to see an example of how to use ListQuickPicks API.
func (client QueryClient) ListQuickPicks(ctx context.Context, request ListQuickPicksRequest) (response ListQuickPicksResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listQuickPicks, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListQuickPicksResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListQuickPicksResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListQuickPicksResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListQuickPicksResponse")
	}
	return
}

// listQuickPicks implements the OCIOperation interface (enables retrying operations)
func (client QueryClient) listQuickPicks(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/queries/quickPicks")
	if err != nil {
		return nil, err
	}

	var response ListQuickPicksResponse
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

// Query Given a query, constructed according to the APM Defined Query Syntax, retrieves the results - selected attributes,
// and aggregations of the queried entity.  Query Results are filtered by the filter criteria specified in the where clause.
// Further query results are grouped by the attributes specified in the group by clause.  Finally,
// ordering (asc/desc) is done by the specified attributes in the order by clause.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/apmtraces/Query.go.html to see an example of how to use Query API.
func (client QueryClient) Query(ctx context.Context, request QueryRequest) (response QueryResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.query, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = QueryResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = QueryResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(QueryResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into QueryResponse")
	}
	return
}

// query implements the OCIOperation interface (enables retrying operations)
func (client QueryClient) query(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/queries/actions/runQuery")
	if err != nil {
		return nil, err
	}

	var response QueryResponse
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

// QueryOld THIS API ENDPOINT WILL BE DEPRECATED AND INSTEAD /queries/actions/runQuery as defined below WILL BE USED GOING FORWARD.  THIS EXISTS JUST
// AS A TEMPORARY PLACEHOLDER SO AS TO BE BACKWARDS COMPATIBLE WITH THE UI BETWEEN RELEASE CYCLES.
// Given a query, constructed according to the APM Defined Query Syntax, retrieves the results - selected attributes,
// and aggregations of the queried entity.  Query Results are filtered by the filter criteria specified in the where clause.
// Further query results are grouped by the attributes specified in the group by clause.  Finally,
// ordering (asc/desc) is done by the specified attributes in the order by clause.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/apmtraces/QueryOld.go.html to see an example of how to use QueryOld API.
func (client QueryClient) QueryOld(ctx context.Context, request QueryOldRequest) (response QueryOldResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.queryOld, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = QueryOldResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = QueryOldResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(QueryOldResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into QueryOldResponse")
	}
	return
}

// queryOld implements the OCIOperation interface (enables retrying operations)
func (client QueryClient) queryOld(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/queries/action/runQuery")
	if err != nil {
		return nil, err
	}

	var response QueryOldResponse
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
