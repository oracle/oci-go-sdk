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

// RelatedDocumentAttachment Related document attachment for the listing revision.
type RelatedDocumentAttachment struct {

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

	// URL of the uploaded document.
	ContentUrl *string `mandatory:"false" json:"contentUrl"`

	// The MIME type of the uploaded data.
	MimeType *string `mandatory:"false" json:"mimeType"`

	// Possible lifecycle states.
	DocumentCategory RelatedDocumentAttachmentDocumentCategoryEnum `mandatory:"false" json:"documentCategory,omitempty"`

	// The current state of the attachment.
	LifecycleState ListingRevisionAttachmentLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`
}

// GetId returns Id
func (m RelatedDocumentAttachment) GetId() *string {
	return m.Id
}

// GetCompartmentId returns CompartmentId
func (m RelatedDocumentAttachment) GetCompartmentId() *string {
	return m.CompartmentId
}

// GetListingRevisionId returns ListingRevisionId
func (m RelatedDocumentAttachment) GetListingRevisionId() *string {
	return m.ListingRevisionId
}

// GetDisplayName returns DisplayName
func (m RelatedDocumentAttachment) GetDisplayName() *string {
	return m.DisplayName
}

// GetDescription returns Description
func (m RelatedDocumentAttachment) GetDescription() *string {
	return m.Description
}

// GetLifecycleState returns LifecycleState
func (m RelatedDocumentAttachment) GetLifecycleState() ListingRevisionAttachmentLifecycleStateEnum {
	return m.LifecycleState
}

// GetTimeCreated returns TimeCreated
func (m RelatedDocumentAttachment) GetTimeCreated() *common.SDKTime {
	return m.TimeCreated
}

// GetTimeUpdated returns TimeUpdated
func (m RelatedDocumentAttachment) GetTimeUpdated() *common.SDKTime {
	return m.TimeUpdated
}

// GetFreeformTags returns FreeformTags
func (m RelatedDocumentAttachment) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

// GetDefinedTags returns DefinedTags
func (m RelatedDocumentAttachment) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

// GetSystemTags returns SystemTags
func (m RelatedDocumentAttachment) GetSystemTags() map[string]map[string]interface{} {
	return m.SystemTags
}

func (m RelatedDocumentAttachment) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m RelatedDocumentAttachment) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingRelatedDocumentAttachmentDocumentCategoryEnum(string(m.DocumentCategory)); !ok && m.DocumentCategory != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DocumentCategory: %s. Supported values are: %s.", m.DocumentCategory, strings.Join(GetRelatedDocumentAttachmentDocumentCategoryEnumStringValues(), ",")))
	}

	if _, ok := GetMappingListingRevisionAttachmentLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetListingRevisionAttachmentLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m RelatedDocumentAttachment) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeRelatedDocumentAttachment RelatedDocumentAttachment
	s := struct {
		DiscriminatorParam string `json:"attachmentType"`
		MarshalTypeRelatedDocumentAttachment
	}{
		"RELATED_DOCUMENT",
		(MarshalTypeRelatedDocumentAttachment)(m),
	}

	return json.Marshal(&s)
}

// RelatedDocumentAttachmentDocumentCategoryEnum Enum with underlying type: string
type RelatedDocumentAttachmentDocumentCategoryEnum string

// Set of constants representing the allowable values for RelatedDocumentAttachmentDocumentCategoryEnum
const (
	RelatedDocumentAttachmentDocumentCategoryCaseStudies          RelatedDocumentAttachmentDocumentCategoryEnum = "CASE_STUDIES"
	RelatedDocumentAttachmentDocumentCategoryCustomizationGuides  RelatedDocumentAttachmentDocumentCategoryEnum = "CUSTOMIZATION_GUIDES"
	RelatedDocumentAttachmentDocumentCategoryDataSheets           RelatedDocumentAttachmentDocumentCategoryEnum = "DATA_SHEETS"
	RelatedDocumentAttachmentDocumentCategoryPressRelease         RelatedDocumentAttachmentDocumentCategoryEnum = "PRESS_RELEASE"
	RelatedDocumentAttachmentDocumentCategoryProductDocumentation RelatedDocumentAttachmentDocumentCategoryEnum = "PRODUCT_DOCUMENTATION"
	RelatedDocumentAttachmentDocumentCategoryUserGuides           RelatedDocumentAttachmentDocumentCategoryEnum = "USER_GUIDES"
	RelatedDocumentAttachmentDocumentCategoryWebinar              RelatedDocumentAttachmentDocumentCategoryEnum = "WEBINAR"
)

var mappingRelatedDocumentAttachmentDocumentCategoryEnum = map[string]RelatedDocumentAttachmentDocumentCategoryEnum{
	"CASE_STUDIES":          RelatedDocumentAttachmentDocumentCategoryCaseStudies,
	"CUSTOMIZATION_GUIDES":  RelatedDocumentAttachmentDocumentCategoryCustomizationGuides,
	"DATA_SHEETS":           RelatedDocumentAttachmentDocumentCategoryDataSheets,
	"PRESS_RELEASE":         RelatedDocumentAttachmentDocumentCategoryPressRelease,
	"PRODUCT_DOCUMENTATION": RelatedDocumentAttachmentDocumentCategoryProductDocumentation,
	"USER_GUIDES":           RelatedDocumentAttachmentDocumentCategoryUserGuides,
	"WEBINAR":               RelatedDocumentAttachmentDocumentCategoryWebinar,
}

var mappingRelatedDocumentAttachmentDocumentCategoryEnumLowerCase = map[string]RelatedDocumentAttachmentDocumentCategoryEnum{
	"case_studies":          RelatedDocumentAttachmentDocumentCategoryCaseStudies,
	"customization_guides":  RelatedDocumentAttachmentDocumentCategoryCustomizationGuides,
	"data_sheets":           RelatedDocumentAttachmentDocumentCategoryDataSheets,
	"press_release":         RelatedDocumentAttachmentDocumentCategoryPressRelease,
	"product_documentation": RelatedDocumentAttachmentDocumentCategoryProductDocumentation,
	"user_guides":           RelatedDocumentAttachmentDocumentCategoryUserGuides,
	"webinar":               RelatedDocumentAttachmentDocumentCategoryWebinar,
}

// GetRelatedDocumentAttachmentDocumentCategoryEnumValues Enumerates the set of values for RelatedDocumentAttachmentDocumentCategoryEnum
func GetRelatedDocumentAttachmentDocumentCategoryEnumValues() []RelatedDocumentAttachmentDocumentCategoryEnum {
	values := make([]RelatedDocumentAttachmentDocumentCategoryEnum, 0)
	for _, v := range mappingRelatedDocumentAttachmentDocumentCategoryEnum {
		values = append(values, v)
	}
	return values
}

// GetRelatedDocumentAttachmentDocumentCategoryEnumStringValues Enumerates the set of values in String for RelatedDocumentAttachmentDocumentCategoryEnum
func GetRelatedDocumentAttachmentDocumentCategoryEnumStringValues() []string {
	return []string{
		"CASE_STUDIES",
		"CUSTOMIZATION_GUIDES",
		"DATA_SHEETS",
		"PRESS_RELEASE",
		"PRODUCT_DOCUMENTATION",
		"USER_GUIDES",
		"WEBINAR",
	}
}

// GetMappingRelatedDocumentAttachmentDocumentCategoryEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingRelatedDocumentAttachmentDocumentCategoryEnum(val string) (RelatedDocumentAttachmentDocumentCategoryEnum, bool) {
	enum, ok := mappingRelatedDocumentAttachmentDocumentCategoryEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
