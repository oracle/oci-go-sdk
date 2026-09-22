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

// AdminListingRevisionAttachment A attachment for the listing revision. User can provide an external URL/upload a file
type AdminListingRevisionAttachment interface {

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

	// Possible values for the publisher listing revision attachments. The source type informs whether the type of attachment for the listing revision is external or internal.
	GetSourceType() ListingRevisionAttachmentSourceTypeEnum

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

type adminlistingrevisionattachment struct {
	JsonData          []byte
	Description       *string                                     `mandatory:"false" json:"description"`
	SourceType        ListingRevisionAttachmentSourceTypeEnum     `mandatory:"false" json:"sourceType,omitempty"`
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
func (m *adminlistingrevisionattachment) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshaleradminlistingrevisionattachment adminlistingrevisionattachment
	s := struct {
		Model Unmarshaleradminlistingrevisionattachment
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
	m.SourceType = s.Model.SourceType
	m.FreeformTags = s.Model.FreeformTags
	m.DefinedTags = s.Model.DefinedTags
	m.SystemTags = s.Model.SystemTags
	m.AttachmentType = s.Model.AttachmentType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *adminlistingrevisionattachment) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.AttachmentType {
	case "RELATED_DOCUMENT":
		mm := AdminRelatedDocumentAttachment{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "SCREENSHOT":
		mm := AdminScreenShotAttachment{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "REVIEW_SUPPORT_DOCUMENT":
		mm := AdminReviewSupportDocumentAttachment{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "SUPPORTED_SERVICES":
		mm := AdminSupportedServiceAttachment{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "CUSTOMER_SUCCESS":
		mm := AdminCustomerSuccessAttachment{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "VIDEO":
		mm := AdminVideoAttachment{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Received unsupported enum value for AdminListingRevisionAttachment: %s.", m.AttachmentType)
		return *m, nil
	}
}

// GetDescription returns Description
func (m adminlistingrevisionattachment) GetDescription() *string {
	return m.Description
}

// GetSourceType returns SourceType
func (m adminlistingrevisionattachment) GetSourceType() ListingRevisionAttachmentSourceTypeEnum {
	return m.SourceType
}

// GetFreeformTags returns FreeformTags
func (m adminlistingrevisionattachment) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

// GetDefinedTags returns DefinedTags
func (m adminlistingrevisionattachment) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

// GetSystemTags returns SystemTags
func (m adminlistingrevisionattachment) GetSystemTags() map[string]map[string]interface{} {
	return m.SystemTags
}

// GetId returns Id
func (m adminlistingrevisionattachment) GetId() *string {
	return m.Id
}

// GetCompartmentId returns CompartmentId
func (m adminlistingrevisionattachment) GetCompartmentId() *string {
	return m.CompartmentId
}

// GetListingRevisionId returns ListingRevisionId
func (m adminlistingrevisionattachment) GetListingRevisionId() *string {
	return m.ListingRevisionId
}

// GetDisplayName returns DisplayName
func (m adminlistingrevisionattachment) GetDisplayName() *string {
	return m.DisplayName
}

// GetLifecycleState returns LifecycleState
func (m adminlistingrevisionattachment) GetLifecycleState() ListingRevisionAttachmentLifecycleStateEnum {
	return m.LifecycleState
}

// GetTimeCreated returns TimeCreated
func (m adminlistingrevisionattachment) GetTimeCreated() *common.SDKTime {
	return m.TimeCreated
}

// GetTimeUpdated returns TimeUpdated
func (m adminlistingrevisionattachment) GetTimeUpdated() *common.SDKTime {
	return m.TimeUpdated
}

func (m adminlistingrevisionattachment) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m adminlistingrevisionattachment) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListingRevisionAttachmentLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetListingRevisionAttachmentLifecycleStateEnumStringValues(), ",")))
	}

	if _, ok := GetMappingListingRevisionAttachmentSourceTypeEnum(string(m.SourceType)); !ok && m.SourceType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SourceType: %s. Supported values are: %s.", m.SourceType, strings.Join(GetListingRevisionAttachmentSourceTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
