// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Support Management API
//
// Use the Support Management API to manage support requests.
// For more information, see Getting Help and Contacting Support (https://docs.cloud.oracle.com/iaas/Content/GSG/Tasks/contactingsupport.htm).
// **Note**: Before you can create service requests with this API,
// complete user registration at My Oracle Cloud Support
// and then ask your tenancy administrator to provide you authorization for the related user groups.
//

package cims

import (
	"strings"
)

// ErrorCodeEnum Enum with underlying type: string
type ErrorCodeEnum string

// Set of constants representing the allowable values for ErrorCodeEnum
const (
	ErrorCodeContentEmpty                         ErrorCodeEnum = "CONTENT_EMPTY"
	ErrorCodeClientException                      ErrorCodeEnum = "CLIENT_EXCEPTION"
	ErrorCodeInvalidFormat                        ErrorCodeEnum = "INVALID_FORMAT"
	ErrorCodeInvalidJsonInput                     ErrorCodeEnum = "INVALID_JSON_INPUT"
	ErrorCodeSslAuthorization                     ErrorCodeEnum = "SSL_AUTHORIZATION"
	ErrorCodeAuthFailed                           ErrorCodeEnum = "AUTH_FAILED"
	ErrorCodeAuthzFailed                          ErrorCodeEnum = "AUTHZ_FAILED"
	ErrorCodeCsiNotAuthorized                     ErrorCodeEnum = "CSI_NOT_AUTHORIZED"
	ErrorCodeUserPolicyNotAuthorized              ErrorCodeEnum = "USER_POLICY_NOT_AUTHORIZED"
	ErrorCodeEmailNotVerified                     ErrorCodeEnum = "EMAIL_NOT_VERIFIED"
	ErrorCodeEmailNotFound                        ErrorCodeEnum = "EMAIL_NOT_FOUND"
	ErrorCodeOciEmailNotFound                     ErrorCodeEnum = "OCI_EMAIL_NOT_FOUND"
	ErrorCodeMosEmailNotFound                     ErrorCodeEnum = "MOS_EMAIL_NOT_FOUND"
	ErrorCodeIdcsEmailNotValid                    ErrorCodeEnum = "IDCS_EMAIL_NOT_VALID"
	ErrorCodeInvalidPath                          ErrorCodeEnum = "INVALID_PATH"
	ErrorCodeMethodNotAllowed                     ErrorCodeEnum = "METHOD_NOT_ALLOWED"
	ErrorCodeJsonProcessing                       ErrorCodeEnum = "JSON_PROCESSING"
	ErrorCodeGenericException                     ErrorCodeEnum = "GENERIC_EXCEPTION"
	ErrorCodeExternalServiceProviderUnavailable   ErrorCodeEnum = "EXTERNAL_SERVICE_PROVIDER_UNAVAILABLE"
	ErrorCodeExternalServiceProviderTimeout       ErrorCodeEnum = "EXTERNAL_SERVICE_PROVIDER_TIMEOUT"
	ErrorCodeTooManyRequests                      ErrorCodeEnum = "TOO_MANY_REQUESTS"
	ErrorCodeIdpScimNotSetup                      ErrorCodeEnum = "IDP_SCIM_NOT_SETUP"
	ErrorCodeIncidentNotFound                     ErrorCodeEnum = "INCIDENT_NOT_FOUND"
	ErrorCodeInvalidUserCsi                       ErrorCodeEnum = "INVALID_USER_CSI"
	ErrorCodeDataAlreadyExists                    ErrorCodeEnum = "DATA_ALREADY_EXISTS"
	ErrorCodeAuthUserNotMatching                  ErrorCodeEnum = "AUTH_USER_NOT_MATCHING"
	ErrorCodeContactNotApproved                   ErrorCodeEnum = "CONTACT_NOT_APPROVED"
	ErrorCodeCreateProfileMosFailure              ErrorCodeEnum = "CREATE_PROFILE_MOS_FAILURE"
	ErrorCodeCreateProfileCreateOssoFailure       ErrorCodeEnum = "CREATE_PROFILE_CREATE_OSSO_FAILURE"
	ErrorCodeCreateProfileIdentityFailure         ErrorCodeEnum = "CREATE_PROFILE_IDENTITY_FAILURE"
	ErrorCodeCreateProfileVerifyOssoFailure       ErrorCodeEnum = "CREATE_PROFILE_VERIFY_OSSO_FAILURE"
	ErrorCodeProfileAccountNotVerified            ErrorCodeEnum = "PROFILE_ACCOUNT_NOT_VERIFIED"
	ErrorCodeSupportAccountNotFound               ErrorCodeEnum = "SUPPORT_ACCOUNT_NOT_FOUND"
	ErrorCodeSupportAccountPendingCsiApproval     ErrorCodeEnum = "SUPPORT_ACCOUNT_PENDING_CSI_APPROVAL"
	ErrorCodeFreeTierCustomerSliUnsupported       ErrorCodeEnum = "FREE_TIER_CUSTOMER_SLI_UNSUPPORTED"
	ErrorCodeProfileAccountVerified               ErrorCodeEnum = "PROFILE_ACCOUNT_VERIFIED"
	ErrorCodeProfileVerifiedCsiRequestPending     ErrorCodeEnum = "PROFILE_VERIFIED_CSI_REQUEST_PENDING"
	ErrorCodeProfileVerifiedCsiRequestNotFound    ErrorCodeEnum = "PROFILE_VERIFIED_CSI_REQUEST_NOT_FOUND"
	ErrorCodeCreateProfileOrganizationNameInvalid ErrorCodeEnum = "CREATE_PROFILE_ORGANIZATION_NAME_INVALID"
	ErrorCodeCreateProfileEmailInvalid            ErrorCodeEnum = "CREATE_PROFILE_EMAIL_INVALID"
)

