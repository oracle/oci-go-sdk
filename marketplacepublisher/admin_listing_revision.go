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

// AdminListingRevision The model for an Oracle Cloud Infrastructure Marketplace Publisher listing revision.
type AdminListingRevision interface {

	// The OCID for the listing revision in Marketplace Publisher.
	GetId() *string

	// The unique identifier for the listing this revision belongs to.
	GetListingId() *string

	// The OCID for the publisher in Marketplace Publisher.
	GetPublisherId() *string

	// The name for the listing revision.
	GetDisplayName() *string

	// Single line introduction for the listing revision.
	GetHeadline() *string

	// The time the listing revision was created. An RFC3339 formatted datetime string.
	GetTimeCreated() *common.SDKTime

	// The time the listing revision was updated. An RFC3339 formatted datetime string.
	GetTimeUpdated() *common.SDKTime

	// The current status for the listing revision.
	GetStatus() ListingRevisionStatusEnum

	// The current state of the listing revision.
	GetLifecycleState() ListingRevisionLifecycleStateEnum

	// The listing's package type. Populated from the listing.
	GetPackageType() PackageTypeEnum

	// The unique legacy identifier for the listing.
	GetLegacyId() *string

	// The unique identifier for the compartment.
	GetCompartmentId() *string

	// The user OCID of the administrator who approved the listing revision.
	GetApproverId() *string

	// The email address of the administrator who approved the listing revision.
	GetApproverEmail() *string

	// The revision number of the listing revision. This is an internal attribute
	GetRevisionNumber() *string

	// The tagline for the listing revision.
	GetTagline() *string

	// Keywords associated with the listing revision.
	GetKeywords() *string

	// A short description of the listing revision.
	GetShortDescription() *string

	// Usage information for the listing revision.
	GetUsageInformation() *string

	// A long description for the listing revision.
	GetLongDescription() *string

	GetContentLanguage() *LanguageItem

	// Languages supported by the publisher for the listing revision.
	GetSupportedlanguages() []LanguageItem

	// Contact information to use to get support from the publisher for the listing revision.
	GetSupportContacts() []SupportContact

	// Links to support resources for the listing revision.
	GetSupportLinks() []NamedLink

	GetIcon() *ListingRevisionIconAttachment

	GetBanner() *ListingRevisionBannerAttachment

	// Status notes for the listing revision.
	GetStatusNotes() *string

	// Additional metadata key/value pairs for the listing revision.
	// For example:
	// `{"oracleStandardTermVersionId": "1234"}`
	GetExtendedMetadata() map[string]string

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	GetFreeformTags() map[string]string

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	GetDefinedTags() map[string]map[string]interface{}

	// System tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	GetSystemTags() map[string]map[string]interface{}
}

