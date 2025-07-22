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

// UpdateMachineImageDetails Machine image details required to update a Image artifact.
type UpdateMachineImageDetails struct {

	// The OCID of source machine image.
	SourceImageId *string `mandatory:"false" json:"sourceImageId"`

	// The username for machine image.
	Username *string `mandatory:"false" json:"username"`

	// Identified if customer can take a snapshot of instance running on the machine image.
	IsSnapshotAllowed *bool `mandatory:"false" json:"isSnapshotAllowed"`

	// List of shape configurations supported by the image.
	ImageShapeCompatibilityEntries []ImageShapeCompatibility `mandatory:"false" json:"imageShapeCompatibilityEntries"`
}

func (m UpdateMachineImageDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpdateMachineImageDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
