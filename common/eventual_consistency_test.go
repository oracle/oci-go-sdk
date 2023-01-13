// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.

package common

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestIsErrorAffectedByEventualConsistency(t *testing.T) {
	var error = servicefailure{
		StatusCode: 400,
		Code:       "InvalidParameter"}
	assert.False(t, IsErrorAffectedByEventualConsistency(error))

	error = servicefailure{
		StatusCode: 400,
		Code:       "RelatedResourceNotAuthorizedOrNotFound"}
	assert.True(t, IsErrorAffectedByEventualConsistency(error))

	error = servicefailure{
		StatusCode: 404,
		Code:       "NotFound"}
	assert.False(t, IsErrorAffectedByEventualConsistency(error))

	error = servicefailure{
		StatusCode: 404,
		Code:       "NotAuthorizedOrNotFound"}
	assert.True(t, IsErrorAffectedByEventualConsistency(error))

	error = servicefailure{
		StatusCode: 409,
		Code:       "Conflict"}
	assert.False(t, IsErrorAffectedByEventualConsistency(error))

	error = servicefailure{
		StatusCode: 409,
		Code:       "NotAuthorizedOrResourceAlreadyExists"}
	assert.True(t, IsErrorAffectedByEventualConsistency(error))

	error = servicefailure{
		StatusCode: 400,
		Code:       "InsufficientServicePermissions"}
	assert.True(t, IsErrorAffectedByEventualConsistency(error))

	error = servicefailure{
		StatusCode: 400,
		Code:       "ResourceDisabled"}
	assert.True(t, IsErrorAffectedByEventualConsistency(error))
}

// tests for EC communication mode: file

func TestRetryGetEcMode(t *testing.T) {
	assert.Equal(t, InProcess, getEcMode("inprocess"))
	assert.Equal(t, InProcess, getEcMode("InProcess"))
	assert.Equal(t, InProcess, getEcMode("INPROCESS"))

	assert.Equal(t, File, getEcMode("file"))
	assert.Equal(t, File, getEcMode("File"))
	assert.Equal(t, File, getEcMode("FILE"))

	assert.Equal(t, InProcess, getEcMode("unknown"))
}

func TestRetryNewEcContext(t *testing.T) {
	savedValue, savedValueProvided := os.LookupEnv(OciGoSdkEcConfigEnvVarName)
	defer func() {
		if savedValueProvided {
			os.Setenv(OciGoSdkEcConfigEnvVarName, savedValue)
		} else {
			os.Unsetenv(OciGoSdkEcConfigEnvVarName)
		}
	}()

	os.Unsetenv(OciGoSdkEcConfigEnvVarName)

	var ecContext = newEcContext()
	assert.Equal(t, InProcess, ecContext.ecMode)

	os.Setenv(OciGoSdkEcConfigEnvVarName, "file,/tmp/foo")

	ecContext = newEcContext()
	assert.Equal(t, File, ecContext.ecMode)

	os.Setenv(OciGoSdkEcConfigEnvVarName, "file")

	ecContext = newEcContext()
	assert.Equal(t, (*EventuallyConsistentContext)(nil), ecContext)

	os.Setenv(OciGoSdkEcConfigEnvVarName, "inprocess")

	ecContext = newEcContext()
	assert.Equal(t, InProcess, ecContext.ecMode)

	os.Setenv(OciGoSdkEcConfigEnvVarName, "unknown")

	ecContext = newEcContext()
	assert.Equal(t, InProcess, ecContext.ecMode)
}