type adminlistingrevision struct {
	JsonData           []byte
	LegacyId           *string                           `mandatory:"false" json:"legacyId"`
	CompartmentId      *string                           `mandatory:"false" json:"compartmentId"`
	ApproverId         *string                           `mandatory:"false" json:"approverId"`
	ApproverEmail      *string                           `mandatory:"false" json:"approverEmail"`
	RevisionNumber     *string                           `mandatory:"false" json:"revisionNumber"`
	Tagline            *string                           `mandatory:"false" json:"tagline"`
	Keywords           *string                           `mandatory:"false" json:"keywords"`
	ShortDescription   *string                           `mandatory:"false" json:"shortDescription"`
	UsageInformation   *string                           `mandatory:"false" json:"usageInformation"`
	LongDescription    *string                           `mandatory:"false" json:"longDescription"`
	ContentLanguage    *LanguageItem                     `mandatory:"false" json:"contentLanguage"`
	Supportedlanguages []LanguageItem                    `mandatory:"false" json:"supportedlanguages"`
	SupportContacts    []SupportContact                  `mandatory:"false" json:"supportContacts"`
	SupportLinks       []NamedLink                       `mandatory:"false" json:"supportLinks"`
	Icon               *ListingRevisionIconAttachment    `mandatory:"false" json:"icon"`
	Banner             *ListingRevisionBannerAttachment  `mandatory:"false" json:"banner"`
	StatusNotes        *string                           `mandatory:"false" json:"statusNotes"`
	ExtendedMetadata   map[string]string                 `mandatory:"false" json:"extendedMetadata"`
	FreeformTags       map[string]string                 `mandatory:"false" json:"freeformTags"`
	DefinedTags        map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
	SystemTags         map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
	Id                 *string                           `mandatory:"true" json:"id"`
	ListingId          *string                           `mandatory:"true" json:"listingId"`
	PublisherId        *string                           `mandatory:"true" json:"publisherId"`
	DisplayName        *string                           `mandatory:"true" json:"displayName"`
	Headline           *string                           `mandatory:"true" json:"headline"`
	TimeCreated        *common.SDKTime                   `mandatory:"true" json:"timeCreated"`
	TimeUpdated        *common.SDKTime                   `mandatory:"true" json:"timeUpdated"`
	Status             ListingRevisionStatusEnum         `mandatory:"true" json:"status"`
	LifecycleState     ListingRevisionLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`
	PackageType        PackageTypeEnum                   `mandatory:"true" json:"packageType"`
	ListingType        string                            `json:"listingType"`
}

// UnmarshalJSON unmarshals json
func (m *adminlistingrevision) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshaleradminlistingrevision adminlistingrevision
	s := struct {
		Model Unmarshaleradminlistingrevision
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Id = s.Model.Id
	m.ListingId = s.Model.ListingId
	m.PublisherId = s.Model.PublisherId
	m.DisplayName = s.Model.DisplayName
	m.Headline = s.Model.Headline
	m.TimeCreated = s.Model.TimeCreated
	m.TimeUpdated = s.Model.TimeUpdated
	m.Status = s.Model.Status
	m.LifecycleState = s.Model.LifecycleState
	m.PackageType = s.Model.PackageType
	m.LegacyId = s.Model.LegacyId
	m.CompartmentId = s.Model.CompartmentId
	m.ApproverId = s.Model.ApproverId
	m.ApproverEmail = s.Model.ApproverEmail
	m.RevisionNumber = s.Model.RevisionNumber
	m.Tagline = s.Model.Tagline
	m.Keywords = s.Model.Keywords
	m.ShortDescription = s.Model.ShortDescription
	m.UsageInformation = s.Model.UsageInformation
	m.LongDescription = s.Model.LongDescription
	m.ContentLanguage = s.Model.ContentLanguage
	m.Supportedlanguages = s.Model.Supportedlanguages
	m.SupportContacts = s.Model.SupportContacts
	m.SupportLinks = s.Model.SupportLinks
	m.Icon = s.Model.Icon
	m.Banner = s.Model.Banner
	m.StatusNotes = s.Model.StatusNotes
	m.ExtendedMetadata = s.Model.ExtendedMetadata
	m.FreeformTags = s.Model.FreeformTags
	m.DefinedTags = s.Model.DefinedTags
	m.SystemTags = s.Model.SystemTags
	m.ListingType = s.Model.ListingType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *adminlistingrevision) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.ListingType {
	case "SERVICE":
		mm := AdminServiceListingRevision{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "AI_AGENT":
		mm := AdminAiAgentListingRevision{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "OCI_APPLICATION":
		mm := AdminOciListingRevision{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "LEAD_GENERATION":
		mm := AdminLeadGenListingRevision{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Received unsupported enum value for AdminListingRevision: %s.", m.ListingType)
		return *m, nil
	}
}

// GetLegacyId returns LegacyId
func (m adminlistingrevision) GetLegacyId() *string {
	return m.LegacyId
}

// GetCompartmentId returns CompartmentId
func (m adminlistingrevision) GetCompartmentId() *string {
	return m.CompartmentId
}

// GetApproverId returns ApproverId
func (m adminlistingrevision) GetApproverId() *string {
	return m.ApproverId
}

// GetApproverEmail returns ApproverEmail
func (m adminlistingrevision) GetApproverEmail() *string {
	return m.ApproverEmail
}

// GetRevisionNumber returns RevisionNumber
func (m adminlistingrevision) GetRevisionNumber() *string {
	return m.RevisionNumber
}

// GetTagline returns Tagline
func (m adminlistingrevision) GetTagline() *string {
	return m.Tagline
}

// GetKeywords returns Keywords
func (m adminlistingrevision) GetKeywords() *string {
	return m.Keywords
}

// GetShortDescription returns ShortDescription
func (m adminlistingrevision) GetShortDescription() *string {
	return m.ShortDescription
}

// GetUsageInformation returns UsageInformation
func (m adminlistingrevision) GetUsageInformation() *string {
	return m.UsageInformation
}

// GetLongDescription returns LongDescription
func (m adminlistingrevision) GetLongDescription() *string {
	return m.LongDescription
}

// GetContentLanguage returns ContentLanguage
func (m adminlistingrevision) GetContentLanguage() *LanguageItem {
	return m.ContentLanguage
}

// GetSupportedlanguages returns Supportedlanguages
func (m adminlistingrevision) GetSupportedlanguages() []LanguageItem {
	return m.Supportedlanguages
}

// GetSupportContacts returns SupportContacts
func (m adminlistingrevision) GetSupportContacts() []SupportContact {
	return m.SupportContacts
}

// GetSupportLinks returns SupportLinks
func (m adminlistingrevision) GetSupportLinks() []NamedLink {
	return m.SupportLinks
}

// GetIcon returns Icon
func (m adminlistingrevision) GetIcon() *ListingRevisionIconAttachment {
	return m.Icon
}

// GetBanner returns Banner
func (m adminlistingrevision) GetBanner() *ListingRevisionBannerAttachment {
	return m.Banner
}

// GetStatusNotes returns StatusNotes
func (m adminlistingrevision) GetStatusNotes() *string {
	return m.StatusNotes
}

// GetExtendedMetadata returns ExtendedMetadata
func (m adminlistingrevision) GetExtendedMetadata() map[string]string {
	return m.ExtendedMetadata
}

// GetFreeformTags returns FreeformTags
func (m adminlistingrevision) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

// GetDefinedTags returns DefinedTags
func (m adminlistingrevision) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

// GetSystemTags returns SystemTags
func (m adminlistingrevision) GetSystemTags() map[string]map[string]interface{} {
	return m.SystemTags
}

// GetId returns Id
func (m adminlistingrevision) GetId() *string {
	return m.Id
}

// GetListingId returns ListingId
func (m adminlistingrevision) GetListingId() *string {
	return m.ListingId
}

// GetPublisherId returns PublisherId
func (m adminlistingrevision) GetPublisherId() *string {
	return m.PublisherId
}

// GetDisplayName returns DisplayName
func (m adminlistingrevision) GetDisplayName() *string {
	return m.DisplayName
}

// GetHeadline returns Headline
func (m adminlistingrevision) GetHeadline() *string {
	return m.Headline
}

// GetTimeCreated returns TimeCreated
func (m adminlistingrevision) GetTimeCreated() *common.SDKTime {
	return m.TimeCreated
}

// GetTimeUpdated returns TimeUpdated
func (m adminlistingrevision) GetTimeUpdated() *common.SDKTime {
	return m.TimeUpdated
}

// GetStatus returns Status
func (m adminlistingrevision) GetStatus() ListingRevisionStatusEnum {
	return m.Status
}

// GetLifecycleState returns LifecycleState
func (m adminlistingrevision) GetLifecycleState() ListingRevisionLifecycleStateEnum {
	return m.LifecycleState
}

// GetPackageType returns PackageType
func (m adminlistingrevision) GetPackageType() PackageTypeEnum {
	return m.PackageType
}

func (m adminlistingrevision) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m adminlistingrevision) ValidateEnumValue() (bool, error) {
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
