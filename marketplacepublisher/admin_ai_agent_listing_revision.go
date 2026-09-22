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

// AdminAiAgentListingRevision Listing revision details for AI agent listings
type AdminAiAgentListingRevision struct {

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

	// Measurable or compelling value of the agent
	KeyBenefits *string `mandatory:"true" json:"keyBenefits"`

	// Capture asset competitive differentiator
	CompetitiveDifferentiators *string `mandatory:"true" json:"competitiveDifferentiators"`

	// Target audience for the asset
	TargetAudience *string `mandatory:"true" json:"targetAudience"`

	// List of industries subscribed by listing.
	Industries []string `mandatory:"true" json:"industries"`

	// Support details based on geographic location
	GeoLocations []GeoLocation `mandatory:"true" json:"geoLocations"`

	// List of products subscribed by listing.
	Products []AdminListingProduct `mandatory:"true" json:"products"`

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

	// Custom link to the listing
	VanityUrl *string `mandatory:"false" json:"vanityUrl"`

	// Agent asset type
	AssetType AdminAiAgentListingRevisionAssetTypeEnum `mandatory:"true" json:"assetType"`

	// Work area that surfaces the AI agent
	Location AdminAiAgentListingRevisionLocationEnum `mandatory:"true" json:"location"`

	// The current status for the listing revision.
	Status ListingRevisionStatusEnum `mandatory:"true" json:"status"`

	// The current state of the listing revision.
	LifecycleState ListingRevisionLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The listing's package type. Populated from the listing.
	PackageType PackageTypeEnum `mandatory:"true" json:"packageType"`
}

// GetId returns Id
func (m AdminAiAgentListingRevision) GetId() *string {
	return m.Id
}

// GetLegacyId returns LegacyId
func (m AdminAiAgentListingRevision) GetLegacyId() *string {
	return m.LegacyId
}

// GetListingId returns ListingId
func (m AdminAiAgentListingRevision) GetListingId() *string {
	return m.ListingId
}

// GetCompartmentId returns CompartmentId
func (m AdminAiAgentListingRevision) GetCompartmentId() *string {
	return m.CompartmentId
}

// GetPublisherId returns PublisherId
func (m AdminAiAgentListingRevision) GetPublisherId() *string {
	return m.PublisherId
}

// GetApproverId returns ApproverId
func (m AdminAiAgentListingRevision) GetApproverId() *string {
	return m.ApproverId
}

// GetApproverEmail returns ApproverEmail
func (m AdminAiAgentListingRevision) GetApproverEmail() *string {
	return m.ApproverEmail
}

// GetDisplayName returns DisplayName
func (m AdminAiAgentListingRevision) GetDisplayName() *string {
	return m.DisplayName
}

// GetRevisionNumber returns RevisionNumber
func (m AdminAiAgentListingRevision) GetRevisionNumber() *string {
	return m.RevisionNumber
}

// GetHeadline returns Headline
func (m AdminAiAgentListingRevision) GetHeadline() *string {
	return m.Headline
}

// GetTagline returns Tagline
func (m AdminAiAgentListingRevision) GetTagline() *string {
	return m.Tagline
}

// GetKeywords returns Keywords
func (m AdminAiAgentListingRevision) GetKeywords() *string {
	return m.Keywords
}

// GetShortDescription returns ShortDescription
func (m AdminAiAgentListingRevision) GetShortDescription() *string {
	return m.ShortDescription
}

// GetUsageInformation returns UsageInformation
func (m AdminAiAgentListingRevision) GetUsageInformation() *string {
	return m.UsageInformation
}

// GetLongDescription returns LongDescription
func (m AdminAiAgentListingRevision) GetLongDescription() *string {
	return m.LongDescription
}

// GetTimeCreated returns TimeCreated
func (m AdminAiAgentListingRevision) GetTimeCreated() *common.SDKTime {
	return m.TimeCreated
}

// GetTimeUpdated returns TimeUpdated
func (m AdminAiAgentListingRevision) GetTimeUpdated() *common.SDKTime {
	return m.TimeUpdated
}

// GetContentLanguage returns ContentLanguage
func (m AdminAiAgentListingRevision) GetContentLanguage() *LanguageItem {
	return m.ContentLanguage
}

