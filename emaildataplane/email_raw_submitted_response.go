// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
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

// EmailRawSubmittedResponse Response object that is returned to sender upon successfully submitting the email request.
type EmailRawSubmittedResponse struct {

	// Email Delivery generated unique Envelope ID of the email submission. If you need to contact Email Delivery about a particular request, please provide the Envelope ID.
	EnvelopeId *string `mandatory:"true" json:"envelopeId"`

	// Return list of suppressed email addresses.
	SuppressedRecipients []string `mandatory:"true" json:"suppressedRecipients"`

	// The unique ID for the email's Message-ID header used for service log correlation.
	// Example: sdiofu234qwermls24fd@mail.example.com
	MessageId *string `mandatory:"false" json:"messageId"`
}

func (m EmailRawSubmittedResponse) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m EmailRawSubmittedResponse) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
