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

// CreateOfferQuoteDetails The information about new offer quotes.
type CreateOfferQuoteDetails struct {

	// Offer quotes identifier.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// Compartment identifier of the reseller.
	ResellerCompartmentId *string `mandatory:"true" json:"resellerCompartmentId"`

	// Compartment identifier of the ISV.
	IsvCompartmentId *string `mandatory:"false" json:"isvCompartmentId"`

	// Publisher identifier of the ISV.
	IsvPublisherId *string `mandatory:"false" json:"isvPublisherId"`

	// Description of the offer quote.
	Description *string `mandatory:"false" json:"description"`

	// A list of resource bundles associated with an offer quote.
	ResourceBundles []ResourceBundle `mandatory:"false" json:"resourceBundles"`

	// A list of buyer tenancies.
	BuyerCompartmentIds []string `mandatory:"false" json:"buyerCompartmentIds"`

	ResellerInformation *ResellerInformation `mandatory:"false" json:"resellerInformation"`

	IsvInformation *IsvInformation `mandatory:"false" json:"isvInformation"`

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
}

func (m CreateOfferQuoteDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateOfferQuoteDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
