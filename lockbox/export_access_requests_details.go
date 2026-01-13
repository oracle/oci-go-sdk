// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
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

// ExportAccessRequestsDetails Details for generating report of Access Requests to export action
type ExportAccessRequestsDetails struct {

	// The unique identifier (OCID) of the lockbox box that the access request is associated with which is immutable.
	LockboxId *string `mandatory:"true" json:"lockboxId"`

	// Date and time after which access requests were created, as described in RFC 3339 (https://tools.ietf.org/rfc/rfc3339)
	TimeCreatedAfter *common.SDKTime `mandatory:"true" json:"timeCreatedAfter"`

	// Date and time before which access requests were created, as described in RFC 3339 (https://tools.ietf.org/rfc/rfc3339)s
	TimeCreatedBefore *common.SDKTime `mandatory:"true" json:"timeCreatedBefore"`
}

func (m ExportAccessRequestsDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ExportAccessRequestsDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
