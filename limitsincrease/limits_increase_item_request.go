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

// LimitsIncreaseItemRequest The properties that define a limit increase item in a limit increase request.
// For information about limit increase requests, see
// Working with Limit Increase Requests (https://docs.oracle.com/iaas/Content/General/service-limits/requests.htm).
type LimitsIncreaseItemRequest struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the limit increase item.
	Id *string `mandatory:"true" json:"id"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment that contains the limit increase item.
	// Note: The tenancy is the root compartment).
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The region code for the limit increase item.
	Region *string `mandatory:"true" json:"region"`

	// The name of the service that owns the limit.
	ServiceName *string `mandatory:"true" json:"serviceName"`

	// The name of the limit to be increased.
	LimitName *string `mandatory:"true" json:"limitName"`

	// The value of the limit for the tenancy at the time of the request. Purely informative.
	CurrentValue *int64 `mandatory:"true" json:"currentValue"`

	// The new value requested for the limit.
	Value *int64 `mandatory:"true" json:"value"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the limit increase request that the limit increase item belongs to.
	LimitsIncreaseRequestId *string `mandatory:"true" json:"limitsIncreaseRequestId"`

	// The time that the limit increase item was created. Format defined by RFC3339.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The current status of the limit increase item.
	LifecycleState LimitsIncreaseItemRequestLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The scope of the limit increase item.
	// For OCI, use the logical availability domain (AD) of the request.
	// For Multicloud, use the availability zone (AZ) of the request
	Scope *string `mandatory:"false" json:"scope"`

	// The time that the limit increase item was resolved. Format defined by RFC3339.
	TimeResolution *common.SDKTime `mandatory:"false" json:"timeResolution"`

	// The time that the limit increase item was last updated. Format defined by RFC3339.
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// List of questionnaire responses for the limit increase item.
	QuestionnaireResponse []LimitsIncreaseItemQuestionResponse `mandatory:"false" json:"questionnaireResponse"`
}

func (m LimitsIncreaseItemRequest) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m LimitsIncreaseItemRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingLimitsIncreaseItemRequestLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetLimitsIncreaseItemRequestLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// LimitsIncreaseItemRequestLifecycleStateEnum Enum with underlying type: string
type LimitsIncreaseItemRequestLifecycleStateEnum string

// Set of constants representing the allowable values for LimitsIncreaseItemRequestLifecycleStateEnum
const (
	LimitsIncreaseItemRequestLifecycleStateAccepted   LimitsIncreaseItemRequestLifecycleStateEnum = "ACCEPTED"
	LimitsIncreaseItemRequestLifecycleStateInProgress LimitsIncreaseItemRequestLifecycleStateEnum = "IN_PROGRESS"
	LimitsIncreaseItemRequestLifecycleStateSucceeded  LimitsIncreaseItemRequestLifecycleStateEnum = "SUCCEEDED"
	LimitsIncreaseItemRequestLifecycleStateCanceled   LimitsIncreaseItemRequestLifecycleStateEnum = "CANCELED"
	LimitsIncreaseItemRequestLifecycleStateFailed     LimitsIncreaseItemRequestLifecycleStateEnum = "FAILED"
)

var mappingLimitsIncreaseItemRequestLifecycleStateEnum = map[string]LimitsIncreaseItemRequestLifecycleStateEnum{
	"ACCEPTED":    LimitsIncreaseItemRequestLifecycleStateAccepted,
	"IN_PROGRESS": LimitsIncreaseItemRequestLifecycleStateInProgress,
	"SUCCEEDED":   LimitsIncreaseItemRequestLifecycleStateSucceeded,
	"CANCELED":    LimitsIncreaseItemRequestLifecycleStateCanceled,
	"FAILED":      LimitsIncreaseItemRequestLifecycleStateFailed,
}

var mappingLimitsIncreaseItemRequestLifecycleStateEnumLowerCase = map[string]LimitsIncreaseItemRequestLifecycleStateEnum{
	"accepted":    LimitsIncreaseItemRequestLifecycleStateAccepted,
	"in_progress": LimitsIncreaseItemRequestLifecycleStateInProgress,
	"succeeded":   LimitsIncreaseItemRequestLifecycleStateSucceeded,
	"canceled":    LimitsIncreaseItemRequestLifecycleStateCanceled,
	"failed":      LimitsIncreaseItemRequestLifecycleStateFailed,
}

// GetLimitsIncreaseItemRequestLifecycleStateEnumValues Enumerates the set of values for LimitsIncreaseItemRequestLifecycleStateEnum
func GetLimitsIncreaseItemRequestLifecycleStateEnumValues() []LimitsIncreaseItemRequestLifecycleStateEnum {
	values := make([]LimitsIncreaseItemRequestLifecycleStateEnum, 0)
	for _, v := range mappingLimitsIncreaseItemRequestLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetLimitsIncreaseItemRequestLifecycleStateEnumStringValues Enumerates the set of values in String for LimitsIncreaseItemRequestLifecycleStateEnum
func GetLimitsIncreaseItemRequestLifecycleStateEnumStringValues() []string {
	return []string{
		"ACCEPTED",
		"IN_PROGRESS",
		"SUCCEEDED",
		"CANCELED",
		"FAILED",
	}
}

// GetMappingLimitsIncreaseItemRequestLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingLimitsIncreaseItemRequestLifecycleStateEnum(val string) (LimitsIncreaseItemRequestLifecycleStateEnum, bool) {
	enum, ok := mappingLimitsIncreaseItemRequestLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
