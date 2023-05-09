// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Access Governance APIs
//
// Use the Oracle Access Governance API to create, view, and manage GovernanceInstances.
//

package accessgovernancecp

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// UpdateSenderConfig Update to a sender information for email notifications sent by GovernanceInstance.
type UpdateSenderConfig struct {

	// The sender's email.
	Email *string `mandatory:"true" json:"email"`

	// Whether the sender email has inbox configured to receive emails.
	IsInboxConfigured *bool `mandatory:"true" json:"isInboxConfigured"`

	// The sender's displayName.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Whether there is a need to resend the verification email.
	IsResendNotificationEmail *bool `mandatory:"false" json:"isResendNotificationEmail"`
}

func (m UpdateSenderConfig) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpdateSenderConfig) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
