// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.

package common

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAllRegions(t *testing.T) {
	fileName := "regions.json"
	content, err := ioutil.ReadFile(fileName)
	assert.True(t, err == nil, "Failed to read regions.json.")
	var jsonArr []map[string]string
	err = json.Unmarshal(content, &jsonArr)
	assert.True(t, err == nil, "Failed to unmarshal contents.")

	for _, jsonItem := range jsonArr {
		potentialRegion := Region(jsonItem["regionIdentifier"])
		_, ok := regionRealm[potentialRegion]
		assert.True(t, ok, fmt.Sprintf("Region %s not found.", jsonItem["regionIdentifier"]))
		_, ok = realm[strings.ToLower(jsonItem["realmKey"])]
		assert.True(t, ok, fmt.Sprintf("Realm %s not found.", jsonItem["realmKey"]))
	}
}

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

	region = StringToRegion("ap-chuncheon-1")
	endpoint = region.Endpoint("bar")
	assert.Equal(t, "bar.ap-chuncheon-1.oraclecloud.com", endpoint)

	region = StringToRegion("us-sanjose-1")
	endpoint = region.Endpoint("bar")
	assert.Equal(t, "bar.us-sanjose-1.oraclecloud.com", endpoint)

	region = StringToRegion("me-dubai-1")
	endpoint = region.Endpoint("bar")
	assert.Equal(t, "bar.me-dubai-1.oraclecloud.com", endpoint)

	region = StringToRegion("uk-cardiff-1")
	endpoint = region.Endpoint("bar")
	assert.Equal(t, "bar.uk-cardiff-1.oraclecloud.com", endpoint)

	region = StringToRegion("sa-santiago-1")
	endpoint = region.Endpoint("bar")
	assert.Equal(t, "bar.sa-santiago-1.oraclecloud.com", endpoint)

	region = StringToRegion("sa-vinhedo-1")
	endpoint = region.Endpoint("bar")
	assert.Equal(t, "bar.sa-vinhedo-1.oraclecloud.com", endpoint)

	os.Setenv("OCI_DEFAULT_REALM", "oraclevaporcloud.space")
	region = StringToRegion("us-unknownregion-1")
	endpoint = region.Endpoint("foo")
	assert.Equal(t, "foo.us-unknownregion-1.oraclevaporcloud.space", endpoint)
	os.Unsetenv("OCI_DEFAULT_REALM")
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

	//OC8
	region = StringToRegion("ap-chiyoda-1")
	endpoint = region.Endpoint("bar")
	assert.Equal(t, "bar.ap-chiyoda-1.oraclecloud8.com", endpoint)
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
			region:           StringToRegion("ap-chuncheon-1"),
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
			region:           StringToRegion("me-dubai-1"),
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
			// template with second level domain
			region:           StringToRegion("us-sanjose-1"),
			service:          "test",
			endpointTemplate: "https://foo.region.{secondLevelDomain}",
			expected:         "https://foo.region.oraclecloud.com",
		},
		{
			// template with second level domain
			region:           StringToRegion("uk-cardiff-1"),
			service:          "test",
			endpointTemplate: "https://foo.region.{secondLevelDomain}",
			expected:         "https://foo.region.oraclecloud.com",
		},
		{
			// template with second level domain
			region:           StringToRegion("sa-santiago-1"),
			service:          "test",
			endpointTemplate: "https://foo.region.{secondLevelDomain}",
			expected:         "https://foo.region.oraclecloud.com",
		},
		{
			// template with second level domain
			region:           StringToRegion("sa-vinhedo-1"),
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
		{
			// template with everything for OC8
			region:           StringToRegion("ap-chiyoda-1"),
			service:          "test",
			endpointTemplate: "https://test.{region}.{secondLevelDomain}",
			expected:         "https://test.ap-chiyoda-1.oraclecloud8.com",
		},
		{
			service:          "identity",
			region:           StringToRegion("broom6.us.oracle.com"),
			endpointTemplate: "https://iaas.{region}.{secondLevelDomain}",
			expected:         "https://iaas.broom6.us.oracle.com",
		},
	}

	for _, testData := range testDataSet {
		endpoint := testData.region.EndpointForTemplate(testData.service, testData.endpointTemplate)
		assert.Equal(t, testData.expected, endpoint)
	}
}

