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

// CreateListingRevisionDetails The model for an Oracle Cloud Infrastructure Marketplace Publisher listing revision.
type CreateListingRevisionDetails interface {

	// The unique identifier for the listing this revision belongs to.
	GetListingId() *string

	// Single line introduction for the listing revision.
	GetHeadline() *string

	// The name for the listing revision.
	GetDisplayName() *string

	// The tagline for the listing revision.
	GetTagline() *string

	// Keywords associated with the listing revision.
	GetKeywords() *string

	// A short description for the listing revision.
	GetShortDescription() *string

	// Usage information for the listing revision.
	GetUsageInformation() *string

	// A long description for the listing revision.
	GetLongDescription() *string

	GetContentLanguage() *LanguageItem

	// Languages supported by the publisher for the listing revision.
	GetSupportedlanguages() []LanguageItem

	// Contact information to use to get support from the publisher for the listing revision.
	GetSupportContacts() []SupportContact

	// Links to support resources for the listing revision.
	GetSupportLinks() []NamedLink

	// The current status of the Listing revision.
	GetStatus() ListingRevisionStatusEnum

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	GetFreeformTags() map[string]string

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	GetDefinedTags() map[string]map[string]interface{}
}

type createlistingrevisiondetails struct {
	JsonData           []byte
	DisplayName        *string                           `mandatory:"false" json:"displayName"`
	Tagline            *string                           `mandatory:"false" json:"tagline"`
	Keywords           *string                           `mandatory:"false" json:"keywords"`
	ShortDescription   *string                           `mandatory:"false" json:"shortDescription"`
	UsageInformation   *string                           `mandatory:"false" json:"usageInformation"`
	LongDescription    *string                           `mandatory:"false" json:"longDescription"`
	ContentLanguage    *LanguageItem                     `mandatory:"false" json:"contentLanguage"`
	Supportedlanguages []LanguageItem                    `mandatory:"false" json:"supportedlanguages"`
	SupportContacts    []SupportContact                  `mandatory:"false" json:"supportContacts"`
	SupportLinks       []NamedLink                       `mandatory:"false" json:"supportLinks"`
	Status             ListingRevisionStatusEnum         `mandatory:"false" json:"status,omitempty"`
	FreeformTags       map[string]string                 `mandatory:"false" json:"freeformTags"`
	DefinedTags        map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
	ListingId          *string                           `mandatory:"true" json:"listingId"`
	Headline           *string                           `mandatory:"true" json:"headline"`
	ListingType        string                            `json:"listingType"`
}

// UnmarshalJSON unmarshals json
func (m *createlistingrevisiondetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalercreatelistingrevisiondetails createlistingrevisiondetails
	s := struct {
		Model Unmarshalercreatelistingrevisiondetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.ListingId = s.Model.ListingId
	m.Headline = s.Model.Headline
	m.DisplayName = s.Model.DisplayName
	m.Tagline = s.Model.Tagline
	m.Keywords = s.Model.Keywords
	m.ShortDescription = s.Model.ShortDescription
	m.UsageInformation = s.Model.UsageInformation
	m.LongDescription = s.Model.LongDescription
	m.ContentLanguage = s.Model.ContentLanguage
	m.Supportedlanguages = s.Model.Supportedlanguages
	m.SupportContacts = s.Model.SupportContacts
	m.SupportLinks = s.Model.SupportLinks
	m.Status = s.Model.Status
	m.FreeformTags = s.Model.FreeformTags
	m.DefinedTags = s.Model.DefinedTags
	m.ListingType = s.Model.ListingType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *createlistingrevisiondetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.ListingType {
	case "OCI_APPLICATION":
		mm := CreateOciListingRevisionDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "SERVICE":
		mm := CreateServiceListingRevisionDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "LEAD_GENERATION":
		mm := CreateLeadGenListingRevisionDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Received unsupported enum value for CreateListingRevisionDetails: %s.", m.ListingType)
		return *m, nil
	}
}

// GetDisplayName returns DisplayName
func (m createlistingrevisiondetails) GetDisplayName() *string {
	return m.DisplayName
}

// GetTagline returns Tagline
func (m createlistingrevisiondetails) GetTagline() *string {
	return m.Tagline
}

// GetKeywords returns Keywords
func (m createlistingrevisiondetails) GetKeywords() *string {
	return m.Keywords
}

// GetShortDescription returns ShortDescription
func (m createlistingrevisiondetails) GetShortDescription() *string {
	return m.ShortDescription
}

// GetUsageInformation returns UsageInformation
func (m createlistingrevisiondetails) GetUsageInformation() *string {
	return m.UsageInformation
}

// GetLongDescription returns LongDescription
func (m createlistingrevisiondetails) GetLongDescription() *string {
	return m.LongDescription
}

// GetContentLanguage returns ContentLanguage
func (m createlistingrevisiondetails) GetContentLanguage() *LanguageItem {
	return m.ContentLanguage
}

// GetSupportedlanguages returns Supportedlanguages
func (m createlistingrevisiondetails) GetSupportedlanguages() []LanguageItem {
	return m.Supportedlanguages
}

// GetSupportContacts returns SupportContacts
func (m createlistingrevisiondetails) GetSupportContacts() []SupportContact {
	return m.SupportContacts
}

// GetSupportLinks returns SupportLinks
func (m createlistingrevisiondetails) GetSupportLinks() []NamedLink {
	return m.SupportLinks
}

// GetStatus returns Status
func (m createlistingrevisiondetails) GetStatus() ListingRevisionStatusEnum {
	return m.Status
}

// GetFreeformTags returns FreeformTags
func (m createlistingrevisiondetails) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

// GetDefinedTags returns DefinedTags
func (m createlistingrevisiondetails) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

// GetListingId returns ListingId
func (m createlistingrevisiondetails) GetListingId() *string {
	return m.ListingId
}

// GetHeadline returns Headline
func (m createlistingrevisiondetails) GetHeadline() *string {
	return m.Headline
}

func (m createlistingrevisiondetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m createlistingrevisiondetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingListingRevisionStatusEnum(string(m.Status)); !ok && m.Status != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", m.Status, strings.Join(GetListingRevisionStatusEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
