// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// GovernanceRulesControlPlane API
//
// A description of the GovernanceRulesControlPlane API
//

package governancerulescontrolplane

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// TenancyAttachment Tenancy attachment associates a tenancy to a governance rule via an inclusion criterion.
type TenancyAttachment struct {

	// The Oracle ID (OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm)) of the tenancy attachment.
	Id *string `mandatory:"true" json:"id"`

	// The Oracle ID (OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm)) of the root compartment containing the tenancy attachment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The Oracle ID (OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm)) of the governance rule. Every tenancy attachment is associated with a governance rule.
	GovernanceRuleId *string `mandatory:"true" json:"governanceRuleId"`

	// The Oracle ID (OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm)) of the tenancy to which the governance rule is attached.
	TenancyId *string `mandatory:"true" json:"tenancyId"`

	// The current state of the tenancy attachment.
	LifecycleState TenancyAttachmentLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// Date and time the tenancy attachment was created. An RFC3339 formatted datetime string.
	// Example: `2016-08-25T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// Date and time the tenancy attachment was updated. An RFC3339 formatted datetime string.
	// Example: `2016-08-25T21:10:29.600Z`
	TimeUpdated *common.SDKTime `mandatory:"true" json:"timeUpdated"`

	// Date and time the tenancy attachment was last attempted. An RFC3339 formatted datetime string.
	// Example: `2016-08-25T21:10:29.600Z`
	TimeLastAttempted *common.SDKTime `mandatory:"false" json:"timeLastAttempted"`
}

func (m TenancyAttachment) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m TenancyAttachment) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingTenancyAttachmentLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetTenancyAttachmentLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// TenancyAttachmentLifecycleStateEnum Enum with underlying type: string
type TenancyAttachmentLifecycleStateEnum string

// Set of constants representing the allowable values for TenancyAttachmentLifecycleStateEnum
const (
	TenancyAttachmentLifecycleStateCreating       TenancyAttachmentLifecycleStateEnum = "CREATING"
	TenancyAttachmentLifecycleStateActive         TenancyAttachmentLifecycleStateEnum = "ACTIVE"
	TenancyAttachmentLifecycleStateUpdating       TenancyAttachmentLifecycleStateEnum = "UPDATING"
	TenancyAttachmentLifecycleStateNeedsAttention TenancyAttachmentLifecycleStateEnum = "NEEDS_ATTENTION"
	TenancyAttachmentLifecycleStateDeleting       TenancyAttachmentLifecycleStateEnum = "DELETING"
	TenancyAttachmentLifecycleStateDeleted        TenancyAttachmentLifecycleStateEnum = "DELETED"
)

var mappingTenancyAttachmentLifecycleStateEnum = map[string]TenancyAttachmentLifecycleStateEnum{
	"CREATING":        TenancyAttachmentLifecycleStateCreating,
	"ACTIVE":          TenancyAttachmentLifecycleStateActive,
	"UPDATING":        TenancyAttachmentLifecycleStateUpdating,
	"NEEDS_ATTENTION": TenancyAttachmentLifecycleStateNeedsAttention,
	"DELETING":        TenancyAttachmentLifecycleStateDeleting,
	"DELETED":         TenancyAttachmentLifecycleStateDeleted,
}

var mappingTenancyAttachmentLifecycleStateEnumLowerCase = map[string]TenancyAttachmentLifecycleStateEnum{
	"creating":        TenancyAttachmentLifecycleStateCreating,
	"active":          TenancyAttachmentLifecycleStateActive,
	"updating":        TenancyAttachmentLifecycleStateUpdating,
	"needs_attention": TenancyAttachmentLifecycleStateNeedsAttention,
	"deleting":        TenancyAttachmentLifecycleStateDeleting,
	"deleted":         TenancyAttachmentLifecycleStateDeleted,
}

// GetTenancyAttachmentLifecycleStateEnumValues Enumerates the set of values for TenancyAttachmentLifecycleStateEnum
func GetTenancyAttachmentLifecycleStateEnumValues() []TenancyAttachmentLifecycleStateEnum {
	values := make([]TenancyAttachmentLifecycleStateEnum, 0)
	for _, v := range mappingTenancyAttachmentLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetTenancyAttachmentLifecycleStateEnumStringValues Enumerates the set of values in String for TenancyAttachmentLifecycleStateEnum
func GetTenancyAttachmentLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"ACTIVE",
		"UPDATING",
		"NEEDS_ATTENTION",
		"DELETING",
		"DELETED",
	}
}

// GetMappingTenancyAttachmentLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingTenancyAttachmentLifecycleStateEnum(val string) (TenancyAttachmentLifecycleStateEnum, bool) {
	enum, ok := mappingTenancyAttachmentLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
