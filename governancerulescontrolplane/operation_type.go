// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// GovernanceRulesControlPlane API
//
// A description of the GovernanceRulesControlPlane API
//

package governancerulescontrolplane

import (
	"strings"
)

// OperationTypeEnum Enum with underlying type: string
type OperationTypeEnum string

// Set of constants representing the allowable values for OperationTypeEnum
const (
	OperationTypeCreateGovernanceRule              OperationTypeEnum = "CREATE_GOVERNANCE_RULE"
	OperationTypeUpdateGovernanceRule              OperationTypeEnum = "UPDATE_GOVERNANCE_RULE"
	OperationTypeDeleteGovernanceRule              OperationTypeEnum = "DELETE_GOVERNANCE_RULE"
	OperationTypeRetryGovernanceRule               OperationTypeEnum = "RETRY_GOVERNANCE_RULE"
	OperationTypeCreateInclusionCriteria           OperationTypeEnum = "CREATE_INCLUSION_CRITERIA"
	OperationTypeDeleteInclusionCriteria           OperationTypeEnum = "DELETE_INCLUSION_CRITERIA"
	OperationTypeRetryTenancyAttachment            OperationTypeEnum = "RETRY_TENANCY_ATTACHMENT"
	OperationTypeApplyTenancyAttachment            OperationTypeEnum = "APPLY_TENANCY_ATTACHMENT"
	OperationTypeCreateEnforcedQuotaGovernanceRule OperationTypeEnum = "CREATE_ENFORCED_QUOTA_GOVERNANCE_RULE"
	OperationTypeCreateEnforcedTagGovernanceRule   OperationTypeEnum = "CREATE_ENFORCED_TAG_GOVERNANCE_RULE"
	OperationTypeUpdateEnforcedQuotaGovernanceRule OperationTypeEnum = "UPDATE_ENFORCED_QUOTA_GOVERNANCE_RULE"
	OperationTypeUpdateEnforcedTagGovernanceRule   OperationTypeEnum = "UPDATE_ENFORCED_TAG_GOVERNANCE_RULE"
	OperationTypeDeleteEnforcedQuotaGovernanceRule OperationTypeEnum = "DELETE_ENFORCED_QUOTA_GOVERNANCE_RULE"
	OperationTypeDeleteEnforcedTagGovernanceRule   OperationTypeEnum = "DELETE_ENFORCED_TAG_GOVERNANCE_RULE"
	OperationTypeTenantManagerOptInEventHandler    OperationTypeEnum = "TENANT_MANAGER_OPT_IN_EVENT_HANDLER"
	OperationTypeTenantManagerOptOutEventHandler   OperationTypeEnum = "TENANT_MANAGER_OPT_OUT_EVENT_HANDLER"
)

var mappingOperationTypeEnum = map[string]OperationTypeEnum{
	"CREATE_GOVERNANCE_RULE":                OperationTypeCreateGovernanceRule,
	"UPDATE_GOVERNANCE_RULE":                OperationTypeUpdateGovernanceRule,
	"DELETE_GOVERNANCE_RULE":                OperationTypeDeleteGovernanceRule,
	"RETRY_GOVERNANCE_RULE":                 OperationTypeRetryGovernanceRule,
	"CREATE_INCLUSION_CRITERIA":             OperationTypeCreateInclusionCriteria,
	"DELETE_INCLUSION_CRITERIA":             OperationTypeDeleteInclusionCriteria,
	"RETRY_TENANCY_ATTACHMENT":              OperationTypeRetryTenancyAttachment,
	"APPLY_TENANCY_ATTACHMENT":              OperationTypeApplyTenancyAttachment,
	"CREATE_ENFORCED_QUOTA_GOVERNANCE_RULE": OperationTypeCreateEnforcedQuotaGovernanceRule,
	"CREATE_ENFORCED_TAG_GOVERNANCE_RULE":   OperationTypeCreateEnforcedTagGovernanceRule,
	"UPDATE_ENFORCED_QUOTA_GOVERNANCE_RULE": OperationTypeUpdateEnforcedQuotaGovernanceRule,
	"UPDATE_ENFORCED_TAG_GOVERNANCE_RULE":   OperationTypeUpdateEnforcedTagGovernanceRule,
	"DELETE_ENFORCED_QUOTA_GOVERNANCE_RULE": OperationTypeDeleteEnforcedQuotaGovernanceRule,
	"DELETE_ENFORCED_TAG_GOVERNANCE_RULE":   OperationTypeDeleteEnforcedTagGovernanceRule,
	"TENANT_MANAGER_OPT_IN_EVENT_HANDLER":   OperationTypeTenantManagerOptInEventHandler,
	"TENANT_MANAGER_OPT_OUT_EVENT_HANDLER":  OperationTypeTenantManagerOptOutEventHandler,
}

var mappingOperationTypeEnumLowerCase = map[string]OperationTypeEnum{
	"create_governance_rule":                OperationTypeCreateGovernanceRule,
	"update_governance_rule":                OperationTypeUpdateGovernanceRule,
	"delete_governance_rule":                OperationTypeDeleteGovernanceRule,
	"retry_governance_rule":                 OperationTypeRetryGovernanceRule,
	"create_inclusion_criteria":             OperationTypeCreateInclusionCriteria,
	"delete_inclusion_criteria":             OperationTypeDeleteInclusionCriteria,
	"retry_tenancy_attachment":              OperationTypeRetryTenancyAttachment,
	"apply_tenancy_attachment":              OperationTypeApplyTenancyAttachment,
	"create_enforced_quota_governance_rule": OperationTypeCreateEnforcedQuotaGovernanceRule,
	"create_enforced_tag_governance_rule":   OperationTypeCreateEnforcedTagGovernanceRule,
	"update_enforced_quota_governance_rule": OperationTypeUpdateEnforcedQuotaGovernanceRule,
	"update_enforced_tag_governance_rule":   OperationTypeUpdateEnforcedTagGovernanceRule,
	"delete_enforced_quota_governance_rule": OperationTypeDeleteEnforcedQuotaGovernanceRule,
	"delete_enforced_tag_governance_rule":   OperationTypeDeleteEnforcedTagGovernanceRule,
	"tenant_manager_opt_in_event_handler":   OperationTypeTenantManagerOptInEventHandler,
	"tenant_manager_opt_out_event_handler":  OperationTypeTenantManagerOptOutEventHandler,
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
		"CREATE_GOVERNANCE_RULE",
		"UPDATE_GOVERNANCE_RULE",
		"DELETE_GOVERNANCE_RULE",
		"RETRY_GOVERNANCE_RULE",
		"CREATE_INCLUSION_CRITERIA",
		"DELETE_INCLUSION_CRITERIA",
		"RETRY_TENANCY_ATTACHMENT",
		"APPLY_TENANCY_ATTACHMENT",
		"CREATE_ENFORCED_QUOTA_GOVERNANCE_RULE",
		"CREATE_ENFORCED_TAG_GOVERNANCE_RULE",
		"UPDATE_ENFORCED_QUOTA_GOVERNANCE_RULE",
		"UPDATE_ENFORCED_TAG_GOVERNANCE_RULE",
		"DELETE_ENFORCED_QUOTA_GOVERNANCE_RULE",
		"DELETE_ENFORCED_TAG_GOVERNANCE_RULE",
		"TENANT_MANAGER_OPT_IN_EVENT_HANDLER",
		"TENANT_MANAGER_OPT_OUT_EVENT_HANDLER",
	}
}

// GetMappingOperationTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingOperationTypeEnum(val string) (OperationTypeEnum, bool) {
	enum, ok := mappingOperationTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
