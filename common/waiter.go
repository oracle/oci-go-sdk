package common

import (
	"context"
	"fmt"
	"time"
)

// OciPollable interface is implemented by models and allows users to inspect the state of a
// server-side operation.
type OciPollable interface {
	// GetStatefulIndicator returns an enum representing the current state of the operation.
	GetStatefulIndicator() string
}

// OciStatefulResponse interface is implemented by responses to indicate that the underlying object
// implements the OciPollable interface.
type OciStatefulResponse interface {
	// Any stateful response must have an underlying response
	OciResponse

	// Returns the underlying model that contains a stateful indicator (such as lifecycleState).
	GetStatefulEntity() OciPollable
}

// OciPollableOperation is the operation performed against
type OciPollableOperation func(context.Context, OciRetryableRequest) (OciStatefulResponse, error)

// OciPredicate is a generalization of any response parsing function that will determine if the polling condition
// has been met. Returns true if we should stop polling, and false otherwise (to continue polling).
type OciPredicate func(OciStatefulResponse, error) bool

// Poll polls a resource until the specified predicate returns true
func (client BaseClient) Poll(ctx context.Context, request OciRetryableRequest, operation OciPollableOperation, predicate OciPredicate, options ...RetryPolicyOption) error {
	policy := request.GetRetryPolicy(options...)
	deadlineContext, deadlineCancel := context.WithTimeout(ctx, GetMaximumTimeout(&policy))
	defer deadlineCancel()

	var response OciStatefulResponse
	var err error

	for currentOperationAttempt := uint(1); ShouldContinueIssuingRequests(currentOperationAttempt, policy.MaximumNumberAttempts); currentOperationAttempt++ {
		response, err = operation(deadlineContext, request)

		select {
		case <-deadlineContext.Done():
			return deadlineContext.Err()
		default:
			// non-blocking select
		}

		if predicate(response, err) {
			return nil
		}

		if policy.ShouldRetryOperation(response, err, currentOperationAttempt) {
			if currentOperationAttempt != policy.MaximumNumberAttempts {
				// sleep before retrying the operation
				duration := policy.GetNextDuration(currentOperationAttempt)
				if deadline, ok := deadlineContext.Deadline(); ok && time.Now().Add(duration).After(deadline) {
					return DeadlineExceededByBackoff
				}
				time.Sleep(duration)
			}
		} else {
			// we should NOT retry operation based on response and/or error => return
			return err
		}
	}
	return fmt.Errorf("maximum number of attempts exceeded (%v)", policy.MaximumNumberAttempts)
}
