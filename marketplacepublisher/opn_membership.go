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

// OpnMembership OPN membership information
type OpnMembership struct {

	// OPN membership start date. An RFC3339 formatted datetime string
	TimeStart *common.SDKTime `mandatory:"false" json:"timeStart"`

	// OPN membership end date. An RFC3339 formatted datetime string
	TimeEnd *common.SDKTime `mandatory:"false" json:"timeEnd"`

	// OPN status
	OpnStatus OpnMembershipOpnStatusEnum `mandatory:"false" json:"opnStatus,omitempty"`

	// OPN Number number
	OpnNumber *string `mandatory:"false" json:"opnNumber"`

	// OPN membership type
	OpnMembershipType *string `mandatory:"false" json:"opnMembershipType"`
}

func (m OpnMembership) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m OpnMembership) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingOpnMembershipOpnStatusEnum(string(m.OpnStatus)); !ok && m.OpnStatus != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for OpnStatus: %s. Supported values are: %s.", m.OpnStatus, strings.Join(GetOpnMembershipOpnStatusEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// OpnMembershipOpnStatusEnum Enum with underlying type: string
type OpnMembershipOpnStatusEnum string

// Set of constants representing the allowable values for OpnMembershipOpnStatusEnum
const (
	OpnMembershipOpnStatusActive            OpnMembershipOpnStatusEnum = "ACTIVE"
	OpnMembershipOpnStatusInactive          OpnMembershipOpnStatusEnum = "INACTIVE"
	OpnMembershipOpnStatusRenewalInProgress OpnMembershipOpnStatusEnum = "RENEWAL_IN_PROGRESS"
)

var mappingOpnMembershipOpnStatusEnum = map[string]OpnMembershipOpnStatusEnum{
	"ACTIVE":              OpnMembershipOpnStatusActive,
	"INACTIVE":            OpnMembershipOpnStatusInactive,
	"RENEWAL_IN_PROGRESS": OpnMembershipOpnStatusRenewalInProgress,
}

var mappingOpnMembershipOpnStatusEnumLowerCase = map[string]OpnMembershipOpnStatusEnum{
	"active":              OpnMembershipOpnStatusActive,
	"inactive":            OpnMembershipOpnStatusInactive,
	"renewal_in_progress": OpnMembershipOpnStatusRenewalInProgress,
}

// GetOpnMembershipOpnStatusEnumValues Enumerates the set of values for OpnMembershipOpnStatusEnum
func GetOpnMembershipOpnStatusEnumValues() []OpnMembershipOpnStatusEnum {
	values := make([]OpnMembershipOpnStatusEnum, 0)
	for _, v := range mappingOpnMembershipOpnStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetOpnMembershipOpnStatusEnumStringValues Enumerates the set of values in String for OpnMembershipOpnStatusEnum
func GetOpnMembershipOpnStatusEnumStringValues() []string {
	return []string{
		"ACTIVE",
		"INACTIVE",
		"RENEWAL_IN_PROGRESS",
	}
}

// GetMappingOpnMembershipOpnStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingOpnMembershipOpnStatusEnum(val string) (OpnMembershipOpnStatusEnum, bool) {
	enum, ok := mappingOpnMembershipOpnStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