func TestEndpointWithServiceName(t *testing.T) {
	//no dot in region with service name and endpoint template
	region := StringToRegion("us-phoenix-1")
	service := "test"
	endpointTemplate := "https://iaas.{region}.{secondLevelDomain}"
	serviceName := "iaas"
	endpoint, err := region.EndpointForTemplateDottedRegion(service, endpointTemplate, serviceName)
	assert.NoError(t, err)
	assert.Equal(t, "https://iaas.us-phoenix-1.oraclecloud.com", endpoint)

	service = "identity"
	region = StringToRegion("broom6.us.oracle.com")
	serviceName = "iam"
	endpointTemplate = "https://identity.{region}.oci.{secondLevelDomain}"
	endpoint, err = region.EndpointForTemplateDottedRegion(service, endpointTemplate, serviceName)
	assert.NoError(t, err)
	assert.Equal(t, "https://iam.broom6.us.oracle.com", endpoint)

	//dot in region with service name and without endpoint template
	region = StringToRegion("some.customerdomain.com")
	service = "analytics"
	serviceName = "test"
	endpoint, err = region.EndpointForTemplateDottedRegion(service, "", serviceName)
	assert.NoError(t, err)
	assert.Equal(t, "https://test.some.customerdomain.com", endpoint)

	region = StringToRegion("broom6.us.oracle.com")
	service = "compute"
	serviceName = "iaas"
	endpoint, err = region.EndpointForTemplateDottedRegion(service, "", serviceName)
	assert.NoError(t, err)
	assert.Equal(t, "https://iaas.broom6.us.oracle.com", endpoint)

	//dot in region without service name and without endpoint template
	service = "analytics"
	region = StringToRegion("broom6.us.oracle.com")
	endpoint, err = region.EndpointForTemplateDottedRegion(service, "", "")
	assert.Error(t, err)
}

func TestStringToRegion(t *testing.T) {
	region := StringToRegion("yyz")
	assert.Equal(t, RegionCAToronto1, region)

	region = StringToRegion("nrt")
	assert.Equal(t, RegionAPTokyo1, region)

	region = StringToRegion("kix")
	assert.Equal(t, RegionAPOsaka1, region)

	region = StringToRegion("nja")
	assert.Equal(t, RegionAPChiyoda1, region)

	region = StringToRegion("syd")
	assert.Equal(t, RegionAPSydney1, region)

	region = StringToRegion("mel")
	assert.Equal(t, RegionAPMelbourne1, region)

	region = StringToRegion("jed")
	assert.Equal(t, RegionMEJeddah1, region)

	region = StringToRegion("dxb")
	assert.Equal(t, RegionMEDubai1, region)

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

	region = StringToRegion("yny")
	assert.Equal(t, RegionAPChuncheon1, region)

	region = StringToRegion("sjc")
	assert.Equal(t, RegionSJC1, region)

	region = StringToRegion("brs")
	assert.Equal(t, RegionUKGovCardiff1, region)

	region = StringToRegion("cwl")
	assert.Equal(t, RegionUKCardiff1, region)

	region = StringToRegion("scl")
	assert.Equal(t, RegionSASantiago1, region)

	region = StringToRegion("vcp")
	assert.Equal(t, RegionSAVinhedo1, region)

	regionMetadataEnvVar := `{"realmKey":"OC0","realmDomainComponent":"testRealm.com","regionKey":"RTK","regionIdentifier":"us-testregion-1"}`
	os.Unsetenv("OCI_REGION_METADATA")
	os.Setenv("OCI_REGION_METADATA", regionMetadataEnvVar)
	defer os.Unsetenv("OCI_REGION_METADATA")
	readEnvVar = true

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
		"realmKey" : "OC100",
		"realmDomainComponent" : "oraclefoobar.com",
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
	defer os.Remove(tmpLocation)
	readEnvVar = true
	readCfgFile = true
	region = StringToRegion("def")
	assert.Equal(t, Region("us-testregion-3"), region)

	region = StringToRegion("us-unknownregion-1")
	assert.Equal(t, Region("us-unknownregion-1"), region)
}

func TestSetRegionMetadataFromEnvVarWithNormalRegionName(t *testing.T) {
	regionMetadataEnvVar := `{"realmKey":"OC0","realmDomainComponent":"testRealm.com","regionKey":"RTK","regionIdentifier":"us-testregion-1"}`
	os.Unsetenv("OCI_REGION_METADATA")
	os.Setenv("OCI_REGION_METADATA", regionMetadataEnvVar)
	defer os.Unsetenv("OCI_REGION_METADATA")
	// once provide region name
	expectedRegion := "us-testregion-1"
	readEnvVar = true
	ok := setRegionMetadataFromEnvVar(&expectedRegion)
	assert.Equal(t, ok, true)
	assert.Equal(t, "oc0", regionRealm[Region("us-testregion-1")])
	assert.Equal(t, "testrealm.com", realm["oc0"])

	ok = setRegionMetadataFromEnvVar(&expectedRegion)
	assert.Equal(t, ok, false)

	r := StringToRegion(expectedRegion)
	assert.Equal(t, Region("us-testregion-1"), r)
	assert.Equal(t, "oc0", regionRealm["us-testregion-1"])
	assert.Equal(t, "testrealm.com", realm["oc0"])
}

