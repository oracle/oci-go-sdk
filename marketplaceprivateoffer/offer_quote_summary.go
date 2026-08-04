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

// OfferQuoteSummary Summary of the offer quotes.
type OfferQuoteSummary struct {

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

	// A list of buyer tenancies.
	BuyerCompartmentIds []string `mandatory:"false" json:"buyerCompartmentIds"`

	// Display name of the reseller publisher.
	ResellerPublisherDisplayName *string `mandatory:"false" json:"resellerPublisherDisplayName"`

	// Display name of the ISV publisher.
	IsvPublisherDisplayName *string `mandatory:"false" json:"isvPublisherDisplayName"`

	// The time the offer quote was updated. An RFC3339 formatted datetime string.
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// Provides detailed information about the current state of the offer quote. For example, if the offer quote is in a failed state this message can include specific validation errors.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// The human readable representation of where the offer quote is at in its contract lifecycle.
	OfferQuoteStatus OfferQuoteOfferQuoteStatusEnum `mandatory:"false" json:"offerQuoteStatus,omitempty"`

	ResellerInformation *ResellerInformation `mandatory:"false" json:"resellerInformation"`

	IsvInformation *IsvInformation `mandatory:"false" json:"isvInformation"`

	Pricing *Pricing `mandatory:"false" json:"pricing"`

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

func (m OfferQuoteSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m OfferQuoteSummary) ValidateEnumValue() (bool, error) {
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
