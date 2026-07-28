// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.

package auth

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"encoding/pem"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/oracle/oci-go-sdk/v65/common"
	"github.com/stretchr/testify/assert"
)

var (
	test_token = testResourcePrincipalV3Token(
		map[string]string{"alg": "HS256", "typ": "JWT"},
		map[string]interface{}{
			"aud":        "kubernetes.default",
			"exp":        time.Now().Add(time.Hour).Unix(),
			"iss":        "system:serviceaccount:test-namespace:test-service-account",
			"res_tenant": "test-tenant",
			"sub":        true,
		},
	)
	rptUrlForParent = "https://sample.test/101010/resource/resourceId/actions/rptUrl"
	rpstLeaf        = testResourcePrincipalV3Token(
		map[string]string{"alg": "RS256", "kid": "test-key", "typ": "JWT"},
		map[string]interface{}{
			"aud":        "oci",
			"exp":        time.Now().Add(time.Hour).Unix(),
			"iss":        "test-auth-service",
			"ptype":      "resource",
			"res_tenant": "test-tenant",
			"res_type":   "database",
			"sub":        "ocid1.datawarehouse.oc1..test-resource",
		},
	)
)

var envVarsrpt3 = map[string]string{
	ResourcePrincipalRptURLForParent:       rptUrlForParent,
	ResourcePrincipalRpstEndpointForParent: "https://identity/101010/v1/rpstEndpoint",
	ResourcePrincipalVersionEnvVar:         "3.0",
	ResourcePrincipalVersionForLeaf:        "2.2",
	ResourcePrincipalRpstForLeaf:           rpstLeaf,
}

func TestResourcePrincipalV3TestTokensAreParseable(t *testing.T) {
	for _, tokenString := range []string{test_token, rpstLeaf} {
		token, err := parseJwt(tokenString)
		if assert.NoError(t, err) {
			assert.False(t, token.expired())
		}
	}
}

func testResourcePrincipalV3Token(header map[string]string, payload map[string]interface{}) string {
	headerBytes, err := json.Marshal(header)
	if err != nil {
		panic(fmt.Sprintf("failed to marshal resource principal V3 token header: %v", err))
	}
	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		panic(fmt.Sprintf("failed to marshal resource principal V3 token payload: %v", err))
	}

	return base64.RawURLEncoding.EncodeToString(headerBytes) + "." +
		base64.RawURLEncoding.EncodeToString(payloadBytes) + ".synthetic-signature"
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
