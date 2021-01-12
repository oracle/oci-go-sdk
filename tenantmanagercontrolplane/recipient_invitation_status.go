// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// TenantManager API
//
// A description of the TenantManager API
//

package tenantmanagercontrolplane

// RecipientInvitationStatusEnum Enum with underlying type: string
type RecipientInvitationStatusEnum string

// Set of constants representing the allowable values for RecipientInvitationStatusEnum
const (
	RecipientInvitationStatusPending  RecipientInvitationStatusEnum = "PENDING"
	RecipientInvitationStatusCanceled RecipientInvitationStatusEnum = "CANCELED"
	RecipientInvitationStatusAccepted RecipientInvitationStatusEnum = "ACCEPTED"
	RecipientInvitationStatusIgnored  RecipientInvitationStatusEnum = "IGNORED"
	RecipientInvitationStatusExpired  RecipientInvitationStatusEnum = "EXPIRED"
	RecipientInvitationStatusFailed   RecipientInvitationStatusEnum = "FAILED"
)

var mappingRecipientInvitationStatus = map[string]RecipientInvitationStatusEnum{
	"PENDING":  RecipientInvitationStatusPending,
	"CANCELED": RecipientInvitationStatusCanceled,
	"ACCEPTED": RecipientInvitationStatusAccepted,
	"IGNORED":  RecipientInvitationStatusIgnored,
	"EXPIRED":  RecipientInvitationStatusExpired,
	"FAILED":   RecipientInvitationStatusFailed,
}

// GetRecipientInvitationStatusEnumValues Enumerates the set of values for RecipientInvitationStatusEnum
func GetRecipientInvitationStatusEnumValues() []RecipientInvitationStatusEnum {
	values := make([]RecipientInvitationStatusEnum, 0)
	for _, v := range mappingRecipientInvitationStatus {
		values = append(values, v)
	}
	return values
}
