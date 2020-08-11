// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Support Management API
//
// Use the Support Management API to manage support requests. For more information, see Getting Help and Contacting Support (https://docs.cloud.oracle.com/iaas/Content/GSG/Tasks/contactingsupport.htm). **Note**: Before you can create service requests with this API, you need to have an Oracle Single Sign On (SSO) account, and you need to register your Customer Support Identifier (CSI) with My Oracle Support.
//

package cims

// ErrorCodeEnum Enum with underlying type: string
type ErrorCodeEnum string

// Set of constants representing the allowable values for ErrorCodeEnum
const (
	ErrorCodeContentEmpty                       ErrorCodeEnum = "CONTENT_EMPTY"
	ErrorCodeClientException                    ErrorCodeEnum = "CLIENT_EXCEPTION"
	ErrorCodeInvalidFormat                      ErrorCodeEnum = "INVALID_FORMAT"
	ErrorCodeInvalidJsonInput                   ErrorCodeEnum = "INVALID_JSON_INPUT"
	ErrorCodeSslAuthorization                   ErrorCodeEnum = "SSL_AUTHORIZATION"
	ErrorCodeAuthFailed                         ErrorCodeEnum = "AUTH_FAILED"
	ErrorCodeCsiNotAuthorized                   ErrorCodeEnum = "CSI_NOT_AUTHORIZED"
	ErrorCodeUserPolicyNotAuthorized            ErrorCodeEnum = "USER_POLICY_NOT_AUTHORIZED"
	ErrorCodeEmailNotVerified                   ErrorCodeEnum = "EMAIL_NOT_VERIFIED"
	ErrorCodeEmailNotFound                      ErrorCodeEnum = "EMAIL_NOT_FOUND"
	ErrorCodeIdcsEmailNotValid                  ErrorCodeEnum = "IDCS_EMAIL_NOT_VALID"
	ErrorCodeInvalidPath                        ErrorCodeEnum = "INVALID_PATH"
	ErrorCodeMethodNotAllowed                   ErrorCodeEnum = "METHOD_NOT_ALLOWED"
	ErrorCodeJsonProcessing                     ErrorCodeEnum = "JSON_PROCESSING"
	ErrorCodeGenericException                   ErrorCodeEnum = "GENERIC_EXCEPTION"
	ErrorCodeExternalServiceProviderUnavailable ErrorCodeEnum = "EXTERNAL_SERVICE_PROVIDER_UNAVAILABLE"
	ErrorCodeExternalServiceProviderTimeout     ErrorCodeEnum = "EXTERNAL_SERVICE_PROVIDER_TIMEOUT"
	ErrorCodeTooManyRequests                    ErrorCodeEnum = "TOO_MANY_REQUESTS"
	ErrorCodeIdpScimNotSetup                    ErrorCodeEnum = "IDP_SCIM_NOT_SETUP"
	ErrorCodeIncidentNotFound                   ErrorCodeEnum = "INCIDENT_NOT_FOUND"
	ErrorCodeInvalidUserCsi                     ErrorCodeEnum = "INVALID_USER_CSI"
	ErrorCodeDataAlreadyExists                  ErrorCodeEnum = "DATA_ALREADY_EXISTS"
	ErrorCodeAuthUserNotMatching                ErrorCodeEnum = "AUTH_USER_NOT_MATCHING"
)

var mappingErrorCode = map[string]ErrorCodeEnum{
	"CONTENT_EMPTY":                         ErrorCodeContentEmpty,
	"CLIENT_EXCEPTION":                      ErrorCodeClientException,
	"INVALID_FORMAT":                        ErrorCodeInvalidFormat,
	"INVALID_JSON_INPUT":                    ErrorCodeInvalidJsonInput,
	"SSL_AUTHORIZATION":                     ErrorCodeSslAuthorization,
	"AUTH_FAILED":                           ErrorCodeAuthFailed,
	"CSI_NOT_AUTHORIZED":                    ErrorCodeCsiNotAuthorized,
	"USER_POLICY_NOT_AUTHORIZED":            ErrorCodeUserPolicyNotAuthorized,
	"EMAIL_NOT_VERIFIED":                    ErrorCodeEmailNotVerified,
	"EMAIL_NOT_FOUND":                       ErrorCodeEmailNotFound,
	"IDCS_EMAIL_NOT_VALID":                  ErrorCodeIdcsEmailNotValid,
	"INVALID_PATH":                          ErrorCodeInvalidPath,
	"METHOD_NOT_ALLOWED":                    ErrorCodeMethodNotAllowed,
	"JSON_PROCESSING":                       ErrorCodeJsonProcessing,
	"GENERIC_EXCEPTION":                     ErrorCodeGenericException,
	"EXTERNAL_SERVICE_PROVIDER_UNAVAILABLE": ErrorCodeExternalServiceProviderUnavailable,
	"EXTERNAL_SERVICE_PROVIDER_TIMEOUT":     ErrorCodeExternalServiceProviderTimeout,
	"TOO_MANY_REQUESTS":                     ErrorCodeTooManyRequests,
	"IDP_SCIM_NOT_SETUP":                    ErrorCodeIdpScimNotSetup,
	"INCIDENT_NOT_FOUND":                    ErrorCodeIncidentNotFound,
	"INVALID_USER_CSI":                      ErrorCodeInvalidUserCsi,
	"DATA_ALREADY_EXISTS":                   ErrorCodeDataAlreadyExists,
	"AUTH_USER_NOT_MATCHING":                ErrorCodeAuthUserNotMatching,
}

// GetErrorCodeEnumValues Enumerates the set of values for ErrorCodeEnum
func GetErrorCodeEnumValues() []ErrorCodeEnum {
	values := make([]ErrorCodeEnum, 0)
	for _, v := range mappingErrorCode {
		values = append(values, v)
	}
	return values
}
