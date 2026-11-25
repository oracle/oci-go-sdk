// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.

package common

import (
	"context"
	"net/http"
	"net/url"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// TestRetry_RaceCondition demonstrates the data race bug in v65.101.1
//
// To reproduce the race condition, run:
//
//	go test -race -run TestRetry_RaceCondition ./common/
//
// Expected output with BUGGY code:
//
//	WARNING: DATA RACE
//	Write at 0x... by goroutine X:
//	  common/retry.go:871
//	Previous read at 0x... by goroutine Y:
//	  common/retry.go:907
//
// The race occurs because:
// 1. Line 807: var response OCIResponse (shared variable)
// 2. Line 871: Goroutine writes to 'response'
// 3. Line 907: Main thread reads 'response' when context is cancelled
//
// This creates a data race where both threads access the same memory
// location concurrently without synchronization.
func TestRetry_RaceCondition(t *testing.T) {
	shouldRetryOperation := func(res OCIOperationResponse) bool { return true }
	getNextDuration := func(res OCIOperationResponse) time.Duration { return 100 * time.Millisecond }

	pol := NewRetryPolicy(uint(5), shouldRetryOperation, getNextDuration)
	req := retryableOCIRequest{retryPolicy: &pol}

	testURL, _ := url.Parse("http://example.com/test")
	fakeOperation := func(ctx context.Context, request OCIRequest, _ *OCIReadSeekCloser, _ map[string]string) (OCIResponse, error) {
		// Simulate slow operation to increase chance of race
		time.Sleep(50 * time.Millisecond)

		httpResponse := http.Response{
			Header:     http.Header{},
			StatusCode: 500,
			Request: &http.Request{
				Method: "GET",
				URL:    testURL,
			},
		}
		return genericOCIResponse{RawResponse: &httpResponse}, nil
	}

	// Cancel context immediately to trigger the race condition
	// This causes the main thread to read 'response' at line 907
	// while the goroutine is writing to it at line 871
	ctx, cancelFn := context.WithCancel(context.Background())
	cancelFn()

	_, err := Retry(ctx, req, fakeOperation, pol)

	// Test may pass, but race detector will catch the bug
	assert.Equal(t, context.Canceled, err)
}