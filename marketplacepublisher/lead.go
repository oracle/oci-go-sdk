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

// Lead Details of Lead
type Lead struct {

	// Unique OCID identifier for the lead.
	Id *string `mandatory:"true" json:"id"`

	// The OCID for the listing.
	ListingId *string `mandatory:"true" json:"listingId"`

	// The Name for the listing.
	ListingName *string `mandatory:"true" json:"listingName"`

	// The OCID for the compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The state for the listing.
	LifecycleState LeadLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The date and time the listing was created, in the format defined by RFC3339 (https://tools.ietf.org/html/rfc3339).
	// Example: `2023-03-27T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	AdditionalInfo *AdditionalInfo `mandatory:"false" json:"additionalInfo"`
}

func (m Lead) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m Lead) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingLeadLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetLeadLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// LeadLifecycleStateEnum Enum with underlying type: string
type LeadLifecycleStateEnum string

// Set of constants representing the allowable values for LeadLifecycleStateEnum
const (
	LeadLifecycleStateActive  LeadLifecycleStateEnum = "ACTIVE"
	LeadLifecycleStateDeleted LeadLifecycleStateEnum = "DELETED"
)

var mappingLeadLifecycleStateEnum = map[string]LeadLifecycleStateEnum{
	"ACTIVE":  LeadLifecycleStateActive,
	"DELETED": LeadLifecycleStateDeleted,
}

var mappingLeadLifecycleStateEnumLowerCase = map[string]LeadLifecycleStateEnum{
	"active":  LeadLifecycleStateActive,
	"deleted": LeadLifecycleStateDeleted,
}

// GetLeadLifecycleStateEnumValues Enumerates the set of values for LeadLifecycleStateEnum
func GetLeadLifecycleStateEnumValues() []LeadLifecycleStateEnum {
	values := make([]LeadLifecycleStateEnum, 0)
	for _, v := range mappingLeadLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetLeadLifecycleStateEnumStringValues Enumerates the set of values in String for LeadLifecycleStateEnum
func GetLeadLifecycleStateEnumStringValues() []string {
	return []string{
		"ACTIVE",
		"DELETED",
	}
}

// GetMappingLeadLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingLeadLifecycleStateEnum(val string) (LeadLifecycleStateEnum, bool) {
	enum, ok := mappingLeadLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
