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

// SubmitEmailDetails Details that are required by the sender to submit a request to send email.
type SubmitEmailDetails struct {
	Sender *Sender `mandatory:"true" json:"sender"`

	Recipients *Recipients `mandatory:"true" json:"recipients"`

	// A short summary of the content, which will appear in the recipient's inbox. UTF-8 supported RFC 2047 (https://www.rfc-editor.org/rfc/rfc2047).
	Subject *string `mandatory:"true" json:"subject"`

	// The unique ID for the email's Message-ID header used for service log correlation. The submission will return an error if the syntax is not a valid RFC 5322 (https://www.rfc-editor.org/rfc/rfc5322) Message-ID. This will be generated if not provided.
	// Example: sdiofu234qwermls24fd@mail.example.com
	MessageId *string `mandatory:"false" json:"messageId"`

	// HTML body content in UTF-8.
	// NOTE: Even though bodytext and bodyhtml are both optional, at least one of them must be provided.
	BodyHtml *string `mandatory:"false" json:"bodyHtml"`

	// Text body content.
	// NOTE: Even though bodytext and bodyhtml are both optional, at least one of them must be provided.
	BodyText *string `mandatory:"false" json:"bodyText"`

	// The email address for the recipient to reply to. If left blank, defaults to the sender address.
	ReplyTo []EmailAddress `mandatory:"false" json:"replyTo"`

	// The header used by the customer for the email sent. Reserved headers are not allowed e.g "subject", "from", and "to" etc.
	// Example: `{"bar-key": "value"}`
	HeaderFields map[string]string `mandatory:"false" json:"headerFields"`
}

func (m SubmitEmailDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m SubmitEmailDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
