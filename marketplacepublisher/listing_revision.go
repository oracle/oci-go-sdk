// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
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

// ListingRevision The model for an Oracle Cloud Infrastructure Marketplace Publisher listing revision.
type ListingRevision struct {

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

	// The categories for the listing revsion.
	Categories []string `mandatory:"true" json:"categories"`

	// The current status for the Listing revision.
	Status ListingRevisionStatusEnum `mandatory:"true" json:"status"`

	// The current state of the listing revision.
	LifecycleState ListingRevisionLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The listing's package type. Populated from the listing.
	PackageType PackageTypeEnum `mandatory:"true" json:"packageType"`

	// The pricing model for the listing revision.
	PricingType ListingRevisionPricingTypeEnum `mandatory:"true" json:"pricingType"`

	// The unique identifier for the compartment.
	CompartmentId *string `mandatory:"false" json:"compartmentId"`

	// The revision number for the listing revision. This is an internal attribute
	RevisionNumber *string `mandatory:"false" json:"revisionNumber"`

	VersionDetails *VersionDetails `mandatory:"false" json:"versionDetails"`

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

	// System requirements for the listing revision.
	SystemRequirements *string `mandatory:"false" json:"systemRequirements"`

	// The markets supported by the listing revision.
	Markets []string `mandatory:"false" json:"markets"`

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

	// Allowed tenancies provided when a listing revision is published as private.
	AllowedTenancies []string `mandatory:"false" json:"allowedTenancies"`

	// Identifies whether publisher allows internal tenancy launches for the listing revision.
	AreInternalTenancyLaunchAllowed *bool `mandatory:"false" json:"areInternalTenancyLaunchAllowed"`

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
}

func (m ListingRevision) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ListingRevision) ValidateEnumValue() (bool, error) {
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

// ListingRevisionPricingTypeEnum Enum with underlying type: string
type ListingRevisionPricingTypeEnum string

// Set of constants representing the allowable values for ListingRevisionPricingTypeEnum
const (
	ListingRevisionPricingTypeFree  ListingRevisionPricingTypeEnum = "FREE"
	ListingRevisionPricingTypeByol  ListingRevisionPricingTypeEnum = "BYOL"
	ListingRevisionPricingTypePaygo ListingRevisionPricingTypeEnum = "PAYGO"
)

var mappingListingRevisionPricingTypeEnum = map[string]ListingRevisionPricingTypeEnum{
	"FREE":  ListingRevisionPricingTypeFree,
	"BYOL":  ListingRevisionPricingTypeByol,
	"PAYGO": ListingRevisionPricingTypePaygo,
}

var mappingListingRevisionPricingTypeEnumLowerCase = map[string]ListingRevisionPricingTypeEnum{
	"free":  ListingRevisionPricingTypeFree,
	"byol":  ListingRevisionPricingTypeByol,
	"paygo": ListingRevisionPricingTypePaygo,
}

// GetListingRevisionPricingTypeEnumValues Enumerates the set of values for ListingRevisionPricingTypeEnum
func GetListingRevisionPricingTypeEnumValues() []ListingRevisionPricingTypeEnum {
	values := make([]ListingRevisionPricingTypeEnum, 0)
	for _, v := range mappingListingRevisionPricingTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetListingRevisionPricingTypeEnumStringValues Enumerates the set of values in String for ListingRevisionPricingTypeEnum
func GetListingRevisionPricingTypeEnumStringValues() []string {
	return []string{
		"FREE",
		"BYOL",
		"PAYGO",
	}
}

// GetMappingListingRevisionPricingTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListingRevisionPricingTypeEnum(val string) (ListingRevisionPricingTypeEnum, bool) {
	enum, ok := mappingListingRevisionPricingTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
