// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Managed Access API
//
// Use the Managed Access API to approve access requests, create and manage templates, and manage resource approval settings. For more information, see Managed Access Overview (https://docs.oracle.com/iaas/Content/managed-access/home.htm).
// Use the table of contents and search tool to explore the Managed Access API.
//

package lockbox

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ApprovalTemplate Group/User OCIDs of those who can approve/deny/revoke operator's request to access associated resources.
type ApprovalTemplate struct {

	// The unique identifier (OCID) of the approval template, which can't be changed after creation.
	Id *string `mandatory:"true" json:"id"`

	// The approval template display name.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The unique identifier (OCID) of the customer compartment where the approval template is located.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The time the the approval template was created. An RFC3339 formatted datetime string
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The current state of the approval template.
	LifecycleState ApprovalTemplateLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	ApproverLevels *ApproverLevels `mandatory:"false" json:"approverLevels"`

	// The auto approval state of the lockbox.
	AutoApprovalState LockboxAutoApprovalStateEnum `mandatory:"false" json:"autoApprovalState,omitempty"`

	// The time the approval template was updated. An RFC3339 formatted datetime string
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// Usage of system tag keys. These predefined keys are scoped to namespaces.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
}

func (m ApprovalTemplate) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ApprovalTemplate) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingApprovalTemplateLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetApprovalTemplateLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingLockboxAutoApprovalStateEnum(string(m.AutoApprovalState)); !ok && m.AutoApprovalState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for AutoApprovalState: %s. Supported values are: %s.", m.AutoApprovalState, strings.Join(GetLockboxAutoApprovalStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ApprovalTemplateLifecycleStateEnum Enum with underlying type: string
type ApprovalTemplateLifecycleStateEnum string

// Set of constants representing the allowable values for ApprovalTemplateLifecycleStateEnum
const (
	ApprovalTemplateLifecycleStateActive   ApprovalTemplateLifecycleStateEnum = "ACTIVE"
	ApprovalTemplateLifecycleStateCreating ApprovalTemplateLifecycleStateEnum = "CREATING"
	ApprovalTemplateLifecycleStateUpdating ApprovalTemplateLifecycleStateEnum = "UPDATING"
	ApprovalTemplateLifecycleStateDeleting ApprovalTemplateLifecycleStateEnum = "DELETING"
	ApprovalTemplateLifecycleStateDeleted  ApprovalTemplateLifecycleStateEnum = "DELETED"
	ApprovalTemplateLifecycleStateFailed   ApprovalTemplateLifecycleStateEnum = "FAILED"
)

var mappingApprovalTemplateLifecycleStateEnum = map[string]ApprovalTemplateLifecycleStateEnum{
	"ACTIVE":   ApprovalTemplateLifecycleStateActive,
	"CREATING": ApprovalTemplateLifecycleStateCreating,
	"UPDATING": ApprovalTemplateLifecycleStateUpdating,
	"DELETING": ApprovalTemplateLifecycleStateDeleting,
	"DELETED":  ApprovalTemplateLifecycleStateDeleted,
	"FAILED":   ApprovalTemplateLifecycleStateFailed,
}

var mappingApprovalTemplateLifecycleStateEnumLowerCase = map[string]ApprovalTemplateLifecycleStateEnum{
	"active":   ApprovalTemplateLifecycleStateActive,
	"creating": ApprovalTemplateLifecycleStateCreating,
	"updating": ApprovalTemplateLifecycleStateUpdating,
	"deleting": ApprovalTemplateLifecycleStateDeleting,
	"deleted":  ApprovalTemplateLifecycleStateDeleted,
	"failed":   ApprovalTemplateLifecycleStateFailed,
}

// GetApprovalTemplateLifecycleStateEnumValues Enumerates the set of values for ApprovalTemplateLifecycleStateEnum
func GetApprovalTemplateLifecycleStateEnumValues() []ApprovalTemplateLifecycleStateEnum {
	values := make([]ApprovalTemplateLifecycleStateEnum, 0)
	for _, v := range mappingApprovalTemplateLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetApprovalTemplateLifecycleStateEnumStringValues Enumerates the set of values in String for ApprovalTemplateLifecycleStateEnum
func GetApprovalTemplateLifecycleStateEnumStringValues() []string {
	return []string{
		"ACTIVE",
		"CREATING",
		"UPDATING",
		"DELETING",
		"DELETED",
		"FAILED",
	}
}

// GetMappingApprovalTemplateLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingApprovalTemplateLifecycleStateEnum(val string) (ApprovalTemplateLifecycleStateEnum, bool) {
	enum, ok := mappingApprovalTemplateLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
