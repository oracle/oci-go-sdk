// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
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

// ContainerImageDetails Container Image details.
type ContainerImageDetails struct {

	// The source registry url of the container image.
	SourceRegistryUrl *string `mandatory:"true" json:"sourceRegistryUrl"`

	// image validation status
	ValidationStatus ValidationStatusEnum `mandatory:"true" json:"validationStatus"`

	// image publication status
	PublicationStatus PublicationStatusEnum `mandatory:"true" json:"publicationStatus"`

	// The source registry OCID of the container image.
	SourceRegistryId *string `mandatory:"false" json:"sourceRegistryId"`

	// image validation failure errors
	ValidationError *string `mandatory:"false" json:"validationError"`

	// image publication failure errors
	PublicationError *string `mandatory:"false" json:"publicationError"`
}

func (m ContainerImageDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ContainerImageDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingValidationStatusEnum(string(m.ValidationStatus)); !ok && m.ValidationStatus != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ValidationStatus: %s. Supported values are: %s.", m.ValidationStatus, strings.Join(GetValidationStatusEnumStringValues(), ",")))
	}
	if _, ok := GetMappingPublicationStatusEnum(string(m.PublicationStatus)); !ok && m.PublicationStatus != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for PublicationStatus: %s. Supported values are: %s.", m.PublicationStatus, strings.Join(GetPublicationStatusEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
