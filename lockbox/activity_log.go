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

// ActivityLog The log of the action taken by different persona on the access request, e.g. approve/deny/revoke
type ActivityLog struct {

	// User OCID of the persona
	UserId *string `mandatory:"false" json:"userId"`

	// Level of the persona
	UserLevel PersonaLevelEnum `mandatory:"false" json:"userLevel,omitempty"`

	// The action take by persona
	Action AccessRequestActionTypeEnum `mandatory:"false" json:"action,omitempty"`

	// The action justification or details.
	Message *string `mandatory:"false" json:"message"`

	// The time the action was taken. Format is defined by RFC3339 (https://tools.ietf.org/html/rfc3339).
	// Example: `2020-01-25T21:10:29.600Z`
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`
}

func (m ActivityLog) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ActivityLog) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingPersonaLevelEnum(string(m.UserLevel)); !ok && m.UserLevel != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for UserLevel: %s. Supported values are: %s.", m.UserLevel, strings.Join(GetPersonaLevelEnumStringValues(), ",")))
	}
	if _, ok := GetMappingAccessRequestActionTypeEnum(string(m.Action)); !ok && m.Action != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Action: %s. Supported values are: %s.", m.Action, strings.Join(GetAccessRequestActionTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
