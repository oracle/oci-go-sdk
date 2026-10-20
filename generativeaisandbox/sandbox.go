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

// Sandbox Sandbox resource representation.
type Sandbox struct {

	// Sandbox identifier.
	Id *string `mandatory:"true" json:"id"`

	// The current state of the Sandbox.
	LifecycleState SandboxLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The sandbox runtime type. For customer-provided containers, use a same-region Oracle Cloud Infrastructure Registry image reference.
	Runtime *string `mandatory:"true" json:"runtime"`

	// Display name for sandbox
	DisplayName *string `mandatory:"false" json:"displayName"`

	// GenerativeAiProject ID that sandbox belongs to
	ProjectId *string `mandatory:"false" json:"projectId"`

	// Name of the source warm pool, when selected explicitly or automatically at creation. The persisted association is retained on retries.
	WarmpoolName *string `mandatory:"false" json:"warmpoolName"`

	// Persisted published-port authorization mode. Absent for legacy sandboxes. Immutable after creation.
	PublishedPortAuthMode SandboxPublishedPortAuthModeEnum `mandatory:"false" json:"publishedPortAuthMode,omitempty"`

	// Published ports for the sandbox proxy endpoint.
	PublishedPorts []PublishedPort `mandatory:"false" json:"publishedPorts"`

	// list of environments for sandbox
	EnvironmentVariables []EnvironmentVariable `mandatory:"false" json:"environmentVariables"`

	// Number of minutes from creation until automatic sandbox expiration. Use `-1` to disable automatic expiration.
	ExpiresInMinutes *int `mandatory:"false" json:"expiresInMinutes"`

	// Deprecated. Use `expiresInMinutes` instead. Existing clients may pass `-1` to disable automatic expiration.
	ExpirationDuration *string `mandatory:"false" json:"expirationDuration"`

	// shape for sandbox
	Shape SandboxShapeEnum `mandatory:"false" json:"shape,omitempty"`

	// The date and time the sandbox was created, in RFC 3339 format.
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// The date and time the sandbox was updated, in RFC 3339 format.
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// The date and time the sandbox expires, in RFC 3339 format.
	TimeExpires *common.SDKTime `mandatory:"false" json:"timeExpires"`

	// The date and time the sandbox was stopped or expired, in RFC 3339 format.
	TimeStopped *common.SDKTime `mandatory:"false" json:"timeStopped"`
}

