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

// AdditionalFilter Additional filters for the product
type AdditionalFilter struct {

	// Filter name
	Name *string `mandatory:"false" json:"name"`

	// Filter code
	Code *string `mandatory:"false" json:"code"`

	// Usage instructions for the properties
	UsageInstructions *string `mandatory:"false" json:"usageInstructions"`

	// Is multiselect available for product code or not
	IsMultiSelect *bool `mandatory:"false" json:"isMultiSelect"`

	// Is the product code mandatory or not
	IsMandatory *bool `mandatory:"false" json:"isMandatory"`

	// Additional filters attached to custom filter
	Properties []FilterProperty `mandatory:"false" json:"properties"`
}

func (m AdditionalFilter) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m AdditionalFilter) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
