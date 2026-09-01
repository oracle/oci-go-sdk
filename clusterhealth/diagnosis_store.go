// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// ITV Control Plane - Diagnosis Store API
//
// Use the ITV Control Plane Diagnosis Store API to manage diagnosis.
//

package clusterhealth

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// DiagnosisStore A DiagnosisStore is a description of a DiagnosisStore.
// To use any of the API operations, you must be authorized in an IAM policy. If you're not authorized, talk to
// an administrator. If you're an administrator who needs to write policies to give users access, see
// Getting Started with Policies (https://docs.oracle.com/iaas/Content/Identity/policiesgs/get-started-with-policies.htm).
type DiagnosisStore struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the DiagnosisStore.
	Id *string `mandatory:"true" json:"id"`

	// A user-friendly name. Does not have to be unique, and it's changeable.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The current state of the Diagnosis.
	LifecycleState DiagnosisStoreLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// A message that describes the current state of the Diagnosis in more detail. For example,
	// can be used to provide actionable information for a resource in the Failed state.
	LifecycleDetails *string `mandatory:"true" json:"lifecycleDetails"`

	// The namespace of the object store.
	ObjectStoreNamespace *string `mandatory:"true" json:"objectStoreNamespace"`

	// The name of the object store bucket.
	ObjectStoreBucket *string `mandatory:"true" json:"objectStoreBucket"`

	// The date and time the Diagnosis Store was created, in the format defined by RFC 3339 (https://tools.ietf
	// .org/html/rfc3339).
	// Example: `2016-08-25T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The date and time the Diagnosis Store was updated, in the format defined by RFC 3339 (https://tools.ietf
	// .org/html/rfc3339).
	// Example: `2016-08-25T21:10:29.600Z`
	TimeUpdated *common.SDKTime `mandatory:"true" json:"timeUpdated"`

	// The last accepted request token.
	LastAcceptedRequestToken *string `mandatory:"true" json:"lastAcceptedRequestToken"`

	// The last completed request token.
	LastCompletedRequestToken *string `mandatory:"true" json:"lastCompletedRequestToken"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"true" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"true" json:"definedTags"`

	// System tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"true" json:"systemTags"`
}

func (m DiagnosisStore) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DiagnosisStore) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingDiagnosisStoreLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetDiagnosisStoreLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// DiagnosisStoreLifecycleStateEnum Enum with underlying type: string
type DiagnosisStoreLifecycleStateEnum string

// Set of constants representing the allowable values for DiagnosisStoreLifecycleStateEnum
const (
	DiagnosisStoreLifecycleStateCreating DiagnosisStoreLifecycleStateEnum = "CREATING"
	DiagnosisStoreLifecycleStateUpdating DiagnosisStoreLifecycleStateEnum = "UPDATING"
	DiagnosisStoreLifecycleStateActive   DiagnosisStoreLifecycleStateEnum = "ACTIVE"
	DiagnosisStoreLifecycleStateDeleting DiagnosisStoreLifecycleStateEnum = "DELETING"
	DiagnosisStoreLifecycleStateDeleted  DiagnosisStoreLifecycleStateEnum = "DELETED"
	DiagnosisStoreLifecycleStateFailed   DiagnosisStoreLifecycleStateEnum = "FAILED"
)

var mappingDiagnosisStoreLifecycleStateEnum = map[string]DiagnosisStoreLifecycleStateEnum{
	"CREATING": DiagnosisStoreLifecycleStateCreating,
	"UPDATING": DiagnosisStoreLifecycleStateUpdating,
	"ACTIVE":   DiagnosisStoreLifecycleStateActive,
	"DELETING": DiagnosisStoreLifecycleStateDeleting,
	"DELETED":  DiagnosisStoreLifecycleStateDeleted,
	"FAILED":   DiagnosisStoreLifecycleStateFailed,
}

var mappingDiagnosisStoreLifecycleStateEnumLowerCase = map[string]DiagnosisStoreLifecycleStateEnum{
	"creating": DiagnosisStoreLifecycleStateCreating,
	"updating": DiagnosisStoreLifecycleStateUpdating,
	"active":   DiagnosisStoreLifecycleStateActive,
	"deleting": DiagnosisStoreLifecycleStateDeleting,
	"deleted":  DiagnosisStoreLifecycleStateDeleted,
	"failed":   DiagnosisStoreLifecycleStateFailed,
}

// GetDiagnosisStoreLifecycleStateEnumValues Enumerates the set of values for DiagnosisStoreLifecycleStateEnum
func GetDiagnosisStoreLifecycleStateEnumValues() []DiagnosisStoreLifecycleStateEnum {
	values := make([]DiagnosisStoreLifecycleStateEnum, 0)
	for _, v := range mappingDiagnosisStoreLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetDiagnosisStoreLifecycleStateEnumStringValues Enumerates the set of values in String for DiagnosisStoreLifecycleStateEnum
func GetDiagnosisStoreLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"UPDATING",
		"ACTIVE",
		"DELETING",
		"DELETED",
		"FAILED",
	}
}

// GetMappingDiagnosisStoreLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDiagnosisStoreLifecycleStateEnum(val string) (DiagnosisStoreLifecycleStateEnum, bool) {
	enum, ok := mappingDiagnosisStoreLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
