// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.

package auth

import (
	"crypto/rsa"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"sync"
	"testing"
	"time"

	"github.com/oracle/oci-go-sdk/v65/common"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

const (
	MetaDataBaseURLEnvVarName = `OCI_METADATA_BASE_URL`
)

func TestInstancePrincipalKeyProvider_getRegionForFederationClient(t *testing.T) {
	t.Parallel()
	regionServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "phx")
	}))
	defer regionServer.Close()

	actualRegion, err := getRegionForFederationClient(&http.Client{}, regionServer.URL)

	assert.NoError(t, err)
	assert.Equal(t, common.RegionPHX, actualRegion)
}

// Test the Federation Client still works when the metadataBaseURL environment
// variable is unset, and thus the default is used.
func TestInstancePrincipalKeyProvider_getRegionForFederationClienURLUnset(t *testing.T) {
	// t.Setenv temporarily sets the value of the environment variable for the duration of the test
	t.Setenv(MetaDataBaseURLEnvVarName, "")
	defer t.Setenv(MetaDataBaseURLEnvVarName, defaultMetadataBaseURL)
	regionServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "phx")
	}))
	defer regionServer.Close()

	actualRegion, err := getRegionForFederationClient(&http.Client{}, regionServer.URL)

	assert.NoError(t, err)
	assert.Equal(t, common.RegionPHX, actualRegion)
}

func TestInstancePrincipalKeyProvider_getRegionForFederationClientNotFound(t *testing.T) {
	t.Parallel()
	regionServer := httptest.NewServer(http.NotFoundHandler())
	defer regionServer.Close()

	_, err := getRegionForFederationClient(&http.Client{}, regionServer.URL)

	assert.Error(t, err)
}

func TestInstancePrincipalKeyProvider_getRegionForFederationClientTimeout(t *testing.T) {
	t.Parallel()
	HandlerFunc := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(100 * time.Millisecond)
	})
	regionServer := httptest.NewServer(http.TimeoutHandler(HandlerFunc, 20*time.Millisecond, "Timeout occured"))
	defer regionServer.Close()

	start := time.Now()
	response, _ := getRegionForFederationClient(&http.Client{}, regionServer.URL)
	assert.NotNil(t, response)
	elapsed := time.Since(start)
	assert.GreaterOrEqual(t, elapsed.Seconds(), 3.0)
}

func TestInstancePrincipalKeyProvider_getRegionForFederationClientNotFoundRetrySuccess(t *testing.T) {
	t.Parallel()
	responses := []func(w http.ResponseWriter, r *http.Request){
		func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "Bad request ", 404)
			fmt.Fprintln(w, "First response")
		},
		func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "Bad request ", 404)
			fmt.Fprintln(w, "Second response")
		},
		func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "Good request ", 200)
			fmt.Fprintln(w, "Third response")
		},
	}
	responseCounter := 0
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		responses[responseCounter](w, r)
		responseCounter++
	}))
	defer ts.Close()
	response, err := getRegionForFederationClient(&http.Client{}, ts.URL)

	assert.NoError(t, err)
	assert.NotEmpty(t, response)
	assert.NotNil(t, response)
}

func TestInstancePrincipalKeyProvider_getRegionForFederationClientNotFoundRetryFailure(t *testing.T) {
	t.Parallel()
	responses := []func(w http.ResponseWriter, r *http.Request){
		func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintln(w, "First response")
		},
		func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintln(w, "Second response")
		},
		func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintln(w, "Third response")
		},
		func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintln(w, "Fourth response")
		},
		func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintln(w, "Fifth response")
		},
		func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintln(w, "Sixth response")
		},
		func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintln(w, "Seventh response")
		},
		func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintln(w, "Eighth response")
		},
		func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintln(w, "Ninth response")
		},
	}
	responseCounter := 0
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Bad request ", 404)
		responses[responseCounter](w, r)
		responseCounter++
	}))
	defer ts.Close()
	response, err := getRegionForFederationClient(&http.Client{}, ts.URL)

	assert.Error(t, err)
	assert.Empty(t, response)
	assert.NotNil(t, response)
}

func TestInstancePrincipalKeyProvider_getRegionForFederationClientRetrySuccess(t *testing.T) {
	t.Parallel()
	statusCodeList := []int{400, 401, 403, 405, 408, 409, 412, 413, 422, 429, 431, 500, 501, 503}
	for _, statusCode := range statusCodeList {
		responses := []func(w http.ResponseWriter, r *http.Request){
			func(w http.ResponseWriter, r *http.Request) {
				http.Error(w, "Bad request ", statusCode)
				fmt.Fprintln(w, "First response")
			},
			func(w http.ResponseWriter, r *http.Request) {
				http.Error(w, "Bad request ", statusCode)
				fmt.Fprintln(w, "Second response")
			},
			func(w http.ResponseWriter, r *http.Request) {
				http.Error(w, "Good request - Third ", 200)
				fmt.Fprintln(w, "Third response")
			},
		}
		responseCounter := 0
		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			responses[responseCounter](w, r)
			responseCounter++
		}))
		defer ts.Close()

		response, err := getRegionForFederationClient(&http.Client{}, ts.URL)

		assert.NoError(t, err)
		assert.NotEmpty(t, response)
		assert.NotNil(t, response)
	}
}