// GetSupportedlanguages returns Supportedlanguages
func (m AdminAiAgentListingRevision) GetSupportedlanguages() []LanguageItem {
	return m.Supportedlanguages
}

// GetSupportContacts returns SupportContacts
func (m AdminAiAgentListingRevision) GetSupportContacts() []SupportContact {
	return m.SupportContacts
}

// GetSupportLinks returns SupportLinks
func (m AdminAiAgentListingRevision) GetSupportLinks() []NamedLink {
	return m.SupportLinks
}

// GetIcon returns Icon
func (m AdminAiAgentListingRevision) GetIcon() *ListingRevisionIconAttachment {
	return m.Icon
}

// GetBanner returns Banner
func (m AdminAiAgentListingRevision) GetBanner() *ListingRevisionBannerAttachment {
	return m.Banner
}

// GetStatus returns Status
func (m AdminAiAgentListingRevision) GetStatus() ListingRevisionStatusEnum {
	return m.Status
}

// GetStatusNotes returns StatusNotes
func (m AdminAiAgentListingRevision) GetStatusNotes() *string {
	return m.StatusNotes
}

// GetLifecycleState returns LifecycleState
func (m AdminAiAgentListingRevision) GetLifecycleState() ListingRevisionLifecycleStateEnum {
	return m.LifecycleState
}

// GetPackageType returns PackageType
func (m AdminAiAgentListingRevision) GetPackageType() PackageTypeEnum {
	return m.PackageType
}

// GetExtendedMetadata returns ExtendedMetadata
func (m AdminAiAgentListingRevision) GetExtendedMetadata() map[string]string {
	return m.ExtendedMetadata
}

// GetFreeformTags returns FreeformTags
func (m AdminAiAgentListingRevision) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

// GetDefinedTags returns DefinedTags
func (m AdminAiAgentListingRevision) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

// GetSystemTags returns SystemTags
func (m AdminAiAgentListingRevision) GetSystemTags() map[string]map[string]interface{} {
	return m.SystemTags
}

func (m AdminAiAgentListingRevision) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m AdminAiAgentListingRevision) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingAdminAiAgentListingRevisionAssetTypeEnum(string(m.AssetType)); !ok && m.AssetType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for AssetType: %s. Supported values are: %s.", m.AssetType, strings.Join(GetAdminAiAgentListingRevisionAssetTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingAdminAiAgentListingRevisionLocationEnum(string(m.Location)); !ok && m.Location != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Location: %s. Supported values are: %s.", m.Location, strings.Join(GetAdminAiAgentListingRevisionLocationEnumStringValues(), ",")))
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
func (m AdminAiAgentListingRevision) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeAdminAiAgentListingRevision AdminAiAgentListingRevision
	s := struct {
		DiscriminatorParam string `json:"listingType"`
		MarshalTypeAdminAiAgentListingRevision
	}{
		"AI_AGENT",
		(MarshalTypeAdminAiAgentListingRevision)(m),
	}

	return json.Marshal(&s)
}

// AdminAiAgentListingRevisionAssetTypeEnum Enum with underlying type: string
type AdminAiAgentListingRevisionAssetTypeEnum string

// Set of constants representing the allowable values for AdminAiAgentListingRevisionAssetTypeEnum
const (
	AdminAiAgentListingRevisionAssetTypeFoundational AdminAiAgentListingRevisionAssetTypeEnum = "FOUNDATIONAL"
	AdminAiAgentListingRevisionAssetTypeTeam         AdminAiAgentListingRevisionAssetTypeEnum = "TEAM"
	AdminAiAgentListingRevisionAssetTypeWorkflow     AdminAiAgentListingRevisionAssetTypeEnum = "WORKFLOW"
	AdminAiAgentListingRevisionAssetTypeApplication  AdminAiAgentListingRevisionAssetTypeEnum = "APPLICATION"
	AdminAiAgentListingRevisionAssetTypeConnector    AdminAiAgentListingRevisionAssetTypeEnum = "CONNECTOR"
)

