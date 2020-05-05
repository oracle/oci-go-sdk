// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.

package common

import (
	"io/ioutil"
	"os"
	"path"
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
			// template with second level domain
			region:           StringToRegion("me-jeddah-1"),
			service:          "test",
			endpointTemplate: "https://foo.region.{secondLevelDomain}",
			expected:         "https://foo.region.oraclecloud.com",
		},
		{
			// template with second level domain
			region:           StringToRegion("eu-amsterdam-1"),
			service:          "test",
			endpointTemplate: "https://foo.region.{secondLevelDomain}",
			expected:         "https://foo.region.oraclecloud.com",
		},
		{
			// template with second level domain
			region:           StringToRegion("ca-montreal-1"),
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

	region = StringToRegion("syd")
	assert.Equal(t, RegionAPSydney1, region)

	region = StringToRegion("mel")
	assert.Equal(t, RegionAPMelbourne1, region)

	region = StringToRegion("jed")
	assert.Equal(t, RegionMEJeddah1, region)

	region = StringToRegion("zrh")
	assert.Equal(t, RegionEUZurich1, region)

	region = StringToRegion("ams")
	assert.Equal(t, RegionEUAmsterdam1, region)

	region = StringToRegion("gru")
	assert.Equal(t, RegionSASaopaulo1, region)

	region = StringToRegion("uk-gov-london-1")
	assert.Equal(t, RegionUKGovLondon1, region)

	region = StringToRegion("ltn")
	assert.Equal(t, RegionUKGovLondon1, region)

	region = StringToRegion("yul")
	assert.Equal(t, RegionCAMontreal1, region)

	regionMetadataEnvVar := `{"realmKey":"OC0","realmDomainComponent":"testRealm.com","regionKey":"RTK","regionIdentifier":"us-testregion-1"}`
	os.Unsetenv("OCI_REGION_METADATA")
	os.Setenv("OCI_REGION_METADATA", regionMetadataEnvVar)
	region = StringToRegion("rtk")
	assert.Equal(t, Region("us-testregion-1"), region)

	fileContent :=
		`[
	{
		"realmKey" : "",
		"realmDomainComponent" : "oraclecloud.com",
		"regionKey" : "ABC",
		"regionIdentifier" : "ap-testregion-2"
	},
	{
		"realmKey" : "OC6",
		"realmDomainComponent" : "oraclensrcloud.com",
		"regionKey" : "DEF",
		"regionIdentifier" : "us-testregion-3"
	}
]`
	tmpLocation := path.Join(getHomeFolder(), ".oci", "regions-config.json")
	tmpPath := path.Join(getHomeFolder(), ".oci")

	if _, err := os.Stat(tmpPath); err != nil && os.IsNotExist(err) {
		if err := os.Mkdir(tmpPath, 0777); err != nil {
			assert.FailNow(t, err.Error())
		}
	}

	if err := ioutil.WriteFile(tmpLocation, []byte(fileContent), 0644); err != nil {
		assert.FailNow(t, err.Error())
	}
	region = StringToRegion("def")
	assert.Equal(t, Region("us-testregion-3"), region)

	region = StringToRegion("us-unknownregion-1")
	assert.Equal(t, Region("us-unknownregion-1"), region)
}

func TestSetRegionMetadataFromEnvVar(t *testing.T) {
	// normal case
	regionMetadataEnvVar := `{"realmKey":"OC0","realmDomainComponent":"testRealm.com","regionKey":"RTK","regionIdentifier":"us-testregion-1"}`
	os.Unsetenv("OCI_REGION_METADATA")
	os.Setenv("OCI_REGION_METADATA", regionMetadataEnvVar)
	// once provide region name
	expectedRegion := "us-testregion-1"
	ok := SetRegionMetadataFromEnvVar(&expectedRegion)
	assert.Equal(t, ok, true)
	assert.Equal(t, "oc0", regionRealm[Region("us-testregion-1")])
	assert.Equal(t, "testrealm.com", realm["oc0"])
	// once provide region short code
	shortCode := "rtk"
	ok = SetRegionMetadataFromEnvVar(&shortCode)
	assert.Equal(t, ok, true)
	assert.Equal(t, expectedRegion, shortCode)
	assert.Equal(t, "oc0", regionRealm[Region("us-testregion-1")])
	assert.Equal(t, "testrealm.com", realm["oc0"])

	// test corner case
	os.Unsetenv("OCI_REGION_METADATA")
	os.Setenv("OCI_REGION_METADATA", `"test": "test"`)
	ok = SetRegionMetadataFromEnvVar(&expectedRegion)
	assert.Equal(t, false, ok)

	os.Unsetenv("OCI_REGION_METADATA")
	os.Setenv("OCI_REGION_METADATA", `{"realmKey":"","realmDomainComponent":"testRealm.com","regionKey":"RTK","regionIdentifier":"us-testregion-1"}`)
	ok = SetRegionMetadataFromEnvVar(&expectedRegion)
	assert.Equal(t, false, ok)

	os.Unsetenv("OCI_REGION_METADATA")
	ok = SetRegionMetadataFromEnvVar(&expectedRegion)
	assert.Equal(t, false, ok)
}

func TestSetRegionMetadataFromCfgFile(t *testing.T) {
	// normal case
	fileContent :=
		`[
	{
		"realmKey" : "",
		"realmDomainComponent" : "oraclecloud.com",
		"regionKey" : "ABC",
		"regionIdentifier" : "ap-testregion-2"
	},
	{
		"realmKey" : "OC6",
		"realmDomainComponent" : "oraclensrcloud.com",
		"regionKey" : "DEF",
		"regionIdentifier" : "us-testregion-3"
	}
]`
	tmpLocation := path.Join(getHomeFolder(), ".oci", "regions-config.json")
	tmpPath := path.Join(getHomeFolder(), ".oci")

	if _, err := os.Stat(tmpPath); err != nil && os.IsNotExist(err) {
		if err := os.Mkdir(tmpPath, 0777); err != nil {
			assert.FailNow(t, err.Error())
		}
	}

	if err := ioutil.WriteFile(tmpLocation, []byte(fileContent), 0644); err != nil {
		assert.FailNow(t, err.Error())
	}
	// once provide region name
	expectedRegion := "us-testregion-3"
	ok := SetRegionMetadataFromCfgFile(&expectedRegion)
	assert.Equal(t, true, ok)
	assert.Equal(t, "", regionRealm[Region("us-testregion-2")])
	assert.Equal(t, "oc6", regionRealm[Region("us-testregion-3")])
	assert.Equal(t, "oraclensrcloud.com", realm["oc6"])
	// once provide region short code
	shortCode := "def"
	ok = SetRegionMetadataFromCfgFile(&shortCode)
	assert.Equal(t, true, ok)
	assert.Equal(t, "us-testregion-3", shortCode)
	assert.Equal(t, "", regionRealm[Region("us-testregion-2")])
	assert.Equal(t, "oc6", regionRealm[Region("us-testregion-3")])
	assert.Equal(t, "oraclensrcloud.com", realm["oc6"])

	//corner case
	fileContent = ""
	if _, err := os.Stat(tmpLocation); err == nil || os.IsExist(err) {
		os.Remove(tmpLocation)
	}
	ok = SetRegionMetadataFromCfgFile(&expectedRegion)
	assert.Equal(t, false, ok)

	if err := ioutil.WriteFile(tmpLocation, []byte(fileContent), 0644); err != nil {
		assert.FailNow(t, err.Error())
	}

	ok = SetRegionMetadataFromCfgFile(&expectedRegion)
	assert.Equal(t, false, ok)

}
