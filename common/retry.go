package common

import (
	"math"
	"net/http"
	"time"
)

const (
	MIN_STATIC_RETRY_INTERVAL     = 100 * time.Millisecond
	MAX_STATIC_RETRY_INTERVAL     = 60 * time.Second
	DEFAULT_STATIC_RETRY_INTERVAL = 5 * time.Second
	DEFAULT_STATIC_RETRY_FACTOR   = float64(1)

	MIN_EXPONENTIAL_RETRY_INTERVAL     = 100 * time.Millisecond
	MAX_EXPONENTIAL_RETRY_INTERVAL     = 10 * time.Second
	MIN_EXPONENTIAL_RETRY_FACTOR       = float64(1)
	MAX_EXPONENTIAL_RETRY_FACTOR       = float64(5)
	DEFAULT_EXPONENTIAL_RETRY_INTERVAL = 1 * time.Second
	DEFAULT_EXPONENTIAL_RETRY_FACTOR   = float64(2)

	ZERO_RETRIES        = 0
	DEFAULT_NUM_RETRIES = 5
)

type Retryer interface {
	// based on the current state (current retry count vs max number of retries),
	// as well as the return code in the response, returns true if we should
	// continue retrying the request, and false otherwise.
	// this method does not mutate internal state.
	ShouldRetryRequest(response http.Response) bool

	// retrieves the duration to pause before retrying the request.
	GetDurationAndIncrementRetryCount() time.Duration
}

type BackoffPolicy struct {
	// Initial interval to wait before retrying a request.
	// If the IntervalFactor is set to '1', then each retry will occur
	// after 'Interval' duration.
	Interval time.Duration

	// Exponential base for computing duration in subsequent retry requests.
	//  duration = <interval> * <base> ** <current retry attempt>
	IntervalFactor float64
}

type RetryPolicy struct {
	// backoff policy to be used when retrying requests
	Backoff BackoffPolicy

	// maximum number of times to retry a request
	Retries int

	// status codes that will not be retried
	StatusCodesToNotRetry []int

	// counter that holds the current request attempt
	current int
}

func (r *RetryPolicy) ShouldRetryRequest(response http.Response) bool {
	shouldRetryRequest := true
	if r.current >= r.Retries {
		shouldRetryRequest = false
	} else {
		if response.StatusCode >= 200 && response.StatusCode < 300 {
			shouldRetryRequest = false
		} else {
			for _, code := range r.StatusCodesToNotRetry {
				if code == response.StatusCode {
					shouldRetryRequest = false
				}
			}
		}
	}
	return shouldRetryRequest
}

func (r *RetryPolicy) GetDurationAndIncrementRetryCount() time.Duration {
	duration := r.Backoff.Interval * time.Duration(math.Pow(r.Backoff.IntervalFactor, float64(r.current)))
	r.current = r.current + 1
	return duration
}

func NoRetryPolicy() Retryer {
	return &RetryPolicy{
		Retries: ZERO_RETRIES,
	}
}

func StandardRetryPolicy() Retryer {
	return StandardRetryPolicyWithIgnored([]int{})
}

func StandardRetryPolicyWithIgnored(ignored []int) Retryer {
	return &RetryPolicy{
		Backoff:               StaticBackoffPolicy(DEFAULT_STATIC_RETRY_INTERVAL),
		Retries:               DEFAULT_NUM_RETRIES,
		StatusCodesToNotRetry: ignored,
	}
}

func ExponentialRetryPolicy() Retryer {
	return ExponentialRetryPolicyWithIgnored([]int{})
}

func ExponentialRetryPolicyWithIgnored(ignored []int) Retryer {
	return &RetryPolicy{
		Backoff: ExponentialBackoffPolicy(
			DEFAULT_EXPONENTIAL_RETRY_INTERVAL,
			DEFAULT_EXPONENTIAL_RETRY_FACTOR),
		Retries:               DEFAULT_NUM_RETRIES,
		StatusCodesToNotRetry: ignored,
	}
}

func StaticBackoffPolicy(interval time.Duration) BackoffPolicy {
	i := interval
	if interval < MIN_STATIC_RETRY_INTERVAL {
		i = MIN_STATIC_RETRY_INTERVAL
	}
	if interval > MAX_STATIC_RETRY_INTERVAL {
		i = MAX_STATIC_RETRY_INTERVAL
	}
	return BackoffPolicy{
		Interval:       time.Duration(i),
		IntervalFactor: DEFAULT_STATIC_RETRY_FACTOR,
	}
}

func ExponentialBackoffPolicy(interval time.Duration, factor float64) BackoffPolicy {
	i := interval
	if interval < MIN_EXPONENTIAL_RETRY_INTERVAL {
		i = MIN_EXPONENTIAL_RETRY_INTERVAL
	}
	if interval > MAX_EXPONENTIAL_RETRY_INTERVAL {
		i = MAX_EXPONENTIAL_RETRY_INTERVAL
	}

	f := factor
	if factor < MIN_EXPONENTIAL_RETRY_FACTOR {
		f = MIN_EXPONENTIAL_RETRY_FACTOR
	}
	if factor > MAX_EXPONENTIAL_RETRY_FACTOR {
		f = MAX_EXPONENTIAL_RETRY_FACTOR
	}
	return BackoffPolicy{
		Interval:       time.Duration(i),
		IntervalFactor: f,
	}
}
