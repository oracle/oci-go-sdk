package common

import (
	"context"
	"math"
	"net/http"
	"reflect"
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

// Retrier interface is implemented at the HTTP request level, allowing for operation-level retry policies.
type Retrier interface {
	Call(ctx context.Context, request *http.Request, options ...Option) (response *http.Response, err error)
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
	ShouldRetryOperation func(*http.Response, error, uint) bool

	// GetNextDuration computes the duration to pause between operation retries.
	GetNextDuration func(attempts uint) time.Duration
}

const (
	retryPolicyOptionIdentifier            string = "RETRY"
	retryPolicyOptionMaximumNumberAttempts string = "MaximumNumberAttempts"
	retryPolicyOptionMaximumTimeout        string = "MaximumTimeout"
	retryPolicyOptionShouldRetryOperation  string = "ShouldRetryOperation"
	retryPolicyOptionGetNextDuration       string = "GetNextDuration"
)

type requestOption struct {
	optionType  string
	optionName  string
	optionValue interface{}
}

// Option exposes a function that allows us to set options associated with the request.
// Currently the only supported functionality surrounds Retry Policy options.
type Option func(options *[]requestOption)

// MaximumNumberAttempts sets the value for the corresponding retry policy option.
func MaximumNumberAttempts(value uint) Option {
	return func(options *[]requestOption) {
		option := requestOption{
			optionType:  retryPolicyOptionIdentifier,
			optionName:  retryPolicyOptionMaximumNumberAttempts,
			optionValue: value,
		}
		*options = append(*options, option)
	}
}

// MaximumTimeout sets the value for the corresponding retry policy option.
func MaximumTimeout(value time.Duration) Option {
	return func(options *[]requestOption) {
		validated := value

		if value < MaximumTimeoutMinimum {
			validated = MaximumTimeoutMinimum
		}

		option := requestOption{
			optionType:  retryPolicyOptionIdentifier,
			optionName:  retryPolicyOptionMaximumTimeout,
			optionValue: validated,
		}
		*options = append(*options, option)
	}
}

// DefaultShouldRetryOperation is the default function for ShouldRetryOperation, if one is not specified.
func DefaultShouldRetryOperation(response *http.Response, e error, currentAttempt uint) bool {
	if e != nil {
		return true
	}

	if response.StatusCode < 405 {
		return false
	}

	return true
}

// DefaultGetNextDuration is the default function for GetNextDuration, if one is not specified.
func DefaultGetNextDuration(attempts uint) time.Duration {
	duration := math.Pow(float64(DefaultExponentialBackoffBase), float64(attempts-1))
	return DefaultExponentialBackoffFactor * time.Duration(duration)
}

// ShouldRetryOperation sets the value for the corresponding retry policy option.
func ShouldRetryOperation(value func(*http.Response, error, uint) bool) Option {
	return func(options *[]requestOption) {
		if value == nil {
			value = DefaultShouldRetryOperation
		}

		option := requestOption{
			optionType:  retryPolicyOptionIdentifier,
			optionName:  retryPolicyOptionShouldRetryOperation,
			optionValue: value,
		}
		*options = append(*options, option)
	}
}

// GetNextDuration sets the value for the corresponding retry policy option.
func GetNextDuration(value func(attempts uint) time.Duration) Option {
	return func(options *[]requestOption) {
		if value == nil {
			value = DefaultGetNextDuration
		}

		option := requestOption{
			optionType:  retryPolicyOptionIdentifier,
			optionName:  retryPolicyOptionGetNextDuration,
			optionValue: value,
		}
		*options = append(*options, option)
	}
}

func updateRetryPolicyWithOptions(validated []requestOption, policy *RetryPolicy) {
	for _, option := range validated {
		if option.optionType == retryPolicyOptionIdentifier {
			policyField := reflect.ValueOf(policy).Elem().FieldByName(option.optionName)
			if policyField.IsValid() && policyField.CanSet() {
				policyField.Set(reflect.ValueOf(option.optionValue))
			}
		}
	}
}

// BuildRetryPolicy accepts a variadic number of retry policy option values and assembles a retry policy. If any
// policy option is not specified, the default is used.
func BuildRetryPolicy(options ...Option) RetryPolicy {
	policy := RetryPolicy{
		MaximumTimeout:        MaximumTimeoutDefault,
		MaximumNumberAttempts: MaximumNumAttemptsDefault,
		ShouldRetryOperation:  DefaultShouldRetryOperation,
		GetNextDuration:       DefaultGetNextDuration,
	}

	validatedOptions := []requestOption{}
	for _, option := range options {
		option(&validatedOptions)
	}

	updateRetryPolicyWithOptions(validatedOptions, &policy)

	return policy
}

// NoRetryPolicy is a helper method that assembles and returns a return policy that indicates an operation should
// never be retried.
func NoRetryPolicy() RetryPolicy {
	return BuildRetryPolicy(
		MaximumNumberAttempts(1),
		ShouldRetryOperation(func(*http.Response, error, uint) bool { return false }),
	)
}
