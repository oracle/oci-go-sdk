// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.

package common

import (
	"bytes"
	"context"
	"crypto/rsa"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"testing"

	"time"

	"github.com/stretchr/testify/assert"
)

type customConfig struct {
	Reg Region
}

func (c customConfig) Region() (string, error) {
	return string(c.Reg), nil
}

func (c customConfig) KeyFingerprint() (string, error) {
	return "a/a/a", nil
}
func (c customConfig) UserOCID() (string, error) {
	return "ocid", nil
}
func (c customConfig) TenancyOCID() (string, error) {
	return "ocid1", nil
}
func (c customConfig) PrivateRSAKey() (*rsa.PrivateKey, error) {
	key, _ := PrivateKeyFromBytes([]byte(testPrivateKeyConf), nil)
	return key, nil
}
func (c customConfig) KeyID() (string, error) {
	return "b/b/b", nil
}

func (c customConfig) AuthType() (AuthConfig, error) {
	return AuthConfig{UserPrincipal, false, nil}, nil
}

func testClientWithRegion(r Region) BaseClient {
	p := customConfig{Reg: r}
	c, _ := NewClientWithConfig(p)
	return c
}

func TestClient_prepareRequestDefScheme(t *testing.T) {
	host := "somehost:9000"
	basePath := "basePath"
	restPath := "somepath"

	c := BaseClient{UserAgent: "asdf"}
	c.Host = host
	c.BasePath = basePath

	request := http.Request{}
	request.URL = &url.URL{Path: restPath}
	c.prepareRequest(&request)
	assert.Equal(t, "https", request.URL.Scheme)
	assert.Equal(t, host, request.URL.Host)
}

func TestClient_prepareRequestCanBeCalledMultipleTimes(t *testing.T) {
	host := "somehost:9000"
	basePath := "basePath"
	restPath := "somepath"

	c := BaseClient{UserAgent: "asdf"}
	c.Host = host
	c.BasePath = basePath

	request := http.Request{}
	request.URL = &url.URL{Path: restPath}
	c.prepareRequest(&request)
	assert.Equal(t, "https", request.URL.Scheme)
	assert.Equal(t, "/basePath/somepath", request.URL.Path)
	assert.Equal(t, host, request.URL.Host)
	c.prepareRequest(&request)
	assert.Equal(t, "https", request.URL.Scheme)
	assert.Equal(t, "/basePath/somepath", request.URL.Path)
}

func TestClient_prepareRequestUpdatesDateHeader(t *testing.T) {
	host := "somehost:9000"
	basePath := "basePath"
	restPath := "somepath"

	c := BaseClient{UserAgent: "asdf"}
	c.Host = host
	c.BasePath = basePath

	request := http.Request{}
	request.URL = &url.URL{Path: restPath}
	c.prepareRequest(&request)
	d1 := request.Header.Get(requestHeaderDate)
	// make sure we wait some time to see that d1 and d2 have different times (set at second-level granularity)
	time.Sleep(2 * time.Second)
	c.prepareRequest(&request)
	d2 := request.Header.Get(requestHeaderDate)
	assert.NotEqual(t, d1, d2)
}

func TestClient_prepareRequestSetScheme(t *testing.T) {
	host := "http://somehost:9000"
	basePath := "basePath"
	restPath := "somepath"

	c := BaseClient{UserAgent: "asdf"}
	c.Host = host
	c.BasePath = basePath

	request := http.Request{}
	request.URL = &url.URL{Path: restPath}
	c.prepareRequest(&request)
	assert.Equal(t, "http", request.URL.Scheme)
	assert.Equal(t, "somehost:9000", request.URL.Host)
}

func TestClient_containsUserAgent(t *testing.T) {
	host := "http://somehost:9000"
	basePath := "basePath"
	restPath := "somepath"
	userAgent := "myuserAgent"

	c := BaseClient{}
	c.Host = host
	c.BasePath = basePath
	c.UserAgent = userAgent

	request := http.Request{}
	request.URL = &url.URL{Path: restPath}
	c.prepareRequest(&request)
	assert.Equal(t, userAgent, request.UserAgent())
}

func TestClient_userAgentBlank(t *testing.T) {
	host := "http://somehost:9000"
	basePath := "basePath"
	restPath := "somepath"
	userAgent := ""

	c := BaseClient{}
	c.Host = host
	c.BasePath = basePath
	c.UserAgent = userAgent

	request := http.Request{}
	request.URL = &url.URL{Path: restPath}
	e := c.prepareRequest(&request)
	assert.Error(t, e)
}

