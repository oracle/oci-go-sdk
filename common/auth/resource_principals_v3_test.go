// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.

package auth

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/oracle/oci-go-sdk/v65/common"
	"github.com/stretchr/testify/assert"
)

var (
	test_token      = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdWQiOiJrdWJlcm5ldGVzLmRlZmF1bHQiLCJleHAiOjQ4NDEwMTAwNTQsImlhdCI6MTY4NzQxMDA1NCwiaXNzIjoic3lzdGVtOnNlcnZpY2VhY2NvdW50Om15LW5hbWVzcGFjZTpteS1zYSIsInN1YiI6dHJ1ZSwicmVzX3RlbmFudCI6InJhbmRvbSJ9._YXkRQbCxdnFfXYYHbLpi6lryxpHPXqGTxWhDTeB0_g"
	rptUrlForParent = "https://sample.test/101010/resource/resourceId/actions/rptUrl"
	rpstLeaf        = "eyJraWQiOiJhc3ciLCJhbGciOiJSUzI1NiJ9.eyJzdWIiOiJvY2lkMS5kYXRhd2FyZWhvdXNlLmN1c3RvbWVyLWJpZy1kYi0xIiwiaXNzIjoiYXV0aFNlcnZpY2Uub3JhY2xlLmNvbSIsInJlc190ZW5hbnQiOiJjdXN0b21lci10ZW5hbnQtMSIsInB0eXBlIjoicmVzb3VyY2UiLCJyZXNfdHlwZSI6ImRhdGF3YXJlaG91c2UiLCJhdWQiOiJvY2kiLCJvcGMtdGFnIjoiVjEsb2NpZDEuZHluYW1pY2dyb3VwLmRldi4uYWFhYWFhYWEsb28ycHA2M3Y3NTMyNXpiNnJuNHYzYWgzYnRlam5ubTJvdDN1NWJyZnd1Y2VvczRpZDZ0YSIsIm9wYy1idW1wIjoiSDRzSUFBQUFBQUFBQUUyUHpRcUNRQlNGTTExTVJSQXRLbHBHYlFRSE1sSm8xVVAwQXVQTWJSeHJISnVmMHA0LUF4SFA5bjczTzV6bFRWSEJqcGcxSlpHQ2NxMWNoUm04TVNaZDVGM3cwMHZGejhaQnduaVpTZjA0eDBWaThuZGVrSlJMY0o4NnF3cWFmRTJkMGx3bE5RbFQ1Qy04N1ZxRFVVNVR3RlRKaW1ncm9iUllzTjJLT21PVkJCME5EdEh4T2dvbktGaE1ROThiLS1FZWpWdkh2SGZZcG9MZG5CRkxQa1JEcnB5Qjl1R0N2SmFhOVZSclAzU1RoaVR1R3pQQkk1Yjl5LUlBZVp2Z0IxX2paUVFDQVFBQSIsInR0eXBlIjoicmVzIiwicmVzX2lkIjoib2NpZDEuZGF0YXdhcmVob3VzZS5jdXN0b21lci1iaWctZGItMSIsIm9wYy1pbnN0YW5jZSI6Im9jaWQxLmluc3RhbmNlLmpveWZ1bC5kYm5vZGUuMSIsInJlc190YWciOiIiLCJyZXNfY29tcGFydG1lbnQiOiJjdXN0b21lci1jb21wYXJ0bWVudC0xIiwiZXhwIjoxNTUyNjgxMTY4LCJvcGMtY29tcGFydG1lbnQiOiJkYmFhcy1jb21wYXJ0bWVudC0xIiwiaWF0IjoxNTUyNjc5OTY4LCJqdGkiOiIzZDQxNDNlOC01ZTMyLTRlMTItYTM4Yy01OTc0NjUwMTA3MDMiLCJ0ZW5hbnQiOiJjdXN0b21lci10ZW5hbnQtMSIsImp3ayI6IntcImtpZFwiOlwiY3VzdG9tZXItZGJub2RlLTFcIixcIm5cIjpcImxzLUFDNGhpS0stMTFVdTFEZ3VLTFE1VGFhZGpNR1hCcDRhMFVFS2w0dnJjcmF3b2V6X3BuUS1pNS1nNV9XTU5xVXdrdUtBcXVTZnlVS25yZEhhV3d4b2RWcmRleTk1T3R4ckIySzNRdzgzaURkcUltSkhfWFp1cERfRHR0SzduS3N6Qy01TFI1Ums3SHF5Y094eEZVNzBNcGduQW9IaVNUM2V0VjJVZlJkNXRtb0dOaTdOSURORWJnSVpmcnczYUVYbHBzaGM2ckpVdUEyOG55ZUNjOVFtOHllMHUwN0UzamlCYmp5RjNFVWhTelNxblFsUlVNVEdaR1ZSZGpfRG9tcEhUVkFPNEJqUnVIZURWWGtWNjh1TzNrSEdTZUVPc2xsZmJZTkpaYUtCQTB3aUxrZkViWVBEWDFwbTM4UFAzcnJFbWhxeElObzdoYVFWakRXRDNKUVwiLFwiZVwiOlwiQVFBQlwiLFwia3R5XCI6XCJSU0FcIixcImFsZ1wiOlwiUlMyNTZcIixcInVzZVwiOlwic2lnXCJ9Iiwib3BjLXRlbmFudCI6ImRiYWFzLXRlbmFudCJ9.LqRt9JXSdcLahdwACjw_p_KHQhKde-NaVZG3zMjzWX6bVad-SRZYWKQSlk6Tq4f1ZNN0uxlP-d2snQAp3Kw-cQRrdCDOmD_0CDgR-yre-YbJDsJbEncczUIbe-ASeq_Sh9zDROVuD_7NdrmUCiVH2g-UYpYkuKKqu_tVjL2uy77W5_DGobPArEFvZ2GnyHT7gVVv12RnINtgr2jJULhegPBfvnp9-fhhZ7_PcsJ7Z5FkPzLtLOwEm3Lbm3veyUVUviu1CSjXnK67KzjS18TVGi723bkxYBf9lYDHfaXh9EEHzPtxeLAl3VrGjwZUv_ih0FRmoM7wgq8HMRjNACMo6g"
)

