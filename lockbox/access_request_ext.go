// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
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

// AccessRequestExt An access request to a customer's resource that includes additional requestor metadata.
// An access request is a subsidiary resource of the Lockbox entity.
type AccessRequestExt struct {

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
	LifecycleState AccessRequestExtLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// Details of access request lifecycle state.
	LifecycleStateDetails AccessRequestExtLifecycleStateDetailsEnum `mandatory:"true" json:"lifecycleStateDetails"`

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

	// The user name i.e. userId of the requestor.
	RequestorUserName *string `mandatory:"true" json:"requestorUserName"`

	// The context object containing the access request specific details.
	Context map[string]string `mandatory:"false" json:"context"`

	// The ticket number raised by external customers
	// Example: `3-37509643121`
	TicketNumber *string `mandatory:"false" json:"ticketNumber"`
}

func (m AccessRequestExt) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m AccessRequestExt) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingAccessRequestExtLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetAccessRequestExtLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingAccessRequestExtLifecycleStateDetailsEnum(string(m.LifecycleStateDetails)); !ok && m.LifecycleStateDetails != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleStateDetails: %s. Supported values are: %s.", m.LifecycleStateDetails, strings.Join(GetAccessRequestExtLifecycleStateDetailsEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// AccessRequestExtLifecycleStateEnum Enum with underlying type: string
type AccessRequestExtLifecycleStateEnum string

// Set of constants representing the allowable values for AccessRequestExtLifecycleStateEnum
const (
	AccessRequestExtLifecycleStateInProgress AccessRequestExtLifecycleStateEnum = "IN_PROGRESS"
	AccessRequestExtLifecycleStateWaiting    AccessRequestExtLifecycleStateEnum = "WAITING"
	AccessRequestExtLifecycleStateSucceeded  AccessRequestExtLifecycleStateEnum = "SUCCEEDED"
	AccessRequestExtLifecycleStateCanceling  AccessRequestExtLifecycleStateEnum = "CANCELING"
	AccessRequestExtLifecycleStateCanceled   AccessRequestExtLifecycleStateEnum = "CANCELED"
	AccessRequestExtLifecycleStateFailed     AccessRequestExtLifecycleStateEnum = "FAILED"
)

var mappingAccessRequestExtLifecycleStateEnum = map[string]AccessRequestExtLifecycleStateEnum{
	"IN_PROGRESS": AccessRequestExtLifecycleStateInProgress,
	"WAITING":     AccessRequestExtLifecycleStateWaiting,
	"SUCCEEDED":   AccessRequestExtLifecycleStateSucceeded,
	"CANCELING":   AccessRequestExtLifecycleStateCanceling,
	"CANCELED":    AccessRequestExtLifecycleStateCanceled,
	"FAILED":      AccessRequestExtLifecycleStateFailed,
}

var mappingAccessRequestExtLifecycleStateEnumLowerCase = map[string]AccessRequestExtLifecycleStateEnum{
	"in_progress": AccessRequestExtLifecycleStateInProgress,
	"waiting":     AccessRequestExtLifecycleStateWaiting,
	"succeeded":   AccessRequestExtLifecycleStateSucceeded,
	"canceling":   AccessRequestExtLifecycleStateCanceling,
	"canceled":    AccessRequestExtLifecycleStateCanceled,
	"failed":      AccessRequestExtLifecycleStateFailed,
}

