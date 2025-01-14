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

// UpdateListingRevisionAttachmentDetails Update the attachment for the listing revision.
type UpdateListingRevisionAttachmentDetails interface {

	// The name for the listing revision attachment.
	GetDisplayName() *string

	// The description for the listing revision attachment.
	GetDescription() *string

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	GetFreeformTags() map[string]string

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	GetDefinedTags() map[string]map[string]interface{}
}

type updatelistingrevisionattachmentdetails struct {
	JsonData       []byte
	DisplayName    *string                           `mandatory:"false" json:"displayName"`
	Description    *string                           `mandatory:"false" json:"description"`
	FreeformTags   map[string]string                 `mandatory:"false" json:"freeformTags"`
	DefinedTags    map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
	AttachmentType string                            `json:"attachmentType"`
}

// UnmarshalJSON unmarshals json
func (m *updatelistingrevisionattachmentdetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerupdatelistingrevisionattachmentdetails updatelistingrevisionattachmentdetails
	s := struct {
		Model Unmarshalerupdatelistingrevisionattachmentdetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.DisplayName = s.Model.DisplayName
	m.Description = s.Model.Description
	m.FreeformTags = s.Model.FreeformTags
	m.DefinedTags = s.Model.DefinedTags
	m.AttachmentType = s.Model.AttachmentType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *updatelistingrevisionattachmentdetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.AttachmentType {
	case "SCREENSHOT":
		mm := UpdateScreenShotAttachmentDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "RELATED_DOCUMENT":
		mm := UpdateRelatedDocumentAttachmentDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "VIDEO":
		mm := UpdateVideoAttachmentDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for UpdateListingRevisionAttachmentDetails: %s.", m.AttachmentType)
		return *m, nil
	}
}

// GetDisplayName returns DisplayName
func (m updatelistingrevisionattachmentdetails) GetDisplayName() *string {
	return m.DisplayName
}

// GetDescription returns Description
func (m updatelistingrevisionattachmentdetails) GetDescription() *string {
	return m.Description
}

// GetFreeformTags returns FreeformTags
func (m updatelistingrevisionattachmentdetails) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

// GetDefinedTags returns DefinedTags
func (m updatelistingrevisionattachmentdetails) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

func (m updatelistingrevisionattachmentdetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m updatelistingrevisionattachmentdetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
