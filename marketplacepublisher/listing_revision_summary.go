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

// ListingRevisionSummary The model for a summary of an Oracle Cloud Infrastructure Marketplace Publisher listing revision.
type ListingRevisionSummary struct {

	// The OCID for the listing revision in Marketplace Publisher.
	Id *string `mandatory:"true" json:"id"`

	// The OCID for the listing in Marketplace Publisher.
	ListingId *string `mandatory:"true" json:"listingId"`

	// The unique identifier for the compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The name of the listing revision.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The current status of the listing revision.
	Status ListingRevisionStatusEnum `mandatory:"true" json:"status"`

	// The current state of the Listing.
	LifecycleState ListingRevisionLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The listing's package type.
	PackageType PackageTypeEnum `mandatory:"true" json:"packageType"`

	// Pricing type of the listing.
	PricingType ListingRevisionPricingTypeEnum `mandatory:"true" json:"pricingType"`

	// Categories that the listing revision belongs to.
	Categories []string `mandatory:"true" json:"categories"`

	// The date and time the listing revision was created, in the format defined by RFC3339 (https://tools.ietf.org/html/rfc3339).
	// Example: `2022-09-15T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The date and time the listing revision was updated, in the format defined by RFC3339 (https://tools.ietf.org/html/rfc3339).
	// Example: `2022-09-15T21:10:29.600Z`
	TimeUpdated *common.SDKTime `mandatory:"true" json:"timeUpdated"`

	// A short description of the listing revision.
	ShortDescription *string `mandatory:"false" json:"shortDescription"`

	// The tagline of the listing revision.
	Tagline *string `mandatory:"false" json:"tagline"`

	Icon *ListingRevisionIconAttachment `mandatory:"false" json:"icon"`

	// The markets where you can deploy the listing.
	Markets []string `mandatory:"false" json:"markets"`

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

func (m ListingRevisionSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ListingRevisionSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListingRevisionStatusEnum(string(m.Status)); !ok && m.Status != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", m.Status, strings.Join(GetListingRevisionStatusEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListingRevisionLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetListingRevisionLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingPackageTypeEnum(string(m.PackageType)); !ok && m.PackageType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for PackageType: %s. Supported values are: %s.", m.PackageType, strings.Join(GetPackageTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListingRevisionPricingTypeEnum(string(m.PricingType)); !ok && m.PricingType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for PricingType: %s. Supported values are: %s.", m.PricingType, strings.Join(GetListingRevisionPricingTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