func TestSetRegionMetadataFromEnvVarWithShortRegionName(t *testing.T) {
	regionMetadataEnvVar := `{"realmKey":"OC0","realmDomainComponent":"testRealm.com","regionKey":"RTK","regionIdentifier":"us-testregion-1"}`
	os.Unsetenv("OCI_REGION_METADATA")
	os.Setenv("OCI_REGION_METADATA", regionMetadataEnvVar)
	defer os.Unsetenv("OCI_REGION_METADATA")

	expectedRegion := "us-testregion-1"
	shortCode := "rtk"
	readEnvVar = true
	ok := setRegionMetadataFromEnvVar(&shortCode)
	assert.Equal(t, ok, true)
	assert.Equal(t, expectedRegion, shortCode)
	assert.Equal(t, "oc0", regionRealm[Region("us-testregion-1")])
	assert.Equal(t, "testrealm.com", realm["oc0"])

	ok = setRegionMetadataFromEnvVar(&shortCode)
	assert.Equal(t, ok, false)
	r := StringToRegion(shortCode)
	assert.Equal(t, Region("us-testregion-1"), r)
	assert.Equal(t, "oc0", regionRealm["us-testregion-1"])
	assert.Equal(t, "testrealm.com", realm["oc0"])
}

func TestSetRegionMetadataFromEnvVarInvalidEnvVar(t *testing.T) {
	os.Unsetenv("OCI_REGION_METADATA")
	os.Setenv("OCI_REGION_METADATA", `"test": "test"`)
	defer os.Unsetenv("OCI_REGION_METADATA")

	expectedRegion := "us-testregion-1"
	readEnvVar = true
	ok := setRegionMetadataFromEnvVar(&expectedRegion)
	assert.Equal(t, false, ok)

	os.Unsetenv("OCI_REGION_METADATA")
	os.Setenv("OCI_REGION_METADATA", `{"realmKey":"","realmDomainComponent":"testRealm.com","regionKey":"RTK","regionIdentifier":"us-testregion-1"}`)
	defer os.Unsetenv("OCI_REGION_METADATA")
	readEnvVar = true
	ok = setRegionMetadataFromEnvVar(&expectedRegion)
	assert.Equal(t, false, ok)
}

func TestSetRegionMetadataFromEnvVarNoEnvVar(t *testing.T) {
	os.Unsetenv("OCI_REGION_METADATA")
	expectedRegion := "us-testregion-1"
	readEnvVar = true
	ok := setRegionMetadataFromEnvVar(&expectedRegion)
	assert.Equal(t, false, ok)
}

func TestSetRegionMetadataFromCfgFileWithNormalRegionName(t *testing.T) {
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
		"realmKey" : "OC100",
		"realmDomainComponent" : "oraclefoobar.com",
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
	defer os.Remove(tmpLocation)

	expectedRegion := "us-testregion-3"
	readCfgFile = true
	ok := setRegionMetadataFromCfgFile(&expectedRegion)
	assert.Equal(t, true, ok)
	assert.Equal(t, "", regionRealm[Region("us-testregion-2")])
	assert.Equal(t, "oc100", regionRealm[Region("us-testregion-3")])
	assert.Equal(t, "oraclefoobar.com", realm["oc100"])

	ok = setRegionMetadataFromCfgFile(&expectedRegion)
	assert.Equal(t, false, ok)

	r := StringToRegion(expectedRegion)
	assert.Equal(t, Region("us-testregion-3"), r)
	assert.Equal(t, "oc100", regionRealm[Region("us-testregion-3")])
	assert.Equal(t, "oraclefoobar.com", realm["oc100"])
}