func (m Sandbox) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m Sandbox) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingSandboxLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetSandboxLifecycleStateEnumStringValues(), ",")))
	}

	if _, ok := GetMappingSandboxPublishedPortAuthModeEnum(string(m.PublishedPortAuthMode)); !ok && m.PublishedPortAuthMode != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for PublishedPortAuthMode: %s. Supported values are: %s.", m.PublishedPortAuthMode, strings.Join(GetSandboxPublishedPortAuthModeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingSandboxShapeEnum(string(m.Shape)); !ok && m.Shape != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Shape: %s. Supported values are: %s.", m.Shape, strings.Join(GetSandboxShapeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// SandboxLifecycleStateEnum Enum with underlying type: string
type SandboxLifecycleStateEnum string

// Set of constants representing the allowable values for SandboxLifecycleStateEnum
const (
	SandboxLifecycleStateRunning    SandboxLifecycleStateEnum = "RUNNING"
	SandboxLifecycleStateCreating   SandboxLifecycleStateEnum = "CREATING"
	SandboxLifecycleStateUpdating   SandboxLifecycleStateEnum = "UPDATING"
	SandboxLifecycleStateDeleted    SandboxLifecycleStateEnum = "DELETED"
	SandboxLifecycleStateFailed     SandboxLifecycleStateEnum = "FAILED"
	SandboxLifecycleStateSuspending SandboxLifecycleStateEnum = "SUSPENDING"
	SandboxLifecycleStateSuspended  SandboxLifecycleStateEnum = "SUSPENDED"
	SandboxLifecycleStateResuming   SandboxLifecycleStateEnum = "RESUMING"
	SandboxLifecycleStateStopping   SandboxLifecycleStateEnum = "STOPPING"
	SandboxLifecycleStateStopped    SandboxLifecycleStateEnum = "STOPPED"
)

var mappingSandboxLifecycleStateEnum = map[string]SandboxLifecycleStateEnum{
	"RUNNING":    SandboxLifecycleStateRunning,
	"CREATING":   SandboxLifecycleStateCreating,
	"UPDATING":   SandboxLifecycleStateUpdating,
	"DELETED":    SandboxLifecycleStateDeleted,
	"FAILED":     SandboxLifecycleStateFailed,
	"SUSPENDING": SandboxLifecycleStateSuspending,
	"SUSPENDED":  SandboxLifecycleStateSuspended,
	"RESUMING":   SandboxLifecycleStateResuming,
	"STOPPING":   SandboxLifecycleStateStopping,
	"STOPPED":    SandboxLifecycleStateStopped,
}

var mappingSandboxLifecycleStateEnumLowerCase = map[string]SandboxLifecycleStateEnum{
	"running":    SandboxLifecycleStateRunning,
	"creating":   SandboxLifecycleStateCreating,
	"updating":   SandboxLifecycleStateUpdating,
	"deleted":    SandboxLifecycleStateDeleted,
	"failed":     SandboxLifecycleStateFailed,
	"suspending": SandboxLifecycleStateSuspending,
	"suspended":  SandboxLifecycleStateSuspended,
	"resuming":   SandboxLifecycleStateResuming,
	"stopping":   SandboxLifecycleStateStopping,
	"stopped":    SandboxLifecycleStateStopped,
}

// GetSandboxLifecycleStateEnumValues Enumerates the set of values for SandboxLifecycleStateEnum
func GetSandboxLifecycleStateEnumValues() []SandboxLifecycleStateEnum {
	values := make([]SandboxLifecycleStateEnum, 0)
	for _, v := range mappingSandboxLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetSandboxLifecycleStateEnumStringValues Enumerates the set of values in String for SandboxLifecycleStateEnum
func GetSandboxLifecycleStateEnumStringValues() []string {
	return []string{
		"RUNNING",
		"CREATING",
		"UPDATING",
		"DELETED",
		"FAILED",
		"SUSPENDING",
		"SUSPENDED",
		"RESUMING",
		"STOPPING",
		"STOPPED",
	}
}

// GetMappingSandboxLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSandboxLifecycleStateEnum(val string) (SandboxLifecycleStateEnum, bool) {
	enum, ok := mappingSandboxLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// SandboxPublishedPortAuthModeEnum Enum with underlying type: string
type SandboxPublishedPortAuthModeEnum string

// Set of constants representing the allowable values for SandboxPublishedPortAuthModeEnum
const (
	SandboxPublishedPortAuthModeOciIam SandboxPublishedPortAuthModeEnum = "OCI_IAM"
)

var mappingSandboxPublishedPortAuthModeEnum = map[string]SandboxPublishedPortAuthModeEnum{
	"OCI_IAM": SandboxPublishedPortAuthModeOciIam,
}

var mappingSandboxPublishedPortAuthModeEnumLowerCase = map[string]SandboxPublishedPortAuthModeEnum{
	"oci_iam": SandboxPublishedPortAuthModeOciIam,
}

// GetSandboxPublishedPortAuthModeEnumValues Enumerates the set of values for SandboxPublishedPortAuthModeEnum
func GetSandboxPublishedPortAuthModeEnumValues() []SandboxPublishedPortAuthModeEnum {
	values := make([]SandboxPublishedPortAuthModeEnum, 0)
	for _, v := range mappingSandboxPublishedPortAuthModeEnum {
		values = append(values, v)
	}
	return values
}

// GetSandboxPublishedPortAuthModeEnumStringValues Enumerates the set of values in String for SandboxPublishedPortAuthModeEnum
func GetSandboxPublishedPortAuthModeEnumStringValues() []string {
	return []string{
		"OCI_IAM",
	}
}

// GetMappingSandboxPublishedPortAuthModeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSandboxPublishedPortAuthModeEnum(val string) (SandboxPublishedPortAuthModeEnum, bool) {
	enum, ok := mappingSandboxPublishedPortAuthModeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// SandboxShapeEnum Enum with underlying type: string
type SandboxShapeEnum string

// Set of constants representing the allowable values for SandboxShapeEnum
const (
	SandboxShapeSmall  SandboxShapeEnum = "SMALL"
	SandboxShapeMedium SandboxShapeEnum = "MEDIUM"
	SandboxShapeLarge  SandboxShapeEnum = "LARGE"
)

var mappingSandboxShapeEnum = map[string]SandboxShapeEnum{
	"SMALL":  SandboxShapeSmall,
	"MEDIUM": SandboxShapeMedium,
	"LARGE":  SandboxShapeLarge,
}

var mappingSandboxShapeEnumLowerCase = map[string]SandboxShapeEnum{
	"small":  SandboxShapeSmall,
	"medium": SandboxShapeMedium,
	"large":  SandboxShapeLarge,
}

// GetSandboxShapeEnumValues Enumerates the set of values for SandboxShapeEnum
func GetSandboxShapeEnumValues() []SandboxShapeEnum {
	values := make([]SandboxShapeEnum, 0)
	for _, v := range mappingSandboxShapeEnum {
		values = append(values, v)
	}
	return values
}

// GetSandboxShapeEnumStringValues Enumerates the set of values in String for SandboxShapeEnum
func GetSandboxShapeEnumStringValues() []string {
	return []string{
		"SMALL",
		"MEDIUM",
		"LARGE",
	}
}

// GetMappingSandboxShapeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSandboxShapeEnum(val string) (SandboxShapeEnum, bool) {
	enum, ok := mappingSandboxShapeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