var mappingErrorCodeEnum = map[string]ErrorCodeEnum{
	"CONTENT_EMPTY":                            ErrorCodeContentEmpty,
	"CLIENT_EXCEPTION":                         ErrorCodeClientException,
	"INVALID_FORMAT":                           ErrorCodeInvalidFormat,
	"INVALID_JSON_INPUT":                       ErrorCodeInvalidJsonInput,
	"SSL_AUTHORIZATION":                        ErrorCodeSslAuthorization,
	"AUTH_FAILED":                              ErrorCodeAuthFailed,
	"AUTHZ_FAILED":                             ErrorCodeAuthzFailed,
	"CSI_NOT_AUTHORIZED":                       ErrorCodeCsiNotAuthorized,
	"USER_POLICY_NOT_AUTHORIZED":               ErrorCodeUserPolicyNotAuthorized,
	"EMAIL_NOT_VERIFIED":                       ErrorCodeEmailNotVerified,
	"EMAIL_NOT_FOUND":                          ErrorCodeEmailNotFound,
	"OCI_EMAIL_NOT_FOUND":                      ErrorCodeOciEmailNotFound,
	"MOS_EMAIL_NOT_FOUND":                      ErrorCodeMosEmailNotFound,
	"IDCS_EMAIL_NOT_VALID":                     ErrorCodeIdcsEmailNotValid,
	"INVALID_PATH":                             ErrorCodeInvalidPath,
	"METHOD_NOT_ALLOWED":                       ErrorCodeMethodNotAllowed,
	"JSON_PROCESSING":                          ErrorCodeJsonProcessing,
	"GENERIC_EXCEPTION":                        ErrorCodeGenericException,
	"EXTERNAL_SERVICE_PROVIDER_UNAVAILABLE":    ErrorCodeExternalServiceProviderUnavailable,
	"EXTERNAL_SERVICE_PROVIDER_TIMEOUT":        ErrorCodeExternalServiceProviderTimeout,
	"TOO_MANY_REQUESTS":                        ErrorCodeTooManyRequests,
	"IDP_SCIM_NOT_SETUP":                       ErrorCodeIdpScimNotSetup,
	"INCIDENT_NOT_FOUND":                       ErrorCodeIncidentNotFound,
	"INVALID_USER_CSI":                         ErrorCodeInvalidUserCsi,
	"DATA_ALREADY_EXISTS":                      ErrorCodeDataAlreadyExists,
	"AUTH_USER_NOT_MATCHING":                   ErrorCodeAuthUserNotMatching,
	"CONTACT_NOT_APPROVED":                     ErrorCodeContactNotApproved,
	"CREATE_PROFILE_MOS_FAILURE":               ErrorCodeCreateProfileMosFailure,
	"CREATE_PROFILE_CREATE_OSSO_FAILURE":       ErrorCodeCreateProfileCreateOssoFailure,
	"CREATE_PROFILE_IDENTITY_FAILURE":          ErrorCodeCreateProfileIdentityFailure,
	"CREATE_PROFILE_VERIFY_OSSO_FAILURE":       ErrorCodeCreateProfileVerifyOssoFailure,
	"PROFILE_ACCOUNT_NOT_VERIFIED":             ErrorCodeProfileAccountNotVerified,
	"SUPPORT_ACCOUNT_NOT_FOUND":                ErrorCodeSupportAccountNotFound,
	"SUPPORT_ACCOUNT_PENDING_CSI_APPROVAL":     ErrorCodeSupportAccountPendingCsiApproval,
	"FREE_TIER_CUSTOMER_SLI_UNSUPPORTED":       ErrorCodeFreeTierCustomerSliUnsupported,
	"PROFILE_ACCOUNT_VERIFIED":                 ErrorCodeProfileAccountVerified,
	"PROFILE_VERIFIED_CSI_REQUEST_PENDING":     ErrorCodeProfileVerifiedCsiRequestPending,
	"PROFILE_VERIFIED_CSI_REQUEST_NOT_FOUND":   ErrorCodeProfileVerifiedCsiRequestNotFound,
	"CREATE_PROFILE_ORGANIZATION_NAME_INVALID": ErrorCodeCreateProfileOrganizationNameInvalid,
	"CREATE_PROFILE_EMAIL_INVALID":             ErrorCodeCreateProfileEmailInvalid,
}

