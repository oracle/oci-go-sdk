// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.

package common

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRedactSensitiveStringForLogs_RedactsSensitiveHeadersFromLoggedText(t *testing.T) {
	value := "Authorization: Signature keyId=\"secret\"\n" +
		"X-Api-Key: api-key-value\n" +
		"opc-obo-token: obo-token-value"

	sanitized := RedactSensitiveStringForLogs(value)

	assert.Equal(t,
		"Authorization: REDACTED\nX-Api-Key: REDACTED\nopc-obo-token: REDACTED",
		sanitized,
	)
}

func TestRedactSensitiveStringForLogs_RedactsSensitiveHeadersUsingNormalizedNamesAndCredentialRules(t *testing.T) {
	value := "Client_Secret: client-secret\n" +
		"service__private-key: private-key\n" +
		"OPC_DELEGATION_TOKEN: delegation-token\n" +
		"proxy-authorization: proxy-authorization\n" +
		"cookie: cookie\n" +
		"set-cookie: set-cookie\n" +
		"security-context: security-context\n" +
		"password: password\n" +
		"passphrase: passphrase\n" +
		"x-token: token\n" +
		"x-token-custom: token\n" +
		"x-authorization: authorization\n" +
		"x-authorization-token: authorization\n" +
		"x-key-secret: key\n" +
		"idcs-access-token: access-token\n" +
		"service_private_key: private-key\n" +
		"tokenization-status: visible\n" +
		"private-key-id: visible\n" +
		"client-secret-version: visible\n" +
		"x-key: visible\n" +
		"x-keyring-name: visible"

	sanitized := RedactSensitiveStringForLogs(value)

	assert.Contains(t, sanitized, "Client_Secret: REDACTED")
	assert.Contains(t, sanitized, "service__private-key: REDACTED")
	assert.Contains(t, sanitized, "OPC_DELEGATION_TOKEN: REDACTED")
	assert.Contains(t, sanitized, "proxy-authorization: REDACTED")
	assert.Contains(t, sanitized, "cookie: REDACTED")
	assert.Contains(t, sanitized, "set-cookie: REDACTED")
	assert.Contains(t, sanitized, "security-context: REDACTED")
	assert.Contains(t, sanitized, "password: REDACTED")
	assert.Contains(t, sanitized, "passphrase: REDACTED")
	assert.Contains(t, sanitized, "x-token: REDACTED")
	assert.Contains(t, sanitized, "x-token-custom: REDACTED")
	assert.Contains(t, sanitized, "x-authorization: REDACTED")
	assert.Contains(t, sanitized, "x-authorization-token: REDACTED")
	assert.Contains(t, sanitized, "x-key-secret: REDACTED")
	assert.Contains(t, sanitized, "idcs-access-token: REDACTED")
	assert.Contains(t, sanitized, "service_private_key: REDACTED")
	assert.Contains(t, sanitized, "tokenization-status: visible")
	assert.Contains(t, sanitized, "private-key-id: visible")
	assert.Contains(t, sanitized, "client-secret-version: visible")
	assert.Contains(t, sanitized, "x-key: visible")
	assert.Contains(t, sanitized, "x-keyring-name: visible")
}

func TestRedactSensitiveStringForLogs_DoesNotRedactNonHeaderJSONFieldsFromLoggedText(t *testing.T) {
	value := `{"token":"token-value","clientSecret":"client-secret","currentToken":"current-token","pass_phrase":"pass-phrase","podKey":"pod-key","name":"visible"}`

	sanitized := RedactSensitiveStringForLogs(value)

	assert.Contains(t, sanitized, `"token":"token-value"`)
	assert.Contains(t, sanitized, `"clientSecret":"client-secret"`)
	assert.Contains(t, sanitized, `"currentToken":"current-token"`)
	assert.Contains(t, sanitized, `"pass_phrase":"pass-phrase"`)
	assert.Contains(t, sanitized, `"podKey":"pod-key"`)
	assert.Contains(t, sanitized, `"name":"visible"`)
}

