// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// MarketplacePublisherService API
//
// Use the Marketplace Publisher API to manage the publishing of applications in Oracle Cloud Infrastructure Marketplace.
//

package marketplacepublisher

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CreateOciListingRevisionDetails Listing revision update details for listings
type CreateOciListingRevisionDetails struct {

	// The unique identifier for the listing this revision belongs to.
	ListingId *string `mandatory:"true" json:"listingId"`

	// Single line introduction for the listing revision.
	Headline *string `mandatory:"true" json:"headline"`

	// List of Products subscribed by listing.
	Products []ListingProduct `mandatory:"true" json:"products"`

	// The name for the listing revision.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// The tagline for the listing revision.
	Tagline *string `mandatory:"false" json:"tagline"`

	// Keywords associated with the listing revision.
	Keywords *string `mandatory:"false" json:"keywords"`

	// A short description for the listing revision.
	ShortDescription *string `mandatory:"false" json:"shortDescription"`

	// Usage information for the listing revision.
	UsageInformation *string `mandatory:"false" json:"usageInformation"`

	// A long description for the listing revision.
	LongDescription *string `mandatory:"false" json:"longDescription"`

	ContentLanguage *LanguageItem `mandatory:"false" json:"contentLanguage"`

	// Languages supported by the publisher for the listing revision.
	Supportedlanguages []LanguageItem `mandatory:"false" json:"supportedlanguages"`

	// Contact information to use to get support from the publisher for the listing revision.
	SupportContacts []SupportContact `mandatory:"false" json:"supportContacts"`

	// Links to support resources for the listing revision.
	SupportLinks []NamedLink `mandatory:"false" json:"supportLinks"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	VersionDetails *VersionDetails `mandatory:"false" json:"versionDetails"`

	// System requirements for the listing revision.
	SystemRequirements *string `mandatory:"false" json:"systemRequirements"`

	// List of Pricing Plans provider by publisher.
	PricingPlans []PricingPlan `mandatory:"false" json:"pricingPlans"`

	// Custom link to the listing
	VanityUrl *string `mandatory:"false" json:"vanityUrl"`

	// OCIDs of service listings attached to lead gen listing
	RecommendedServiceProviderListingIds []string `mandatory:"false" json:"recommendedServiceProviderListingIds"`

	// Listing availability and Pricing Policy statement.
	AvailabilityAndPricingPolicy *string `mandatory:"false" json:"availabilityAndPricingPolicy"`

	// Is this listing rover exportable
	IsRoverExportable *bool `mandatory:"false" json:"isRoverExportable"`

	// The current status of the Listing revision.
	Status ListingRevisionStatusEnum `mandatory:"false" json:"status,omitempty"`

	// The pricing model for the listing revision.
	PricingType OciListingRevisionPricingTypeEnum `mandatory:"true" json:"pricingType"`
}

// GetDisplayName returns DisplayName
func (m CreateOciListingRevisionDetails) GetDisplayName() *string {
	return m.DisplayName
}

// GetListingId returns ListingId
func (m CreateOciListingRevisionDetails) GetListingId() *string {
	return m.ListingId
}

// GetHeadline returns Headline
func (m CreateOciListingRevisionDetails) GetHeadline() *string {
	return m.Headline
}

// GetTagline returns Tagline
func (m CreateOciListingRevisionDetails) GetTagline() *string {
	return m.Tagline
}

// GetKeywords returns Keywords
func (m CreateOciListingRevisionDetails) GetKeywords() *string {
	return m.Keywords
}

// GetShortDescription returns ShortDescription
func (m CreateOciListingRevisionDetails) GetShortDescription() *string {
	return m.ShortDescription
}

// GetUsageInformation returns UsageInformation
func (m CreateOciListingRevisionDetails) GetUsageInformation() *string {
	return m.UsageInformation
}

// GetLongDescription returns LongDescription
func (m CreateOciListingRevisionDetails) GetLongDescription() *string {
	return m.LongDescription
}

// GetContentLanguage returns ContentLanguage
func (m CreateOciListingRevisionDetails) GetContentLanguage() *LanguageItem {
	return m.ContentLanguage
}

// GetSupportedlanguages returns Supportedlanguages
func (m CreateOciListingRevisionDetails) GetSupportedlanguages() []LanguageItem {
	return m.Supportedlanguages
}

// GetSupportContacts returns SupportContacts
func (m CreateOciListingRevisionDetails) GetSupportContacts() []SupportContact {
	return m.SupportContacts
}

// GetSupportLinks returns SupportLinks
func (m CreateOciListingRevisionDetails) GetSupportLinks() []NamedLink {
	return m.SupportLinks
}

// GetStatus returns Status
func (m CreateOciListingRevisionDetails) GetStatus() ListingRevisionStatusEnum {
	return m.Status
}

// GetFreeformTags returns FreeformTags
func (m CreateOciListingRevisionDetails) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

// GetDefinedTags returns DefinedTags
func (m CreateOciListingRevisionDetails) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

func (m CreateOciListingRevisionDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateOciListingRevisionDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingListingRevisionStatusEnum(string(m.Status)); !ok && m.Status != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", m.Status, strings.Join(GetListingRevisionStatusEnumStringValues(), ",")))
	}
	if _, ok := GetMappingOciListingRevisionPricingTypeEnum(string(m.PricingType)); !ok && m.PricingType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for PricingType: %s. Supported values are: %s.", m.PricingType, strings.Join(GetOciListingRevisionPricingTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m CreateOciListingRevisionDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeCreateOciListingRevisionDetails CreateOciListingRevisionDetails
	s := struct {
		DiscriminatorParam string `json:"listingType"`
		MarshalTypeCreateOciListingRevisionDetails
	}{
		"OCI_APPLICATION",
		(MarshalTypeCreateOciListingRevisionDetails)(m),
	}

	return json.Marshal(&s)
}

// UnmarshalJSON unmarshals from json
func (m *CreateOciListingRevisionDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		DisplayName                          *string                           `json:"displayName"`
		Tagline                              *string                           `json:"tagline"`
		Keywords                             *string                           `json:"keywords"`
		ShortDescription                     *string                           `json:"shortDescription"`
		UsageInformation                     *string                           `json:"usageInformation"`
		LongDescription                      *string                           `json:"longDescription"`
		ContentLanguage                      *LanguageItem                     `json:"contentLanguage"`
		Supportedlanguages                   []LanguageItem                    `json:"supportedlanguages"`
		SupportContacts                      []SupportContact                  `json:"supportContacts"`
		SupportLinks                         []NamedLink                       `json:"supportLinks"`
		Status                               ListingRevisionStatusEnum         `json:"status"`
		FreeformTags                         map[string]string                 `json:"freeformTags"`
		DefinedTags                          map[string]map[string]interface{} `json:"definedTags"`
		VersionDetails                       *VersionDetails                   `json:"versionDetails"`
		SystemRequirements                   *string                           `json:"systemRequirements"`
		PricingPlans                         []pricingplan                     `json:"pricingPlans"`
		VanityUrl                            *string                           `json:"vanityUrl"`
		RecommendedServiceProviderListingIds []string                          `json:"recommendedServiceProviderListingIds"`
		AvailabilityAndPricingPolicy         *string                           `json:"availabilityAndPricingPolicy"`
		IsRoverExportable                    *bool                             `json:"isRoverExportable"`
		ListingId                            *string                           `json:"listingId"`
		Headline                             *string                           `json:"headline"`
		PricingType                          OciListingRevisionPricingTypeEnum `json:"pricingType"`
		Products                             []ListingProduct                  `json:"products"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.DisplayName = model.DisplayName

	m.Tagline = model.Tagline

	m.Keywords = model.Keywords

	m.ShortDescription = model.ShortDescription

	m.UsageInformation = model.UsageInformation

	m.LongDescription = model.LongDescription

	m.ContentLanguage = model.ContentLanguage

	m.Supportedlanguages = make([]LanguageItem, len(model.Supportedlanguages))
	copy(m.Supportedlanguages, model.Supportedlanguages)
	m.SupportContacts = make([]SupportContact, len(model.SupportContacts))
	copy(m.SupportContacts, model.SupportContacts)
	m.SupportLinks = make([]NamedLink, len(model.SupportLinks))
	copy(m.SupportLinks, model.SupportLinks)
	m.Status = model.Status

	m.FreeformTags = model.FreeformTags

	m.DefinedTags = model.DefinedTags

	m.VersionDetails = model.VersionDetails

	m.SystemRequirements = model.SystemRequirements

	m.PricingPlans = make([]PricingPlan, len(model.PricingPlans))
	for i, n := range model.PricingPlans {
		nn, e = n.UnmarshalPolymorphicJSON(n.JsonData)
		if e != nil {
			return e
		}
		if nn != nil {
			m.PricingPlans[i] = nn.(PricingPlan)
		} else {
			m.PricingPlans[i] = nil
		}
	}
	m.VanityUrl = model.VanityUrl

	m.RecommendedServiceProviderListingIds = make([]string, len(model.RecommendedServiceProviderListingIds))
	copy(m.RecommendedServiceProviderListingIds, model.RecommendedServiceProviderListingIds)
	m.AvailabilityAndPricingPolicy = model.AvailabilityAndPricingPolicy

	m.IsRoverExportable = model.IsRoverExportable

	m.ListingId = model.ListingId

	m.Headline = model.Headline

	m.PricingType = model.PricingType

	m.Products = make([]ListingProduct, len(model.Products))
	copy(m.Products, model.Products)
	return
}