var mappingErrorCodeEnumLowerCase = map[string]ErrorCodeEnum{
	"content_empty":                            ErrorCodeContentEmpty,
	"client_exception":                         ErrorCodeClientException,
	"invalid_format":                           ErrorCodeInvalidFormat,
	"invalid_json_input":                       ErrorCodeInvalidJsonInput,
	"ssl_authorization":                        ErrorCodeSslAuthorization,
	"auth_failed":                              ErrorCodeAuthFailed,
	"authz_failed":                             ErrorCodeAuthzFailed,
	"csi_not_authorized":                       ErrorCodeCsiNotAuthorized,
	"user_policy_not_authorized":               ErrorCodeUserPolicyNotAuthorized,
	"email_not_verified":                       ErrorCodeEmailNotVerified,
	"email_not_found":                          ErrorCodeEmailNotFound,
	"oci_email_not_found":                      ErrorCodeOciEmailNotFound,
	"mos_email_not_found":                      ErrorCodeMosEmailNotFound,
	"idcs_email_not_valid":                     ErrorCodeIdcsEmailNotValid,
	"invalid_path":                             ErrorCodeInvalidPath,
	"method_not_allowed":                       ErrorCodeMethodNotAllowed,
	"json_processing":                          ErrorCodeJsonProcessing,
	"generic_exception":                        ErrorCodeGenericException,
	"external_service_provider_unavailable":    ErrorCodeExternalServiceProviderUnavailable,
	"external_service_provider_timeout":        ErrorCodeExternalServiceProviderTimeout,
	"too_many_requests":                        ErrorCodeTooManyRequests,
	"idp_scim_not_setup":                       ErrorCodeIdpScimNotSetup,
	"incident_not_found":                       ErrorCodeIncidentNotFound,
	"invalid_user_csi":                         ErrorCodeInvalidUserCsi,
	"data_already_exists":                      ErrorCodeDataAlreadyExists,
	"auth_user_not_matching":                   ErrorCodeAuthUserNotMatching,
	"contact_not_approved":                     ErrorCodeContactNotApproved,
	"create_profile_mos_failure":               ErrorCodeCreateProfileMosFailure,
	"create_profile_create_osso_failure":       ErrorCodeCreateProfileCreateOssoFailure,
	"create_profile_identity_failure":          ErrorCodeCreateProfileIdentityFailure,
	"create_profile_verify_osso_failure":       ErrorCodeCreateProfileVerifyOssoFailure,
	"profile_account_not_verified":             ErrorCodeProfileAccountNotVerified,
	"support_account_not_found":                ErrorCodeSupportAccountNotFound,
	"support_account_pending_csi_approval":     ErrorCodeSupportAccountPendingCsiApproval,
	"free_tier_customer_sli_unsupported":       ErrorCodeFreeTierCustomerSliUnsupported,
	"profile_account_verified":                 ErrorCodeProfileAccountVerified,
	"profile_verified_csi_request_pending":     ErrorCodeProfileVerifiedCsiRequestPending,
	"profile_verified_csi_request_not_found":   ErrorCodeProfileVerifiedCsiRequestNotFound,
	"create_profile_organization_name_invalid": ErrorCodeCreateProfileOrganizationNameInvalid,
	"create_profile_email_invalid":             ErrorCodeCreateProfileEmailInvalid,
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
		"AUTHZ_FAILED",
		"CSI_NOT_AUTHORIZED",
		"USER_POLICY_NOT_AUTHORIZED",
		"EMAIL_NOT_VERIFIED",
		"EMAIL_NOT_FOUND",
		"OCI_EMAIL_NOT_FOUND",
		"MOS_EMAIL_NOT_FOUND",
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
		"CONTACT_NOT_APPROVED",
		"CREATE_PROFILE_MOS_FAILURE",
		"CREATE_PROFILE_CREATE_OSSO_FAILURE",
		"CREATE_PROFILE_IDENTITY_FAILURE",
		"CREATE_PROFILE_VERIFY_OSSO_FAILURE",
		"PROFILE_ACCOUNT_NOT_VERIFIED",
		"SUPPORT_ACCOUNT_NOT_FOUND",
		"SUPPORT_ACCOUNT_PENDING_CSI_APPROVAL",
		"FREE_TIER_CUSTOMER_SLI_UNSUPPORTED",
		"PROFILE_ACCOUNT_VERIFIED",
		"PROFILE_VERIFIED_CSI_REQUEST_PENDING",
		"PROFILE_VERIFIED_CSI_REQUEST_NOT_FOUND",
		"CREATE_PROFILE_ORGANIZATION_NAME_INVALID",
		"CREATE_PROFILE_EMAIL_INVALID",
	}
}

// GetMappingErrorCodeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingErrorCodeEnum(val string) (ErrorCodeEnum, bool) {
	enum, ok := mappingErrorCodeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
