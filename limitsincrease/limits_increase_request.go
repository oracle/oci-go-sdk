// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Limits Increase API
//
// Use the Limits Increase API to work with limit increase requests.
// For more information, see
// Working with Limit Increase Requests (https://docs.oracle.com/iaas/Content/General/service-limits/requests.htm).
//

package limitsincrease

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// LimitsIncreaseRequest The properties that define a limit increase request.
// For information about limit increase requests, see
// Working with Limit Increase Requests (https://docs.oracle.com/iaas/Content/General/service-limits/requests.htm).
type LimitsIncreaseRequest struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the limit increase request.
	Id *string `mandatory:"true" json:"id"`

	// The user-friendly name of the limit increase request. Avoid entering confidential information.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment that contains the limit increase request.
	// Note: The tenancy is the root compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The current status of the limit increase request.
	LifecycleState LimitsIncreaseRequestLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// List of limit increase items within the limit increase request.
	LimitsIncreaseItemRequests []LimitsIncreaseItemRequest `mandatory:"true" json:"limitsIncreaseItemRequests"`

	// The time that the limit increase request was created. Format defined by RFC3339.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// Simple key-value pair that is applied without any predefined name, type, or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"true" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"true" json:"definedTags"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the tenancy subscription for the limit increase request.
	SubscriptionId *string `mandatory:"false" json:"subscriptionId"`

	// Reason for the limit increase request.
	Justification *string `mandatory:"false" json:"justification"`

	// The time that the limit increase request was updated. Format defined by RFC3339.
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// The time that the limit increase request was resolved (reached a terminal state). Format defined by RFC3339.
	TimeResolution *common.SDKTime `mandatory:"false" json:"timeResolution"`

	// List of comments in the limit increase request. A comment is typically between the requester and OCI Support.
	CustomerComments []LimitsIncreaseRequestComment `mandatory:"false" json:"customerComments"`

	// System tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
}

func (m LimitsIncreaseRequest) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m LimitsIncreaseRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingLimitsIncreaseRequestLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetLimitsIncreaseRequestLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// LimitsIncreaseRequestLifecycleStateEnum Enum with underlying type: string
type LimitsIncreaseRequestLifecycleStateEnum string

// Set of constants representing the allowable values for LimitsIncreaseRequestLifecycleStateEnum
const (
	LimitsIncreaseRequestLifecycleStateAccepted           LimitsIncreaseRequestLifecycleStateEnum = "ACCEPTED"
	LimitsIncreaseRequestLifecycleStateInProgress         LimitsIncreaseRequestLifecycleStateEnum = "IN_PROGRESS"
	LimitsIncreaseRequestLifecycleStateSucceeded          LimitsIncreaseRequestLifecycleStateEnum = "SUCCEEDED"
	LimitsIncreaseRequestLifecycleStateCanceled           LimitsIncreaseRequestLifecycleStateEnum = "CANCELED"
	LimitsIncreaseRequestLifecycleStatePartiallySucceeded LimitsIncreaseRequestLifecycleStateEnum = "PARTIALLY_SUCCEEDED"
	LimitsIncreaseRequestLifecycleStateFailed             LimitsIncreaseRequestLifecycleStateEnum = "FAILED"
)

var mappingLimitsIncreaseRequestLifecycleStateEnum = map[string]LimitsIncreaseRequestLifecycleStateEnum{
	"ACCEPTED":            LimitsIncreaseRequestLifecycleStateAccepted,
	"IN_PROGRESS":         LimitsIncreaseRequestLifecycleStateInProgress,
	"SUCCEEDED":           LimitsIncreaseRequestLifecycleStateSucceeded,
	"CANCELED":            LimitsIncreaseRequestLifecycleStateCanceled,
	"PARTIALLY_SUCCEEDED": LimitsIncreaseRequestLifecycleStatePartiallySucceeded,
	"FAILED":              LimitsIncreaseRequestLifecycleStateFailed,
}

var mappingLimitsIncreaseRequestLifecycleStateEnumLowerCase = map[string]LimitsIncreaseRequestLifecycleStateEnum{
	"accepted":            LimitsIncreaseRequestLifecycleStateAccepted,
	"in_progress":         LimitsIncreaseRequestLifecycleStateInProgress,
	"succeeded":           LimitsIncreaseRequestLifecycleStateSucceeded,
	"canceled":            LimitsIncreaseRequestLifecycleStateCanceled,
	"partially_succeeded": LimitsIncreaseRequestLifecycleStatePartiallySucceeded,
	"failed":              LimitsIncreaseRequestLifecycleStateFailed,
}

// GetLimitsIncreaseRequestLifecycleStateEnumValues Enumerates the set of values for LimitsIncreaseRequestLifecycleStateEnum
func GetLimitsIncreaseRequestLifecycleStateEnumValues() []LimitsIncreaseRequestLifecycleStateEnum {
	values := make([]LimitsIncreaseRequestLifecycleStateEnum, 0)
	for _, v := range mappingLimitsIncreaseRequestLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetLimitsIncreaseRequestLifecycleStateEnumStringValues Enumerates the set of values in String for LimitsIncreaseRequestLifecycleStateEnum
func GetLimitsIncreaseRequestLifecycleStateEnumStringValues() []string {
	return []string{
		"ACCEPTED",
		"IN_PROGRESS",
		"SUCCEEDED",
		"CANCELED",
		"PARTIALLY_SUCCEEDED",
		"FAILED",
	}
}

// GetMappingLimitsIncreaseRequestLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingLimitsIncreaseRequestLifecycleStateEnum(val string) (LimitsIncreaseRequestLifecycleStateEnum, bool) {
	enum, ok := mappingLimitsIncreaseRequestLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
