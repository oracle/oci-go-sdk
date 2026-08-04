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

// QuoteAttachment Description of an offer quote attachment.
type QuoteAttachment struct {

	// Unique identifier that is immutable on creation.
	Id *string `mandatory:"true" json:"id"`

	// Unique identifier of the associated offer quote that is immutable on creation.
	OfferQuoteId *string `mandatory:"true" json:"offerQuoteId"`

	// OCID of the reseller's tenancy (root compartment).
	ResellerCompartmentId *string `mandatory:"true" json:"resellerCompartmentId"`

	// OCID of the ISV's tenancy (root compartment).
	IsvCompartmentId *string `mandatory:"true" json:"isvCompartmentId"`

	// The name used to refer to the uploaded data.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The type of offer quote attachment.
	Type QuoteAttachmentTypeEnum `mandatory:"true" json:"type"`

	// The time the offer quote attachment was created. An RFC3339 formatted datetime string.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The current state of the offer quote attachment.
	LifecycleState QuoteAttachmentLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The recipient of the offer quote attachment.
	Recipient QuoteAttachmentRecipientEnum `mandatory:"true" json:"recipient"`

	// The time the offer quote attachment was updated. An RFC3339 formatted datetime string.
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// Provides detailed information about the current state of the offer quote attachment. For example, if the offer quote attachment is in a failed state this message can include specific validation errors.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// The MIME type of the uploaded data.
	MimeType *string `mandatory:"false" json:"mimeType"`

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

func (m QuoteAttachment) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m QuoteAttachment) ValidateEnumValue() (bool, error) {
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

// QuoteAttachmentTypeEnum Enum with underlying type: string
type QuoteAttachmentTypeEnum string

// Set of constants representing the allowable values for QuoteAttachmentTypeEnum
const (
	QuoteAttachmentTypeContractTAndC QuoteAttachmentTypeEnum = "CONTRACT_T_AND_C"
	QuoteAttachmentTypeQuote         QuoteAttachmentTypeEnum = "QUOTE"
	QuoteAttachmentTypeEula          QuoteAttachmentTypeEnum = "EULA"
	QuoteAttachmentTypeTermsOfUse    QuoteAttachmentTypeEnum = "TERMS_OF_USE"
	QuoteAttachmentTypeMisc          QuoteAttachmentTypeEnum = "MISC"
)

var mappingQuoteAttachmentTypeEnum = map[string]QuoteAttachmentTypeEnum{
	"CONTRACT_T_AND_C": QuoteAttachmentTypeContractTAndC,
	"QUOTE":            QuoteAttachmentTypeQuote,
	"EULA":             QuoteAttachmentTypeEula,
	"TERMS_OF_USE":     QuoteAttachmentTypeTermsOfUse,
	"MISC":             QuoteAttachmentTypeMisc,
}

var mappingQuoteAttachmentTypeEnumLowerCase = map[string]QuoteAttachmentTypeEnum{
	"contract_t_and_c": QuoteAttachmentTypeContractTAndC,
	"quote":            QuoteAttachmentTypeQuote,
	"eula":             QuoteAttachmentTypeEula,
	"terms_of_use":     QuoteAttachmentTypeTermsOfUse,
	"misc":             QuoteAttachmentTypeMisc,
}

// GetQuoteAttachmentTypeEnumValues Enumerates the set of values for QuoteAttachmentTypeEnum
func GetQuoteAttachmentTypeEnumValues() []QuoteAttachmentTypeEnum {
	values := make([]QuoteAttachmentTypeEnum, 0)
	for _, v := range mappingQuoteAttachmentTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetQuoteAttachmentTypeEnumStringValues Enumerates the set of values in String for QuoteAttachmentTypeEnum
func GetQuoteAttachmentTypeEnumStringValues() []string {
	return []string{
		"CONTRACT_T_AND_C",
		"QUOTE",
		"EULA",
		"TERMS_OF_USE",
		"MISC",
	}
}

// GetMappingQuoteAttachmentTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingQuoteAttachmentTypeEnum(val string) (QuoteAttachmentTypeEnum, bool) {
	enum, ok := mappingQuoteAttachmentTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// QuoteAttachmentLifecycleStateEnum Enum with underlying type: string
type QuoteAttachmentLifecycleStateEnum string

// Set of constants representing the allowable values for QuoteAttachmentLifecycleStateEnum
const (
	QuoteAttachmentLifecycleStateCreating QuoteAttachmentLifecycleStateEnum = "CREATING"
	QuoteAttachmentLifecycleStateUpdating QuoteAttachmentLifecycleStateEnum = "UPDATING"
	QuoteAttachmentLifecycleStateActive   QuoteAttachmentLifecycleStateEnum = "ACTIVE"
	QuoteAttachmentLifecycleStateDeleting QuoteAttachmentLifecycleStateEnum = "DELETING"
	QuoteAttachmentLifecycleStateDeleted  QuoteAttachmentLifecycleStateEnum = "DELETED"
	QuoteAttachmentLifecycleStateFailed   QuoteAttachmentLifecycleStateEnum = "FAILED"
)

var mappingQuoteAttachmentLifecycleStateEnum = map[string]QuoteAttachmentLifecycleStateEnum{
	"CREATING": QuoteAttachmentLifecycleStateCreating,
	"UPDATING": QuoteAttachmentLifecycleStateUpdating,
	"ACTIVE":   QuoteAttachmentLifecycleStateActive,
	"DELETING": QuoteAttachmentLifecycleStateDeleting,
	"DELETED":  QuoteAttachmentLifecycleStateDeleted,
	"FAILED":   QuoteAttachmentLifecycleStateFailed,
}

var mappingQuoteAttachmentLifecycleStateEnumLowerCase = map[string]QuoteAttachmentLifecycleStateEnum{
	"creating": QuoteAttachmentLifecycleStateCreating,
	"updating": QuoteAttachmentLifecycleStateUpdating,
	"active":   QuoteAttachmentLifecycleStateActive,
	"deleting": QuoteAttachmentLifecycleStateDeleting,
	"deleted":  QuoteAttachmentLifecycleStateDeleted,
	"failed":   QuoteAttachmentLifecycleStateFailed,
}

// GetQuoteAttachmentLifecycleStateEnumValues Enumerates the set of values for QuoteAttachmentLifecycleStateEnum
func GetQuoteAttachmentLifecycleStateEnumValues() []QuoteAttachmentLifecycleStateEnum {
	values := make([]QuoteAttachmentLifecycleStateEnum, 0)
	for _, v := range mappingQuoteAttachmentLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetQuoteAttachmentLifecycleStateEnumStringValues Enumerates the set of values in String for QuoteAttachmentLifecycleStateEnum
func GetQuoteAttachmentLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"UPDATING",
		"ACTIVE",
		"DELETING",
		"DELETED",
		"FAILED",
	}
}

// GetMappingQuoteAttachmentLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingQuoteAttachmentLifecycleStateEnum(val string) (QuoteAttachmentLifecycleStateEnum, bool) {
	enum, ok := mappingQuoteAttachmentLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// QuoteAttachmentRecipientEnum Enum with underlying type: string
type QuoteAttachmentRecipientEnum string

// Set of constants representing the allowable values for QuoteAttachmentRecipientEnum
const (
	QuoteAttachmentRecipientReseller QuoteAttachmentRecipientEnum = "RESELLER"
	QuoteAttachmentRecipientBuyer    QuoteAttachmentRecipientEnum = "BUYER"
)

var mappingQuoteAttachmentRecipientEnum = map[string]QuoteAttachmentRecipientEnum{
	"RESELLER": QuoteAttachmentRecipientReseller,
	"BUYER":    QuoteAttachmentRecipientBuyer,
}

var mappingQuoteAttachmentRecipientEnumLowerCase = map[string]QuoteAttachmentRecipientEnum{
	"reseller": QuoteAttachmentRecipientReseller,
	"buyer":    QuoteAttachmentRecipientBuyer,
}

// GetQuoteAttachmentRecipientEnumValues Enumerates the set of values for QuoteAttachmentRecipientEnum
func GetQuoteAttachmentRecipientEnumValues() []QuoteAttachmentRecipientEnum {
	values := make([]QuoteAttachmentRecipientEnum, 0)
	for _, v := range mappingQuoteAttachmentRecipientEnum {
		values = append(values, v)
	}
	return values
}

// GetQuoteAttachmentRecipientEnumStringValues Enumerates the set of values in String for QuoteAttachmentRecipientEnum
func GetQuoteAttachmentRecipientEnumStringValues() []string {
	return []string{
		"RESELLER",
		"BUYER",
	}
}

// GetMappingQuoteAttachmentRecipientEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingQuoteAttachmentRecipientEnum(val string) (QuoteAttachmentRecipientEnum, bool) {
	enum, ok := mappingQuoteAttachmentRecipientEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
