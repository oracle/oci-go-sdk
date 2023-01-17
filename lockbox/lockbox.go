// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Managed Access API
//
// Use the Managed Access API to approve access requests, create and manage templates, and manage resource approval settings. For more information, see Managed Access Overview (https://docs.oracle.com/en-us/iaas/managed-access/overview.htm).
// Use the table of contents and search tool to explore the Managed Access API.
//

package lockbox

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// Lockbox Description of Lockbox.
type Lockbox struct {

	// Unique identifier that is immutable on creation
	Id *string `mandatory:"true" json:"id"`

	// Lockbox Identifier, can be renamed
	DisplayName *string `mandatory:"true" json:"displayName"`

	// Compartment Identifier
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// Compartment Identifier
	PartnerCompartmentId *string `mandatory:"true" json:"partnerCompartmentId"`

	// The unique identifier (OCID) of associated resource that the lockbox is created for.
	ResourceId *string `mandatory:"true" json:"resourceId"`

	// The partner using this lockbox to lock a resource.
	LockboxPartner LockboxPartnerEnum `mandatory:"true" json:"lockboxPartner"`

	// The time the the Lockbox was created. An RFC3339 formatted datetime string
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The current state of the Lockbox.
	LifecycleState LockboxLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"true" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"true" json:"definedTags"`

	// The time the Lockbox was updated. An RFC3339 formatted datetime string
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	AccessContextAttributes *AccessContextAttributeCollection `mandatory:"false" json:"accessContextAttributes"`

	// Approval template ID
	ApprovalTemplateId *string `mandatory:"false" json:"approvalTemplateId"`

	// The maximum amount of time operator has access to associated resources.
	MaxAccessDuration *string `mandatory:"false" json:"maxAccessDuration"`

	// A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// Usage of system tag keys. These predefined keys are scoped to namespaces.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
}

func (m Lockbox) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m Lockbox) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingLockboxPartnerEnum(string(m.LockboxPartner)); !ok && m.LockboxPartner != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LockboxPartner: %s. Supported values are: %s.", m.LockboxPartner, strings.Join(GetLockboxPartnerEnumStringValues(), ",")))
	}
	if _, ok := GetMappingLockboxLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetLockboxLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// LockboxLifecycleStateEnum Enum with underlying type: string
type LockboxLifecycleStateEnum string

// Set of constants representing the allowable values for LockboxLifecycleStateEnum
const (
	LockboxLifecycleStateActive   LockboxLifecycleStateEnum = "ACTIVE"
	LockboxLifecycleStateCreating LockboxLifecycleStateEnum = "CREATING"
	LockboxLifecycleStateUpdating LockboxLifecycleStateEnum = "UPDATING"
	LockboxLifecycleStateDeleting LockboxLifecycleStateEnum = "DELETING"
	LockboxLifecycleStateDeleted  LockboxLifecycleStateEnum = "DELETED"
	LockboxLifecycleStateFailed   LockboxLifecycleStateEnum = "FAILED"
)

var mappingLockboxLifecycleStateEnum = map[string]LockboxLifecycleStateEnum{
	"ACTIVE":   LockboxLifecycleStateActive,
	"CREATING": LockboxLifecycleStateCreating,
	"UPDATING": LockboxLifecycleStateUpdating,
	"DELETING": LockboxLifecycleStateDeleting,
	"DELETED":  LockboxLifecycleStateDeleted,
	"FAILED":   LockboxLifecycleStateFailed,
}

var mappingLockboxLifecycleStateEnumLowerCase = map[string]LockboxLifecycleStateEnum{
	"active":   LockboxLifecycleStateActive,
	"creating": LockboxLifecycleStateCreating,
	"updating": LockboxLifecycleStateUpdating,
	"deleting": LockboxLifecycleStateDeleting,
	"deleted":  LockboxLifecycleStateDeleted,
	"failed":   LockboxLifecycleStateFailed,
}

// GetLockboxLifecycleStateEnumValues Enumerates the set of values for LockboxLifecycleStateEnum
func GetLockboxLifecycleStateEnumValues() []LockboxLifecycleStateEnum {
	values := make([]LockboxLifecycleStateEnum, 0)
	for _, v := range mappingLockboxLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetLockboxLifecycleStateEnumStringValues Enumerates the set of values in String for LockboxLifecycleStateEnum
func GetLockboxLifecycleStateEnumStringValues() []string {
	return []string{
		"ACTIVE",
		"CREATING",
		"UPDATING",
		"DELETING",
		"DELETED",
		"FAILED",
	}
}

// GetMappingLockboxLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingLockboxLifecycleStateEnum(val string) (LockboxLifecycleStateEnum, bool) {
	enum, ok := mappingLockboxLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
