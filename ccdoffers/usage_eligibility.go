// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Continuous Consumption Discounts API
//
// Use the Continuous Consumption Discounts API to manage continuous consumption discount operations.
//

package ccdoffers

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// UsageEligibility The CCD commitment applicability restrictions. If the CCD offer has been created with any region- or AD-specific offer restriction,
// then it will be honored first over the usage eligibility criteria.
type UsageEligibility struct {

	// The compartment OCID the CCD commitment is applicable to.
	CompartmentIds []string `mandatory:"false" json:"compartmentIds"`

	// The region the CCD commitment is applicable to.
	// If the CCD offer has already been restricted to any specific region, then the offer will supersede this input.
	Region *string `mandatory:"false" json:"region"`

	// The Availability Domain the CCD is applicable to.
	// If you're setting the `availabilityDomain`, ensure to also set the `region` property value.
	// If the CCD offer has already been restricted to a specific Availability Domain, then the offer will supersede this input.
	AvailabilityDomain *string `mandatory:"false" json:"availabilityDomain"`

	// The tenancy IDs the CCD commitment is applicable to.
	TenancyIds []string `mandatory:"false" json:"tenancyIds"`
}

func (m UsageEligibility) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UsageEligibility) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
