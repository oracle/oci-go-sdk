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

// CreateMachineImageDetails Machine image details required to create an Image artifact.
type CreateMachineImageDetails struct {

	// The OCID of source Machine Image.
	SourceImageId *string `mandatory:"true" json:"sourceImageId"`

	// Indicates if the customer is allowed to take a snapshot of running instance created from the machine image.
	IsSnapshotAllowed *bool `mandatory:"true" json:"isSnapshotAllowed"`

	// List of shape configurations compatible with the image.
	ImageShapeCompatibilityEntries []ImageShapeCompatibility `mandatory:"true" json:"imageShapeCompatibilityEntries"`

	// The username for machine image.
	Username *string `mandatory:"false" json:"username"`
}

func (m CreateMachineImageDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateMachineImageDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
