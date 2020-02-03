// Copyright (c) 2016, 2018, Oracle and/or its affiliates. All rights reserved.

package common

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEndpoint(t *testing.T) {
	// OC1
	region := StringToRegion("us-phoenix-1")
	endpoint := region.Endpoint("foo")
	assert.Equal(t, "foo.us-phoenix-1.oraclecloud.com", endpoint)

	region = StringToRegion("us-ashburn-1")
	endpoint = region.Endpoint("bar")
	assert.Equal(t, "bar.us-ashburn-1.oraclecloud.com", endpoint)

	region = StringToRegion("ca-toronto-1")
	endpoint = region.Endpoint("bar")
	assert.Equal(t, "bar.ca-toronto-1.oraclecloud.com", endpoint)

	// OC2
	region = StringToRegion("us-langley-1")
	endpoint = region.Endpoint("bar")
	assert.Equal(t, "bar.us-langley-1.oraclegovcloud.com", endpoint)

	// OC3
	region = StringToRegion("us-gov-ashburn-1")
	endpoint = region.Endpoint("bar")
	assert.Equal(t, "bar.us-gov-ashburn-1.oraclegovcloud.com", endpoint)

	// OC4
	region = StringToRegion("uk-gov-london-1")
	endpoint = region.Endpoint("bar")
	assert.Equal(t, "bar.uk-gov-london-1.oraclegovcloud.uk", endpoint)
}

func TestEndpointForTemplate(t *testing.T) {
	type testData struct {
		region           Region
		service          string
		endpointTemplate string
		expected         string
	}
	testDataSet := []testData{
		{
			// template with region
			region:           StringToRegion("us-phoenix-1"),
			service:          "test",
			endpointTemplate: "https://foo.{region}.bar.com",
			expected:         "https://foo.us-phoenix-1.bar.com",
		},
		{
			// template with second level domain
			region:           StringToRegion("us-phoenix-1"),
			service:          "test",
			endpointTemplate: "https://foo.region.{secondLevelDomain}",
			expected:         "https://foo.region.oraclecloud.com",
		},
		{
			// template with second level domain
			region:           StringToRegion("ap-sydney-1"),
			service:          "test",
			endpointTemplate: "https://foo.region.{secondLevelDomain}",
			expected:         "https://foo.region.oraclecloud.com",
		},
		{
			// template with everything for OC2
			region:           StringToRegion("us-langley-1"),
			service:          "test",
			endpointTemplate: "https://test.{region}.{secondLevelDomain}",
			expected:         "https://test.us-langley-1.oraclegovcloud.com",
		},
		{
			// template with everything for OC3
			region:           StringToRegion("us-gov-ashburn-1"),
			service:          "test",
			endpointTemplate: "https://test.{region}.{secondLevelDomain}",
			expected:         "https://test.us-gov-ashburn-1.oraclegovcloud.com",
		},
		{
			// template with everything for OC4
			region:           StringToRegion("uk-gov-london-1"),
			service:          "test",
			endpointTemplate: "https://test.{region}.{secondLevelDomain}",
			expected:         "https://test.uk-gov-london-1.oraclegovcloud.uk",
		},
	}

	for _, testData := range testDataSet {
		endpoint := testData.region.EndpointForTemplate(testData.service, testData.endpointTemplate)
		assert.Equal(t, testData.expected, endpoint)
	}
}

func TestStringToRegion(t *testing.T) {
	region := StringToRegion("yyz")
	assert.Equal(t, RegionCAToronto1, region)

	region = StringToRegion("nrt")
	assert.Equal(t, RegionAPTokyo1, region)

	region = StringToRegion("kix")
	assert.Equal(t, RegionAPOsaka1, region)

	region = StringToRegion("gru")
	assert.Equal(t, RegionSASaopaulo1, region)

	region = StringToRegion("syd")
	assert.Equal(t, RegionAPSydney1, region)

	region = StringToRegion("mel")
	assert.Equal(t, RegionAPMelbourne1, region)

	region = StringToRegion("uk-gov-london-1")
	assert.Equal(t, RegionUKGovLondon1, region)

	region = StringToRegion("ltn")
	assert.Equal(t, RegionUKGovLondon1, region)
}
