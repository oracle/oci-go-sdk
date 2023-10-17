// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
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

// Term Base model object for the term.
type Term struct {

	// The name for the term.
	Name *string `mandatory:"true" json:"name"`

	// Who authored the term. Publisher terms will be defaulted to 'PARTNER'.
	Author TermAuthorEnum `mandatory:"true" json:"author"`

	// The unique identifier for the compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The unique identifier for the publisher.
	PublisherId *string `mandatory:"true" json:"publisherId"`

	// The current state for the Term.
	LifecycleState TermLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The date and time the term was created, in the format defined by RFC3339 (https://tools.ietf.org/html/rfc3339).
	// Example: `2022-09-15T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The date and time the term was updated, in the format defined by RFC3339 (https://tools.ietf.org/html/rfc3339).
	// Example: `2022-09-15T21:10:29.600Z`
	TimeUpdated *common.SDKTime `mandatory:"true" json:"timeUpdated"`

	// Unique OCID identifier for the term.
	Id *string `mandatory:"false" json:"id"`

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

func (m Term) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m Term) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingTermAuthorEnum(string(m.Author)); !ok && m.Author != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Author: %s. Supported values are: %s.", m.Author, strings.Join(GetTermAuthorEnumStringValues(), ",")))
	}
	if _, ok := GetMappingTermLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetTermLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// TermAuthorEnum Enum with underlying type: string
type TermAuthorEnum string

// Set of constants representing the allowable values for TermAuthorEnum
const (
	TermAuthorOracle  TermAuthorEnum = "ORACLE"
	TermAuthorPartner TermAuthorEnum = "PARTNER"
)

var mappingTermAuthorEnum = map[string]TermAuthorEnum{
	"ORACLE":  TermAuthorOracle,
	"PARTNER": TermAuthorPartner,
}

var mappingTermAuthorEnumLowerCase = map[string]TermAuthorEnum{
	"oracle":  TermAuthorOracle,
	"partner": TermAuthorPartner,
}

// GetTermAuthorEnumValues Enumerates the set of values for TermAuthorEnum
func GetTermAuthorEnumValues() []TermAuthorEnum {
	values := make([]TermAuthorEnum, 0)
	for _, v := range mappingTermAuthorEnum {
		values = append(values, v)
	}
	return values
}

// GetTermAuthorEnumStringValues Enumerates the set of values in String for TermAuthorEnum
func GetTermAuthorEnumStringValues() []string {
	return []string{
		"ORACLE",
		"PARTNER",
	}
}

// GetMappingTermAuthorEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingTermAuthorEnum(val string) (TermAuthorEnum, bool) {
	enum, ok := mappingTermAuthorEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// TermLifecycleStateEnum Enum with underlying type: string
type TermLifecycleStateEnum string

// Set of constants representing the allowable values for TermLifecycleStateEnum
const (
	TermLifecycleStateActive   TermLifecycleStateEnum = "ACTIVE"
	TermLifecycleStateInactive TermLifecycleStateEnum = "INACTIVE"
)

var mappingTermLifecycleStateEnum = map[string]TermLifecycleStateEnum{
	"ACTIVE":   TermLifecycleStateActive,
	"INACTIVE": TermLifecycleStateInactive,
}

var mappingTermLifecycleStateEnumLowerCase = map[string]TermLifecycleStateEnum{
	"active":   TermLifecycleStateActive,
	"inactive": TermLifecycleStateInactive,
}

// GetTermLifecycleStateEnumValues Enumerates the set of values for TermLifecycleStateEnum
func GetTermLifecycleStateEnumValues() []TermLifecycleStateEnum {
	values := make([]TermLifecycleStateEnum, 0)
	for _, v := range mappingTermLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetTermLifecycleStateEnumStringValues Enumerates the set of values in String for TermLifecycleStateEnum
func GetTermLifecycleStateEnumStringValues() []string {
	return []string{
		"ACTIVE",
		"INACTIVE",
	}
}

// GetMappingTermLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingTermLifecycleStateEnum(val string) (TermLifecycleStateEnum, bool) {
	enum, ok := mappingTermLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
