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

// AccessRequestSummary Summary information for an access request.
type AccessRequestSummary struct {

	// The unique identifier (OCID) of the access request, which can't be changed after creation.
	Id *string `mandatory:"true" json:"id"`

	// The unique identifier (OCID) of the lockbox box that the access request is associated with, which can't be changed after creation.
	LockboxId *string `mandatory:"true" json:"lockboxId"`

	// The name of the access request.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The rationale for requesting the access request.
	Description *string `mandatory:"true" json:"description"`

	// The unique identifier of the requestor.
	RequestorId *string `mandatory:"true" json:"requestorId"`

	// The current state of the access request.
	LifecycleState AccessRequestLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The time the access request was created. Format is defined by RFC3339 (https://tools.ietf.org/html/rfc3339).
	// Example: `2020-01-25T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The time the access request was last updated. Format is defined by RFC3339 (https://tools.ietf.org/html/rfc3339).
	// Example: `2020-01-25T21:10:29.600Z`
	TimeUpdated *common.SDKTime `mandatory:"true" json:"timeUpdated"`

	// The time the access request expired. Format is defined by RFC3339 (https://tools.ietf.org/html/rfc3339).
	// Example: `2020-01-25T21:10:29.600Z`
	TimeExpired *common.SDKTime `mandatory:"true" json:"timeExpired"`

	// The two-char country code of the requestor while creating the access request
	// Example: `US`
	RequestorLocation *string `mandatory:"false" json:"requestorLocation"`

	// The maximum amount of time operator has access to associated resources.
	AccessDuration *string `mandatory:"false" json:"accessDuration"`

	// The ticket number raised by external customers
	// Example: `3-37509643121`
	TicketNumber *string `mandatory:"false" json:"ticketNumber"`

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

func (m AccessRequestSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m AccessRequestSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingAccessRequestLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetAccessRequestLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