func TestClient_clientForRegion(t *testing.T) {
	region := RegionPHX
	c := testClientWithRegion(region)
	assert.Equal(t, defaultUserAgent(), c.UserAgent)
	assert.NotNil(t, c.HTTPClient)
	assert.Nil(t, c.Interceptor)
	assert.NotNil(t, c.Signer)

}

func TestClient_customClientForRegion(t *testing.T) {
	host := "http://somehost:9000"
	basePath := "basePath"
	restPath := "somepath"
	userAgent := "suseragent"

	region := RegionPHX
	c := testClientWithRegion(region)
	c.Host = host
	c.UserAgent = userAgent
	c.BasePath = basePath

	request := http.Request{}
	request.URL = &url.URL{Path: restPath}
	c.prepareRequest(&request)

	assert.Equal(t, userAgent, c.UserAgent)
	assert.NotNil(t, c.HTTPClient)
	assert.Nil(t, c.Interceptor)
	assert.NotNil(t, c.Signer)
	assert.Equal(t, "http", request.URL.Scheme)
	assert.Equal(t, "somehost:9000", request.URL.Host)
}

type fakeCaller struct {
	CustomResponse *http.Response
	Customcall     func(r *http.Request) (*http.Response, error)
}

func (f fakeCaller) Do(req *http.Request) (*http.Response, error) {
	if f.CustomResponse != nil {
		return f.CustomResponse, nil
	}
	return f.Customcall(req)
}

func TestBaseClient_Call(t *testing.T) {
	response := http.Response{
		Header:     http.Header{},
		StatusCode: 200,
	}
	body := `{"key" : "RegionFRA","name" : "eu-frankfurt-1"}`
	c := testClientWithRegion(RegionIAD)
	host := "http://somehost:9000"
	basePath := "basePath/"
	restPath := "/somepath"
	caller := fakeCaller{
		Customcall: func(r *http.Request) (*http.Response, error) {
			assert.Equal(t, "somehost:9000", r.URL.Host)
			assert.Equal(t, defaultUserAgent(), r.UserAgent())
			assert.Contains(t, r.Header.Get(requestHeaderAuthorization), "signature")
			assert.Contains(t, r.URL.Path, "basePath/somepath")
			bodyBuffer := bytes.NewBufferString(body)
			response.Body = ioutil.NopCloser(bodyBuffer)
			return &response, nil
		},
	}

	c.Host = host
	c.BasePath = basePath
	c.HTTPClient = caller

	request := http.Request{}
	request.URL = &url.URL{Path: restPath}
	retRes, err := c.Call(context.Background(), &request)
	assert.Equal(t, &response, retRes)
	assert.NoError(t, err)

}

func TestBaseClient_CallWithInterceptor(t *testing.T) {
	response := http.Response{
		Header:     http.Header{},
		StatusCode: 200,
	}
	body := `{"key" : "RegionFRA","name" : "eu-frankfurt-1"}`
	c := testClientWithRegion(RegionIAD)
	c.Interceptor = func(request *http.Request) error {
		request.Header.Set("Custom-Header", "CustomValue")
		return nil
	}
	host := "http://somehost:9000"
	basePath := "basePath/"
	restPath := "/somepath"
	caller := fakeCaller{
		Customcall: func(r *http.Request) (*http.Response, error) {
			assert.Equal(t, "somehost:9000", r.URL.Host)
			assert.Equal(t, defaultUserAgent(), r.UserAgent())
			assert.Contains(t, r.Header.Get(requestHeaderAuthorization), "signature")
			assert.Contains(t, r.URL.Path, "basePath/somepath")
			assert.Equal(t, "CustomValue", r.Header.Get("Custom-Header"))
			bodyBuffer := bytes.NewBufferString(body)
			response.Body = ioutil.NopCloser(bodyBuffer)
			return &response, nil
		},
	}

	c.Host = host
	c.BasePath = basePath
	c.HTTPClient = caller

	request := http.Request{}
	request.URL = &url.URL{Path: restPath}
	retRes, err := c.Call(context.Background(), &request)
	assert.Equal(t, &response, retRes)
	assert.NoError(t, err)

}

type genericOCIResponse struct {
	RawResponse *http.Response
}

type retryableOCIRequest struct {
	retryPolicy *RetryPolicy
}

