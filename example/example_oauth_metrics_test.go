// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.

// Example code for using the OAuth2ConfigurationProvider to publish T2 Metrics. For more information on using OAuth2 type authorization, please refer to https://confluence.oraclecorp.com/confluence/display/IDOCI/Onboard+to+OAuth+token

package example

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/oracle/oci-go-sdk/v65/common"
	"github.com/oracle/oci-go-sdk/v65/common/auth"
	"github.com/oracle/oci-go-sdk/v65/example/helpers"
	"github.com/oracle/oci-go-sdk/v65/monitoring"
)

func Example_t2PostMetricsWithOAuth() {

	// Set environment variables as needed. OCI_SDK_AUTH_CLIENT_REGION_URL will be used for retrieving auth tokens
	// for both Instance principals (if used) and OAuth
	os.Setenv("OCI_SDK_AUTH_CLIENT_REGION_URL", "https://auth.r1.oracleiaas.com")
	os.Setenv("OCI_DEFAULT_CLIENT_CERTS_PATH", "/etc/pki/ca-trust/extracted/pem/tls-ca-bundle.pem")

	// The OAuth2ConfigurationProvider requires an existing config provider of type
	// Instance Principals, Resource Principals, or Service Principals
	instanceProvider, err := auth.InstancePrincipalConfigurationProvider()
	helpers.FatalIfError(err)

	_, keyErr := instanceProvider.KeyID()
	helpers.FatalIfError(keyErr)

	targetCompartment := "ocid1.compartment.region1..aaaaaaaasfm6ym4xn4o7bk27vwt6mswiibyqwqxo654hkf4l27whzlhcbvka"
	scope := fmt.Sprintf("monitoring_write_urn-%s:instance", targetCompartment)
	oAuthProvider, err := auth.NewOAuth2ConfigurationProvider(instanceProvider, scope, targetCompartment)
	helpers.FatalIfError(err)
	_, keyErr = oAuthProvider.KeyID()
	helpers.FatalIfError(keyErr)
	fmt.Println("Created OAuth Provider")

	// Set desired metric data. This example posts a single MetricData object with placeholder field values
	metricValue := float64(15)
	metricCount := 1
	namespace := "exampleNamespace"
	resourceGroup := "exampleResourceGroup"
	metricName := "exampleMetricName"
	dimensionResource := "exampleDimensionResource"
	ctx := context.Background()
	postMetricDetails := monitoring.PostMetricDataDetails{
		MetricData: []monitoring.MetricDataDetails{
			{
				Namespace:     common.String(namespace),
				ResourceGroup: common.String(resourceGroup),
				CompartmentId: common.String(targetCompartment),
				Name:          common.String(metricName),
				Dimensions:    map[string]string{"resourceId": dimensionResource},
				Metadata:      map[string]string{"unit": "count"},
				Datapoints:    []monitoring.Datapoint{{Timestamp: &common.SDKTime{time.Now()}, Value: &metricValue, Count: &metricCount}},
			},
		},
	}

	postMetricRequest := monitoring.PostMetricDataRequest{
		PostMetricDataDetails: postMetricDetails,
	}

	// Set the endpoint for the Monitoring Client to use. Note that the endpoing used for PostMetricsData is different
	// then the endpoint for other Monitoring Service APIs. Be sure to replace `telemetry` with `telemetry-ingestion` in the endpoint
	metricsClient, err := monitoring.NewMonitoringClientWithConfigurationProvider(oAuthProvider)
	metricsClient.Host = "https://telemetry-ingestion.r1.oracleiaas.com"
	for i := 1; i <= 5; i++ {
		newTimestamp := time.Now().Add(time.Duration(2*i) * time.Second)
		postMetricRequest.PostMetricDataDetails.MetricData[0].Datapoints[0].Timestamp = &common.SDKTime{newTimestamp}
		_, err := metricsClient.PostMetricData(ctx, postMetricRequest)
		helpers.FatalIfError(err)
		fmt.Printf("Posted metric #%v\n", i)
	}
	fmt.Println("Finished posting metrics")

	// Output:
	// Created OAuth Provider
	// Posted metric #1
	// Posted metric #2
	// Posted metric #3
	// Posted metric #4
	// Posted metric #5
	// Finished posting metrics
}
