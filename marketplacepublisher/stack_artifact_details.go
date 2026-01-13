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

// StackArtifactDetails Stack details required to create a Stack artifact.
type StackArtifactDetails struct {

	// The source stack OCID.
	SourceStackId *string `mandatory:"true" json:"sourceStackId"`

	// Stack validation status.
	ValidationStatus ValidationStatusEnum `mandatory:"true" json:"validationStatus"`

	// Image listing OCIDs that are referred in the Stack.
	ImageListingIds []string `mandatory:"false" json:"imageListingIds"`

	// The description of the stack validation failure errors.
	ValidationError *string `mandatory:"false" json:"validationError"`
}

func (m StackArtifactDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m StackArtifactDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingValidationStatusEnum(string(m.ValidationStatus)); !ok && m.ValidationStatus != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ValidationStatus: %s. Supported values are: %s.", m.ValidationStatus, strings.Join(GetValidationStatusEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
