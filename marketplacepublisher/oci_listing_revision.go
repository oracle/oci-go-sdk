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

// OciListingRevision Listing revision details for listings
type OciListingRevision struct {

	// Unique OCID identifier for the listing revision in Marketplace Publisher.
	Id *string `mandatory:"true" json:"id"`

	// The unique identifier for the listing this revision belongs to.
	ListingId *string `mandatory:"true" json:"listingId"`

	// The name for the listing revision.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// Single line introduction for the listing revision.
	Headline *string `mandatory:"true" json:"headline"`

	// The time the listing revision was created. An RFC3339 formatted datetime string.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The time the listing revision was updated. An RFC3339 formatted datetime string.
	TimeUpdated *common.SDKTime `mandatory:"true" json:"timeUpdated"`

	// List of products subscribed by listing.
	Products []ListingProduct `mandatory:"true" json:"products"`

	// The unique identifier for the compartment.
	CompartmentId *string `mandatory:"false" json:"compartmentId"`

	// The revision number for the listing revision. This is an internal attribute
	RevisionNumber *string `mandatory:"false" json:"revisionNumber"`

	// The tagline of the listing revision.
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

	Icon *ListingRevisionIconAttachment `mandatory:"false" json:"icon"`

	// Status notes for the listing revision.
	StatusNotes *string `mandatory:"false" json:"statusNotes"`

	// Additional metadata key/value pairs for the listing revision summary.
	ExtendedMetadata map[string]string `mandatory:"false" json:"extendedMetadata"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// System tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`

	VersionDetails *VersionDetails `mandatory:"false" json:"versionDetails"`

	// System requirements for the listing revision.
	SystemRequirements *string `mandatory:"false" json:"systemRequirements"`

	// List of Pricing Plans provided by publisher.
	PricingPlans []PricingPlan `mandatory:"false" json:"pricingPlans"`

	// Listing availability and Pricing Policy statement.
	AvailabilityAndPricingPolicy *string `mandatory:"false" json:"availabilityAndPricingPolicy"`

	// Allowed tenancies provided when a listing revision is published as private.
	AllowedTenancies []string `mandatory:"false" json:"allowedTenancies"`

	// Custom link to the listing
	VanityUrl *string `mandatory:"false" json:"vanityUrl"`

	// OCIDs of service listings attached to lead gen listing
	RecommendedServiceProviderListingIds []string `mandatory:"false" json:"recommendedServiceProviderListingIds"`

	// Identifies whether publisher allows internal tenancy launches for the listing revision.
	AreInternalTenancyLaunchAllowed *bool `mandatory:"false" json:"areInternalTenancyLaunchAllowed"`

	// Is this listing rover exportable
	IsRoverExportable *bool `mandatory:"false" json:"isRoverExportable"`

	// The pricing model for the listing revision.
	PricingType OciListingRevisionPricingTypeEnum `mandatory:"true" json:"pricingType"`

	// The current status for the Listing revision.
	Status ListingRevisionStatusEnum `mandatory:"true" json:"status"`

	// The current state of the listing revision.
	LifecycleState ListingRevisionLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The listing's package type. Populated from the listing.
	PackageType PackageTypeEnum `mandatory:"false" json:"packageType,omitempty"`
}

// GetId returns Id
func (m OciListingRevision) GetId() *string {
	return m.Id
}

// GetListingId returns ListingId
func (m OciListingRevision) GetListingId() *string {
	return m.ListingId
}

// GetCompartmentId returns CompartmentId
func (m OciListingRevision) GetCompartmentId() *string {
	return m.CompartmentId
}

// GetDisplayName returns DisplayName
func (m OciListingRevision) GetDisplayName() *string {
	return m.DisplayName
}

// GetRevisionNumber returns RevisionNumber
func (m OciListingRevision) GetRevisionNumber() *string {
	return m.RevisionNumber
}

// GetHeadline returns Headline
func (m OciListingRevision) GetHeadline() *string {
	return m.Headline
}

// GetTagline returns Tagline
func (m OciListingRevision) GetTagline() *string {
	return m.Tagline
}

// GetKeywords returns Keywords
func (m OciListingRevision) GetKeywords() *string {
	return m.Keywords
}

// GetShortDescription returns ShortDescription
func (m OciListingRevision) GetShortDescription() *string {
	return m.ShortDescription
}

// GetUsageInformation returns UsageInformation
func (m OciListingRevision) GetUsageInformation() *string {
	return m.UsageInformation
}

// GetLongDescription returns LongDescription
func (m OciListingRevision) GetLongDescription() *string {
	return m.LongDescription
}

// GetTimeCreated returns TimeCreated
func (m OciListingRevision) GetTimeCreated() *common.SDKTime {
	return m.TimeCreated
}

