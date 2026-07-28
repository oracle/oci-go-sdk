// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.

package auth

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

const testTokenPathPrefix = "/tmp/temp-test-golang-sdk-kubernetesio-sa-token-"

func TestDefaultServiceAccountTokenProviderCanReadByPath(t *testing.T) {
	var tokenPath = testTokenPathPrefix + uuid()
	tokenString := testServiceAccountToken(t, time.Now().Add(time.Hour))
	createSaToken(tokenPath, tokenString)
	provider := NewDefaultServiceAccountTokenProvider().WithSaTokenPath(tokenPath)
	fmt.Println(provider.tokenPath)
	assert.NotNil(t, provider)

	_, err := provider.ServiceAccountToken()
	assert.NoError(t, err)
	deleteSaToken(tokenPath)
}

func TestDefaultServiceAccountTokenProviderCanDetectWrongPath(t *testing.T) {
	provider := NewDefaultServiceAccountTokenProvider().WithSaTokenPath("/wrong_path")
	fmt.Println(provider.tokenPath)
	assert.NotNil(t, provider)

	_, err := provider.ServiceAccountToken()
	assert.Error(t, err)
}

func TestDefaultServiceAccountTokenProviderCanDetectMalformedSaToken(t *testing.T) {
	var tokenPath = testTokenPathPrefix + uuid()
	createSaToken(tokenPath, "malformed string")
	provider := NewDefaultServiceAccountTokenProvider().WithSaTokenPath(tokenPath)
	fmt.Println(provider.tokenPath)
	assert.NotNil(t, provider)

	_, err := provider.ServiceAccountToken()
	assert.Error(t, err)
	deleteSaToken(tokenPath)
}

func TestDefaultServiceAccountTokenProviderCanDetectExpiredSaToken(t *testing.T) {
	var tokenPath = testTokenPathPrefix + uuid()
	tokenString := testServiceAccountToken(t, time.Now().Add(-time.Hour))
	createSaToken(tokenPath, tokenString)
	provider := NewDefaultServiceAccountTokenProvider().WithSaTokenPath(tokenPath)
	fmt.Println(provider.tokenPath)
	assert.NotNil(t, provider)

	_, err := provider.ServiceAccountToken()
	assert.Error(t, err)
	deleteSaToken(tokenPath)
}

func TestSuppliedServiceAccountTokenProviderCanDetectExpiredSaToken(t *testing.T) {
	tokenString := testServiceAccountToken(t, time.Now().Add(-time.Hour))
	provider := NewSuppliedServiceAccountTokenProvider(tokenString)
	assert.NotNil(t, provider)
	_, err := provider.ServiceAccountToken()
	assert.Error(t, err)
}

func TestSuppliedServiceAccountTokenProviderCanDetectMalformedSaToken(t *testing.T) {
	provider := NewSuppliedServiceAccountTokenProvider("malformed string")
	assert.NotNil(t, provider)
	_, err := provider.ServiceAccountToken()
	assert.Error(t, err)
}

func testServiceAccountToken(t *testing.T, expiration time.Time) string {
	t.Helper()

	header, err := json.Marshal(map[string]string{
		"alg": "none",
		"typ": "JWT",
	})
	if err != nil {
		t.Fatalf("failed to marshal service account token header: %v", err)
	}
	payload, err := json.Marshal(map[string]interface{}{
		"exp": expiration.Unix(),
	})
	if err != nil {
		t.Fatalf("failed to marshal service account token payload: %v", err)
	}

	return base64.RawURLEncoding.EncodeToString(header) + "." +
		base64.RawURLEncoding.EncodeToString(payload) + ".synthetic-signature"
}

func createSaToken(filePath string, content string) {
	f, err := os.Create(filePath)

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	_, err2 := f.WriteString(content)

	if err2 != nil {
		log.Fatal(err2)
	}

	fmt.Println("SA Token created at " + filePath + "\n")
}

func deleteSaToken(filePath string) {
	e := os.Remove(filePath)
	if e != nil {
		log.Fatal(e)
	}
}
func uuid() string {
	ts := time.Now().Unix()
	return strconv.FormatInt(ts, 10)
}
