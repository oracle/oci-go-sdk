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

// AccessRequest An access request to a customer's resource.
// An access request is a subsidiary resource of the Lockbox entity.
type AccessRequest struct {

	// The unique identifier (OCID) of the access request, which can't be changed after creation.
	Id *string `mandatory:"true" json:"id"`

	// The unique identifier (OCID) of the lockbox box that the access request is associated with, which can't be changed after creation.
	LockboxId *string `mandatory:"true" json:"lockboxId"`

	// The name of the access request.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The rationale for requesting the access request and any other related details..
	Description *string `mandatory:"true" json:"description"`

	// The unique identifier of the requestor.
	RequestorId *string `mandatory:"true" json:"requestorId"`

	// Possible access request lifecycle states.
	LifecycleState AccessRequestLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// Details of access request lifecycle state.
	LifecycleStateDetails AccessRequestLifecycleStateDetailsEnum `mandatory:"true" json:"lifecycleStateDetails"`

	// The maximum amount of time operator has access to associated resources.
	AccessDuration *string `mandatory:"true" json:"accessDuration"`

	// The actions taken by different persona on the access request, e.g. approve/deny/revoke
	ActivityLogs []ActivityLog `mandatory:"true" json:"activityLogs"`

	// The time the access request was created. Format is defined by RFC3339 (https://tools.ietf.org/html/rfc3339).
	// Example: `2020-01-25T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The time the access request was last updated. Format is defined by RFC3339 (https://tools.ietf.org/html/rfc3339).
	// Example: `2020-01-25T21:10:29.600Z`
	TimeUpdated *common.SDKTime `mandatory:"true" json:"timeUpdated"`

	// The time the access request expired. Format is defined by RFC3339 (https://tools.ietf.org/html/rfc3339).
	// Example: `2020-01-25T21:10:29.600Z`
	TimeExpired *common.SDKTime `mandatory:"true" json:"timeExpired"`

	// The time the access request was last reminded. Format is defined by RFC3339 (https://tools.ietf.org/html/rfc3339).
	// Example: `2020-01-25T21:10:29.600Z`
	TimeReminded *common.SDKTime `mandatory:"true" json:"timeReminded"`

	// The count of times the access request was reminded.
	ReminderCount *int `mandatory:"true" json:"reminderCount"`

	// The location of the requestor. Format with be two letters indicatiog operator's country code defined by https://jira-sd.mc1.oracleiaas.com/browse/SSD-17880
	// Example: `US`
	RequestorLocation *string `mandatory:"true" json:"requestorLocation"`

	// The context object containing the access request specific details.
	Context map[string]string `mandatory:"false" json:"context"`

	// The ticket number raised by external customers
	// Example: `3-37509643121`
	TicketNumber *string `mandatory:"false" json:"ticketNumber"`
}

func (m AccessRequest) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m AccessRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingAccessRequestLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetAccessRequestLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingAccessRequestLifecycleStateDetailsEnum(string(m.LifecycleStateDetails)); !ok && m.LifecycleStateDetails != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleStateDetails: %s. Supported values are: %s.", m.LifecycleStateDetails, strings.Join(GetAccessRequestLifecycleStateDetailsEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// AccessRequestLifecycleStateEnum Enum with underlying type: string
type AccessRequestLifecycleStateEnum string

// Set of constants representing the allowable values for AccessRequestLifecycleStateEnum
const (
	AccessRequestLifecycleStateInProgress AccessRequestLifecycleStateEnum = "IN_PROGRESS"
	AccessRequestLifecycleStateWaiting    AccessRequestLifecycleStateEnum = "WAITING"
	AccessRequestLifecycleStateSucceeded  AccessRequestLifecycleStateEnum = "SUCCEEDED"
	AccessRequestLifecycleStateCanceling  AccessRequestLifecycleStateEnum = "CANCELING"
	AccessRequestLifecycleStateCanceled   AccessRequestLifecycleStateEnum = "CANCELED"
	AccessRequestLifecycleStateFailed     AccessRequestLifecycleStateEnum = "FAILED"
)

var mappingAccessRequestLifecycleStateEnum = map[string]AccessRequestLifecycleStateEnum{
	"IN_PROGRESS": AccessRequestLifecycleStateInProgress,
	"WAITING":     AccessRequestLifecycleStateWaiting,
	"SUCCEEDED":   AccessRequestLifecycleStateSucceeded,
	"CANCELING":   AccessRequestLifecycleStateCanceling,
	"CANCELED":    AccessRequestLifecycleStateCanceled,
	"FAILED":      AccessRequestLifecycleStateFailed,
}

var mappingAccessRequestLifecycleStateEnumLowerCase = map[string]AccessRequestLifecycleStateEnum{
	"in_progress": AccessRequestLifecycleStateInProgress,
	"waiting":     AccessRequestLifecycleStateWaiting,
	"succeeded":   AccessRequestLifecycleStateSucceeded,
	"canceling":   AccessRequestLifecycleStateCanceling,
	"canceled":    AccessRequestLifecycleStateCanceled,
	"failed":      AccessRequestLifecycleStateFailed,
}

