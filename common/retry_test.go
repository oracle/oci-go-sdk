package common

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
	"time"
)

func getMockedResponseWithStatusCode(code int) *http.Response {
	response := http.Response{
		Header:     http.Header{},
		StatusCode: code,
	}
	return &response
}

func TestBackoff_staticDefaultsWhenValuesSetTooHigh(t *testing.T) {
	duration := MAX_STATIC_RETRY_INTERVAL + 10
	b := StaticBackoffPolicy(duration)
	assert.Equal(t, MAX_STATIC_RETRY_INTERVAL, b.Interval)
	assert.Equal(t, DEFAULT_STATIC_RETRY_FACTOR, b.IntervalFactor)
}

func TestBackoff_staticDefaultsWhenValuesSetTooLow(t *testing.T) {
	duration := MIN_STATIC_RETRY_INTERVAL - 10
	b := StaticBackoffPolicy(duration)
	assert.Equal(t, MIN_STATIC_RETRY_INTERVAL, b.Interval)
	assert.Equal(t, DEFAULT_STATIC_RETRY_FACTOR, b.IntervalFactor)
}

func TestBackoff_exponentialDefaultsWhenValuesSetTooLow(t *testing.T) {
	duration := MIN_EXPONENTIAL_RETRY_INTERVAL - 10
	factor := MIN_EXPONENTIAL_RETRY_FACTOR - 10
	b := ExponentialBackoffPolicy(duration, factor)
	assert.Equal(t, MIN_EXPONENTIAL_RETRY_INTERVAL, b.Interval)
	assert.Equal(t, MIN_EXPONENTIAL_RETRY_FACTOR, b.IntervalFactor)
}

func TestBackoff_exponentialDefaultsWhenValuesSetTooHigh(t *testing.T) {
	duration := MAX_EXPONENTIAL_RETRY_INTERVAL + 10
	factor := MAX_EXPONENTIAL_RETRY_FACTOR + 10
	b := ExponentialBackoffPolicy(duration, factor)
	assert.Equal(t, MAX_EXPONENTIAL_RETRY_INTERVAL, b.Interval)
	assert.Equal(t, MAX_EXPONENTIAL_RETRY_FACTOR, b.IntervalFactor)
}

func TestRetry_noRetryStatusCode100(t *testing.T) {
	response := getMockedResponseWithStatusCode(100)
	r := NoRetryPolicy()
	assert.False(t, r.ShouldRetryRequest(*response))
}

func TestRetry_noRetryStatusCode200(t *testing.T) {
	response := getMockedResponseWithStatusCode(200)
	r := NoRetryPolicy()
	assert.False(t, r.ShouldRetryRequest(*response))
}

func TestRetry_noRetryStatusCode300(t *testing.T) {
	response := getMockedResponseWithStatusCode(300)
	r := NoRetryPolicy()
	assert.False(t, r.ShouldRetryRequest(*response))
}

func TestRetry_standardRetryStatusCode100(t *testing.T) {
	response := getMockedResponseWithStatusCode(100)
	r := StandardRetryPolicy()
	for i := 0; i < DEFAULT_NUM_RETRIES; i++ {
		assert.True(t, r.ShouldRetryRequest(*response))
		d := r.GetDurationAndIncrementRetryCount()
		assert.Equal(t, DEFAULT_STATIC_RETRY_INTERVAL, d)
	}
	assert.False(t, r.ShouldRetryRequest(*response))
}

func TestRetry_standardRetryStatusCode200(t *testing.T) {
	response := getMockedResponseWithStatusCode(200)
	r := StandardRetryPolicy()
	assert.False(t, r.ShouldRetryRequest(*response))
}

func TestRetry_standardRetryStatusCode300(t *testing.T) {
	response := getMockedResponseWithStatusCode(300)
	r := StandardRetryPolicy()
	for i := 0; i < DEFAULT_NUM_RETRIES; i++ {
		assert.True(t, r.ShouldRetryRequest(*response))
		d := r.GetDurationAndIncrementRetryCount()
		assert.Equal(t, DEFAULT_STATIC_RETRY_INTERVAL, d)
	}
	assert.False(t, r.ShouldRetryRequest(*response))
}

func TestRetry_standardRetryStatusCode400(t *testing.T) {
	response := getMockedResponseWithStatusCode(400)
	r := StandardRetryPolicy()
	for i := 0; i < DEFAULT_NUM_RETRIES; i++ {
		assert.True(t, r.ShouldRetryRequest(*response))
		d := r.GetDurationAndIncrementRetryCount()
		assert.Equal(t, DEFAULT_STATIC_RETRY_INTERVAL, d)
	}
	assert.False(t, r.ShouldRetryRequest(*response))
}

func TestRetry_standardRetryStatusCode500(t *testing.T) {
	response := getMockedResponseWithStatusCode(500)
	r := StandardRetryPolicy()
	for i := 0; i < DEFAULT_NUM_RETRIES; i++ {
		assert.True(t, r.ShouldRetryRequest(*response))
		d := r.GetDurationAndIncrementRetryCount()
		assert.Equal(t, DEFAULT_STATIC_RETRY_INTERVAL, d)
	}
	assert.False(t, r.ShouldRetryRequest(*response))
}

func TestRetry_standardRetryStatusCode500Ignored(t *testing.T) {
	response := getMockedResponseWithStatusCode(500)
	r := StandardRetryPolicyWithIgnored([]int{500})
	assert.False(t, r.ShouldRetryRequest(*response))
}

func TestRetry_exponentialRetryStatusCode500(t *testing.T) {
	response := getMockedResponseWithStatusCode(500)
	r := ExponentialRetryPolicy()
	// unroll the exponential retry test so it's more obvious what's happening
	// request #1
	assert.True(t, r.ShouldRetryRequest(*response))
	d := r.GetDurationAndIncrementRetryCount()
	assert.Equal(t, 1*time.Second, d)
	// request #2
	assert.True(t, r.ShouldRetryRequest(*response))
	d = r.GetDurationAndIncrementRetryCount()
	assert.Equal(t, 2*time.Second, d)
	// request #3
	assert.True(t, r.ShouldRetryRequest(*response))
	d = r.GetDurationAndIncrementRetryCount()
	assert.Equal(t, 4*time.Second, d)
	// request #4
	assert.True(t, r.ShouldRetryRequest(*response))
	d = r.GetDurationAndIncrementRetryCount()
	assert.Equal(t, 8*time.Second, d)
	// request #5
	assert.True(t, r.ShouldRetryRequest(*response))
	d = r.GetDurationAndIncrementRetryCount()
	assert.Equal(t, 16*time.Second, d)
	// done
	assert.False(t, r.ShouldRetryRequest(*response))
}
