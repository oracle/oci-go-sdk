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

// Offer Description of Offer.
type Offer struct {

	// Unique identifier that is immutable on creation
	Id *string `mandatory:"true" json:"id"`

	// Offer Identifier, can be renamed
	DisplayName *string `mandatory:"true" json:"displayName"`

	// OCID of the seller's tenancy (root compartment).
	SellerCompartmentId *string `mandatory:"true" json:"sellerCompartmentId"`

	// The time the the Offer was created. An RFC3339 formatted datetime string
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The current state of the Offer.
	LifecycleState OfferLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"true" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"true" json:"definedTags"`

	// OCID of the buyer's tenancy (root compartment).
	BuyerCompartmentId *string `mandatory:"false" json:"buyerCompartmentId"`

	// The description of the offer
	Description *string `mandatory:"false" json:"description"`

	// The time the Offer will become active after it has been accepted by the Buyer. An RFC3339 formatted datetime string
	TimeStartDate *common.SDKTime `mandatory:"false" json:"timeStartDate"`

	// Duration the Offer will be active after its start date. An ISO8601 extended formatted string.
	Duration *string `mandatory:"false" json:"duration"`

	// The time the Offer was updated. An RFC3339 formatted datetime string
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// The time the Offer must be accepted by the Buyer before the Offer becomes invalid. An RFC3339 formatted datetime string
	TimeAcceptBy *common.SDKTime `mandatory:"false" json:"timeAcceptBy"`

	// The time the Offer was accepted by the Buyer of the Offer. An RFC3339 formatted datetime string
	TimeAccepted *common.SDKTime `mandatory:"false" json:"timeAccepted"`

	// The time the accepted Offer will end. An RFC3339 formatted datetime string
	TimeOfferEnd *common.SDKTime `mandatory:"false" json:"timeOfferEnd"`

	// A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// A substate for lifeCycleStatus to give a more human readable version of the offer
	OfferStatus OfferOfferStatusEnum `mandatory:"false" json:"offerStatus,omitempty"`

	PublisherSummary *PublisherSummary `mandatory:"false" json:"publisherSummary"`

	Pricing *Pricing `mandatory:"false" json:"pricing"`

	BuyerInformation *BuyerInformation `mandatory:"false" json:"buyerInformation"`

	SellerInformation *SellerInformation `mandatory:"false" json:"sellerInformation"`

	// A list of Resource Bundles associated with an Offer.
	ResourceBundles []ResourceBundle `mandatory:"false" json:"resourceBundles"`
}

func (m Offer) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m Offer) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingOfferLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetOfferLifecycleStateEnumStringValues(), ",")))
	}

	if _, ok := GetMappingOfferOfferStatusEnum(string(m.OfferStatus)); !ok && m.OfferStatus != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for OfferStatus: %s. Supported values are: %s.", m.OfferStatus, strings.Join(GetOfferOfferStatusEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// OfferLifecycleStateEnum Enum with underlying type: string
type OfferLifecycleStateEnum string

// Set of constants representing the allowable values for OfferLifecycleStateEnum
const (
	OfferLifecycleStateCreating OfferLifecycleStateEnum = "CREATING"
	OfferLifecycleStateUpdating OfferLifecycleStateEnum = "UPDATING"
	OfferLifecycleStateActive   OfferLifecycleStateEnum = "ACTIVE"
	OfferLifecycleStateDeleting OfferLifecycleStateEnum = "DELETING"
	OfferLifecycleStateDeleted  OfferLifecycleStateEnum = "DELETED"
	OfferLifecycleStateFailed   OfferLifecycleStateEnum = "FAILED"
)

var mappingOfferLifecycleStateEnum = map[string]OfferLifecycleStateEnum{
	"CREATING": OfferLifecycleStateCreating,
	"UPDATING": OfferLifecycleStateUpdating,
	"ACTIVE":   OfferLifecycleStateActive,
	"DELETING": OfferLifecycleStateDeleting,
	"DELETED":  OfferLifecycleStateDeleted,
	"FAILED":   OfferLifecycleStateFailed,
}

var mappingOfferLifecycleStateEnumLowerCase = map[string]OfferLifecycleStateEnum{
	"creating": OfferLifecycleStateCreating,
	"updating": OfferLifecycleStateUpdating,
	"active":   OfferLifecycleStateActive,
	"deleting": OfferLifecycleStateDeleting,
	"deleted":  OfferLifecycleStateDeleted,
	"failed":   OfferLifecycleStateFailed,
}

// GetOfferLifecycleStateEnumValues Enumerates the set of values for OfferLifecycleStateEnum
func GetOfferLifecycleStateEnumValues() []OfferLifecycleStateEnum {
	values := make([]OfferLifecycleStateEnum, 0)
	for _, v := range mappingOfferLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetOfferLifecycleStateEnumStringValues Enumerates the set of values in String for OfferLifecycleStateEnum
func GetOfferLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"UPDATING",
		"ACTIVE",
		"DELETING",
		"DELETED",
		"FAILED",
	}
}

// GetMappingOfferLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingOfferLifecycleStateEnum(val string) (OfferLifecycleStateEnum, bool) {
	enum, ok := mappingOfferLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// OfferOfferStatusEnum Enum with underlying type: string
type OfferOfferStatusEnum string

// Set of constants representing the allowable values for OfferOfferStatusEnum
const (
	OfferOfferStatusDraft              OfferOfferStatusEnum = "DRAFT"
	OfferOfferStatusPendingMarketplace OfferOfferStatusEnum = "PENDING_MARKETPLACE"
	OfferOfferStatusPendingBuyer       OfferOfferStatusEnum = "PENDING_BUYER"
	OfferOfferStatusExpired            OfferOfferStatusEnum = "EXPIRED"
	OfferOfferStatusAccepted           OfferOfferStatusEnum = "ACCEPTED"
	OfferOfferStatusActive             OfferOfferStatusEnum = "ACTIVE"
	OfferOfferStatusEnded              OfferOfferStatusEnum = "ENDED"
	OfferOfferStatusFailedSend         OfferOfferStatusEnum = "FAILED_SEND"
	OfferOfferStatusFailedAccept       OfferOfferStatusEnum = "FAILED_ACCEPT"
)

var mappingOfferOfferStatusEnum = map[string]OfferOfferStatusEnum{
	"DRAFT":               OfferOfferStatusDraft,
	"PENDING_MARKETPLACE": OfferOfferStatusPendingMarketplace,
	"PENDING_BUYER":       OfferOfferStatusPendingBuyer,
	"EXPIRED":             OfferOfferStatusExpired,
	"ACCEPTED":            OfferOfferStatusAccepted,
	"ACTIVE":              OfferOfferStatusActive,
	"ENDED":               OfferOfferStatusEnded,
	"FAILED_SEND":         OfferOfferStatusFailedSend,
	"FAILED_ACCEPT":       OfferOfferStatusFailedAccept,
}

var mappingOfferOfferStatusEnumLowerCase = map[string]OfferOfferStatusEnum{
	"draft":               OfferOfferStatusDraft,
	"pending_marketplace": OfferOfferStatusPendingMarketplace,
	"pending_buyer":       OfferOfferStatusPendingBuyer,
	"expired":             OfferOfferStatusExpired,
	"accepted":            OfferOfferStatusAccepted,
	"active":              OfferOfferStatusActive,
	"ended":               OfferOfferStatusEnded,
	"failed_send":         OfferOfferStatusFailedSend,
	"failed_accept":       OfferOfferStatusFailedAccept,
}

// GetOfferOfferStatusEnumValues Enumerates the set of values for OfferOfferStatusEnum
func GetOfferOfferStatusEnumValues() []OfferOfferStatusEnum {
	values := make([]OfferOfferStatusEnum, 0)
	for _, v := range mappingOfferOfferStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetOfferOfferStatusEnumStringValues Enumerates the set of values in String for OfferOfferStatusEnum
func GetOfferOfferStatusEnumStringValues() []string {
	return []string{
		"DRAFT",
		"PENDING_MARKETPLACE",
		"PENDING_BUYER",
		"EXPIRED",
		"ACCEPTED",
		"ACTIVE",
		"ENDED",
		"FAILED_SEND",
		"FAILED_ACCEPT",
	}
}

// GetMappingOfferOfferStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingOfferOfferStatusEnum(val string) (OfferOfferStatusEnum, bool) {
	enum, ok := mappingOfferOfferStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