func TestSetRegionMetadataFromCfgFileWithShortRegionName(t *testing.T) {
	fileContent :=
		`[
	{
		"realmKey" : "",
		"realmDomainComponent" : "oraclecloud.com",
		"regionKey" : "ABC",
		"regionIdentifier" : "ap-testregion-2"
	},
	{
		"realmKey" : "OC100",
		"realmDomainComponent" : "oraclefoobar.com",
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
	defer os.Remove(tmpLocation)

	shortCode := "def"
	readCfgFile = true
	ok := setRegionMetadataFromCfgFile(&shortCode)
	assert.Equal(t, true, ok)
	assert.Equal(t, "us-testregion-3", shortCode)
	assert.Equal(t, "", regionRealm[Region("us-testregion-2")])
	assert.Equal(t, "oc100", regionRealm[Region("us-testregion-3")])
	assert.Equal(t, "oraclefoobar.com", realm["oc100"])

	ok = setRegionMetadataFromCfgFile(&shortCode)
	assert.Equal(t, false, ok)

	r := StringToRegion("Def")
	assert.Equal(t, Region("us-testregion-3"), r)
	assert.Equal(t, "oc100", regionRealm[Region("us-testregion-3")])
	assert.Equal(t, "oraclefoobar.com", realm["oc100"])
}

func TestSetRegionMetadataFromCfgFileWithInvalidFileContent(t *testing.T) {
	fileContent := ""
	tmpLocation := path.Join(getHomeFolder(), ".oci", "regions-config.json")

	if _, err := os.Stat(tmpLocation); err == nil || os.IsExist(err) {
		os.Remove(tmpLocation)
	}
	defer os.Remove(tmpLocation)

	expectedRegion := "us-testregion-3"
	readCfgFile = true
	ok := setRegionMetadataFromCfgFile(&expectedRegion)
	assert.Equal(t, false, ok)

	if err := ioutil.WriteFile(tmpLocation, []byte(fileContent), 0644); err != nil {
		assert.FailNow(t, err.Error())
	}

	ok = setRegionMetadataFromCfgFile(&expectedRegion)
	assert.Equal(t, false, ok)
}

func getRegionInfoFromInstanceMetadataServiceSucceed() ([]byte, error) {
	contentString :=
		`{ 
	"realmKey" : "OC-test",
	"realmDomainComponent" : "test.com",
	"regionKey" : "key",
	"regionIdentifier" : "us-test-1"
}`
	return []byte(contentString), nil
}

func getRegionInfoFromInstanceMetadataServiceInvalidContent() ([]byte, error) {
	contentString :=
		`{ 
	"realmKey" : "",
	"realmDomainComponent" : "",
	"regionKey" : "",
	"regionIdentifier" : ""
}`
	return []byte(contentString), nil
}

func getRegionInfoFromInstanceMetadataServiceFail() ([]byte, error) {
	return nil, errors.New("test error")
}

func TestSetRegionFromInstanceMetadataService(t *testing.T) {
	expectedRegion := "us-test-1"
	getRegionInfoFromInstanceMetadataService = getRegionInfoFromInstanceMetadataServiceSucceed
	EnableInstanceMetadataServiceLookup()
	ok := setRegionFromInstanceMetadataService(&expectedRegion)
	assert.Equal(t, true, ok)
	assert.Equal(t, "oc-test", regionRealm[Region("us-test-1")])
	assert.Equal(t, "test.com", realm["oc-test"])

	shortCode := "testRegionKey"
	getRegionInfoFromInstanceMetadataService = getRegionInfoFromInstanceMetadataServiceSucceed
	EnableInstanceMetadataServiceLookup()
	ok = setRegionFromInstanceMetadataService(&shortCode)
	assert.Equal(t, true, ok)
	assert.Equal(t, "oc-test", regionRealm[Region("us-test-1")])
	assert.Equal(t, "test.com", realm["oc-test"])

	getRegionInfoFromInstanceMetadataService = getRegionInfoFromInstanceMetadataServiceInvalidContent
	ok = setRegionFromInstanceMetadataService(&expectedRegion)
	assert.Equal(t, false, ok)

	getRegionInfoFromInstanceMetadataService = getRegionInfoFromInstanceMetadataServiceFail
	ok = setRegionFromInstanceMetadataService(&expectedRegion)
	assert.Equal(t, false, ok)
}

func TestGetRealmIDFromRegion(t *testing.T) {
	// normal region:
	region := Region("us-langley-1")
	realmID, err := region.RealmID()
	assert.NoError(t, err)
	assert.Equal(t, "oc2", realmID)

	// not existed region:
	region = Region("us-whatever-1")
	realmID, err = region.RealmID()
	assert.Equal(t, "", realmID)
	assert.Error(t, err)

	// defined in env var region
	regionMetadataEnvVar := `{"realmKey":"OC0","realmDomainComponent":"testRealm.com","regionKey":"RTK","regionIdentifier":"us-testregion-1"}`
	os.Unsetenv("OCI_REGION_METADATA")
	os.Setenv("OCI_REGION_METADATA", regionMetadataEnvVar)
	defer os.Unsetenv("OCI_REGION_METADATA")

	region = StringToRegion("rtk")
	assert.Equal(t, Region("us-testregion-1"), region)
	realmID, err = region.RealmID()
	assert.NoError(t, err)
	assert.Equal(t, "oc0", realmID)
}
