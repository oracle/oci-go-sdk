package common

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
	"time"
)

func getMockedResponseWithStatusCode(code int) (*http.Response, error) {
	response := http.Response{
		Header:     http.Header{},
		StatusCode: code,
	}
	return &response, nil
}

func TestRetryPolicyOptionsMaxNumAttemptsSetToUnlimited(t *testing.T) {
	value := UnlimitedNumAttemptsValue
	p := BuildRetryPolicy(MaximumNumberAttempts(value))
	assert.Equal(t, MaximumNumAttemptsMinimum, p.MaximumNumberAttempts)
}

func TestRetryPolicyOptionsMaxNumAttemptsDefaultsValueWhenNoneProvided(t *testing.T) {
	p := BuildRetryPolicy()
	assert.Equal(t, MaximumNumAttemptsDefault, p.MaximumNumberAttempts)
}

func TestRetryPolicyOptionsMaxNumAttemptsSetsValueCorrectly(t *testing.T) {
	value := uint(9)
	p := BuildRetryPolicy(MaximumNumberAttempts(value))
	assert.Equal(t, value, p.MaximumNumberAttempts)
}

func TestRetryPolicyOptionsCumulativeTimeoutValueSetTooLow(t *testing.T) {
	value := MaximumTimeoutMinimum - 10
	p := BuildRetryPolicy(MaximumTimeout(value))
	assert.Equal(t, MaximumTimeoutMinimum, p.MaximumTimeout)
}

func TestRetryPolicyOptionsCumulativeTimeoutDefaultsValueWhenNoneProvided(t *testing.T) {
	p := BuildRetryPolicy()
	assert.Equal(t, MaximumTimeoutDefault, p.MaximumTimeout)
}

func TestRetryPolicyOptionsCumulativeTimeoutSetsValueCorrectly(t *testing.T) {
	value := 22 * time.Second
	p := BuildRetryPolicy(MaximumTimeout(value))
	assert.Equal(t, value, p.MaximumTimeout)
}

func TestRetryPolicyDefaultStatusCode100(t *testing.T) {
	response, err := getMockedResponseWithStatusCode(100)
	r := BuildRetryPolicy()
	assert.False(t, r.ShouldRetryOperation(response, err, uint(1)))
}

func TestRetryPolicyDefaultStatusCode200(t *testing.T) {
	response, err := getMockedResponseWithStatusCode(200)
	r := BuildRetryPolicy()
	assert.False(t, r.ShouldRetryOperation(response, err, uint(1)))
}

func TestRetryPolicyDefaultStatusCode300(t *testing.T) {
	response, err := getMockedResponseWithStatusCode(300)
	r := BuildRetryPolicy()
	assert.False(t, r.ShouldRetryOperation(response, err, uint(1)))
}

func TestRetryPolicyDefaultStatusCode400(t *testing.T) {
	response, err := getMockedResponseWithStatusCode(400)
	r := BuildRetryPolicy()
	assert.False(t, r.ShouldRetryOperation(response, err, uint(1)))
}

func TestRetryPolicyDefaultStatusCode405(t *testing.T) {
	response, err := getMockedResponseWithStatusCode(405)
	r := BuildRetryPolicy()
	assert.True(t, r.ShouldRetryOperation(response, err, uint(1)))
}

func TestRetryPolicyDefaultStatusCode500(t *testing.T) {
	response, err := getMockedResponseWithStatusCode(500)
	r := BuildRetryPolicy()
	assert.True(t, r.ShouldRetryOperation(response, err, uint(1)))
}

func TestRetryPolicyDefaultGetNextDurationUnrolled(t *testing.T) {
	response, err := getMockedResponseWithStatusCode(500)
	r := BuildRetryPolicy(MaximumNumberAttempts(5))
	// unroll the default (exponential) retry policy with a specified maximum
	// number of attempts so it's more obvious what's happening
	// request #1
	assert.True(t, r.ShouldRetryOperation(response, err, uint(1)))
	d := r.GetNextDuration(1)
	assert.Equal(t, 1*time.Second, d)
	// request #2
	assert.True(t, r.ShouldRetryOperation(response, err, uint(2)))
	d = r.GetNextDuration(2)
	assert.Equal(t, 2*time.Second, d)
	// request #3
	assert.True(t, r.ShouldRetryOperation(response, err, uint(3)))
	d = r.GetNextDuration(3)
	assert.Equal(t, 4*time.Second, d)
	// request #4
	assert.True(t, r.ShouldRetryOperation(response, err, uint(4)))
	d = r.GetNextDuration(4)
	assert.Equal(t, 8*time.Second, d)
	// request #5
	assert.True(t, r.ShouldRetryOperation(response, err, uint(5)))
	d = r.GetNextDuration(5)
	assert.Equal(t, 16*time.Second, d)
	// done
}
