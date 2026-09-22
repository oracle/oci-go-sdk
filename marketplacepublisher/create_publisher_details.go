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

// CreatePublisherDetails Details to create a publisher.
type CreatePublisherDetails struct {

	// Email address of the publisher.
	Email *string `mandatory:"true" json:"email"`

	// The business phone number of the publisher.
	BusinessPhoneNumber *string `mandatory:"true" json:"businessPhoneNumber"`

	// The company name of the publisher.
	CompanyName *string `mandatory:"true" json:"companyName"`

	// The company name of the publisher.
	CompanyDescription *string `mandatory:"true" json:"companyDescription"`

	// Count of employees in publisher's company
	EmployeeCount *int64 `mandatory:"true" json:"employeeCount"`

	// OPN membership number of the publisher
	OpnNumber *string `mandatory:"true" json:"opnNumber"`

	// The year the publisher's company or organization was founded.
	YearFounded *int64 `mandatory:"true" json:"yearFounded"`

	// City
	City *string `mandatory:"true" json:"city"`

	// State
	State *string `mandatory:"true" json:"state"`

	// Country
	Country *string `mandatory:"true" json:"country"`

	// The publisher's website.
	Website *string `mandatory:"true" json:"website"`

	// The contact phone number of the publisher.
	ContactPhoneNumber *string `mandatory:"true" json:"contactPhoneNumber"`

	// The contact email address of the publisher.
	ContactEmail *string `mandatory:"true" json:"contactEmail"`

	// The tenancy(compartment) OCID of the publisher.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// A description of the publisher solutions.
	SolutionDescription *string `mandatory:"false" json:"solutionDescription"`
}

func (m CreatePublisherDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreatePublisherDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
