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

// AdminLeadGenListingRevision Listing revision details for lead gen listings
type AdminLeadGenListingRevision struct {

	// The OCID for the listing revision in Marketplace Publisher.
	Id *string `mandatory:"true" json:"id"`

	// The unique identifier for the listing this revision belongs to.
	ListingId *string `mandatory:"true" json:"listingId"`

	// The OCID for the publisher in Marketplace Publisher.
	PublisherId *string `mandatory:"true" json:"publisherId"`

	// The name for the listing revision.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// Single line introduction for the listing revision.
	Headline *string `mandatory:"true" json:"headline"`

	// The time the listing revision was created. An RFC3339 formatted datetime string.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The time the listing revision was updated. An RFC3339 formatted datetime string.
	TimeUpdated *common.SDKTime `mandatory:"true" json:"timeUpdated"`

	// The unique legacy identifier for the listing.
	LegacyId *string `mandatory:"false" json:"legacyId"`

	// The unique identifier for the compartment.
	CompartmentId *string `mandatory:"false" json:"compartmentId"`

	// The user OCID of the administrator who approved the listing revision.
	ApproverId *string `mandatory:"false" json:"approverId"`

	// The email address of the administrator who approved the listing revision.
	ApproverEmail *string `mandatory:"false" json:"approverEmail"`

	// The revision number of the listing revision. This is an internal attribute
	RevisionNumber *string `mandatory:"false" json:"revisionNumber"`

	// The tagline for the listing revision.
	Tagline *string `mandatory:"false" json:"tagline"`

	// Keywords associated with the listing revision.
	Keywords *string `mandatory:"false" json:"keywords"`

	// A short description of the listing revision.
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

	Banner *ListingRevisionBannerAttachment `mandatory:"false" json:"banner"`

	// Status notes for the listing revision.
	StatusNotes *string `mandatory:"false" json:"statusNotes"`

	// Additional metadata key/value pairs for the listing revision.
	// For example:
	// `{"oracleStandardTermVersionId": "1234"}`
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

	// Url to demo of the listing
	DemoUrl *string `mandatory:"false" json:"demoUrl"`

	// Url to training resources of the listing
	SelfPacedTrainingUrl *string `mandatory:"false" json:"selfPacedTrainingUrl"`

	// OCIDs of service listings attached to lead gen listing
	RecommendedServiceProviderListingIds []string `mandatory:"false" json:"recommendedServiceProviderListingIds"`

	// Custom link to the listing
	VanityUrl *string `mandatory:"false" json:"vanityUrl"`

	DownloadInfo *DownloadInfo `mandatory:"false" json:"downloadInfo"`

	// Pricing details for lead gen listing
	PricingPlans *string `mandatory:"false" json:"pricingPlans"`

	// List of Products subscribed by listing.
	Products []AdminListingProduct `mandatory:"false" json:"products"`

	// The current status for the listing revision.
	Status ListingRevisionStatusEnum `mandatory:"true" json:"status"`

	// The current state of the listing revision.
	LifecycleState ListingRevisionLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The listing's package type. Populated from the listing.
	PackageType PackageTypeEnum `mandatory:"true" json:"packageType"`

	// The pricing model for the listing revision.
	PricingType OciListingRevisionPricingTypeEnum `mandatory:"false" json:"pricingType,omitempty"`
}

// GetId returns Id
func (m AdminLeadGenListingRevision) GetId() *string {
	return m.Id
}

// GetLegacyId returns LegacyId
func (m AdminLeadGenListingRevision) GetLegacyId() *string {
	return m.LegacyId
}

// GetListingId returns ListingId
func (m AdminLeadGenListingRevision) GetListingId() *string {
	return m.ListingId
}

// GetCompartmentId returns CompartmentId
func (m AdminLeadGenListingRevision) GetCompartmentId() *string {
	return m.CompartmentId
}

// GetPublisherId returns PublisherId
func (m AdminLeadGenListingRevision) GetPublisherId() *string {
	return m.PublisherId
}

// GetApproverId returns ApproverId
func (m AdminLeadGenListingRevision) GetApproverId() *string {
	return m.ApproverId
}

// GetApproverEmail returns ApproverEmail
func (m AdminLeadGenListingRevision) GetApproverEmail() *string {
	return m.ApproverEmail
}

// GetDisplayName returns DisplayName
func (m AdminLeadGenListingRevision) GetDisplayName() *string {
	return m.DisplayName
}

