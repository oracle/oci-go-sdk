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

// TermVersion Model object for the term version details.
type TermVersion struct {

	// The unique identifier for the term.
	TermId *string `mandatory:"true" json:"termId"`

	// Who authored the term. Publisher terms will be defaulted to 'PARTNER'.
	TermAuthor TermAuthorEnum `mandatory:"true" json:"termAuthor"`

	// The name for the term version.
	DisplayName *string `mandatory:"true" json:"displayName"`

	Attachment *TermVersionAttachment `mandatory:"true" json:"attachment"`

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

	// The unique identifier for the compartment.
	CompartmentId *string `mandatory:"false" json:"compartmentId"`

	// Who authored the term. Publisher terms will be defaulted to 'PARTNER'.
	Author TermAuthorEnum `mandatory:"false" json:"author,omitempty"`

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

func (m TermVersion) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m TermVersion) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingTermAuthorEnum(string(m.TermAuthor)); !ok && m.TermAuthor != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for TermAuthor: %s. Supported values are: %s.", m.TermAuthor, strings.Join(GetTermAuthorEnumStringValues(), ",")))
	}
	if _, ok := GetMappingTermVersionStatusEnum(string(m.Status)); !ok && m.Status != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", m.Status, strings.Join(GetTermVersionStatusEnumStringValues(), ",")))
	}
	if _, ok := GetMappingTermVersionLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetTermVersionLifecycleStateEnumStringValues(), ",")))
	}

	if _, ok := GetMappingTermAuthorEnum(string(m.Author)); !ok && m.Author != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Author: %s. Supported values are: %s.", m.Author, strings.Join(GetTermAuthorEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// TermVersionStatusEnum Enum with underlying type: string
type TermVersionStatusEnum string

// Set of constants representing the allowable values for TermVersionStatusEnum
const (
	TermVersionStatusAvailable    TermVersionStatusEnum = "AVAILABLE"
	TermVersionStatusNotAvailable TermVersionStatusEnum = "NOT_AVAILABLE"
	TermVersionStatusDeleted      TermVersionStatusEnum = "DELETED"
)

var mappingTermVersionStatusEnum = map[string]TermVersionStatusEnum{
	"AVAILABLE":     TermVersionStatusAvailable,
	"NOT_AVAILABLE": TermVersionStatusNotAvailable,
	"DELETED":       TermVersionStatusDeleted,
}

var mappingTermVersionStatusEnumLowerCase = map[string]TermVersionStatusEnum{
	"available":     TermVersionStatusAvailable,
	"not_available": TermVersionStatusNotAvailable,
	"deleted":       TermVersionStatusDeleted,
}

// GetTermVersionStatusEnumValues Enumerates the set of values for TermVersionStatusEnum
func GetTermVersionStatusEnumValues() []TermVersionStatusEnum {
	values := make([]TermVersionStatusEnum, 0)
	for _, v := range mappingTermVersionStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetTermVersionStatusEnumStringValues Enumerates the set of values in String for TermVersionStatusEnum
func GetTermVersionStatusEnumStringValues() []string {
	return []string{
		"AVAILABLE",
		"NOT_AVAILABLE",
		"DELETED",
	}
}

// GetMappingTermVersionStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingTermVersionStatusEnum(val string) (TermVersionStatusEnum, bool) {
	enum, ok := mappingTermVersionStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// TermVersionLifecycleStateEnum Enum with underlying type: string
type TermVersionLifecycleStateEnum string

// Set of constants representing the allowable values for TermVersionLifecycleStateEnum
const (
	TermVersionLifecycleStateActive   TermVersionLifecycleStateEnum = "ACTIVE"
	TermVersionLifecycleStateInactive TermVersionLifecycleStateEnum = "INACTIVE"
)

var mappingTermVersionLifecycleStateEnum = map[string]TermVersionLifecycleStateEnum{
	"ACTIVE":   TermVersionLifecycleStateActive,
	"INACTIVE": TermVersionLifecycleStateInactive,
}

var mappingTermVersionLifecycleStateEnumLowerCase = map[string]TermVersionLifecycleStateEnum{
	"active":   TermVersionLifecycleStateActive,
	"inactive": TermVersionLifecycleStateInactive,
}

// GetTermVersionLifecycleStateEnumValues Enumerates the set of values for TermVersionLifecycleStateEnum
func GetTermVersionLifecycleStateEnumValues() []TermVersionLifecycleStateEnum {
	values := make([]TermVersionLifecycleStateEnum, 0)
	for _, v := range mappingTermVersionLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetTermVersionLifecycleStateEnumStringValues Enumerates the set of values in String for TermVersionLifecycleStateEnum
func GetTermVersionLifecycleStateEnumStringValues() []string {
	return []string{
		"ACTIVE",
		"INACTIVE",
	}
}

// GetMappingTermVersionLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingTermVersionLifecycleStateEnum(val string) (TermVersionLifecycleStateEnum, bool) {
	enum, ok := mappingTermVersionLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
