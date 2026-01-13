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

// UpdateListingRevisionPackageDetails The model for an Oracle Cloud Infrastructure Marketplace Listing revison package.
type UpdateListingRevisionPackageDetails struct {

	// The version for the package.
	PackageVersion *string `mandatory:"false" json:"packageVersion"`

	// The name for the listing revision package.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// The description for this package.
	Description *string `mandatory:"false" json:"description"`

	// The unique identifier for the artifact.
	ArtifactId *string `mandatory:"false" json:"artifactId"`

	// The unique term identifier.
	TermId *string `mandatory:"false" json:"termId"`

	// Identifies that this will be default package for the listing revision.
	IsDefault *bool `mandatory:"false" json:"isDefault"`

	// Identifies whether security upgrades will be provided for this package.
	AreSecurityUpgradesProvided *bool `mandatory:"false" json:"areSecurityUpgradesProvided"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m UpdateListingRevisionPackageDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpdateListingRevisionPackageDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
