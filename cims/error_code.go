// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Support Management API
//
// Use the Support Management API to manage support requests. For more information, see Getting Help and Contacting Support (https://docs.cloud.oracle.com/iaas/Content/GSG/Tasks/contactingsupport.htm). **Note**: Before you can create service requests with this API, you need to have an Oracle Single Sign On (SSO) account, and you need to register your Customer Support Identifier (CSI) with My Oracle Support.
//

package cims

import (
	"strings"
)

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

var mappingErrorCodeEnum = map[string]ErrorCodeEnum{
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

var mappingErrorCodeEnumLowerCase = map[string]ErrorCodeEnum{
	"content_empty":                         ErrorCodeContentEmpty,
	"client_exception":                      ErrorCodeClientException,
	"invalid_format":                        ErrorCodeInvalidFormat,
	"invalid_json_input":                    ErrorCodeInvalidJsonInput,
	"ssl_authorization":                     ErrorCodeSslAuthorization,
	"auth_failed":                           ErrorCodeAuthFailed,
	"csi_not_authorized":                    ErrorCodeCsiNotAuthorized,
	"user_policy_not_authorized":            ErrorCodeUserPolicyNotAuthorized,
	"email_not_verified":                    ErrorCodeEmailNotVerified,
	"email_not_found":                       ErrorCodeEmailNotFound,
	"idcs_email_not_valid":                  ErrorCodeIdcsEmailNotValid,
	"invalid_path":                          ErrorCodeInvalidPath,
	"method_not_allowed":                    ErrorCodeMethodNotAllowed,
	"json_processing":                       ErrorCodeJsonProcessing,
	"generic_exception":                     ErrorCodeGenericException,
	"external_service_provider_unavailable": ErrorCodeExternalServiceProviderUnavailable,
	"external_service_provider_timeout":     ErrorCodeExternalServiceProviderTimeout,
	"too_many_requests":                     ErrorCodeTooManyRequests,
	"idp_scim_not_setup":                    ErrorCodeIdpScimNotSetup,
	"incident_not_found":                    ErrorCodeIncidentNotFound,
	"invalid_user_csi":                      ErrorCodeInvalidUserCsi,
	"data_already_exists":                   ErrorCodeDataAlreadyExists,
	"auth_user_not_matching":                ErrorCodeAuthUserNotMatching,
}

// GetErrorCodeEnumValues Enumerates the set of values for ErrorCodeEnum
func GetErrorCodeEnumValues() []ErrorCodeEnum {
	values := make([]ErrorCodeEnum, 0)
	for _, v := range mappingErrorCodeEnum {
		values = append(values, v)
	}
	return values
}

// GetErrorCodeEnumStringValues Enumerates the set of values in String for ErrorCodeEnum
func GetErrorCodeEnumStringValues() []string {
	return []string{
		"CONTENT_EMPTY",
		"CLIENT_EXCEPTION",
		"INVALID_FORMAT",
		"INVALID_JSON_INPUT",
		"SSL_AUTHORIZATION",
		"AUTH_FAILED",
		"CSI_NOT_AUTHORIZED",
		"USER_POLICY_NOT_AUTHORIZED",
		"EMAIL_NOT_VERIFIED",
		"EMAIL_NOT_FOUND",
		"IDCS_EMAIL_NOT_VALID",
		"INVALID_PATH",
		"METHOD_NOT_ALLOWED",
		"JSON_PROCESSING",
		"GENERIC_EXCEPTION",
		"EXTERNAL_SERVICE_PROVIDER_UNAVAILABLE",
		"EXTERNAL_SERVICE_PROVIDER_TIMEOUT",
		"TOO_MANY_REQUESTS",
		"IDP_SCIM_NOT_SETUP",
		"INCIDENT_NOT_FOUND",
		"INVALID_USER_CSI",
		"DATA_ALREADY_EXISTS",
		"AUTH_USER_NOT_MATCHING",
	}
}

// GetMappingErrorCodeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingErrorCodeEnum(val string) (ErrorCodeEnum, bool) {
	enum, ok := mappingErrorCodeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
