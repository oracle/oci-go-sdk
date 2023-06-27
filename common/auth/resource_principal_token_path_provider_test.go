// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.

package auth

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"log"
	"os"
	"strconv"
	"testing"
	"time"
)

const testTokenPathPrefix = "/tmp/temp-test-golang-sdk-kubernetesio-sa-token-"

func TestDefaultServiceAccountTokenProviderCanReadByPath(t *testing.T) {
	var tokenPath = testTokenPathPrefix + uuid()
	tokenString := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdWQiOiJrdWJlcm5ldGVzLmRlZmF1bHQiLCJleHAiOjQ4NDEwMTAwNTQsImlhdCI6MTY4NzQxMDA1NCwiaXNzIjoic3lzdGVtOnNlcnZpY2VhY2NvdW50Om15LW5hbWVzcGFjZTpteS1zYSIsInN1YiI6dHJ1ZX0.lOCsshvqcfX6I-FQcxAigrD7-KGhsKXuk97mmak5FFQ"
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
	tokenString := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdWQiOiJrdWJlcm5ldGVzLmRlZmF1bHQiLCJleHAiOjE2ODc0MDAwNTQsImlhdCI6MTY4NzQxMDA1NCwiaXNzIjoic3lzdGVtOnNlcnZpY2VhY2NvdW50Om15LW5hbWVzcGFjZTpteS1zYSIsInN1YiI6dHJ1ZX0.UpnEGk95K9TQaem9u-uz8-uMl3lF1Yd0rgHwT6p0A6w"
	createSaToken(tokenPath, tokenString)
	provider := NewDefaultServiceAccountTokenProvider().WithSaTokenPath(tokenPath)
	fmt.Println(provider.tokenPath)
	assert.NotNil(t, provider)

	_, err := provider.ServiceAccountToken()
	assert.Error(t, err)
	deleteSaToken(tokenPath)
}

func TestSuppliedServiceAccountTokenProviderCanDetectExpiredSaToken(t *testing.T) {
	tokenString := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdWQiOiJrdWJlcm5ldGVzLmRlZmF1bHQiLCJleHAiOjE2ODc0MDAwNTQsImlhdCI6MTY4NzQxMDA1NCwiaXNzIjoic3lzdGVtOnNlcnZpY2VhY2NvdW50Om15LW5hbWVzcGFjZTpteS1zYSIsInN1YiI6dHJ1ZX0.UpnEGk95K9TQaem9u-uz8-uMl3lF1Yd0rgHwT6p0A6w"
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
