// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// MarketplacePrivateOffer API
//
// Use the Marketplace Publisher API to manage the publishing of applications in Oracle Cloud Infrastructure Marketplace.
//

package marketplaceprivateoffer

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CreateQuoteAttachmentDetails The details for creating a new offer quote attachment.
type CreateQuoteAttachmentDetails struct {

	// Base64-encoded file to attach to the offer quote.
	FileBase64Encoded []byte `mandatory:"true" json:"fileBase64Encoded"`

	// The name used to refer to the uploaded data.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The type of offer quote attachment.
	Type QuoteAttachmentTypeEnum `mandatory:"true" json:"type"`

	// The recipient of the offer quote attachment.
	Recipient QuoteAttachmentRecipientEnum `mandatory:"true" json:"recipient"`
}

func (m CreateQuoteAttachmentDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateQuoteAttachmentDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingQuoteAttachmentTypeEnum(string(m.Type)); !ok && m.Type != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Type: %s. Supported values are: %s.", m.Type, strings.Join(GetQuoteAttachmentTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingQuoteAttachmentRecipientEnum(string(m.Recipient)); !ok && m.Recipient != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Recipient: %s. Supported values are: %s.", m.Recipient, strings.Join(GetQuoteAttachmentRecipientEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
