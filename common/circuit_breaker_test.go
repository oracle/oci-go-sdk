// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.

package common

import (
	"testing"
	"time"

	"github.com/sony/gobreaker"
	"github.com/stretchr/testify/assert"
)

func TestCircuitBreaker_ReadyToTrip(t *testing.T) {
	type testData struct {
		counts   gobreaker.Counts
		expected bool
	}
	testDataSet := []testData{
		{
			counts: gobreaker.Counts{
				Requests:             30,
				TotalSuccesses:       3,
				TotalFailures:        27,
				ConsecutiveSuccesses: 1,
				ConsecutiveFailures:  15,
			},
			expected: true,
		},
		{
			counts: gobreaker.Counts{
				Requests:             9,
				TotalSuccesses:       0,
				TotalFailures:        9,
				ConsecutiveSuccesses: 0,
				ConsecutiveFailures:  9,
			},
			expected: false,
		},
		{
			counts: gobreaker.Counts{
				Requests:             20,
				TotalSuccesses:       18,
				TotalFailures:        2,
				ConsecutiveSuccesses: 17,
				ConsecutiveFailures:  1,
			},
			expected: false,
		},
	}
	st := gobreaker.Settings{}
	st.ReadyToTrip = func(counts gobreaker.Counts) bool {
		failureRatio := float64(counts.TotalFailures) / float64(counts.Requests)
		return counts.Requests >= CircuitBreakerDefaultVolumeThreshold && failureRatio >= CircuitBreakerDefaultFailureRateThreshold
	}
	for _, testData := range testDataSet {
		assert.Equal(t, testData.expected, st.ReadyToTrip(testData.counts))
	}
}

func TestCircuitBreaker_IsSuccessful(t *testing.T) {
	type testData struct {
		err      servicefailure
		expected bool
	}

	testDataSet := []testData{
		{
			err: servicefailure{
				StatusCode: 409,
				Code:       "IncorrectState",
			},
			expected: false,
		},
		{
			err: servicefailure{
				StatusCode: 409,
				Code:       "otherErrorCode",
			},
			expected: true,
		},
		{
			err: servicefailure{
				StatusCode: 500,
				Code:       "whatever",
			},
			expected: false,
		},
	}
	st := gobreaker.Settings{}
	successStatErrCodeMap := map[StatErrCode]bool{
		{409, "IncorrectState"}: false,
	}
	successStatCodeMap := map[int]bool{
		429: false,
		500: false,
		502: false,
		503: false,
		504: false,
	}
	st.IsSuccessful = func(err error) bool {
		if serviceErr, ok := IsServiceError(err); ok {
			if isSuccessful, ok := successStatCodeMap[serviceErr.GetHTTPStatusCode()]; ok {
				return isSuccessful
			}
			if isSuccessful, ok := successStatErrCodeMap[StatErrCode{serviceErr.GetHTTPStatusCode(), serviceErr.GetCode()}]; ok {
				return isSuccessful
			}
		}
		return true
	}
	for _, testData := range testDataSet {
		assert.Equal(t, testData.expected, st.IsSuccessful(testData.err))
	}
}

func TestCircuitBreaker_CustomizeGoBreakerSetting(t *testing.T) {

	cbSetting := DefaultCircuitBreakerSetting()
	assert.True(t, cbSetting.isEnabled)
	st := gobreaker.Settings{}

	customizeGoBreakerSetting(&st, cbSetting)
	assert.Equal(t, st.Name, DefaultCircuitBreakerName)
	assert.Equal(t, st.Timeout, CircuitBreakerDefaultResetTimeout)
	assert.Equal(t, st.Interval, CircuitBreakerDefaultClosedWindow)
	assert.Equal(t, cbSetting.failureRateThreshold, CircuitBreakerDefaultFailureRateThreshold)
	counts := gobreaker.Counts{
		Requests:             12,
		TotalSuccesses:       2,
		TotalFailures:        10,
		ConsecutiveSuccesses: 1,
		ConsecutiveFailures:  6,
	}
	assert.True(t, st.ReadyToTrip(counts))

	type testData struct {
		err      servicefailure
		expected bool
	}
	testDataSet := []testData{
		{
			err: servicefailure{
				StatusCode: 400,
				Code:       "InvalidParameter"},
			expected: true,
		},
		{
			err: servicefailure{
				StatusCode: 409,
				Code:       "IncorrectState"},
			expected: false,
		},
		{
			err: servicefailure{
				StatusCode: 429,
				Code:       "TooManyRequests"},
			expected: false,
		},
		{
			err: servicefailure{
				StatusCode: 504,
				Code:       "whatever"},
			expected: false,
		},
	}

	for _, testData := range testDataSet {
		assert.Equal(t, testData.expected, st.IsSuccessful(testData.err))
	}
}

func TestAuthClient_CircuitBreaker_ReadyToTrip(t *testing.T) {
	cbSetting := DefaultAuthClientCircuitBreakerSetting()
	assert.True(t, cbSetting.isEnabled)
	st := gobreaker.Settings{}

	customizeGoBreakerSetting(&st, cbSetting)
	assert.Equal(t, st.Name, AuthClientCircuitBreakerName)
	assert.True(t, st.Timeout >= time.Duration(MinAuthClientCircuitBreakerResetTimeout)*time.Second && st.Timeout <= time.Duration(MaxAuthClientCircuitBreakerResetTimeout)*time.Second)
	assert.Equal(t, st.Interval, CircuitBreakerDefaultClosedWindow)
	assert.Equal(t, cbSetting.failureRateThreshold, AuthClientCircuitBreakerDefaultFailureThreshold)

	counts := gobreaker.Counts{
		Requests:             1,
		TotalSuccesses:       0,
		TotalFailures:        1,
		ConsecutiveSuccesses: 0,
		ConsecutiveFailures:  1,
	}
	assert.False(t, st.ReadyToTrip(counts))

	counts = gobreaker.Counts{
		Requests:             2,
		TotalSuccesses:       0,
		TotalFailures:        2,
		ConsecutiveSuccesses: 0,
		ConsecutiveFailures:  2,
	}
	assert.False(t, st.ReadyToTrip(counts))

	counts = gobreaker.Counts{
		Requests:             3,
		TotalSuccesses:       0,
		TotalFailures:        3,
		ConsecutiveSuccesses: 0,
		ConsecutiveFailures:  3,
	}
	assert.True(t, st.ReadyToTrip(counts))

	type testData struct {
		err      servicefailure
		expected bool
	}
	testDataSet := []testData{
		{
			err: servicefailure{
				StatusCode: 400,
				Code:       "InvalidParameter"},
			expected: true,
		},
		{
			err: servicefailure{
				StatusCode: 409,
				Code:       "IncorrectState"},
			expected: false,
		},
		{
			err: servicefailure{
				StatusCode: 429,
				Code:       "TooManyRequests"},
			expected: false,
		},
		{
			err: servicefailure{
				StatusCode: 504,
				Code:       "whatever"},
			expected: false,
		},
	}

	for _, testData := range testDataSet {
		assert.Equal(t, testData.expected, st.IsSuccessful(testData.err))
	}
}
