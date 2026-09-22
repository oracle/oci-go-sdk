// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
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

// AdminListingRevisionSummary The model for a summary of an Oracle Cloud Infrastructure Marketplace Publisher listing revision.
type AdminListingRevisionSummary struct {

	// The OCID for the listing revision in Marketplace Publisher.
	Id *string `mandatory:"true" json:"id"`

	// The OCID for the Listing in Marketplace Publisher.
	ListingId *string `mandatory:"true" json:"listingId"`

	// The unique identifier for the compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The OCID for the publisher in Marketplace Publisher.
	PublisherId *string `mandatory:"true" json:"publisherId"`

	// The display name for the listing revision.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The current status for the listing revision.
	Status ListingRevisionStatusEnum `mandatory:"true" json:"status"`

	// The current state for the listing revision.
	LifecycleState ListingRevisionLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The listing's package type.
	PackageType PackageTypeEnum `mandatory:"true" json:"packageType"`

	// The listing's type. Populated from the listing.
	ListingType ListingTypeEnum `mandatory:"true" json:"listingType"`

	// The date and time the listing revision was created, in the format defined by RFC3339 (https://tools.ietf.org/html/rfc3339).
	// Example: `2022-09-15T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The date and time the listing revision was updated, in the format defined by RFC3339 (https://tools.ietf.org/html/rfc3339).
	// Example: `2022-09-15T21:10:29.600Z`
	TimeUpdated *common.SDKTime `mandatory:"true" json:"timeUpdated"`

	// The unique legacy identifier for the listing.
	LegacyId *string `mandatory:"false" json:"legacyId"`

	// A short description for the listing revision.
	ShortDescription *string `mandatory:"false" json:"shortDescription"`

	// A long description for the listing revision.
	LongDescription *string `mandatory:"false" json:"longDescription"`

	// List of Pricing Plans provided by publisher.
	PricingPlans []PricingPlan `mandatory:"false" json:"pricingPlans"`

	// Languages supported by the publisher for the listing revision.
	Supportedlanguages []LanguageItem `mandatory:"false" json:"supportedlanguages"`

	// Links to support resources for the listing revision.
	SupportLinks []NamedLink `mandatory:"false" json:"supportLinks"`

	// System requirements for the listing revision.
	SystemRequirements *string `mandatory:"false" json:"systemRequirements"`

	VersionDetails *VersionDetails `mandatory:"false" json:"versionDetails"`

	// Number of trained professional per product
	TrainedProfessionals []TrainedProfessionals `mandatory:"false" json:"trainedProfessionals"`

	// Support details based on geographic location
	GeoLocations []GeoLocation `mandatory:"false" json:"geoLocations"`

	// The pricing model for the listing revision.
	PricingType AdminListingRevisionSummaryPricingTypeEnum `mandatory:"false" json:"pricingType,omitempty"`

	// Markets associated with Listing revision.
	Markets []string `mandatory:"false" json:"markets"`

	// The tagline for the listing revision.
	Tagline *string `mandatory:"false" json:"tagline"`

	// List of Admin Products subscribed by listing.
	Products []AdminListingProduct `mandatory:"false" json:"products"`

	Icon *ListingRevisionIconAttachment `mandatory:"false" json:"icon"`

	Banner *ListingRevisionBannerAttachment `mandatory:"false" json:"banner"`

	// Url to demo of the listing
	DemoUrl *string `mandatory:"false" json:"demoUrl"`

	// List of industries subscribed by listing.
	Industries []string `mandatory:"false" json:"industries"`

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

func (m AdminListingRevisionSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m AdminListingRevisionSummary) ValidateEnumValue() (bool, error) {
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
	if _, ok := GetMappingListingTypeEnum(string(m.ListingType)); !ok && m.ListingType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ListingType: %s. Supported values are: %s.", m.ListingType, strings.Join(GetListingTypeEnumStringValues(), ",")))
	}

	if _, ok := GetMappingAdminListingRevisionSummaryPricingTypeEnum(string(m.PricingType)); !ok && m.PricingType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for PricingType: %s. Supported values are: %s.", m.PricingType, strings.Join(GetAdminListingRevisionSummaryPricingTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *AdminListingRevisionSummary) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		LegacyId             *string                                    `json:"legacyId"`
		ShortDescription     *string                                    `json:"shortDescription"`
		LongDescription      *string                                    `json:"longDescription"`
		PricingPlans         []pricingplan                              `json:"pricingPlans"`
		Supportedlanguages   []LanguageItem                             `json:"supportedlanguages"`
		SupportLinks         []NamedLink                                `json:"supportLinks"`
		SystemRequirements   *string                                    `json:"systemRequirements"`
		VersionDetails       *VersionDetails                            `json:"versionDetails"`
		TrainedProfessionals []TrainedProfessionals                     `json:"trainedProfessionals"`
		GeoLocations         []GeoLocation                              `json:"geoLocations"`
		PricingType          AdminListingRevisionSummaryPricingTypeEnum `json:"pricingType"`
		Markets              []string                                   `json:"markets"`
		Tagline              *string                                    `json:"tagline"`
		Products             []AdminListingProduct                      `json:"products"`
		Icon                 *ListingRevisionIconAttachment             `json:"icon"`
		Banner               *ListingRevisionBannerAttachment           `json:"banner"`
		DemoUrl              *string                                    `json:"demoUrl"`
		Industries           []string                                   `json:"industries"`
		FreeformTags         map[string]string                          `json:"freeformTags"`
		DefinedTags          map[string]map[string]interface{}          `json:"definedTags"`
		SystemTags           map[string]map[string]interface{}          `json:"systemTags"`
		Id                   *string                                    `json:"id"`
		ListingId            *string                                    `json:"listingId"`
		CompartmentId        *string                                    `json:"compartmentId"`
		PublisherId          *string                                    `json:"publisherId"`
		DisplayName          *string                                    `json:"displayName"`
		Status               ListingRevisionStatusEnum                  `json:"status"`
		LifecycleState       ListingRevisionLifecycleStateEnum          `json:"lifecycleState"`
		PackageType          PackageTypeEnum                            `json:"packageType"`
		ListingType          ListingTypeEnum                            `json:"listingType"`
		TimeCreated          *common.SDKTime                            `json:"timeCreated"`
		TimeUpdated          *common.SDKTime                            `json:"timeUpdated"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.LegacyId = model.LegacyId

	m.ShortDescription = model.ShortDescription

	m.LongDescription = model.LongDescription

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
	m.Supportedlanguages = make([]LanguageItem, len(model.Supportedlanguages))
	copy(m.Supportedlanguages, model.Supportedlanguages)
	m.SupportLinks = make([]NamedLink, len(model.SupportLinks))
	copy(m.SupportLinks, model.SupportLinks)
	m.SystemRequirements = model.SystemRequirements

	m.VersionDetails = model.VersionDetails

	m.TrainedProfessionals = make([]TrainedProfessionals, len(model.TrainedProfessionals))
	copy(m.TrainedProfessionals, model.TrainedProfessionals)
	m.GeoLocations = make([]GeoLocation, len(model.GeoLocations))
	copy(m.GeoLocations, model.GeoLocations)
	m.PricingType = model.PricingType

	m.Markets = make([]string, len(model.Markets))
	copy(m.Markets, model.Markets)
	m.Tagline = model.Tagline

	m.Products = make([]AdminListingProduct, len(model.Products))
	copy(m.Products, model.Products)
	m.Icon = model.Icon

	m.Banner = model.Banner

	m.DemoUrl = model.DemoUrl

	m.Industries = make([]string, len(model.Industries))
	copy(m.Industries, model.Industries)
	m.FreeformTags = model.FreeformTags

	m.DefinedTags = model.DefinedTags

	m.SystemTags = model.SystemTags

	m.Id = model.Id

	m.ListingId = model.ListingId

	m.CompartmentId = model.CompartmentId

	m.PublisherId = model.PublisherId

	m.DisplayName = model.DisplayName

	m.Status = model.Status

	m.LifecycleState = model.LifecycleState

	m.PackageType = model.PackageType

	m.ListingType = model.ListingType

	m.TimeCreated = model.TimeCreated

	m.TimeUpdated = model.TimeUpdated

	return
}

// AdminListingRevisionSummaryPricingTypeEnum Enum with underlying type: string
type AdminListingRevisionSummaryPricingTypeEnum string

// Set of constants representing the allowable values for AdminListingRevisionSummaryPricingTypeEnum
const (
	AdminListingRevisionSummaryPricingTypeFree  AdminListingRevisionSummaryPricingTypeEnum = "FREE"
	AdminListingRevisionSummaryPricingTypeByol  AdminListingRevisionSummaryPricingTypeEnum = "BYOL"
	AdminListingRevisionSummaryPricingTypePaygo AdminListingRevisionSummaryPricingTypeEnum = "PAYGO"
)

var mappingAdminListingRevisionSummaryPricingTypeEnum = map[string]AdminListingRevisionSummaryPricingTypeEnum{
	"FREE":  AdminListingRevisionSummaryPricingTypeFree,
	"BYOL":  AdminListingRevisionSummaryPricingTypeByol,
	"PAYGO": AdminListingRevisionSummaryPricingTypePaygo,
}

var mappingAdminListingRevisionSummaryPricingTypeEnumLowerCase = map[string]AdminListingRevisionSummaryPricingTypeEnum{
	"free":  AdminListingRevisionSummaryPricingTypeFree,
	"byol":  AdminListingRevisionSummaryPricingTypeByol,
	"paygo": AdminListingRevisionSummaryPricingTypePaygo,
}

// GetAdminListingRevisionSummaryPricingTypeEnumValues Enumerates the set of values for AdminListingRevisionSummaryPricingTypeEnum
func GetAdminListingRevisionSummaryPricingTypeEnumValues() []AdminListingRevisionSummaryPricingTypeEnum {
	values := make([]AdminListingRevisionSummaryPricingTypeEnum, 0)
	for _, v := range mappingAdminListingRevisionSummaryPricingTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetAdminListingRevisionSummaryPricingTypeEnumStringValues Enumerates the set of values in String for AdminListingRevisionSummaryPricingTypeEnum
func GetAdminListingRevisionSummaryPricingTypeEnumStringValues() []string {
	return []string{
		"FREE",
		"BYOL",
		"PAYGO",
	}
}

// GetMappingAdminListingRevisionSummaryPricingTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAdminListingRevisionSummaryPricingTypeEnum(val string) (AdminListingRevisionSummaryPricingTypeEnum, bool) {
	enum, ok := mappingAdminListingRevisionSummaryPricingTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
