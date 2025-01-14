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

// ListingRevisionNote The model for the listing revision notes.
type ListingRevisionNote struct {

	// The OCID of the listing revision note.
	Id *string `mandatory:"true" json:"id"`

	// The unique identifier of the listing revision that the specified note belongs to.
	ListingRevisionId *string `mandatory:"true" json:"listingRevisionId"`

	// The unique identifier for the compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// type of the note.
	NoteSource ListingRevisionNoteNoteSourceEnum `mandatory:"true" json:"noteSource"`

	// Notes provided for the listing revision.
	NoteDetails *string `mandatory:"true" json:"noteDetails"`

	// The date and time the listing revision note was created, expressed in RFC 3339 (https://tools.ietf.org/html/rfc3339)
	// timestamp format.
	// Example: `2022-09-24T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The date and time the listing revision note was updated, expressed in RFC 3339 (https://tools.ietf.org/html/rfc3339)
	// timestamp format.
	// Example: `2022-09-24T21:10:29.600Z`
	TimeUpdated *common.SDKTime `mandatory:"true" json:"timeUpdated"`

	// The current state of the listing revision note.
	LifecycleState ListingRevisionNoteLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`

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

func (m ListingRevisionNote) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ListingRevisionNote) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListingRevisionNoteNoteSourceEnum(string(m.NoteSource)); !ok && m.NoteSource != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for NoteSource: %s. Supported values are: %s.", m.NoteSource, strings.Join(GetListingRevisionNoteNoteSourceEnumStringValues(), ",")))
	}

	if _, ok := GetMappingListingRevisionNoteLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetListingRevisionNoteLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListingRevisionNoteNoteSourceEnum Enum with underlying type: string
type ListingRevisionNoteNoteSourceEnum string

// Set of constants representing the allowable values for ListingRevisionNoteNoteSourceEnum
const (
	ListingRevisionNoteNoteSourcePublisher     ListingRevisionNoteNoteSourceEnum = "PUBLISHER"
	ListingRevisionNoteNoteSourceAdministrator ListingRevisionNoteNoteSourceEnum = "ADMINISTRATOR"
)

var mappingListingRevisionNoteNoteSourceEnum = map[string]ListingRevisionNoteNoteSourceEnum{
	"PUBLISHER":     ListingRevisionNoteNoteSourcePublisher,
	"ADMINISTRATOR": ListingRevisionNoteNoteSourceAdministrator,
}

var mappingListingRevisionNoteNoteSourceEnumLowerCase = map[string]ListingRevisionNoteNoteSourceEnum{
	"publisher":     ListingRevisionNoteNoteSourcePublisher,
	"administrator": ListingRevisionNoteNoteSourceAdministrator,
}

// GetListingRevisionNoteNoteSourceEnumValues Enumerates the set of values for ListingRevisionNoteNoteSourceEnum
func GetListingRevisionNoteNoteSourceEnumValues() []ListingRevisionNoteNoteSourceEnum {
	values := make([]ListingRevisionNoteNoteSourceEnum, 0)
	for _, v := range mappingListingRevisionNoteNoteSourceEnum {
		values = append(values, v)
	}
	return values
}

// GetListingRevisionNoteNoteSourceEnumStringValues Enumerates the set of values in String for ListingRevisionNoteNoteSourceEnum
func GetListingRevisionNoteNoteSourceEnumStringValues() []string {
	return []string{
		"PUBLISHER",
		"ADMINISTRATOR",
	}
}

// GetMappingListingRevisionNoteNoteSourceEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListingRevisionNoteNoteSourceEnum(val string) (ListingRevisionNoteNoteSourceEnum, bool) {
	enum, ok := mappingListingRevisionNoteNoteSourceEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListingRevisionNoteLifecycleStateEnum Enum with underlying type: string
type ListingRevisionNoteLifecycleStateEnum string

// Set of constants representing the allowable values for ListingRevisionNoteLifecycleStateEnum
const (
	ListingRevisionNoteLifecycleStateActive  ListingRevisionNoteLifecycleStateEnum = "ACTIVE"
	ListingRevisionNoteLifecycleStateDeleted ListingRevisionNoteLifecycleStateEnum = "DELETED"
)

var mappingListingRevisionNoteLifecycleStateEnum = map[string]ListingRevisionNoteLifecycleStateEnum{
	"ACTIVE":  ListingRevisionNoteLifecycleStateActive,
	"DELETED": ListingRevisionNoteLifecycleStateDeleted,
}

var mappingListingRevisionNoteLifecycleStateEnumLowerCase = map[string]ListingRevisionNoteLifecycleStateEnum{
	"active":  ListingRevisionNoteLifecycleStateActive,
	"deleted": ListingRevisionNoteLifecycleStateDeleted,
}

// GetListingRevisionNoteLifecycleStateEnumValues Enumerates the set of values for ListingRevisionNoteLifecycleStateEnum
func GetListingRevisionNoteLifecycleStateEnumValues() []ListingRevisionNoteLifecycleStateEnum {
	values := make([]ListingRevisionNoteLifecycleStateEnum, 0)
	for _, v := range mappingListingRevisionNoteLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetListingRevisionNoteLifecycleStateEnumStringValues Enumerates the set of values in String for ListingRevisionNoteLifecycleStateEnum
func GetListingRevisionNoteLifecycleStateEnumStringValues() []string {
	return []string{
		"ACTIVE",
		"DELETED",
	}
}

// GetMappingListingRevisionNoteLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListingRevisionNoteLifecycleStateEnum(val string) (ListingRevisionNoteLifecycleStateEnum, bool) {
	enum, ok := mappingListingRevisionNoteLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
