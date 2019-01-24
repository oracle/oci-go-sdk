// Copyright (c) 2016, 2018, Oracle and/or its affiliates. All rights reserved.

package common

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEndpoint(t *testing.T) {
	region := StringToRegion("us-phoenix-1")
	endpoint := region.Endpoint("foo")
	assert.Equal(t, "foo.us-phoenix-1.oraclecloud.com", endpoint)

	region = StringToRegion("us-ashburn-1")
	endpoint = region.Endpoint("bar")
	assert.Equal(t, "bar.us-ashburn-1.oraclecloud.com", endpoint)

	region = StringToRegion("ca-toronto-1")
	endpoint = region.Endpoint("bar")
	assert.Equal(t, "bar.ca-toronto-1.oraclecloud.com", endpoint)
}

func TestEndpointForTemplate(t *testing.T) {
	region := StringToRegion("us-phoenix-1")

	type testData struct {
		service          string
		endpointTemplate string
		expected         string
	}
	testDataSet := []testData{
		{
			// template with service prefix
			service:          "foo",
			endpointTemplate: "https://{serviceEndpointPrefix}.region.bar.com",
			expected:         "https://foo.region.bar.com",
		},
		{
			// template with region
			service:          "test",
			endpointTemplate: "https://foo.{region}.bar.com",
			expected:         "https://foo.us-phoenix-1.bar.com",
		},
		{
			// template with second level domain
			service:          "test",
			endpointTemplate: "https://foo.region.{secondLevelDomain}",
			expected:         "https://foo.region.oraclecloud.com",
		},
		{
			// template with everything
			service:          "test",
			endpointTemplate: "https://{serviceEndpointPrefix}.{region}.{secondLevelDomain}",
			expected:         "https://test.us-phoenix-1.oraclecloud.com",
		},
	}

	for _, testData := range testDataSet {
		endpoint := region.EndpointForTemplate(testData.service, testData.endpointTemplate)
		assert.Equal(t, testData.expected, endpoint)
	}
}
