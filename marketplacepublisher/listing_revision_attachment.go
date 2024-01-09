// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
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

// ListingRevisionAttachment A attachment for the listing revision. User can provide an external URL/upload a file
type ListingRevisionAttachment interface {

	// Unique OCID identifier for the listing revision attachment.
	GetId() *string

	// The unique identifier for the compartment.
	GetCompartmentId() *string

	// The unique identifier of the listing revision that the specified attachment belongs to.
	GetListingRevisionId() *string

	// Name of the listing revision attachment.
	GetDisplayName() *string

	// The current state of the attachment.
	GetLifecycleState() ListingRevisionAttachmentLifecycleStateEnum

	// The time the attachment was created. An RFC3339 formatted datetime string.
	GetTimeCreated() *common.SDKTime

	// The time the attachment was updated. An RFC3339 formatted datetime string.
	GetTimeUpdated() *common.SDKTime

	// Description of the listing revision attachment.
	GetDescription() *string

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

type listingrevisionattachment struct {
	JsonData          []byte
	Description       *string                                     `mandatory:"false" json:"description"`
	FreeformTags      map[string]string                           `mandatory:"false" json:"freeformTags"`
	DefinedTags       map[string]map[string]interface{}           `mandatory:"false" json:"definedTags"`
	SystemTags        map[string]map[string]interface{}           `mandatory:"false" json:"systemTags"`
	Id                *string                                     `mandatory:"true" json:"id"`
	CompartmentId     *string                                     `mandatory:"true" json:"compartmentId"`
	ListingRevisionId *string                                     `mandatory:"true" json:"listingRevisionId"`
	DisplayName       *string                                     `mandatory:"true" json:"displayName"`
	LifecycleState    ListingRevisionAttachmentLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`
	TimeCreated       *common.SDKTime                             `mandatory:"true" json:"timeCreated"`
	TimeUpdated       *common.SDKTime                             `mandatory:"true" json:"timeUpdated"`
	AttachmentType    string                                      `json:"attachmentType"`
}

// UnmarshalJSON unmarshals json
func (m *listingrevisionattachment) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerlistingrevisionattachment listingrevisionattachment
	s := struct {
		Model Unmarshalerlistingrevisionattachment
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Id = s.Model.Id
	m.CompartmentId = s.Model.CompartmentId
	m.ListingRevisionId = s.Model.ListingRevisionId
	m.DisplayName = s.Model.DisplayName
	m.LifecycleState = s.Model.LifecycleState
	m.TimeCreated = s.Model.TimeCreated
	m.TimeUpdated = s.Model.TimeUpdated
	m.Description = s.Model.Description
	m.FreeformTags = s.Model.FreeformTags
	m.DefinedTags = s.Model.DefinedTags
	m.SystemTags = s.Model.SystemTags
	m.AttachmentType = s.Model.AttachmentType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *listingrevisionattachment) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.AttachmentType {
	case "RELATED_DOCUMENT":
		mm := RelatedDocumentAttachment{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "SCREENSHOT":
		mm := ScreenShotAttachment{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "VIDEO":
		mm := VideoAttachment{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for ListingRevisionAttachment: %s.", m.AttachmentType)
		return *m, nil
	}
}

// GetDescription returns Description
func (m listingrevisionattachment) GetDescription() *string {
	return m.Description
}

// GetFreeformTags returns FreeformTags
func (m listingrevisionattachment) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

// GetDefinedTags returns DefinedTags
func (m listingrevisionattachment) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

// GetSystemTags returns SystemTags
func (m listingrevisionattachment) GetSystemTags() map[string]map[string]interface{} {
	return m.SystemTags
}

// GetId returns Id
func (m listingrevisionattachment) GetId() *string {
	return m.Id
}

// GetCompartmentId returns CompartmentId
func (m listingrevisionattachment) GetCompartmentId() *string {
	return m.CompartmentId
}

// GetListingRevisionId returns ListingRevisionId
func (m listingrevisionattachment) GetListingRevisionId() *string {
	return m.ListingRevisionId
}

// GetDisplayName returns DisplayName
func (m listingrevisionattachment) GetDisplayName() *string {
	return m.DisplayName
}

// GetLifecycleState returns LifecycleState
func (m listingrevisionattachment) GetLifecycleState() ListingRevisionAttachmentLifecycleStateEnum {
	return m.LifecycleState
}

// GetTimeCreated returns TimeCreated
func (m listingrevisionattachment) GetTimeCreated() *common.SDKTime {
	return m.TimeCreated
}

// GetTimeUpdated returns TimeUpdated
func (m listingrevisionattachment) GetTimeUpdated() *common.SDKTime {
	return m.TimeUpdated
}

func (m listingrevisionattachment) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m listingrevisionattachment) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListingRevisionAttachmentLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetListingRevisionAttachmentLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListingRevisionAttachmentLifecycleStateEnum Enum with underlying type: string
type ListingRevisionAttachmentLifecycleStateEnum string

// Set of constants representing the allowable values for ListingRevisionAttachmentLifecycleStateEnum
const (
	ListingRevisionAttachmentLifecycleStateActive   ListingRevisionAttachmentLifecycleStateEnum = "ACTIVE"
	ListingRevisionAttachmentLifecycleStateInactive ListingRevisionAttachmentLifecycleStateEnum = "INACTIVE"
	ListingRevisionAttachmentLifecycleStateDeleted  ListingRevisionAttachmentLifecycleStateEnum = "DELETED"
)

var mappingListingRevisionAttachmentLifecycleStateEnum = map[string]ListingRevisionAttachmentLifecycleStateEnum{
	"ACTIVE":   ListingRevisionAttachmentLifecycleStateActive,
	"INACTIVE": ListingRevisionAttachmentLifecycleStateInactive,
	"DELETED":  ListingRevisionAttachmentLifecycleStateDeleted,
}

var mappingListingRevisionAttachmentLifecycleStateEnumLowerCase = map[string]ListingRevisionAttachmentLifecycleStateEnum{
	"active":   ListingRevisionAttachmentLifecycleStateActive,
	"inactive": ListingRevisionAttachmentLifecycleStateInactive,
	"deleted":  ListingRevisionAttachmentLifecycleStateDeleted,
}

// GetListingRevisionAttachmentLifecycleStateEnumValues Enumerates the set of values for ListingRevisionAttachmentLifecycleStateEnum
func GetListingRevisionAttachmentLifecycleStateEnumValues() []ListingRevisionAttachmentLifecycleStateEnum {
	values := make([]ListingRevisionAttachmentLifecycleStateEnum, 0)
	for _, v := range mappingListingRevisionAttachmentLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetListingRevisionAttachmentLifecycleStateEnumStringValues Enumerates the set of values in String for ListingRevisionAttachmentLifecycleStateEnum
func GetListingRevisionAttachmentLifecycleStateEnumStringValues() []string {
	return []string{
		"ACTIVE",
		"INACTIVE",
		"DELETED",
	}
}

// GetMappingListingRevisionAttachmentLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListingRevisionAttachmentLifecycleStateEnum(val string) (ListingRevisionAttachmentLifecycleStateEnum, bool) {
	enum, ok := mappingListingRevisionAttachmentLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListingRevisionAttachmentAttachmentTypeEnum Enum with underlying type: string
type ListingRevisionAttachmentAttachmentTypeEnum string

// Set of constants representing the allowable values for ListingRevisionAttachmentAttachmentTypeEnum
const (
	ListingRevisionAttachmentAttachmentTypeRelatedDocument ListingRevisionAttachmentAttachmentTypeEnum = "RELATED_DOCUMENT"
	ListingRevisionAttachmentAttachmentTypeScreenshot      ListingRevisionAttachmentAttachmentTypeEnum = "SCREENSHOT"
	ListingRevisionAttachmentAttachmentTypeVideo           ListingRevisionAttachmentAttachmentTypeEnum = "VIDEO"
)

var mappingListingRevisionAttachmentAttachmentTypeEnum = map[string]ListingRevisionAttachmentAttachmentTypeEnum{
	"RELATED_DOCUMENT": ListingRevisionAttachmentAttachmentTypeRelatedDocument,
	"SCREENSHOT":       ListingRevisionAttachmentAttachmentTypeScreenshot,
	"VIDEO":            ListingRevisionAttachmentAttachmentTypeVideo,
}

var mappingListingRevisionAttachmentAttachmentTypeEnumLowerCase = map[string]ListingRevisionAttachmentAttachmentTypeEnum{
	"related_document": ListingRevisionAttachmentAttachmentTypeRelatedDocument,
	"screenshot":       ListingRevisionAttachmentAttachmentTypeScreenshot,
	"video":            ListingRevisionAttachmentAttachmentTypeVideo,
}

// GetListingRevisionAttachmentAttachmentTypeEnumValues Enumerates the set of values for ListingRevisionAttachmentAttachmentTypeEnum
func GetListingRevisionAttachmentAttachmentTypeEnumValues() []ListingRevisionAttachmentAttachmentTypeEnum {
	values := make([]ListingRevisionAttachmentAttachmentTypeEnum, 0)
	for _, v := range mappingListingRevisionAttachmentAttachmentTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetListingRevisionAttachmentAttachmentTypeEnumStringValues Enumerates the set of values in String for ListingRevisionAttachmentAttachmentTypeEnum
func GetListingRevisionAttachmentAttachmentTypeEnumStringValues() []string {
	return []string{
		"RELATED_DOCUMENT",
		"SCREENSHOT",
		"VIDEO",
	}
}

// GetMappingListingRevisionAttachmentAttachmentTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListingRevisionAttachmentAttachmentTypeEnum(val string) (ListingRevisionAttachmentAttachmentTypeEnum, bool) {
	enum, ok := mappingListingRevisionAttachmentAttachmentTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
