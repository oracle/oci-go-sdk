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

// CreateOfferDetails The information about new Offers.
type CreateOfferDetails struct {

	// Offers Identifier
	DisplayName *string `mandatory:"true" json:"displayName"`

	// Compartment Identifier of the seller
	SellerCompartmentId *string `mandatory:"true" json:"sellerCompartmentId"`

	// Compartment Identifier of the buyer
	BuyerCompartmentId *string `mandatory:"false" json:"buyerCompartmentId"`

	// Description of the Offer
	Description *string `mandatory:"false" json:"description"`

	// Internal notes of the Offer
	InternalNotes *string `mandatory:"false" json:"internalNotes"`

	// The time the Offer will become active after it has been accepted by the Buyer. An RFC3339 formatted datetime string
	TimeStartDate *common.SDKTime `mandatory:"false" json:"timeStartDate"`

	// Duration the Offer will be active after its start date. An ISO8601 extended formatted string.
	Duration *string `mandatory:"false" json:"duration"`

	// The time the Offer must be accepted by the Buyer before the Offer becomes invalid. An RFC3339 formatted datetime string
	TimeAcceptBy *common.SDKTime `mandatory:"false" json:"timeAcceptBy"`

	Pricing *Pricing `mandatory:"false" json:"pricing"`

	BuyerInformation *BuyerInformation `mandatory:"false" json:"buyerInformation"`

	SellerInformation *SellerInformation `mandatory:"false" json:"sellerInformation"`

	// A list of Resource Bundles associated with an Offer.
	ResourceBundles []ResourceBundle `mandatory:"false" json:"resourceBundles"`

	// A list of key value pairs specified by the seller
	CustomFields []CustomField `mandatory:"false" json:"customFields"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m CreateOfferDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateOfferDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
