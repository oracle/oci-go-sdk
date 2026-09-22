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

// AdminSupportedServiceAttachment Supported service attachment for the listing revision.
type AdminSupportedServiceAttachment struct {

	// Unique OCID identifier for the listing revision attachment.
	Id *string `mandatory:"true" json:"id"`

	// The unique identifier for the compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The unique identifier of the listing revision that the specified attachment belongs to.
	ListingRevisionId *string `mandatory:"true" json:"listingRevisionId"`

	// Name of the listing revision attachment.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The time the attachment was created. An RFC3339 formatted datetime string.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The time the attachment was updated. An RFC3339 formatted datetime string.
	TimeUpdated *common.SDKTime `mandatory:"true" json:"timeUpdated"`

	// Name of the service
	ServiceName *string `mandatory:"true" json:"serviceName"`

	// Description of the listing revision attachment.
	Description *string `mandatory:"false" json:"description"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// System tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`

	// Identifies whether the attachment is for Internal Oracle Users or external users as well.
	IsOracleUsersOnly *bool `mandatory:"false" json:"isOracleUsersOnly"`

	// Optional url to service
	Url *string `mandatory:"false" json:"url"`

	// URL of the uploaded document.
	ContentUrl *string `mandatory:"false" json:"contentUrl"`

	// The MIME type of the uploaded data.
	MimeType *string `mandatory:"false" json:"mimeType"`

	// Possible values for the publisher listing revision attachments. The source type informs whether the type of attachment for the listing revision is external or internal.
	SourceType ListingRevisionAttachmentSourceTypeEnum `mandatory:"false" json:"sourceType,omitempty"`

	// The current state of the attachment.
	LifecycleState ListingRevisionAttachmentLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// Type of service
	Type SupportedServiceAttachmentTypeEnum `mandatory:"true" json:"type"`
}

// GetId returns Id
func (m AdminSupportedServiceAttachment) GetId() *string {
	return m.Id
}

// GetCompartmentId returns CompartmentId
func (m AdminSupportedServiceAttachment) GetCompartmentId() *string {
	return m.CompartmentId
}

// GetListingRevisionId returns ListingRevisionId
func (m AdminSupportedServiceAttachment) GetListingRevisionId() *string {
	return m.ListingRevisionId
}

// GetDisplayName returns DisplayName
func (m AdminSupportedServiceAttachment) GetDisplayName() *string {
	return m.DisplayName
}

// GetDescription returns Description
func (m AdminSupportedServiceAttachment) GetDescription() *string {
	return m.Description
}

// GetSourceType returns SourceType
func (m AdminSupportedServiceAttachment) GetSourceType() ListingRevisionAttachmentSourceTypeEnum {
	return m.SourceType
}

// GetLifecycleState returns LifecycleState
func (m AdminSupportedServiceAttachment) GetLifecycleState() ListingRevisionAttachmentLifecycleStateEnum {
	return m.LifecycleState
}

// GetTimeCreated returns TimeCreated
func (m AdminSupportedServiceAttachment) GetTimeCreated() *common.SDKTime {
	return m.TimeCreated
}

// GetTimeUpdated returns TimeUpdated
func (m AdminSupportedServiceAttachment) GetTimeUpdated() *common.SDKTime {
	return m.TimeUpdated
}

// GetFreeformTags returns FreeformTags
func (m AdminSupportedServiceAttachment) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

// GetDefinedTags returns DefinedTags
func (m AdminSupportedServiceAttachment) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

// GetSystemTags returns SystemTags
func (m AdminSupportedServiceAttachment) GetSystemTags() map[string]map[string]interface{} {
	return m.SystemTags
}

func (m AdminSupportedServiceAttachment) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m AdminSupportedServiceAttachment) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingListingRevisionAttachmentSourceTypeEnum(string(m.SourceType)); !ok && m.SourceType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SourceType: %s. Supported values are: %s.", m.SourceType, strings.Join(GetListingRevisionAttachmentSourceTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListingRevisionAttachmentLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetListingRevisionAttachmentLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingSupportedServiceAttachmentTypeEnum(string(m.Type)); !ok && m.Type != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Type: %s. Supported values are: %s.", m.Type, strings.Join(GetSupportedServiceAttachmentTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m AdminSupportedServiceAttachment) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeAdminSupportedServiceAttachment AdminSupportedServiceAttachment
	s := struct {
		DiscriminatorParam string `json:"attachmentType"`
		MarshalTypeAdminSupportedServiceAttachment
	}{
		"SUPPORTED_SERVICES",
		(MarshalTypeAdminSupportedServiceAttachment)(m),
	}

	return json.Marshal(&s)
}