func TestRedactSensitiveStringForLogs_RedactsSensitiveHeaderKeysFromSerializedLoggerObjects(t *testing.T) {
	value := `{"headers":{"authorization":"signature","cookie":"cookie","client_secret":"client-secret","x-token-custom":"token","set-cookie":["first-cookie","second-cookie"],"security-token":null,"private-key-id":"visible"}}`

	sanitized := RedactSensitiveStringForLogs(value)

	assert.Contains(t, sanitized, `"authorization":"REDACTED"`)
	assert.Contains(t, sanitized, `"cookie":"REDACTED"`)
	assert.Contains(t, sanitized, `"client_secret":"REDACTED"`)
	assert.Contains(t, sanitized, `"x-token-custom":"REDACTED"`)
	assert.Contains(t, sanitized, `"set-cookie":["REDACTED","REDACTED"]`)
	assert.Contains(t, sanitized, `"security-token":null`)
	assert.Contains(t, sanitized, `"private-key-id":"visible"`)
}

func TestRedactSensitiveStringForLogs_DoesNotRedactSensitiveLookingQueryParametersFromLoggedText(t *testing.T) {
	value := "GET https://example.com/path?mediaAssetId=asset&token=token-value&client_secret=client-secret subject_token=subject-token"

	sanitized := RedactSensitiveStringForLogs(value)

	assert.Equal(t, value, sanitized)
}

func TestRedactSensitiveStringForLogs_DoesNotRedactNonHeaderFieldsFromSerializedLoggerObjects(t *testing.T) {
	value := `{"response":{"url":"https://example.com/path?token=token-value"},"errorObject":{"requestEndpoint":"GET https://example.com/path?password=password-value","message":{"privateKey":"private-key","displayName":"visible"}}}`

	sanitized := RedactSensitiveStringForLogs(value)

	assert.Contains(t, sanitized, `"url":"https://example.com/path?token=token-value"`)
	assert.Contains(t, sanitized, `"requestEndpoint":"GET https://example.com/path?password=password-value"`)
	assert.Contains(t, sanitized, `"privateKey":"private-key"`)
	assert.Contains(t, sanitized, `"displayName":"visible"`)
	assert.Equal(t, `{"response":{"url":"https://example.com/path?token=token-value"},"errorObject":{"requestEndpoint":"GET https://example.com/path?password=password-value","message":{"privateKey":"private-key","displayName":"visible"}}}`, value)
}

func TestRedactSensitiveStringForLogs_DoesNotRequireCallersToMutateStoredValues(t *testing.T) {
	endpoint := "GET https://example.com/path?token=token-value"

	sanitized := RedactSensitiveStringForLogs(endpoint)

	assert.Equal(t, "GET https://example.com/path?token=token-value", endpoint)
	assert.Equal(t, endpoint, sanitized)
}

func TestRedactSensitiveHeadersForLogs(t *testing.T) {
	servicePrivateKeyHeader := http.CanonicalHeaderKey("Service_Private_Key")
	headers := http.Header{
		"Authorization":         {"Bearer top-secret"},
		"Set-Cookie":            {"a=1", "b=2"},
		servicePrivateKeyHeader: {"private-key"},
		"Session-Token":         nil,
		"X-Key":                 {"keep"},
	}

	redacted := RedactSensitiveHeadersForLogs(headers)

	assert.Nil(t, RedactSensitiveHeadersForLogs(nil))
	assert.Empty(t, RedactSensitiveHeadersForLogs(http.Header{}))
	assert.Equal(t, []string{"REDACTED"}, redacted["Authorization"])
	assert.Equal(t, []string{"REDACTED", "REDACTED"}, redacted["Set-Cookie"])
	assert.Equal(t, []string{"REDACTED"}, redacted[servicePrivateKeyHeader])
	assert.Nil(t, redacted["Session-Token"])
	assert.Equal(t, []string{"keep"}, redacted["X-Key"])
	assert.Equal(t, []string{"Bearer top-secret"}, headers["Authorization"])
}
