package common

import (
	"bytes"
	"context"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"net/url"
	"testing"
)

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
	region := REGION_PHX
	c := NewClientForRegion(region)
	assert.Equal(t, defaultUserAgent(), c.UserAgent)
	assert.NotNil(t, c.HttpClient)
	assert.Equal(t, region, c.Region)
	assert.Nil(t, c.Interceptor)
	assert.NotNil(t, c.Signer)

}

func TestClient_customClientForRegion(t *testing.T) {
	host := "http://somehost:9000"
	basePath := "basePath"
	restPath := "somepath"
	userAgent := "suseragent"

	region := REGION_PHX
	c := NewClientForRegion(region)
	c.Host = host
	c.UserAgent = userAgent
	c.BasePath = basePath
	c.Region = REGION_IAD

	request := http.Request{}
	request.URL = &url.URL{Path: restPath}
	c.prepareRequest(&request)

	assert.Equal(t, userAgent, c.UserAgent)
	assert.NotNil(t, c.HttpClient)
	assert.Equal(t, REGION_IAD, c.Region)
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
	body := `{"key" : "REGION_FRA","name" : "eu-frankfurt-1"}`
	c := NewClientForRegion(REGION_IAD)
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
	c.HttpClient = caller

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
	c := NewClientForRegion(REGION_IAD)
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
	c.HttpClient = caller

	request := http.Request{}
	request.URL = &url.URL{Path: restPath}
	retRes, err := c.Call(context.Background(), &request)
	assert.Error(t, err)
	assert.Equal(t, &response, retRes)

}
