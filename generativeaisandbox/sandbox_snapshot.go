// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// GenAI Sandbox API
//
// API for managing sandbox resources and exchanging sandbox access tokens.
//

package generativeaisandbox

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// SandboxSnapshot Sandbox snapshot resource representation.
type SandboxSnapshot struct {

	// Sandbox snapshot identifier.
	Id *string `mandatory:"true" json:"id"`

	// The current state of the SandboxSnapshot.
	LifecycleState SandboxSnapshotLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// Display name for the sandbox snapshot.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Project identifier that the sandbox snapshot belongs to.
	ProjectId *string `mandatory:"false" json:"projectId"`

	// Identifier of the sandbox used to create the sandbox snapshot.
	SandboxId *string `mandatory:"false" json:"sandboxId"`

	// Runtime identifier for the sandbox snapshot.
	Runtime *string `mandatory:"false" json:"runtime"`

	// Duration after creation before the sandbox snapshot expires, in ISO 8601 duration format. For example, `P7D` for 7 days or `PT24H` for 24 hours.
	ExpirationDuration *string `mandatory:"false" json:"expirationDuration"`

	// Size of the sandbox snapshot archive in bytes.
	SizeBytes *int64 `mandatory:"false" json:"sizeBytes"`

	// The date and time the sandbox snapshot was created, in RFC 3339 format.
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// The date and time the sandbox snapshot was last updated, in RFC 3339 format.
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// The date and time the sandbox snapshot expires, in RFC 3339 format.
	TimeExpires *common.SDKTime `mandatory:"false" json:"timeExpires"`
}

func (m SandboxSnapshot) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m SandboxSnapshot) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingSandboxSnapshotLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetSandboxSnapshotLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// SandboxSnapshotLifecycleStateEnum Enum with underlying type: string
type SandboxSnapshotLifecycleStateEnum string

// Set of constants representing the allowable values for SandboxSnapshotLifecycleStateEnum
const (
	SandboxSnapshotLifecycleStateCreating SandboxSnapshotLifecycleStateEnum = "CREATING"
	SandboxSnapshotLifecycleStateActive   SandboxSnapshotLifecycleStateEnum = "ACTIVE"
	SandboxSnapshotLifecycleStateFailed   SandboxSnapshotLifecycleStateEnum = "FAILED"
	SandboxSnapshotLifecycleStateDeleting SandboxSnapshotLifecycleStateEnum = "DELETING"
	SandboxSnapshotLifecycleStateDeleted  SandboxSnapshotLifecycleStateEnum = "DELETED"
)

var mappingSandboxSnapshotLifecycleStateEnum = map[string]SandboxSnapshotLifecycleStateEnum{
	"CREATING": SandboxSnapshotLifecycleStateCreating,
	"ACTIVE":   SandboxSnapshotLifecycleStateActive,
	"FAILED":   SandboxSnapshotLifecycleStateFailed,
	"DELETING": SandboxSnapshotLifecycleStateDeleting,
	"DELETED":  SandboxSnapshotLifecycleStateDeleted,
}

var mappingSandboxSnapshotLifecycleStateEnumLowerCase = map[string]SandboxSnapshotLifecycleStateEnum{
	"creating": SandboxSnapshotLifecycleStateCreating,
	"active":   SandboxSnapshotLifecycleStateActive,
	"failed":   SandboxSnapshotLifecycleStateFailed,
	"deleting": SandboxSnapshotLifecycleStateDeleting,
	"deleted":  SandboxSnapshotLifecycleStateDeleted,
}

// GetSandboxSnapshotLifecycleStateEnumValues Enumerates the set of values for SandboxSnapshotLifecycleStateEnum
func GetSandboxSnapshotLifecycleStateEnumValues() []SandboxSnapshotLifecycleStateEnum {
	values := make([]SandboxSnapshotLifecycleStateEnum, 0)
	for _, v := range mappingSandboxSnapshotLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetSandboxSnapshotLifecycleStateEnumStringValues Enumerates the set of values in String for SandboxSnapshotLifecycleStateEnum
func GetSandboxSnapshotLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"ACTIVE",
		"FAILED",
		"DELETING",
		"DELETED",
	}
}

// GetMappingSandboxSnapshotLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSandboxSnapshotLifecycleStateEnum(val string) (SandboxSnapshotLifecycleStateEnum, bool) {
	enum, ok := mappingSandboxSnapshotLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
