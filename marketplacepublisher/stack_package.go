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

// StackPackage A package for container image listings.
type StackPackage struct {

	// The name of the listing revision package.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The unique identifier for the listing revision.
	ListingRevisionId *string `mandatory:"true" json:"listingRevisionId"`

	// The unique identifier for the compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The unique identifier for the artifact.
	ArtifactId *string `mandatory:"true" json:"artifactId"`

	// The unique identifier for the term.
	TermId *string `mandatory:"true" json:"termId"`

	// The version for the package.
	PackageVersion *string `mandatory:"true" json:"packageVersion"`

	// Identifies whether security upgrades will be provided for this package.
	AreSecurityUpgradesProvided *bool `mandatory:"true" json:"areSecurityUpgradesProvided"`

	// Identifies that this will be default package for the listing revision.
	IsDefault *bool `mandatory:"true" json:"isDefault"`

	// The date and time this listing revision package was created, expressed in RFC 3339 (https://tools.ietf.org/html/rfc3339)
	// timestamp format.
	// Example: `2016-08-25T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The date and time this listing revision package was updated, expressed in RFC 3339 (https://tools.ietf.org/html/rfc3339)
	// timestamp format.
	// Example: `2016-08-25T21:10:29.600Z`
	TimeUpdated *common.SDKTime `mandatory:"true" json:"timeUpdated"`

	// The OCID for the listing revision package in Marketplace Publisher.
	Id *string `mandatory:"false" json:"id"`

	// The description of this package.
	Description *string `mandatory:"false" json:"description"`

	// Additional metadata key/value pairs for the listing revision package summary.
	// For example:
	// `{"partnerListingRevisionPackageStatus": "Published","parentListingRevisionPackageId": "1" }`
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

	// The current state for the listing revision package.
	LifecycleState ListingRevisionPackageLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The current status for the listing revision package.
	Status ListingRevisionPackageStatusEnum `mandatory:"true" json:"status"`
}

// GetId returns Id
func (m StackPackage) GetId() *string {
	return m.Id
}

// GetDisplayName returns DisplayName
func (m StackPackage) GetDisplayName() *string {
	return m.DisplayName
}

// GetDescription returns Description
func (m StackPackage) GetDescription() *string {
	return m.Description
}

// GetListingRevisionId returns ListingRevisionId
func (m StackPackage) GetListingRevisionId() *string {
	return m.ListingRevisionId
}

// GetCompartmentId returns CompartmentId
func (m StackPackage) GetCompartmentId() *string {
	return m.CompartmentId
}

// GetArtifactId returns ArtifactId
func (m StackPackage) GetArtifactId() *string {
	return m.ArtifactId
}

// GetTermId returns TermId
func (m StackPackage) GetTermId() *string {
	return m.TermId
}

// GetPackageVersion returns PackageVersion
func (m StackPackage) GetPackageVersion() *string {
	return m.PackageVersion
}

// GetLifecycleState returns LifecycleState
func (m StackPackage) GetLifecycleState() ListingRevisionPackageLifecycleStateEnum {
	return m.LifecycleState
}

// GetStatus returns Status
func (m StackPackage) GetStatus() ListingRevisionPackageStatusEnum {
	return m.Status
}

// GetAreSecurityUpgradesProvided returns AreSecurityUpgradesProvided
func (m StackPackage) GetAreSecurityUpgradesProvided() *bool {
	return m.AreSecurityUpgradesProvided
}

// GetIsDefault returns IsDefault
func (m StackPackage) GetIsDefault() *bool {
	return m.IsDefault
}

// GetTimeCreated returns TimeCreated
func (m StackPackage) GetTimeCreated() *common.SDKTime {
	return m.TimeCreated
}

// GetTimeUpdated returns TimeUpdated
func (m StackPackage) GetTimeUpdated() *common.SDKTime {
	return m.TimeUpdated
}

// GetExtendedMetadata returns ExtendedMetadata
func (m StackPackage) GetExtendedMetadata() map[string]string {
	return m.ExtendedMetadata
}

// GetFreeformTags returns FreeformTags
func (m StackPackage) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

// GetDefinedTags returns DefinedTags
func (m StackPackage) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

// GetSystemTags returns SystemTags
func (m StackPackage) GetSystemTags() map[string]map[string]interface{} {
	return m.SystemTags
}

func (m StackPackage) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m StackPackage) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingListingRevisionPackageLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetListingRevisionPackageLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListingRevisionPackageStatusEnum(string(m.Status)); !ok && m.Status != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", m.Status, strings.Join(GetListingRevisionPackageStatusEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m StackPackage) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeStackPackage StackPackage
	s := struct {
		DiscriminatorParam string `json:"packageType"`
		MarshalTypeStackPackage
	}{
		"STACK",
		(MarshalTypeStackPackage)(m),
	}

	return json.Marshal(&s)
}