// GetAccessRequestExtLifecycleStateEnumValues Enumerates the set of values for AccessRequestExtLifecycleStateEnum
func GetAccessRequestExtLifecycleStateEnumValues() []AccessRequestExtLifecycleStateEnum {
	values := make([]AccessRequestExtLifecycleStateEnum, 0)
	for _, v := range mappingAccessRequestExtLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetAccessRequestExtLifecycleStateEnumStringValues Enumerates the set of values in String for AccessRequestExtLifecycleStateEnum
func GetAccessRequestExtLifecycleStateEnumStringValues() []string {
	return []string{
		"IN_PROGRESS",
		"WAITING",
		"SUCCEEDED",
		"CANCELING",
		"CANCELED",
		"FAILED",
	}
}

// GetMappingAccessRequestExtLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAccessRequestExtLifecycleStateEnum(val string) (AccessRequestExtLifecycleStateEnum, bool) {
	enum, ok := mappingAccessRequestExtLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// AccessRequestExtLifecycleStateDetailsEnum Enum with underlying type: string
type AccessRequestExtLifecycleStateDetailsEnum string

// Set of constants representing the allowable values for AccessRequestExtLifecycleStateDetailsEnum
const (
	AccessRequestExtLifecycleStateDetailsProcessing          AccessRequestExtLifecycleStateDetailsEnum = "PROCESSING"
	AccessRequestExtLifecycleStateDetailsWaitingForApprovals AccessRequestExtLifecycleStateDetailsEnum = "WAITING_FOR_APPROVALS"
	AccessRequestExtLifecycleStateDetailsApproved            AccessRequestExtLifecycleStateDetailsEnum = "APPROVED"
	AccessRequestExtLifecycleStateDetailsAutoApproved        AccessRequestExtLifecycleStateDetailsEnum = "AUTO_APPROVED"
	AccessRequestExtLifecycleStateDetailsCancellingAccess    AccessRequestExtLifecycleStateDetailsEnum = "CANCELLING_ACCESS"
	AccessRequestExtLifecycleStateDetailsExpired             AccessRequestExtLifecycleStateDetailsEnum = "EXPIRED"
	AccessRequestExtLifecycleStateDetailsRevoked             AccessRequestExtLifecycleStateDetailsEnum = "REVOKED"
	AccessRequestExtLifecycleStateDetailsDenied              AccessRequestExtLifecycleStateDetailsEnum = "DENIED"
	AccessRequestExtLifecycleStateDetailsError               AccessRequestExtLifecycleStateDetailsEnum = "ERROR"
)

var mappingAccessRequestExtLifecycleStateDetailsEnum = map[string]AccessRequestExtLifecycleStateDetailsEnum{
	"PROCESSING":            AccessRequestExtLifecycleStateDetailsProcessing,
	"WAITING_FOR_APPROVALS": AccessRequestExtLifecycleStateDetailsWaitingForApprovals,
	"APPROVED":              AccessRequestExtLifecycleStateDetailsApproved,
	"AUTO_APPROVED":         AccessRequestExtLifecycleStateDetailsAutoApproved,
	"CANCELLING_ACCESS":     AccessRequestExtLifecycleStateDetailsCancellingAccess,
	"EXPIRED":               AccessRequestExtLifecycleStateDetailsExpired,
	"REVOKED":               AccessRequestExtLifecycleStateDetailsRevoked,
	"DENIED":                AccessRequestExtLifecycleStateDetailsDenied,
	"ERROR":                 AccessRequestExtLifecycleStateDetailsError,
}

var mappingAccessRequestExtLifecycleStateDetailsEnumLowerCase = map[string]AccessRequestExtLifecycleStateDetailsEnum{
	"processing":            AccessRequestExtLifecycleStateDetailsProcessing,
	"waiting_for_approvals": AccessRequestExtLifecycleStateDetailsWaitingForApprovals,
	"approved":              AccessRequestExtLifecycleStateDetailsApproved,
	"auto_approved":         AccessRequestExtLifecycleStateDetailsAutoApproved,
	"cancelling_access":     AccessRequestExtLifecycleStateDetailsCancellingAccess,
	"expired":               AccessRequestExtLifecycleStateDetailsExpired,
	"revoked":               AccessRequestExtLifecycleStateDetailsRevoked,
	"denied":                AccessRequestExtLifecycleStateDetailsDenied,
	"error":                 AccessRequestExtLifecycleStateDetailsError,
}

// GetAccessRequestExtLifecycleStateDetailsEnumValues Enumerates the set of values for AccessRequestExtLifecycleStateDetailsEnum
func GetAccessRequestExtLifecycleStateDetailsEnumValues() []AccessRequestExtLifecycleStateDetailsEnum {
	values := make([]AccessRequestExtLifecycleStateDetailsEnum, 0)
	for _, v := range mappingAccessRequestExtLifecycleStateDetailsEnum {
		values = append(values, v)
	}
	return values
}

// GetAccessRequestExtLifecycleStateDetailsEnumStringValues Enumerates the set of values in String for AccessRequestExtLifecycleStateDetailsEnum
func GetAccessRequestExtLifecycleStateDetailsEnumStringValues() []string {
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

// GetMappingAccessRequestExtLifecycleStateDetailsEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAccessRequestExtLifecycleStateDetailsEnum(val string) (AccessRequestExtLifecycleStateDetailsEnum, bool) {
	enum, ok := mappingAccessRequestExtLifecycleStateDetailsEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
