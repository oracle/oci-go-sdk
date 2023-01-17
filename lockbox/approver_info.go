// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Managed Access API
//
// Use the Managed Access API to approve access requests, create and manage templates, and manage resource approval settings. For more information, see Managed Access Overview (https://docs.oracle.com/en-us/iaas/managed-access/overview.htm).
// Use the table of contents and search tool to explore the Managed Access API.
//

package lockbox

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ApproverInfo The approver data for this approver level.
type ApproverInfo struct {

	// The approver type of this approver level.
	ApproverType ApproverTypeEnum `mandatory:"true" json:"approverType"`

	// The group or user ocid of the approver for this approver level.
	ApproverId *string `mandatory:"true" json:"approverId"`
}

func (m ApproverInfo) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ApproverInfo) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingApproverTypeEnum(string(m.ApproverType)); !ok && m.ApproverType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ApproverType: %s. Supported values are: %s.", m.ApproverType, strings.Join(GetApproverTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