// GetAccessRequestLifecycleStateEnumValues Enumerates the set of values for AccessRequestLifecycleStateEnum
func GetAccessRequestLifecycleStateEnumValues() []AccessRequestLifecycleStateEnum {
	values := make([]AccessRequestLifecycleStateEnum, 0)
	for _, v := range mappingAccessRequestLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetAccessRequestLifecycleStateEnumStringValues Enumerates the set of values in String for AccessRequestLifecycleStateEnum
func GetAccessRequestLifecycleStateEnumStringValues() []string {
	return []string{
		"IN_PROGRESS",
		"WAITING",
		"SUCCEEDED",
		"CANCELING",
		"CANCELED",
		"FAILED",
	}
}

// GetMappingAccessRequestLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAccessRequestLifecycleStateEnum(val string) (AccessRequestLifecycleStateEnum, bool) {
	enum, ok := mappingAccessRequestLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// AccessRequestLifecycleStateDetailsEnum Enum with underlying type: string
type AccessRequestLifecycleStateDetailsEnum string

// Set of constants representing the allowable values for AccessRequestLifecycleStateDetailsEnum
const (
	AccessRequestLifecycleStateDetailsProcessing          AccessRequestLifecycleStateDetailsEnum = "PROCESSING"
	AccessRequestLifecycleStateDetailsWaitingForApprovals AccessRequestLifecycleStateDetailsEnum = "WAITING_FOR_APPROVALS"
	AccessRequestLifecycleStateDetailsApproved            AccessRequestLifecycleStateDetailsEnum = "APPROVED"
	AccessRequestLifecycleStateDetailsAutoApproved        AccessRequestLifecycleStateDetailsEnum = "AUTO_APPROVED"
	AccessRequestLifecycleStateDetailsCancellingAccess    AccessRequestLifecycleStateDetailsEnum = "CANCELLING_ACCESS"
	AccessRequestLifecycleStateDetailsExpired             AccessRequestLifecycleStateDetailsEnum = "EXPIRED"
	AccessRequestLifecycleStateDetailsRevoked             AccessRequestLifecycleStateDetailsEnum = "REVOKED"
	AccessRequestLifecycleStateDetailsDenied              AccessRequestLifecycleStateDetailsEnum = "DENIED"
	AccessRequestLifecycleStateDetailsError               AccessRequestLifecycleStateDetailsEnum = "ERROR"
)

var mappingAccessRequestLifecycleStateDetailsEnum = map[string]AccessRequestLifecycleStateDetailsEnum{
	"PROCESSING":            AccessRequestLifecycleStateDetailsProcessing,
	"WAITING_FOR_APPROVALS": AccessRequestLifecycleStateDetailsWaitingForApprovals,
	"APPROVED":              AccessRequestLifecycleStateDetailsApproved,
	"AUTO_APPROVED":         AccessRequestLifecycleStateDetailsAutoApproved,
	"CANCELLING_ACCESS":     AccessRequestLifecycleStateDetailsCancellingAccess,
	"EXPIRED":               AccessRequestLifecycleStateDetailsExpired,
	"REVOKED":               AccessRequestLifecycleStateDetailsRevoked,
	"DENIED":                AccessRequestLifecycleStateDetailsDenied,
	"ERROR":                 AccessRequestLifecycleStateDetailsError,
}

var mappingAccessRequestLifecycleStateDetailsEnumLowerCase = map[string]AccessRequestLifecycleStateDetailsEnum{
	"processing":            AccessRequestLifecycleStateDetailsProcessing,
	"waiting_for_approvals": AccessRequestLifecycleStateDetailsWaitingForApprovals,
	"approved":              AccessRequestLifecycleStateDetailsApproved,
	"auto_approved":         AccessRequestLifecycleStateDetailsAutoApproved,
	"cancelling_access":     AccessRequestLifecycleStateDetailsCancellingAccess,
	"expired":               AccessRequestLifecycleStateDetailsExpired,
	"revoked":               AccessRequestLifecycleStateDetailsRevoked,
	"denied":                AccessRequestLifecycleStateDetailsDenied,
	"error":                 AccessRequestLifecycleStateDetailsError,
}

// GetAccessRequestLifecycleStateDetailsEnumValues Enumerates the set of values for AccessRequestLifecycleStateDetailsEnum
func GetAccessRequestLifecycleStateDetailsEnumValues() []AccessRequestLifecycleStateDetailsEnum {
	values := make([]AccessRequestLifecycleStateDetailsEnum, 0)
	for _, v := range mappingAccessRequestLifecycleStateDetailsEnum {
		values = append(values, v)
	}
	return values
}

// GetAccessRequestLifecycleStateDetailsEnumStringValues Enumerates the set of values in String for AccessRequestLifecycleStateDetailsEnum
func GetAccessRequestLifecycleStateDetailsEnumStringValues() []string {
	return []string{
		"PROCESSING",
		"WAITING_FOR_APPROVALS",
		"APPROVED",
		"AUTO_APPROVED",
		"CANCELLING_ACCESS",
		"EXPIRED",
		"REVOKED",
		"DENIED",
		"ERROR",
	}
}

// GetMappingAccessRequestLifecycleStateDetailsEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAccessRequestLifecycleStateDetailsEnum(val string) (AccessRequestLifecycleStateDetailsEnum, bool) {
	enum, ok := mappingAccessRequestLifecycleStateDetailsEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