func TestRetryFileEcContext(t *testing.T) {
	savedValue, savedValueProvided := os.LookupEnv(OciGoSdkEcConfigEnvVarName)
	defer func() {
		if savedValueProvided {
			os.Setenv(OciGoSdkEcConfigEnvVarName, savedValue)
		} else {
			os.Unsetenv(OciGoSdkEcConfigEnvVarName)
		}
	}()

	// File timestamps don't maintain monotonic clock.
	// The canonical way to strip a monotonic clock reading is to use t = t.Round(0).
	now := time.Now().Round(0)
	fileName := filepath.Join(os.TempDir(), fmt.Sprintf("TestRetryFileEcContext-%d", now.Unix()))
	var f, err = os.Create(fileName)
	assert.Equal(t, nil, err)
	f.Close()
	err = os.Remove(fileName)
	assert.Equal(t, nil, err)

	os.Setenv(OciGoSdkEcConfigEnvVarName, "file,"+f.Name())

	var ecContext = newEcContext()
	assert.Equal(t, File, ecContext.ecMode)

	ecContext.readLock(ecContext)
	ecContext.readUnlock(ecContext)

	ecContext.writeLock(ecContext)
	ecContext.writeUnlock(ecContext)

	// test no EC
	var endOfWindowTime = ecContext.GetEndOfWindow()
	assert.Equal(t, (*time.Time)(nil), endOfWindowTime)

	// test EC window ends in the past
	var timeNowProviderValues = make(chan time.Time, 300)
	timeNowProviderValues <- now                      // most recent EC
	timeNowProviderValues <- now.Add(5 * time.Minute) // initial request, already outside the EC window
	ecContext.timeNowProvider = func() time.Time {
		nextTime := <-timeNowProviderValues
		return nextTime
	}
	endOfWindowTime = ecContext.UpdateEndOfWindow(eventuallyConsistentWindowSize)
	assert.Equal(t, now.Add(eventuallyConsistentWindowSize), *endOfWindowTime)
	assert.Equal(t, nil, err)

	endOfWindowTime = ecContext.GetEndOfWindow()
	assert.Equal(t, now.Add(eventuallyConsistentWindowSize), *endOfWindowTime)
}

// EC

func setupFileMode(t *testing.T, testName string) {
	// File timestamps don't maintain monotonic clock.
	// The canonical way to strip a monotonic clock reading is to use t = t.Round(0).
	now := time.Now().Round(0)
	fileName := filepath.Join(os.TempDir(), fmt.Sprintf("%s-%d", testName, now.Unix()))
	var f, err = os.Create(fileName)
	assert.Equal(t, nil, err)
	f.Close()
	err = os.Remove(fileName)
	assert.Equal(t, nil, err)

	os.Setenv(OciGoSdkEcConfigEnvVarName, "file,"+f.Name())

	EcContext = newEcContext()
	assert.Equal(t, File, EcContext.ecMode)
}

func TestRetryFileDeterminePolicyToUse_EcRetryPolicy_NoEc(t *testing.T) {
	savedValue, savedValueProvided := os.LookupEnv(OciGoSdkEcConfigEnvVarName)
	defer func() {
		if savedValueProvided {
			os.Setenv(OciGoSdkEcConfigEnvVarName, savedValue)
		} else {
			os.Unsetenv(OciGoSdkEcConfigEnvVarName)
		}
		EcContext = newEcContext()
	}()
	setupFileMode(t, "TestRetryFileDeterminePolicyToUse_EcRetryPolicy_NoEc")
	resetTimes()

	drp := DefaultRetryPolicyWithoutEventualConsistency()
	rp := EventuallyConsistentRetryPolicy(drp)

	// no eventually consistent effects
	var policyToUse, endOfWindowTime, backoffScalingFactor = rp.DeterminePolicyToUse(rp)
	assertEqualRetryPolicies(t, drp, policyToUse)
	assert.Equal(t, (*time.Time)(nil), endOfWindowTime)
	assert.Equal(t, 1.0, backoffScalingFactor)
}

