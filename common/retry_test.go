// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.

package common

import (
	"bytes"
	"context"
	"io"
	"math"
	"net/http"
	"reflect"
	"strings"
	"sync"
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
	now := time.Now().Round(0)
	return NewOCIOperationResponseExtended(response, nil, attemptNumber, &now, 1.0, now)
}

func getMockedOCIOperationResponseWithError(status int, statusCode string) OCIOperationResponse {
	return getMockedOCIOperationResponseWithErrorFull(status, statusCode, (*time.Time)(nil), 1.0)
}

func getMockedOCIOperationResponseWithErrorFull(status int, statusCode string, endOfWindowTime *time.Time, backoffScalingFactor float64) OCIOperationResponse {
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
	return NewOCIOperationResponseExtended(response, err, uint(1), endOfWindowTime, backoffScalingFactor, time.Now().Round(0))
}

func getExponentialBackoffRetryPolicy(attempts uint) RetryPolicy {
	// we don't want jitter here
	exponentialBackoff := func(r OCIOperationResponse) time.Duration {
		return time.Duration(math.Pow(float64(2), float64(r.AttemptNumber-1))) * time.Second
	}
	return NewRetryPolicyWithOptions(
		ReplaceWithValuesFromRetryPolicy(getExponentialBackoffRetryPolicyWithJitter(attempts)),
		WithNextDuration(exponentialBackoff))
}

