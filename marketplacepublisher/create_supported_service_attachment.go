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

// CreateSupportedServiceAttachment Supported service attachment for the listing revision.
type CreateSupportedServiceAttachment struct {

	// The OCID for the listing revision in Marketplace Publisher.
	ListingRevisionId *string `mandatory:"true" json:"listingRevisionId"`

	// Name of the service
	ServiceName *string `mandatory:"true" json:"serviceName"`

	// The name for the listing revision attachment.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Description for this specified attachment.
	Description *string `mandatory:"false" json:"description"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// Optional url to service information
	Url *string `mandatory:"false" json:"url"`

	// Type of the service
	Type SupportedServiceAttachmentTypeEnum `mandatory:"true" json:"type"`
}

// GetListingRevisionId returns ListingRevisionId
func (m CreateSupportedServiceAttachment) GetListingRevisionId() *string {
	return m.ListingRevisionId
}

// GetDisplayName returns DisplayName
func (m CreateSupportedServiceAttachment) GetDisplayName() *string {
	return m.DisplayName
}

// GetDescription returns Description
func (m CreateSupportedServiceAttachment) GetDescription() *string {
	return m.Description
}

// GetFreeformTags returns FreeformTags
func (m CreateSupportedServiceAttachment) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

// GetDefinedTags returns DefinedTags
func (m CreateSupportedServiceAttachment) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

func (m CreateSupportedServiceAttachment) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateSupportedServiceAttachment) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingSupportedServiceAttachmentTypeEnum(string(m.Type)); !ok && m.Type != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Type: %s. Supported values are: %s.", m.Type, strings.Join(GetSupportedServiceAttachmentTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m CreateSupportedServiceAttachment) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeCreateSupportedServiceAttachment CreateSupportedServiceAttachment
	s := struct {
		DiscriminatorParam string `json:"attachmentType"`
		MarshalTypeCreateSupportedServiceAttachment
	}{
		"SUPPORTED_SERVICES",
		(MarshalTypeCreateSupportedServiceAttachment)(m),
	}

	return json.Marshal(&s)
}
