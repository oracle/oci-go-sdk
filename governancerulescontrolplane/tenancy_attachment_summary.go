// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
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

// TenancyAttachmentSummary Summary of the tenancy attachment.
type TenancyAttachmentSummary struct {

	// The Oracle ID (OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm)) of the tenancy attachment.
	Id *string `mandatory:"true" json:"id"`

	// The Oracle ID (OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm)) of the root compartment containing the tenancy attachment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The Oracle ID (OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm)) of the governance rule. Every tenancy attachment is associated with a governance rule.
	GovernanceRuleId *string `mandatory:"true" json:"governanceRuleId"`

	// The Oracle ID (OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm)) of the tenancy to which the governance rule is attached.
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

func (m TenancyAttachmentSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m TenancyAttachmentSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingTenancyAttachmentLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetTenancyAttachmentLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
