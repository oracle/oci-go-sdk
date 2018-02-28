package common

import (
	"math"
	"time"
)

const (
	// DefaultExponentialBackoffBase is the default value for the base in the following equation:
	// duration = factor * (base ^ (attempts - 1))
	DefaultExponentialBackoffBase = 2

	// DefaultExponentialBackoffFactor is the default value for the factor in the following equation:
	// duration = factor * (base ^ (attempts - 1))
	DefaultExponentialBackoffFactor = 1 * time.Second

	// UnlimitedNumAttemptsValue is the value for indicating unlimited attempts for reaching success
	UnlimitedNumAttemptsValue = uint(0)

	// MaximumNumAttemptsMinimum is the minimum value that can be set for the MaximumNumAttempts policy option.
	// This option accepts zero, which indicates unlimited number of attempts.
	MaximumNumAttemptsMinimum = uint(0)

	// MaximumNumAttemptsDefault is the default value set for the MaximumNumAttempts policy option (assigned if no
	// MaximumNumAttempts option is passed to the operation). There is no maximum value parameter, as it is possible to
	// specify 'unlimited', however there is an implicit maximum by the upper bound of uint in golang.
	MaximumNumAttemptsDefault = uint(10)

	// MaximumTimeoutMinimum is the minimum value that can be set for the MaximumTimeout policy option.
	// This option accepts zero, which indicates an unlimited timeout.
	MaximumTimeoutMinimum = 0 * time.Second

	// MaximumTimeoutDefault is the default value set for the MaximumTimeout policy option (assigned if no
	// MaximumTimeout option is passed to the operation).
	MaximumTimeoutDefault = 20 * time.Minute
)

// OciRetryableRequest represents a request that can be reissued according to the specified policy.
type OciRetryableRequest interface {
	// Any retryable request must have an underlying request
	OciRequest

	// Each operation specifies default retry behavior. By passing no arguments to this method, the default retry
	// behavior, as determined on a per-operation-basis, will be honored. Variadic retry policy option arguments
	// passed to this method will override the default behavior.
	GetRetryPolicy() *RetryPolicy
}

// RetryPolicy is the class that holds all relevant information for retrying operations.
type RetryPolicy struct {
	// MaximumNumberAttempts is the maximum number of times to retry a request. Zero indicates an unlimited
	// number of attempts.
	MaximumNumberAttempts uint

	// MaximumTimeout is the total duration to wait for an operation, including all retries.
	MaximumTimeout time.Duration

	// ShouldRetryOperation inspects the http response, error, and operation attempt number, and
	// - returns true if we should retry the operation
	// - returns false otherwise
	ShouldRetryOperation func(OciResponse, error, uint) bool

	// GetNextDuration computes the duration to pause between operation retries.
	GetNextDuration func(uint) time.Duration
}

// RetryPolicyOption exposes a function that allows us to set values on the underlying RetryPolicy.
type RetryPolicyOption func(policy *RetryPolicy)

// MaximumNumberAttempts sets the value for the corresponding retry policy option.
func MaximumNumberAttempts(value uint) RetryPolicyOption {
	return func(policy *RetryPolicy) {
		validated := value

		if value < MaximumNumAttemptsMinimum {
			validated = MaximumNumAttemptsMinimum
		}

		policy.MaximumNumberAttempts = validated
	}
}

// MaximumTimeout sets the value for the corresponding retry policy option.
func MaximumTimeout(value time.Duration) RetryPolicyOption {
	return func(policy *RetryPolicy) {
		validated := value

		if value < MaximumTimeoutMinimum {
			validated = MaximumTimeoutMinimum
		}

		policy.MaximumTimeout = validated
	}
}

// DefaultShouldRetryOperation is the default function for ShouldRetryOperation, if one is not specified.
func DefaultShouldRetryOperation(response OciResponse, e error, currentAttempt uint) bool {
	if e != nil {
		return true
	}

	if response.GetRawResponse().StatusCode > 404 {
		return true
	}

	return false
}

// DefaultGetNextDuration is the default function for GetNextDuration, if one is not specified.
func DefaultGetNextDuration(attempts uint) time.Duration {
	duration := math.Pow(float64(DefaultExponentialBackoffBase), float64(attempts-1))
	return DefaultExponentialBackoffFactor * time.Duration(duration)
}

// ShouldRetryOperation sets the value for the corresponding retry policy option.
func ShouldRetryOperation(value func(OciResponse, error, uint) bool) RetryPolicyOption {
	if value == nil {
		value = DefaultShouldRetryOperation
	}
	return func(policy *RetryPolicy) {
		policy.ShouldRetryOperation = value
	}
}

// GetNextDuration sets the value for the corresponding retry policy option.
func GetNextDuration(value func(attempts uint) time.Duration) RetryPolicyOption {
	if value == nil {
		value = DefaultGetNextDuration
	}
	return func(policy *RetryPolicy) {
		policy.GetNextDuration = value
	}
}

// BuildRetryPolicy accepts a variadic number of retry policy option values and assembles a retry policy. If any
// policy option is not specified, the default is used.
func BuildRetryPolicy(options ...RetryPolicyOption) RetryPolicy {
	policy := RetryPolicy{
		MaximumTimeout:        MaximumTimeoutDefault,
		MaximumNumberAttempts: MaximumNumAttemptsDefault,
		ShouldRetryOperation:  DefaultShouldRetryOperation,
		GetNextDuration:       DefaultGetNextDuration,
	}

	for _, option := range options {
		option(&policy)
	}

	return policy
}

// NoRetryPolicy is a helper method that assembles and returns a return policy that indicates an operation should
// never be retried.
func NoRetryPolicy() *RetryPolicy {
	return &BuildRetryPolicy(
		MaximumNumberAttempts(1),
		ShouldRetryOperation(func(OciResponse, error, uint) bool { return false }),
	)
}
