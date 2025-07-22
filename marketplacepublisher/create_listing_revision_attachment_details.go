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

// CreateListingRevisionAttachmentDetails Attachment uploaded by the publisher for the listing revision.
type CreateListingRevisionAttachmentDetails interface {

	// The OCID for the listing revision in Marketplace Publisher.
	GetListingRevisionId() *string

	// The name for the listing revision attachment.
	GetDisplayName() *string

	// Description for this specified attachment.
	GetDescription() *string

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	GetFreeformTags() map[string]string

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	GetDefinedTags() map[string]map[string]interface{}
}

type createlistingrevisionattachmentdetails struct {
	JsonData          []byte
	DisplayName       *string                           `mandatory:"false" json:"displayName"`
	Description       *string                           `mandatory:"false" json:"description"`
	FreeformTags      map[string]string                 `mandatory:"false" json:"freeformTags"`
	DefinedTags       map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
	ListingRevisionId *string                           `mandatory:"true" json:"listingRevisionId"`
	AttachmentType    string                            `json:"attachmentType"`
}

// UnmarshalJSON unmarshals json
func (m *createlistingrevisionattachmentdetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalercreatelistingrevisionattachmentdetails createlistingrevisionattachmentdetails
	s := struct {
		Model Unmarshalercreatelistingrevisionattachmentdetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.ListingRevisionId = s.Model.ListingRevisionId
	m.DisplayName = s.Model.DisplayName
	m.Description = s.Model.Description
	m.FreeformTags = s.Model.FreeformTags
	m.DefinedTags = s.Model.DefinedTags
	m.AttachmentType = s.Model.AttachmentType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *createlistingrevisionattachmentdetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.AttachmentType {
	case "SUPPORTED_SERVICES":
		mm := CreateSupportedServiceAttachment{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "SCREENSHOT":
		mm := CreateScreenShotAttachmentDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "CUSTOMER_SUCCESS":
		mm := CreateCustomerSuccessAttachment{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "REVIEW_SUPPORT_DOCUMENT":
		mm := CreateReviewSupportDocumentAttachment{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "VIDEO":
		mm := CreateVideoAttachmentDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "RELATED_DOCUMENT":
		mm := CreateRelatedDocumentAttachmentDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Received unsupported enum value for CreateListingRevisionAttachmentDetails: %s.", m.AttachmentType)
		return *m, nil
	}
}

// GetDisplayName returns DisplayName
func (m createlistingrevisionattachmentdetails) GetDisplayName() *string {
	return m.DisplayName
}

// GetDescription returns Description
func (m createlistingrevisionattachmentdetails) GetDescription() *string {
	return m.Description
}

// GetFreeformTags returns FreeformTags
func (m createlistingrevisionattachmentdetails) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

// GetDefinedTags returns DefinedTags
func (m createlistingrevisionattachmentdetails) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

// GetListingRevisionId returns ListingRevisionId
func (m createlistingrevisionattachmentdetails) GetListingRevisionId() *string {
	return m.ListingRevisionId
}

func (m createlistingrevisionattachmentdetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m createlistingrevisionattachmentdetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
