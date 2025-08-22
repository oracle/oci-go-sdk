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

// OfferSummary Summary of the Offers.
type OfferSummary struct {

	// Unique identifier that is immutable on creation
	Id *string `mandatory:"true" json:"id"`

	// Offer Identifier, can be renamed
	DisplayName *string `mandatory:"true" json:"displayName"`

	// Ocid of the buyer tenancy.
	BuyerCompartmentId *string `mandatory:"true" json:"buyerCompartmentId"`

	// Ocid of the seller tenancy.
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

	// The time the Offer was updated. An RFC3339 formatted datetime string
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// The time the Offer must be accepted by the Buyer before the Offer becomes invalid. An RFC3339 formatted datetime string
	TimeAcceptBy *common.SDKTime `mandatory:"false" json:"timeAcceptBy"`

	// The time the Offer was accepted by the Buyer of the Offer. An RFC3339 formatted datetime string
	TimeAccepted *common.SDKTime `mandatory:"false" json:"timeAccepted"`

	// The time the Offer will become active after it has been accepted by the Buyer. An RFC3339 formatted datetime string
	TimeStartDate *common.SDKTime `mandatory:"false" json:"timeStartDate"`

	// The time the accepted Offer will end. An RFC3339 formatted datetime string
	TimeOfferEnd *common.SDKTime `mandatory:"false" json:"timeOfferEnd"`

	// A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// The human readable representation of where the offer is at in it's contract lifecycle.
	OfferStatus OfferOfferStatusEnum `mandatory:"false" json:"offerStatus,omitempty"`

	BuyerInformation *BuyerInformation `mandatory:"false" json:"buyerInformation"`

	SellerInformation *SellerInformation `mandatory:"false" json:"sellerInformation"`

	Pricing *Pricing `mandatory:"false" json:"pricing"`
}

func (m OfferSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m OfferSummary) ValidateEnumValue() (bool, error) {
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
