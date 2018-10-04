// Copyright (c) 2016, 2018, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Core Services API
//
// APIs for Networking Service, Compute Service, and Block Volume Service.
//

package core

import (
	"github.com/oracle/oci-go-sdk/common"
)

// AppCatalogSubscriptionSummary a subscription summary for a listing resource version.
type AppCatalogSubscriptionSummary struct {

	// Name of the publisher who published this listing.
	PublisherName *string `mandatory:"false" json:"publisherName"`

	// The ocid of the listing resource.
	ListingId *string `mandatory:"false" json:"listingId"`

	// Listing resource version.
	ListingResourceVersion *string `mandatory:"false" json:"listingResourceVersion"`

	// Listing resource id.
	ListingResourceId *string `mandatory:"false" json:"listingResourceId"`

	// The display name of the listing.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// The short summary to the listing.
	Summary *string `mandatory:"false" json:"summary"`

	// The compartmentID of the subscription.
	CompartmentId *string `mandatory:"false" json:"compartmentId"`

	// Date and time at which the subscription was created, in RFC3339 format.
	// Example: `2018-03-20T12:32:53.532Z`
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`
}

func (m AppCatalogSubscriptionSummary) String() string {
	return common.PointerString(m)
}
