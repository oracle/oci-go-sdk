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

// AvailableServiceSummary Available service provider listing for lead generation listings
type AvailableServiceSummary struct {

	// Listing Id for the listing
	ListingId *string `mandatory:"true" json:"listingId"`

	Icon *ListingRevisionIconAttachment `mandatory:"false" json:"icon"`

	// The name of the listing revision.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// A short description of the listing revision.
	ShortDescription *string `mandatory:"false" json:"shortDescription"`

	// The tagline of the listing revision.
	Tagline *string `mandatory:"false" json:"tagline"`
}

func (m AvailableServiceSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m AvailableServiceSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
