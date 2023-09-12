// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.

package common

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"net/url"
	"os"
	"strings"
	"syscall"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestErrors_ServiceFailureFromResponse(t *testing.T) {
	header := http.Header{}
	opcID := "111"
	header.Set("opc-request-id", opcID)
	sampleResponse := `{"key" : "test"}`

	httpRequest := http.Request{Method: http.MethodGet, URL: &url.URL{}}
	httpResponse := http.Response{Header: header, Request: &httpRequest}
	bodyBuffer := bytes.NewBufferString(sampleResponse)
	httpResponse.Body = ioutil.NopCloser(bodyBuffer)
	httpResponse.StatusCode = 200
	err := newServiceFailureFromResponse(&httpResponse)
	assert.Equal(t, err.(ServiceError).GetOpcRequestID(), opcID)
	assert.Equal(t, strings.Contains(err.Error(), "Error returned by"), true)

	failure, ok := IsServiceError(err)
	assert.Equal(t, ok, true)
	assert.Equal(t, failure.GetHTTPStatusCode(), httpResponse.StatusCode)
}

func TestErrors_FailedToParseJson(t *testing.T) {
	// invalid json
	sampleResponse := `{"key" : test"}`

	httpRequest := http.Request{Method: http.MethodGet, URL: &url.URL{}}
	httpResponse := http.Response{Request: &httpRequest}
	bodyBuffer := bytes.NewBufferString(sampleResponse)
	httpResponse.Body = ioutil.NopCloser(bodyBuffer)
	err := newServiceFailureFromResponse(&httpResponse)

	failure, ok := IsServiceError(err)
	assert.Equal(t, ok, true)
	assert.Equal(t, failure.GetCode(), "BadErrorResponse")
	assert.Equal(t, strings.Contains(failure.GetMessage(), "Failed to parse json from response body due to"), true)
}

func TestComputeOutOfCapacityErrors(t *testing.T) {
	header := http.Header{}
	opcID := "111"
	header.Set("opc-request-id", opcID)
	sampleResponse := `{"message" : "Out of host capacity.","originalMessage" : "Out of capacity for shape VM.Standard2.1 in dedicated pool vmidp_stress_test.","originalMessageTemplate" : "Out of capacity for shape {shape} in dedicated pool {dedicatedPool}.","messageArguments" : {"problemType": "OUT_OF_CAPACITY","shape": "VM.Standard2.1","dedicatedPool": "vmidp_stress_test"}}`
	httpRequest := http.Request{Method: http.MethodGet, URL: &url.URL{}}
	httpResponse := http.Response{Header: header, Request: &httpRequest}
	bodyBuffer := bytes.NewBufferString(sampleResponse)
	httpResponse.Body = ioutil.NopCloser(bodyBuffer)
	httpResponse.StatusCode = 500
	err := newServiceFailureFromResponse(&httpResponse)
	assert.Equal(t, err.(ServiceError).GetMessage(), "Out of host capacity.")
	assert.Equal(t, err.(ServiceErrorLocalizationMessage).GetOriginalMessage(), "Out of capacity for shape VM.Standard2.1 in dedicated pool vmidp_stress_test.")
	assert.Equal(t, err.(ServiceErrorLocalizationMessage).GetOriginalMessageTemplate(), "Out of capacity for shape {shape} in dedicated pool {dedicatedPool}.")
	assert.Equal(t, err.(ServiceError).GetOpcRequestID(), opcID)
	assert.Equal(t, strings.Contains(err.Error(), "Error returned by"), true)

	failure, ok := IsServiceError(err)
	assert.Equal(t, ok, true)
	assert.Equal(t, failure.GetHTTPStatusCode(), httpResponse.StatusCode)
}

func TestPartialComputeOutOfCapacityErrors(t *testing.T) {
	header := http.Header{}
	opcID := "111"
	header.Set("opc-request-id", opcID)
	sampleResponse := `{"message" : "Out of host capacity.", "messageArguments" : {"problemType": "OUT_OF_CAPACITY","shape": "VM.Standard2.1","dedicatedPool": "vmidp_stress_test"}}`
	httpRequest := http.Request{Method: http.MethodGet, URL: &url.URL{}}
	httpResponse := http.Response{Header: header, Request: &httpRequest}
	bodyBuffer := bytes.NewBufferString(sampleResponse)
	httpResponse.Body = ioutil.NopCloser(bodyBuffer)
	httpResponse.StatusCode = 500
	err := newServiceFailureFromResponse(&httpResponse)
	assert.Equal(t, err.(ServiceError).GetMessage(), "Out of host capacity.")
	assert.Equal(t, err.(ServiceErrorLocalizationMessage).GetOriginalMessage(), "")
	assert.Equal(t, err.(ServiceErrorLocalizationMessage).GetOriginalMessageTemplate(), "")
	assert.Equal(t, err.(ServiceError).GetOpcRequestID(), opcID)
	assert.Equal(t, strings.Contains(err.Error(), "Error returned by"), true)

	failure, ok := IsServiceError(err)
	assert.Equal(t, ok, true)
	assert.Equal(t, failure.GetHTTPStatusCode(), httpResponse.StatusCode)
}

func TestNetworkErrors(t *testing.T) {
	netErr := net.OpError{
		Op:  "accept",
		Net: "tcp",
		Err: syscall.ECONNRESET,
	}
	ok := IsNetworkError(&netErr)
	assert.Equal(t, ok, true)

	netErr1 := net.OpError{
		Op:  "write",
		Net: "tcp",
		Err: syscall.ETIMEDOUT,
	}
	success := IsNetworkError(&netErr1)
	assert.Equal(t, success, true)

	netErr2 := net.OpError{
		Op:  "write",
		Net: "tcp",
		Err: fmt.Errorf("net/http: HTTP/1.x transport connection broken"),
	}
	valid := IsNetworkError(&netErr2)
	assert.Equal(t, valid, true)

}

func TestConnectionReset(t *testing.T) {
	go server()

	time.Sleep(3 * time.Second) // wait for server to run

	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		log.Fatal("client", err)
	}

	if _, err := conn.Write([]byte("ab")); err != nil {
		log.Printf("client: %v", err)
	}

	time.Sleep(1 * time.Second) // wait for close on the server side

	data := make([]byte, 1)

	_, resetErr := conn.Read(data)

	success := IsNetworkError(resetErr)

	assert.Equal(t, success, true)

}

func server() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}

	defer listener.Close()

	conn, err := listener.Accept()
	if err != nil {
		log.Fatal("server", err)
		os.Exit(1)
	}
	data := make([]byte, 1)
	if _, err := conn.Read(data); err != nil {
		log.Fatal("server", err)
	}

	conn.Close()
}