func getExponentialBackoffRetryPolicyWithJitter(attempts uint) RetryPolicy {
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

func TestRetryPolicyExponentialBackoffNextDurationUnrolledWithJitter(t *testing.T) {
	responses := []OCIOperationResponse{
		getMockedOCIOperationResponse(500, 1),
		getMockedOCIOperationResponse(500, 2),
		getMockedOCIOperationResponse(500, 3),
		getMockedOCIOperationResponse(500, 4),
		getMockedOCIOperationResponse(500, 5),
	}
	policy := getExponentialBackoffRetryPolicyWithJitter(5)
	// unroll an exponential retry policy with a specified maximum
	// number of attempts so it's more obvious what's happening
	// request #1
	assert.True(t, shouldContinueIssuingRequests(1, policy.MaximumNumberAttempts))
	assert.True(t, policy.ShouldRetryOperation(responses[0]))
	assert.True(t, 1*time.Second <= policy.NextDuration(responses[0]) && policy.NextDuration(responses[0]) < 2*time.Second)
	// request #2
	assert.True(t, shouldContinueIssuingRequests(2, policy.MaximumNumberAttempts))
	assert.True(t, policy.ShouldRetryOperation(responses[1]))
	assert.True(t, 2*time.Second <= policy.NextDuration(responses[1]) && policy.NextDuration(responses[1]) < 3*time.Second)
	// request #3
	assert.True(t, shouldContinueIssuingRequests(3, policy.MaximumNumberAttempts))
	assert.True(t, policy.ShouldRetryOperation(responses[2]))
	assert.True(t, 4*time.Second <= policy.NextDuration(responses[2]) && policy.NextDuration(responses[2]) < 5*time.Second)
	// request #4
	assert.True(t, shouldContinueIssuingRequests(4, policy.MaximumNumberAttempts))
	assert.True(t, policy.ShouldRetryOperation(responses[3]))
	assert.True(t, 8*time.Second <= policy.NextDuration(responses[3]) && policy.NextDuration(responses[3]) < 9*time.Second)
	// request #5
	assert.True(t, shouldContinueIssuingRequests(5, policy.MaximumNumberAttempts))
	assert.True(t, policy.ShouldRetryOperation(responses[4]))
	assert.True(t, 16*time.Second <= policy.NextDuration(responses[4]) && policy.NextDuration(responses[4]) < 17*time.Second)
	// done
	assert.False(t, shouldContinueIssuingRequests(6, policy.MaximumNumberAttempts))
}

func buildResponsesNoRetry(endOfWindowTime *time.Time, backoffScalingFactor float64) []OCIOperationResponse {
	responses := []OCIOperationResponse{
		getMockedOCIOperationResponseWithErrorFull(400, "CannotParseRequest", endOfWindowTime, backoffScalingFactor),
		getMockedOCIOperationResponseWithErrorFull(400, "InvalidParameter", endOfWindowTime, backoffScalingFactor),
		getMockedOCIOperationResponseWithErrorFull(400, "MissingParameter", endOfWindowTime, backoffScalingFactor),
		getMockedOCIOperationResponseWithErrorFull(400, "QuotaExceeded", endOfWindowTime, backoffScalingFactor),
		getMockedOCIOperationResponseWithErrorFull(400, "LimitExceeded", endOfWindowTime, backoffScalingFactor),
		getMockedOCIOperationResponseWithErrorFull(400, "RelatedResourceNotAuthorizedOrNotFound", endOfWindowTime, backoffScalingFactor),
		getMockedOCIOperationResponseWithErrorFull(400, "InsufficientServicePermissions", endOfWindowTime, backoffScalingFactor),
		getMockedOCIOperationResponseWithErrorFull(401, "NotAuthenticated", endOfWindowTime, backoffScalingFactor),
		getMockedOCIOperationResponseWithErrorFull(403, "SignUpRequired", endOfWindowTime, backoffScalingFactor),
		getMockedOCIOperationResponseWithErrorFull(403, "NotAllowed", endOfWindowTime, backoffScalingFactor),
		getMockedOCIOperationResponseWithErrorFull(403, "NotAuthorized", endOfWindowTime, backoffScalingFactor),
		getMockedOCIOperationResponseWithErrorFull(404, "NotFound", endOfWindowTime, backoffScalingFactor),
		getMockedOCIOperationResponseWithErrorFull(404, "InvalidParameter", endOfWindowTime, backoffScalingFactor),
		getMockedOCIOperationResponseWithErrorFull(404, "NotAuthorizedOrNotFound", endOfWindowTime, backoffScalingFactor),
		getMockedOCIOperationResponseWithErrorFull(405, "MethodNotAllowed", endOfWindowTime, backoffScalingFactor),
		getMockedOCIOperationResponseWithErrorFull(409, "NotAuthorizedOrResourceAlreadyExists", endOfWindowTime, backoffScalingFactor),
		getMockedOCIOperationResponseWithErrorFull(409, "InvalidatedRetryToken", endOfWindowTime, backoffScalingFactor),
		getMockedOCIOperationResponseWithErrorFull(412, "NoEtagMatch", endOfWindowTime, backoffScalingFactor),
		getMockedOCIOperationResponseWithErrorFull(413, "PayloadTooLarge", endOfWindowTime, backoffScalingFactor),
		getMockedOCIOperationResponseWithErrorFull(422, "UnprocessableEntity", endOfWindowTime, backoffScalingFactor),
		getMockedOCIOperationResponseWithErrorFull(431, "RequestHeaderFieldsTooLarge", endOfWindowTime, backoffScalingFactor),
		getMockedOCIOperationResponseWithErrorFull(501, "MethodNotImplemented", endOfWindowTime, backoffScalingFactor),
		getMockedOCIOperationResponseWithErrorFull(599, "Unknown 500 Error", endOfWindowTime, backoffScalingFactor),

		getMockedOCIOperationResponse(200, 1),
	}
	return responses
}

func buildResponsesWantRetry(endOfWindowTime *time.Time, backoffScalingFactor float64) []OCIOperationResponse {
	responses := []OCIOperationResponse{
		getMockedOCIOperationResponseWithErrorFull(409, "IncorrectState", endOfWindowTime, backoffScalingFactor),
		getMockedOCIOperationResponseWithErrorFull(429, "TooManyRequests", endOfWindowTime, backoffScalingFactor),
		getMockedOCIOperationResponseWithErrorFull(500, "InternalServerError", endOfWindowTime, backoffScalingFactor),
		getMockedOCIOperationResponseWithErrorFull(500, "OutOfCapacity", endOfWindowTime, backoffScalingFactor),
		getMockedOCIOperationResponseWithErrorFull(503, "ServiceUnavailable", endOfWindowTime, backoffScalingFactor),
	}
	return responses
}

func buildEcResponsesNoRetry(endOfWindowTime *time.Time, backoffScalingFactor float64) []OCIOperationResponse {
	responses := []OCIOperationResponse{
		getMockedOCIOperationResponseWithErrorFull(400, "CannotParseRequest", endOfWindowTime, backoffScalingFactor),
		getMockedOCIOperationResponseWithErrorFull(400, "InvalidParameter", endOfWindowTime, backoffScalingFactor),
		getMockedOCIOperationResponseWithErrorFull(400, "MissingParameter", endOfWindowTime, backoffScalingFactor),
		getMockedOCIOperationResponseWithErrorFull(400, "QuotaExceeded", endOfWindowTime, backoffScalingFactor),
		getMockedOCIOperationResponseWithErrorFull(400, "LimitExceeded", endOfWindowTime, backoffScalingFactor),
		getMockedOCIOperationResponseWithErrorFull(401, "NotAuthenticated", endOfWindowTime, backoffScalingFactor),
		getMockedOCIOperationResponseWithErrorFull(403, "SignUpRequired", endOfWindowTime, backoffScalingFactor),
		getMockedOCIOperationResponseWithErrorFull(403, "NotAllowed", endOfWindowTime, backoffScalingFactor),
		getMockedOCIOperationResponseWithErrorFull(403, "NotAuthorized", endOfWindowTime, backoffScalingFactor),
		getMockedOCIOperationResponseWithErrorFull(404, "NotFound", endOfWindowTime, backoffScalingFactor),
		getMockedOCIOperationResponseWithErrorFull(404, "InvalidParameter", endOfWindowTime, backoffScalingFactor),
		getMockedOCIOperationResponseWithErrorFull(405, "MethodNotAllowed", endOfWindowTime, backoffScalingFactor),
		getMockedOCIOperationResponseWithErrorFull(409, "InvalidatedRetryToken", endOfWindowTime, backoffScalingFactor),
		getMockedOCIOperationResponseWithErrorFull(412, "NoEtagMatch", endOfWindowTime, backoffScalingFactor),
		getMockedOCIOperationResponseWithErrorFull(413, "PayloadTooLarge", endOfWindowTime, backoffScalingFactor),
		getMockedOCIOperationResponseWithErrorFull(422, "UnprocessableEntity", endOfWindowTime, backoffScalingFactor),
		getMockedOCIOperationResponseWithErrorFull(431, "RequestHeaderFieldsTooLarge", endOfWindowTime, backoffScalingFactor),
		getMockedOCIOperationResponseWithErrorFull(501, "MethodNotImplemented", endOfWindowTime, backoffScalingFactor),

		getMockedOCIOperationResponse(200, 1),
	}
	return responses
}

func buildEcResponsesWantRetry(endOfWindowTime *time.Time, backoffScalingFactor float64) []OCIOperationResponse {
	responses := []OCIOperationResponse{
		getMockedOCIOperationResponseWithErrorFull(400, "RelatedResourceNotAuthorizedOrNotFound", endOfWindowTime, backoffScalingFactor),
		getMockedOCIOperationResponseWithErrorFull(404, "NotAuthorizedOrNotFound", endOfWindowTime, backoffScalingFactor),
		getMockedOCIOperationResponseWithErrorFull(409, "IncorrectState", endOfWindowTime, backoffScalingFactor),
		getMockedOCIOperationResponseWithErrorFull(409, "NotAuthorizedOrResourceAlreadyExists", endOfWindowTime, backoffScalingFactor),
		getMockedOCIOperationResponseWithErrorFull(429, "TooManyRequests", endOfWindowTime, backoffScalingFactor),
		getMockedOCIOperationResponseWithErrorFull(500, "InternalServerError", endOfWindowTime, backoffScalingFactor),
		getMockedOCIOperationResponseWithErrorFull(500, "OutOfCapacity", endOfWindowTime, backoffScalingFactor),
		getMockedOCIOperationResponseWithErrorFull(503, "ServiceUnavailable", endOfWindowTime, backoffScalingFactor),
		getMockedOCIOperationResponseWithErrorFull(400, "InsufficientServicePermissions", endOfWindowTime, backoffScalingFactor),
	}
	return responses
}

var (
	responsesNoRetry   = buildResponsesNoRetry((*time.Time)(nil), 1.0)
	responsesWantRetry = buildResponsesWantRetry((*time.Time)(nil), 1.0)
)

func TestDefaultRetryPolicyWithoutEventualConsistency(t *testing.T) {
	policy := DefaultRetryPolicyWithoutEventualConsistency()
	// unroll an exponential retry policy with a specified maximum
	// number of attempts so it's more obvious what's happening
	for i := uint(1); i <= 8; i++ {
		assert.True(t, shouldContinueIssuingRequests(i, policy.MaximumNumberAttempts))
	}
	assert.False(t, shouldContinueIssuingRequests(9, policy.MaximumNumberAttempts))

	for _, r := range responsesNoRetry {
		assert.False(t, policy.ShouldRetryOperation(r))
	}
	for _, r := range responsesWantRetry {
		assert.True(t, policy.ShouldRetryOperation(r))
	}
}

func TestDefaultRetryPolicy(t *testing.T) {
	policy := DefaultRetryPolicy()

	// unroll an exponential retry policy with a specified maximum
	// number of attempts so it's more obvious what's happening
	for i := uint(1); i <= 9; i++ {
		assert.True(t, shouldContinueIssuingRequests(i, policy.MaximumNumberAttempts))
	}
	assert.False(t, shouldContinueIssuingRequests(10, policy.MaximumNumberAttempts))

	for _, r := range responsesNoRetry {
		assert.False(t, policy.ShouldRetryOperation(r))
	}
	for _, r := range responsesWantRetry {
		assert.True(t, policy.ShouldRetryOperation(r))
	}
}

func TestRetryEOF(t *testing.T) {
	policy := DefaultRetryPolicy()
	// unroll an exponential retry policy with a specified maximum
	// number of attempts so it's more obvious what's happening
	for i := uint(1); i <= 9; i++ {
		assert.True(t, shouldContinueIssuingRequests(i, policy.MaximumNumberAttempts))
	}
	assert.False(t, shouldContinueIssuingRequests(10, policy.MaximumNumberAttempts))
	for _, r := range responsesNoRetry {
		assert.False(t, policy.ShouldRetryOperation(r))
	}
	for _, r := range responsesWantRetry {
		r.Error = io.EOF
		assert.True(t, policy.ShouldRetryOperation(r))
	}
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

const (
	eventuallyConsistentWindowSize = time.Duration(4 * time.Minute)
)

func TestGetEndTimeOfEventuallyConsistentWindow(t *testing.T) {
	resetTimes()

	var eowt = EcContext.GetEndOfWindow()
	assert.Equal(t, (*time.Time)(nil), eowt)

	eowt = EcContext.UpdateEndOfWindow(eventuallyConsistentWindowSize)
	assert.NotEqual(t, (*time.Time)(nil), eowt)

	eowt = EcContext.GetEndOfWindow()
	assert.NotEqual(t, (*time.Time)(nil), eowt)

	value1 := EcContext.GetEndOfWindow()
	time.Sleep(1 * time.Second)
	assert.True(t, value1.Before(time.Now().Round(0).Add(eventuallyConsistentWindowSize)))
	value2 := EcContext.GetEndOfWindow()
	assert.Equal(t, value1, value2)

	eowt = EcContext.UpdateEndOfWindow(eventuallyConsistentWindowSize)

	value3 := EcContext.GetEndOfWindow()
	time.Sleep(1 * time.Second)
	assert.True(t, value3.Before(time.Now().Round(0).Add(eventuallyConsistentWindowSize)))
	assert.True(t, value1.Before(*value3))
	assert.True(t, value2.Before(*value3))
}

func TestGetBackoffWithoutJitter(t *testing.T) {
	rp := DefaultRetryPolicyWithoutEventualConsistency()
	assert.Equal(t, time.Duration(1)*time.Second, GetBackoffWithoutJitter(rp, 1))
	assert.Equal(t, time.Duration(2)*time.Second, GetBackoffWithoutJitter(rp, 2))
	assert.Equal(t, time.Duration(4)*time.Second, GetBackoffWithoutJitter(rp, 3))
	assert.Equal(t, time.Duration(8)*time.Second, GetBackoffWithoutJitter(rp, 4))
	assert.Equal(t, time.Duration(16)*time.Second, GetBackoffWithoutJitter(rp, 5))
	assert.Equal(t, time.Duration(30)*time.Second, GetBackoffWithoutJitter(rp, 6))
	assert.Equal(t, time.Duration(30)*time.Second, GetBackoffWithoutJitter(rp, 7))
}

func TestGetMaximumCumulativeBackoffWithoutJitter(t *testing.T) {
	rp := DefaultRetryPolicyWithoutEventualConsistency()
	assert.Equal(t, time.Duration(91)*time.Second, GetMaximumCumulativeBackoffWithoutJitter(rp))
	assert.Equal(t, time.Duration(91)*time.Second, rp.GetMaximumCumulativeBackoffWithoutJitter())
}

func TestGetMaximumCumulativeBackoffWithoutJitter_UnlimitedAttempts(t *testing.T) {
	var rp = NewRetryPolicyWithOptions(
		ReplaceWithValuesFromRetryPolicy(DefaultRetryPolicyWithoutEventualConsistency()),
		WithMaximumNumberAttempts(0),
	)
	var validated, validateError = rp.validate()
	assert.True(t, validated)
	assert.Equal(t, nil, validateError)

	rp = NewRetryPolicyWithOptions(
		ReplaceWithValuesFromRetryPolicy(DefaultRetryPolicyWithoutEventualConsistency()),
		WithUnlimitedAttempts(time.Duration(10)*time.Minute),
	)
	validated, validateError = rp.validate()
	assert.True(t, validated)
	assert.Equal(t, nil, validateError)

	assert.Equal(t, time.Duration(10)*time.Minute, GetMaximumCumulativeBackoffWithoutJitter(rp))
	assert.Equal(t, time.Duration(10)*time.Minute, rp.GetMaximumCumulativeBackoffWithoutJitter())
}

func TestGetEventuallyConsistentBackoffWithoutJitter(t *testing.T) {
	rp := DefaultRetryPolicy()
	assert.Equal(t, time.Duration(1)*time.Second, GetEventuallyConsistentBackoffWithoutJitter(rp, 1, 1.0))
	assert.Equal(t, time.Duration(3520)*time.Millisecond, GetEventuallyConsistentBackoffWithoutJitter(rp, 2, 1.0))
	assert.Equal(t, time.Duration(12390)*time.Millisecond, GetEventuallyConsistentBackoffWithoutJitter(rp, 3, 1.0))
	assert.Equal(t, time.Duration(43614)*time.Millisecond, GetEventuallyConsistentBackoffWithoutJitter(rp, 4, 1.0))
	assert.Equal(t, time.Duration(45)*time.Second, GetEventuallyConsistentBackoffWithoutJitter(rp, 5, 1.0))
	assert.Equal(t, time.Duration(45)*time.Second, GetEventuallyConsistentBackoffWithoutJitter(rp, 6, 1.0))
	assert.Equal(t, time.Duration(45)*time.Second, GetEventuallyConsistentBackoffWithoutJitter(rp, 7, 1.0))
	assert.Equal(t, time.Duration(45)*time.Second, GetEventuallyConsistentBackoffWithoutJitter(rp, 8, 1.0))

	assert.Equal(t, time.Duration(1)*time.Second, GetEventuallyConsistentBackoffWithoutJitter(rp, 1, 0.5))
	assert.Equal(t, time.Duration(2)*time.Second, GetEventuallyConsistentBackoffWithoutJitter(rp, 2, 0.5))
	assert.Equal(t, time.Duration(12390/2)*time.Millisecond, GetEventuallyConsistentBackoffWithoutJitter(rp, 3, 0.5))
	assert.Equal(t, time.Duration(43614/2)*time.Millisecond, GetEventuallyConsistentBackoffWithoutJitter(rp, 4, 0.5))
	assert.Equal(t, time.Duration(45000/2)*time.Millisecond, GetEventuallyConsistentBackoffWithoutJitter(rp, 5, 0.5))
	assert.Equal(t, time.Duration(30)*time.Second, GetEventuallyConsistentBackoffWithoutJitter(rp, 6, 0.5))
	assert.Equal(t, time.Duration(30)*time.Second, GetEventuallyConsistentBackoffWithoutJitter(rp, 7, 0.5))
	assert.Equal(t, time.Duration(30)*time.Second, GetEventuallyConsistentBackoffWithoutJitter(rp, 8, 0.5))
}

func TestGetMaximumCumulativeEventuallyConsistentBackoffWithoutJitter(t *testing.T) {
	rp := DefaultRetryPolicy()
	assert.Equal(t, time.Duration(240524)*time.Millisecond, GetMaximumCumulativeEventuallyConsistentBackoffWithoutJitter(rp))
	assert.Equal(t, time.Duration(240524)*time.Millisecond, rp.GetMaximumCumulativeBackoffWithoutJitter())
}

func TestGetMaximumCumulativeEventuallyConsistentBackoffWithoutJitter_UnlimitedAttempts(t *testing.T) {
	var rp = NewRetryPolicyWithOptions(
		WithMaximumNumberAttempts(0),
	)
	var validated, validateError = rp.validate()
	assert.False(t, validated)
	assert.True(t, strings.Contains(validateError.Error(), "Unlimited"))

	rp = NewRetryPolicyWithOptions(
		WithUnlimitedAttempts(time.Duration(10) * time.Minute),
	)
	validated, validateError = rp.validate()
	assert.True(t, validated)
	assert.Equal(t, nil, validateError)

	assert.Equal(t, time.Duration(10)*time.Minute, GetMaximumCumulativeEventuallyConsistentBackoffWithoutJitter(rp))
	assert.Equal(t, time.Duration(10)*time.Minute, rp.GetMaximumCumulativeBackoffWithoutJitter())

	// non-EC policy has unlimited attempts
	rp = NewRetryPolicyWithOptions(
		ReplaceWithValuesFromRetryPolicy(DefaultRetryPolicyWithoutEventualConsistency()),
		WithMaximumNumberAttempts(0),
	)
	rp = EventuallyConsistentRetryPolicy(rp)
	validated, validateError = rp.validate()
	assert.False(t, validated)
	assert.True(t, strings.Contains(validateError.Error(), "Unlimited"))

	rp = NewRetryPolicyWithOptions(
		ReplaceWithValuesFromRetryPolicy(DefaultRetryPolicyWithoutEventualConsistency()),
		WithUnlimitedAttempts(time.Duration(10)*time.Minute),
	)
	rp = EventuallyConsistentRetryPolicy(rp)
	validated, validateError = rp.validate()
	assert.True(t, validated)
	assert.Equal(t, nil, validateError)

	assert.Equal(t, time.Duration(10)*time.Minute, GetMaximumCumulativeEventuallyConsistentBackoffWithoutJitter(*rp.NonEventuallyConsistentPolicy))
	assert.Equal(t, time.Duration(10)*time.Minute, rp.NonEventuallyConsistentPolicy.GetMaximumCumulativeBackoffWithoutJitter())
}

func TestDeterminePolicyToUse_DefaultRetryPolicyWithoutEventualConsistency(t *testing.T) {
	resetTimes()

	rp := DefaultRetryPolicyWithoutEventualConsistency()
	var policyToUse, endOfWindowTime, backoffScalingFactor = rp.DeterminePolicyToUse(rp)
	assertEqualRetryPolicies(t, rp, policyToUse)
	assert.Equal(t, (*time.Time)(nil), endOfWindowTime)
	assert.Equal(t, 1.0, backoffScalingFactor)

	currentEndOfWindowTime := EcContext.UpdateEndOfWindow(eventuallyConsistentWindowSize)

	policyToUse, endOfWindowTime, backoffScalingFactor = rp.DeterminePolicyToUse(rp)
	assertEqualRetryPolicies(t, rp, policyToUse)
	assert.Equal(t, currentEndOfWindowTime, endOfWindowTime)
	assert.Equal(t, 1.0, backoffScalingFactor)
}

func TestDeterminePolicyToUse_NoRetryPolicy(t *testing.T) {
	resetTimes()

	rp := NoRetryPolicy()
	var policyToUse, endOfWindowTime, backoffScalingFactor = rp.DeterminePolicyToUse(rp)
	assertEqualRetryPolicies(t, rp, policyToUse)
	assert.Equal(t, (*time.Time)(nil), endOfWindowTime)
	assert.Equal(t, 1.0, backoffScalingFactor)

	currentEndOfWindowTime := EcContext.UpdateEndOfWindow(eventuallyConsistentWindowSize)

	policyToUse, endOfWindowTime, backoffScalingFactor = rp.DeterminePolicyToUse(rp)
	assertEqualRetryPolicies(t, rp, policyToUse)
	assert.Equal(t, currentEndOfWindowTime, endOfWindowTime)
	assert.Equal(t, 1.0, backoffScalingFactor)
}

// EC

func TestDeterminePolicyToUse_EcRetryPolicy_NoEc(t *testing.T) {
	resetTimes()

	drp := DefaultRetryPolicyWithoutEventualConsistency()
	rp := EventuallyConsistentRetryPolicy(drp)

	// no eventually consistent effects
	var policyToUse, endOfWindowTime, backoffScalingFactor = rp.DeterminePolicyToUse(rp)
	assertEqualRetryPolicies(t, drp, policyToUse)
	assert.Equal(t, (*time.Time)(nil), endOfWindowTime)
	assert.Equal(t, 1.0, backoffScalingFactor)
}

func TestDeterminePolicyToUse_EcRetryPolicy_EcEndsSoonerThanDefaultRetries(t *testing.T) {
	resetTimes()

	drp := DefaultRetryPolicyWithoutEventualConsistency()
	rp := EventuallyConsistentRetryPolicy(drp)
	drpWithEc := NewRetryPolicyWithOptions(
		ReplaceWithValuesFromRetryPolicy(drp),
		WithShouldRetryOperation(rp.ShouldRetryOperation))

	now := setupTimesEcEndsSoonerThanDefaultRetries((*time.Time)(nil))

	var policyToUse, endOfWindowTime, backoffScalingFactor = rp.DeterminePolicyToUse(rp)
	assertEqualRetryPolicies(t, drpWithEc, policyToUse)
	assert.Equal(t, now.Add(eventuallyConsistentWindowSize), *endOfWindowTime)
	assert.Equal(t, 1.0, backoffScalingFactor)
}

func TestDeterminePolicyToUse_EcRetryPolicy_NeedEcRetries(t *testing.T) {
	resetTimes()

	rp := DefaultRetryPolicy()

	now := setupTimesNeedEcRetries((*time.Time)(nil), rp)

	var policyToUse, endOfWindowTime, backoffScalingFactor = rp.DeterminePolicyToUse(rp)
	assertEqualRetryPolicies(t, rp, policyToUse)
	assert.Equal(t, now.Add(eventuallyConsistentWindowSize), *endOfWindowTime)
	// since the initial request is exactly half-way through the maximum, expect factor of 0.5
	assert.Equal(t, 0.5, backoffScalingFactor)
}

func TestDeterminePolicyToUse_EcRetryPolicy_NeedFullEcRetries(t *testing.T) {
	resetTimes()

	rp := DefaultRetryPolicy()
	maxCumulativeBackoffWithoutJitter := GetMaximumCumulativeEventuallyConsistentBackoffWithoutJitter(rp)
	assert.Equal(t, maxCumulativeBackoffWithoutJitter, rp.GetMaximumCumulativeBackoffWithoutJitter())

	// EC window ends after default retries, with almost the full max cumulative backoff left
	timeNowProviderValues := make(chan time.Time, 300)
	now := time.Now().Round(0)
	// most recent EC
	timeNowProviderValues <- now
	// initial request
	timeNowProviderValues <- now
	EcContext.timeNowProvider = func() time.Time {
		nextTime := <-timeNowProviderValues
		return nextTime
	}
	EcContext.UpdateEndOfWindow(eventuallyConsistentWindowSize)

	var policyToUse, endOfWindowTime, backoffScalingFactor = rp.DeterminePolicyToUse(rp)
	assertEqualRetryPolicies(t, rp, policyToUse)
	assert.Equal(t, now.Add(eventuallyConsistentWindowSize), *endOfWindowTime)
	// since the initial request is 4 minutes / (4 minutes 0.524 seconds), the factor
	// should be really close to 1
	assert.Equal(t, float64(4*time.Minute)/float64(maxCumulativeBackoffWithoutJitter), backoffScalingFactor)
	assert.True(t, backoffScalingFactor > 0.99)
	assert.True(t, backoffScalingFactor < 1.0)
}

func TestDeterminePolicyToUse_EcRetryPolicy_EndOfWindowInThePast(t *testing.T) {
	resetTimes()

	drp := DefaultRetryPolicyWithoutEventualConsistency()
	rp := EventuallyConsistentRetryPolicy(drp)

	setupTimesEndOfEcWindowInThePast((*time.Time)(nil))

	var policyToUse, endOfWindowTime, backoffScalingFactor = rp.DeterminePolicyToUse(rp)
	assertEqualRetryPolicies(t, drp, policyToUse)
	assert.Equal(t, (*time.Time)(nil), endOfWindowTime)
	assert.Equal(t, 1.0, backoffScalingFactor)
}

// EC with unlimited attempts

func TestDeterminePolicyToUse_EcRetryPolicy_NoEc_UnlimitedAttempts(t *testing.T) {
	resetTimes()

	drp := NewRetryPolicyWithOptions(
		ReplaceWithValuesFromRetryPolicy(DefaultRetryPolicyWithoutEventualConsistency()),
		WithUnlimitedAttempts(time.Duration(10)*time.Minute),
	)
	rp := EventuallyConsistentRetryPolicy(drp)

	// no eventually consistent effects
	var policyToUse, endOfWindowTime, backoffScalingFactor = rp.DeterminePolicyToUse(rp)
	assertEqualRetryPolicies(t, drp, policyToUse)
	assert.Equal(t, (*time.Time)(nil), endOfWindowTime)
	assert.Equal(t, 1.0, backoffScalingFactor)
}

func TestDeterminePolicyToUse_EcRetryPolicy_EcEndsSoonerThanDefaultRetries_UnlimitedAttempts(t *testing.T) {
	resetTimes()

	drp := NewRetryPolicyWithOptions(
		ReplaceWithValuesFromRetryPolicy(DefaultRetryPolicyWithoutEventualConsistency()),
		WithUnlimitedAttempts(time.Duration(10)*time.Minute),
	)
	rp := EventuallyConsistentRetryPolicy(drp)
	drpWithEc := NewRetryPolicyWithOptions(
		ReplaceWithValuesFromRetryPolicy(DefaultRetryPolicyWithoutEventualConsistency()),
		WithUnlimitedAttempts(time.Duration(10)*time.Minute),
		WithShouldRetryOperation(rp.ShouldRetryOperation))

	now := setupTimesEcEndsSoonerThanDefaultRetries((*time.Time)(nil))

	var policyToUse, endOfWindowTime, backoffScalingFactor = rp.DeterminePolicyToUse(rp)
	assertEqualRetryPolicies(t, drpWithEc, policyToUse)
	assert.Equal(t, now.Add(eventuallyConsistentWindowSize), *endOfWindowTime)
	assert.Equal(t, 1.0, backoffScalingFactor)
}

func TestDeterminePolicyToUse_EcRetryPolicy_NeedEcRetries_UnlimitedAttempts(t *testing.T) {
	resetTimes()

	drp := NewRetryPolicyWithOptions(
		ReplaceWithValuesFromRetryPolicy(DefaultRetryPolicyWithoutEventualConsistency()),
		WithUnlimitedAttempts(time.Duration(2)*time.Minute),
	)
	rp := EventuallyConsistentRetryPolicy(drp)

	now := setupTimesNeedEcRetries((*time.Time)(nil), rp)

	var policyToUse, endOfWindowTime, backoffScalingFactor = rp.DeterminePolicyToUse(rp)
	assertEqualRetryPolicies(t, rp, policyToUse)
	assert.Equal(t, now.Add(eventuallyConsistentWindowSize), *endOfWindowTime)
	// since the initial request is exactly half-way through the maximum, expect factor of 0.5
	assert.Equal(t, 0.5, backoffScalingFactor)
}

func TestDeterminePolicyToUse_EcRetryPolicy_NeedFullEcRetries_UnlimitedAttempts(t *testing.T) {
	resetTimes()

	drp := NewRetryPolicyWithOptions(
		ReplaceWithValuesFromRetryPolicy(DefaultRetryPolicyWithoutEventualConsistency()),
		WithUnlimitedAttempts(time.Duration(2)*time.Minute),
	)
	rp := EventuallyConsistentRetryPolicy(drp)

	maxCumulativeBackoffWithoutJitter := GetMaximumCumulativeEventuallyConsistentBackoffWithoutJitter(rp)
	assert.Equal(t, maxCumulativeBackoffWithoutJitter, rp.GetMaximumCumulativeBackoffWithoutJitter())

	// EC window ends after default retries, with almost the full max cumulative backoff left
	timeNowProviderValues := make(chan time.Time, 300)
	now := time.Now().Round(0)
	// most recent EC
	timeNowProviderValues <- now
	// initial request
	timeNowProviderValues <- now
	EcContext.timeNowProvider = func() time.Time {
		nextTime := <-timeNowProviderValues
		return nextTime
	}
	EcContext.UpdateEndOfWindow(eventuallyConsistentWindowSize)

	var policyToUse, endOfWindowTime, backoffScalingFactor = rp.DeterminePolicyToUse(rp)
	assertEqualRetryPolicies(t, rp, policyToUse)
	assert.Equal(t, now.Add(eventuallyConsistentWindowSize), *endOfWindowTime)
	// since the initial request is 4 minutes / (4 minutes 0.524 seconds), the factor
	// should be really close to 1
	assert.Equal(t, float64(4*time.Minute)/float64(maxCumulativeBackoffWithoutJitter), backoffScalingFactor)
	assert.True(t, backoffScalingFactor > 0.99)
	assert.True(t, backoffScalingFactor < 1.0)
}

func TestDeterminePolicyToUse_EcRetryPolicy_EndOfWindowInThePast_UnlimitedAttempts(t *testing.T) {
	resetTimes()

	drp := NewRetryPolicyWithOptions(
		ReplaceWithValuesFromRetryPolicy(DefaultRetryPolicyWithoutEventualConsistency()),
		WithUnlimitedAttempts(time.Duration(2)*time.Minute),
	)
	rp := EventuallyConsistentRetryPolicy(drp)

	setupTimesEndOfEcWindowInThePast((*time.Time)(nil))

	var policyToUse, endOfWindowTime, backoffScalingFactor = rp.DeterminePolicyToUse(rp)
	assertEqualRetryPolicies(t, drp, policyToUse)
	assert.Equal(t, (*time.Time)(nil), endOfWindowTime)
	assert.Equal(t, 1.0, backoffScalingFactor)
}

// more tests

func assertEqualRetryPolicies(t *testing.T, this RetryPolicy, that RetryPolicy) {
	assert.Equal(t, this.MaximumNumberAttempts, that.MaximumNumberAttempts)
	assert.Equal(t, this.MinSleepBetween, that.MinSleepBetween)
	assert.Equal(t, this.MaxSleepBetween, that.MaxSleepBetween)
	assert.Equal(t, this.ExponentialBackoffBase, that.ExponentialBackoffBase)
	assert.Equal(t, this.MaximumCumulativeBackoffWithoutJitter, that.MaximumCumulativeBackoffWithoutJitter)
	assert.True(t, reflect.ValueOf(this.NextDuration).Pointer() == reflect.ValueOf(that.NextDuration).Pointer())
	assert.True(t, reflect.ValueOf(this.ShouldRetryOperation).Pointer() == reflect.ValueOf(that.ShouldRetryOperation).Pointer())
	if this.NonEventuallyConsistentPolicy != nil && that.NonEventuallyConsistentPolicy != nil {
		assertEqualRetryPolicies(t, *this.NonEventuallyConsistentPolicy, *that.NonEventuallyConsistentPolicy)
	} else {
		assert.True(t, reflect.ValueOf(this.NonEventuallyConsistentPolicy).Pointer() == reflect.ValueOf(that.NonEventuallyConsistentPolicy).Pointer())
	}
}

func TestEventuallyConsistentRetryPolicy_AlreadyEC(t *testing.T) {
	policy := DefaultRetryPolicy()
	policy2 := EventuallyConsistentRetryPolicy(policy)

	assertEqualRetryPolicies(t, policy, policy2)
}

func TestEventuallyConsistentRetryPolicy_NoEc(t *testing.T) {
	resetTimes()

	policy := DefaultRetryPolicy()
	assert.Equal(t, ecMaximumNumberAttempts, policy.MaximumNumberAttempts)
	assert.Equal(t, ecExponentialBackoffBase, policy.ExponentialBackoffBase)
	assert.Equal(t, ecMaxSleepBetween, policy.MaxSleepBetween)

	// unroll an exponential retry policy with a specified maximum
	// number of attempts so it's more obvious what's happening
	for i := uint(1); i <= 9; i++ {
		assert.True(t, shouldContinueIssuingRequests(i, policy.MaximumNumberAttempts))
	}
	assert.False(t, shouldContinueIssuingRequests(10, policy.MaximumNumberAttempts))

	for _, r := range responsesNoRetry {
		assert.False(t, policy.ShouldRetryOperation(r))
	}
	for _, r := range responsesWantRetry {
		assert.True(t, policy.ShouldRetryOperation(r))
	}
}

func TestEventuallyConsistentRetryPolicy_EndOfEcWindowInThePast(t *testing.T) {
	resetTimes()

	policy := DefaultRetryPolicy()
	assert.Equal(t, ecMaximumNumberAttempts, policy.MaximumNumberAttempts)
	assert.Equal(t, ecExponentialBackoffBase, policy.ExponentialBackoffBase)
	assert.Equal(t, ecMaxSleepBetween, policy.MaxSleepBetween)

	// unroll an exponential retry policy with a specified maximum
	// number of attempts so it's more obvious what's happening
	for i := uint(1); i <= 9; i++ {
		assert.True(t, shouldContinueIssuingRequests(i, policy.MaximumNumberAttempts))
	}
	assert.False(t, shouldContinueIssuingRequests(10, policy.MaximumNumberAttempts))

	now := time.Now().Round(0)
	nowPlusOne := now.Add(1 * time.Minute)
	endOfWindowTime := now.Add(5 * time.Minute)

	for _, r := range buildResponsesNoRetry(&endOfWindowTime, 1.0) {
		setupTimesEndOfEcWindowInThePast(&nowPlusOne)
		assert.False(t, policy.ShouldRetryOperation(r))
	}
	for _, r := range buildResponsesWantRetry(&endOfWindowTime, 1.0) {
		setupTimesEndOfEcWindowInThePast(&nowPlusOne)
		assert.True(t, policy.ShouldRetryOperation(r))
	}
}

func TestEventuallyConsistentRetryPolicy_EcEndsSoonerThanDefaultRetries(t *testing.T) {
	resetTimes()

	policy := DefaultRetryPolicy()
	assert.Equal(t, ecMaximumNumberAttempts, policy.MaximumNumberAttempts)
	assert.Equal(t, ecExponentialBackoffBase, policy.ExponentialBackoffBase)
	assert.Equal(t, ecMaxSleepBetween, policy.MaxSleepBetween)

	// unroll an exponential retry policy with a specified maximum
	// number of attempts so it's more obvious what's happening
	for i := uint(1); i <= 9; i++ {
		assert.True(t, shouldContinueIssuingRequests(i, policy.MaximumNumberAttempts))
	}
	assert.False(t, shouldContinueIssuingRequests(10, policy.MaximumNumberAttempts))

	now := time.Now().Round(0)
	endOfWindowTime := now.Add(5 * time.Minute)

	for _, r := range buildEcResponsesNoRetry(&endOfWindowTime, 1.0) {
		setupTimesEcEndsSoonerThanDefaultRetries(&now)
		assert.False(t, policy.ShouldRetryOperation(r))
	}
	for _, r := range buildEcResponsesWantRetry(&endOfWindowTime, 1.0) {
		setupTimesEcEndsSoonerThanDefaultRetries(&now)
		assert.True(t, policy.ShouldRetryOperation(r))
	}
}

func TestEventuallyConsistentRetryPolicy_NeedEcRetries(t *testing.T) {
	resetTimes()

	policy := DefaultRetryPolicy()
	assert.Equal(t, ecMaximumNumberAttempts, policy.MaximumNumberAttempts)
	assert.Equal(t, ecExponentialBackoffBase, policy.ExponentialBackoffBase)
	assert.Equal(t, ecMaxSleepBetween, policy.MaxSleepBetween)

	// unroll an exponential retry policy with a specified maximum
	// number of attempts so it's more obvious what's happening
	for i := uint(1); i <= 9; i++ {
		assert.True(t, shouldContinueIssuingRequests(i, policy.MaximumNumberAttempts))
	}
	assert.False(t, shouldContinueIssuingRequests(10, policy.MaximumNumberAttempts))

	now := time.Now().Round(0)
	endOfWindowTime := now.Add(5 * time.Minute)

	for _, r := range buildEcResponsesNoRetry(&endOfWindowTime, 1.0) {
		setupTimesNeedEcRetries(&now, policy)
		assert.False(t, policy.ShouldRetryOperation(r))
	}
	for _, r := range buildEcResponsesWantRetry(&endOfWindowTime, 1.0) {
		setupTimesNeedEcRetries(&now, policy)
		assert.True(t, policy.ShouldRetryOperation(r))
	}
}

func resetTimes() {
	EcContext.setEndOfWindow((*time.Time)(nil))
	EcContext.timeNowProvider = func() time.Time { return time.Now().Round(0) }
}

func setupTimesEndOfEcWindowInThePast(now *time.Time) time.Time {
	// EC window ends in the past
	timeNowProviderValues := make(chan time.Time, 300)

	if now == (*time.Time)(nil) {
		n := time.Now().Round(0)
		now = &n
	}

	// most recent EC
	timeNowProviderValues <- *now
	// initial request
	timeNowProviderValues <- now.Add(5 * time.Minute)
	EcContext.timeNowProvider = func() time.Time {
		nextTime := <-timeNowProviderValues
		return nextTime
	}
	EcContext.UpdateEndOfWindow(eventuallyConsistentWindowSize)

	return *now
}

func setupTimesEcEndsSoonerThanDefaultRetries(now *time.Time) time.Time {
	// EC window ends sooner than default retries
	timeNowProviderValues := make(chan time.Time, 300)

	if now == (*time.Time)(nil) {
		n := time.Now().Round(0)
		now = &n
	}

	// most recent EC
	timeNowProviderValues <- *now
	// initial request, set up to be 90 seconds before the end of the EC period
	// the default backoff is 91 seconds long, so it's longer
	timeNowProviderValues <- now.Add((240 - 90) * time.Second)
	// checking after first attempt
	timeNowProviderValues <- now.Add((240 - 89) * time.Second)
	EcContext.timeNowProvider = func() time.Time {
		nextTime := <-timeNowProviderValues
		return nextTime
	}
	EcContext.UpdateEndOfWindow(eventuallyConsistentWindowSize)

	return *now
}

func setupTimesNeedEcRetries(now *time.Time, policy RetryPolicy) time.Time {
	// EC window ends after default retries, with half the maximum cumulative backoff left
	timeNowProviderValues := make(chan time.Time, 300)

	if now == (*time.Time)(nil) {
		n := time.Now().Round(0)
		now = &n
	}

	maxCumulativeBackoffWithoutJitter := GetMaximumCumulativeEventuallyConsistentBackoffWithoutJitter(policy)

	// most recent EC
	timeNowProviderValues <- *now
	// initial request, set up to be exactly half-way through the max cumulative backoff
	timeNowProviderValues <- now.Add(240*time.Second - maxCumulativeBackoffWithoutJitter/2)
	EcContext.timeNowProvider = func() time.Time {
		nextTime := <-timeNowProviderValues
		return nextTime
	}
	EcContext.UpdateEndOfWindow(eventuallyConsistentWindowSize)

	return *now
}

// EC

func TestEcRetry_NoEc(t *testing.T) {
	resetTimes()

	body := bytes.NewBufferString("YES")
	req, _ := http.NewRequest("POST", "/some", body)
	policy := DefaultRetryPolicy()
	r := mockedRequest{Request: *req, Policy: &policy}

	var timesCalled = 0

	operation := func(i context.Context, request OCIRequest, binaryRequestBody *OCIReadSeekCloser, extraHeaders map[string]string) (OCIResponse, error) {
		timesCalled++

		request.HTTPRequest("POST", "/some", NewOCIReadSeekCloser(nil), nil)

		httpResponse := http.Response{
			Header:     http.Header{},
			StatusCode: 404,
		}
		response := mockedResponse{
			RawResponse: &httpResponse,
		}
		err := servicefailure{
			StatusCode: 404,
			Code:       "NotAuthorizedOrNotFound",
		}

		return response, err
	}

	Retry(context.Background(), r, operation, *r.Policy)

	assert.Equal(t, 1, timesCalled)
}

func TestEcRetry_EndOfEcWindowInThePast(t *testing.T) {
	resetTimes()
	setupTimesEndOfEcWindowInThePast((*time.Time)(nil))

	body := bytes.NewBufferString("YES")
	req, _ := http.NewRequest("POST", "/some", body)
	policy := DefaultRetryPolicy()
	r := mockedRequest{Request: *req, Policy: &policy}

	var timesCalled = 0

	operation := func(i context.Context, request OCIRequest, binaryRequestBody *OCIReadSeekCloser, extraHeaders map[string]string) (OCIResponse, error) {
		timesCalled++

		request.HTTPRequest("POST", "/some", NewOCIReadSeekCloser(nil), nil)

		httpResponse := http.Response{
			Header:     http.Header{},
			StatusCode: 404,
		}
		response := mockedResponse{
			RawResponse: &httpResponse,
		}
		err := servicefailure{
			StatusCode: 404,
			Code:       "NotAuthorizedOrNotFound",
		}

		return response, err
	}

	Retry(context.Background(), r, operation, *r.Policy)

	assert.Equal(t, 1, timesCalled)
}

func TestDefaultRetryWithoutEventualConsistency_NoRetryOn404(t *testing.T) {
	body := bytes.NewBufferString("YES")
	req, _ := http.NewRequest("POST", "/some", body)
	policy := DefaultRetryPolicyWithoutEventualConsistency()
	r := mockedRequest{Request: *req, Policy: &policy}

	var timesCalled = 0

	operation := func(i context.Context, request OCIRequest, binaryRequestBody *OCIReadSeekCloser, extraHeaders map[string]string) (OCIResponse, error) {
		timesCalled++

		request.HTTPRequest("POST", "/some", NewOCIReadSeekCloser(nil), nil)

		if timesCalled == 1 {
			httpResponse := http.Response{
				Header:     http.Header{},
				StatusCode: 404,
			}
			response := mockedResponse{
				RawResponse: &httpResponse,
			}
			err := servicefailure{
				StatusCode: 404,
				Code:       "NotAuthorizedOrNotFound",
			}
			return response, err
		}
		httpResponse := http.Response{
			Header:     http.Header{},
			StatusCode: 200,
		}
		return mockedResponse{RawResponse: &httpResponse}, nil
	}

	Retry(context.Background(), r, operation, *r.Policy)

	assert.Equal(t, 1, timesCalled)
}

func TestEcRetry_EcEndsSoonerThanDefaultRetries(t *testing.T) {
	resetTimes()
	setupTimesEcEndsSoonerThanDefaultRetries((*time.Time)(nil))

	body := bytes.NewBufferString("YES")
	req, _ := http.NewRequest("POST", "/some", body)
	policy := DefaultRetryPolicy()
	r := mockedRequest{Request: *req, Policy: &policy}

	var timesCalled = 0

	operation := func(i context.Context, request OCIRequest, binaryRequestBody *OCIReadSeekCloser, extraHeaders map[string]string) (OCIResponse, error) {
		timesCalled++

		request.HTTPRequest("POST", "/some", NewOCIReadSeekCloser(nil), nil)

		if timesCalled == 1 {
			httpResponse := http.Response{
				Header:     http.Header{},
				StatusCode: 404,
			}
			response := mockedResponse{
				RawResponse: &httpResponse,
			}
			err := servicefailure{
				StatusCode: 404,
				Code:       "NotAuthorizedOrNotFound",
			}
			return response, err
		}
		httpResponse := http.Response{
			Header:     http.Header{},
			StatusCode: 200,
		}
		return mockedResponse{RawResponse: &httpResponse}, nil
	}

	Retry(context.Background(), r, operation, *r.Policy)

	assert.Equal(t, 2, timesCalled)
}

// EC with unlimited attempts

func TestEcRetry_NoEc_UnlimitedAttempts(t *testing.T) {
	resetTimes()

	body := bytes.NewBufferString("YES")
	req, _ := http.NewRequest("POST", "/some", body)
	policy := NewRetryPolicyWithOptions(
		WithUnlimitedAttempts(time.Duration(10)*time.Minute),
		WithNextDuration(func(r OCIOperationResponse) time.Duration {
			return time.Duration(math.Pow(float64(2), float64(r.AttemptNumber-1))) * time.Second
		}),
	)
	r := mockedRequest{Request: *req, Policy: &policy}

	var timesCalled = 0

	operation := func(i context.Context, request OCIRequest, binaryRequestBody *OCIReadSeekCloser, extraHeaders map[string]string) (OCIResponse, error) {
		timesCalled++

		request.HTTPRequest("POST", "/some", NewOCIReadSeekCloser(nil), nil)

		httpResponse := http.Response{
			Header:     http.Header{},
			StatusCode: 404,
		}
		response := mockedResponse{
			RawResponse: &httpResponse,
		}
		err := servicefailure{
			StatusCode: 404,
			Code:       "NotAuthorizedOrNotFound",
		}

		return response, err
	}

	Retry(context.Background(), r, operation, *r.Policy)

	assert.Equal(t, 1, timesCalled)
}

func TestEcRetry_EndOfEcWindowInThePast_UnlimitedAttempts(t *testing.T) {
	resetTimes()
	setupTimesEndOfEcWindowInThePast((*time.Time)(nil))

	body := bytes.NewBufferString("YES")
	req, _ := http.NewRequest("POST", "/some", body)
	policy := NewRetryPolicyWithOptions(
		WithUnlimitedAttempts(time.Duration(10)*time.Minute),
		WithNextDuration(func(r OCIOperationResponse) time.Duration {
			return time.Duration(math.Pow(float64(2), float64(r.AttemptNumber-1))) * time.Second
		}),
	)
	r := mockedRequest{Request: *req, Policy: &policy}

	var timesCalled = 0

	operation := func(i context.Context, request OCIRequest, binaryRequestBody *OCIReadSeekCloser, extraHeaders map[string]string) (OCIResponse, error) {
		timesCalled++

		request.HTTPRequest("POST", "/some", NewOCIReadSeekCloser(nil), nil)

		httpResponse := http.Response{
			Header:     http.Header{},
			StatusCode: 404,
		}
		response := mockedResponse{
			RawResponse: &httpResponse,
		}
		err := servicefailure{
			StatusCode: 404,
			Code:       "NotAuthorizedOrNotFound",
		}

		return response, err
	}

	Retry(context.Background(), r, operation, *r.Policy)

	assert.Equal(t, 1, timesCalled)
}

func TestDefaultRetryWithoutEventualConsistency_NoRetryOn404_UnlimitedAttempts(t *testing.T) {
	body := bytes.NewBufferString("YES")
	req, _ := http.NewRequest("POST", "/some", body)
	policy := NewRetryPolicyWithOptions(
		ReplaceWithValuesFromRetryPolicy(DefaultRetryPolicyWithoutEventualConsistency()),
		WithUnlimitedAttempts(time.Duration(10)*time.Minute),
		WithNextDuration(func(r OCIOperationResponse) time.Duration {
			return time.Duration(math.Pow(float64(2), float64(r.AttemptNumber-1))) * time.Second
		}),
	)

	r := mockedRequest{Request: *req, Policy: &policy}

	var timesCalled = 0

	operation := func(i context.Context, request OCIRequest, binaryRequestBody *OCIReadSeekCloser, extraHeaders map[string]string) (OCIResponse, error) {
		timesCalled++

		request.HTTPRequest("POST", "/some", NewOCIReadSeekCloser(nil), nil)

		if timesCalled == 1 {
			httpResponse := http.Response{
				Header:     http.Header{},
				StatusCode: 404,
			}
			response := mockedResponse{
				RawResponse: &httpResponse,
			}
			err := servicefailure{
				StatusCode: 404,
				Code:       "NotAuthorizedOrNotFound",
			}
			return response, err
		}
		httpResponse := http.Response{
			Header:     http.Header{},
			StatusCode: 200,
		}
		return mockedResponse{RawResponse: &httpResponse}, nil
	}

	Retry(context.Background(), r, operation, *r.Policy)

	assert.Equal(t, 1, timesCalled)
}

func TestEcRetry_EcEndsSoonerThanDefaultRetries_UnlimitedAttempts(t *testing.T) {
	resetTimes()
	setupTimesEcEndsSoonerThanDefaultRetries((*time.Time)(nil))

	body := bytes.NewBufferString("YES")
	req, _ := http.NewRequest("POST", "/some", body)
	policy := NewRetryPolicyWithOptions(
		WithUnlimitedAttempts(time.Duration(2)*time.Minute),
		WithNextDuration(func(r OCIOperationResponse) time.Duration {
			return time.Duration(math.Pow(float64(2), float64(r.AttemptNumber-1))) * time.Second
		}),
	)
	r := mockedRequest{Request: *req, Policy: &policy}

	var timesCalled = 0

	operation := func(i context.Context, request OCIRequest, binaryRequestBody *OCIReadSeekCloser, extraHeaders map[string]string) (OCIResponse, error) {
		timesCalled++

		request.HTTPRequest("POST", "/some", NewOCIReadSeekCloser(nil), nil)

		if timesCalled == 1 {
			httpResponse := http.Response{
				Header:     http.Header{},
				StatusCode: 404,
			}
			response := mockedResponse{
				RawResponse: &httpResponse,
			}
			err := servicefailure{
				StatusCode: 404,
				Code:       "NotAuthorizedOrNotFound",
			}
			return response, err
		}
		httpResponse := http.Response{
			Header:     http.Header{},
			StatusCode: 200,
		}
		return mockedResponse{RawResponse: &httpResponse}, nil
	}

	Retry(context.Background(), r, operation, *r.Policy)

	assert.Equal(t, 2, timesCalled)
}

// more tests

func TestRetryPolicyBuilder(t *testing.T) {
	var rp = DefaultRetryPolicyWithoutEventualConsistency()
	assert.Equal(t, defaultMaximumNumberAttempts, rp.MaximumNumberAttempts)
	assert.Equal(t, defaultMinSleepBetween, rp.MinSleepBetween)
	assert.Equal(t, defaultMaxSleepBetween, rp.MaxSleepBetween)
	assert.Equal(t, defaultExponentialBackoffBase, rp.ExponentialBackoffBase)
	assert.Equal(t, (*RetryPolicy)(nil), rp.NonEventuallyConsistentPolicy)

	rp = NewRetryPolicyWithOptions(
		ReplaceWithValuesFromRetryPolicy(rp),
		WithConditionalOption(true, WithMaximumNumberAttempts(4)))
	assert.Equal(t, uint(4), rp.MaximumNumberAttempts)
	assert.Equal(t, defaultMinSleepBetween, rp.MinSleepBetween)
	assert.Equal(t, defaultMaxSleepBetween, rp.MaxSleepBetween)
	assert.Equal(t, defaultExponentialBackoffBase, rp.ExponentialBackoffBase)
	assert.Equal(t, (*RetryPolicy)(nil), rp.NonEventuallyConsistentPolicy)

	rp = NewRetryPolicyWithOptions(
		ReplaceWithValuesFromRetryPolicy(rp),
		WithConditionalOption(false, WithMaximumNumberAttempts(5)))
	assert.Equal(t, uint(4), rp.MaximumNumberAttempts)
	assert.Equal(t, defaultMinSleepBetween, rp.MinSleepBetween)
	assert.Equal(t, defaultMaxSleepBetween, rp.MaxSleepBetween)
	assert.Equal(t, defaultExponentialBackoffBase, rp.ExponentialBackoffBase)
	assert.Equal(t, (*RetryPolicy)(nil), rp.NonEventuallyConsistentPolicy)

	rp = NewRetryPolicyWithOptions(
		ReplaceWithValuesFromRetryPolicy(rp),
		WithMaximumNumberAttempts(5))
	assert.Equal(t, uint(5), rp.MaximumNumberAttempts)
	assert.Equal(t, defaultMinSleepBetween, rp.MinSleepBetween)
	assert.Equal(t, defaultMaxSleepBetween, rp.MaxSleepBetween)
	assert.Equal(t, defaultExponentialBackoffBase, rp.ExponentialBackoffBase)
	assert.Equal(t, (*RetryPolicy)(nil), rp.NonEventuallyConsistentPolicy)

	assertNextDurationIsBetween(t, rp, 1, 1*time.Second, 2*time.Second)
	assertNextDurationIsBetween(t, rp, 2, 2*time.Second, 3*time.Second)
	assertNextDurationIsBetween(t, rp, 3, 4*time.Second, 5*time.Second)
	assertNextDurationIsBetween(t, rp, 4, 8*time.Second, 9*time.Second)
	assertNextDurationIsBetween(t, rp, 5, 16*time.Second, 17*time.Second)
	assertNextDurationIsBetween(t, rp, 6, 30*time.Second, 31*time.Second) // hit the max
	assertNextDurationIsBetween(t, rp, 7, 30*time.Second, 31*time.Second)

	rp = NewRetryPolicyWithOptions(
		ReplaceWithValuesFromRetryPolicy(rp),
		WithExponentialBackoff(45*time.Second, rp.ExponentialBackoffBase))
	assert.Equal(t, uint(5), rp.MaximumNumberAttempts)
	assert.Equal(t, float64(45), rp.MaxSleepBetween)
	assert.Equal(t, defaultExponentialBackoffBase, rp.ExponentialBackoffBase)
	assert.Equal(t, (*RetryPolicy)(nil), rp.NonEventuallyConsistentPolicy)

	assertNextDurationIsBetween(t, rp, 1, 1*time.Second, 2*time.Second)
	assertNextDurationIsBetween(t, rp, 2, 2*time.Second, 3*time.Second)
	assertNextDurationIsBetween(t, rp, 3, 4*time.Second, 5*time.Second)
	assertNextDurationIsBetween(t, rp, 4, 8*time.Second, 9*time.Second)
	assertNextDurationIsBetween(t, rp, 5, 16*time.Second, 17*time.Second)
	assertNextDurationIsBetween(t, rp, 6, 32*time.Second, 33*time.Second)
	assertNextDurationIsBetween(t, rp, 7, 45*time.Second, 46*time.Second) // hit the max

	rp = NewRetryPolicyWithOptions(
		ReplaceWithValuesFromRetryPolicy(rp),
		WithExponentialBackoff(45*time.Second, 3))
	assert.Equal(t, uint(5), rp.MaximumNumberAttempts)
	assert.Equal(t, defaultMinSleepBetween, rp.MinSleepBetween)
	assert.Equal(t, float64(45), rp.MaxSleepBetween)
	assert.Equal(t, 3.0, rp.ExponentialBackoffBase)
	assert.Equal(t, (*RetryPolicy)(nil), rp.NonEventuallyConsistentPolicy)

	assertNextDurationIsBetween(t, rp, 1, 1*time.Second, 2*time.Second)
	assertNextDurationIsBetween(t, rp, 2, 3*time.Second, 4*time.Second)
	assertNextDurationIsBetween(t, rp, 3, 9*time.Second, 10*time.Second)
	assertNextDurationIsBetween(t, rp, 4, 27*time.Second, 28*time.Second)
	assertNextDurationIsBetween(t, rp, 5, 45*time.Second, 46*time.Second) // hit the max
	assertNextDurationIsBetween(t, rp, 6, 45*time.Second, 46*time.Second)
	assertNextDurationIsBetween(t, rp, 7, 45*time.Second, 46*time.Second)

	assert.Equal(t, (1+3+9+27)*time.Second, GetMaximumCumulativeBackoffWithoutJitter(rp))
	assert.Equal(t, (1+3+9+27)*time.Second, rp.GetMaximumCumulativeBackoffWithoutJitter())

	rp = NewRetryPolicyWithOptions(
		ReplaceWithValuesFromRetryPolicy(rp),
		WithFixedBackoff(5*time.Second))
	assert.Equal(t, uint(5), rp.MaximumNumberAttempts)
	assert.Equal(t, float64(5), rp.MinSleepBetween)
	assert.Equal(t, float64(5), rp.MaxSleepBetween)
	assert.Equal(t, 1.0, rp.ExponentialBackoffBase)
	assert.Equal(t, (*RetryPolicy)(nil), rp.NonEventuallyConsistentPolicy)

	assertNextDurationIsEqual(t, rp, 1, 5*time.Second)
	assertNextDurationIsEqual(t, rp, 2, 5*time.Second)
	assertNextDurationIsEqual(t, rp, 3, 5*time.Second)
	assertNextDurationIsEqual(t, rp, 4, 5*time.Second)
	assertNextDurationIsEqual(t, rp, 5, 5*time.Second)
	assertNextDurationIsEqual(t, rp, 6, 5*time.Second)
	assertNextDurationIsEqual(t, rp, 7, 5*time.Second)

	assert.Equal(t, (5+5+5+5)*time.Second, GetMaximumCumulativeBackoffWithoutJitter(rp))
	assert.Equal(t, (5+5+5+5)*time.Second, rp.GetMaximumCumulativeBackoffWithoutJitter())

	rp = NewRetryPolicyWithOptions(
		ReplaceWithValuesFromRetryPolicy(rp),
		WithExponentialBackoff(45*time.Second, defaultExponentialBackoffBase))

	var ecrp = NewRetryPolicyWithOptions(
		ReplaceWithValuesFromRetryPolicy(rp),
		WithEventualConsistency())
	assert.Equal(t, ecMaximumNumberAttempts, ecrp.MaximumNumberAttempts)
	assert.Equal(t, ecMinSleepBetween, ecrp.MinSleepBetween)
	assert.Equal(t, ecMaxSleepBetween, ecrp.MaxSleepBetween)
	assert.Equal(t, ecExponentialBackoffBase, ecrp.ExponentialBackoffBase)
	assertEqualRetryPolicies(t, rp, *ecrp.NonEventuallyConsistentPolicy)

	rp = NewRetryPolicyWithOptions()
	assertEqualRetryPolicies(t, DefaultRetryPolicy(), rp)
}

func assertNextDurationIsBetween(t *testing.T, rp RetryPolicy, attempt uint, min time.Duration, max time.Duration) {
	response := getMockedOCIOperationResponse(500, attempt)
	nextDuration := rp.NextDuration(response)
	assert.True(t, min <= nextDuration && nextDuration < max)
}

func assertNextDurationIsEqual(t *testing.T, rp RetryPolicy, attempt uint, duration time.Duration) {
	response := getMockedOCIOperationResponse(500, attempt)
	nextDuration := rp.NextDuration(response)
	assert.Equal(t, duration, nextDuration)
}

func TestRetryPolicy_DeterminePolicyToUseNil(t *testing.T) {
	var shouldRetry = true
	retryPolicy := &RetryPolicy{
		MaximumNumberAttempts: 0,
		ShouldRetryOperation: func(response OCIOperationResponse) bool {
			prev := shouldRetry
			shouldRetry = false
			return prev
		},
		NextDuration: func(response OCIOperationResponse) time.Duration {
			return 1 * time.Second
		},
	}

	resetTimes()

	body := bytes.NewBufferString("YES")
	req, _ := http.NewRequest("POST", "/some", body)
	r := mockedRequest{Request: *req, Policy: retryPolicy}

	var timesCalled = 0

	operation := func(i context.Context, request OCIRequest, binaryRequestBody *OCIReadSeekCloser, extraHeaders map[string]string) (OCIResponse, error) {
		timesCalled++

		request.HTTPRequest("POST", "/some", NewOCIReadSeekCloser(nil), nil)

		httpResponse := http.Response{
			Header:     http.Header{},
			StatusCode: 404,
		}
		response := mockedResponse{
			RawResponse: &httpResponse,
		}
		err := servicefailure{
			StatusCode: 404,
			Code:       "NotAuthorizedOrNotFound",
		}

		return response, err
	}

	Retry(context.Background(), r, operation, *r.Policy)

	assert.Equal(t, 2, timesCalled)
}

func TestRetryPolicy_Validate(t *testing.T) {
	var rp = DefaultRetryPolicyWithoutEventualConsistency()
	var result, err = rp.validate()
	assert.True(t, result)
	assert.Equal(t, nil, err)

	rp = NoRetryPolicy()
	result, err = rp.validate()
	assert.True(t, result)
	assert.Equal(t, nil, err)

	rp = DefaultRetryPolicy()
	result, err = rp.validate()
	assert.True(t, result)
	assert.Equal(t, nil, err)

	rp = RetryPolicy{
		MaximumNumberAttempts: 0,
		ShouldRetryOperation: func(response OCIOperationResponse) bool {
			return false
		},
		NextDuration: func(response OCIOperationResponse) time.Duration {
			return 0 * time.Second
		},
	}
	result, err = rp.validate()
	assert.True(t, result)
	assert.Equal(t, nil, err)

	rp = RetryPolicy{
		MaximumNumberAttempts: 0,
		NextDuration: func(response OCIOperationResponse) time.Duration {
			return 0 * time.Second
		},
	}
	result, err = rp.validate()
	assert.False(t, result)
	assert.True(t, strings.Contains(err.Error(), "ShouldRetryOperation"))

	rp = RetryPolicy{
		MaximumNumberAttempts: 0,
		ShouldRetryOperation: func(response OCIOperationResponse) bool {
			return false
		},
	}
	result, err = rp.validate()
	assert.False(t, result)
	assert.True(t, strings.Contains(err.Error(), "NextDuration"))

	rp = RetryPolicy{
		MaximumNumberAttempts: 0,
	}
	result, err = rp.validate()
	assert.False(t, result)
	assert.True(t, strings.Contains(err.Error(), "ShouldRetryOperation"))
	assert.True(t, strings.Contains(err.Error(), "NextDuration"))
}

func TestRetryPolicy_RetryWithBadRetryPolicy(t *testing.T) {
	retryPolicy := &RetryPolicy{
		MaximumNumberAttempts: 0,
		NextDuration: func(response OCIOperationResponse) time.Duration {
			return 1 * time.Second
		},
	}

	resetTimes()

	body := bytes.NewBufferString("YES")
	req, _ := http.NewRequest("POST", "/some", body)
	r := mockedRequest{Request: *req, Policy: retryPolicy}

	var timesCalled = 0

	operation := func(i context.Context, request OCIRequest, binaryRequestBody *OCIReadSeekCloser, extraHeaders map[string]string) (OCIResponse, error) {
		timesCalled++

		request.HTTPRequest("POST", "/some", NewOCIReadSeekCloser(nil), nil)

		httpResponse := http.Response{
			Header:     http.Header{},
			StatusCode: 404,
		}
		response := mockedResponse{
			RawResponse: &httpResponse,
		}
		err := servicefailure{
			StatusCode: 404,
			Code:       "NotAuthorizedOrNotFound",
		}

		return response, err
	}

	_, error := Retry(context.Background(), r, operation, *r.Policy)

	assert.True(t, strings.Contains(error.Error(), "ShouldRetryOperation"))
}

func helperTestGetEndTimeOfEventuallyConsistentWindowMultiThreadedSet(t *testing.T, wg2 *sync.WaitGroup) {
	EcContext.UpdateEndOfWindow(eventuallyConsistentWindowSize)
	var eowt = EcContext.GetEndOfWindow()
	assert.NotEqual(t, (*time.Time)(nil), eowt)

	wg2.Done()

	value1 := EcContext.GetEndOfWindow()
	time.Sleep(1 * time.Second)
	assert.True(t, value1.Before(time.Now().Round(0).Add(eventuallyConsistentWindowSize)))
	value2 := EcContext.GetEndOfWindow()
	assert.Equal(t, value1, value2)
	eowt = EcContext.UpdateEndOfWindow(eventuallyConsistentWindowSize)
	assert.NotEqual(t, (*time.Time)(nil), eowt)

	value3 := EcContext.GetEndOfWindow()
	time.Sleep(1 * time.Second)
	assert.True(t, value3.Before(time.Now().Round(0).Add(eventuallyConsistentWindowSize)))
	assert.True(t, value1.Before(*value3))
	assert.True(t, value2.Before(*value3))
}

func helperTestGetEndTimeOfEventuallyConsistentWindowMultiThreadedGet(t *testing.T, wg2 *sync.WaitGroup) {
	wg2.Wait()

	eowt := EcContext.GetEndOfWindow()
	assert.NotEqual(t, (*time.Time)(nil), eowt)
}

func TestGetEndTimeOfEventuallyConsistentWindow_MultiThreaded(t *testing.T) {
	var eowt = EcContext.GetEndOfWindow()
	assert.Equal(t, (*time.Time)(nil), eowt)

	var wg sync.WaitGroup
	wg.Add(2)

	var wg2 sync.WaitGroup
	wg2.Add(1)

	go func() {
		defer wg.Done()
		helperTestGetEndTimeOfEventuallyConsistentWindowMultiThreadedSet(t, &wg2)
	}()

	go func() {
		defer wg.Done()
		helperTestGetEndTimeOfEventuallyConsistentWindowMultiThreadedGet(t, &wg2)
	}()

	wg.Wait()

	eowt = EcContext.GetEndOfWindow()
	assert.NotEqual(t, (*time.Time)(nil), eowt)
}
