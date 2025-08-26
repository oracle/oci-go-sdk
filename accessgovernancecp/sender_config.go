// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
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

// SenderConfig The sender information for email notifications sent by GovernanceInstance.
type SenderConfig struct {

	// The sender's displayName.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// The sender's email.
	Email *string `mandatory:"false" json:"email"`

	// Whether or not the sender's email has been verified.
	IsVerified *bool `mandatory:"false" json:"isVerified"`

	// The time when the verify response needs to be received by.
	TimeVerifyResponseExpiry *common.SDKTime `mandatory:"false" json:"timeVerifyResponseExpiry"`

	// Whether the sender email has inbox configured to receive emails.
	IsInboxConfigured *bool `mandatory:"false" json:"isInboxConfigured"`
}

func (m SenderConfig) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m SenderConfig) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