var mappingAdminAiAgentListingRevisionAssetTypeEnum = map[string]AdminAiAgentListingRevisionAssetTypeEnum{
	"FOUNDATIONAL": AdminAiAgentListingRevisionAssetTypeFoundational,
	"TEAM":         AdminAiAgentListingRevisionAssetTypeTeam,
	"WORKFLOW":     AdminAiAgentListingRevisionAssetTypeWorkflow,
	"APPLICATION":  AdminAiAgentListingRevisionAssetTypeApplication,
	"CONNECTOR":    AdminAiAgentListingRevisionAssetTypeConnector,
}

var mappingAdminAiAgentListingRevisionAssetTypeEnumLowerCase = map[string]AdminAiAgentListingRevisionAssetTypeEnum{
	"foundational": AdminAiAgentListingRevisionAssetTypeFoundational,
	"team":         AdminAiAgentListingRevisionAssetTypeTeam,
	"workflow":     AdminAiAgentListingRevisionAssetTypeWorkflow,
	"application":  AdminAiAgentListingRevisionAssetTypeApplication,
	"connector":    AdminAiAgentListingRevisionAssetTypeConnector,
}

// GetAdminAiAgentListingRevisionAssetTypeEnumValues Enumerates the set of values for AdminAiAgentListingRevisionAssetTypeEnum
func GetAdminAiAgentListingRevisionAssetTypeEnumValues() []AdminAiAgentListingRevisionAssetTypeEnum {
	values := make([]AdminAiAgentListingRevisionAssetTypeEnum, 0)
	for _, v := range mappingAdminAiAgentListingRevisionAssetTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetAdminAiAgentListingRevisionAssetTypeEnumStringValues Enumerates the set of values in String for AdminAiAgentListingRevisionAssetTypeEnum
func GetAdminAiAgentListingRevisionAssetTypeEnumStringValues() []string {
	return []string{
		"FOUNDATIONAL",
		"TEAM",
		"WORKFLOW",
		"APPLICATION",
		"CONNECTOR",
	}
}

// GetMappingAdminAiAgentListingRevisionAssetTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAdminAiAgentListingRevisionAssetTypeEnum(val string) (AdminAiAgentListingRevisionAssetTypeEnum, bool) {
	enum, ok := mappingAdminAiAgentListingRevisionAssetTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// AdminAiAgentListingRevisionLocationEnum Enum with underlying type: string
type AdminAiAgentListingRevisionLocationEnum string

// Set of constants representing the allowable values for AdminAiAgentListingRevisionLocationEnum
const (
	AdminAiAgentListingRevisionLocationSeeded  AdminAiAgentListingRevisionLocationEnum = "SEEDED"
	AdminAiAgentListingRevisionLocationLeadGen AdminAiAgentListingRevisionLocationEnum = "LEAD_GEN"
)

var mappingAdminAiAgentListingRevisionLocationEnum = map[string]AdminAiAgentListingRevisionLocationEnum{
	"SEEDED":   AdminAiAgentListingRevisionLocationSeeded,
	"LEAD_GEN": AdminAiAgentListingRevisionLocationLeadGen,
}

var mappingAdminAiAgentListingRevisionLocationEnumLowerCase = map[string]AdminAiAgentListingRevisionLocationEnum{
	"seeded":   AdminAiAgentListingRevisionLocationSeeded,
	"lead_gen": AdminAiAgentListingRevisionLocationLeadGen,
}

// GetAdminAiAgentListingRevisionLocationEnumValues Enumerates the set of values for AdminAiAgentListingRevisionLocationEnum
func GetAdminAiAgentListingRevisionLocationEnumValues() []AdminAiAgentListingRevisionLocationEnum {
	values := make([]AdminAiAgentListingRevisionLocationEnum, 0)
	for _, v := range mappingAdminAiAgentListingRevisionLocationEnum {
		values = append(values, v)
	}
	return values
}

// GetAdminAiAgentListingRevisionLocationEnumStringValues Enumerates the set of values in String for AdminAiAgentListingRevisionLocationEnum
func GetAdminAiAgentListingRevisionLocationEnumStringValues() []string {
	return []string{
		"SEEDED",
		"LEAD_GEN",
	}
}

// GetMappingAdminAiAgentListingRevisionLocationEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAdminAiAgentListingRevisionLocationEnum(val string) (AdminAiAgentListingRevisionLocationEnum, bool) {
	enum, ok := mappingAdminAiAgentListingRevisionLocationEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
