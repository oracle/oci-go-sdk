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

// MachineImageDetails Machine image details.
type MachineImageDetails struct {

	// The source image ocid of the machine image.
	SourceImageId *string `mandatory:"true" json:"sourceImageId"`

	// Identified if customer can take a snapshot of instance running on the machine image.
	IsSnapshotAllowed *bool `mandatory:"true" json:"isSnapshotAllowed"`

	// List of shape configurations supported by the image.
	ImageShapeCompatibilityEntries []ImageShapeCompatibility `mandatory:"true" json:"imageShapeCompatibilityEntries"`

	// Image validation status.
	ValidationStatus ValidationStatusEnum `mandatory:"true" json:"validationStatus"`

	// The username for machine image.
	Username *string `mandatory:"false" json:"username"`

	// Description of image validation errors.
	ValidationError *string `mandatory:"false" json:"validationError"`
}

func (m MachineImageDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m MachineImageDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingValidationStatusEnum(string(m.ValidationStatus)); !ok && m.ValidationStatus != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ValidationStatus: %s. Supported values are: %s.", m.ValidationStatus, strings.Join(GetValidationStatusEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
