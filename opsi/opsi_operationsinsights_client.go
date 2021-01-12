// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Operations Insights API
//
// Use the Operations Insights API to perform data extraction operations to obtain database
// resource utilization, performance statistics, and reference information. For more information,
// see About Oracle Cloud Infrastructure Operations Insights (https://docs.cloud.oracle.com/en-us/iaas/operations-insights/doc/operations-insights.html).
//

package opsi

import (
	"context"
	"fmt"
	"github.com/oracle/oci-go-sdk/v32/common"
	"github.com/oracle/oci-go-sdk/v32/common/auth"
	"net/http"
)

//OperationsInsightsClient a client for OperationsInsights
type OperationsInsightsClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewOperationsInsightsClientWithConfigurationProvider Creates a new default OperationsInsights client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewOperationsInsightsClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client OperationsInsightsClient, err error) {
	provider, err := auth.GetGenericConfigurationProvider(configProvider)
	if err != nil {
		return client, err
	}
	baseClient, e := common.NewClientWithConfig(provider)
	if e != nil {
		return client, e
	}
	return newOperationsInsightsClientFromBaseClient(baseClient, provider)
}

// NewOperationsInsightsClientWithOboToken Creates a new default OperationsInsights client with the given configuration provider.
// The obotoken will be added to default headers and signed; the configuration provider will be used for the signer
//  as well as reading the region
func NewOperationsInsightsClientWithOboToken(configProvider common.ConfigurationProvider, oboToken string) (client OperationsInsightsClient, err error) {
	baseClient, err := common.NewClientWithOboToken(configProvider, oboToken)
	if err != nil {
		return client, err
	}

	return newOperationsInsightsClientFromBaseClient(baseClient, configProvider)
}

func newOperationsInsightsClientFromBaseClient(baseClient common.BaseClient, configProvider common.ConfigurationProvider) (client OperationsInsightsClient, err error) {
	client = OperationsInsightsClient{BaseClient: baseClient}
	client.BasePath = "20200630"
	err = client.setConfigurationProvider(configProvider)
	return
}