// GetTimeUpdated returns TimeUpdated
func (m OciListingRevision) GetTimeUpdated() *common.SDKTime {
	return m.TimeUpdated
}

// GetContentLanguage returns ContentLanguage
func (m OciListingRevision) GetContentLanguage() *LanguageItem {
	return m.ContentLanguage
}

// GetSupportedlanguages returns Supportedlanguages
func (m OciListingRevision) GetSupportedlanguages() []LanguageItem {
	return m.Supportedlanguages
}

// GetSupportContacts returns SupportContacts
func (m OciListingRevision) GetSupportContacts() []SupportContact {
	return m.SupportContacts
}

// GetSupportLinks returns SupportLinks
func (m OciListingRevision) GetSupportLinks() []NamedLink {
	return m.SupportLinks
}

// GetIcon returns Icon
func (m OciListingRevision) GetIcon() *ListingRevisionIconAttachment {
	return m.Icon
}

// GetStatus returns Status
func (m OciListingRevision) GetStatus() ListingRevisionStatusEnum {
	return m.Status
}

// GetStatusNotes returns StatusNotes
func (m OciListingRevision) GetStatusNotes() *string {
	return m.StatusNotes
}

// GetLifecycleState returns LifecycleState
func (m OciListingRevision) GetLifecycleState() ListingRevisionLifecycleStateEnum {
	return m.LifecycleState
}

// GetPackageType returns PackageType
func (m OciListingRevision) GetPackageType() PackageTypeEnum {
	return m.PackageType
}

// GetExtendedMetadata returns ExtendedMetadata
func (m OciListingRevision) GetExtendedMetadata() map[string]string {
	return m.ExtendedMetadata
}

// GetFreeformTags returns FreeformTags
func (m OciListingRevision) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

// GetDefinedTags returns DefinedTags
func (m OciListingRevision) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

// GetSystemTags returns SystemTags
func (m OciListingRevision) GetSystemTags() map[string]map[string]interface{} {
	return m.SystemTags
}

func (m OciListingRevision) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m OciListingRevision) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingOciListingRevisionPricingTypeEnum(string(m.PricingType)); !ok && m.PricingType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for PricingType: %s. Supported values are: %s.", m.PricingType, strings.Join(GetOciListingRevisionPricingTypeEnumStringValues(), ",")))
	}

	if _, ok := GetMappingListingRevisionStatusEnum(string(m.Status)); !ok && m.Status != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", m.Status, strings.Join(GetListingRevisionStatusEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListingRevisionLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetListingRevisionLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingPackageTypeEnum(string(m.PackageType)); !ok && m.PackageType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for PackageType: %s. Supported values are: %s.", m.PackageType, strings.Join(GetPackageTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m OciListingRevision) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeOciListingRevision OciListingRevision
	s := struct {
		DiscriminatorParam string `json:"listingType"`
		MarshalTypeOciListingRevision
	}{
		"OCI_APPLICATION",
		(MarshalTypeOciListingRevision)(m),
	}

	return json.Marshal(&s)
}

