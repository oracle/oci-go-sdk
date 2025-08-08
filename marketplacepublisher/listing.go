// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
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

// Listing The model for the Marketplace Publisher listing.
type Listing struct {

	// Unique OCID identifier for the listing.
	Id *string `mandatory:"true" json:"id"`

	// The unique identifier for the compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The unique identifier for the publisher.
	PublisherId *string `mandatory:"true" json:"publisherId"`

	// The listing type for the listing.
	ListingType ListingTypeEnum `mandatory:"true" json:"listingType"`

	// Name for the listing.
	Name *string `mandatory:"true" json:"name"`

	// The package type for the listing.
	PackageType PackageTypeEnum `mandatory:"true" json:"packageType"`

	// The date and time the listing was created, in the format defined by RFC3339 (https://tools.ietf.org/html/rfc3339).
	// Example: `2023-03-27T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The date and time the listing was updated, in the format defined by RFC3339 (https://tools.ietf.org/html/rfc3339).
	// Example: `2023-03-27T21:10:29.600Z`
	TimeUpdated *common.SDKTime `mandatory:"true" json:"timeUpdated"`

	// Additional metadata key/value pairs for the listing summary.
	// For example:
	// `{"listingRevisionStatus": "Published","listingRevision": "1" }`
	ExtendedMetadata map[string]string `mandatory:"false" json:"extendedMetadata"`

	// The current state of the Listing.
	LifecycleState ListingLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`

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

func (m Listing) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m Listing) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListingTypeEnum(string(m.ListingType)); !ok && m.ListingType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ListingType: %s. Supported values are: %s.", m.ListingType, strings.Join(GetListingTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingPackageTypeEnum(string(m.PackageType)); !ok && m.PackageType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for PackageType: %s. Supported values are: %s.", m.PackageType, strings.Join(GetPackageTypeEnumStringValues(), ",")))
	}

	if _, ok := GetMappingListingLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetListingLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListingLifecycleStateEnum Enum with underlying type: string
type ListingLifecycleStateEnum string

// Set of constants representing the allowable values for ListingLifecycleStateEnum
const (
	ListingLifecycleStateCreating ListingLifecycleStateEnum = "CREATING"
	ListingLifecycleStateUpdating ListingLifecycleStateEnum = "UPDATING"
	ListingLifecycleStateActive   ListingLifecycleStateEnum = "ACTIVE"
	ListingLifecycleStateInactive ListingLifecycleStateEnum = "INACTIVE"
	ListingLifecycleStateDeleting ListingLifecycleStateEnum = "DELETING"
	ListingLifecycleStateDeleted  ListingLifecycleStateEnum = "DELETED"
	ListingLifecycleStateFailed   ListingLifecycleStateEnum = "FAILED"
)

var mappingListingLifecycleStateEnum = map[string]ListingLifecycleStateEnum{
	"CREATING": ListingLifecycleStateCreating,
	"UPDATING": ListingLifecycleStateUpdating,
	"ACTIVE":   ListingLifecycleStateActive,
	"INACTIVE": ListingLifecycleStateInactive,
	"DELETING": ListingLifecycleStateDeleting,
	"DELETED":  ListingLifecycleStateDeleted,
	"FAILED":   ListingLifecycleStateFailed,
}

var mappingListingLifecycleStateEnumLowerCase = map[string]ListingLifecycleStateEnum{
	"creating": ListingLifecycleStateCreating,
	"updating": ListingLifecycleStateUpdating,
	"active":   ListingLifecycleStateActive,
	"inactive": ListingLifecycleStateInactive,
	"deleting": ListingLifecycleStateDeleting,
	"deleted":  ListingLifecycleStateDeleted,
	"failed":   ListingLifecycleStateFailed,
}

// GetListingLifecycleStateEnumValues Enumerates the set of values for ListingLifecycleStateEnum
func GetListingLifecycleStateEnumValues() []ListingLifecycleStateEnum {
	values := make([]ListingLifecycleStateEnum, 0)
	for _, v := range mappingListingLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetListingLifecycleStateEnumStringValues Enumerates the set of values in String for ListingLifecycleStateEnum
func GetListingLifecycleStateEnumStringValues() []string {
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

// GetMappingListingLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListingLifecycleStateEnum(val string) (ListingLifecycleStateEnum, bool) {
	enum, ok := mappingListingLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
