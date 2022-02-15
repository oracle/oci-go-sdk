// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Organizations API
//
// The Organizations API allows you to consolidate multiple OCI tenancies into an organization, and centrally manage your tenancies and its resources.
//

package tenantmanagercontrolplane

import (
	"strings"
)

// OperationTypeEnum Enum with underlying type: string
type OperationTypeEnum string

// Set of constants representing the allowable values for OperationTypeEnum
const (
	OperationTypeCreateSenderInvitation       OperationTypeEnum = "CREATE_SENDER_INVITATION"
	OperationTypeAcceptRecipientInvitation    OperationTypeEnum = "ACCEPT_RECIPIENT_INVITATION"
	OperationTypeCancelSenderInvitation       OperationTypeEnum = "CANCEL_SENDER_INVITATION"
	OperationTypeCompleteOrderActivation      OperationTypeEnum = "COMPLETE_ORDER_ACTIVATION"
	OperationTypeActivateOrderExistingTenancy OperationTypeEnum = "ACTIVATE_ORDER_EXISTING_TENANCY"
	OperationTypeRegisterDomain               OperationTypeEnum = "REGISTER_DOMAIN"
	OperationTypeReleaseDomain                OperationTypeEnum = "RELEASE_DOMAIN"
	OperationTypeCreateChildTenancy           OperationTypeEnum = "CREATE_CHILD_TENANCY"
	OperationTypeAssignDefaultSubscription    OperationTypeEnum = "ASSIGN_DEFAULT_SUBSCRIPTION"
	OperationTypeManualLinkCreation           OperationTypeEnum = "MANUAL_LINK_CREATION"
	OperationTypeTerminateOrganizationTenancy OperationTypeEnum = "TERMINATE_ORGANIZATION_TENANCY"
)

var mappingOperationTypeEnum = map[string]OperationTypeEnum{
	"CREATE_SENDER_INVITATION":        OperationTypeCreateSenderInvitation,
	"ACCEPT_RECIPIENT_INVITATION":     OperationTypeAcceptRecipientInvitation,
	"CANCEL_SENDER_INVITATION":        OperationTypeCancelSenderInvitation,
	"COMPLETE_ORDER_ACTIVATION":       OperationTypeCompleteOrderActivation,
	"ACTIVATE_ORDER_EXISTING_TENANCY": OperationTypeActivateOrderExistingTenancy,
	"REGISTER_DOMAIN":                 OperationTypeRegisterDomain,
	"RELEASE_DOMAIN":                  OperationTypeReleaseDomain,
	"CREATE_CHILD_TENANCY":            OperationTypeCreateChildTenancy,
	"ASSIGN_DEFAULT_SUBSCRIPTION":     OperationTypeAssignDefaultSubscription,
	"MANUAL_LINK_CREATION":            OperationTypeManualLinkCreation,
	"TERMINATE_ORGANIZATION_TENANCY":  OperationTypeTerminateOrganizationTenancy,
}

// GetOperationTypeEnumValues Enumerates the set of values for OperationTypeEnum
func GetOperationTypeEnumValues() []OperationTypeEnum {
	values := make([]OperationTypeEnum, 0)
	for _, v := range mappingOperationTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetOperationTypeEnumStringValues Enumerates the set of values in String for OperationTypeEnum
func GetOperationTypeEnumStringValues() []string {
	return []string{
		"CREATE_SENDER_INVITATION",
		"ACCEPT_RECIPIENT_INVITATION",
		"CANCEL_SENDER_INVITATION",
		"COMPLETE_ORDER_ACTIVATION",
		"ACTIVATE_ORDER_EXISTING_TENANCY",
		"REGISTER_DOMAIN",
		"RELEASE_DOMAIN",
		"CREATE_CHILD_TENANCY",
		"ASSIGN_DEFAULT_SUBSCRIPTION",
		"MANUAL_LINK_CREATION",
		"TERMINATE_ORGANIZATION_TENANCY",
	}
}

// GetMappingOperationTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingOperationTypeEnum(val string) (OperationTypeEnum, bool) {
	mappingOperationTypeEnumIgnoreCase := make(map[string]OperationTypeEnum)
	for k, v := range mappingOperationTypeEnum {
		mappingOperationTypeEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingOperationTypeEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