// UnmarshalJSON unmarshals from json
func (m *OciListingRevision) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		CompartmentId                        *string                           `json:"compartmentId"`
		RevisionNumber                       *string                           `json:"revisionNumber"`
		Tagline                              *string                           `json:"tagline"`
		Keywords                             *string                           `json:"keywords"`
		ShortDescription                     *string                           `json:"shortDescription"`
		UsageInformation                     *string                           `json:"usageInformation"`
		LongDescription                      *string                           `json:"longDescription"`
		ContentLanguage                      *LanguageItem                     `json:"contentLanguage"`
		Supportedlanguages                   []LanguageItem                    `json:"supportedlanguages"`
		SupportContacts                      []SupportContact                  `json:"supportContacts"`
		SupportLinks                         []NamedLink                       `json:"supportLinks"`
		Icon                                 *ListingRevisionIconAttachment    `json:"icon"`
		StatusNotes                          *string                           `json:"statusNotes"`
		PackageType                          PackageTypeEnum                   `json:"packageType"`
		ExtendedMetadata                     map[string]string                 `json:"extendedMetadata"`
		FreeformTags                         map[string]string                 `json:"freeformTags"`
		DefinedTags                          map[string]map[string]interface{} `json:"definedTags"`
		SystemTags                           map[string]map[string]interface{} `json:"systemTags"`
		VersionDetails                       *VersionDetails                   `json:"versionDetails"`
		SystemRequirements                   *string                           `json:"systemRequirements"`
		PricingPlans                         []pricingplan                     `json:"pricingPlans"`
		AvailabilityAndPricingPolicy         *string                           `json:"availabilityAndPricingPolicy"`
		AllowedTenancies                     []string                          `json:"allowedTenancies"`
		VanityUrl                            *string                           `json:"vanityUrl"`
		RecommendedServiceProviderListingIds []string                          `json:"recommendedServiceProviderListingIds"`
		AreInternalTenancyLaunchAllowed      *bool                             `json:"areInternalTenancyLaunchAllowed"`
		IsRoverExportable                    *bool                             `json:"isRoverExportable"`
		Id                                   *string                           `json:"id"`
		ListingId                            *string                           `json:"listingId"`
		DisplayName                          *string                           `json:"displayName"`
		Headline                             *string                           `json:"headline"`
		TimeCreated                          *common.SDKTime                   `json:"timeCreated"`
		TimeUpdated                          *common.SDKTime                   `json:"timeUpdated"`
		Status                               ListingRevisionStatusEnum         `json:"status"`
		LifecycleState                       ListingRevisionLifecycleStateEnum `json:"lifecycleState"`
		PricingType                          OciListingRevisionPricingTypeEnum `json:"pricingType"`
		Products                             []ListingProduct                  `json:"products"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.CompartmentId = model.CompartmentId

	m.RevisionNumber = model.RevisionNumber

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
	m.Icon = model.Icon

	m.StatusNotes = model.StatusNotes

	m.PackageType = model.PackageType

	m.ExtendedMetadata = model.ExtendedMetadata

	m.FreeformTags = model.FreeformTags

	m.DefinedTags = model.DefinedTags

	m.SystemTags = model.SystemTags

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
	m.AvailabilityAndPricingPolicy = model.AvailabilityAndPricingPolicy

	m.AllowedTenancies = make([]string, len(model.AllowedTenancies))
	copy(m.AllowedTenancies, model.AllowedTenancies)
	m.VanityUrl = model.VanityUrl

	m.RecommendedServiceProviderListingIds = make([]string, len(model.RecommendedServiceProviderListingIds))
	copy(m.RecommendedServiceProviderListingIds, model.RecommendedServiceProviderListingIds)
	m.AreInternalTenancyLaunchAllowed = model.AreInternalTenancyLaunchAllowed

	m.IsRoverExportable = model.IsRoverExportable

	m.Id = model.Id

	m.ListingId = model.ListingId

	m.DisplayName = model.DisplayName

	m.Headline = model.Headline

	m.TimeCreated = model.TimeCreated

	m.TimeUpdated = model.TimeUpdated

	m.Status = model.Status

	m.LifecycleState = model.LifecycleState

	m.PricingType = model.PricingType

	m.Products = make([]ListingProduct, len(model.Products))
	copy(m.Products, model.Products)
	return
}

// OciListingRevisionPricingTypeEnum Enum with underlying type: string
type OciListingRevisionPricingTypeEnum string

// Set of constants representing the allowable values for OciListingRevisionPricingTypeEnum
const (
	OciListingRevisionPricingTypeFree  OciListingRevisionPricingTypeEnum = "FREE"
	OciListingRevisionPricingTypeByol  OciListingRevisionPricingTypeEnum = "BYOL"
	OciListingRevisionPricingTypePaygo OciListingRevisionPricingTypeEnum = "PAYGO"
)

var mappingOciListingRevisionPricingTypeEnum = map[string]OciListingRevisionPricingTypeEnum{
	"FREE":  OciListingRevisionPricingTypeFree,
	"BYOL":  OciListingRevisionPricingTypeByol,
	"PAYGO": OciListingRevisionPricingTypePaygo,
}

var mappingOciListingRevisionPricingTypeEnumLowerCase = map[string]OciListingRevisionPricingTypeEnum{
	"free":  OciListingRevisionPricingTypeFree,
	"byol":  OciListingRevisionPricingTypeByol,
	"paygo": OciListingRevisionPricingTypePaygo,
}

// GetOciListingRevisionPricingTypeEnumValues Enumerates the set of values for OciListingRevisionPricingTypeEnum
func GetOciListingRevisionPricingTypeEnumValues() []OciListingRevisionPricingTypeEnum {
	values := make([]OciListingRevisionPricingTypeEnum, 0)
	for _, v := range mappingOciListingRevisionPricingTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetOciListingRevisionPricingTypeEnumStringValues Enumerates the set of values in String for OciListingRevisionPricingTypeEnum
func GetOciListingRevisionPricingTypeEnumStringValues() []string {
	return []string{
		"FREE",
		"BYOL",
		"PAYGO",
	}
}

// GetMappingOciListingRevisionPricingTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingOciListingRevisionPricingTypeEnum(val string) (OciListingRevisionPricingTypeEnum, bool) {
	enum, ok := mappingOciListingRevisionPricingTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
