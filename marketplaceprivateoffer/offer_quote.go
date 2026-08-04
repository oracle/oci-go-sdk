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

// OfferQuote The model for the offer quote details.
type OfferQuote struct {

	// Unique identifier that is immutable on creation.
	Id *string `mandatory:"true" json:"id"`

	// Offer quote identifier, can be renamed.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// OCID of the reseller tenancy.
	ResellerCompartmentId *string `mandatory:"true" json:"resellerCompartmentId"`

	// The time the offer quote was created. An RFC3339 formatted datetime string.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The current state of the offer quote.
	LifecycleState OfferQuoteLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// OCID of the ISV tenancy.
	IsvCompartmentId *string `mandatory:"false" json:"isvCompartmentId"`

	// Publisher identifier of the ISV.
	IsvPublisherId *string `mandatory:"false" json:"isvPublisherId"`

	// Display name of the reseller publisher.
	ResellerPublisherDisplayName *string `mandatory:"false" json:"resellerPublisherDisplayName"`

	// Display name of the ISV publisher.
	IsvPublisherDisplayName *string `mandatory:"false" json:"isvPublisherDisplayName"`

	// A list of buyer tenancies.
	BuyerCompartmentIds []string `mandatory:"false" json:"buyerCompartmentIds"`

	// The description of the offer quote.
	Description *string `mandatory:"false" json:"description"`

	// The time the offer quote was updated. An RFC3339 formatted datetime string.
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// Provides detailed information about the current state of the offer quote. For example, if the offer quote is in a failed state this message can include specific validation errors.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// The human readable representation of where the offer quote is at in its contract lifecycle.
	OfferQuoteStatus OfferQuoteOfferQuoteStatusEnum `mandatory:"false" json:"offerQuoteStatus,omitempty"`

	ResellerInformation *ResellerInformation `mandatory:"false" json:"resellerInformation"`

	IsvInformation *IsvInformation `mandatory:"false" json:"isvInformation"`

	Pricing *Pricing `mandatory:"false" json:"pricing"`

	// A list of resource bundles associated with an offer quote.
	ResourceBundles []ResourceBundle `mandatory:"false" json:"resourceBundles"`

	// The time the offer for the buyer must be accepted by before the offer becomes invalid. An RFC3339 formatted datetime string.
	BuyerAcceptanceDeadline *common.SDKTime `mandatory:"false" json:"buyerAcceptanceDeadline"`

	// Duration the offer for the buyer will be active after its start date. An ISO8601 extended formatted string.
	BuyerOfferDuration *string `mandatory:"false" json:"buyerOfferDuration"`

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

func (m OfferQuote) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m OfferQuote) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingOfferQuoteLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetOfferQuoteLifecycleStateEnumStringValues(), ",")))
	}

	if _, ok := GetMappingOfferQuoteOfferQuoteStatusEnum(string(m.OfferQuoteStatus)); !ok && m.OfferQuoteStatus != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for OfferQuoteStatus: %s. Supported values are: %s.", m.OfferQuoteStatus, strings.Join(GetOfferQuoteOfferQuoteStatusEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// OfferQuoteLifecycleStateEnum Enum with underlying type: string
type OfferQuoteLifecycleStateEnum string

// Set of constants representing the allowable values for OfferQuoteLifecycleStateEnum
const (
	OfferQuoteLifecycleStateCreating OfferQuoteLifecycleStateEnum = "CREATING"
	OfferQuoteLifecycleStateUpdating OfferQuoteLifecycleStateEnum = "UPDATING"
	OfferQuoteLifecycleStateActive   OfferQuoteLifecycleStateEnum = "ACTIVE"
	OfferQuoteLifecycleStateDeleting OfferQuoteLifecycleStateEnum = "DELETING"
	OfferQuoteLifecycleStateDeleted  OfferQuoteLifecycleStateEnum = "DELETED"
	OfferQuoteLifecycleStateFailed   OfferQuoteLifecycleStateEnum = "FAILED"
)

var mappingOfferQuoteLifecycleStateEnum = map[string]OfferQuoteLifecycleStateEnum{
	"CREATING": OfferQuoteLifecycleStateCreating,
	"UPDATING": OfferQuoteLifecycleStateUpdating,
	"ACTIVE":   OfferQuoteLifecycleStateActive,
	"DELETING": OfferQuoteLifecycleStateDeleting,
	"DELETED":  OfferQuoteLifecycleStateDeleted,
	"FAILED":   OfferQuoteLifecycleStateFailed,
}

var mappingOfferQuoteLifecycleStateEnumLowerCase = map[string]OfferQuoteLifecycleStateEnum{
	"creating": OfferQuoteLifecycleStateCreating,
	"updating": OfferQuoteLifecycleStateUpdating,
	"active":   OfferQuoteLifecycleStateActive,
	"deleting": OfferQuoteLifecycleStateDeleting,
	"deleted":  OfferQuoteLifecycleStateDeleted,
	"failed":   OfferQuoteLifecycleStateFailed,
}

// GetOfferQuoteLifecycleStateEnumValues Enumerates the set of values for OfferQuoteLifecycleStateEnum
func GetOfferQuoteLifecycleStateEnumValues() []OfferQuoteLifecycleStateEnum {
	values := make([]OfferQuoteLifecycleStateEnum, 0)
	for _, v := range mappingOfferQuoteLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetOfferQuoteLifecycleStateEnumStringValues Enumerates the set of values in String for OfferQuoteLifecycleStateEnum
func GetOfferQuoteLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"UPDATING",
		"ACTIVE",
		"DELETING",
		"DELETED",
		"FAILED",
	}
}

// GetMappingOfferQuoteLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingOfferQuoteLifecycleStateEnum(val string) (OfferQuoteLifecycleStateEnum, bool) {
	enum, ok := mappingOfferQuoteLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// OfferQuoteOfferQuoteStatusEnum Enum with underlying type: string
type OfferQuoteOfferQuoteStatusEnum string

// Set of constants representing the allowable values for OfferQuoteOfferQuoteStatusEnum
const (
	OfferQuoteOfferQuoteStatusDraft            OfferQuoteOfferQuoteStatusEnum = "DRAFT"
	OfferQuoteOfferQuoteStatusPendingIsv       OfferQuoteOfferQuoteStatusEnum = "PENDING_ISV"
	OfferQuoteOfferQuoteStatusPendingIsvUpdate OfferQuoteOfferQuoteStatusEnum = "PENDING_ISV_UPDATE"
	OfferQuoteOfferQuoteStatusPendingReseller  OfferQuoteOfferQuoteStatusEnum = "PENDING_RESELLER"
	OfferQuoteOfferQuoteStatusResellerAccepted OfferQuoteOfferQuoteStatusEnum = "RESELLER_ACCEPTED"
	OfferQuoteOfferQuoteStatusBuyerAccepted    OfferQuoteOfferQuoteStatusEnum = "BUYER_ACCEPTED"
	OfferQuoteOfferQuoteStatusActive           OfferQuoteOfferQuoteStatusEnum = "ACTIVE"
	OfferQuoteOfferQuoteStatusEnded            OfferQuoteOfferQuoteStatusEnum = "ENDED"
	OfferQuoteOfferQuoteStatusExpired          OfferQuoteOfferQuoteStatusEnum = "EXPIRED"
	OfferQuoteOfferQuoteStatusFailedSend       OfferQuoteOfferQuoteStatusEnum = "FAILED_SEND"
	OfferQuoteOfferQuoteStatusFailedAccept     OfferQuoteOfferQuoteStatusEnum = "FAILED_ACCEPT"
	OfferQuoteOfferQuoteStatusFailedRespond    OfferQuoteOfferQuoteStatusEnum = "FAILED_RESPOND"
)

var mappingOfferQuoteOfferQuoteStatusEnum = map[string]OfferQuoteOfferQuoteStatusEnum{
	"DRAFT":              OfferQuoteOfferQuoteStatusDraft,
	"PENDING_ISV":        OfferQuoteOfferQuoteStatusPendingIsv,
	"PENDING_ISV_UPDATE": OfferQuoteOfferQuoteStatusPendingIsvUpdate,
	"PENDING_RESELLER":   OfferQuoteOfferQuoteStatusPendingReseller,
	"RESELLER_ACCEPTED":  OfferQuoteOfferQuoteStatusResellerAccepted,
	"BUYER_ACCEPTED":     OfferQuoteOfferQuoteStatusBuyerAccepted,
	"ACTIVE":             OfferQuoteOfferQuoteStatusActive,
	"ENDED":              OfferQuoteOfferQuoteStatusEnded,
	"EXPIRED":            OfferQuoteOfferQuoteStatusExpired,
	"FAILED_SEND":        OfferQuoteOfferQuoteStatusFailedSend,
	"FAILED_ACCEPT":      OfferQuoteOfferQuoteStatusFailedAccept,
	"FAILED_RESPOND":     OfferQuoteOfferQuoteStatusFailedRespond,
}

var mappingOfferQuoteOfferQuoteStatusEnumLowerCase = map[string]OfferQuoteOfferQuoteStatusEnum{
	"draft":              OfferQuoteOfferQuoteStatusDraft,
	"pending_isv":        OfferQuoteOfferQuoteStatusPendingIsv,
	"pending_isv_update": OfferQuoteOfferQuoteStatusPendingIsvUpdate,
	"pending_reseller":   OfferQuoteOfferQuoteStatusPendingReseller,
	"reseller_accepted":  OfferQuoteOfferQuoteStatusResellerAccepted,
	"buyer_accepted":     OfferQuoteOfferQuoteStatusBuyerAccepted,
	"active":             OfferQuoteOfferQuoteStatusActive,
	"ended":              OfferQuoteOfferQuoteStatusEnded,
	"expired":            OfferQuoteOfferQuoteStatusExpired,
	"failed_send":        OfferQuoteOfferQuoteStatusFailedSend,
	"failed_accept":      OfferQuoteOfferQuoteStatusFailedAccept,
	"failed_respond":     OfferQuoteOfferQuoteStatusFailedRespond,
}

// GetOfferQuoteOfferQuoteStatusEnumValues Enumerates the set of values for OfferQuoteOfferQuoteStatusEnum
func GetOfferQuoteOfferQuoteStatusEnumValues() []OfferQuoteOfferQuoteStatusEnum {
	values := make([]OfferQuoteOfferQuoteStatusEnum, 0)
	for _, v := range mappingOfferQuoteOfferQuoteStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetOfferQuoteOfferQuoteStatusEnumStringValues Enumerates the set of values in String for OfferQuoteOfferQuoteStatusEnum
func GetOfferQuoteOfferQuoteStatusEnumStringValues() []string {
	return []string{
		"DRAFT",
		"PENDING_ISV",
		"PENDING_ISV_UPDATE",
		"PENDING_RESELLER",
		"RESELLER_ACCEPTED",
		"BUYER_ACCEPTED",
		"ACTIVE",
		"ENDED",
		"EXPIRED",
		"FAILED_SEND",
		"FAILED_ACCEPT",
		"FAILED_RESPOND",
	}
}

// GetMappingOfferQuoteOfferQuoteStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingOfferQuoteOfferQuoteStatusEnum(val string) (OfferQuoteOfferQuoteStatusEnum, bool) {
	enum, ok := mappingOfferQuoteOfferQuoteStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
