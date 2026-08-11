// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.

package common

import (
	"encoding/json"
	"io"
	"reflect"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestStructToString(t *testing.T) {
	ints := []int{1, 2, 4}

	s := struct {
		Anint        *int
		AString      *string
		AFloat       *float32
		SimpleString string
		SimepleInt   int
		IntSlice     *[]int
	}{Int(1),
		String("one"),
		Float32(2.3),
		"simple",
		2, &ints}

	str := PointerString(s)
	assert.Contains(t, str, "one")
	assert.Contains(t, str, "1")
	assert.Contains(t, str, "[1 2 4]")
}

type sample struct {
	Anint        *int
	AString      *string
	AFloat       *float32
	SimpleString string
	SimepleInt   int
	Nested       sampleNested
}

func (s sample) String() string {
	str := PointerString(s)
	return str
}

type sampleNested struct {
	NestedInt    *int
	NestedBool   *bool
	NestedString *string
	Thestring    string
}

func (s sampleNested) String() string {
	str := PointerString(s)
	return str
}

func TestStructToString_Nested(t *testing.T) {
	s := sample{Anint: Int(1),
		AString:      nil,
		AFloat:       Float32(2.3),
		SimpleString: "simple",
		SimepleInt:   2,
	}
	s.Nested.NestedBool = Bool(true)
	s.Nested.NestedString = nil
	s.Nested.Thestring = "somestring"
	s.Nested.NestedInt = Int(2)

	str := s.String()
	assert.Contains(t, str, "1")
	assert.Contains(t, str, "somestring")
	assert.Contains(t, str, "<nil>")
}

type sensitiveTestPayload struct {
	Password       *string           `json:"password" sensitive:"true"`
	Nested         sensitiveTestItem `json:"nested"`
	Passwords      []string          `json:"passwords" sensitive:"true"`
	PasswordByName map[string]string `json:"passwordByName" sensitive:"true"`
}

func (p sensitiveTestPayload) String() string {
	return PointerString(p)
}

type sensitiveTestItem struct {
	Password *string `json:"password" sensitive:"true"`
}

func (i sensitiveTestItem) String() string {
	return PointerString(i)
}

type sensitiveTestRequest struct {
	Details sensitiveTestPayload `contributesTo:"body"`
}

func (r sensitiveTestRequest) String() string {
	return PointerString(r)
}

func TestPointerStringRedactsSensitiveValuesWhileRequestSerializationPreservesThem(t *testing.T) {
	const canary = "go-sdk-password-canary"

	payload := sensitiveTestPayload{
		Password:  String(canary),
		Nested:    sensitiveTestItem{Password: String(canary)},
		Passwords: []string{canary},
		PasswordByName: map[string]string{
			"primary": canary,
		},
	}

	diagnosticOutput := sensitiveTestRequest{Details: payload}.String()
	assert.NotContains(t, diagnosticOutput, canary)
	assert.Contains(t, diagnosticOutput, "Password=<redacted>")
	assert.Contains(t, diagnosticOutput, "Passwords=<redacted>")
	assert.Contains(t, diagnosticOutput, "PasswordByName=<redacted>")

	httpRequest, err := MakeDefaultHTTPRequestWithTaggedStruct("POST", "/", sensitiveTestRequest{Details: payload})
	assert.NoError(t, err)

	requestBody, err := io.ReadAll(httpRequest.Body)
	assert.NoError(t, err)
	assert.Equal(t, 4, strings.Count(string(requestBody), canary))
}

func TestDateParsing_LastModifiedHeaderDate(t *testing.T) {
	data := []string{"Tue, 2 Jan 2018 17:49:29 GMT", "Tue, 02 Jan 2018 17:49:29 GMT"}
	for _, val := range data {
		tt, err := tryParsingTimeWithValidFormatsForHeaders([]byte(val), "lastmodified")
		assert.NoError(t, err)
		assert.Equal(t, tt.Day(), 2)
	}
}

func TestFormattedTimeMarshaling(t *testing.T) {
	sampleTime, _ := time.Parse(time.UnixDate, "Mon Jan 02 15:04:05 MST 2006")
	testIO := []struct {
		name          string
		t             *SDKDate
		expectedJSON  string
		expectedError error
	}{
		{
			name:          "formatting time to simple date format",
			t:             &SDKDate{Date: sampleTime},
			expectedJSON:  `"2006-01-02"`,
			expectedError: nil,
		},
		{
			name:          "formatting nil",
			t:             nil,
			expectedJSON:  `null`,
			expectedError: nil,
		},
	}

	for _, tc := range testIO {
		t.Run(tc.name, func(t *testing.T) {
			bytes, e := json.Marshal(&tc.t)
			assert.Equal(t, tc.expectedError, e)
			assert.Equal(t, tc.expectedJSON, string(bytes))
		})
	}

}

func TestFormattedTimeUnMarshaling(t *testing.T) {
	sampleTime, _ := time.Parse(time.UnixDate, "Mon Jan 02 15:04:05 MST 2006")
	testIO := []struct {
		name          string
		json          string
		expectedTime  *SDKDate
		expectedError error
	}{
		{
			name:          "unmarshaling time to simple date format",
			expectedTime:  &SDKDate{Date: sampleTime},
			json:          `"2006-01-02"`,
			expectedError: nil,
		},
		{
			name:          "unmarshaling time to simple RFC3339 format",
			expectedTime:  &SDKDate{Date: sampleTime},
			json:          `"2006-01-02T15:04:05Z"`,
			expectedError: &time.ParseError{},
		},
		{
			name:          "unmarshaling null",
			expectedTime:  &SDKDate{},
			json:          `"null"`,
			expectedError: nil,
		},
	}

	for _, tc := range testIO {
		t.Run(tc.name, func(t *testing.T) {
			newTime := SDKDate{}
			e := json.Unmarshal([]byte(tc.json), &newTime)
			assert.IsType(t, reflect.TypeOf(tc.expectedError), reflect.TypeOf(e))
			if tc.expectedError != nil {
				return
			}
			assert.Equal(t, tc.expectedTime.Date.Format(sdkDateFormat), newTime.Date.Format(sdkDateFormat))
		})
	}

}

func TestSDKDateToAndFromString(t *testing.T) {
	_, err := NewSDKDateFromString("InvalidFormat")
	s, _ := NewSDKDateFromString("2018-09-13")
	str := s.String()

	assert.Equal(t, "2018-09-13", str)
	assert.IsType(t, &time.ParseError{}, err)
}

func TestMakeACopy(t *testing.T) {
	original := []string{"a", "b", "c"}

	copy := makeACopy(original)

	assert.Equal(t, original, copy)
	copy[0] = "mutate"
	assert.NotEqual(t, original, copy)
}
