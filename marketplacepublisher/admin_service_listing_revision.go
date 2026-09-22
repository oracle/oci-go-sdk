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

// AdminServiceListingRevision Listing revision details for service listings
type AdminServiceListingRevision struct {

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

	// Custom text by partner regarding contact information
	ContactUs *string `mandatory:"false" json:"contactUs"`

	// List of product code subscribed by listing.
	ProductCodes []string `mandatory:"false" json:"productCodes"`

	// List of industries subscribed by listing.
	Industries []string `mandatory:"false" json:"industries"`

	// Number of trained professional per product
	TrainedProfessionals []TrainedProfessionals `mandatory:"false" json:"trainedProfessionals"`

	// Custom link to the listing
	VanityUrl *string `mandatory:"false" json:"vanityUrl"`

	// Support details based on geographic location
	GeoLocations []GeoLocation `mandatory:"false" json:"geoLocations"`

	// List of Products subscribed by listing.
	Products []AdminListingProduct `mandatory:"false" json:"products"`

	// The current status for the listing revision.
	Status ListingRevisionStatusEnum `mandatory:"true" json:"status"`

	// The current state of the listing revision.
	LifecycleState ListingRevisionLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The listing's package type. Populated from the listing.
	PackageType PackageTypeEnum `mandatory:"true" json:"packageType"`
}

// GetId returns Id
func (m AdminServiceListingRevision) GetId() *string {
	return m.Id
}

// GetLegacyId returns LegacyId
func (m AdminServiceListingRevision) GetLegacyId() *string {
	return m.LegacyId
}

// GetListingId returns ListingId
func (m AdminServiceListingRevision) GetListingId() *string {
	return m.ListingId
}

// GetCompartmentId returns CompartmentId
func (m AdminServiceListingRevision) GetCompartmentId() *string {
	return m.CompartmentId
}

// GetPublisherId returns PublisherId
func (m AdminServiceListingRevision) GetPublisherId() *string {
	return m.PublisherId
}

// GetApproverId returns ApproverId
func (m AdminServiceListingRevision) GetApproverId() *string {
	return m.ApproverId
}

// GetApproverEmail returns ApproverEmail
func (m AdminServiceListingRevision) GetApproverEmail() *string {
	return m.ApproverEmail
}

// GetDisplayName returns DisplayName
func (m AdminServiceListingRevision) GetDisplayName() *string {
	return m.DisplayName
}

// GetRevisionNumber returns RevisionNumber
func (m AdminServiceListingRevision) GetRevisionNumber() *string {
	return m.RevisionNumber
}

// GetHeadline returns Headline
func (m AdminServiceListingRevision) GetHeadline() *string {
	return m.Headline
}

// GetTagline returns Tagline
func (m AdminServiceListingRevision) GetTagline() *string {
	return m.Tagline
}

// GetKeywords returns Keywords
func (m AdminServiceListingRevision) GetKeywords() *string {
	return m.Keywords
}

// GetShortDescription returns ShortDescription
func (m AdminServiceListingRevision) GetShortDescription() *string {
	return m.ShortDescription
}

// GetUsageInformation returns UsageInformation
func (m AdminServiceListingRevision) GetUsageInformation() *string {
	return m.UsageInformation
}

// GetLongDescription returns LongDescription
func (m AdminServiceListingRevision) GetLongDescription() *string {
	return m.LongDescription
}

// GetTimeCreated returns TimeCreated
func (m AdminServiceListingRevision) GetTimeCreated() *common.SDKTime {
	return m.TimeCreated
}

// GetTimeUpdated returns TimeUpdated
func (m AdminServiceListingRevision) GetTimeUpdated() *common.SDKTime {
	return m.TimeUpdated
}

// GetContentLanguage returns ContentLanguage
func (m AdminServiceListingRevision) GetContentLanguage() *LanguageItem {
	return m.ContentLanguage
}

// GetSupportedlanguages returns Supportedlanguages
func (m AdminServiceListingRevision) GetSupportedlanguages() []LanguageItem {
	return m.Supportedlanguages
}

// GetSupportContacts returns SupportContacts
func (m AdminServiceListingRevision) GetSupportContacts() []SupportContact {
	return m.SupportContacts
}

// GetSupportLinks returns SupportLinks
func (m AdminServiceListingRevision) GetSupportLinks() []NamedLink {
	return m.SupportLinks
}

// GetIcon returns Icon
func (m AdminServiceListingRevision) GetIcon() *ListingRevisionIconAttachment {
	return m.Icon
}

// GetBanner returns Banner
func (m AdminServiceListingRevision) GetBanner() *ListingRevisionBannerAttachment {
	return m.Banner
}

// GetStatus returns Status
func (m AdminServiceListingRevision) GetStatus() ListingRevisionStatusEnum {
	return m.Status
}

// GetStatusNotes returns StatusNotes
func (m AdminServiceListingRevision) GetStatusNotes() *string {
	return m.StatusNotes
}

// GetLifecycleState returns LifecycleState
func (m AdminServiceListingRevision) GetLifecycleState() ListingRevisionLifecycleStateEnum {
	return m.LifecycleState
}

// GetPackageType returns PackageType
func (m AdminServiceListingRevision) GetPackageType() PackageTypeEnum {
	return m.PackageType
}

// GetExtendedMetadata returns ExtendedMetadata
func (m AdminServiceListingRevision) GetExtendedMetadata() map[string]string {
	return m.ExtendedMetadata
}

// GetFreeformTags returns FreeformTags
func (m AdminServiceListingRevision) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

// GetDefinedTags returns DefinedTags
func (m AdminServiceListingRevision) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

// GetSystemTags returns SystemTags
func (m AdminServiceListingRevision) GetSystemTags() map[string]map[string]interface{} {
	return m.SystemTags
}

func (m AdminServiceListingRevision) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m AdminServiceListingRevision) ValidateEnumValue() (bool, error) {
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
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m AdminServiceListingRevision) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeAdminServiceListingRevision AdminServiceListingRevision
	s := struct {
		DiscriminatorParam string `json:"listingType"`
		MarshalTypeAdminServiceListingRevision
	}{
		"SERVICE",
		(MarshalTypeAdminServiceListingRevision)(m),
	}

	return json.Marshal(&s)
}
