// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
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

// ListingRevisionAttachmentSummary The model for a summary of a listing revision related attachments.
type ListingRevisionAttachmentSummary struct {

	// The OCID of the listing revision attachment.
	Id *string `mandatory:"true" json:"id"`

	// The ID of the listing revision.
	ListingRevisionId *string `mandatory:"true" json:"listingRevisionId"`

	// The unique identifier for the compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The name of the specified document.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The specified attachment type.
	AttachmentType ListingRevisionAttachmentAttachmentTypeEnum `mandatory:"true" json:"attachmentType"`

	// The URL of the specified attachment.
	ContentUrl *string `mandatory:"true" json:"contentUrl"`

	// The date and time the related document was created, expressed in RFC 3339 (https://tools.ietf.org/html/rfc3339)
	// timestamp format.
	// Example: `2022-09-24T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The date and time the related document was updated, expressed in RFC 3339 (https://tools.ietf.org/html/rfc3339)
	// timestamp format.
	// Example: `2022-09-24T21:10:29.600Z`
	TimeUpdated *common.SDKTime `mandatory:"true" json:"timeUpdated"`

	// Possible lifecycle states.
	DocumentCategory ListingRevisionAttachmentSummaryDocumentCategoryEnum `mandatory:"false" json:"documentCategory,omitempty"`

	// The MIME type of the screenshot.
	MimeType *string `mandatory:"false" json:"mimeType"`

	// The current state of the document.
	LifecycleState ListingRevisionAttachmentLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`

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

func (m ListingRevisionAttachmentSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ListingRevisionAttachmentSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListingRevisionAttachmentAttachmentTypeEnum(string(m.AttachmentType)); !ok && m.AttachmentType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for AttachmentType: %s. Supported values are: %s.", m.AttachmentType, strings.Join(GetListingRevisionAttachmentAttachmentTypeEnumStringValues(), ",")))
	}

	if _, ok := GetMappingListingRevisionAttachmentSummaryDocumentCategoryEnum(string(m.DocumentCategory)); !ok && m.DocumentCategory != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DocumentCategory: %s. Supported values are: %s.", m.DocumentCategory, strings.Join(GetListingRevisionAttachmentSummaryDocumentCategoryEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListingRevisionAttachmentLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetListingRevisionAttachmentLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListingRevisionAttachmentSummaryDocumentCategoryEnum Enum with underlying type: string
type ListingRevisionAttachmentSummaryDocumentCategoryEnum string

// Set of constants representing the allowable values for ListingRevisionAttachmentSummaryDocumentCategoryEnum
const (
	ListingRevisionAttachmentSummaryDocumentCategoryCaseStudies          ListingRevisionAttachmentSummaryDocumentCategoryEnum = "CASE_STUDIES"
	ListingRevisionAttachmentSummaryDocumentCategoryCustomizationGuides  ListingRevisionAttachmentSummaryDocumentCategoryEnum = "CUSTOMIZATION_GUIDES"
	ListingRevisionAttachmentSummaryDocumentCategoryDataSheets           ListingRevisionAttachmentSummaryDocumentCategoryEnum = "DATA_SHEETS"
	ListingRevisionAttachmentSummaryDocumentCategoryPressRelease         ListingRevisionAttachmentSummaryDocumentCategoryEnum = "PRESS_RELEASE"
	ListingRevisionAttachmentSummaryDocumentCategoryProductDocumentation ListingRevisionAttachmentSummaryDocumentCategoryEnum = "PRODUCT_DOCUMENTATION"
	ListingRevisionAttachmentSummaryDocumentCategoryUserGuides           ListingRevisionAttachmentSummaryDocumentCategoryEnum = "USER_GUIDES"
	ListingRevisionAttachmentSummaryDocumentCategoryWebinars             ListingRevisionAttachmentSummaryDocumentCategoryEnum = "WEBINARS"
	ListingRevisionAttachmentSummaryDocumentCategoryWhitepapers          ListingRevisionAttachmentSummaryDocumentCategoryEnum = "WHITEPAPERS"
)

var mappingListingRevisionAttachmentSummaryDocumentCategoryEnum = map[string]ListingRevisionAttachmentSummaryDocumentCategoryEnum{
	"CASE_STUDIES":          ListingRevisionAttachmentSummaryDocumentCategoryCaseStudies,
	"CUSTOMIZATION_GUIDES":  ListingRevisionAttachmentSummaryDocumentCategoryCustomizationGuides,
	"DATA_SHEETS":           ListingRevisionAttachmentSummaryDocumentCategoryDataSheets,
	"PRESS_RELEASE":         ListingRevisionAttachmentSummaryDocumentCategoryPressRelease,
	"PRODUCT_DOCUMENTATION": ListingRevisionAttachmentSummaryDocumentCategoryProductDocumentation,
	"USER_GUIDES":           ListingRevisionAttachmentSummaryDocumentCategoryUserGuides,
	"WEBINARS":              ListingRevisionAttachmentSummaryDocumentCategoryWebinars,
	"WHITEPAPERS":           ListingRevisionAttachmentSummaryDocumentCategoryWhitepapers,
}

var mappingListingRevisionAttachmentSummaryDocumentCategoryEnumLowerCase = map[string]ListingRevisionAttachmentSummaryDocumentCategoryEnum{
	"case_studies":          ListingRevisionAttachmentSummaryDocumentCategoryCaseStudies,
	"customization_guides":  ListingRevisionAttachmentSummaryDocumentCategoryCustomizationGuides,
	"data_sheets":           ListingRevisionAttachmentSummaryDocumentCategoryDataSheets,
	"press_release":         ListingRevisionAttachmentSummaryDocumentCategoryPressRelease,
	"product_documentation": ListingRevisionAttachmentSummaryDocumentCategoryProductDocumentation,
	"user_guides":           ListingRevisionAttachmentSummaryDocumentCategoryUserGuides,
	"webinars":              ListingRevisionAttachmentSummaryDocumentCategoryWebinars,
	"whitepapers":           ListingRevisionAttachmentSummaryDocumentCategoryWhitepapers,
}

// GetListingRevisionAttachmentSummaryDocumentCategoryEnumValues Enumerates the set of values for ListingRevisionAttachmentSummaryDocumentCategoryEnum
func GetListingRevisionAttachmentSummaryDocumentCategoryEnumValues() []ListingRevisionAttachmentSummaryDocumentCategoryEnum {
	values := make([]ListingRevisionAttachmentSummaryDocumentCategoryEnum, 0)
	for _, v := range mappingListingRevisionAttachmentSummaryDocumentCategoryEnum {
		values = append(values, v)
	}
	return values
}

// GetListingRevisionAttachmentSummaryDocumentCategoryEnumStringValues Enumerates the set of values in String for ListingRevisionAttachmentSummaryDocumentCategoryEnum
func GetListingRevisionAttachmentSummaryDocumentCategoryEnumStringValues() []string {
	return []string{
		"CASE_STUDIES",
		"CUSTOMIZATION_GUIDES",
		"DATA_SHEETS",
		"PRESS_RELEASE",
		"PRODUCT_DOCUMENTATION",
		"USER_GUIDES",
		"WEBINARS",
		"WHITEPAPERS",
	}
}

// GetMappingListingRevisionAttachmentSummaryDocumentCategoryEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListingRevisionAttachmentSummaryDocumentCategoryEnum(val string) (ListingRevisionAttachmentSummaryDocumentCategoryEnum, bool) {
	enum, ok := mappingListingRevisionAttachmentSummaryDocumentCategoryEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
