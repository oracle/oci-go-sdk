// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.

package common

import (
	"bytes"
	"context"
	"math"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// testing resource for mocking responses
type mockedResponse struct {
	RawResponse *http.Response
}

// HTTPResponse implements the OCIResponse interface
func (response mockedResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

func getMockedOCIOperationResponse(statusCode int, attemptNumber uint) OCIOperationResponse {
	httpResponse := http.Response{
		Header:     http.Header{},
		StatusCode: statusCode,
	}
	response := mockedResponse{
		RawResponse: &httpResponse,
	}
	return NewOCIOperationResponse(response, nil, attemptNumber)
}

func getMockedOCIOperationResponseWithError(status int, statusCode string) OCIOperationResponse {
	httpResponse := http.Response{
		Header:     http.Header{},
		StatusCode: status,
	}
	response := mockedResponse{
		RawResponse: &httpResponse,
	}
	err := servicefailure{
		StatusCode: status,
		Code:       statusCode,
	}
	return NewOCIOperationResponse(response, err, uint(1))
}

func getExponentialBackoffRetryPolicy(attempts uint) RetryPolicy {
	shouldRetry := func(OCIOperationResponse) bool {
		return true
	}
	nextDuration := func(response OCIOperationResponse) time.Duration {
		return time.Duration(math.Pow(float64(2), float64(response.AttemptNumber-1))) * time.Second
	}
	return NewRetryPolicy(attempts, shouldRetry, nextDuration)
}

func TestNoRetryPolicyDefaults(t *testing.T) {
	response := getMockedOCIOperationResponse(200, 1)
	policy := NoRetryPolicy()
	assert.False(t, policy.ShouldRetryOperation(response))
}

func TestShouldContinueIssuingRequests(t *testing.T) {
	assert.True(t, shouldContinueIssuingRequests(uint(1), uint(2)))
	assert.True(t, shouldContinueIssuingRequests(uint(2), uint(2)))
	assert.True(t, shouldContinueIssuingRequests(uint(150), UnlimitedNumAttemptsValue))
}

func TestRetryPolicyExponentialBackoffNextDurationUnrolled(t *testing.T) {
	responses := []OCIOperationResponse{
		getMockedOCIOperationResponse(500, 1),
		getMockedOCIOperationResponse(500, 2),
		getMockedOCIOperationResponse(500, 3),
		getMockedOCIOperationResponse(500, 4),
		getMockedOCIOperationResponse(500, 5),
	}
	policy := getExponentialBackoffRetryPolicy(5)
	// unroll an exponential retry policy with a specified maximum
	// number of attempts so it's more obvious what's happening
	// request #1
	assert.True(t, shouldContinueIssuingRequests(1, policy.MaximumNumberAttempts))
	assert.True(t, policy.ShouldRetryOperation(responses[0]))
	assert.Equal(t, 1*time.Second, policy.NextDuration(responses[0]))
	// request #2
	assert.True(t, shouldContinueIssuingRequests(2, policy.MaximumNumberAttempts))
	assert.True(t, policy.ShouldRetryOperation(responses[1]))
	assert.Equal(t, 2*time.Second, policy.NextDuration(responses[1]))
	// request #3
	assert.True(t, shouldContinueIssuingRequests(3, policy.MaximumNumberAttempts))
	assert.True(t, policy.ShouldRetryOperation(responses[2]))
	assert.Equal(t, 4*time.Second, policy.NextDuration(responses[2]))
	// request #4
	assert.True(t, shouldContinueIssuingRequests(4, policy.MaximumNumberAttempts))
	assert.True(t, policy.ShouldRetryOperation(responses[3]))
	assert.Equal(t, 8*time.Second, policy.NextDuration(responses[3]))
	// request #5
	assert.True(t, shouldContinueIssuingRequests(5, policy.MaximumNumberAttempts))
	assert.True(t, policy.ShouldRetryOperation(responses[4]))
	assert.Equal(t, 16*time.Second, policy.NextDuration(responses[4]))
	// done
	assert.False(t, shouldContinueIssuingRequests(6, policy.MaximumNumberAttempts))
}

func TestDefaultRetryPolicy(t *testing.T) {
	responses := []OCIOperationResponse{
		getMockedOCIOperationResponseWithError(400, "CannotParseRequest"),
		getMockedOCIOperationResponseWithError(400, "InvalidParameter"),
		getMockedOCIOperationResponseWithError(400, "MissingParameter"),
		getMockedOCIOperationResponseWithError(400, "QuotaExceeded"),
		getMockedOCIOperationResponseWithError(400, "LimitExceeded"),
		getMockedOCIOperationResponseWithError(400, "RelatedResourceNotAuthorizedOrNotFound"),
		getMockedOCIOperationResponseWithError(401, "NotAuthenticated"),
		getMockedOCIOperationResponseWithError(403, "SignUpRequired"),
		getMockedOCIOperationResponseWithError(403, "NotAllowed"),
		getMockedOCIOperationResponseWithError(403, "NotAuthorized"),

		getMockedOCIOperationResponseWithError(404, "NotFound"),
		getMockedOCIOperationResponseWithError(404, "InvalidParameter"),
		getMockedOCIOperationResponseWithError(404, "NotAuthorizedOrNotFound"),
		getMockedOCIOperationResponseWithError(405, "MethodNotAllowed"),
		getMockedOCIOperationResponseWithError(409, "NotAuthorizedOrResourceAlreadyExists"),
		getMockedOCIOperationResponseWithError(409, "InvalidatedRetryToken"),
		getMockedOCIOperationResponseWithError(409, "IncorrectState"),
		getMockedOCIOperationResponseWithError(409, "Conflict"),
		getMockedOCIOperationResponseWithError(412, "NoEtagMatch"),
		getMockedOCIOperationResponseWithError(413, "PayloadTooLarge"),

		getMockedOCIOperationResponseWithError(422, "UnprocessableEntity"),
		getMockedOCIOperationResponseWithError(429, "TooManyRequests"),
		getMockedOCIOperationResponseWithError(431, "RequestHeaderFieldsTooLarge"),
		getMockedOCIOperationResponseWithError(500, "InternalServerError"),
		getMockedOCIOperationResponseWithError(500, "OutOfCapacity"),
		getMockedOCIOperationResponseWithError(501, "MethodNotImplemented"),
		getMockedOCIOperationResponseWithError(503, "ServiceUnavailable"),
		getMockedOCIOperationResponseWithError(599, "Unknown 500 Error"),

		getMockedOCIOperationResponse(200, 1),
	}
	policy := DefaultRetryPolicy()
	// unroll an exponential retry policy with a specified maximum
	// number of attempts so it's more obvious what's happening
	assert.True(t, shouldContinueIssuingRequests(1, policy.MaximumNumberAttempts))
	assert.True(t, shouldContinueIssuingRequests(2, policy.MaximumNumberAttempts))
	assert.True(t, shouldContinueIssuingRequests(3, policy.MaximumNumberAttempts))
	assert.True(t, shouldContinueIssuingRequests(4, policy.MaximumNumberAttempts))
	assert.True(t, shouldContinueIssuingRequests(5, policy.MaximumNumberAttempts))
	assert.True(t, shouldContinueIssuingRequests(6, policy.MaximumNumberAttempts))
	assert.True(t, shouldContinueIssuingRequests(7, policy.MaximumNumberAttempts))
	assert.True(t, shouldContinueIssuingRequests(8, policy.MaximumNumberAttempts))

	assert.False(t, shouldContinueIssuingRequests(9, policy.MaximumNumberAttempts))

	// request #1
	assert.False(t, policy.ShouldRetryOperation(responses[0]))
	// request #2
	assert.False(t, policy.ShouldRetryOperation(responses[1]))
	// request #3
	assert.False(t, policy.ShouldRetryOperation(responses[2]))
	// request #4
	assert.False(t, policy.ShouldRetryOperation(responses[3]))
	// request #5
	assert.False(t, policy.ShouldRetryOperation(responses[4]))
	// request #6
	assert.False(t, policy.ShouldRetryOperation(responses[5]))
	// request #7
	assert.False(t, policy.ShouldRetryOperation(responses[6]))
	// request #8
	assert.False(t, policy.ShouldRetryOperation(responses[7]))
	// request #9
	assert.False(t, policy.ShouldRetryOperation(responses[8]))
	// request #10
	assert.False(t, policy.ShouldRetryOperation(responses[9]))

	// request #11
	assert.False(t, policy.ShouldRetryOperation(responses[10]))
	// request #12
	assert.False(t, policy.ShouldRetryOperation(responses[11]))
	// request #13
	assert.False(t, policy.ShouldRetryOperation(responses[12]))
	// request #14
	assert.False(t, policy.ShouldRetryOperation(responses[13]))
	// request #15
	assert.False(t, policy.ShouldRetryOperation(responses[14]))
	// request #16
	assert.False(t, policy.ShouldRetryOperation(responses[15]))
	// request #17
	assert.True(t, policy.ShouldRetryOperation(responses[16]))
	// request #18
	assert.False(t, policy.ShouldRetryOperation(responses[17]))
	// request #19
	assert.False(t, policy.ShouldRetryOperation(responses[18]))
	// request #20
	assert.False(t, policy.ShouldRetryOperation(responses[19]))

	// request #21
	assert.False(t, policy.ShouldRetryOperation(responses[20]))
	// request #22
	assert.True(t, policy.ShouldRetryOperation(responses[21]))
	// request #23
	assert.False(t, policy.ShouldRetryOperation(responses[22]))
	// request #24
	assert.True(t, policy.ShouldRetryOperation(responses[23]))
	// request #25
	assert.True(t, policy.ShouldRetryOperation(responses[24]))
	// request #26
	assert.False(t, policy.ShouldRetryOperation(responses[25]))
	// request #27
	assert.True(t, policy.ShouldRetryOperation(responses[26]))
	// request #28
	assert.True(t, policy.ShouldRetryOperation(responses[27]))

	// request #29
	assert.False(t, policy.ShouldRetryOperation(responses[28]))
}

type mockedRequest struct {
	Request http.Request
	Policy  *RetryPolicy
}

func (m mockedRequest) HTTPRequest(method, path string, binaryRequestBody *OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {
	return m.Request, nil
}

func (m mockedRequest) RetryPolicy() *RetryPolicy {
	return m.Policy
}

func (m mockedRequest) BinaryRequestBody() (*OCIReadSeekCloser, bool) {
	return NewOCIReadSeekCloser(nil), false
}

func TestRetryTokenPersists(t *testing.T) {
	body := bytes.NewBufferString("YES")
	req, _ := http.NewRequest("POST", "/some", body)
	token := RetryToken()
	req.Header.Set(requestHeaderOpcRetryToken, token)
	policy := getExponentialBackoffRetryPolicy(2)
	r := mockedRequest{Request: *req, Policy: &policy}
	operation := func(i context.Context, request OCIRequest, binaryRequestBody *OCIReadSeekCloser, extraHeaders map[string]string) (OCIResponse, error) {
		httpResponse := http.Response{
			Header:     http.Header{},
			StatusCode: 200,
		}
		httpReq, _ := request.HTTPRequest("POST", "/some", NewOCIReadSeekCloser(nil), nil)
		headerToken := httpReq.Header.Get(requestHeaderOpcRetryToken)

		assert.Equal(t, token, headerToken)
		return mockedResponse{RawResponse: &httpResponse}, nil
	}

	Retry(context.Background(), r, operation, *r.Policy)
}
func TestRetryWithPanicInOperation(t *testing.T) {
	body := bytes.NewBufferString("YES")
	req, _ := http.NewRequest("POST", "/some", body)
	token := RetryToken()
	req.Header.Set(requestHeaderOpcRetryToken, token)
	policy := getExponentialBackoffRetryPolicy(3)
	r := mockedRequest{Request: *req, Policy: &policy}
	times := 0
	operation := func(i context.Context, request OCIRequest, binaryRequestBody *OCIReadSeekCloser, extraHeaders map[string]string) (OCIResponse, error) {
		httpResponse := http.Response{
			Header:     http.Header{},
			StatusCode: 200,
		}

		if times <= 0 {
			times++
			return mockedResponse{RawResponse: &httpResponse}, nil
		}
		panic("test panic")
	}

	resp, err := Retry(context.Background(), r, operation, *r.Policy)
	assert.Nil(t, resp)
	assert.Error(t, err)
}