var envVarsrpt3 = map[string]string{
	ResourcePrincipalRptURLForParent:       rptUrlForParent,
	ResourcePrincipalRpstEndpointForParent: "https://identity/101010/v1/rpstEndpoint",
	ResourcePrincipalVersionEnvVar:         "3.0",
	ResourcePrincipalVersionForLeaf:        "2.2",
	ResourcePrincipalRpstForLeaf:           rpstLeaf,
}

func TestNewResourcePrinicipalV3ConfigurationProviderLeaf(t *testing.T) {

	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(fmt.Sprintf(`{"fake": "data","resourcePrincipalToken": "%s","servicePrincipalSessionToken": "%s","token":"%s"}`, test_token, test_token, test_token)))
	}))
	defer mockServer.Close()

	unsetAllVars()

	setupResourcePrincipalsEnvsWithValues(envVarsrpt3)

	t.Setenv(ResourcePrincipalRptURLForParent, mockServer.URL)
	t.Setenv(ResourcePrincipalRpstEndpointForParent, mockServer.URL)

	leafResourcePrincipalKeyProvider := &fakeConfigProviderWithClaimAccess{}

	provider, err := ResourcePrincipalConfigurationProviderV3(leafResourcePrincipalKeyProvider)

	assert.NoError(t, err)

	assert.NotNil(t, provider)

	valid, err := common.IsConfigurationProviderValid(provider)
	assert.NoError(t, err)
	assert.True(t, valid)

}

func TestNewResourcePrinicipalV3ConfigurationProviderWithDepth(t *testing.T) {
	serverCalls := 0

	//mock server providing rpt
	mockParentServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, serverCalls, 2)
		serverCalls += 1
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(fmt.Sprintf(`{"fake": "data","resourcePrincipalToken": "%s","servicePrincipalSessionToken": "%s"}`, test_token, test_token)))
	}))
	defer mockParentServer.Close()

	//mock server to get parent rpt url
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, serverCalls, 0)
		serverCalls += 1
		w.Header().Set("opc-parent-rpt-url", mockParentServer.URL)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(fmt.Sprintf(`{"fake": "data","resourcePrincipalToken": "%s","servicePrincipalSessionToken": "%s"}`, test_token, test_token)))
	}))
	defer mockServer.Close()

	// exchange server to exchange rpt for rpst
	mockExchangeServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		serverCalls += 1
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(fmt.Sprintf(`{"fake": "data","token":"%s"}`, test_token)))
	}))
	defer mockExchangeServer.Close()

	unsetAllVars()

	setupResourcePrincipalsEnvsWithValues(envVarsrpt3)

	t.Setenv(ResourcePrincipalRptURLForParent, mockServer.URL)
	t.Setenv(ResourcePrincipalRpstEndpointForParent, mockExchangeServer.URL)

	leafResourcePrincipalKeyProvider := &fakeConfigProviderWithClaimAccess{}

	provider, err := ResourcePrincipalConfigurationProviderV3(leafResourcePrincipalKeyProvider)

	assert.NoError(t, err)

	assert.NotNil(t, provider)

	valid, err := common.IsConfigurationProviderValid(provider)
	assert.NoError(t, err)
	assert.True(t, valid)

	assert.Equal(t, 4, serverCalls)

}

type fakeConfigProviderWithClaimAccess struct {
}

func (f *fakeConfigProviderWithClaimAccess) AuthType() (common.AuthConfig, error) {
	return common.AuthConfig{}, nil
}

func (f *fakeConfigProviderWithClaimAccess) GetClaim(key string) (interface{}, error) {
	return test_token, nil
}

func (f *fakeConfigProviderWithClaimAccess) KeyFingerprint() (string, error) {
	return "fingerprint", nil
}

func (f *fakeConfigProviderWithClaimAccess) KeyID() (string, error) {
	return "keyId", nil
}

func (f *fakeConfigProviderWithClaimAccess) PrivateRSAKey() (*rsa.PrivateKey, error) {
	block, _ := pem.Decode([]byte(testPrivateKey))
	key, _ := x509.ParsePKCS1PrivateKey(block.Bytes)
	return key, nil
}

func (f *fakeConfigProviderWithClaimAccess) Region() (string, error) {
	return "region", nil
}

func (f *fakeConfigProviderWithClaimAccess) TenancyOCID() (string, error) {
	return "tenancy", nil
}

func (f *fakeConfigProviderWithClaimAccess) UserOCID() (string, error) {
	return "userId", nil
}
