// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// MarketplacePublisherService API
//
// Use the Marketplace Publisher API to manage the publishing of applications in Oracle Cloud Infrastructure Marketplace.
//

package marketplacepublisher

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// Attachment Description of Attachment.
type Attachment struct {

	// Unique identifier that is immutable on creation
	Id *string `mandatory:"true" json:"id"`

	// OCID of the seller's tenancy (root compartment).
	SellerCompartmentId *string `mandatory:"true" json:"sellerCompartmentId"`

	// Unique identifier of the associated offer that is immutable on creation
	OfferId *string `mandatory:"true" json:"offerId"`

	// The name used to refer to the uploaded data.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The type of offer attachment.
	Type AttachmentTypeEnum `mandatory:"true" json:"type"`

	// The time the the Offer was created. An RFC3339 formatted datetime string
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The current state of the Offer.
	LifecycleState AttachmentLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"true" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"true" json:"definedTags"`

	// OCID of the buyer's tenancy (root compartment).
	BuyerCompartmentId *string `mandatory:"false" json:"buyerCompartmentId"`

	// The MIME type of the uploaded data.
	MimeType *string `mandatory:"false" json:"mimeType"`
}

func (m Attachment) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m Attachment) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingAttachmentTypeEnum(string(m.Type)); !ok && m.Type != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Type: %s. Supported values are: %s.", m.Type, strings.Join(GetAttachmentTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingAttachmentLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetAttachmentLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// AttachmentTypeEnum Enum with underlying type: string
type AttachmentTypeEnum string

// Set of constants representing the allowable values for AttachmentTypeEnum
const (
	AttachmentTypeContractTAndC AttachmentTypeEnum = "CONTRACT_T_AND_C"
	AttachmentTypeQuote         AttachmentTypeEnum = "QUOTE"
	AttachmentTypeEula          AttachmentTypeEnum = "EULA"
	AttachmentTypeTermsOfUse    AttachmentTypeEnum = "TERMS_OF_USE"
	AttachmentTypeMisc          AttachmentTypeEnum = "MISC"
)

var mappingAttachmentTypeEnum = map[string]AttachmentTypeEnum{
	"CONTRACT_T_AND_C": AttachmentTypeContractTAndC,
	"QUOTE":            AttachmentTypeQuote,
	"EULA":             AttachmentTypeEula,
	"TERMS_OF_USE":     AttachmentTypeTermsOfUse,
	"MISC":             AttachmentTypeMisc,
}

var mappingAttachmentTypeEnumLowerCase = map[string]AttachmentTypeEnum{
	"contract_t_and_c": AttachmentTypeContractTAndC,
	"quote":            AttachmentTypeQuote,
	"eula":             AttachmentTypeEula,
	"terms_of_use":     AttachmentTypeTermsOfUse,
	"misc":             AttachmentTypeMisc,
}

// GetAttachmentTypeEnumValues Enumerates the set of values for AttachmentTypeEnum
func GetAttachmentTypeEnumValues() []AttachmentTypeEnum {
	values := make([]AttachmentTypeEnum, 0)
	for _, v := range mappingAttachmentTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetAttachmentTypeEnumStringValues Enumerates the set of values in String for AttachmentTypeEnum
func GetAttachmentTypeEnumStringValues() []string {
	return []string{
		"CONTRACT_T_AND_C",
		"QUOTE",
		"EULA",
		"TERMS_OF_USE",
		"MISC",
	}
}

// GetMappingAttachmentTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAttachmentTypeEnum(val string) (AttachmentTypeEnum, bool) {
	enum, ok := mappingAttachmentTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// AttachmentLifecycleStateEnum Enum with underlying type: string
type AttachmentLifecycleStateEnum string

// Set of constants representing the allowable values for AttachmentLifecycleStateEnum
const (
	AttachmentLifecycleStateCreating AttachmentLifecycleStateEnum = "CREATING"
	AttachmentLifecycleStateUpdating AttachmentLifecycleStateEnum = "UPDATING"
	AttachmentLifecycleStateActive   AttachmentLifecycleStateEnum = "ACTIVE"
	AttachmentLifecycleStateDeleting AttachmentLifecycleStateEnum = "DELETING"
	AttachmentLifecycleStateDeleted  AttachmentLifecycleStateEnum = "DELETED"
	AttachmentLifecycleStateFailed   AttachmentLifecycleStateEnum = "FAILED"
)

var mappingAttachmentLifecycleStateEnum = map[string]AttachmentLifecycleStateEnum{
	"CREATING": AttachmentLifecycleStateCreating,
	"UPDATING": AttachmentLifecycleStateUpdating,
	"ACTIVE":   AttachmentLifecycleStateActive,
	"DELETING": AttachmentLifecycleStateDeleting,
	"DELETED":  AttachmentLifecycleStateDeleted,
	"FAILED":   AttachmentLifecycleStateFailed,
}

var mappingAttachmentLifecycleStateEnumLowerCase = map[string]AttachmentLifecycleStateEnum{
	"creating": AttachmentLifecycleStateCreating,
	"updating": AttachmentLifecycleStateUpdating,
	"active":   AttachmentLifecycleStateActive,
	"deleting": AttachmentLifecycleStateDeleting,
	"deleted":  AttachmentLifecycleStateDeleted,
	"failed":   AttachmentLifecycleStateFailed,
}

// GetAttachmentLifecycleStateEnumValues Enumerates the set of values for AttachmentLifecycleStateEnum
func GetAttachmentLifecycleStateEnumValues() []AttachmentLifecycleStateEnum {
	values := make([]AttachmentLifecycleStateEnum, 0)
	for _, v := range mappingAttachmentLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetAttachmentLifecycleStateEnumStringValues Enumerates the set of values in String for AttachmentLifecycleStateEnum
func GetAttachmentLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"UPDATING",
		"ACTIVE",
		"DELETING",
		"DELETED",
		"FAILED",
	}
}

// GetMappingAttachmentLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAttachmentLifecycleStateEnum(val string) (AttachmentLifecycleStateEnum, bool) {
	enum, ok := mappingAttachmentLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
