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

// Product The model for the product details.
type Product struct {

	// The name for the product.
	Name *string `mandatory:"true" json:"name"`

	// The code for the product.
	Code *string `mandatory:"true" json:"code"`

	// The product group for the product.
	ProductGroup *string `mandatory:"true" json:"productGroup"`

	// The current state for the product.
	LifecycleState ProductLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The date and time the product was created, in the format defined by RFC3339 (https://tools.ietf.org/html/rfc3339).
	// Example: `2022-09-15T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The date and time the product was updated, in the format defined by RFC3339 (https://tools.ietf.org/html/rfc3339).
	// Example: `2022-09-15T21:10:29.600Z`
	TimeUpdated *common.SDKTime `mandatory:"true" json:"timeUpdated"`

	// Additional filter properties for product
	AdditionalFilters []AdditionalFilter `mandatory:"false" json:"additionalFilters"`
}

func (m Product) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m Product) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingProductLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetProductLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ProductLifecycleStateEnum Enum with underlying type: string
type ProductLifecycleStateEnum string

// Set of constants representing the allowable values for ProductLifecycleStateEnum
const (
	ProductLifecycleStateActive   ProductLifecycleStateEnum = "ACTIVE"
	ProductLifecycleStateInactive ProductLifecycleStateEnum = "INACTIVE"
)

var mappingProductLifecycleStateEnum = map[string]ProductLifecycleStateEnum{
	"ACTIVE":   ProductLifecycleStateActive,
	"INACTIVE": ProductLifecycleStateInactive,
}

var mappingProductLifecycleStateEnumLowerCase = map[string]ProductLifecycleStateEnum{
	"active":   ProductLifecycleStateActive,
	"inactive": ProductLifecycleStateInactive,
}

// GetProductLifecycleStateEnumValues Enumerates the set of values for ProductLifecycleStateEnum
func GetProductLifecycleStateEnumValues() []ProductLifecycleStateEnum {
	values := make([]ProductLifecycleStateEnum, 0)
	for _, v := range mappingProductLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetProductLifecycleStateEnumStringValues Enumerates the set of values in String for ProductLifecycleStateEnum
func GetProductLifecycleStateEnumStringValues() []string {
	return []string{
		"ACTIVE",
		"INACTIVE",
	}
}

// GetMappingProductLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingProductLifecycleStateEnum(val string) (ProductLifecycleStateEnum, bool) {
	enum, ok := mappingProductLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
