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

// AdminTermVersion Model object for the term version details.
type AdminTermVersion struct {

	// The unique identifier for the term.
	TermId *string `mandatory:"true" json:"termId"`

	// The name for the term version.
	DisplayName *string `mandatory:"true" json:"displayName"`

	Attachment *AdminTermVersionAttachment `mandatory:"true" json:"attachment"`

	// The current status for the term version.
	Status TermVersionStatusEnum `mandatory:"true" json:"status"`

	// The current state for the term version.
	LifecycleState TermVersionLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The date and time the term version was created, in the format defined by RFC3339 (https://tools.ietf.org/html/rfc3339).
	// Example: `2022-09-15T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The date and time the term version was updated, in the format defined by RFC3339 (https://tools.ietf.org/html/rfc3339).
	// Example: `2022-09-15T21:10:29.600Z`
	TimeUpdated *common.SDKTime `mandatory:"true" json:"timeUpdated"`

	// Unique OCID identifier for the term version.
	Id *string `mandatory:"false" json:"id"`

	// The unique legacy identifier for the term version.
	LegacyId *string `mandatory:"false" json:"legacyId"`

	// Who authored the term. Publisher terms will be defaulted to 'PARTNER'.
	TermAuthor AdminTermVersionTermAuthorEnum `mandatory:"false" json:"termAuthor,omitempty"`

	// The unique identifier for the compartment.
	CompartmentId *string `mandatory:"false" json:"compartmentId"`

	// Who authored the term. Publisher terms will be defaulted to 'PARTNER'.
	Author AdminTermVersionAuthorEnum `mandatory:"false" json:"author,omitempty"`

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

func (m AdminTermVersion) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m AdminTermVersion) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingTermVersionStatusEnum(string(m.Status)); !ok && m.Status != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", m.Status, strings.Join(GetTermVersionStatusEnumStringValues(), ",")))
	}
	if _, ok := GetMappingTermVersionLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetTermVersionLifecycleStateEnumStringValues(), ",")))
	}

	if _, ok := GetMappingAdminTermVersionTermAuthorEnum(string(m.TermAuthor)); !ok && m.TermAuthor != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for TermAuthor: %s. Supported values are: %s.", m.TermAuthor, strings.Join(GetAdminTermVersionTermAuthorEnumStringValues(), ",")))
	}
	if _, ok := GetMappingAdminTermVersionAuthorEnum(string(m.Author)); !ok && m.Author != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Author: %s. Supported values are: %s.", m.Author, strings.Join(GetAdminTermVersionAuthorEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// AdminTermVersionTermAuthorEnum Enum with underlying type: string
type AdminTermVersionTermAuthorEnum string

// Set of constants representing the allowable values for AdminTermVersionTermAuthorEnum
const (
	AdminTermVersionTermAuthorOracle  AdminTermVersionTermAuthorEnum = "ORACLE"
	AdminTermVersionTermAuthorPartner AdminTermVersionTermAuthorEnum = "PARTNER"
)

var mappingAdminTermVersionTermAuthorEnum = map[string]AdminTermVersionTermAuthorEnum{
	"ORACLE":  AdminTermVersionTermAuthorOracle,
	"PARTNER": AdminTermVersionTermAuthorPartner,
}

var mappingAdminTermVersionTermAuthorEnumLowerCase = map[string]AdminTermVersionTermAuthorEnum{
	"oracle":  AdminTermVersionTermAuthorOracle,
	"partner": AdminTermVersionTermAuthorPartner,
}

// GetAdminTermVersionTermAuthorEnumValues Enumerates the set of values for AdminTermVersionTermAuthorEnum
func GetAdminTermVersionTermAuthorEnumValues() []AdminTermVersionTermAuthorEnum {
	values := make([]AdminTermVersionTermAuthorEnum, 0)
	for _, v := range mappingAdminTermVersionTermAuthorEnum {
		values = append(values, v)
	}
	return values
}

// GetAdminTermVersionTermAuthorEnumStringValues Enumerates the set of values in String for AdminTermVersionTermAuthorEnum
func GetAdminTermVersionTermAuthorEnumStringValues() []string {
	return []string{
		"ORACLE",
		"PARTNER",
	}
}

// GetMappingAdminTermVersionTermAuthorEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAdminTermVersionTermAuthorEnum(val string) (AdminTermVersionTermAuthorEnum, bool) {
	enum, ok := mappingAdminTermVersionTermAuthorEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// AdminTermVersionAuthorEnum Enum with underlying type: string
type AdminTermVersionAuthorEnum string

// Set of constants representing the allowable values for AdminTermVersionAuthorEnum
const (
	AdminTermVersionAuthorOracle  AdminTermVersionAuthorEnum = "ORACLE"
	AdminTermVersionAuthorPartner AdminTermVersionAuthorEnum = "PARTNER"
)

var mappingAdminTermVersionAuthorEnum = map[string]AdminTermVersionAuthorEnum{
	"ORACLE":  AdminTermVersionAuthorOracle,
	"PARTNER": AdminTermVersionAuthorPartner,
}

var mappingAdminTermVersionAuthorEnumLowerCase = map[string]AdminTermVersionAuthorEnum{
	"oracle":  AdminTermVersionAuthorOracle,
	"partner": AdminTermVersionAuthorPartner,
}

// GetAdminTermVersionAuthorEnumValues Enumerates the set of values for AdminTermVersionAuthorEnum
func GetAdminTermVersionAuthorEnumValues() []AdminTermVersionAuthorEnum {
	values := make([]AdminTermVersionAuthorEnum, 0)
	for _, v := range mappingAdminTermVersionAuthorEnum {
		values = append(values, v)
	}
	return values
}

// GetAdminTermVersionAuthorEnumStringValues Enumerates the set of values in String for AdminTermVersionAuthorEnum
func GetAdminTermVersionAuthorEnumStringValues() []string {
	return []string{
		"ORACLE",
		"PARTNER",
	}
}

// GetMappingAdminTermVersionAuthorEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAdminTermVersionAuthorEnum(val string) (AdminTermVersionAuthorEnum, bool) {
	enum, ok := mappingAdminTermVersionAuthorEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
