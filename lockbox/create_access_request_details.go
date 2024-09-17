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

// CreateAccessRequestDetails The configuration details for a new access request.
// We don't accept a compartmentId parameter because it is implied to be the same as the lockbox as a subsidiary resource.
// The requestorId is also based on the caller user info.
type CreateAccessRequestDetails struct {

	// The unique identifier (OCID) of the lockbox box that the access request is associated with which is immutable.
	LockboxId *string `mandatory:"true" json:"lockboxId"`

	// The rationale for requesting the access request.
	Description *string `mandatory:"true" json:"description"`

	// The maximum amount of time operator has access to associated resources.
	AccessDuration *string `mandatory:"true" json:"accessDuration"`

	// The name of the access request.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// The context object containing the access request specific details.
	Context map[string]string `mandatory:"false" json:"context"`

	// The ticket number raised by external customers
	// Example: `3-37509643121`
	TicketNumber *string `mandatory:"false" json:"ticketNumber"`
}

func (m CreateAccessRequestDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateAccessRequestDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