func TestInstancePrincipalKeyProvider_getRegionForFederationClientRetryFailure(t *testing.T) {
	t.Parallel()
	statusCodeList := []int{400, 401, 403, 405, 408, 409, 412, 413, 422, 429, 431, 500, 501, 503}
	responses := []func(w http.ResponseWriter, r *http.Request){
		func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintln(w, "First response")
		},
		func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintln(w, "Second response")
		},
		func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintln(w, "Third response")
		},
		func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintln(w, "Fourth response")
		},
		func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintln(w, "Fifth response")
		},
		func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintln(w, "Sixth response")
		},
		func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintln(w, "Seventh response")
		},
		func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintln(w, "Eighth response")
		},
	}
	var wg sync.WaitGroup
	for _, statusCode := range statusCodeList {
		wg.Add(1)
		go func() {
			responseCounter := 0
			ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				http.Error(w, "Bad request ", statusCode)
				responses[responseCounter](w, r)
				responseCounter++
			}))
			defer ts.Close()

			response, err := getRegionForFederationClient(&http.Client{}, ts.URL)
			assert.Error(t, err)
			assert.Empty(t, response)
			wg.Done()
		}()
	}
	wg.Wait()
}

func TestInstancePrincipalKeyProvider_getRegionForFederationClientInternalServerError(t *testing.T) {
	t.Parallel()
	regionServer := httptest.NewServer(http.HandlerFunc(internalServerError))
	defer regionServer.Close()

	_, err := getRegionForFederationClient(&http.Client{}, regionServer.URL)

	assert.Error(t, err)
}

func TestInstancePrincipalKeyProvider_RegionForFederationClient(t *testing.T) {
	t.Parallel()
	expectedRegion := common.StringToRegion("sea")
	keyProvider := &instancePrincipalKeyProvider{Region: expectedRegion}
	returnedRegion := keyProvider.RegionForFederationClient()
	assert.Equal(t, returnedRegion, expectedRegion)
}

func TestInstancePrincipalKeyProvider_PrivateRSAKey(t *testing.T) {
	t.Parallel()
	mockFederationClient := new(mockFederationClient)
	expectedPrivateKey := new(rsa.PrivateKey)
	mockFederationClient.On("PrivateKey").Return(expectedPrivateKey, nil).Once()

	keyProvider := &instancePrincipalKeyProvider{FederationClient: mockFederationClient}

	actualPrivateKey, err := keyProvider.PrivateRSAKey()

	assert.NoError(t, err)
	assert.Equal(t, expectedPrivateKey, actualPrivateKey)
	mockFederationClient.AssertExpectations(t)
}

func TestInstancePrincipalKeyProvider_PrivateRSAKeyError(t *testing.T) {
	t.Parallel()
	mockFederationClient := new(mockFederationClient)
	var nilPtr *rsa.PrivateKey
	expectedErrorMessage := "TestPrivateRSAKeyError"
	mockFederationClient.On("PrivateKey").Return(nilPtr, fmt.Errorf("%s", expectedErrorMessage)).Once()

	keyProvider := &instancePrincipalKeyProvider{FederationClient: mockFederationClient}

	actualPrivateKey, actualError := keyProvider.PrivateRSAKey()

	assert.Nil(t, actualPrivateKey)
	assert.Equal(t, strings.Contains(actualError.Error(), fmt.Sprintf("failed to get private key: %s", expectedErrorMessage)), true)
	mockFederationClient.AssertExpectations(t)
}

func TestInstancePrincipalKeyProvider_KeyID(t *testing.T) {
	t.Parallel()
	mockFederationClient := new(mockFederationClient)
	mockFederationClient.On("SecurityToken").Return("TestSecurityTokenString", nil).Once()

	keyProvider := &instancePrincipalKeyProvider{FederationClient: mockFederationClient}

	actualKeyID, err := keyProvider.KeyID()

	assert.NoError(t, err)
	assert.Equal(t, "ST$TestSecurityTokenString", actualKeyID)
}

func TestInstancePrincipalKeyProvider_KeyIDError(t *testing.T) {
	t.Parallel()
	mockFederationClient := new(mockFederationClient)
	expectedErrorMessage := "TestSecurityTokenError"
	mockFederationClient.On("SecurityToken").Return("", fmt.Errorf("%s", expectedErrorMessage)).Once()

	keyProvider := &instancePrincipalKeyProvider{FederationClient: mockFederationClient}

	actualKeyID, actualError := keyProvider.KeyID()

	assert.Equal(t, "", actualKeyID)
	assert.Equal(t, strings.Contains(actualError.Error(), fmt.Sprintf("failed to get security token: %s", expectedErrorMessage)), true)
	mockFederationClient.AssertExpectations(t)
}

type mockFederationClient struct {
	mock.Mock
}

func (m *mockFederationClient) PrivateKey() (*rsa.PrivateKey, error) {
	args := m.Called()
	return args.Get(0).(*rsa.PrivateKey), args.Error(1)
}

func (m *mockFederationClient) SecurityToken() (string, error) {
	args := m.Called()
	return args.String(0), args.Error(1)
}

func (m *mockFederationClient) GetClaim(key string) (interface{}, error) {
	args := m.Called(key)
	return args.Get(0), args.Error(1)
}