// SetRegion overrides the region of this client.
func (client *OperationsInsightsClient) SetRegion(region string) {
	client.Host = common.StringToRegion(region).EndpointForTemplate("opsi", "https://operationsinsights.{region}.oci.{secondLevelDomain}")
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *OperationsInsightsClient) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
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
func (client *OperationsInsightsClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// IngestSqlBucket The sqlbucket endpoint takes in a JSON payload, persists it in Operations Insights ingest pipeline.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/IngestSqlBucket.go.html to see an example of how to use IngestSqlBucket API.
func (client OperationsInsightsClient) IngestSqlBucket(ctx context.Context, request IngestSqlBucketRequest) (response IngestSqlBucketResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.ingestSqlBucket, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = IngestSqlBucketResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = IngestSqlBucketResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(IngestSqlBucketResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into IngestSqlBucketResponse")
	}
	return
}

// ingestSqlBucket implements the OCIOperation interface (enables retrying operations)
func (client OperationsInsightsClient) ingestSqlBucket(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/databaseInsights/actions/ingestSqlBucket")
	if err != nil {
		return nil, err
	}

	var response IngestSqlBucketResponse
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

// IngestSqlPlanLines The SqlPlanLines endpoint takes in a JSON payload, persists it in Operation Insights ingest pipeline.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/IngestSqlPlanLines.go.html to see an example of how to use IngestSqlPlanLines API.
func (client OperationsInsightsClient) IngestSqlPlanLines(ctx context.Context, request IngestSqlPlanLinesRequest) (response IngestSqlPlanLinesResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.ingestSqlPlanLines, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = IngestSqlPlanLinesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = IngestSqlPlanLinesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(IngestSqlPlanLinesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into IngestSqlPlanLinesResponse")
	}
	return
}

// ingestSqlPlanLines implements the OCIOperation interface (enables retrying operations)
func (client OperationsInsightsClient) ingestSqlPlanLines(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/databaseInsights/actions/ingestSqlPlanLines")
	if err != nil {
		return nil, err
	}

	var response IngestSqlPlanLinesResponse
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

// IngestSqlText The SqlText endpoint takes in a JSON payload, persists it in Operation Insights ingest pipeline.
// Disclaimer: SQL text being uploaded explicitly via APIs is not masked. Any sensitive literals contained in the sqlFullText column should be masked prior to ingestion.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/IngestSqlText.go.html to see an example of how to use IngestSqlText API.
func (client OperationsInsightsClient) IngestSqlText(ctx context.Context, request IngestSqlTextRequest) (response IngestSqlTextResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.ingestSqlText, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = IngestSqlTextResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = IngestSqlTextResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(IngestSqlTextResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into IngestSqlTextResponse")
	}
	return
}

// ingestSqlText implements the OCIOperation interface (enables retrying operations)
func (client OperationsInsightsClient) ingestSqlText(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/databaseInsights/actions/ingestSqlText")
	if err != nil {
		return nil, err
	}

	var response IngestSqlTextResponse
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

// ListDatabaseInsights Lists database insight resources
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/ListDatabaseInsights.go.html to see an example of how to use ListDatabaseInsights API.
func (client OperationsInsightsClient) ListDatabaseInsights(ctx context.Context, request ListDatabaseInsightsRequest) (response ListDatabaseInsightsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listDatabaseInsights, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListDatabaseInsightsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListDatabaseInsightsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListDatabaseInsightsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListDatabaseInsightsResponse")
	}
	return
}

// listDatabaseInsights implements the OCIOperation interface (enables retrying operations)
func (client OperationsInsightsClient) listDatabaseInsights(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/databaseInsights")
	if err != nil {
		return nil, err
	}

	var response ListDatabaseInsightsResponse
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

// ListSqlPlans Query SQL Warehouse to list the plan xml for a given SQL execution plan. This returns a SqlPlanCollection object, but is currently limited to a single plan.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/ListSqlPlans.go.html to see an example of how to use ListSqlPlans API.
func (client OperationsInsightsClient) ListSqlPlans(ctx context.Context, request ListSqlPlansRequest) (response ListSqlPlansResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listSqlPlans, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListSqlPlansResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListSqlPlansResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListSqlPlansResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListSqlPlansResponse")
	}
	return
}

// listSqlPlans implements the OCIOperation interface (enables retrying operations)
func (client OperationsInsightsClient) listSqlPlans(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/databaseInsights/sqlPlans")
	if err != nil {
		return nil, err
	}

	var response ListSqlPlansResponse
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

// ListSqlSearches Search SQL by SQL Identifier across databases and get the SQL Text and the details of the databases executing the SQL for a given time period.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/ListSqlSearches.go.html to see an example of how to use ListSqlSearches API.
func (client OperationsInsightsClient) ListSqlSearches(ctx context.Context, request ListSqlSearchesRequest) (response ListSqlSearchesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listSqlSearches, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListSqlSearchesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListSqlSearchesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListSqlSearchesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListSqlSearchesResponse")
	}
	return
}

// listSqlSearches implements the OCIOperation interface (enables retrying operations)
func (client OperationsInsightsClient) listSqlSearches(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/databaseInsights/sqlSearches")
	if err != nil {
		return nil, err
	}

	var response ListSqlSearchesResponse
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

// ListSqlTexts Query SQL Warehouse to get the full SQL Text for a SQL.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/ListSqlTexts.go.html to see an example of how to use ListSqlTexts API.
func (client OperationsInsightsClient) ListSqlTexts(ctx context.Context, request ListSqlTextsRequest) (response ListSqlTextsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listSqlTexts, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListSqlTextsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListSqlTextsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListSqlTextsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListSqlTextsResponse")
	}
	return
}

// listSqlTexts implements the OCIOperation interface (enables retrying operations)
func (client OperationsInsightsClient) listSqlTexts(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/databaseInsights/sqlTexts")
	if err != nil {
		return nil, err
	}

	var response ListSqlTextsResponse
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

// SummarizeDatabaseInsightResourceCapacityTrend Returns response with time series data (endTimestamp, capacity, baseCapacity) for the time period specified.
// The maximum time range for analysis is 2 years, hence this is intentionally not paginated.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/SummarizeDatabaseInsightResourceCapacityTrend.go.html to see an example of how to use SummarizeDatabaseInsightResourceCapacityTrend API.
func (client OperationsInsightsClient) SummarizeDatabaseInsightResourceCapacityTrend(ctx context.Context, request SummarizeDatabaseInsightResourceCapacityTrendRequest) (response SummarizeDatabaseInsightResourceCapacityTrendResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.summarizeDatabaseInsightResourceCapacityTrend, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = SummarizeDatabaseInsightResourceCapacityTrendResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = SummarizeDatabaseInsightResourceCapacityTrendResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(SummarizeDatabaseInsightResourceCapacityTrendResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into SummarizeDatabaseInsightResourceCapacityTrendResponse")
	}
	return
}

// summarizeDatabaseInsightResourceCapacityTrend implements the OCIOperation interface (enables retrying operations)
func (client OperationsInsightsClient) summarizeDatabaseInsightResourceCapacityTrend(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/databaseInsights/resourceCapacityTrend")
	if err != nil {
		return nil, err
	}

	var response SummarizeDatabaseInsightResourceCapacityTrendResponse
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

// SummarizeDatabaseInsightResourceForecastTrend Get Forecast predictions for CPU and Storage resources since a time in the past.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/SummarizeDatabaseInsightResourceForecastTrend.go.html to see an example of how to use SummarizeDatabaseInsightResourceForecastTrend API.
func (client OperationsInsightsClient) SummarizeDatabaseInsightResourceForecastTrend(ctx context.Context, request SummarizeDatabaseInsightResourceForecastTrendRequest) (response SummarizeDatabaseInsightResourceForecastTrendResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.summarizeDatabaseInsightResourceForecastTrend, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = SummarizeDatabaseInsightResourceForecastTrendResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = SummarizeDatabaseInsightResourceForecastTrendResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(SummarizeDatabaseInsightResourceForecastTrendResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into SummarizeDatabaseInsightResourceForecastTrendResponse")
	}
	return
}

// summarizeDatabaseInsightResourceForecastTrend implements the OCIOperation interface (enables retrying operations)
func (client OperationsInsightsClient) summarizeDatabaseInsightResourceForecastTrend(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/databaseInsights/resourceForecastTrend")
	if err != nil {
		return nil, err
	}

	var response SummarizeDatabaseInsightResourceForecastTrendResponse
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

// SummarizeDatabaseInsightResourceStatistics Lists the Resource statistics (usage,capacity, usage change percent, utilization percent, base capacity, isAutoScalingEnabled) for each database filtered by utilization level
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/SummarizeDatabaseInsightResourceStatistics.go.html to see an example of how to use SummarizeDatabaseInsightResourceStatistics API.
func (client OperationsInsightsClient) SummarizeDatabaseInsightResourceStatistics(ctx context.Context, request SummarizeDatabaseInsightResourceStatisticsRequest) (response SummarizeDatabaseInsightResourceStatisticsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.summarizeDatabaseInsightResourceStatistics, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = SummarizeDatabaseInsightResourceStatisticsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = SummarizeDatabaseInsightResourceStatisticsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(SummarizeDatabaseInsightResourceStatisticsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into SummarizeDatabaseInsightResourceStatisticsResponse")
	}
	return
}

// summarizeDatabaseInsightResourceStatistics implements the OCIOperation interface (enables retrying operations)
func (client OperationsInsightsClient) summarizeDatabaseInsightResourceStatistics(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/databaseInsights/resourceStatistics")
	if err != nil {
		return nil, err
	}

	var response SummarizeDatabaseInsightResourceStatisticsResponse
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

// SummarizeDatabaseInsightResourceUsage A cumulative distribution function is used to rank the usage data points per database within the specified time period.
// For each database, the minimum data point with a ranking > the percentile value is included in the summation.
// Linear regression functions are used to calculate the usage change percentage.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/SummarizeDatabaseInsightResourceUsage.go.html to see an example of how to use SummarizeDatabaseInsightResourceUsage API.
func (client OperationsInsightsClient) SummarizeDatabaseInsightResourceUsage(ctx context.Context, request SummarizeDatabaseInsightResourceUsageRequest) (response SummarizeDatabaseInsightResourceUsageResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.summarizeDatabaseInsightResourceUsage, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = SummarizeDatabaseInsightResourceUsageResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = SummarizeDatabaseInsightResourceUsageResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(SummarizeDatabaseInsightResourceUsageResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into SummarizeDatabaseInsightResourceUsageResponse")
	}
	return
}

// summarizeDatabaseInsightResourceUsage implements the OCIOperation interface (enables retrying operations)
func (client OperationsInsightsClient) summarizeDatabaseInsightResourceUsage(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/databaseInsights/resourceUsageSummary")
	if err != nil {
		return nil, err
	}

	var response SummarizeDatabaseInsightResourceUsageResponse
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

// SummarizeDatabaseInsightResourceUsageTrend Returns response with time series data (endTimestamp, usage, capacity) for the time period specified.
// The maximum time range for analysis is 2 years, hence this is intentionally not paginated.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/SummarizeDatabaseInsightResourceUsageTrend.go.html to see an example of how to use SummarizeDatabaseInsightResourceUsageTrend API.
func (client OperationsInsightsClient) SummarizeDatabaseInsightResourceUsageTrend(ctx context.Context, request SummarizeDatabaseInsightResourceUsageTrendRequest) (response SummarizeDatabaseInsightResourceUsageTrendResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.summarizeDatabaseInsightResourceUsageTrend, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = SummarizeDatabaseInsightResourceUsageTrendResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = SummarizeDatabaseInsightResourceUsageTrendResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(SummarizeDatabaseInsightResourceUsageTrendResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into SummarizeDatabaseInsightResourceUsageTrendResponse")
	}
	return
}

// summarizeDatabaseInsightResourceUsageTrend implements the OCIOperation interface (enables retrying operations)
func (client OperationsInsightsClient) summarizeDatabaseInsightResourceUsageTrend(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/databaseInsights/resourceUsageTrend")
	if err != nil {
		return nil, err
	}

	var response SummarizeDatabaseInsightResourceUsageTrendResponse
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

// SummarizeDatabaseInsightResourceUtilizationInsight Gets resources with current utilization (high and low) and projected utilization (high and low) for a resource type over specified time period.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/SummarizeDatabaseInsightResourceUtilizationInsight.go.html to see an example of how to use SummarizeDatabaseInsightResourceUtilizationInsight API.
func (client OperationsInsightsClient) SummarizeDatabaseInsightResourceUtilizationInsight(ctx context.Context, request SummarizeDatabaseInsightResourceUtilizationInsightRequest) (response SummarizeDatabaseInsightResourceUtilizationInsightResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.summarizeDatabaseInsightResourceUtilizationInsight, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = SummarizeDatabaseInsightResourceUtilizationInsightResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = SummarizeDatabaseInsightResourceUtilizationInsightResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(SummarizeDatabaseInsightResourceUtilizationInsightResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into SummarizeDatabaseInsightResourceUtilizationInsightResponse")
	}
	return
}

// summarizeDatabaseInsightResourceUtilizationInsight implements the OCIOperation interface (enables retrying operations)
func (client OperationsInsightsClient) summarizeDatabaseInsightResourceUtilizationInsight(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/databaseInsights/resourceUtilizationInsight")
	if err != nil {
		return nil, err
	}

	var response SummarizeDatabaseInsightResourceUtilizationInsightResponse
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

// SummarizeSqlInsights Query SQL Warehouse to get the performance insights for SQLs taking greater than X% database time for a given time period across the given databases or database types.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/SummarizeSqlInsights.go.html to see an example of how to use SummarizeSqlInsights API.
func (client OperationsInsightsClient) SummarizeSqlInsights(ctx context.Context, request SummarizeSqlInsightsRequest) (response SummarizeSqlInsightsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.summarizeSqlInsights, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = SummarizeSqlInsightsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = SummarizeSqlInsightsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(SummarizeSqlInsightsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into SummarizeSqlInsightsResponse")
	}
	return
}

// summarizeSqlInsights implements the OCIOperation interface (enables retrying operations)
func (client OperationsInsightsClient) summarizeSqlInsights(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/databaseInsights/sqlInsights")
	if err != nil {
		return nil, err
	}

	var response SummarizeSqlInsightsResponse
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

// SummarizeSqlPlanInsights Query SQL Warehouse to get the performance insights on the execution plans for a given SQL for a given time period.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/SummarizeSqlPlanInsights.go.html to see an example of how to use SummarizeSqlPlanInsights API.
func (client OperationsInsightsClient) SummarizeSqlPlanInsights(ctx context.Context, request SummarizeSqlPlanInsightsRequest) (response SummarizeSqlPlanInsightsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.summarizeSqlPlanInsights, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = SummarizeSqlPlanInsightsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = SummarizeSqlPlanInsightsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(SummarizeSqlPlanInsightsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into SummarizeSqlPlanInsightsResponse")
	}
	return
}

// summarizeSqlPlanInsights implements the OCIOperation interface (enables retrying operations)
func (client OperationsInsightsClient) summarizeSqlPlanInsights(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/databaseInsights/sqlPlanInsights")
	if err != nil {
		return nil, err
	}

	var response SummarizeSqlPlanInsightsResponse
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

// SummarizeSqlResponseTimeDistributions Query SQL Warehouse to summarize the response time distribution of query executions for a given SQL for a given time period.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/SummarizeSqlResponseTimeDistributions.go.html to see an example of how to use SummarizeSqlResponseTimeDistributions API.
func (client OperationsInsightsClient) SummarizeSqlResponseTimeDistributions(ctx context.Context, request SummarizeSqlResponseTimeDistributionsRequest) (response SummarizeSqlResponseTimeDistributionsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.summarizeSqlResponseTimeDistributions, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = SummarizeSqlResponseTimeDistributionsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = SummarizeSqlResponseTimeDistributionsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(SummarizeSqlResponseTimeDistributionsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into SummarizeSqlResponseTimeDistributionsResponse")
	}
	return
}

// summarizeSqlResponseTimeDistributions implements the OCIOperation interface (enables retrying operations)
func (client OperationsInsightsClient) summarizeSqlResponseTimeDistributions(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/databaseInsights/sqlResponseTimeDistributions")
	if err != nil {
		return nil, err
	}

	var response SummarizeSqlResponseTimeDistributionsResponse
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

// SummarizeSqlStatistics Query SQL Warehouse to get the performance statistics for SQLs taking greater than X% database time for a given time period across the given databases or database types.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/SummarizeSqlStatistics.go.html to see an example of how to use SummarizeSqlStatistics API.
func (client OperationsInsightsClient) SummarizeSqlStatistics(ctx context.Context, request SummarizeSqlStatisticsRequest) (response SummarizeSqlStatisticsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.summarizeSqlStatistics, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = SummarizeSqlStatisticsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = SummarizeSqlStatisticsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(SummarizeSqlStatisticsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into SummarizeSqlStatisticsResponse")
	}
	return
}

// summarizeSqlStatistics implements the OCIOperation interface (enables retrying operations)
func (client OperationsInsightsClient) summarizeSqlStatistics(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/databaseInsights/sqlStatistics")
	if err != nil {
		return nil, err
	}

	var response SummarizeSqlStatisticsResponse
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

// SummarizeSqlStatisticsTimeSeries Query SQL Warehouse to get the performance statistics time series for a given SQL across given databases for a given time period.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/SummarizeSqlStatisticsTimeSeries.go.html to see an example of how to use SummarizeSqlStatisticsTimeSeries API.
func (client OperationsInsightsClient) SummarizeSqlStatisticsTimeSeries(ctx context.Context, request SummarizeSqlStatisticsTimeSeriesRequest) (response SummarizeSqlStatisticsTimeSeriesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.summarizeSqlStatisticsTimeSeries, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = SummarizeSqlStatisticsTimeSeriesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = SummarizeSqlStatisticsTimeSeriesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(SummarizeSqlStatisticsTimeSeriesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into SummarizeSqlStatisticsTimeSeriesResponse")
	}
	return
}

// summarizeSqlStatisticsTimeSeries implements the OCIOperation interface (enables retrying operations)
func (client OperationsInsightsClient) summarizeSqlStatisticsTimeSeries(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/databaseInsights/sqlStatisticsTimeSeries")
	if err != nil {
		return nil, err
	}

	var response SummarizeSqlStatisticsTimeSeriesResponse
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

// SummarizeSqlStatisticsTimeSeriesByPlan Query SQL Warehouse to get the performance statistics time series for a given SQL by execution plans for a given time period.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/SummarizeSqlStatisticsTimeSeriesByPlan.go.html to see an example of how to use SummarizeSqlStatisticsTimeSeriesByPlan API.
func (client OperationsInsightsClient) SummarizeSqlStatisticsTimeSeriesByPlan(ctx context.Context, request SummarizeSqlStatisticsTimeSeriesByPlanRequest) (response SummarizeSqlStatisticsTimeSeriesByPlanResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.summarizeSqlStatisticsTimeSeriesByPlan, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = SummarizeSqlStatisticsTimeSeriesByPlanResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = SummarizeSqlStatisticsTimeSeriesByPlanResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(SummarizeSqlStatisticsTimeSeriesByPlanResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into SummarizeSqlStatisticsTimeSeriesByPlanResponse")
	}
	return
}

// summarizeSqlStatisticsTimeSeriesByPlan implements the OCIOperation interface (enables retrying operations)
func (client OperationsInsightsClient) summarizeSqlStatisticsTimeSeriesByPlan(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/databaseInsights/sqlStatisticsTimeSeriesByPlan")
	if err != nil {
		return nil, err
	}

	var response SummarizeSqlStatisticsTimeSeriesByPlanResponse
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
