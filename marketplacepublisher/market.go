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

// Market The model for the market details.
type Market struct {

	// The name of the market.
	Name *string `mandatory:"true" json:"name"`

	// The code of the market.
	Code *string `mandatory:"true" json:"code"`

	// The category code of the market.
	CategoryCode *string `mandatory:"true" json:"categoryCode"`

	// bill to countries for the market.
	BillToCountries []string `mandatory:"true" json:"billToCountries"`

	// The current state for the market.
	LifecycleState MarketLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The date and time the market was created, in the format defined by RFC3339 (https://tools.ietf.org/html/rfc3339).
	// Example: `2022-09-15T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The date and time the market was updated, in the format defined by RFC3339 (https://tools.ietf.org/html/rfc3339).
	// Example: `2022-09-15T21:10:29.600Z`
	TimeUpdated *common.SDKTime `mandatory:"true" json:"timeUpdated"`

	// The realm code of the market.
	RealmCode *string `mandatory:"false" json:"realmCode"`
}

func (m Market) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m Market) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingMarketLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetMarketLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarketLifecycleStateEnum Enum with underlying type: string
type MarketLifecycleStateEnum string

// Set of constants representing the allowable values for MarketLifecycleStateEnum
const (
	MarketLifecycleStateActive   MarketLifecycleStateEnum = "ACTIVE"
	MarketLifecycleStateInactive MarketLifecycleStateEnum = "INACTIVE"
)

var mappingMarketLifecycleStateEnum = map[string]MarketLifecycleStateEnum{
	"ACTIVE":   MarketLifecycleStateActive,
	"INACTIVE": MarketLifecycleStateInactive,
}

var mappingMarketLifecycleStateEnumLowerCase = map[string]MarketLifecycleStateEnum{
	"active":   MarketLifecycleStateActive,
	"inactive": MarketLifecycleStateInactive,
}

// GetMarketLifecycleStateEnumValues Enumerates the set of values for MarketLifecycleStateEnum
func GetMarketLifecycleStateEnumValues() []MarketLifecycleStateEnum {
	values := make([]MarketLifecycleStateEnum, 0)
	for _, v := range mappingMarketLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetMarketLifecycleStateEnumStringValues Enumerates the set of values in String for MarketLifecycleStateEnum
func GetMarketLifecycleStateEnumStringValues() []string {
	return []string{
		"ACTIVE",
		"INACTIVE",
	}
}

// GetMappingMarketLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingMarketLifecycleStateEnum(val string) (MarketLifecycleStateEnum, bool) {
	enum, ok := mappingMarketLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
