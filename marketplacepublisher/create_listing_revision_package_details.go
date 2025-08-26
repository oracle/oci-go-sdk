// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
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

// CreateListingRevisionPackageDetails The model for an Oracle Cloud Infrastructure Marketplace Publisher listing revision package.
type CreateListingRevisionPackageDetails struct {

	// The OCID for the listing revision in Marketplace Publisher.
	ListingRevisionId *string `mandatory:"true" json:"listingRevisionId"`

	// The version for the package.
	PackageVersion *string `mandatory:"true" json:"packageVersion"`

	// The unique identifier for the artifact.
	ArtifactId *string `mandatory:"true" json:"artifactId"`

	// The unique identifier for the term.
	TermId *string `mandatory:"true" json:"termId"`

	// Identifies whether security upgrades will be provided for this package.
	AreSecurityUpgradesProvided *bool `mandatory:"true" json:"areSecurityUpgradesProvided"`

	// The name for the listing revision package.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Description for this package.
	Description *string `mandatory:"false" json:"description"`

	// Identifies that this will be default package for the listing revision.
	IsDefault *bool `mandatory:"false" json:"isDefault"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m CreateListingRevisionPackageDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateListingRevisionPackageDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
