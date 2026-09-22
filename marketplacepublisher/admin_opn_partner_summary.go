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

// AdminOpnPartnerSummary OPN partner summary.
type AdminOpnPartnerSummary struct {

	// The id for company.
	CompanyId *string `mandatory:"true" json:"companyId"`

	// Partner Name.
	PartnerName *string `mandatory:"true" json:"partnerName"`

	// Partner Name.
	LevelProgramName *string `mandatory:"true" json:"levelProgramName"`

	// The OPN membership status.
	MembershipStatus AdminOpnPartnerSummaryMembershipStatusEnum `mandatory:"true" json:"membershipStatus"`

	// OPN membership type
	MembershipType *string `mandatory:"true" json:"membershipType"`

	// The start date.
	StartDate *string `mandatory:"false" json:"startDate"`

	// The end date.
	EndDate *string `mandatory:"false" json:"endDate"`

	// The address of the partner.
	StreetAddress1 *string `mandatory:"false" json:"streetAddress1"`

	// The address of the partner.
	StreetAddress2 *string `mandatory:"false" json:"streetAddress2"`

	// City
	City *string `mandatory:"false" json:"city"`

	// State
	State *string `mandatory:"false" json:"state"`

	// Zip
	Zipcode *string `mandatory:"false" json:"zipcode"`

	// Country
	Country *string `mandatory:"false" json:"country"`

	// The phone number of the contact.
	PhoneNumber *string `mandatory:"false" json:"phoneNumber"`

	// The URL for partner's service.
	PartnerUrl *string `mandatory:"false" json:"partnerUrl"`

	// The total number of employees.
	NumberOfEmployees *string `mandatory:"false" json:"numberOfEmployees"`

	// The last updated timestamp in format "yyyy-MM-dd HH:mm:ss.S".
	LastUpdated *string `mandatory:"false" json:"lastUpdated"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// System tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
}

func (m AdminOpnPartnerSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m AdminOpnPartnerSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingAdminOpnPartnerSummaryMembershipStatusEnum(string(m.MembershipStatus)); !ok && m.MembershipStatus != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for MembershipStatus: %s. Supported values are: %s.", m.MembershipStatus, strings.Join(GetAdminOpnPartnerSummaryMembershipStatusEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// AdminOpnPartnerSummaryMembershipStatusEnum Enum with underlying type: string
type AdminOpnPartnerSummaryMembershipStatusEnum string

// Set of constants representing the allowable values for AdminOpnPartnerSummaryMembershipStatusEnum
const (
	AdminOpnPartnerSummaryMembershipStatusActive            AdminOpnPartnerSummaryMembershipStatusEnum = "ACTIVE"
	AdminOpnPartnerSummaryMembershipStatusInactive          AdminOpnPartnerSummaryMembershipStatusEnum = "INACTIVE"
	AdminOpnPartnerSummaryMembershipStatusRenewalInProgress AdminOpnPartnerSummaryMembershipStatusEnum = "RENEWAL_IN_PROGRESS"
)

var mappingAdminOpnPartnerSummaryMembershipStatusEnum = map[string]AdminOpnPartnerSummaryMembershipStatusEnum{
	"ACTIVE":              AdminOpnPartnerSummaryMembershipStatusActive,
	"INACTIVE":            AdminOpnPartnerSummaryMembershipStatusInactive,
	"RENEWAL_IN_PROGRESS": AdminOpnPartnerSummaryMembershipStatusRenewalInProgress,
}

var mappingAdminOpnPartnerSummaryMembershipStatusEnumLowerCase = map[string]AdminOpnPartnerSummaryMembershipStatusEnum{
	"active":              AdminOpnPartnerSummaryMembershipStatusActive,
	"inactive":            AdminOpnPartnerSummaryMembershipStatusInactive,
	"renewal_in_progress": AdminOpnPartnerSummaryMembershipStatusRenewalInProgress,
}

// GetAdminOpnPartnerSummaryMembershipStatusEnumValues Enumerates the set of values for AdminOpnPartnerSummaryMembershipStatusEnum
func GetAdminOpnPartnerSummaryMembershipStatusEnumValues() []AdminOpnPartnerSummaryMembershipStatusEnum {
	values := make([]AdminOpnPartnerSummaryMembershipStatusEnum, 0)
	for _, v := range mappingAdminOpnPartnerSummaryMembershipStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetAdminOpnPartnerSummaryMembershipStatusEnumStringValues Enumerates the set of values in String for AdminOpnPartnerSummaryMembershipStatusEnum
func GetAdminOpnPartnerSummaryMembershipStatusEnumStringValues() []string {
	return []string{
		"ACTIVE",
		"INACTIVE",
		"RENEWAL_IN_PROGRESS",
	}
}

// GetMappingAdminOpnPartnerSummaryMembershipStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAdminOpnPartnerSummaryMembershipStatusEnum(val string) (AdminOpnPartnerSummaryMembershipStatusEnum, bool) {
	enum, ok := mappingAdminOpnPartnerSummaryMembershipStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