func (request retryableOCIRequest) HTTPRequest(method, path string, binaryRequestBody *OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {
	r := http.Request{}
	r.Method = method
	r.URL = &url.URL{Path: path}
	return r, nil
}

func (request retryableOCIRequest) BinaryRequestBody() (*OCIReadSeekCloser, bool) {
	return nil, false
}

func (request retryableOCIRequest) RetryPolicy() *RetryPolicy {
	return request.retryPolicy
}

// HTTPResponse implements the OCIResponse interface
func (response genericOCIResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

func TestRetry_NeverGetSuccessfulResponse(t *testing.T) {
	errorResponse := genericOCIResponse{
		RawResponse: &http.Response{
			Header:     http.Header{},
			StatusCode: 400,
		},
	}
	errorMessage := "TestRetryError"
	totalNumberAttempts := uint(5)
	numberOfTimesWeEnterShouldRetry := uint(0)
	numberOfTimesWeEnterGetNextDuration := uint(0)
	shouldRetryOperation := func(operationResponse OCIOperationResponse) bool {
		numberOfTimesWeEnterShouldRetry = numberOfTimesWeEnterShouldRetry + 1
		return true
	}
	getNextDuration := func(operationResponse OCIOperationResponse) time.Duration {
		numberOfTimesWeEnterGetNextDuration = numberOfTimesWeEnterGetNextDuration + 1
		return 0
	}
	retryPolicy := NewRetryPolicy(totalNumberAttempts, shouldRetryOperation, getNextDuration)
	retryableRequest := retryableOCIRequest{
		retryPolicy: &retryPolicy,
	}
	// type OCIOperation func(context.Context, OCIRequest, OCIReadSeekCloser) (OCIResponse, error)
	fakeOperation := func(context.Context, OCIRequest, *OCIReadSeekCloser, map[string]string) (OCIResponse, error) {
		return errorResponse, errors.New(errorMessage)
	}

	response, err := Retry(context.Background(), retryableRequest, fakeOperation, retryPolicy)
	assert.Equal(t, totalNumberAttempts, numberOfTimesWeEnterShouldRetry)
	assert.NotNil(t, response)
	assert.Equal(t, 400, response.HTTPResponse().StatusCode)
	assert.Equal(t, err.Error(), errorMessage)
}

func TestRetry_ImmediatelyGetsSuccessfulResponse(t *testing.T) {
	successResponse := genericOCIResponse{
		RawResponse: &http.Response{
			Header:     http.Header{},
			StatusCode: 200,
		},
	}
	totalNumberAttempts := uint(5)
	numberOfTimesWeEnterShouldRetry := uint(0)
	numberOfTimesWeEnterGetNextDuration := uint(0)
	shouldRetryOperation := func(operationResponse OCIOperationResponse) bool {
		numberOfTimesWeEnterShouldRetry = numberOfTimesWeEnterShouldRetry + 1
		return false
	}
	getNextDuration := func(operationResponse OCIOperationResponse) time.Duration {
		numberOfTimesWeEnterGetNextDuration = numberOfTimesWeEnterGetNextDuration + 1
		return 0
	}
	retryPolicy := NewRetryPolicy(totalNumberAttempts, shouldRetryOperation, getNextDuration)
	retryableRequest := retryableOCIRequest{
		retryPolicy: &retryPolicy,
	}
	// type OCIOperation func(context.Context, OCIRequest, OCIReadSeekCloser) (OCIResponse, error)
	fakeOperation := func(context.Context, OCIRequest, *OCIReadSeekCloser, map[string]string) (OCIResponse, error) {
		return successResponse, nil
	}

	response, err := Retry(context.Background(), retryableRequest, fakeOperation, retryPolicy)
	assert.Equal(t, uint(1), numberOfTimesWeEnterShouldRetry)
	assert.Equal(t, uint(0), numberOfTimesWeEnterGetNextDuration)
	assert.Equal(t, response, successResponse)
	assert.Nil(t, err)
}

func TestRetry_RaisesDeadlineExceededException(t *testing.T) {
	errorResponse := genericOCIResponse{
		RawResponse: &http.Response{
			Header:     http.Header{},
			StatusCode: 400,
		},
	}
	totalNumberAttempts := uint(5)
	numberOfTimesWeEnterShouldRetry := uint(0)
	numberOfTimesWeEnterGetNextDuration := uint(0)
	shouldRetryOperation := func(operationResponse OCIOperationResponse) bool {
		numberOfTimesWeEnterShouldRetry = numberOfTimesWeEnterShouldRetry + 1
		return true
	}
	getNextDuration := func(operationResponse OCIOperationResponse) time.Duration {
		numberOfTimesWeEnterGetNextDuration = numberOfTimesWeEnterGetNextDuration + 1
		return 10 * time.Second
	}
	retryPolicy := NewRetryPolicy(totalNumberAttempts, shouldRetryOperation, getNextDuration)
	retryableRequest := retryableOCIRequest{
		retryPolicy: &retryPolicy,
	}
	// type OCIOperation func(context.Context, OCIReadSeekCloser) (OCIResponse, error)
	fakeOperation := func(context.Context, OCIRequest, *OCIReadSeekCloser, map[string]string) (OCIResponse, error) {
		return errorResponse, nil
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	response, err := Retry(ctx, retryableRequest, fakeOperation, retryPolicy)
	assert.Equal(t, uint(1), numberOfTimesWeEnterShouldRetry)
	assert.Equal(t, uint(1), numberOfTimesWeEnterGetNextDuration)
	assert.Equal(t, response, errorResponse)
	assert.Equal(t, DeadlineExceededByBackoff, err)
}

func TestRetry_GetsSuccessfulResponseAfterMultipleAttempts(t *testing.T) {
	errorResponse := genericOCIResponse{
		RawResponse: &http.Response{
			Header:     http.Header{},
			StatusCode: 400,
		},
	}
	successResponse := genericOCIResponse{
		RawResponse: &http.Response{
			Header:     http.Header{},
			StatusCode: 200,
		},
	}
	totalNumberAttempts := uint(10)
	numberOfTimesWeEnterShouldRetry := uint(0)
	numberOfTimesWeEnterGetNextDuration := uint(0)
	shouldRetryOperation := func(operationResponse OCIOperationResponse) bool {
		numberOfTimesWeEnterShouldRetry = numberOfTimesWeEnterShouldRetry + 1
		return operationResponse.Response.HTTPResponse().StatusCode == 400
	}
	getNextDuration := func(operationResponse OCIOperationResponse) time.Duration {
		numberOfTimesWeEnterGetNextDuration = numberOfTimesWeEnterGetNextDuration + 1
		return 0 * time.Second
	}
	retryPolicy := NewRetryPolicy(totalNumberAttempts, shouldRetryOperation, getNextDuration)
	retryableRequest := retryableOCIRequest{
		retryPolicy: &retryPolicy,
	}
	// type OCIOperation func(context.Context, OCIRequest, OCIReadSeekCloser) (OCIResponse, error)
	fakeOperation := func(context.Context, OCIRequest, *OCIReadSeekCloser, map[string]string) (OCIResponse, error) {
		if numberOfTimesWeEnterShouldRetry < 7 {
			return errorResponse, nil
		}
		return successResponse, nil
	}

	response, err := Retry(context.Background(), retryableRequest, fakeOperation, retryPolicy)
	assert.Equal(t, uint(8), numberOfTimesWeEnterShouldRetry)
	assert.Equal(t, uint(7), numberOfTimesWeEnterGetNextDuration)
	assert.Equal(t, response, successResponse)
	assert.Nil(t, err)
}

func TestRetry_CancelContextWhileSleeping(t *testing.T) {

	shouldRetryOperation := func(res OCIOperationResponse) bool { return true }
	getNextDuration := func(res OCIOperationResponse) time.Duration { return 4 * time.Second }

	errorResponse := genericOCIResponse{
		RawResponse: &http.Response{
			Header:     http.Header{},
			StatusCode: 400,
		},
	}
	pol := NewRetryPolicy(uint(10), shouldRetryOperation, getNextDuration)
	req := retryableOCIRequest{retryPolicy: &pol}
	fakeOperation := func(context.Context, OCIRequest, *OCIReadSeekCloser, map[string]string) (OCIResponse, error) {
		return errorResponse, nil
	}

	ctx, cancelFn := context.WithDeadline(context.Background(), time.Now().Add(10*time.Second))

	//cancel context while sleeping
	go func() {
		time.Sleep(1 * time.Second)
		cancelFn()
	}()

	res, err := Retry(ctx, req, fakeOperation, pol)
	assert.Equal(t, errorResponse, res)
	assert.Equal(t, context.Canceled, err)
}

func TestRetryToken_GenerateMultipleTimes(t *testing.T) {
	token1 := RetryToken()
	token2 := RetryToken()
	assert.NotEqual(t, token1, token2)
}

func TestBaseClient_CreateWithInvalidConfig(t *testing.T) {
	dataTpl := `[DEFAULT]
user=someuser
fingerprint=somefingerprint
key_file=%s
region=us-ashburn-1
`

	keyFile := writeTempFile(testPrivateKeyConf)
	data := fmt.Sprintf(dataTpl, keyFile)
	tmpConfFile := writeTempFile(data)

	defer removeFileFn(tmpConfFile)
	defer removeFileFn(keyFile)

	configurationProvider, _ := ConfigurationProviderFromFile(tmpConfFile, "")

	_, err := NewClientWithConfig(configurationProvider)
	assert.Error(t, err)
}

func TestBaseClient_CreateWithConfig(t *testing.T) {
	dataTpl := `[DEFAULT]
tenancy=sometenancy
user=someuser
fingerprint=somefingerprint
key_file=%s
region=us-ashburn-1
`

	keyFile := writeTempFile(testPrivateKeyConf)
	data := fmt.Sprintf(dataTpl, keyFile)
	tmpConfFile := writeTempFile(data)

	defer removeFileFn(tmpConfFile)
	defer removeFileFn(keyFile)

	configurationProvider, errConf := ConfigurationProviderFromFile(tmpConfFile, "")
	assert.NoError(t, errConf)
	authConfig, err := configurationProvider.AuthType()
	assert.Nil(t, err)
	assert.Equal(t, authConfig.AuthType, UserPrincipal)

	client, err := NewClientWithConfig(configurationProvider)
	assert.NotNil(t, client)
	assert.NoError(t, err)
}

func TestBaseClient_CreateWithAuthConfig(t *testing.T) {
	delegationTokenContent := `fakeContent`
	tokenFile := writeTempFile(delegationTokenContent)
	dataTpl := `[DEFAULT]
tenancy=sometenancy
user=someuser
fingerprint=somefingerprint
key_file=%s
region=us-ashburn-1
authentication_type=instance_principal
delegation_token_file=%s
`

	keyFile := writeTempFile(testPrivateKeyConf)
	data := fmt.Sprintf(dataTpl, keyFile, tokenFile)
	tmpConfFile := writeTempFile(data)

	defer removeFileFn(tmpConfFile)
	defer removeFileFn(keyFile)
	defer removeFileFn(tokenFile)

	configurationProvider, errConf := ConfigurationProviderFromFile(tmpConfFile, "")
	assert.NoError(t, errConf)
	authConfig, err := configurationProvider.AuthType()
	assert.Nil(t, err)
	assert.Equal(t, *authConfig.OboToken, delegationTokenContent)
	assert.Equal(t, authConfig.AuthType, InstancePrincipalDelegationToken)

	client, err := NewClientWithConfig(configurationProvider)
	assert.NotNil(t, client)
	assert.NotNil(t, client.Signer)
	assert.NotNil(t, client.Interceptor)
	assert.NoError(t, err)
}

func TestBaseClient_CreateWithErrorAuthConfig(t *testing.T) {
	dataTpl := `[DEFAULT]
tenancy=sometenancy
user=someuser
fingerprint=somefingerprint
key_file=%s
region=us-ashburn-1
authentication_type=instance_principal
delegation_token_file=/nothingThere
`
	keyFile := writeTempFile(testPrivateKeyConf)
	data := fmt.Sprintf(dataTpl, keyFile)
	tmpConfFile := writeTempFile(data)

	defer removeFileFn(tmpConfFile)
	defer removeFileFn(keyFile)

	configurationProvider, errConf := ConfigurationProviderFromFile(tmpConfFile, "")
	assert.NoError(t, errConf)
	authConfig, err := configurationProvider.AuthType()
	assert.Nil(t, err)
	assert.Nil(t, authConfig.OboToken)
	assert.Equal(t, authConfig.AuthType, UnknownAuthenticationType)

	client, err := NewClientWithConfig(configurationProvider)
	assert.NotNil(t, client)
	assert.NotNil(t, client.Signer)
	assert.Nil(t, client.Interceptor)
	assert.NoError(t, err)
}

func TestBaseClient_CreateWithBadRegion(t *testing.T) {
	dataTpl := `[DEFAULT]
tenancy=sometenancy
user=someuser
fingerprint=somefingerprint
key_file=%s
region=noregion
`

	keyFile := writeTempFile(testPrivateKeyConf)
	data := fmt.Sprintf(dataTpl, keyFile)
	tmpConfFile := writeTempFile(data)

	defer removeFileFn(tmpConfFile)
	defer removeFileFn(keyFile)

	configurationProvider, errConf := ConfigurationProviderFromFile(tmpConfFile, "")
	assert.NoError(t, errConf)

	_, err := NewClientWithConfig(configurationProvider)
	assert.NoError(t, err)
}

func TestCustomProfileClient_CreateWithBadRegion(t *testing.T) {
	dataTpl := `[DEFAULT]
region=someregion
[PROFILE1]
tenancy=sometenancy
user=someuser
fingerprint=somefingerprint
key_file=%s
region=noregion
`

	keyFile := writeTempFile(testPrivateKeyConf)
	data := fmt.Sprintf(dataTpl, keyFile)
	tmpConfFile := writeTempFile(data)

	defer removeFileFn(tmpConfFile)
	defer removeFileFn(keyFile)

	configurationProvider := CustomProfileConfigProvider(tmpConfFile, "PROFILE1")

	_, err := NewClientWithConfig(configurationProvider)
	region, _ := configurationProvider.Region()
	assert.Equal(t, "noregion", region)
	assert.NoError(t, err)
}

func TestBaseClient_CreateWithoutRegion(t *testing.T) {
	dataTpl := `[DEFAULT]
tenancy=sometenancy
user=someuser
fingerprint=somefingerprint
key_file=%s
`

	keyFile := writeTempFile(testPrivateKeyConf)
	data := fmt.Sprintf(dataTpl, keyFile)
	tmpConfFile := writeTempFile(data)

	defer removeFileFn(tmpConfFile)
	defer removeFileFn(keyFile)

	configurationProvider, errConf := ConfigurationProviderFromFile(tmpConfFile, "")
	assert.NoError(t, errConf)

	_, err := NewClientWithConfig(configurationProvider)
	assert.Error(t, err)
}

func TestHomeDir(t *testing.T) {
	h := getHomeFolder()
	_, e := os.Stat(h)
	assert.NoError(t, e)
}

func TestSeekable(t *testing.T) {
	file, err := ioutil.TempFile("", "TEMPFILE")
	assert.NoError(t, err)
	_, err = file.WriteString("Hello World")
	assert.NoError(t, err)
	defer file.Close()
	ocirsc := NewOCIReadSeekCloser(file)
	assert.True(t, ocirsc.Seekable())

	wrappedFile := ioutil.NopCloser(file)
	wrappedOcirsc := NewOCIReadSeekCloser(wrappedFile)
	assert.True(t, wrappedOcirsc.Seekable())

	byteArray := []byte("test for section reader")
	reader := bytes.NewReader(byteArray)

	offset := int64(2)
	sectionLength := int64(5)
	sectionReader := io.NewSectionReader(reader, offset, sectionLength)
	ocirsc = NewOCIReadSeekCloser(ioutil.NopCloser(sectionReader))
	assert.True(t, ocirsc.Seekable())

	ocirsc = NewOCIReadSeekCloser(ioutil.NopCloser(reader))
	assert.True(t, ocirsc.Seekable())

	nilOcirsc := NewOCIReadSeekCloser(nil)
	assert.False(t, nilOcirsc.Seekable())

}

func TestSeek(t *testing.T) {
	file, err := ioutil.TempFile("", "TEMPFILE")
	assert.NoError(t, err)
	_, err = file.WriteString("Hello World1")
	assert.NoError(t, err)
	defer file.Close()
	offset := int64(0)
	ocirsc := NewOCIReadSeekCloser(file)
	curPos, e := ocirsc.Seek(offset, io.SeekCurrent)
	if info, err := file.Stat(); err == nil {
		length := info.Size()
		assert.Equal(t, curPos, length)
	}
	offset = int64(2)
	curPos, e = ocirsc.Seek(offset, io.SeekStart)
	assert.Nil(t, e)
	assert.Equal(t, curPos, offset)

	byteArray := []byte("test for bytes.reader")
	reader := bytes.NewReader(byteArray)
	ocirsc = NewOCIReadSeekCloser(ioutil.NopCloser(reader))
	offset = int64(6)
	curPos, e = ocirsc.Seek(offset, io.SeekCurrent)
	assert.Nil(t, e)
	assert.Equal(t, curPos, offset)

	offset = int64(0)
	curPos, e = ocirsc.Seek(offset, io.SeekStart)
	assert.Nil(t, e)
	assert.Equal(t, curPos, offset)
}
