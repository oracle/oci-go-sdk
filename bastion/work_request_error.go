// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Bastion API
//
// Use the Bastion API to provide restricted and time-limited access to target resources that don't have public endpoints. Bastions let authorized users connect from specific IP addresses to target resources using Secure Shell (SSH) sessions. For more information, see the Bastion documentation (https://docs.oracle.com/iaas/Content/Bastion/home.htm).
//

package bastion

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// WorkRequestError An error encountered while executing a work request.
type WorkRequestError struct {

	// A machine-usable code for the error that occurred. Error codes are listed on
	// (https://docs.oracle.com/iaas/Content/API/References/apierrors.htm)
	Code *string `mandatory:"true" json:"code"`

	// A human readable description of the issue encountered.
	Message *string `mandatory:"true" json:"message"`

	// The time the error occurred. An RFC3339 formatted datetime string.
	Timestamp *common.SDKTime `mandatory:"true" json:"timestamp"`
}

func (m WorkRequestError) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m WorkRequestError) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
