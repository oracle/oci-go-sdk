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

// SandboxSummary Sandbox resource representation.
type SandboxSummary struct {

	// Sandbox identifier.
	Id *string `mandatory:"true" json:"id"`

	// The current state of the Sandbox.
	// Allowed values are:
	//   - RUNNING
	//   - CREATING
	//   - UPDATING
	//   - DELETED
	//   - FAILED
	//   - SUSPENDING
	//   - SUSPENDED
	//   - RESUMING
	//   - STOPPING
	//   - STOPPED
	LifecycleState SandboxLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The sandbox runtime type. For customer-provided containers, use a same-region Oracle Cloud Infrastructure Registry image reference.
	Runtime *string `mandatory:"true" json:"runtime"`

	// Display name for sandbox
	DisplayName *string `mandatory:"false" json:"displayName"`

	// GenerativeAiProject ID that sandbox belongs to
	ProjectId *string `mandatory:"false" json:"projectId"`

	// Name of the source warm pool, when selected explicitly or automatically at creation. The persisted association is retained on retries.
	WarmpoolName *string `mandatory:"false" json:"warmpoolName"`

	// Persisted published-port authorization mode. Absent for legacy sandboxes.
	PublishedPortAuthMode SandboxPublishedPortAuthModeEnum `mandatory:"false" json:"publishedPortAuthMode,omitempty"`

	// Published ports for the sandbox proxy endpoint.
	PublishedPorts []PublishedPort `mandatory:"false" json:"publishedPorts"`

	// Number of minutes from creation until automatic sandbox expiration. Use `-1` to disable automatic expiration.
	ExpiresInMinutes *int `mandatory:"false" json:"expiresInMinutes"`

	// Deprecated. Use `expiresInMinutes` instead. Existing clients may pass `-1` to disable automatic expiration.
	ExpirationDuration *string `mandatory:"false" json:"expirationDuration"`

	// shape for sandbox
	Shape SandboxSummaryShapeEnum `mandatory:"false" json:"shape,omitempty"`

	// The date and time the sandbox was created, in RFC 3339 format.
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// The date and time the sandbox was updated, in RFC 3339 format.
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// The date and time the sandbox expires, in RFC 3339 format.
	TimeExpires *common.SDKTime `mandatory:"false" json:"timeExpires"`

	// The date and time the sandbox was stopped or expired, in RFC 3339 format.
	TimeStopped *common.SDKTime `mandatory:"false" json:"timeStopped"`
}

func (m SandboxSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m SandboxSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingSandboxLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetSandboxLifecycleStateEnumStringValues(), ",")))
	}

	if _, ok := GetMappingSandboxPublishedPortAuthModeEnum(string(m.PublishedPortAuthMode)); !ok && m.PublishedPortAuthMode != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for PublishedPortAuthMode: %s. Supported values are: %s.", m.PublishedPortAuthMode, strings.Join(GetSandboxPublishedPortAuthModeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingSandboxSummaryShapeEnum(string(m.Shape)); !ok && m.Shape != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Shape: %s. Supported values are: %s.", m.Shape, strings.Join(GetSandboxSummaryShapeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// SandboxSummaryShapeEnum Enum with underlying type: string
type SandboxSummaryShapeEnum string

// Set of constants representing the allowable values for SandboxSummaryShapeEnum
const (
	SandboxSummaryShapeSmall  SandboxSummaryShapeEnum = "SMALL"
	SandboxSummaryShapeMedium SandboxSummaryShapeEnum = "MEDIUM"
	SandboxSummaryShapeLarge  SandboxSummaryShapeEnum = "LARGE"
)

var mappingSandboxSummaryShapeEnum = map[string]SandboxSummaryShapeEnum{
	"SMALL":  SandboxSummaryShapeSmall,
	"MEDIUM": SandboxSummaryShapeMedium,
	"LARGE":  SandboxSummaryShapeLarge,
}

var mappingSandboxSummaryShapeEnumLowerCase = map[string]SandboxSummaryShapeEnum{
	"small":  SandboxSummaryShapeSmall,
	"medium": SandboxSummaryShapeMedium,
	"large":  SandboxSummaryShapeLarge,
}

// GetSandboxSummaryShapeEnumValues Enumerates the set of values for SandboxSummaryShapeEnum
func GetSandboxSummaryShapeEnumValues() []SandboxSummaryShapeEnum {
	values := make([]SandboxSummaryShapeEnum, 0)
	for _, v := range mappingSandboxSummaryShapeEnum {
		values = append(values, v)
	}
	return values
}

// GetSandboxSummaryShapeEnumStringValues Enumerates the set of values in String for SandboxSummaryShapeEnum
func GetSandboxSummaryShapeEnumStringValues() []string {
	return []string{
		"SMALL",
		"MEDIUM",
		"LARGE",
	}
}

// GetMappingSandboxSummaryShapeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSandboxSummaryShapeEnum(val string) (SandboxSummaryShapeEnum, bool) {
	enum, ok := mappingSandboxSummaryShapeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
