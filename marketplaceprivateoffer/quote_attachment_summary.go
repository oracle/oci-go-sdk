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

// QuoteAttachmentSummary Summary of offer quote attachment.
type QuoteAttachmentSummary struct {

	// Unique identifier that is immutable on creation.
	Id *string `mandatory:"true" json:"id"`

	// Unique identifier of the associated offer quote that is immutable on creation.
	OfferQuoteId *string `mandatory:"true" json:"offerQuoteId"`

	// The name used to refer to the uploaded data.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The type of offer quote attachment.
	Type QuoteAttachmentTypeEnum `mandatory:"true" json:"type"`

	// The current state of the offer quote attachment.
	LifecycleState QuoteAttachmentLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// OCID of the reseller's tenancy (root compartment).
	ResellerCompartmentId *string `mandatory:"false" json:"resellerCompartmentId"`

	// OCID of the ISV's tenancy (root compartment).
	IsvCompartmentId *string `mandatory:"false" json:"isvCompartmentId"`

	// The time the offer quote attachment was created. An RFC3339 formatted datetime string.
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// The time the offer quote attachment was updated. An RFC3339 formatted datetime string.
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// Provides detailed information about the current state of the offer quote attachment. For example, if the offer quote attachment is in a failed state this message can include specific validation errors.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// The MIME type of the uploaded data.
	MimeType *string `mandatory:"false" json:"mimeType"`

	// The current state of the offer quote attachment.
	Recipient QuoteAttachmentRecipientEnum `mandatory:"false" json:"recipient,omitempty"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// System tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
}

func (m QuoteAttachmentSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m QuoteAttachmentSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingQuoteAttachmentTypeEnum(string(m.Type)); !ok && m.Type != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Type: %s. Supported values are: %s.", m.Type, strings.Join(GetQuoteAttachmentTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingQuoteAttachmentLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetQuoteAttachmentLifecycleStateEnumStringValues(), ",")))
	}

	if _, ok := GetMappingQuoteAttachmentRecipientEnum(string(m.Recipient)); !ok && m.Recipient != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Recipient: %s. Supported values are: %s.", m.Recipient, strings.Join(GetQuoteAttachmentRecipientEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
