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

// SupportedServiceAttachment Supported services attachment for the service listing revision.
type SupportedServiceAttachment struct {

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

	// Optional url to service
	Url *string `mandatory:"false" json:"url"`

	// URL of the uploaded document.
	ContentUrl *string `mandatory:"false" json:"contentUrl"`

	// The MIME type of the uploaded data.
	MimeType *string `mandatory:"false" json:"mimeType"`

	// Type of service
	Type SupportedServiceAttachmentTypeEnum `mandatory:"true" json:"type"`

	// The current state of the attachment.
	LifecycleState ListingRevisionAttachmentLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`
}

// GetId returns Id
func (m SupportedServiceAttachment) GetId() *string {
	return m.Id
}

// GetCompartmentId returns CompartmentId
func (m SupportedServiceAttachment) GetCompartmentId() *string {
	return m.CompartmentId
}

// GetListingRevisionId returns ListingRevisionId
func (m SupportedServiceAttachment) GetListingRevisionId() *string {
	return m.ListingRevisionId
}

// GetDisplayName returns DisplayName
func (m SupportedServiceAttachment) GetDisplayName() *string {
	return m.DisplayName
}

// GetDescription returns Description
func (m SupportedServiceAttachment) GetDescription() *string {
	return m.Description
}

// GetLifecycleState returns LifecycleState
func (m SupportedServiceAttachment) GetLifecycleState() ListingRevisionAttachmentLifecycleStateEnum {
	return m.LifecycleState
}

// GetTimeCreated returns TimeCreated
func (m SupportedServiceAttachment) GetTimeCreated() *common.SDKTime {
	return m.TimeCreated
}

// GetTimeUpdated returns TimeUpdated
func (m SupportedServiceAttachment) GetTimeUpdated() *common.SDKTime {
	return m.TimeUpdated
}

// GetFreeformTags returns FreeformTags
func (m SupportedServiceAttachment) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

// GetDefinedTags returns DefinedTags
func (m SupportedServiceAttachment) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

// GetSystemTags returns SystemTags
func (m SupportedServiceAttachment) GetSystemTags() map[string]map[string]interface{} {
	return m.SystemTags
}

func (m SupportedServiceAttachment) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m SupportedServiceAttachment) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingSupportedServiceAttachmentTypeEnum(string(m.Type)); !ok && m.Type != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Type: %s. Supported values are: %s.", m.Type, strings.Join(GetSupportedServiceAttachmentTypeEnumStringValues(), ",")))
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
func (m SupportedServiceAttachment) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeSupportedServiceAttachment SupportedServiceAttachment
	s := struct {
		DiscriminatorParam string `json:"attachmentType"`
		MarshalTypeSupportedServiceAttachment
	}{
		"SUPPORTED_SERVICES",
		(MarshalTypeSupportedServiceAttachment)(m),
	}

	return json.Marshal(&s)
}

// SupportedServiceAttachmentTypeEnum Enum with underlying type: string
type SupportedServiceAttachmentTypeEnum string

// Set of constants representing the allowable values for SupportedServiceAttachmentTypeEnum
const (
	SupportedServiceAttachmentTypeImplementationService SupportedServiceAttachmentTypeEnum = "IMPLEMENTATION_SERVICE"
	SupportedServiceAttachmentTypeConsultingService     SupportedServiceAttachmentTypeEnum = "CONSULTING_SERVICE"
	SupportedServiceAttachmentTypeAppDevelopmentService SupportedServiceAttachmentTypeEnum = "APP_DEVELOPMENT_SERVICE"
	SupportedServiceAttachmentTypeManagedService        SupportedServiceAttachmentTypeEnum = "MANAGED_SERVICE"
	SupportedServiceAttachmentTypeOther                 SupportedServiceAttachmentTypeEnum = "OTHER"
)

var mappingSupportedServiceAttachmentTypeEnum = map[string]SupportedServiceAttachmentTypeEnum{
	"IMPLEMENTATION_SERVICE":  SupportedServiceAttachmentTypeImplementationService,
	"CONSULTING_SERVICE":      SupportedServiceAttachmentTypeConsultingService,
	"APP_DEVELOPMENT_SERVICE": SupportedServiceAttachmentTypeAppDevelopmentService,
	"MANAGED_SERVICE":         SupportedServiceAttachmentTypeManagedService,
	"OTHER":                   SupportedServiceAttachmentTypeOther,
}

var mappingSupportedServiceAttachmentTypeEnumLowerCase = map[string]SupportedServiceAttachmentTypeEnum{
	"implementation_service":  SupportedServiceAttachmentTypeImplementationService,
	"consulting_service":      SupportedServiceAttachmentTypeConsultingService,
	"app_development_service": SupportedServiceAttachmentTypeAppDevelopmentService,
	"managed_service":         SupportedServiceAttachmentTypeManagedService,
	"other":                   SupportedServiceAttachmentTypeOther,
}

// GetSupportedServiceAttachmentTypeEnumValues Enumerates the set of values for SupportedServiceAttachmentTypeEnum
func GetSupportedServiceAttachmentTypeEnumValues() []SupportedServiceAttachmentTypeEnum {
	values := make([]SupportedServiceAttachmentTypeEnum, 0)
	for _, v := range mappingSupportedServiceAttachmentTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetSupportedServiceAttachmentTypeEnumStringValues Enumerates the set of values in String for SupportedServiceAttachmentTypeEnum
func GetSupportedServiceAttachmentTypeEnumStringValues() []string {
	return []string{
		"IMPLEMENTATION_SERVICE",
		"CONSULTING_SERVICE",
		"APP_DEVELOPMENT_SERVICE",
		"MANAGED_SERVICE",
		"OTHER",
	}
}

// GetMappingSupportedServiceAttachmentTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSupportedServiceAttachmentTypeEnum(val string) (SupportedServiceAttachmentTypeEnum, bool) {
	enum, ok := mappingSupportedServiceAttachmentTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
