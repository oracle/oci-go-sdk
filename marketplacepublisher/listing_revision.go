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

// ListingRevision The model for an Oracle Cloud Infrastructure Marketplace Publisher listing revision.
type ListingRevision interface {

	// Unique OCID identifier for the listing revision in Marketplace Publisher.
	GetId() *string

	// The unique identifier for the listing this revision belongs to.
	GetListingId() *string

	// The name for the listing revision.
	GetDisplayName() *string

	// Single line introduction for the listing revision.
	GetHeadline() *string

	// The time the listing revision was created. An RFC3339 formatted datetime string.
	GetTimeCreated() *common.SDKTime

	// The time the listing revision was updated. An RFC3339 formatted datetime string.
	GetTimeUpdated() *common.SDKTime

	// The current status for the Listing revision.
	GetStatus() ListingRevisionStatusEnum

	// The current state of the listing revision.
	GetLifecycleState() ListingRevisionLifecycleStateEnum

	// The unique identifier for the compartment.
	GetCompartmentId() *string

	// The revision number for the listing revision. This is an internal attribute
	GetRevisionNumber() *string

	// The tagline of the listing revision.
	GetTagline() *string

	// Keywords associated with the listing revision.
	GetKeywords() *string

	// A short description for the listing revision.
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

	// Status notes for the listing revision.
	GetStatusNotes() *string

	// The listing's package type. Populated from the listing.
	GetPackageType() PackageTypeEnum

	// Additional metadata key/value pairs for the listing revision summary.
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

type listingrevision struct {
	JsonData           []byte
	CompartmentId      *string                           `mandatory:"false" json:"compartmentId"`
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
	StatusNotes        *string                           `mandatory:"false" json:"statusNotes"`
	PackageType        PackageTypeEnum                   `mandatory:"false" json:"packageType,omitempty"`
	ExtendedMetadata   map[string]string                 `mandatory:"false" json:"extendedMetadata"`
	FreeformTags       map[string]string                 `mandatory:"false" json:"freeformTags"`
	DefinedTags        map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
	SystemTags         map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
	Id                 *string                           `mandatory:"true" json:"id"`
	ListingId          *string                           `mandatory:"true" json:"listingId"`
	DisplayName        *string                           `mandatory:"true" json:"displayName"`
	Headline           *string                           `mandatory:"true" json:"headline"`
	TimeCreated        *common.SDKTime                   `mandatory:"true" json:"timeCreated"`
	TimeUpdated        *common.SDKTime                   `mandatory:"true" json:"timeUpdated"`
	Status             ListingRevisionStatusEnum         `mandatory:"true" json:"status"`
	LifecycleState     ListingRevisionLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`
	ListingType        string                            `json:"listingType"`
}

// UnmarshalJSON unmarshals json
func (m *listingrevision) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerlistingrevision listingrevision
	s := struct {
		Model Unmarshalerlistingrevision
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Id = s.Model.Id
	m.ListingId = s.Model.ListingId
	m.DisplayName = s.Model.DisplayName
	m.Headline = s.Model.Headline
	m.TimeCreated = s.Model.TimeCreated
	m.TimeUpdated = s.Model.TimeUpdated
	m.Status = s.Model.Status
	m.LifecycleState = s.Model.LifecycleState
	m.CompartmentId = s.Model.CompartmentId
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
	m.StatusNotes = s.Model.StatusNotes
	m.PackageType = s.Model.PackageType
	m.ExtendedMetadata = s.Model.ExtendedMetadata
	m.FreeformTags = s.Model.FreeformTags
	m.DefinedTags = s.Model.DefinedTags
	m.SystemTags = s.Model.SystemTags
	m.ListingType = s.Model.ListingType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *listingrevision) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.ListingType {
	case "SERVICE":
		mm := ServiceListingRevision{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "LEAD_GENERATION":
		mm := LeadGenListingRevision{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "OCI_APPLICATION":
		mm := OciListingRevision{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Received unsupported enum value for ListingRevision: %s.", m.ListingType)
		return *m, nil
	}
}

// GetCompartmentId returns CompartmentId
func (m listingrevision) GetCompartmentId() *string {
	return m.CompartmentId
}

// GetRevisionNumber returns RevisionNumber
func (m listingrevision) GetRevisionNumber() *string {
	return m.RevisionNumber
}

// GetTagline returns Tagline
func (m listingrevision) GetTagline() *string {
	return m.Tagline
}

// GetKeywords returns Keywords
func (m listingrevision) GetKeywords() *string {
	return m.Keywords
}

// GetShortDescription returns ShortDescription
func (m listingrevision) GetShortDescription() *string {
	return m.ShortDescription
}

// GetUsageInformation returns UsageInformation
func (m listingrevision) GetUsageInformation() *string {
	return m.UsageInformation
}

// GetLongDescription returns LongDescription
func (m listingrevision) GetLongDescription() *string {
	return m.LongDescription
}

// GetContentLanguage returns ContentLanguage
func (m listingrevision) GetContentLanguage() *LanguageItem {
	return m.ContentLanguage
}

// GetSupportedlanguages returns Supportedlanguages
func (m listingrevision) GetSupportedlanguages() []LanguageItem {
	return m.Supportedlanguages
}

// GetSupportContacts returns SupportContacts
func (m listingrevision) GetSupportContacts() []SupportContact {
	return m.SupportContacts
}

// GetSupportLinks returns SupportLinks
func (m listingrevision) GetSupportLinks() []NamedLink {
	return m.SupportLinks
}

// GetIcon returns Icon
func (m listingrevision) GetIcon() *ListingRevisionIconAttachment {
	return m.Icon
}

// GetStatusNotes returns StatusNotes
func (m listingrevision) GetStatusNotes() *string {
	return m.StatusNotes
}

// GetPackageType returns PackageType
func (m listingrevision) GetPackageType() PackageTypeEnum {
	return m.PackageType
}

// GetExtendedMetadata returns ExtendedMetadata
func (m listingrevision) GetExtendedMetadata() map[string]string {
	return m.ExtendedMetadata
}

// GetFreeformTags returns FreeformTags
func (m listingrevision) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

// GetDefinedTags returns DefinedTags
func (m listingrevision) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

// GetSystemTags returns SystemTags
func (m listingrevision) GetSystemTags() map[string]map[string]interface{} {
	return m.SystemTags
}

// GetId returns Id
func (m listingrevision) GetId() *string {
	return m.Id
}

// GetListingId returns ListingId
func (m listingrevision) GetListingId() *string {
	return m.ListingId
}

// GetDisplayName returns DisplayName
func (m listingrevision) GetDisplayName() *string {
	return m.DisplayName
}

// GetHeadline returns Headline
func (m listingrevision) GetHeadline() *string {
	return m.Headline
}

// GetTimeCreated returns TimeCreated
func (m listingrevision) GetTimeCreated() *common.SDKTime {
	return m.TimeCreated
}

// GetTimeUpdated returns TimeUpdated
func (m listingrevision) GetTimeUpdated() *common.SDKTime {
	return m.TimeUpdated
}

// GetStatus returns Status
func (m listingrevision) GetStatus() ListingRevisionStatusEnum {
	return m.Status
}

// GetLifecycleState returns LifecycleState
func (m listingrevision) GetLifecycleState() ListingRevisionLifecycleStateEnum {
	return m.LifecycleState
}

func (m listingrevision) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m listingrevision) ValidateEnumValue() (bool, error) {
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

// ListingRevisionStatusEnum Enum with underlying type: string
type ListingRevisionStatusEnum string

// Set of constants representing the allowable values for ListingRevisionStatusEnum
const (
	ListingRevisionStatusNew                        ListingRevisionStatusEnum = "NEW"
	ListingRevisionStatusPendingReview              ListingRevisionStatusEnum = "PENDING_REVIEW"
	ListingRevisionStatusReviewInProgress           ListingRevisionStatusEnum = "REVIEW_IN_PROGRESS"
	ListingRevisionStatusRejected                   ListingRevisionStatusEnum = "REJECTED"
	ListingRevisionStatusApproved                   ListingRevisionStatusEnum = "APPROVED"
	ListingRevisionStatusPublishInProgress          ListingRevisionStatusEnum = "PUBLISH_IN_PROGRESS"
	ListingRevisionStatusPublishFailed              ListingRevisionStatusEnum = "PUBLISH_FAILED"
	ListingRevisionStatusPublished                  ListingRevisionStatusEnum = "PUBLISHED"
	ListingRevisionStatusPublishAsPrivateFailed     ListingRevisionStatusEnum = "PUBLISH_AS_PRIVATE_FAILED"
	ListingRevisionStatusPublishedAsPrivate         ListingRevisionStatusEnum = "PUBLISHED_AS_PRIVATE"
	ListingRevisionStatusPublishAsPrivateInProgress ListingRevisionStatusEnum = "PUBLISH_AS_PRIVATE_IN_PROGRESS"
	ListingRevisionStatusUnpublishInProgress        ListingRevisionStatusEnum = "UNPUBLISH_IN_PROGRESS"
	ListingRevisionStatusUnpublished                ListingRevisionStatusEnum = "UNPUBLISHED"
)

var mappingListingRevisionStatusEnum = map[string]ListingRevisionStatusEnum{
	"NEW":                            ListingRevisionStatusNew,
	"PENDING_REVIEW":                 ListingRevisionStatusPendingReview,
	"REVIEW_IN_PROGRESS":             ListingRevisionStatusReviewInProgress,
	"REJECTED":                       ListingRevisionStatusRejected,
	"APPROVED":                       ListingRevisionStatusApproved,
	"PUBLISH_IN_PROGRESS":            ListingRevisionStatusPublishInProgress,
	"PUBLISH_FAILED":                 ListingRevisionStatusPublishFailed,
	"PUBLISHED":                      ListingRevisionStatusPublished,
	"PUBLISH_AS_PRIVATE_FAILED":      ListingRevisionStatusPublishAsPrivateFailed,
	"PUBLISHED_AS_PRIVATE":           ListingRevisionStatusPublishedAsPrivate,
	"PUBLISH_AS_PRIVATE_IN_PROGRESS": ListingRevisionStatusPublishAsPrivateInProgress,
	"UNPUBLISH_IN_PROGRESS":          ListingRevisionStatusUnpublishInProgress,
	"UNPUBLISHED":                    ListingRevisionStatusUnpublished,
}

var mappingListingRevisionStatusEnumLowerCase = map[string]ListingRevisionStatusEnum{
	"new":                            ListingRevisionStatusNew,
	"pending_review":                 ListingRevisionStatusPendingReview,
	"review_in_progress":             ListingRevisionStatusReviewInProgress,
	"rejected":                       ListingRevisionStatusRejected,
	"approved":                       ListingRevisionStatusApproved,
	"publish_in_progress":            ListingRevisionStatusPublishInProgress,
	"publish_failed":                 ListingRevisionStatusPublishFailed,
	"published":                      ListingRevisionStatusPublished,
	"publish_as_private_failed":      ListingRevisionStatusPublishAsPrivateFailed,
	"published_as_private":           ListingRevisionStatusPublishedAsPrivate,
	"publish_as_private_in_progress": ListingRevisionStatusPublishAsPrivateInProgress,
	"unpublish_in_progress":          ListingRevisionStatusUnpublishInProgress,
	"unpublished":                    ListingRevisionStatusUnpublished,
}

// GetListingRevisionStatusEnumValues Enumerates the set of values for ListingRevisionStatusEnum
func GetListingRevisionStatusEnumValues() []ListingRevisionStatusEnum {
	values := make([]ListingRevisionStatusEnum, 0)
	for _, v := range mappingListingRevisionStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetListingRevisionStatusEnumStringValues Enumerates the set of values in String for ListingRevisionStatusEnum
func GetListingRevisionStatusEnumStringValues() []string {
	return []string{
		"NEW",
		"PENDING_REVIEW",
		"REVIEW_IN_PROGRESS",
		"REJECTED",
		"APPROVED",
		"PUBLISH_IN_PROGRESS",
		"PUBLISH_FAILED",
		"PUBLISHED",
		"PUBLISH_AS_PRIVATE_FAILED",
		"PUBLISHED_AS_PRIVATE",
		"PUBLISH_AS_PRIVATE_IN_PROGRESS",
		"UNPUBLISH_IN_PROGRESS",
		"UNPUBLISHED",
	}
}

// GetMappingListingRevisionStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListingRevisionStatusEnum(val string) (ListingRevisionStatusEnum, bool) {
	enum, ok := mappingListingRevisionStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListingRevisionLifecycleStateEnum Enum with underlying type: string
type ListingRevisionLifecycleStateEnum string

// Set of constants representing the allowable values for ListingRevisionLifecycleStateEnum
const (
	ListingRevisionLifecycleStateCreating ListingRevisionLifecycleStateEnum = "CREATING"
	ListingRevisionLifecycleStateUpdating ListingRevisionLifecycleStateEnum = "UPDATING"
	ListingRevisionLifecycleStateActive   ListingRevisionLifecycleStateEnum = "ACTIVE"
	ListingRevisionLifecycleStateDeleting ListingRevisionLifecycleStateEnum = "DELETING"
	ListingRevisionLifecycleStateDeleted  ListingRevisionLifecycleStateEnum = "DELETED"
	ListingRevisionLifecycleStateFailed   ListingRevisionLifecycleStateEnum = "FAILED"
)

var mappingListingRevisionLifecycleStateEnum = map[string]ListingRevisionLifecycleStateEnum{
	"CREATING": ListingRevisionLifecycleStateCreating,
	"UPDATING": ListingRevisionLifecycleStateUpdating,
	"ACTIVE":   ListingRevisionLifecycleStateActive,
	"DELETING": ListingRevisionLifecycleStateDeleting,
	"DELETED":  ListingRevisionLifecycleStateDeleted,
	"FAILED":   ListingRevisionLifecycleStateFailed,
}

var mappingListingRevisionLifecycleStateEnumLowerCase = map[string]ListingRevisionLifecycleStateEnum{
	"creating": ListingRevisionLifecycleStateCreating,
	"updating": ListingRevisionLifecycleStateUpdating,
	"active":   ListingRevisionLifecycleStateActive,
	"deleting": ListingRevisionLifecycleStateDeleting,
	"deleted":  ListingRevisionLifecycleStateDeleted,
	"failed":   ListingRevisionLifecycleStateFailed,
}

// GetListingRevisionLifecycleStateEnumValues Enumerates the set of values for ListingRevisionLifecycleStateEnum
func GetListingRevisionLifecycleStateEnumValues() []ListingRevisionLifecycleStateEnum {
	values := make([]ListingRevisionLifecycleStateEnum, 0)
	for _, v := range mappingListingRevisionLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetListingRevisionLifecycleStateEnumStringValues Enumerates the set of values in String for ListingRevisionLifecycleStateEnum
func GetListingRevisionLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"UPDATING",
		"ACTIVE",
		"DELETING",
		"DELETED",
		"FAILED",
	}
}

// GetMappingListingRevisionLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListingRevisionLifecycleStateEnum(val string) (ListingRevisionLifecycleStateEnum, bool) {
	enum, ok := mappingListingRevisionLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
