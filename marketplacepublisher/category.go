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

// Category The model for the category details.
type Category struct {

	// The name for the category.
	Name *string `mandatory:"true" json:"name"`

	// The code of the category.
	Code *string `mandatory:"true" json:"code"`

	// The product that the category belongs to.
	ProductCode *string `mandatory:"true" json:"productCode"`

	// The current state for the category.
	LifecycleState CategoryLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The date and time the category was created, in the format defined by RFC3339 (https://tools.ietf.org/html/rfc3339).
	// Example: `2022-09-15T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The date and time the category was updated, in the format defined by RFC3339 (https://tools.ietf.org/html/rfc3339).
	// Example: `2022-09-15T21:10:29.600Z`
	TimeUpdated *common.SDKTime `mandatory:"true" json:"timeUpdated"`
}

func (m Category) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m Category) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingCategoryLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetCategoryLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// CategoryLifecycleStateEnum Enum with underlying type: string
type CategoryLifecycleStateEnum string

// Set of constants representing the allowable values for CategoryLifecycleStateEnum
const (
	CategoryLifecycleStateActive   CategoryLifecycleStateEnum = "ACTIVE"
	CategoryLifecycleStateInactive CategoryLifecycleStateEnum = "INACTIVE"
)

var mappingCategoryLifecycleStateEnum = map[string]CategoryLifecycleStateEnum{
	"ACTIVE":   CategoryLifecycleStateActive,
	"INACTIVE": CategoryLifecycleStateInactive,
}

var mappingCategoryLifecycleStateEnumLowerCase = map[string]CategoryLifecycleStateEnum{
	"active":   CategoryLifecycleStateActive,
	"inactive": CategoryLifecycleStateInactive,
}

// GetCategoryLifecycleStateEnumValues Enumerates the set of values for CategoryLifecycleStateEnum
func GetCategoryLifecycleStateEnumValues() []CategoryLifecycleStateEnum {
	values := make([]CategoryLifecycleStateEnum, 0)
	for _, v := range mappingCategoryLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetCategoryLifecycleStateEnumStringValues Enumerates the set of values in String for CategoryLifecycleStateEnum
func GetCategoryLifecycleStateEnumStringValues() []string {
	return []string{
		"ACTIVE",
		"INACTIVE",
	}
}

// GetMappingCategoryLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCategoryLifecycleStateEnum(val string) (CategoryLifecycleStateEnum, bool) {
	enum, ok := mappingCategoryLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
