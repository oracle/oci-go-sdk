// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Organizations API
//
// The Organizations API allows you to consolidate multiple OCI tenancies into an organization, and centrally manage your tenancies and its resources.
//

package tenantmanagercontrolplane

// SenderInvitationStatusEnum Enum with underlying type: string
type SenderInvitationStatusEnum string

// Set of constants representing the allowable values for SenderInvitationStatusEnum
const (
	SenderInvitationStatusPending  SenderInvitationStatusEnum = "PENDING"
	SenderInvitationStatusCanceled SenderInvitationStatusEnum = "CANCELED"
	SenderInvitationStatusAccepted SenderInvitationStatusEnum = "ACCEPTED"
	SenderInvitationStatusExpired  SenderInvitationStatusEnum = "EXPIRED"
	SenderInvitationStatusFailed   SenderInvitationStatusEnum = "FAILED"
)

var mappingSenderInvitationStatusEnum = map[string]SenderInvitationStatusEnum{
	"PENDING":  SenderInvitationStatusPending,
	"CANCELED": SenderInvitationStatusCanceled,
	"ACCEPTED": SenderInvitationStatusAccepted,
	"EXPIRED":  SenderInvitationStatusExpired,
	"FAILED":   SenderInvitationStatusFailed,
}

// GetSenderInvitationStatusEnumValues Enumerates the set of values for SenderInvitationStatusEnum
func GetSenderInvitationStatusEnumValues() []SenderInvitationStatusEnum {
	values := make([]SenderInvitationStatusEnum, 0)
	for _, v := range mappingSenderInvitationStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetSenderInvitationStatusEnumStringValues Enumerates the set of values in String for SenderInvitationStatusEnum
func GetSenderInvitationStatusEnumStringValues() []string {
	return []string{
		"PENDING",
		"CANCELED",
		"ACCEPTED",
		"EXPIRED",
		"FAILED",
	}
}