func TestRetryFileDeterminePolicyToUse_EcRetryPolicy_EcEndsSoonerThanDefaultRetries(t *testing.T) {
	savedValue, savedValueProvided := os.LookupEnv(OciGoSdkEcConfigEnvVarName)
	defer func() {
		if savedValueProvided {
			os.Setenv(OciGoSdkEcConfigEnvVarName, savedValue)
		} else {
			os.Unsetenv(OciGoSdkEcConfigEnvVarName)
		}
		EcContext = newEcContext()
	}()
	setupFileMode(t, "TestRetryFileDeterminePolicyToUse_EcRetryPolicy_EcEndsSoonerThanDefaultRetries")
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

func TestRetryFileDeterminePolicyToUse_EcRetryPolicy_NeedEcRetries(t *testing.T) {
	savedValue, savedValueProvided := os.LookupEnv(OciGoSdkEcConfigEnvVarName)
	defer func() {
		if savedValueProvided {
			os.Setenv(OciGoSdkEcConfigEnvVarName, savedValue)
		} else {
			os.Unsetenv(OciGoSdkEcConfigEnvVarName)
		}
		EcContext = newEcContext()
	}()
	setupFileMode(t, "TestRetryFileDeterminePolicyToUse_EcRetryPolicy_NeedEcRetries")
	resetTimes()

	rp := DefaultRetryPolicy()

	now := setupTimesNeedEcRetries((*time.Time)(nil), rp)

	var policyToUse, endOfWindowTime, backoffScalingFactor = rp.DeterminePolicyToUse(rp)
	assertEqualRetryPolicies(t, rp, policyToUse)
	assert.Equal(t, now.Add(eventuallyConsistentWindowSize), *endOfWindowTime)
	// since the initial request is exactly half-way through the maximum, expect factor of 0.5
	assert.Equal(t, 0.5, backoffScalingFactor)
}

func TestRetryFileDeterminePolicyToUse_EcRetryPolicy_NeedFullEcRetries(t *testing.T) {
	savedValue, savedValueProvided := os.LookupEnv(OciGoSdkEcConfigEnvVarName)
	defer func() {
		if savedValueProvided {
			os.Setenv(OciGoSdkEcConfigEnvVarName, savedValue)
		} else {
			os.Unsetenv(OciGoSdkEcConfigEnvVarName)
		}
		EcContext = newEcContext()
	}()
	setupFileMode(t, "TestRetryFileDeterminePolicyToUse_EcRetryPolicy_NeedFullEcRetries")
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

func TestRetryFileDeterminePolicyToUse_EcRetryPolicy_EndOfWindowInThePast(t *testing.T) {
	savedValue, savedValueProvided := os.LookupEnv(OciGoSdkEcConfigEnvVarName)
	defer func() {
		if savedValueProvided {
			os.Setenv(OciGoSdkEcConfigEnvVarName, savedValue)
		} else {
			os.Unsetenv(OciGoSdkEcConfigEnvVarName)
		}
		EcContext = newEcContext()
	}()
	setupFileMode(t, "TestRetryFileDeterminePolicyToUse_EcRetryPolicy_EndOfWindowInThePast")
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

func TestRetryFileDeterminePolicyToUse_EcRetryPolicy_NoEc_UnlimitedAttempts(t *testing.T) {
	savedValue, savedValueProvided := os.LookupEnv(OciGoSdkEcConfigEnvVarName)
	defer func() {
		if savedValueProvided {
			os.Setenv(OciGoSdkEcConfigEnvVarName, savedValue)
		} else {
			os.Unsetenv(OciGoSdkEcConfigEnvVarName)
		}
		EcContext = newEcContext()
	}()
	setupFileMode(t, "TestRetryFileDeterminePolicyToUse_EcRetryPolicy_NoEc_UnlimitedAttempts")
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

func TestRetryFileDeterminePolicyToUse_EcRetryPolicy_EcEndsSoonerThanDefaultRetries_UnlimitedAttempts(t *testing.T) {
	savedValue, savedValueProvided := os.LookupEnv(OciGoSdkEcConfigEnvVarName)
	defer func() {
		if savedValueProvided {
			os.Setenv(OciGoSdkEcConfigEnvVarName, savedValue)
		} else {
			os.Unsetenv(OciGoSdkEcConfigEnvVarName)
		}
		EcContext = newEcContext()
	}()
	setupFileMode(t, "TestRetryFileDeterminePolicyToUse_EcRetryPolicy_EcEndsSoonerThanDefaultRetries_UnlimitedAttempts")
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

func TestRetryFileDeterminePolicyToUse_EcRetryPolicy_NeedEcRetries_UnlimitedAttempts(t *testing.T) {
	savedValue, savedValueProvided := os.LookupEnv(OciGoSdkEcConfigEnvVarName)
	defer func() {
		if savedValueProvided {
			os.Setenv(OciGoSdkEcConfigEnvVarName, savedValue)
		} else {
			os.Unsetenv(OciGoSdkEcConfigEnvVarName)
		}
		EcContext = newEcContext()
	}()
	setupFileMode(t, "TestRetryFileDeterminePolicyToUse_EcRetryPolicy_NeedEcRetries_UnlimitedAttempts")
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

func TestRetryFileDeterminePolicyToUse_EcRetryPolicy_NeedFullEcRetries_UnlimitedAttempts(t *testing.T) {
	savedValue, savedValueProvided := os.LookupEnv(OciGoSdkEcConfigEnvVarName)
	defer func() {
		if savedValueProvided {
			os.Setenv(OciGoSdkEcConfigEnvVarName, savedValue)
		} else {
			os.Unsetenv(OciGoSdkEcConfigEnvVarName)
		}
		EcContext = newEcContext()
	}()
	setupFileMode(t, "TestRetryFileDeterminePolicyToUse_EcRetryPolicy_NeedFullEcRetries_UnlimitedAttempts")
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

func TestRetryFileDeterminePolicyToUse_EcRetryPolicy_EndOfWindowInThePast_UnlimitedAttempts(t *testing.T) {
	savedValue, savedValueProvided := os.LookupEnv(OciGoSdkEcConfigEnvVarName)
	defer func() {
		if savedValueProvided {
			os.Setenv(OciGoSdkEcConfigEnvVarName, savedValue)
		} else {
			os.Unsetenv(OciGoSdkEcConfigEnvVarName)
		}
		EcContext = newEcContext()
	}()
	setupFileMode(t, "TestRetryFileDeterminePolicyToUse_EcRetryPolicy_EndOfWindowInThePast_UnlimitedAttempts")
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

// Sanity test: acquire read lock, sleep 30 seconds, release read lock
// Can be used in conjunction with SanityTestRetryFileEcContextLockFileWriteLock.
// To use as test, remove the "Sanity" from the function name.
// The behavior that we should then observe is this:
// - If a write lock is held, no other lock can be acquired.
// - If a read lock is held, another read lock can be acquired.
// - If a read lock is held, no write lock can be acquired.
func SanityTestRetryFileEcContextLockFileReadLock(t *testing.T) {
	savedValue, savedValueProvided := os.LookupEnv(OciGoSdkEcConfigEnvVarName)
	defer func() {
		if savedValueProvided {
			os.Setenv(OciGoSdkEcConfigEnvVarName, savedValue)
		} else {
			os.Unsetenv(OciGoSdkEcConfigEnvVarName)
		}
	}()

	fileName := fmt.Sprintf("%s/TestRetryFileEcContextLockFile", os.TempDir())

	var f, err = os.Create(fileName)
	assert.Equal(t, nil, err)
	err = os.Remove(fileName)
	assert.Equal(t, nil, err)

	os.Setenv(OciGoSdkEcConfigEnvVarName, "file,"+f.Name())

	var ecContext = newEcContext()
	assert.Equal(t, File, ecContext.ecMode)

	ecContext.readLock(ecContext)

	fmt.Printf("Acquired read lock\n")
	time.Sleep(30 * time.Second)

	ecContext.readUnlock(ecContext)
	fmt.Printf("Released read lock\n")
}

// Sanity test: acquire write lock, sleep 30 seconds, release write lock
// Can be used in conjunction with SanityTestRetryFileEcContextLockFileReadLock.
// To use as test, remove the "Sanity" from the function name.
// The behavior that we should then observe is this:
// - If a write lock is held, no other lock can be acquired.
// - If a read lock is held, another read lock can be acquired.
// - If a read lock is held, no write lock can be acquired.
func SanityTestRetryFileEcContextLockFileWriteLock(t *testing.T) {
	savedValue, savedValueProvided := os.LookupEnv(OciGoSdkEcConfigEnvVarName)
	defer func() {
		if savedValueProvided {
			os.Setenv(OciGoSdkEcConfigEnvVarName, savedValue)
		} else {
			os.Unsetenv(OciGoSdkEcConfigEnvVarName)
		}
	}()

	fileName := fmt.Sprintf("%s/TestRetryFileEcContextLockFile", os.TempDir())

	var f, err = os.Create(fileName)
	assert.Equal(t, nil, err)
	err = os.Remove(fileName)
	assert.Equal(t, nil, err)

	os.Setenv(OciGoSdkEcConfigEnvVarName, "file,"+f.Name())

	var ecContext = newEcContext()
	assert.Equal(t, File, ecContext.ecMode)

	ecContext.writeLock(ecContext)
	fmt.Printf("Acquired write lock\n")

	time.Sleep(30 * time.Second)

	ecContext.writeUnlock(ecContext)
	fmt.Printf("Released write lock\n")
}
