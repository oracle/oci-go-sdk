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

// ListingRevisionPackage A base object for all types of listing revision packages.
type ListingRevisionPackage interface {

	// The name of the listing revision package.
	GetDisplayName() *string

	// The unique identifier for the listing revision.
	GetListingRevisionId() *string

	// The unique identifier for the compartment.
	GetCompartmentId() *string

	// The unique identifier for the artifact.
	GetArtifactId() *string

	// The unique identifier for the term.
	GetTermId() *string

	// The version for the package.
	GetPackageVersion() *string

	// The current state for the listing revision package.
	GetLifecycleState() ListingRevisionPackageLifecycleStateEnum

	// The current status for the listing revision package.
	GetStatus() ListingRevisionPackageStatusEnum

	// Identifies whether security upgrades will be provided for this package.
	GetAreSecurityUpgradesProvided() *bool

	// Identifies that this will be default package for the listing revision.
	GetIsDefault() *bool

	// The date and time this listing revision package was created, expressed in RFC 3339 (https://tools.ietf.org/html/rfc3339)
	// timestamp format.
	// Example: `2016-08-25T21:10:29.600Z`
	GetTimeCreated() *common.SDKTime

	// The date and time this listing revision package was updated, expressed in RFC 3339 (https://tools.ietf.org/html/rfc3339)
	// timestamp format.
	// Example: `2016-08-25T21:10:29.600Z`
	GetTimeUpdated() *common.SDKTime

	// The OCID for the listing revision package in Marketplace Publisher.
	GetId() *string

	// The description of this package.
	GetDescription() *string

	// Additional metadata key/value pairs for the listing revision package summary.
	// For example:
	// `{"partnerListingRevisionPackageStatus": "Published","parentListingRevisionPackageId": "1" }`
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

type listingrevisionpackage struct {
	JsonData                    []byte
	Id                          *string                                  `mandatory:"false" json:"id"`
	Description                 *string                                  `mandatory:"false" json:"description"`
	ExtendedMetadata            map[string]string                        `mandatory:"false" json:"extendedMetadata"`
	FreeformTags                map[string]string                        `mandatory:"false" json:"freeformTags"`
	DefinedTags                 map[string]map[string]interface{}        `mandatory:"false" json:"definedTags"`
	SystemTags                  map[string]map[string]interface{}        `mandatory:"false" json:"systemTags"`
	DisplayName                 *string                                  `mandatory:"true" json:"displayName"`
	ListingRevisionId           *string                                  `mandatory:"true" json:"listingRevisionId"`
	CompartmentId               *string                                  `mandatory:"true" json:"compartmentId"`
	ArtifactId                  *string                                  `mandatory:"true" json:"artifactId"`
	TermId                      *string                                  `mandatory:"true" json:"termId"`
	PackageVersion              *string                                  `mandatory:"true" json:"packageVersion"`
	LifecycleState              ListingRevisionPackageLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`
	Status                      ListingRevisionPackageStatusEnum         `mandatory:"true" json:"status"`
	AreSecurityUpgradesProvided *bool                                    `mandatory:"true" json:"areSecurityUpgradesProvided"`
	IsDefault                   *bool                                    `mandatory:"true" json:"isDefault"`
	TimeCreated                 *common.SDKTime                          `mandatory:"true" json:"timeCreated"`
	TimeUpdated                 *common.SDKTime                          `mandatory:"true" json:"timeUpdated"`
	PackageType                 string                                   `json:"packageType"`
}

// UnmarshalJSON unmarshals json
func (m *listingrevisionpackage) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerlistingrevisionpackage listingrevisionpackage
	s := struct {
		Model Unmarshalerlistingrevisionpackage
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.DisplayName = s.Model.DisplayName
	m.ListingRevisionId = s.Model.ListingRevisionId
	m.CompartmentId = s.Model.CompartmentId
	m.ArtifactId = s.Model.ArtifactId
	m.TermId = s.Model.TermId
	m.PackageVersion = s.Model.PackageVersion
	m.LifecycleState = s.Model.LifecycleState
	m.Status = s.Model.Status
	m.AreSecurityUpgradesProvided = s.Model.AreSecurityUpgradesProvided
	m.IsDefault = s.Model.IsDefault
	m.TimeCreated = s.Model.TimeCreated
	m.TimeUpdated = s.Model.TimeUpdated
	m.Id = s.Model.Id
	m.Description = s.Model.Description
	m.ExtendedMetadata = s.Model.ExtendedMetadata
	m.FreeformTags = s.Model.FreeformTags
	m.DefinedTags = s.Model.DefinedTags
	m.SystemTags = s.Model.SystemTags
	m.PackageType = s.Model.PackageType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *listingrevisionpackage) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.PackageType {
	case "HELM_CHART":
		mm := HelmChartPackage{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "CONTAINER_IMAGE":
		mm := ContainerPackage{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for ListingRevisionPackage: %s.", m.PackageType)
		return *m, nil
	}
}

// GetId returns Id
func (m listingrevisionpackage) GetId() *string {
	return m.Id
}

// GetDescription returns Description
func (m listingrevisionpackage) GetDescription() *string {
	return m.Description
}

// GetExtendedMetadata returns ExtendedMetadata
func (m listingrevisionpackage) GetExtendedMetadata() map[string]string {
	return m.ExtendedMetadata
}

// GetFreeformTags returns FreeformTags
func (m listingrevisionpackage) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

// GetDefinedTags returns DefinedTags
func (m listingrevisionpackage) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

// GetSystemTags returns SystemTags
func (m listingrevisionpackage) GetSystemTags() map[string]map[string]interface{} {
	return m.SystemTags
}

// GetDisplayName returns DisplayName
func (m listingrevisionpackage) GetDisplayName() *string {
	return m.DisplayName
}

// GetListingRevisionId returns ListingRevisionId
func (m listingrevisionpackage) GetListingRevisionId() *string {
	return m.ListingRevisionId
}

// GetCompartmentId returns CompartmentId
func (m listingrevisionpackage) GetCompartmentId() *string {
	return m.CompartmentId
}

// GetArtifactId returns ArtifactId
func (m listingrevisionpackage) GetArtifactId() *string {
	return m.ArtifactId
}

// GetTermId returns TermId
func (m listingrevisionpackage) GetTermId() *string {
	return m.TermId
}

// GetPackageVersion returns PackageVersion
func (m listingrevisionpackage) GetPackageVersion() *string {
	return m.PackageVersion
}

// GetLifecycleState returns LifecycleState
func (m listingrevisionpackage) GetLifecycleState() ListingRevisionPackageLifecycleStateEnum {
	return m.LifecycleState
}

// GetStatus returns Status
func (m listingrevisionpackage) GetStatus() ListingRevisionPackageStatusEnum {
	return m.Status
}

// GetAreSecurityUpgradesProvided returns AreSecurityUpgradesProvided
func (m listingrevisionpackage) GetAreSecurityUpgradesProvided() *bool {
	return m.AreSecurityUpgradesProvided
}

// GetIsDefault returns IsDefault
func (m listingrevisionpackage) GetIsDefault() *bool {
	return m.IsDefault
}

// GetTimeCreated returns TimeCreated
func (m listingrevisionpackage) GetTimeCreated() *common.SDKTime {
	return m.TimeCreated
}

// GetTimeUpdated returns TimeUpdated
func (m listingrevisionpackage) GetTimeUpdated() *common.SDKTime {
	return m.TimeUpdated
}

func (m listingrevisionpackage) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m listingrevisionpackage) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListingRevisionPackageLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetListingRevisionPackageLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListingRevisionPackageStatusEnum(string(m.Status)); !ok && m.Status != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", m.Status, strings.Join(GetListingRevisionPackageStatusEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListingRevisionPackageLifecycleStateEnum Enum with underlying type: string
type ListingRevisionPackageLifecycleStateEnum string

// Set of constants representing the allowable values for ListingRevisionPackageLifecycleStateEnum
const (
	ListingRevisionPackageLifecycleStateCreating ListingRevisionPackageLifecycleStateEnum = "CREATING"
	ListingRevisionPackageLifecycleStateUpdating ListingRevisionPackageLifecycleStateEnum = "UPDATING"
	ListingRevisionPackageLifecycleStateActive   ListingRevisionPackageLifecycleStateEnum = "ACTIVE"
	ListingRevisionPackageLifecycleStateInactive ListingRevisionPackageLifecycleStateEnum = "INACTIVE"
	ListingRevisionPackageLifecycleStateDeleting ListingRevisionPackageLifecycleStateEnum = "DELETING"
	ListingRevisionPackageLifecycleStateDeleted  ListingRevisionPackageLifecycleStateEnum = "DELETED"
	ListingRevisionPackageLifecycleStateFailed   ListingRevisionPackageLifecycleStateEnum = "FAILED"
)

var mappingListingRevisionPackageLifecycleStateEnum = map[string]ListingRevisionPackageLifecycleStateEnum{
	"CREATING": ListingRevisionPackageLifecycleStateCreating,
	"UPDATING": ListingRevisionPackageLifecycleStateUpdating,
	"ACTIVE":   ListingRevisionPackageLifecycleStateActive,
	"INACTIVE": ListingRevisionPackageLifecycleStateInactive,
	"DELETING": ListingRevisionPackageLifecycleStateDeleting,
	"DELETED":  ListingRevisionPackageLifecycleStateDeleted,
	"FAILED":   ListingRevisionPackageLifecycleStateFailed,
}

var mappingListingRevisionPackageLifecycleStateEnumLowerCase = map[string]ListingRevisionPackageLifecycleStateEnum{
	"creating": ListingRevisionPackageLifecycleStateCreating,
	"updating": ListingRevisionPackageLifecycleStateUpdating,
	"active":   ListingRevisionPackageLifecycleStateActive,
	"inactive": ListingRevisionPackageLifecycleStateInactive,
	"deleting": ListingRevisionPackageLifecycleStateDeleting,
	"deleted":  ListingRevisionPackageLifecycleStateDeleted,
	"failed":   ListingRevisionPackageLifecycleStateFailed,
}

// GetListingRevisionPackageLifecycleStateEnumValues Enumerates the set of values for ListingRevisionPackageLifecycleStateEnum
func GetListingRevisionPackageLifecycleStateEnumValues() []ListingRevisionPackageLifecycleStateEnum {
	values := make([]ListingRevisionPackageLifecycleStateEnum, 0)
	for _, v := range mappingListingRevisionPackageLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetListingRevisionPackageLifecycleStateEnumStringValues Enumerates the set of values in String for ListingRevisionPackageLifecycleStateEnum
func GetListingRevisionPackageLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"UPDATING",
		"ACTIVE",
		"INACTIVE",
		"DELETING",
		"DELETED",
		"FAILED",
	}
}

// GetMappingListingRevisionPackageLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListingRevisionPackageLifecycleStateEnum(val string) (ListingRevisionPackageLifecycleStateEnum, bool) {
	enum, ok := mappingListingRevisionPackageLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListingRevisionPackageStatusEnum Enum with underlying type: string
type ListingRevisionPackageStatusEnum string

// Set of constants representing the allowable values for ListingRevisionPackageStatusEnum
const (
	ListingRevisionPackageStatusNew                 ListingRevisionPackageStatusEnum = "NEW"
	ListingRevisionPackageStatusPublishInProgress   ListingRevisionPackageStatusEnum = "PUBLISH_IN_PROGRESS"
	ListingRevisionPackageStatusUnpublishInProgress ListingRevisionPackageStatusEnum = "UNPUBLISH_IN_PROGRESS"
	ListingRevisionPackageStatusPublishFailed       ListingRevisionPackageStatusEnum = "PUBLISH_FAILED"
	ListingRevisionPackageStatusPublished           ListingRevisionPackageStatusEnum = "PUBLISHED"
	ListingRevisionPackageStatusPublishedAsPrivate  ListingRevisionPackageStatusEnum = "PUBLISHED_AS_PRIVATE"
	ListingRevisionPackageStatusUnpublished         ListingRevisionPackageStatusEnum = "UNPUBLISHED"
)

var mappingListingRevisionPackageStatusEnum = map[string]ListingRevisionPackageStatusEnum{
	"NEW":                   ListingRevisionPackageStatusNew,
	"PUBLISH_IN_PROGRESS":   ListingRevisionPackageStatusPublishInProgress,
	"UNPUBLISH_IN_PROGRESS": ListingRevisionPackageStatusUnpublishInProgress,
	"PUBLISH_FAILED":        ListingRevisionPackageStatusPublishFailed,
	"PUBLISHED":             ListingRevisionPackageStatusPublished,
	"PUBLISHED_AS_PRIVATE":  ListingRevisionPackageStatusPublishedAsPrivate,
	"UNPUBLISHED":           ListingRevisionPackageStatusUnpublished,
}

var mappingListingRevisionPackageStatusEnumLowerCase = map[string]ListingRevisionPackageStatusEnum{
	"new":                   ListingRevisionPackageStatusNew,
	"publish_in_progress":   ListingRevisionPackageStatusPublishInProgress,
	"unpublish_in_progress": ListingRevisionPackageStatusUnpublishInProgress,
	"publish_failed":        ListingRevisionPackageStatusPublishFailed,
	"published":             ListingRevisionPackageStatusPublished,
	"published_as_private":  ListingRevisionPackageStatusPublishedAsPrivate,
	"unpublished":           ListingRevisionPackageStatusUnpublished,
}

// GetListingRevisionPackageStatusEnumValues Enumerates the set of values for ListingRevisionPackageStatusEnum
func GetListingRevisionPackageStatusEnumValues() []ListingRevisionPackageStatusEnum {
	values := make([]ListingRevisionPackageStatusEnum, 0)
	for _, v := range mappingListingRevisionPackageStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetListingRevisionPackageStatusEnumStringValues Enumerates the set of values in String for ListingRevisionPackageStatusEnum
func GetListingRevisionPackageStatusEnumStringValues() []string {
	return []string{
		"NEW",
		"PUBLISH_IN_PROGRESS",
		"UNPUBLISH_IN_PROGRESS",
		"PUBLISH_FAILED",
		"PUBLISHED",
		"PUBLISHED_AS_PRIVATE",
		"UNPUBLISHED",
	}
}

// GetMappingListingRevisionPackageStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListingRevisionPackageStatusEnum(val string) (ListingRevisionPackageStatusEnum, bool) {
	enum, ok := mappingListingRevisionPackageStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