// GetRevisionNumber returns RevisionNumber
func (m AdminLeadGenListingRevision) GetRevisionNumber() *string {
	return m.RevisionNumber
}

// GetHeadline returns Headline
func (m AdminLeadGenListingRevision) GetHeadline() *string {
	return m.Headline
}

// GetTagline returns Tagline
func (m AdminLeadGenListingRevision) GetTagline() *string {
	return m.Tagline
}

// GetKeywords returns Keywords
func (m AdminLeadGenListingRevision) GetKeywords() *string {
	return m.Keywords
}

// GetShortDescription returns ShortDescription
func (m AdminLeadGenListingRevision) GetShortDescription() *string {
	return m.ShortDescription
}

// GetUsageInformation returns UsageInformation
func (m AdminLeadGenListingRevision) GetUsageInformation() *string {
	return m.UsageInformation
}

// GetLongDescription returns LongDescription
func (m AdminLeadGenListingRevision) GetLongDescription() *string {
	return m.LongDescription
}

// GetTimeCreated returns TimeCreated
func (m AdminLeadGenListingRevision) GetTimeCreated() *common.SDKTime {
	return m.TimeCreated
}

// GetTimeUpdated returns TimeUpdated
func (m AdminLeadGenListingRevision) GetTimeUpdated() *common.SDKTime {
	return m.TimeUpdated
}

// GetContentLanguage returns ContentLanguage
func (m AdminLeadGenListingRevision) GetContentLanguage() *LanguageItem {
	return m.ContentLanguage
}

// GetSupportedlanguages returns Supportedlanguages
func (m AdminLeadGenListingRevision) GetSupportedlanguages() []LanguageItem {
	return m.Supportedlanguages
}

// GetSupportContacts returns SupportContacts
func (m AdminLeadGenListingRevision) GetSupportContacts() []SupportContact {
	return m.SupportContacts
}

// GetSupportLinks returns SupportLinks
func (m AdminLeadGenListingRevision) GetSupportLinks() []NamedLink {
	return m.SupportLinks
}

// GetIcon returns Icon
func (m AdminLeadGenListingRevision) GetIcon() *ListingRevisionIconAttachment {
	return m.Icon
}

// GetBanner returns Banner
func (m AdminLeadGenListingRevision) GetBanner() *ListingRevisionBannerAttachment {
	return m.Banner
}

// GetStatus returns Status
func (m AdminLeadGenListingRevision) GetStatus() ListingRevisionStatusEnum {
	return m.Status
}

// GetStatusNotes returns StatusNotes
func (m AdminLeadGenListingRevision) GetStatusNotes() *string {
	return m.StatusNotes
}

// GetLifecycleState returns LifecycleState
func (m AdminLeadGenListingRevision) GetLifecycleState() ListingRevisionLifecycleStateEnum {
	return m.LifecycleState
}

// GetPackageType returns PackageType
func (m AdminLeadGenListingRevision) GetPackageType() PackageTypeEnum {
	return m.PackageType
}

// GetExtendedMetadata returns ExtendedMetadata
func (m AdminLeadGenListingRevision) GetExtendedMetadata() map[string]string {
	return m.ExtendedMetadata
}

// GetFreeformTags returns FreeformTags
func (m AdminLeadGenListingRevision) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

// GetDefinedTags returns DefinedTags
func (m AdminLeadGenListingRevision) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

// GetSystemTags returns SystemTags
func (m AdminLeadGenListingRevision) GetSystemTags() map[string]map[string]interface{} {
	return m.SystemTags
}

func (m AdminLeadGenListingRevision) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m AdminLeadGenListingRevision) ValidateEnumValue() (bool, error) {
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
	if _, ok := GetMappingOciListingRevisionPricingTypeEnum(string(m.PricingType)); !ok && m.PricingType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for PricingType: %s. Supported values are: %s.", m.PricingType, strings.Join(GetOciListingRevisionPricingTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m AdminLeadGenListingRevision) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeAdminLeadGenListingRevision AdminLeadGenListingRevision
	s := struct {
		DiscriminatorParam string `json:"listingType"`
		MarshalTypeAdminLeadGenListingRevision
	}{
		"LEAD_GENERATION",
		(MarshalTypeAdminLeadGenListingRevision)(m),
	}

	return json.Marshal(&s)
}
