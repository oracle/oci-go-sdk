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

// UpdateListingRevisionDetails The model for an Oracle Cloud Infrastructure Marketplace Publisher listing revision.
type UpdateListingRevisionDetails interface {

	// The name for the listing revision.
	GetDisplayName() *string

	// Single line introduction for the listing revision.
	GetHeadline() *string

	// The tagline for the listing revision.
	GetTagline() *string

	// Keywords associated for the listing revision.
	GetKeywords() *string

	// A short description for the listing revision.
	GetShortDescription() *string

	// Usage information for the listing revision.
	GetUsageInformation() *string

	// A long description for the listing revision.
	GetLongDescription() *string

	GetContentLanguage() *LanguageItem

	// Languages supported by the listing revision.
	GetSupportedlanguages() []LanguageItem

	// Contact information to use to get support from the publisher for the listing revision.
	GetSupportContacts() []SupportContact

	// Links to support resources for the listing revision.
	GetSupportLinks() []NamedLink

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	GetFreeformTags() map[string]string

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	GetDefinedTags() map[string]map[string]interface{}
}

type updatelistingrevisiondetails struct {
	JsonData           []byte
	DisplayName        *string                           `mandatory:"false" json:"displayName"`
	Headline           *string                           `mandatory:"false" json:"headline"`
	Tagline            *string                           `mandatory:"false" json:"tagline"`
	Keywords           *string                           `mandatory:"false" json:"keywords"`
	ShortDescription   *string                           `mandatory:"false" json:"shortDescription"`
	UsageInformation   *string                           `mandatory:"false" json:"usageInformation"`
	LongDescription    *string                           `mandatory:"false" json:"longDescription"`
	ContentLanguage    *LanguageItem                     `mandatory:"false" json:"contentLanguage"`
	Supportedlanguages []LanguageItem                    `mandatory:"false" json:"supportedlanguages"`
	SupportContacts    []SupportContact                  `mandatory:"false" json:"supportContacts"`
	SupportLinks       []NamedLink                       `mandatory:"false" json:"supportLinks"`
	FreeformTags       map[string]string                 `mandatory:"false" json:"freeformTags"`
	DefinedTags        map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
	ListingType        string                            `json:"listingType"`
}

// UnmarshalJSON unmarshals json
func (m *updatelistingrevisiondetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerupdatelistingrevisiondetails updatelistingrevisiondetails
	s := struct {
		Model Unmarshalerupdatelistingrevisiondetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.DisplayName = s.Model.DisplayName
	m.Headline = s.Model.Headline
	m.Tagline = s.Model.Tagline
	m.Keywords = s.Model.Keywords
	m.ShortDescription = s.Model.ShortDescription
	m.UsageInformation = s.Model.UsageInformation
	m.LongDescription = s.Model.LongDescription
	m.ContentLanguage = s.Model.ContentLanguage
	m.Supportedlanguages = s.Model.Supportedlanguages
	m.SupportContacts = s.Model.SupportContacts
	m.SupportLinks = s.Model.SupportLinks
	m.FreeformTags = s.Model.FreeformTags
	m.DefinedTags = s.Model.DefinedTags
	m.ListingType = s.Model.ListingType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *updatelistingrevisiondetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.ListingType {
	case "SERVICE":
		mm := UpdateServiceListingRevisionDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "OCI_APPLICATION":
		mm := UpdateOciListingRevisionDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "LEAD_GENERATION":
		mm := UpdateLeadGenListingRevisionDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Received unsupported enum value for UpdateListingRevisionDetails: %s.", m.ListingType)
		return *m, nil
	}
}

// GetDisplayName returns DisplayName
func (m updatelistingrevisiondetails) GetDisplayName() *string {
	return m.DisplayName
}

// GetHeadline returns Headline
func (m updatelistingrevisiondetails) GetHeadline() *string {
	return m.Headline
}

// GetTagline returns Tagline
func (m updatelistingrevisiondetails) GetTagline() *string {
	return m.Tagline
}

// GetKeywords returns Keywords
func (m updatelistingrevisiondetails) GetKeywords() *string {
	return m.Keywords
}

// GetShortDescription returns ShortDescription
func (m updatelistingrevisiondetails) GetShortDescription() *string {
	return m.ShortDescription
}

// GetUsageInformation returns UsageInformation
func (m updatelistingrevisiondetails) GetUsageInformation() *string {
	return m.UsageInformation
}

// GetLongDescription returns LongDescription
func (m updatelistingrevisiondetails) GetLongDescription() *string {
	return m.LongDescription
}

// GetContentLanguage returns ContentLanguage
func (m updatelistingrevisiondetails) GetContentLanguage() *LanguageItem {
	return m.ContentLanguage
}

// GetSupportedlanguages returns Supportedlanguages
func (m updatelistingrevisiondetails) GetSupportedlanguages() []LanguageItem {
	return m.Supportedlanguages
}

// GetSupportContacts returns SupportContacts
func (m updatelistingrevisiondetails) GetSupportContacts() []SupportContact {
	return m.SupportContacts
}

// GetSupportLinks returns SupportLinks
func (m updatelistingrevisiondetails) GetSupportLinks() []NamedLink {
	return m.SupportLinks
}

// GetFreeformTags returns FreeformTags
func (m updatelistingrevisiondetails) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

// GetDefinedTags returns DefinedTags
func (m updatelistingrevisiondetails) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

func (m updatelistingrevisiondetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m updatelistingrevisiondetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
