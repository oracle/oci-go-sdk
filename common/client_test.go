package common

import (
	"bytes"
	"context"
	"crypto/rsa"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"testing"

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

func TestDefaultHTTPDispatcher_transportNotSet(t *testing.T) {
	client := defaultHTTPDispatcher()

	if client.Transport != nil {
		t.Errorf("Expecting default http transport to be nil")
	}
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
			assert.Contains(t, r.Header.Get("Authorization"), "signature")
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
			assert.Contains(t, r.Header.Get("Authorization"), "signature")
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

func TestBaseClient_CallError(t *testing.T) {
	response := http.Response{
		Header:     http.Header{},
		StatusCode: 400,
	}
	body := `{"code" : "some fake error","message" : "fake error not here"}`
	c := testClientWithRegion(RegionIAD)
	host := "http://somehost:9000"
	basePath := "basePath/"
	restPath := "/somepath"
	caller := fakeCaller{
		Customcall: func(r *http.Request) (*http.Response, error) {
			assert.Equal(t, "somehost:9000", r.URL.Host)
			assert.Equal(t, defaultUserAgent(), r.UserAgent())
			assert.Contains(t, r.Header.Get("Authorization"), "signature")
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
	assert.Error(t, err)
	assert.Equal(t, &response, retRes)

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

	client, err := NewClientWithConfig(configurationProvider)
	assert.NotNil(t, client)
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
