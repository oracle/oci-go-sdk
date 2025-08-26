// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Email Delivery Submission API
//
// Use the Email Delivery API to send high-volume and application-generated emails.
// For more information, see Overview of the Email Delivery Service (https://docs.oracle.com/iaas/Content/Email/Concepts/overview.htm).
//

package emaildataplane

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// Recipients The destination for the email, composed of To, CC, and BCC fields.
// NOTE: At least one of To, CC, and BCC must be provided. And max 50 recipients are allowed across the To:, CC: and BCC: fields.
type Recipients struct {

	// Array of To address.
	To []EmailAddress `mandatory:"false" json:"to"`

	// Array of CC address.
	Cc []EmailAddress `mandatory:"false" json:"cc"`

	// Array of BCC address. Bcc headers can only be viewed by non bcc recipients.
	Bcc []EmailAddress `mandatory:"false" json:"bcc"`
}

func (m Recipients) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m Recipients) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
